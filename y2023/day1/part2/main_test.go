package main_test

import (
	testcom "aoc/y2023/com/test_com"
	"aoc/y2023/day1/part2/solution"
	"testing"
)

func TestSolveP2Short(t *testing.T) {
	testcom.SolveAOC(t, 281, testcom.Part2ShortFilepath, solution.SolveChallenge)
}

func TestSolveP2Long(t *testing.T) {
	testcom.SolveAOC(t, 54078, testcom.LongFilepath, solution.SolveChallenge)
}
