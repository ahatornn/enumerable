package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestToMap(t *testing.T) {
	t.Run("convert non-empty slice to map", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{1, 2, 3, 4, 5}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}
	})

	t.Run("convert slice with duplicates to map", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4, 4, 5})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{1, 2, 3, 4, 5}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d (unique elements), got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}

		if len(set) != 5 {
			t.Errorf("Expected 5 unique elements, got %d", len(set))
		}
	})

	t.Run("convert single element to map", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 1 {
			t.Fatalf("Expected length 1, got %d", len(set))
		}

		if _, exists := set[42]; !exists {
			t.Errorf("Expected key 42 to exist in map")
		}
	})

	t.Run("convert empty slice to map", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected empty map, got nil")
		}

		if len(set) != 0 {
			t.Errorf("Expected empty map, got length %d", len(set))
		}
	})

	t.Run("convert nil enumerator to map", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		set := enumerator.ToMap()

		if set == nil {
			t.Errorf("Expected empty map for nil enumerator, got nil")
		}

		if len(set) != 0 {
			t.Errorf("Expected empty map, got map with length %d", len(set))
		}

		set[42] = struct{}{}
		if len(set) != 1 {
			t.Errorf("Expected to be able to add element to returned map")
		}
	})

	t.Run("convert string slice to map", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go", "hello"})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []string{"hello", "world", "go"}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %s to exist in map", key)
			}
		}
	})
}

func TestToMapStruct(t *testing.T) {
	t.Run("convert struct slice to map", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Alice", Age: 30},
			{Name: "Charlie", Age: 35},
		}

		enumerator := FromSlice(people)
		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 3 {
			t.Errorf("Expected 3 unique elements, got %d", len(set))
		}

		alice := Person{Name: "Alice", Age: 30}
		bob := Person{Name: "Bob", Age: 25}
		charlie := Person{Name: "Charlie", Age: 35}

		if _, exists := set[alice]; !exists {
			t.Error("Expected Alice to be in map")
		}

		if _, exists := set[bob]; !exists {
			t.Error("Expected Bob to be in map")
		}

		if _, exists := set[charlie]; !exists {
			t.Error("Expected Charlie to be in map")
		}
	})
}

func TestToMapBoolean(t *testing.T) {
	t.Run("convert boolean slice to map", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 2 {
			t.Errorf("Expected 2 unique elements, got %d", len(set))
		}

		if _, exists := set[true]; !exists {
			t.Error("Expected true to be in map")
		}

		if _, exists := set[false]; !exists {
			t.Error("Expected false to be in map")
		}
	})
}

func TestToMapWithOperations(t *testing.T) {
	t.Run("to map after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		set := filtered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{2, 4, 6}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}
	})

	t.Run("to map after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		set := distinct.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{1, 2, 3, 4}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}
	})
}

func TestToMapEdgeCases(t *testing.T) {
	t.Run("to map with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 1 {
			t.Errorf("Expected 1 unique element, got %d", len(set))
		}

		if _, exists := set[0]; !exists {
			t.Error("Expected 0 to be in map")
		}
	})

	t.Run("to map with empty strings", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"", "", "hello", ""})

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 2 {
			t.Errorf("Expected 2 unique elements, got %d", len(set))
		}

		if _, exists := set[""]; !exists {
			t.Error("Expected empty string to be in map")
		}

		if _, exists := set["hello"]; !exists {
			t.Error("Expected 'hello' to be in map")
		}
	})

	t.Run("to map with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 5)

		set := enumerator.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 1 {
			t.Errorf("Expected 1 unique element, got %d", len(set))
		}

		if _, exists := set["test"]; !exists {
			t.Error("Expected 'test' to be in map")
		}
	})
}

