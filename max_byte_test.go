package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMaxByte(t *testing.T) {
	t.Run("max byte from mixed values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 200, 25, 150})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200, got %d", max)
		}
	})

	t.Run("max byte from all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 100, 100})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 100 {
			t.Errorf("Expected max 100, got %d", max)
		}
	})

	t.Run("max byte with maximum value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 255, 25, 150})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 255 {
			t.Errorf("Expected max 255, got %d", max)
		}
	})

	t.Run("max byte single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{42})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 42 {
			t.Errorf("Expected max 42, got %d", max)
		}
	})

	t.Run("max byte minimum value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{0, 1, 2})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 2 {
			t.Errorf("Expected max 2, got %d", max)
		}
	})

	t.Run("max byte empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{})

		max, ok := enumerator.MaxByte(selector.Byte)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for empty slice, got %d", max)
		}
	})

	t.Run("max byte nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[byte] = nil

		max, ok := enumerator.MaxByte(selector.Byte)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil enumerator, got %d", max)
		}
	})

	t.Run("max byte with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 25})

		max, ok := enumerator.MaxByte(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil keySelector, got %d", max)
		}
	})

	t.Run("max byte with custom key selector", func(t *testing.T) {
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
		max, ok := enumerator.MaxByte(func(d Data) byte { return d.Size })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200, got %d", max)
		}
	})

	t.Run("max byte with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 200, 50, 200, 150})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200, got %d", max)
		}
	})
}

func TestMaxByteStruct(t *testing.T) {
	t.Run("max byte from struct field", func(t *testing.T) {
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
		max, ok := enumerator.MaxByte(func(f File) byte { return f.Size })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200, got %d", max)
		}
	})

	t.Run("max byte from struct with slice field", func(t *testing.T) {
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
		max, ok := enumerator.MaxByte(func(c Config) byte { return c.BufferSize })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200, got %d", max)
		}
	})
}

func TestMaxByteEdgeCases(t *testing.T) {
	t.Run("max byte with all maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{255, 255, 255, 255})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 255 {
			t.Errorf("Expected max 255, got %d", max)
		}
	})

	t.Run("max byte with all minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{0, 0, 0, 0})

		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 0 {
			t.Errorf("Expected max 0, got %d", max)
		}
	})

	t.Run("max byte early termination optimization", func(t *testing.T) {
		t.Parallel()
		bytes := make([]byte, 1000)
		bytes[0] = 255
		for i := 1; i < 1000; i++ {
			bytes[i] = byte(i % 254)
		}

		enumerator := FromSlice(bytes)
		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 255 {
			t.Errorf("Expected max 255, got %d", max)
		}
	})

	t.Run("max byte without early termination", func(t *testing.T) {
		t.Parallel()
		bytes := make([]byte, 1000)
		for i := 0; i < 999; i++ {
			bytes[i] = byte(i % 254)
		}
		bytes[999] = 255

		enumerator := FromSlice(bytes)
		max, ok := enumerator.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 255 {
			t.Errorf("Expected max 255, got %d", max)
		}
	})
}

func TestMaxByteWithOperations(t *testing.T) {
	t.Run("max byte after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 200, 25, 150, 10})
		filtered := enumerator.Where(func(b byte) bool { return b < 180 })

		max, ok := filtered.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 150 {
			t.Errorf("Expected max 150 (from filtered values < 180), got %d", max)
		}
	})

	t.Run("max byte after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 50, 200, 25, 150, 240})
		taken := enumerator.Take(4)

		max, ok := taken.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200 (from first 4 elements), got %d", max)
		}
	})

	t.Run("max byte after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]byte{100, 244, 200, 25, 150, 10})
		skipped := enumerator.Skip(2)

		max, ok := skipped.MaxByte(selector.Byte)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200 (after skipping 2 elements), got %d", max)
		}
	})
}

func TestMaxByteCustomKeySelector(t *testing.T) {
	t.Run("max byte by string length", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "golang"})

		max, ok := enumerator.MaxByte(func(s string) byte { return byte(len(s)) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 6 {
			t.Errorf("Expected max 6, got %d", max)
		}
	})

	t.Run("max byte by array element", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Values [5]byte
		}

		data := []Data{
			{Values: [5]byte{100, 50, 200, 25, 150}},
			{Values: [5]byte{200, 100, 50, 75, 125}},
			{Values: [5]byte{25, 75, 125, 175, 225}},
		}
		enumerator := FromSlice(data)
		max, ok := enumerator.MaxByte(func(d Data) byte { return d.Values[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 200 {
			t.Errorf("Expected max 200, got %d", max)
		}
	})
}

func TestMaxByteNonComparable(t *testing.T) {
	t.Run("max byte from non-comparable struct slice", func(t *testing.T) {
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
		max, ok := enumerator.MaxByte(func(p Person) byte { return p.Age })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 35 {
			t.Errorf("Expected max 35, got %d", max)
		}
	})

	t.Run("max byte with map field struct", func(t *testing.T) {
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
		max, ok := enumerator.MaxByte(func(c Config) byte { return c.Timeout })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 60 {
			t.Errorf("Expected max 60, got %d", max)
		}
	})

	t.Run("max byte with function field struct", func(t *testing.T) {
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
		max, ok := enumerator.MaxByte(func(h Handler) byte { return h.Priority })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 5 {
			t.Errorf("Expected max 5, got %d", max)
		}
	})
}

func BenchmarkMaxByte(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []byte{100, 50, 200, 25, 150}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxByte(selector.Byte)
			if !ok || result != 200 {
				b.Fatalf("Expected 200, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		// 255 в начале - должно быть быстро
		items := make([]byte, 1000)
		items[0] = 255
		for i := 1; i < 1000; i++ {
			items[i] = byte(i % 254)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxByte(selector.Byte)
			if !ok || result != 255 {
				b.Fatalf("Expected 255, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		// 255 в конце - должно пройти все элементы
		items := make([]byte, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = byte(i % 254)
		}
		items[9999] = 255

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxByte(selector.Byte)
			if !ok || result != 255 {
				b.Fatalf("Expected 255, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]byte{})
			result, ok := enumerator.MaxByte(selector.Byte)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %d, ok: %v", result, ok)
			}
		}
	})
}
