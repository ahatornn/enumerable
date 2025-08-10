package enumerable

import (
	"testing"
)

func TestExcept(t *testing.T) {
	t.Run("basic except operation", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{2, 4})

		result := first.Except(second)

		expected := []int{1, 3, 5}
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

	t.Run("except with duplicates in first", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 2, 3, 3, 4, 5, 5})
		second := FromSlice([]int{2, 4})

		result := first.Except(second)

		expected := []int{1, 3, 5}
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

	t.Run("except with duplicates in second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{2, 2, 4, 4, 6, 6})

		result := first.Except(second)

		expected := []int{1, 3, 5}
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

	t.Run("except empty second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{})

		result := first.Except(second)

		expected := []int{1, 2, 3}
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

	t.Run("except with nil second", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		var second Enumerator[int] = nil

		result := first.Except(second)

		expected := []int{1, 2, 3}
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

	t.Run("except everything", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{1, 2, 3, 4, 5})

		result := first.Except(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items when excluding everything, got %d", count)
		}
	})
}

func TestExceptString(t *testing.T) {
	t.Run("string except operation", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"hello", "world", "foo", "bar", "baz"})
		second := FromSlice([]string{"world", "bar"})

		result := first.Except(second)

		expected := []string{"hello", "foo", "baz"}
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

	t.Run("string except with empty strings", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]string{"", "hello", "world", ""})
		second := FromSlice([]string{""})

		result := first.Except(second)

		expected := []string{"hello", "world"}
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

func TestExceptEarlyTermination(t *testing.T) {
	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		second := FromSlice([]int{2, 4})

		result := first.Except(second)

		actual := []int{}
		result(func(item int) bool {
			if len(actual) >= 2 {
				return false
			}
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 3}
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

func TestExceptStruct(t *testing.T) {
	t.Run("struct except operation", func(t *testing.T) {
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
		})

		result := first.Except(second)

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Charlie", Age: 35},
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

func TestExceptEdgeCases(t *testing.T) {
	t.Run("nil first enumerator", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		second := FromSlice([]int{1, 2, 3})

		result := first.Except(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil first enumerator, got %d", count)
		}
	})

	t.Run("both nil enumerators", func(t *testing.T) {
		t.Parallel()
		var first Enumerator[int] = nil
		var second Enumerator[int] = nil

		result := first.Except(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from both nil enumerators, got %d", count)
		}
	})

	t.Run("no intersection", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{1, 2, 3})
		second := FromSlice([]int{4, 5, 6})

		result := first.Except(second)

		expected := []int{1, 2, 3}
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

	t.Run("empty first", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]int{})
		second := FromSlice([]int{1, 2, 3})

		result := first.Except(second)

		count := 0
		result(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty first enumeration, got %d", count)
		}
	})
}

func TestExceptBoolean(t *testing.T) {
	t.Run("boolean except operation", func(t *testing.T) {
		t.Parallel()
		first := FromSlice([]bool{true, false, true, false})
		second := FromSlice([]bool{false})

		result := first.Except(second)

		expected := []bool{true}
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
func BenchmarkExcept(b *testing.B) {
	b.Run("small except", func(b *testing.B) {
		first := FromSlice([]int{1, 2, 3, 4, 5})
		second := FromSlice([]int{2, 4})

		for i := 0; i < b.N; i++ {
			result := first.Except(second)
			result(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large except", func(b *testing.B) {
		firstItems := make([]int, 10000)
		secondItems := make([]int, 1000)
		for i := 0; i < 10000; i++ {
			firstItems[i] = i
		}
		for i := 0; i < 1000; i++ {
			secondItems[i] = i * 10
		}

		first := FromSlice(firstItems)
		second := FromSlice(secondItems)

		for i := 0; i < b.N; i++ {
			result := first.Except(second)
			result(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
