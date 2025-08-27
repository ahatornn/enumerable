package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMaxRune(t *testing.T) {
	t.Run("max rune from latin letters", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'a', 'm', 'b', 'y'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})

	t.Run("max rune from mixed unicode", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'я', 'а', 'м', 'б', 'ю'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'я' {
			t.Errorf("Expected max 'я', got %c", max)
		}
	})

	t.Run("max rune from numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'9', '1', '5', '3', '7'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != '9' {
			t.Errorf("Expected max '9', got %c", max)
		}
	})

	t.Run("max rune single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'A'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'A' {
			t.Errorf("Expected max 'A', got %c", max)
		}
	})

	t.Run("max rune with maximum unicode value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 0x10FFFF, 'a', 'm'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 0x10FFFF {
			t.Errorf("Expected max 0x10FFFF, got %U", max)
		}
	})

	t.Run("max rune empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{})

		max, ok := enumerator.MaxRune(selector.Rune)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for empty slice, got %c", max)
		}
	})

	t.Run("max rune nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[rune] = nil

		max, ok := enumerator.MaxRune(selector.Rune)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil enumerator, got %c", max)
		}
	})

	t.Run("max rune with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'a', 'b', 'c'})

		max, ok := enumerator.MaxRune(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if max != 0 {
			t.Errorf("Expected max 0 for nil keySelector, got %c", max)
		}
	})

	t.Run("max rune with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Character struct {
			Name string
			Code rune
		}

		chars := []Character{
			{Name: "Alpha", Code: 'α'},
			{Name: "Beta", Code: 'β'},
			{Name: "Gamma", Code: 'γ'},
		}

		enumerator := FromSlice(chars)
		max, ok := enumerator.MaxRune(func(c Character) rune { return c.Code })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'γ' {
			t.Errorf("Expected max 'γ', got %c", max)
		}
	})

	t.Run("max rune with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'a', 'm', 'z', 'y'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})

	t.Run("max rune with special characters", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'!', '@', '#', '$', '%'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != '@' {
			t.Errorf("Expected max '@', got %c", max)
		}
	})
}

func TestMaxRuneStruct(t *testing.T) {
	t.Run("max rune from struct field", func(t *testing.T) {
		t.Parallel()
		type Symbol struct {
			Name string
			Char rune
		}

		symbols := []Symbol{
			{Name: "First", Char: 'z'},
			{Name: "Second", Char: 'a'},
			{Name: "Third", Char: 'm'},
		}

		enumerator := FromSlice(symbols)
		max, ok := enumerator.MaxRune(func(s Symbol) rune { return s.Char })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})

	t.Run("max rune from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type TextElement struct {
			Tags []string
			Char rune
		}

		elements := []TextElement{
			{Tags: []string{"latin"}, Char: 'z'},
			{Tags: []string{"latin"}, Char: 'a'},
			{Tags: []string{"latin"}, Char: 'm'},
		}

		enumerator := FromSliceAny(elements)
		max, ok := enumerator.MaxRune(func(e TextElement) rune { return e.Char })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})
}

func TestMaxRuneEdgeCases(t *testing.T) {
	t.Run("max rune with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'x', 'x', 'x', 'x'})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'x' {
			t.Errorf("Expected max 'x', got %c", max)
		}
	})

	t.Run("max rune with unicode ranges", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{0x1F600, 0x1F601, 0x1F602})

		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 0x1F602 {
			t.Errorf("Expected max 0x1F602, got %U", max)
		}
	})

	t.Run("max rune early termination optimization", func(t *testing.T) {
		t.Parallel()
		runes := make([]rune, 1000)
		runes[0] = 0x10FFFF
		for i := 1; i < 1000; i++ {
			runes[i] = rune(i + 65)
		}

		enumerator := FromSlice(runes)
		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 0x10FFFF {
			t.Errorf("Expected max 0x10FFFF, got %U", max)
		}
	})

	t.Run("max rune without early termination", func(t *testing.T) {
		t.Parallel()
		runes := make([]rune, 1000)
		for i := 0; i < 999; i++ {
			runes[i] = rune(i + 65)
		}
		runes[999] = 0x10FFFF

		enumerator := FromSlice(runes)
		max, ok := enumerator.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 0x10FFFF {
			t.Errorf("Expected max 0x10FFFF, got %U", max)
		}
	})
}

