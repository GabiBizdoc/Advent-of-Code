package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_12/part2/solution"
	"testing"
)

const solutionShort = 82000210
const solutionLong = 649862989626

func TestSolveShort1(t *testing.T) {
	testcom.SolveAOC(t, solutionShort, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolveLong(t *testing.T) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkLong(t *testing.B) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}
