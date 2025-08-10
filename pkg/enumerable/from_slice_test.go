package enumerable

import (
	"testing"
)

func TestFromSlice(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		enumerator := FromSlice(slice)

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items, got %d", count)
		}
	})

	t.Run("slice with items", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(slice)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		enumerator(func(item int) bool {
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

	t.Run("early termination", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(slice)

		actual := []int{}
		enumerator(func(item int) bool {
			actual = append(actual, item)
			return len(actual) < 3
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

	t.Run("string slice", func(t *testing.T) {
		slice := []string{"hello", "world", "test"}
		enumerator := FromSlice(slice)

		actual := []string{}
		enumerator(func(item string) bool {
			actual = append(actual, item)
			return true
		})

		expected := []string{"hello", "world", "test"}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
			}
		}
	})

	t.Run("single element", func(t *testing.T) {
		slice := []int{42}
		enumerator := FromSlice(slice)

		var result int
		found := false

		enumerator(func(item int) bool {
			result = item
			found = true
			return true
		})

		if !found {
			t.Error("Expected to find one element, but found none")
		}

		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
	})
}
