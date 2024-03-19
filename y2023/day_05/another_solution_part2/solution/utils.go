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

func (r *Row) DestinationStart() int {
	return r[1]
}
func (r *Row) DestinationEnd() int {
	return r.DestinationStart() + r.Range()
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

func NewMap(from string, to string) *Map {
	return &Map{From: from, To: to, Rows: make([]Row, 0)}
}
func (m *Map) next(x int) int {
	for _, row := range m.Rows {
		if x >= row.SourceStart() && x < row.SourceEnd() {
			next := x - row.SourceStart() + row.DestinationStart()
			if DebugEnabled {
				log.Println(
					m.From, x, "to", m.To,
					"start", row.SourceStart(),
					"destination", row.DestinationStart(),
					"range", row.Range(),
					"next", next,
				)
			}
			return next
		}
	}
	return x
}
func (m *Map) nextInterval(iStart, iEnd int) (response [][2]int) {
	response = make([][2]int, 0)
	var closest *Row = nil
	for _, row := range m.Rows {
		if row.SourceStart() > iStart && (closest == nil || closest.SourceStart() > row.SourceStart()) {
			closest = &row
		}

		if iStart >= row.SourceStart() && iStart <= row.SourceEnd() {
			next := iStart - row.SourceStart() + row.DestinationStart()
			if iEnd <= row.SourceEnd() {
				response = append(response, [2]int{next, next + iEnd - iStart})
				return response
			}
			response = append(response, [2]int{next, row.DestinationEnd()})
			response = append(response, m.nextInterval(row.SourceEnd()+1, iEnd)...)
			return response
		}
	}

	// we passed all the intervals
	if closest != nil && iEnd > closest.SourceStart() {
		response = append(response, [2]int{iStart, closest.SourceStart() - 1})
		response = append(response, m.nextInterval(closest.SourceStart(), iEnd)...)
		return response
	}

	// we passed all the intervals
	response = append(response, [2]int{iStart, iEnd})
	return response
}
func (m *Map) AddRow(s, d, r int) {
	r = r - 1
	m.Rows = append(m.Rows, Row{s, d, r})
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

type AlmanacPath []*Map

func (a AlmanacPath) Traverse(x int) int {
	if DebugEnabled {
		sb := strings.Builder{}
		sb.WriteString(strconv.Itoa(x))
		for _, m := range a {
			x = m.next(x)
			sb.WriteString(" -> ")
			sb.WriteString(strconv.Itoa(x))
		}
		log.Println(sb.String())
		return x
	}

	for _, m := range a {
		x = m.next(x)
	}
	return x
}
func (a AlmanacPath) TraverseInterval(x1, x2 int) [][2]int {
	initial := [][2]int{{x1, x2}}
	for _, m := range a {
		nextIntervals := make([][2]int, 0)
		for _, next := range initial {
			intervals := m.nextInterval(next[0], next[1])
			nextIntervals = append(nextIntervals, intervals...)
		}
		if DebugEnabled {
			log.Println(initial, m.From, " -> ", m.To, nextIntervals)
		}
		initial = nextIntervals
	}
	return initial
}

var NotFoundErr = errors.New(" not found")
