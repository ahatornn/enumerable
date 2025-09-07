package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestLastOrDefault(t *testing.T) {
	t.Run("last element from non-empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.LastOrDefault(-1)

		if result != 5 {
			t.Errorf("Expected value 5, got %d", result)
		}
	})

	t.Run("last element from non-empty slice for non-comparable slice", func(t *testing.T) {
		t.Parallel()

		defaultSlice := []int{-1, -2}

		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
		})

		result := enumerator.LastOrDefault(defaultSlice)

		if len(result) != 2 {
			t.Errorf("Expected last element length 2, got %d", len(result))
		}

		if result[0] != 5 || result[1] != 6 {
			t.Errorf("Expected last element [5,6], got %v", result)
		}
	})

	t.Run("last element from single element slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result := enumerator.LastOrDefault(-1)

		if result != 42 {
			t.Errorf("Expected value 42, got %d", result)
		}
	})

	t.Run("last element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.LastOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1, got %d", result)
		}
	})

	t.Run("last element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.LastOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1 for nil enumerator, got %d", result)
		}
	})

	t.Run("last element from nil enumerator any", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		result := enumerator.LastOrDefault(-1)

		if result != -1 {
			t.Errorf("Expected default value -1 for nil enumerator any, got %d", result)
		}
	})

	t.Run("last string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		result := enumerator.LastOrDefault("default")

		if result != "go" {
			t.Errorf("Expected value 'go', got '%s'", result)
		}
	})

	t.Run("last empty string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", ""})

		result := enumerator.LastOrDefault("default")

		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}
	})

	t.Run("last string from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})

		result := enumerator.LastOrDefault("default")

		if result != "default" {
			t.Errorf("Expected default value 'default', got '%s'", result)
		}
	})
}

func TestLastOrDefaultStruct(t *testing.T) {
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

		defaultPerson := Person{Name: "Unknown", Age: 0}
		enumerator := FromSlice(people)
		result := enumerator.LastOrDefault(defaultPerson)

		expected := Person{Name: "Charlie", Age: 35}
		if result != expected {
			t.Errorf("Expected %+v, got %+v", expected, result)
		}
	})

	t.Run("last struct from empty slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{}
		defaultPerson := Person{Name: "Unknown", Age: 0}

		enumerator := FromSlice(people)
		result := enumerator.LastOrDefault(defaultPerson)

		if result != defaultPerson {
			t.Errorf("Expected default person %+v, got %+v", defaultPerson, result)
		}
	})
}

func TestLastOrDefaultBoolean(t *testing.T) {
	t.Run("last false element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		result := enumerator.LastOrDefault(false)

		if result != true {
			t.Errorf("Expected true, got %t", result)
		}
	})

	t.Run("last true element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{false, true, false})

		result := enumerator.LastOrDefault(true)

		if result != false {
			t.Errorf("Expected false, got %t", result)
		}
	})

	t.Run("last boolean from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{})

		result := enumerator.LastOrDefault(true)

		if result != true {
			t.Errorf("Expected default value true, got %t", result)
		}
	})
}

func TestLastOrDefaultWithOperations(t *testing.T) {
	t.Run("last after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		result := filtered.LastOrDefault(-1)

		if result != 6 {
			t.Errorf("Expected value 6, got %d", result)
		}
	})

	t.Run("last after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(5)

		result := taken.LastOrDefault(-1)

		if result != 5 {
			t.Errorf("Expected value 5, got %d", result)
		}
	})
}

func TestLastOrDefaultEdgeCases(t *testing.T) {
	t.Run("last zero value with non-zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 0})

		result := enumerator.LastOrDefault(-1)

		if result != 0 {
			t.Errorf("Expected value 0, got %d", result)
		}
	})

	t.Run("last zero value with zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 0})

		result := enumerator.LastOrDefault(0)

		if result != 0 {
			t.Errorf("Expected value 0, got %d", result)
		}
	})

	t.Run("empty slice with zero default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.LastOrDefault(0)

		if result != 0 {
			t.Errorf("Expected default value 0, got %d", result)
		}
	})

	t.Run("distinguishing zero value from default", func(t *testing.T) {
		t.Parallel()
		withZero := FromSlice([]int{1, 2, 0})
		empty := FromSlice([]int{})

		zeroResult := withZero.LastOrDefault(-1)
		emptyResult := empty.LastOrDefault(-1)

		if zeroResult != 0 {
			t.Errorf("Expected 0 from slice with zero, got %d", zeroResult)
		}

		if emptyResult != -1 {
			t.Errorf("Expected -1 from empty slice, got %d", emptyResult)
		}
	})

	t.Run("last with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		result := enumerator.LastOrDefault("default")

		if result != "test" {
			t.Errorf("Expected 'test', got '%s'", result)
		}
	})

	t.Run("last with range", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 5) // 1, 2, 3, 4, 5

		result := enumerator.LastOrDefault(-1)

		if result != 5 {
			t.Errorf("Expected value 5, got %d", result)
		}
	})
}

