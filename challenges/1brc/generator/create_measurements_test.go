package main

import (
	"os"
	"strings"
	"testing"
)

const chunkSize = 100_000_00

func Benchmark_appendRowV1(t *testing.B) {
	w := strings.Builder{}
	stations := GetHardcodedWeatherStations()
	t.ResetTimer()

	for i := 0; i < chunkSize; i++ {
		err := appendRowV1(w, stations)
		if err != nil {
			panic(err)
		}
	}
}
func Benchmark_appendRowV2(t *testing.B) {
	w := strings.Builder{}
	stations := GetHardcodedWeatherStations()
	t.ResetTimer()

	for i := 0; i < chunkSize; i++ {
		err := appendRowV2(w, stations)
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_appendRowV3(t *testing.B) {
	w := strings.Builder{}
	stations := GetHardcodedWeatherStations()
	t.ResetTimer()

	for i := 0; i < chunkSize; i++ {
		err := appendRowV3(w, stations)
		if err != nil {
			panic(err)
		}
	}
}

func Benchmark_appendRowV4(t *testing.B) {
	w := make([]byte, 0)
	stations := GetHardcodedWeatherStations()
	t.ResetTimer()

	for i := 0; i < chunkSize; i++ {
		w = appendRowV4(w, stations)
	}
}

func Benchmark_appendRowV4x2(t *testing.B) {
	w := make([]byte, 0, chunkSize*10)
	stations := GetHardcodedWeatherStations()
	t.ResetTimer()

	for i := 0; i < chunkSize; i++ {
		w = appendRowV4(w, stations)
	}
}

func Benchmark_generateAndWrite(t *testing.B) {
	const size = 1_000_000_000
	file, _ := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	t.ResetTimer()
	writeFile(size, file)
}
