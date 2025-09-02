package comparer

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
)

func TestByField(t *testing.T) {
	t.Run("compare by int field", func(t *testing.T) {
		t.Parallel()
		type User struct {
			ID   int
			Name string
		}

		comparer := ByField(func(u User) int { return u.ID })

		user1 := User{ID: 1, Name: "Alice"}
		user2 := User{ID: 1, Name: "Bob"}
		user3 := User{ID: 2, Name: "Alice"}

		if !comparer(user1, user2) {
			t.Error("Expected users with same ID to be equal")
		}

		if comparer(user1, user3) {
			t.Error("Expected users with different IDs to be unequal")
		}
	})

	t.Run("compare by string field", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float64
		}

		comparer := ByField(func(p Product) string { return p.Name })

		product1 := Product{Name: "Laptop", Price: 1000}
		product2 := Product{Name: "Laptop", Price: 1500}
		product3 := Product{Name: "Phone", Price: 800}

		if !comparer(product1, product2) {
			t.Error("Expected products with same name to be equal")
		}

		if comparer(product1, product3) {
			t.Error("Expected products with different names to be unequal")
		}
	})

	t.Run("compare by pointer field", func(t *testing.T) {
		t.Parallel()
		type Node struct {
			ID   *int
			Data string
		}

		comparer := ByField(func(n Node) *int { return n.ID })

		id1, id2 := 1, 2
		node1 := Node{ID: &id1, Data: "data1"}
		node2 := Node{ID: &id1, Data: "data2"}
		node3 := Node{ID: &id2, Data: "data1"}

		if !comparer(node1, node2) {
			t.Error("Expected nodes with same pointer ID to be equal")
		}

		if comparer(node1, node3) {
			t.Error("Expected nodes with different pointer IDs to be unequal")
		}
	})

	t.Run("compare by struct with non-comparable fields", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			ID    int
			Roles []string               // non-comparable
			Meta  map[string]interface{} // non-comparable
		}

		comparer := ByField(func(c Config) int { return c.ID })

		config1 := Config{ID: 42, Roles: []string{"admin"}, Meta: map[string]interface{}{"key": "value"}}
		config2 := Config{ID: 42, Roles: []string{"user"}, Meta: map[string]interface{}{"other": "data"}}
		config3 := Config{ID: 43, Roles: []string{"admin"}, Meta: map[string]interface{}{"key": "value"}}

		if !comparer(config1, config2) {
			t.Error("Expected configs with same ID to be equal despite different non-comparable fields")
		}

		if comparer(config1, config3) {
			t.Error("Expected configs with different IDs to be unequal")
		}
	})
}

func TestComposite(t *testing.T) {
	t.Run("composite with multiple field comparers", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}

		firstNameComparer := ByField(func(p Person) string { return p.FirstName })
		lastNameComparer := ByField(func(p Person) string { return p.LastName })
		ageComparer := ByField(func(p Person) int { return p.Age })

		compositeComparer := Composite(firstNameComparer, lastNameComparer, ageComparer)

		person1 := Person{FirstName: "John", LastName: "Doe", Age: 30}
		person2 := Person{FirstName: "John", LastName: "Doe", Age: 30}
		person3 := Person{FirstName: "John", LastName: "Smith", Age: 30}
		person4 := Person{FirstName: "John", LastName: "Doe", Age: 31}

		if !compositeComparer(person1, person2) {
			t.Error("Expected identical persons to be equal")
		}

		if compositeComparer(person1, person3) {
			t.Error("Expected persons with different last names to be unequal")
		}

		if compositeComparer(person1, person4) {
			t.Error("Expected persons with different ages to be unequal")
		}
	})

	t.Run("composite with single comparer", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			ID int
		}

		idComparer := ByField(func(i Item) int { return i.ID })
		compositeComparer := Composite(idComparer)

		item1 := Item{ID: 1}
		item2 := Item{ID: 1}
		item3 := Item{ID: 2}

		if !compositeComparer(item1, item2) {
			t.Error("Expected items with same ID to be equal")
		}

		if compositeComparer(item1, item3) {
			t.Error("Expected items with different IDs to be unequal")
		}
	})

	t.Run("empty composite comparer", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Value int
		}

		emptyComparer := Composite[Data]()

		data1 := Data{Value: 1}
		data2 := Data{Value: 2}

		if !emptyComparer(data1, data2) {
			t.Error("Expected empty composite comparer to always return true")
		}
	})

	t.Run("composite with no matching comparer", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Field1 string
			Field2 int
		}

		neverEqual := Custom(func(a, b Record) bool { return false })
		alwaysEqual := Custom(func(a, b Record) bool { return true })

		composite := Composite(neverEqual, alwaysEqual)

		record1 := Record{Field1: "test", Field2: 1}
		record2 := Record{Field1: "test", Field2: 1}

		if composite(record1, record2) {
			t.Error("Expected composite to return false when first comparer returns false")
		}
	})

	t.Run("composite short-circuit behavior", func(t *testing.T) {
		t.Parallel()
		type TestStruct struct {
			A int
			B int
		}

		callCount := 0
		expensiveComparer := Custom(func(a, b TestStruct) bool {
			callCount++
			return a.B == b.B
		})

		cheapComparer := ByField(func(s TestStruct) int { return s.A })

		composite := Composite(cheapComparer, expensiveComparer)

		struct1 := TestStruct{A: 1, B: 10}
		struct2 := TestStruct{A: 2, B: 20}

		callCount = 0
		result := composite(struct1, struct2)

		if result {
			t.Error("Expected composite to return false")
		}

		if callCount != 0 {
			t.Error("Expected expensive comparer to not be called due to short-circuit")
		}
	})
}

