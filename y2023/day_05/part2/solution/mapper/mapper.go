package mapper

import "sort"

type Mapper struct {
	Source      string
	Destination string

	rules []*Rule
}

func (m *Mapper) addRule(rule *Rule) {
	m.rules = append(m.rules, rule)

	sort.Slice(m.rules, func(i, j int) bool {
		return m.rules[i].sourceStart < m.rules[j].sourceStart
	})

	// check consistency
	for i := 1; i <= len(m.rules); i++ {
		first, second := m.rules[i-1], m.rules[i]
		if first.sourceStart+first.size > second.sourceStart {
			panic("rules overlap!")
		}
	}
}

func (m *Mapper) AddRule(sourceStart, destStart, size int) {
	m.addRule(NewMapperRule(sourceStart, destStart, size))
}

func (m *Mapper) UpdateMissingRules() {
	start := 0
	for _, rule := range m.rules {
		if rule.sourceStart > start {
			m.AddRule(start, start, rule.sourceStart-start)
		}
		start = rule.sourceStart + rule.size
	}
}

// FindInterval end == 0 means the whole interval
func (m *Mapper) FindInterval(x int) *Rule {

	var start *int

	for _, rule := range m.rules {
		if _, ok := rule.Read(x); ok {
			return rule
		}
		if start != nil && rule.sourceStart >= x {

		}
	}

	//start := 0

	//if len(m.rules) > 0 {
	//
	//	//end = m.rules[0].sourceStart
	//}

	//for _, rule := range m.rules {
	//	if rule.Contains(x) {
	//		return rule
	//	}
	//}
	return nil
}

func (m *Mapper) GetValue(x int) int {
	for _, rule := range m.rules {
		if value, ok := rule.Read(x); ok {
			return value
		}
	}
	return x
}

func NewMapper(source string, destination string) *Mapper {
	rules := make([]*Rule, 0)
	return &Mapper{Source: source, Destination: destination, rules: rules}
}
