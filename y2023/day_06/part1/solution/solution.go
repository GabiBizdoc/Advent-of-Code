package solution

import (
	"aoc/com"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve(file io.Reader) (solution int, err error) {
	solution = 1
	scanner := bufio.NewScanner(file)

	var timings []int
	var distances []int

	if scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Time: ") {
			return 0, fmt.Errorf("the first line should contain timings")
		}
		timings, err = com.ExtractIntSliceFromString(strings.TrimPrefix(line, "Time: "))
		if err != nil {
			return 0, err
		}
	}
	if scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Distance: ") {
			return 0, fmt.Errorf("the seccond line should contain distances")
		}
		distances, err = com.ExtractIntSliceFromString(strings.TrimPrefix(line, "Distance: "))
		if err != nil {
			return 0, err
		}
	}

	for i, timing := range timings {
		distance := distances[i]
		solution *= FindCoefficientUsingRoots(timing, distance)
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
