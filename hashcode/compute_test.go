package hashcode

import (
	"math"
	"testing"
)

func TestCompute(t *testing.T) {
	t.Run("compute int values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(42)
		hash2 := Compute(42)
		hash3 := Compute(43)

		if hash1 != hash2 {
			t.Error("Expected same int values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different int values to produce different hash")
		}
	})

	t.Run("compute different integer types", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(int(42))
		hash2 := Compute(int8(42))
		hash3 := Compute(int16(42))
		hash4 := Compute(int32(42))
		hash5 := Compute(int64(42))

		hashes := []uint64{hash1, hash2, hash3, hash4, hash5}
		uniqueHashes := make(map[uint64]bool)
		for _, h := range hashes {
			uniqueHashes[h] = true
		}
		_ = uniqueHashes
	})

	t.Run("compute unsigned integer types", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(uint(42))
		hash5 := Compute(uint64(42))

		if hash1 != Compute(uint(42)) {
			t.Error("Expected consistency for uint")
		}

		if hash5 != Compute(uint64(42)) {
			t.Error("Expected consistency for uint64")
		}
	})

	t.Run("compute float values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(1.5)
		hash2 := Compute(1.5)
		hash3 := Compute(2.5)

		if hash1 != hash2 {
			t.Error("Expected same float values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different float values to produce different hash")
		}
	})

	t.Run("compute float32 vs float64", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(float32(1.5))
		hash2 := Compute(float64(1.5))

		if hash1 == hash2 {
			t.Log("Warning: float32 and float64 with same value produced same hash")
		}
	})

	t.Run("compute string values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute("hello")
		hash2 := Compute("hello")
		hash3 := Compute("world")

		if hash1 != hash2 {
			t.Error("Expected same strings to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different strings to produce different hash")
		}
	})

	t.Run("compute empty string", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute("")
		hash2 := Compute("")
		hash3 := Compute("a")

		if hash1 != hash2 {
			t.Error("Expected same empty strings to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected empty and non-empty strings to produce different hash")
		}
	})

	t.Run("compute boolean values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(true)
		hash2 := Compute(true)
		hash3 := Compute(false)
		hash4 := Compute(false)

		if hash1 != hash2 {
			t.Error("Expected same true values to produce same hash")
		}

		if hash3 != hash4 {
			t.Error("Expected same false values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected true and false to produce different hash")
		}
	})

	t.Run("compute byte slice values", func(t *testing.T) {
		t.Parallel()

		data1 := []byte("hello")
		data2 := []byte("hello")
		data3 := []byte("world")

		hash1 := Compute(data1)
		hash2 := Compute(data2)
		hash3 := Compute(data3)

		if hash1 != hash2 {
			t.Error("Expected same byte slices to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different byte slices to produce different hash")
		}
	})

	t.Run("compute empty byte slice", func(t *testing.T) {
		t.Parallel()

		empty1 := []byte{}
		empty2 := []byte{}

		hash1 := Compute(empty1)
		hash2 := Compute(empty2)
		hash3 := Compute([]byte("a"))

		if hash1 != hash2 {
			t.Error("Expected same empty byte slices to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected empty and non-empty byte slices to produce different hash")
		}
	})

	t.Run("compute nil value", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(nil)
		hash2 := Compute(nil)
		hash3 := Compute(0)

		if hash1 != hash2 {
			t.Error("Expected same nil values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected nil and 0 to produce different hash")
		}
	})

	t.Run("compute struct values", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
			Age  int
		}

		person1 := Person{Name: "Alice", Age: 30}
		person2 := Person{Name: "Alice", Age: 30}
		person3 := Person{Name: "Bob", Age: 30}

		hash1 := Compute(person1)
		hash2 := Compute(person2)
		hash3 := Compute(person3)

		if hash1 != hash2 {
			t.Error("Expected same structs to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different structs to produce different hash")
		}
	})

	t.Run("compute pointer values", func(t *testing.T) {
		t.Parallel()

		value := 42
		ptr1 := &value
		ptr2 := &value
		value2 := 42
		ptr3 := &value2

		hash1 := Compute(ptr1)
		hash2 := Compute(ptr2)
		hash3 := Compute(ptr3)

		if hash1 != hash2 {
			t.Error("Expected same pointers to produce same hash")
		}

		if hash1 == hash3 {
			t.Log("Note: Different pointers to equal values may produce same hash")
		}
	})

	t.Run("compute slice values", func(t *testing.T) {
		t.Parallel()

		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 3}
		slice3 := []int{1, 2, 4}

		hash1 := Compute(slice1)
		hash2 := Compute(slice2)
		hash3 := Compute(slice3)

		if hash1 != hash2 {
			t.Error("Expected same slices to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different slices to produce different hash")
		}
	})

	t.Run("compute map values", func(t *testing.T) {
		t.Parallel()

		map1 := map[string]int{"a": 1, "b": 2}
		map2 := map[string]int{"a": 1, "b": 2}
		map3 := map[string]int{"a": 1, "b": 3}

		hash1 := Compute(map1)
		hash2 := Compute(map2)
		hash3 := Compute(map3)

		if hash1 != hash2 {
			t.Error("Expected same maps to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different maps to produce different hash")
		}
	})

	t.Run("compute consistency across calls", func(t *testing.T) {
		t.Parallel()

		value := "test string value"

		hashes := make([]uint64, 10)
		for i := 0; i < 10; i++ {
			hashes[i] = Compute(value)
		}

		for i := 1; i < 10; i++ {
			if hashes[i] != hashes[0] {
				t.Error("Expected consistent hash codes across multiple calls")
			}
		}
	})

	t.Run("compute zero values of different types", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(0)
		hash2 := Compute("")
		hash3 := Compute(false)
		hash4 := Compute(0.0)

		if hash1 == hash2 || hash1 == hash3 || hash1 == hash4 ||
			hash2 == hash3 || hash2 == hash4 || hash3 == hash4 {
			t.Log("Note: Different zero types may produce some hash collisions")
		}
	})

	t.Run("compute special float values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(0.0)
		hash3 := Compute(math.NaN())
		if hash1 != Compute(0.0) {
			t.Error("Expected 0.0 to be consistent")
		}

		if hash3 != Compute(math.NaN()) {
			t.Error("Expected NaN to be consistent (within same process)")
		}
	})

	t.Run("compute complex nested structures", func(t *testing.T) {
		t.Parallel()

		type Address struct {
			Street string
			City   string
		}

		type Person struct {
			Name    string
			Age     int
			Address Address
			Tags    []string
		}

		person1 := Person{
			Name: "Alice",
			Age:  30,
			Address: Address{
				Street: "123 Main St",
				City:   "Anytown",
			},
			Tags: []string{"developer", "go"},
		}

		person2 := Person{
			Name: "Alice",
			Age:  30,
			Address: Address{
				Street: "123 Main St",
				City:   "Anytown",
			},
			Tags: []string{"developer", "go"},
		}

		hash1 := Compute(person1)
		hash2 := Compute(person2)

		if hash1 != hash2 {
			t.Error("Expected identical complex structures to produce same hash")
		}
	})

	t.Run("compute uint8 values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(uint8(42))
		hash2 := Compute(uint8(42))
		hash3 := Compute(uint8(43))

		if hash1 != hash2 {
			t.Error("Expected same uint8 values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different uint8 values to produce different hash")
		}
	})

	t.Run("compute uint16 values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(uint16(1000))
		hash2 := Compute(uint16(1000))
		hash3 := Compute(uint16(1001))

		if hash1 != hash2 {
			t.Error("Expected same uint16 values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different uint16 values to produce different hash")
		}
	})

	t.Run("compute uint32 values", func(t *testing.T) {
		t.Parallel()

		hash1 := Compute(uint32(1000000))
		hash2 := Compute(uint32(1000000))
		hash3 := Compute(uint32(1000001))

		if hash1 != hash2 {
			t.Error("Expected same uint32 values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different uint32 values to produce different hash")
		}
	})

	t.Run("compute different unsigned integer types with same value", func(t *testing.T) {
		t.Parallel()

		hash8 := Compute(uint8(42))
		hash16 := Compute(uint16(42))
		hash32 := Compute(uint32(42))
		hash64 := Compute(uint64(42))
		hashUint := Compute(uint(42))

		hashes := []uint64{hash8, hash16, hash32, hash64, hashUint}
		uniqueHashes := make(map[uint64]bool)
		for _, h := range hashes {
			uniqueHashes[h] = true
		}

		_ = uniqueHashes
	})

	t.Run("compute unsigned integer zero values", func(t *testing.T) {
		t.Parallel()

		hash8 := Compute(uint8(0))
		hash16 := Compute(uint16(0))
		hash32 := Compute(uint32(0))
		hash64 := Compute(uint64(0))
		hashUint := Compute(uint(0))

		if hash8 != Compute(uint8(0)) {
			t.Error("Expected consistency for uint8 zero")
		}

		if hash16 != Compute(uint16(0)) {
			t.Error("Expected consistency for uint16 zero")
		}

		if hash32 != Compute(uint32(0)) {
			t.Error("Expected consistency for uint32 zero")
		}

		if hash64 != Compute(uint64(0)) {
			t.Error("Expected consistency for uint64 zero")
		}

		if hashUint != Compute(uint(0)) {
			t.Error("Expected consistency for uint zero")
		}
	})

	t.Run("compute maximum values for unsigned integer types", func(t *testing.T) {
		t.Parallel()

		hash8 := Compute(uint8(math.MaxUint8))
		hash16 := Compute(uint16(math.MaxUint16))
		hash32 := Compute(uint32(math.MaxUint32))
		hash64 := Compute(uint64(math.MaxUint64))

		if hash8 != Compute(uint8(math.MaxUint8)) {
			t.Error("Expected consistency for uint8 max")
		}

		if hash16 != Compute(uint16(math.MaxUint16)) {
			t.Error("Expected consistency for uint16 max")
		}

		if hash32 != Compute(uint32(math.MaxUint32)) {
			t.Error("Expected consistency for uint32 max")
		}

		if hash64 != Compute(uint64(math.MaxUint64)) {
			t.Error("Expected consistency for uint64 max")
		}
	})

	t.Run("compute unsigned integer edge cases", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			name  string
			value interface{}
		}{
			{"uint8 max", uint8(math.MaxUint8)},
			{"uint16 max", uint16(math.MaxUint16)},
			{"uint32 max", uint32(math.MaxUint32)},
			{"uint64 max", uint64(math.MaxUint64)},
			{"uint8 mid", uint8(128)},
			{"uint16 mid", uint16(32768)},
			{"uint32 mid", uint32(2147483648)},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				hash1 := Compute(tc.value)
				hash2 := Compute(tc.value)

				if hash1 != hash2 {
					t.Errorf("Expected consistency for %s", tc.name)
				}
			})
		}
	})
}

func BenchmarkCompute(b *testing.B) {
	b.Run("compute int", func(b *testing.B) {
		value := 42

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})

	b.Run("compute string", func(b *testing.B) {
		value := "hello world test string"

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})

	b.Run("compute bool", func(b *testing.B) {
		value := true

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})

	b.Run("compute float64", func(b *testing.B) {
		value := 3.141592653589793

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})

	b.Run("compute byte slice", func(b *testing.B) {
		value := []byte("binary data for testing hash computation performance")

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})

	b.Run("compute struct", func(b *testing.B) {
		type TestStruct struct {
			A int
			B string
			C bool
			D float64
		}

		value := TestStruct{
			A: 42,
			B: "test",
			C: true,
			D: 3.14,
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})

	b.Run("compute nil", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(nil)
		}
	})

	b.Run("compute slice", func(b *testing.B) {
		value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Compute(value)
		}
	})
}
