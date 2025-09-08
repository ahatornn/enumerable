package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestTakeWhile(t *testing.T) {
	t.Run("basic take while", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 1, 2})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 4 })

		expected := []int{1, 2, 3}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("basic take while for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{2, 3},
			{3, 4, 5},
			{4, 5},
			{5, 6},
			{1, 2},
			{2, 3},
		})

		taken := enumerator.TakeWhile(func(slice []int) bool {
			return len(slice) < 3
		})

		expected := [][]int{
			{1, 2},
			{2, 3},
		}
		actual := [][]int{}

		taken(func(item []int) bool {
			copy := make([]int, len(item))
			for i, v := range item {
				copy[i] = v
			}
			actual = append(actual, copy)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if len(actual[i]) != len(v) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(v), i, len(actual[i]))
				continue
			}
			for j, val := range v {
				if actual[i][j] != val {
					t.Errorf("Expected %d at index [%d][%d], got %d", val, i, j, actual[i][j])
				}
			}
		}
	})

	t.Run("take while even numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 7, 8, 10, 11, 12})

		taken := enumerator.TakeWhile(func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6}
		actual := []int{}

		taken(func(item int) bool {
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

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when predicate never matches, got %d", count)
		}
	})

	t.Run("predicate always matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 10 })

		expected := []int{1, 2, 3, 4}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
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

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
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

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator any, got %d", count)
		}
	})
}

