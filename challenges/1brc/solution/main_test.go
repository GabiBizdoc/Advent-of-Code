package main

import (
	"cmp"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"sync"
	"testing"
)

func getFileName() string {
	name, err := filepath.Abs(filepath.Join("../../../", weatherCSVFile))
	if err != nil {
		panic(err)
	}
	return name
}

func Benchmark_ReadFileLineByLine(b *testing.B) {
	file, err := os.Open(getFileName())
	if err != nil {
		panic(err)
	}
	ReadFileLineByLine(file)
}

func Benchmark_ReadFileOs(b *testing.B) {
	b.Skip("not enough memory but buffer the whole file in memory")
	_, err := os.ReadFile(getFileName())
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadFileIntoChan(b *testing.B) {
	runReadFile(ReadFile)
}

func TestReadFile2(t *testing.T) {
	runReadFile(ReadFile)
}

func runReadFile(cb func(file *os.File) (out chan []byte, pool *sync.Pool)) {
	file, err := os.Open(getFileName())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	out, pool := cb(file)

	for data := range out {
		pool.Put(data[:0])
	}
}

func BenchmarkReadAndSplit(b *testing.B) {
	file, err := os.Open(getFileName())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	readOut, pool := ReadFile(file)
	splitterOut := SplitIntoChunks(readOut, pool)
	parserOut := ParseRows(splitterOut, runtime.NumCPU())
	respSlice := AggregateResponses(parserOut)
	slices.SortFunc(respSlice, func(a, b *Stats) int {
		return cmp.Compare(a.Name, b.Name)
	})
	PrintResponse(respSlice, CustomWriter{})
}
