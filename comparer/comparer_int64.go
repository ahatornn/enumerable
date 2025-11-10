package comparer

// ComparerInt64 is a predefined ComparerFunc for comparing two int64 values.
// It performs a natural numeric comparison between two int64 integers and returns:
//
//	-1 if the first integer is less than the second
//	 0 if both integers are equal
//	+1 if the first integer is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerInt64 ComparerFunc[int64] = func(a, b int64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// KeyComparerInt64 creates a ComparerFunc[T] that compares elements of type T based on an int64 key extracted by a key selector function.
// It uses the predefined ComparerInt64 to perform the comparison on the extracted int64 values.
// This is a convenience function for sorting or ordering operations where the sort key is an 'int64'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts an int64 key (int64) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their int64 keys using keySelector
//	and then applying ComparerInt64 to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerInt64[T any](keySelector func(T) int64) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerInt64)
}
