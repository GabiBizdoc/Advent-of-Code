package solution

func CalculateLoad(grid [][]rune) (solution int) {
	rocksCount := 0
	for j := 0; j < len(grid[0]); j++ {
		for i := len(grid) - 1; i >= 0; i-- {
			switch grid[i][j] {
			case '.':
				continue
			case 'O':
				rocksCount += 1
			case '#':
				tmp := rocksCount * (rocksCount + 1) / 2
				solution += (len(grid)-i)*rocksCount - tmp
				rocksCount = 0
			}
		}
		tmp := rocksCount * (rocksCount - 1) / 2
		solution += len(grid)*rocksCount - tmp
		rocksCount = 0
	}
	return solution
}
