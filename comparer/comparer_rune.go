package comparer

// ComparerRune is a predefined ComparerFunc for comparing two rune values.
// It performs a numerical comparison between two Unicode code points and returns:
//
//	-1 if the first rune is less than the second
//	 0 if both runes are equal
//	+1 if the first rune is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerRune ComparerFunc[rune] = func(a, b rune) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// KeyComparerRune creates a ComparerFunc[T] that compares elements of type T based on a rune key extracted by a key selector function.
// It uses the predefined ComparerRune to perform the comparison on the extracted rune values.
// This is a convenience function for sorting or ordering operations where the sort key is a 'rune'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts a rune key (rune) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their rune keys using keySelector
//	and then applying ComparerRune to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerRune[T any](keySelector func(T) rune) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerRune)
}
