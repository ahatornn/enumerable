# AGENTS.md — enumerable

Go library: `github.com/ahatornn/enumerable` — lazy, generic iterators (LINQ-style) for Go.

## Commands

```
go test ./...                          # run all tests
go test ./... -bench=^$                # run tests, skip benchmarks (CI uses this)
go test ./... -coverprofile=coverage.txt  # with coverage
go test -run TestName ./...            # run a single test
gofmt -w .                             # fix formatting (CI fails on bad format)
```

CI matrix: Go 1.24 and 1.25 on ubuntu-latest.

## Architecture

**Core types** (`enumerator.go`):
- `Enumerator[T comparable]` — func-based lazy iterator for comparable types
- `EnumeratorAny[T any]` — same, for any types (including slices, maps, functions)

Both are `func(yield func(T) bool)` — Go 1.22+ range-over-func.

**Dual-type pattern**: nearly every operation has two variants:
- `FromSlice`, `Where`, `Take`, etc. — for `T comparable` (uses `==`, maps)
- `FromSliceAny`, `WhereAny`, `TakeAny`, etc. — for `T any`

**Subpackages**:
- `comparer/` — `ComparerFunc[T]` (returns -1/0/1) and `EqualityComparer[T]` (like C# IEqualityComparer)
- `grouping/` — `Group[K, T]` and `groupingBuilder` for GroupBy
- `hashcode/` — hash computation for non-comparable types
- `selector/` — selector helper functions

**Key entry points**: `FromSlice[T]`, `FromSliceAny[T]`, `FromChannel[T]`, `Empty[T]`, `Range[T]`, `Repeat[T]`

**Errors** (`infrastructure.go`): `ErrNoElements`, `ErrMultipleElements`

**Sorting** (`order_enumerator.go`): `OrderEnumerator[T]` / `OrderEnumeratorAny[T]` — lazy, accumulates sort levels, sorts on first enumeration via `sort.SliceStable`.

## Conventions

- Tests use `t.Parallel()` in subtests
- Nil enumerators are handled gracefully (return empty/default, never panic)
- Lazy methods return new enumerators; materialize methods (`ToSlice`, `ToMap`, `All`, `Count`, etc.) consume the enumeration
- Set operations (`Distinct`, `Except`, `Intersect`, `Union`) buffer the second enumeration into memory — documented with warnings in README
