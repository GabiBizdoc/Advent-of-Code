package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func extractNumbersFromString(line string) ([]int, error) {
	result := make([]int, 0)
	for _, n := range strings.Split(line, " ") {
		n = strings.TrimSpace(n)
		if n != "" {
			number, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			result = append(result, number)
		}
	}
	return result, nil
}

func parseLine(line string) (*RowData, error) {
	data := &RowData{}
	var err error

	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "Card ")
	parts := strings.Split(line, ": ")
	data.GameID, err = strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, err
	}

	gameParts := strings.Split(parts[1], " | ")
	data.WinningNumbers, err = extractNumbersFromString(gameParts[0])
	if err != nil {
		return nil, err
	}

	data.ExtractedNumbers, err = extractNumbersFromString(gameParts[1])
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Solve(file io.Reader) (int, error) {
	solution := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		rowData, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		solution += rowData.FindCardValue()
	}

	return solution, scanner.Err()
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
