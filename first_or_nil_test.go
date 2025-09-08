package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestFirstOrNil(t *testing.T) {
	t.Run("first element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 1 {
			t.Errorf("Expected value 1, got %d", *result)
		}
	})

	t.Run("first element from non-empty for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if len(*result) != 2 {
			t.Errorf("Expected first element length 2, got %d", len(*result))
		}

		if (*result)[0] != 1 || (*result)[1] != 2 {
			t.Errorf("Expected first element [1,2], got %v", *result)
		}
	})

	t.Run("first element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 42 {
			t.Errorf("Expected value 42, got %d", *result)
		}
	})

	t.Run("first element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty slice, got pointer to %d", *result)
		}
	})

	t.Run("first element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %d", *result)
		}
	})

	t.Run("first string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != "hello" {
			t.Errorf("Expected value 'hello', got '%s'", *result)
		}
	})

	t.Run("first empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "world", "go"})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != "" {
			t.Errorf("Expected empty string, got '%s'", *result)
		}
	})
}

func TestFirstOrNilStruct(t *testing.T) {
	t.Run("first struct element", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}

		enumerator := FromSlice(people)
		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		expected := Person{Name: "Alice", Age: 30}
		if *result != expected {
			t.Errorf("Expected %+v, got %+v", expected, *result)
		}
	})

	t.Run("first struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty struct slice, got %+v", *result)
		}
	})
}

func TestFirstOrNilBoolean(t *testing.T) {
	t.Run("first true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != true {
			t.Errorf("Expected true, got %t", *result)
		}
	})

	t.Run("first false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != false {
			t.Errorf("Expected false, got %t", *result)
		}
	})
}

func TestFirstOrNilEarlyTermination(t *testing.T) {
	t.Run("stops after first element", func(t *testing.T) {
		t.Parallel()
		callCount := 0

		enumerator := func(yield func(int) bool) {
			for i := 1; i <= 100; i++ {
				callCount++
				if !yield(i) {
					return
				}
			}
		}

		var enum Enumerator[int] = enumerator
		result := enum.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 1 {
			t.Errorf("Expected value 1, got %d", *result)
		}

		if callCount != 1 {
			t.Errorf("Expected exactly 1 call, got %d", callCount)
		}
	})
}

func TestFirstOrNilEdgeCases(t *testing.T) {
	t.Run("first zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 1, 2, 3})

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != 0 {
			t.Errorf("Expected value 0, got %d", *result)
		}
	})

	t.Run("distinguishing nil from zero value", func(t *testing.T) {
		t.Parallel()
		empty := FromSlice([]int{})
		withZero := FromSlice([]int{0})

		emptyResult := empty.FirstOrNil()
		zeroResult := withZero.FirstOrNil()

		// emptyResult должен быть nil
		if emptyResult != nil {
			t.Errorf("Expected nil for empty slice, got pointer")
		}

		if zeroResult == nil {
			t.Errorf("Expected pointer for slice with zero value, got nil")
		} else if *zeroResult != 0 {
			t.Errorf("Expected 0, got %d", *zeroResult)
		}
	})

	t.Run("first with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.FirstOrNil()

		if result == nil {
			t.Fatal("Expected pointer to first element, got nil")
		}

		if *result != "test" {
			t.Errorf("Expected 'test', got '%s'", *result)
		}
	})
}

