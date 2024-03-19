package solution

import (
	"aoc/com"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var DebugEnabled bool

func Solve(file io.Reader) (int, error) {
	solution := math.MaxInt

	seeds, a, err := readAlmanac(file)
	almanacPath := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	if err != nil {
		return 0, err
	}
	maps, err := a.GetPath(almanacPath...)
	if err != nil {
		return 0, fmt.Errorf("invalid path: %w", err)
	}
	for _, seed := range seeds {
		location := maps.Traverse(seed)
		if location < solution {
			solution = location
		}
	}
	return solution, nil
}

func readAlmanac(file io.Reader) (seeds []int, a *Almanac, err error) {
	scanner := bufio.NewScanner(file)
	seeds = make([]int, 0, 4)
	a = NewAlmanac()

	var am *Map
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			line = strings.TrimPrefix(line, "seeds:")
			for _, seed := range strings.Fields(line) {
				value, err := strconv.Atoi(seed)
				if err != nil {
					return seeds, a, err
				}
				seeds = append(seeds, value)
			}
			continue
		}

		if strings.HasSuffix(line, "map:") {
			if am != nil {
				err := a.AddMap(am)
				if err != nil {
					return seeds, a, err
				}
			}
			info := strings.Split(strings.Fields(line)[0], "-")
			am = NewMap(info[0], info[2])
			continue
		}

		row, err := com.StringsToInts(strings.Fields(line))
		if err != nil {
			return seeds, a, err
		}

		am.AddRow(row[1], row[0], row[2])
	}
	if err := scanner.Err(); err != nil {
		return seeds, a, err
	}
	err = a.AddMap(am)
	if err != nil {
		return seeds, a, err
	}

	return seeds, a, err
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return Solve(file)
}
