package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestToBatch(t *testing.T) {
	t.Run("enumerator to batch with exact batches", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})

		batches := enumerator.ToBatch(3)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		expectedBatch0 := []int{1, 2, 3}
		expectedBatch1 := []int{4, 5, 6}

		if len(batches[0]) != len(expectedBatch0) {
			t.Errorf("Expected batch 0 length %d, got %d", len(expectedBatch0), len(batches[0]))
		}
		for i, v := range expectedBatch0 {
			if batches[0][i] != v {
				t.Errorf("Expected %d at batch 0 index %d, got %d", v, i, batches[0][i])
			}
		}

		if len(batches[1]) != len(expectedBatch1) {
			t.Errorf("Expected batch 1 length %d, got %d", len(expectedBatch1), len(batches[1]))
		}
		for i, v := range expectedBatch1 {
			if batches[1][i] != v {
				t.Errorf("Expected %d at batch 1 index %d, got %d", v, i, batches[1][i])
			}
		}
	})

	t.Run("enumerator any to batch with partial last batch", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{"a", "b", "c", "d", "e"})

		batches := enumerator.ToBatch(2)

		if len(batches) != 3 {
			t.Fatalf("Expected 3 batches, got %d", len(batches))
		}

		expectedBatch0 := []string{"a", "b"}
		expectedBatch1 := []string{"c", "d"}
		expectedBatch2 := []string{"e"}

		if len(batches[0]) != len(expectedBatch0) {
			t.Errorf("Expected batch 0 length %d, got %d", len(expectedBatch0), len(batches[0]))
		}
		for i, v := range expectedBatch0 {
			if batches[0][i] != v {
				t.Errorf("Expected %s at batch 0 index %d, got %s", v, i, batches[0][i])
			}
		}

		if len(batches[1]) != len(expectedBatch1) {
			t.Errorf("Expected batch 1 length %d, got %d", len(expectedBatch1), len(batches[1]))
		}
		for i, v := range expectedBatch1 {
			if batches[1][i] != v {
				t.Errorf("Expected %s at batch 1 index %d, got %s", v, i, batches[1][i])
			}
		}

		if len(batches[2]) != len(expectedBatch2) {
			t.Errorf("Expected batch 2 length %d, got %d", len(expectedBatch2), len(batches[2]))
		}
		if batches[2][0] != expectedBatch2[0] {
			t.Errorf("Expected %s at batch 2 index 0, got %s", expectedBatch2[0], batches[2][0])
		}
	})

	t.Run("order enumerator to batch with sorting", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 2, 8, 1, 9, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		expectedBatch0 := []int{1, 2, 3}
		expectedBatch1 := []int{5, 8, 9}

		if len(batches[0]) != len(expectedBatch0) {
			t.Errorf("Expected batch 0 length %d, got %d", len(expectedBatch0), len(batches[0]))
		}
		for i, v := range expectedBatch0 {
			if batches[0][i] != v {
				t.Errorf("Expected %d at batch 0 index %d, got %d", v, i, batches[0][i])
			}
		}

		if len(batches[1]) != len(expectedBatch1) {
			t.Errorf("Expected batch 1 length %d, got %d", len(expectedBatch1), len(batches[1]))
		}
		for i, v := range expectedBatch1 {
			if batches[1][i] != v {
				t.Errorf("Expected %d at batch 1 index %d, got %d", v, i, batches[1][i])
			}
		}
	})

	t.Run("order enumerator any to batch with complex struct and sorting", func(t *testing.T) {
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
			{Name: "Eve", Age: 22},
		}
		var enumerator = FromSliceAny(people)

		ordered := enumerator.OrderBy(func(a, b Person) int { return a.Age - b.Age })
		batches := ordered.ToBatch(2)

		if len(batches) != 3 {
			t.Fatalf("Expected 3 batches, got %d", len(batches))
		}

		if len(batches[0]) != 2 {
			t.Errorf("Expected batch 0 length 2, got %d", len(batches[0]))
		}
		if batches[0][0].Name != "Eve" || batches[0][0].Age != 22 {
			t.Errorf("Expected {Eve,22} at batch 0 index 0, got %+v", batches[0][0])
		}
		if batches[0][1].Name != "Alice" || batches[0][1].Age != 25 {
			t.Errorf("Expected {Alice,25} at batch 0 index 1, got %+v", batches[0][1])
		}

		if len(batches[1]) != 2 {
			t.Errorf("Expected batch 1 length 2, got %d", len(batches[1]))
		}
		if batches[1][0].Name != "Diana" || batches[1][0].Age != 28 {
			t.Errorf("Expected {Diana,28} at batch 1 index 0, got %+v", batches[1][0])
		}
		if batches[1][1].Name != "Charlie" || batches[1][1].Age != 30 {
			t.Errorf("Expected {Charlie,30} at batch 1 index 1, got %+v", batches[1][1])
		}

		if len(batches[2]) != 1 {
			t.Errorf("Expected batch 2 length 1, got %d", len(batches[2]))
		}
		if batches[2][0].Name != "Bob" || batches[2][0].Age != 35 {
			t.Errorf("Expected {Bob,35} at batch 2 index 0, got %+v", batches[2][0])
		}
	})

	t.Run("order enumerator to batch with multiple sorting levels", func(t *testing.T) {
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
			{Category: "A", Value: 25, Name: "Fifth"},
		}
		enumerator := FromSlice(records)

		ordered := enumerator.OrderBy(func(a, b Record) int {
			return compareStrings(a.Category, b.Category)
		}).ThenBy(func(a, b Record) int {
			return a.Value - b.Value
		})
		batches := ordered.ToBatch(3)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		if len(batches[0]) != 3 {
			t.Errorf("Expected batch 0 length 3, got %d", len(batches[0]))
		}
		if batches[0][0].Category != "A" || batches[0][0].Value != 15 || batches[0][0].Name != "Third" {
			t.Errorf("Expected {A,15,Third} at batch 0 index 0, got %+v", batches[0][0])
		}
		if batches[0][1].Category != "A" || batches[0][1].Value != 20 || batches[0][1].Name != "First" {
			t.Errorf("Expected {A,20,First} at batch 0 index 1, got %+v", batches[0][1])
		}
		if batches[0][2].Category != "A" || batches[0][2].Value != 25 || batches[0][2].Name != "Fifth" {
			t.Errorf("Expected {A,25,Fifth} at batch 0 index 2, got %+v", batches[0][2])
		}

		if len(batches[1]) != 2 {
			t.Errorf("Expected batch 1 length 2, got %d", len(batches[1]))
		}
		if batches[1][0].Category != "B" || batches[1][0].Value != 10 || batches[1][0].Name != "Second" {
			t.Errorf("Expected {B,10,Second} at batch 1 index 0, got %+v", batches[1][0])
		}
		if batches[1][1].Category != "B" || batches[1][1].Value != 30 || batches[1][1].Name != "Fourth" {
			t.Errorf("Expected {B,30,Fourth} at batch 1 index 1, got %+v", batches[1][1])
		}
	})

	t.Run("order enumerator to batch with zero batch size", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(0)

		if len(batches) != 0 {
			t.Errorf("Expected empty batches for zero batch size, got length %d", len(batches))
		}
	})

	t.Run("order enumerator any to batch with negative batch size", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]int{1, 2, 3, 4, 5})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(-1)

		if len(batches) != 0 {
			t.Errorf("Expected empty batches for negative batch size, got length %d", len(batches))
		}
	})

	t.Run("order enumerator to batch with batch size larger than total elements", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(10)

		if len(batches) != 1 {
			t.Fatalf("Expected 1 batch, got %d", len(batches))
		}

		expected := []int{1, 2, 3}
		if len(batches[0]) != len(expected) {
			t.Errorf("Expected batch length %d, got %d", len(expected), len(batches[0]))
		}
		for i, v := range expected {
			if batches[0][i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, batches[0][i])
			}
		}
	})

	t.Run("order enumerator to batch with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if len(batches) != 0 {
			t.Errorf("Expected empty batches from empty slice, got length %d", len(batches))
		}
	})

	t.Run("order enumerator any to batch with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if len(batches) != 0 {
			t.Errorf("Expected empty batches from nil enumerator, got length %d", len(batches))
		}
	})

	t.Run("order enumerator to batch with single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(5)

		if len(batches) != 1 {
			t.Fatalf("Expected 1 batch, got %d", len(batches))
		}

		if len(batches[0]) != 1 {
			t.Errorf("Expected batch length 1, got %d", len(batches[0]))
		}
		if batches[0][0] != 42 {
			t.Errorf("Expected 42 at batch index 0, got %d", batches[0][0])
		}
	})

	t.Run("order enumerator to batch preserves stability", func(t *testing.T) {
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
		batches := ordered.ToBatch(2)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		if len(batches[0]) != 2 {
			t.Errorf("Expected batch 0 length 2, got %d", len(batches[0]))
		}
		if batches[0][0].Value != 1 || batches[0][0].Index != 2 {
			t.Errorf("Expected {1,2} at batch 0 index 0, got %+v", batches[0][0])
		}
		if batches[0][1].Value != 1 || batches[0][1].Index != 4 {
			t.Errorf("Expected {1,4} at batch 0 index 1, got %+v", batches[0][1])
		}

		if len(batches[1]) != 2 {
			t.Errorf("Expected batch 1 length 2, got %d", len(batches[1]))
		}
		if batches[1][0].Value != 2 || batches[1][0].Index != 1 {
			t.Errorf("Expected {2,1} at batch 1 index 0, got %+v", batches[1][0])
		}
		if batches[1][1].Value != 2 || batches[1][1].Index != 3 {
			t.Errorf("Expected {2,3} at batch 1 index 1, got %+v", batches[1][1])
		}
	})

	t.Run("order enumerator to batch with descending order", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 5, 3, 9, 2, 8})

		ordered := enumerator.OrderByDescending(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		expectedBatch0 := []int{9, 8, 5}
		expectedBatch1 := []int{3, 2, 1}

		if len(batches[0]) != len(expectedBatch0) {
			t.Errorf("Expected batch 0 length %d, got %d", len(expectedBatch0), len(batches[0]))
		}
		for i, v := range expectedBatch0 {
			if batches[0][i] != v {
				t.Errorf("Expected %d at batch 0 index %d, got %d", v, i, batches[0][i])
			}
		}

		if len(batches[1]) != len(expectedBatch1) {
			t.Errorf("Expected batch 1 length %d, got %d", len(expectedBatch1), len(batches[1]))
		}
		for i, v := range expectedBatch1 {
			if batches[1][i] != v {
				t.Errorf("Expected %d at batch 1 index %d, got %d", v, i, batches[1][i])
			}
		}
	})

	t.Run("order enumerator to batch with duplicate values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{3, 1, 3, 2, 1, 2, 4, 4})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(4)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		expectedBatch0 := []int{1, 1, 2, 2}
		expectedBatch1 := []int{3, 3, 4, 4}

		if len(batches[0]) != len(expectedBatch0) {
			t.Errorf("Expected batch 0 length %d, got %d", len(expectedBatch0), len(batches[0]))
		}
		for i, v := range expectedBatch0 {
			if batches[0][i] != v {
				t.Errorf("Expected %d at batch 0 index %d, got %d", v, i, batches[0][i])
			}
		}

		if len(batches[1]) != len(expectedBatch1) {
			t.Errorf("Expected batch 1 length %d, got %d", len(expectedBatch1), len(batches[1]))
		}
		for i, v := range expectedBatch1 {
			if batches[1][i] != v {
				t.Errorf("Expected %d at batch 1 index %d, got %d", v, i, batches[1][i])
			}
		}
	})

	t.Run("order enumerator any to batch with complex struct and custom sorting", func(t *testing.T) {
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
			{Name: "Critical", Priority: 0, Options: []string{"opt7", "opt8", "opt9"}},
		}
		var enumerator = FromSliceAny(configs)

		ordered := enumerator.OrderBy(func(a, b Config) int { return a.Priority - b.Priority })
		batches := ordered.ToBatch(3)

		if len(batches) != 2 {
			t.Fatalf("Expected 2 batches, got %d", len(batches))
		}

		// Check first batch
		if len(batches[0]) != 3 {
			t.Errorf("Expected batch 0 length 3, got %d", len(batches[0]))
		}
		if batches[0][0].Name != "Critical" || batches[0][0].Priority != 0 {
			t.Errorf("Expected {Critical,0} at batch 0 index 0, got %+v", batches[0][0])
		}
		if batches[0][1].Name != "High" || batches[0][1].Priority != 1 {
			t.Errorf("Expected {High,1} at batch 0 index 1, got %+v", batches[0][1])
		}
		if batches[0][2].Name != "Medium" || batches[0][2].Priority != 2 {
			t.Errorf("Expected {Medium,2} at batch 0 index 2, got %+v", batches[0][2])
		}

		if len(batches[1]) != 2 {
			t.Errorf("Expected batch 1 length 2, got %d", len(batches[1]))
		}
		if batches[1][0].Name != "Low" || batches[1][0].Priority != 3 {
			t.Errorf("Expected {Low,3} at batch 1 index 0, got %+v", batches[1][0])
		}
		if batches[1][1].Name != "VeryLow" || batches[1][1].Priority != 4 {
			t.Errorf("Expected {VeryLow,4} at batch 1 index 1, got %+v", batches[1][1])
		}
	})
}

