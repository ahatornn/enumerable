package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestElementAt(t *testing.T) {
	t.Run("Enumerator[int] element at valid index", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{10, 20, 30, 40, 50})

		result, ok := enumerator.ElementAt(2)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != 30 {
			t.Errorf("Expected result 30, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] element at valid index", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"a", "b", "c", "d"})

		result, ok := enumerator.ElementAt(1)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != "b" {
			t.Errorf("Expected result 'b', got %s", result)
		}
	})

	t.Run("Enumerator[int] element at index 0", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42, 43, 44})

		result, ok := enumerator.ElementAt(0)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("EnumeratorAny[int] element at last index", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = FromSliceAny([]int{1, 2, 3, 4, 5})

		result, ok := enumerator.ElementAt(4)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != 5 {
			t.Errorf("Expected result 5, got %d", result)
		}
	})

	t.Run("Enumerator[int] negative index", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		result, ok := enumerator.ElementAt(-1)

		if ok {
			t.Error("Expected negative index to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] negative index", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"a", "b"})

		result, ok := enumerator.ElementAt(-5)

		if ok {
			t.Error("Expected negative index to return false")
		}
		if result != "" {
			t.Errorf("Expected zero value '', got %s", result)
		}
	})

	t.Run("Enumerator[int] index out of bounds", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		result, ok := enumerator.ElementAt(5)

		if ok {
			t.Error("Expected out of bounds index to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] index out of bounds", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"hello", "world"})

		result, ok := enumerator.ElementAt(10)

		if ok {
			t.Error("Expected out of bounds index to return false")
		}
		if result != "" {
			t.Errorf("Expected zero value '', got %s", result)
		}
	})

	t.Run("Enumerator[int] empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result, ok := enumerator.ElementAt(0)

		if ok {
			t.Error("Expected empty slice to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{})

		result, ok := enumerator.ElementAt(0)

		if ok {
			t.Error("Expected empty slice to return false")
		}
		if result != "" {
			t.Errorf("Expected zero value '', got %s", result)
		}
	})

	t.Run("Enumerator[int] nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result, ok := enumerator.ElementAt(0)

		if ok {
			t.Error("Expected nil enumerator to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("EnumeratorAny[int] nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		result, ok := enumerator.ElementAt(5)

		if ok {
			t.Error("Expected nil enumerator to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("Enumerator[bool] various indices", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true, false, true})

		if result, ok := enumerator.ElementAt(0); !ok || result != true {
			t.Errorf("Expected true at index 0, got %v, ok: %v", result, ok)
		}

		if result, ok := enumerator.ElementAt(2); !ok || result != true {
			t.Errorf("Expected true at index 2, got %v, ok: %v", result, ok)
		}

		if result, ok := enumerator.ElementAt(4); !ok || result != true {
			t.Errorf("Expected true at index 4, got %v, ok: %v", result, ok)
		}

		if result, ok := enumerator.ElementAt(5); ok {
			t.Errorf("Expected false for out of bounds, got %v, ok: %v", result, ok)
		}
	})

	t.Run("EnumeratorAny[struct] with comparable fields", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}
		var enumerator EnumeratorAny[User] = FromSliceAny(users)

		result, ok := enumerator.ElementAt(1)

		if !ok {
			t.Error("Expected user to be found")
		}
		expected := User{ID: 2, Name: "Bob"}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("EnumeratorAny[struct with slice] non-comparable type", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name    string
			Options []string
			Enabled bool
		}

		configs := []Config{
			{Name: "Config1", Options: []string{"opt1", "opt2"}, Enabled: true},
			{Name: "Config2", Options: []string{"opt3", "opt4"}, Enabled: false},
			{Name: "Config3", Options: []string{"opt5"}, Enabled: true},
		}
		var enumerator EnumeratorAny[Config] = FromSliceAny(configs)

		result, ok := enumerator.ElementAt(1)

		if !ok {
			t.Error("Expected config to be found")
		}
		expected := Config{Name: "Config2", Options: []string{"opt3", "opt4"}, Enabled: false}
		if result.Name != expected.Name || result.Enabled != expected.Enabled {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
		if len(result.Options) != len(expected.Options) {
			t.Errorf("Expected %d options, got %d", len(expected.Options), len(result.Options))
		}
	})

	t.Run("EnumeratorAny[struct with map] non-comparable type", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			ID   int
			Meta map[string]interface{}
			Tags []string
		}

		data := []Data{
			{
				ID:   1,
				Meta: map[string]interface{}{"key1": "value1"},
				Tags: []string{"tag1", "tag2"},
			},
			{
				ID:   2,
				Meta: map[string]interface{}{"key2": "value2"},
				Tags: []string{"tag3"},
			},
		}
		var enumerator EnumeratorAny[Data] = FromSliceAny(data)

		result, ok := enumerator.ElementAt(0)

		if !ok {
			t.Error("Expected data to be found")
		}
		if result.ID != 1 {
			t.Errorf("Expected ID 1, got %d", result.ID)
		}
		if len(result.Meta) != 1 || result.Meta["key1"] != "value1" {
			t.Errorf("Expected meta data not found")
		}
		if len(result.Tags) != 2 {
			t.Errorf("Expected 2 tags, got %d", len(result.Tags))
		}
	})

	t.Run("EnumeratorAny[struct with function] non-comparable type", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Name     string
			Callback func() error
			Active   bool
		}

		handlers := []Handler{
			{
				Name:     "Handler1",
				Callback: func() error { return nil },
				Active:   true,
			},
			{
				Name:     "Handler2",
				Callback: func() error { return fmt.Errorf("error") },
				Active:   false,
			},
		}
		var enumerator EnumeratorAny[Handler] = FromSliceAny(handlers)

		result, ok := enumerator.ElementAt(1)

		if !ok {
			t.Error("Expected handler to be found")
		}
		if result.Name != "Handler2" {
			t.Errorf("Expected Handler2, got %s", result.Name)
		}
		if result.Active != false {
			t.Errorf("Expected Active false, got %v", result.Active)
		}
		if result.Callback == nil {
			t.Error("Expected Callback to be not nil")
		}
	})

	t.Run("EnumeratorAny[slice of ints] non-comparable type", func(t *testing.T) {
		t.Parallel()
		type DataSet struct {
			Name   string
			Values []int
		}

		datasets := []DataSet{
			{Name: "Dataset1", Values: []int{1, 2, 3}},
			{Name: "Dataset2", Values: []int{4, 5, 6, 7}},
			{Name: "Dataset3", Values: []int{8, 9}},
		}
		var enumerator EnumeratorAny[DataSet] = FromSliceAny(datasets)

		result, ok := enumerator.ElementAt(2)

		if !ok {
			t.Error("Expected dataset to be found")
		}
		if result.Name != "Dataset3" {
			t.Errorf("Expected Dataset3, got %s", result.Name)
		}
		if len(result.Values) != 2 {
			t.Errorf("Expected 2 values, got %d", len(result.Values))
		}
		if result.Values[0] != 8 || result.Values[1] != 9 {
			t.Errorf("Expected values [8,9], got %v", result.Values)
		}
	})

	t.Run("EnumeratorAny[complex nested non-comparable] type", func(t *testing.T) {
		t.Parallel()
		type NestedStruct struct {
			Data []map[string][]int
			Name string
		}

		nested := []NestedStruct{
			{
				Name: "First",
				Data: []map[string][]int{
					{"group1": []int{1, 2}},
					{"group2": []int{3, 4, 5}},
				},
			},
			{
				Name: "Second",
				Data: []map[string][]int{
					{"group3": []int{6, 7, 8, 9}},
				},
			},
		}
		var enumerator EnumeratorAny[NestedStruct] = FromSliceAny(nested)

		result, ok := enumerator.ElementAt(0)

		if !ok {
			t.Error("Expected nested struct to be found")
		}
		if result.Name != "First" {
			t.Errorf("Expected First, got %s", result.Name)
		}
		if len(result.Data) != 2 {
			t.Errorf("Expected 2 data items, got %d", len(result.Data))
		}
	})

	t.Run("EnumeratorAny[struct] with zero values", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Text  string
			Flags []bool
		}

		items := []Item{
			{Value: 0, Text: "", Flags: []bool{}},
			{Value: 1, Text: "test", Flags: []bool{true}},
		}
		var enumerator EnumeratorAny[Item] = FromSliceAny(items)

		result, ok := enumerator.ElementAt(0)

		if !ok {
			t.Error("Expected item to be found")
		}
		if result.Value != 0 {
			t.Errorf("Expected Value 0, got %d", result.Value)
		}
		if result.Text != "" {
			t.Errorf("Expected Text '', got %s", result.Text)
		}
		if result.Flags == nil {
			t.Error("Expected Flags to be empty slice, not nil")
		}
	})

	t.Run("Enumerator[int] large index performance", func(t *testing.T) {
		t.Parallel()
		items := make([]int, 10000)
		items[5] = 42
		enumerator := FromSlice(items)

		result, ok := enumerator.ElementAt(5)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != 42 {
			t.Errorf("Expected result 42, got %d", result)
		}
	})

	t.Run("EnumeratorAny[string] sequential access", func(t *testing.T) {
		t.Parallel()
		letters := []string{"a", "b", "c", "d", "e", "f", "g"}
		var enumerator EnumeratorAny[string] = FromSliceAny(letters)

		for i, expected := range letters {
			result, ok := enumerator.ElementAt(i)
			if !ok {
				t.Errorf("Expected element at index %d to be found", i)
			}
			if result != expected {
				t.Errorf("Expected %s at index %d, got %s", expected, i, result)
			}
		}

		if result, ok := enumerator.ElementAt(len(letters)); ok {
			t.Errorf("Expected out of bounds to return false, got %s, ok: %v", result, ok)
		}
	})
}

