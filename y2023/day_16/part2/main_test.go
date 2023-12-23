package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_16/part2/solution"
	"testing"
)

const solutionShort = 51
const solutionLong = 8318

func TestSolveShort(t *testing.T) {
	testcom.SolveAOC(t, solutionShort, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolveLong(t *testing.T) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkLong(t *testing.B) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}
