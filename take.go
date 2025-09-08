package enumerable

// Take returns a specified number of contiguous elements from the start of an enumeration.
// This operation is useful for pagination, limiting results, or taking samples from sequences.
//
// The take operation will:
//   - Yield the first n elements from the enumeration
//   - Stop enumeration once n elements have been yielded
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields at most n elements from the start
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements
//   - If the original enumerator is nil, returns an empty enumerator
//   - Lazy evaluation - elements are processed and yielded during iteration
//   - No elements are buffered - memory efficient
//   - Negative values of n are treated as 0
//   - Early termination by the consumer stops further enumeration
//   - The enumeration stops as soon as n elements are yielded or the source is exhausted
func (q Enumerator[T]) Take(n int) Enumerator[T] {
	if q == nil || n <= 0 {
		return Empty[T]()
	}
	return takeInternal(q, n)
}

// Take returns a specified number of contiguous elements from the start of an enumeration.
// This operation is useful for pagination, limiting results, or taking samples from sequences.
//
// The take operation will:
//   - Yield the first n elements from the enumeration
//   - Stop enumeration once n elements have been yielded
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields at most n elements from the start
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements
//   - If the original enumerator is nil, returns an empty enumerator
//   - Lazy evaluation - elements are processed and yielded during iteration
//   - No elements are buffered - memory efficient
//   - Negative values of n are treated as 0
//   - Early termination by the consumer stops further enumeration
//   - The enumeration stops as soon as n elements are yielded or the source is exhausted
func (q EnumeratorAny[T]) Take(n int) EnumeratorAny[T] {
	if q == nil || n <= 0 {
		return EmptyAny[T]()
	}
	return takeInternal(q, n)
}

// Take returns a specified number of contiguous elements from the start of a sorted enumeration
// in sorted order. This operation is useful for pagination, limiting results, or taking samples
// from sorted sequences.
//
// The take operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Yield the first n elements from the sorted enumeration
//   - Stop enumeration once n elements have been yielded
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields at most n elements from the start in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The take operation itself is O(1) - elements are taken during iteration.
// For large enumerations where n << total, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements in sorted order
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and yielded during iteration
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - Negative values of n are treated as 0
//   - Early termination by the consumer stops further enumeration
//   - The enumeration stops as soon as n elements are yielded or the source is exhausted
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) Take(n int) Enumerator[T] {
	if n <= 0 {
		return Empty[T]()
	}
	return takeInternal(o.getSortedEnumerator(), n)
}

// Take returns a specified number of contiguous elements from the start of a sorted enumeration
// in sorted order. This operation is useful for pagination, limiting results, or taking samples
// from sorted sequences.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The take operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Yield the first n elements from the sorted enumeration
//   - Stop enumeration once n elements have been yielded
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields at most n elements from the start in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The take operation itself is O(1) - elements are taken during iteration.
// For large enumerations where n << total, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements in sorted order
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and yielded during iteration
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - Negative values of n are treated as 0
//   - Early termination by the consumer stops further enumeration
//   - The enumeration stops as soon as n elements are yielded or the source is exhausted
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) Take(n int) EnumeratorAny[T] {
	if n <= 0 {
		return EmptyAny[T]()
	}
	return takeInternal(o.getSortedEnumerator(), n)
}

func takeInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	return func(yield func(T) bool) {
		var taken int
		enumerator(func(item T) bool {
			if taken >= n {
				return false
			}
			if !yield(item) {
				return false
			}
			taken++
			return true
		})
	}
}
