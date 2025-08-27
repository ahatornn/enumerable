package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxString returns the lexicographically largest string value extracted from elements of the enumeration
// using a key selector function and natural string ordering.
// This operation is useful for finding the maximum string key derived from complex elements.
//
// The MaxString operation will:
//   - Apply the keySelector function to each element to extract a string value
//   - Compare extracted keys using natural lexicographic ordering
//   - Return the largest string key and true if the enumeration is non-empty
//   - Return empty string ("") and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a string key from each element of type T.
//	              Must be non-nil; if nil, the operation returns ("", false).
//
// Returns:
//
//	The maximum string key value extracted from elements and true if found,
//	empty string ("") and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the maximum element.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns ("", false)
//   - If keySelector is nil, returns ("", false)
//   - If the enumeration is empty, returns ("", false)
//   - Processes elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, all elements may be processed (no early termination optimization)
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxString(keySelector func(T) string) (string, bool) {
	return extremumStringInternal(e, keySelector, reverseStringComparer)
}

// MaxString returns the lexicographically largest string value extracted from elements of the enumeration
// using a key selector function and natural string ordering.
// This operation is useful for finding the maximum string key derived from complex elements.
//
// The MaxString operation will:
//   - Apply the keySelector function to each element to extract a string value
//   - Compare extracted keys using natural lexicographic ordering
//   - Return the largest string key and true if the enumeration is non-empty
//   - Return empty string ("") and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a string key from each element of type T.
//	              Must be non-nil; if nil, the operation returns ("", false).
//
// Returns:
//
//	The maximum string key value extracted from elements and true if found,
//	empty string ("") and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the maximum element.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns ("", false)
//   - If keySelector is nil, returns ("", false)
//   - If the enumeration is empty, returns ("", false)
//   - Processes elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For large enumerations, all elements may be processed (no early termination optimization)
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxString(keySelector func(T) string) (string, bool) {
	return extremumStringInternal(e, keySelector, reverseStringComparer)
}

var reverseStringComparer = comparer.ComparerFunc[string](func(a, b string) int {
	return comparer.ComparerString(b, a)
})
