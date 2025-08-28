package enumerable

import (
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestAverageInt(t *testing.T) {
	t.Run("average int from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 3.0 {
			t.Errorf("Expected average 3.0, got %f", avg)
		}
	})

	t.Run("average int from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-1, -2, -3, -4, -5})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != -3.0 {
			t.Errorf("Expected average -3.0, got %f", avg)
		}
	})

	t.Run("average int from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{-5, -2, 0, 2, 5})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 0.0 {
			t.Errorf("Expected average 0.0, got %f", avg)
		}
	})

	t.Run("average int single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{42})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 42.0 {
			t.Errorf("Expected average 42.0, got %f", avg)
		}
	})

	t.Run("average int empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{})

		avg, ok := enumerator.AverageInt(selector.Int)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if avg != 0 {
			t.Errorf("Expected average 0 for empty slice, got %f", avg)
		}
	})

	t.Run("average int nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[int] = nil

		avg, ok := enumerator.AverageInt(selector.Int)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if avg != 0 {
			t.Errorf("Expected average 0 for nil enumerator, got %f", avg)
		}
	})

	t.Run("average int with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3})

		avg, ok := enumerator.AverageInt(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if avg != 0 {
			t.Errorf("Expected average 0 for nil keySelector, got %f", avg)
		}
	})

	t.Run("average int with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price int
		}

		products := []Product{
			{Name: "A", Price: 100},
			{Name: "B", Price: 200},
			{Name: "C", Price: 300},
		}

		enumerator := FromSlice(products)
		avg, ok := enumerator.AverageInt(func(p Product) int { return p.Price })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 200.0 {
			t.Errorf("Expected average 200.0, got %f", avg)
		}
	})

	t.Run("average int with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{0, 0, 0, 0})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 0.0 {
			t.Errorf("Expected average 0.0, got %f", avg)
		}
	})

	t.Run("average int with large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1000000, 2000000, 3000000})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 2000000.0 {
			t.Errorf("Expected average 2000000.0, got %f", avg)
		}
	})
}

func TestAverageInt64(t *testing.T) {
	t.Run("average int64 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{1, 2, 3, 4, 5})

		avg, ok := enumerator.AverageInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 3.0 {
			t.Errorf("Expected average 3.0, got %f", avg)
		}
	})

	t.Run("average int64 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{-1, -2, -3, -4, -5})

		avg, ok := enumerator.AverageInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != -3.0 {
			t.Errorf("Expected average -3.0, got %f", avg)
		}
	})

	t.Run("average int64 single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int64{9223372036854775807}) // max int64

		avg, ok := enumerator.AverageInt64(selector.Int64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 9223372036854775807.0 {
			t.Errorf("Expected average 9223372036854775807.0, got %f", avg)
		}
	})

	t.Run("average int64 with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			ID    int
			Value int64
		}

		data := []Data{
			{ID: 1, Value: 1000000000},
			{ID: 2, Value: 2000000000},
			{ID: 3, Value: 3000000000},
		}

		enumerator := FromSlice(data)
		avg, ok := enumerator.AverageInt64(func(d Data) int64 { return d.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 2000000000.0 {
			t.Errorf("Expected average 2000000000.0, got %f", avg)
		}
	})
}

func TestAverageFloat(t *testing.T) {
	t.Run("average float32 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.5, 2.5, 3.5, 4.5, 5.5})

		avg, ok := enumerator.AverageFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 3.5 {
			t.Errorf("Expected average 3.5, got %f", avg)
		}
	})

	t.Run("average float32 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{-1.5, -2.5, -3.5, -4.5, -5.5})

		avg, ok := enumerator.AverageFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != -3.5 {
			t.Errorf("Expected average -3.5, got %f", avg)
		}
	})

	t.Run("average float32 with decimal precision", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.1, 2.2, 3.3, 4.4, 5.5})

		avg, ok := enumerator.AverageFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 3.3
		if avg < expected-0.001 || avg > expected+0.001 {
			t.Errorf("Expected average ~3.3, got %f", avg)
		}
	})

	t.Run("average float32 with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Measurement struct {
			Name  string
			Value float32
		}

		measurements := []Measurement{
			{Name: "Temp1", Value: 25.5},
			{Name: "Temp2", Value: 22.3},
			{Name: "Temp3", Value: 28.7},
		}

		enumerator := FromSlice(measurements)
		avg, ok := enumerator.AverageFloat(func(m Measurement) float32 { return m.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 25.5
		if avg < expected-0.001 || avg > expected+0.001 {
			t.Errorf("Expected average ~25.5, got %f", avg)
		}
	})
}

func TestAverageFloat64(t *testing.T) {
	t.Run("average float64 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{1.5, 2.5, 3.5, 4.5, 5.5})

		avg, ok := enumerator.AverageFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 3.5 {
			t.Errorf("Expected average 3.5, got %f", avg)
		}
	})

	t.Run("average float64 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{-1.5, -2.5, -3.5, -4.5, -5.5})

		avg, ok := enumerator.AverageFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != -3.5 {
			t.Errorf("Expected average -3.5, got %f", avg)
		}
	})

	t.Run("average float64 with high precision", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{1.111111111, 2.222222222, 3.333333333})

		avg, ok := enumerator.AverageFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 2.222222222
		if avg < expected-0.000000001 || avg > expected+0.000000001 {
			t.Errorf("Expected average ~2.222222222, got %f", avg)
		}
	})

	t.Run("average float64 with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Price struct {
			Name  string
			Value float64
		}

		prices := []Price{
			{Name: "Item1", Value: 100.99},
			{Name: "Item2", Value: 200.50},
			{Name: "Item3", Value: 150.75},
		}

		enumerator := FromSlice(prices)
		avg, ok := enumerator.AverageFloat64(func(p Price) float64 { return p.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 150.74666666666667
		if avg < expected-0.000000001 || avg > expected+0.000000001 {
			t.Errorf("Expected average ~150.74666666666667, got %f", avg)
		}
	})
}

