package solution

import (
	"aoc/com/dsa"
)

type Direction int

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	}
	panic("invalid!")
}

const (
	UP Direction = 1 << iota
	DOWN
	LEFT
	RIGHT
)

type Grid struct {
	data       [][]rune
	marks      [][]rune
	history    [][]Direction
	rows, cols int
}

func (g *Grid) ResetMarksAndHistory() {
	for i, line := range g.history {
		for j := range line {
			g.history[i][j] = 0
			g.marks[i][j] = '.'
		}
	}
}

func NewGrid(data [][]rune) *Grid {
	rows, cols := len(data), len(data[0])

	marks := make([][]rune, rows)
	for i := range marks {
		marks[i] = make([]rune, cols)
		for j := range marks[i] {
			marks[i][j] = '.'
		}
	}

	history := make([][]Direction, rows)
	for i := range history {
		history[i] = make([]Direction, cols)
		for j := range history[i] {
			history[i][j] = 0
		}
	}

	return &Grid{data: data, marks: marks, history: history, rows: rows, cols: cols}
}

func (g *Grid) Traverse(start *dsa.Point, direction Direction) {
	g.moveForward(start, direction)
}

func (g *Grid) checkHistory(p *dsa.Point, d Direction) bool {
	v := g.history[p.Line][p.Col]
	if v&d != 0 {
		return true
	}
	g.history[p.Line][p.Col] = v | d
	return false
}

func (g *Grid) moveForward(p *dsa.Point, d Direction) {
	if p.IsInGrid(g.rows, g.cols) {
		if g.checkHistory(p, d) {
			return
		}
		g.moveForward(p, d)
	}

	for {
		if !p.IsInGrid(g.rows, g.cols) {
			return
		}

		if g.marks[p.Line][p.Col] == '.' {
			g.marks[p.Line][p.Col] = '#'
		}

		switch d {
		case RIGHT:
			switch g.data[p.Line][p.Col] {
			case '|':
				g.moveForward(p.Cone().MoveUp(), UP)
				g.moveForward(p.MoveDown(), DOWN)
				return
			case '/':
				g.moveForward(p.MoveUp(), UP)
				return
			case '\\':
				g.moveForward(p.MoveDown(), DOWN)
				return
			case '-':
				g.moveForward(p.MoveRight(), RIGHT)
				return
			case '.':
				p.MoveRight()
			default:
				panic("unknown direction")
			}

		case LEFT:
			switch g.data[p.Line][p.Col] {
			case '|':
				g.moveForward(p.Cone().MoveUp(), UP)
				g.moveForward(p.MoveDown(), DOWN)
				return
			case '/':
				g.moveForward(p.MoveDown(), DOWN)
				return
			case '\\':
				g.moveForward(p.MoveUp(), UP)
				return
			case '-':
				g.moveForward(p.MoveLeft(), LEFT)
				return
			case '.':
				p.MoveLeft()
			default:
				panic("unknown direction")
			}

		case DOWN:
			switch g.data[p.Line][p.Col] {
			case '|':
				g.moveForward(p.MoveDown(), DOWN)
				return
			case '/':
				g.moveForward(p.MoveLeft(), LEFT)
				return
			case '\\':
				g.moveForward(p.MoveRight(), RIGHT)
				return
			case '-':
				g.moveForward(p.Cone().MoveLeft(), LEFT)
				g.moveForward(p.MoveRight(), RIGHT)
				return
			case '.':
				p.MoveDown()
			default:
				panic("unknown direction")
			}

		case UP:
			switch g.data[p.Line][p.Col] {
			case '|':
				g.moveForward(p.MoveUp(), UP)
				return
			case '/':
				g.moveForward(p.MoveRight(), RIGHT)
				return
			case '\\':
				g.moveForward(p.MoveLeft(), LEFT)
				return
			case '-':
				g.moveForward(p.Cone().MoveLeft(), LEFT)
				g.moveForward(p.MoveRight(), RIGHT)
				return
			case '.':
				p.MoveUp()
			default:
				panic("unknown direction")
			}
		}
	}
}

func (g *Grid) CountEnergized(start *dsa.Point, direction Direction) (solution int) {
	g.Traverse(start, direction)
	defer g.ResetMarksAndHistory()

	for _, line := range g.marks {
		for _, c := range line {
			if c == '#' {
				solution += 1
			}
		}
		//fmt.Println(string(line))
	}
	//fmt.Println()
	//fmt.Println()

	return solution
}

func CountEnergized(data [][]rune) (solution int) {
	grid := NewGrid(data)
	for i := 0; i < grid.rows; i++ {
		energized := grid.CountEnergized(&dsa.Point{Line: i, Col: 0}, RIGHT)
		energized2 := grid.CountEnergized(&dsa.Point{Line: i, Col: grid.cols - 1}, LEFT)
		solution = max(solution, energized, energized2)
	}
	for i := 0; i < grid.cols; i++ {
		energized := grid.CountEnergized(&dsa.Point{Line: 0, Col: i}, DOWN)
		energized2 := grid.CountEnergized(&dsa.Point{Line: grid.rows - 1, Col: i}, UP)
		solution = max(solution, energized, energized2)
	}
	return solution
}
