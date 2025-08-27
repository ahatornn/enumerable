package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MinRune returns the smallest rune value extracted from elements of the enumeration
// using a key selector function and natural Unicode code point ordering.
// This operation is useful for finding the minimum Unicode character key derived from complex elements.
//
// The MinRune operation will:
//   - Apply the keySelector function to each element to extract a rune value
//   - Compare extracted keys using natural Unicode code point ordering
//   - Return the smallest rune key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when 0 is found (since 0 is the minimum Unicode code point)
//
// Parameters:
//
//	keySelector - a function that extracts a rune key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum rune key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that iterates
// through the enumeration to find the minimum element. Performance is
// optimized with early termination when 0 is found (since 0 is the minimum Unicode code point), but
// worst-case scenario processes all elements.
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
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when 0 is found (since 0 is the minimum Unicode code point)
//   - For large enumerations without 0 values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinRune(keySelector func(T) rune) (rune, bool) {
	return extremumRuneInternal(e, keySelector, comparer.ComparerRune)
}

// MinRune returns the smallest rune value extracted from elements of the enumeration
// using a key selector function and natural Unicode code point ordering.
// This operation is useful for finding the minimum Unicode character key derived from complex elements.
//
// The MinRune operation will:
//   - Apply the keySelector function to each element to extract a rune value
//   - Compare extracted keys using natural Unicode code point ordering
//   - Return the smallest rune key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when 0 is found (since 0 is the minimum Unicode code point)
//
// Parameters:
//
//	keySelector - a function that extracts a rune key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum rune key value extracted from elements and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that iterates
// through the enumeration to find the minimum element. Performance is
// optimized with early termination when 0 is found (since 0 is the minimum Unicode code point), but
// worst-case scenario processes all elements.
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
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when 0 is found (since 0 is the minimum Unicode code point)
//   - For large enumerations without 0 values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinRune(keySelector func(T) rune) (rune, bool) {
	return extremumRuneInternal(e, keySelector, comparer.ComparerRune)
}

func extremumRuneInternal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) rune,
	cmp comparer.ComparerFunc[rune],
) (rune, bool) {
	var resultKey rune
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	var extremeValue rune
	if cmp(0, 1) < 0 {
		extremeValue = rune(0)
	} else {
		extremeValue = rune(0x10FFFF)
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
