package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MinByte returns the smallest byte value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric byte key derived from complex elements.
//
// The MinByte operation will:
//   - Apply the keySelector function to each element to extract a byte value
//   - Compare extracted keys using natural numeric ordering (0 to 255)
//   - Return the smallest byte key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when 0 is found (since 0 is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a byte key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum byte key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when 0 is found (since 0 is the minimum byte value)
//   - For large enumerations without 0 values, all elements may be processed
//   - To use custom ordering, use MinByteBy with a custom comparer
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinByte(keySelector func(T) byte) (byte, bool) {
	return extremumByteInternal(e, keySelector, comparer.ComparerByte)
}

// MinByte returns the smallest byte value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric byte key derived from complex elements.
//
// The MinByte operation will:
//   - Apply the keySelector function to each element to extract a byte value
//   - Compare extracted keys using natural numeric ordering (0 to 255)
//   - Return the smallest byte key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when 0 is found (since 0 is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a byte key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum byte key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when 0 is found (since 0 is the minimum byte value)
//   - For large enumerations without 0 values, all elements may be processed
//   - To use custom ordering, use MinByteBy with a custom comparer
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinByte(keySelector func(T) byte) (byte, bool) {
	return extremumByteInternal(e, keySelector, comparer.ComparerByte)
}

func extremumByteInternal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) byte,
	cmp comparer.ComparerFunc[byte],
) (byte, bool) {
	var resultKey byte
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	var extremeValue byte
	if cmp(0, 255) < 0 {
		extremeValue = 0
	} else {
		extremeValue = 255
	}

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !found {
			resultKey = key
			found = true
		} else if cmp(key, resultKey) < 0 {
			resultKey = key
		}

		if key == extremeValue {
			return false
		}

		return true
	})

	return resultKey, found
}
