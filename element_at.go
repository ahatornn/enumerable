package enumerable

// ElementAt returns the element at the specified index in the sequence.
// This operation is useful when you need to access an element at a specific position
// without iterating through the entire sequence.
//
// The ElementAt operation will:
//   - Return the element at the specified zero-based index and true if found
//   - Return the zero value of T and false if index is negative
//   - Return the zero value of T and false if index is beyond the sequence bounds
//   - Return the zero value of T and false if the enumerator is nil
//   - Process elements sequentially until the specified index is reached
//
// Parameters:
//
//	index - the zero-based index of the element to retrieve
//
// Returns:
//
//	The element at the specified index and true if found,
//	zero value of T and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until the specified index is reached.
// Time complexity: O(n) where n is index + 1.
//
// ⚠️ Memory note: This operation does not buffer elements, only tracks current position.
//
// Notes:
//   - Index must be non-negative (negative indices return zero value and false)
//   - If index is beyond sequence bounds, returns zero value and false
//   - If enumerator is nil, returns zero value and false
//   - Processes elements sequentially - early termination when index is found
//   - No elements are buffered - memory efficient
//   - This is a terminal operation that materializes the enumeration
func (e Enumerator[T]) ElementAt(index int) (T, bool) {
	return elementAtInternal(e, index)
}

// ElementAt returns the element at the specified index in the sequence.
// This operation is useful when you need to access an element at a specific position
// without iterating through the entire sequence.
//
// The ElementAt operation will:
//   - Return the element at the specified zero-based index and true if found
//   - Return the zero value of T and false if index is negative
//   - Return the zero value of T and false if index is beyond the sequence bounds
//   - Return the zero value of T and false if the enumerator is nil
//   - Process elements sequentially until the specified index is reached
//
// Parameters:
//
//	index - the zero-based index of the element to retrieve
//
// Returns:
//
//	The element at the specified index and true if found,
//	zero value of T and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until the specified index is reached.
// Time complexity: O(n) where n is index + 1.
//
// ⚠️ Memory note: This operation does not buffer elements, only tracks current position.
//
// Notes:
//   - Index must be non-negative (negative indices return zero value and false)
//   - If index is beyond sequence bounds, returns zero value and false
//   - If enumerator is nil, returns zero value and false
//   - Processes elements sequentially - early termination when index is found
//   - No elements are buffered - memory efficient
//   - This is a terminal operation that materializes the enumeration
//   - Works with any type T (including non-comparable types)
func (e EnumeratorAny[T]) ElementAt(index int) (T, bool) {
	return elementAtInternal(e, index)
}

// ElementAt returns the element at the specified index in the sorted sequence.
// This operation is useful when you need to access an element at a specific position
// in the sorted order without iterating through the entire sequence.
//
// The ElementAt operation will:
//   - Execute deferred sorting rules and iterate through the sorted sequence
//   - Return the element at the specified zero-based index and true if found
//   - Return the zero value of T and false if index is negative
//   - Return the zero value of T and false if index is beyond the sequence bounds
//   - Return the zero value of T and false if the enumerator is nil
//   - Process elements sequentially until the specified index is reached in sorted order
//
// Parameters:
//
//	index - the zero-based index of the element to retrieve from the sorted sequence
//
// Returns:
//
//	The element at the specified index in sorted order and true if found,
//	zero value of T and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must perform sorting and then iterate through the sorted sequence
// until the specified index is reached. Time complexity: O(n log n) for sorting +
// O(index + 1) for iteration = O(n log n) where n is the total number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// but does not buffer elements beyond the specified index during iteration.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - Index must be non-negative (negative indices return zero value and false)
//   - If index is beyond sequence bounds, returns zero value and false
//   - If enumerator is nil, returns zero value and false
//   - Actual sorting computation occurs only during this first operation
//   - Processes elements sequentially in sorted order - early termination when index is found
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - No elements are buffered beyond the target index - memory efficient for early indices
//   - This is a terminal operation that materializes the enumeration
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) ElementAt(index int) (T, bool) {
	return elementAtInternal(o.getSortedEnumerator(), index)
}

// ElementAt returns the element at the specified index in the sorted sequence.
// This operation is useful when you need to access an element at a specific position
// in the sorted order without iterating through the entire sequence.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The ElementAt operation will:
//   - Execute deferred sorting rules and iterate through the sorted sequence
//   - Return the element at the specified zero-based index and true if found
//   - Return the zero value of T and false if index is negative
//   - Return the zero value of T and false if index is beyond the sequence bounds
//   - Return the zero value of T and false if the enumerator is nil
//   - Process elements sequentially until the specified index is reached in sorted order
//
// Parameters:
//
//	index - the zero-based index of the element to retrieve from the sorted sequence
//
// Returns:
//
//	The element at the specified index in sorted order and true if found,
//	zero value of T and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// The operation must perform sorting and then iterate through the sorted sequence
// until the specified index is reached. Time complexity: O(n log n) for sorting +
// O(index + 1) for iteration = O(n log n) where n is the total number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// but does not buffer elements beyond the specified index during iteration.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - Index must be non-negative (negative indices return zero value and false)
//   - If index is beyond sequence bounds, returns zero value and false
//   - If enumerator is nil, returns zero value and false
//   - Actual sorting computation occurs only during this first operation
//   - Processes elements sequentially in sorted order - early termination when index is found
//   - Custom comparer functions are used for element comparison during sorting
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - No elements are buffered beyond the target index - memory efficient for early indices
//   - This is a terminal operation that materializes the enumeration
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) ElementAt(index int) (T, bool) {
	return elementAtInternal(o.getSortedEnumerator(), index)
}

func elementAtInternal[T any](enumerator func(func(T) bool), index int) (T, bool) {
	var result T

	if enumerator == nil || index < 0 {
		return result, false
	}

	currentIndex := 0
	found := false

	enumerator(func(item T) bool {
		if currentIndex == index {
			result = item
			found = true
			return false
		}
		currentIndex++
		return true
	})

	return result, found
}
