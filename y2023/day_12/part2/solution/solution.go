package solution

import (
	"aoc/com"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

func unpackLine(line string) ([]rune, []int, error) {
	fields := strings.Fields(line)
	ints, err := com.ExtractIntSliceFromStringFunc(fields[1], func(r rune) bool {
		return r == ','
	})

	// simplify
	springs := make([]rune, 0, len(fields[0]))
	var prev rune = ' '
	for _, x := range fields[0] {
		if x == '.' && prev == '.' {
			continue
		}

		springs = append(springs, rune(x))
		prev = x
	}

	// clone
	springs2 := make([]rune, 0, len(springs)*5+5)
	ints2 := make([]int, 0, len(ints)*5)

	for i := 0; i < 5; i++ {
		springs2 = append(springs2, springs...)
		springs2 = append(springs2, '?')
		ints2 = append(ints2, ints...)
	}

	// adding an extra dot to the end to simplify recursion
	springs2[len(springs2)-1] = '.'
	return springs2, ints2, err
}

func Solve(file io.Reader) (solution int, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		pattern, match, err := unpackLine(line)
		if err != nil {
			return 0, err
		}

		solution += solveLine(pattern, match)
	}

	if err := scanner.Err(); err != nil {
		return solution, err
	}

	return solution, nil
}

func worker(wg *sync.WaitGroup, lineChan <-chan string, solution *atomic.Int64) {
	defer wg.Done()

	for line := range lineChan {
		springs, pattern, err := unpackLine(line)
		if err != nil {
			panic(err)
		}
		sol := solveLine(springs, pattern)
		solution.Add(int64(sol))
	}
}

func Solve2(file io.Reader) (solution int, err error) {
	scanner := bufio.NewScanner(file)

	var maxWorkers = 30
	var atomicSolution atomic.Int64
	var wg sync.WaitGroup
	lineChan := make(chan string, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(&wg, lineChan, &atomicSolution)
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lineChan <- line
	}
	close(lineChan)

	if err := scanner.Err(); err != nil {
		return solution, err
	}
	wg.Wait()

	solution = int(atomicSolution.Load())
	return solution, nil
}

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Solve2(file)
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
