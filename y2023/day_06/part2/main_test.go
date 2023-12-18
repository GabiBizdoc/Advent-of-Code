package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_06/part2/solution"
	"testing"
)

func TestSolveShort(t *testing.T) {
	testcom.SolveAOC(t, 71503, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolveLong(t *testing.T) {
	testcom.SolveAOC(t, 32583852, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolveLong(t *testing.B) {
	testcom.SolveAOC(t, 32583852, testcom.LongFilepath, solution.SolveChallenge)
}
