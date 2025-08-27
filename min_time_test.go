package enumerable

import (
	"testing"
	"time"

	"github.com/ahatornn/enumerable/selector"
)

func TestMinTime(t *testing.T) {
	t.Run("min time from mixed timestamps", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			time.Date(2023, 3, 8, 8, 20, 15, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time with zero time", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			time.Time{}, // zero time
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.IsZero() {
			t.Errorf("Expected min zero time, got %v", min)
		}
	})

	t.Run("min time single element", func(t *testing.T) {
		t.Parallel()
		singleTime := time.Date(2023, 7, 4, 12, 0, 0, 0, time.UTC)
		enumerator := FromSlice([]time.Time{singleTime})

		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.Equal(singleTime) {
			t.Errorf("Expected min %v, got %v", singleTime, min)
		}
	})

	t.Run("min time with equal timestamps", func(t *testing.T) {
		t.Parallel()
		sameTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
		times := []time.Time{
			sameTime,
			sameTime,
			sameTime,
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.Equal(sameTime) {
			t.Errorf("Expected min %v, got %v", sameTime, min)
		}
	})

	t.Run("min time empty slice", func(t *testing.T) {
		t.Parallel()
		enumerator := FromSlice([]time.Time{})

		min, ok := enumerator.MinTime(selector.Time)

		if ok {
			t.Error("Expected ok to be false for empty slice")
		}
		if !min.IsZero() {
			t.Errorf("Expected min zero time for empty slice, got %v", min)
		}
	})

	t.Run("min time nil enumerator", func(t *testing.T) {
		t.Parallel()
		var enumerator Enumerator[time.Time] = nil

		min, ok := enumerator.MinTime(selector.Time)

		if ok {
			t.Error("Expected ok to be false for nil enumerator")
		}
		if !min.IsZero() {
			t.Errorf("Expected min zero time for nil enumerator, got %v", min)
		}
	})

	t.Run("min time with nil keySelector", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(nil)

		if ok {
			t.Error("Expected ok to be false for nil keySelector")
		}
		if !min.IsZero() {
			t.Errorf("Expected min zero time for nil keySelector, got %v", min)
		}
	})

	t.Run("min time with custom key selector", func(t *testing.T) {
		t.Parallel()
		type Event struct {
			Name      string
			Timestamp time.Time
		}

		events := []Event{
			{Name: "Event1", Timestamp: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Name: "Event2", Timestamp: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Name: "Event3", Timestamp: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSlice(events)
		min, ok := enumerator.MinTime(func(e Event) time.Time { return e.Timestamp })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time with duplicate minimum values", func(t *testing.T) {
		t.Parallel()
		minTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			minTime,
			time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			minTime,
			time.Date(2023, 3, 8, 8, 20, 15, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.Equal(minTime) {
			t.Errorf("Expected min %v, got %v", minTime, min)
		}
	})
}

func TestMinTimeStruct(t *testing.T) {
	t.Run("min time from struct field", func(t *testing.T) {
		t.Parallel()
		type LogEntry struct {
			Message   string
			Timestamp time.Time
		}

		logs := []LogEntry{
			{Message: "First", Timestamp: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Message: "Second", Timestamp: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Message: "Third", Timestamp: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSlice(logs)
		min, ok := enumerator.MinTime(func(l LogEntry) time.Time { return l.Timestamp })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time from struct with slice field", func(t *testing.T) {
		t.Parallel()
		type DataPoint struct {
			Tags      []string
			CreatedAt time.Time
		}

		dataPoints := []DataPoint{
			{Tags: []string{"sensor1"}, CreatedAt: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Tags: []string{"sensor2"}, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Tags: []string{"sensor3"}, CreatedAt: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSliceAny(dataPoints)
		min, ok := enumerator.MinTime(func(d DataPoint) time.Time { return d.CreatedAt })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})
}

func TestMinTimeEdgeCases(t *testing.T) {
	t.Run("min time with all same values", func(t *testing.T) {
		t.Parallel()
		sameTime := time.Date(2023, 7, 4, 12, 0, 0, 0, time.UTC)
		times := []time.Time{
			sameTime,
			sameTime,
			sameTime,
			sameTime,
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.Equal(sameTime) {
			t.Errorf("Expected min %v, got %v", sameTime, min)
		}
	})

	t.Run("min time early termination optimization", func(t *testing.T) {
		t.Parallel()
		times := make([]time.Time, 1000)
		times[0] = time.Time{}
		for i := 1; i < 1000; i++ {
			times[i] = time.Date(2023, 1, i%31+1, i%24, i%60, i%60, 0, time.UTC)
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.IsZero() {
			t.Errorf("Expected min zero time, got %v", min)
		}
	})

	t.Run("min time without early termination", func(t *testing.T) {
		t.Parallel()
		times := make([]time.Time, 1000)
		for i := 0; i < 999; i++ {
			times[i] = time.Date(2023, 1, i%31+1, i%24, i%60, i%60, 0, time.UTC)
		}
		times[999] = time.Time{}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		if !min.IsZero() {
			t.Errorf("Expected min zero time, got %v", min)
		}
	})

	t.Run("min time with nanosecond precision", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 1, 1, 0, 0, 0, 100, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 50, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 200, time.UTC),
		}

		enumerator := FromSlice(times)
		min, ok := enumerator.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 50, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})
}

func TestMinTimeWithOperations(t *testing.T) {
	t.Run("min time after filter", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			time.Date(2022, 3, 8, 8, 20, 15, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		filtered := enumerator.Where(func(t time.Time) bool { return t.Year() == 2023 }) // только 2023 год

		min, ok := filtered.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v (from 2023 year), got %v", expected, min)
		}
	})

	t.Run("min time after take", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			time.Date(2023, 3, 8, 8, 20, 15, 0, time.UTC),
			time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		taken := enumerator.Take(3)

		min, ok := taken.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v (from first 3 elements), got %v", expected, min)
		}
	})

	t.Run("min time after skip", func(t *testing.T) {
		t.Parallel()
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			time.Date(2023, 3, 8, 8, 20, 15, 0, time.UTC),
			time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC),
		}

		enumerator := FromSlice(times)
		skipped := enumerator.Skip(2)

		min, ok := skipped.MinTime(selector.Time)

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 3, 8, 8, 20, 15, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v (after skipping 2 elements), got %v", expected, min)
		}
	})
}

