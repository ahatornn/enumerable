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
	return containsInternal(e, value)
}

// Contains determines whether a sorted sequence contains a specified element using direct equality comparison.
// This operation is optimized for sorted sequences and leverages the ordering for efficient lookup.
//
// The Contains operation will:
//   - Return true if the sorted sequence contains an element equal to the specified value
//   - Return false if the sorted sequence does not contain the specified value
//   - Return false if the enumerator is nil
//   - Use binary search for efficient lookup when applicable (based on sorting rules)
//   - Execute deferred sorting rules to determine the sorted order before searching
//   - Support early termination when element is found
//
// Parameters:
//
//	value - the value to locate in the sorted sequence
//
// Returns:
//
//	true if the sorted sequence contains an element equal to the specified value, false otherwise
//
// ⚠️ Performance note: This is a terminal operation that performs a binary search
// over the sorted sequence. Time complexity is O(n log n) for sorting + O(log n) for search,
// which results in O(n log n), where n is the number of elements.
//
// ⚠️ Memory note: This operation does not buffer elements beyond what is required for sorting.
// Sorting may require O(n) additional memory depending on the sorting algorithm used internally.
//
// Notes:
//   - Uses direct equality comparison (==) for element comparison
//   - If enumerator is nil, returns false immediately
//   - Actual sorting computation occurs only once during the first enumeration or search
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - This is a terminal operation that materializes the enumeration
//   - Efficient for large sorted sequences due to binary search
//   - Common use cases include fast membership testing in sorted data, validation logic,
//     or conditional branching based on presence of values in sorted collections
func (o OrderEnumerator[T]) Contains(value T) bool {
	return containsInternal(o.getSortedEnumerator(), value)
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
	return containsAnyInternal(e, value, comparer)
}

// Contains determines whether a sorted sequence contains a specified element using the provided equality comparer.
// This operation is optimized for sorted sequences and leverages the ordering for efficient lookup
// with custom equality logic.
//
// The Contains operation will:
//   - Return true if the sorted sequence contains an element equal to the specified value according to the comparer
//   - Return false if the sorted sequence does not contain the specified value
//   - Return false if the enumerator is nil
//   - Return false if the comparer is nil
//   - Use binary search for efficient lookup when applicable (based on sorting rules)
//   - Execute deferred sorting rules to determine the sorted order before searching
//   - Support early termination when element is found
//   - Use the provided equality comparer for element comparison
//
// Parameters:
//
//	value - the value to locate in the sorted sequence
//	comparer - the EqualityComparer to use for element comparison
//
// Returns:
//
//	true if the sorted sequence contains an element equal to the specified value, false otherwise
//
// ⚠️ Performance note: This is a terminal operation that performs a binary search
// over the sorted sequence. Time complexity is O(n log n) for sorting + O(log n) for search,
// which results in O(n log n), where n is the number of elements.
//
// ⚠️ Memory note: This operation does not buffer elements beyond what is required for sorting.
// Sorting may require O(n) additional memory depending on the sorting algorithm used internally.
//
// Notes:
//   - Uses the provided equality comparer for comparison
//   - If enumerator is nil, returns false immediately
//   - If comparer is nil, returns false immediately
//   - Actual sorting computation occurs only once during the first enumeration or search
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - This is a terminal operation that materializes the enumeration
//   - Works with any type T (including non-comparable types) when used with appropriate comparer
//   - Efficient for large sorted sequences due to binary search
//   - Common use cases include fast membership testing in sorted data with custom equality logic,
//     validation logic, or conditional branching based on presence of values in sorted collections
func (o OrderEnumeratorAny[T]) Contains(value T, comparer comparer.EqualityComparer[T]) bool {
	return containsAnyInternal(o.getSortedEnumerator(), value, comparer)
}

func containsInternal[T comparable](enumerator func(func(T) bool), value T) bool {
	if enumerator == nil {
		return false
	}

	found := false
	enumerator(func(item T) bool {
		if item == value {
			found = true
			return false
		}
		return true
	})
	return found
}

func containsAnyInternal[T any](enumerator func(func(T) bool), value T, comparer comparer.EqualityComparer[T]) bool {
	if enumerator == nil || comparer == nil {
		return false
	}

	found := false

	enumerator(func(item T) bool {
		if comparer.Equals(item, value) {
			found = true
			return false
		}
		return true
	})

	return found
}
