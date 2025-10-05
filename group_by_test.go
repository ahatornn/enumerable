package enumerable

import (
	"fmt"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
	"github.com/ahatornn/enumerable/grouping"
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

func TestOrderEnumerator_GroupBy(t *testing.T) {
	t.Run("group by with sorted data preserves order", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{4, 1, 3, 2, 5, 6})

		groups := source.OrderBy(comparer.ComparerInt).
			GroupBy(func(x int) any { return x % 2 }).
			ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		if groups[0].Key() != 1 || groups[1].Key() != 0 {
			t.Errorf("Expected group keys [1, 0], got [%v, %v]", groups[0].Key(), groups[1].Key())
		}

		expectedItems := [][]int{
			{1, 3, 5},
			{2, 4, 6},
		}

		for i, element := range groups {
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

	t.Run("group by string with sorted data", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}
		people := []Person{
			{"Charlie", 30},
			{"Alice", 25},
			{"Bob", 30},
			{"David", 25},
		}
		source := FromSlice(people)

		groups := source.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
			GroupBy(func(p Person) any { return p.Age }).
			ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		if groups[0].Key() != 25 || groups[1].Key() != 30 {
			t.Errorf("Expected keys [25, 30], got [%v, %v]", groups[0].Key(), groups[1].Key())
		}

		expected := [][]Person{
			{{"Alice", 25}, {"David", 25}},
			{{"Charlie", 30}, {"Bob", 30}},
		}

		for i, group := range groups {
			items := group.Items()
			exp := expected[i]
			if len(items) != len(exp) {
				t.Errorf("Group %d: expected %d items, got %d", i, len(exp), len(items))
				continue
			}
			for j, p := range exp {
				if items[j].Name != p.Name || items[j].Age != p.Age {
					t.Errorf("Group %d, item %d: expected %+v, got %+v", i, j, p, items[j])
				}
			}
		}
	})

	t.Run("group by with ThenBy sorting", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}
		people := []Person{
			{"Bob", 30},
			{"Alice", 30},
			{"Charlie", 25},
			{"David", 25},
		}
		source := FromSlice(people)

		groups := source.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
			ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) }).
			GroupBy(func(p Person) any { return p.Age }).
			ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		if groups[0].Key() != 25 || groups[1].Key() != 30 {
			t.Errorf("Expected keys [25, 30], got [%v, %v]", groups[0].Key(), groups[1].Key())
		}

		expected := [][]Person{
			{{"Charlie", 25}, {"David", 25}},
			{{"Alice", 30}, {"Bob", 30}},
		}

		for i, group := range groups {
			items := group.Items()
			exp := expected[i]
			if len(items) != len(exp) {
				t.Errorf("Group %d: expected %d items, got %d", i, len(exp), len(items))
				continue
			}
			for j, p := range exp {
				if items[j].Name != p.Name || items[j].Age != p.Age {
					t.Errorf("Group %d, item %d: expected %+v, got %+v", i, j, p, items[j])
				}
			}
		}
	})

	t.Run("empty sorted source", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{})

		groups := source.OrderBy(comparer.ComparerInt).
			GroupBy(func(x int) any { return x }).
			ToSlice()

		if len(groups) != 0 {
			t.Errorf("Expected 0 groups from empty sorted source, got %d", len(groups))
		}
	})

	t.Run("single element sorted", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]string{"hello"})

		groups := source.OrderBy(comparer.ComparerString).
			GroupBy(func(s string) any { return len(s) }).
			ToSlice()

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

	t.Run("nil keySelector on sorted data", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3})

		groups := source.OrderBy(comparer.ComparerInt).
			GroupBy(nil).
			ToSlice()

		if len(groups) != 0 {
			t.Errorf("Expected 0 groups when keySelector is nil, got %d", len(groups))
		}
	})

	t.Run("all same key with sorted data", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{30, 10, 20})

		groups := source.OrderBy(comparer.ComparerInt).
			GroupBy(func(x int) any { return "same" }).
			ToSlice()

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

	t.Run("descending sort before grouping", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 4, 2, 3, 6, 5})

		groups := source.OrderByDescending(comparer.ComparerInt).
			GroupBy(func(x int) any { return x % 2 }).
			ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		if groups[0].Key() != 0 || groups[1].Key() != 1 {
			t.Errorf("Expected keys [0, 1], got [%v, %v]", groups[0].Key(), groups[1].Key())
		}

		expectedItems := [][]int{
			{6, 4, 2},
			{5, 3, 1},
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
					t.Errorf("Group %d, item %d: expected %d, got %d", i, j, v, items[j])
				}
			}
		}
	})

	t.Run("complex struct sorting before grouping", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name     string
			Price    float64
			Category string
		}

		products := []Product{
			{"Laptop", 1200.0, "Electronics"},
			{"Phone", 800.0, "Electronics"},
			{"Desk", 300.0, "Furniture"},
			{"Chair", 200.0, "Furniture"},
		}
		source := FromSlice(products)

		groups := source.OrderByDescending(func(a, b Product) int {
			if a.Price < b.Price {
				return -1
			}
			if a.Price > b.Price {
				return 1
			}
			return 0
		}).GroupBy(func(p Product) any { return p.Category }).ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		var electronicsGroup, furnitureGroup *grouping.Group[any, Product]
		for _, group := range groups {
			if group.Key() == "Electronics" {
				electronicsGroup = &group
			} else if group.Key() == "Furniture" {
				furnitureGroup = &group
			}
		}

		if electronicsGroup == nil || furnitureGroup == nil {
			t.Fatal("Missing expected groups")
		}

		electronicsItems := electronicsGroup.Items()
		if len(electronicsItems) != 2 {
			t.Errorf("Expected 2 electronics items, got %d", len(electronicsItems))
		} else {
			if electronicsItems[0].Name != "Laptop" || electronicsItems[1].Name != "Phone" {
				t.Errorf("Expected Laptop, Phone in electronics group, got %s, %s",
					electronicsItems[0].Name, electronicsItems[1].Name)
			}
		}

		furnitureItems := furnitureGroup.Items()
		if len(furnitureItems) != 2 {
			t.Errorf("Expected 2 furniture items, got %d", len(furnitureItems))
		} else {
			if furnitureItems[0].Name != "Desk" || furnitureItems[1].Name != "Chair" {
				t.Errorf("Expected Desk, Chair in furniture group, got %s, %s",
					furnitureItems[0].Name, furnitureItems[1].Name)
			}
		}
	})

	t.Run("stability of sorting preserved in grouping", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
			ID    int
		}

		items := []Item{
			{Value: 2, ID: 1},
			{Value: 1, ID: 2},
			{Value: 2, ID: 3},
			{Value: 1, ID: 4},
		}
		source := FromSlice(items)

		groups := source.OrderBy(func(a, b Item) int { return a.Value - b.Value }).
			GroupBy(func(i Item) any { return i.Value }).
			ToSlice()

		if len(groups) != 2 {
			t.Fatalf("Expected 2 groups, got %d", len(groups))
		}

		var group1, group2 *grouping.Group[any, Item]
		for _, group := range groups {
			if group.Key() == 1 {
				group1 = &group
			} else if group.Key() == 2 {
				group2 = &group
			}
		}

		if group1 == nil || group2 == nil {
			t.Fatal("Missing expected groups")
		}

		items1 := group1.Items()
		if len(items1) != 2 || items1[0].ID != 2 || items1[1].ID != 4 {
			t.Errorf("Group 1 should preserve order: expected IDs [2, 4], got [%d, %d]",
				items1[0].ID, items1[1].ID)
		}

		items2 := group2.Items()
		if len(items2) != 2 || items2[0].ID != 1 || items2[1].ID != 3 {
			t.Errorf("Group 2 should preserve order: expected IDs [1, 3], got [%d, %d]",
				items2[0].ID, items2[1].ID)
		}
	})
}

