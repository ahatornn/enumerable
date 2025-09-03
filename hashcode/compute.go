package hashcode

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

// Compute generates a hash code for any value using type-specific serialization
// and the FNV-1a 64-bit hash algorithm. This function provides a generic way
// to compute hash codes for values of any type without using reflection.
//
// Parameters:
//
//	value - any value to compute hash code for (can be nil)
//
// Returns:
//
//	A uint64 FNV-1a hash code for the provided value
//
// ⚠️ Performance note: Performance varies by type:
//   - Primitive types (int, float, string, bool): O(1) - very fast
//   - Byte slices: O(n) where n is slice length
//   - Complex types: O(m) where m is string representation length
//   - Uses FNV-1a which is fast but not cryptographically secure
//
// ⚠️ Memory note: Minimal allocations for primitive types. Byte slices are
// included directly. Complex types may allocate during fmt.Sprintf conversion.
//
// Notes:
//   - Thread safe - can be called concurrently
//   - Consistent results for equal values within the same process
//   - Handles nil values gracefully with consistent hash codes
//   - Uses FNV-1a 64-bit hash algorithm for good distribution and speed
//   - Little-endian byte order for numeric types ensures consistency
//   - Common use cases include generic collections, caching, and debugging
func Compute(value any) uint64 {
	h := fnv.New64a()

	switch v := value.(type) {
	case int:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case int8:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case int16:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case int32:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case int64:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case uint:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case uint8:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case uint16:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case uint32:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case uint64:
		binary.Write(h, binary.LittleEndian, v)
	case float32:
		binary.Write(h, binary.LittleEndian, uint32(v))
	case float64:
		binary.Write(h, binary.LittleEndian, uint64(v))
	case string:
		h.Write([]byte(v))
	case bool:
		if v {
			h.Write([]byte{1})
		} else {
			h.Write([]byte{0})
		}
	case []byte:
		h.Write(v)
	case nil:
		h.Write([]byte("nil"))
	default:
		h.Write([]byte(fmt.Sprintf("%v", v)))
	}

	return h.Sum64()
}
