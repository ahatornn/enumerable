package enumerable

import (
	"github.com/ahatornn/enumerable/comparer"
)

// SingleBy returns the single element of a sequence that satisfies the provided equality comparer.
// This operation is useful when you expect exactly one element in the sequence and need custom equality logic.
//
// The SingleBy operation will:
//   - Return the single element if the sequence contains exactly one distinct element according to the comparer
//   - Return an error if the sequence is empty ("sequence contains no elements")
//   - Return an error if the sequence contains more than one distinct element according to the comparer
//   - Process elements sequentially until completion or second distinct element found
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	The single element of the sequence and nil error if successful,
//	zero value of T and error otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until completion or second distinct element found.
// For large sequences with multiple distinct elements, this may process more elements than necessary.
// Average time complexity: O(n) where n is the number of elements processed.
//
// ⚠️ Memory note: This operation buffers elements to perform equality comparison using hash-based lookup.
// Memory usage depends on the number of distinct elements encountered during processing.
// Worst case memory complexity: O(n) where n is the number of distinct elements.
//
// Notes:
//   - If the enumerator is nil, returns zero value and error ("sequence contains no elements")
//   - If the enumeration is empty, returns zero value and error ("sequence contains no elements")
//   - If the enumeration contains more than one distinct element, returns zero value and error ("sequence contains more than one element")
//   - Uses the provided comparer's Equals and GetHashCode methods for efficient element comparison
//   - Elements are buffered in hash buckets during comparison for efficient duplicate detection
//   - This is a terminal operation that materializes the enumeration
//   - Works with any type T (including non-comparable types) when used with appropriate comparer
//   - The comparer functions must be consistent and deterministic
//   - For large enumerations, processing stops early when second distinct element is found
//   - This method should be used when you are certain the sequence contains exactly one distinct element
//   - Common use cases include lookup by custom equality logic, complex object comparison, or unique constraint validation
func (e Enumerator[T]) SingleBy(comparer comparer.EqualityComparer[T]) (T, error) {
	return singleByInternal(e, comparer)
}

// SingleBy returns the single element of a sequence that satisfies the provided equality comparer.
// This operation is useful when you expect exactly one element in the sequence and need custom equality logic.
//
// The SingleBy operation will:
//   - Return the single element if the sequence contains exactly one distinct element according to the comparer
//   - Return an error if the sequence is empty ("sequence contains no elements")
//   - Return an error if the sequence contains more than one distinct element according to the comparer
//   - Process elements sequentially until completion or second distinct element found
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	The single element of the sequence and nil error if successful,
//	zero value of T and error otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the sequence until completion or second distinct element found.
// For large sequences with multiple distinct elements, this may process more elements than necessary.
// Average time complexity: O(n) where n is the number of elements processed.
//
// ⚠️ Memory note: This operation buffers elements to perform equality comparison using hash-based lookup.
// Memory usage depends on the number of distinct elements encountered during processing.
// Worst case memory complexity: O(n) where n is the number of distinct elements.
//
// Notes:
//   - If the enumerator is nil, returns zero value and error ("sequence contains no elements")
//   - If the enumeration is empty, returns zero value and error ("sequence contains no elements")
//   - If the enumeration contains more than one distinct element, returns zero value and error ("sequence contains more than one element")
//   - Uses the provided comparer's Equals and GetHashCode methods for efficient element comparison
//   - Elements are buffered in hash buckets during comparison for efficient duplicate detection
//   - This is a terminal operation that materializes the enumeration
//   - Works with any type T (including non-comparable types) when used with appropriate comparer
//   - The comparer functions must be consistent and deterministic
//   - For large enumerations, processing stops early when second distinct element is found
//   - This method should be used when you are certain the sequence contains exactly one distinct element
//   - Common use cases include lookup by custom equality logic, complex object comparison, or unique constraint validation
func (e EnumeratorAny[T]) SingleBy(comparer comparer.EqualityComparer[T]) (T, error) {
	return singleByInternal(e, comparer)
}

