package enumerable

// TakeLast returns a specified number of contiguous elements from the end of an enumeration.
// This operation is useful for getting the final elements, such as last N records,
// trailing averages, or end-of-sequence markers.
//
// The take last operation will:
//   - Buffer elements to track the last n elements seen so far
//   - Yield the final n elements when enumeration completes
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take from the end (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields the last n elements
//
// ⚠️ Performance note: This operation buffers up to n elements in memory
// to track which elements should be yielded. For large values of n,
// this may consume significant memory. All elements must be processed
// before any are yielded.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// the entire source enumeration must be consumed before yielding begins.
// Elements are yielded in order of their appearance in the original enumeration.
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements
//   - If the original enumerator is nil, returns an empty enumerator
//   - Negative values of n are treated as 0
//   - The enumeration stops as soon as the consumer returns false
func (q Enumerator[T]) TakeLast(n int) Enumerator[T] {
	if q == nil || n <= 0 {
		return Empty[T]()
	}

	return takeLastInternal(q, n)
}

// TakeLast returns a specified number of contiguous elements from the end of an enumeration.
// This operation is useful for getting the final elements, such as last N records,
// trailing averages, or end-of-sequence markers.
//
// The take last operation will:
//   - Buffer elements to track the last n elements seen so far
//   - Yield the final n elements when enumeration completes
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take from the end (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields the last n elements
//
// ⚠️ Performance note: This operation buffers up to n elements in memory
// to track which elements should be yielded. For large values of n,
// this may consume significant memory. All elements must be processed
// before any are yielded.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// the entire source enumeration must be consumed before yielding begins.
// Elements are yielded in order of their appearance in the original enumeration.
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements
//   - If the original enumerator is nil, returns an empty enumerator
//   - Negative values of n are treated as 0
//   - The enumeration stops as soon as the consumer returns false
func (q EnumeratorAny[T]) TakeLast(n int) EnumeratorAny[T] {
	if q == nil || n <= 0 {
		return EmptyAny[T]()
	}

	return takeLastInternal(q, n)
}

// TakeLast returns a specified number of contiguous elements from the end of a sorted enumeration
// in sorted order. This operation is useful for getting the final elements, such as last N records,
// trailing averages, or end-of-sequence markers from sorted sequences.
//
// The take last operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Buffer elements to track the last n elements seen so far in sorted order
//   - Yield the final n elements when enumeration completes
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take from the end (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields the last n elements in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// This operation buffers up to n elements in memory to track which elements should be yielded.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// For large values of n, this may consume significant memory. All elements must be processed
// before any are yielded.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// the entire source enumeration must be consumed and sorted before yielding begins.
// Elements are yielded in sorted order according to all accumulated sorting rules.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus additional buffering of up to n elements for the take operation.
// Space complexity: O(m + n) where m is total elements and n is take count.
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements in sorted order
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Negative values of n are treated as 0
//   - The enumeration stops as soon as the consumer returns false
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) TakeLast(n int) Enumerator[T] {
	if n <= 0 {
		return Empty[T]()
	}

	return takeLastInternal(o.getSortedEnumerator(), n)
}

// TakeLast returns a specified number of contiguous elements from the end of a sorted enumeration
// in sorted order. This operation is useful for getting the final elements, such as last N records,
// trailing averages, or end-of-sequence markers from sorted sequences.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The take last operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Buffer elements to track the last n elements seen so far in sorted order
//   - Yield the final n elements when enumeration completes
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to take from the end (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields the last n elements in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// This operation buffers up to n elements in memory to track which elements should be yielded.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// For large values of n, this may consume significant memory. All elements must be processed
// before any are yielded.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// the entire source enumeration must be consumed and sorted before yielding begins.
// Elements are yielded in sorted order according to all accumulated sorting rules.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus additional buffering of up to n elements for the take operation.
// Space complexity: O(m + n) where m is total elements and n is take count.
//
// Notes:
//   - If n <= 0, returns an empty enumerator
//   - If n >= total number of elements, returns all available elements in sorted order
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Negative values of n are treated as 0
//   - The enumeration stops as soon as the consumer returns false
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) TakeLast(n int) EnumeratorAny[T] {
	if n <= 0 {
		return EmptyAny[T]()
	}

	return takeLastInternal(o.getSortedEnumerator(), n)
}

func takeLastInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	return func(yield func(T) bool) {
		buffer := make([]T, n)
		count := 0
		index := 0

		enumerator(func(item T) bool {
			buffer[index] = item
			index = (index + 1) % n
			if count < n {
				count++
			}
			return true
		})

		if count == 0 {
			return
		}

		startIndex := (index - count + n) % n
		for i := 0; i < count; i++ {
			itemIndex := (startIndex + i) % n
			if !yield(buffer[itemIndex]) {
				return
			}
		}
	}
}
