package solution

import (
	"errors"
	"testing"
)

type TestCaseData struct {
	Line     string
	Err      error
	Expected int
}

var testData = [...]TestCaseData{
	{Line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", Expected: 48},
	{Line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", Expected: 12},
	{Line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", Err: InvalidGameSet, Expected: 1560},
	{Line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", Err: InvalidGameSet, Expected: 630},
	{Line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", Expected: 36},
}

func TestLineParser(t *testing.T) {
	for i, testCase := range testData {
		game, err := LineParser(testCase.Line)
		if err != nil {
			if !errors.Is(err, testCase.Err) {
				t.Errorf("%d: expected error %v, got %v", i, testCase.Err, err)
			}
		} else {
			solution := game.ComputeCubePower()
			if solution != testCase.Expected {
				t.Errorf("%d: expected sum %d, got %d", i, testCase.Expected, solution)
			}
		}
	}
}
