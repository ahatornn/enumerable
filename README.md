# enumerable: Lazy, Generic Iterators for Go
[![GoDoc](https://godoc.org/github.com/ahatornn/enumerable?status.svg)](https://godoc.org/github.com/ahatornn/enumerable)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/ahatornn/enumerable)](https://github.com/ahatornn/enumerable)
![GitHub release](https://img.shields.io/github/v/tag/ahatornn/enumerable)

[![Tests](https://github.com/ahatornn/enumerable/actions/workflows/test.yml/badge.svg)](https://github.com/ahatornn/enumerable/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ahatornn/enumerable)](https://goreportcard.com/report/github.com/ahatornn/enumerable)
[![codecov](https://codecov.io/github/ahatornn/enumerable/graph/badge.svg?token=UADCDVFYP3)](https://codecov.io/github/ahatornn/enumerable)

enumerable is a Go library for functional-style data processing using **generics** and **lazy iterators**.

It lets you filter, transform, and aggregate data in a clean, readable way — without for loops or intermediate slices. Operations are evaluated lazily, meaning no unnecessary allocations or eager processing, making it efficient even for large or streaming datasets.

> Composable. Lazy. Generic. Type-safe.

---

## ✨ Features

- ✅ **Functional pipeline operators**: Chainable methods like Where, Select, Take, Skip, First, Any, All, Count, and more - all evaluated lazily and fully type-safe.
- ✅ **Lazy iterators** — data is processed on-demand.
- ✅ **Generics (Go 1.18+)** — type safety without type casting.
- ✅ **Method chaining** — fluent, readable APIs.
- ✅ **Efficiency** — minimal allocations, suitable for large datasets.
- ✅ **Nil-safe operations** — methods safely handle nil sequences without panics.
- ✅ **Two enumerator types** — EnumeratorAny[T any] for any type and Enumerator[T comparable] for comparable types with enhanced functionality.
- ✅ **Set operations for comparable types** — additional methods like Distinct, Except, Intersect, and Union available when working with comparable data.

---

## 📦 Installation

Use go get to add the enumerable to your project:

bash
go get github.com/ahatornn/enumerable@latest


## ⚖️ Comparison: Enumerable vs Vanilla Go

Let’s say you want to: Merge two slices, keep numbers greater than 50, skip the first 3, and take the next 2.

### ✅ Using enumerable

```go
nums1 := []int{19, 63, 25, 19, 47, 91, 58}
nums2 := []int{80, 25, 91, 36, 19, 52, 74, 36}

result := enumerable.FromSlice(nums1).
	Union(enumerable.FromSlice(nums2)).
	Where(func(x int) bool { return x > 50 }).
	Skip(3).
	Take(2).
	ToSlice()

fmt.Println(result) // Result: [80 52]
```

### 🛠 Manual

```go
seen := make(map[int]bool)
var filtered []int
for _, x := range a {
	if x > 50 && !seen[x] {
		seen[x] = true
		filtered = append(filtered, x)
	}
}
for _, x := range b {
	if x > 50 && !seen[x] {
		seen[x] = true
		filtered = append(filtered, x)
	}
}
skip := 3
if skip >= len(filtered) {
	return nil
}
afterSkip := filtered[skip:]
take := 2
if take > len(afterSkip) {
	take = len(afterSkip)
}
return afterSkip[:take]
```

### 🚀 Performance Benchmarks

Comparison between the enumerable and manual implementation across different dataset sizes (go 1.24.1, cpu: 12th Gen Intel(R) Core(TM) i7-12700F, goos: windows, goarch: amd64):
| Records | enumerable (ns/op) | Manual (ns/op) | Speed Ratio | Memory (enumerable → Manual) | Allocs (enumerable → Manual) |
|-|-:|-:|-:|-:|-:|
| 10 | 809 | 382 | 0.47x | 1.2KB → 0.6KB | 12 → 8 |
| 50 | 2,104 | 1,893 | 0.9x  | 2.4KB → 3.1KB | 14 → 14 |
| 100 | 2,945 | 6,611 | 2.24x | 4.7KB → 13.4KB | 16 → 20 |
| 1,000 | 3,023 | 93,675 | 31x | 4.7KB → 208KB | 16 → 43 |
| 5,000 | 3,092 | 389,795 | 126x | 4.7KB → 817KB | 16 → 93 |
#### Key Insights
1. **Small datasets (≤50 records)**:
   - Minimal difference (manual 0.9-2x faster)
   - enumerable uses less memory (-23% at 50 records)

2. **Medium/large datasets (≥100 records)**:
   - enumerable is **2-126x faster**
   - Saves **up to 99% memory** (at 5,000 records)
   - Consistent execution time (~3 ns) regardless of size

## ⚙️ Methods
### 🌱 Creation Methods
Initialize new sequences from scratch.
- Often act as entry points to the API.
- May accept minimal input (or none) to generate data.
- For each method that works with comparable types, there is a corresponding [name]Any variant that works with any types.

| Method | Description |
|--|--|
| [Empty](./empty.go) | Returns an empty [Enumerator[T]](./enumerator.go) of type T that yields no values with comparable types only. |
| [EmptyAny](./empty.go) | Returns an empty [EnumeratorAny[T]](./enumerator.go) of type T that yields no values with any types. |
| [FromChannel](./from_channel.go) | Creates an [Enumerator[T]](./enumerator.go) that yields values received from a channel.<br>The enumeration continues until the channel is closed or the consumer stops iteration with comparable types only. |
| [FromChannelAny](./from_channel.go) | Creates an [EnumeratorAny[T]](./enumerator.go) that yields values received from a channel.<br>The enumeration continues until the channel is closed or the consumer stops iteration with any types. |
| [FromSlice](./from_slice.go) | Creates an [Enumerator[T]](./enumerator.go) that yields all elements from the input slice in order with comparable types only. |
| [FromSliceAny](./from_slice.go) | Creates an [EnumeratorAny[T]](./enumerator.go) that yields all elements from the input slice in order with any types. |
| [Range](./range.go) | Generates a sequence of consecutive integers starting at 'start', producing exactly 'count' values in ascending order (with step +1) with comparable types only. |
| [RangeAny](./range.go) | Generates an [EnumeratorAny[T]](./enumerator.go) of consecutive integers starting at 'start', producing exactly 'count' values in ascending order (with step +1) with any types. |
| [Repeat](./repeat.go) | Generates a sequence containing the same item repeated 'count' times with comparable types only. |
| [RepeatAny](./repeat.go) | Generates an [EnumeratorAny[T]](./enumerator.go) containing the same item repeated 'count' times with any types. |

## ⏳ Lazy Methods
Perform deferred computations, evaluating only when needed
- Chainable without intermediate allocations.
- Save memory/CPU until iteration.

| Method | Description |
|--|--|
| [Concat](./concat.go) | Combines two enumerations into a single enumeration. The resulting enumeration yields all elements from the first enumeration, followed by all elements from the second enumeration. |
| [DefaultIfEmpty](./default_if_empty.go) | Returns an enumeration that contains the elements of the current enumeration, or a single default value if the current enumeration is empty. |
| [Distinct](./distinct.go) | Returns an enumerator that yields only unique elements from the original enumeration. Each element appears only once in the result, regardless of how many times it appears in the source.<br/><br/>⚠️ Performance note: This operation buffers all unique elements encountered so far in memory. For enumerations with many unique elements, memory usage can become significant. The operation is not memory-bounded<br/>ℹ️ Available only for comparable types. |
| [Except](./except.go) | Returns an enumerator that yields elements from the first enumeration that are not present in the second enumeration. This is equivalent to set difference operation (first - second).<br/><br/>⚠️ Performance note: This operation completely buffers the second enumerator into memory (creates a map for fast lookup). For large second enumerations, this may consume significant memory. The memory usage is proportional to the number of unique elements in the second enumerator.<br/>ℹ️ Available only for comparable types. |
| [Intersect](./intersect.go) |Returns an enumerator that yields elements present in both enumerations. This is equivalent to set intersection operation (first ∩ second).<br/><br/>⚠️ Performance note: The second enumeration is completely loaded into memory to enable fast lookups. Be cautious when using this with very large second enumerations as it may cause high memory usage.<br/>ℹ️ Available only for comparable types. |
| [Skip](./skip.go) | Bypasses a specified number of elements in an enumeration and then yields the remaining elements. This operation is useful for pagination, skipping headers, or bypassing initial elements. |
| [SkipLast](./skip_last.go) | Bypasses a specified number of elements at the end of an enumeration and yields the remaining elements. This operation is useful for removing trailing elements like footers, summaries, or fixed-size endings from sequences.<br/><br/>⚠️ Performance note: This operation buffers up to n elements in memory using a circular buffer for efficient memory usage. For large values of n, this may consume significant memory.<br/>⚠️ Evaluation note: This operation is partially lazy - elements are processed as the enumeration proceeds, but the last n elements are buffered and never yielded. The enumeration must progress beyond n elements to yield earlier ones. |
| [SkipWhile](./skip_while.go) | Bypasses elements in an enumeration as long as a specified condition is true and then yields the remaining elements. This operation is useful for skipping elements until a certain condition is met. |
| [Take](./take.go) | Returns a specified number of contiguous elements from the start of an enumeration. This operation is useful for pagination, limiting results, or taking samples from sequences. |
| [TakeLast](./take_last.go) | Returns a specified number of contiguous elements from the end of an enumeration. This operation is useful for getting the final elements, such as last N records, trailing averages, or end-of-sequence markers.<br/><br/>⚠️ Performance note: This operation buffers up to n elements in memory to track which elements should be yielded. For large values of n, this may consume significant memory. All elements must be processed before any are yielded.<br/>⚠️ Evaluation note: This operation is not lazy in the traditional sense - the entire source enumeration must be consumed before yielding begins. Elements are yielded in order of their appearance in the original enumeration. |
| [TakeWhile](./take_while.go) | Returns elements from an enumeration as long as a specified condition is true. This operation is useful for taking elements until a certain condition is met, such as taking elements while they are valid or within a range. |
| [Union](./union.go) | Produces the set union of two enumerations by using the default equality comparer. This operation returns unique elements that appear in either enumeration.<br/><br/>⚠️ Performance note: This operation buffers all unique elements encountered so far in memory. For enumerations with many unique elements, memory usage can become significant. The operation is not memory-bounded.<br/>ℹ️ Available only for comparable types. |
| [Where](./where.go) | Filters an enumeration based on a predicate function. This operation returns elements that satisfy the specified condition. |

## 🪄 Materialize Methods
Convert lazy sequences into concrete data structures.
- Trigger all pending computations.
- Require memory to store results.

| Method | Description |
|--|--|
| [All](./all.go) | Determines whether all elements in the enumeration satisfy a predicate. Returns true if every element matches the predicate, or if the enumeration is empty. |
| [Any](./any.go) | Determines whether an enumeration contains any elements. This operation is useful for checking if a sequence is non-empty. |
| [AverageInt](./average.go) | AverageInt returns the arithmetic mean of integer values extracted from elements of the enumeration using a key selector function.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to calculate the sum and count. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [AverageInt64](./average.go) | AverageInt returns the arithmetic mean of integer64 values extracted from elements of the enumeration using a key selector function.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to calculate the sum and count. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [AverageFloat](./average.go) | AverageInt returns the arithmetic mean of float32 values extracted from elements of the enumeration using a key selector function.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to calculate the sum and count. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [AverageFloat64](./average.go) | AverageInt returns the arithmetic mean of float64 values extracted from elements of the enumeration using a key selector function.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to calculate the sum and count. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [Count](./any.go) | Returns the number of elements in an enumeration. This operation is useful for determining the size of a sequence.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to count all elements. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [ElementAt](./element_at.go) | Returns the element at the specified index in the sequence. This operation is useful when you need to access an element at a specific position without iterating through the entire sequence.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the sequence until the specified index is reached.<br/>⚠️ Memory note: This operation does not buffer elements, only tracks current position. |
| [FirstOrDefault](./first_or_default.go) | Returns the first element of an enumeration, or a default value if the enumeration is empty or nil. |
| [FirstOrNil](./first_or_nil.go) | Returns a pointer to the first element of an enumeration. This operation is useful for getting the first element when it exists, with the ability to distinguish between "no elements" and "zero value" cases. |
| [ForEach](./foreach.go) | Executes the specified action for each element in the enumeration. This operation is useful for performing side effects like printing, logging, or modifying external state for each element.<br/><br/>⚠️ Performance note: This operation must iterate through the entire enumeration, which may be expensive for large enumerations.<br/>⚠️ Side effects warning: The action function may have side effects. Use with caution in functional programming contexts. |
| [LastOrDefault](./last_or_default.go) | Returns the last element of an enumeration, or a default value if the enumeration is empty or nil.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the last element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [LastOrNil](./last_or_nil.go) | Returns a pointer to the last element of an enumeration, or nil if the enumeration is empty or nil.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the last element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [LongCount](./long_count.go) | Returns the number of elements in an enumeration as an int64. This operation is useful for determining the size of large sequences where the count might exceed the range of int.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to count all elements. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxBool](./max_bool.go) | Returns the largest boolean value extracted from elements of the enumeration using a key selector function and natural boolean ordering (false < true).<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the maximum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxBy](./max_by.go) | Returns the largest element in the enumeration according to a custom comparison function.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the maximum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxByte](./max_byte.go) | Returns the largest byte value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element. Performance is optimized with early termination when 255 is found, but worst-case scenario processes all elements.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxFloat](./max_float.go) | Returns the largest float32 value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxFloat64](./max_float.go) | Returns the largest float64 value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxInt](./max_int.go) | Returns the largest integer value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxInt64](./max_int.go) | Returns the largest int64 value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxRune](./max_rune.go) | Returns the largest rune value extracted from elements of the enumeration using a key selector function and natural Unicode code point ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxString](./max_string.go) | Returns the lexicographically largest string value extracted from elements of the enumeration using a key selector function and natural string ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MaxTime](./max_time.go) | Returns the latest time.Time value extracted from elements of the enumeration using a key selector function and natural time ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the maximum element.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinBool](./min_bool.go) | Returns the smallest boolean value extracted from elements of the enumeration using a key selector function and natural boolean ordering (false < true).<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the minimum element. Performance is optimized with early termination when false is found, but worst-case scenario processes all elements.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinBy](./min_by.go) | Returns the smallest element in the enumeration according to a custom comparison function.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the minimum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinByte](./min_byte.go) | Returns the smallest byte value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the minimum element. Performance is optimized with early termination when 0 is found, but worst-case scenario processes all elements.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinFloat](./min_float.go) | Returns the smallest float32 value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the minimum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinFloat64](./min_float.go) | Returns the smallest float64 value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the minimum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinInt](./min_int.go) | Returns the smallest integer value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the minimum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinInt64](./min_int.go) | Returns the smallest int64 value extracted from elements of the enumeration using a key selector function and natural numeric ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to find the minimum element. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinRune](./min_rune.go) | Returns the smallest rune value extracted from elements of the enumeration using a key selector function and natural Unicode code point ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the minimum element. Performance is optimized with early termination when 0 is found (since 0 is the minimum Unicode code point), but worst-case scenario processes all elements.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinString](./min_string.go) | Returns the lexicographically smallest string value extracted from elements of the enumeration using a key selector function and natural string ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the minimum element. Performance is optimized with early termination when empty string is found, but worst-case scenario processes all elements.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [MinTime](./min_time.go) | Returns the earliest time.Time value extracted from elements of the enumeration using a key selector function and natural time ordering.<br/><br/>⚠️ Performance note: This is a terminal operation that iterates through the enumeration to find the minimum element. Performance is optimized with early termination when zero time is found, but worst-case scenario processes all elements.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations. |
| [Single](./single.go) | Returns the single element of a sequence using default equality comparison. This operation is useful when you expect exactly one element in the sequence.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the sequence until completion or error. For large sequences with multiple elements, this may process more elements than necessary.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations.<br/>ℹ️ Available only for comparable types. |
| [SingleBy](./single_by.go) | Returns the single element of a sequence that satisfies the provided [equality comparer](comparer/equality_comparer.go).<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the sequence until completion or second distinct element found. For large sequences with multiple distinct elements, this may process more elements than necessary.<br/>⚠️ Memory note: This operation buffers elements to perform equality comparison using hash-based lookup. Memory usage depends on the number of distinct elements encountered during processing. |
| [SingleOrDefault](./single.go) | Returns the single element of a sequence using default equality comparison, or a specified default value if the sequence is empty.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the sequence until completion or second element found. For large sequences with multiple elements, this may process more elements than necessary.<br/>⚠️ Memory note: This operation does not buffer elements, but it must process the entire enumeration, which may trigger upstream operations.<br/>ℹ️ Available only for comparable types. |
| [SumFloat](./sum_float.go) | Computes the sum of float32 values obtained by applying a selector function to each element in the enumeration. This operation is useful for calculating totals, aggregates, or numeric summaries with floating-point precision.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to sum all values. For large enumerations, this may be expensive.<br/>⚠️ Precision warning: Floating-point arithmetic may introduce small rounding errors. Consider using appropriate rounding for display. |
| [SumInt](./sum_int.go) | Computes the sum of integers obtained by applying a selector function to each element in the enumeration. This operation is useful for calculating totals, aggregates, or numeric summaries.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to sum all values. For large enumerations, this may be expensive.<br/>⚠️ Overflow warning: Integer overflow may occur with very large sums. Consider using SumInt64 for larger ranges. |
| [ToChannel](./to_channel.go) | Converts an enumeration to a channel that yields all elements. This operation is useful for converting lazy enumerations into channel-based processing pipelines or for interoperability with goroutines.<br/><br/>⚠️ Resource management: This operation starts a goroutine that runs until the enumeration is complete. Always consume all elements or the goroutine may block indefinitely.<br/>⚠️ Blocking behavior: The goroutine will block when trying to send to a full channel if there's no reader. Use appropriate buffer size.<br/>⚠️ Warning: If the returned channel is not fully consumed, the goroutine may leak. Always range over the entire channel or ensure proper cleanup. |
| [ToMap](./to_map.go) | Converts an enumeration to a map[T]struct{} for efficient set-like operations. This operation is useful for creating memory-efficient lookup collections or for removing duplicates while materializing an enumeration.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration. For large enumerations, this may be expensive.<br/>⚠️ Memory note: This operation buffers all unique elements in memory. The memory usage depends on the number of unique elements.<br/>ℹ️ Available only for comparable types. |
| [ToSlice](./to_slice.go) | Converts an enumeration to a slice containing all elements. This operation is useful for materializing lazy enumerations into concrete collections.<br/><br/>⚠️ Performance note: This is a terminal operation that must iterate through the entire enumeration to collect all elements. For large enumerations, this may be expensive in both time and memory.<br/>⚠️ Memory note: This operation buffers all elements in memory. |
