package gslice

import (
	"reflect"
	"testing"

	"github.com/geebos/gocraft/pkg/gvalue"
)

func TestCmpWith(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		cmp      func(int, int) bool
		value    int
		expected []int
	}{
		{
			name:     "greater than 3",
			input:    []int{1, 2, 3, 4, 5},
			cmp:      gvalue.GT[int],
			value:    3,
			expected: []int{4, 5},
		},
		{
			name:     "less than or equal to 2",
			input:    []int{1, 2, 3, 4, 5},
			cmp:      gvalue.LTE[int],
			value:    2,
			expected: []int{1, 2},
		},
		{
			name:     "equal to 3",
			input:    []int{1, 2, 3, 4, 3, 5},
			cmp:      gvalue.EQ[int],
			value:    3,
			expected: []int{3, 3},
		},
		{
			name:     "greater than or equal to 4",
			input:    []int{1, 2, 3, 4, 5},
			cmp:      gvalue.GTE[int],
			value:    4,
			expected: []int{4, 5},
		},
		{
			name:     "less than 3",
			input:    []int{1, 2, 3, 4, 5},
			cmp:      gvalue.LT[int],
			value:    3,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			predicate := CmpWith(tt.cmp, tt.value)
			result := Filter(tt.input, predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CmpWith() with Filter() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCmpWithFind(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	predicate := CmpWith(gvalue.LTE[int], 2)
	value, found := Find(input, predicate)

	if !found {
		t.Errorf("Find() with CmpWith() found = false, want true")
	}
	if value != 1 {
		t.Errorf("Find() with CmpWith() = %v, want 1", value)
	}
}

func TestCmpWithAny(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	predicate := CmpWith(gvalue.GT[int], 3)
	result := Any(input, predicate)

	if !result {
		t.Errorf("Any() with CmpWith() = false, want true")
	}
}

func TestCmpWithAll(t *testing.T) {
	input := []int{2, 4, 6, 8}
	predicate := CmpWith(gvalue.GTE[int], 2)
	result := All(input, predicate)

	if !result {
		t.Errorf("All() with CmpWith() = false, want true")
	}

	input2 := []int{1, 2, 3, 4}
	predicate2 := CmpWith(gvalue.GT[int], 2)
	result2 := All(input2, predicate2)

	if result2 {
		t.Errorf("All() with CmpWith() = true, want false")
	}
}
