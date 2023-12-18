package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_07/part2/solution"
	"testing"
)

func TestSolvePart2Short(t *testing.T) {
	testcom.SolveAOC(t, 5905, testcom.Part2ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart2Long(t *testing.T) {
	testcom.SolveAOC(t, 249666369, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart2Long(t *testing.B) {
	testcom.SolveAOC(t, 76008, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart2Short(t *testing.B) {
	testcom.SolveAOC(t, 46, testcom.Part2ShortFilepath, solution.SolveChallenge)
}
