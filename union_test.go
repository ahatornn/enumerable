package enumerable

import (
	"testing"
)

func TestUnion(t *testing.T) {
	t.Run("basic union", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4})
		second := FromSlice([]int{3, 4, 5, 6})

		union := first.Union(second)

		expected := []int{1, 2, 3, 4, 5, 6}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with duplicates", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 1, 2, 2, 3})
		second := FromSlice([]int{2, 3, 3, 4, 4})

		union := first.Union(second)

		expected := []int{1, 2, 3, 4}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with no overlap", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{4, 5, 6})

		union := first.Union(second)

		expected := []int{1, 2, 3, 4, 5, 6}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with complete overlap", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{1, 2, 3})

		union := first.Union(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with empty first", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{})
		second := FromSlice([]int{1, 2, 3})

		union := first.Union(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with empty second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{})

		union := first.Union(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with nil first", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		second := FromSlice([]int{1, 2, 3})

		union := first.Union(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with nil second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		var second Enumerator[int] = nil

		union := first.Union(second)

		expected := []int{1, 2, 3}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("union with both nil", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		var second Enumerator[int] = nil

		union := first.Union(second)

		count := 0
		union(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from union of nil enumerators, got %d", count)
		}
	})
}

func TestUnionString(t *testing.T) {
	t.Run("union strings", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"a", "b", "c", "d"})
		second := FromSlice([]string{"c", "d", "e", "f"})

		union := first.Union(second)

		expected := []string{"a", "b", "c", "d", "e", "f"}
		actual := []string{}

		union(func(item string) bool {
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

	t.Run("union strings with empty strings", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"", "hello", "world"})
		second := FromSlice([]string{"world", "", "go"})

		union := first.Union(second)

		expected := []string{"", "hello", "world", "go"}
		actual := []string{}

		union(func(item string) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected '%s' at index %d, got '%s'", v, i, actual[i])
			}
		}
	})
}

func TestUnionEarlyTermination(t *testing.T) {
	t.Run("early termination by consumer", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{6, 7, 8, 9, 10})

		union := first.Union(second)

		actual := []int{}
		union(func(item int) bool {
			if len(actual) >= 4 {
				return false // Early termination by consumer
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 2, 3, 4}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("early termination stops second enumeration", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := Repeat(42, 1000) // Large second enumeration

		union := first.Union(second)

		actual := []int{}
		union(func(item int) bool {
			if len(actual) >= 4 {
				return false // Should stop processing second enumeration
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 2, 3, 42}
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

func TestUnionStruct(t *testing.T) {
	t.Run("union structs", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		first := FromSlice([]Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		})

		second := FromSlice([]Person{
			{Name: "Bob", Age: 25}, // Duplicate
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
		})

		union := first.Union(second)

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
		}

		actual := []Person{}
		union(func(item Person) bool {
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

func TestUnionEdgeCases(t *testing.T) {
	t.Run("single element each", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{42})
		second := FromSlice([]int{24})

		union := first.Union(second)

		expected := []int{42, 24}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("single element duplicate", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{42})
		second := FromSlice([]int{42})

		union := first.Union(second)

		expected := []int{42}
		actual := []int{}

		union(func(item int) bool {
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

	t.Run("zero values", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{0, 0, 1})
		second := FromSlice([]int{1, 0, 2})

		union := first.Union(second)

		expected := []int{0, 1, 2}
		actual := []int{}

		union(func(item int) bool {
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

func TestUnionBoolean(t *testing.T) {
	t.Run("union booleans", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]bool{true, false, true})
		second := FromSlice([]bool{false, true, false})

		union := first.Union(second)

		expected := []bool{true, false}
		actual := []bool{}

		union(func(item bool) bool {
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

// Benchmark для проверки производительности
func BenchmarkUnion(b *testing.B) {
	b.Run("small union", func(b *testing.B) {
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{4, 5, 6, 7, 8})

		for i := 0; i < b.N; i++ {
			union := first.Union(second)
			union(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large union with overlap", func(b *testing.B) {
		firstItems := make([]int, 5000)
		secondItems := make([]int, 5000)
		for i := 0; i < 5000; i++ {
			firstItems[i] = i
			secondItems[i] = i + 2500 // 50% overlap
		}

		first := FromSlice(firstItems)
		second := FromSlice(secondItems)

		for i := 0; i < b.N; i++ {
			union := first.Union(second)
			union(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
