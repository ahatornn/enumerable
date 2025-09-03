package comparer

import (
	"math"
	"strings"
	"testing"

	"github.com/ahatornn/enumerable/hashcode"
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

		if !comparer.Equals(user1, user2) {
			t.Error("Expected users with same ID to be equal")
		}

		if comparer.Equals(user1, user3) {
			t.Error("Expected users with different IDs to be unequal")
		}

		if comparer.GetHashCode(user1) != comparer.GetHashCode(user2) {
			t.Error("Expected same hash codes for equal elements")
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

		if !comparer.Equals(product1, product2) {
			t.Error("Expected products with same name to be equal")
		}

		if comparer.Equals(product1, product3) {
			t.Error("Expected products with different names to be unequal")
		}

		if comparer.GetHashCode(product1) != comparer.GetHashCode(product2) {
			t.Error("Expected same hash codes for equal elements")
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

		if !comparer.Equals(node1, node2) {
			t.Error("Expected nodes with same pointer ID to be equal")
		}

		if comparer.Equals(node1, node3) {
			t.Error("Expected nodes with different pointer IDs to be unequal")
		}
	})

	t.Run("compare by struct with non-comparable fields", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			ID    int
			Roles []string
			Meta  map[string]interface{}
		}

		comparer := ByField(func(c Config) int { return c.ID })

		config1 := Config{ID: 42, Roles: []string{"admin"}, Meta: map[string]interface{}{"key": "value"}}
		config2 := Config{ID: 42, Roles: []string{"user"}, Meta: map[string]interface{}{"other": "data"}}
		config3 := Config{ID: 43, Roles: []string{"admin"}, Meta: map[string]interface{}{"key": "value"}}

		if !comparer.Equals(config1, config2) {
			t.Error("Expected configs with same ID to be equal despite different non-comparable fields")
		}

		if comparer.Equals(config1, config3) {
			t.Error("Expected configs with different IDs to be unequal")
		}

		if comparer.GetHashCode(config1) != comparer.GetHashCode(config2) {
			t.Error("Expected same hash codes for equal elements")
		}
	})

	t.Run("hash code consistency", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			Value int
		}

		comparer := ByField(func(i Item) int { return i.Value })

		item := Item{Value: 42}

		hash1 := comparer.GetHashCode(item)
		hash2 := comparer.GetHashCode(item)

		if hash1 != hash2 {
			t.Error("Hash code should be consistent for same object")
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

		if !compositeComparer.Equals(person1, person2) {
			t.Error("Expected identical persons to be equal")
		}

		if compositeComparer.Equals(person1, person3) {
			t.Error("Expected persons with different last names to be unequal")
		}

		if compositeComparer.Equals(person1, person4) {
			t.Error("Expected persons with different ages to be unequal")
		}

		if compositeComparer.GetHashCode(person1) != compositeComparer.GetHashCode(person2) {
			t.Error("Expected same hash codes for equal elements")
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

		if !compositeComparer.Equals(item1, item2) {
			t.Error("Expected items with same ID to be equal")
		}

		if compositeComparer.Equals(item1, item3) {
			t.Error("Expected items with different IDs to be unequal")
		}

		if compositeComparer.GetHashCode(item1) != compositeComparer.GetHashCode(item2) {
			t.Error("Expected same hash codes for equal elements")
		}
	})

	t.Run("composite hash code combination", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			A int
			B string
		}

		comparerA := ByField(func(r Record) int { return r.A })
		comparerB := ByField(func(r Record) string { return r.B })

		composite := Composite(comparerA, comparerB)

		record1 := Record{A: 1, B: "test"}
		record2 := Record{A: 1, B: "test"}
		record3 := Record{A: 1, B: "other"}

		if composite.GetHashCode(record1) != composite.GetHashCode(record2) {
			t.Error("Expected same hash codes for equal elements")
		}

		if composite.GetHashCode(record1) == composite.GetHashCode(record3) {
			t.Log("Warning: Hash codes are the same for different elements (possible collision)")
		}
	})

	t.Run("composite with different comparer types", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			ID   int
			Name string
		}

		idComparer := ByField(func(d Data) int { return d.ID })
		customComparer := Custom(
			func(a, b Data) bool { return a.Name == b.Name },
			func(d Data) uint64 { return hashcode.Compute(d.Name) },
		)

		composite := Composite(idComparer, customComparer)

		data1 := Data{ID: 1, Name: "test"}
		data2 := Data{ID: 1, Name: "test"}
		data3 := Data{ID: 1, Name: "other"}

		if !composite.Equals(data1, data2) {
			t.Error("Expected equal data to be equal")
		}

		if composite.Equals(data1, data3) {
			t.Error("Expected different names to make data unequal")
		}
	})
}

