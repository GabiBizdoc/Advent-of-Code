package solution

import (
	"errors"
)

func cmpLine(a, b []rune) int {
	diffs := 0
	if len(a) != len(b) {
		panic("lines must have the same len")
	}

	for i, c := range a {
		if b[i] != c {
			diffs += 1
		}
	}
	return diffs
}

func isMirror(k int, grid [][]rune, nudges int) bool {
	if k == len(grid)-1 {
		return false
	}
	for i := 0; i < len(grid); i++ {
		left, right := k-i, k+i+1
		if left < 0 || right >= len(grid) {
			return nudges == 0
		}
		line1 := grid[left]
		line2 := grid[right]
		nudges -= cmpLine(line1, line2)
	}
	return nudges == 0
}

func findPointOfIncidence(grid [][]rune, nudges int) (int, bool) {
	for i := 0; i < len(grid); i++ {
		ok := isMirror(i, grid, nudges)
		if ok {
			return i, true
		}
	}
	return 0, false
}

// SolveFor finds an incidence point by checking each line of the grid and its transpose.
// It counts the required nudges for a possible incidence point, then compares with the provided nudges
// Bails early if the provided nudges are exceeded; benchmarks show no performance impact.
func SolveFor(grid [][]rune, nudges int) int {
	s, err := solveFor(grid, nudges)
	if err != nil {
		panic(err)
	}
	return s.Solution
}

func solveFor(grid [][]rune, nudges int) (*response, error) {
	i, ok := findPointOfIncidence(grid, nudges)
	if ok {
		solution := i + 1
		solution *= 100
		return &response{solution, i, "h", grid}, nil
	}

	grid2 := transpose(grid)
	j, ok := findPointOfIncidence(grid2, nudges)
	if ok {
		solution := j + 1
		return &response{solution, i, "v", grid2}, nil
	}

	return nil, errors.New("solution not found")
}

// Response object is used for debugging
type response struct {
	Solution int
	Line     int
	// `h` or `v`
	Orientation string
	GridUsed    [][]rune
}
