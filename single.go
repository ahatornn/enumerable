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
	var result T
	count := 0

	if e != nil {
		e(func(item T) bool {
			if count == 0 {
				result = item
			}
			count++
			return count <= 1
		})
	}

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
