package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
)

func TestHashSet_Add(t *testing.T) {
	t.Run("Add new items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		result1 := hs.add("apple")
		result2 := hs.add("banana")
		result3 := hs.add("cherry")

		if !result1 {
			t.Error("expected add to return true for new item")
		}
		if !result2 {
			t.Error("expected add to return true for new item")
		}
		if !result3 {
			t.Error("expected add to return true for new item")
		}
	})

	t.Run("Add duplicate items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		result1 := hs.add("apple")
		result2 := hs.add("apple")

		if !result1 {
			t.Error("expected add to return true for new item")
		}
		if result2 {
			t.Error("expected add to return false for duplicate")
		}
	})

	t.Run("Add items with hash collision", func(t *testing.T) {
		t.Parallel()
		collisionComparer := comparer.Custom(
			func(a, b string) bool {
				return a == b
			},
			func(s string) uint64 {
				return 42
			},
		)

		hs := newHashSet(collisionComparer)

		result1 := hs.add("apple")
		result2 := hs.add("banana")
		result3 := hs.add("cherry")

		if !result1 {
			t.Error("expected add to return true for new item")
		}
		if !result2 {
			t.Error("expected add to return true for new item")
		}
		if !result3 {
			t.Error("expected add to return true for new item")
		}

		result4 := hs.add("apple")
		if result4 {
			t.Error("expected add to return false for duplicate")
		}
	})

	t.Run("Add struct", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Id   int
			Name string
		}
		productComparer := comparer.Custom(
			func(a, b Product) bool {
				return a.Name == b.Name
			},
			func(p Product) uint64 {
				return uint64(p.Id)
			},
		)

		hs := newHashSet(productComparer)

		result1 := hs.add(Product{
			Id:   1,
			Name: "apple",
		})
		result2 := hs.add(Product{
			Id:   1,
			Name: "banana",
		})
		result3 := hs.add(Product{
			Id:   2,
			Name: "cherry",
		})
		result4 := hs.add(Product{
			Id:   1,
			Name: "banana",
		})

		if !result1 {
			t.Error("expected add to return true for new item")
		}
		if !result2 {
			t.Error("expected add to return true for new item")
		}
		if !result3 {
			t.Error("expected add to return true for new item")
		}
		if result4 {
			t.Error("expected add to return false for duplicate")
		}

		result5 := hs.add(Product{
			Id:   2,
			Name: "cherry",
		})
		if result5 {
			t.Error("expected add to return false for duplicate")
		}
	})

	t.Run("Add integer items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[int]()
		hs := newHashSet(eqComparer)

		result1 := hs.add(1)
		result2 := hs.add(2)
		result3 := hs.add(1)

		if !result1 {
			t.Error("expected add to return true for new item")
		}
		if !result2 {
			t.Error("expected add to return true for new item")
		}
		if result3 {
			t.Error("expected add to return false for duplicate")
		}
	})
}

func TestHashSet_Contains(t *testing.T) {
	t.Run("Contains existing items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		hs.add("apple")
		hs.add("banana")
		hs.add("cherry")

		if !hs.contains("apple") {
			t.Error("expected contains to return true for existing item")
		}
		if !hs.contains("banana") {
			t.Error("expected contains to return true for existing item")
		}
		if !hs.contains("cherry") {
			t.Error("expected contains to return true for existing item")
		}
	})

	t.Run("Contains non-existing items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		hs.add("apple")
		hs.add("banana")

		if hs.contains("cherry") {
			t.Error("expected contains to return false for non-existing item")
		}
		if hs.contains("date") {
			t.Error("expected contains to return false for non-existing item")
		}
		if hs.contains("") {
			t.Error("expected contains to return false for non-existing item")
		}
	})

	t.Run("Contains with hash collision", func(t *testing.T) {
		t.Parallel()
		collisionComparer := comparer.Custom(
			func(a, b string) bool {
				return a == b
			},
			func(s string) uint64 {
				return 42
			},
		)

		hs := newHashSet(collisionComparer)

		hs.add("apple")
		hs.add("banana")
		hs.add("cherry")

		if !hs.contains("apple") {
			t.Error("expected contains to return true for existing item")
		}
		if !hs.contains("banana") {
			t.Error("expected contains to return true for existing item")
		}
		if !hs.contains("cherry") {
			t.Error("expected contains to return true for existing item")
		}

		if hs.contains("date") {
			t.Error("expected contains to return false for non-existing item")
		}
		if hs.contains("grape") {
			t.Error("expected contains to return false for non-existing item")
		}
	})

	t.Run("Contains empty hashset", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		if hs.contains("anything") {
			t.Error("expected contains to return false for empty hashset")
		}
		if hs.contains("") {
			t.Error("expected contains to return false for empty hashset")
		}
	})

	t.Run("Contains struct", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Id   int
			Name string
		}
		productComparer := comparer.Custom(
			func(a, b Product) bool {
				return a.Name == b.Name
			},
			func(p Product) uint64 {
				return uint64(p.Id)
			},
		)

		hs := newHashSet(productComparer)

		product1 := Product{
			Id:   1,
			Name: "apple",
		}
		hs.add(product1)
		product2 := Product{
			Id:   1,
			Name: "banana",
		}
		hs.add(product2)
		product3 := Product{
			Id:   2,
			Name: "cherry",
		}
		hs.add(product3)
		product4 := Product{
			Id:   3,
			Name: "plum",
		}
		product5 := Product{
			Id:   2,
			Name: "grape",
		}

		if !hs.contains(product1) {
			t.Error("expected contains to return true for existing item")
		}
		if !hs.contains(product2) {
			t.Error("expected contains to return true for existing item")
		}
		if !hs.contains(product3) {
			t.Error("expected contains to return true for existing item")
		}
		if hs.contains(product4) {
			t.Error("expected contains to return false for non-existing item")
		}
		if hs.contains(product5) {
			t.Error("expected contains to return false for non-existing item")
		}
	})
}
