package solution

func transpose(grid [][]rune) [][]rune {
	transposedGrid := make([][]rune, len(grid[0]))
	for _, line := range grid {
		for i, c := range line {
			transposedGrid[i] = append(transposedGrid[i], c)
		}
	}
	return transposedGrid
}
