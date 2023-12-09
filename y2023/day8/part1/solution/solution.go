package solution

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solveChallenge(inputFilePath string) (int, error) {
	var solution int

	const START = "AAA"
	const END = "ZZZ"

	data, err := readData(inputFilePath)
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

func readData(inputFilePath string) (*ProblemData, error) {
	data := NewProblemData()
	fmt.Println(inputFilePath)

	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
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

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
