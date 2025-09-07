package enumerable

import (
	"errors"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestSingle(t *testing.T) {
	t.Run("single element from slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("single string element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello"})

		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "hello" {
			t.Errorf("Expected result 'hello', got %s", result)
		}
	})

	t.Run("single boolean element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true})

		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != true {
			t.Errorf("Expected result true, got %v", result)
		}
	})

	t.Run("single element from empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for empty slice")
		}
		if err != nil && err.Error() != "sequence contains no elements" {
			t.Errorf("Expected 'sequence contains no elements' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("single element from nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for nil enumerator")
		}
		if err != nil && err.Error() != "sequence contains no elements" {
			t.Errorf("Expected 'sequence contains no elements' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("single element from multiple elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("single element from two identical elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42, 42})

		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("single element with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "Item", Price: 100},
		}

		enumerator := FromSlice(products)
		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := Product{Name: "Item", Price: 100}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("single element from multiple struct elements", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "Item1", Price: 100},
			{Name: "Item2", Price: 200},
		}

		enumerator := FromSlice(products)
		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var expected Product
		if result != expected {
			t.Errorf("Expected zero value %+v, got %+v", expected, result)
		}
	})
}

func TestSingleOrDefault(t *testing.T) {
	t.Run("single element from slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})
		defaultValue := -1

		result := enumerator.SingleOrDefault(defaultValue)

		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("empty slice returns default value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		defaultValue := -1

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("nil enumerator returns default value", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil
		defaultValue := -1

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("multiple elements return default value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})
		defaultValue := -1

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("two identical elements return default value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42, 42})
		defaultValue := -1

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("empty slice with zero default value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		defaultValue := 0

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("multiple elements with zero default value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2})
		defaultValue := 0

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("struct with default value on empty", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		enumerator := FromSlice([]User{})
		defaultValue := User{ID: -1, Name: "default"}

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %+v, got %+v", defaultValue, result)
		}
	})

	t.Run("struct with default value on multiple elements", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		}
		enumerator := FromSlice(users)
		defaultValue := User{ID: -1, Name: "default"}

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %+v, got %+v", defaultValue, result)
		}
	})

	t.Run("string slice empty returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{})
		defaultValue := "not found"

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %s, got %s", defaultValue, result)
		}
	})

	t.Run("string slice multiple returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world"})
		defaultValue := "not found"

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %s, got %s", defaultValue, result)
		}
	})

	t.Run("boolean slice empty returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{})
		defaultValue := true

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %v, got %v", defaultValue, result)
		}
	})

	t.Run("boolean slice multiple returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})
		defaultValue := false

		result := enumerator.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %v, got %v", defaultValue, result)
		}
	})
}

func TestSingleStruct(t *testing.T) {
	t.Run("single element from struct field", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		users := []User{
			{ID: 1, Name: "Alice"},
		}

		enumerator := FromSlice(users)
		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := User{ID: 1, Name: "Alice"}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("single element from struct with comparable fields", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name    string
			Enabled bool
			Count   int
		}

		configs := []Config{
			{Name: "Config1", Enabled: true, Count: 5},
		}

		enumerator := FromSlice(configs)
		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := Config{Name: "Config1", Enabled: true, Count: 5}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})
}

func TestSingleEdgeCases(t *testing.T) {
	t.Run("single element with zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0})

		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected result 0, got %d", result)
		}
	})

	t.Run("single element with maximum value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{9223372036854775807}) // max int64

		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 9223372036854775807 {
			t.Errorf("Expected result 9223372036854775807, got %d", result)
		}
	})

	t.Run("single element with minimum value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-9223372036854775808}) // min int64

		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != -9223372036854775808 {
			t.Errorf("Expected result -9223372036854775808, got %d", result)
		}
	})

	t.Run("single element with duplicate zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0})

		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})
}

