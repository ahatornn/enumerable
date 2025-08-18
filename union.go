package enumerable

// Union produces the set union of two enumerations by using the default equality comparer.
// This operation returns unique elements that appear in either enumeration.
//
// The union operation will:
// - Yield all unique elements from the first enumeration
// - Then yield unique elements from the second enumeration that haven't been seen yet
// - Remove duplicates both within each enumeration and between enumerations
// - Preserve the order of first occurrence of each element
// - Handle nil enumerators gracefully
// - Support early termination when consumer returns false
//
// Parameters:
//
//	second - the enumerator to union with the current one
//
// Returns:
//
//	An Enumerator[T] that yields unique elements from both enumerations
//
// ⚠️ Performance note: This operation buffers all unique elements encountered
// so far in memory. For enumerations with many unique elements, memory usage
// can become significant. The operation is not memory-bounded.
//
// Notes:
//   - Requires T to be comparable (supports == operator)
//   - Uses map[T]bool internally for tracking seen elements
//   - Memory usage grows with number of unique elements from both enumerations
//   - For nil enumerators, treats them as empty enumerations
//   - Lazy evaluation - elements are processed during iteration
//   - Elements are yielded in order: first unique elements from the first enumeration,
//     then unique elements from the second enumeration that weren't in the first
//   - Early termination by the consumer stops further enumeration of both sources
//
// Union produces the set union of two enumerations by using the default equality comparer.
func (q Enumerator[T]) Union(second Enumerator[T]) Enumerator[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]bool)
		stopped := false

		tryYield := func(item T) bool {
			if stopped {
				return false
			}
			if !yield(item) {
				stopped = true
				return false
			}
			return true
		}

		if q != nil {
			q(func(item T) bool {
				if stopped {
					return false
				}
				if !seen[item] {
					seen[item] = true
					return tryYield(item)
				}
				return true
			})
		}

		if second != nil && !stopped {
			second(func(item T) bool {
				if stopped {
					return false
				}
				if !seen[item] {
					seen[item] = true
					return tryYield(item)
				}
				return true
			})
		}
	}
}
