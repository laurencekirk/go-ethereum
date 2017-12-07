package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
)



func Test(t *testing.T) {
	/*cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := calculateSignaturesRealValue(c.in)
		assertValueInCorrectRange(got, t)
		if got != c.want {
			t.Errorf("calculateSignaturesRealValue(%q) == %q, want %q", c.in, got, c.want)
		}
	}*/
	hexString := "0x49e58681b5cb510332393cad62722dd3374773f1cd37d75066e1b0880d73f9b273ab2913bf4160fd066b9a77d1b8b690631ec7dc3db012741a7ba24d9c25edff00"
	sig := types.HexToSignature(hexString)
	foo := calculateSignaturesRealValue(&sig)
	assertValueInCorrectRange(foo, t)
}

func assertValueInCorrectRange(value float64, t *testing.T) {
	if value <= 0 || value >= 1 {
		t.Error("The signature's real value should be in the interval (0, 1)")
	}
}

