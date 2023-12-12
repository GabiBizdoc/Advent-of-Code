package dsa

import "fmt"

type Point struct {
	Line int
	Col  int
}

func (p *Point) MoveLeft() *Point {
	p.Col -= 1
	return p
}
func (p *Point) MoveRight() *Point {
	p.Col += 1
	return p
}
func (p *Point) MoveUp() *Point {
	p.Line -= 1
	return p
}
func (p *Point) MoveDown() *Point {
	p.Line += 1
	return p
}
func (p *Point) Cone() *Point {
	//return &*p
	return &Point{p.Line, p.Col}
}
func (p *Point) IsInGrid(rows, columns int) bool {
	return 0 <= p.Line && p.Line < rows && 0 <= p.Col && p.Col < columns
}

func (p *Point) Neighbours() []*Point {
	neighbours := make([]*Point, 0, 4)
	neighbours = append(neighbours, p.Cone().MoveUp())
	neighbours = append(neighbours, p.Cone().MoveDown())
	neighbours = append(neighbours, p.Cone().MoveLeft())
	neighbours = append(neighbours, p.Cone().MoveRight())
	return neighbours
}

func (p *Point) String() string {
	return fmt.Sprintf("%v", *p)
}
