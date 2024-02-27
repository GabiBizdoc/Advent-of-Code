package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type WeatherStation struct {
	ID              string
	MeanTemperature float64
}

func (w *WeatherStation) Measurement() float64 {
	m := rand.NormFloat64()*10 + w.MeanTemperature
	return math.Round(m*10) / 10
}

var customItoa = NewCustomItoa()

func main() {
	start := time.Now()

	var fileFlag string
	var size int
	flag.StringVar(&fileFlag, "file", "", "Output file")
	flag.IntVar(&size, "size", 0, "Number of lines")
	flag.Parse()
	//fileFlag = os.DevNull
	//size = 1_000_000_000

	if size <= 0 {
		fmt.Println("Usage: go create_measurements.go -size <number of records to create>")
		os.Exit(1)
	}
	fmt.Println("Using size: ", size)

	var file *os.File

	if fileFlag == "" {
		fmt.Println("Using file: ", "Stdout")
		file = os.Stdout
	} else {
		var err error
		fileFlag = os.DevNull
		file, err = os.OpenFile(fileFlag, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		fmt.Println("Using file: ", fileFlag)

		absPath, err := filepath.Abs(fileFlag)
		if err != nil {
			fmt.Println("Error getting absolute path:", err)
			return
		}

		fmt.Println("Absolute path of the file:", absPath)
		fmt.Println("Name of the file:", file.Name())

		defer func(file *os.File) {
			fmt.Println("closing file: ", fileFlag)

			err := file.Close()
			if err != nil {
				log.Println(err)
			}
		}(file)
	}
	fmt.Println("===========")

	writeFile(size, file)
	elapsed := time.Since(start)
	fmt.Printf("Created file with %d measurements in %s\n", size, elapsed)
}

func writeFile(size int, file *os.File) {
	stations := GetHardcodedWeatherStations()

	chunkSize := 1_000_000
	numberOfWorkers := size/chunkSize + 1

	//fmt.Println("numberOfWorkers:", numberOfWorkers)
	var wg sync.WaitGroup

	in := make(chan []byte, 80)

	var reader FileWriterHandler
	out := reader.Write(in, file)
	wg.Add(numberOfWorkers)

	for n := 0; n < numberOfWorkers; n++ {
		start := n * chunkSize
		end := start + chunkSize

		if end > size {
			end = size
		}
		go func(wg *sync.WaitGroup, start, end int) {
			defer wg.Done()
			w := make([]byte, 0, 5000)
			for i := start; i < end; i++ {
				w = appendRowV4(w, stations)
			}
			in <- w
		}(&wg, start, end)
	}

	wg.Wait()
	close(in)
	<-out
	//fmt.Println("RowCount", reader.RowCount)
}

func appendRowV1(w strings.Builder, stations []WeatherStation) (err error) {
	station := stations[rand.IntN(len(stations))]
	_, err = w.WriteString(fmt.Sprintf("%s;%.1f\n", station.ID, station.Measurement()))
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return err
	}
	return err
}
func appendRowV2(w strings.Builder, stations []WeatherStation) (err error) {
	station := stations[rand.IntN(len(stations))]

	_, err = w.WriteString(station.ID)
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	_, err = w.WriteRune(';')
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	measurement := strconv.FormatFloat(station.Measurement(), 'f', 1, 64)
	_, err = w.WriteString(measurement)
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	_, err = w.WriteRune('\n')
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	return
}
func appendRowV3(w strings.Builder, stations []WeatherStation) (err error) {
	station := stations[rand.IntN(len(stations))]

	_, err = w.WriteString(station.ID)
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	_, err = w.WriteRune(';')
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	measurement := strconv.Itoa(int(station.Measurement() * 10))
	_, err = w.WriteString(measurement[0 : len(measurement)-1])
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	_, err = w.WriteRune('.')
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	_, err = w.WriteRune(rune(measurement[len(measurement)-1]))
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	_, err = w.WriteRune('\n')
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	return
}
func appendRowV4(b []byte, stations []WeatherStation) []byte {
	station := stations[rand.IntN(len(stations))]
	b = append(b, station.ID...)
	b = append(b, ';')
	msm := int(station.Measurement() * 10)
	//measurement := strconv.Itoa(msm)
	measurement := customItoa.Parse(msm)

	b = append(b, measurement[:len(measurement)-1]...)
	b = append(b, '.')
	b = append(b, measurement[len(measurement)-1])
	b = append(b, '\n')
	return b
}

type FileWriterHandler struct {
	RowCount int
}

func (r *FileWriterHandler) Write(in chan []byte, file *os.File) chan struct{} {
	out := make(chan struct{})
	go r.write(in, out, file)
	return out
}

func (r *FileWriterHandler) write(in chan []byte, out chan struct{}, file *os.File) {
	for data := range in {
		r.RowCount += bytes.Count(data, []byte{'\n'})
		_, err := file.Write(data)
		if err != nil {
			panic(err)
		}
	}
	out <- struct{}{}
}
