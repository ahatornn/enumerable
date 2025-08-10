package enumerable

import (
	"testing"
	"time"
)

func TestFromChannel(t *testing.T) {
	t.Run("basic channel enumeration", func(t *testing.T) {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		enumerator := FromChannel(ch)

		expected := []int{1, 2, 3}
		actual := []int{}

		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("empty channel", func(t *testing.T) {
		ch := make(chan string)
		close(ch)

		enumerator := FromChannel(ch)

		count := 0
		enumerator(func(item string) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from empty channel, got %d", count)
		}
	})

	t.Run("nil channel", func(t *testing.T) {
		var ch chan int // nil channel

		enumerator := FromChannel(ch)

		count := 0
		enumerator(func(item int) bool {
			count++
			return true
		})

		if count != 0 {
			t.Errorf("Expected 0 items from nil channel, got %d", count)
		}
	})

	t.Run("early termination", func(t *testing.T) {
		ch := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)

		enumerator := FromChannel(ch)

		actual := []int{}
		enumerator(func(item int) bool {
			actual = append(actual, item)
			return len(actual) < 3 // Останавливаемся после 3 элементов
		})

		expected := []int{1, 2, 3}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("string channel", func(t *testing.T) {
		ch := make(chan string, 2)
		ch <- "hello"
		ch <- "world"
		close(ch)

		enumerator := FromChannel(ch)

		actual := []string{}
		enumerator(func(item string) bool {
			actual = append(actual, item)
			return true
		})

		expected := []string{"hello", "world"}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %s at index %d, got %s", v, i, actual[i])
			}
		}
	})
}

func TestFromChannelConcurrent(t *testing.T) {
	t.Run("concurrent channel send", func(t *testing.T) {
		ch := make(chan int)

		// Горутина для отправки данных
		go func() {
			defer close(ch)
			for i := 1; i <= 3; i++ {
				ch <- i
				time.Sleep(10 * time.Millisecond) // Небольшая задержка
			}
		}()

		enumerator := FromChannel(ch)

		actual := []int{}
		enumerator(func(item int) bool {
			actual = append(actual, item)
			return true
		})

		expected := []int{1, 2, 3}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %d at index %d, got %d", v, i, actual[i])
			}
		}
	})

	t.Run("channel closed during iteration", func(t *testing.T) {
		ch := make(chan int, 10)
		ch <- 1
		ch <- 2

		enumerator := FromChannel(ch)

		actual := []int{}
		go func() {
			time.Sleep(50 * time.Millisecond)
			close(ch) // Закрываем канал во время итерации
		}()

		enumerator(func(item int) bool {
			actual = append(actual, item)
			time.Sleep(20 * time.Millisecond) // Задержка между итерациями
			return true
		})

		if len(actual) != 2 {
			t.Errorf("Expected 2 items, got %d", len(actual))
		}
	})
}

func TestFromChannelStruct(t *testing.T) {
	t.Run("struct channel", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		ch := make(chan Person, 2)
		ch <- Person{Name: "Alice", Age: 30}
		ch <- Person{Name: "Bob", Age: 25}
		close(ch)

		enumerator := FromChannel(ch)

		actual := []Person{}
		enumerator(func(item Person) bool {
			actual = append(actual, item)
			return true
		})

		expected := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}

		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %+v at index %d, got %+v", v, i, actual[i])
			}
		}
	})
}

func TestFromChannelEdgeCases(t *testing.T) {
	t.Run("single item channel", func(t *testing.T) {
		ch := make(chan int, 1)
		ch <- 42
		close(ch)

		enumerator := FromChannel(ch)

		var result int
		found := false

		enumerator(func(item int) bool {
			result = item
			found = true
			return true
		})

		if !found {
			t.Error("Expected to find one element, but found none")
		}

		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
	})

	t.Run("boolean channel", func(t *testing.T) {
		ch := make(chan bool, 3)
		ch <- true
		ch <- false
		ch <- true
		close(ch)

		enumerator := FromChannel(ch)

		actual := []bool{}
		enumerator(func(item bool) bool {
			actual = append(actual, item)
			return true
		})

		expected := []bool{true, false, true}
		if len(actual) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(actual))
		}

		for i, v := range expected {
			if actual[i] != v {
				t.Errorf("Expected %t at index %d, got %t", v, i, actual[i])
			}
		}
	})
}

// Benchmark для проверки производительности
func BenchmarkFromChannel(b *testing.B) {
	b.Run("small channel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch := make(chan int, 10)
			for j := 0; j < 10; j++ {
				ch <- j
			}
			close(ch)

			enumerator := FromChannel(ch)
			enumerator(func(item int) bool {
				_ = item
				return true
			})
		}
	})

	b.Run("large channel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ch := make(chan int, 1000)
			for j := 0; j < 1000; j++ {
				ch <- j
			}
			close(ch)

			enumerator := FromChannel(ch)
			enumerator(func(item int) bool {
				_ = item
				return true
			})
		}
	})
}
