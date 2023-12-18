package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_02/part1/solution"
	"testing"
)

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, 8, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, 2348, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart1Long(t *testing.B) {
	testcom.SolveAOC(t, 2348, testcom.LongFilepath, solution.SolveChallenge)
}
