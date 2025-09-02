package comparer

// ByField creates an EqualityComparer that compares two values of type T by comparing
// a specific field selected by the provided field selector function.
//
// This comparer is particularly useful when:
//   - T is a struct type containing non-comparable fields (slices, maps, etc.)
//   - You want to compare instances based on a single comparable field
//   - You need custom equality logic for complex types
//
// The returned EqualityComparer will:
//   - Extract the field value from both input values using fieldSelector
//   - Compare the extracted field values using == operator
//   - Return true if field values are equal, false otherwise
//
// Type Parameters:
//   T - the type of values to compare (can contain non-comparable fields)
//   F - the type of the field to compare (must be comparable)
//
// Parameters:
//   fieldSelector - a function that extracts the comparable field from type T
//
// Returns:
//   An EqualityComparer[T] that compares values by the selected field
//
// ⚠️ Important: The field type F must be comparable (no slices, maps, functions)
// If F is non-comparable, the == operation will cause a compile-time error
//
// ⚠️ Performance note: The field selector function is called twice per comparison
// Consider the cost of field extraction when using expensive selectors
//
// Notes:
//   - The field selector should be deterministic and side-effect free
//   - For nested field access, you can use chained selectors
//   - Works with pointer receivers and value receivers
//   - Thread safety depends on the field selector implementation
//   - Common use cases include ID-based comparison, name-based comparison
//   - Can be composed with other comparers for complex logic
//
// ⚠️ Limitations:
//   - Cannot compare fields of non-comparable types (slices, maps, funcs)
//   - Field selector must not panic or have side effects
//   - No automatic handling of nil pointers in field selection
func ByField[T any, F comparable](fieldSelector func(T) F) EqualityComparer[T] {
	return func(a, b T) bool {
		return fieldSelector(a) == fieldSelector(b)
	}
}

// Composite creates an EqualityComparer that combines multiple EqualityComparer instances
// using logical AND operation. All provided comparers must return true for the
// composite comparer to return true.
//
// This comparer is useful when:
//   - You need to compare complex objects by multiple fields
//   - You want to combine different comparison strategies
//   - You need to create hierarchical equality logic
//   - You want to reuse existing comparers in combinations
//
// The returned EqualityComparer will:
//   - Call each provided comparer in the order they were passed
//   - Return false immediately if any comparer returns false (short-circuit)
//   - Return true only if all comparers return true
//   - Handle empty comparer slice (returns true for any inputs)
//
// Type Parameters:
//   T - the type of values to compare
//
// Parameters:
//   comparers - variadic list of EqualityComparer[T] to combine
//
// Returns:
//   An EqualityComparer[T] that combines all provided comparers with AND logic
//
// ⚠️ Performance note: Comparers are evaluated in order until one returns false
// Place the most selective or fastest comparers first for better performance
//
// Notes:
//   - If no comparers are provided, the result is always true
//   - If one comparer is provided, it behaves identically to that comparer
//   - Each comparer should be deterministic and side-effect free
//   - Thread safety depends on the individual comparer implementations
//   - Common use cases include multi-field comparison, composite key comparison
//   - Can be nested to create complex comparison hierarchies
//
// ⚠️ Important considerations:
//   - Order matters for performance but not for correctness
//   - Each comparer is called with the same input parameters
//   - No duplicate comparer detection is performed
//   - Be careful with expensive comparers in the chain
func Composite[T any](comparers ...EqualityComparer[T]) EqualityComparer[T] {
	return func(a, b T) bool {
		for _, comparer := range comparers {
			if !comparer(a, b) {
				return false
			}
		}
		return true
	}
}

// Custom creates an EqualityComparer from a custom equality function.
// This is useful when you need complete control over the equality comparison logic
// or when working with complex types that require special comparison handling.
//
// This comparer is particularly useful when:
//   - Built-in comparers (ByField, Composite) don't meet your needs
//   - You need complex comparison logic that spans multiple fields
//   - You want to implement domain-specific equality rules
//   - You need to compare types with special equality semantics
//   - You want to wrap existing comparison functions
//
// The returned EqualityComparer will:
//   - Delegate all equality checks to the provided equalFunc
//   - Pass both input parameters directly to equalFunc
//   - Return the result of equalFunc unchanged
//   - Preserve the behavior and performance characteristics of equalFunc
//
// Type Parameters:
//   T - the type of values to compare
//
// Parameters:
//   equalFunc - a function that determines equality between two values of type T
//
// Returns:
//   An EqualityComparer[T] that uses the provided function for equality comparison
//
// ⚠️ Important: The equalFunc must be:
//   - Deterministic (same inputs always produce same output)
//   - Reflexive (equalFunc(x, x) should return true)
//   - Symmetric (equalFunc(x, y) should equal equalFunc(y, x))
//   - Transitive (if equalFunc(x, y) and equalFunc(y, z), then equalFunc(x, z))
//
// ⚠️ Performance note: The performance characteristics of the returned comparer
// are identical to those of the provided equalFunc. No additional overhead is added.
//
// Notes:
//   - The equalFunc should be side-effect free
//   - Thread safety depends on the equalFunc implementation
//   - Can be used to wrap existing comparison logic
//   - Useful for implementing comparison with tolerance (floating point, time, etc.)
//   - Can handle nil values gracefully if equalFunc supports them
//   - Common use cases include custom business logic, approximate equality, complex struct comparison
//
// ⚠️ Important considerations:
//   - No validation is performed on the equalFunc - incorrect implementation may cause unexpected behavior
//   - The equalFunc is called directly - ensure it handles all possible input values
//   - Consider using ByField or Composite for simpler cases before resorting to Custom
//   - Be careful with recursive types to avoid infinite loops
func Custom[T any](equalFunc func(T, T) bool) EqualityComparer[T] {
	return equalFunc
}
