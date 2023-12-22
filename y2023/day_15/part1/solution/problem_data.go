package solution

import "strings"

func CustomHash(s string) (solution int) {
	for _, c := range s {
		solution += int(c)
		solution *= 17
		solution %= 256
	}
	return solution
}

func SolveLine(line string) (solution int) {
	for _, str := range strings.Split(line, ",") {
		solution += CustomHash(str)
	}
	return solution
}
