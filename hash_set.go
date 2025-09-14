package enumerable

import "github.com/ahatornn/enumerable/comparer"

const (
	defaultCapacity = 8
)

type hashSet[T any] struct {
	comparer comparer.EqualityComparer[T]
	items    map[uint64]hashSetItem[T]
}

type hashSetItem[T any] struct {
	item  T
	other []T
}

func newHashSet[T any](comparer comparer.EqualityComparer[T]) *hashSet[T] {
	return &hashSet[T]{
		comparer: comparer,
		items:    make(map[uint64]hashSetItem[T], defaultCapacity),
	}
}

func (hs *hashSet[T]) add(item T) bool {
	hash := hs.comparer.GetHashCode(item)

	if existingItem, exists := hs.items[hash]; exists {
		if hs.comparer.Equals(item, existingItem.item) {
			return false
		}

		for _, collisionItem := range existingItem.other {
			if hs.comparer.Equals(item, collisionItem) {
				return false
			}
		}

		existingItem.other = append(existingItem.other, item)
		hs.items[hash] = existingItem
	} else {
		newItem := hashSetItem[T]{
			item:  item,
			other: make([]T, 0),
		}
		hs.items[hash] = newItem
	}

	return true
}

func (hs *hashSet[T]) contains(item T) bool {
	hash := hs.comparer.GetHashCode(item)

	if existingItem, exists := hs.items[hash]; exists {
		if hs.comparer.Equals(item, existingItem.item) {
			return true
		}
		for _, collisionItem := range existingItem.other {
			if hs.comparer.Equals(item, collisionItem) {
				return true
			}
		}
	}

	return false
}
