package enumerable

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	t.Run("basic intersection", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{3, 4, 5, 6, 7})

		result := first.Intersect(second)

		expected := []int{3, 4, 5}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("intersection with duplicates in first", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 2, 3, 3, 4, 5})
		second := FromSlice([]int{2, 3, 4, 6})

		result := first.Intersect(second)

		expected := []int{2, 3, 4}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("intersection with duplicates in second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4})
		second := FromSlice([]int{2, 2, 3, 3, 5, 5})

		result := first.Intersect(second)

		expected := []int{2, 3}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("no intersection", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{4, 5, 6})

		result := first.Intersect(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when no intersection, got %d", count)
		}
	})

	t.Run("empty first enumeration", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{})
		second := FromSlice([]int{1, 2, 3})

		result := first.Intersect(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty first enumeration, got %d", count)
		}
	})

	t.Run("empty second enumeration", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{})

		result := first.Intersect(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty second enumeration, got %d", count)
		}
	})
}

func TestIntersectString(t *testing.T) {
	t.Run("string intersection", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"hello", "world", "foo", "bar"})
		second := FromSlice([]string{"world", "bar", "baz"})

		result := first.Intersect(second)

		expected := []string{"world", "bar"}
		actual := []string{}

		result(func(item string) bool {
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

	t.Run("string intersection with empty strings", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"", "hello", "world"})
		second := FromSlice([]string{"hello", "", "foo"})

		result := first.Intersect(second)

		expected := []string{"", "hello"}
		actual := []string{}

		result(func(item string) bool {
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

func TestIntersectEarlyTermination(t *testing.T) {
	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5, 6, 7})
		second := FromSlice([]int{2, 4, 6, 8})

		result := first.Intersect(second)

		actual := []int{}
		result(func(item int) bool {
			if len(actual) >= 2 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{2, 4}
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

func TestIntersectStruct(t *testing.T) {
	t.Run("struct intersection", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		first := FromSlice([]Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
		})

		second := FromSlice([]Person{
			{Name: "Bob", Age: 25},
			{Name: "Diana", Age: 28},
			{Name: "Eve", Age: 32},
		})

		result := first.Intersect(second)

		expected := []Person{
			{Name: "Bob", Age: 25},
			{Name: "Diana", Age: 28},
		}

		actual := []Person{}
		result(func(item Person) bool {
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

func TestIntersectEdgeCases(t *testing.T) {
	t.Run("nil first enumerator", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		second := FromSlice([]int{1, 2, 3})

		result := first.Intersect(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil first enumerator, got %d", count)
		}
	})

	t.Run("nil second enumerator", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		var second Enumerator[int] = nil

		result := first.Intersect(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil second enumerator, got %d", count)
		}
	})

	t.Run("both nil enumerators", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		var second Enumerator[int] = nil

		result := first.Intersect(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from both nil enumerators, got %d", count)
		}
	})

	t.Run("identical enumerations", func(t *testing.T) {
		t.Parallel()
		items := []int{1, 2, 3, 4, 5}
		first := FromSlice(items)
		second := FromSlice(items)

		result := first.Intersect(second)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		result(func(item int) bool {
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

	t.Run("first contains all of second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{2, 4})

		result := first.Intersect(second)

		expected := []int{2, 4}
		actual := []int{}

		result(func(item int) bool {
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

func TestIntersectBoolean(t *testing.T) {
	t.Run("boolean intersection", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]bool{true, false, true})
		second := FromSlice([]bool{false, true, false})

		result := first.Intersect(second)

		expected := []bool{true, false}
		actual := []bool{}

		result(func(item bool) bool {
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
func BenchmarkIntersect(b *testing.B) {
	b.Run("small intersection", func(b *testing.B) {
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{3, 4, 5, 6, 7})

		for i := 0; i < b.N; i++ {
			result := first.Intersect(second)
			result(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large intersection", func(b *testing.B) {
		firstItems := make([]int, 10000)
		secondItems := make([]int, 8000)
		for i := 0; i < 10000; i++ {
			firstItems[i] = i
		}
		for i := 0; i < 8000; i++ {
			secondItems[i] = i + 2000 // Пересечение с 2000 по 9999
		}

		first := FromSlice(firstItems)
		second := FromSlice(secondItems)

		for i := 0; i < b.N; i++ {
			result := first.Intersect(second)
			result(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
