package solution

import (
	"errors"
	"testing"
)

type TestCaseData struct {
	Line  string
	First int
	Last  int
	Err   error
}

var testData = [...]TestCaseData{
	{Line: "1abc2", First: 1, Last: 2},
	{Line: "pqr3stu8vwx", First: 3, Last: 8},
	{Line: "a1b2c3d4e5f", First: 1, Last: 5},
	{Line: "treb7uchet", First: 7, Last: 7},
	{Line: "qwertyuiop", Err: DigitNotFoundErr},
}

func TestFindFirstDigit(t *testing.T) {
	for i, testCase := range testData {
		digit, err := findFirstDigit([]rune(testCase.Line))
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

func TestFindLastDigit(t *testing.T) {
	for i, testCase := range testData {
		digit, err := findLastDigit([]rune(testCase.Line))
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
