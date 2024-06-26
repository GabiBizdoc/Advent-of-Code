package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

var weatherCSVFile = "challenges/1brc/dist/weather_stations.csv"

const ioBufferSize = 64 << 10

//const ioBufferSize = 1 << 20

var workerCnt = runtime.NumCPU() * 4

func readArgs() {
	var workers int
	var csvFile string
	var err error

	flag.IntVar(&workers, "workers", 0, "Number of goroutines, default is NumCPU")
	flag.StringVar(&csvFile, "file", "", "Input file: weather_data.csv")
	flag.Parse()

	if csvFile != "" {
		weatherCSVFile, err = filepath.Abs(csvFile)
		if err != nil {
			panic(err)
		}
	}
	if workers > 0 {
		workerCnt = workers
	}
}

type CustomWriter struct {
}

func (c CustomWriter) Close() error {
	return nil
}

func (c CustomWriter) Write(p []byte) (n int, err error) {
	//return os.Stdout.Write(p)
	return len(p), err
}

func main() {
	start := time.Now()
	readArgs()

	file, err := os.Open(weatherCSVFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	compute(file)
	fmt.Println(time.Since(start))
}
func compute(file *os.File) {
	readOutChan, pool := ReadFile(file)
	splitterOutChan := SplitIntoChunks(readOutChan, pool)
	parserOutChan := ParseRows(splitterOutChan, workerCnt)
	respMap := AggregateResponsesMap(parserOutChan)
	PrintResponseMap(respMap, CustomWriter{})

	//we don't need to sort the output
	//respSlice := AggregateResponsesSlice(respMap)
	//slices.SortFunc(respSlice, func(a, b *Stats) int {
	//	return cmp.Compare(a.Name, b.Name)
	//})
	//PrintResponseSlice(respSlice, CustomWriter{})
}

func PrintResponseMap(stats map[string]*Stats, w io.Writer) {
	var sb bytes.Buffer
	sb.WriteByte('{')
	separator := false
	for name, v := range stats {
		v.SetName(name)
		if separator {
			sb.WriteString(", ")
		}
		separator = true
		sb.WriteString(v.Fmt())
	}
	sb.WriteByte('}')
	w.Write(sb.Bytes())
}

func PrintResponseSlice(stats []*Stats, w io.Writer) {
	var sb bytes.Buffer
	sb.WriteByte('{')
	separator := false
	for _, v := range stats {
		if separator {
			sb.WriteString(", ")
		}
		separator = true
		sb.WriteString(v.Fmt())
	}
	sb.WriteByte('}')
	w.Write(sb.Bytes())
}

func ReadFileLineByLine(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_ = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func ReadFile(f *os.File) (out chan []byte, pool *sync.Pool) {
	pool = &sync.Pool{New: func() any {
		return make([]byte, 0, ioBufferSize)
	}}
	out = make(chan []byte, 1000)
	go readFile(f, pool, out)
	return out, pool
}

func readFile(f *os.File, pool *sync.Pool, out chan []byte) {
	defer func() {
		close(out)
	}()

	pushData := func(data []byte) {
		out <- data
	}
	for {
		data := pool.Get().([]byte)
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				if len(data) > 0 {
					pushData(data)
				}
				return
			}
			panic(err)
		}
		pushData(data)
	}
}

//var totalLines = 0

func ParseRows(in chan []byte, workerCnt int) chan map[string]*Stats {
	out := make(chan map[string]*Stats, 80)
	go func() {
		var wg sync.WaitGroup
		wg.Add(workerCnt)
		for i := 0; i < workerCnt; i++ {
			go func() {
				defer wg.Done()
				results := make(map[string]*Stats, 10_000)
				for data := range in {
					var sep, start int
					for i, c := range data {
						switch c {
						case ';':
							sep = i
						case '\n':
							// copying the string is −28.57% slower
							// name := string(data[start:sep])
							// value := string(data[sep+1 : i])

							nameSlice := data[start:sep]

							// the float becomes int
							valueSlice := data[sep+1 : i-1]
							valueSlice[len(valueSlice)-1] = data[i-1]

							name := unsafe.String(unsafe.SliceData(nameSlice), len(nameSlice))
							value := unsafe.String(unsafe.SliceData(valueSlice), len(valueSlice))
							start = i + 1

							temp1, err := strconv.Atoi(value)
							if err != nil {
								fmt.Println(string(data[sep:i]))
								panic(err)
							}
							stats, ok := results[name]
							temp := float64(temp1) / 10
							if !ok {
								results[name] = NewStats(temp)
							} else {
								stats.Update(temp)
							}
						}
					}
				}
				out <- results
			}()
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func AggregateResponsesMap(in chan map[string]*Stats) map[string]*Stats {
	final := make(map[string]*Stats, 10000)
	for results := range in {
		for k, v := range results {
			stats, ok := final[k]
			if !ok {
				final[k] = v
			} else {
				stats.Combine(v)
			}
		}
	}
	return final
}
func AggregateResponsesSlice(finalMap map[string]*Stats) []*Stats {
	respSlice := make([]*Stats, 0, len(finalMap))
	for k, v := range finalMap {
		v.SetName(k)
		respSlice = append(respSlice, v)
	}
	return respSlice
}

func SplitIntoChunks(in chan []byte, pool *sync.Pool) chan []byte {
	var last []byte
	out := make(chan []byte, 80)
	go func() {
		for data := range in {
			i := len(data) - 1
			for i >= 0 && data[i] != '\n' {
				i -= 1
			}
			i += 1
			out <- append(last, data[:i]...)
			last = slices.Clone(data[i:])
			pool.Put(data[:0])
		}
		if len(last) > 0 {
			if last[len(last)-1] != '\n' {
				last = append(last, '\n')
			}
			out <- last
		}
		close(out)
	}()

	return out
}
