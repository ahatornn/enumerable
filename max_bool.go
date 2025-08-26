package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MaxBool returns the largest boolean value extracted from elements of the enumeration
// using a key selector function and natural boolean ordering (false < true).
// This operation is useful for finding the maximum boolean key derived from complex elements.
//
// The MaxBool operation will:
//   - Apply the keySelector function to each element to extract a bool value
//   - Compare extracted keys using natural boolean ordering (true > false)
//   - Return the largest bool key and true if the enumeration is non-empty
//   - Return false and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when true is found (since true is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a bool key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (false, false).
//
// Returns:
//
//	The maximum bool key value extracted from elements and true if found,
//	false and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (false, false)
//   - If keySelector is nil, returns (false, false)
//   - If the enumeration is empty, returns (false, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when true is found (since true > false)
//   - For large enumerations without true values, all elements may be processed
//   - To use custom ordering, use MaxBoolBy with a custom comparer
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e Enumerator[T]) MaxBool(keySelector func(T) bool) (bool, bool) {
	return extremumBoolInternal(e, keySelector, reverseBoolComparer)
}

// MaxBool returns the largest boolean value extracted from elements of the enumeration
// using a key selector function and natural boolean ordering (false < true).
// This operation is useful for finding the maximum boolean key derived from complex elements.
//
// The MaxBool operation will:
//   - Apply the keySelector function to each element to extract a bool value
//   - Compare extracted keys using natural boolean ordering (true > false)
//   - Return the largest bool key and true if the enumeration is non-empty
//   - Return false and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when true is found (since true is the maximum value)
//
// Parameters:
//
//	keySelector - a function that extracts a bool key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (false, false).
//
// Returns:
//
//	The maximum bool key value extracted from elements and true if found,
//	false and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (false, false)
//   - If keySelector is nil, returns (false, false)
//   - If the enumeration is empty, returns (false, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same maximal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when true is found (since true > false)
//   - For large enumerations without true values, all elements may be processed
//   - To use custom ordering, use MaxBoolBy with a custom comparer
//   - To get the original element associated with the maximum key, use MaxBy instead
func (e EnumeratorAny[T]) MaxBool(keySelector func(T) bool) (bool, bool) {
	return extremumBoolInternal(e, keySelector, reverseBoolComparer)
}

var reverseBoolComparer = comparer.ComparerFunc[bool](func(a, b bool) int {
	return comparer.ComparerBool(b, a)
})
