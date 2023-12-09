package main_test

import (
	testcom "aoc/y2023/com/test_com"
	"aoc/y2023/day1/part1/solution"
	"testing"
)

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, 142, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, 54601, testcom.LongFilepath, solution.SolveChallenge)
}
