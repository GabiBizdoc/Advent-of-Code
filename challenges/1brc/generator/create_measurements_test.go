package main

import (
	"bytes"
	"strings"
	"testing"
)

const numberOfRows = 100_000_00

func Benchmark_appendRowFormattedString(t *testing.B) {
	w := strings.Builder{}
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < numberOfRows; i++ {
		err := appendRowFormattedString(w, weather.RandomStation())
		if err != nil {
			panic(err)
		}
	}
}
func Benchmark_appendRowWithoutFormat(t *testing.B) {
	w := strings.Builder{}
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < numberOfRows; i++ {
		appendRowWithoutFormat(w, weather.RandomStation())
	}
}

func Benchmark_appendRowWithoutFormatScaled(t *testing.B) {
	w := strings.Builder{}
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < numberOfRows; i++ {
		appendRowWithoutFormatScaled(w, weather.RandomStation())
	}
}

func Benchmark_appendRowToByteSlice(t *testing.B) {
	w := make([]byte, 0)
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < numberOfRows; i++ {
		w = appendRowToByteSlice(w, weather.RandomStation())
	}
}

func Benchmark_appendRowToByteSliceWithCapacity(t *testing.B) {
	w := make([]byte, 0, numberOfRows*10)
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < numberOfRows; i++ {
		w = appendRowToByteSlice(w, weather.RandomStation())
	}
}

func Benchmark_appendRowUsingBuffer(t *testing.B) {
	var w bytes.Buffer
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < numberOfRows; i++ {
		appendRowUsingBuffer(&w, weather.RandomStation())
	}
}

func Benchmark_generate(t *testing.B) {
	t.ResetTimer()
	out := generateData(numberOfRows, 100_000, 1000, 0)
	for data := range out {
		_ = data
	}
}

func Benchmark_generate2(t *testing.B) {
	t.ResetTimer()
	out := generateData(numberOfRows, 100_000, 1, 5)
	for data := range out {
		_ = data
	}
}
