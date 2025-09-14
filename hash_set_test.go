package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/comparer"
	"github.com/stretchr/testify/assert"
)

func TestHashSet_Add(t *testing.T) {
	t.Run("Add new items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		result1 := hs.add("apple")
		result2 := hs.add("banana")
		result3 := hs.add("cherry")

		assert.True(t, result1)
		assert.True(t, result2)
		assert.True(t, result3)
	})

	t.Run("Add duplicate items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		result1 := hs.add("apple")
		result2 := hs.add("apple")

		assert.True(t, result1)
		assert.False(t, result2)
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

		assert.True(t, result1)
		assert.True(t, result2)
		assert.True(t, result3)

		result4 := hs.add("apple")
		assert.False(t, result4)
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

		assert.True(t, result1)
		assert.True(t, result2)
		assert.True(t, result3)
		assert.False(t, result4)

		result5 := hs.add(Product{
			Id:   2,
			Name: "cherry",
		})
		assert.False(t, result5)
	})

	t.Run("Add integer items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[int]()
		hs := newHashSet(eqComparer)

		result1 := hs.add(1)
		result2 := hs.add(2)
		result3 := hs.add(1)

		assert.True(t, result1)
		assert.True(t, result2)
		assert.False(t, result3)
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

		assert.True(t, hs.contains("apple"))
		assert.True(t, hs.contains("banana"))
		assert.True(t, hs.contains("cherry"))
	})

	t.Run("Contains non-existing items", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		hs.add("apple")
		hs.add("banana")

		assert.False(t, hs.contains("cherry"))
		assert.False(t, hs.contains("date"))
		assert.False(t, hs.contains(""))
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

		assert.True(t, hs.contains("apple"))
		assert.True(t, hs.contains("banana"))
		assert.True(t, hs.contains("cherry"))

		assert.False(t, hs.contains("date"))
		assert.False(t, hs.contains("grape"))
	})

	t.Run("Contains empty hashset", func(t *testing.T) {
		t.Parallel()
		eqComparer := comparer.Default[string]()
		hs := newHashSet(eqComparer)

		assert.False(t, hs.contains("anything"))
		assert.False(t, hs.contains(""))
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

		assert.True(t, hs.contains(product1))
		assert.True(t, hs.contains(product2))
		assert.True(t, hs.contains(product3))
		assert.False(t, hs.contains(product4))
		assert.False(t, hs.contains(product5))
	})
}
