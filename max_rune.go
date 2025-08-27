package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxRune returns the largest rune value extracted from elements of the enumeration
// using a key selector function and natural Unicode code point ordering.
// This operation is useful for finding the maximum Unicode character key derived from complex elements.
//
// The MaxRune operation will:
//   - Apply the keySelector function to each element to extract a rune value
//   - Compare extracted keys using natural Unicode code point ordering
//   - Return the largest rune key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when maximum Unicode code point is found
//
// Parameters:
//
//	keySelector - a function that extracts a rune key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum rune key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the maximum element.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when maximum Unicode code point is found
//   - For large enumerations without maximum values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxRune(keySelector func(T) rune) (rune, bool) {
	return extremumRuneInternal(e, keySelector, reverseRuneComparer)
}

// MaxRune returns the largest rune value extracted from elements of the enumeration
// using a key selector function and natural Unicode code point ordering.
// This operation is useful for finding the maximum Unicode character key derived from complex elements.
//
// The MaxRune operation will:
//   - Apply the keySelector function to each element to extract a rune value
//   - Compare extracted keys using natural Unicode code point ordering
//   - Return the largest rune key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when maximum Unicode code point is found
//
// Parameters:
//
//	keySelector - a function that extracts a rune key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum rune key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to find the maximum element.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it may
// trigger upstream operations during enumeration.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when maximum Unicode code point is found
//   - For large enumerations without maximum values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxRune(keySelector func(T) rune) (rune, bool) {
	return extremumRuneInternal(e, keySelector, reverseRuneComparer)
}

var reverseRuneComparer = comparer.ComparerFunc[rune](func(a, b rune) int {
	return comparer.ComparerRune(b, a)
})
