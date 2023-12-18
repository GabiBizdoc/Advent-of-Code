package board

import (
	"unicode"
)

type Point struct {
	Line int
	Col  int
}

type GridNumber struct {
	points []*Point
}

func (g *GridNumber) Add(p *Point) {
	g.points = append(g.points, p)
}

func NewBoard() *Board {
	return &Board{}
}

type Board struct {
	board [][]rune

	// array of squares containing numbers
	numbers []*GridNumber
}

func (b *Board) AppendLine(line string) error {
	runicLine := []rune(line)
	currentLine := len(b.board)
	b.board = append(b.board, runicLine)

	var gridNumber *GridNumber

	for i, r := range runicLine {
		if unicode.IsNumber(r) {
			if gridNumber == nil {
				gridNumber = &GridNumber{}
			}
			gridNumber.Add(&Point{Line: currentLine, Col: i})
		} else if gridNumber != nil {
			b.numbers = append(b.numbers, gridNumber)
			gridNumber = nil
		}
	}

	if gridNumber != nil {
		b.numbers = append(b.numbers, gridNumber)
		gridNumber = nil
	}
	return nil
}

func (b *Board) ReadPoint(point *Point) rune {
	return b.board[point.Line][point.Col]
}

func (b *Board) FindStarsWithNeighbours() map[Point][]*GridNumber {
	stars := make(map[Point][]*GridNumber)
	for _, gridNum := range b.numbers {
		if ok, point := b.isPartNumberWithStar(gridNum); ok {
			// check if the symbol is a star
			if b.ReadPoint(point) == '*' {
				stars[*point] = append(stars[*point], gridNum)
			}
		}
	}

	return stars
}

func (b *Board) ReadGridNumber(n *GridNumber) int {
	value := 0
	for _, point := range n.points {
		digit := b.ReadPoint(point) - '0'
		value *= 10
		value += int(digit)
	}
	return value
}

func (b *Board) isPartNumberWithStar(n *GridNumber) (bool, *Point) {
	// we don't want to change the original points
	startPoint := *n.points[0]
	endPoint := *n.points[len(n.points)-1]

	linesCount := len(b.board)
	colsCount := len(b.board[startPoint.Line])

	if startPoint.Col > 0 {
		startPoint.Col -= 1
	}

	if endPoint.Col < colsCount-1 {
		endPoint.Col += 1
	}

	if p := b.ReadPoint(&startPoint); isSymbol(p) {
		return true, &startPoint
	}

	if isSymbol(b.ReadPoint(&endPoint)) {
		return true, &endPoint
	}

	// check the above the line
	if startPoint.Line > 0 {
		above := startPoint.Line - 1
		for i := startPoint.Col; i <= endPoint.Col; i++ {
			if isSymbol(b.board[above][i]) {
				return true, &Point{Line: above, Col: i}
			}
		}
	}

	// check the under the line
	if startPoint.Line < linesCount-1 {
		under := startPoint.Line + 1

		for i := startPoint.Col; i <= endPoint.Col; i++ {
			if isSymbol(b.board[under][i]) {
				return true, &Point{Line: under, Col: i}
			}
		}
	}

	return false, nil
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}
