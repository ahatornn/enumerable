package enumerable

// ToChannel converts an enumeration to a channel that yields all elements.
// This operation is useful for converting lazy enumerations into channel-based
// processing pipelines or for interoperability with goroutines.
//
// The to channel operation will:
//   - Create a new channel with specified buffer size
//   - Start a goroutine that iterates through the enumeration
//   - Send each element to the channel
//   - Close the channel when enumeration is complete
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	bufferSize - the buffer size for the returned channel (0 for unbuffered)
//
// Returns:
//
//	A read-only channel containing all elements from the enumeration
//
// ⚠️ Resource management: This operation starts a goroutine that runs
// until the enumeration is complete. Always consume all elements or
// the goroutine may block indefinitely.
//
// ⚠️ Blocking behavior: The goroutine will block when trying to send
// to a full channel if there's no reader. Use appropriate buffer size.
//
// ⚠️ Warning: If the returned channel is not fully consumed, the goroutine
// may leak. Always range over the entire channel or ensure proper cleanup.
//
// Notes:
//   - If the enumerator is nil, returns a closed channel
//   - If the enumeration is empty, returns an empty but closed channel
//   - The goroutine automatically closes the channel when done
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(bufferSize) plus any upstream buffering
//   - The enumeration runs in a separate goroutine, enabling concurrent processing
func (q Enumerator[T]) ToChannel(bufferSize int) <-chan T {
	return toChannelInternal(q, bufferSize)
}

// ToChannel converts an enumeration to a channel that yields all elements.
// This operation is useful for converting lazy enumerations into channel-based
// processing pipelines or for interoperability with goroutines.
//
// The to channel operation will:
//   - Create a new channel with specified buffer size
//   - Start a goroutine that iterates through the enumeration
//   - Send each element to the channel
//   - Close the channel when enumeration is complete
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	bufferSize - the buffer size for the returned channel (0 for unbuffered)
//
// Returns:
//
//	A read-only channel containing all elements from the enumeration
//
// ⚠️ Resource management: This operation starts a goroutine that runs
// until the enumeration is complete. Always consume all elements or
// the goroutine may block indefinitely.
//
// ⚠️ Blocking behavior: The goroutine will block when trying to send
// to a full channel if there's no reader. Use appropriate buffer size.
//
// ⚠️ Warning: If the returned channel is not fully consumed, the goroutine
// may leak. Always range over the entire channel or ensure proper cleanup.
//
// Notes:
//   - If the enumerator is nil, returns a closed channel
//   - If the enumeration is empty, returns an empty but closed channel
//   - The goroutine automatically closes the channel when done
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(bufferSize) plus any upstream buffering
//   - The enumeration runs in a separate goroutine, enabling concurrent processing
func (q EnumeratorAny[T]) ToChannel(bufferSize int) <-chan T {
	return toChannelInternal(q, bufferSize)
}

// ToChannel converts a sorted enumeration to a channel that yields all elements in sorted order.
// This operation is useful for converting lazy sorted enumerations into channel-based
// processing pipelines or for interoperability with goroutines.
//
// The to channel operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Create a new channel with specified buffer size
//   - Start a goroutine that iterates through the sorted enumeration
//   - Send each element to the channel in sorted order
//   - Close the channel when enumeration is complete
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	bufferSize - the buffer size for the returned channel (0 for unbuffered)
//
// Returns:
//
//	A read-only channel containing all elements from the enumeration in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The channel creation and goroutine startup is O(1).
// For large enumerations, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// Plus O(bufferSize) for the channel buffer.
//
// ⚠️ Resource management: This operation starts a goroutine that runs
// until the enumeration is complete. Always consume all elements or
// the goroutine may block indefinitely.
//
// ⚠️ Blocking behavior: The goroutine will block when trying to send
// to a full channel if there's no reader. Use appropriate buffer size.
//
// ⚠️ Warning: If the returned channel is not fully consumed, the goroutine
// may leak. Always range over the entire channel or ensure proper cleanup.
//
// Notes:
//   - If the enumerator is nil, returns a closed channel
//   - If the enumeration is empty, returns an empty but closed channel
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - The goroutine automatically closes the channel when done
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n + bufferSize) for sorting plus channel buffering
//   - The enumeration runs in a separate goroutine, enabling concurrent processing
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) ToChannel(bufferSize int) <-chan T {
	return toChannelInternal(o.getSortedEnumerator(), bufferSize)
}

// ToChannel converts a sorted enumeration to a channel that yields all elements in sorted order.
// This operation is useful for converting lazy sorted enumerations into channel-based
// processing pipelines or for interoperability with goroutines.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The to channel operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Create a new channel with specified buffer size
//   - Start a goroutine that iterates through the sorted enumeration
//   - Send each element to the channel in sorted order
//   - Close the channel when enumeration is complete
//   - Handle nil enumerators gracefully
//
// Parameters:
//
//	bufferSize - the buffer size for the returned channel (0 for unbuffered)
//
// Returns:
//
//	A read-only channel containing all elements from the enumeration in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The channel creation and goroutine startup is O(1).
// For large enumerations, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// Plus O(bufferSize) for the channel buffer.
//
// ⚠️ Resource management: This operation starts a goroutine that runs
// until the enumeration is complete. Always consume all elements or
// the goroutine may block indefinitely.
//
// ⚠️ Blocking behavior: The goroutine will block when trying to send
// to a full channel if there's no reader. Use appropriate buffer size.
//
// ⚠️ Warning: If the returned channel is not fully consumed, the goroutine
// may leak. Always range over the entire channel or ensure proper cleanup.
//
// Notes:
//   - If the enumerator is nil, returns a closed channel
//   - If the enumeration is empty, returns an empty but closed channel
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - The goroutine automatically closes the channel when done
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n + bufferSize) for sorting plus channel buffering
//   - The enumeration runs in a separate goroutine, enabling concurrent processing
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) ToChannel(bufferSize int) <-chan T {
	return toChannelInternal(o.getSortedEnumerator(), bufferSize)
}

func toChannelInternal[T any](enumerator func(func(T) bool), bufferSize int) <-chan T {
	if enumerator == nil {
		ch := make(chan T)
		close(ch)
		return ch
	}

	ch := make(chan T, bufferSize)

	go func() {
		defer close(ch)
		enumerator(func(item T) bool {
			ch <- item
			return true
		})
	}()

	return ch
}
