package enumerable

import (
	"math"
	"testing"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinFloat32(t *testing.T) {
	t.Run("min float32 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.5, 2.2, 8.8, 1.1, 9.9})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min float32 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{-5.5, -2.2, -8.8, -1.1, -9.9})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9.9 {
			t.Errorf("Expected min -9.9, got %f", min)
		}
	})

	t.Run("min float32 from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{-5.5, 2.2, -8.8, 1.1, 0.0})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -8.8 {
			t.Errorf("Expected min -8.8, got %f", min)
		}
	})

	t.Run("min float32 single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{42.5})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 42.5 {
			t.Errorf("Expected min 42.5, got %f", min)
		}
	})

	t.Run("min float32 with infinity", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{100.0, float32(math.Inf(-1)), 50.0})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !math.IsInf(float64(min), -1) {
			t.Errorf("Expected min -Inf, got %f", min)
		}
	})

	t.Run("min float32 with NaN values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.0, float32(math.NaN()), 3.0, 1.0})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.0 {
			t.Errorf("Expected valid minimum, got %f", min)
		}
	})

	t.Run("min float32 empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{})

		min, ok := enumerator.MinFloat(selector.Float32)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for empty slice, got %f", min)
		}
	})

	t.Run("min float32 nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[float32] = nil

		min, ok := enumerator.MinFloat(selector.Float32)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil enumerator, got %f", min)
		}
	})

	t.Run("min float32 with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.0, 2.0, 3.0})

		min, ok := enumerator.MinFloat(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil keySelector, got %f", min)
		}
	})

	t.Run("min float32 with custom key selector", func(t *testing.T) {
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
		min, ok := enumerator.MinFloat(func(m Measurement) float32 { return m.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 22.3 {
			t.Errorf("Expected min 22.3, got %f", min)
		}
	})

	t.Run("min float32 with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.5, 1.1, 8.8, 1.1, 9.9})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min float32 with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.5, 0.0, 8.8, 0.0, 9.9})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0.0 {
			t.Errorf("Expected min 0.0, got %f", min)
		}
	})
}

func TestMinFloat64(t *testing.T) {
	t.Run("min float64 from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{5.5, 2.2, 8.8, 1.1, 9.9})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min float64 for any from positive numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]float64{5.5, 2.2, 8.8, 1.1, 9.9})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min float64 from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{-5.5, -2.2, -8.8, -1.1, -9.9})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9.9 {
			t.Errorf("Expected min -9.9, got %f", min)
		}
	})

	t.Run("min float64 for any from negative numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSliceAny([]float64{-5.5, -2.2, -8.8, -1.1, -9.9})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -9.9 {
			t.Errorf("Expected min -9.9, got %f", min)
		}
	})

	t.Run("min float64 from mixed numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{-5.5, 2.2, -8.8, 1.1, 0.0})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != -8.8 {
			t.Errorf("Expected min -8.8, got %f", min)
		}
	})

	t.Run("min float64 single element", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{42.5})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 42.5 {
			t.Errorf("Expected min 42.5, got %f", min)
		}
	})

	t.Run("min float64 with infinity", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{100.0, math.Inf(-1), 50.0})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !math.IsInf(min, -1) {
			t.Errorf("Expected min -Inf, got %f", min)
		}
	})

	t.Run("min float64 with NaN values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{5.0, math.NaN(), 3.0, 1.0})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.0 {
			t.Errorf("Expected valid minimum, got %f", min)
		}
	})

	t.Run("min float64 empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for empty slice, got %f", min)
		}
	})

	t.Run("min float64 nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[float64] = nil

		min, ok := enumerator.MinFloat64(selector.Float64)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil enumerator, got %f", min)
		}
	})

	t.Run("min float64 with nil keySelector", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{1.0, 2.0, 3.0})

		min, ok := enumerator.MinFloat64(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if min != 0 {
			t.Errorf("Expected min 0 for nil keySelector, got %f", min)
		}
	})

	t.Run("min float64 with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Measurement struct {
			Name  string
			Value float64
		}

		measurements := []Measurement{
			{Name: "Temp1", Value: 25.5},
			{Name: "Temp2", Value: 22.3},
			{Name: "Temp3", Value: 28.7},
		}

		enumerator := FromSlice(measurements)
		min, ok := enumerator.MinFloat64(func(m Measurement) float64 { return m.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 22.3 {
			t.Errorf("Expected min 22.3, got %f", min)
		}
	})

	t.Run("min float64 with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{5.5, 1.1, 8.8, 1.1, 9.9})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min float64 with zero values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float64{5.5, 0.0, 8.8, 0.0, 9.9})

		min, ok := enumerator.MinFloat64(selector.Float64)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0.0 {
			t.Errorf("Expected min 0.0, got %f", min)
		}
	})
}