func TestOrderEnumeratorFirstOrNil(t *testing.T) {
	t.Run("order enumerator first or nil with integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 1 {
			t.Errorf("Expected result 1, got %d", *result)
		}
	})

	t.Run("order enumerator any first or nil with strings", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date"})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != "apple" {
			t.Errorf("Expected result 'apple', got %s", *result)
		}
	})

	t.Run("order enumerator first or nil with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil, got pointer to %d", *result)
		}
	})

	t.Run("order enumerator any first or nil with empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil, got pointer to %s", *result)
		}
	})

	t.Run("order enumerator first or nil with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %d", *result)
		}
	})

	t.Run("order enumerator any first or nil with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = nil

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.FirstOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %s", *result)
		}
	})

	t.Run("order enumerator first or nil with struct and multiple sorting levels", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Charlie", Age: 30},
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Diana", Age: 25},
		}
		enumerator := FromSlice(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
			ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) })

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			expected := Person{Name: "Alice", Age: 25}
			if *result != expected {
				t.Errorf("Expected result %+v, got %+v", expected, *result)
			}
		}
	})

	t.Run("order enumerator any first or nil with complex struct", func(t *testing.T) {
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
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Name != "High" || result.Priority != 1 {
				t.Errorf("Expected config with name 'High' and priority 1, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator first or nil with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 9 {
			t.Errorf("Expected result 9 (maximum), got %d", *result)
		}
	})

	t.Run("order enumerator any first or nil with descending order", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float64
		}

		products := []Product{
			{Name: "Laptop", Price: 1200.0},
			{Name: "Phone", Price: 800.0},
			{Name: "Tablet", Price: 500.0},
		}
		var enumerator = FromSliceAny(products)

		ordered := enumerator.OrderByDescending(func(a, b Product) int {
			if a.Price < b.Price {
				return -1
			}
			if a.Price > b.Price {
				return 1
			}
			return 0
		})

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Name != "Laptop" || result.Price != 1200.0 {
				t.Errorf("Expected most expensive product, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator first or nil distinguishes zero value from empty", func(t *testing.T) {
		t.Parallel()
		enumeratorWithZero := FromSlice([]int{0, 1, 2})
		orderedWithZero := enumeratorWithZero.OrderBy(comparer.ComparerInt)

		resultWithZero := orderedWithZero.FirstOrNil()

		if resultWithZero == nil {
			t.Error("Expected pointer for slice with zero value, got nil")
		} else if *resultWithZero != 0 {
			t.Errorf("Expected zero value, got %d", *resultWithZero)
		}

		emptyEnumerator := FromSlice([]int{})
		orderedEmpty := emptyEnumerator.OrderBy(comparer.ComparerInt)

		resultEmpty := orderedEmpty.FirstOrNil()

		if resultEmpty != nil {
			t.Errorf("Expected nil for empty slice, got pointer to %d", *resultEmpty)
		}
	})

	t.Run("order enumerator any first or nil with single element", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Data  []string
		}

		items := []Item{{Value: 42, Data: []string{"test"}}}
		var enumerator = FromSliceAny(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Value != 42 {
				t.Errorf("Expected value 42, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator first or nil with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 1 {
			t.Errorf("Expected minimum value 1, got %d", *result)
		}
	})

	t.Run("order enumerator first or nil preserves stability", func(t *testing.T) {
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

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Value != 1 || result.Index != 2 {
				t.Errorf("Expected {Value: 1, Index: 2}, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator first or nil with multiple then by levels", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Type     string
			Name     string
		}

		records := []Record{
			{Category: "B", Type: "X", Name: "Second"},
			{Category: "A", Type: "Y", Name: "First"},
			{Category: "B", Type: "Y", Name: "Third"},
			{Category: "A", Type: "X", Name: "Fourth"},
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return compareStrings(a.Type, b.Type)
		}).ThenBy(func(a, b Record) int {
			return compareStrings(a.Name, b.Name)
		})

		result := ordered.FirstOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Category != "A" || result.Type != "X" {
				t.Errorf("Expected first record to be Category A, Type X, got %+v", *result)
			}
		}
	})
}

func BenchmarkOrderEnumeratorFirstOrNil(b *testing.B) {
	b.Run("order enumerator first or nil small", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = 100 - i
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.FirstOrNil()
			if result == nil || *result != 1 {
				b.Fatalf("Expected pointer to 1, got %v", result)
			}
		}
	})

	b.Run("order enumerator any first or nil medium", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}

		people := make([]Person, 1000)
		for i := 0; i < 1000; i++ {
			people[i] = Person{Name: fmt.Sprintf("Person%d", i), Age: i}
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.FirstOrNil()
			if result == nil || result.Age != 0 {
				b.Fatalf("Expected pointer to person with age 0, got %v", result)
			}
		}
	})

	b.Run("order enumerator first or nil with multiple levels", func(b *testing.B) {
		type Record struct {
			Category string
			Value    int
		}

		records := make([]Record, 500)
		for i := 0; i < 500; i++ {
			records[i] = Record{
				Category: fmt.Sprintf("Cat%d", i%10),
				Value:    i,
			}
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int { return compareStrings(a.Category, b.Category) }).
			ThenBy(func(a, b Record) int { return a.Value - b.Value })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.FirstOrNil()
			if result == nil {
				b.Fatal("Expected pointer to record, got nil")
			}
			_ = result
		}
	})

	b.Run("order enumerator first or nil empty", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.FirstOrNil()
			if result != nil {
				b.Fatalf("Expected nil, got %v", result)
			}
		}
	})

	b.Run("order enumerator any first or nil descending", func(b *testing.B) {
		items := make([]int, 200)
		for i := 0; i < 200; i++ {
			items[i] = i
		}
		var enumerator = FromSliceAny(items)

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.FirstOrNil()
			if result == nil || *result != 199 {
				b.Fatalf("Expected pointer to 199, got %v", result)
			}
		}
	})

	b.Run("order enumerator first or nil with nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.FirstOrNil()
			if result != nil {
				b.Fatalf("Expected nil, got %v", result)
			}
		}
	})
}

func BenchmarkFirstOrNil(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrNil()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrNil()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.FirstOrNil()
		}
	})
}
