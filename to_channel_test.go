package enumerable

import (
	"testing"
	"time"

	"github.com/ahatornn/enumerable/comparer"
)

func TestToChannel(t *testing.T) {
	t.Run("convert non-empty slice to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("convert non-empty slice to channel for non-comparable", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		})

		ch := enumerator.ToChannel(2)

		var result [][]int
		for item := range ch {
			copy := make([]int, len(item))
			for i, v := range item {
				copy[i] = v
			}
			result = append(result, copy)
		}

		expected := [][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if len(result[i]) != len(v) {
				t.Errorf("Expected slice length %d at index %d, got %d", len(v), i, len(result[i]))
				continue
			}
			for j, val := range v {
				if result[i][j] != val {
					t.Errorf("Expected %d at index [%d][%d], got %d", val, i, j, result[i][j])
				}
			}
		}
	})

	t.Run("convert single element to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 1 {
			t.Fatalf("Expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("Expected 42, got %d", result[0])
		}
	})

	t.Run("convert empty slice to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(result))
		}
	})

	t.Run("convert nil enumerator to channel", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice from nil enumerator, got length %d", len(result))
		}
	})

	t.Run("convert string slice to channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "world", "go"})

		ch := enumerator.ToChannel(0)

		var result []string
		for item := range ch {
			result = append(result, item)
		}

		expected := []string{"hello", "world", "go"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})
}

func TestToChannelBuffered(t *testing.T) {
	t.Run("buffered channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ch := enumerator.ToChannel(3) // Buffer size 3

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func TestToChannelStruct(t *testing.T) {
	t.Run("convert struct slice to channel", func(t *testing.T) {
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
		ch := enumerator.ToChannel(0)

		var result []Person
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != len(people) {
			t.Fatalf("Expected length %d, got %d", len(people), len(result))
		}

		for i, v := range people {
			if result[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, result[i])
			}
		}
	})
}

func TestToChannelWithOperations(t *testing.T) {
	t.Run("to channel after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		ch := filtered.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{2, 4, 6}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to channel after distinct", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 2, 3, 3, 4})
		distinct := enumerator.Distinct()

		ch := distinct.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 4}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func TestToChannelConcurrent(t *testing.T) {
	t.Run("concurrent channel consumption", func(t *testing.T) {
		t.Parallel()
		enumerator := Range(1, 10)

		ch := enumerator.ToChannel(5)

		resultChan := make(chan []int)
		go func() {
			var result []int
			for item := range ch {
				result = append(result, item)
			}
			resultChan <- result
		}()

		select {
		case result := <-resultChan:
			expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			if len(result) != len(expected) {
				t.Fatalf("Expected length %d, got %d", len(expected), len(result))
			}

			for i, v := range expected {
				if result[i] != v {
					t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
				}
			}
		case <-time.After(1 * time.Second):
			t.Fatal("Timeout waiting for channel consumption")
		}
	})
}

func TestToChannelEdgeCases(t *testing.T) {
	t.Run("to channel with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0})

		ch := enumerator.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{0, 0, 0}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("to channel with repeat", func(t *testing.T) {
		t.Parallel()
		enumerator := Repeat("test", 3)

		ch := enumerator.ToChannel(0)

		var result []string
		for item := range ch {
			result = append(result, item)
		}

		expected := []string{"test", "test", "test"}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
			}
		}
	})
}

