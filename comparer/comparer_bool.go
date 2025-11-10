package comparer

// ComparerBool is a predefined ComparerFunc for comparing two bool values.
// It performs a logical comparison where false is considered less than true and returns:
//
//	-1 if the first boolean is false and the second is true
//	 0 if both booleans are equal
//	+1 if the first boolean is true and the second is false
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerBool ComparerFunc[bool] = func(a, b bool) int {
	if !a && b {
		return -1
	} else if a && !b {
		return 1
	}
	return 0
}

// KeyComparerBool creates a ComparerFunc[T] that compares elements of type T based on a boolean key extracted by a key selector function.
// It uses the predefined ComparerBool to perform the comparison on the extracted boolean values.
// This is a convenience function for sorting or ordering operations where the sort key is a 'bool'.
// In the standard boolean ordering, false is considered less than true.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts a boolean key (bool) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their boolean keys using keySelector
//	and then applying ComparerBool to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerBool[T any](keySelector func(T) bool) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerBool)
}
