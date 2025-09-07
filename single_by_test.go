package enumerable

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
	"github.com/ahatornn/enumerable/hashcode"
)

func TestSingleBy(t *testing.T) {
	t.Run("Enumerator[int] single element with default comparer", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] single element with default comparer", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[string] = FromSlice([]string{"hello"})
		comparer := comparer.Default[string]()

		result, err := enumerator.SingleBy(comparer)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != "hello" {
			t.Errorf("Expected result 'hello', got %s", result)
		}
	})

	t.Run("Enumerator[int] empty slice returns error", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

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

	t.Run("EnumeratorAny[int] nil enumerator returns error", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

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

	t.Run("Enumerator[int] multiple distinct elements return error", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

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

	t.Run("EnumeratorAny[int] duplicate elements (same according to comparer) return single element", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = FromSliceAny([]int{42, 42, 42})
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

		if err != nil {
			t.Errorf("Expected no error for duplicate elements, got %v", err)
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("Enumerator[struct] comparison by field", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 1, Name: "Alice Updated"},
			{ID: 1, Name: "Alice Final"},
		}
		enumerator := FromSlice(users)
		idComparer := comparer.ByField(func(u User) int { return u.ID })

		result, err := enumerator.SingleBy(idComparer)

		if err != nil {
			t.Errorf("Expected no error for users with same ID, got %v", err)
		}
		expected := User{ID: 1, Name: "Alice"}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("EnumeratorAny[struct] comparison by field with multiple distinct IDs", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		var enumerator Enumerator[User] = FromSlice([]User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		})
		idComparer := comparer.ByField(func(u User) int { return u.ID })

		result, err := enumerator.SingleBy(idComparer)

		if err == nil {
			t.Error("Expected error for multiple distinct IDs")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero User
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})

	t.Run("EnumeratorAny[struct] custom comparer with case-insensitive string comparison", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name string
			Age  int
		}

		var enumerator Enumerator[User] = FromSlice([]User{
			{Name: "Alice", Age: 25},
			{Name: "alice", Age: 25},
			{Name: "ALICE", Age: 25},
		})

		caseInsensitiveComparer := comparer.Custom(
			func(a, b User) bool {
				return strings.EqualFold(a.Name, b.Name) && a.Age == b.Age
			},
			func(u User) uint64 {
				return hashcode.Combine(strings.ToLower(u.Name), u.Age)
			},
		)

		result, err := enumerator.SingleBy(caseInsensitiveComparer)

		if err != nil {
			t.Errorf("Expected no error for case-insensitive matches, got %v", err)
		}
		expected := User{Name: "Alice", Age: 25}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("Enumerator[struct] composite comparer with multiple fields", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}

		people := []Person{
			{FirstName: "John", LastName: "Doe", Age: 30},
			{FirstName: "John", LastName: "Doe", Age: 30},
			{FirstName: "John", LastName: "Doe", Age: 30},
		}
		enumerator := FromSlice(people)

		firstNameComparer := comparer.ByField(func(p Person) string { return p.FirstName })
		lastNameComparer := comparer.ByField(func(p Person) string { return p.LastName })
		ageComparer := comparer.ByField(func(p Person) int { return p.Age })
		compositeComparer := comparer.Composite(firstNameComparer, lastNameComparer, ageComparer)

		result, err := enumerator.SingleBy(compositeComparer)

		if err != nil {
			t.Errorf("Expected no error for identical people, got %v", err)
		}
		expected := Person{FirstName: "John", LastName: "Doe", Age: 30}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("EnumeratorAny[struct] composite comparer with multiple distinct elements", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}

		var enumerator Enumerator[Person] = FromSlice([]Person{
			{FirstName: "John", LastName: "Doe", Age: 30},
			{FirstName: "Jane", LastName: "Smith", Age: 25},
			{FirstName: "Bob", LastName: "Johnson", Age: 35},
		})

		firstNameComparer := comparer.ByField(func(p Person) string { return p.FirstName })
		lastNameComparer := comparer.ByField(func(p Person) string { return p.LastName })
		ageComparer := comparer.ByField(func(p Person) int { return p.Age })
		compositeComparer := comparer.Composite(firstNameComparer, lastNameComparer, ageComparer)

		result, err := enumerator.SingleBy(compositeComparer)

		if err == nil {
			t.Error("Expected error for multiple distinct people")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero Person
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})

	t.Run("Enumerator[int] early termination with multiple distinct elements", func(t *testing.T) {
		t.Parallel()
		items := make([]int, 1000)
		items[0] = 1
		items[1] = 2
		for i := 2; i < 1000; i++ {
			items[i] = 1
		}

		enumerator := FromSlice(items)
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

		if err == nil {
			t.Error("Expected error for multiple distinct elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("EnumeratorAny[struct] hash collision handling", func(t *testing.T) {
		t.Parallel()
		type TestStruct struct {
			Value int
		}

		collidingComparer := comparer.Custom(
			func(a, b TestStruct) bool {
				return a.Value == b.Value
			},
			func(s TestStruct) uint64 {
				return 42
			},
		)

		var enumerator Enumerator[TestStruct] = FromSlice([]TestStruct{
			{Value: 1},
			{Value: 1},
			{Value: 1},
		})

		result, err := enumerator.SingleBy(collidingComparer)

		if err != nil {
			t.Errorf("Expected no error for equal elements with hash collisions, got %v", err)
		}
		expected := TestStruct{Value: 1}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("Enumerator[struct] hash collision with different elements", func(t *testing.T) {
		t.Parallel()
		type TestStruct struct {
			Value int
		}

		collidingComparer := comparer.Custom(
			func(a, b TestStruct) bool {
				return a.Value == b.Value
			},
			func(s TestStruct) uint64 {
				return 42
			},
		)

		enumerator := FromSlice([]TestStruct{
			{Value: 1},
			{Value: 2},
		})

		result, err := enumerator.SingleBy(collidingComparer)

		if err == nil {
			t.Error("Expected error for different elements with hash collisions")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero TestStruct
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})

	t.Run("Enumerator[int] zero value element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0})
		comparer := comparer.Default[int]()

		result, err := enumerator.SingleBy(comparer)

		if err != nil {
			t.Errorf("Expected no error for zero value element, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected result 0, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] empty string element", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[string] = FromSlice([]string{""})
		comparer := comparer.Default[string]()

		result, err := enumerator.SingleBy(comparer)

		if err != nil {
			t.Errorf("Expected no error for empty string element, got %v", err)
		}
		if result != "" {
			t.Errorf("Expected result '', got %s", result)
		}
	})
}

func TestOrderEnumeratorSingleBy(t *testing.T) {
	t.Run("order enumerator single by with one element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})
		eqComparer := comparer.Default[int]()

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.SingleBy(eqComparer)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("order enumerator any single by with one element", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name    string
			Options []string
		}

		configs := []Config{{Name: "TestConfig", Options: []string{"opt1"}}}
		var enumerator = FromSliceAny(configs)
		nameComparer := comparer.ByField(func(c Config) string { return c.Name })

		ordered := enumerator.OrderBy(func(a, b Config) int { return compareStrings(a.Name, b.Name) })

		result, err := ordered.SingleBy(nameComparer)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := Config{Name: "TestConfig", Options: []string{"opt1"}}
		if result.Name != expected.Name {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator single by with multiple distinct elements returns error", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})
		eqComparer := comparer.Default[int]()

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.SingleBy(eqComparer)

		if err == nil {
			t.Error("Expected error for multiple distinct elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator any single by with multiple distinct elements returns error", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
			Tags []string
		}

		users := []User{
			{ID: 1, Name: "Alice", Tags: []string{"dev"}},
			{ID: 2, Name: "Bob", Tags: []string{"qa"}},
			{ID: 3, Name: "Charlie", Tags: []string{"ops"}},
		}
		var enumerator = FromSliceAny(users)
		idComparer := comparer.ByField(func(u User) int { return u.ID })

		ordered := enumerator.OrderBy(func(a, b User) int { return a.ID - b.ID })

		_, err := ordered.SingleBy(idComparer)

		if err == nil {
			t.Error("Expected error for multiple distinct elements")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
	})

	t.Run("order enumerator single by with duplicate elements according to comparer", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Category string
			Name     string
		}

		items := []Item{
			{Category: "A", Name: "Item1"},
			{Category: "A", Name: "Item2"},
			{Category: "B", Name: "Item3"},
		}
		enumerator := FromSlice(items)
		categoryComparer := comparer.ByField(func(i Item) string { return i.Category })

		ordered := enumerator.OrderBy(func(a, b Item) int { return compareStrings(a.Category, b.Category) })

		result, err := ordered.SingleBy(categoryComparer)

		if err == nil {
			t.Error("Expected error for multiple distinct elements by category")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero Item
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})

	t.Run("order enumerator single by with empty slice returns error", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})
		eqComparer := comparer.Default[int]()

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, err := ordered.SingleBy(eqComparer)

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

	t.Run("order enumerator single by with custom comparer and sorting", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float64
		}

		products := []Product{
			{Name: "Laptop", Price: 1200.0},
			{Name: "Phone", Price: 800.0},
		}
		enumerator := FromSlice(products)

		priceRangeComparer := comparer.Custom(
			func(a, b Product) bool {
				rangeA := int(a.Price/100) * 100
				rangeB := int(b.Price/100) * 100
				return rangeA == rangeB
			},
			func(p Product) uint64 {
				rangeVal := int(p.Price/100) * 100
				return hashcode.Compute(rangeVal)
			},
		)

		ordered := enumerator.OrderBy(func(a, b Product) int {
			if a.Price < b.Price {
				return -1
			}
			if a.Price > b.Price {
				return 1
			}
			return 0
		})

		result, err := ordered.SingleBy(priceRangeComparer)

		if err == nil {
			t.Error("Expected error for products in same price range")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero Product
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})

	t.Run("order enumerator single by preserves sorting order in error case", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Priority int
			Name     string
		}

		records := []Record{
			{Priority: 1, Name: "High1"},
			{Priority: 1, Name: "High2"},
			{Priority: 2, Name: "Low1"},
		}
		enumerator := FromSlice(records)
		priorityComparer := comparer.ByField(func(r Record) int { return r.Priority })

		ordered := enumerator.OrderBy(func(a, b Record) int { return a.Priority - b.Priority })

		result, err := ordered.SingleBy(priorityComparer)

		if err == nil {
			t.Error("Expected error for records with same priority")
		}
		if err != nil && err.Error() != "sequence contains more than one element" {
			t.Errorf("Expected 'sequence contains more than one element' error, got %v", err)
		}
		var zero Record
		if result != zero {
			t.Errorf("Expected zero value %+v, got %+v", zero, result)
		}
	})
}

func BenchmarkSingleBy(b *testing.B) {
	b.Run("Enumerator[int] single element success", func(b *testing.B) {
		items := []int{42}
		enumerator := FromSlice(items)
		comparer := comparer.Default[int]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err != nil || result != 42 {
				b.Fatalf("Expected 42, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("EnumeratorAny[int] single element success", func(b *testing.B) {
		items := []int{42}
		var enumerator Enumerator[int] = FromSlice(items)
		comparer := comparer.Default[int]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err != nil || result != 42 {
				b.Fatalf("Expected 42, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("Enumerator[int] single element empty", func(b *testing.B) {
		enumerator := FromSlice([]int{})
		comparer := comparer.Default[int]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("EnumeratorAny[int] single element multiple distinct", func(b *testing.B) {
		var enumerator Enumerator[int] = FromSlice([]int{1, 2, 3, 4, 5})
		comparer := comparer.Default[int]()

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("Enumerator[int] nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil
		comparer := comparer.Default[int]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("EnumeratorAny[int] nil enumerator", func(b *testing.B) {
		var enumerator EnumeratorAny[int] = nil
		comparer := comparer.Default[int]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err == nil || result != 0 {
				b.Fatalf("Expected error and 0, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("Enumerator[int] duplicate elements", func(b *testing.B) {
		items := make([]int, 1000)
		for i := range items {
			items[i] = 42
		}
		enumerator := FromSlice(items)
		comparer := comparer.Default[int]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(comparer)
			if err != nil || result != 42 {
				b.Fatalf("Expected 42, got %d, err: %v", result, err)
			}
		}
	})

	b.Run("EnumeratorAny[struct] struct comparison by field", func(b *testing.B) {
		type User struct {
			ID   int
			Name string
		}

		users := make([]User, 100)
		for i := range users {
			users[i] = User{ID: 1, Name: fmt.Sprintf("User%d", i)}
		}
		var enumerator Enumerator[User] = FromSlice(users)
		idComparer := comparer.ByField(func(u User) int { return u.ID })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, err := enumerator.SingleBy(idComparer)
			if err != nil {
				b.Fatalf("Unexpected error: %v", err)
			}
			expected := User{ID: 1, Name: "User0"}
			if result != expected {
				b.Fatalf("Expected %+v, got %+v", expected, result)
			}
		}
	})
}
