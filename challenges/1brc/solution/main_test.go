package main

import (
	"os"
	"path/filepath"
	"strings"
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

func BenchmarkEverything(b *testing.B) {
	file, err := os.Open(getFileName())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	compute(file)
}

func Benchmark_ReadFileOs(b *testing.B) {
	b.Skip("not enough memory but buffer the whole file in memory")
	_, err := os.ReadFile(getFileName())
	if err != nil {
		panic(err)
	}
}

func Benchmark_ReadFileLineByLine(b *testing.B) {
	file, err := os.Open(getFileName())
	if err != nil {
		panic(err)
	}
	ReadFileLineByLine(file)
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

func cmpOutputs(t *testing.T, expected, result string) {
	expectedMap := parseOutputString(expected)
	ourMap := parseOutputString(result)

	t.Log(len(expectedMap), len(ourMap))
	if len(expectedMap) != len(ourMap) {
		t.Fail()
	}
	for k, v := range expectedMap {
		ourValue, ok := ourMap[k]
		if !ok {
			t.Logf("missing key: `%s`", k)
			t.Fail()
		} else if ourValue != v {
			t.Logf("key: `%s`: expected `%s` but got `%s`", k, v, ourValue)
			t.Fail()
		}
	}
}

func parseOutputString(s string) map[string]string {
	buf := s[1 : len(s)-1]
	entries := strings.Split(buf, ", ")
	m := make(map[string]string, len(entries))

	for _, entry := range entries {
		parts := strings.Split(entry, "=")
		m[parts[0]] = parts[1]
	}
	return m
}
