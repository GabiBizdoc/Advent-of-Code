package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_04/part2/solution"
	"testing"
)

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, 30, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, 10212704, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart1Long(t *testing.B) {
	testcom.SolveAOC(t, 261668924, testcom.LongFilepath, solution.SolveChallenge)
}
