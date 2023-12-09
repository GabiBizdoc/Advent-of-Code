package main_test

import (
	testcom "aoc/y2023/com/test_com"
	"aoc/y2023/day8/part2/solution"
	"testing"
)

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, 6, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, 14935034899483, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart1Long(t *testing.B) {
	testcom.SolveAOC(t, 14935034899483, testcom.LongFilepath, solution.SolveChallenge)
}
