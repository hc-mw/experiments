package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
