package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestLastOrNil(t *testing.T) {
	t.Run("last element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 5 {
			t.Errorf("Expected value 5, got %d", *result)
		}
	})

	t.Run("last element from non-empty slice for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if len(*result) != 2 {
			t.Errorf("Expected last element length 2, got %d", len(*result))
		}

		if (*result)[0] != 5 || (*result)[1] != 6 {
			t.Errorf("Expected last element [5,6], got %v", *result)
		}
	})

	t.Run("last element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 42 {
			t.Errorf("Expected value 42, got %d", *result)
		}
	})

	t.Run("last element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty slice, got pointer to %d", *result)
		}
	})

	t.Run("last element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %d", *result)
		}
	})

	t.Run("last string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != "go" {
			t.Errorf("Expected value 'go', got '%s'", *result)
		}
	})

	t.Run("last empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", ""})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != "" {
			t.Errorf("Expected empty string, got '%s'", *result)
		}
	})
}

func TestLastOrNilStruct(t *testing.T) {
	t.Run("last struct element", func(t *testing.T) {
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
		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		expected := Person{Name: "Charlie", Age: 35}
		if *result != expected {
			t.Errorf("Expected %+v, got %+v", expected, *result)
		}
	})

	t.Run("last struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}

		enumerator := FromSlice(people)
		result := enumerator.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for empty struct slice, got %+v", *result)
		}
	})
}

func TestLastOrNilBoolean(t *testing.T) {
	t.Run("last true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != false {
			t.Errorf("Expected false, got %t", *result)
		}
	})

	t.Run("last false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != true {
			t.Errorf("Expected true, got %t", *result)
		}
	})
}

func TestLastOrNilWithOperations(t *testing.T) {
	t.Run("last after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		result := filtered.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 6 {
			t.Errorf("Expected value 6, got %d", *result)
		}
	})

	t.Run("last after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(5)

		result := taken.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 5 {
			t.Errorf("Expected value 5, got %d", *result)
		}
	})
}

func TestLastOrNilEdgeCases(t *testing.T) {
	t.Run("last zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 0})

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 0 {
			t.Errorf("Expected value 0, got %d", *result)
		}
	})

	t.Run("distinguishing nil from zero value", func(t *testing.T) {
		t.Parallel()
		empty := FromSlice([]int{})
		withZero := FromSlice([]int{0})

		emptyResult := empty.LastOrNil()
		zeroResult := withZero.LastOrNil()

		if emptyResult != nil {
			t.Errorf("Expected nil for empty slice, got pointer")
		}

		if zeroResult == nil {
			t.Errorf("Expected pointer for slice with zero value, got nil")
		} else if *zeroResult != 0 {
			t.Errorf("Expected 0, got %d", *zeroResult)
		}
	})

	t.Run("last with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != "test" {
			t.Errorf("Expected 'test', got '%s'", *result)
		}
	})

	t.Run("last with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		result := enumerator.LastOrNil()

		if result == nil {
			t.Fatal("Expected pointer to last element, got nil")
		}

		if *result != 5 {
			t.Errorf("Expected value 5, got %d", *result)
		}
	})
}

