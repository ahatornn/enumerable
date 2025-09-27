package enumerable

import (
	"github.com/ahatornn/enumerable/comparer"
	"github.com/ahatornn/enumerable/grouping"
)

// GroupBy groups the elements of an enumeration according to a specified key selector function.
// This method transforms a flat sequence into a sequence of groups, where each group contains
// all elements that share the same key value.
//
// The GroupBy operation:
//   - Groups elements by the key returned from the keySelector function
//   - Returns an EnumeratorAny containing Group[any, T] objects (one per unique key)
//   - Uses comparer.Default[any]() internally for key comparison and hashing
//   - Preserves the order of first occurrence for groups (order of first element with each key)
//   - Maintains insertion order of elements within each group
//   - Performs immediate execution (eager) — all grouping happens during enumeration
//
// Parameters:
//
//	keySelector - a function that extracts a key from each element of type T
//	              The function should return any type (including non-comparable types like slices or structs)
//	              Keys are compared using comparer.Default[any]() which handles most common types
//	              For custom key comparison logic, consider using GroupByWithComparer
//
// Returns:
//
//	An EnumeratorAny[grouping.Group[any, T]] that yields groups of elements sharing the same key
//	Each Group contains:
//	  - Key() — the grouping key (type any)
//	  - Items() — slice of all elements in this group (type []T, in order of appearance)
//
// ⚠️ Performance note: This is an eager operation that fully materializes all groups during
// first enumeration. Time complexity: O(n) for grouping, where n is number of elements.
// Memory complexity: O(n + k) where k is number of unique keys. Be cautious with large
// datasets as all elements are loaded into memory during grouping.
//
// Notes:
//   - This is an eager operation — grouping occurs immediately during enumeration
//   - Groups are yielded in order of first occurrence of each key in the source enumeration
//   - Elements within each group maintain their original relative order
//   - Empty groups are never created — each Group contains at least one element
//   - For nil input enumerator or nil keySelector, returns empty enumeration
//   - The returned EnumeratorAny shares internal storage with groups — do not modify Group.Items() slices
//   - Key type 'any' supports slices, maps, structs, and other non-comparable types
//   - Uses comparer.Default[any]() which provides reasonable defaults for common types
//   - For production use with complex keys, consider implementing custom comparer and using GroupByWithComparer
//   - Thread-safe for single enumeration, but not designed for concurrent enumeration
func (e Enumerator[T]) GroupBy(keySelector func(T) any) EnumeratorAny[grouping.Group[any, T]] {
	if e == nil || keySelector == nil {
		return EmptyAny[grouping.Group[any, T]]()
	}

	return func(yield func(grouping.Group[any, T]) bool) {
		builder := grouping.NewGroupingBuilder[any, T](comparer.Default[any]())
		e(func(item T) bool {
			key := keySelector(item)
			builder.Add(key, item)
			return true
		})

		groups := builder.Result()
		for _, group := range groups {
			if !yield(group) {
				break
			}
		}
	}
}