func TestMinTimeCustomKeySelector(t *testing.T) {
	t.Run("min time by struct field with time calculation", func(t *testing.T) {
		t.Parallel()
		type Event struct {
			Name string
			Date time.Time
		}

		events := []Event{
			{Name: "Event1", Date: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)},
			{Name: "Event2", Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Name: "Event3", Date: time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)},
		}

		enumerator := FromSlice(events)
		min, ok := enumerator.MinTime(func(e Event) time.Time { return e.Date })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time by array element", func(t *testing.T) {
		t.Parallel()
		type Data struct {
			Timestamps [3]time.Time
		}

		data := []Data{
			{
				Timestamps: [3]time.Time{
					time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			{
				Timestamps: [3]time.Time{
					time.Date(2023, 3, 8, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2023, 11, 11, 0, 0, 0, 0, time.UTC),
				},
			},
		}

		enumerator := FromSlice(data)
		min, ok := enumerator.MinTime(func(d Data) time.Time { return d.Timestamps[0] })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 3, 8, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})
}

func TestMinTimeNonComparable(t *testing.T) {
	t.Run("min time from non-comparable struct slice", func(t *testing.T) {
		t.Parallel()
		type Record struct {
			Name    []string
			Created time.Time
		}

		records := []Record{
			{Name: []string{"Record1"}, Created: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Name: []string{"Record2"}, Created: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Name: []string{"Record3"}, Created: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSliceAny(records)
		min, ok := enumerator.MinTime(func(r Record) time.Time { return r.Created })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time with map field struct", func(t *testing.T) {
		t.Parallel()
		type Config struct {
			Settings map[string]interface{}
			Updated  time.Time
		}

		configs := []Config{
			{Settings: map[string]interface{}{"a": 1}, Updated: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Settings: map[string]interface{}{"b": 2}, Updated: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Settings: map[string]interface{}{"c": 3}, Updated: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSliceAny(configs)
		min, ok := enumerator.MinTime(func(c Config) time.Time { return c.Updated })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time with function field struct", func(t *testing.T) {
		t.Parallel()
		type Handler struct {
			Callback func()
			LastRun  time.Time
		}

		handlers := []Handler{
			{Callback: func() {}, LastRun: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Callback: func() {}, LastRun: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Callback: func() {}, LastRun: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSliceAny(handlers)
		min, ok := enumerator.MinTime(func(h Handler) time.Time { return h.LastRun })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time with complex non-comparable struct", func(t *testing.T) {
		t.Parallel()
		type ComplexStruct struct {
			Data      map[string][]time.Time
			Callback  func(int) bool
			Channel   chan string
			Timestamp time.Time
		}

		complexData := []ComplexStruct{
			{
				Data:      map[string][]time.Time{"a": {time.Now()}},
				Callback:  func(i int) bool { return i > 0 },
				Channel:   make(chan string, 1),
				Timestamp: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			},
			{
				Data:      map[string][]time.Time{"b": {time.Now()}},
				Callback:  func(i int) bool { return i < 0 },
				Channel:   make(chan string, 1),
				Timestamp: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				Data:      map[string][]time.Time{"c": {time.Now()}},
				Callback:  func(i int) bool { return i == 0 },
				Channel:   make(chan string, 1),
				Timestamp: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			},
		}

		enumerator := FromSliceAny(complexData)
		min, ok := enumerator.MinTime(func(c ComplexStruct) time.Time { return c.Timestamp })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})

	t.Run("min time with interface field struct", func(t *testing.T) {
		t.Parallel()
		type Service struct {
			Config   interface{}
			Metadata []interface{}
			LastPing time.Time
		}

		services := []Service{
			{Config: "config1", Metadata: []interface{}{"a", 1}, LastPing: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)},
			{Config: 42, Metadata: []interface{}{"b", 2}, LastPing: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
			{Config: map[string]string{"key": "value"}, Metadata: []interface{}{"c", 3}, LastPing: time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC)},
		}

		enumerator := FromSliceAny(services)
		min, ok := enumerator.MinTime(func(s Service) time.Time { return s.LastPing })

		if !ok {
			t.Error("Expected ok to be true")
		}
		expected := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		if !min.Equal(expected) {
			t.Errorf("Expected min %v, got %v", expected, min)
		}
	})
}

func BenchmarkMinTime(b *testing.B) {
	b.Run("small enumeration", func(b *testing.B) {
		times := []time.Time{
			time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 15, 15, 45, 30, 0, time.UTC),
			time.Date(2023, 3, 8, 8, 20, 15, 0, time.UTC),
			time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC),
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(times)
			result, ok := enumerator.MinTime(selector.Time)
			if !ok || !result.Equal(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)) {
				b.Fatalf("Expected 2023-01-01, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("medium enumeration early termination", func(b *testing.B) {
		times := make([]time.Time, 1000)
		times[0] = time.Time{}
		for i := 1; i < 1000; i++ {
			times[i] = time.Date(2023, 1, i%31+1, i%24, i%60, i%60, 0, time.UTC)
		}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(times)
			result, ok := enumerator.MinTime(selector.Time)
			if !ok || !result.IsZero() {
				b.Fatalf("Expected zero time, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("large enumeration no early termination", func(b *testing.B) {
		times := make([]time.Time, 10000)
		for i := 0; i < 9999; i++ {
			times[i] = time.Date(2023, 1, i%31+1, i%24, i%60, i%60, 0, time.UTC)
		}
		times[9999] = time.Time{}

		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice(times)
			result, ok := enumerator.MinTime(selector.Time)
			if !ok || !result.IsZero() {
				b.Fatalf("Expected zero time, got %v, ok: %v", result, ok)
			}
		}
	})

	b.Run("empty enumeration", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			enumerator := FromSlice([]time.Time{})
			result, ok := enumerator.MinTime(selector.Time)
			if ok || !result.IsZero() {
				b.Fatalf("Expected zero time and false, got %v, ok: %v", result, ok)
			}
		}
	})
}
