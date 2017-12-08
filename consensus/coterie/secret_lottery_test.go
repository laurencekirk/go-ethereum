package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
)

const TOLERANCE = 0.000000000000001

func Test(t *testing.T) {
	cases := []struct {
		in string
		want float64
	}{
		{"0xFFFFFFFFFFFFF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.9999999999999999},
		{"0x0000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.0000000000000002},
		{"0x1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.0625},
		{"0x1111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.0666666666666666},
		{"0x49e58681b5cb510332393cad62722dd3374773f1cd37d75066e1b0880d73f9b273ab2913bf4160fd066b9a77d1b8b690631ec7dc3db012741a7ba24d9c25edff00", 0.2886585299182063},
	}
	for _, c := range cases {
		sigature := types.HexToSignature(c.in)
		got, err := calculateSignaturesRealValue(sigature)
		if err != nil {
			t.Error(err)
		}
		assertValueInCorrectRange(got, t)
		assertNearEquals(c.want, got, t)
	}
}

func assertValueInCorrectRange(value float64, t *testing.T) {
	if value <= 0 || value >= 1 {
		t.Error("The signature's real value should be in the interval (0, 1)")
	}
}

func assertNearEquals(expected float64, actual float64, t *testing.T) {
	if math.Abs(expected - actual) > TOLERANCE && math.Abs(actual - expected) > TOLERANCE {
		t.Errorf("Expected %q, Actual %q. Was not within the tolerence `%q`", expected, actual, TOLERANCE)
	}
}

