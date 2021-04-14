package iteration

import (
    "testing"
    "fmt"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeating 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeating 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeating -3 times", func(t *testing.T) {
		repeated := Repeat("a", -3)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
    repeat := Repeat("a", 6)
    fmt.Println(repeat)
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 0)
    }
}