func TestToBatchEdgeCases(t *testing.T) {
	t.Run("order enumerator to batch with empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches from empty slice, got length %d", len(batches))
		}
	})

	t.Run("order enumerator any to batch with empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator = FromSliceAny([]string{})

		ordered := enumerator.OrderBy(comparer.ComparerString)
		batches := ordered.ToBatch(3)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches from empty slice, got length %d", len(batches))
		}
	})

	t.Run("order enumerator to batch with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches from nil enumerator, got length %d", len(batches))
		}
	})

	t.Run("order enumerator any to batch with nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(3)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches from nil enumerator, got length %d", len(batches))
		}
	})

	t.Run("order enumerator to batch with zero batch size and empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(0)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches for zero batch size, got length %d", len(batches))
		}
	})

	t.Run("order enumerator any to batch with zero batch size and nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(0)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches for zero batch size with nil enumerator, got length %d", len(batches))
		}
	})

	t.Run("order enumerator to batch with negative batch size and empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(-1)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches for negative batch size, got length %d", len(batches))
		}
	})

	t.Run("order enumerator any to batch with negative batch size and nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		ordered := enumerator.OrderBy(comparer.ComparerInt)
		batches := ordered.ToBatch(-5)

		if batches == nil {
			t.Fatal("Expected empty slice, got nil")
		}
		if len(batches) != 0 {
			t.Errorf("Expected empty batches for negative batch size with nil enumerator, got length %d", len(batches))
		}
	})
}

