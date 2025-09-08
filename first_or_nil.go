package enumerable

// FirstOrNil returns a pointer to the first element of an enumeration.
// This operation is useful for getting the first element when it exists,
// with the ability to distinguish between "no elements" and "zero value" cases.
//
// The first operation will:
//   - Return a pointer to the first element if the enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the first element, or nil if enumeration is empty or nil
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Lazy evaluation stops immediately after finding the first element
//   - This is a terminal operation that materializes the enumeration
//   - Very efficient - O(1) time complexity for non-empty enumerations
//   - No elements are buffered - memory efficient
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
func (q Enumerator[T]) FirstOrNil() *T {
	return firstOrNilInternal(q)
}

// FirstOrNil returns a pointer to the first element of an enumeration.
// This operation is useful for getting the first element when it exists,
// with the ability to distinguish between "no elements" and "zero value" cases.
//
// The first operation will:
//   - Return a pointer to the first element if the enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the first element, or nil if enumeration is empty or nil
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Lazy evaluation stops immediately after finding the first element
//   - This is a terminal operation that materializes the enumeration
//   - Very efficient - O(1) time complexity for non-empty enumerations
//   - No elements are buffered - memory efficient
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
func (q EnumeratorAny[T]) FirstOrNil() *T {
	return firstOrNilInternal(q)
}

// FirstOrNil returns a pointer to the first element of a sorted enumeration.
// This operation is useful for getting the first element in sorted order when it exists,
// with the ability to distinguish between "no elements" and "zero value" cases.
//
// The first operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return a pointer to the first element in sorted order if the enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the first element in sorted order, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must perform sorting to determine the first element in sorted order.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Actual sorting computation occurs only during this first operation
//   - Lazy evaluation stops immediately after finding the first element in sorted order
//   - This is a terminal operation that materializes the enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - No elements are buffered beyond the first element - memory efficient after sorting
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the minimum element according to sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) FirstOrNil() *T {
	return firstOrNilInternal(o.getSortedEnumerator())
}

// FirstOrNil returns a pointer to the first element of a sorted enumeration.
// with the ability to distinguish between "no elements" and "zero value" cases.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The first operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return a pointer to the first element in sorted order if the enumeration contains elements
//   - Return nil if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Returns:
//
//	A pointer to the first element in sorted order, or nil if enumeration is empty or nil
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must perform sorting to determine the first element in sorted order.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns nil
//   - If the enumeration is empty, returns nil
//   - Actual sorting computation occurs only during this first operation
//   - Lazy evaluation stops immediately after finding the first element in sorted order
//   - This is a terminal operation that materializes the enumeration
//   - Custom comparer functions are used for element comparison during sorting
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - No elements are buffered beyond the first element - memory efficient after sorting
//   - Returns pointer to allow distinction between "no elements" (nil) and "zero value" element
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the minimum element according to sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) FirstOrNil() *T {
	return firstOrNilInternal(o.getSortedEnumerator())
}

func firstOrNilInternal[T any](enumerator func(func(T) bool)) *T {
	if enumerator == nil {
		return nil
	}
	var result *T
	enumerator(func(item T) bool {
		temp := item
		result = &temp
		return false
	})
	return result
}
