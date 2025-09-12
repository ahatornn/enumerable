package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestDistinct(t *testing.T) {
	t.Run("remove duplicates", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 1, 4, 3})

		distinct := enumerator.Distinct()

		expected := []int{1, 2, 3, 4}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("no duplicates", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		distinct := enumerator.Distinct()

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("all duplicates", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 1, 1, 1})

		distinct := enumerator.Distinct()

		expected := []int{1}
		actual := []int{}

		distinct(func(item int) bool {
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

		distinct := enumerator.Distinct()

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty enumeration, got %d", count)
		}
	})

	t.Run("nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		distinct := enumerator.Distinct()

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})
}

func TestDistinctString(t *testing.T) {
	t.Run("string duplicates", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "hello", "go", "world"})

		distinct := enumerator.Distinct()

		expected := []string{"hello", "world", "go"}
		actual := []string{}

		distinct(func(item string) bool {
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

	t.Run("empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "hello", "", "world", ""})

		distinct := enumerator.Distinct()

		expected := []string{"", "hello", "world"}
		actual := []string{}

		distinct(func(item string) bool {
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

func TestDistinctEarlyTermination(t *testing.T) {
	t.Run("early termination", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4, 4, 5})

		distinct := enumerator.Distinct()

		actual := []int{}
		distinct(func(item int) bool {
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

func TestDistinctStruct(t *testing.T) {
	t.Run("struct distinct", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Alice", Age: 30}, // Duplicate
			{Name: "Charlie", Age: 35},
			{Name: "Bob", Age: 25}, // Duplicate
		}

		enumerator := FromSlice(people)
		distinct := enumerator.Distinct()

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}

		actual := []Person{}
		distinct(func(item Person) bool {
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

func TestDistinctBoolean(t *testing.T) {
	t.Run("boolean distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		distinct := enumerator.Distinct()

		expected := []bool{true, false}
		actual := []bool{}

		distinct(func(item bool) bool {
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

func TestDistinctEdgeCases(t *testing.T) {
	t.Run("single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		distinct := enumerator.Distinct()

		expected := []int{42}
		actual := []int{}

		distinct(func(item int) bool {
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
		enumerator := FromSlice([]int{0, 0, 0})

		distinct := enumerator.Distinct()

		expected := []int{0}
		actual := []int{}

		distinct(func(item int) bool {
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

func TestDistinctForAny(t *testing.T) {
	t.Run("remove duplicates for any", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{1, 2, 2, 3, 1, 4, 3})
		comparer := comparer.Default[int]()
		distinct := enumerator.Distinct(comparer)

		expected := []int{1, 2, 3, 4}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("no duplicates for any", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{1, 2, 3, 4, 5})
		comparer := comparer.Default[int]()
		distinct := enumerator.Distinct(comparer)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("all duplicates for any", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{1, 1, 1, 1})
		comparer := comparer.Default[int]()
		distinct := enumerator.Distinct(comparer)

		expected := []int{1}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("empty slice for any", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{})
		comparer := comparer.Default[int]()
		distinct := enumerator.Distinct(comparer)

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty enumeration, got %d", count)
		}
	})

	t.Run("nil enumerator for any", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil
		comparer := comparer.Default[int]()
		distinct := enumerator.Distinct(comparer)

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})
}

func TestOrderEnumeratorDistinct(t *testing.T) {
	t.Run("remove duplicates from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 2, 1, 3, 2, 1})
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct()

		expected := []int{1, 2, 3}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("no duplicates from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 4, 3, 2, 1})
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct()

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("all duplicates from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 1, 1, 1})
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct()

		expected := []int{1}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("empty slice from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct()

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty enumeration, got %d", count)
		}
	})

	t.Run("nil enumerator from ordered items", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct()

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})
}

func TestOrderEnumeratorAnyDistinct(t *testing.T) {
	t.Run("remove duplicates from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{3, 2, 1, 3, 2, 1})
		eqComparer := comparer.Default[int]()
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct(eqComparer)

		expected := []int{1, 2, 3}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("no duplicates from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{5, 4, 3, 2, 1})
		eqComparer := comparer.Default[int]()
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct(eqComparer)

		expected := []int{1, 2, 3, 4, 5}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("all duplicates from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{1, 1, 1, 1})
		eqComparer := comparer.Default[int]()
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct(eqComparer)

		expected := []int{1}
		actual := []int{}

		distinct(func(item int) bool {
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

	t.Run("empty slice from ordered items", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]int{})
		eqComparer := comparer.Default[int]()
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct(eqComparer)

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty enumeration, got %d", count)
		}
	})

	t.Run("nil enumerator from ordered items", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil
		eqComparer := comparer.Default[int]()
		distinct := enumerator.OrderBy(comparer.ComparerInt).Distinct(eqComparer)

		count := 0
		distinct(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil enumerator, got %d", count)
		}
	})
}

func BenchmarkDistinct(b *testing.B) {
	b.Run("small distinct", func(b *testing.B) {
		items := []int{1, 2, 2, 3, 3, 4, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			distinct := enumerator.Distinct()
			distinct(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large with duplicates", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i % 1000
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			distinct := enumerator.Distinct()
			distinct(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large no duplicates", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			distinct := enumerator.Distinct()
			distinct(func(item int) bool {
				_ = item
				return true
			})
		}
	})
	b.Run("small distinct for any", func(b *testing.B) {
		items := []int{1, 2, 2, 3, 3, 4, 4, 5}
		enumerator := FromSliceAny(items)
		eqComparer := comparer.Default[int]()
		for i := 0; i < b.N; i++ {
			distinct := enumerator.Distinct(eqComparer)
			distinct(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large with duplicates for any", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i % 1000
		}
		enumerator := FromSliceAny(items)
		eqComparer := comparer.Default[int]()
		for i := 0; i < b.N; i++ {
			distinct := enumerator.Distinct(eqComparer)
			distinct(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large no duplicates for any", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSliceAny(items)
		eqComparer := comparer.Default[int]()
		for i := 0; i < b.N; i++ {
			distinct := enumerator.Distinct(eqComparer)
			distinct(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
