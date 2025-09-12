package enumerable

import (
	"github.com/ahatornn/enumerable/comparer"
)

// Distinct returns an enumerator that yields only unique elements from the original enumeration.
// Each element appears only once in the result, regardless of how many times it appears in the source.
//
// The distinct operation will:
//   - Yield each unique element exactly once
//   - Preserve the order of first occurrence of each element
//   - Use equality comparison (==) to determine uniqueness
//   - Support early termination when consumer returns false
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
//   - Requires T to be comparable (supports == operator)
//   - Uses map[T]bool internally for tracking seen elements
//   - Memory usage grows with number of unique elements
//   - For nil enumerators, returns empty enumerator
//   - Lazy evaluation - elements processed during iteration
//   - Elements are compared using Go's built-in equality
func (q Enumerator[T]) Distinct() Enumerator[T] {
	return distinctInternal(q)
}

// Distinct returns an enumerator that yields only unique elements from the original enumeration
// using a custom equality comparer to determine uniqueness.
// Each element appears only once in the result, regardless of how many times it appears in the source.
//
// The distinct operation will:
//   - Yield each unique element exactly once
//   - Preserve the order of first occurrence of each element
//   - Use the provided EqualityComparer to determine uniqueness
//   - Support early termination when consumer returns false
//
// Returns:
//
//	An EnumeratorAny[T] that yields unique elements in order of first appearance
//
// ⚠️ Performance note: This operation buffers all unique elements encountered
// so far in memory. For enumerations with many unique elements, memory usage
// can become significant. The operation is not memory-bounded.
//
// ⚠️ Performance warning: This operation can be significantly slower than Distinct()
// for comparable types due to the overhead of EqualityComparer interface calls and
// hash-based collision resolution. Consider using Distinct() with comparable types
// when possible for better performance.
//
// ⚠️ Allocation warning: This operation performs multiple memory allocations during
// execution, especially for enumerations with many unique elements. Each unique element
// requires storage in internal hash buckets, leading to increased GC pressure.
//
// Notes:
//   - Uses map[uint64][]T internally for tracking seen elements with collision chaining
//   - Memory usage grows with number of unique elements
//   - For nil enumerators, returns empty enumerator
//   - Lazy evaluation - elements processed during iteration
//   - Elements are compared using the provided EqualityComparer's Equals and GetHashCode methods
//   - The comparer must not be nil, or the operation will panic
//   - Time complexity: O(n) average case, O(n²) worst case with many hash collisions
//   - Space complexity: O(k) where k is the number of unique elements
func (q EnumeratorAny[T]) Distinct(comparer comparer.EqualityComparer[T]) EnumeratorAny[T] {
	return distinctAnyInternal(q, comparer)
}

// Distinct returns an enumerator that yields only unique elements from a sorted sequence.
// This operation is optimized for sorted sequences and yields distinct elements in sorted order.
//
// The Distinct operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Yield each unique element exactly once in sorted order
//   - Preserve the order of elements according to the applied sorting rules
//   - Support early termination when consumer returns false
//   - Process elements sequentially with O(n) time complexity
//
// Returns:
//
//	An Enumerator[T] that yields unique elements in sorted order
//
// ⚠️ Performance note: This is a streaming operation that processes elements sequentially
// and maintains only the previous element for comparison. Time complexity: O(n log n) for sorting
// + O(n) for distinct processing = O(n log n) where n is the number of elements.
//
// ⚠️ Memory note: This operation uses minimal additional memory - only stores the previous
// element for comparison. Memory usage is O(1) for processing, plus memory required for sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - Requires T to be comparable (supports == operator)
//   - Uses Go's built-in equality operator (==) for element comparison
//   - For nil enumerators, returns empty enumerator
//   - Lazy evaluation - elements processed during iteration
//   - Elements are compared using Go's built-in equality after sorting
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - This is a streaming operation that does not buffer all elements in memory
//   - Works efficiently with sorted sequences by comparing only adjacent elements
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Common use cases include removing duplicates from sorted data, unique value extraction,
//     or preparing data for further processing that requires uniqueness
func (o OrderEnumerator[T]) Distinct() Enumerator[T] {
	return distinctInternal(o.getSortedEnumerator())
}

