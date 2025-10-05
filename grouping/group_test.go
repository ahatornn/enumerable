package grouping

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestGroup(t *testing.T) {
	t.Run("group key retrieval", func(t *testing.T) {
		t.Parallel()
		group := Group[string, int]{
			key:   "test-key",
			items: []int{1, 2, 3},
		}

		key := group.Key()
		if key != "test-key" {
			t.Errorf("Expected key 'test-key', got %q", key)
		}
	})

	t.Run("group items retrieval", func(t *testing.T) {
		t.Parallel()
		expectedItems := []int{10, 20, 30}
		group := Group[string, int]{
			key:   "numbers",
			items: expectedItems,
		}

		items := group.Items()
		if len(items) != len(expectedItems) {
			t.Errorf("Expected %d items, got %d", len(expectedItems), len(items))
		}

		for i, expected := range expectedItems {
			if items[i] != expected {
				t.Errorf("Item %d: expected %d, got %d", i, expected, items[i])
			}
		}
	})

	t.Run("group with empty items", func(t *testing.T) {
		t.Parallel()
		group := Group[string, int]{
			key:   "empty",
			items: []int{},
		}

		if group.Key() != "empty" {
			t.Errorf("Expected key 'empty', got %q", group.Key())
		}

		items := group.Items()
		if len(items) != 0 {
			t.Errorf("Expected empty items slice, got length %d", len(items))
		}
	})

	t.Run("group with complex key", func(t *testing.T) {
		t.Parallel()
		type ComplexKey struct {
			ID   int
			Name string
		}
		key := ComplexKey{ID: 123, Name: "test"}
		group := Group[ComplexKey, string]{
			key:   key,
			items: []string{"a", "b"},
		}

		retrievedKey := group.Key()
		if retrievedKey.ID != key.ID || retrievedKey.Name != key.Name {
			t.Errorf("Expected key %+v, got %+v", key, retrievedKey)
		}

		items := group.Items()
		expectedItems := []string{"a", "b"}
		if len(items) != len(expectedItems) {
			t.Errorf("Expected %d items, got %d", len(expectedItems), len(items))
		} else {
			for i, expected := range expectedItems {
				if items[i] != expected {
					t.Errorf("Item %d: expected %q, got %q", i, expected, items[i])
				}
			}
		}
	})

	t.Run("group preserves item order", func(t *testing.T) {
		t.Parallel()
		originalOrder := []int{5, 3, 8, 1, 9}
		group := Group[string, int]{
			key:   "ordered",
			items: originalOrder,
		}

		items := group.Items()
		if len(items) != len(originalOrder) {
			t.Errorf("Expected %d items, got %d", len(originalOrder), len(items))
		} else {
			for i, expected := range originalOrder {
				if items[i] != expected {
					t.Errorf("Item %d: expected %d, got %d (order not preserved)", i, expected, items[i])
				}
			}
		}
	})
}

func TestNewGroupingBuilder(t *testing.T) {
	t.Run("create new grouping builder", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		if builder == nil {
			t.Fatal("Expected non-nil builder, got nil")
		}

		if builder.order == nil {
			t.Error("Expected non-nil order slice, got nil")
		}

		if builder.lookup == nil {
			t.Error("Expected non-nil lookup map, got nil")
		}

		if len(builder.order) != 0 {
			t.Errorf("Expected empty order slice, got length %d", len(builder.order))
		}

		if len(builder.lookup) != 0 {
			t.Errorf("Expected empty lookup map, got size %d", len(builder.lookup))
		}

		if builder.comparer == nil {
			t.Error("Expected non-nil comparer, got nil")
		}
	})

	t.Run("create builder with custom comparer", func(t *testing.T) {
		t.Parallel()
		customComparer := comparer.Default[string]()
		builder := NewGroupingBuilder[string, int](customComparer)

		if builder.comparer == nil {
			t.Error("Expected non-nil comparer, got nil")
		}
	})
}

