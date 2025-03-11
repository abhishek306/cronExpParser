package main

import "testing"

func TestFieldParser(t *testing.T) {
	tests := []struct {
		input      string
		fieldRange []int
		expected   []int
	}{
		{"*", minuteRange, generateRange(0, 59)},
		{"*/15", minuteRange, []int{0, 15, 30, 45}},
		{"1-5", hourRange, []int{1, 2, 3, 4, 5}},
		{"1,15", dayOfMonthRange, []int{1, 15}},
		{"5", monthRange, []int{5}},
	}

	for _, test := range tests {
		result := fieldParser(test.input, test.fieldRange)
		if !equalSlices(result, test.expected) {
			t.Errorf("parseField(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func generateRange(start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
