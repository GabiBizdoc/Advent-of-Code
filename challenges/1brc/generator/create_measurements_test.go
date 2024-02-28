package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var defaultConfig = NewConfig()

func init() {
	defaultConfig.Lines.Value = 1_000_000_000
	defaultConfig.Generators.Value = 10
	defaultConfig.WriterChannelSize.Value = 10
}

func Benchmark_appendRowFormattedString(t *testing.B) {
	w := strings.Builder{}
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < defaultConfig.Lines.Value; i++ {
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
	for i := 0; i < defaultConfig.Lines.Value; i++ {
		appendRowWithoutFormat(w, weather.RandomStation())
	}
}

func Benchmark_appendRowWithoutFormatScaled(t *testing.B) {
	w := strings.Builder{}
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < defaultConfig.Lines.Value; i++ {
		appendRowWithoutFormatScaled(w, weather.RandomStation())
	}
}

func Benchmark_appendRowToByteSlice(t *testing.B) {
	w := make([]byte, 0)
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < defaultConfig.Lines.Value; i++ {
		w = appendRowToByteSlice(w, weather.RandomStation())
	}
}

func Benchmark_appendRowToByteSliceWithCapacity(t *testing.B) {
	w := make([]byte, 0, defaultConfig.Lines.Value*10)
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < defaultConfig.Lines.Value; i++ {
		w = appendRowToByteSlice(w, weather.RandomStation())
	}
}

func Benchmark_appendRowUsingBuffer(t *testing.B) {
	var w bytes.Buffer
	weather := NewWeatherStationsGenerator()
	t.ResetTimer()
	for i := 0; i < defaultConfig.Lines.Value; i++ {
		appendRowUsingBuffer(&w, weather.RandomStation())
	}
}

func Test_Generate(t *testing.T) {
	f, err := os.OpenFile("/dev/null", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	generate(defaultConfig, f)
}

func Benchmark_generateData(t *testing.B) {
	out := generateData(defaultConfig.Lines.Value, 100_000, 1000, 0)
	for data := range out {
		_ = data
	}
}

func Benchmark_generateData2(t *testing.B) {
	out := generateData(defaultConfig.Lines.Value, 100_000, 20, 5)
	for data := range out {
		_ = data
	}
}
