package solution

import (
	"aoc/y2023/day_03/part1/solution/board"
	"bufio"
	"io"
	"os"
)

func solveChallenge(inputFilePath string) (int, error) {
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

	b := board.NewBoard()
	for scanner.Scan() {
		err := b.AppendLine(scanner.Text())
		if err != nil {
			return 0, err
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	for _, number := range b.FindPartNumbers() {
		solution += number
	}

	return solution, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
