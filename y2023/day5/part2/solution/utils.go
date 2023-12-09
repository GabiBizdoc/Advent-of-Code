package solution

import (
	"aoc/com"
	. "aoc/y2023/day5/part2/solution/almanac"
	"aoc/y2023/day5/part2/solution/mapper"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Solution struct {
	value int
	mu    sync.Mutex
}

func NewSolution() *Solution {
	return &Solution{value: math.MaxInt}
}

func (s *Solution) UnsafeUpdate(value int) {
	if value == math.MaxInt {
		panic("value was max int")
	}
	if s.value > value {
		s.value = value
	}
}

func (s *Solution) Value() int {
	return s.value
}

func (s *Solution) SafeUpdate(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.UnsafeUpdate(value)
}

func readAlmanac(inputFilePath string) (*Almanac, error) {
	almanac := &Almanac{}

	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var m *mapper.Mapper
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "seeds: "):
			line = strings.TrimPrefix(line, "seeds: ")
			seeds := strings.Split(line, " ")
			for _, seed := range seeds {
				seedValue, err := strconv.Atoi(seed)
				if err != nil {
					return nil, err
				}
				almanac.Seeds = append(almanac.Seeds, seedValue)
			}

		case strings.HasSuffix(line, "map:"):
			if m != nil {
				if err := almanac.AddMapper(m); err != nil {
					return nil, err
				}
			}
			line = strings.TrimSuffix(line, "map:")
			parts := strings.Split(line, "-to-")
			source := strings.TrimSpace(parts[0])
			dest := strings.TrimSpace(parts[1])
			m = mapper.NewMapper(source, dest)

		case line == "":
			continue
		default:
			parts := strings.Split(line, " ")
			if len(parts) != 3 {
				fmt.Println(line)
				return nil, fmt.Errorf("invalid number of arguments: expected=3  got=%d", len(parts))
			}

			ints, err := com.StringsToInts(parts)
			if err != nil {
				return nil, err
			}

			m.AddRule(ints[1], ints[0], ints[2])
		}
	}

	if m != nil {
		if err := almanac.AddMapper(m); err != nil {
			return nil, err
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return almanac, nil
}
