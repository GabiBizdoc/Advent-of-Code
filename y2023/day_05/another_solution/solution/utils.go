package solution

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Row [3]int

func (r *Row) Destination() int {
	return r[1]
}
func (r *Row) SourceStart() int {
	return r[0]
}
func (r *Row) SourceEnd() int {
	return r[0] + r.Range()
}
func (r *Row) Range() int {
	return r[2]
}

type Map struct {
	From, To string

	// destination; source; range
	Rows []Row
}

func (m *Map) Next(x int) int {
	for _, row := range m.Rows {
		if x >= row.SourceStart() && x < row.SourceEnd() {
			next := x - row.SourceStart() + row.Destination()
			if DebugEnabled {
				log.Println("seed", x,
					"start", row.SourceStart(),
					"destination", row.Destination(),
					"range", row.Range(),
					"next", next,
				)
			}
			return next
		}
	}
	return x
}

func (m *Map) AddRow(d, s, r int) {
	m.Rows = append(m.Rows, Row{d, s, r})
}

func NewMap(from string, to string) *Map {
	return &Map{From: from, To: to, Rows: make([]Row, 0)}
}

type Almanac struct {
	maps []*Map
}

func NewAlmanac() *Almanac {
	return &Almanac{maps: make([]*Map, 0)}
}

func (a *Almanac) AddMap(m *Map) error {
	_, err := a.Find(m.From, m.To)
	if errors.Is(err, NotFoundErr) {
		a.maps = append(a.maps, m)
		return nil
	}
	return errors.New("multiple maps with the same name were found")
}
func (a *Almanac) Find(from, to string) (*Map, error) {
	i := slices.IndexFunc(a.maps, func(m *Map) bool {
		return m.From == from && m.To == to
	})
	if i == -1 {
		return nil, fmt.Errorf("%w: map not foud: from = %s, to = %s", NotFoundErr, from, to)
	}
	return a.maps[i], nil
}
func (a *Almanac) GetPath(names ...string) (AlmanacPath, error) {
	var err error
	maps := make([]*Map, len(names)-1)
	for i := 1; i < len(names); i++ {
		prev := i - 1
		maps[prev], err = a.Find(names[prev], names[i])
		if err != nil {
			return nil, err
		}
	}
	return maps, nil
}

var NotFoundErr = errors.New(" not found")

type AlmanacPath []*Map

func (a AlmanacPath) Traverse(x int) int {
	if DebugEnabled {
		sb := strings.Builder{}
		sb.WriteString(strconv.Itoa(x))
		for _, m := range a {
			x = m.Next(x)
			sb.WriteString(" -> ")
			sb.WriteString(strconv.Itoa(x))
		}
		log.Println(sb.String())
		return x
	}

	for _, m := range a {
		x = m.Next(x)
	}
	return x
}
