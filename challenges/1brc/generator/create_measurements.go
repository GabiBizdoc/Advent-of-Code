package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/GabiBizdoc/golang-playground/pkg/progressbar"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func readArgs() (int, *os.File) {
	var fileFlag string
	var size int
	var help bool
	flag.BoolVar(&help, "h", false, "Help")
	flag.StringVar(&fileFlag, "output", "", "Output file: Example -output ./file.csv")
	flag.IntVar(&size, "size", 0, "Number of lines: Example -size 1_000_000_000")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(fileFlag, size)
	if size <= 0 {
		flag.Usage()
		//fmt.Println("Usage: go create_measurements.go -size <number of records to create> -file <output file>")
		os.Exit(1)
	}
	fmt.Println("Using size: ", size)

	var file *os.File

	if fileFlag == "" {
		fmt.Println("Using file: ", "Stdout")
		file = os.Stdout
	} else {
		var err error

		fmt.Println("Using file: ", fileFlag)
		absPath, err := filepath.Abs(fileFlag)
		if err != nil {
			fmt.Println("Error getting absolute path:", err)
			return 0, nil
		}
		fmt.Println("Absolute path of the file:", absPath)

		file, err = os.OpenFile(fileFlag, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("Name of the file:", file.Name())
	}
	return size, file
}

func main() {
	start := time.Now()
	size, file := readArgs()
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()

	pb := progressbar.NewProgressBar(size)
	pb.Label = "Progress: "
	fmt.Println(strings.Repeat("==", 50))
	const chunkSize = 100_000
	out := generateData(size, chunkSize, 2, 10)
	afterChunk := WriteToFile(out, file)
	for range afterChunk {
		pb.Update(chunkSize)
	}
	pb.Done()

	elapsed := time.Since(start)
	fmt.Printf("Created file with %d measurements in %s\n", size, elapsed)
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
	fmt.Println("writing to file, ", file.Name())
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
