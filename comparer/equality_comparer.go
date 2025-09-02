package comparer

// EqualityComparer defines a function type that determines whether two values of type T are equal.
// This is useful for comparing non-comparable types such as structs containing slices or maps.
//
// Type Parameters:
//
//	T - the type of values to compare (can be any type including non-comparable ones)
//
// The function should return:
//
//	true if the values are considered equal
//	false if the values are considered different
//
// Notes:
//   - For comparable types, simple equality (==) can be used
//   - For non-comparable types, custom equality logic must be implemented
//   - The function should be consistent and deterministic
//   - Thread safety depends on the implementation
type EqualityComparer[T any] func(x, y T) bool
