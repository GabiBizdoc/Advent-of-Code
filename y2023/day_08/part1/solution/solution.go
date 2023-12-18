package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve(file io.Reader) (int, error) {
	var solution int

	const START = "AAA"
	const END = "ZZZ"

	data, err := readData(file)
	if err != nil {
		return 0, err
	}

	c := START
	for c != END {
		for _, x := range data.Path {
			solution += 1
			next := data.Rows[c]
			switch x {
			case 'L':
				c = next.Left
			case 'R':
				c = next.Right
			}
		}
	}

	fmt.Println(solution)
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

func readData(file io.Reader) (*ProblemData, error) {
	data := NewProblemData()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		data.Path = strings.TrimSpace(line)
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " = (")

		leftRight := strings.Split(parts[1], ", ")
		index := parts[0]
		leftPart := leftRight[0]
		rightPart := leftRight[1][:len(leftRight[1])-1]
		data.Rows[index] = &RowItem{Left: leftPart, Right: rightPart}
	}

	return data, scanner.Err()
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
