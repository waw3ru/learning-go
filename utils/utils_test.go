package utils

import "testing"

const Pi = 3.141592653589793

func TestRoundTo(t *testing.T) {
	t.Run("should round off PI to 3 decimal places", func(t *testing.T) {
		got := RoundTo(Pi, 3)
		want := 3.142

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
