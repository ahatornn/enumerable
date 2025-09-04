package enumerable

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/ahatornn/enumerable/comparer"
	"github.com/ahatornn/enumerable/hashcode"
)

func TestContains(t *testing.T) {
	t.Run("Enumerator[int] contains existing element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.Contains(3)

		if !result {
			t.Error("Expected to find element 3")
		}
	})

	t.Run("EnumeratorAny[string] contains existing element", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"hello", "world", "go"})

		result := enumerator.Contains("world", comparer.Default[string]())

		if !result {
			t.Error("Expected to find element 'world'")
		}
	})

	t.Run("Enumerator[int] does not contain element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		result := enumerator.Contains(10)

		if result {
			t.Error("Expected not to find element 10")
		}
	})

	t.Run("EnumeratorAny[string] does not contain element", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"hello", "world", "go"})

		result := enumerator.Contains("java", comparer.Default[string]())

		if result {
			t.Error("Expected not to find element 'java'")
		}
	})

	t.Run("Enumerator[int] contains first element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42, 1, 2, 3})

		result := enumerator.Contains(42)

		if !result {
			t.Error("Expected to find first element 42")
		}
	})

	t.Run("EnumeratorAny[string] contains last element", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"a", "b", "c", "final"})

		result := enumerator.Contains("final", comparer.Default[string]())

		if !result {
			t.Error("Expected to find last element 'final'")
		}
	})

	t.Run("Enumerator[int] empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		result := enumerator.Contains(1)

		if result {
			t.Error("Expected false for empty slice")
		}
	})

	t.Run("EnumeratorAny[string] empty slice", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{})

		result := enumerator.Contains("test", comparer.Default[string]())

		if result {
			t.Error("Expected false for empty slice")
		}
	})

	t.Run("Enumerator[int] nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		result := enumerator.Contains(1)

		if result {
			t.Error("Expected false for nil enumerator")
		}
	})

	t.Run("EnumeratorAny[int] nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = nil

		result := enumerator.Contains(1, comparer.Default[int]())

		if result {
			t.Error("Expected false for nil enumerator")
		}
	})

	t.Run("EnumeratorAny[int] nil comparer", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[int] = FromSliceAny([]int{1, 2, 3})
		var nilComparer comparer.EqualityComparer[int] = nil

		result := enumerator.Contains(2, nilComparer)

		if result {
			t.Error("Expected false for nil comparer")
		}
	})

	t.Run("Enumerator[bool] contains various values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]bool{true, false, true})

		if !enumerator.Contains(true) {
			t.Error("Expected to find true")
		}

		if !enumerator.Contains(false) {
			t.Error("Expected to find false")
		}
	})

	t.Run("Enumerator[float64] contains with exact values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{1.1, 2.2, 3.3, 4.4})

		if !enumerator.Contains(2.2) {
			t.Error("Expected to find 2.2")
		}

		if enumerator.Contains(5.5) {
			t.Error("Expected not to find 5.5")
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

		idComparer := comparer.ByField(func(u User) int { return u.ID })
		targetUser := User{ID: 2, Name: "Unknown"}

		result := enumerator.Contains(targetUser, idComparer)

		if !result {
			t.Error("Expected to find user with ID 2")
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

		nameComparer := comparer.ByField(func(c Config) string { return c.Name })
		targetConfig := Config{Name: "Config2", Options: nil, Enabled: true}

		result := enumerator.Contains(targetConfig, nameComparer)

		if !result {
			t.Error("Expected to find config with name 'Config2'")
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

		idComparer := comparer.ByField(func(d Data) int { return d.ID })
		targetData := Data{ID: 2, Meta: nil, Tags: nil}

		result := enumerator.Contains(targetData, idComparer)

		if !result {
			t.Error("Expected to find data with ID 2")
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

		nameComparer := comparer.ByField(func(h Handler) string { return h.Name })
		targetHandler := Handler{Name: "Handler2", Callback: nil, Active: true}

		result := enumerator.Contains(targetHandler, nameComparer)

		if !result {
			t.Error("Expected to find handler with name 'Handler2'")
		}
	})

	t.Run("Enumerator[struct] custom comparer with tolerance", func(t *testing.T) {
		t.Parallel()
		type Point struct {
			X, Y float64
		}

		points := []Point{
			{X: 1.0, Y: 2.0},
			{X: 3.0, Y: 4.0},
			{X: 5.0, Y: 6.0},
		}
		var enumerator = FromSliceAny(points)

		toleranceComparer := comparer.Custom(
			func(a, b Point) bool {
				const epsilon = 1e-9
				return math.Abs(a.X-b.X) < epsilon && math.Abs(a.Y-b.Y) < epsilon
			},
			func(p Point) uint64 {
				x := math.Round(p.X*1e6) / 1e6
				y := math.Round(p.Y*1e6) / 1e6
				return hashcode.Combine(x, y)
			},
		)

		targetPoint := Point{X: 3.0 + 1e-10, Y: 4.0 - 1e-10} // Within tolerance

		result := enumerator.Contains(targetPoint, toleranceComparer)

		if !result {
			t.Error("Expected to find point within tolerance")
		}
	})

	t.Run("EnumeratorAny[string] case-insensitive comparison", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"Hello", "WORLD", "Go", "Programming"})

		// Case-insensitive comparer
		caseInsensitiveComparer := comparer.Custom(
			func(a, b string) bool {
				return strings.EqualFold(a, b)
			},
			func(s string) uint64 {
				return hashcode.Compute(strings.ToLower(s))
			},
		)

		result := enumerator.Contains("world", caseInsensitiveComparer)

		if !result {
			t.Error("Expected to find 'world' case-insensitively")
		}
	})

	t.Run("EnumeratorAny[struct] composite comparer", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}

		people := []Person{
			{FirstName: "John", LastName: "Doe", Age: 30},
			{FirstName: "Jane", LastName: "Smith", Age: 25},
			{FirstName: "Bob", LastName: "Johnson", Age: 35},
		}
		var enumerator EnumeratorAny[Person] = FromSliceAny(people)

		firstNameComparer := comparer.ByField(func(p Person) string { return p.FirstName })
		lastNameComparer := comparer.ByField(func(p Person) string { return p.LastName })
		compositeComparer := comparer.Composite(firstNameComparer, lastNameComparer)

		targetPerson := Person{FirstName: "Jane", LastName: "Smith", Age: 99}

		result := enumerator.Contains(targetPerson, compositeComparer)

		if !result {
			t.Error("Expected to find person with same first and last name")
		}
	})

	t.Run("Enumerator[int] contains zero value", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 0, 2, 3})

		result := enumerator.Contains(0)

		if !result {
			t.Error("Expected to find zero value")
		}
	})

	t.Run("EnumeratorAny[string] contains empty string", func(t *testing.T) {
		t.Parallel()
		var enumerator EnumeratorAny[string] = FromSliceAny([]string{"hello", "", "world"})

		result := enumerator.Contains("", comparer.Default[string]())

		if !result {
			t.Error("Expected to find empty string")
		}
	})

	t.Run("EnumeratorAny[struct] multiple matches", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Category string
			Name     string
		}

		items := []Item{
			{Category: "A", Name: "Item1"},
			{Category: "B", Name: "Item2"},
			{Category: "A", Name: "Item3"},
		}
		var enumerator EnumeratorAny[Item] = FromSliceAny(items)

		categoryComparer := comparer.ByField(func(i Item) string { return i.Category })
		targetItem := Item{Category: "A", Name: "Target"}

		result := enumerator.Contains(targetItem, categoryComparer)

		if !result {
			t.Error("Expected to find item with category 'A'")
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

		nameComparer := comparer.ByField(func(n NestedStruct) string { return n.Name })
		targetNested := NestedStruct{Name: "Second", Data: nil}

		result := enumerator.Contains(targetNested, nameComparer)

		if !result {
			t.Error("Expected to find nested struct with name 'Second'")
		}
	})
}