func TestSingleWithOperations(t *testing.T) {
	t.Run("single element after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})
		filtered := enumerator.Where(func(n int) bool { return n == 3 })

		result, err := filtered.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 3 {
			t.Errorf("Expected result 3, got %d", result)
		}
	})

	t.Run("single element after filter with no matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})
		filtered := enumerator.Where(func(n int) bool { return n > 10 })

		result, err := filtered.Single()

		if err == nil {
			t.Error("Expected error for empty filtered sequence")
		}
		if err != nil && err.Error() != "sequence contains no elements" {
			t.Errorf("Expected 'sequence contains no elements' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("single element after filter with multiple matches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 2, 5})
		filtered := enumerator.Where(func(n int) bool { return n == 2 })

		result, err := filtered.Single()

		if err == nil {
			t.Error("Expected error for multiple filtered elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("single element after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})
		taken := enumerator.Take(1)

		result, err := taken.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 1 {
			t.Errorf("Expected result 1, got %d", result)
		}
	})

	t.Run("single element after skip and take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})
		skipped := enumerator.Skip(2).Take(1)

		result, err := skipped.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 3 {
			t.Errorf("Expected result 3, got %d", result)
		}
	})
}

func TestSingleOrDefaultWithOperations(t *testing.T) {
	t.Run("empty filtered sequence returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})
		filtered := enumerator.Where(func(n int) bool { return n > 10 })
		defaultValue := -1

		result := filtered.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("multiple filtered elements return default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 2, 5})
		filtered := enumerator.Where(func(n int) bool { return n == 2 })
		defaultValue := -1

		result := filtered.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("successful filter returns single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})
		filtered := enumerator.Where(func(n int) bool { return n == 3 })
		defaultValue := -1

		result := filtered.SingleOrDefault(defaultValue)

		if result != 3 {
			t.Errorf("Expected result 3, got %d", result)
		}
	})
}

func TestSingleNonComparable(t *testing.T) {
	t.Run("single element with comparable struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name    string
			Enabled bool
		}

		configs := []Config{
			{Name: "Config1", Enabled: true},
		}

		enumerator := FromSlice(configs)
		result, err := enumerator.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := Config{Name: "Config1", Enabled: true}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("single element with comparable struct multiple elements", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name    string
			Enabled bool
		}

		configs := []Config{
			{Name: "Config1", Enabled: true},
			{Name: "Config2", Enabled: false},
		}

		enumerator := FromSlice(configs)
		result, err := enumerator.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var expected Config
		if result != expected {
			t.Errorf("Expected zero value %+v, got %+v", expected, result)
		}
	})
}

func TestSingleErrorMessages(t *testing.T) {
	t.Run("empty sequence error message", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		_, err := enumerator.Single()

		if err == nil {
			t.Fatal("Expected error")
		}
		expectedMsg := "sequence contains no elements"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("multiple elements error message", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2})

		_, err := enumerator.Single()

		if err == nil {
			t.Fatal("Expected error")
		}
		expectedMsg := "sequence contains more than one element"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("nil enumerator error message", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		_, err := enumerator.Single()

		if err == nil {
			t.Fatal("Expected error")
		}
		expectedMsg := "sequence contains no elements"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})
}

