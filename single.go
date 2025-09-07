package enumerable

// Single returns the single element of a sequence using default equality comparison.
// This operation is useful when you expect exactly one element in the sequence.
//
// The Single operation will:
//   - Return the single element if the sequence contains exactly one element
//   - Return an error if the sequence is empty ("sequence contains no elements")
//   - Return an error if the sequence contains more than one element ("sequence contains more than one element")
//   - Process elements sequentially until completion or error
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	The single element of the sequence and nil error if successful,
//	zero value of T and error otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until completion or error. For large sequences
// with multiple elements, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// process multiple elements to ensure uniqueness, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns zero value and error ("sequence contains no elements")
//   - If the enumeration is empty, returns zero value and error ("sequence contains no elements")
//   - If the enumeration contains more than one element, returns zero value and error ("sequence contains more than one element")
//   - Processes elements in the enumeration - O(n) time complexity in worst case
//   - No elements are buffered - memory efficient
//   - This is a terminal operation that materializes the enumeration
//   - Works only with comparable types (no slices, maps, functions in struct fields)
//   - The type T must be comparable for equality comparison to work
//   - For large enumerations, all elements may be processed to ensure uniqueness
//   - This method should be used when you are certain the sequence contains exactly one element
//   - Common use cases include lookup by unique identifier or configuration validation
func (e Enumerator[T]) Single() (T, error) {
	return singleInternal(e)
}

// Single returns the single element of a sorted sequence using default equality comparison.
// This operation is useful when you expect exactly one element in the sorted sequence.
//
// The Single operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return the single element if the sequence contains exactly one element
//   - Return an error if the sequence is empty ("sequence contains no elements")
//   - Return an error if the sequence contains more than one element ("sequence contains more than one element")
//   - Process elements sequentially until completion or error
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	The single element of the sequence and nil error if successful,
//	zero value of T and error otherwise
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the sequence until completion or error. For large sequences
// with multiple elements, this may process more elements than necessary.
// Time complexity: O(n log n) for sorting + O(n) for validation = O(n log n)
// where n is the number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// but it does not buffer elements beyond validation during iteration.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns zero value and error ("sequence contains no elements")
//   - If the enumeration is empty, returns zero value and error ("sequence contains no elements")
//   - If the enumeration contains more than one element, returns zero value and error ("sequence contains more than one element")
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Processes elements in the enumeration - O(n log n) time complexity in worst case
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - This is a terminal operation that materializes the enumeration
//   - Works only with comparable types (no slices, maps, functions in struct fields)
//   - The type T must be comparable for equality comparison to work
//   - For large enumerations, all elements may be processed to ensure uniqueness
//   - This method should be used when you are certain the sequence contains exactly one element
//   - Common use cases include lookup by unique identifier or configuration validation
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) Single() (T, error) {
	return singleInternal(o.getSortedEnumerator())
}

// SingleOrDefault returns the single element of a sequence using default equality comparison,
// or a specified default value if the sequence is empty.
//
// The SingleOrDefault operation will:
//   - Return the single element if the sequence contains exactly one element
//   - Return the specified default value if the sequence is empty (including nil enumerator)
//   - Return the specified default value if the sequence contains more than one element
//   - Process elements sequentially until completion or second element found
//   - Handle nil enumerators gracefully (treated as empty)
//
// Returns:
//
//	The single element of the sequence if successful,
//	or the specified default value otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until completion or second element found.
// For large sequences with multiple elements, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// process multiple elements to ensure uniqueness, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns the specified default value
//   - If the enumeration is empty, returns the specified default value
//   - If the enumeration contains more than one element, returns the specified default value
//   - Processes elements in the enumeration - O(n) time complexity in worst case
//   - No elements are buffered - memory efficient
//   - This is a terminal operation that materializes the enumeration
//   - Works only with comparable types (no slices, maps, functions in struct fields)
//   - The type T must be comparable for equality comparison to work
//   - For large enumerations, elements may be processed to detect multiple items
//   - This method should be used when you expect zero or one element, with fallback for other cases
//   - Common use cases include optional lookup by unique identifier or configuration with defaults
//
// ⚠️ Important: Unlike Single(), this method never returns an error.
// All error conditions (empty sequence, multiple elements) result in returning the default value.
// If you need to distinguish between these cases, use Single() instead.
func (e Enumerator[T]) SingleOrDefault(defaultValue T) T {
	result, err := singleInternal(e)
	if err != nil {
		return defaultValue
	}
	return result
}

// SingleOrDefault returns the single element of a sorted sequence using default equality comparison,
// or a specified default value if the sequence is empty.
//
// The SingleOrDefault operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return the single element if the sequence contains exactly one element
//   - Return the specified default value if the sequence is empty (including nil enumerator)
//   - Return the specified default value if the sequence contains more than one element
//   - Process elements sequentially until completion or second element found
//   - Handle nil enumerators gracefully (treated as empty)
//
// Returns:
//
//	The single element of the sequence if successful,
//	or the specified default value otherwise
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the sequence until completion or second element found.
// For large sequences with multiple elements, this may process more elements than necessary.
// Time complexity: O(n log n) for sorting + O(n) for validation = O(n log n)
// where n is the number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// but it does not buffer elements beyond validation during iteration.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns the specified default value
//   - If the enumeration is empty, returns the specified default value
//   - If the enumeration contains more than one element, returns the specified default value
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Processes elements in the enumeration - O(n log n) time complexity in worst case
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - This is a terminal operation that materializes the enumeration
//   - Works only with comparable types (no slices, maps, functions in struct fields)
//   - The type T must be comparable for equality comparison to work
//   - For large enumerations, elements may be processed to detect multiple items
//   - This method should be used when you expect zero or one element, with fallback for other cases
//   - Common use cases include optional lookup by unique identifier or configuration with defaults
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//
// ⚠️ Important: Unlike Single(), this method never returns an error.
// All error conditions (empty sequence, multiple elements) result in returning the default value.
// If you need to distinguish between these cases, use Single() instead.
func (o OrderEnumerator[T]) SingleOrDefault(defaultValue T) T {
	result, err := singleInternal(o.getSortedEnumerator())
	if err != nil {
		return defaultValue
	}
	return result
}

func singleInternal[T any](enumerator func(func(T) bool)) (T, error) {
	var result T
	count := 0

	if enumerator == nil {
		var zero T
		return zero, ErrNoElements
	}

	enumerator(func(item T) bool {
		if count == 0 {
			result = item
		}
		count++
		return count <= 1
	})

	switch count {
	case 0:
		var zero T
		return zero, ErrNoElements
	case 1:
		return result, nil
	default:
		var zero T
		return zero, ErrMultipleElements
	}
}
