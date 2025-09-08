package enumerable

// SkipLast bypasses a specified number of elements at the end of an enumeration
// and yields the remaining elements.
// This operation is useful for removing trailing elements like footers, summaries,
// or fixed-size endings from sequences.
//
// The skip last operation will:
//   - Buffer elements to track which are the final n elements
//   - Yield elements only after confirming they are not among the last n
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip from the end (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields elements except the last n elements
//
// ⚠️ Performance note: This operation buffers up to n elements in memory
// using a circular buffer for efficient memory usage. For large values of n,
// this may consume significant memory.
//
// ⚠️ Evaluation note: This operation is partially lazy - elements are
// processed as the enumeration proceeds, but the last n elements are
// buffered and never yielded. The enumeration must progress beyond
// n elements to yield earlier ones.
//
// Notes:
//   - If n <= 0, returns the original enumerator unchanged
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator
//   - Elements are yielded in order of their appearance in the original enumeration
//   - Negative values of n are treated as 0
func (q Enumerator[T]) SkipLast(n int) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}

	return skipLastInternal(q, n)
}

// SkipLast bypasses a specified number of elements at the end of an enumeration
// and yields the remaining elements.
// This operation is useful for removing trailing elements like footers, summaries,
// or fixed-size endings from sequences.
//
// The skip last operation will:
//   - Buffer elements to track which are the final n elements
//   - Yield elements only after confirming they are not among the last n
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip from the end (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements except the last n elements
//
// ⚠️ Performance note: This operation buffers up to n elements in memory
// using a circular buffer for efficient memory usage. For large values of n,
// this may consume significant memory.
//
// ⚠️ Evaluation note: This operation is partially lazy - elements are
// processed as the enumeration proceeds, but the last n elements are
// buffered and never yielded. The enumeration must progress beyond
// n elements to yield earlier ones.
//
// Notes:
//   - If n <= 0, returns the original enumerator unchanged
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator
//   - Elements are yielded in order of their appearance in the original enumeration
//   - Negative values of n are treated as 0
func (q EnumeratorAny[T]) SkipLast(n int) EnumeratorAny[T] {
	if q == nil {
		return EmptyAny[T]()
	}

	return skipLastInternal(q, n)
}

// SkipLast bypasses a specified number of elements at the end of a sorted enumeration
// and yields the remaining elements in sorted order.
// This operation is useful for removing trailing elements like footers, summaries,
// or fixed-size endings from sorted sequences.
//
// The skip last operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Buffer elements to track which are the final n elements in sorted order
//   - Yield elements only after confirming they are not among the last n in sorted order
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip from the end (must be non-negative)
//
// Returns:
//
//	An Enumerator[T] that yields elements in sorted order except the last n elements
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// This operation buffers up to n elements in memory using a circular buffer for efficient memory usage.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// For large values of n, this may consume significant memory.
//
// ⚠️ Evaluation note: This operation is partially lazy - elements are
// processed as the enumeration proceeds, but the last n elements are
// buffered and never yielded. The enumeration must progress beyond
// n elements to yield earlier ones. Actual sorting computation occurs only during enumeration.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus additional buffering of up to n elements for the skip operation.
// Space complexity: O(m + n) where m is total elements and n is skip count.
//
// Notes:
//   - If n <= 0, returns an enumerator that yields all elements in sorted order
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) SkipLast(n int) Enumerator[T] {
	return skipLastInternal(o.getSortedEnumerator(), n)
}

// SkipLast bypasses a specified number of elements at the end of a sorted enumeration
// and yields the remaining elements in sorted order.
// This operation is useful for removing trailing elements like footers, summaries,
// or fixed-size endings from sorted sequences.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The skip last operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Buffer elements to track which are the final n elements in sorted order
//   - Yield elements only after confirming they are not among the last n in sorted order
//   - Handle edge cases gracefully (n <= 0, n >= count)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	n - the number of elements to skip from the end (must be non-negative)
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements in sorted order except the last n elements
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// This operation buffers up to n elements in memory using a circular buffer for efficient memory usage.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// For large values of n, this may consume significant memory.
//
// ⚠️ Evaluation note: This operation is partially lazy - elements are
// processed as the enumeration proceeds, but the last n elements are
// buffered and never yielded. The enumeration must progress beyond
// n elements to yield earlier ones. Actual sorting computation occurs only during enumeration.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus additional buffering of up to n elements for the skip operation.
// Space complexity: O(m + n) where m is total elements and n is skip count.
//
// Notes:
//   - If n <= 0, returns an enumerator that yields all elements in sorted order
//   - If n >= total number of elements, returns an empty enumerator
//   - If the original enumerator is nil, returns an empty enumerator
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Actual sorting computation occurs only during first enumeration
//   - Custom comparer functions are used for element comparison during sorting
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) SkipLast(n int) EnumeratorAny[T] {
	return skipLastInternal(o.getSortedEnumerator(), n)
}

func skipLastInternal[T any](enumerator func(func(T) bool), n int) func(func(T) bool) {
	if n <= 0 {
		return enumerator
	}
	return func(yield func(T) bool) {
		buffer := make([]T, n)
		count := 0
		index := 0

		enumerator(func(item T) bool {
			if count < n {
				buffer[index] = item
				count++
				index++
				return true
			}

			oldestIndex := index % n
			oldest := buffer[oldestIndex]

			buffer[oldestIndex] = item
			index++

			return yield(oldest)
		})
	}
}
