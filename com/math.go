package com

func Abs[T ~int](a T) T {
	if a < 0 {
		return a * -1
	}
	return a
}
