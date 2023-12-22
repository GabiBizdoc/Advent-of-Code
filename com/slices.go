package com

import (
	"fmt"
	"strconv"
	"strings"
)

func CloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

func CloneGrid[T any](a [][]T) [][]T {
	b := make([][]T, len(a))
	for i := range b {
		b[i] = CloneSlice(a[i])
	}
	return b
}

func CloneReverseSlice[T any](s []T) []T {
	b := make([]T, len(s))

	for i, t := range s {
		b[len(s)-i-1] = t
	}

	return b
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CloneReverseString(s string) string {
	return ReverseString(strings.Clone(s))
}

func StringsToInts(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))

	for i, str := range stringSlice {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, fmt.Errorf("error converting string to int: %v", err)
		}
		intSlice[i] = num
	}

	return intSlice, nil
}

func IntsToString(intSlice []int) []string {
	strSlice := make([]string, len(intSlice))
	for i, str := range intSlice {
		num := strconv.Itoa(str)
		strSlice[i] = num
	}
	return strSlice
}

func MapSlice[T, K any](s []T, pred func(x T) K) []K {
	next := make([]K, len(s))
	for i, t := range s {
		next[i] = pred(t)
	}
	return next
}
