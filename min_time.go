package enumerable

import (
	"time"

	"github.com/ahatornn/enumerable/comparer"
)

// MinTime returns the earliest time.Time value extracted from elements of the enumeration
// using a key selector function and natural time ordering.
// This operation is useful for finding the minimum timestamp key derived from complex elements.
//
// The MinTime operation will:
//   - Apply the keySelector function to each element to extract a time.Time value
//   - Compare extracted keys using natural time ordering
//   - Return the earliest time.Time key and true if the enumeration is non-empty
//   - Return zero time (time.Time{}) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when zero time is found
//
// Parameters:
//
//	keySelector - a function that extracts a time.Time key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (time.Time{}, false).
//
// Returns:
//
//	The minimum time.Time key value extracted from elements and true if found,
//	zero time (time.Time{}) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that iterates
// through the enumeration to find the minimum element. Performance is
// optimized with early termination when zero time is found, but
// worst-case scenario processes all elements.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns (time.Time{}, false)
//   - If keySelector is nil, returns (time.Time{}, false)
//   - If the enumeration is empty, returns (time.Time{}, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when zero time is found (since time.Time{} is the earliest possible time)
//   - For large enumerations without zero times, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinTime(keySelector func(T) time.Time) (time.Time, bool) {
	return extremumTimeInternal(e, keySelector, comparer.ComparerTime)
}

// MinTime returns the earliest time.Time value extracted from elements of the enumeration
// using a key selector function and natural time ordering.
// This operation is useful for finding the minimum timestamp key derived from complex elements.
//
// The MinTime operation will:
//   - Apply the keySelector function to each element to extract a time.Time value
//   - Compare extracted keys using natural time ordering
//   - Return the earliest time.Time key and true if the enumeration is non-empty
//   - Return zero time (time.Time{}) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when zero time is found
//
// Parameters:
//
//	keySelector - a function that extracts a time.Time key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (time.Time{}, false).
//
// Returns:
//
//	The minimum time.Time key value extracted from elements and true if found,
//	zero time (time.Time{}) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that iterates
// through the enumeration to find the minimum element. Performance is
// optimized with early termination when zero time is found, but
// worst-case scenario processes all elements.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns (time.Time{}, false)
//   - If keySelector is nil, returns (time.Time{}, false)
//   - If the enumeration is empty, returns (time.Time{}, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when zero time is found (since time.Time{} is the earliest possible time)
//   - For large enumerations without zero times, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinTime(keySelector func(T) time.Time) (time.Time, bool) {
	return extremumTimeInternal(e, keySelector, comparer.ComparerTime)
}

func extremumTimeInternal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) time.Time,
	cmp comparer.ComparerFunc[time.Time],
) (time.Time, bool) {
	var resultKey time.Time
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return time.Time{}, false
	}

	isMinOperation := cmp(time.Time{}, time.Unix(1, 0)) < 0

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !found {
			resultKey = key
			found = true
		} else if cmp(key, resultKey) < 0 {
			resultKey = key
		}

		if isMinOperation && key.IsZero() {
			return false
		}

		return true
	})

	return resultKey, found
}
