package solution

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type lineHandler = func(line string) (*Game, error)

func solveChallenge(inputFilePath string, onLine lineHandler) (int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Solve(file)
}

func Solve(file io.Reader) (int, error) {
	var solution int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if game, err := LineParser(scanner.Text()); err != nil {
			if !errors.Is(err, InvalidGameSet) {
				return 0, err
			}
		} else {
			solution += game.GameID
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return solution, nil
}

func LineParser(line string) (*Game, error) {
	parts := strings.Split(line, ": ")
	gameId, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(parts[0], "Game ")))
	if err != nil {
		return nil, err
	}
	game := NewGame(gameId)
	games := strings.Split(strings.TrimSpace(parts[1]), "; ")
	for _, pairs := range games {
		gameSet := NewGameSet()
		for _, p := range strings.Split(strings.TrimSpace(pairs), ", ") {
			a := strings.Split(p, " ")
			if len(a) != 2 {
				return nil, fmt.Errorf("failed to parse [count, color]")
			}

			countStr, color := a[0], a[1]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return nil, err
			}
			gameSet.AddCube(color, count)
		}

		if gameSet.IsValid() {
			game.Sets = append(game.Sets, gameSet)
		} else {
			return nil, InvalidGameSet
		}
	}

	return game, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath, LineParser)
}
