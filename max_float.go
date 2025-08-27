package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxFloat returns the largest float32 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric float32 key derived from complex elements.
//
// The MaxFloat operation will:
//   - Apply the keySelector function to each element to extract a float32 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the largest float32 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when +Inf is found (since +Inf is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float32 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum float32 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when +Inf is found (since +Inf is the maximum float32 value)
//   - For large enumerations without +Inf values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxFloat(keySelector func(T) float32) (float32, bool) {
	return extremumFloat32Internal(e, keySelector, reverseFloatComparer)
}

// MaxFloat returns the largest float32 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric float32 key derived from complex elements.
//
// The MaxFloat operation will:
//   - Apply the keySelector function to each element to extract a float32 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the largest float32 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when +Inf is found (since +Inf is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float32 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum float32 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when +Inf is found (since +Inf is the maximum float32 value)
//   - For large enumerations without +Inf values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxFloat(keySelector func(T) float32) (float32, bool) {
	return extremumFloat32Internal(e, keySelector, reverseFloatComparer)
}

// MaxFloat64 returns the largest float64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric float64 key derived from complex elements.
//
// The MaxFloat64 operation will:
//   - Apply the keySelector function to each element to extract a float64 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the largest float64 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when +Inf is found (since +Inf is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum float64 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when +Inf is found (since +Inf is the maximum float64 value)
//   - For large enumerations without +Inf values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxFloat64(keySelector func(T) float64) (float64, bool) {
	return extremumFloat64Internal(e, keySelector, reverseFloat64Comparer)
}

// MaxFloat64 returns the largest float64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the maximum numeric float64 key derived from complex elements.
//
// The MaxFloat64 operation will:
//   - Apply the keySelector function to each element to extract a float64 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the largest float64 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when +Inf is found (since +Inf is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The maximum float64 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when +Inf is found (since +Inf is the maximum float64 value)
//   - For large enumerations without +Inf values, all elements may be processed
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxFloat64(keySelector func(T) float64) (float64, bool) {
	return extremumFloat64Internal(e, keySelector, reverseFloat64Comparer)
}

var reverseFloatComparer = comparer.ComparerFunc[float32](func(a, b float32) int {
	return comparer.ComparerFloat32(b, a)
})

var reverseFloat64Comparer = comparer.ComparerFunc[float64](func(a, b float64) int {
	return comparer.ComparerFloat64(b, a)
})
