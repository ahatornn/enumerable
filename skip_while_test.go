package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestSkipWhile(t *testing.T) {
	t.Run("basic skip while", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 1, 2})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		expected := []int{3, 4, 5, 1, 2}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("basic skip while for non-comparable", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
			Age  int
		}

		enumerator := FromSliceAny([]Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
			{Name: "David", Age: 40},
			{Name: "Eve", Age: 45},
			{Name: "Frank", Age: 25},
			{Name: "Grace", Age: 30},
		})

		skipped := enumerator.SkipWhile(func(p Person) bool {
			return p.Age < 35
		})

		expected := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "David", Age: 40},
			{Name: "Eve", Age: 45},
			{Name: "Frank", Age: 25},
			{Name: "Grace", Age: 30},
		}
		actual := []Person{}

		skipped(func(item Person) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i].Name != v.Name || actual[i].Age != v.Age {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, actual[i])
			}
		}
	})

	t.Run("skip while even numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 7, 8, 10, 11, 12})

		skipped := enumerator.SkipWhile(func(n int) bool { return n%2 == 0 })

		expected := []int{7, 8, 10, 11, 12}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("predicate never matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 4, 5, 6})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		expected := []int{3, 4, 5, 6}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("predicate always matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 10 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when predicate always matches, got %d", count)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty slice, got %d", count)
		}
	})

	t.Run("nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})

	t.Run("nil enumerator any", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator any, got %d", count)
		}
	})
}

func TestSkipWhileString(t *testing.T) {
	t.Run("skip while strings start with letter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "banana", "123start", "cherry", "456end"})

		skipped := enumerator.SkipWhile(func(s string) bool {
			return len(s) > 0 && s[0] >= 'a' && s[0] <= 'z'
		})

		expected := []string{"123start", "cherry", "456end"}
		actual := []string{}

		skipped(func(item string) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
			}
		}
	})
}

func TestSkipWhileEarlyTermination(t *testing.T) {
	t.Run("early termination after skip while", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 5 })

		actual := []int{}
		skipped(func(item int) bool {
			if len(actual) >= 3 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{5, 6, 7}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})
}

func TestSkipWhileStruct(t *testing.T) {
	t.Run("skip while struct field condition", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 20},
			{Name: "Eve", Age: 40},
		}

		enumerator := FromSlice(people)
		skipped := enumerator.SkipWhile(func(p Person) bool { return p.Age < 35 })

		expected := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 20},
			{Name: "Eve", Age: 40},
		}

		actual := []Person{}
		skipped(func(item Person) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, actual[i])
			}
		}
	})
}

func TestSkipWhileEdgeCases(t *testing.T) {
	t.Run("single element matches predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when single element matches, got %d", count)
		}
	})

	t.Run("single element doesn't match predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5})

		skipped := enumerator.SkipWhile(func(n int) bool { return n < 3 })

		expected := []int{5}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("predicate returns false then true", func(t *testing.T) {
		t.Parallel()
		// Important: once we stop skipping, we yield everything after
		enumerator := FromSlice([]int{2, 1, 3, 1, 4}) // 2 matches, 1 doesn't, so yield 1 and all after

		skipped := enumerator.SkipWhile(func(n int) bool { return n > 1 })

		expected := []int{1, 3, 1, 4}
		actual := []int{}

		skipped(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})
}

func TestSkipWhileBoolean(t *testing.T) {
	t.Run("skip while true values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, false, true, false})

		skipped := enumerator.SkipWhile(func(b bool) bool { return b })

		expected := []bool{false, true, false}
		actual := []bool{}

		skipped(func(item bool) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %t at index %d, got %t", v, i, actual[i])
			}
		}
	})
}

