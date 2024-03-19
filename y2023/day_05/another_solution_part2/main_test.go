package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_05/another_solution_part2/solution"
	"testing"
)

func TestMain(m *testing.M) {
	//solution.DebugEnabled = true
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	m.Run()
}

func TestSolvePart2Short(t *testing.T) {
	testcom.SolveAOC(t, 46, testcom.Part2ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart2Long(t *testing.T) {
	testcom.SolveAOC(t, 24261545, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart2Long(t *testing.B) {
	testcom.SolveAOC(t, 76008, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart2Short(t *testing.B) {
	testcom.SolveAOC(t, 46, testcom.Part2ShortFilepath, solution.SolveChallenge)
}
