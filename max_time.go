package enumerable

import (
	"time"

	"github.com/ahatornn/enumerable/comparer"
)

// MaxTime returns the latest time.Time value extracted from elements of the enumeration
// using a key selector function and natural time ordering.
// This operation is useful for finding the maximum timestamp key derived from complex elements.
//
// The MaxTime operation will:
//   - Apply the keySelector function to each element to extract a time.Time value
//   - Compare extracted keys using natural time ordering
//   - Return the latest time.Time key and true if the enumeration is non-empty
//   - Return zero time (time.Time{}) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a time.Time key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (time.Time{}, false).
//
// Returns:
//
//	The maximum time.Time key value extracted from elements and true if found,
//	zero time (time.Time{}) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the maximum element.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns (time.Time{}, false)
//   - If keySelector is nil, returns (time.Time{}, false)
//   - If the enumeration is empty, returns (time.Time{}, false)
//   - Processes elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, all elements may be processed (no early termination optimization)
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxTime(keySelector func(T) time.Time) (time.Time, bool) {
	return extremumTimeInternal(e, keySelector, reverseTimeComparer)
}

// MaxTime returns the latest time.Time value extracted from elements of the enumeration
// using a key selector function and natural time ordering.
// This operation is useful for finding the maximum timestamp key derived from complex elements.
//
// The MaxTime operation will:
//   - Apply the keySelector function to each element to extract a time.Time value
//   - Compare extracted keys using natural time ordering
//   - Return the latest time.Time key and true if the enumeration is non-empty
//   - Return zero time (time.Time{}) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a time.Time key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (time.Time{}, false).
//
// Returns:
//
//	The maximum time.Time key value extracted from elements and true if found,
//	zero time (time.Time{}) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the maximum element.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns (time.Time{}, false)
//   - If keySelector is nil, returns (time.Time{}, false)
//   - If the enumeration is empty, returns (time.Time{}, false)
//   - Processes elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, all elements may be processed (no early termination optimization)
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxTime(keySelector func(T) time.Time) (time.Time, bool) {
	return extremumTimeInternal(e, keySelector, reverseTimeComparer)
}

var reverseTimeComparer = comparer.ComparerFunc[time.Time](func(a, b time.Time) int {
	return comparer.ComparerTime(b, a)
})
