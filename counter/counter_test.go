package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	assert := func(t testing.TB, got *Counter, want int) {
		t.Helper()
		if got.Value() != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		wantCount := 20000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}

		wg.Wait()

		assert(t, counter, wantCount)
	})
}