func TestGroupingBuilder_Add(t *testing.T) {
	t.Run("add single key-value pair", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		builder.Add("key1", 100)

		result := builder.Result()
		if len(result) != 1 {
			t.Fatalf("Expected 1 group, got %d", len(result))
		}

		if result[0].Key() != "key1" {
			t.Errorf("Expected key 'key1', got %q", result[0].Key())
		}

		items := result[0].Items()
		if len(items) != 1 || items[0] != 100 {
			t.Errorf("Expected item [100], got %v", items)
		}
	})

	t.Run("add multiple values to same key", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		builder.Add("key1", 100)
		builder.Add("key1", 200)
		builder.Add("key1", 300)

		result := builder.Result()
		if len(result) != 1 {
			t.Fatalf("Expected 1 group, got %d", len(result))
		}

		if result[0].Key() != "key1" {
			t.Errorf("Expected key 'key1', got %q", result[0].Key())
		}

		items := result[0].Items()
		expected := []int{100, 200, 300}
		if len(items) != len(expected) {
			t.Errorf("Expected %d items, got %d", len(expected), len(items))
		} else {
			for i, expectedVal := range expected {
				if items[i] != expectedVal {
					t.Errorf("Item %d: expected %d, got %d", i, expectedVal, items[i])
				}
			}
		}
	})

	t.Run("add values to different keys", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		builder.Add("key1", 100)
		builder.Add("key2", 200)
		builder.Add("key1", 150)
		builder.Add("key3", 300)

		result := builder.Result()
		if len(result) != 3 {
			t.Fatalf("Expected 3 groups, got %d", len(result))
		}

		var group1, group2, group3 *Group[string, int]
		for i := range result {
			switch result[i].Key() {
			case "key1":
				g := result[i]
				group1 = &g
			case "key2":
				g := result[i]
				group2 = &g
			case "key3":
				g := result[i]
				group3 = &g
			}
		}

		if group1 == nil || group2 == nil || group3 == nil {
			t.Fatal("Missing expected groups")
		}

		key1Items := group1.Items()
		if len(key1Items) != 2 || key1Items[0] != 100 || key1Items[1] != 150 {
			t.Errorf("Expected key1 items [100, 150], got %v", key1Items)
		}

		key2Items := group2.Items()
		if len(key2Items) != 1 || key2Items[0] != 200 {
			t.Errorf("Expected key2 items [200], got %v", key2Items)
		}

		key3Items := group3.Items()
		if len(key3Items) != 1 || key3Items[0] != 300 {
			t.Errorf("Expected key3 items [300], got %v", key3Items)
		}
	})

	t.Run("add with integer keys", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[int, string](comparer.Default[int]())

		builder.Add(1, "one")
		builder.Add(2, "two")
		builder.Add(1, "first")
		builder.Add(3, "three")

		result := builder.Result()
		if len(result) != 3 {
			t.Fatalf("Expected 3 groups, got %d", len(result))
		}

		var group1, group2, group3 *Group[int, string]
		for i := range result {
			switch result[i].Key() {
			case 1:
				g := result[i]
				group1 = &g
			case 2:
				g := result[i]
				group2 = &g
			case 3:
				g := result[i]
				group3 = &g
			}
		}

		if group1 == nil || group2 == nil || group3 == nil {
			t.Fatal("Missing expected groups")
		}

		key1Items := group1.Items()
		if len(key1Items) != 2 || key1Items[0] != "one" || key1Items[1] != "first" {
			t.Errorf("Expected key1 items [\"one\", \"first\"], got %v", key1Items)
		}

		key2Items := group2.Items()
		if len(key2Items) != 1 || key2Items[0] != "two" {
			t.Errorf("Expected key2 items [\"two\"], got %v", key2Items)
		}

		key3Items := group3.Items()
		if len(key3Items) != 1 || key3Items[0] != "three" {
			t.Errorf("Expected key3 items [\"three\"], got %v", key3Items)
		}
	})

	t.Run("add with struct key", func(t *testing.T) {
		t.Parallel()
		type KeyStruct struct {
			ID   int
			Name string
		}

		builder := NewGroupingBuilder[KeyStruct, int](comparer.Default[KeyStruct]())

		key1 := KeyStruct{ID: 1, Name: "A"}
		key2 := KeyStruct{ID: 2, Name: "B"}

		builder.Add(key1, 100)
		builder.Add(key2, 200)
		builder.Add(key1, 150)

		result := builder.Result()
		if len(result) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(result))
		}

		var group1, group2 *Group[KeyStruct, int]
		for i := range result {
			if result[i].Key().ID == 1 && result[i].Key().Name == "A" {
				g := result[i]
				group1 = &g
			} else if result[i].Key().ID == 2 && result[i].Key().Name == "B" {
				g := result[i]
				group2 = &g
			}
		}

		if group1 == nil || group2 == nil {
			t.Fatal("Missing expected groups")
		}

		key1Items := group1.Items()
		if len(key1Items) != 2 || key1Items[0] != 100 || key1Items[1] != 150 {
			t.Errorf("Expected key1 items [100, 150], got %v", key1Items)
		}

		key2Items := group2.Items()
		if len(key2Items) != 1 || key2Items[0] != 200 {
			t.Errorf("Expected key2 items [200], got %v", key2Items)
		}
	})

	t.Run("add preserves insertion order", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		builder.Add("key", 3)
		builder.Add("key", 1)
		builder.Add("key", 2)

		result := builder.Result()
		if len(result) != 1 {
			t.Fatalf("Expected 1 group, got %d", len(result))
		}

		items := result[0].Items()
		expected := []int{3, 1, 2}
		if len(items) != len(expected) {
			t.Errorf("Expected %d items, got %d", len(expected), len(items))
		} else {
			for i, expectedVal := range expected {
				if items[i] != expectedVal {
					t.Errorf("Item %d: expected %d, got %d (insertion order not preserved)", i, expectedVal, items[i])
				}
			}
		}
	})
}

