package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
)

const TOLERANCE = 0.000000000000001

/**
 * calculateSignaturesRealValue tests START
 */
func TestConversionToRealValueOutput(t *testing.T) {
	cases := []struct {
		in string
		want float64
	}{
		{"0xFFFFFFFFFFFFF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.9999999999999999},
		{"0x0000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.0000000000000002},
		{"0x1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.0625},
		{"0x1111111111111000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0.0666666666666666},
		{"0x49e58681b5cb510332393cad62722dd3374773f1cd37d75066e1b0880d73f9b273ab2913bf4160fd066b9a77d1b8b690631ec7dc3db012741a7ba24d9c25edff00", 0.2886585299182063},
		{"0x6C80e492308f051EBA48D03bCC04625682aE3E07507824e5c292cb8bc02a7736458e5e8cfc3d046d0cacd21a629011188340e2d06ea1dd857d28505c7892ed5800", 0.4238417488964466},
		{"0xf3da58f390703e9219b309e756302a7db7576177507824e5c292cb8bc02a7736458e5e8cfc3d046d0cacd21a629011188340e2d06ea1dd857d28505c7892ed5800", 0.9525504679335777},
		{"0xb04a55bc2d24894879f4b79f7c3fc7e2701b4ec961f56b2badf2107a1072570e3f623f1f403b8fb48ade56b5d39e95c5bb5ad9b260de451d94aa5afcfe58335f01", 0.688634260598649},
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
/**
 * calculateSignaturesRealValue tests END
 */

/**
 * removeLeadingZeroDigits tests START
 */
func TestRemoveLeadingZeroDigits(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"0x49e58681b5cb5103", "0x49e58681b5cb5103"},
		{"0x0000000000001000", "0x1000"},
		{"0x0000000000000001", "0x1"},
		{"0x1000000000000000", "0x1000000000000000"},
		{"0x0000000000000000", "0x0"},
	}
	for _, c := range cases {
		got, err := removeLeadingZeroDigits(c.in)
		if err != nil {
			t.Error(err)
		}
		if got != c.want {
			t.Errorf("Expected removeLeadingZeroDigits would output the correct value: expected %v, got %v", c.want, got)
		}
	}
}
/**
 * removeLeadingZeroDigits tests END
 */

/**
 * Testing utility functions START
 */
func assertValueInCorrectRange(value float64, t *testing.T) {
	if value <= 0 || value >= 1 {
		t.Errorf("The signature's real value should be in the interval (0, 1): got %v", value)
	}
}

func assertNearEquals(expected float64, actual float64, t *testing.T) {
	if math.Abs(expected - actual) > TOLERANCE && math.Abs(actual - expected) > TOLERANCE {
		t.Errorf("Expected %q, Actual %q. Was not within the tolerence `%q`", expected, actual, TOLERANCE)
	}
}

/**
 * Testing utility functions END
 */
