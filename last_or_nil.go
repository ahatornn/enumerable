package enumerable

// LastOrNil returns a pointer to the last element of an enumeration,
// or nil if the enumeration is empty or nil.
//
// The last or nil operation will:
//   - Iterate through all elements in the enumeration
//   - Keep track of the most recent element encountered
//   - Return a pointer to the last element if enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the last element, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(1) - constant space usage
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
func (q Enumerator[T]) LastOrNil() *T {
	return lastOrNilInternal(q)
}

// LastOrNil returns a pointer to the last element of an enumeration,
// or nil if the enumeration is empty or nil.
//
// The last or nil operation will:
//   - Iterate through all elements in the enumeration
//   - Keep track of the most recent element encountered
//   - Return a pointer to the last element if enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the last element, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(1) - constant space usage
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
func (q EnumeratorAny[T]) LastOrNil() *T {
	return lastOrNilInternal(q)
}

// LastOrNil returns a pointer to the last element of a sorted enumeration in sorted order,
// or nil if the enumeration is empty or nil.
//
// The last or nil operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Iterate through all elements in the sorted enumeration
//   - Keep track of the most recent element encountered in sorted order
//   - Return a pointer to the last element if enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the last element in sorted order, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the entire sorted enumeration to find the last element.
// Time complexity: O(n log n) for sorting + O(n) for iteration = O(n log n)
// where n is the number of elements. For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// but does not buffer elements beyond tracking the last element during iteration.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n) during sorting, O(1) during iteration
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the maximum element according to sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) LastOrNil() *T {
	return lastOrNilInternal(o.getSortedEnumerator())
}

// LastOrNil returns a pointer to the last element of a sorted enumeration in sorted order,
// or nil if the enumeration is empty or nil.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The last or nil operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Iterate through all elements in the sorted enumeration
//   - Keep track of the most recent element encountered in sorted order
//   - Return a pointer to the last element if enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the last element in sorted order, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the entire sorted enumeration to find the last element.
// Time complexity: O(n log n) for sorting + O(n) for iteration = O(n log n)
// where n is the number of elements. For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// but does not buffer elements beyond tracking the last element during iteration.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Actual sorting computation occurs only during this first operation
//   - Custom comparer functions are used for element comparison during sorting
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n) during sorting, O(1) during iteration
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the maximum element according to sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) LastOrNil() *T {
	return lastOrNilInternal(o.getSortedEnumerator())
}

func lastOrNilInternal[T any](enumerator func(func(T) bool)) *T {
	if enumerator == nil {
		return nil
	}

	var lastItem T
	var found bool

	enumerator(func(item T) bool {
		lastItem = item
		found = true
		return true
	})

	if !found {
		return nil
	}

	return &lastItem
}