func TestAverageStruct(t *testing.T) {
	t.Run("average int from struct field", func(t *testing.T) {
		t.Parallel()
		type Student struct {
			Name  string
			Score int
		}

		students := []Student{
			{Name: "Alice", Score: 85},
			{Name: "Bob", Score: 90},
			{Name: "Charlie", Score: 75},
		}

		enumerator := FromSlice(students)
		avg, ok := enumerator.AverageInt(func(s Student) int { return s.Score })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 83.33333333333333 {
			t.Errorf("Expected average ~83.33, got %f", avg)
		}
	})

	t.Run("average int64 from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type DataPoint struct {
			Tags   []string
			Amount int64
		}

		dataPoints := []DataPoint{
			{Tags: []string{"sales"}, Amount: 1000},
			{Tags: []string{"marketing"}, Amount: 2000},
			{Tags: []string{"development"}, Amount: 3000},
		}

		enumerator := FromSliceAny(dataPoints)
		avg, ok := enumerator.AverageInt64(func(d DataPoint) int64 { return d.Amount })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 2000.0 {
			t.Errorf("Expected average 2000.0, got %f", avg)
		}
	})
}

func TestAverageEdgeCases(t *testing.T) {
	t.Run("average int with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{5, 5, 5, 5})

		avg, ok := enumerator.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 5.0 {
			t.Errorf("Expected average 5.0, got %f", avg)
		}
	})

	t.Run("average float64 with very small numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{1e-10, 2e-10, 3e-10})

		avg, ok := enumerator.AverageFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 2e-10
		if avg < expected*0.999 || avg > expected*1.001 {
			t.Errorf("Expected average ~2e-10, got %e", avg)
		}
	})

	t.Run("average float64 with very large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{1e10, 2e10, 3e10})

		avg, ok := enumerator.AverageFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 2e10
		if avg < expected*0.999 || avg > expected*1.001 {
			t.Errorf("Expected average ~2e10, got %e", avg)
		}
	})
}

func TestAverageWithOperations(t *testing.T) {
	t.Run("average int after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6})
		filtered := enumerator.Where(func(n int) bool { return n%2 == 0 })

		avg, ok := filtered.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 4.0 {
			t.Errorf("Expected average 4.0, got %f", avg)
		}
	})

	t.Run("average int after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		taken := enumerator.Take(4)

		avg, ok := taken.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 2.5 {
			t.Errorf("Expected average 2.5, got %f", avg)
		}
	})

	t.Run("average int after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		skipped := enumerator.Skip(4)

		avg, ok := skipped.AverageInt(selector.Int)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 6.5 {
			t.Errorf("Expected average 6.5, got %f", avg)
		}
	})
}

func TestAverageNonComparable(t *testing.T) {
	t.Run("average int from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Name  []string
			Score int
		}

		records := []Record{
			{Name: []string{"Record1"}, Score: 80},
			{Name: []string{"Record2"}, Score: 90},
			{Name: []string{"Record3"}, Score: 70},
		}

		enumerator := FromSliceAny(records)
		avg, ok := enumerator.AverageInt(func(r Record) int { return r.Score })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if avg != 80.0 {
			t.Errorf("Expected average 80.0, got %f", avg)
		}
	})

	t.Run("average float64 with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Rate     float64
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Rate: 0.8},
			{Settings: map[string]interface{}{"b": 2}, Rate: 0.9},
			{Settings: map[string]interface{}{"c": 3}, Rate: 0.7},
		}

		enumerator := FromSliceAny(configs)
		avg, ok := enumerator.AverageFloat64(func(c Config) float64 { return c.Rate })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := 0.8
		if avg < expected-0.001 || avg > expected+0.001 {
			t.Errorf("Expected average ~0.8, got %f", avg)
		}
	})
}

func BenchmarkAverageInt(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []int{1, 2, 3, 4, 5}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.AverageInt(selector.Int)
			if !ok || result != 3.0 {
				b.Fatalf("Expected 3.0, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.AverageInt(selector.Int)
			if !ok || result != 499.5 {
				b.Fatalf("Expected 499.5, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = i
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.AverageInt(selector.Int)
			if !ok || result != 4999.5 {
				b.Fatalf("Expected 4999.5, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]int{})
			result, ok := enumerator.AverageInt(selector.Int)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %f, ok: %v", result, ok)
			}
		}
	})
}

func BenchmarkAverageFloat64(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []float64{1.5, 2.5, 3.5, 4.5, 5.5}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.AverageFloat64(selector.Float64)
			if !ok || result != 3.5 {
				b.Fatalf("Expected 3.5, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration", func(b *testing.B) {
		items := make([]float64, 1000)
		for i := 0; i < 1000; i++ {
			items[i] = float64(i) + 0.5
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.AverageFloat64(selector.Float64)
			if !ok || result != 500.0 {
				b.Fatalf("Expected 500.0, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration", func(b *testing.B) {
		items := make([]float64, 10000)
		for i := 0; i < 10000; i++ {
			items[i] = float64(i) + 0.5
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.AverageFloat64(selector.Float64)
			if !ok || result != 5000.0 {
				b.Fatalf("Expected 5000.0, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]float64{})
			result, ok := enumerator.AverageFloat64(selector.Float64)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %f, ok: %v", result, ok)
			}
		}
	})
}
