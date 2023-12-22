package solution

import (
	"strconv"
	"strings"
)

func CustomHash(s string) (solution int) {
	for _, c := range s {
		solution += int(c)
		solution *= 17
		solution %= 256
	}
	return solution
}

func SolveLine(line string) (solution int) {
	h := NewHashMap()
	for _, str := range strings.Split(line, ",") {
		var sb strings.Builder
		for i, c := range str {
			if c == '=' {
				if v, err := strconv.Atoi(str[i+1:]); err != nil {
					panic(err)
				} else {
					h.Set(sb.String(), v)
				}
			} else if c == '-' {
				h.Remove(sb.String())
			}
			sb.WriteRune(c)
		}
	}

	for i, datum := range h.data {
		for j, value := range datum {
			boxNum := i + 1
			slotNum := j + 1

			power := boxNum * slotNum * value.FocalLen
			//fmt.Println(value.Label, boxNum, slotNum, value.FocalLen, power)
			solution += power
		}
	}
	return solution
}
