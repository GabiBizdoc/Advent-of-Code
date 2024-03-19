package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve(file io.Reader) (int, error) {
	solution := 0
	maze := NewUniverse()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		maze.AddLine(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	solution = maze.Solve(1_000_000)
	fmt.Println("solution: ", solution)
	return solution, nil
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
