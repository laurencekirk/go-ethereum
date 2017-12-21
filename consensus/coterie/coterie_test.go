package coterie

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/ethereum/go-ethereum/consensus/coterie/mocks"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
)

/**
 * AuthorisedBy tests START
 */
func TestAuthorisedByCorrectlyReturnsAuthor(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	consensus := &Coterie{}

	cases := []struct {
		authorisation, parentSeed, addressExpected string
		blockNumber int64
	}{
		{
			authorisation: "0x4a363879a18bf248703d1158b70cb06c5467769f97fd747e864a86dad1c8db79029fac3d3e1e20f29e5e9faf1272666009667c9171b08afb74aefdab25fc943700",
			parentSeed: "0xfcd1b63a1e3021fd26d1cae9f042e22edc36026e56f0417a3d22e9f214b77269048931d6b1ac59ea1280c7b21d2944e7f25c3528ef800034f74bba1681c7209400",
			addressExpected: "0x28955e2e5584194939af9223702482298e08055d",
			blockNumber: 48778,
		},
		{
			authorisation: "0x9376b4faed0bf0b95f75848fcaea79b176c84ff5cd3e35193487da814a3c86c030019a83011df27e2be8a98f57a414bdb0143e311ec733def7e1108fd6e9d55b00",
			parentSeed: "0x01b2510568c61e3b5acb048d4ed497244d0f247a7aaf3b4e92d2bae95b2611ae7a915d7b2095d379665cb4137c0b0746722c60e5acf89793f25bf811c8e18f8000",
			addressExpected: "0x46dfb921f8f7edbbd8100458b7c1beefeabf6e15",
			blockNumber: 18278,
		},
	}
	for _, c := range cases {
		// Setup
		parentHeader := getMockedParentHeader()
		currentBlockHeader := getMockBlockHeaderForAuthenticating(parentHeader, c.blockNumber)

		parentHeader.ExtendedHeader.Seed = *types.HexToSignature(c.parentSeed)

		currentBlockHeader.Number = big.NewInt(c.blockNumber)
		currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(c.authorisation)

		mockChaiReader := mocks.NewMockChainReader(ctrl)
		mockChaiReader.EXPECT().GetHeader(currentBlockHeader.ParentHash, uint64(c.blockNumber)-1).Return(parentHeader).Times(1)

		// Test
		retrievedAddress, err := consensus.AuthorisedBy(mockChaiReader, currentBlockHeader)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		if retrievedAddress != common.HexToAddress(c.addressExpected) {
			t.Errorf("Expected that the retrieved address would match the expected address: expected %v, got %v", c.addressExpected, retrievedAddress.String())
		}
	}

}


/**
 * AuthorisedBy tests END
 */

/**
* VerifySeal tests START
*/
func TestVerifySealValidScenarios(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		authorisation, curentSeed, parentSeed, addressExpected string
		blockNumber int64
	}{
		{
			authorisation: "0x4a363879a18bf248703d1158b70cb06c5467769f97fd747e864a86dad1c8db79029fac3d3e1e20f29e5e9faf1272666009667c9171b08afb74aefdab25fc943700",
			curentSeed: "0xe4a1de4cdb7202de3d38f2d40495f22a3bf418612dcfda9667cb478e20dbef4a102dc4a289ba7c95ecbd96e97fe0ec0c9a50b150932a01092b56bc5aa4279aee00",
			parentSeed: "0xfcd1b63a1e3021fd26d1cae9f042e22edc36026e56f0417a3d22e9f214b77269048931d6b1ac59ea1280c7b21d2944e7f25c3528ef800034f74bba1681c7209400",
			addressExpected: "0x28955e2e5584194939af9223702482298e08055d",
			blockNumber: 48778,
		},
		{
			authorisation: "0x9376b4faed0bf0b95f75848fcaea79b176c84ff5cd3e35193487da814a3c86c030019a83011df27e2be8a98f57a414bdb0143e311ec733def7e1108fd6e9d55b00",
			curentSeed: "0x02207acfb438d0bc93ab24fe8c39f6787a9a14ae1e77e0662b1236a88e9919a6062692e741e2164b1f5c07a8d07117669e4ed90d5f01b6100251098705c0e28d01",
			parentSeed: "0x01b2510568c61e3b5acb048d4ed497244d0f247a7aaf3b4e92d2bae95b2611ae7a915d7b2095d379665cb4137c0b0746722c60e5acf89793f25bf811c8e18f8000",
			addressExpected: "0x46dfb921f8f7edbbd8100458b7c1beefeabf6e15",
			blockNumber: 18278,
		},
	}
	for _, c := range cases {

		// Setup
		parentHeader := getMockedParentHeader()
		currentBlockHeader := getMockBlockHeaderForAuthenticating(parentHeader, c.blockNumber)

		parentHeader.ExtendedHeader.Seed = *types.HexToSignature(c.parentSeed)

		currentBlockHeader.Number = big.NewInt(c.blockNumber)
		currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(c.authorisation)
		currentBlockHeader.ExtendedHeader.Seed = *types.HexToSignature(c.curentSeed)

		whitelist := mocks.NewMockAuthorisedWhitelist(ctrl)
		params := mocks.NewMockConsensusParameters(ctrl)
		mockChaiReader := mocks.NewMockChainReader(ctrl)
		mockConsensus := mocks.NewMockEngine(ctrl)

		consensus := GetMockCoterieForValidation(params, whitelist)
		consensus.secondLayerConsensusEngine = mockConsensus

		mockChaiReader.EXPECT().GetHeader(currentBlockHeader.ParentHash, uint64(c.blockNumber)-1).Return(parentHeader).Times(1)
		whitelist.EXPECT().IsMinerInWhitelist(common.HexToAddress(c.addressExpected)).Return(true, nil).Times(1)
		params.EXPECT().GetTargetCommitteeSize().Return(uint(100), nil).Times(1)
		whitelist.EXPECT().GetWhitelistSize().Return(uint(99), nil).Times(1)	// Create a threshold of 0.99 - an almost certainty
		mockConsensus.EXPECT().VerifySeal(mockChaiReader, currentBlockHeader).Return(nil).Times(1)

		// Test
		err := consensus.VerifySeal(mockChaiReader, currentBlockHeader)

		// Verify
		if err != nil {
			t.Fatal(err)
		}
	}
}
/**
 * VerifySeal tests END
 */