func TestCustom(t *testing.T) {
	t.Run("custom comparer with floating point tolerance", func(t *testing.T) {
		t.Parallel()
		type Point struct {
			X, Y float64
		}

		toleranceComparer := Custom(func(a, b Point) bool {
			const epsilon = 1e-9
			return math.Abs(a.X-b.X) < epsilon && math.Abs(a.Y-b.Y) < epsilon
		})

		point1 := Point{X: 1.0, Y: 2.0}
		point2 := Point{X: 1.0 + 1e-10, Y: 2.0 - 1e-10}
		point3 := Point{X: 1.1, Y: 2.0}

		if !toleranceComparer(point1, point2) {
			t.Error("Expected points within tolerance to be equal")
		}

		if toleranceComparer(point1, point3) {
			t.Error("Expected points outside tolerance to be unequal")
		}
	})

	t.Run("custom comparer with case-insensitive string comparison", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name string
			Age  int
		}

		caseInsensitiveComparer := Custom(func(a, b User) bool {
			return strings.EqualFold(a.Name, b.Name) && a.Age == b.Age
		})

		user1 := User{Name: "Alice", Age: 25}
		user2 := User{Name: "alice", Age: 25}
		user3 := User{Name: "Alice", Age: 26}

		if !caseInsensitiveComparer(user1, user2) {
			t.Error("Expected case-insensitive name match to be equal")
		}

		if caseInsensitiveComparer(user1, user3) {
			t.Error("Expected different ages to make users unequal")
		}
	})

	t.Run("custom comparer with business logic", func(t *testing.T) {
		t.Parallel()
		type Order struct {
			ID          int
			ExternalRef string
			Status      string
		}

		orderComparer := Custom(func(a, b Order) bool {
			return a.ID == b.ID || (a.ExternalRef != "" && a.ExternalRef == b.ExternalRef)
		})

		order1 := Order{ID: 1, ExternalRef: "EXT001", Status: "pending"}
		order2 := Order{ID: 1, ExternalRef: "EXT002", Status: "shipped"}
		order3 := Order{ID: 2, ExternalRef: "EXT001", Status: "cancelled"}
		order4 := Order{ID: 3, ExternalRef: "EXT003", Status: "pending"}

		if !orderComparer(order1, order2) {
			t.Error("Expected orders with same ID to be equal")
		}

		if !orderComparer(order1, order3) {
			t.Error("Expected orders with same external reference to be equal")
		}

		if orderComparer(order1, order4) {
			t.Error("Expected orders with different ID and external reference to be unequal")
		}
	})

	t.Run("custom comparer with time tolerance", func(t *testing.T) {
		t.Parallel()
		type Event struct {
			Timestamp time.Time
			Data      string
		}

		timeToleranceComparer := Custom(func(a, b Event) bool {
			diff := a.Timestamp.Sub(b.Timestamp)
			return diff >= -time.Second && diff <= time.Second
		})

		now := time.Now()
		event1 := Event{Timestamp: now, Data: "data1"}
		event2 := Event{Timestamp: now.Add(500 * time.Millisecond), Data: "data2"}
		event3 := Event{Timestamp: now.Add(2 * time.Second), Data: "data3"}

		if !timeToleranceComparer(event1, event2) {
			t.Error("Expected events within time tolerance to be equal")
		}

		if timeToleranceComparer(event1, event3) {
			t.Error("Expected events outside time tolerance to be unequal")
		}
	})

	t.Run("custom comparer wrapping existing function", func(t *testing.T) {
		t.Parallel()
		type Simple struct {
			Value int
		}

		simpleComparer := Custom(func(a, b Simple) bool {
			return a.Value == b.Value
		})

		s1 := Simple{Value: 42}
		s2 := Simple{Value: 42}
		s3 := Simple{Value: 43}

		if !simpleComparer(s1, s2) {
			t.Error("Expected simples with same value to be equal")
		}

		if simpleComparer(s1, s3) {
			t.Error("Expected simples with different values to be unequal")
		}
	})
}

