package main

import "fmt"

type Stats struct {
	Min, Max, Sum, Cnt float64
	Name               string
}

func NewStats(v float64) *Stats {
	return &Stats{
		Max: v,
		Min: v,
		Sum: v,
		Cnt: 1,
	}
}

func (s *Stats) Mean() float64 {
	return s.Sum / s.Cnt
}

func (s *Stats) SetName(name string) {
	s.Name = name
}

func (s *Stats) Fmt() string {
	return fmt.Sprintf(`%s=%.1f/%.1f/%.1f`, s.Name, s.Min, s.Mean(), s.Max)
}

func (s *Stats) Update(v float64) {
	s.Cnt += 1
	s.Sum += v
	if s.Max < v {
		s.Max = v
	}
	if s.Min > v {
		s.Min = v
	}
}

func (s *Stats) Combine(o *Stats) {
	s.Cnt += o.Cnt
	s.Sum += o.Sum
	if s.Max < o.Max {
		s.Max = o.Max
	}
	if s.Min > o.Min {
		s.Min = o.Min
	}
}
