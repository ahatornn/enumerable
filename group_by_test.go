package enumerable

import (
	"fmt"
	"testing"
)

func TestGroupBy(t *testing.T) {
	t.Run("group by int key", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3, 4, 5, 6})

		groups := source.GroupBy(func(x int) any { return x % 2 })
		slice := groups.ToSlice()

		if len(slice) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(slice))
		}

		if slice[0].Key() != 1 || slice[1].Key() != 0 {
			t.Errorf("Expected group keys [1, 0], got [%v, %v]", slice[0].Key(), slice[1].Key())
		}

		expectedItems := [][]int{
			{1, 3, 5},
			{2, 4, 6},
		}

		for i, element := range slice {
			items := element.Items()
			exp := expectedItems[i]
			if len(items) != len(exp) {
				t.Errorf("Group %d: expected %d items, got %d", i, len(exp), len(items))
				continue
			}
			for j, v := range exp {
				if items[j] != v {
					t.Errorf("Group %d, item %d: expected %d, got %d", i, j, v, items[j])
				}
			}
		}
	})
	t.Run("group by string key", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			City string
		}
		people := []Person{
			{"Alice", "NYC"},
			{"Bob", "LA"},
			{"Charlie", "NYC"},
			{"Diana", "LA"},
		}
		source := FromSlice(people)

		groups := source.GroupBy(func(p Person) any { return p.City }).ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		if groups[0].Key() != "NYC" || groups[1].Key() != "LA" {
			t.Errorf("Expected keys [NYC, LA], got [%v, %v]", groups[0].Key(), groups[1].Key())
		}

		expected := [][]Person{
			{{"Alice", "NYC"}, {"Charlie", "NYC"}},
			{{"Bob", "LA"}, {"Diana", "LA"}},
		}

		for i, group := range groups {
			items := group.Items()
			exp := expected[i]
			if len(items) != len(exp) {
				t.Errorf("Group %d: expected %d items, got %d", i, len(exp), len(items))
				continue
			}
			for j, p := range exp {
				if items[j].Name != p.Name || items[j].City != p.City {
					t.Errorf("Group %d, item %d: expected %+v, got %+v", i, j, p, items[j])
				}
			}
		}
	})

	t.Run("empty source", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{})

		groups := source.GroupBy(func(x int) any { return x }).ToSlice()

		if len(groups) != 0 {
			t.Errorf("Expected 0 groups from empty source, got %d", len(groups))
		}
	})

	t.Run("nil enumerator", func(t *testing.T) {
		t.Parallel()
		var source Enumerator[int] = nil

		groups := source.GroupBy(func(x int) any { return x }).ToSlice()

		if len(groups) != 0 {
			t.Errorf("Expected 0 groups from nil enumerator, got %d", len(groups))
		}
	})

	t.Run("nil keySelector", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3})

		groups := source.GroupBy(nil).ToSlice()

		if len(groups) != 0 {
			t.Errorf("Expected 0 groups when keySelector is nil, got %d", len(groups))
		}
	})

	t.Run("single element", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]string{"hello"})

		groups := source.GroupBy(func(s string) any { return len(s) }).ToSlice()

		if len(groups) != 1 {
			t.Fatalf("Expected 1 group, got %d", len(groups))
		}

		if groups[0].Key() != 5 {
			t.Errorf("Expected key 5, got %v", groups[0].Key())
		}

		items := groups[0].Items()
		if len(items) != 1 {
			t.Fatalf("Expected 1 item in group, got %d", len(items))
		}
		if items[0] != "hello" {
			t.Errorf("Expected item 'hello', got %q", items[0])
		}
	})

	t.Run("all same key", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{10, 20, 30})

		groups := source.GroupBy(func(x int) any { return "same" }).ToSlice()

		if len(groups) != 1 {
			t.Fatalf("Expected 1 group, got %d", len(groups))
		}

		if groups[0].Key() != "same" {
			t.Errorf("Expected key 'same', got %v", groups[0].Key())
		}

		items := groups[0].Items()
		expected := []int{10, 20, 30}
		if len(items) != len(expected) {
			t.Fatalf("Expected %d items, got %d", len(expected), len(items))
		}
		for i, v := range expected {
			if items[i] != v {
				t.Errorf("Item %d: expected %d, got %d", i, v, items[i])
			}
		}
	})

	t.Run("groups preserve insertion order", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]string{"x", "y", "z", "w", "v"})

		keySelector := func(s string) any {
			switch s {
			case "x", "z":
				return "a"
			case "y", "v":
				return "b"
			case "w":
				return "c"
			default:
				return "unknown"
			}
		}

		groups := source.GroupBy(keySelector).ToSlice()

		if len(groups) != 3 {
			t.Fatalf("Expected 3 groups, got %d", len(groups))
		}

		expectedKeys := []any{"a", "b", "c"}
		for i, key := range expectedKeys {
			if groups[i].Key() != key {
				t.Errorf("Group %d: expected key %v, got %v", i, key, groups[i].Key())
			}
		}

		expectedItems := [][]string{
			{"x", "z"},
			{"y", "v"},
			{"w"},
		}

		for i, group := range groups {
			items := group.Items()
			exp := expectedItems[i]
			if len(items) != len(exp) {
				t.Errorf("Group %d: expected %d items, got %d", i, len(exp), len(items))
				continue
			}
			for j, v := range exp {
				if items[j] != v {
					t.Errorf("Group %d, item %d: expected %q, got %q", i, j, v, items[j])
				}
			}
		}
	})

	t.Run("group with nil key", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Name string
			Cat  *string
		}
		catA := "A"
		items := []Item{
			{"X", &catA},
			{"Y", nil},
			{"Z", &catA},
			{"W", nil},
		}
		source := FromSlice(items)

		groups := source.GroupBy(func(i Item) any { return i.Cat }).ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups (A and nil), got %d", len(groups))
		}

		key0 := groups[0].Key()
		key1 := groups[1].Key()

		if key0 == nil {
			t.Error("First key is nil")
		} else {
			if ptr, ok := key0.(*string); !ok {
				t.Error("First key is not *string")
			} else if *ptr != "A" {
				t.Errorf("First key value is %q, expected 'A'", *ptr)
			}
		}

		if fmt.Sprintf("%v", key1) != "<nil>" {
			t.Error("Second key is not nil")
		}

		if ptr, ok := key0.(*string); !ok || *ptr != "A" {
			t.Error("First key is not pointing to 'A'")
		}

		items0 := groups[0].Items()
		if len(items0) != 2 || items0[0].Name != "X" || items0[1].Name != "Z" {
			t.Errorf("Group 0 items: expected X,Z; got %s,%s", items0[0].Name, items0[1].Name)
		}

		items1 := groups[1].Items()
		if len(items1) != 2 || items1[0].Name != "Y" || items1[1].Name != "W" {
			t.Errorf("Group 1 items: expected Y,W; got %s,%s", items1[0].Name, items1[1].Name)
		}
	})

	t.Run("early termination not applicable to ToSlice", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3, 4, 5})

		groups := source.GroupBy(func(x int) any { return x }).ToSlice()

		if len(groups) != 5 {
			t.Fatalf("Expected 5 groups, got %d", len(groups))
		}

		for i, group := range groups {
			key := group.Key()
			items := group.Items()
			expectedKey := i + 1
			if key != expectedKey {
				t.Errorf("Group %d: expected key %d, got %v", i, expectedKey, key)
			}
			if len(items) != 1 || items[0] != expectedKey {
				t.Errorf("Group %d: expected item [%d], got %v", i, expectedKey, items)
			}
		}
	})
}

func BenchmarkGroupBy(b *testing.B) {
	b.Run("small group by", func(b *testing.B) {
		items := make([]int, 100)
		for i := range items {
			items[i] = i % 10
		}
		source := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = source.GroupBy(func(x int) any { return x % 10 }).ToSlice()
		}
	})

	b.Run("large group by with few keys", func(b *testing.B) {
		items := make([]int, 10000)
		for i := range items {
			items[i] = i % 5
		}
		source := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = source.GroupBy(func(x int) any { return x % 5 }).ToSlice()
		}
	})
}
