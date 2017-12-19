package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"math"
	"github.com/golang/mock/gomock"
	"github.com/ethereum/go-ethereum/consensus/coterie/mocks"
	"github.com/ethereum/go-ethereum/common"
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

func TestRemoveLeadingZeroDigitsHandlesErrors(t *testing.T) {
	cases := []struct {
		in string
	}{
		{""},
		{" 0x1"},
		{"0"},
		{"G"},
		{"0xG"},
	}
	for _, c := range cases {
		_, err := removeLeadingZeroDigits(c.in)
		if err == nil {
			t.Error("Expected that removeLeadingZeroDigits would reject the invalid string")
		}
	}
}
/**
 * removeLeadingZeroDigits tests END
 */

/**
 * calculateWinningThreshold tests START
 */
func TestCalculateWinningThreshold(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		committeeSize, whitelistSize uint
		expectedThreshold float64
	}{
		{committeeSize: uint(20), whitelistSize: uint(100), expectedThreshold: float64(0.2)},
		{committeeSize: uint(7), whitelistSize: uint(8), expectedThreshold: float64(0.875)},
		{committeeSize: uint(3926), whitelistSize: uint(93716), expectedThreshold: float64(0.04189252635622519)},
	}
	for _, c := range cases {
		// Setup
		whitelist := mocks.NewMockAuthorisedWhitelist(ctrl)
		params := mocks.NewMockConsensusParameters(ctrl)

		params.EXPECT().GetTargetCommitteeSize().Return(c.committeeSize, nil).Times(1)
		whitelist.EXPECT().GetWhitelistSize().Return(c.whitelistSize, nil).Times(1)

		// Test
		threshold, err := calculateWinningThreshold(params, whitelist)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		assertValueInCorrectRange(threshold, t)
		assertNearEquals(c.expectedThreshold, threshold, t)
	}
}
/**
 * calculateWinningThreshold tests START
 */

/**
 * hasWonSecretLottery tests START
 */
func TestHasWonSecretLottery(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		committeeSize, whitelistSize uint
		authorisation string
		expectedToBeSelected bool
	}{
		{	committeeSize: uint(20),
			whitelistSize: uint(100), // committee threshold of ~ 0.2
			authorisation: "0x49e58681b5cb510332393cad62722dd3374773f1cd37d75066e1b0880d73f9b273ab2913bf4160fd066b9a77d1b8b690631ec7dc3db012741a7ba24d9c25edff00",  // ~ 0.28
			expectedToBeSelected: false,
		},
		{	committeeSize: uint(20),
			whitelistSize: uint(100), // committee threshold of ~ 0.2
			authorisation: "0x30A3D70A4FFFF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  // ~ 0.19
			expectedToBeSelected: true,
		},
		{	committeeSize: uint(9),
			whitelistSize: uint(10), // committee threshold of ~ 0.9
			authorisation: "0xFEFEFEFEFEFEF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  // ~ 0.996078431
			expectedToBeSelected: false,
		},
		{	committeeSize: uint(9),
			whitelistSize: uint(10), // committee threshold of ~ 0.9
			authorisation: "0xE666666666666000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  // ~ 0.9
			expectedToBeSelected: true,
		},
		{	committeeSize: uint(12987123),
			whitelistSize: uint(94982709308), // committee threshold of ~ 0.000136731
			authorisation: "0x0008E9B390000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  // ~ 0.000136
			expectedToBeSelected: true,
		},
		{	committeeSize: uint(12987123),
			whitelistSize: uint(94982709308), // committee threshold of ~ 0.000136731
			authorisation: "0x00092CCF70000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",  // ~ 0.00014
			expectedToBeSelected: false,
		},

	}
	for _, c := range cases {
		// Setup
		whitelist := mocks.NewMockAuthorisedWhitelist(ctrl)
		params := mocks.NewMockConsensusParameters(ctrl)

		params.EXPECT().GetTargetCommitteeSize().Return(c.committeeSize, nil).Times(1)
		whitelist.EXPECT().GetWhitelistSize().Return(c.whitelistSize, nil).Times(1)

		signature := types.HexToSignature(c.authorisation)
		consensus := GetMockCoterieForValidation(params, whitelist)

		// Test
		selected, err := consensus.hasWonSecretLottery(signature)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		if c.expectedToBeSelected != selected {
			t.Errorf("Expected that the authorisation would result in the block producer being selected: auth %v, whitelist size %v, committee size %v", c.authorisation, c.whitelistSize, c.committeeSize)
		}
	}
}
/**
 * hasWonSecretLottery tests END
 */

