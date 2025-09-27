package grouping

import (
	"github.com/ahatornn/enumerable/comparer"
)

// Group represents a collection of elements that share the same key.
//
// A Group contains:
//   - A Key: the value by which elements were grouped
//   - A collection of Items: all elements that matched this key during grouping
//
// Groups are typically created by GroupBy operations and should not be constructed manually.
// They are immutable after creation — the key and items cannot be modified.
//
// Type Parameters:
//
//	K - the type of the grouping key (can be any type, including non-comparable types)
//	T - the type of the grouped elements (must be comparable to support Enumerator operations)
//
// Methods:
//
//	Key() K - returns the grouping key
//	Items() []T - returns a slice containing all elements in this group
//
// Notes:
//   - The Items() slice is a direct reference to internal storage — do not modify it
//   - Groups are designed to be consumed, not mutated
//   - Key type K can be any type, including slices, maps, or structs (with appropriate EqualityComparer)
//   - Element type T must be comparable to support operations in Enumerator chains
//   - Groups preserve the order of element insertion (order from original enumeration)
//   - Each Group represents exactly one key and all its associated values
//   - Memory layout is efficient — no unnecessary copying or wrapping
//   - Safe for concurrent read access (but not designed for concurrent modification)
type Group[K any, T comparable] struct {
	key   K
	items []T
}

// Key returns the grouping key associated with this group.
// The key is the value that was returned by the keySelector function during GroupBy.
//
// Returns:
//
//	The grouping key of type K
//
// Notes:
//   - The returned key is the original value, not a copy
//   - If K is a reference type (slice, map, pointer), modifications may affect original data
//   - Key type can be any type, including non-comparable types (when used with appropriate EqualityComparer)
func (gi Group[K, T]) Key() K {
	return gi.key
}

// Items returns a slice containing all elements that belong to this group.
// The elements are in the order they were encountered during the original enumeration.
//
// Returns:
//
//	A slice of type []T containing all elements in this group
//
// ⚠️ Important: The returned slice is a direct reference to internal storage.
// Modifying this slice will affect the Group's internal state.
//
// Notes:
//   - The slice preserves the original order of elements from the source enumeration
//   - Duplicates are preserved — if the same element appeared multiple times, it appears multiple times here
//   - The slice may be empty if no elements matched this key (though GroupBy typically doesn't create empty groups)
//   - For performance reasons, no copy is made — if you need to modify the slice, make a copy first
//   - Element type T must be comparable to support Enumerator operations in chained queries
func (gi Group[K, T]) Items() []T {
	return gi.items
}

type groupingBuilder[K any, V comparable] struct {
	order    []*uniqueKeyEntry[K, V]
	lookup   map[uint64][]*uniqueKeyEntry[K, V]
	comparer comparer.EqualityComparer[K]
}

type uniqueKeyEntry[K any, V comparable] struct {
	key   K
	items []V
}

// NewGroupingBuilder creates a new grouping builder that can accumulate elements into groups
// based on their keys. This is primarily used internally by GroupBy operations and should
// rarely be used directly by end users.
//
// The grouping builder:
//   - Uses the provided EqualityComparer to determine key equality and compute hash codes
//   - Efficiently groups elements by key using hash buckets
//   - Supports any key type K (including non-comparable types like slices or maps)
//   - Requires value type V to be comparable (for compatibility with Enumerator chains)
//   - Maintains insertion order of elements within each group
//
// Parameters:
//
//	comparer - an EqualityComparer[K] that defines how to compare and hash keys of type K
//
// Returns:
//
//	A pointer to a new groupingBuilder[K, V] ready to accept elements via Add()
func NewGroupingBuilder[K any, V comparable](comparer comparer.EqualityComparer[K]) *groupingBuilder[K, V] {
	return &groupingBuilder[K, V]{
		order:    []*uniqueKeyEntry[K, V]{},
		lookup:   make(map[uint64][]*uniqueKeyEntry[K, V]),
		comparer: comparer,
	}
}

// Add inserts a key-value pair into the grouping builder, accumulating values under their respective keys.
// If the key already exists, the value is appended to the existing group. If not, a new group is created.
//
// The operation:
//   - Computes hash code for the key using the builder's EqualityComparer
//   - Searches for existing key in the corresponding hash bucket
//   - Uses comparer.Equals to confirm key equality (hash collisions are resolved by full comparison)
//   - Appends value to existing group or creates new group as needed
//   - Preserves insertion order of values within each group
//
// Parameters:
//
//	key   - the grouping key (type K, can be any type including non-comparable)
//	value - the value to add to the group (type V, must be comparable)
func (gb *groupingBuilder[K, V]) Add(key K, value V) {
	hash := gb.comparer.GetHashCode(key)
	bucket := gb.lookup[hash]

	var entry *uniqueKeyEntry[K, V]
	for _, e := range bucket {
		if gb.comparer.Equals(e.key, key) {
			entry = e
			break
		}
	}

	if entry != nil {
		entry.items = append(entry.items, value)
	} else {
		newEntry := &uniqueKeyEntry[K, V]{
			key:   key,
			items: []V{value},
		}
		gb.lookup[hash] = append(bucket, newEntry)
		gb.order = append(gb.order, newEntry)
	}
}

// Result finalizes the grouping process and returns all accumulated groups as a slice.
// This method should be called after all Add() operations are complete.
//
// The operation:
//   - Iterates through all hash buckets and their entries
//   - Converts each bucketEntry into a Group[K, V] (zero-cost conversion when fields match)
//   - Returns a new slice containing all groups
//   - Does not modify or clear the builder — can be called multiple times (but not recommended)
//
// Returns:
//
//	A slice of type []Group[K, V] containing all groups created by previous Add() calls
func (gb *groupingBuilder[K, V]) Result() []Group[K, V] {
	var result []Group[K, V]
	for _, entry := range gb.order {
		result = append(result, Group[K, V]{
			key:   entry.key,
			items: entry.items,
		})
	}
	return result
}