func TestOrderEnumeratorToChannel(t *testing.T) {
	t.Run("order enumerator to channel with sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 2, 3, 5, 8, 9}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator any to channel with complex struct", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{Name: "Charlie", Age: 30},
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 35},
			{Name: "Diana", Age: 28},
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })
		ch := ordered.ToChannel(2)

		var result []Person
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		expectedAges := []int{25, 28, 30, 35}
		for i, expectedAge := range expectedAges {
			if result[i].Age != expectedAge {
				t.Errorf("Expected age %d at index %d, got %d", expectedAge, i, result[i].Age)
			}
		}

		expectedNames := []string{"Alice", "Diana", "Charlie", "Bob"}
		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}
	})

	t.Run("order enumerator to channel with multiple sorting levels", func(t *testing.T) {
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
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return a.Value - b.Value
		})
		ch := ordered.ToChannel(1)

		var result []Record
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		if result[0].Category != "A" || result[0].Value != 15 || result[0].Name != "Third" {
			t.Errorf("Expected {A,15,Third}, got %+v", result[0])
		}

		if result[1].Category != "A" || result[1].Value != 20 || result[1].Name != "First" {
			t.Errorf("Expected {A,20,First}, got %+v", result[1])
		}

		if result[2].Category != "B" || result[2].Value != 10 || result[2].Name != "Second" {
			t.Errorf("Expected {B,10,Second}, got %+v", result[2])
		}

		if result[3].Category != "B" || result[3].Value != 30 || result[3].Name != "Fourth" {
			t.Errorf("Expected {B,30,Fourth}, got %+v", result[3])
		}
	})

	t.Run("order enumerator to channel with buffered channel", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 4, 1, 5, 9, 2, 6})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(5)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator any to channel with complex struct and custom sorting", func(t *testing.T) {
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
			{Name: "VeryLow", Priority: 4, Options: []string{"opt6"}},
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
		ch := ordered.ToChannel(3)

		var result []Config
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		expectedPriorities := []int{1, 2, 3, 4}
		for i, expectedPriority := range expectedPriorities {
			if result[i].Priority != expectedPriority {
				t.Errorf("Expected priority %d at index %d, got %d", expectedPriority, i, result[i].Priority)
			}
		}

		expectedNames := []string{"High", "Medium", "Low", "VeryLow"}
		for i, expectedName := range expectedNames {
			if result[i].Name != expectedName {
				t.Errorf("Expected name %s at index %d, got %s", expectedName, i, result[i].Name)
			}
		}
	})

	t.Run("order enumerator to channel preserves stability", func(t *testing.T) {
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
		ch := ordered.ToChannel(0)

		var result []Item
		for item := range ch {
			result = append(result, item)
		}
		if len(result) != 4 {
			t.Fatalf("Expected length 4, got %d", len(result))
		}

		if result[0].Value != 1 || result[0].Index != 2 {
			t.Errorf("Expected {1,2} at index 0, got %+v", result[0])
		}
		if result[1].Value != 1 || result[1].Index != 4 {
			t.Errorf("Expected {1,4} at index 1, got %+v", result[1])
		}
		if result[2].Value != 2 || result[2].Index != 1 {
			t.Errorf("Expected {2,1} at index 2, got %+v", result[2])
		}
		if result[3].Value != 2 || result[3].Index != 3 {
			t.Errorf("Expected {2,3} at index 3, got %+v", result[3])
		}
	})

	t.Run("order enumerator to channel with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		ch := ordered.ToChannel(2)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{9, 8, 5, 3, 2, 1}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator to channel with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2, 4, 4})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(4)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		expected := []int{1, 1, 2, 2, 3, 3, 4, 4}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})

	t.Run("order enumerator to channel with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice from empty enumerator, got length %d", len(result))
		}
	})

	t.Run("order enumerator to channel with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(1)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 1 {
			t.Fatalf("Expected length 1, got %d", len(result))
		}

		if result[0] != 42 {
			t.Errorf("Expected 42, got %d", result[0])
		}
	})

	t.Run("order enumerator to channel with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(0)

		var result []int
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice from nil enumerator, got length %d", len(result))
		}
	})

	t.Run("order enumerator any to channel with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = nil

		ordered := enumerator.OrderBy(comparer.ComparerString)
		ch := ordered.ToChannel(0)

		var result []string
		for item := range ch {
			result = append(result, item)
		}

		if len(result) != 0 {
			t.Errorf("Expected empty slice from nil enumerator, got length %d", len(result))
		}
	})

	t.Run("order enumerator to channel with concurrent reading", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3, 7, 4, 6})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		ch := ordered.ToChannel(3)
		done := make(chan bool)
		var result []int
		go func() {
			defer close(done)
			for item := range ch {
				result = append(result, item)
			}
		}()

		<-done
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i, v := range expected {
			if result[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
			}
		}
	})
}

func BenchmarkToChannel(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ch := enumerator.ToChannel(0)
			for range ch {
				// Consume all items
			}
		}
	})

	b.Run("medium enumeration buffered", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ch := enumerator.ToChannel(100)
			for range ch {
				// Consume all items
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)

		for i := 0; i < b.N; i++ {
			ch := enumerator.ToChannel(1000)
			for range ch {
				// Consume all items
			}
		}
	})
}
