package testcom

import (
	"fmt"
	"path"
	"testing"
	"time"
)

type solutionSolver func(inputFilePath string) (int, error)

func SolveAOC(t testing.TB, expected int, filepath string, solver solutionSolver) {
	if path.Ext(filepath) == "" {
		filepath += ".txt"
	}

	start := time.Now()
	result, err := solver(filepath)
	if err != nil {
		t.Error(err)
	}

	if result == expected {
		t.Log(result)
	} else {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
	fmt.Println("execution time: ", time.Since(start))
}

const Part1ShortFilepath = "../part1-short.txt"
const Part2ShortFilepath = "../part2-short.txt"
const TMP = "../tmp.txt"
const LongFilepath = "../input-long.txt"