func TestComparerCombinations(t *testing.T) {
	t.Run("ByField with Custom field selector", func(t *testing.T) {
		t.Parallel()
		type Complex struct {
			Data []int
		}

		lengthComparer := ByField(func(c Complex) int { return len(c.Data) })

		c1 := Complex{Data: []int{1, 2, 3}}
		c2 := Complex{Data: []int{4, 5, 6}}
		c3 := Complex{Data: []int{1, 2}}

		if !lengthComparer(c1, c2) {
			t.Error("Expected complexes with same data length to be equal")
		}

		if lengthComparer(c1, c3) {
			t.Error("Expected complexes with different data lengths to be unequal")
		}
	})

	t.Run("Composite with Custom and ByField", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			ID   int
			Name string
			Tags []string
		}

		idComparer := ByField(func(r Record) int { return r.ID })
		customTagComparer := Custom(func(a, b Record) bool {
			return len(a.Tags) == len(b.Tags)
		})

		composite := Composite(idComparer, customTagComparer)

		r1 := Record{ID: 1, Name: "Record1", Tags: []string{"tag1", "tag2"}}
		r2 := Record{ID: 1, Name: "Record2", Tags: []string{"tag3", "tag4"}}
		r3 := Record{ID: 2, Name: "Record3", Tags: []string{"tag1", "tag2"}}
		r4 := Record{ID: 1, Name: "Record4", Tags: []string{"tag1"}}

		if !composite(r1, r2) {
			t.Error("Expected records with same ID and tag count to be equal")
		}

		if composite(r1, r3) {
			t.Error("Expected records with different IDs to be unequal")
		}

		if composite(r1, r4) {
			t.Error("Expected records with different tag counts to be unequal")
		}
	})
}

func BenchmarkComparers(b *testing.B) {
	b.Run("ByField comparer", func(b *testing.B) {
		type User struct {
			ID   int
			Name string
		}

		comparer := ByField(func(u User) int { return u.ID })

		users := make([]User, b.N)
		for i := 0; i < b.N; i++ {
			users[i] = User{ID: i % 1000, Name: fmt.Sprintf("User%d", i)}
		}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = comparer(users[i], User{ID: i % 1000, Name: "Test"})
		}
		b.StopTimer()
		if !result {
			b.Log("Unexpected result")
		}
	})

	b.Run("Composite comparer", func(b *testing.B) {
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}

		firstNameComparer := ByField(func(p Person) string { return p.FirstName })
		lastNameComparer := ByField(func(p Person) string { return p.LastName })
		composite := Composite(firstNameComparer, lastNameComparer)

		persons := make([]Person, b.N)
		for i := 0; i < b.N; i++ {
			persons[i] = Person{
				FirstName: fmt.Sprintf("First%d", i%100),
				LastName:  fmt.Sprintf("Last%d", i%100),
				Age:       i,
			}
		}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = composite(persons[i], Person{
				FirstName: fmt.Sprintf("First%d", i%100),
				LastName:  fmt.Sprintf("Last%d", i%100),
				Age:       0,
			})
		}
		b.StopTimer()
		if !result {
			b.Log("Unexpected result")
		}
	})

	b.Run("Custom comparer", func(b *testing.B) {
		type Point struct {
			X, Y float64
		}

		custom := Custom(func(a, b Point) bool {
			return a.X == b.X && a.Y == b.Y
		})

		points := make([]Point, b.N)
		for i := 0; i < b.N; i++ {
			points[i] = Point{X: float64(i % 1000), Y: float64(i % 500)}
		}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = custom(points[i], Point{X: float64(i % 1000), Y: float64(i % 500)})
		}
		b.StopTimer()
		if !result {
			b.Log("Unexpected result")
		}
	})
}
