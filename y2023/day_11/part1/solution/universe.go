package solution

import (
	"fmt"
	"slices"
)

type Universe struct {
	grid       [][]CosmicEntity
	rows, cols int
}

func (u *Universe) isPointInsideGrid(p *Point) bool {
	rows := len(u.grid)
	columns := len(u.grid[0])
	return 0 <= p.Line && p.Line < rows && 0 <= p.Col && p.Col < columns
}

func (u *Universe) AddLine(line string) {
	runicLine := []CosmicEntity(line)
	u.grid = append(u.grid, runicLine)
}

func (u *Universe) Read(p *Point) CosmicEntity {
	return u.grid[p.Line][p.Col]
}

func NewMaze() *Universe {
	maze := make([][]CosmicEntity, 0)
	return &Universe{grid: maze}
}

func (u *Universe) Solve() int {
	u.rows = len(u.grid)
	u.cols = len(u.grid[0])

	u.Expand()
	//printGrid(u.grid)

	galaxies := u.findGalaxies()

	total := 0
	for i, g1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			d := abs(g1.Col-g2.Col) + abs(g1.Line-g2.Line)

			total += d
		}
	}
	return total
}

func (u *Universe) findGalaxies() []Point {
	galaxies := make([]Point, 0)
	for i, line := range u.grid {
		for j, x := range line {
			if x.IsGalaxy() {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}
	return galaxies
}

func (u *Universe) Expand() {
	// todo: dont expand the galaxy... expand the points. then copy to another grid if needed
	linesToAdd := make([]int, 0)
	colsToAdd := make([]int, 0)

	for i := 0; i < u.rows; i++ {
		isEmpty := true
		for j := 0; j < u.cols; j++ {
			if u.grid[i][j].IsGalaxy() {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			linesToAdd = append(linesToAdd, i)
		}
	}

	for j := 0; j < u.cols; j++ {
		isEmpty := true
		for i := 0; i < u.rows; i++ {
			if u.grid[i][j].IsGalaxy() {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			colsToAdd = append(colsToAdd, j)
		}
	}

	fmt.Println("linesToAdd", linesToAdd, "colsToAdd", colsToAdd)
	for i, lineNumber := range linesToAdd {
		next := make([]CosmicEntity, u.cols)
		for i2 := range next {
			next[i2] = '.' // empty space
		}

		u.grid = slices.Insert(u.grid, i+lineNumber, next)
	}
	u.rows += len(linesToAdd)

	for i, colNumber := range colsToAdd {
		for k, line := range u.grid {
			u.grid[k] = slices.Insert(line, i+colNumber, '.')
		}
	}
	u.cols += len(colsToAdd)
}

// utils

func printGrid(g [][]CosmicEntity) {
	fmt.Print("\n")
	for _, lines := range g {
		for _, v := range lines {
			fmt.Printf("%2c", v)
		}
		fmt.Print("\n")
	}
}

func abs[T ~int](a T) T {
	if a < 0 {
		return a * -1
	}
	return a
}