/**
 * HasBeenSelectedToCommittee tests START
 */
func TestHasBeenSelectedToCommittee(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		committeeSize, whitelistSize uint
		address, authorisation string
		addressInWhitelist, expectedToBeSelected bool
	}{
		{	committeeSize: uint(400),
			whitelistSize: uint(500), // committee threshold of ~ 0.8
			address: "0x9e5e939fb0a23529934c061d6ecf4c93e7893d4e",
			authorisation: "0xcb49c9f1a18340b5a8806db212d0e7d0be3605d257fe7a4814a5b3c84f90e53323d84e8a317651dbd69a05d3038518a01842001cd9024b0f4642182e359809ff00",  // ~ 0.794094678
			addressInWhitelist: true,
			expectedToBeSelected: true,
		},
		{	committeeSize: uint(98),
			whitelistSize: uint(274), // committee threshold of ~ 0.362
			address: "0x30ff130a7d11ef9d1efbdf19d5309556acd129cf",
			authorisation: "0x5c7a424a53c7e82625e727d23c676d1dd307b79d1ed3708847723dd91b66d1ae08f56c9efc6b8422c302bb941f12c114179d86a08cb2aedf7d29ef1d060f858b01",  // ~ 0.361
			addressInWhitelist: true,
			expectedToBeSelected: false,
		},
		{	committeeSize: uint(2378),
			whitelistSize: uint(62363), // committee threshold of ~ 0.03813
			address: "0x46dfb921f8f7edbbd8100458b7c1beefeabf6e15",
			authorisation: "0x09c336764c106475a1f5ea01abc7fe39d7474ec461a8a88a04fd6bdaf066f3080c98eaf50de01be86afbdf64a47d4e1d36fd8c341666e0372b98516fe7ae89cd00",  // ~ 0.0381
			addressInWhitelist: true,
			expectedToBeSelected: false,
		},
		{	committeeSize: uint(0),
			whitelistSize: uint(0),
			address: "0xea30250dd7263a4783c66463c236a2153d6b88b4",
			authorisation: "0x1335dd6905c3f7a49d570fe7d31596801a6a566819b6c0443ae76b47c5f65be517d1bdfccb54f3e90d40a2654dbad98d1d9096482e8d06e9f9437ec5e8b7df9400",
			addressInWhitelist: false,
			expectedToBeSelected: false,
		},
	}
	for _, c := range cases {
		// Setup
		whitelist := mocks.NewMockAuthorisedWhitelist(ctrl)
		params := mocks.NewMockConsensusParameters(ctrl)

		signer := common.HexToAddress(c.address)

		whitelist.EXPECT().IsMinerInWhitelist(signer).Return(c.addressInWhitelist, nil).Times(1)

		if c.addressInWhitelist {
			params.EXPECT().GetTargetCommitteeSize().Return(c.committeeSize, nil).Times(1)
			whitelist.EXPECT().GetWhitelistSize().Return(c.whitelistSize, nil).Times(1)
		}

		signature := types.HexToSignature(c.authorisation)
		consensus := GetMockCoterieForValidation(params, whitelist)

		// Test
		selected, err := consensus.HasBeenSelectedToCommittee(signer, signature)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		if c.expectedToBeSelected != selected {
			if selected {
				t.Errorf("Expected that the authorisation would result in the block producer being selected: auth %v, whitelist size %v, committee size %v", c.authorisation, c.whitelistSize, c.committeeSize)
			} else {
				t.Errorf("Expected that the authorisation would result in the block producer NOT being selected: auth %v, whitelist size %v, committee size %v", c.authorisation, c.whitelistSize, c.committeeSize)
			}

		}
	}
}
/**
 * HasBeenSelectedToCommittee tests START
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
	if math.Abs(expected-actual) > TOLERANCE && math.Abs(actual-expected) > TOLERANCE {
		t.Errorf("Expected %q, Actual %q. Was not within the tolerence `%q`", expected, actual, TOLERANCE)
	}
}
/**
 * Testing utility functions END
 */
