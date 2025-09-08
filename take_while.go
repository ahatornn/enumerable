package enumerable

// TakeWhile returns elements from an enumeration as long as a specified condition is true.
// This operation is useful for taking elements until a certain condition is met,
// such as taking elements while they are valid or within a range.
//
// The take while operation will:
//   - Yield elements from the start while the predicate returns true
//   - Stop enumeration as soon as the predicate returns false for an element
//   - Handle edge cases gracefully (nil enumerator, always false predicate)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to take an element
//
// Returns:
//
//	An Enumerator[T] that yields elements while the condition is true
//
// Notes:
//   - If the predicate immediately returns false for the first element, returns empty enumerator
//   - If the predicate never returns false, returns all elements from the enumeration
//   - If the original enumerator is nil, returns an empty enumerator
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - No elements are buffered - memory efficient
//   - The enumeration stops as soon as the predicate returns false or consumer returns false
//   - Once the predicate returns false for any element, no subsequent elements are evaluated
func (q Enumerator[T]) TakeWhile(predicate func(T) bool) Enumerator[T] {
	if q == nil {
		return Empty[T]()
	}
	return takeWhileInternal(q, predicate)
}

// TakeWhile returns elements from an enumeration as long as a specified condition is true.
// This operation is useful for taking elements until a certain condition is met,
// such as taking elements while they are valid or within a range.
//
// The take while operation will:
//   - Yield elements from the start while the predicate returns true
//   - Stop enumeration as soon as the predicate returns false for an element
//   - Handle edge cases gracefully (nil enumerator, always false predicate)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to take an element
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements while the condition is true
//
// Notes:
//   - If the predicate immediately returns false for the first element, returns empty enumerator
//   - If the predicate never returns false, returns all elements from the enumeration
//   - If the original enumerator is nil, returns an empty enumerator
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - No elements are buffered - memory efficient
//   - The enumeration stops as soon as the predicate returns false or consumer returns false
//   - Once the predicate returns false for any element, no subsequent elements are evaluated
func (q EnumeratorAny[T]) TakeWhile(predicate func(T) bool) EnumeratorAny[T] {
	if q == nil {
		return EmptyAny[T]()
	}
	return takeWhileInternal(q, predicate)
}

// TakeWhile returns elements from a sorted enumeration as long as a specified condition is true.
// This operation is useful for taking elements until a certain condition is met,
// such as taking elements while they are valid or within a range, in the context of sorted sequences.
//
// The take while operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Yield elements from the start while the predicate returns true
//   - Stop enumeration as soon as the predicate returns false for an element
//   - Handle edge cases gracefully (nil enumerator, always false predicate)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to take an element
//
// Returns:
//
//	An Enumerator[T] that yields elements in sorted order while the condition is true
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The predicate evaluation is O(1) per element during iteration.
// For large enumerations, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If the predicate immediately returns false for the first element, returns empty enumerator
//   - If the predicate never returns false, returns all elements from the sorted enumeration
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - The enumeration stops as soon as the predicate returns false or consumer returns false
//   - Once the predicate returns false for any element, no subsequent elements are evaluated
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) TakeWhile(predicate func(T) bool) Enumerator[T] {
	return takeWhileInternal(o.getSortedEnumerator(), predicate)
}

// TakeWhile returns elements from a sorted enumeration as long as a specified condition is true.
// This operation is useful for taking elements until a certain condition is met,
// such as taking elements while they are valid or within a range, in the context of sorted sequences.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The take while operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Yield elements from the start while the predicate returns true
//   - Stop enumeration as soon as the predicate returns false for an element
//   - Handle edge cases gracefully (nil enumerator, always false predicate)
//   - Support early termination when consumer returns false
//
// Parameters:
//
//	predicate - a function that determines whether to take an element
//
// Returns:
//
//	An EnumeratorAny[T] that yields elements in sorted order while the condition is true
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// Time complexity: O(n log n) for sorting where n is the total number of elements.
// The predicate evaluation is O(1) per element during iteration.
// For large enumerations, this may process more elements than necessary.
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// Notes:
//   - If the predicate immediately returns false for the first element, returns empty enumerator
//   - If the predicate never returns false, returns all elements from the sorted enumeration
//   - If the original enumerator is nil, returns an empty enumerator
//   - Actual sorting computation occurs only during first enumeration
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - Lazy evaluation - elements are processed and evaluated during iteration
//   - No elements are buffered beyond sorting phase - memory efficient after sorting
//   - The enumeration stops as soon as the predicate returns false or consumer returns false
//   - Once the predicate returns false for any element, no subsequent elements are evaluated
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent enumerations will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
func (o OrderEnumeratorAny[T]) TakeWhile(predicate func(T) bool) EnumeratorAny[T] {
	return takeWhileInternal(o.getSortedEnumerator(), predicate)
}

func takeWhileInternal[T any](enumerator func(func(T) bool), predicate func(T) bool) func(func(T) bool) {
	return func(yield func(T) bool) {
		enumerator(func(item T) bool {
			if !predicate(item) {
				return false
			}
			return yield(item)
		})
	}
}
