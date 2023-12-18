package solution

type CosmicEntity rune

func (t CosmicEntity) IsEmpty() bool {
	return t == '.'
}

func (t CosmicEntity) IsGalaxy() bool {
	return t == '#'
}