// Distinct returns an enumerator that yields only unique elements from a sorted sequence
// using a custom equality comparer to determine uniqueness.
// This operation is optimized for sorted sequences and yields distinct elements in sorted order.
//
// The Distinct operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Yield each unique element exactly once in sorted order
//   - Preserve the order of elements according to the applied sorting rules
//   - Support early termination when consumer returns false
//   - Process elements sequentially with O(n) time complexity
//
// Parameters:
//
//	comparer - an EqualityComparer that defines when two elements are considered equal
//
// Returns:
//
//	An EnumeratorAny[T] that yields unique elements in sorted order
//
// ⚠️ Performance note: This is a streaming operation that processes elements sequentially
// and maintains only the previous element for comparison. Time complexity: O(n log n) for sorting
// + O(n) for distinct processing = O(n log n) where n is the number of elements.
//
// ⚠️ Performance warning: This operation can be significantly slower than Distinct()
// for comparable types due to the overhead of EqualityComparer interface calls and
// hash-based collision resolution. Consider using Distinct() with comparable types
// when possible for better performance.
//
// ⚠️ Allocation warning: This operation performs multiple memory allocations during
// execution, especially for enumerations with many unique elements. Each unique element
// requires storage in internal hash buckets, leading to increased GC pressure.
//
// ⚠️ Memory note: This operation uses minimal additional memory - only stores the previous
// element for comparison and hash buckets for efficient equality checking.
// Memory usage depends on the comparer implementation and element size.
// Space complexity: O(n) - allocates memory for all elements during sorting.
//
// Notes:
//   - Uses the provided comparer's Equals and GetHashCode methods for element comparison
//   - For nil enumerators, returns empty enumerator
//   - Lazy evaluation - elements processed during iteration
//   - Elements are compared using the provided EqualityComparer after sorting
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - This is a streaming operation that does not buffer all elements in memory
//   - Works efficiently with sorted sequences by comparing only adjacent elements
//   - The comparer must not be nil, or the operation will panic
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Common use cases include removing duplicates from sorted data with custom equality logic,
//     unique value extraction for complex objects, or preparing data for further processing
//   - Elements are buffered in hash buckets during comparison for efficient duplicate detection
//   - This method is particularly useful when default equality comparison is insufficient
//   - Time complexity: O(n log n) average case for sorting + O(n) for distinct processing
//   - Space complexity: O(k) where k is the number of unique elements during distinct processing
func (o OrderEnumeratorAny[T]) Distinct(comparer comparer.EqualityComparer[T]) EnumeratorAny[T] {
	return distinctAnyInternal(o.getSortedEnumerator(), comparer)
}

func distinctInternal[T comparable](enumerator func(func(T) bool)) Enumerator[T] {
	if enumerator == nil {
		return Empty[T]()
	}
	return func(yield func(T) bool) {
		seen := make(map[T]bool)

		enumerator(func(item T) bool {
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

func distinctAnyInternal[T any](enumerator func(func(T) bool), comparer comparer.EqualityComparer[T]) EnumeratorAny[T] {
	if enumerator == nil {
		return EmptyAny[T]()
	}
	return func(yield func(T) bool) {
		seen := make(map[uint64][]T, 8)

		enumerator(func(item T) bool {
			hash := comparer.GetHashCode(item)

			if bucket, exists := seen[hash]; exists {
				for _, existing := range bucket {
					if comparer.Equals(item, existing) {
						return true
					}
				}
			}

			seen[hash] = append(seen[hash], item)
			return yield(item)
		})
	}
}
