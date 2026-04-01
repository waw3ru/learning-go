package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := Sum(input)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, input)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{0, 3}, []int{1, 2})
	want := []int{3, 3}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTotalSum(t *testing.T) {
	got := TotalSum([]int{0, 3}, []int{1, 2})
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumTail(t *testing.T) {
	t.Run("make the sums of slice except for the first value", func(t *testing.T) {
		got := SumTail([]int{0, 3, 6, 8})
		want := 17

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("make the sums of empty slice", func(t *testing.T) {
		got := SumTail([]int{})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("return pair sums from slice of slices", func(t *testing.T) {
		got := SumAllTails([]int{0, 3, 6, 8}, []int{1, 2})
		want := []int{17, 2}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("return a pair of slices with 0", func(t *testing.T) {
		got := SumTail([]int{})
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
