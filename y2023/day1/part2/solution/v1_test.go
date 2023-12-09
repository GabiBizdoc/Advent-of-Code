package solution

import (
	"testing"
)

func TestFindFirstDigit(t *testing.T) {
	testFirstDigit(t, findFirstDigit)
}

func TestFindLastDigit(t *testing.T) {
	testLastDigit(t, findLastDigit)
}

func TestSolution(t *testing.T) {
	testProcessLine(t, processLine)
}

func BenchmarkFindFirstDigit(b *testing.B) {
	testFirstDigit(b, findFirstDigit)
}

func BenchmarkFindLastDigit(b *testing.B) {
	testLastDigit(b, findLastDigit)
}

func BenchmarkSolution(t *testing.B) {
	testProcessLine(t, processLine)
}
