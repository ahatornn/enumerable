package comparer

// ComparerFloat64 is a predefined ComparerFunc for comparing two float64 values.
// It performs a numerical comparison between two float64 values and returns:
//
//	-1 if the first value is less than the second
//	 0 if both values are equal
//	+1 if the first value is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
//
// Note: This comparison does not handle NaN values specially - NaN comparisons
// follow Go's built-in comparison rules where NaN is not equal to anything,
// including itself.
var ComparerFloat64 ComparerFunc[float64] = func(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// KeyComparerFloat64 creates a ComparerFunc[T] that compares elements of type T based on a float64 key extracted by a key selector function.
// It uses the predefined ComparerFloat64 to perform the comparison on the extracted float64 values.
// This is a convenience function for sorting or ordering operations where the sort key is a 'float64'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key (float64) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their float64 keys using keySelector
//	and then applying ComparerFloat64 to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerFloat64[T any](keySelector func(T) float64) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerFloat64)
}
