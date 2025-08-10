package enumerable

// SkipLast bypasses a specified number of elements at the end of an enumeration
// and yields the remaining elements.
// This operation is useful for removing trailing elements like footers, summaries,
// or fixed-size endings from sequences.
//
// The skip last operation will:
// - Buffer elements to track which are the final n elements
// - Yield elements only after confirming they are not among the last n
// - Handle edge cases gracefully (n <= 0, n >= count)
// - Support early termination when consumer returns false
//
// Parameters:
//   n - the number of elements to skip from the end (must be non-negative)
//
// Returns:
//   An Enumerator[T] that yields elements except the last n elements
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
// - If n <= 0, returns the original enumerator unchanged
// - If n >= total number of elements, returns an empty enumerator
// - If the original enumerator is nil, returns an empty enumerator
// - Elements are yielded in order of their appearance in the original enumeration
// - Negative values of n are treated as 0
func (q Enumerator[T]) SkipLast(n int) Enumerator[T] {
	if n <= 0 {
		return q
	}
	if q == nil {
		return Empty[T]()
	}

	return func(yield func(T) bool) {
		buffer := make([]T, n)
		count := 0
		index := 0

		q(func(item T) bool {
			if n <= 0 {
				return yield(item)
			}

			if count < n {
				buffer[index] = item
				index = (index + 1) % n
				count++
				return true
			}

			oldestIndex := index
			oldest := buffer[oldestIndex]

			buffer[oldestIndex] = item
			index = (index + 1) % n

			return yield(oldest)
		})
	}
}
