package enumerable

// Intersect returns an enumerator that yields elements present in both enumerations.
// This is equivalent to set intersection operation (first ∩ second).
//
// The intersect operation will:
//   - Yield elements that exist in both the first and second enumerations
//   - Remove duplicates from the result (each element appears only once)
//   - Preserve the order of first occurrence from the first enumeration
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	second - the enumerator to intersect with
//
// Returns:
//
//	An Enumerator[T] that yields elements present in both enumerations
//
// ⚠️ Performance note: The second enumeration is completely loaded into memory
// to enable fast lookups. Be cautious when using this with very large second
// enumerations as it may cause high memory usage.
//
// Notes:
//   - Requires T to be comparable (supports == operator)
//   - Uses map[T]bool internally for efficient lookup
//   - Result contains only unique elements (duplicates removed)
//   - For nil first enumerator, returns empty enumeration
//   - For nil second enumerator, returns empty enumeration
//   - Lazy evaluation - elements processed during iteration
//   - Memory usage depends on size of second enumeration and unique elements in first
//   - Elements are yielded in order of their first appearance in the first enumeration
func (q Enumerator[T]) Intersect(second Enumerator[T]) Enumerator[T] {
	return func(yield func(T) bool) {
		secondSet := make(map[T]bool)
		if second != nil {
			second(func(item T) bool {
				secondSet[item] = true
				return true
			})
		}

		if q == nil {
			return
		}

		seen := make(map[T]bool)
		q(func(item T) bool {
			if secondSet[item] && !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
