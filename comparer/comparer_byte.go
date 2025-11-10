package comparer

// ComparerByte is a predefined ComparerFunc for comparing two byte values.
// It performs a numerical comparison between two unsigned 8-bit integers and returns:
//
//	-1 if the first byte is less than the second
//	 0 if both bytes are equal
//	+1 if the first byte is greater than the second
//
// This comparer adheres to the required mathematical properties of consistency,
// antisymmetry, transitivity, and equality as defined by the ComparerFunc type.
var ComparerByte ComparerFunc[byte] = func(a, b byte) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// KeyComparerByte creates a ComparerFunc[T] that compares elements of type T based on a byte key extracted by a key selector function.
// It uses the predefined ComparerByte to perform the comparison on the extracted byte values.
// This is a convenience function for sorting or ordering operations where the sort key is a 'byte'.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//
// Parameters:
//
//	keySelector - a function that extracts a byte key (byte) from an element of type T
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by extracting their byte keys using keySelector
//	and then applying ComparerByte to those keys.
//
// Notes:
//   - This function is a specialized wrapper around the generic KeyComparer function.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementation of the keySelector function.
func KeyComparerByte[T any](keySelector func(T) byte) ComparerFunc[T] {
	return KeyComparer(keySelector, ComparerByte)
}
