package enumerable

// Where filters an enumeration based on a predicate function.
// This operation returns elements that satisfy the specified condition.
//
// The where operation will:
// - Apply the predicate function to each element in the enumeration
// - Yield only elements for which the predicate returns true
// - Preserve the original order of elements that pass the filter
// - Handle nil enumerators gracefully
// - Support early termination when consumer returns false
//
// Parameters:
//   predicate - a function that determines whether to include an element
//
// Returns:
//   An Enumerator[T] that yields elements satisfying the predicate
//
// Notes:
// - If the original enumerator is nil, returns nil
// - Lazy evaluation - elements are processed and filtered during iteration
// - No elements are buffered - memory efficient
// - The enumeration stops when the source is exhausted or consumer returns false
// - Predicate function should be pure (no side effects) for predictable behavior
// - Elements for which predicate returns false are simply skipped
func (q Enumerator[T]) Where(predicate func(T) bool) Enumerator[T] {
	if q == nil {
		return nil
	}
	return func(yield func(T) bool) {
		q(func(item T) bool {
			if predicate(item) {
				return yield(item)
			}
			return true
		})
	}
}
