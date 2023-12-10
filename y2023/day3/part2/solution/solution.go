package solution

import (
	"aoc/y2023/day3/part2/solution/board"
	"bufio"
	"os"
)

func solveChallenge(inputFilePath string) (int, error) {
	var solution int

	file, err := os.Open(inputFilePath)
	if err != nil {
		return solution, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	b := board.NewBoard()
	for scanner.Scan() {
		err := b.AppendLine(scanner.Text())
		if err != nil {
			return 0, err
		}
	}

	if err = scanner.Err(); err != nil {
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
