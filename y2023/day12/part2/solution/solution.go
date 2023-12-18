package solution

import (
	"aoc/com"
	"bufio"
	"cmp"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var globalMusk []int

type PatterBuilder struct {
	springs []Spring
	str     string
	musk    []int
	m       int

	cnt int
}

func NewPatterBuilder(musk []int) *PatterBuilder {
	return &PatterBuilder{musk: musk}
}

func (b *PatterBuilder) Clone() *PatterBuilder {
	return &PatterBuilder{
		springs: com.CloneSlice(b.springs),
		musk:    com.CloneSlice(b.musk),
		m:       b.m,
		cnt:     b.cnt,
	}
}

func (b *PatterBuilder) IsSolution() bool {
	return b.m == len(b.musk)
}

func (b *PatterBuilder) Prefix() string {
	return string(b.springs)
}

func (b *PatterBuilder) Write(s Spring) bool {
	defer func() {
		b.str = string(b.springs)
	}()

	b.springs = append(b.springs, s)
	switch {
	case s.IsDamaged():
		b.cnt += 1
		if b.m == len(b.musk) || b.cnt > b.musk[b.m] {
			return false
		}
	case s.IsOperational():
		if b.cnt == 0 {
			return true // we don't care!
		}
		if b.m < len(b.musk) && b.cnt == b.musk[b.m] {
			b.m += 1
			b.cnt = 0
		} else {
			return false
		}
	}

	return true
}

var invalidSprings map[string]int

func init() {
	invalidSprings = make(map[string]int)
}
func solveFunction22(pattern []Spring, musk []int) int {
	springs := make([]*PatterBuilder, 0)
	springs = append(springs, NewPatterBuilder(musk))

	for _, spring := range pattern {
		if spring.IsUnknown() {
			next := make([]*PatterBuilder, 0)
			for _, builder := range springs {
				b2 := builder.Clone()
				if builder.Write('.') {
					next = append(next, builder)
					//if invalidSprings[builder.Prefix()] == 0 {
					//	next = append(next, builder)
					//	invalidSprings[builder.Prefix()] += 1
					//}
				}
				if b2.Write('#') {
					next = append(next, b2)
					//if invalidSprings[b2.Prefix()] == 0 {
					//	next = append(next, b2)
					//	invalidSprings[b2.Prefix()] += 1
					//}
				}
			}
			springs = next
		} else {
			next := make([]*PatterBuilder, 0)
			for _, builder := range springs {
				if builder.Write(spring) {
					next = append(next, builder)
					//if invalidSprings[builder.Prefix()] == 0 {
					//	next = append(next, builder)
					//	invalidSprings[builder.Prefix()] += 1
					//}
				}
			}
			springs = next
		}

		sort.Slice(springs, func(i, j int) bool {
			return cmp.Compare(springs[i].str, springs[j].str) > 0
		})

		fmt.Println(" ")
		for i, builder := range springs {
			//fmt.Println(i, string(builder.springs))
			fmt.Println(i, builder.str)
		}
		fmt.Println(" ")

	}

	solution := 0
	for _, spring := range springs {
		if spring.IsSolution() {
			solution += 1
			//fmt.Println(string(spring.springs))
		}
	}

	return solution
}

func unpackLine(line string) ([]Spring, []int, error) {
	fields := strings.Fields(line)
	ints, err := com.ExtractIntSliceFromStringFunc(fields[1], func(r rune) bool {
		return r == ','
	})

	// simplify
	//springs := make([]Spring, 0)
	//var prev rune = ' '
	//for _, x := range fields[0] {
	//	prev = x
	//	if x != '.' && prev != '.' {
	//		springs = append(springs, Spring(x))
	//	}
	//}
	//fmt.Println("==-", springs)
	springs := []Spring(fields[0])

	// clone
	springs2 := make([]Spring, 0, len(springs)*5+5)
	ints2 := make([]int, 0, len(ints)*5)

	for i := 0; i < 5; i++ {
		springs2 = append(springs2, springs...)
		springs2 = append(springs2, '?')
		ints2 = append(ints2, ints...)
	}
	springs2[len(springs2)-1] = '.'
	//springs2 = append(springs2, springs...)
	//springs2 = append(springs, '.')
	return springs2, ints2, err
}

func Solve(file io.Reader) (solution int, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		pattern, match, err := unpackLine(line)
		if err != nil {
			return 0, err
		}

		solution += solve(pattern, match)
	}

	if err := scanner.Err(); err != nil {
		return solution, err
	}

	return solution, nil
}

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Solve(file)
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
