package enumerable

import (
	"github.com/ahatornn/enumerable/comparer"
)

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

// Intersect returns an enumerator that yields elements present in both enumerations
// using the provided equality comparer. This is equivalent to set intersection operation (first ∩ second).
//
// The intersect operation will:
//   - Yield elements that exist in both the first and second enumerations according to the comparer
//   - Remove duplicates from the result (each element appears only once)
//   - Preserve the order of first occurrence from the first enumeration
//   - Handle nil enumerators gracefully
//   - Use the provided equality comparer for element comparison
//
// Parameters:
//
//	second - the enumerator to intersect with
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements present in both enumerations
//
// ⚠️ Performance note: The second enumeration is completely loaded into memory
// to enable fast lookups using the provided comparer. Be cautious when using this
// with very large second enumerations as it may cause high memory usage.
// Time complexity: O(n + m) where n is first enumeration size, m is second enumeration size.
// Space complexity: O(m) for buffering second enumeration elements.
//
// ⚠️ Memory note: This operation buffers all elements from the second enumeration
// in memory during intersection computation. Memory usage scales with the number
// of unique elements in the second enumeration. For large second enumerations,
// this may consume significant memory.
//
// Notes:
//   - Uses the provided equality comparer's Equals and GetHashCode methods for element comparison
//   - Result contains only unique elements (duplicates removed)
//   - For nil first enumerator, returns empty enumeration
//   - For nil second enumerator, returns empty enumeration
//   - For nil comparer, returns empty enumeration
//   - Lazy evaluation - elements processed during iteration
//   - Memory usage depends on size of second enumeration and unique elements in first
//   - Elements are yielded in order of their first appearance in the first enumeration
//   - Hash codes are used for efficient O(1) average case lookup performance
//   - Hash collisions are resolved using exact equality comparison
//   - Supports any type T when used with appropriate comparer
//   - Works with non-comparable types through custom equality logic
//   - Stable intersection preserves relative order of equal elements
//   - The enumeration stops when first enumeration is exhausted or consumer returns false
//   - No elements are buffered beyond intersection computation
//   - This is a terminal operation that materializes the second enumeration
//   - All accumulated intersection rules are applied during computation
//   - Subsequent calls will re-execute intersection (intersection is not cached)
//   - Thread safe for rule accumulation, but enumeration should be single-threaded
func (q EnumeratorAny[T]) Intersect(second EnumeratorAny[T], comparer comparer.EqualityComparer[T]) EnumeratorAny[T] {
	return func(yield func(T) bool) {
		secondSet := newHashSet(comparer)
		if second != nil {
			second(func(item T) bool {
				if !secondSet.contains(item) {
					secondSet.add(item)
				}
				return true
			})
		}

		if q == nil {
			return
		}

		seen := newHashSet(comparer)
		q(func(item T) bool {
			if secondSet.contains(item) && !seen.contains(item) {
				seen.add(item)
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
