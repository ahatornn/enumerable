package enumerable

// Concat combines two enumerations into a single enumeration.
// The resulting enumeration yields all elements from the first enumeration,
// followed by all elements from the second enumeration.
//
// The concatenation will:
// - Yield all elements from the first enumeration in order
// - Then yield all elements from the second enumeration in order
// - Handle nil enumerators gracefully (treated as empty)
// - Support early termination when consumer returns false
//
// Parameters:
//   second - the enumerator to concatenate after the current one
//
// Returns:
//   A new Enumerator[T] that yields elements from both enumerations in sequence
//
// Notes:
// - Nil enumerators are treated as empty (no elements yielded)
// - Both enumerations are consumed in order during iteration
// - If the first enumeration is infinite, second will never be reached
// - Lazy evaluation - elements are produced on-demand during iteration
// - No elements are buffered - memory efficient
// - Safe for use with any combination of nil and non-nil enumerators
func (q Enumerator[T]) Concat(second Enumerator[T]) Enumerator[T] {
	return func(yield func(T) bool) {
		if q != nil {
			q(func(item T) bool {
				return yield(item)
			})
		}

		if second != nil {
			second(func(item T) bool {
				return yield(item)
			})
		}
	}
}
