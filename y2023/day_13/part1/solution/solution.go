package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(file io.Reader, strategy func([][]rune) int) (solution int, err error) {
	scanner := bufio.NewScanner(file)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			solution += strategy(grid)
			grid = make([][]rune, 0)
		} else {
			grid = append(grid, []rune(line))
		}
	}
	solution += strategy(grid)
	grid = make([][]rune, 0)
	if err := scanner.Err(); err != nil {
		return solution, err
	}

	return solution, nil
}

func Solve(file io.Reader) (solution int, err error) {
	return solve(file, func(data [][]rune) int {
		return SolveFor(data, 0)
	})
}

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Solve(file)
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