func BenchmarkOrderEnumerator_GroupBy(b *testing.B) {
	b.Run("group by after ascending sort", func(b *testing.B) {
		items := make([]int, 1000)
		for i := range items {
			items[i] = i % 100 // 100 different groups
		}
		source := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = source.OrderBy(comparer.ComparerInt).
				GroupBy(func(x int) any { return x % 10 }).
				ToSlice()
		}
	})

	b.Run("group by after descending sort", func(b *testing.B) {
		items := make([]int, 1000)
		for i := range items {
			items[i] = i % 100
		}
		source := FromSlice(items)

		for i := 0; i < b.N; i++ {
			_ = source.OrderByDescending(comparer.ComparerInt).
				GroupBy(func(x int) any { return x % 5 }).
				ToSlice()
		}
	})

	b.Run("group by after ThenBy sort", func(b *testing.B) {
		type Person struct {
			Age  int
			Name string
		}
		people := make([]Person, 1000)
		for i := range people {
			people[i] = Person{
				Age:  i % 10,
				Name: fmt.Sprintf("Person%d", i%100),
			}
		}
		source := FromSlice(people)

		for i := 0; i < b.N; i++ {
			_ = source.OrderBy(func(a, b Person) int { return a.Age - b.Age }).
				ThenBy(func(a, b Person) int { return compareStrings(a.Name, b.Name) }).
				GroupBy(func(p Person) any { return p.Age }).
				ToSlice()
		}
	})
}

