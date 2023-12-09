package solution

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

var DigitNotFoundErr = errors.New("digit not found in string")

func findFirstDigit(line []rune) (int, error) {
	for _, c := range line {
		if unicode.IsNumber(c) {
			return int(c - '0'), nil
		}
	}
	return 0, DigitNotFoundErr
}

func findLastDigit(line []rune) (int, error) {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if unicode.IsNumber(c) {
			return int(c - '0'), nil
		}
	}
	return 0, DigitNotFoundErr
}

func processLine(line string) (int, error) {
	runicLine := []rune(line)
	var calibration int
	if digit, err := findFirstDigit(runicLine); err != nil {
		return calibration, err
	} else {
		calibration = digit * 10
	}

	if digit, err := findLastDigit(runicLine); err != nil {
		return calibration, err
	} else {
		calibration += digit
	}
	return calibration, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	var solution int

	file, err := os.Open(inputFilePath)
	if err != nil {
		return solution, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if calibration, err := processLine(scanner.Text()); err != nil {
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
