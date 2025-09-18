package enumerable

// ForEach executes the specified action for each element in the enumeration.
// This operation is useful for performing side effects like printing, logging,
// or modifying external state for each element.
//
// The for each operation will:
//   - Execute the action function for each element in the enumeration
//   - Process all elements in the enumeration
//   - Handle nil enumerators gracefully
//   - Not return any value (void operation)
//
// Parameters:
//
//	action - the action to execute for each element
//
// ⚠️ Performance note: This operation must iterate through the entire
// enumeration, which may be expensive for large enumerations.
//
// ⚠️ Side effects warning: The action function may have side effects.
// Use with caution in functional programming contexts.
//
// Notes:
//   - If the enumerator is nil, no action is executed
//   - This is a terminal operation that materializes the enumeration
//   - All elements are processed regardless of action behavior
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(1) - constant space usage
//   - The enumeration stops only when exhausted or if upstream operations stop it
//   - Action function should handle all possible values including zero values
func (q Enumerator[T]) ForEach(action func(T)) {
	forEachInternal(q, action)
}

// ForEach executes the specified action for each element in the enumeration.
// This operation is useful for performing side effects like printing, logging,
// or modifying external state for each element.
//
// The for each operation will:
//   - Execute the action function for each element in the enumeration
//   - Process all elements in the enumeration
//   - Handle nil enumerators gracefully
//   - Not return any value (void operation)
//
// Parameters:
//
//	action - the action to execute for each element
//
// ⚠️ Performance note: This operation must iterate through the entire
// enumeration, which may be expensive for large enumerations.
//
// ⚠️ Side effects warning: The action function may have side effects.
// Use with caution in functional programming contexts.
//
// Notes:
//   - If the enumerator is nil, no action is executed
//   - This is a terminal operation that materializes the enumeration
//   - All elements are processed regardless of action behavior
//   - Time complexity: O(n) where n is the number of elements
//   - Space complexity: O(1) - constant space usage
//   - The enumeration stops only when exhausted or if upstream operations stop it
//   - Action function should handle all possible values including zero values
func (q EnumeratorAny[T]) ForEach(action func(T)) {
	forEachInternal(q, action)
}

// ForEach executes the specified action for each element in the sorted enumeration.
// This operation is useful for performing side effects like printing, logging,
// or modifying external state for each element in sorted order.
//
// The for each operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Execute the action function for each element in sorted order
//   - Process all elements in the enumeration in sorted order
//   - Handle nil enumerators gracefully
//   - Not return any value (void operation)
//
// Parameters:
//
//	action - the action to execute for each element in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// This operation must iterate through the entire sorted enumeration, which may be expensive
// for large enumerations. Time complexity: O(n log n) for sorting + O(n) for iteration = O(n log n).
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// ⚠️ Side effects warning: The action function may have side effects.
// Use with caution in functional programming contexts.
// Actions are executed in sorted order according to all accumulated sorting rules.
//
// Notes:
//   - If the enumerator is nil, no action is executed
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - This is a terminal operation that materializes the enumeration
//   - Elements are processed in sorted order according to all accumulated sorting rules
//   - All elements are processed regardless of action behavior
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n) during sorting, O(1) during iteration
//   - The enumeration stops only when exhausted or if upstream operations stop it
//   - Action function should handle all possible values including zero values
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
func (o OrderEnumerator[T]) ForEach(action func(T)) {
	forEachInternal(o.getSortedEnumerator(), action)
}

// ForEach executes the specified action for each element in the sorted enumeration.
// This operation is useful for performing side effects like printing, logging,
// or modifying external state for each element in sorted order.
// This method supports any type T, including non-comparable types with custom sorting logic.
//
// The for each operation will:
//   - Execute deferred sorting rules to determine the sorted order
//   - Execute the action function for each element in sorted order
//   - Process all elements in the enumeration in sorted order
//   - Handle nil enumerators gracefully
//   - Not return any value (void operation)
//
// Parameters:
//
//	action - the action to execute for each element in sorted order
//
// ⚠️ Performance note: This is a terminal operation that triggers actual sorting computation.
// This operation must iterate through the entire sorted enumeration, which may be expensive
// for large enumerations. Time complexity: O(n log n) for sorting + O(n) for iteration = O(n log n).
//
// ⚠️ Memory note: This operation buffers all elements in memory during sorting.
// Space complexity: O(n) - allocates memory for all elements during sorting.
// No additional buffering beyond sorting phase.
//
// ⚠️ Side effects warning: The action function may have side effects.
// Use with caution in functional programming contexts.
// Actions are executed in sorted order according to all accumulated sorting rules.
//
// Notes:
//   - If the enumerator is nil, no action is executed
//   - Actual sorting computation occurs only during this first operation
//   - All accumulated sorting rules (OrderBy + ThenBy levels) are applied during sorting
//   - This is a terminal operation that materializes the enumeration
//   - Elements are processed in sorted order according to all accumulated sorting rules
//   - All elements are processed regardless of action behavior
//   - Time complexity: O(n log n) where n is the number of elements
//   - Space complexity: O(n) during sorting, O(1) during iteration
//   - The enumeration stops only when exhausted or if upstream operations stop it
//   - Action function should handle all possible values including zero values
//   - Elements are yielded in sorted order according to all accumulated sorting rules
//   - Subsequent calls will re-execute sorting (sorting is not cached)
//   - Stable sorting preserves relative order of equal elements according to sorting rules
//   - Supports complex types including structs with non-comparable fields
//   - Custom comparer functions are used for element comparison during sorting
func (o OrderEnumeratorAny[T]) ForEach(action func(T)) {
	forEachInternal(o.getSortedEnumerator(), action)
}

func forEachInternal[T any](enumerator func(func(T) bool), action func(T)) {
	if enumerator == nil {
		return
	}
	enumerator(func(item T) bool {
		action(item)
		return true
	})
}
