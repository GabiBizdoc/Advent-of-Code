package solution

import (
	"aoc/com"
	"strconv"
	"strings"
)

type Solver struct {
	cache                 map[string]int
	springs               []rune
	pattern               []int
	springInd, patternInd int
	cnt                   int
}

func (s *Solver) SolveCached() int {
	key := s.Hash()
	sol, ok := s.cache[key]
	if !ok {
		sol = s.Solve()
		s.cache[key] = sol
	}
	return sol
}

func (s *Solver) Hash() string {
	sb := strings.Builder{}
	sb.WriteString(string(s.springs[s.patternInd:]))
	sb.WriteRune('|')
	sb.WriteString(strconv.Itoa(s.springInd))
	sb.WriteRune('|')
	sb.WriteString(strconv.Itoa(s.cnt))
	return sb.String()
}

func (s *Solver) Clone() *Solver {
	return &Solver{
		springs:    com.CloneSlice(s.springs),
		springInd:  s.springInd,
		patternInd: s.patternInd,
		cnt:        s.cnt,

		// readonly data!
		cache:   s.cache,
		pattern: s.pattern,
	}
}

func (s *Solver) Solve() (sol int) {
	for ; s.patternInd < len(s.springs); s.patternInd += 1 {
		if s.springs[s.patternInd] == '?' {
			for _, x := range ".#" {
				next := s.Clone()
				next.springs[s.patternInd] = x

				res := next.SolveCached()
				sol += res
			}
			return sol
		}

		isLast := len(s.pattern) == s.springInd
		if s.springs[s.patternInd] == '#' {
			s.cnt += 1
			if isLast || s.cnt > s.pattern[s.springInd] {
				return sol
			}
		} else if s.springs[s.patternInd] == '.' && s.cnt > 0 {
			if isLast || s.cnt != s.pattern[s.springInd] {
				return sol
			}

			s.springInd += 1
			s.cnt = 0
		}
	}

	if s.springInd == len(s.pattern) {
		return sol + 1
	}

	return sol
}

func NewSolver(pattern []rune, match []int) *Solver {
	cache := make(map[string]int)
	return &Solver{springs: com.CloneSlice(pattern), pattern: match, cache: cache}
}

func solveLine(pattern []rune, match []int) int {
	return NewSolver(pattern, match).Solve()
}
