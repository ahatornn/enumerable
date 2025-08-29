package enumerable

// AverageInt returns the arithmetic mean of integer values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageInt operation will:
//   - Apply the keySelector function to each element to extract an int value
//   - Calculate the sum of all extracted integer values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted integer values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Returns float64 to accommodate fractional results from integer division
//   - Uses int64 internally to prevent integer overflow during summation
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt64, AverageFloat, or AverageFloat64
func (e Enumerator[T]) AverageInt(keySelector func(T) int) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageInt returns the arithmetic mean of integer values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageInt operation will:
//   - Apply the keySelector function to each element to extract an int value
//   - Calculate the sum of all extracted integer values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted integer values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Returns float64 to accommodate fractional results from integer division
//   - Uses int64 internally to prevent integer overflow during summation
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt64, AverageFloat, or AverageFloat64
func (e EnumeratorAny[T]) AverageInt(keySelector func(T) int) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageInt64 returns the arithmetic mean of int64 values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageInt64 operation will:
//   - Apply the keySelector function to each element to extract an int64 value
//   - Calculate the sum of all extracted int64 values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted int64 values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Returns float64 to accommodate fractional results from integer division
//   - Uses float64 internally to handle large int64 values and prevent overflow
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt, AverageFloat, or AverageFloat64
func (e Enumerator[T]) AverageInt64(keySelector func(T) int64) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageInt64 returns the arithmetic mean of int64 values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageInt64 operation will:
//   - Apply the keySelector function to each element to extract an int64 value
//   - Calculate the sum of all extracted int64 values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts an int64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted int64 values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Returns float64 to accommodate fractional results from integer division
//   - Uses float64 internally to handle large int64 values and prevent overflow
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt, AverageFloat, or AverageFloat64
func (e EnumeratorAny[T]) AverageInt64(keySelector func(T) int64) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageFloat returns the arithmetic mean of float32 values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageFloat operation will:
//   - Apply the keySelector function to each element to extract a float32 value
//   - Calculate the sum of all extracted float32 values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a float32 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted float32 values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Returns float64 for higher precision in result
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt, AverageInt64, or AverageFloat64
func (e Enumerator[T]) AverageFloat(keySelector func(T) float32) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageFloat returns the arithmetic mean of float32 values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageFloat operation will:
//   - Apply the keySelector function to each element to extract a float32 value
//   - Calculate the sum of all extracted float32 values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a float32 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted float32 values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - Returns float64 for higher precision in result
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt, AverageInt64, or AverageFloat64
func (e EnumeratorAny[T]) AverageFloat(keySelector func(T) float32) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageFloat64 returns the arithmetic mean of float64 values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageFloat64 operation will:
//   - Apply the keySelector function to each element to extract a float64 value
//   - Calculate the sum of all extracted float64 values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted float64 values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt, AverageInt64, or AverageFloat
func (e Enumerator[T]) AverageFloat64(keySelector func(T) float64) (float64, bool) {
	return averageInternal(e, keySelector)
}

// AverageFloat64 returns the arithmetic mean of float64 values extracted from elements of the enumeration
// using a key selector function.
// This operation is useful for calculating the average numeric value derived from complex elements.
//
// The AverageFloat64 operation will:
//   - Apply the keySelector function to each element to extract a float64 value
//   - Calculate the sum of all extracted float64 values
//   - Divide the sum by the count of elements to compute the arithmetic mean
//   - Return the average as float64 and true if the enumeration is non-empty
//   - Return zero value (0) and false if the enumeration is empty, nil, or keySelector is nil
//   - Process elements sequentially until the end of the enumeration
//   - Handle nil enumerators and nil functions gracefully
//
// Parameters:
//
//	keySelector - a function that extracts a float64 key from each element of type T.
//	              Must be non-nil; if nil, the operation returns (0, false).
//
// Returns:
//
//	The arithmetic mean of extracted float64 values as float64 and true if found,
//	zero value (0) and false otherwise
//
// ⚠️ Performance note: This is a terminal operation that must iterate
// through the entire enumeration to calculate the sum and count.
// For large enumerations, this may be expensive.
//
// ⚠️ Memory note: This operation does not buffer elements, but it must
// process the entire enumeration, which may trigger upstream operations.
//
// Notes:
//   - If the enumerator is nil, returns (0, false)
//   - If keySelector is nil, returns (0, false)
//   - If the enumeration is empty, returns (0, false)
//   - Processes all elements in the enumeration - O(n) time complexity
//   - No elements are buffered - memory efficient
//   - The keySelector function is called exactly once per element
//   - This is a terminal operation that materializes the enumeration
//   - The keySelector function should be deterministic for consistent results
//   - For empty enumerations, no division by zero occurs (returns false)
//   - To calculate average of other numeric types, use AverageInt, AverageInt64, or AverageFloat
func (e EnumeratorAny[T]) AverageFloat64(keySelector func(T) float64) (float64, bool) {
	return averageInternal(e, keySelector)
}

func averageInternal[T any, N allNumber](
	enumerator func(yield func(T) bool),
	keySelector func(T) N,
) (float64, bool) {
	var sum float64
	count := 0

	if enumerator == nil || keySelector == nil {
		return 0, false
	}

	enumerator(func(item T) bool {
		sum += float64(keySelector(item))
		count++
		return true
	})

	if count == 0 {
		return 0, false
	}
	return sum / float64(count), true
}
