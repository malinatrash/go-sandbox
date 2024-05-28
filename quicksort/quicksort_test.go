package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{input: []int{10, 7, 8, 9, 1, 5}, output: []int{1, 5, 7, 8, 9, 10}},
		{input: []int{4, 3, 2, 10, 12, 1, 5, 6}, output: []int{1, 2, 3, 4, 5, 6, 10, 12}},
		{input: []int{3, 7, 8, 5, 2, 1, 9, 5, 4}, output: []int{1, 2, 3, 4, 5, 5, 7, 8, 9}},
		{input: []int{1}, output: []int{1}},
		{input: []int{}, output: []int{}},
		{input: []int{2, 1}, output: []int{1, 2}},
	}

	for _, test := range tests {
		got := QuickSort(test.input)
		want := test.output
		if !reflect.DeepEqual(got, want) {
			t.Errorf("QuickSort(%v) = %v; want %v", test.input, got, test.output)
		}
	}
}
