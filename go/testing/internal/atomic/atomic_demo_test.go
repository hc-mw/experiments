package atomic

import (
	"testing"
)

func TestAtomicIncrease(t *testing.T) {
	counter := &Counter{}

	// Increase the counter atomically
	counter.AtomicIncrease()

	// Check if the value is 1
	if got := counter.Value(); got != 1 {
		t.Errorf("got %d, want %d", got, 1)
	}

	// Increase the counter again atomically
	counter.AtomicIncrease()

	// Check if the value is now 2
	if got := counter.Value(); got != 2 {
		t.Errorf("got %d, want %d", got, 2)
	}
}

func BenchmarkAtomicIncrease(b *testing.B) {
	counter := &Counter{}

	for b.Loop() {
		counter.AtomicIncrease()
	}
}

func BenchmarkIncrease(b *testing.B) {
	counter := &Counter{}

	for b.Loop() {
		counter.Increase()
	}
}

func TestIncrease(t *testing.T) {
	counter := &Counter{}

	// Increase the counter using the non-atomic method
	counter.Increase()

	// Check if the value is 1
	if got := counter.Value(); got != 1 {
		t.Errorf("got %d, want %d", got, 1)
	}

	// Increase the counter again using the non-atomic method
	counter.Increase()

	// Check if the value is now 2
	if got := counter.Value(); got != 2 {
		t.Errorf("got %d, want %d", got, 2)
	}
}
