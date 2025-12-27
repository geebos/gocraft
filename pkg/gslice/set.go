package gslice

// Unique removes duplicate elements from a slice.
//
// Unique uses the provided equality function to determine if two elements are equal.
// It returns a new slice containing only unique elements, preserving the order
// of the first occurrence of each element. The original slice is not modified.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4}
//	unique := Unique(numbers, func(a, b int) bool { return a == b })
//	// unique is []int{1, 2, 3, 4}
func Unique[T any](s []T, eq func(T, T) bool) []T {
	if s == nil {
		return nil
	}
	if len(s) == 0 {
		return []T{}
	}
	result := make([]T, 0, len(s))

	for i, v := range s {
		isDuplicate := false
		for j := 0; j < i; j++ {
			if eq(s[j], v) {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			result = append(result, v)
		}
	}
	return result
}

// Union merges multiple slices and removes duplicates.
//
// Union combines all elements from all input slices and returns a new slice
// containing only unique elements. The order is preserved based on the first
// occurrence of each element across all slices. The input slices are not modified.
//
// Union requires T to be comparable, allowing efficient duplicate detection using a map.
//
// Example:
//
//	s1 := []int{1, 2, 3}
//	s2 := []int{2, 3, 4}
//	s3 := []int{3, 4, 5}
//	unified := Union(s1, s2, s3)
//	// unified is []int{1, 2, 3, 4, 5}
func Union[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	// Calculate total capacity
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}
	if totalLen == 0 {
		return []T{}
	}

	// Use a map to track seen elements for O(1) lookup
	seen := make(map[T]struct{}, totalLen)
	result := make([]T, 0, totalLen)

	// Iterate through all slices
	for _, s := range slices {
		for _, elem := range s {
			// Check if element has been seen using map lookup
			if _, exists := seen[elem]; !exists {
				// Mark as seen and add to result
				seen[elem] = struct{}{}
				result = append(result, elem)
			}
		}
	}

	return result
}

// Intersection returns elements that exist in all input slices.
//
// Intersection finds elements that are present in every slice provided.
// It returns a new slice containing these common elements, with duplicates
// removed. The order is based on the first slice. The input slices are not modified.
//
// Intersection requires T to be comparable, allowing efficient lookup using maps.
//
// Example:
//
//	s1 := []int{1, 2, 3, 4}
//	s2 := []int{2, 3, 4, 5}
//	s3 := []int{3, 4, 5, 6}
//	common := Intersection(s1, s2, s3)
//	// common is []int{3, 4}
func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}
	if len(slices) == 1 {
		// Use Union to remove duplicates from a single slice
		return Union(slices[0])
	}

	// Count occurrences of each element across all slices
	// Use a map to count how many slices contain each element
	countMap := make(map[T]int)

	// Iterate through each slice
	for _, s := range slices {
		// Use a map to deduplicate elements within a single slice
		// to avoid counting the same element multiple times in one slice
		seenInSlice := make(map[T]struct{})
		for _, elem := range s {
			// Only count each element once per slice
			if _, exists := seenInSlice[elem]; !exists {
				seenInSlice[elem] = struct{}{}
				countMap[elem]++
			}
		}
	}

	// Collect elements that appear in all slices (count == number of slices)
	sliceCount := len(slices)
	result := make([]T, 0)

	// Iterate through countMap to collect elements that appear in all slices
	for elem, count := range countMap {
		if count >= sliceCount {
			result = append(result, elem)
		}
	}

	return result
}

// Difference returns elements from the first slice that are not in any of the other slices.
//
// Difference finds elements that exist in the first slice but not in any subsequent slice.
// It returns a new slice containing these elements, with duplicates removed.
// The order is based on the first slice. The input slices are not modified.
//
// Difference requires T to be comparable, allowing efficient lookup using maps.
//
// Example:
//
//	s1 := []int{1, 2, 3, 4, 5}
//	s2 := []int{2, 4}
//	s3 := []int{3}
//	diff := Difference(s1, s2, s3)
//	// diff is []int{1, 5}
func Difference[T comparable](first []T, others ...[]T) []T {
	if len(first) == 0 {
		return []T{}
	}
	if len(others) == 0 {
		// Use Union to remove duplicates from a single slice
		return Union(first)
	}

	// Build a map of all elements in other slices for O(1) lookup
	otherSet := make(map[T]struct{})
	for _, other := range others {
		for _, elem := range other {
			otherSet[elem] = struct{}{}
		}
	}

	// Track seen elements to avoid duplicates
	seen := make(map[T]struct{}, len(first))
	result := make([]T, 0, len(first))

	// For each element in the first slice, check if it exists in other slices
	for _, elem := range first {
		// Skip if already added
		if _, exists := seen[elem]; exists {
			continue
		}

		// Check if element exists in other slices using map lookup
		if _, exists := otherSet[elem]; !exists {
			seen[elem] = struct{}{}
			result = append(result, elem)
		}
	}

	return result
}
