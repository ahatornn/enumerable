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
