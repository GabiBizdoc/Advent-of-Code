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

type Config struct {
	OutputFile        string
	Lines             int
	Generators        int
	WriterChannelSize int
	Help              bool
}

func NewConfig() *Config {
	return &Config{Lines: 1000, Generators: 10, WriterChannelSize: 10, Help: false}
}

func parseArgs() (*Config, error) {
	args := NewConfig()

	flag.StringVar(&args.OutputFile, "output", "", "Output file. Skip for stdout: Example --output ./file.csv")
	flag.IntVar(&args.Lines, "lines", 0, "Number of lines. Must be bigger than 0: Example --size 1_000_000_000")
	flag.IntVar(&args.Generators, "generators", 0, "Number of goroutines used to generated data: --generators 10")
	flag.IntVar(&args.WriterChannelSize, "writer-channel-size", 0, "Number of chunks buffered. Must be bigger than 0: --writer-channel-size 10")
	flag.BoolVar(&args.Help, "h", false, "Help")

	flag.Parse()

	if args.Help {
		flag.Usage()
		os.Exit(0)
	}
	if args.Lines <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	var err error
	if args.OutputFile != "" {
		args.OutputFile, err = filepath.Abs(args.OutputFile)
		if err != nil {
			return nil, err
		}
	}

	if args.Generators <= 0 {
		const defaultValue = 10
		fmt.Printf("generators must be at least 1. Changing it to %d. it was %d\n",
			defaultValue, args.Generators)
		args.Generators = defaultValue
	}

	if args.Lines <= 0 {
		fmt.Println("lines must be at least 1")
		os.Exit(1)
	}

	if args.WriterChannelSize < 0 {
		const defaultValue = 5
		fmt.Printf("writer-channel-size must be at least 1. Changing it to %d. it was %d\n",
			defaultValue, args.WriterChannelSize)
		args.WriterChannelSize = defaultValue
	}

	return args, nil
}

func main() {
	start := time.Now()
	config, err := parseArgs()
	if err != nil {
		panic(err)
	}
	var file *os.File
	if config.OutputFile == "" {
		file = os.Stdout
	} else {
		file, err = os.OpenFile(config.OutputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
	}

	defer file.Close()

	generate(config, file)

	elapsed := time.Since(start)
	fmt.Printf("Created file with %d measurements in %s\n", config.Lines, elapsed)
}

func generate(config *Config, outputFile *os.File) {
	pb := progressbar.NewProgressBar(config.Lines)
	pb.Label = "Progress"
	fmt.Println(strings.Repeat("=", 50))

	chunkSize := min(100_000, 1+config.Lines/config.Generators)
	out := generateData(config.Lines, chunkSize, config.Generators, config.WriterChannelSize)
	afterChunk := WriteToFile(out, outputFile)

	for range afterChunk {
		pb.Update(chunkSize)
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
