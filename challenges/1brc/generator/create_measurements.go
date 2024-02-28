package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/GabiBizdoc/golang-playground/pkg/progressbar"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	config, err := parseArgs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", config)
	fmt.Println(PreviewArguments(config))
	if config.DryRun.Value {
		return
	}
	var file *os.File
	if config.OutputFile.Value == "" {
		file = os.Stdout
	} else {
		file, err = os.OpenFile(config.OutputFile.Value, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
	}

	defer file.Close()

	generate(config, file)

	elapsed := time.Since(start)
	fmt.Printf("Created file with %d measurements in %s\n", config.Lines.Value, elapsed)
}

func generate(config *Config, outputFile *os.File) {
	pb := progressbar.NewProgressBar(config.Lines.Value)
	pb.Label = "Progress"
	fmt.Println(strings.Repeat("=", 50))

	chunkSize := min(100_000, 1+config.Lines.Value/config.Generators.Value)
	out := generateData(config.Lines.Value, chunkSize, config.Generators.Value, config.WriterChannelSize.Value)
	afterChunk := WriteToFile(out, outputFile)

	for range afterChunk {
		pb.Update(min(chunkSize, config.Lines.Value-pb.Current))
	}
	pb.Done()

}
func appendRowFormattedString(w strings.Builder, station WeatherStation) (err error) {
	w.WriteString(fmt.Sprintf("%s;%.1f\n", station.ID, station.Measurement()))
	return err
}
func appendRowWithoutFormat(w strings.Builder, station WeatherStation) {
	w.WriteString(station.ID)
	w.WriteByte(';')
	measurement := strconv.FormatFloat(station.Measurement(), 'f', 1, 64)
	w.WriteString(measurement)
	w.WriteByte('\n')
}
func appendRowWithoutFormatScaled(w strings.Builder, station WeatherStation) {
	w.WriteString(station.ID)
	w.WriteByte(';')
	measurement := strconv.Itoa(int(station.Measurement() * 10))
	w.WriteString(measurement[0 : len(measurement)-1])
	w.WriteByte('.')
	w.WriteByte(measurement[len(measurement)-1])
	w.WriteByte('\n')
}

var customItoa = NewCustomItoa()

func appendRowToByteSlice(b []byte, station WeatherStation) []byte {
	b = append(b, station.ID...)
	b = append(b, ';')
	msm := int(station.Measurement() * 10)
	measurement := customItoa.Parse(msm)
	b = append(b, measurement[:len(measurement)-1]...)
	b = append(b, '.')
	b = append(b, measurement[len(measurement)-1])
	b = append(b, '\n')
	return b
}
func appendRowUsingBuffer(b *bytes.Buffer, station WeatherStation) {
	b.WriteString(station.ID)
	b.WriteByte(';')
	msm := int(station.Measurement() * 10)
	measurement := customItoa.Parse(msm)
	b.WriteString(measurement[:len(measurement)-1])
	b.WriteByte('.')
	b.WriteByte(measurement[len(measurement)-1])
	b.WriteByte('\n')
}

func generateData(size, chunkSize, numberOfWorkers, bufferSize int) chan []byte {
	weatherGenerator := NewWeatherStationsGenerator()
	in := make(chan []byte, bufferSize)
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	for n := 0; n < numberOfWorkers; n++ {
		batch := size / numberOfWorkers
		if n == 0 {
			batch = batch + size - (numberOfWorkers * batch)
		}
		go func(wg *sync.WaitGroup, batch int) {
			defer wg.Done()
			for progress := 0; progress < batch; progress += chunkSize {
				w := make([]byte, 0, chunkSize*10)

				for i := 0; i < min(batch-progress, chunkSize); i++ {
					w = appendRowToByteSlice(w, weatherGenerator.RandomStation())
				}
				in <- w
			}
		}(&wg, batch)
	}

	go func() {
		wg.Wait()
		close(in)
	}()

	return in
}

func WriteToFile(in chan []byte, file *os.File) chan struct{} {
	out := make(chan struct{})
	go writeToFile(in, out, file)
	return out
}

func writeToFile(in chan []byte, out chan struct{}, file *os.File) {
	w := bufio.NewWriter(file)
	for data := range in {
		_, err := w.Write(data)
		if err != nil {
			panic(err)
		}
		out <- struct{}{}
	}
	err := w.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
	close(out)
}
