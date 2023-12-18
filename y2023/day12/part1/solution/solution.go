package solution

import (
	"aoc/com"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var globalMusk []int

type PatterBuilder struct {
	springs []Spring
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

//var invalidSprings map[string]struct{}
//
//func init() {
//	invalidSprings = make(map[string]struct{})
//}

func solveFunction(pattern []Spring, musk []int) int {
	springs := make([]*PatterBuilder, 0)
	springs = append(springs, NewPatterBuilder(musk))

	for _, spring := range pattern {
		if spring.IsUnknown() {
			next := make([]*PatterBuilder, 0)
			for _, builder := range springs {
				b2 := builder.Clone()
				if builder.Write('.') {
					next = append(next, builder)
				}
				if b2.Write('#') {
					next = append(next, b2)
				}
			}
			springs = next
		} else {
			next := make([]*PatterBuilder, 0)
			for _, builder := range springs {
				if builder.Write(spring) {
					next = append(next, builder)
				}
			}
			springs = next
		}

		//fmt.Println("   ", string(pattern), musk)
		//for i, builder := range springs {
		//	fmt.Println(fmt.Sprintf("%3d", i), string(builder.springs), builder.cnt)
		//}
		//fmt.Println()
	}

	solution := 0
	for _, spring := range springs {
		if spring.IsSolution() {
			solution += 1
		}
	}

	//fmt.Println("   ", string(pattern), musk)
	//for i, builder := range springs {
	//	fmt.Println(fmt.Sprintf("%3d", i), string(builder.springs), builder.cnt, builder.IsSolution())
	//}
	//fmt.Println()
	return solution

	//return len(springs)
}

func findArrangements(pattern []Spring, musk []int) int {
	globalMusk = musk
	//fmt.Println(string(pattern), musk)
	return solveFunction(pattern, musk)
}

func unpackLine(line string) ([]Spring, []int, error) {
	fields := strings.Fields(line)
	ints, err := com.ExtractIntSliceFromStringFunc(fields[1], func(r rune) bool {
		return r == ','
	})
	springs := []Spring(fields[0])
	springs = append(springs, '.')
	return springs, ints, err
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
		res := findArrangements(pattern, match)
		fmt.Println(pattern, match, res)
		solution += res
		//return 0, err
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
