package solution

import (
	"aoc/com"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	pattern    []Spring
	match      []int
	mInd, pInd int

	hc  *string
	cnt int

	checks *int
}

func (s *Solver) SolveCached(sc int) int {
	key := s.HashKey()
	sol, ok := cache[key]
	if ok {
		hits += 1
		//newSol := s.Solve(sc)
		//if cache[key] != newSol {
		//	//cache[key] = sol
		//	fmt.Println("key:", "(", key, ")", "m: ", s.match[s.mInd:])
		//	fmt.Println("new sol:", newSol, "cache", sol)
		//	fmt.Println("pattern:", string(s.pattern), s.pInd, s.mInd, s.match, *s.checks)
		//	fmt.Println(len(key), len(s.pattern))
		//	panic("failed to match something")
		//} else {
		//	//fmt.Println("--")
		//}
		//panic("cache hit: " + key)
	} else {
		sol = s.Solve(sc)
		cache[key] = sol
	}
	return sol
}

func (s *Solver) HashKey() string {
	if s.hc == nil {
		sb := strings.Builder{}

		//sb.WriteString(string(s.pattern[:s.pInd+1]))
		sb.WriteString(string(s.pattern[s.pInd:]))
		sb.WriteRune('|')
		sb.WriteString(strconv.Itoa(s.mInd))
		sb.WriteRune('|')
		sb.WriteString(strconv.Itoa(s.cnt))
		//for i := s.pInd + 1; i < len(s.pattern); i++ {
		//	sb.WriteRune(rune(s.pattern[i]))
		//}
		//sb.WriteRune('|')
		//for i := s.mInd + 1; i < len(s.match); i++ {
		//	sb.WriteString(strconv.Itoa(s.match[i]))
		//	sb.WriteString(",")
		//}
		//sb.WriteRune('|')
		//sb.WriteString(strconv.Itoa(s.mInd))

		hash := sb.String()
		s.hc = &hash

		if *s.hc == ".||15" {
			fmt.Println("wtf: ", s.cnt, s.pInd, s.mInd, string(s.pattern), "&", string(s.pattern[:s.pInd+1]))
		}
	}
	return *s.hc
}

func (s *Solver) HashKeyOLD() string {
	if s.hc == nil {
		sb := strings.Builder{}
		//sb.WriteString(string(s.pattern[s.pInd]))
		//sb.WriteString(strings.Join(com.IntsToString(s.match[:s.mInd]), ","))
		//sb.WriteString(string(s.pattern))

		//sb.WriteString(strconv.Itoa(s.pInd))
		//sb.WriteRune('|')
		sb.WriteString(strconv.Itoa(s.mInd))
		sb.WriteRune('|')
		//sb.WriteString(strings.Join(com.IntsToString(s.match[:s.mInd]), ","))
		//sb.WriteRune('|')
		//sb.WriteString(string(s.pattern[s.pInd:]), )
		//sb.WriteString()
		//rightMatch := string(s.match[:s.pInd])

		sb.WriteRune('|')
		leftPatterns := string(s.pattern[s.pInd+1:])
		sb.WriteString(leftPatterns)
		sb.WriteRune('|')
		//sb.WriteRune('|')
		//rightPatterns := string(s.pattern[:s.pInd])
		//sb.WriteString(rightPatterns)
		//sb.WriteRune('|')

		leftMatch := strings.Join(com.IntsToString(s.match[s.mInd:]), ",")
		sb.WriteString(leftMatch)
		sb.WriteRune('|')

		rightMatch := strings.Join(com.IntsToString(s.match[:s.mInd]), ",")
		sb.WriteString(rightMatch)
		sb.WriteRune('|')

		//sb.WriteRune('|')
		//sb.WriteString(string(s.pattern))
		//fmt.Println(string(s.pattern))
		hash := sb.String()
		//s.hc = &hash
		//hash := fmt.Sprintf("%d-%d", s.pInd, s.mInd)
		//hash := "asdf"
		s.hc = &hash
	}
	return *s.hc
}

//func hash(s *Solver) {}

func (s *Solver) Clone() *Solver {
	return &Solver{pattern: com.CloneSlice(s.pattern), match: s.match,
		mInd: s.mInd, pInd: s.pInd, cnt: s.cnt, checks: s.checks}
}

func NewSolver(pattern []Spring, match []int) *Solver {
	checks := 1
	return &Solver{pattern: com.CloneSlice(pattern), match: match, checks: &checks}
}

var cache map[string]int
var hits int

func solve(pattern []Spring, match []int) int {
	cache = make(map[string]int)
	s := NewSolver(pattern, match)
	//fmt.Println("            ", string(pattern), match)

	res := s.Solve(1)
	//fmt.Println(string(pattern), match, res, "recursive calls: ", *s.checks, "cache hits: ", hits)

	hits = 0
	cache = make(map[string]int)
	return res
}

func (s *Solver) Solve(sc int) (sol int) {
	*s.checks += 1
	for ; s.pInd < len(s.pattern); s.pInd += 1 {
		if s.pattern[s.pInd] == '?' {
			{
				ns := s.Clone()
				ns.pattern[s.pInd] = '#'

				res := ns.SolveCached(sc + 1)
				sol += res
			}
			{
				ns := s.Clone()
				ns.pattern[s.pInd] = '.'

				res := ns.SolveCached(sc + 1)
				sol += res
			}
			return sol
		}

		isLast := len(s.match) == s.mInd
		if s.pattern[s.pInd] == '#' {
			s.cnt += 1
			if isLast || s.cnt > s.match[s.mInd] {
				return sol
			}
		} else if s.pattern[s.pInd] == '.' && s.cnt > 0 {
			if isLast || s.cnt != s.match[s.mInd] {
				return sol
			}

			s.mInd += 1
			s.cnt = 0
		}
	}

	if s.mInd == len(s.match) {
		return sol + 1
	}

	return sol
}
