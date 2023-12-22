package solution

import (
	"strings"
)

const movableRock = 'O'
const immovableRock = '#'
const emptySpace = '.'

func CalculateLoad(grid [][]rune) (solution int) {
	for i, line := range grid {
		for _, c := range line {
			if c == movableRock {
				solution += len(grid) - i
			}
		}
	}
	return solution
}

func tiltNorth(grid [][]rune) {
	rocksCount := 0
	start := 0
	for j := 0; j < len(grid[0]); j++ {
		for i, _ := range grid {
			switch grid[i][j] {
			case emptySpace:
				continue
			case movableRock:
				rocksCount += 1
				grid[i][j] = emptySpace
			case immovableRock:
				for k := 0; k < rocksCount; k++ {
					grid[k+start][j] = movableRock
				}
				start = i + 1
				rocksCount = 0
			}
		}
		for k := 0; k < rocksCount; k++ {
			grid[k+start][j] = movableRock
		}
		start = 0
		rocksCount = 0
	}
}
func tiltWest(grid [][]rune) {
	rocksCount := 0
	start := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case emptySpace:
				continue
			case movableRock:
				rocksCount += 1
				grid[i][j] = emptySpace
			case immovableRock:
				for k := 0; k < rocksCount; k++ {
					grid[i][k+start] = movableRock
				}
				start = j + 1
				rocksCount = 0
			}
		}
		for k := 0; k < rocksCount; k++ {
			grid[i][k+start] = movableRock
		}
		start = 0
		rocksCount = 0
	}
}
func tiltSouth(grid [][]rune) {
	rocksCount := 0
	for j := 0; j < len(grid[0]); j++ {
		for i, _ := range grid {
			switch grid[i][j] {
			case emptySpace:
				continue
			case movableRock:
				rocksCount += 1
				grid[i][j] = emptySpace
			case immovableRock:
				start := i - rocksCount
				for k := 0; k < rocksCount; k++ {
					grid[k+start][j] = movableRock
				}
				start = i + 1
				rocksCount = 0
			}
		}
		start := len(grid) - rocksCount

		for k := 0; k < rocksCount; k++ {
			grid[k+start][j] = movableRock
		}
		start = 0
		rocksCount = 0
	}
}

func tiltEast(grid [][]rune) {
	rocksCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case emptySpace:
				continue
			case movableRock:
				rocksCount += 1
				grid[i][j] = emptySpace
			case immovableRock:
				start := j - rocksCount
				for k := 0; k < rocksCount; k++ {
					grid[i][k+start] = movableRock
				}
				start = j + 1
				rocksCount = 0
			}
		}
		start := len(grid[i]) - rocksCount
		for k := 0; k < rocksCount; k++ {
			grid[i][k+start] = movableRock
		}
		start = 0
		rocksCount = 0
	}
}

func cycle(grid [][]rune) {
	tiltNorth(grid)
	tiltWest(grid)
	tiltSouth(grid)
	tiltEast(grid)
}

func CalculateLoadAfterCycles(grid [][]rune, cycles int) int {
	cache := make(map[string]int)
	for i := 0; i <= cycles; i++ {
		key := gridToString(grid)
		if v, ok := cache[key]; ok {
			loopSize := i - v
			leftoverCycles := cycles - v - 1
			return CalculateLoadAfterCycles(grid, leftoverCycles%loopSize)
		} else {
			cache[key] = i
		}
		cycle(grid)
	}
	return CalculateLoad(grid)
}

// attempt to cache values. Doesn't work
func gridToString(grid [][]rune) string {
	var sb strings.Builder
	sb.Grow(len(grid) * len(grid[0]))
	for _, line := range grid {
		for _, c := range line {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
