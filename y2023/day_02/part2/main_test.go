package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_02/part2/solution"
	"testing"
)

func TestSolvePart2Short(t *testing.T) {
	testcom.SolveAOC(t, 2286, testcom.Part2ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart2Long(t *testing.T) {
	testcom.SolveAOC(t, 76008, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart2Long(t *testing.B) {
	testcom.SolveAOC(t, 76008, testcom.LongFilepath, solution.SolveChallenge)
}
