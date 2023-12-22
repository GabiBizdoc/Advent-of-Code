package solution

import "testing"

func TestCustomHash(t *testing.T) {
	testCases := map[string]int{
		"HASH": 52,
	}

	for input, expected := range testCases {
		result := CustomHash(input)
		if result != expected {
			t.Errorf("CustomHash(%s) = %d, expected %d", input, result, expected)
		}
	}
}