func TestOrderEnumeratorToMap(t *testing.T) {
	t.Run("order enumerator to map with sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3, 5, 2})
		ordered := enumerator.OrderBy(comparer.ComparerInt)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{1, 2, 3, 5, 8, 9}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}

		if len(set) != 6 {
			t.Errorf("Expected 6 unique elements, got %d", len(set))
		}
	})

	t.Run("order enumerator to map with string sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"zebra", "apple", "banana", "apple", "cherry", "banana"})

		ordered := enumerator.OrderBy(comparer.ComparerString)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []string{"apple", "banana", "cherry", "zebra"}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %s to exist in map", key)
			}
		}
	})

	t.Run("order enumerator to map with multiple sorting levels", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Value    int
			Name     string
		}

		records := []Record{
			{Category: "B", Value: 10, Name: "Second"},
			{Category: "A", Value: 20, Name: "First"},
			{Category: "B", Value: 30, Name: "Fourth"},
			{Category: "A", Value: 15, Name: "Third"},
			{Category: "A", Value: 20, Name: "Duplicate"},
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return a.Value - b.Value
		})
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 5 {
			t.Errorf("Expected 5 unique records, got %d", len(set))
		}

		for record := range set {
			found := false
			for _, expected := range records {
				if record == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Unexpected record in map: %+v", record)
			}
		}
	})

	t.Run("order enumerator to map preserves stability", func(t *testing.T) {
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
			{Value: 1, Index: 2},
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 4 {
			t.Errorf("Expected 4 unique items, got %d", len(set))
		}

		expectedItems := []Item{
			{Value: 1, Index: 2},
			{Value: 1, Index: 4},
			{Value: 2, Index: 1},
			{Value: 2, Index: 3},
		}

		for _, expected := range expectedItems {
			if _, exists := set[expected]; !exists {
				t.Errorf("Expected item %+v to exist in map", expected)
			}
		}
	})

	t.Run("order enumerator to map with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8, 5, 3})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{9, 8, 5, 3, 2, 1}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}
	})

	t.Run("order enumerator to map with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2, 4, 4, 1, 2, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		expectedKeys := []int{1, 2, 3, 4}
		if len(set) != len(expectedKeys) {
			t.Fatalf("Expected length %d, got %d", len(expectedKeys), len(set))
		}

		for _, key := range expectedKeys {
			if _, exists := set[key]; !exists {
				t.Errorf("Expected key %d to exist in map", key)
			}
		}

		if len(set) != 4 {
			t.Errorf("Expected 4 unique elements, got %d", len(set))
		}
	})

	t.Run("order enumerator to map with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected empty map, got nil")
		}

		if len(set) != 0 {
			t.Errorf("Expected empty map, got length %d", len(set))
		}
	})

	t.Run("order enumerator to map with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 1 {
			t.Fatalf("Expected length 1, got %d", len(set))
		}

		if _, exists := set[42]; !exists {
			t.Errorf("Expected key 42 to exist in map")
		}
	})

	t.Run("order enumerator to map with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected empty map, got nil")
		}

		if len(set) != 0 {
			t.Errorf("Expected empty map from nil enumerator, got length %d", len(set))
		}
	})

	t.Run("order enumerator to map membership testing", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		set := ordered.ToMap()

		existingElements := []int{1, 2, 3, 5, 8, 9}
		for _, element := range existingElements {
			if _, exists := set[element]; !exists {
				t.Errorf("Expected element %d to exist in map", element)
			}
		}

		nonExistingElements := []int{0, 4, 6, 7, 10}
		for _, element := range nonExistingElements {
			if _, exists := set[element]; exists {
				t.Errorf("Expected element %d to not exist in map", element)
			}
		}
	})

	t.Run("order enumerator to map with struct keys", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Alice", Age: 25},
			{Name: "Charlie", Age: 35},
			{Name: "Bob", Age: 30},
		}
		enumerator := FromSlice(people)

		ordered := enumerator.OrderBy(func(a, b Person) int {
			nameCompare := compareStrings(a.Name, b.Name)
			if nameCompare != 0 {
				return nameCompare
			}
			return a.Age - b.Age
		})
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 3 {
			t.Errorf("Expected 3 unique people, got %d", len(set))
		}

		expectedPeople := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
			{Name: "Charlie", Age: 35},
		}

		for _, expected := range expectedPeople {
			if _, exists := set[expected]; !exists {
				t.Errorf("Expected person %+v to exist in map", expected)
			}
		}
	})

	t.Run("order enumerator to map with complex sorting and deduplication", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Value    int
			Name     string
		}

		records := []Record{
			{Category: "A", Value: 10, Name: "First"},
			{Category: "A", Value: 20, Name: "Second"},
			{Category: "B", Value: 10, Name: "Third"},
			{Category: "A", Value: 10, Name: "First"},
			{Category: "B", Value: 20, Name: "Fourth"},
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			categoryCompare := compareStrings(a.Category, b.Category)
			if categoryCompare != 0 {
				return categoryCompare
			}
			valueCompare := a.Value - b.Value
			if valueCompare != 0 {
				return valueCompare
			}
			return compareStrings(a.Name, b.Name)
		})
		set := ordered.ToMap()

		if set == nil {
			t.Fatal("Expected map, got nil")
		}

		if len(set) != 4 {
			t.Errorf("Expected 4 unique records, got %d", len(set))
		}

		for record := range set {
			count := 0
			for _, original := range records {
				if record == original {
					count++
				}
			}
			if count == 0 {
				t.Errorf("Unexpected record in map: %+v", record)
			}
		}
	})
}

func BenchmarkToMap(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToMap()
		}
	})

	b.Run("medium enumeration with duplicates", func(b *testing.B) {
		items := make([]int, 2000)
		for i := 0; i < 2000; i++ {
			items[i] = i % 1000 // Создаем дубликаты
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToMap()
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToMap()
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		for i := 0; i < b.N; i++ {
			_ = enumerator.ToMap()
		}
	})
}
