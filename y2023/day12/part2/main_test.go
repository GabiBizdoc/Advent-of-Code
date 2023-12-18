package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day12/part2/solution"
	"testing"
)

const solutionShort = 525152
const solutionLong = 1909291258644

func TestSolveShort1(t *testing.T) {
	testcom.SolveAOC(t, solutionShort, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolveShort2(t *testing.T) {
	testcom.SolveAOC(t, 506250, testcom.Part2ShortFilepath, solution.SolveChallenge)
}

func TestSolveTmp(t *testing.T) {
	testcom.SolveAOC(t, 6537520, testcom.TMP, solution.SolveChallenge)
}

func TestSolveLong(t *testing.T) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkLong(t *testing.B) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}
