package hashcode

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	t.Run("combine single primitive values", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine(42)
		hash2 := Combine(42)
		hash3 := Combine(43)

		if hash1 != hash2 {
			t.Error("Expected same values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different values to produce different hash")
		}
	})

	t.Run("combine multiple primitive values", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine(42, "hello", true)
		hash2 := Combine(42, "hello", true)
		hash3 := Combine(42, "hello", false)

		if hash1 != hash2 {
			t.Error("Expected same values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different values to produce different hash")
		}
	})

	t.Run("combine order matters", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine(1, 2, 3)
		hash2 := Combine(3, 2, 1)

		if hash1 == hash2 {
			t.Error("Expected different order to produce different hash")
		}
	})

	t.Run("combine empty slice", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine()
		hash2 := Combine()

		if hash1 != hash2 {
			t.Error("Expected empty calls to produce same hash")
		}
	})

	t.Run("combine nil values", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine(nil)
		hash2 := Combine(nil)
		hash3 := Combine(0)

		if hash1 != hash2 {
			t.Error("Expected same nil values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected nil and 0 to produce different hash")
		}
	})

	t.Run("combine string values", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine("hello", "world")
		hash2 := Combine("hello", "world")
		hash3 := Combine("hello", "world!")

		if hash1 != hash2 {
			t.Error("Expected same strings to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different strings to produce different hash")
		}
	})

	t.Run("combine boolean values", func(t *testing.T) {
		t.Parallel()

		hash1 := Combine(true, false, true)
		hash2 := Combine(true, false, true)
		hash3 := Combine(true, true, true)

		if hash1 != hash2 {
			t.Error("Expected same booleans to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different booleans to produce different hash")
		}
	})

	t.Run("combine float values", func(t *testing.T) {
		t.Parallel()

		val1 := float64(1.5)
		val2 := float64(2.5)
		val3 := float64(3.5)

		hash1 := Combine(val1, val2)
		hash2 := Combine(val1, val2)
		hash3 := Combine(val1, val3)

		if hash1 != hash2 {
			t.Error("Expected same floats to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different floats to produce different hash")
		}
	})

	t.Run("combine byte slices", func(t *testing.T) {
		t.Parallel()

		data1 := []byte("hello")
		data2 := []byte("hello")
		data3 := []byte("world")

		hash1 := Combine(data1)
		hash2 := Combine(data2)
		hash3 := Combine(data3)

		if hash1 != hash2 {
			t.Error("Expected same byte slices to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different byte slices to produce different hash")
		}
	})

	t.Run("combine struct values", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
			Age  int
		}

		person1 := Person{Name: "Alice", Age: 30}
		person2 := Person{Name: "Alice", Age: 30}
		person3 := Person{Name: "Bob", Age: 30}

		hash1 := Combine(person1)
		hash2 := Combine(person2)
		hash3 := Combine(person3)

		if hash1 != hash2 {
			t.Error("Expected same structs to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different structs to produce different hash")
		}
	})

	t.Run("combine mixed types", func(t *testing.T) {
		t.Parallel()

		type Config struct {
			Name    string
			Enabled bool
			Count   int
		}

		config := Config{Name: "test", Enabled: true, Count: 42}
		data := []byte("binary")

		hash1 := Combine(config, data, float64(1.5), "settings")
		hash2 := Combine(config, data, float64(1.5), "settings")
		hash3 := Combine(config, data, float64(2.5), "settings")

		if hash1 != hash2 {
			t.Error("Expected same mixed values to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different mixed values to produce different hash")
		}
	})

	t.Run("combine pointer values", func(t *testing.T) {
		t.Parallel()

		value := 42
		ptr1 := &value
		ptr2 := &value
		value2 := 42
		ptr3 := &value2

		hash1 := Combine(ptr1)
		hash2 := Combine(ptr2)
		hash3 := Combine(ptr3)

		if hash1 != hash2 {
			t.Error("Expected same pointers to produce same hash")
		}
		_ = hash3
	})

	t.Run("combine slice values (converted to string)", func(t *testing.T) {
		t.Parallel()

		slice1 := []int{1, 2, 3}
		slice2 := []int{1, 2, 3}
		slice3 := []int{1, 2, 4}

		hash1 := Combine(slice1)
		hash2 := Combine(slice2)
		hash3 := Combine(slice3)

		if hash1 != hash2 {
			t.Error("Expected same slices to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different slices to produce different hash")
		}
	})

	t.Run("combine map values (converted to string)", func(t *testing.T) {
		t.Parallel()

		map1 := map[string]int{"a": 1, "b": 2}
		map2 := map[string]int{"a": 1, "b": 2}
		map3 := map[string]int{"a": 1, "b": 3}

		hash1 := Combine(map1)
		hash2 := Combine(map2)
		hash3 := Combine(map3)

		if hash1 != hash2 {
			t.Error("Expected same maps to produce same hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different maps to produce different hash")
		}
	})

	t.Run("consistency across calls", func(t *testing.T) {
		t.Parallel()

		values := []interface{}{42, "hello", true, 3.14}

		hashes := make([]uint64, 10)
		for i := 0; i < 10; i++ {
			hashes[i] = Combine(values...)
		}

		for i := 1; i < 10; i++ {
			if hashes[i] != hashes[0] {
				t.Error("Expected consistent hash codes across multiple calls")
			}
		}
	})

	t.Run("hash distribution", func(t *testing.T) {
		t.Parallel()

		hashes := make(map[uint64]bool)
		collisions := 0

		for i := 0; i < 1000; i++ {
			hash := Combine(fmt.Sprintf("value%d", i), i, float64(i)/3.0)
			if hashes[hash] {
				collisions++
			} else {
				hashes[hash] = true
			}
		}

		if collisions > 50 {
			t.Errorf("Too many hash collisions: %d out of 1000", collisions)
		}
	})
}

