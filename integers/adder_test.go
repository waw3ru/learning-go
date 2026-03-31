package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Adder(2, 5)
	want := 7

	if sum != want {
		t.Errorf("expected '%d' but got '%d'", want, sum)
	}
}