func TestOrderEnumeratorLastOrNil(t *testing.T) {
	t.Run("order enumerator last or nil with integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 9 {
			t.Errorf("Expected result 9, got %d", *result)
		}
	})

	t.Run("order enumerator any last or nil with strings", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date"})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != "date" {
			t.Errorf("Expected result 'date', got %s", *result)
		}
	})

	t.Run("order enumerator last or nil with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil, got pointer to %d", *result)
		}
	})

	t.Run("order enumerator any last or nil with empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil, got pointer to %s", *result)
		}
	})

	t.Run("order enumerator last or nil with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %d", *result)
		}
	})

	t.Run("order enumerator any last or nil with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = nil

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.LastOrNil()

		if result != nil {
			t.Errorf("Expected nil for nil enumerator, got pointer to %s", *result)
		}
	})

	t.Run("order enumerator last or nil with struct and multiple sorting levels", func(t *testing.T) {
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

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			expected := Person{Name: "Charlie", Age: 30}
			if *result != expected {
				t.Errorf("Expected result %+v, got %+v", expected, *result)
			}
		}
	})

	t.Run("order enumerator any last or nil with complex struct", func(t *testing.T) {
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

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Name != "Low" || result.Priority != 3 {
				t.Errorf("Expected config with name 'Low' and priority 3, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator last or nil with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 1 {
			t.Errorf("Expected result 1 (minimum), got %d", *result)
		}
	})

	t.Run("order enumerator any last or nil with descending order", func(t *testing.T) {
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

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Name != "Tablet" || result.Price != 500.0 {
				t.Errorf("Expected least expensive product, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator last or nil distinguishes zero value from empty", func(t *testing.T) {
		t.Parallel()
		enumeratorWithZero := FromSlice([]int{0, 1, 2})
		orderedWithZero := enumeratorWithZero.OrderBy(comparer.ComparerInt)

		resultWithZero := orderedWithZero.LastOrNil()

		if resultWithZero == nil {
			t.Error("Expected pointer for slice with zero value, got nil")
		} else if *resultWithZero != 2 {
			t.Errorf("Expected last value 2, got %d", *resultWithZero)
		}

		emptyEnumerator := FromSlice([]int{})
		orderedEmpty := emptyEnumerator.OrderBy(comparer.ComparerInt)

		resultEmpty := orderedEmpty.LastOrNil()

		if resultEmpty != nil {
			t.Errorf("Expected nil for empty slice, got pointer to %d", *resultEmpty)
		}
	})

	t.Run("order enumerator any last or nil with single element", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Data  []string
		}

		items := []Item{{Value: 42, Data: []string{"test"}}}
		var enumerator = FromSliceAny(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Value != 42 {
				t.Errorf("Expected value 42, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator last or nil with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 3 {
			t.Errorf("Expected maximum value 3, got %d", *result)
		}
	})

	t.Run("order enumerator last or nil preserves stability", func(t *testing.T) {
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

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			if result.Value != 2 || result.Index != 3 {
				t.Errorf("Expected {Value: 2, Index: 3}, got %+v", *result)
			}
		}
	})

	t.Run("order enumerator last or nil with multiple then by levels", func(t *testing.T) {
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

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else {
			expected := Record{Category: "B", Type: "Y", Name: "Third"}
			if *result != expected {
				t.Errorf("Expected last record to be %+v, got %+v", expected, *result)
			}
		}
	})

	t.Run("order enumerator last or nil with boolean values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		ordered := enumerator.OrderBy(comparer.ComparerBool)
		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != true {
			t.Errorf("Expected true (maximum), got %v", *result)
		}
	})

	t.Run("order enumerator any last or nil with floating point values", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]float64{3.14, 1.41, 2.71, 0.57})

		ordered := enumerator.OrderBy(comparer.ComparerFloat64)

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 3.14 {
			t.Errorf("Expected maximum value 3.14, got %f", *result)
		}
	})

	t.Run("order enumerator last or nil with reverse sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.LastOrNil()

		if result == nil {
			t.Error("Expected pointer to element, got nil")
		} else if *result != 1 {
			t.Errorf("Expected result 1 (minimum in reverse order), got %d", *result)
		}
	})
}

func BenchmarkOrderEnumeratorLastOrNil(b *testing.B) {
	b.Run("order enumerator last or nil small", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrNil()
			if result == nil || *result != 99 {
				b.Fatalf("Expected pointer to 99, got %v", result)
			}
		}
	})

	b.Run("order enumerator any last or nil medium", func(b *testing.B) {
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
			result := ordered.LastOrNil()
			if result == nil || result.Age != 999 {
				b.Fatalf("Expected pointer to person with age 999, got %v", result)
			}
		}
	})

	b.Run("order enumerator last or nil with multiple levels", func(b *testing.B) {
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
			result := ordered.LastOrNil()
			if result == nil {
				b.Fatal("Expected pointer to record, got nil")
			}
			_ = result
		}
	})

	b.Run("order enumerator last or nil empty", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrNil()
			if result != nil {
				b.Fatalf("Expected nil, got %v", result)
			}
		}
	})

	b.Run("order enumerator any last or nil descending", func(b *testing.B) {
		items := make([]int, 200)
		for i := 0; i < 200; i++ {
			items[i] = i
		}
		var enumerator = FromSliceAny(items)

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrNil()
			if result == nil || *result != 0 {
				b.Fatalf("Expected pointer to 0, got %v", result)
			}
		}
	})

	b.Run("order enumerator last or nil with nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrNil()
			if result != nil {
				b.Fatalf("Expected nil, got %v", result)
			}
		}
	})

	b.Run("order enumerator last or nil with struct", func(b *testing.B) {
		type Config struct {
			Name     string
			Priority int
		}

		configs := make([]Config, 100)
		for i := 0; i < 100; i++ {
			configs[i] = Config{Name: fmt.Sprintf("Config%d", i), Priority: i}
		}
		enumerator := FromSlice(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrNil()
			if result == nil || result.Priority != 99 {
				b.Fatalf("Expected config with priority 99, got %v", result)
			}
		}
	})

	b.Run("order enumerator last or nil with duplicate values", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i % 100
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrNil()
			if result == nil || *result != 99 {
				b.Fatalf("Expected pointer to 99, got %v", result)
			}
		}
	})
}

func BenchmarkLastOrNil(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrNil()
		}
	})
}
