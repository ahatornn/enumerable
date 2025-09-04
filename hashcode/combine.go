package hashcode

const (
	seed       = 17
	multiplier = 31
)

// Combine computes a combined hash code from multiple values using the same algorithm
// as Composite comparer.
//
// Parameters:
//
//	values - variadic list of values to combine into a single hash code
//
// Returns:
//
//	A uint64 hash code computed by combining all input values
//
// ⚠️ Performance note: Time complexity is O(n) where n is the number of values.
func Combine(values ...any) uint64 {
	var hash uint64 = seed
	for _, value := range values {
		hash = hash*multiplier + Compute(value)
	}
	return hash
}

// CombineHashes computes a combined hash code from multiple individual hash codes
// using a prime-based combination algorithm. This function provides a standardized
// way to combine multiple hash codes into a single hash value with good distribution
// properties, similar to HashCode.Combine() in C# or Guava's HashCode.combine().
//
// Parameters:
//
//	values - variadic list of uint64 hash codes to combine
//
// Returns:
//
//	A uint64 hash code computed by combining all input hash codes
//
// ⚠️ Performance note: Time complexity is O(n) where n is the number of values.
// This is a very fast operation involving only simple arithmetic operations.
// Memory allocation is zero - no additional memory is required.
func CombineHashes(values ...uint64) uint64 {
	var hash uint64 = seed
	for _, value := range values {
		hash = hash*multiplier + value
	}
	return hash
}
