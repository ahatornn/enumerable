package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinString(t *testing.T) {
	t.Run("min string from words", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "banana", "cherry"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple', got '%s'", min)
		}
	})

	t.Run("min string from mixed case", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"Zebra", "apple", "Banana", "cherry"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "Banana" {
			t.Errorf("Expected min 'Banana', got '%s'", min)
		}
	})

	t.Run("min string with empty string", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "", "apple", "monkey"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "" {
			t.Errorf("Expected min '', got '%s'", min)
		}
	})

	t.Run("min string single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "hello" {
			t.Errorf("Expected min 'hello', got '%s'", min)
		}
	})

	t.Run("min string with numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"999", "1", "55", "3", "777"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "1" {
			t.Errorf("Expected min '1', got '%s'", min)
		}
	})

	t.Run("min string empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		min, ok := enumerator.MinString(selector.String)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != "" {
			t.Errorf("Expected min '' for empty slice, got '%s'", min)
		}
	})

	t.Run("min string nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[string] = nil

		min, ok := enumerator.MinString(selector.String)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != "" {
			t.Errorf("Expected min '' for nil enumerator, got '%s'", min)
		}
	})

	t.Run("min string with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c"})

		min, ok := enumerator.MinString(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != "" {
			t.Errorf("Expected min '' for nil keySelector, got '%s'", min)
		}
	})

	t.Run("min string with custom key selector", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(p Person) string { return p.ID })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "A002" {
			t.Errorf("Expected min 'A002', got '%s'", min)
		}
	})

	t.Run("min string with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "apple", "cherry"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple', got '%s'", min)
		}
	})

	t.Run("min string with special characters", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"!", "@", "#", "$", "%"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "!" {
			t.Errorf("Expected min '!', got '%s'", min)
		}
	})
}

func TestMinStringStruct(t *testing.T) {
	t.Run("min string from struct field", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(f File) string { return f.Path })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "/apple" {
			t.Errorf("Expected min '/apple', got '%s'", min)
		}
	})

	t.Run("min string from struct with slice field", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(d Document) string { return d.Title })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "Apple Sports" {
			t.Errorf("Expected min 'Apple Sports', got '%s'", min)
		}
	})
}

func TestMinStringEdgeCases(t *testing.T) {
	t.Run("min string with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hello", "hello", "hello"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "hello" {
			t.Errorf("Expected min 'hello', got '%s'", min)
		}
	})

	t.Run("min string with unicode", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"яблоко", "апельсин", "мандарин", "банан"})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "апельсин" {
			t.Errorf("Expected min 'апельсин', got '%s'", min)
		}
	})

	t.Run("min string early termination optimization", func(t *testing.T) {
		t.Parallel()
		strings := make([]string, 1000)
		strings[0] = ""
		for i := 1; i < 1000; i++ {
			strings[i] = string(rune(i%26 + 65))
		}

		enumerator := FromSlice(strings)
		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "" {
			t.Errorf("Expected min '', got '%s'", min)
		}
	})

	t.Run("min string without early termination", func(t *testing.T) {
		t.Parallel()
		strings := make([]string, 1000)
		for i := 0; i < 999; i++ {
			strings[i] = string(rune(i%26 + 65))
		}
		strings[999] = ""

		enumerator := FromSlice(strings)
		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "" {
			t.Errorf("Expected min '', got '%s'", min)
		}
	})

	t.Run("min string with very long strings", func(t *testing.T) {
		t.Parallel()
		long1 := string(make([]byte, 100, 10000))
		long2 := string(make([]byte, 101, 10001))
		long3 := string(make([]byte, 99, 9999))

		enumerator := FromSlice([]string{long1, long2, long3})

		min, ok := enumerator.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != long3 {
			t.Errorf("Expected min long3, got different string")
		}
	})
}

func TestMinStringWithOperations(t *testing.T) {
	t.Run("min string after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "Banana", "cherry", "Date", "elderberry", "Fig"})
		filtered := enumerator.Where(func(s string) bool { return s[0] >= 'a' })

		min, ok := filtered.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple' (from lowercase strings), got '%s'", min)
		}
	})

	t.Run("min string after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "banana", "cherry", "abc"})
		taken := enumerator.Take(4)

		min, ok := taken.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple' (from first 4 elements), got '%s'", min)
		}
	})

	t.Run("min string after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey", "banana", "cherry", "date"})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinString(selector.String)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "banana" {
			t.Errorf("Expected min 'banana' (after skipping 2 elements), got '%s'", min)
		}
	})
}

func TestMinStringCustomKeySelector(t *testing.T) {
	t.Run("min string by first character", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey"})

		min, ok := enumerator.MinString(func(s string) string { return string(s[0]) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "a" {
			t.Errorf("Expected min 'a', got '%s'", min)
		}
	})

	t.Run("min string by array element", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(d Data) string { return d.Words[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "banana" {
			t.Errorf("Expected min 'banana', got '%s'", min)
		}
	})
}

func TestMinStringNonComparable(t *testing.T) {
	t.Run("min string from non-comparable struct slice", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(d Document) string { return d.Content })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "Apple content" {
			t.Errorf("Expected min 'Apple content', got '%s'", min)
		}
	})

	t.Run("min string with map field struct", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(c Config) string { return c.Name })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple', got '%s'", min)
		}
	})

	t.Run("min string with function field struct", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(h Handler) string { return h.Pattern })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "apple" {
			t.Errorf("Expected min 'apple', got '%s'", min)
		}
	})

	t.Run("min string with complex non-comparable struct", func(t *testing.T) {
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
		min, ok := enumerator.MinString(func(c ComplexStruct) string { return c.Label })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != "Apple" {
			t.Errorf("Expected min 'Apple', got '%s'", min)
		}
	})
}

func BenchmarkMinString(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []string{"zebra", "apple", "monkey", "banana", "cherry"}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinString(selector.String)
			if !ok || result != "apple" {
				b.Fatalf("Expected 'apple', got '%s', ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		items := make([]string, 1000)
		items[0] = ""
		for i := 1; i < 1000; i++ {
			items[i] = string(rune(i%26 + 65))
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinString(selector.String)
			if !ok || result != "" {
				b.Fatalf("Expected '', got '%s', ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		items := make([]string, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = string(rune(i%26 + 65))
		}
		items[9999] = ""

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinString(selector.String)
			if !ok || result != "" {
				b.Fatalf("Expected '', got '%s', ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]string{})
			result, ok := enumerator.MinString(selector.String)
			if ok || result != "" {
				b.Fatalf("Expected '' and false, got '%s', ok: %v", result, ok)
			}
		}
	})
}
