package enumerable

import (
	"github.com/ahatornn/enumerable/comparer"
)

// Contains determines whether a sequence contains a specified element using direct equality comparison.
// This operation is useful when you need to check for the existence of a specific value in the sequence.
//
// The Contains operation will:
//   - Return true if the sequence contains an element equal to the specified value
//   - Return false if the sequence does not contain the specified value
//   - Return false if the enumerator is nil
//   - Process elements sequentially until the element is found or sequence ends
//   - Use direct equality comparison (==) for optimal performance
//
// Parameters:
//
//	value - the value to locate in the sequence
//
// Returns:
//
//	true if the sequence contains an element equal to the specified value, false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until the element is found or sequence ends.
//
// ⚠️ Memory note: This operation does not buffer elements and requires zero memory allocations.
//
// Notes:
//   - Uses direct equality comparison (==) for maximum performance
//   - If enumerator is nil, returns false immediately
//   - Processes elements sequentially with early termination when element is found
//   - No elements are buffered - memory efficient
//   - This is a terminal operation that materializes the enumeration
//   - Zero allocations for all operations
//   - Common use cases include membership testing, validation, conditional logic
func (e Enumerator[T]) Contains(value T) bool {
	if e == nil {
		return false
	}

	found := false
	e(func(item T) bool {
		if item == value {
			found = true
			return false
		}
		return true
	})
	return found
}

// Contains determines whether a sequence contains a specified element using the provided equality comparer.
// This operation is useful when you need to check for the existence of a specific value in the sequence
// with custom equality logic.
//
// The Contains operation will:
//   - Return true if the sequence contains an element equal to the specified value according to the comparer
//   - Return false if the sequence does not contain the specified value
//   - Return false if the enumerator is nil
//   - Return false if the comparer is nil
//   - Process elements sequentially until the element is found or sequence ends
//   - Use the provided equality comparer for element comparison
//
// Parameters:
//
//	value - the value to locate in the sequence
//	comparer - the EqualityComparer to use for element comparison
//
// Returns:
//
//	true if the sequence contains an element equal to the specified value, false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until the element is found or sequence ends.
//
// ⚠️ Memory note: This operation does not buffer elements.
//
// Notes:
//   - Uses the provided equality comparer for comparison
//   - If enumerator is nil, returns false
//   - If comparer is nil, returns false
//   - Processes elements sequentially - early termination when element is found
//   - No elements are buffered - memory efficient
//   - This is a terminal operation that materializes the enumeration
//   - Works with any type T (including non-comparable types) when used with appropriate comparer
func (e EnumeratorAny[T]) Contains(value T, comparer comparer.EqualityComparer[T]) bool {
	if e == nil || comparer == nil {
		return false
	}

	found := false

	e(func(item T) bool {
		if comparer.Equals(item, value) {
			found = true
			return false
		}
		return true
	})

	return found
}
