package solution

import (
	"aoc/y2023/day3/part2/solution/board"
	"bufio"
	"fmt"
	"io"
	"os"
)

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
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

	for _, neighbours := range b.FindStarsWithNeighbours() {
		if len(neighbours) == 2 {
			solution += b.ReadGridNumber(neighbours[0]) * b.ReadGridNumber(neighbours[1])
		}
	}

	return solution, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
