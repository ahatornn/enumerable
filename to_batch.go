package enumerable

// ToBatch groups elements from an enumeration into batches of specified size
// and returns them as a slice of slices. This operation is useful for processing
// large sequences in chunks, implementing pagination, or optimizing batch operations.
//
// The to batch operation will:
//   - Group consecutive elements into slices of at most batchSize elements
//   - Return all batches as a slice when enumeration is complete
//   - Handle edge cases gracefully (batchSize <= 0, empty source)
//   - Materialize the entire enumeration (terminal operation)
//
// Parameters:
//
//	batchSize - the maximum number of elements per batch (must be positive)
//
// Returns:
//
//	A slice containing batches of elements, or empty slice if no batches
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to create all batches.
//
// ⚠️ Memory note: This operation buffers all elements in memory during processing.
//
// Notes:
//   - If batchSize <= 0, returns empty slice
//   - If the original enumerator is nil, returns empty slice
//   - Last batch may contain fewer elements than batchSize
//   - Terminal operation - consumes entire enumeration
//   - Each batch is a separate slice - safe for modification
//   - No elements are lost - all elements from source are included in batches
func (e Enumerator[T]) ToBatch(batchSize int) [][]T {
	return toBatchInternal(e, batchSize)
}

// ToBatch groups elements from an enumeration into batches of specified size
// and returns them as a slice of slices. This operation is useful for processing
// large sequences in chunks, implementing pagination, or optimizing batch operations.
// This method supports any type T, including non-comparable types.
//
// The to batch operation will:
//   - Group consecutive elements into slices of at most batchSize elements
//   - Return all batches as a slice when enumeration is complete
//   - Handle edge cases gracefully (batchSize <= 0, empty source)
//   - Materialize the entire enumeration (terminal operation)
//
// Parameters:
//
//	batchSize - the maximum number of elements per batch (must be positive)
//
// Returns:
//
//	A slice containing batches of elements, or empty slice if no batches
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to create all batches.
//
// ⚠️ Memory note: This operation buffers all elements in memory during processing.
//
// Notes:
//   - If batchSize <= 0, returns empty slice
//   - If the original enumerator is nil, returns empty slice
//   - Last batch may contain fewer elements than batchSize
//   - Terminal operation - consumes entire enumeration
//   - Each batch is a separate slice - safe for modification
//   - No elements are lost - all elements from source are included in batches
//   - Supports complex types including structs with non-comparable fields
func (e EnumeratorAny[T]) ToBatch(batchSize int) [][]T {
	return toBatchInternal(e, batchSize)
}

// ToBatch groups elements from a sorted enumeration into batches of specified size
// and returns them as a slice of slices in sorted order. This operation is useful
// for processing large sorted sequences in chunks, implementing pagination,
// or optimizing batch operations on sorted data.
//
// The to batch operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Group consecutive elements in sorted order into slices of at most batchSize elements
//   - Return all batches as a slice when enumeration is complete
//   - Handle edge cases gracefully (batchSize <= 0, empty source)
//   - Materialize the entire sorted enumeration (terminal operation)
//
// Parameters:
//
//	batchSize - the maximum number of elements per batch (must be positive)
//
// Returns:
//
//	A slice containing batches of elements in sorted order, or empty slice if no batches
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the entire sorted enumeration to create all batches.
// Time complexity: O(n log n) for sorting + O(n) for batching = O(n log n)
// where n is the number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus additional buffering for batching. Space complexity: O(n) where n is total elements.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// actual sorting computation occurs immediately upon calling this method.
// All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting.
//
// Notes:
//   - If batchSize <= 0, returns empty slice
//   - If the original enumerator is nil, returns empty slice
//   - Last batch may contain fewer elements than batchSize
//   - Terminal operation - consumes entire sorted enumeration
//   - Each batch is a separate slice - safe for modification
//   - No elements are lost - all elements from source are included in batches
//   - Actual sorting computation occurs only during this operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Elements are grouped in sorted order according to all accumulated sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (e OrderEnumerator[T]) ToBatch(batchSize int) [][]T {
	return toBatchInternal(e.getSortedEnumerator(), batchSize)
}

// ToBatch groups elements from a sorted enumeration into batches of specified size
// and returns them as a slice of slices in sorted order. This operation is useful
// for processing large sorted sequences in chunks, implementing pagination,
// or optimizing batch operations on sorted data.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The to batch operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Group consecutive elements in sorted order into slices of at most batchSize elements
//   - Return all batches as a slice when enumeration is complete
//   - Handle edge cases gracefully (batchSize <= 0, empty source)
//   - Materialize the entire sorted enumeration (terminal operation)
//
// Parameters:
//
//	batchSize - the maximum number of elements per batch (must be positive)
//
// Returns:
//
//	A slice containing batches of elements in sorted order, or empty slice if no batches
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation
// and must iterate through the entire sorted enumeration to create all batches.
// Time complexity: O(n log n) for sorting + O(n) for batching = O(n log n)
// where n is the number of elements.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting,
// plus additional buffering for batching. Space complexity: O(n) where n is total elements.
//
// ⚠️ Evaluation note: This operation is not lazy in the traditional sense -
// actual sorting computation occurs immediately upon calling this method.
// All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting.
//
// Notes:
//   - If batchSize <= 0, returns empty slice
//   - If the original enumerator is nil, returns empty slice
//   - Last batch may contain fewer elements than batchSize
//   - Terminal operation - consumes entire sorted enumeration
//   - Each batch is a separate slice - safe for modification
//   - No elements are lost - all elements from source are included in batches
//   - Actual sorting computation occurs only during this operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Elements are grouped in sorted order according to all accumulated sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
//   - Custom comparer functions are used for element comparison during sorting
func (e OrderEnumeratorAny[T]) ToBatch(batchSize int) [][]T {
	return toBatchInternal(e.getSortedEnumerator(), batchSize)
}

func toBatchInternal[T any](enumerator func(func(T) bool), batchSize int) [][]T {
	if enumerator == nil || batchSize <= 0 {
		return [][]T{}
	}
	batches := make([][]T, 0)
	buffer := make([]T, batchSize)
	index := 0

	enumerator(func(item T) bool {
		buffer[index] = item
		index++
		if index == batchSize {
			batchCopy := make([]T, batchSize)
			copy(batchCopy, buffer)
			index = 0
			batches = append(batches, batchCopy)
		}
		return true
	})

	if index > 0 {
		batchCopy := make([]T, index)
		copy(batchCopy, buffer[:index])
		batches = append(batches, batchCopy)
	}

	return batches
}