func TestMinFloat32Struct(t *testing.T) {
	t.Run("min float32 from struct field", func(t *testing.T) {
		t.Parallel()
		type Product struct {
			Name  string
			Price float32
		}

		products := []Product{
			{Name: "Laptop", Price: 999.99},
			{Name: "Mouse", Price: 25.50},
			{Name: "Keyboard", Price: 75.00},
		}

		enumerator := FromSlice(products)
		min, ok := enumerator.MinFloat(func(p Product) float32 { return p.Price })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25.50 {
			t.Errorf("Expected min 25.50, got %f", min)
		}
	})

	t.Run("min float32 from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type DataPoint struct {
			Tags  []string
			Value float32
		}

		dataPoints := []DataPoint{
			{Tags: []string{"sensor1"}, Value: 25.5},
			{Tags: []string{"sensor2"}, Value: 22.3},
			{Tags: []string{"sensor3"}, Value: 28.7},
		}

		enumerator := FromSliceAny(dataPoints)
		min, ok := enumerator.MinFloat(func(d DataPoint) float32 { return d.Value })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 22.3 {
			t.Errorf("Expected min 22.3, got %f", min)
		}
	})
}

func TestMinFloat32EdgeCases(t *testing.T) {
	t.Run("min float32 with all same values", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.5, 5.5, 5.5, 5.5})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 5.5 {
			t.Errorf("Expected min 5.5, got %f", min)
		}
	})

	t.Run("min float32 with very large numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1e30, 1e31, 1e29})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1e29 {
			t.Errorf("Expected min 1e29, got %e", min)
		}
	})

	t.Run("min float32 with very small numbers", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1e-30, 1e-31, 1e-29})

		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1e-31 {
			t.Errorf("Expected min 1e-31, got %e", min)
		}
	})

	t.Run("min float32 early termination optimization", func(t *testing.T) {
		t.Parallel()
		floats := make([]float32, 1000)
		floats[0] = float32(math.Inf(-1))
		for i := 1; i < 1000; i++ {
			floats[i] = float32(i)
		}

		enumerator := FromSlice(floats)
		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !math.IsInf(float64(min), -1) {
			t.Errorf("Expected min -Inf, got %f", min)
		}
	})

	t.Run("min float32 without early termination", func(t *testing.T) {
		t.Parallel()
		floats := make([]float32, 1000)
		for i := 0; i < 999; i++ {
			floats[i] = float32(i + 1)
		}
		floats[999] = float32(math.Inf(-1))

		enumerator := FromSlice(floats)
		min, ok := enumerator.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !math.IsInf(float64(min), -1) {
			t.Errorf("Expected min -Inf, got %f", min)
		}
	})
}

func TestMinFloat32WithOperations(t *testing.T) {
	t.Run("min float32 after filter", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{1.5, 2.7, 3.9, 4.1, 5.3, 6.5})
		filtered := enumerator.Where(func(f float32) bool { return f > 3.0 })

		min, ok := filtered.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 3.9 {
			t.Errorf("Expected min 3.9 (from filtered values > 3.0), got %f", min)
		}
	})

	t.Run("min float32 after take", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 0.0})
		taken := enumerator.Take(4)

		min, ok := taken.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1 (from first 4 elements), got %f", min)
		}
	})

	t.Run("min float32 after skip", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]float32{5.5, 0.2, 8.8, 1.1, 9.9, 3.3})
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinFloat(selector.Float32)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1 (after skipping 2 elements), got %f", min)
		}
	})
}

func TestMinFloat32CustomKeySelector(t *testing.T) {
	t.Run("min float32 by string length converted to float", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]string{"hello", "hi", "world", "golang"})

		min, ok := enumerator.MinFloat(func(s string) float32 { return float32(len(s)) })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 2.0 {
			t.Errorf("Expected min 2.0, got %f", min)
		}
	})

	t.Run("min float32 by array element", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Values [5]float32
		}

		data := []Data{
			{Values: [5]float32{100.5, 50.2, 200.8, 25.1, 150.3}},
			{Values: [5]float32{200.1, 100.9, 50.7, 75.4, 125.6}},
			{Values: [5]float32{25.3, 75.8, 125.2, 175.5, 225.9}},
		}
		enumerator := FromSlice(data)
		min, ok := enumerator.MinFloat(func(d Data) float32 { return d.Values[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 25.3 {
			t.Errorf("Expected min 25.3, got %f", min)
		}
	})
}

