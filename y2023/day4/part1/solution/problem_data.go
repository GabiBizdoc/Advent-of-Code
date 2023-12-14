package solution

import "slices"

type RowData struct {
	GameID           int
	WinningNumbers   []int
	ExtractedNumbers []int
}

func (r *RowData) FindCardValue() int {
	var value int
	for _, number := range r.WinningNumbers {
		if slices.Contains(r.ExtractedNumbers, number) {
			if value == 0 {
				value = 1
			} else {
				value *= 2
			}
		}
	}
	return value
}
