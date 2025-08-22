package enumerable

import "github.com/ahatornn/enumerable/comparer"

func (e Enumerator[T]) MinInt(keySelector func(T) int) (int, bool) {
	return minIntInternal(e, keySelector, comparer.ComparerInt)
}

func (e EnumeratorAny[T]) MinInt(keySelector func(T) int) (int, bool) {
	return minIntInternal(e, keySelector, comparer.ComparerInt)
}

func minIntInternal[T any](enumerator func(yield func(T) bool), keySelector func(T) int, cmp comparer.ComparerFunc[int]) (int, bool) {
	var minKey int
	found := false

	if enumerator == nil || keySelector == nil || cmp == nil {
		return 0, false
	}

	enumerator(func(item T) bool {
		key := keySelector(item)

		if !found {
			minKey = key
			found = true
		} else if cmp(key, minKey) < 0 {
			minKey = key
		}

		return true
	})

	return minKey, found
}
