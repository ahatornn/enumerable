package enumerable

import (
	"fmt"
	"testing"
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
