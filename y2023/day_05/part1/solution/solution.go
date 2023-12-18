package solution

import (
	"aoc/com"
	"aoc/y2023/day_05/part1/solution/mapper"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Seeds   []int
	mappers map[string]map[string]*mapper.Mapper
}

func (d *Data) Get(source, dest string, value int) (int, error) {
	m, ok := d.mappers[source][dest]
	if !ok {
		return 0, fmt.Errorf("mapper not found for m[%s][%s]", source, dest)
	}

	next := m.GetValue(value)
	//fmt.Printf("%-20s:%d \t %-20s:%d\n", source, value, dest, next)

	return next, nil
}

func (d *Data) GetPath(value int, path ...string) (int, error) {
	start, tail := path[0], path[1:]
	fmt.Println(value, start, tail)
	var err error

	//var paths = ""
	//paths += start + ": " + strconv.Itoa(value)

	for _, next := range tail {
		value, err = d.Get(start, next, value)
		if err != nil {
			return 0, err
		}
		//paths += "\t" + next + ": " + strconv.Itoa(value)
		start = next
	}

	//fmt.Println(paths + "\n\n")
	return value, nil
}

func (d *Data) AddMapper(m *mapper.Mapper) error {
	if d.mappers == nil {
		d.mappers = make(map[string]map[string]*mapper.Mapper)
	}

	if d.mappers[m.Source] == nil {
		d.mappers[m.Source] = make(map[string]*mapper.Mapper)
	}

	if _, ok := d.mappers[m.Source][m.Destination]; ok {
		return fmt.Errorf("attempt to override mappers")
	}

	d.mappers[m.Source][m.Destination] = m
	return nil
}

func solveChallenge(inputFilePath string) (int, error) {
	var solution = math.MaxInt

	data, err := readData(inputFilePath)
	if err != nil {
		return 0, err
	}

	for _, seed := range data.Seeds {
		next, err := data.GetPath(seed,
			"seed", "soil", "fertilizer", "water",
			"light", "temperature", "humidity", "location")
		if err != nil {
			return 0, err
		}
		solution = min(solution, next)
	}

	fmt.Println(solution)
	return solution, nil
}

func readData(inputFilePath string) (*Data, error) {
	data := &Data{}

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
		//fmt.Println(line)
		switch {
		case strings.HasPrefix(line, "seeds: "):
			line = strings.TrimPrefix(line, "seeds: ")
			seeds := strings.Split(line, " ")
			for _, seed := range seeds {
				seedValue, err := strconv.Atoi(seed)
				if err != nil {
					return nil, err
				}
				data.Seeds = append(data.Seeds, seedValue)
			}

			//fmt.Println("line]= ", line)
			//fmt.Println("read seeds", data.Seeds)
		case strings.HasSuffix(line, "map:"):
			if m != nil {
				if err := data.AddMapper(m); err != nil {
					return nil, err
				}
			}
			line = strings.TrimSuffix(line, "map:")
			parts := strings.Split(line, "-to-")
			source := strings.TrimSpace(parts[0])
			dest := strings.TrimSpace(parts[1])
			m = mapper.NewMapper(source, dest)

			//fmt.Println("line]= ", line)
			//fmt.Println("read map", m)
		case line == "":
			continue
			// skip
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

			// source, destination, range
			m.AddRule(ints[1], ints[0], ints[2])

			//fmt.Println("line]= ", line)
			//fmt.Println("read rule", ints)
		}
	}
	if m != nil {
		if err := data.AddMapper(m); err != nil {
			return nil, err
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
