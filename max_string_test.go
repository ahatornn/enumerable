package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMaxString(t *testing.T) {
	t.Run("max string from words", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "banana", "cherry"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})

	t.Run("max string from mixed case", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"Zebra", "apple", "Banana", "cherry"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "cherry" {
			t.Errorf("Expected max 'cherry', got '%s'", max)
		}
	})

	t.Run("max string with empty string", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "", "apple", "monkey"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})

	t.Run("max string single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "hello" {
			t.Errorf("Expected max 'hello', got '%s'", max)
		}
	})

	t.Run("max string with numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"999", "1", "55", "3", "777"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "999" {
			t.Errorf("Expected max '999', got '%s'", max)
		}
	})

	t.Run("max string empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		max, ok := enumerator.MaxString(selector.String)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if max != "" {
			t.Errorf("Expected max '' for empty slice, got '%s'", max)
		}
	})

	t.Run("max string nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[string] = nil

		max, ok := enumerator.MaxString(selector.String)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if max != "" {
			t.Errorf("Expected max '' for nil enumerator, got '%s'", max)
		}
	})

	t.Run("max string with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c"})

		max, ok := enumerator.MaxString(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if max != "" {
			t.Errorf("Expected max '' for nil keySelector, got '%s'", max)
		}
	})

	t.Run("max string with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			ID   string
		}

		people := []Person{
			{Name: "Alice", ID: "Z001"},
			{Name: "Bob", ID: "A002"},
			{Name: "Charlie", ID: "M003"},
		}

		enumerator := FromSlice(people)
		max, ok := enumerator.MaxString(func(p Person) string { return p.ID })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "Z001" {
			t.Errorf("Expected max 'Z001', got '%s'", max)
		}
	})

	t.Run("max string with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "zebra", "cherry"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})

	t.Run("max string with special characters", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"!", "@", "#", "$", "%"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "@" {
			t.Errorf("Expected max '@', got '%s'", max)
		}
	})
}

func TestMaxStringStruct(t *testing.T) {
	t.Run("max string from struct field", func(t *testing.T) {
		t.Parallel()
		type File struct {
			Name string
			Path string
		}

		files := []File{
			{Name: "First", Path: "/zebra"},
			{Name: "Second", Path: "/apple"},
			{Name: "Third", Path: "/monkey"},
		}

		enumerator := FromSlice(files)
		max, ok := enumerator.MaxString(func(f File) string { return f.Path })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "/zebra" {
			t.Errorf("Expected max '/zebra', got '%s'", max)
		}
	})

	t.Run("max string from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type Document struct {
			Tags  []string
			Title string
		}

		documents := []Document{
			{Tags: []string{"news"}, Title: "Zebra News"},
			{Tags: []string{"sports"}, Title: "Apple Sports"},
			{Tags: []string{"tech"}, Title: "Monkey Tech"},
		}

		enumerator := FromSliceAny(documents)
		max, ok := enumerator.MaxString(func(d Document) string { return d.Title })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "Zebra News" {
			t.Errorf("Expected max 'Zebra News', got '%s'", max)
		}
	})
}

func TestMaxStringEdgeCases(t *testing.T) {
	t.Run("max string with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hello", "hello", "hello"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "hello" {
			t.Errorf("Expected max 'hello', got '%s'", max)
		}
	})

	t.Run("max string with unicode", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"яблоко", "апельсин", "мандарин", "банан"})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "яблоко" {
			t.Errorf("Expected max 'яблоко', got '%s'", max)
		}
	})

	t.Run("max string without early termination", func(t *testing.T) {
		t.Parallel()
		strings := make([]string, 1000)
		for i := 0; i < 999; i++ {
			strings[i] = string(rune(i%26 + 97))
		}
		strings[999] = "zzz"

		enumerator := FromSlice(strings)
		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zzz" {
			t.Errorf("Expected max 'zzz', got '%s'", max)
		}
	})

	t.Run("max string with very long strings", func(t *testing.T) {
		t.Parallel()
		long1 := string(make([]byte, 100, 10000))
		long2 := string(make([]byte, 101, 10001))
		long3 := string(make([]byte, 99, 9999))

		enumerator := FromSlice([]string{long1, long2, long3})

		max, ok := enumerator.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != long2 {
			t.Errorf("Expected max long2, got different string")
		}
	})
}

