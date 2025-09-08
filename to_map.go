package enumerable

// ToMap converts an enumeration to a map[T]struct{} for efficient set-like operations.
// This operation is useful for creating memory-efficient lookup collections or
// for removing duplicates while materializing an enumeration.
//
// The to map operation will:
//   - Iterate through all elements in the enumeration
//   - Add each element as a key in the resulting map with empty struct value
//   - Automatically remove duplicates (map key uniqueness)
//   - Return the resulting map
//   - Handle nil enumerators gracefully by returning an empty map
//
// Returns:
//
//	A map[T]struct{} containing all unique elements as keys,
//	or an empty map if enumerator is nil or enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration. For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation buffers all unique elements in memory.
// The memory usage depends on the number of unique elements.
//
// Notes:
//   - If the enumerator is nil, returns an empty map (not nil)
//   - If the enumeration is empty, returns an empty map
//   - Automatically removes duplicates due to map key uniqueness
//   - Uses map[T]struct{} for memory efficiency (struct{} takes zero memory)
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(k) where k is the number of unique elements
//   - Keys in the returned map are the unique elements from the enumeration
//   - Values in the returned map are empty structs (zero memory footprint)
//   - Check for element presence using: if _, exists := map[key]; exists { ... }
//   - More memory-efficient than map[T]bool since struct{} takes no memory
//   - Always returns a valid map, never nil - safe for immediate use
func (q Enumerator[T]) ToMap() map[T]struct{} {
	if q == nil {
		return make(map[T]struct{})
	}

	result := make(map[T]struct{})
	q(func(item T) bool {
		result[item] = struct{}{}
		return true
	})
	return result
}

// ToMap converts a sorted enumeration to a map[T]struct{} for efficient set-like operations.
// This operation is useful for creating memory-efficient lookup collections or
// for removing duplicates while materializing a sorted enumeration.
//
// The to map operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Iterate through all elements in the sorted enumeration
//   - Add each element as a key in the resulting map with empty struct value
//   - Automatically remove duplicates (map key uniqueness)
//   - Return the resulting map
//   - Handle nil enumerators gracefully by returning an empty map
//
// Returns:
//
//	A map[T]struct{} containing all unique elements as keys in sorted order,
//	or an empty map if enumerator is nil or enumeration is empty
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the entire sorted enumeration.
// Time complexity: O(n log n) for sorting + O(n) for map creation = O(n log n)
// where n is the number of elements. For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus all unique elements in the resulting map.
// Space complexity: O(n) for sorting + O(k) for map where k is unique elements.
//
// ⚠️ Evaluation note: This operation is not lazy - actual sorting computation
// occurs immediately upon calling this method.
// All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting.
//
// Notes:
//   - If the enumerator is nil, returns an empty map (not nil)
//   - If the enumeration is empty, returns an empty map
//   - Automatically removes duplicates due to map key uniqueness
//   - Uses map[T]struct{} for memory efficiency (struct{} takes zero memory)
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n + k) where n is total elements and k is unique elements
//   - Keys in the returned map are the unique elements from the sorted enumeration
//   - Values in the returned map are empty structs (zero memory footprint)
//   - Check for element presence using: if _, exists := map[key]; exists { ... }
//   - More memory-efficient than map[T]bool since struct{} takes no memory
//   - Always returns a valid map, never nil - safe for immediate use
//   - Actual sorting computation occurs only during this operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Elements are added to map in sorted order according to all accumulated sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) ToMap() map[T]struct{} {
	result := make(map[T]struct{})
	o.getSortedEnumerator()(func(item T) bool {
		result[item] = struct{}{}
		return true
	})
	return result
}