func TestMinFloat32NonComparable(t *testing.T) {
	t.Run("min float32 from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name   []string
			Salary float32
		}

		people := []Person{
			{Name: []string{"Alice"}, Salary: 75000.50},
			{Name: []string{"Bob"}, Salary: 65000.75},
			{Name: []string{"Charlie"}, Salary: 85000.25},
		}

		enumerator := FromSliceAny(people)
		min, ok := enumerator.MinFloat(func(p Person) float32 { return p.Salary })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 65000.75 {
			t.Errorf("Expected min 65000.75, got %f", min)
		}
	})

	t.Run("min float32 with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Timeout  float32
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Timeout: 30.5},
			{Settings: map[string]interface{}{"b": 2}, Timeout: 10.2},
			{Settings: map[string]interface{}{"c": 3}, Timeout: 60.8},
		}

		enumerator := FromSliceAny(configs)
		min, ok := enumerator.MinFloat(func(c Config) float32 { return c.Timeout })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 10.2 {
			t.Errorf("Expected min 10.2, got %f", min)
		}
	})

	t.Run("min float32 with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			Priority float32
		}

		handlers := []Handler{
			{Callback: func() {}, Priority: 5.5},
			{Callback: func() {}, Priority: 1.1},
			{Callback: func() {}, Priority: 3.3},
		}

		enumerator := FromSliceAny(handlers)
		min, ok := enumerator.MinFloat(func(h Handler) float32 { return h.Priority })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 1.1 {
			t.Errorf("Expected min 1.1, got %f", min)
		}
	})

	t.Run("min float32 with complex non-comparable struct", func(t *testing.T) {
		t.Parallel()
		type ComplexStruct struct {
			Data     map[string][]float32
			Callback func(int) bool
			Channel  chan string
			Score    float32
		}

		complexData := []ComplexStruct{
			{
				Data:     map[string][]float32{"a": {1.1, 2.2}},
				Callback: func(i int) bool { return i > 0 },
				Channel:  make(chan string, 1),
				Score:    80.5,
			},
			{
				Data:     map[string][]float32{"b": {3.3, 4.4}},
				Callback: func(i int) bool { return i < 0 },
				Channel:  make(chan string, 1),
				Score:    95.7,
			},
			{
				Data:     map[string][]float32{"c": {5.5, 6.6}},
				Callback: func(i int) bool { return i == 0 },
				Channel:  make(chan string, 1),
				Score:    70.3,
			},
		}

		enumerator := FromSliceAny(complexData)
		min, ok := enumerator.MinFloat(func(c ComplexStruct) float32 { return c.Score })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 70.3 {
			t.Errorf("Expected min 70.3, got %f", min)
		}
	})

	t.Run("min float32 with interface field struct", func(t *testing.T) {
		t.Parallel()
		type Service struct {
			Config   interface{}
			Metadata []interface{}
			Rate     float32
		}

		services := []Service{
			{Config: "config1", Metadata: []interface{}{"a", 1}, Rate: 0.8},
			{Config: 42, Metadata: []interface{}{"b", 2}, Rate: 0.3},
			{Config: map[string]string{"key": "value"}, Metadata: []interface{}{"c", 3}, Rate: 0.9},
		}

		enumerator := FromSliceAny(services)
		min, ok := enumerator.MinFloat(func(s Service) float32 { return s.Rate })

		if !ok {
			t.Error("Expected ok to be true")
		}
		if min != 0.3 {
			t.Errorf("Expected min 0.3, got %f", min)
		}
	})
}

func BenchmarkMinFloat32(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		items := []float32{5.5, 2.2, 8.8, 1.1, 9.9}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinFloat(selector.Float32)
			if !ok || result != 1.1 {
				b.Fatalf("Expected 1.1, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		// -Inf в начале - должно быть быстро
		items := make([]float32, 1000)
		items[0] = float32(math.Inf(-1))
		for i := 1; i < 1000; i++ {
			items[i] = float32(i)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinFloat(selector.Float32)
			if !ok || !math.IsInf(float64(result), -1) {
				b.Fatalf("Expected -Inf, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		// -Inf в конце - должно пройти все элементы
		items := make([]float32, 10000)
		for i := 0; i < 9999; i++ {
			items[i] = float32(i + 1)
		}
		items[9999] = float32(math.Inf(-1))

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(items)
			result, ok := enumerator.MinFloat(selector.Float32)
			if !ok || !math.IsInf(float64(result), -1) {
				b.Fatalf("Expected -Inf, got %f, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]float32{})
			result, ok := enumerator.MinFloat(selector.Float32)
			if ok || result != 0 {
				b.Fatalf("Expected 0 and false, got %f, ok: %v", result, ok)
			}
		}
	})
}