func TestGroupByEarlyTermination(t *testing.T) {
	t.Run("early termination during enumeration", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		groups := source.GroupBy(func(x int) any { return x % 3 })

		yieldCount := 0
		shouldStopAfter := 2

		resultGroups := make([]grouping.Group[any, int], 0)

		groups(func(group grouping.Group[any, int]) bool {
			yieldCount++
			resultGroups = append(resultGroups, group)

			return yieldCount < shouldStopAfter
		})

		if len(resultGroups) != shouldStopAfter {
			t.Errorf("Expected %d groups due to early termination, got %d", shouldStopAfter, len(resultGroups))
		}

		if len(resultGroups) >= 1 && resultGroups[0].Key() != 1 {
			t.Errorf("Expected first group key to be 1, got %v", resultGroups[0].Key())
		}
		if len(resultGroups) >= 2 && resultGroups[1].Key() != 2 {
			t.Errorf("Expected second group key to be 2, got %v", resultGroups[1].Key())
		}

		if yieldCount != shouldStopAfter {
			t.Errorf("Expected yield to be called %d times, got %d", shouldStopAfter, yieldCount)
		}
	})

	t.Run("early termination with Take after GroupBy", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{1, 2, 3, 4, 5, 6})
		result := source.GroupBy(func(x int) any { return x % 2 }).
			Take(1).
			ToSlice()

		if len(result) != 1 {
			t.Errorf("Expected 1 group after Take(1), got %d", len(result))
		}

		if result[0].Key() == 1 {
			expectedItems := []int{1, 3, 5}
			actualItems := result[0].Items()
			if len(actualItems) != len(expectedItems) {
				t.Errorf("Expected %d items in first group, got %d", len(expectedItems), len(actualItems))
			} else {
				for i, v := range expectedItems {
					if actualItems[i] != v {
						t.Errorf("Item %d: expected %d, got %d", i, v, actualItems[i])
					}
				}
			}
		} else if result[0].Key() == 0 {
			expectedItems := []int{2, 4, 6}
			actualItems := result[0].Items()
			if len(actualItems) != len(expectedItems) {
				t.Errorf("Expected %d items in first group, got %d", len(expectedItems), len(actualItems))
			} else {
				for i, v := range expectedItems {
					if actualItems[i] != v {
						t.Errorf("Item %d: expected %d, got %d", i, v, actualItems[i])
					}
				}
			}
		} else {
			t.Errorf("Expected first group key to be 1 or 0, got %v", result[0].Key())
		}
	})

	t.Run("Take with OrderEnumerator GroupBy", func(t *testing.T) {
		t.Parallel()
		source := FromSlice([]int{4, 1, 3, 2, 5, 6})
		result := source.OrderBy(comparer.ComparerInt).
			GroupBy(func(x int) any { return x % 2 }).
			Take(1).
			ToSlice()

		if len(result) != 1 {
			t.Errorf("Expected 1 group after Take(1) on sorted GroupBy, got %d", len(result))
		}

		if result[0].Key() != 1 {
			t.Errorf("Expected first group key to be 1 (odd), got %v", result[0].Key())
		}

		expectedItems := []int{1, 3, 5}
		actualItems := result[0].Items()
		if len(actualItems) != len(expectedItems) {
			t.Errorf("Expected %d items in first group, got %d", len(expectedItems), len(actualItems))
		} else {
			for i, v := range expectedItems {
				if actualItems[i] != v {
					t.Errorf("Item %d: expected %d, got %d", i, v, actualItems[i])
				}
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
