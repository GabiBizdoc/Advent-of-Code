package solution

import (
	"aoc/y2023/com"
	"fmt"
	"regexp"
	"strconv"
)

var literalDigitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

const pattern = `one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9`

func findFirstDigitRegex(line string) (int, error) {
	re := regexp.MustCompile(fmt.Sprintf(`(%s)`, pattern))
	match := re.FindString(line)
	if len(match) == 1 {
		return strconv.Atoi(match)
	}
	if result, ok := literalDigitsMap[match]; ok {
		return result, nil
	}
	return 0, DigitNotFoundErr
}

func findLastDigitRegex(line string) (int, error) {
	re := regexp.MustCompile(fmt.Sprintf(`(%s)`, com.CloneReverseString(pattern)))
	match := re.FindString(com.CloneReverseString(line))
	if len(match) == 1 {
		return strconv.Atoi(match)
	}
	if result, ok := literalDigitsMap[com.ReverseString(match)]; ok {
		return result, nil
	}
	return 0, DigitNotFoundErr
}

func firstAndLastRegex(line string) (int, int, error) {
	first, err := findFirstDigitRegex(line)
	if err != nil {
		return 0, 0, err
	}
	last, err := findLastDigitRegex(line)
	if err != nil {
		return 0, 0, err
	}
	return first, last, nil
}

func processLineRegex(line string) (int, error) {
	fist, last, err := firstAndLastRegex(line)
	return fist*10 + last, err
}