func TestOrderEnumeratorLastOrDefault(t *testing.T) {
	t.Run("order enumerator last or default with integers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrDefault(defaultValue)

		if result != 9 {
			t.Errorf("Expected result 9, got %d", result)
		}
	})

	t.Run("order enumerator any last or default with strings", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"banana", "apple", "cherry", "date"})
		defaultValue := "none"

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.LastOrDefault(defaultValue)

		if result != "date" {
			t.Errorf("Expected result 'date', got %s", result)
		}
	})

	t.Run("order enumerator last or default with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("order enumerator any last or default with empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{})
		defaultValue := "empty"

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.LastOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %s, got %s", defaultValue, result)
		}
	})

	t.Run("order enumerator last or default with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil
		defaultValue := 42

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d for nil enumerator, got %d", defaultValue, result)
		}
	})

	t.Run("order enumerator any last or default with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = nil
		defaultValue := "nil"

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result := ordered.LastOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %s for nil enumerator, got %s", defaultValue, result)
		}
	})

	t.Run("order enumerator last or default with struct and multiple sorting levels", func(t *testing.T) {
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
		defaultValue := Person{Name: "Default", Age: 0}

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
			ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) })

		result := ordered.LastOrDefault(defaultValue)

		expected := Person{Name: "Charlie", Age: 30}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator any last or default with complex struct", func(t *testing.T) {
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
		defaultValue := Config{Name: "Default", Priority: -1, Options: []string{}}

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })

		result := ordered.LastOrDefault(defaultValue)

		if result.Name != "Low" || result.Priority != 3 {
			t.Errorf("Expected config with name 'Low' and priority 3, got %+v", result)
		}
	})

	t.Run("order enumerator last or default with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})
		defaultValue := 0

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.LastOrDefault(defaultValue)

		if result != 1 {
			t.Errorf("Expected result 1 (minimum), got %d", result)
		}
	})

	t.Run("order enumerator any last or default with descending order", func(t *testing.T) {
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
		defaultValue := Product{Name: "None", Price: 0.0}

		ordered := enumerator.OrderByDescending(func(a, b Product) int {
			if a.Price < b.Price {
				return -1
			}
			if a.Price > b.Price {
				return 1
			}
			return 0
		})

		result := ordered.LastOrDefault(defaultValue)

		if result.Name != "Tablet" || result.Price != 500.0 {
			t.Errorf("Expected least expensive product, got %+v", result)
		}
	})

	t.Run("order enumerator last or default with zero value as default", func(t *testing.T) {
		t.Parallel()
		enumeratorWithZero := FromSlice([]int{0, 1, 2})
		defaultValue := -1

		orderedWithZero := enumeratorWithZero.OrderBy(comparer.ComparerInt)

		resultWithZero := orderedWithZero.LastOrDefault(defaultValue)

		if resultWithZero != 2 {
			t.Errorf("Expected last value 2, got %d", resultWithZero)
		}

		emptyEnumerator := FromSlice([]int{})
		zeroDefault := 0

		orderedEmpty := emptyEnumerator.OrderBy(comparer.ComparerInt)

		resultEmpty := orderedEmpty.LastOrDefault(zeroDefault)

		if resultEmpty != 0 {
			t.Errorf("Expected zero default value, got %d", resultEmpty)
		}
	})

	t.Run("order enumerator any last or default with single element", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Data  []string
		}

		items := []Item{{Value: 42, Data: []string{"test"}}}
		var enumerator = FromSliceAny(items)
		defaultValue := Item{Value: 0, Data: []string{}}

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result := ordered.LastOrDefault(defaultValue)

		if result.Value != 42 {
			t.Errorf("Expected value 42, got %+v", result)
		}
	})

	t.Run("order enumerator last or default with duplicate maximum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.LastOrDefault(defaultValue)

		if result != 3 {
			t.Errorf("Expected maximum value 3, got %d", result)
		}
	})

	t.Run("order enumerator last or default preserves stability", func(t *testing.T) {
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
		defaultValue := Item{Value: 0, Index: 0}

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result := ordered.LastOrDefault(defaultValue)

		if result.Value != 2 || result.Index != 3 {
			t.Errorf("Expected {Value: 2, Index: 3}, got %+v", result)
		}
	})

	t.Run("order enumerator last or default with multiple then by levels", func(t *testing.T) {
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
		defaultValue := Record{Category: "Default", Type: "Default", Name: "Default"}

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return compareStrings(a.Type, b.Type)
		}).ThenBy(func(a, b Record) int {
			return compareStrings(a.Name, b.Name)
		})

		result := ordered.LastOrDefault(defaultValue)

		expected := Record{Category: "B", Type: "Y", Name: "Third"}
		if result != expected {
			t.Errorf("Expected last record to be %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator last or default with struct zero value default", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Value int
			Text  string
		}

		emptyEnumerator := FromSlice([]Data{})
		zeroDefault := Data{}

		ordered := emptyEnumerator.OrderBy(func(a, b Data) int {
			return a.Value - b.Value
		})

		result := ordered.LastOrDefault(zeroDefault)

		if result.Value != 0 || result.Text != "" {
			t.Errorf("Expected zero value struct, got %+v", result)
		}
	})

	t.Run("order enumerator any last or default with boolean values", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]bool{true, false, true})
		defaultValue := false

		ordered := enumerator.OrderBy(comparer.ComparerBool)

		result := ordered.LastOrDefault(defaultValue)

		if result != true {
			t.Errorf("Expected true (maximum), got %v", result)
		}
	})

	t.Run("order enumerator last or default with floating point values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{3.14, 1.41, 2.71, 0.57})
		defaultValue := -1.0

		ordered := enumerator.OrderBy(comparer.ComparerFloat64)

		result := ordered.LastOrDefault(defaultValue)

		if result != 3.14 {
			t.Errorf("Expected maximum value 3.14, got %f", result)
		}
	})

	t.Run("order enumerator last or default with reverse sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})
		defaultValue := -1

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result := ordered.LastOrDefault(defaultValue)

		if result != 1 {
			t.Errorf("Expected result 1 (minimum in reverse order), got %d", result)
		}
	})
}

