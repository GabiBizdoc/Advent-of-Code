package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_03/part1/solution"
	"testing"
)

const solutionShort = 4361
const solutionLong = 543867

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, solutionShort, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart1Long(t *testing.B) {
	testcom.SolveAOC(t, solutionLong, testcom.LongFilepath, solution.SolveChallenge)
}