func TestOrderEnumeratorElementAt(t *testing.T) {
	t.Run("order enumerator element at valid index", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		result, ok := ordered.ElementAt(2)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != 3 {
			t.Errorf("Expected result 3, got %d", result)
		}
	})

	t.Run("order enumerator any element at valid index", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })
		result, ok := ordered.ElementAt(1)

		if !ok {
			t.Error("Expected element to be found")
		}
		expected := Person{Name: "Alice", Age: 30}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator element at index 0", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, ok := ordered.ElementAt(0)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result != 1 {
			t.Errorf("Expected result 1, got %d", result)
		}
	})

	t.Run("order enumerator any element at last index", func(t *testing.T) {
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

		ordered := enumerator.OrderBy(func(a, b Product) int {
			if a.Price < b.Price {
				return -1
			}
			if a.Price > b.Price {
				return 1
			}
			return 0
		})

		result, ok := ordered.ElementAt(2)

		if !ok {
			t.Error("Expected element to be found")
		}
		expected := Product{Name: "Laptop", Price: 1200.0}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator element at negative index", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, ok := ordered.ElementAt(-1)

		if ok {
			t.Error("Expected negative index to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator any element at negative index", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			Data  []string
		}

		items := []Item{{Value: 1, Data: []string{"test"}}}
		var enumerator = FromSliceAny(items)

		ordered := enumerator.OrderBy(func(a, b Item) int { return a.Value - b.Value })

		_, ok := ordered.ElementAt(-5)

		if ok {
			t.Error("Expected negative index to return false")
		}
	})

	t.Run("order enumerator element at out of bounds index", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, ok := ordered.ElementAt(5)

		if ok {
			t.Error("Expected out of bounds index to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator any element at out of bounds index", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Name    string
			Options []string
		}

		configs := []Config{
			{Name: "Config1", Options: []string{"opt1"}},
			{Name: "Config2", Options: []string{"opt2"}},
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return compareStrings(a.Name, b.Name) })

		_, ok := ordered.ElementAt(10)

		if ok {
			t.Error("Expected out of bounds index to return false")
		}
	})

	t.Run("order enumerator element at index with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, ok := ordered.ElementAt(0)

		if ok {
			t.Error("Expected empty slice to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator any element at index with empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{})

		ordered := enumerator.OrderBy(comparer.ComparerString)

		result, ok := ordered.ElementAt(0)

		if ok {
			t.Error("Expected empty slice to return false")
		}
		if result != "" {
			t.Errorf("Expected zero value '', got %s", result)
		}
	})

	t.Run("order enumerator element at index with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, ok := ordered.ElementAt(0)

		if ok {
			t.Error("Expected nil enumerator to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator any element at index with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result, ok := ordered.ElementAt(5)

		if ok {
			t.Error("Expected nil enumerator to return false")
		}
		if result != 0 {
			t.Errorf("Expected zero value 0, got %d", result)
		}
	})

	t.Run("order enumerator element at index with multiple sorting levels", func(t *testing.T) {
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

		result, ok := ordered.ElementAt(2)

		if !ok {
			t.Error("Expected element to be found")
		}
		expected := Person{Name: "Bob", Age: 30}
		if result != expected {
			t.Errorf("Expected result %+v, got %+v", expected, result)
		}
	})

	t.Run("order enumerator any element at index with complex struct", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Category string
			Name     string
			Value    int
			Tags     []string
		}

		records := []Record{
			{Category: "B", Name: "Second", Value: 20, Tags: []string{"tag1"}},
			{Category: "A", Name: "First", Value: 10, Tags: []string{"tag2"}},
			{Category: "B", Name: "Third", Value: 30, Tags: []string{"tag3"}},
		}
		var enumerator = FromSliceAny(records)

		ordered := enumerator.OrderBy(func(a, b Record) int { return compareStrings(a.Category, b.Category) }).
			ThenBy(func(a, b Record) int { return a.Value - b.Value })

		result, ok := ordered.ElementAt(1)

		if !ok {
			t.Error("Expected element to be found")
		}
		if result.Category != "B" || result.Name != "Second" || result.Value != 20 {
			t.Errorf("Expected category B, name Second, value 20, got %+v", result)
		}
	})

	t.Run("order enumerator element at index with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2})

		ordered := enumerator.OrderBy(comparer.ComparerInt)

		result0, ok0 := ordered.ElementAt(0)
		result2, ok2 := ordered.ElementAt(2)
		result4, ok4 := ordered.ElementAt(4)

		if !ok0 || result0 != 1 {
			t.Errorf("Expected element at index 0 to be 1, got %d, ok: %v", result0, ok0)
		}
		if !ok2 || result2 != 2 {
			t.Errorf("Expected element at index 2 to be 2, got %d, ok: %v", result2, ok2)
		}
		if !ok4 || result4 != 3 {
			t.Errorf("Expected element at index 4 to be 3, got %d, ok: %v", result4, ok4)
		}
	})

	t.Run("order enumerator element at index preserves stability", func(t *testing.T) {
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

		result0, ok0 := ordered.ElementAt(0)
		result1, ok1 := ordered.ElementAt(1)

		if !ok0 || result0.Value != 1 || result0.Index != 2 {
			t.Errorf("Expected element at index 0 to be {Value: 1, Index: 2}, got %+v, ok: %v", result0, ok0)
		}
		if !ok1 || result1.Value != 1 || result1.Index != 4 {
			t.Errorf("Expected element at index 1 to be {Value: 1, Index: 4}, got %+v, ok: %v", result1, ok1)
		}
	})

	t.Run("order enumerator element at index with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)

		result0, ok0 := ordered.ElementAt(0)
		result2, ok2 := ordered.ElementAt(2)

		if !ok0 || result0 != 9 {
			t.Errorf("Expected element at index 0 to be 9, got %d, ok: %v", result0, ok0)
		}
		if !ok2 || result2 != 5 {
			t.Errorf("Expected element at index 2 to be 5, got %d, ok: %v", result2, ok2)
		}
	})
}

