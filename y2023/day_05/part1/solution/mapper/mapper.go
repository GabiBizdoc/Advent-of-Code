package mapper

type Mapper struct {
	Source      string
	Destination string

	rules []*Rule
}

func (m *Mapper) addRule(rule *Rule) {
	m.rules = append(m.rules, rule)
}

func (m *Mapper) AddRule(sourceStart, destStart, size int) {
	m.addRule(NewMapperRule(sourceStart, destStart, size))
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
