package comparer

import (
	"github.com/ahatornn/enumerable/hashcode"
)

// EqualityComparer defines a type that determines equality and hash codes for values of type T.
// This approach is similar to IEqualityComparer<T> in C# and allows efficient comparison
// and hashing of non-comparable types.
//
// Type Parameters:
//
//	T - the type of values to compare (can be any type including non-comparable ones)
type EqualityComparer[T any] interface {
	// Equals determines whether two values of type T are equal.
	// Returns true if the values are considered equal, false otherwise.
	Equals(x, y T) bool

	// GetHashCode returns a hash code for the specified value.
	// Values that are equal according to Equals() should return the same hash code.
	GetHashCode(x T) uint64
}

// New creates an EqualityComparer from separate equals and hash code functions.
// This is the recommended way to create custom equality comparers.
//
// Parameters:
//
//	equals - function that determines equality between two values
//	getHashCode - function that computes hash code for a value
//
// Returns:
//
//	An EqualityComparer[T] that uses the provided functions
//
// ⚠️ Important: The functions must follow these rules:
//   - Equals must be reflexive, symmetric, and transitive
//   - GetHashCode must return the same value for equal objects
//   - GetHashCode should distribute well to minimize collisions
func New[T any](equals func(T, T) bool, getHashCode func(T) uint64) EqualityComparer[T] {
	return &comparer[T]{
		equals:      equals,
		getHashCode: getHashCode,
	}
}

// Default creates an EqualityComparer that uses built-in equality (==) and hash operations
// for comparable types. This is the most efficient comparer for types that support
// direct comparison and is equivalent to the default equality behavior in Go.
//
// This comparer is particularly useful when:
//   - Working with built-in comparable types (int, string, bool, etc.)
//   - Working with struct types containing only comparable fields
//   - You want the standard Go equality semantics
//   - Maximum performance is required for simple comparisons
//   - You need a fallback comparer for comparable types
//
// The returned EqualityComparer will:
//   - Compare values using the built-in == operator
//   - Generate hash codes using simple hashing of the value's memory representation
//   - Provide O(1) average time complexity for both Equals and GetHashCode operations
//   - Require zero memory allocations during comparison or hashing
//
// Type Parameters:
//
//	T - the comparable type to compare (must satisfy comparable constraint)
//
// Returns:
//
//	An EqualityComparer[T] that uses built-in == and hash operations
//
// ⚠️ Important: The type T must be comparable (no slices, maps, functions, or structs
// containing non-comparable fields). Attempting to use this with non-comparable types
// will result in a compile-time error.
//
// ⚠️ Performance note: This is the fastest possible comparer for comparable types
// as it uses direct language operations with no function call overhead.
// Both Equals and GetHashCode operations are O(1) with 0 allocations.
//
// Notes:
//   - For struct types, all fields must be comparable
//   - Pointer equality is used for pointer types (same memory address)
//   - Interface equality follows Go's interface equality rules
//   - Arrays are compared element-wise if all elements are comparable
//   - Channel equality compares the channel references, not contents
//   - Thread safe - can be used concurrently without synchronization
//   - Hash codes are consistent for the lifetime of the process
//   - Common use cases include primitive types, simple structs, and ID-based comparisons
//
// ⚠️ Limitations:
//   - Cannot be used with non-comparable types (slices, maps, functions)
//   - Structs containing non-comparable fields are not supported
//   - Hash distribution depends on the value's memory representation
//   - Interface values are compared by both dynamic type and value
//   - Complex types may have less optimal hash distribution
//
// ⚠️ Important considerations:
//   - The comparable constraint is enforced at compile time
//   - Hash codes are generated from the value's memory representation
//   - For best performance, use with simple types and small structs
//   - Consider ByField for structs where you only need to compare specific fields
func Default[T comparable]() EqualityComparer[T] {
	return New(
		func(val1, val2 T) bool { return val1 == val2 },
		func(val T) uint64 { return hashcode.Compute(val) },
	)
}

// comparer is the internal implementation of EqualityComparer interface
type comparer[T any] struct {
	equals      func(T, T) bool
	getHashCode func(T) uint64
}

func (c *comparer[T]) Equals(x, y T) bool {
	return c.equals(x, y)
}

func (c *comparer[T]) GetHashCode(x T) uint64 {
	return c.getHashCode(x)
}