func TestCustom(t *testing.T) {
	t.Run("custom comparer with floating point tolerance", func(t *testing.T) {
		t.Parallel()
		type Point struct {
			X, Y float64
		}

		toleranceComparer := Custom(
			func(a, b Point) bool {
				const epsilon = 1e-9
				return math.Abs(a.X-b.X) < epsilon && math.Abs(a.Y-b.Y) < epsilon
			},
			func(p Point) uint64 {
				x := int64(p.X * 1e6)
				y := int64(p.Y * 1e6)
				return uint64(x*31 + y)
			},
		)

		point1 := Point{X: 1.0, Y: 2.0}
		point2 := Point{X: 1.0 + 1e-10, Y: 2.0 - 1e-10}
		point3 := Point{X: 1.1, Y: 2.0}

		if !toleranceComparer.Equals(point1, point2) {
			t.Error("Expected points within tolerance to be equal")
		}

		if toleranceComparer.Equals(point1, point3) {
			t.Error("Expected points outside tolerance to be unequal")
		}
	})

	t.Run("custom comparer with case-insensitive string comparison", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name string
			Age  int
		}

		caseInsensitiveComparer := Custom(
			func(a, b User) bool {
				return strings.EqualFold(a.Name, b.Name) && a.Age == b.Age
			},
			func(u User) uint64 {
				return hashcode.Combine(strings.ToLower(u.Name), u.Age)
			},
		)

		user1 := User{Name: "Alice", Age: 25}
		user2 := User{Name: "alice", Age: 25}
		user3 := User{Name: "Alice", Age: 26}

		if !caseInsensitiveComparer.Equals(user1, user2) {
			t.Error("Expected case-insensitive name match to be equal")
		}

		if caseInsensitiveComparer.Equals(user1, user3) {
			t.Error("Expected different ages to make users unequal")
		}

		if caseInsensitiveComparer.GetHashCode(user1) != caseInsensitiveComparer.GetHashCode(user2) {
			t.Error("Expected same hash codes for equal elements")
		}
	})

	t.Run("custom comparer with business logic", func(t *testing.T) {
		t.Parallel()
		type Order struct {
			ID          int
			ExternalRef string
			Status      string
		}

		orderComparer := Custom(
			func(a, b Order) bool {
				return a.ID == b.ID || (a.ExternalRef != "" && a.ExternalRef == b.ExternalRef)
			},
			func(o Order) uint64 {
				return hashcode.Compute(o.ID)
			},
		)

		order1 := Order{ID: 1, ExternalRef: "EXT001", Status: "pending"}
		order2 := Order{ID: 1, ExternalRef: "EXT002", Status: "shipped"}
		order3 := Order{ID: 2, ExternalRef: "EXT001", Status: "cancelled"}
		order4 := Order{ID: 3, ExternalRef: "EXT003", Status: "pending"}

		if !orderComparer.Equals(order1, order2) {
			t.Error("Expected orders with same ID to be equal")
		}

		if !orderComparer.Equals(order1, order3) {
			t.Error("Expected orders with same external reference to be equal")
		}

		if orderComparer.Equals(order1, order4) {
			t.Error("Expected orders with different ID and external reference to be unequal")
		}
	})

	t.Run("custom comparer hash consistency", func(t *testing.T) {
		t.Parallel()
		type Simple struct {
			Value int
		}

		custom := Custom(
			func(a, b Simple) bool { return a.Value == b.Value },
			func(s Simple) uint64 { return hashcode.Compute(s.Value) },
		)

		s1 := Simple{Value: 42}
		s2 := Simple{Value: 42}

		if !custom.Equals(s1, s2) {
			t.Error("Expected simples with same value to be equal")
		}

		if custom.GetHashCode(s1) != custom.GetHashCode(s2) {
			t.Error("Expected same hash codes for equal elements")
		}
	})
}