func TestTakeWhileString(t *testing.T) {
	t.Run("take while strings start with letter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"apple", "banana", "123start", "cherry", "456end"})

		taken := enumerator.TakeWhile(func(s string) bool {
			return len(s) > 0 && s[0] >= 'a' && s[0] <= 'z'
		})

		expected := []string{"apple", "banana"}
		actual := []string{}

		taken(func(item string) bool {
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

func TestTakeWhileEarlyTermination(t *testing.T) {
	t.Run("early termination by consumer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 8 })

		actual := []int{}
		taken(func(item int) bool {
			if len(actual) >= 3 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 2, 3}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("early termination by predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 4, 6, 7, 8, 10})

		taken := enumerator.TakeWhile(func(n int) bool { return n%2 == 0 })

		actual := []int{}
		taken(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		expected := []int{2, 4, 6}
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

func TestTakeWhileStruct(t *testing.T) {
	t.Run("take while struct field condition", func(t *testing.T) {
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
		taken := enumerator.TakeWhile(func(p Person) bool { return p.Age < 35 })

		expected := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		actual := []Person{}
		taken(func(item Person) bool {
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

func TestTakeWhileEdgeCases(t *testing.T) {
	t.Run("single element matches predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		expected := []int{2}
		actual := []int{}

		taken(func(item int) bool {
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

	t.Run("single element doesn't match predicate", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5})

		taken := enumerator.TakeWhile(func(n int) bool { return n < 3 })

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when single element doesn't match, got %d", count)
		}
	})

	t.Run("predicate returns true then false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{2, 1, 3, 1, 4})

		taken := enumerator.TakeWhile(func(n int) bool { return n > 1 })

		expected := []int{2}
		actual := []int{}

		taken(func(item int) bool {
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

func TestTakeWhileBoolean(t *testing.T) {
	t.Run("take while true values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, true, false, true, false})

		taken := enumerator.TakeWhile(func(b bool) bool { return b })

		expected := []bool{true, true}
		actual := []bool{}

		taken(func(item bool) bool {
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

func TestOrderEnumeratorTakeWhile(t *testing.T) {
	t.Run("order enumerator take while with sorting", func(t *testing.T) {
		t.Parallel()
		actual := FromSlice([]int{5, 2, 8, 1, 9, 3}).
			OrderBy(comparer.ComparerInt).
			TakeWhile(func(n int) bool { return n < 5 }).
			ToSlice()
		expected := []int{1, 2, 3}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator any take while with complex struct", func(t *testing.T) {
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
		taken := ordered.TakeWhile(func(p Person) bool { return p.Age < 30 })

		actual := taken.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		expectedAges := []int{22, 25, 28}
		for i, expectedAge := range expectedAges {
			if actual[i].Age != expectedAge {
				t.Errorf("Expected age %d at index %d, got %d", expectedAge, i, actual[i].Age)
			}
		}

		expectedNames := []string{"Eve", "Alice", "Diana"}
		for i, expectedName := range expectedNames {
			if actual[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, actual[i].Name)
			}
		}
	})

	t.Run("order enumerator take while with multiple sorting levels", func(t *testing.T) {
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
		taken := ordered.TakeWhile(func(r Record) bool { return r.Category == "A" })

		actual := taken.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		if actual[0].Category != "A" || actual[0].Value != 15 || actual[0].Name != "Third" {
			t.Errorf("Expected {A,15,Third}, got %+v", actual[0])
		}

		if actual[1].Category != "A" || actual[1].Value != 20 || actual[1].Name != "First" {
			t.Errorf("Expected {A,20,First}, got %+v", actual[1])
		}

		if actual[2].Category != "A" || actual[2].Value != 25 || actual[2].Name != "Fifth" {
			t.Errorf("Expected {A,25,Fifth}, got %+v", actual[2])
		}
	})

	t.Run("order enumerator take while predicate immediately false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeWhile(func(n int) bool { return n < 0 })

		actual := taken.ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice when predicate immediately false, got length %d", len(actual))
		}
	})

	t.Run("order enumerator take while predicate never false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeWhile(func(n int) bool { return n < 20 })

		actual := taken.ToSlice()
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

	t.Run("order enumerator any take while with complex struct and custom sorting", func(t *testing.T) {
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
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
		taken := ordered.TakeWhile(func(c Config) bool { return c.Priority < 3 })

		actual := taken.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		expectedPriorities := []int{0, 1, 2}
		for i, expectedPriority := range expectedPriorities {
			if actual[i].Priority != expectedPriority {
				t.Errorf("Expected priority %d at index %d, got %d", expectedPriority, i, actual[i].Priority)
			}
		}

		expectedNames := []string{"Critical", "High", "Medium"}
		for i, expectedName := range expectedNames {
			if actual[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, actual[i].Name)
			}
		}
	})

	t.Run("order enumerator take while with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2, 4, 4})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeWhile(func(n int) bool { return n < 3 })

		actual := taken.ToSlice()
		expected := []int{1, 1, 2, 2}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator take while preserves stability", func(t *testing.T) {
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
		taken := ordered.TakeWhile(func(i Item) bool { return i.Value < 3 })

		actual := taken.ToSlice()

		if len(actual) != 4 {
			t.Fatalf("Expected length 4, got %d", len(actual))
		}

		if actual[0].Value != 1 || actual[0].Index != 2 {
			t.Errorf("Expected {1,2} at index 0, got %+v", actual[0])
		}
		if actual[1].Value != 1 || actual[1].Index != 4 {
			t.Errorf("Expected {1,4} at index 1, got %+v", actual[1])
		}

		if actual[2].Value != 2 || actual[2].Index != 1 {
			t.Errorf("Expected {2,1} at index 2, got %+v", actual[2])
		}
		if actual[3].Value != 2 || actual[3].Index != 3 {
			t.Errorf("Expected {2,3} at index 3, got %+v", actual[3])
		}
	})

	t.Run("order enumerator take while with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		taken := ordered.TakeWhile(func(n int) bool { return n > 3 })

		actual := taken.ToSlice()
		expected := []int{9, 8, 5}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator take while stops at first false", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 1, 2})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeWhile(func(n int) bool { return n < 4 })

		actual := taken.ToSlice()
		expected := []int{1, 1, 2, 2, 3}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator take while with complex predicate logic", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Value    int
			Name     string
		}

		records := []Record{
			{Category: "A", Value: 10, Name: "First"},
			{Category: "A", Value: 20, Name: "Second"},
			{Category: "B", Value: 15, Name: "Third"},
			{Category: "A", Value: 25, Name: "Fourth"},
			{Category: "B", Value: 30, Name: "Fifth"},
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return a.Value - b.Value
		})
		taken := ordered.TakeWhile(func(r Record) bool {
			return r.Category == "A" && r.Value < 25
		})

		actual := taken.ToSlice()

		if len(actual) != 2 {
			t.Fatalf("Expected length 2, got %d", len(actual))
		}

		if actual[0].Category != "A" || actual[0].Value != 10 || actual[0].Name != "First" {
			t.Errorf("Expected {A,10,First}, got %+v", actual[0])
		}

		if actual[1].Category != "A" || actual[1].Value != 20 || actual[1].Name != "Second" {
			t.Errorf("Expected {A,20,Second}, got %+v", actual[1])
		}
	})
}

func BenchmarkTakeWhile(b *testing.B) {
	b.Run("small take while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeWhile(func(n int) bool { return n < 500 })
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("no take while", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i + 1000
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeWhile(func(n int) bool { return n < 500 })
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
