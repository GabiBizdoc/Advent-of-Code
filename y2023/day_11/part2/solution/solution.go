package solution

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	solution := 0
	maze := NewUniverse()

	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		maze.AddLine(line)

	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	solution = maze.Solve(1_000_000)
	fmt.Println("solution: ", solution)
	return solution, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
