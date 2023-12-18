package solution

import "slices"

type RowData struct {
	GameID           int
	WinningNumbers   []int
	ExtractedNumbers []int
}

func (r *RowData) FindMatchings() int {
	var matching int
	for _, number := range r.WinningNumbers {
		if slices.Contains(r.ExtractedNumbers, number) {
			matching += 1
		}
	}
	return matching
}