func TestMaxRuneWithOperations(t *testing.T) {
	t.Run("max rune after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'a', 'B', 'c', 'D', 'e', 'F'})
		filtered := enumerator.Where(func(r rune) bool { return r >= 'a' })

		max, ok := filtered.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'e' {
			t.Errorf("Expected max 'e' (from lowercase letters), got %c", max)
		}
	})

	t.Run("max rune after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'y', 'a', 'm', 'b', 'y', 'z'})
		taken := enumerator.Take(4)

		max, ok := taken.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'y' {
			t.Errorf("Expected max 'z' (from first 4 elements), got %c", max)
		}
	})

	t.Run("max rune after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'a', 'm', 'b', 'y', 'c'})
		skipped := enumerator.Skip(2)

		max, ok := skipped.MaxRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'y' {
			t.Errorf("Expected max 'y' (after skipping 2 elements), got %c", max)
		}
	})
}

func TestMaxRuneCustomKeySelector(t *testing.T) {
	t.Run("max rune by string first rune", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey"})

		max, ok := enumerator.MaxRune(func(s string) rune { return rune(s[0]) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})

	t.Run("max rune by array element", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Chars [3]rune
		}

		data := []Data{
			{Chars: [3]rune{'z', 'a', 'm'}},
			{Chars: [3]rune{'b', 'y', 'c'}},
			{Chars: [3]rune{'d', 'x', 'e'}},
		}
		enumerator := FromSlice(data)
		max, ok := enumerator.MaxRune(func(d Data) rune { return d.Chars[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})
}

func TestMaxRuneNonComparable(t *testing.T) {
	t.Run("max rune from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Letter struct {
			Name    []string
			Unicode rune
		}

		letters := []Letter{
			{Name: []string{"Alpha"}, Unicode: 'α'},
			{Name: []string{"Beta"}, Unicode: 'β'},
			{Name: []string{"Gamma"}, Unicode: 'γ'},
		}

		enumerator := FromSliceAny(letters)
		max, ok := enumerator.MaxRune(func(l Letter) rune { return l.Unicode })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'γ' {
			t.Errorf("Expected max 'γ', got %c", max)
		}
	})

	t.Run("max rune with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Code     rune
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Code: 'z'},
			{Settings: map[string]interface{}{"b": 2}, Code: 'a'},
			{Settings: map[string]interface{}{"c": 3}, Code: 'm'},
		}

		enumerator := FromSliceAny(configs)
		max, ok := enumerator.MaxRune(func(c Config) rune { return c.Code })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})

	t.Run("max rune with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Symbol   rune
		}

		handlers := []Handler{
			{Callback: func() {}, Symbol: 'z'},
			{Callback: func() {}, Symbol: 'a'},
			{Callback: func() {}, Symbol: 'm'},
		}

		enumerator := FromSliceAny(handlers)
		max, ok := enumerator.MaxRune(func(h Handler) rune { return h.Symbol })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'z' {
			t.Errorf("Expected max 'z', got %c", max)
		}
	})

	t.Run("max rune with complex non-comparable struct", func(t *testing.T) {
		t.Parallel()
		type ComplexStruct struct {
			Data     map[string][]rune
			Callback func(int) bool
			Channel  chan string
			CharCode rune
		}

		complexData := []ComplexStruct{
			{
				Data:     map[string][]rune{"a": {'α', 'β'}},
				Callback: func(i int) bool { return i > 0 },
				Channel:  make(chan string, 1),
				CharCode: 'A',
			},
			{
				Data:     map[string][]rune{"b": {'γ', 'δ'}},
				Callback: func(i int) bool { return i < 0 },
				Channel:  make(chan string, 1),
				CharCode: 'B',
			},
			{
				Data:     map[string][]rune{"c": {'ε', 'ζ'}},
				Callback: func(i int) bool { return i == 0 },
				Channel:  make(chan string, 1),
				CharCode: 'C',
			},
		}

		enumerator := FromSliceAny(complexData)
		max, ok := enumerator.MaxRune(func(c ComplexStruct) rune { return c.CharCode })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != 'C' {
			t.Errorf("Expected max 'C', got %c", max)
		}
	})
}

func BenchmarkMaxRune(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []rune{'z', 'a', 'm', 'b', 'y'}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxRune(selector.Rune)
			if !ok || result != 'z' {
				b.Fatalf("Expected 'z', got %c, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		items := make([]rune, 1000)
		items[0] = 0x10FFFF
		for i := 1; i < 1000; i++ {
			items[i] = rune(i + 65)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxRune(selector.Rune)
			if !ok || result != 0x10FFFF {
				b.Fatalf("Expected 0x10FFFF, got %U, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		items := make([]rune, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = rune(i + 65)
		}
		items[9999] = 0x10FFFF

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxRune(selector.Rune)
			if !ok || result != 0x10FFFF {
				b.Fatalf("Expected 0x10FFFF, got %U, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]rune{})
			result, ok := enumerator.MaxRune(selector.Rune)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %c, ok: %v", result, ok)
			}
		}
	})
}
