package com

import (
	"strconv"
	"strings"
)

func ExtractIntSliceFromString(line string) ([]int, error) {
	var result []int
	for _, n := range strings.Fields(line) {
		if number, err := strconv.Atoi(n); err == nil {
			result = append(result, number)
		} else {
			return nil, err
		}
	}
	return result, nil
}

func ExtractIntSliceFromStringFunc(line string, f func(rune) bool) ([]int, error) {
	var result []int
	for _, n := range strings.FieldsFunc(line, f) {
		if number, err := strconv.Atoi(n); err == nil {
			result = append(result, number)
		} else {
			return nil, err
		}
	}
	return result, nil
}
