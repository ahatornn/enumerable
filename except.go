package enumerable

import (
	"github.com/ahatornn/enumerable/comparer"
)

// Except returns an enumerator that yields elements from the first enumeration
// that are not present in the second enumeration.
// This is equivalent to set difference operation (first - second).
//
// The except operation will:
//   - Yield elements that exist in the first enumeration but not in the second
//   - Remove duplicates from the result (each element appears only once)
//   - Preserve the order of first occurrence from the first enumeration
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	second - the enumerator containing elements to exclude
//
// Returns:
//
//	An Enumerator[T] that yields elements from first enumeration not in second
//
// ⚠️ Performance note: This operation completely buffers the `second` enumerator
// into memory (creates a map for fast lookup). For large second enumerations,
// this may consume significant memory. The memory usage is proportional to
// the number of unique elements in the second enumerator.
//
// Notes:
//   - Requires T to be comparable (supports == operator)
//   - Uses map[T]bool internally for efficient lookup
//   - Result contains only unique elements (duplicates removed)
//   - For nil first enumerator, returns empty enumeration
//   - For nil second enumerator, returns distinct elements from first
//   - Lazy evaluation - elements processed during iteration
//   - Memory usage depends on size of second enumeration and unique elements in first
func (q Enumerator[T]) Except(second Enumerator[T]) Enumerator[T] {
	return func(yield func(T) bool) {
		excludeSet := make(map[T]bool)
		if second != nil {
			second(func(item T) bool {
				excludeSet[item] = true
				return true
			})
		}

		if q == nil {
			return
		}

		seen := make(map[T]bool)
		q(func(item T) bool {
			if !excludeSet[item] && !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}

// Except returns an enumerator that yields elements from the first enumeration
// that are not present in the second enumeration using the provided equality comparer.
// This is equivalent to set difference operation (first - second).
//
// The except operation will:
//   - Yield elements that exist in the first enumeration but not in the second
//   - Remove duplicates from the result (each element appears only once)
//   - Preserve the order of first occurrence from the first enumeration
//   - Handle nil enumerators gracefully
//   - Use the provided equality comparer for element comparison
//
// Parameters:
//
//	second - the enumerator containing elements to exclude
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	An Enumerator[T] that yields elements from first enumeration not in second
//
// ⚠️ Performance note: This operation completely buffers the `second` enumerator
// into memory (creates a hash set for fast lookup). For large second enumerations,
// this may consume significant memory. The memory usage is proportional to
// the number of unique elements in the second enumerator.
//
// ⚠️ Memory note: This operation buffers elements from both enumerations
// for duplicate detection and exclusion checking. Memory usage scales with
// the number of unique elements in both enumerations.
//
// Notes:
//   - Uses the provided equality comparer's Equals and GetHashCode methods
//   - Result contains only unique elements (duplicates removed)
//   - For nil first enumerator, returns empty enumeration
//   - For nil second enumerator, returns distinct elements from first
//   - For nil comparer, returns empty enumeration
//   - Lazy evaluation - elements processed during iteration
//   - Memory usage depends on size of second enumeration and unique elements in first
//   - Supports any type T when used with appropriate comparer
//   - Works with non-comparable types through custom equality logic
//   - Hash codes are used for efficient O(1) average case lookup performance
//   - Hash collisions are resolved using exact equality comparison
func (e EnumeratorAny[T]) Except(second EnumeratorAny[T], comparer comparer.EqualityComparer[T]) EnumeratorAny[T] {
	if e == nil || comparer == nil {
		return EmptyAny[T]()
	}

	return func(yield func(T) bool) {
		excludeSet := newHashSet(comparer)
		if second != nil {
			second(func(item T) bool {
				excludeSet.add(item)
				return true
			})
		}

		if e == nil {
			return
		}

		seenSet := newHashSet(comparer)
		e(func(item T) bool {
			if !excludeSet.contains(item) && !seenSet.contains(item) {
				seenSet.add(item)
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
