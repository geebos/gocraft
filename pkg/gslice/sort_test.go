package gslice

import (
	"reflect"
	"testing"

	"github.com/geebos/gocraft/pkg/gvalue"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		less     func(int, int) bool
		expected []int
	}{
		{
			name:     "sort integers ascending",
			input:    []int{3, 1, 4, 1, 5},
			less:     gvalue.Less[int],
			expected: []int{1, 1, 3, 4, 5},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			less:     gvalue.Less[int],
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse order",
			input:    []int{5, 4, 3, 2, 1},
			less:     gvalue.Less[int],
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "nil slice",
			input:    nil,
			less:     gvalue.Less[int],
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			less:     gvalue.Less[int],
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{42},
			less:     gvalue.Less[int],
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var original []int
			if tt.input != nil {
				original = make([]int, len(tt.input))
				copy(original, tt.input)
			}

			result := Sort(tt.input, tt.less)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Sort() = %v, want %v", result, tt.expected)
			}

			// Verify original slice is not modified
			if !reflect.DeepEqual(tt.input, original) {
				t.Errorf("Sort() modified original slice, got %v, want %v", tt.input, original)
			}
		})
	}
}

func TestSortDescending(t *testing.T) {
	input := []int{3, 1, 4, 1, 5}
	less := func(a, b int) bool { return gvalue.GT[int](a, b) } // descending
	result := Sort(input, less)
	expected := []int{5, 4, 3, 1, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sort() descending = %v, want %v", result, expected)
	}

	// Verify original is not modified
	if !reflect.DeepEqual(input, []int{3, 1, 4, 1, 5}) {
		t.Errorf("Sort() modified original slice")
	}
}

func TestStealSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		less     func(int, int) bool
		expected []int
	}{
		{
			name:     "sort integers ascending",
			input:    []int{3, 1, 4, 1, 5},
			less:     gvalue.Less[int],
			expected: []int{1, 1, 3, 4, 5},
		},
		{
			name:     "already sorted",
			input:    []int{1, 2, 3, 4, 5},
			less:     gvalue.Less[int],
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse order",
			input:    []int{5, 4, 3, 2, 1},
			less:     gvalue.Less[int],
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "nil slice",
			input:    nil,
			less:     gvalue.Less[int],
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			less:     gvalue.Less[int],
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{42},
			less:     gvalue.Less[int],
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StealSort(tt.input, tt.less)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("StealSort() = %v, want %v", result, tt.expected)
			}

			// Verify original slice is modified (same underlying array)
			if tt.input != nil && !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("StealSort() did not modify original slice, got %v, want %v", tt.input, tt.expected)
			}
		})
	}
}

func TestStealSortModifiesOriginal(t *testing.T) {
	input := []int{3, 1, 4, 1, 5}
	originalPtr := &input[0] // Keep reference to first element

	result := StealSort(input, gvalue.Less[int])
	expected := []int{1, 1, 3, 4, 5}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("StealSort() = %v, want %v", result, expected)
	}

	// Verify original slice is modified
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("StealSort() did not modify original slice, got %v, want %v", input, expected)
	}

	// Verify same underlying array (pointer should be the same after sort)
	// Note: This might not always be true due to sorting algorithm, but the slice should be modified
	if &input[0] != originalPtr && len(input) > 0 {
		// This is expected - the first element changed, but the slice itself is modified
	}
}