func TestCombineHashes(t *testing.T) {
	t.Run("combine single hash value", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(42)
		hash2 := CombineHashes(42)
		hash3 := CombineHashes(43)

		if hash1 != hash2 {
			t.Error("Expected same values to produce same combined hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different values to produce different combined hash")
		}
	})

	t.Run("combine multiple hash values", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(1, 2, 3)
		hash2 := CombineHashes(1, 2, 3)
		hash3 := CombineHashes(1, 2, 4)

		if hash1 != hash2 {
			t.Error("Expected same values to produce same combined hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different values to produce different combined hash")
		}
	})

	t.Run("combine order matters", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(1, 2, 3)
		hash2 := CombineHashes(3, 2, 1)

		if hash1 == hash2 {
			t.Error("Expected different order to produce different combined hash")
		}
	})

	t.Run("combine empty slice", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes()
		hash2 := CombineHashes()

		if hash1 != hash2 {
			t.Error("Expected empty calls to produce same combined hash")
		}

		if hash1 != seed {
			t.Errorf("Expected empty call to return seed value %d, got %d", seed, hash1)
		}
	})

	t.Run("combine with zero values", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(0)
		hash2 := CombineHashes(0)
		hash3 := CombineHashes(1)

		if hash1 != hash2 {
			t.Error("Expected same zero values to produce same combined hash")
		}

		if hash1 == hash3 {
			t.Error("Expected zero and non-zero to produce different combined hash")
		}
	})

	t.Run("combine multiple zero values", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(0, 0, 0)
		hash2 := CombineHashes(0, 0, 0)
		hash3 := CombineHashes(0, 0, 1)

		if hash1 != hash2 {
			t.Error("Expected same zero values to produce same combined hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different values to produce different combined hash")
		}
	})

	t.Run("combine large values", func(t *testing.T) {
		t.Parallel()

		large1 := uint64(0xFFFFFFFFFFFFFFFF)
		large2 := uint64(0xAAAAAAAAAAAAAAAA)
		large3 := uint64(0x5555555555555555)

		hash1 := CombineHashes(large1, large2)
		hash2 := CombineHashes(large1, large2)
		hash3 := CombineHashes(large1, large3)

		if hash1 != hash2 {
			t.Error("Expected same large values to produce same combined hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different large values to produce different combined hash")
		}
	})

	t.Run("combine consistent results", func(t *testing.T) {
		t.Parallel()

		values := []uint64{1, 2, 3, 4, 5}

		hashes := make([]uint64, 10)
		for i := 0; i < 10; i++ {
			hashes[i] = CombineHashes(values...)
		}

		for i := 1; i < 10; i++ {
			if hashes[i] != hashes[0] {
				t.Error("Expected consistent combined hash codes across multiple calls")
			}
		}
	})

	t.Run("combine mathematical properties", func(t *testing.T) {
		t.Parallel()

		singleValue := uint64(42)
		expected := seed*multiplier + singleValue
		actual := CombineHashes(singleValue)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
		value1, value2 := uint64(10), uint64(20)
		expected = (seed*multiplier+value1)*multiplier + value2
		actual = CombineHashes(value1, value2)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("combine overflow handling", func(t *testing.T) {
		t.Parallel()

		largeValues := []uint64{
			0xFFFFFFFFFFFFFFFF,
			0xFFFFFFFFFFFFFFFE,
			0xFFFFFFFFFFFFFFFD,
		}
		hash1 := CombineHashes(largeValues...)
		hash2 := CombineHashes(largeValues...)

		if hash1 != hash2 {
			t.Error("Expected consistent results even with large values")
		}
	})

	t.Run("combine different lengths", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(1, 2, 3)
		hash2 := CombineHashes(1, 2, 3, 4)

		if hash1 == hash2 {
			t.Error("Expected different lengths to produce different combined hash")
		}
	})

	t.Run("combine with repeated values", func(t *testing.T) {
		t.Parallel()

		hash1 := CombineHashes(1, 1, 1)
		hash2 := CombineHashes(1, 1, 1)
		hash3 := CombineHashes(1, 1, 2)

		if hash1 != hash2 {
			t.Error("Expected same repeated values to produce same combined hash")
		}

		if hash1 == hash3 {
			t.Error("Expected different repeated values to produce different combined hash")
		}
	})

	t.Run("combine hash distribution", func(t *testing.T) {
		t.Parallel()

		hashes := make(map[uint64]bool)
		collisions := 0

		for i := 0; i < 1000; i++ {
			values := []uint64{uint64(i), uint64(i * 2), uint64(i * 3)}
			hash := CombineHashes(values...)
			if hashes[hash] {
				collisions++
			} else {
				hashes[hash] = true
			}
		}

		if collisions > 50 {
			t.Errorf("Too many hash collisions: %d out of 1000", collisions)
		}
	})

	t.Run("combine deterministic with same inputs", func(t *testing.T) {
		t.Parallel()

		inputs := []uint64{42, 123, 456, 789, 999}

		result1 := CombineHashes(inputs...)
		result2 := CombineHashes(inputs...)

		if result1 != result2 {
			t.Error("CombineHashes should be deterministic with same inputs")
		}
	})
}

