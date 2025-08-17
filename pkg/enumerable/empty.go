package enumerable

// Empty returns an empty enumerator of type T that yields no values.
//
// The returned Enumerator[T] will immediately terminate any range loop
// without executing the loop body, as there are no values to enumerate.
//
// Returns:
//   An empty Enumerator[T] that can be used in range loops (Go 1.22+).
//
// Notes:
// - Can represent "no results" in a type-safe way
// - Works with any comparable type T
func Empty[T comparable]() Enumerator[T] {
	return func(yield func(T) bool) {}
}

// EmptyAny returns an empty enumerator of type T that yields no values.
//
// The returned EnumeratorAny[T] will immediately terminate any range loop
// without executing the loop body, as there are no values to enumerate.
//
// Returns:
//   An empty EnumeratorAny[T] that can be used in range loops (Go 1.22+).
//
// Notes:
// - Can represent "no results" in a type-safe way
// - Works with any type T (no constraints)
func EmptyAny[T any]() EnumeratorAny[T] {
	return func(yield func(T) bool) {}
}
