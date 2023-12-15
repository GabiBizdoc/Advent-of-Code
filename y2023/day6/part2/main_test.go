package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day6/part2/solution"
	"testing"
)

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, 71503, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, 32583852, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart1Long(t *testing.B) {
	testcom.SolveAOC(t, 1624896, testcom.LongFilepath, solution.SolveChallenge)
}
