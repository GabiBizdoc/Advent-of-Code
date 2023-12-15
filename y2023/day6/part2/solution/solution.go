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

func ReadData(file io.Reader) (pd *ProblemData, err error) {
	scanner := bufio.NewScanner(file)
	pd = &ProblemData{}

	if scanner.Scan() {
		line := scanner.Text()
		var sb strings.Builder
		if !strings.HasPrefix(line, "Time: ") {
			return nil, fmt.Errorf("the first line should contain timings")
		}
		for _, c := range strings.TrimPrefix(line, "Time: ") {
			if unicode.IsDigit(c) {
				sb.WriteRune(c)
			}
		}
		pd.Timing, err = strconv.Atoi(sb.String())
		if err != nil {
			return nil, err
		}
	}
	if scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Distance: ") {
			return nil, fmt.Errorf("the seccond line should contain distances")
		}
		var sb strings.Builder
		for _, c := range strings.TrimPrefix(line, "Distance: ") {
			if unicode.IsDigit(c) {
				sb.WriteRune(c)
			}
		}
		pd.Distance, err = strconv.Atoi(sb.String())
		if err != nil {
			return nil, err
		}
	}

	return pd, nil
}

func Solve(file io.Reader) (solution int, err error) {
	scanner := bufio.NewScanner(file)
	pd, err := ReadData(file)
	if err != nil {
		return 0, err
	}
	solution = FindCoefficientUsingRoots(pd.Timing, pd.Distance)
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
