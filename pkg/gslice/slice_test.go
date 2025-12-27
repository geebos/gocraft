package gslice

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) int
		expected []int
	}{
		{
			name:     "double numbers",
			input:    []int{1, 2, 3, 4},
			fn:       func(n int) int { return n * 2 },
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "square numbers",
			input:    []int{1, 2, 3, 4},
			fn:       func(n int) int { return n * n },
			expected: []int{1, 4, 9, 16},
		},
		{
			name:     "nil slice",
			input:    nil,
			fn:       func(n int) int { return n * 2 },
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			fn:       func(n int) int { return n * 2 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.input, tt.fn)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMapStringToInt(t *testing.T) {
	input := []string{"a", "ab", "abc"}
	fn := func(s string) int { return len(s) }
	result := Map(input, fn)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) bool
		expected []int
	}{
		{
			name:     "filter even numbers",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: []int{2, 4},
		},
		{
			name:     "filter greater than 3",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(n int) bool { return n > 3 },
			expected: []int{4, 5},
		},
		{
			name:     "nil slice",
			input:    nil,
			fn:       func(n int) bool { return n%2 == 0 },
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []int{},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: []int{},
		},
		{
			name:     "no matches",
			input:    []int{1, 3, 5},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.input, tt.fn)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Filter() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		initial  int
		fn       func(int, int) int
		expected int
	}{
		{
			name:     "sum numbers",
			input:    []int{1, 2, 3, 4},
			initial:  0,
			fn:       func(acc, n int) int { return acc + n },
			expected: 10,
		},
		{
			name:     "multiply numbers",
			input:    []int{1, 2, 3, 4},
			initial:  1,
			fn:       func(acc, n int) int { return acc * n },
			expected: 24,
		},
		{
			name:     "empty slice",
			input:    []int{},
			initial:  10,
			fn:       func(acc, n int) int { return acc + n },
			expected: 10,
		},
		{
			name:     "max value",
			input:    []int{3, 1, 4, 1, 5},
			initial:  0,
			fn:       func(acc, n int) int { if n > acc { return n }; return acc },
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce(tt.input, tt.initial, tt.fn)
			if result != tt.expected {
				t.Errorf("Reduce() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		fn            func(int) bool
		expectedValue int
		expectedFound bool
	}{
		{
			name:          "find first even",
			input:         []int{1, 2, 3, 4, 5},
			fn:            func(n int) bool { return n%2 == 0 },
			expectedValue: 2,
			expectedFound: true,
		},
		{
			name:          "find greater than 3",
			input:         []int{1, 2, 3, 4, 5},
			fn:            func(n int) bool { return n > 3 },
			expectedValue: 4,
			expectedFound: true,
		},
		{
			name:          "not found",
			input:         []int{1, 2, 3},
			fn:            func(n int) bool { return n > 10 },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "nil slice",
			input:         nil,
			fn:            func(n int) bool { return n > 0 },
			expectedValue: 0,
			expectedFound: false,
		},
		{
			name:          "empty slice",
			input:         []int{},
			fn:            func(n int) bool { return n > 0 },
			expectedValue: 0,
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := Find(tt.input, tt.fn)
			if value != tt.expectedValue || found != tt.expectedFound {
				t.Errorf("Find() = (%v, %v), want (%v, %v)", value, found, tt.expectedValue, tt.expectedFound)
			}
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) bool
		expected bool
	}{
		{
			name:     "has even number",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: true,
		},
		{
			name:     "has number greater than 3",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(n int) bool { return n > 3 },
			expected: true,
		},
		{
			name:     "no matches",
			input:    []int{1, 3, 5},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: false,
		},
		{
			name:     "nil slice",
			input:    nil,
			fn:       func(n int) bool { return n > 0 },
			expected: false,
		},
		{
			name:     "empty slice",
			input:    []int{},
			fn:       func(n int) bool { return n > 0 },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Any(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("Any() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		fn       func(int) bool
		expected bool
	}{
		{
			name:     "all even",
			input:    []int{2, 4, 6, 8},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: true,
		},
		{
			name:     "all greater than 0",
			input:    []int{1, 2, 3, 4},
			fn:       func(n int) bool { return n > 0 },
			expected: true,
		},
		{
			name:     "not all even",
			input:    []int{2, 3, 4, 6},
			fn:       func(n int) bool { return n%2 == 0 },
			expected: false,
		},
		{
			name:     "nil slice",
			input:    nil,
			fn:       func(n int) bool { return n > 0 },
			expected: true,
		},
		{
			name:     "empty slice",
			input:    []int{},
			fn:       func(n int) bool { return n > 0 },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := All(tt.input, tt.fn)
			if result != tt.expected {
				t.Errorf("All() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "multiple slices",
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
			name:     "mixed empty and non-empty",
			slices:   [][]int{{1}, {}, {2, 3}},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Concat(tt.slices...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Concat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		start    int
		end      int
		expected []int
	}{
		{
			name:     "normal slice",
			input:    []int{1, 2, 3, 4, 5},
			start:    1,
			end:      4,
			expected: []int{2, 3, 4},
		},
		{
			name:     "full slice",
			input:    []int{1, 2, 3, 4, 5},
			start:    0,
			end:      5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "single element",
			input:    []int{1, 2, 3, 4, 5},
			start:    2,
			end:      3,
			expected: []int{3},
		},
		{
			name:     "nil slice",
			input:    nil,
			start:    0,
			end:      0,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Slice(tt.input, tt.start, tt.end)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Slice() = %v, want %v", result, tt.expected)
			}
		})
	}
}

