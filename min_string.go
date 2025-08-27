package enumerable

import "github.com/ahatornn/enumerable/comparer"

// MinString returns the lexicographically smallest string value extracted from elements of the enumeration
// using a key selector function and natural string ordering.
// This operation is useful for finding the minimum string key derived from complex elements.
//
// The MinString operation will:
//   - Apply the keySelector function to each element to extract a string value
//   - Compare extracted keys using natural lexicographic ordering
//   - Return the smallest string key and true if the enumeration is non-empty
//   - Return empty string ("") and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when empty string is found
//
// Parameters:
//
//	keySelector - a function that extracts a string key from each element of type T.
//	              Must be non-nil; if nil, the operation returns ("", false).
//
// Returns:
//
//	The minimum string key value extracted from elements and true if found,
//	empty string ("") and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns ("", false)
//   - If keySelector is nil, returns ("", false)
//   - If the enumeration is empty, returns ("", false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when empty string is found (since "" < any non-empty string)
//   - For large enumerations without empty strings, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinString(keySelector func(T) string) (string, bool) {
	return extremumStringInternal(e, keySelector, comparer.ComparerString)
}

// MinString returns the lexicographically smallest string value extracted from elements of the enumeration
// using a key selector function and natural string ordering.
// This operation is useful for finding the minimum string key derived from complex elements.
//
// The MinString operation will:
//   - Apply the keySelector function to each element to extract a string value
//   - Compare extracted keys using natural lexicographic ordering
//   - Return the smallest string key and true if the enumeration is non-empty
//   - Return empty string ("") and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when empty string is found
//
// Parameters:
//
//	keySelector - a function that extracts a string key from each element of type T.
//	              Must be non-nil; if nil, the operation returns ("", false).
//
// Returns:
//
//	The minimum string key value extracted from elements and true if found,
//	empty string ("") and false otherwise
//
// Notes:
//   - If the enumerator is nil, returns ("", false)
//   - If keySelector is nil, returns ("", false)
//   - If the enumeration is empty, returns ("", false)
//   - Processes elements in the enumeration - O(n) worst-case time complexity, but may terminate early
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element until termination
//   - If multiple elements yield the same minimal key, the first one encountered is used
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Performance is optimized: terminates early when empty string is found (since "" < any non-empty string)
//   - For large enumerations without empty strings, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinString(keySelector func(T) string) (string, bool) {
	return extremumStringInternal(e, keySelector, comparer.ComparerString)
}

func extremumStringInternal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) string,
	cmp comparer.ComparerFunc[string],
) (string, bool) {
	var resultKey string
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return "", false
	}

	isNaturalOrder := cmp("", "a") < 0

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !found {
			resultKey = key
			found = true
			if isNaturalOrder && key == "" {
				return false
			}
		} else if cmp(key, resultKey) < 0 {
			resultKey = key
			if isNaturalOrder && key == "" {
				return false
			}
		}

		return true
	})

	return resultKey, found
}
