package enumerable

// SkipWhile bypasses elements in an enumeration as long as a specified condition is true
// and then yields the remaining elements.
// This operation is useful for skipping elements until a certain condition is met.
//
// The skip while operation will:
//   - Bypass elements from the beginning while the predicate returns true
//   - Once the predicate returns false for an element, yield that element and all subsequent elements
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to skip an element
//
// Returns:
//
//	An Enumerator[T] that yields elements after skipping initial elements that match the condition
//
// Notes:
//   - If the predicate never returns false, returns an empty enumerator
//   - If the predicate immediately returns false for the first element, returns the original enumerator
//   - If the original enumerator is nil, returns an empty enumerator
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - No elements are buffered - memory efficient
//   - The enumeration stops skipping as soon as the first non-matching element is found
//   - Once skipping stops, all remaining elements are yielded (even if they would match the predicate)
func (q Enumerator[T]) SkipWhile(predicate func(T) bool) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return skipWhileInternal(q, predicate)
}

// SkipWhile bypasses elements in an enumeration as long as a specified condition is true
// and then yields the remaining elements.
// This operation is useful for skipping elements until a certain condition is met.
//
// The skip while operation will:
//   - Bypass elements from the beginning while the predicate returns true
//   - Once the predicate returns false for an element, yield that element and all subsequent elements
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to skip an element
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements after skipping initial elements that match the condition
//
// Notes:
//   - If the predicate never returns false, returns an empty enumerator
//   - If the predicate immediately returns false for the first element, returns the original enumerator
//   - If the original enumerator is nil, returns an empty enumerator
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - No elements are buffered - memory efficient
//   - The enumeration stops skipping as soon as the first non-matching element is found
//   - Once skipping stops, all remaining elements are yielded (even if they would match the predicate)
func (q EnumeratorAny[T]) SkipWhile(predicate func(T) bool) EnumeratorAny[T] {
	if q == nil {
		return EmptyAny[T]()
	}
	return skipWhileInternal(q, predicate)
}

// SkipWhile bypasses elements in a sorted enumeration as long as a specified condition is true
// and then yields the remaining elements in sorted order.
// This operation is useful for skipping elements until a certain condition is met
// in the context of sorted sequences.
//
// The skip while operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Bypass elements from the beginning while the predicate returns true
//   - Once the predicate returns false for an element, yield that element and all subsequent elements
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to skip an element
//
// Returns:
//
//	An Enumerator[T] that yields elements in sorted order after skipping initial elements that match the condition
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The predicate is evaluated during iteration after sorting is complete.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If the predicate never returns false, returns an empty enumerator
//   - If the predicate immediately returns false for the first element, returns all elements in sorted order
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - The enumeration stops skipping as soon as the first non-matching element is found
//   - Once skipping stops, all remaining elements are yielded (even if they would match the predicate)
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) SkipWhile(predicate func(T) bool) Enumerator[T] {
	return skipWhileInternal(o.getSortedEnumerator(), predicate)
}

// SkipWhile bypasses elements in a sorted enumeration as long as a specified condition is true
// and then yields the remaining elements in sorted order.
// This operation is useful for skipping elements until a certain condition is met
// in the context of sorted sequences.
// This method supports any type T, including non-comparable types.
//
// The skip while operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Bypass elements from the beginning while the predicate returns true
//   - Once the predicate returns false for an element, yield that element and all subsequent elements
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to skip an element
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements in sorted order after skipping initial elements that match the condition
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The predicate is evaluated during iteration after sorting is complete.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If the predicate never returns false, returns an empty enumerator
//   - If the predicate immediately returns false for the first element, returns all elements in sorted order
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - The enumeration stops skipping as soon as the first non-matching element is found
//   - Once skipping stops, all remaining elements are yielded (even if they would match the predicate)
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) SkipWhile(predicate func(T) bool) EnumeratorAny[T] {
	return skipWhileInternal(o.getSortedEnumerator(), predicate)
}

func skipWhileInternal[T any](enumerator func(func(T) bool), predicate func(T) bool) func(func(T) bool) {
	return func(yield func(T) bool) {
		skipping := true
		enumerator(func(item T) bool {
			if skipping {
				if predicate(item) {
					return true
				}
				skipping = false
			}
			return yield(item)
		})
	}
}
