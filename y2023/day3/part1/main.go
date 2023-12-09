package main

import (
	"aoc/y2023/day3/part1/solution"
	"flag"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	inputFile := flag.String("file", "", "Path to the input file")
	flag.Parse()

	result, err := solution.SolveChallenge(*inputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(result, time.Since(start))
}
