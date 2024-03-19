package main_test

import (
	testcom "aoc/com/test_com"
	"aoc/y2023/day_05/another_solution/solution"
	"testing"
)

func TestMain(m *testing.M) {
	solution.DebugEnabled = true
	m.Run()
}

func TestSolvePart1Short(t *testing.T) {
	testcom.SolveAOC(t, 35, testcom.Part1ShortFilepath, solution.SolveChallenge)
}

func TestSolvePart1Long(t *testing.T) {
	testcom.SolveAOC(t, 261668924, testcom.LongFilepath, solution.SolveChallenge)
}

func BenchmarkSolvePart1Long(t *testing.B) {
	testcom.SolveAOC(t, 261668924, testcom.LongFilepath, solution.SolveChallenge)
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
