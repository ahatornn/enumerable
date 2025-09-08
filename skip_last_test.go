package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestSkipLast(t *testing.T) {
	t.Run("basic skip last", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(2)

		expected := []int{1, 2, 3}
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

	t.Run("basic skip last for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		skipped := enumerator.SkipLast(2)

		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		actual := [][]int{}

		skipped(func(item []int) bool {
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

	t.Run("skip last zero elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(0)

		expected := []int{1, 2, 3, 4, 5}
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

	t.Run("skip more than available", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		skipped := enumerator.SkipLast(5)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when skipping more than available, got %d", count)
		}
	})

	t.Run("skip exactly all elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(5)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when skipping exactly all elements, got %d", count)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		skipped := enumerator.SkipLast(3)

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

		skipped := enumerator.SkipLast(3)

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

		skipped := enumerator.SkipLast(3)

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

func TestSkipLastString(t *testing.T) {
	t.Run("skip last strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c", "d", "e"})

		skipped := enumerator.SkipLast(2)

		expected := []string{"a", "b", "c"}
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

func TestSkipLastEarlyTermination(t *testing.T) {
	t.Run("early termination after skip last", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})

		skipped := enumerator.SkipLast(2)

		actual := []int{}
		skipped(func(item int) bool {
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
}

func TestSkipLastStruct(t *testing.T) {
	t.Run("skip last structs", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
		}

		enumerator := FromSlice(people)
		skipped := enumerator.SkipLast(2)

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
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

func TestSkipLastEdgeCases(t *testing.T) {
	t.Run("single element skip zero", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		skipped := enumerator.SkipLast(0)

		expected := []int{42}
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

	t.Run("single element skip one", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		skipped := enumerator.SkipLast(1)

		count := 0
		skipped(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when skipping single element, got %d", count)
		}
	})

	t.Run("two elements skip one", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{10, 20})

		skipped := enumerator.SkipLast(1)

		expected := []int{10}
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

	t.Run("negative skip count", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		skipped := enumerator.SkipLast(-1)

		expected := []int{1, 2, 3, 4, 5}
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

func TestSkipLastBoolean(t *testing.T) {
	t.Run("skip last booleans", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		skipped := enumerator.SkipLast(2)

		expected := []bool{true, false, true}
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

func TestOrderEnumeratorSkipLast(t *testing.T) {
	t.Run("order enumerator skip last with sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		skipped := ordered.SkipLast(2)

		expected := []int{1, 2, 3, 5}
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

	t.Run("order enumerator any skip last with complex sorting", func(t *testing.T) {
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
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })
		skipped := ordered.SkipLast(1)

		actual := skipped.ToSlice()
		expectedAges := []int{25, 28, 30}

		if len(actual) != len(expectedAges) {
			t.Fatalf("Expected length %d, got %d", len(expectedAges), len(actual))
		}

		for i, expectedAge := range expectedAges {
			if actual[i].Age != expectedAge {
				t.Errorf("Expected age %d at index %d, got %d", expectedAge, i, actual[i].Age)
			}
		}

		expectedNames := []string{"Alice", "Diana", "Charlie"}
		for i, expectedName := range expectedNames {
			if actual[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, actual[i].Name)
			}
		}
	})

	t.Run("order enumerator skip last with multiple sorting levels", func(t *testing.T) {
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
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return a.Value - b.Value
		})
		skipped := ordered.SkipLast(1)

		actual := skipped.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		if actual[0].Category != "A" || actual[0].Value != 15 || actual[0].Name != "Third" {
			t.Errorf("Expected first record {A,15,Third}, got %+v", actual[0])
		}

		if actual[1].Category != "A" || actual[1].Value != 20 || actual[1].Name != "First" {
			t.Errorf("Expected second record {A,20,First}, got %+v", actual[1])
		}

		if actual[2].Category != "B" || actual[2].Value != 10 || actual[2].Name != "Second" {
			t.Errorf("Expected third record {B,10,Second}, got %+v", actual[2])
		}
	})

	t.Run("order enumerator skip last with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		skipped := ordered.SkipLast(2)

		actual := skipped.ToSlice()
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

	t.Run("order enumerator skip last preserves stability", func(t *testing.T) {
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
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })
		skipped := ordered.SkipLast(1)

		actual := skipped.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		// Check stability for Value=1
		if actual[0].Value != 1 || actual[0].Index != 2 {
			t.Errorf("Expected {1,2} at index 0, got %+v", actual[0])
		}
		if actual[1].Value != 1 || actual[1].Index != 4 {
			t.Errorf("Expected {1,4} at index 1, got %+v", actual[1])
		}

		// Check first element of Value=2
		if actual[2].Value != 2 || actual[2].Index != 1 {
			t.Errorf("Expected {2,1} at index 2, got %+v", actual[2])
		}
	})

	t.Run("order enumerator skip last with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		skipped := ordered.SkipLast(2)

		actual := skipped.ToSlice()
		expected := []int{9, 8, 5, 3}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator any skip last with complex struct and custom sorting", func(t *testing.T) {
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
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
		skipped := ordered.SkipLast(1)

		actual := skipped.ToSlice()
		expectedPriorities := []int{1, 2, 3}

		if len(actual) != len(expectedPriorities) {
			t.Fatalf("Expected length %d, got %d", len(expectedPriorities), len(actual))
		}

		for i, expectedPriority := range expectedPriorities {
			if actual[i].Priority != expectedPriority {
				t.Errorf("Expected priority %d at index %d, got %d", expectedPriority, i, actual[i].Priority)
			}
		}

		expectedNames := []string{"High", "Medium", "Low"}
		for i, expectedName := range expectedNames {
			if actual[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, actual[i].Name)
			}
		}
	})
}

func BenchmarkSkipLast(b *testing.B) {
	b.Run("small skip last", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipLast(10)
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large skip last", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			skipped := enumerator.SkipLast(1000)
			skipped(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
