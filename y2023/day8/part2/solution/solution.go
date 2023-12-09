package solution

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// GDC greatest common divisor
func GDC(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM least common multiple
func LCM(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(GDC(a, b)))
}

func findLCMOfSlice(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result = LCM(result, num)
	}
	return result
}

func solveFor(data *ProblemData, start string) int {
	var solution int

	for !strings.HasSuffix(start, "Z") {
		for _, x := range data.Path {
			solution += 1
			next := data.Rows[start]
			switch x {
			case 'L':
				start = next.Left
			case 'R':
				start = next.Right
			}
		}
	}

	return solution
}

func solveChallenge(inputFilePath string) (int, error) {
	var solution int

	data, err := readData(inputFilePath)
	if err != nil {
		return 0, err
	}

	currentPositions := make([]string, 0)
	for key := range data.Rows {
		if strings.HasSuffix(key, "A") {
			currentPositions = append(currentPositions, key)
		}
	}

	depths := make([]int, len(currentPositions))
	for i, position := range currentPositions {
		depths[i] = solveFor(data, position)
	}

	solution = findLCMOfSlice(depths)
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
