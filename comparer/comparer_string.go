package comparer

// ComparerString is a predefined ComparerFunc for comparing two string values.
// It performs a lexicographic comparison between two strings and returns:
//
//	-1 if the first string is lexicographically less than the second
//	 0 if both strings are equal
//	+1 if the first string is lexicographically greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerString ComparerFunc[string] = func(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// KeyComparerString creates a ComparerFunc[T] that compares elements of type T based on a string key extracted by a key selector function.
// It uses the predefined ComparerString to perform the comparison on the extracted string values.
// This is a convenience function for sorting or ordering operations where the sort key is a 'string'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts a string key (string) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their string keys using keySelector
//	and then applying ComparerString to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerString[T any](keySelector func(T) string) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerString)
}
