package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Friend")

	got := buf.String()
	want := "Hello, Friend"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
