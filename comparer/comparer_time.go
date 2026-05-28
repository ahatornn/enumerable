package comparer

import "time"

// ComparerTime is a predefined ComparerFunc for comparing two time.Time values.
// It performs a chronological comparison between two time values and returns:
//
//	-1 if the first time is before the second
//	 0 if both times are equal
//	+1 if the first time is after the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
//
// Note: This comparison uses time.Time's built-in Before() and After() methods
// for accurate chronological ordering.
var ComparerTime ComparerFunc[time.Time] = func(a, b time.Time) int {
	if a.Before(b) {
		return -1
	} else if a.After(b) {
		return 1
	}
	return 0
}

// KeyComparerTime creates a ComparerFunc[T] that compares elements of type T based on a time.Time key extracted by a key selector function.
// It uses the predefined ComparerTime to perform the comparison on the extracted time.Time values.
// This is a convenience function for sorting or ordering operations where the sort key is a 'time.Time'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts a time.Time key (time.Time) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their time.Time keys using keySelector
//	and then applying ComparerTime to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerTime[T any](keySelector func(T) time.Time) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerTime)
}
