package solution

import (
	"strings"
	"unicode"
)

func findFirstDigit09(line []rune) (int, int) {
	for i, c := range line {
		if unicode.IsNumber(c) {
			return int(c - '0'), i
		}
	}
	return 0, -1
}

func findLastDigit09(line []rune) (int, int) {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if unicode.IsNumber(c) {
			return int(c - '0'), i
		}
	}
	return 0, -1
}

var literalDigits = [...]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func findFirstDigit(lineStr string) (int, error) {
	var index int
	var value int

	lineRunes := []rune(lineStr)
	value, index = findFirstDigit09(lineRunes)

	for i, digit := range literalDigits {
		ind := strings.Index(lineStr, digit)
		if ind == -1 {
			continue
		}

		if index == -1 || ind < index {
			index = ind
			value = i + 1
		}
	}

	if index == -1 {
		return 0, DigitNotFoundErr
	}
	return value, nil
}

func findLastDigit(lineStr string) (int, error) {
	var index int
	var value int

	lineRunes := []rune(lineStr)
	value, index = findLastDigit09(lineRunes)

	for i, digit := range literalDigits {
		ind := strings.LastIndex(lineStr, digit)
		if ind == -1 {
			continue
		}

		if index == -1 || ind > index {
			index = ind
			value = i + 1
		}
	}

	if index == -1 {
		return 0, DigitNotFoundErr
	}
	return value, nil
}

func firstAndLastDigit(line string) (int, int, error) {
	first, err := findFirstDigit(line)
	if err != nil {
		return 0, 0, err
	}

	last, err := findLastDigit(line)
	if err != nil {
		return 0, 0, err
	}

	return first, last, nil
}

func processLine(line string) (int, error) {
	fist, last, err := firstAndLastDigit(line)
	return fist*10 + last, err
}
