package gslice

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		eq       func(int, int) bool
		expected []int
	}{
		{
			name:     "remove duplicates",
			input:    []int{1, 2, 2, 3, 3, 3, 4},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "all duplicates",
			input:    []int{1, 1, 1, 1},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{1},
		},
		{
			name:     "nil slice",
			input:    nil,
			eq:       func(a, b int) bool { return a == b },
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			eq:       func(a, b int) bool { return a == b },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unique(tt.input, tt.eq)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unique() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "multiple slices with duplicates",
			slices:   [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "no duplicates",
			slices:   [][]int{{1, 2}, {3, 4}, {5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "single slice",
			slices:   [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no slices",
			slices:   [][]int{},
			expected: nil,
		},
		{
			name:     "empty slices",
			slices:   [][]int{{}, {}, {}},
			expected: []int{},
		},
		{
			name:     "all same elements",
			slices:   [][]int{{1, 1}, {1, 1}, {1}},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Union(tt.slices...)
			// Union may return elements in different order, so we need to check if all elements are present
			if len(result) != len(tt.expected) {
				t.Errorf("Union() length = %v, want %v", len(result), len(tt.expected))
			}

			// Check if all expected elements are in result
			resultMap := make(map[int]struct{})
			for _, v := range result {
				resultMap[v] = struct{}{}
			}
			for _, v := range tt.expected {
				if _, exists := resultMap[v]; !exists {
					t.Errorf("Union() missing element %v", v)
				}
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "common elements",
			slices:   [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}},
			expected: []int{3, 4},
		},
		{
			name:     "no common elements",
			slices:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{},
		},
		{
			name:     "single slice",
			slices:   [][]int{{1, 2, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "all same elements",
			slices:   [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no slices",
			slices:   [][]int{},
			expected: nil,
		},
		{
			name:     "empty slice in input",
			slices:   [][]int{{1, 2, 3}, {}, {2, 3}},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.slices...)
			// Intersection may return elements in different order, so we need to check if all elements are present
			if len(result) != len(tt.expected) {
				t.Errorf("Intersection() length = %v, want %v, got %v", len(result), len(tt.expected), result)
			}

			// Check if all expected elements are in result
			resultMap := make(map[int]struct{})
			for _, v := range result {
				resultMap[v] = struct{}{}
			}
			for _, v := range tt.expected {
				if _, exists := resultMap[v]; !exists {
					t.Errorf("Intersection() missing element %v, got %v", v, result)
				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		others   [][]int
		expected []int
	}{
		{
			name:     "difference with multiple others",
			first:    []int{1, 2, 3, 4, 5},
			others:   [][]int{{2, 4}, {3}},
			expected: []int{1, 5},
		},
		{
			name:     "no difference",
			first:    []int{1, 2, 3},
			others:   [][]int{{1, 2, 3}},
			expected: []int{},
		},
		{
			name:     "all different",
			first:    []int{1, 2, 3},
			others:   [][]int{{4, 5, 6}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no others",
			first:    []int{1, 2, 2, 3},
			others:   [][]int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty first",
			first:    []int{},
			others:   [][]int{{1, 2, 3}},
			expected: []int{},
		},
		{
			name:     "duplicates in first",
			first:    []int{1, 1, 2, 2, 3},
			others:   [][]int{{2}},
			expected: []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.first, tt.others...)
			// Difference may return elements in different order, so we need to check if all elements are present
			if len(result) != len(tt.expected) {
				t.Errorf("Difference() length = %v, want %v, got %v", len(result), len(tt.expected), result)
			}

			// Check if all expected elements are in result
			resultMap := make(map[int]struct{})
			for _, v := range result {
				resultMap[v] = struct{}{}
			}
			for _, v := range tt.expected {
				if _, exists := resultMap[v]; !exists {
					t.Errorf("Difference() missing element %v, got %v", v, result)
				}
			}
		})
	}
}

