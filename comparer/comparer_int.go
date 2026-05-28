package comparer

// ComparerInt is a predefined ComparerFunc for comparing two int values.
// It performs a natural numeric comparison between two integers and returns:
//
//	-1 if the first integer is less than the second
//	 0 if both integers are equal
//	+1 if the first integer is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerInt ComparerFunc[int] = func(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// KeyComparerInt creates a ComparerFunc[T] that compares elements of type T based on an integer key extracted by a key selector function.
// It uses the predefined ComparerInt to perform the comparison on the extracted integer values.
// This is a convenience function for sorting or ordering operations where the sort key is an 'int'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts an integer key (int) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their integer keys using keySelector
//	and then applying ComparerInt to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerInt[T any](keySelector func(T) int) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerInt)
}
