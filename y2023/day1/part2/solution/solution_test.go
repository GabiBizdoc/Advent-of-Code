package solution

import (
	"errors"
	"strings"
	"testing"
)

type TestCaseData struct {
	Line  string
	First int
	Last  int
	Err   error
}

// createLongString don't forget to update testData if you change the string
func createLongString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteString("something something bla bla twone bla bla something something. ")
	}
	return builder.String()
}

var testData = [...]TestCaseData{
	{Line: "two1nine", First: 2, Last: 9},
	{Line: "eightwothree", First: 8, Last: 3},
	{Line: "abcone2threexyz", First: 1, Last: 3},
	{Line: "xtwone3four", First: 2, Last: 4},
	{Line: "4nineeightseven2", First: 4, Last: 2},
	{Line: "zoneight234", First: 1, Last: 4},
	{Line: "7pqrstsixteen", First: 7, Last: 6},
	{Line: "7pqrstsixteen", First: 7, Last: 6},
	{Line: "twone3twone", First: 2, Last: 1},
	{Line: createLongString(1_000_000), First: 2, Last: 1},
}

type digitFinder = func(lineStr string) (int, error)

func testFirstDigit(t testing.TB, df digitFinder) {
	for i, testCase := range testData {
		digit, err := df(testCase.Line)
		expected := testCase.First

		if err != nil {
			if !errors.Is(err, testCase.Err) {
				t.Errorf("%d: expected error %v, got %v", i, testCase.Err, err)
			}
		} else if digit != expected {
			t.Errorf("%d: `%s` expected %d, but got %d", i, testCase.Line, expected, digit)
		}
	}
}

func testLastDigit(t testing.TB, df digitFinder) {
	for i, testCase := range testData {
		digit, err := df(testCase.Line)
		expected := testCase.Last

		if err != nil {
			if !errors.Is(err, testCase.Err) {
				t.Errorf("%d: expected error %v, got %v", i, testCase.Err, err)
			}
		} else if digit != expected {
			t.Errorf("%d: `%s` expected %d, but got %d", i, testCase.Line, expected, digit)
		}
	}
}

func testProcessLine(t testing.TB, df digitFinder) {
	for i, testCase := range testData {
		sol, err := df(testCase.Line)
		expected := testCase.First*10 + testCase.Last

		if err != nil {
			if !errors.Is(err, testCase.Err) {
				t.Errorf("%d: expected error %v, got %v", i, testCase.Err, err)
			}
		} else if sol != expected {
			t.Errorf("%d: `%s` expected %d, but got %d", i, testCase.Line, expected, sol)
		}
	}
}
