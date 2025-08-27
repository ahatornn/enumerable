package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMaxBool(t *testing.T) {
	t.Run("max bool from mixed values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool from all true values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, true})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool from all false values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, false, false})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != false {
			t.Errorf("Expected max false, got %v", max)
		}
	})

	t.Run("max bool single true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool single false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != false {
			t.Errorf("Expected max false, got %v", max)
		}
	})

	t.Run("max bool empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{})

		max, ok := enumerator.MaxBool(selector.Bool)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if max != false {
			t.Errorf("Expected max false for empty slice, got %v", max)
		}
	})

	t.Run("max bool nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[bool] = nil

		max, ok := enumerator.MaxBool(selector.Bool)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if max != false {
			t.Errorf("Expected max false for nil enumerator, got %v", max)
		}
	})

	t.Run("max bool with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false})

		max, ok := enumerator.MaxBool(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if max != false {
			t.Errorf("Expected max false for nil keySelector, got %v", max)
		}
	})

	t.Run("max bool with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Feature struct {
			Name     string
			Enabled  bool
			Priority int
		}

		features := []Feature{
			{Name: "FeatureA", Enabled: true, Priority: 1},
			{Name: "FeatureB", Enabled: false, Priority: 2},
			{Name: "FeatureC", Enabled: true, Priority: 3},
		}

		enumerator := FromSlice(features)
		max, ok := enumerator.MaxBool(func(f Feature) bool { return f.Enabled })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})
}

func TestMaxBoolStruct(t *testing.T) {
	t.Run("max bool from struct field", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name     string
			IsActive bool
			Age      int
		}

		users := []User{
			{Name: "Alice", IsActive: true, Age: 30},
			{Name: "Bob", IsActive: false, Age: 25},
			{Name: "Charlie", IsActive: true, Age: 35},
		}

		enumerator := FromSlice(users)
		max, ok := enumerator.MaxBool(func(u User) bool { return u.IsActive })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings []string
			IsValid  bool
		}

		configs := []Config{
			{Settings: []string{"a", "b"}, IsValid: true},
			{Settings: []string{"c"}, IsValid: false},
			{Settings: []string{"d", "e", "f"}, IsValid: true},
		}

		enumerator := FromSliceAny(configs)
		max, ok := enumerator.MaxBool(func(c Config) bool { return c.IsValid })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})
}

func TestMaxBoolEdgeCases(t *testing.T) {
	t.Run("max bool with all same values true", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, true, true})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool with all same values false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, false, false, false})

		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != false {
			t.Errorf("Expected max false, got %v", max)
		}
	})

	t.Run("max bool early termination optimization", func(t *testing.T) {
		t.Parallel()

		bools := make([]bool, 1000)
		bools[0] = true
		for i := 1; i < 1000; i++ {
			bools[i] = false
		}

		enumerator := FromSlice(bools)
		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool without early termination", func(t *testing.T) {
		t.Parallel()
		bools := make([]bool, 1000)
		for i := 0; i < 999; i++ {
			bools[i] = false
		}
		bools[999] = true

		enumerator := FromSlice(bools)
		max, ok := enumerator.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})
}

func TestMaxBoolWithOperations(t *testing.T) {
	t.Run("max bool after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})
		filtered := enumerator.Where(func(b bool) bool { return !b }) // только false

		max, ok := filtered.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != false {
			t.Errorf("Expected max false (from filtered false values), got %v", max)
		}
	})

	t.Run("max bool after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, false, true, false, true, false})
		taken := enumerator.Take(3)

		max, ok := taken.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true (from first 3 elements), got %v", max)
		}
	})

	t.Run("max bool after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, false, true, false, true, false})
		skipped := enumerator.Skip(2)

		max, ok := skipped.MaxBool(selector.Bool)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true (after skipping 2 elements), got %v", max)
		}
	})
}

func TestMaxBoolCustomKeySelector(t *testing.T) {
	t.Run("max bool by condition", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		max, ok := enumerator.MaxBool(func(x int) bool { return x%2 == 0 })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool by string empty check", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "", "world", "", "go"})

		max, ok := enumerator.MaxBool(func(s string) bool { return s != "" })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})
}

func TestMaxBoolNonComparable(t *testing.T) {
	t.Run("max bool from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name    []string
			IsAdmin bool
		}

		people := []Person{
			{Name: []string{"Alice"}, IsAdmin: true},
			{Name: []string{"Bob"}, IsAdmin: false},
			{Name: []string{"Charlie"}, IsAdmin: true},
		}

		enumerator := FromSliceAny(people)
		max, ok := enumerator.MaxBool(func(p Person) bool { return p.IsAdmin })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Enabled  bool
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Enabled: true},
			{Settings: map[string]interface{}{"b": 2}, Enabled: false},
			{Settings: map[string]interface{}{"c": 3}, Enabled: true},
		}

		enumerator := FromSliceAny(configs)
		max, ok := enumerator.MaxBool(func(c Config) bool { return c.Enabled })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})

	t.Run("max bool with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Active   bool
		}

		handlers := []Handler{
			{Callback: func() {}, Active: true},
			{Callback: func() {}, Active: false},
			{Callback: func() {}, Active: true},
		}

		enumerator := FromSliceAny(handlers)
		max, ok := enumerator.MaxBool(func(h Handler) bool { return h.Active })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if max != true {
			t.Errorf("Expected max true, got %v", max)
		}
	})
}

func BenchmarkMaxBool(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []bool{false, true, false, true, false}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxBool(selector.Bool)
			if !ok || result != true {
				b.Fatalf("Expected true, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		// true в начале - должно быть быстро
		items := make([]bool, 1000)
		items[0] = true
		for i := 1; i < 1000; i++ {
			items[i] = false
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxBool(selector.Bool)
			if !ok || result != true {
				b.Fatalf("Expected true, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		// true в конце - должно пройти все элементы
		items := make([]bool, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = false
		}
		items[9999] = true

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MaxBool(selector.Bool)
			if !ok || result != true {
				b.Fatalf("Expected true, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]bool{})
			result, ok := enumerator.MaxBool(selector.Bool)
			if ok || result != false {
				b.Fatalf("Expected false and false, got %v, ok: %v", result, ok)
			}
		}
	})
}
