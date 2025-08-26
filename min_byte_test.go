package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinByte(t *testing.T) {
	t.Run("min byte from mixed values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 200, 25, 150})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min 25, got %d", min)
		}
	})

	t.Run("min byte from all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 100, 100})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 100 {
			t.Errorf("Expected min 100, got %d", min)
		}
	})

	t.Run("min byte with zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 0, 25, 150})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})

	t.Run("min byte single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{42})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 42 {
			t.Errorf("Expected min 42, got %d", min)
		}
	})

	t.Run("min byte maximum value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{255, 254, 253})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 253 {
			t.Errorf("Expected min 253, got %d", min)
		}
	})

	t.Run("min byte empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{})

		min, ok := enumerator.MinByte(selector.Byte)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for empty slice, got %d", min)
		}
	})

	t.Run("min byte nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[byte] = nil

		min, ok := enumerator.MinByte(selector.Byte)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil enumerator, got %d", min)
		}
	})

	t.Run("min byte with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 25})

		min, ok := enumerator.MinByte(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil keySelector, got %d", min)
		}
	})

	t.Run("min byte with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			ID   int
			Size byte
		}

		data := []Data{
			{ID: 1, Size: 100},
			{ID: 2, Size: 50},
			{ID: 3, Size: 200},
		}

		enumerator := FromSlice(data)
		min, ok := enumerator.MinByte(func(d Data) byte { return d.Size })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 50 {
			t.Errorf("Expected min 50, got %d", min)
		}
	})

	t.Run("min byte with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 25, 200, 25, 150})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min 25, got %d", min)
		}
	})
}

func TestMinByteStruct(t *testing.T) {
	t.Run("min byte from struct field", func(t *testing.T) {
		t.Parallel()
		type File struct {
			Name string
			Size byte
		}

		files := []File{
			{Name: "file1.txt", Size: 100},
			{Name: "file2.txt", Size: 50},
			{Name: "file3.txt", Size: 200},
		}

		enumerator := FromSlice(files)
		min, ok := enumerator.MinByte(func(f File) byte { return f.Size })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 50 {
			t.Errorf("Expected min 50, got %d", min)
		}
	})

	t.Run("min byte from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings   []string
			BufferSize byte
		}

		configs := []Config{
			{Settings: []string{"a", "b"}, BufferSize: 100},
			{Settings: []string{"c"}, BufferSize: 50},
			{Settings: []string{"d", "e", "f"}, BufferSize: 200},
		}

		enumerator := FromSliceAny(configs)
		min, ok := enumerator.MinByte(func(c Config) byte { return c.BufferSize })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 50 {
			t.Errorf("Expected min 50, got %d", min)
		}
	})
}

func TestMinByteEdgeCases(t *testing.T) {
	t.Run("min byte with all maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{255, 255, 255, 255})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 255 {
			t.Errorf("Expected min 255, got %d", min)
		}
	})

	t.Run("min byte with all minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{0, 0, 0, 0})

		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})

	t.Run("min byte early termination optimization", func(t *testing.T) {
		t.Parallel()
		bytes := make([]byte, 1000)
		bytes[0] = 0
		for i := 1; i < 1000; i++ {
			bytes[i] = byte(i%255 + 1)
		}

		enumerator := FromSlice(bytes)
		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})

	t.Run("min byte without early termination", func(t *testing.T) {
		t.Parallel()
		bytes := make([]byte, 1000)
		for i := 0; i < 999; i++ {
			bytes[i] = byte((i % 254) + 1)
		}
		bytes[999] = 0

		enumerator := FromSlice(bytes)
		min, ok := enumerator.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %d", min)
		}
	})
}

func TestMinByteWithOperations(t *testing.T) {
	t.Run("min byte after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 200, 25, 150, 10})
		filtered := enumerator.Where(func(b byte) bool { return b > 30 })

		min, ok := filtered.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 50 {
			t.Errorf("Expected min 50 (from filtered values > 30), got %d", min)
		}
	})

	t.Run("min byte after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 200, 25, 150, 10})
		taken := enumerator.Take(4)

		min, ok := taken.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min 25 (from first 4 elements), got %d", min)
		}
	})

	t.Run("min byte after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 5, 200, 25, 150, 10})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 10 {
			t.Errorf("Expected min 10 (after skipping 2 elements), got %d", min)
		}
	})
}

func TestMinByteCustomKeySelector(t *testing.T) {
	t.Run("min byte by string length", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "golang"})

		min, ok := enumerator.MinByte(func(s string) byte { return byte(len(s)) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2 {
			t.Errorf("Expected min 2, got %d", min)
		}
	})

	t.Run("min byte by array element", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Values [5]byte
		}

		data := []Data{
			{Values: [5]byte{100, 50, 200, 25, 150}},
			{Values: [5]byte{200, 100, 50, 75, 125}},
			{Values: [5]byte{25, 75, 125, 175, 225}},
		}
		enumerator := FromSliceAny(data)
		min, ok := enumerator.MinByte(func(d Data) byte { return d.Values[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min 25, got %d", min)
		}
	})
}

func TestMinByteNonComparable(t *testing.T) {
	t.Run("min byte from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name []string
			Age  byte
		}

		people := []Person{
			{Name: []string{"Alice"}, Age: 30},
			{Name: []string{"Bob"}, Age: 25},
			{Name: []string{"Charlie"}, Age: 35},
		}

		enumerator := FromSliceAny(people)
		min, ok := enumerator.MinByte(func(p Person) byte { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25 {
			t.Errorf("Expected min 25, got %d", min)
		}
	})

	t.Run("min byte with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Timeout  byte
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Timeout: 30},
			{Settings: map[string]interface{}{"b": 2}, Timeout: 10},
			{Settings: map[string]interface{}{"c": 3}, Timeout: 60},
		}

		enumerator := FromSliceAny(configs)
		min, ok := enumerator.MinByte(func(c Config) byte { return c.Timeout })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 10 {
			t.Errorf("Expected min 10, got %d", min)
		}
	})

	t.Run("min byte with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Priority byte
		}

		handlers := []Handler{
			{Callback: func() {}, Priority: 5},
			{Callback: func() {}, Priority: 1},
			{Callback: func() {}, Priority: 3},
		}

		enumerator := FromSliceAny(handlers)
		min, ok := enumerator.MinByte(func(h Handler) byte { return h.Priority })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1 {
			t.Errorf("Expected min 1, got %d", min)
		}
	})
}

func BenchmarkMinByte(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []byte{100, 50, 200, 25, 150}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinByte(selector.Byte)
			if !ok || result != 25 {
				b.Fatalf("Expected 25, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		items := make([]byte, 1000)
		items[0] = 0
		for i := 1; i < 1000; i++ {
			items[i] = byte(i%255 + 1)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinByte(selector.Byte)
			if !ok || result != 0 {
				b.Fatalf("Expected 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		items := make([]byte, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = byte((i % 254) + 1)
		}
		items[9999] = 0

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinByte(selector.Byte)
			if !ok || result != 0 {
				b.Fatalf("Expected 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]byte{})
			result, ok := enumerator.MinByte(selector.Byte)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %d, ok: %v", result, ok)
			}
		}
	})
}