func BenchmarkToBatch(b *testing.B) {
	b.Run("to batch small enumeration small batch size", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 10

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 10 {
				b.Fatalf("Expected 10 batches, got %d", len(result))
			}
			if len(result[0]) != 10 {
				b.Fatalf("Expected batch size 10, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch medium enumeration medium batch size", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 50

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 20 {
				b.Fatalf("Expected 20 batches, got %d", len(result))
			}
			if len(result[0]) != 50 {
				b.Fatalf("Expected batch size 50, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch large enumeration large batch size", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 1000

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 10 {
				b.Fatalf("Expected 10 batches, got %d", len(result))
			}
			if len(result[0]) != 1000 {
				b.Fatalf("Expected batch size 1000, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with batch size larger than total elements", func(b *testing.B) {
		items := make([]int, 50)
		for i := 0; i < 50; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 100

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 1 {
				b.Fatalf("Expected 1 batch, got %d", len(result))
			}
			if len(result[0]) != 50 {
				b.Fatalf("Expected batch size 50, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with batch size of 1", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 1

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 100 {
				b.Fatalf("Expected 100 batches, got %d", len(result))
			}
			if len(result[0]) != 1 {
				b.Fatalf("Expected batch size 1, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with zero batch size", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 0

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 0 {
				b.Fatalf("Expected 0 batches for zero batch size, got %d", len(result))
			}
		}
	})

	b.Run("to batch with negative batch size", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := -5

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 0 {
				b.Fatalf("Expected 0 batches for negative batch size, got %d", len(result))
			}
		}
	})

	b.Run("to batch with empty enumeration", func(b *testing.B) {
		enumerator := FromSlice([]int{})
		batchSize := 10

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 0 {
				b.Fatalf("Expected 0 batches from empty enumeration, got %d", len(result))
			}
		}
	})

	b.Run("to batch with nil enumerator", func(b *testing.B) {
		var enumerator Enumerator[int] = nil
		batchSize := 10

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 0 {
				b.Fatalf("Expected 0 batches from nil enumerator, got %d", len(result))
			}
		}
	})

	b.Run("to batch with single element", func(b *testing.B) {
		enumerator := FromSlice([]int{42})
		batchSize := 5

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 1 {
				b.Fatalf("Expected 1 batch, got %d", len(result))
			}
			if len(result[0]) != 1 {
				b.Fatalf("Expected batch size 1, got %d", len(result[0]))
			}
			if result[0][0] != 42 {
				b.Fatalf("Expected 42 in batch, got %d", result[0][0])
			}
		}
	})

	b.Run("to batch with duplicate elements", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i % 10
		}
		enumerator := FromSlice(items)
		batchSize := 100

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 10 {
				b.Fatalf("Expected 10 batches, got %d", len(result))
			}
			if len(result[0]) != 100 {
				b.Fatalf("Expected batch size 100, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with strings", func(b *testing.B) {
		items := make([]string, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = fmt.Sprintf("item_%d", i)
		}
		enumerator := FromSlice(items)
		batchSize := 50

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 20 {
				b.Fatalf("Expected 20 batches, got %d", len(result))
			}
			if len(result[0]) != 50 {
				b.Fatalf("Expected batch size 50, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with complex structs", func(b *testing.B) {
		type ComplexStruct struct {
			ID      int
			Name    string
			Data    []int
			Options []string
		}

		items := make([]ComplexStruct, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = ComplexStruct{
				ID:      i,
				Name:    fmt.Sprintf("Name_%d", i),
				Data:    []int{i, i + 1, i + 2},
				Options: []string{fmt.Sprintf("opt_%d_a", i), fmt.Sprintf("opt_%d_b", i)},
			}
		}
		enumerator := FromSliceAny(items)
		batchSize := 25

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 40 {
				b.Fatalf("Expected 40 batches, got %d", len(result))
			}
			if len(result[0]) != 25 {
				b.Fatalf("Expected batch size 25, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with very small batch size", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 2

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 50 {
				b.Fatalf("Expected 50 batches, got %d", len(result))
			}
			if len(result[0]) != 2 {
				b.Fatalf("Expected batch size 2, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with very large batch size", func(b *testing.B) {
		items := make([]int, 100)
		for i := 0; i < 100; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 10000

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 1 {
				b.Fatalf("Expected 1 batch, got %d", len(result))
			}
			if len(result[0]) != 100 {
				b.Fatalf("Expected batch size 100, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch with alternating pattern", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			if i%2 == 0 {
				items[i] = 0
			} else {
				items[i] = 1
			}
		}
		enumerator := FromSlice(items)
		batchSize := 50

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 20 {
				b.Fatalf("Expected 20 batches, got %d", len(result))
			}
			if len(result[0]) != 50 {
				b.Fatalf("Expected batch size 50, got %d", len(result[0]))
			}
		}
	})

	b.Run("to batch memory allocation test", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 100

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 100 {
				b.Fatalf("Expected 100 batches, got %d", len(result))
			}
		}
	})

	b.Run("to batch with sequential access pattern", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}
		enumerator := FromSlice(items)
		batchSize := 100

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.ToBatch(batchSize)
			if len(result) != 10 { // 1000/100 = 10 batches
				b.Fatalf("Expected 10 batches, got %d", len(result))
			}

			for batchIdx, batch := range result {
				for itemIdx, item := range batch {
					expected := batchIdx*batchSize + itemIdx
					if item != expected {
						b.Fatalf("Expected %d at batch[%d][%d], got %d", expected, batchIdx, itemIdx, item)
					}
				}
			}
		}
	})
}