func BenchmarkOrderEnumeratorLastOrDefault(b *testing.B) {
	b.Run("order enumerator last or default small", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		defaultValue := -1

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result != 99 {
				b.Fatalf("Expected 99, got %d", result)
			}
		}
	})

	b.Run("order enumerator any last or default medium", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}

		people := make([]Person, 1000)
		for i := 0; i < 1000; i++ {
			people[i] = Person{Name: fmt.Sprintf("Person%d", i), Age: i}
		}
		var enumerator = FromSliceAny(people)
		defaultValue := Person{Name: "Default", Age: -1}

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result.Age != 999 {
				b.Fatalf("Expected person with age 999, got %+v", result)
			}
		}
	})

	b.Run("order enumerator last or default with multiple levels", func(b *testing.B) {
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
		defaultValue := Record{Category: "Default", Value: -1}

		ordered := enumerator.OrderBy(func(a, b Record) int { return compareStrings(a.Category, b.Category) }).
			ThenBy(func(a, b Record) int { return a.Value - b.Value })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result.Category == "Default" {
				b.Fatal("Expected non-default record")
			}
			_ = result
		}
	})

	b.Run("order enumerator last or default empty", func(b *testing.B) {
		enumerator := FromSlice([]int{})
		defaultValue := 0

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result != 0 {
				b.Fatalf("Expected default value 0, got %d", result)
			}
		}
	})

	b.Run("order enumerator any last or default descending", func(b *testing.B) {
		items := make([]int, 200)
		for i := 0; i < 200; i++ {
			items[i] = i
		}
		var enumerator = FromSliceAny(items)
		defaultValue := -1

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result != 0 {
				b.Fatalf("Expected 0, got %d", result)
			}
		}
	})

	b.Run("order enumerator last or default with nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil
		defaultValue := 42

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result != 42 {
				b.Fatalf("Expected default value 42, got %d", result)
			}
		}
	})

	b.Run("order enumerator last or default with struct default", func(b *testing.B) {
		type Config struct {
			Name     string
			Priority int
		}

		configs := make([]Config, 100)
		for i := 0; i < 100; i++ {
			configs[i] = Config{Name: fmt.Sprintf("Config%d", i), Priority: i}
		}
		enumerator := FromSlice(configs)
		defaultValue := Config{Name: "Default", Priority: -1}

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result.Priority != 99 {
				b.Fatalf("Expected config with priority 99, got %+v", result)
			}
		}
	})

	b.Run("order enumerator last or default with duplicate values", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i % 100
		}
		enumerator := FromSlice(items)
		defaultValue := -1

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result != 99 {
				b.Fatalf("Expected 99, got %d", result)
			}
		}
	})

	b.Run("order enumerator last or default with floating point", func(b *testing.B) {
		items := make([]float64, 100)
		for i := 0; i < 100; i++ {
			items[i] = float64(i) * 0.1
		}
		enumerator := FromSlice(items)
		defaultValue := -1.0

		ordered := enumerator.OrderBy(func(a, b float64) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := ordered.LastOrDefault(defaultValue)
			if result != 9.9 {
				b.Fatalf("Expected 9.9, got %f", result)
			}
		}
	})
}

func BenchmarkLastOrDefault(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.LastOrDefault(-1)
		}
	})
}
