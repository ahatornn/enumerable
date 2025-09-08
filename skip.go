package enumerable

// Skip bypasses a specified number of elements in an enumeration and then yields the remaining elements.
// This operation is useful for pagination, skipping headers, or bypassing initial elements.
//
// The skip operation will:
//   - Bypass the first n elements from the enumeration
//   - Yield all remaining elements in order
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields elements after skipping the first n elements
//
// Notes:
//   - If n <= 0, returns the original enumerator unchanged
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator (not nil)
//   - Lazy evaluation - elements are processed and skipped during iteration
//   - No elements are buffered - memory efficient
//   - Negative values of n are treated as 0
//   - The enumeration is consumed sequentially, so skipped elements are still processed
func (q Enumerator[T]) Skip(n int) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return skipInternal(q, n)
}

// Skip bypasses a specified number of elements in an enumeration and then yields the remaining elements.
// This operation is useful for pagination, skipping headers, or bypassing initial elements.
//
// The skip operation will:
//   - Bypass the first n elements from the enumeration
//   - Yield all remaining elements in order
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements after skipping the first n elements
//
// Notes:
//   - If n <= 0, returns the original enumerator unchanged
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator (not nil)
//   - Lazy evaluation - elements are processed and skipped during iteration
//   - No elements are buffered - memory efficient
//   - Negative values of n are treated as 0
//   - The enumeration is consumed sequentially, so skipped elements are still processed
func (q EnumeratorAny[T]) Skip(n int) EnumeratorAny[T] {
	if q == nil {
		return EmptyAny[T]()
	}
	return skipInternal(q, n)
}

// Skip bypasses a specified number of elements in a sorted enumeration and then yields the remaining elements
// in sorted order. This operation is useful for pagination, skipping headers, or bypassing initial elements
// in the context of sorted sequences.
//
// The skip operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Bypass the first n elements from the sorted enumeration
//   - Yield all remaining elements in sorted order
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields elements in sorted order after skipping the first n elements
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The skip operation itself is O(1) - elements are skipped during iteration.
// For large enumerations, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If n <= 0, returns an enumerator that yields all elements in sorted order
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator (not nil)
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and skipped during iteration
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - Negative values of n are treated as 0
//   - The enumeration is consumed sequentially, so skipped elements are still processed during sorting
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) Skip(n int) Enumerator[T] {
	return skipInternal(o.getSortedEnumerator(), n)
}

// Skip bypasses a specified number of elements in a sorted enumeration and then yields the remaining elements
// in sorted order. This operation is useful for pagination, skipping headers, or bypassing initial elements
// in the context of sorted sequences.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The skip operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Bypass the first n elements from the sorted enumeration
//   - Yield all remaining elements in sorted order
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements in sorted order after skipping the first n elements
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The skip operation itself is O(1) - elements are skipped during iteration.
// For large enumerations, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If n <= 0, returns an enumerator that yields all elements in sorted order
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator (not nil)
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and skipped during iteration
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - Negative values of n are treated as 0
//   - The enumeration is consumed sequentially, so skipped elements are still processed during sorting
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) Skip(n int) EnumeratorAny[T] {
	return skipInternal(o.getSortedEnumerator(), n)
}

func skipInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	if n <= 0 {
		return enumerator
	}
	return func(yield func(T) bool) {
		var skipped int
		enumerator(func(item T) bool {
			if skipped < n {
				skipped++
				return true
			}
			return yield(item)
		})
	}
}
