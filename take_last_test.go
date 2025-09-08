package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestTakeLast(t *testing.T) {
	t.Run("basic take last", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})

		taken := enumerator.TakeLast(3)

		expected := []int{6, 7, 8}
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

	t.Run("basic take last for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
			{11, 12},
			{13, 14},
			{15, 16},
		})

		taken := enumerator.TakeLast(3)

		expected := [][]int{
			{11, 12},
			{13, 14},
			{15, 16},
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

	t.Run("take last zero elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(0)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when taking zero elements, got %d", count)
		}
	})

	t.Run("take negative number", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(-1)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when taking negative number, got %d", count)
		}
	})

	t.Run("take more than available", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		taken := enumerator.TakeLast(5)

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

	t.Run("take exactly all elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(5)

		expected := []int{1, 2, 3, 4, 5}
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

		taken := enumerator.TakeLast(3)

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

		taken := enumerator.TakeLast(3)

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

		taken := enumerator.TakeLast(3)

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

func TestTakeLastString(t *testing.T) {
	t.Run("take last strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"a", "b", "c", "d", "e"})

		taken := enumerator.TakeLast(3)

		expected := []string{"c", "d", "e"}
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

func TestTakeLastEarlyTermination(t *testing.T) {
	t.Run("early termination by consumer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		taken := enumerator.TakeLast(3)

		actual := []int{}
		taken(func(item int) bool {
			if len(actual) >= 2 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{3, 4}
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

func TestTakeLastStruct(t *testing.T) {
	t.Run("take last structs", func(t *testing.T) {
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
		taken := enumerator.TakeLast(3)

		expected := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
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

func TestTakeLastEdgeCases(t *testing.T) {
	t.Run("single element take one", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		taken := enumerator.TakeLast(1)

		expected := []int{42}
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

	t.Run("single element take zero", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		taken := enumerator.TakeLast(0)

		count := 0
		taken(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when taking zero elements, got %d", count)
		}
	})

	t.Run("large take from small collection", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		taken := enumerator.TakeLast(1000)

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

	t.Run("take last with ring buffer wraparound", func(t *testing.T) {
		t.Parallel()
		// Test case where ring buffer wraps around
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		taken := enumerator.TakeLast(5)

		expected := []int{95, 96, 97, 98, 99}
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

func TestTakeLastBoolean(t *testing.T) {
	t.Run("take last booleans", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		taken := enumerator.TakeLast(3)

		expected := []bool{true, false, true}
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

func TestOrderEnumeratorTakeLast(t *testing.T) {
	t.Run("order enumerator take last with sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(3)
		expected := []int{5, 8, 9}
		actual := taken.ToSlice()

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator any take last with complex struct", func(t *testing.T) {
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
		taken := ordered.TakeLast(3)
		actual := taken.ToSlice()

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

	t.Run("order enumerator take last with multiple sorting levels", func(t *testing.T) {
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
		taken := ordered.TakeLast(3)
		actual := taken.ToSlice()

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

	t.Run("order enumerator take last zero elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(0)
		actual := taken.ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice when taking 0 elements, got length %d", len(actual))
		}
	})

	t.Run("order enumerator take last negative number", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(-1)

		actual := taken.ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice when taking negative elements, got length %d", len(actual))
		}
	})

	t.Run("order enumerator take last more than available", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(5)
		actual := taken.ToSlice()
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

	t.Run("order enumerator any take last with complex struct and custom sorting", func(t *testing.T) {
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
		taken := ordered.TakeLast(3)
		actual := taken.ToSlice()

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

	t.Run("order enumerator take last with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2, 4, 4})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(5)
		actual := taken.ToSlice()

		if len(actual) != 5 {
			t.Fatalf("Expected length 5, got %d", len(actual))
		}

		expectedLastFive := []int{2, 3, 3, 4, 4}
		for i, v := range expectedLastFive {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator take last preserves stability", func(t *testing.T) {
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
		taken := ordered.TakeLast(3)

		actual := taken.ToSlice()

		if len(actual) != 3 {
			t.Fatalf("Expected length 3, got %d", len(actual))
		}

		if actual[0].Value != 2 || actual[0].Index != 1 {
			t.Errorf("Expected {2,1} at index 0, got %+v", actual[0])
		}
		if actual[1].Value != 2 || actual[1].Index != 3 {
			t.Errorf("Expected {2,3} at index 1, got %+v", actual[1])
		}

		if actual[2].Value != 3 || actual[2].Index != 5 {
			t.Errorf("Expected {3,5} at index 2, got %+v", actual[2])
		}
	})

	t.Run("order enumerator take last with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		taken := ordered.TakeLast(3)

		actual := taken.ToSlice()
		expected := []int{3, 2, 1}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator take last exactly all elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(5)
		actual := taken.ToSlice()
		expected := []int{1, 2, 3, 4, 5}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("order enumerator take last with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(1)

		actual := taken.ToSlice()
		expected := []int{42}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		if actual[0] != expected[0] {
			t.Errorf("Expected %d, got %d", expected[0], actual[0])
		}
	})

	t.Run("order enumerator take last with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		taken := ordered.TakeLast(3)

		actual := taken.ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice from empty enumerator, got length %d", len(actual))
		}
	})

	t.Run("order enumerator take last with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		actual := enumerator.
			OrderBy(comparer.ComparerInt).
			TakeLast(3).
			ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice from nil enumerator, got length %d", len(actual))
		}
	})

	t.Run("order enumerator any take last with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = nil

		ordered := enumerator.OrderBy(comparer.ComparerString)
		taken := ordered.TakeLast(3)

		actual := taken.ToSlice()

		if len(actual) != 0 {
			t.Errorf("Expected empty slice from nil enumerator, got length %d", len(actual))
		}
	})
}

func BenchmarkTakeLast(b *testing.B) {
	b.Run("small take last", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeLast(10)
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large take last", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			taken := enumerator.TakeLast(1000)
			taken(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
