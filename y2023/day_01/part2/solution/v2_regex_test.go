package solution

import "testing"

func TestFindFirstDigitRegex(t *testing.T) {
	testFirstDigit(t, findFirstDigitRegex)
}

func TestFindLastDigitRegex(t *testing.T) {
	testLastDigit(t, findLastDigitRegex)
}

func TestSolutionRegex(t *testing.T) {
	testProcessLine(t, processLineRegex)
}

func BenchmarkFindFirstDigitRegex(b *testing.B) {
	testFirstDigit(b, findFirstDigitRegex)
}

func BenchmarkFindLastDigitRegex(b *testing.B) {
	testLastDigit(b, findLastDigitRegex)
}

func BenchmarkSolutionRegex(t *testing.B) {
	testProcessLine(t, processLineRegex)
}