func BenchmarkCombineHashes(b *testing.B) {
	b.Run("combine single value", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = CombineHashes(42)
		}
	})

	b.Run("combine multiple values", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = CombineHashes(1, 2, 3, 4, 5)
		}
	})

	b.Run("combine many values", func(b *testing.B) {
		values := make([]uint64, 100)
		for i := 0; i < 100; i++ {
			values[i] = uint64(i)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = CombineHashes(values...)
		}
	})

	b.Run("combine large values", func(b *testing.B) {
		largeValues := []uint64{
			0xFFFFFFFFFFFFFFFF,
			0xAAAAAAAAAAAAAAAA,
			0x5555555555555555,
			0x123456789ABCDEF0,
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = CombineHashes(largeValues...)
		}
	})
}

func BenchmarkCombine(b *testing.B) {
	b.Run("combine single value", func(b *testing.B) {
		value := 42

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Combine(value)
		}
	})

	b.Run("combine multiple primitives", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Combine(42, "hello", true, 3.14)
		}
	})

	b.Run("combine struct", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}
		person := Person{Name: "Alice", Age: 30}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Combine(person)
		}
	})

	b.Run("combine slice", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Combine(slice)
		}
	})

	b.Run("combine many values", func(b *testing.B) {
		values := make([]interface{}, 100)
		for i := 0; i < 100; i++ {
			values[i] = fmt.Sprintf("value%d", i)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Combine(values...)
		}
	})
}
