package mapper

type Rule struct {
	sourceStart, destStart, size int
}

func (r *Rule) Contains(x int) bool {
	return x >= r.sourceStart && x <= r.sourceStart+r.size
}

func (r *Rule) Read(x int) (int, bool) {
	if r.Contains(x) {
		of1 := x - r.sourceStart
		return of1 + r.destStart, true
	}
	return 0, false
}

func NewMapperRule(sourceStart, destStart, size int) *Rule {
	return &Rule{sourceStart: sourceStart, destStart: destStart, size: size}
}
