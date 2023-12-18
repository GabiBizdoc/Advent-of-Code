package solution

import (
	"fmt"
)

type Universe struct {
	grid       [][]CosmicEntity
	rows, cols int
}

func (u *Universe) AddLine(line string) {
	runicLine := []CosmicEntity(line)
	u.grid = append(u.grid, runicLine)
}

func NewUniverse() *Universe {
	gird := make([][]CosmicEntity, 0)
	return &Universe{grid: gird}
}

func (u *Universe) Solve(universeMultiplier int) int {
	u.rows = len(u.grid)
	u.cols = len(u.grid[0])

	linesToAdd, colsToAdd := u.Expand()
	//printGrid(u.grid)

	galaxies := u.findGalaxies()

	// todo:
	for i, g := range galaxies {
		for _, lineN := range linesToAdd {
			if g.Line > lineN {
				galaxies[i].Line += universeMultiplier - 1
			}
		}
		for _, colN := range colsToAdd {
			if g.Col > colN {
				galaxies[i].Col += universeMultiplier - 1
			}
		}
	}

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

func (u *Universe) Expand() (linesToAdd []int, colsToAdd []int) {
	linesToAdd = make([]int, 0)
	colsToAdd = make([]int, 0)

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
	return linesToAdd, colsToAdd
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
