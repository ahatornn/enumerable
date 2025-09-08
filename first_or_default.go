package enumerable

// FirstOrDefault returns the first element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The first or default operation will:
//   - Return the first element from the enumeration if it exists
//   - Return the provided default value if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The first element of the enumeration, or the default value if enumeration is empty
//
// Notes:
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Lazy evaluation stops immediately after finding the first element
//   - This is a terminal operation that materializes the enumeration
//   - Very efficient - O(1) time complexity for non-empty enumerations
//   - No elements are buffered - memory efficient
//   - Unlike FirstOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - When using zero value as default, consider using FirstOrNil() for distinction
func (q Enumerator[T]) FirstOrDefault(defaultValue T) T {
	if first := q.FirstOrNil(); first != nil {
		return *first
	}
	return defaultValue
}

// FirstOrDefault returns the first element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The first or default operation will:
//   - Return the first element from the enumeration if it exists
//   - Return the provided default value if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The first element of the enumeration, or the default value if enumeration is empty
//
// Notes:
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Lazy evaluation stops immediately after finding the first element
//   - This is a terminal operation that materializes the enumeration
//   - Very efficient - O(1) time complexity for non-empty enumerations
//   - No elements are buffered - memory efficient
//   - Unlike FirstOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - When using zero value as default, consider using FirstOrNil() for distinction
func (q EnumeratorAny[T]) FirstOrDefault(defaultValue T) T {
	if first := q.FirstOrNil(); first != nil {
		return *first
	}
	return defaultValue
}

// FirstOrDefault returns the first element of a sorted enumeration in sorted order,
// or a default value if the enumeration is empty or nil.
//
// The first or default operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return the first element from the enumeration in sorted order if it exists
//   - Return the provided default value if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The first element of the enumeration in sorted order, or the default value if enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must perform sorting to determine the first element in sorted order.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Actual sorting computation occurs only during this first operation
//   - Lazy evaluation stops immediately after finding the first element in sorted order
//   - This is a terminal operation that materializes the enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - No elements are buffered beyond the first element - memory efficient after sorting
//   - Unlike FirstOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the minimum element according to sorting rules
//   - When using zero value as default, consider using FirstOrNil() for distinction
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) FirstOrDefault(defaultValue T) T {
	if first := o.FirstOrNil(); first != nil {
		return *first
	}
	return defaultValue
}

// FirstOrDefault returns the first element of a sorted enumeration in sorted order,
// or a default value if the enumeration is empty or nil.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The first or default operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return the first element from the enumeration in sorted order if it exists
//   - Return the provided default value if the enumeration is empty or nil
//   - Stop enumeration immediately after finding the first element
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The first element of the enumeration in sorted order, or the default value if enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must perform sorting to determine the first element in sorted order.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Actual sorting computation occurs only during this first operation
//   - Lazy evaluation stops immediately after finding the first element in sorted order
//   - This is a terminal operation that materializes the enumeration
//   - Custom comparer functions are used for element comparison during sorting
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - No elements are buffered beyond the first element - memory efficient after sorting
//   - Unlike FirstOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the minimum element according to sorting rules
//   - When using zero value as default, consider using FirstOrNil() for distinction
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) FirstOrDefault(defaultValue T) T {
	if first := o.FirstOrNil(); first != nil {
		return *first
	}
	return defaultValue
}
