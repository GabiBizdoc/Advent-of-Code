package solution

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