func TestSingleTypedErrors(t *testing.T) {
	t.Run("typed error for no elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		_, err := enumerator.Single()

		if err == nil {
			t.Fatal("Expected error")
		}

		if !errors.Is(err, ErrNoElements) {
			t.Errorf("Expected ErrNoElements, got %v", err)
		}

		if _, ok := err.(*NoElementsError); !ok {
			t.Errorf("Expected *NoElementsError, got %T", err)
		}

		expectedMsg := "sequence contains no elements"
		if err.Error() != expectedMsg {
			t.Errorf("Expected message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("typed error for multiple elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2})

		_, err := enumerator.Single()

		if err == nil {
			t.Fatal("Expected error")
		}

		if !errors.Is(err, ErrMultipleElements) {
			t.Errorf("Expected ErrMultipleElements, got %v", err)
		}

		if _, ok := err.(*MultipleElementsError); !ok {
			t.Errorf("Expected *MultipleElementsError, got %T", err)
		}

		expectedMsg := "sequence contains more than one element"
		if err.Error() != expectedMsg {
			t.Errorf("Expected message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("error handling with switch", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		_, err := enumerator.Single()

		if err == nil {
			t.Fatal("Expected error")
		}

		switch err.(type) {
		case *NoElementsError:
		case *MultipleElementsError:
			t.Error("Expected NoElementsError, got MultipleElementsError")
		default:
			t.Errorf("Unexpected error type: %T", err)
		}
	})
}

func BenchmarkSingle(b *testing.B) {
	b.Run("single element success", func(b *testing.B) {
		items := []int{42}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, err := enumerator.Single()
			if err != nil || result != 42 {
				b.Fatalf("Expected 42, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("single element empty", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result, err := enumerator.Single()
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("single element multiple", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, err := enumerator.Single()
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("nil enumerator", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var enumerator Enumerator[int] = nil
			result, err := enumerator.Single()
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})
}

func TestOrderEnumeratorSingle(t *testing.T) {
	t.Run("order enumerator single with one element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("order enumerator single with multiple elements returns error", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator single with empty slice returns error", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.Single()

		if err == nil {
			t.Error("Expected error for empty slice")
		}
		if err != nil && err.Error() != "sequence contains no elements" {
			t.Errorf("Expected 'sequence contains no elements' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator single with struct and sorting", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{{Name: "Alice", Age: 25}}
		enumerator := FromSlice(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })

		result, err := ordered.Single()

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := Person{Name: "Alice", Age: 25}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator single with multiple elements after sorting", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
		}

		items := []Item{{Value: 1}, {Value: 2}, {Value: 3}}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result, err := ordered.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero Item
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})

	t.Run("order enumerator single with duplicate elements after sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.Single()

		if err == nil {
			t.Error("Expected error for multiple elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})
}

func TestOrderEnumeratorSingleOrDefault(t *testing.T) {
	t.Run("order enumerator single or default with one element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.SingleOrDefault(defaultValue)

		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("order enumerator single or default with multiple elements returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("order enumerator single or default with empty slice returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("order enumerator single or default with struct and sorting", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{{Name: "Alice", Age: 25}}
		enumerator := FromSlice(people)
		defaultValue := Person{Name: "Default", Age: 0}

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })

		result := ordered.SingleOrDefault(defaultValue)

		expected := Person{Name: "Alice", Age: 25}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator single or default with multiple elements after sorting returns default", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
		}

		items := []Item{{Value: 1}, {Value: 2}, {Value: 3}}
		enumerator := FromSlice(items)
		defaultValue := Item{Value: -1}

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		result := ordered.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %+v, got %+v", defaultValue, result)
		}
	})

	t.Run("order enumerator single or default with duplicate elements returns default", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5})
		defaultValue := -1

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result := ordered.SingleOrDefault(defaultValue)

		if result != defaultValue {
			t.Errorf("Expected default value %d, got %d", defaultValue, result)
		}
	})

	t.Run("order enumerator single or default with zero default value", func(t *testing.T) {
		t.Parallel()
		emptyEnumerator := FromSlice([]int{})
		zeroDefault := 0

		orderedEmpty := emptyEnumerator.OrderBy(comparer.ComparerInt)

		resultEmpty := orderedEmpty.SingleOrDefault(zeroDefault)

		if resultEmpty != 0 {
			t.Errorf("Expected zero default value, got %d", resultEmpty)
		}

		multipleEnumerator := FromSlice([]int{1, 2})
		orderedMultiple := multipleEnumerator.OrderBy(comparer.ComparerInt)

		resultMultiple := orderedMultiple.SingleOrDefault(zeroDefault)

		if resultMultiple != 0 {
			t.Errorf("Expected zero default value, got %d", resultMultiple)
		}
	})
}

func BenchmarkSingleOrDefault(b *testing.B) {
	b.Run("single element success", func(b *testing.B) {
		items := []int{42}
		defaultValue := -1

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result := enumerator.SingleOrDefault(defaultValue)
			if result != 42 {
				b.Fatalf("Expected 42, got %d", result)
			}
		}
	})

	b.Run("empty slice with default", func(b *testing.B) {
		defaultValue := -1

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result := enumerator.SingleOrDefault(defaultValue)
			if result != defaultValue {
				b.Fatalf("Expected %d, got %d", defaultValue, result)
			}
		}
	})

	b.Run("multiple elements with default", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		defaultValue := -1

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result := enumerator.SingleOrDefault(defaultValue)
			if result != defaultValue {
				b.Fatalf("Expected %d, got %d", defaultValue, result)
			}
		}
	})

	b.Run("nil enumerator with default", func(b *testing.B) {
		defaultValue := -1

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var enumerator Enumerator[int] = nil
			result := enumerator.SingleOrDefault(defaultValue)
			if result != defaultValue {
				b.Fatalf("Expected %d, got %d", defaultValue, result)
			}
		}
	})
}
