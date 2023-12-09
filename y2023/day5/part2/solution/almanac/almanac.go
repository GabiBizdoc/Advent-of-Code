package almanac

import (
	"aoc/y2023/day5/part2/solution/mapper"
	"fmt"
)

type Almanac struct {
	Seeds   []int
	mappers map[string]map[string]*mapper.Mapper
}

func (d *Almanac) Get(source, dest string, value int) (int, error) {
	m, ok := d.mappers[source][dest]
	if !ok {
		return 0, fmt.Errorf("mapper not found for m[%s][%s]", source, dest)
	}

	next := m.GetValue(value)
	return next, nil
}

func (d *Almanac) GetPath(value int, path ...string) (int, error) {
	source, tail := path[0], path[1:]
	var err error
	for _, destination := range tail {
		value, err = d.Get(source, destination, value)
		if err != nil {
			return 0, err
		}
		source = destination
	}
	return value, nil
}

func (d *Almanac) AddMapper(m *mapper.Mapper) error {
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
