package solution

import (
	"bufio"
	"errors"
	"os"
)

var DigitNotFoundErr = errors.New("digit not found in string")

func solveChallenge(inputFilePath string, onLine func(string) (int, error)) (int, error) {
	var solution int

	file, err := os.Open(inputFilePath)
	if err != nil {
		return solution, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if calibration, err := onLine(scanner.Text()); err != nil {
			return 0, err
		} else {
			solution += calibration
		}
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	return solution, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath, processLine)
}

func SolveChallengeRegex(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath, processLineRegex)
}
