package comparer

// ComparerFunc defines a function type that compares two values of type T and returns their relative ordering.
// It is used throughout the enumerable library for sorting, ordering, and comparison operations.
//
// Type Parameters:
//
//	T - the type of values to compare (can be any type)
//
// The function should return:
//
//	-1 if x is less than y
//	 0 if x is equal to y
//	+1 if x is greater than y
//
// Implementations should ensure the following mathematical properties:
//   - Consistency: f(x, y) should always return the same result for identical inputs
//   - Antisymmetry: if f(x, y) < 0 then f(y, x) > 0
//   - Transitivity: if f(x, y) < 0 and f(y, z) < 0 then f(x, z) < 0
//   - Equality: f(x, y) == 0 if and only if x and y are considered equal
//
// Notes:
//   - For natural ordering of built-in types, use predefined comparer functions
//   - For custom ordering logic, create ComparerFunc instances
//   - Thread safety depends on the function implementation
//   - Nil handling should be consistent within the function
type ComparerFunc[T any] func(x, y T) int

// KeyComparer creates a ComparerFunc[T] based on a key selector function and a comparer for the key type.
// This is useful for defining custom sort orders for complex types by comparing a specific extracted field or property.
//
// Type Parameters:
//
//	T - the type of the elements to be compared (can be any type)
//	K - the type of the key extracted from T (must be a comparable type)
//
// Parameters:
//
//	keySelector - a function that extracts the key value of type K from an element of type T
//	keyComparer - a ComparerFunc[K] that defines the comparison logic for the extracted keys
//
// Returns:
//
//	A new ComparerFunc[T] that compares elements by first applying the keySelector to each,
//	and then using the provided keyComparer to compare the resulting keys.
//
// The returned function adheres to the ComparerFunc contract by returning:
//   - A negative value if the key of x is less than the key of y according to keyComparer
//   - Zero if the key of x is equal to the key of y according to keyComparer
//   - A positive value if the key of x is greater than the key of y according to keyComparer
//
// Notes:
//   - The stability and properties (like transitivity) of the returned comparer depend on the keyComparer provided.
//   - The keySelector function should be deterministic and side-effect free for consistent results.
//   - Thread safety depends on the implementations of keySelector and keyComparer.
func KeyComparer[T any, K comparable](keySelector func(T) K, keyComparer ComparerFunc[K]) ComparerFunc[T] {
	return func(a, b T) int {
		keyA := keySelector(a)
		keyB := keySelector(b)
		return keyComparer(keyA, keyB)
	}
}
