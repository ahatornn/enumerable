package enumerable

// LastOrDefault returns the last element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The last or default operation will:
//   - Iterate through all elements in the enumeration
//   - Keep track of the most recent element encountered
//   - Return the last element if enumeration contains elements
//   - Return the provided default value if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The last element of the enumeration, or the default value if enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(1) - constant space usage
//   - Unlike LastOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - When using zero value as default, consider using LastOrNil() for distinction
func (q Enumerator[T]) LastOrDefault(defaultValue T) T {
	if last := q.LastOrNil(); last != nil {
		return *last
	}
	return defaultValue
}

// LastOrDefault returns the last element of an enumeration, or a default value
// if the enumeration is empty or nil.
//
// The last or default operation will:
//   - Iterate through all elements in the enumeration
//   - Keep track of the most recent element encountered
//   - Return the last element if enumeration contains elements
//   - Return the provided default value if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The last element of the enumeration, or the default value if enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the last element. For large
// enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(1) - constant space usage
//   - Unlike LastOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - When using zero value as default, consider using LastOrNil() for distinction
func (q EnumeratorAny[T]) LastOrDefault(defaultValue T) T {
	if last := q.LastOrNil(); last != nil {
		return *last
	}
	return defaultValue
}

// LastOrDefault returns the last element of a sorted enumeration in sorted order,
// or a default value if the enumeration is empty or nil.
//
// The last or default operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Iterate through all elements in the sorted enumeration
//   - Keep track of the most recent element encountered in sorted order
//   - Return the last element if enumeration contains elements
//   - Return the provided default value if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The last element of the enumeration in sorted order, or the default value if enumeration is empty
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
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n) during sorting, O(1) during iteration
//   - Unlike LastOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the maximum element according to sorting rules
//   - When using zero value as default, consider using LastOrNil() for distinction
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) LastOrDefault(defaultValue T) T {
	if last := o.LastOrNil(); last != nil {
		return *last
	}
	return defaultValue
}

// LastOrDefault returns the last element of a sorted enumeration in sorted order,
// or a default value if the enumeration is empty or nil.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The last or default operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Iterate through all elements in the sorted enumeration
//   - Keep track of the most recent element encountered in sorted order
//   - Return the last element if enumeration contains elements
//   - Return the provided default value if the enumeration is empty or nil
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	defaultValue - the value to return if the enumeration is empty or nil
//
// Returns:
//
//	The last element of the enumeration in sorted order, or the default value if enumeration is empty
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
//   - If the enumerator is nil, returns the defaultValue
//   - If the enumeration is empty, returns the defaultValue
//   - Actual sorting computation occurs only during this first operation
//   - Custom comparer functions are used for element comparison during sorting
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n) during sorting, O(1) during iteration
//   - Unlike LastOrNil(), this method returns the value directly, not a pointer
//   - Safe for all types including those with zero values like 0, "", false, etc.
//   - For sorted enumerations, this typically returns the maximum element according to sorting rules
//   - When using zero value as default, consider using LastOrNil() for distinction
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) LastOrDefault(defaultValue T) T {
	if last := o.LastOrNil(); last != nil {
		return *last
	}
	return defaultValue
}
