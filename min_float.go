package enumerable

import (
	"math"

	"github.com/ahatornn/enumerable/comparer"
)

// MinFloat returns the smallest float32 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric float32 key derived from complex elements.
//
// The MinFloat operation will:
//   - Apply the keySelector function to each element to extract a float32 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the smallest float32 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when -Inf is found (since -Inf is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float32 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum float32 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when -Inf is found (since -Inf is the minimum float32 value)
//   - For large enumerations without -Inf values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinFloat(keySelector func(T) float32) (float32, bool) {
	return extremumFloat32Internal(e, keySelector, comparer.ComparerFloat32)
}

// MinFloat returns the smallest float32 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric float32 key derived from complex elements.
//
// The MinFloat operation will:
//   - Apply the keySelector function to each element to extract a float32 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the smallest float32 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when -Inf is found (since -Inf is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float32 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum float32 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when -Inf is found (since -Inf is the minimum float32 value)
//   - For large enumerations without -Inf values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinFloat(keySelector func(T) float32) (float32, bool) {
	return extremumFloat32Internal(e, keySelector, comparer.ComparerFloat32)
}

// MinFloat64 returns the smallest float64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric float64 key derived from complex elements.
//
// The MinFloat64 operation will:
//   - Apply the keySelector function to each element to extract a float64 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the smallest float64 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when -Inf is found (since -Inf is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum float64 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when -Inf is found (since -Inf is the minimum float64 value)
//   - For large enumerations without -Inf values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e Enumerator[T]) MinFloat64(keySelector func(T) float64) (float64, bool) {
	return extremumFloat64Internal(e, keySelector, comparer.ComparerFloat64)
}

// MinFloat64 returns the smallest float64 value extracted from elements of the enumeration
// using a key selector function and natural numeric ordering.
// This operation is useful for finding the minimum numeric float64 key derived from complex elements.
//
// The MinFloat64 operation will:
//   - Apply the keySelector function to each element to extract a float64 value
//   - Compare extracted keys using natural numeric ordering
//   - Return the smallest float64 key and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//   - Optimize performance by early termination when -Inf is found (since -Inf is the minimum value)
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The minimum float64 key value extracted from elements and true if found,
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
//   - Performance is optimized: terminates early when -Inf is found (since -Inf is the minimum float64 value)
//   - For large enumerations without -Inf values, all elements may be processed
//   - To get the original element associated with the minimum key, use MinBy instead
func (e EnumeratorAny[T]) MinFloat64(keySelector func(T) float64) (float64, bool) {
	return extremumFloat64Internal(e, keySelector, comparer.ComparerFloat64)
}

func extremumFloat32Internal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) float32,
	cmp comparer.ComparerFunc[float32],
) (float32, bool) {
	var resultKey float32
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	var extremeValue float32
	if cmp(0.0, 1.0) < 0 {
		extremeValue = float32(math.Inf(-1))
	} else {
		extremeValue = float32(math.Inf(1))
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

func extremumFloat64Internal[T any](
	enumerator func(yield func(T) bool),
	keySelector func(T) float64,
	cmp comparer.ComparerFunc[float64],
) (float64, bool) {
	var resultKey float64
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	var extremeValue float64
	if cmp(0.0, 1.0) < 0 {
		extremeValue = math.Inf(-1)
	} else {
		extremeValue = math.Inf(1)
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
