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
