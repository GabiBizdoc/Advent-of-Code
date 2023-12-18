package solution

type Spring rune

func (i Spring) String() string {
	return string(i)
}

func (i Spring) IsDamaged() bool {
	return i == '#'
}

func (i Spring) IsOperational() bool {
	return i == '.'
}

func (i Spring) IsUnknown() bool {
	return i == '.'
}