func BenchmarkContains(b *testing.B) {
	b.Run("Enumerator[int] contains at beginning", func(b *testing.B) {
		items := make([]int, 1000)
		items[0] = 42
		enumerator := FromSlice(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains(42)
			if !result {
				b.Fatal("Expected to find 42")
			}
		}
	})

	b.Run("EnumeratorAny[int] contains at end", func(b *testing.B) {
		items := make([]int, 1000)
		items[999] = 999
		var enumerator EnumeratorAny[int] = FromSliceAny(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains(999, comparer.Default[int]())
			if !result {
				b.Fatal("Expected to find 999")
			}
		}
	})

	b.Run("Enumerator[int] does not contain", func(b *testing.B) {
		items := make([]int, 1000)
		for i := range items {
			items[i] = i
		}
		enumerator := FromSlice(items)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains(-1)
			if result {
				b.Fatal("Expected not to find -1")
			}
		}
	})

	b.Run("EnumeratorAny[string] contains with custom comparer", func(b *testing.B) {
		items := make([]string, 100)
		for i := range items {
			items[i] = fmt.Sprintf("item%d", i)
		}
		var enumerator EnumeratorAny[string] = FromSliceAny(items)

		comparer := comparer.Default[string]()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains("item50", comparer)
			if !result {
				b.Fatal("Expected to find item50")
			}
		}
	})

	b.Run("Enumerator[int] empty slice", func(b *testing.B) {
		enumerator := FromSlice([]int{})

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains(1)
			if result {
				b.Fatal("Expected false for empty slice")
			}
		}
	})

	b.Run("EnumeratorAny[int] nil enumerator", func(b *testing.B) {
		var enumerator EnumeratorAny[int] = nil

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains(1, comparer.Default[int]())
			if result {
				b.Fatal("Expected false for nil enumerator")
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

		nameComparer := comparer.ByField(func(c Config) string { return c.Name })

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			target := Config{Name: "Config50", Options: nil, Enabled: true}
			result := enumerator.Contains(target, nameComparer)
			if !result {
				b.Fatal("Expected to find Config50")
			}
		}
	})

	b.Run("EnumeratorAny[int] nil comparer", func(b *testing.B) {
		var enumerator EnumeratorAny[int] = FromSliceAny([]int{1, 2, 3})
		var nilComparer comparer.EqualityComparer[int] = nil

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result := enumerator.Contains(2, nilComparer)
			if result {
				b.Fatal("Expected false for nil comparer")
			}
		}
	})
}