// SingleBy returns the single element of a sorted sequence that satisfies the provided equality comparer.
// This operation is useful when you expect exactly one element in the sorted sequence and need custom equality logic.
//
// The SingleBy operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return the single element if the sequence contains exactly one distinct element according to the comparer
//   - Return an error if the sequence is empty ("sequence contains no elements")
//   - Return an error if the sequence contains more than one distinct element according to the comparer
//   - Process elements sequentially until completion or second distinct element found
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	The single element of the sequence and nil error if successful,
//	zero value of T and error otherwise
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the sequence until completion or second distinct element found.
// For large sequences with multiple distinct elements, this may process more elements than necessary.
// Time complexity: O(n log n) for sorting + O(n) for validation = O(n log n)
// where n is the number of elements processed.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// and buffers elements to perform equality comparison using hash-based lookup.
// Memory usage depends on the number of distinct elements encountered during processing.
// Space complexity: O(n) - allocates memory for all elements during sorting and comparison.
//
// Notes:
//   - If the enumerator is nil, returns zero value and error ("sequence contains no elements")
//   - If the enumeration is empty, returns zero value and error ("sequence contains no elements")
//   - If the enumeration contains more than one distinct element, returns zero value and error ("sequence contains more than one element")
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Uses the provided comparer's Equals and GetHashCode methods for efficient element comparison
//   - Elements are buffered in hash buckets during comparison for efficient duplicate detection
//   - This is a terminal operation that materializes the enumeration
//   - Works with comparable types when used with appropriate comparer
//   - The comparer functions must be consistent and deterministic
//   - For large enumerations, processing stops early when second distinct element is found
//   - This method should be used when you are certain the sequence contains exactly one distinct element
//   - Common use cases include lookup by custom equality logic, complex object comparison, or unique constraint validation
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) SingleBy(comparer comparer.EqualityComparer[T]) (T, error) {
	return singleByInternal(o.getSortedEnumerator(), comparer)
}

// SingleBy returns the single element of a sorted sequence that satisfies the provided equality comparer.
// This operation is useful when you expect exactly one element in the sorted sequence and need custom equality logic.
// This method supports any type T, including non-comparable types with custom equality logic.
//
// The SingleBy operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Return the single element if the sequence contains exactly one distinct element according to the comparer
//   - Return an error if the sequence is empty ("sequence contains no elements")
//   - Return an error if the sequence contains more than one distinct element according to the comparer
//   - Process elements sequentially until completion or second distinct element found
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	The single element of the sequence and nil error if successful,
//	zero value of T and error otherwise
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the sequence until completion or second distinct element found.
// For large sequences with multiple distinct elements, this may process more elements than necessary.
// Time complexity: O(n log n) for sorting + O(n) for validation = O(n log n)
// where n is the number of elements processed.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// and buffers elements to perform equality comparison using hash-based lookup.
// Memory usage depends on the number of distinct elements encountered during processing.
// Space complexity: O(n) - allocates memory for all elements during sorting and comparison.
//
// Notes:
//   - If the enumerator is nil, returns zero value and error ("sequence contains no elements")
//   - If the enumeration is empty, returns zero value and error ("sequence contains no elements")
//   - If the enumeration contains more than one distinct element, returns zero value and error ("sequence contains more than one element")
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Uses the provided comparer's Equals and GetHashCode methods for efficient element comparison
//   - Elements are buffered in hash buckets during comparison for efficient duplicate detection
//   - This is a terminal operation that materializes the enumeration
//   - Works with any type T (including non-comparable types) when used with appropriate comparer
//   - The comparer functions must be consistent and deterministic
//   - For large enumerations, processing stops early when second distinct element is found
//   - This method should be used when you are certain the sequence contains exactly one distinct element
//   - Common use cases include lookup by custom equality logic, complex object comparison, or unique constraint validation
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) SingleBy(comparer comparer.EqualityComparer[T]) (T, error) {
	return singleByInternal(o.getSortedEnumerator(), comparer)
}

func singleByInternal[T any](enumerator func(func(T) bool), comparer comparer.EqualityComparer[T]) (T, error) {
	var result T
	count := 0

	seenHashes := make(map[uint64][]T)

	if enumerator == nil {
		var zero T
		return zero, ErrNoElements
	}

	enumerator(func(item T) bool {
		hash := comparer.GetHashCode(item)
		if bucket, exists := seenHashes[hash]; exists {
			for _, existing := range bucket {
				if comparer.Equals(item, existing) {
					return true
				}
			}
		}
		seenHashes[hash] = append(seenHashes[hash], item)
		count++

		if count > 1 {
			return false
		}

		result = item
		return true
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
