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
