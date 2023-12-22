package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_14/part2_brute_force/solution"
	"testing"
)

const solutionShort = 64
const solutionLong = 90795

func TestSolveShort(t *testing.T) {
	testcom.SolveAOC(t, solutionShort, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolveLong(t *testing.T) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkLong(t *testing.B) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}