func TestEqualityComparerCombinations(t *testing.T) {
	t.Run("ByField with Custom field selector", func(t *testing.T) {
		t.Parallel()
		type Complex struct {
			Data []int
		}

		lengthComparer := ByField(func(c Complex) int { return len(c.Data) })

		c1 := Complex{Data: []int{1, 2, 3}}
		c2 := Complex{Data: []int{4, 5, 6}}
		c3 := Complex{Data: []int{1, 2}}

		if !lengthComparer.Equals(c1, c2) {
			t.Error("Expected complexes with same data length to be equal")
		}

		if lengthComparer.Equals(c1, c3) {
			t.Error("Expected complexes with different data lengths to be unequal")
		}

		if lengthComparer.GetHashCode(c1) != lengthComparer.GetHashCode(c2) {
			t.Error("Expected same hash codes for equal elements")
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
		customTagComparer := Custom(
			func(a, b Record) bool {
				return len(a.Tags) == len(b.Tags)
			},
			func(r Record) uint64 {
				return hashcode.Compute(len(r.Tags))
			},
		)

		composite := Composite(idComparer, customTagComparer)

		r1 := Record{ID: 1, Name: "Record1", Tags: []string{"tag1", "tag2"}}
		r2 := Record{ID: 1, Name: "Record2", Tags: []string{"tag3", "tag4"}}
		r3 := Record{ID: 2, Name: "Record3", Tags: []string{"tag1", "tag2"}}
		r4 := Record{ID: 1, Name: "Record4", Tags: []string{"tag1"}}

		if !composite.Equals(r1, r2) {
			t.Error("Expected records with same ID and tag count to be equal")
		}

		if composite.Equals(r1, r3) {
			t.Error("Expected records with different IDs to be unequal")
		}

		if composite.Equals(r1, r4) {
			t.Error("Expected records with different tag counts to be unequal")
		}
	})

	t.Run("hash code distribution", func(t *testing.T) {
		t.Parallel()
		type Item struct {
			ID int
		}

		comparer := ByField(func(i Item) int { return i.ID })

		hashes := make(map[uint64]bool)
		collisions := 0

		for i := 0; i < 1000; i++ {
			item := Item{ID: i}
			hash := comparer.GetHashCode(item)

			if hashes[hash] {
				collisions++
			} else {
				hashes[hash] = true
			}
		}

		if collisions > 100 {
			t.Errorf("Too many hash collisions: %d out of 1000", collisions)
		}
	})
}

func TestDefault(t *testing.T) {
	t.Run("default comparer for int", func(t *testing.T) {
		t.Parallel()
		comparer := Default[int]()

		if !comparer.Equals(42, 42) {
			t.Error("Expected same integers to be equal")
		}

		if comparer.Equals(42, 43) {
			t.Error("Expected different integers to be unequal")
		}

		if comparer.GetHashCode(42) != comparer.GetHashCode(42) {
			t.Error("Expected same hash codes for equal integers")
		}
	})

	t.Run("default comparer for string", func(t *testing.T) {
		t.Parallel()
		comparer := Default[string]()

		if !comparer.Equals("hello", "hello") {
			t.Error("Expected same strings to be equal")
		}

		if comparer.Equals("hello", "world") {
			t.Error("Expected different strings to be unequal")
		}

		if comparer.GetHashCode("hello") != comparer.GetHashCode("hello") {
			t.Error("Expected same hash codes for equal strings")
		}
	})

	t.Run("default comparer for bool", func(t *testing.T) {
		t.Parallel()
		comparer := Default[bool]()

		if !comparer.Equals(true, true) {
			t.Error("Expected same booleans to be equal")
		}

		if comparer.Equals(true, false) {
			t.Error("Expected different booleans to be unequal")
		}

		if comparer.GetHashCode(true) != comparer.GetHashCode(true) {
			t.Error("Expected same hash codes for equal booleans")
		}

		if comparer.GetHashCode(true) == comparer.GetHashCode(false) {
			t.Log("Warning: true and false have same hash code (possible but not ideal)")
		}
	})

	t.Run("default comparer for float", func(t *testing.T) {
		t.Parallel()
		comparer := Default[float64]()

		if !comparer.Equals(1.5, 1.5) {
			t.Error("Expected same floats to be equal")
		}

		if comparer.Equals(1.5, 2.5) {
			t.Error("Expected different floats to be unequal")
		}

		if comparer.GetHashCode(1.5) != comparer.GetHashCode(1.5) {
			t.Error("Expected same hash codes for equal floats")
		}
	})

	t.Run("default comparer for struct", func(t *testing.T) {
		t.Parallel()
		type SimpleStruct struct {
			A int
			B string
		}

		comparer := Default[SimpleStruct]()

		s1 := SimpleStruct{A: 1, B: "test"}
		s2 := SimpleStruct{A: 1, B: "test"}
		s3 := SimpleStruct{A: 2, B: "test"}

		if !comparer.Equals(s1, s2) {
			t.Error("Expected identical structs to be equal")
		}

		if comparer.Equals(s1, s3) {
			t.Error("Expected different structs to be unequal")
		}

		if comparer.GetHashCode(s1) != comparer.GetHashCode(s2) {
			t.Error("Expected same hash codes for equal structs")
		}
	})

	t.Run("default comparer for pointer", func(t *testing.T) {
		t.Parallel()
		comparer := Default[*int]()

		value1, value2 := 42, 42
		ptr1 := &value1
		ptr2 := &value2
		var nilPtr *int

		if !comparer.Equals(ptr1, ptr1) {
			t.Error("Expected same pointer to be equal")
		}

		if comparer.Equals(ptr1, ptr2) {
			t.Error("Expected different pointers to be unequal")
		}

		if comparer.Equals(ptr1, nilPtr) {
			t.Error("Expected pointer and nil to be unequal")
		}

		if comparer.GetHashCode(ptr1) != comparer.GetHashCode(ptr1) {
			t.Error("Expected same hash codes for equal pointers")
		}
	})

	t.Run("default comparer for array", func(t *testing.T) {
		t.Parallel()
		comparer := Default[[3]int]()

		arr1 := [3]int{1, 2, 3}
		arr2 := [3]int{1, 2, 3}
		arr3 := [3]int{1, 2, 4}

		if !comparer.Equals(arr1, arr2) {
			t.Error("Expected identical arrays to be equal")
		}

		if comparer.Equals(arr1, arr3) {
			t.Error("Expected different arrays to be unequal")
		}

		if comparer.GetHashCode(arr1) != comparer.GetHashCode(arr2) {
			t.Error("Expected same hash codes for equal arrays")
		}
	})

	t.Run("default comparer for complex number", func(t *testing.T) {
		t.Parallel()
		comparer := Default[complex128]()

		c1 := complex(1.5, 2.5)
		c2 := complex(1.5, 2.5)
		c3 := complex(1.5, 3.5)

		if !comparer.Equals(c1, c2) {
			t.Error("Expected same complex numbers to be equal")
		}

		if comparer.Equals(c1, c3) {
			t.Error("Expected different complex numbers to be unequal")
		}

		if comparer.GetHashCode(c1) != comparer.GetHashCode(c2) {
			t.Error("Expected same hash codes for equal complex numbers")
		}
	})

	t.Run("default comparer hash consistency", func(t *testing.T) {
		t.Parallel()
		comparer := Default[int]()

		value := 12345
		hash1 := comparer.GetHashCode(value)
		hash2 := comparer.GetHashCode(value)
		hash3 := comparer.GetHashCode(value)

		if hash1 != hash2 || hash2 != hash3 {
			t.Error("Hash codes should be consistent for same value")
		}
	})

	t.Run("default comparer reflexive property", func(t *testing.T) {
		t.Parallel()
		comparer := Default[string]()

		value := "test value"
		if !comparer.Equals(value, value) {
			t.Error("Equals should be reflexive")
		}
	})

	t.Run("default comparer symmetric property", func(t *testing.T) {
		t.Parallel()
		comparer := Default[int]()

		a, b := 42, 42
		if comparer.Equals(a, b) != comparer.Equals(b, a) {
			t.Error("Equals should be symmetric")
		}

		a, b = 42, 43
		if comparer.Equals(a, b) != comparer.Equals(b, a) {
			t.Error("Equals should be symmetric for unequal values")
		}
	})

	t.Run("default comparer transitive property", func(t *testing.T) {
		t.Parallel()
		comparer := Default[string]()

		a := "same"
		b := "same"
		c := "same"

		if comparer.Equals(a, b) && comparer.Equals(b, c) {
			if !comparer.Equals(a, c) {
				t.Error("Equals should be transitive")
			}
		}
	})

	t.Run("default comparer hash-equals contract", func(t *testing.T) {
		t.Parallel()
		comparer := Default[complex128]()

		val1 := complex(1.5, 2.5)
		val2 := complex(1.5, 2.5)

		if comparer.Equals(val1, val2) {
			if comparer.GetHashCode(val1) != comparer.GetHashCode(val2) {
				t.Error("Equal values must have equal hash codes")
			}
		}
	})

	t.Run("default comparer with zero values", func(t *testing.T) {
		t.Parallel()

		intComparer := Default[int]()
		if !intComparer.Equals(0, 0) {
			t.Error("Expected zero integers to be equal")
		}

		stringComparer := Default[string]()
		if !stringComparer.Equals("", "") {
			t.Error("Expected empty strings to be equal")
		}

		boolComparer := Default[bool]()
		if !boolComparer.Equals(false, false) {
			t.Error("Expected false booleans to be equal")
		}

		type TestStruct struct {
			A int
			B string
		}
		structComparer := Default[TestStruct]()
		var zeroStruct TestStruct
		if !structComparer.Equals(zeroStruct, zeroStruct) {
			t.Error("Expected zero structs to be equal")
		}
	})

	t.Run("default comparer different types consistency", func(t *testing.T) {
		t.Parallel()

		intComparer := Default[int]()
		stringComparer := Default[string]()
		boolComparer := Default[bool]()

		if !intComparer.Equals(1, 1) || intComparer.Equals(1, 2) {
			t.Error("Int comparer not working correctly")
		}

		if !stringComparer.Equals("a", "a") || stringComparer.Equals("a", "b") {
			t.Error("String comparer not working correctly")
		}

		if !boolComparer.Equals(true, true) || boolComparer.Equals(true, false) {
			t.Error("Bool comparer not working correctly")
		}
	})

	t.Run("default comparer array of structs", func(t *testing.T) {
		t.Parallel()

		type Point struct {
			X, Y int
		}

		comparer := Default[[2]Point]()

		arr1 := [2]Point{{1, 2}, {3, 4}}
		arr2 := [2]Point{{1, 2}, {3, 4}}
		arr3 := [2]Point{{1, 2}, {3, 5}}

		if !comparer.Equals(arr1, arr2) {
			t.Error("Expected identical arrays of structs to be equal")
		}

		if comparer.Equals(arr1, arr3) {
			t.Error("Expected different arrays of structs to be unequal")
		}
	})

	t.Run("default comparer channel types", func(t *testing.T) {
		t.Parallel()

		comparer := Default[chan int]()

		ch1 := make(chan int)
		ch2 := make(chan int)
		var nilCh chan int

		if !comparer.Equals(ch1, ch1) {
			t.Error("Expected same channel to be equal")
		}

		if comparer.Equals(ch1, ch2) {
			t.Error("Expected different channels to be unequal")
		}

		if comparer.Equals(ch1, nilCh) {
			t.Error("Expected channel and nil to be unequal")
		}
	})
}

func BenchmarkDefault(b *testing.B) {
	b.Run("Default int comparer Equals", func(b *testing.B) {
		comparer := Default[int]()
		a, bVal := 42, 42

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = comparer.Equals(a, bVal)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
		}
	})

	b.Run("Default string comparer Equals", func(b *testing.B) {
		comparer := Default[string]()
		a, bVal := "test", "test"

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = comparer.Equals(a, bVal)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
		}
	})

	b.Run("Default int comparer GetHashCode", func(b *testing.B) {
		comparer := Default[int]()
		value := 12345

		var hash uint64
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hash = comparer.GetHashCode(value)
		}
		b.StopTimer()
		_ = hash // Prevent optimization
	})

	b.Run("Default string comparer GetHashCode", func(b *testing.B) {
		comparer := Default[string]()
		value := "benchmark test string"

		var hash uint64
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hash = comparer.GetHashCode(value)
		}
		b.StopTimer()
		_ = hash // Prevent optimization
	})

	b.Run("Default struct comparer Equals", func(b *testing.B) {
		type SimpleStruct struct {
			A int
			B string
		}

		comparer := Default[SimpleStruct]()
		s1 := SimpleStruct{A: 1, B: "test"}
		s2 := SimpleStruct{A: 1, B: "test"}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = comparer.Equals(s1, s2)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
		}
	})

	b.Run("Default struct comparer GetHashCode", func(b *testing.B) {
		type SimpleStruct struct {
			A int
			B string
		}

		comparer := Default[SimpleStruct]()
		s := SimpleStruct{A: 1, B: "test"}

		var hash uint64
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hash = comparer.GetHashCode(s)
		}
		b.StopTimer()
		_ = hash // Prevent optimization
	})

	b.Run("Default array comparer", func(b *testing.B) {
		comparer := Default[[5]int]()
		arr1 := [5]int{1, 2, 3, 4, 5}
		arr2 := [5]int{1, 2, 3, 4, 5}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = comparer.Equals(arr1, arr2)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
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
		user1 := User{ID: 42, Name: "Alice"}
		user2 := User{ID: 42, Name: "Bob"}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = comparer.Equals(user1, user2)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
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

		person1 := Person{FirstName: "John", LastName: "Doe", Age: 30}
		person2 := Person{FirstName: "John", LastName: "Doe", Age: 35}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = composite.Equals(person1, person2)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
		}
	})

	b.Run("Custom comparer", func(b *testing.B) {
		type Point struct {
			X, Y float64
		}

		custom := Custom(
			func(a, b Point) bool {
				return a.X == b.X && a.Y == b.Y
			},
			func(p Point) uint64 {
				return hashcode.Combine(p.X, p.Y)
			},
		)

		point1 := Point{X: 1.0, Y: 2.0}
		point2 := Point{X: 1.0, Y: 2.0}

		result := false
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = custom.Equals(point1, point2)
		}
		b.StopTimer()
		if !result {
			b.Fatal("Expected true")
		}
	})

	b.Run("Hash code generation", func(b *testing.B) {
		type Data struct {
			ID   int
			Name string
		}

		comparer := ByField(func(d Data) int { return d.ID })
		data := Data{ID: 42, Name: "test"}

		var hash uint64
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hash = comparer.GetHashCode(data)
		}
		b.StopTimer()
		_ = hash // Prevent optimization
	})
}
