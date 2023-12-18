package solution

import (
	"errors"
	"testing"
)

type TestCaseData struct {
	Line string
	Err  error
	ID   int
}

var testData = [...]TestCaseData{
	{Line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", ID: 1},
	{Line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", ID: 2},
	{Line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", Err: InvalidGameSet, ID: 3},
	{Line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", Err: InvalidGameSet, ID: 4},
	{Line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", ID: 5},
}

func TestLineParser(t *testing.T) {
	for i, testCase := range testData {
		game, err := LineParser(testCase.Line)
		if err != nil {
			if !errors.Is(err, testCase.Err) {
				t.Errorf("%d: expected error %v, got %v", i, testCase.Err, err)
			}
		} else if game.GameID != testCase.ID {
			t.Errorf("%d: invalid game id. expected gameid %d, got %d", i, testCase.ID, game.GameID)
		}
	}
}