func TestGroupingBuilder_Result(t *testing.T) {
	t.Run("result with empty builder", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		result := builder.Result()
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %d groups", len(result))
		}
	})

	t.Run("result preserves group order", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		// Add in specific order
		builder.Add("first", 1)
		builder.Add("second", 2)
		builder.Add("first", 11)
		builder.Add("third", 3)

		result := builder.Result()
		if len(result) != 3 {
			t.Fatalf("Expected 3 groups, got %d", len(result))
		}

		expectedKeys := []string{"first", "second", "third"}
		for i, expectedKey := range expectedKeys {
			if result[i].Key() != expectedKey {
				t.Errorf("Group %d: expected key %q, got %q", i, expectedKey, result[i].Key())
			}
		}
	})

	t.Run("result can be called multiple times", func(t *testing.T) {
		t.Parallel()
		builder := NewGroupingBuilder[string, int](comparer.Default[string]())

		builder.Add("key", 100)
		builder.Add("key", 200)

		// Call Result multiple times
		result1 := builder.Result()
		result2 := builder.Result()

		if len(result1) != len(result2) {
			t.Errorf("Multiple Result() calls returned different lengths: %d vs %d", len(result1), len(result2))
		}

		if len(result1) != 1 {
			t.Fatalf("Expected 1 group, got %d", len(result1))
		}

		if result1[0].Key() != result2[0].Key() {
			t.Errorf("Keys differ between Result() calls: %v vs %v", result1[0].Key(), result2[0].Key())
		}

		items1 := result1[0].Items()
		items2 := result2[0].Items()
		if len(items1) != len(items2) {
			t.Errorf("Items lengths differ between Result() calls: %d vs %d", len(items1), len(items2))
		} else {
			for i := range items1 {
				if items1[i] != items2[i] {
					t.Errorf("Item %d differs between Result() calls: %d vs %d", i, items1[i], items2[i])
				}
			}
		}
	})

	t.Run("result with complex types", func(t *testing.T) {
		t.Parallel()
		type ComplexKey struct {
			ID    int
			Value string
		}
		type ComplexValue struct {
			Name  string
			Count int
		}

		builder := NewGroupingBuilder[ComplexKey, ComplexValue](comparer.Default[ComplexKey]())

		key1 := ComplexKey{ID: 1, Value: "test1"}
		key2 := ComplexKey{ID: 2, Value: "test2"}

		val1 := ComplexValue{Name: "A", Count: 10}
		val2 := ComplexValue{Name: "B", Count: 20}
		val3 := ComplexValue{Name: "C", Count: 30}

		builder.Add(key1, val1)
		builder.Add(key2, val2)
		builder.Add(key1, val3)

		result := builder.Result()
		if len(result) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(result))
		}

		var group1, group2 *Group[ComplexKey, ComplexValue]
		for i := range result {
			if result[i].Key().ID == 1 {
				g := result[i]
				group1 = &g
			} else if result[i].Key().ID == 2 {
				g := result[i]
				group2 = &g
			}
		}

		if group1 == nil || group2 == nil {
			t.Fatal("Missing expected groups")
		}

		if group1.Key().ID != 1 || group1.Key().Value != "test1" {
			t.Errorf("Expected group1 key {1, 'test1'}, got %+v", group1.Key())
		}

		items1 := group1.Items()
		if len(items1) != 2 || items1[0] != val1 || items1[1] != val3 {
			t.Errorf("Expected group1 items [val1, val3], got %v", items1)
		}

		// Check group2 (key2)
		if group2.Key().ID != 2 || group2.Key().Value != "test2" {
			t.Errorf("Expected group2 key {2, 'test2'}, got %+v", group2.Key())
		}

		items2 := group2.Items()
		if len(items2) != 1 || items2[0] != val2 {
			t.Errorf("Expected group2 items [val2], got %v", items2)
		}
	})
}

func BenchmarkGroupingBuilder(b *testing.B) {
	b.Run("add many different keys", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builder := NewGroupingBuilder[string, int](comparer.Default[string]())
			for j := 0; j < 100; j++ {
				builder.Add(string(rune('A'+j%26)), j)
			}
			_ = builder.Result()
		}
	})

	b.Run("add many same keys", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builder := NewGroupingBuilder[string, int](comparer.Default[string]())
			for j := 0; j < 100; j++ {
				builder.Add("same-key", j)
			}
			_ = builder.Result()
		}
	})

	b.Run("mixed add operations", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			builder := NewGroupingBuilder[int, string](comparer.Default[int]())
			for j := 0; j < 1000; j++ {
				builder.Add(j%100, "value")
			}
			_ = builder.Result()
		}
	})
}
