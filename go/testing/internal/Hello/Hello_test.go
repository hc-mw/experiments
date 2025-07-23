package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		got := Hello("world")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Empty name", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