func TestMaxStringWithOperations(t *testing.T) {
	t.Run("max string after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "Banana", "cherry", "Date", "elderberry", "Fig"})
		filtered := enumerator.Where(func(s string) bool { return s[0] >= 'a' })

		max, ok := filtered.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "elderberry" {
			t.Errorf("Expected max 'elderberry' (from lowercase strings), got '%s'", max)
		}
	})

	t.Run("max string after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "banana", "cherry", "date"})
		taken := enumerator.Take(4)

		max, ok := taken.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra' (from first 4 elements), got '%s'", max)
		}
	})

	t.Run("max string after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "banana", "cherry", "date"})
		skipped := enumerator.Skip(2)

		max, ok := skipped.MaxString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "monkey" {
			t.Errorf("Expected max 'monkey' (after skipping 2 elements), got '%s'", max)
		}
	})
}

func TestMaxStringCustomKeySelector(t *testing.T) {
	t.Run("max string by first character", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey"})

		max, ok := enumerator.MaxString(func(s string) string { return string(s[0]) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "z" {
			t.Errorf("Expected max 'z', got '%s'", max)
		}
	})

	t.Run("max string by array element", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Words [3]string
		}

		data := []Data{
			{Words: [3]string{"zebra", "apple", "monkey"}},
			{Words: [3]string{"banana", "cherry", "date"}},
			{Words: [3]string{"elderberry", "fig", "grape"}},
		}
		enumerator := FromSlice(data)
		max, ok := enumerator.MaxString(func(d Data) string { return d.Words[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})
}

func TestMaxStringNonComparable(t *testing.T) {
	t.Run("max string from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Document struct {
			Name    []string
			Content string
		}

		documents := []Document{
			{Name: []string{"Doc1"}, Content: "Zebra content"},
			{Name: []string{"Doc2"}, Content: "Apple content"},
			{Name: []string{"Doc3"}, Content: "Monkey content"},
		}

		enumerator := FromSliceAny(documents)
		max, ok := enumerator.MaxString(func(d Document) string { return d.Content })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "Zebra content" {
			t.Errorf("Expected max 'Zebra content', got '%s'", max)
		}
	})

	t.Run("max string with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Name     string
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Name: "zebra"},
			{Settings: map[string]interface{}{"b": 2}, Name: "apple"},
			{Settings: map[string]interface{}{"c": 3}, Name: "monkey"},
		}

		enumerator := FromSliceAny(configs)
		max, ok := enumerator.MaxString(func(c Config) string { return c.Name })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})

	t.Run("max string with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Pattern  string
		}

		handlers := []Handler{
			{Callback: func() {}, Pattern: "zebra"},
			{Callback: func() {}, Pattern: "apple"},
			{Callback: func() {}, Pattern: "monkey"},
		}

		enumerator := FromSliceAny(handlers)
		max, ok := enumerator.MaxString(func(h Handler) string { return h.Pattern })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "zebra" {
			t.Errorf("Expected max 'zebra', got '%s'", max)
		}
	})

	t.Run("max string with complex non-comparable struct", func(t *testing.T) {
		t.Parallel()
		type ComplexStruct struct {
			Data     map[string][]string
			Callback func(int) bool
			Channel  chan string
			Label    string
		}

		complexData := []ComplexStruct{
			{
				Data:     map[string][]string{"a": {"alpha", "beta"}},
				Callback: func(i int) bool { return i > 0 },
				Channel:  make(chan string, 1),
				Label:    "Zebra",
			},
			{
				Data:     map[string][]string{"b": {"gamma", "delta"}},
				Callback: func(i int) bool { return i < 0 },
				Channel:  make(chan string, 1),
				Label:    "Apple",
			},
			{
				Data:     map[string][]string{"c": {"epsilon", "zeta"}},
				Callback: func(i int) bool { return i == 0 },
				Channel:  make(chan string, 1),
				Label:    "Monkey",
			},
		}

		enumerator := FromSliceAny(complexData)
		max, ok := enumerator.MaxString(func(c ComplexStruct) string { return c.Label })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != "Zebra" {
			t.Errorf("Expected max 'Zebra', got '%s'", max)
		}
	})
}

func BenchmarkMaxString(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []string{"zebra", "apple", "monkey", "banana", "cherry"}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxString(selector.String)
			if !ok || result != "zebra" {
				b.Fatalf("Expected 'zebra', got '%s', ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]string, 1000)
		for i := 0; i < 999; i++ {
			items[i] = string(rune(i%26 + 97))
		}
		items[999] = "zzz"

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxString(selector.String)
			if !ok || result != "zzz" {
				b.Fatalf("Expected 'zzz', got '%s', ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]string, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = string(rune(i%26 + 97))
		}
		items[9999] = "zzz"

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxString(selector.String)
			if !ok || result != "zzz" {
				b.Fatalf("Expected 'zzz', got '%s', ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]string{})
			result, ok := enumerator.MaxString(selector.String)
			if ok || result != "" {
				b.Fatalf("Expected '' and false, got '%s', ok: %v", result, ok)
			}
		}
	})
}
