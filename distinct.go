package enumerable

// Distinct returns an enumerator that yields only unique elements from the original enumeration.
// Each element appears only once in the result, regardless of how many times it appears in the source.
//
// The distinct operation will:
// - Yield each unique element exactly once
// - Preserve the order of first occurrence of each element
// - Use equality comparison (==) to determine uniqueness
// - Support early termination when consumer returns false
//
// Returns:
//
//	An Enumerator[T] that yields unique elements in order of first appearance
//
// ⚠️ Performance note: This operation buffers all unique elements encountered
// so far in memory. For enumerations with many unique elements, memory usage
// can become significant. The operation is not memory-bounded.
//
// Notes:
// - Requires T to be comparable (supports == operator)
// - Uses map[T]bool internally for tracking seen elements
// - Memory usage grows with number of unique elements
// - For nil enumerators, returns empty enumerator
// - Lazy evaluation - elements processed during iteration
// - Elements are compared using Go's built-in equality
func (q Enumerator[T]) Distinct() Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return func(yield func(T) bool) {
		seen := make(map[T]bool)

		q(func(item T) bool {
			if !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
