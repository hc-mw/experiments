package syncdemo

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("increment counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("concurrently", func(t *testing.T) {
		wantedCount := 1000
		cntr := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				defer wg.Done()
				cntr.Inc()
			}()
		}
		wg.Wait()

		assertCounter(t, &cntr, wantedCount)
	})
}

func assertCounter(t testing.TB, cntr *Counter, wantedCount int) {
	t.Helper()

	if cntr.Value() != wantedCount {
		t.Errorf("got %d, want %d\n", cntr.Value(), wantedCount)
	}
}