func BenchmarkOrderEnumeratorElementAt(b *testing.B) {
	b.Run("order enumerator element at beginning", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = 1000 - i
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := ordered.ElementAt(0)
			if !ok || result != 1 {
				b.Fatalf("Expected 1, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("order enumerator element at middle", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = 1000 - i
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := ordered.ElementAt(500)
			if !ok || result != 501 {
				b.Fatalf("Expected 501, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("order enumerator any element at end", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}

		people := make([]Person, 100)
		for i := 0; i < 100; i++ {
			people[i] = Person{Name: fmt.Sprintf("Person%d", i), Age: i}
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := ordered.ElementAt(99)
			if !ok || result.Age != 99 {
				b.Fatalf("Expected age 99, got %+v, ok: %v", result, ok)
			}
		}
	})

	b.Run("order enumerator element at out of bounds", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := ordered.ElementAt(200)
			if ok || result != 0 {
				b.Fatalf("Expected false and 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("order enumerator element at negative index", func(b *testing.B) {
		enumerator := FromSlice([]int{1, 2, 3})

		ordered := enumerator.OrderBy(func(a, b int) int { return a - b })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := ordered.ElementAt(-1)
			if ok || result != 0 {
				b.Fatalf("Expected false and 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("order enumerator with multiple levels element at index", func(b *testing.B) {
		type Record struct {
			Category string
			Value    int
		}

		records := make([]Record, 200)
		for i := 0; i < 200; i++ {
			records[i] = Record{
				Category: fmt.Sprintf("Cat%d", i%5),
				Value:    i,
			}
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int { return compareStrings(a.Category, b.Category) }).
			ThenBy(func(a, b Record) int { return a.Value - b.Value })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := ordered.ElementAt(100)
			if !ok {
				b.Fatalf("Expected element to be found, ok: %v", ok)
			}
			_ = result
		}
	})
}

func BenchmarkElementAt(b *testing.B) {
	b.Run("Enumerator[int] element at beginning", func(b *testing.B) {
		items := make([]int, 1000)
		items[0] = 42
		enumerator := FromSlice(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(0)
			if !ok || result != 42 {
				b.Fatalf("Expected 42, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("EnumeratorAny[int] element at middle", func(b *testing.B) {
		items := make([]int, 1000)
		items[500] = 123
		var enumerator EnumeratorAny[int] = FromSliceAny(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(500)
			if !ok || result != 123 {
				b.Fatalf("Expected 123, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("Enumerator[int] element at end", func(b *testing.B) {
		items := make([]int, 1000)
		items[999] = 999
		enumerator := FromSlice(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(999)
			if !ok || result != 999 {
				b.Fatalf("Expected 999, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("EnumeratorAny[int] out of bounds", func(b *testing.B) {
		items := make([]int, 100)
		var enumerator EnumeratorAny[int] = FromSliceAny(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(200)
			if ok || result != 0 {
				b.Fatalf("Expected false and 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("Enumerator[int] negative index", func(b *testing.B) {
		enumerator := FromSlice([]int{1, 2, 3})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(-1)
			if ok || result != 0 {
				b.Fatalf("Expected false and 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("EnumeratorAny[struct] non-comparable type", func(b *testing.B) {
		type Config struct {
			Name    string
			Options []string
			Enabled bool
		}

		configs := make([]Config, 100)
		for i := range configs {
			configs[i] = Config{
				Name:    fmt.Sprintf("Config%d", i),
				Options: []string{fmt.Sprintf("opt%d_1", i), fmt.Sprintf("opt%d_2", i)},
				Enabled: i%2 == 0,
			}
		}
		var enumerator EnumeratorAny[Config] = FromSliceAny(configs)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(50)
			if !ok || result.Name != "Config50" {
				b.Fatalf("Expected Config50, got %+v, ok: %v", result, ok)
			}
		}
	})

	b.Run("Enumerator[int] nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(0)
			if ok || result != 0 {
				b.Fatalf("Expected false and 0, got %d, ok: %v", result, ok)
			}
		}
	})

	b.Run("Enumerator[int] empty slice", func(b *testing.B) {
		var enumerator Enumerator[int] = FromSlice([]int{})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result, ok := enumerator.ElementAt(0)
			if ok || result != 0 {
				b.Fatalf("Expected false and 0, got %d, ok: %v", result, ok)
			}
		}
	})
}