func TestOrderEnumeratorSkipWhile(t *testing.T) {
	t.Run("order enumerator skip while with sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		skipped := ordered.SkipWhile(func(n int) bool { return n < 4 })

		expected := []int{5, 8, 9}
		actual := skipped.ToSlice()

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator any skip while with complex struct", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Charlie", Age: 30},
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 35},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 22},
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })
		skipped := ordered.SkipWhile(func(p Person) bool { return p.Age < 28 })

		actual := skipped.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		expectedAges := []int{28, 30, 35}
		for i, expectedAge := range expectedAges {
			if actual[i].Age != expectedAge {
				t.Errorf("Expected age %d at index %d, got %d", expectedAge, i, actual[i].Age)
			}
		}

		expectedNames := []string{"Diana", "Charlie", "Bob"}
		for i, expectedName := range expectedNames {
			if actual[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, actual[i].Name)
			}
		}
	})

	t.Run("order enumerator skip while with multiple sorting levels", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Value    int
			Name     string
		}

		records := []Record{
			{Category: "B", Value: 10, Name: "Second"},
			{Category: "A", Value: 20, Name: "First"},
			{Category: "B", Value: 30, Name: "Fourth"},
			{Category: "A", Value: 15, Name: "Third"},
			{Category: "A", Value: 25, Name: "Fifth"},
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return a.Value - b.Value
		})
		skipped := ordered.SkipWhile(func(r Record) bool { return r.Category == "A" && r.Value <= 20 })

		actual := skipped.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		if actual[0].Category != "A" || actual[0].Value != 25 || actual[0].Name != "Fifth" {
			t.Errorf("Expected {A,25,Fifth}, got %+v", actual[0])
		}

		if actual[1].Category != "B" || actual[1].Value != 10 || actual[1].Name != "Second" {
			t.Errorf("Expected {B,10,Second}, got %+v", actual[1])
		}

		if actual[2].Category != "B" || actual[2].Value != 30 || actual[2].Name != "Fourth" {
			t.Errorf("Expected {B,30,Fourth}, got %+v", actual[2])
		}
	})

	t.Run("order enumerator skip while with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2, 4, 4})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		skipped := ordered.SkipWhile(func(n int) bool { return n < 3 })

		actual := skipped.ToSlice()
		expected := []int{3, 3, 4, 4}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator skip while predicate never matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		skipped := ordered.SkipWhile(func(n int) bool { return n < 0 })

		actual := skipped.ToSlice()
		expected := []int{1, 2, 3, 5, 8, 9}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator skip while predicate always matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		skipped := ordered.SkipWhile(func(n int) bool { return n < 20 })

		actual := skipped.ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice when predicate always matches, got length %d", len(actual))
		}
	})

	t.Run("order enumerator any skip while with complex struct and custom sorting", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name     string
			Priority int
			Options  []string
		}

		configs := []Config{
			{Name: "High", Priority: 1, Options: []string{"opt1", "opt2"}},
			{Name: "Low", Priority: 3, Options: []string{"opt3"}},
			{Name: "Medium", Priority: 2, Options: []string{"opt4", "opt5"}},
			{Name: "VeryLow", Priority: 4, Options: []string{"opt6"}},
			{Name: "Critical", Priority: 0, Options: []string{"opt7", "opt8", "opt9"}},
		}
		var enumerator EnumeratorAny[Config] = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
		skipped := ordered.SkipWhile(func(c Config) bool { return c.Priority < 2 })

		actual := skipped.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		expectedPriorities := []int{2, 3, 4}
		for i, expectedPriority := range expectedPriorities {
			if actual[i].Priority != expectedPriority {
				t.Errorf("Expected priority %d at index %d, got %d", expectedPriority, i, actual[i].Priority)
			}
		}

		expectedNames := []string{"Medium", "Low", "VeryLow"}
		for i, expectedName := range expectedNames {
			if actual[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, actual[i].Name)
			}
		}
	})

	t.Run("order enumerator skip while preserves stability", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Index int
		}

		items := []Item{
			{Value: 2, Index: 1},
			{Value: 1, Index: 2},
			{Value: 2, Index: 3},
			{Value: 1, Index: 4},
			{Value: 3, Index: 5},
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })
		skipped := ordered.SkipWhile(func(i Item) bool { return i.Value < 2 })

		actual := skipped.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		if actual[0].Value != 2 || actual[0].Index != 1 {
			t.Errorf("Expected {2,1} at index 0, got %+v", actual[0])
		}
		if actual[1].Value != 2 || actual[1].Index != 3 {
			t.Errorf("Expected {2,3} at index 1, got %+v", actual[1])
		}

		// Check last element
		if actual[2].Value != 3 || actual[2].Index != 5 {
			t.Errorf("Expected {3,5} at index 2, got %+v", actual[2])
		}
	})
}

func BenchmarkSkipWhile(b *testing.B) {
	b.Run("small skip while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipWhile(func(n int) bool { return n < 500 })
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("no skip while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i + 1000 // All items >= 1000
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipWhile(func(n int) bool { return n < 500 })
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
