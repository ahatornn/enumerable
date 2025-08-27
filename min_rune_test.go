package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinRune(t *testing.T) {
	t.Run("min rune from latin letters", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'a', 'm', 'b', 'y'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})

	t.Run("min rune from mixed unicode", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'я', 'а', 'м', 'б', 'ю'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'а' {
			t.Errorf("Expected min 'а', got %c", min)
		}
	})

	t.Run("min rune from numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'9', '1', '5', '3', '7'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != '1' {
			t.Errorf("Expected min '1', got %c", min)
		}
	})

	t.Run("min rune single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'A'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'A' {
			t.Errorf("Expected min 'A', got %c", min)
		}
	})

	t.Run("min rune with zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 0, 'a', 'm'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %c", min)
		}
	})

	t.Run("min rune empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{})

		min, ok := enumerator.MinRune(selector.Rune)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for empty slice, got %c", min)
		}
	})

	t.Run("min rune nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[rune] = nil

		min, ok := enumerator.MinRune(selector.Rune)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil enumerator, got %c", min)
		}
	})

	t.Run("min rune with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'a', 'b', 'c'})

		min, ok := enumerator.MinRune(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil keySelector, got %c", min)
		}
	})

	t.Run("min rune with custom key selector", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(c Character) rune { return c.Code })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'α' {
			t.Errorf("Expected min 'α', got %c", min)
		}
	})

	t.Run("min rune with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'a', 'm', 'a', 'y'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})

	t.Run("min rune with special characters", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'!', '@', '#', '$', '%'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != '!' {
			t.Errorf("Expected min '!', got %c", min)
		}
	})
}

func TestMinRuneStruct(t *testing.T) {
	t.Run("min rune from struct field", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(s Symbol) rune { return s.Char })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})

	t.Run("min rune from struct with slice field", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(e TextElement) rune { return e.Char })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})
}

func TestMinRuneEdgeCases(t *testing.T) {
	t.Run("min rune with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'x', 'x', 'x', 'x'})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'x' {
			t.Errorf("Expected min 'x', got %c", min)
		}
	})

	t.Run("min rune with unicode ranges", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{0x1F600, 0x1F601, 0x1F602})

		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0x1F600 {
			t.Errorf("Expected min 0x1F600, got %U", min)
		}
	})

	t.Run("min rune early termination optimization", func(t *testing.T) {
		t.Parallel()
		runes := make([]rune, 1000)
		runes[0] = 0
		for i := 1; i < 1000; i++ {
			runes[i] = rune(i + 65)
		}

		enumerator := FromSlice(runes)
		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %c", min)
		}
	})

	t.Run("min rune without early termination", func(t *testing.T) {
		t.Parallel()
		runes := make([]rune, 1000)
		for i := 0; i < 999; i++ {
			runes[i] = rune(i + 65)
		}
		runes[999] = 0

		enumerator := FromSlice(runes)
		min, ok := enumerator.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0 {
			t.Errorf("Expected min 0, got %c", min)
		}
	})
}

func TestMinRuneWithOperations(t *testing.T) {
	t.Run("min rune after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'a', 'B', 'c', 'D', 'e', 'F'})
		filtered := enumerator.Where(func(r rune) bool { return r >= 'a' })

		min, ok := filtered.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a' (from lowercase letters), got %c", min)
		}
	})

	t.Run("min rune after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'c', 'm', 'b', 'y', 'a'})
		taken := enumerator.Take(4)

		min, ok := taken.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'b' {
			t.Errorf("Expected min 'a' (from first 4 elements), got %c", min)
		}
	})

	t.Run("min rune after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]rune{'z', 'a', 'm', 'b', 'y', 'c'})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinRune(selector.Rune)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'b' {
			t.Errorf("Expected min 'b' (after skipping 2 elements), got %c", min)
		}
	})
}

func TestMinRuneCustomKeySelector(t *testing.T) {
	t.Run("min rune by string first rune", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "monkey"})

		min, ok := enumerator.MinRune(func(s string) rune { return rune(s[0]) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})

	t.Run("min rune by array element", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(d Data) rune { return d.Chars[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'b' {
			t.Errorf("Expected min 'b', got %c", min)
		}
	})
}

func TestMinRuneNonComparable(t *testing.T) {
	t.Run("min rune from non-comparable struct slice", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(l Letter) rune { return l.Unicode })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'α' {
			t.Errorf("Expected min 'α', got %c", min)
		}
	})

	t.Run("min rune with map field struct", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(c Config) rune { return c.Code })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})

	t.Run("min rune with function field struct", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(h Handler) rune { return h.Symbol })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'a' {
			t.Errorf("Expected min 'a', got %c", min)
		}
	})

	t.Run("min rune with complex non-comparable struct", func(t *testing.T) {
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
		min, ok := enumerator.MinRune(func(c ComplexStruct) rune { return c.CharCode })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 'A' {
			t.Errorf("Expected min 'A', got %c", min)
		}
	})
}

func BenchmarkMinRune(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []rune{'z', 'a', 'm', 'b', 'y'}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinRune(selector.Rune)
			if !ok || result != 'a' {
				b.Fatalf("Expected 'a', got %c, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		items := make([]rune, 1000)
		items[0] = 0
		for i := 1; i < 1000; i++ {
			items[i] = rune(i + 65)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinRune(selector.Rune)
			if !ok || result != 0 {
				b.Fatalf("Expected 0, got %c, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		items := make([]rune, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = rune(i + 65)
		}
		items[9999] = 0

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinRune(selector.Rune)
			if !ok || result != 0 {
				b.Fatalf("Expected 0, got %c, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]rune{})
			result, ok := enumerator.MinRune(selector.Rune)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %c, ok: %v", result, ok)
			}
		}
	})
}
