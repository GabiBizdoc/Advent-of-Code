package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve(file io.Reader) (solution int, err error) {
	scanner := bufio.NewScanner(file)

	var timing int
	var distance int

	if scanner.Scan() {
		line := scanner.Text()
		var sb strings.Builder
		if !strings.HasPrefix(line, "Time: ") {
			return 0, fmt.Errorf("the first line should contain timings")
		}
		for _, c := range strings.TrimPrefix(line, "Time: ") {
			if unicode.IsDigit(c) {
				sb.WriteRune(c)
			}
		}
		timing, err = strconv.Atoi(sb.String())
		if err != nil {
			return 0, err
		}
	}
	if scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Distance: ") {
			return 0, fmt.Errorf("the seccond line should contain distances")
		}
		var sb strings.Builder
		for _, c := range strings.TrimPrefix(line, "Distance: ") {
			if unicode.IsDigit(c) {
				sb.WriteRune(c)
			}
		}
		distance, err = strconv.Atoi(sb.String())
		if err != nil {
			return 0, err
		}
	}

	//k := FindCoefficient(timing, distance)

	x1, x2 := FindRoots(float64(timing), float64(distance))
	k := int(x2) - int(x1)

	solution = k
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
