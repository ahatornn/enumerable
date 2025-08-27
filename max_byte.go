package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxByte returns the largest byte value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric byte key derived from complex elements.
//
// The MaxByte operation will:
//   - Apply the keySelector function to each element to extract a byte value
//   - Compare extracted keys using natural numeric ordering (0 to 255)
//   - Return the largest byte key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when 255 is found (since 255 is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a byte key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum byte key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that iterates
// through the enumeration to find the maximum element.
// Performance is optimized with early termination when 255 is
// found, but worst-case scenario processes all elements.
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
//   - Performance is optimized: terminates early when 255 is found (since 255 is the maximum byte value)
//   - For large enumerations without 255 values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxByte(keySelector func(T) byte) (byte, bool) {
	return extremumByteInternal(e, keySelector, reverseByteComparer)
}

// MaxByte returns the largest byte value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric byte key derived from complex elements.
//
// The MaxByte operation will:
//   - Apply the keySelector function to each element to extract a byte value
//   - Compare extracted keys using natural numeric ordering (0 to 255)
//   - Return the largest byte key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when 255 is found (since 255 is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a byte key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum byte key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that iterates
// through the enumeration to find the maximum element.
// Performance is optimized with early termination when 255 is
// found, but worst-case scenario processes all elements.
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
//   - Performance is optimized: terminates early when 255 is found (since 255 is the maximum byte value)
//   - For large enumerations without 255 values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxByte(keySelector func(T) byte) (byte, bool) {
	return extremumByteInternal(e, keySelector, reverseByteComparer)
}

var reverseByteComparer = comparer.ComparerFunc[byte](func(a, b byte) int {
	return comparer.ComparerByte(b, a)
})
