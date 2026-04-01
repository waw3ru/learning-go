package shapes

import (
	"testing"

	"github.com/waw3ru/learning-go/utils"
)

type AssetType string

const (
	Perimeter AssetType = "Perimeter"
	Area      AssetType = "Area"
)

type TableTest struct {
	assert AssetType
	in     Shape
	out    float64
}

func TestShapes(t *testing.T) {
	assertHelper := func(t testing.TB, tt TableTest) {
		t.Helper()
		var got float64
		switch tt.assert {
		case Perimeter:
			got = utils.RoundTo(tt.in.Perimeter(), 2)
		case Area:
			got = utils.RoundTo(tt.in.Area(), 1)
		}

		if got != tt.out {
			t.Errorf("[AssertionType %q] got %.2f want %.2f", tt.assert, got, tt.out)
		}
	}

	tableTest := map[string]TableTest{
		"should get the perimeter of a rectangle": {
			assert: Perimeter,
			in:     Rectangle{10.0, 20.0},
			out:    60.0,
		},
		"should get the area of a rectangle": {
			assert: Area,
			in:     Rectangle{Width: 10.0, Height: 40.0},
			out:    400.0,
		},
		"should get the perimeter of a circle": {
			assert: Perimeter,
			in:     Circle{7},
			out:    43.98,
		},
		"should get the area of a circle": {
			assert: Area,
			in:     Circle{10.0},
			out:    314.20,
		},
	}

	for name, tc := range tableTest {
		t.Run(name, func(t *testing.T) {
			assertHelper(t, tc)
		})
	}
}
