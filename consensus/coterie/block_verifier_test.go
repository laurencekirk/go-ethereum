package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

/**
 * RetrieveBlockAuthor tests START
 */

 // To verify authorisation signature We require the seed from the parent's header and the block number - and the address that we expect to have authorised the block.
func TestRetrieveBlockAuthorValidSignature(t *testing.T) {
	cases := []struct {
		authorisation, parentSeed, addressExpected string
		blockNumber int64
	}{
		{
			"0xd817f3ba3956cbccd769afd17ef2b08a8ab7997d52cf5db55974c38c28fa24b90b814957b864a5adcca553cba2157574c39a7c10c342cb86c09d9603cdf037e200",
			"0x8c68c72f23975111ed783d3351730266489971f0fea49d43196ab7b482ce6fb6648d2f50dcfc5619c0dfa3f023882de27ee0f01fb98bc9c124c47cb11f757fe300",
			"0x30ff130a7d11ef9d1efbdf19d5309556acd129cf",
			2,
		},
		{
			"0x0860c7eeeaa9070d7c3b1116122671171846e2b5e8a29ae3a0b9b7637e281f173a30a148ba876379723c13a7de4a948645dd39b1b66e705fc2c5fe3eb7aa8a1301",
			"0xc723f849846dfb865134a28b7bb2e9b19b2c454ee8945ea41641e6809bb710da380da3edeecc0bde2a52e6245f3e2841bbdbaab95a6cb9828d20861d02c7f68b00",
			"0xea30250dd7263a4783c66463c236a2153d6b88b4",
			42,
			},
		{
			"0x1f31fb2878cf6cbefbecad4b05664a2af87d5ec725cfe9e6a2123bdd16cb1ece134a3647dd52d72d2ea4c010454d22fe39e5c875f666e33dfa33dabcf5bad5aa01",
			"0x884a3ad3e55ea7b9917b07128dc58c544d7a27611336e9ee445219fbadd413943677b96221c39a708c55fb4deaa0a80c3d12f9bf1d059028e958d73e31522b0101",
			"0x6c80e492308f051eba48d03bcc04625682ae3e07",
			100,
		},
	}
	for _, c := range cases {
		// Set up
		parentHeader := getMockedParentHeader()
		currentBlockHeader := getMockedBlockHeader()

		parentHeader.ExtendedHeader.Seed = *types.HexToSignature(c.parentSeed)

		currentBlockHeader.Number = big.NewInt(c.blockNumber)
		currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(c.authorisation)

		expectedAddress := common.HexToAddress(c.addressExpected)

		// Test
		retrievedAddress, err := RetrieveBlockAuthor(parentHeader, currentBlockHeader)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		if expectedAddress != retrievedAddress {
			t.Errorf("Retrieve address does not matched the expected address: expected %v, got %v", expectedAddress.String(), retrievedAddress.String())
		}
	}
}

/*
 * Scenario:
 * A valid signature has been taken from block Br, but the subsequently propagated block at the same block-height (r) as Br has been modified to point to a
 * different parent e.g. not Br-1.
 * The retrieved address should, therefore, not be that of the original authoriser
 */
func TestRetrieveBlockAuthorInvalidParentSeedValidBlockNumberIsNotValidated(t *testing.T) {
	cases := []struct {
		authorisation, parentSeed, addressExpected string
		blockNumber int64
	}{
		{
			"0xd817f3ba3956cbccd769afd17ef2b08a8ab7997d52cf5db55974c38c28fa24b90b814957b864a5adcca553cba2157574c39a7c10c342cb86c09d9603cdf037e200",
			"0xc723f849846dfb865134a28b7bb2e9b19b2c454ee8945ea41641e6809bb710da380da3edeecc0bde2a52e6245f3e2841bbdbaab95a6cb9828d20861d02c7f68b00",
			"0x30ff130a7d11ef9d1efbdf19d5309556acd129cf",
			2,
		},
		{
			"0x0860c7eeeaa9070d7c3b1116122671171846e2b5e8a29ae3a0b9b7637e281f173a30a148ba876379723c13a7de4a948645dd39b1b66e705fc2c5fe3eb7aa8a1301",
			"0x884a3ad3e55ea7b9917b07128dc58c544d7a27611336e9ee445219fbadd413943677b96221c39a708c55fb4deaa0a80c3d12f9bf1d059028e958d73e31522b0101",
			"0xea30250dd7263a4783c66463c236a2153d6b88b4",
			42,
		},
		{
			"0x1f31fb2878cf6cbefbecad4b05664a2af87d5ec725cfe9e6a2123bdd16cb1ece134a3647dd52d72d2ea4c010454d22fe39e5c875f666e33dfa33dabcf5bad5aa01",
			"0x8c68c72f23975111ed783d3351730266489971f0fea49d43196ab7b482ce6fb6648d2f50dcfc5619c0dfa3f023882de27ee0f01fb98bc9c124c47cb11f757fe300",
			"0x6c80e492308f051eba48d03bcc04625682ae3e07",
			100,
		},
	}
	for _, c := range cases {
		// Set up
		parentHeader := getMockedParentHeader()
		currentBlockHeader := getMockedBlockHeader()

		parentHeader.ExtendedHeader.Seed = *types.HexToSignature(c.parentSeed)

		currentBlockHeader.Number = big.NewInt(c.blockNumber)
		currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(c.authorisation)

		expectedAddress := common.HexToAddress(c.addressExpected)

		// Test
		retrievedAddress, err := RetrieveBlockAuthor(parentHeader, currentBlockHeader)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		if expectedAddress == retrievedAddress {
			t.Errorf("Retrieve address does not matched the expected address: expected %v, got %v", expectedAddress.String(), retrievedAddress.String())
		}
	}
}

/*
 * Scenario:
 * In the unlikely situation that a valid seed is produced in more than one block (say that the seed Sa was first seen in Block Sa),
 * then we want to make sure that re-using the authorisation signature (Aa) in block (Br) doesn't result in the consensus layer thinking that the
 * authorised account has produced block Br.
 * The retrieved address should, therefore, not be that of the original authoriser
 */
func TestRetrieveBlockAuthorValidParentSeedInvalidBlockNumberIsNotValidated(t *testing.T) {
	cases := []struct {
		authorisation, parentSeed, addressExpected string
		blockNumber int64
	}{
		{
			"0xd817f3ba3956cbccd769afd17ef2b08a8ab7997d52cf5db55974c38c28fa24b90b814957b864a5adcca553cba2157574c39a7c10c342cb86c09d9603cdf037e200",
			"0x8c68c72f23975111ed783d3351730266489971f0fea49d43196ab7b482ce6fb6648d2f50dcfc5619c0dfa3f023882de27ee0f01fb98bc9c124c47cb11f757fe300",
			"0x30ff130a7d11ef9d1efbdf19d5309556acd129cf",
			3,
		},
		{
			"0x0860c7eeeaa9070d7c3b1116122671171846e2b5e8a29ae3a0b9b7637e281f173a30a148ba876379723c13a7de4a948645dd39b1b66e705fc2c5fe3eb7aa8a1301",
			"0xc723f849846dfb865134a28b7bb2e9b19b2c454ee8945ea41641e6809bb710da380da3edeecc0bde2a52e6245f3e2841bbdbaab95a6cb9828d20861d02c7f68b00",
			"0xea30250dd7263a4783c66463c236a2153d6b88b4",
			41,
		},
		{
			"0x1f31fb2878cf6cbefbecad4b05664a2af87d5ec725cfe9e6a2123bdd16cb1ece134a3647dd52d72d2ea4c010454d22fe39e5c875f666e33dfa33dabcf5bad5aa01",
			"0x884a3ad3e55ea7b9917b07128dc58c544d7a27611336e9ee445219fbadd413943677b96221c39a708c55fb4deaa0a80c3d12f9bf1d059028e958d73e31522b0101",
			"0x6c80e492308f051eba48d03bcc04625682ae3e07",
			9001,
		},
	}
	for _, c := range cases {
		// Set up
		parentHeader := getMockedParentHeader()
		currentBlockHeader := getMockedBlockHeader()

		parentHeader.ExtendedHeader.Seed = *types.HexToSignature(c.parentSeed)

		currentBlockHeader.Number = big.NewInt(c.blockNumber)
		currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(c.authorisation)

		expectedAddress := common.HexToAddress(c.addressExpected)

		// Test
		retrievedAddress, err := RetrieveBlockAuthor(parentHeader, currentBlockHeader)

		// Verify
		if err != nil {
			t.Fatal(err)
		}

		if expectedAddress == retrievedAddress {
			t.Errorf("Retrieve address does not matched the expected address: expected %v, got %v", expectedAddress.String(), retrievedAddress.String())
		}
	}
}

/**
 * RetrieveBlockAuthor tests END
 */

/*
 * isSeedValid Tests Start
 */
func TestValidSeedIsCorrectlyVerified(t *testing.T) {
	// Set up
	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Valid seed for the address 0x343Baf7D51e5aeB0D5939926EBF93c9404A01d83 based on the parent seed "FFF...FFF"
	validSeed := "0xf10430be4b33312ff53245336e4a6e5df462268cb152374deebfdabccdfc4cbe18847c370c43fa37554c3bdb2d2a724a4199cb30042ea727b529824dc809a13601"
	// Signature for the address 0x343Baf7D51e5aeB0D5939926EBF93c9404A01d83
	validSignature := "0x14a625c5a7452ac86f852e8f29e854b8cde398e7037346cc6face77f51a2f0e62141f4d74b921f4e0cafe8e51315fdc2d6198d0fa7a2d1ccb9cef64b8252c63f00"

	currentBlockHeader.ExtendedHeader.Seed = *types.HexToSignature(validSeed)
	currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(validSignature)

	// Test
	valid, err := isSeedValid(parentHeader, currentBlockHeader)

	// Verify
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Error("Expected that the seed would be considered valid")
	}
}


/*
 * Scenario:
 * The previous seed Qr-1 has been taken from block Br(pk1) signed by the address pk1 and used in a different block (at he same block height) Br(pk2).
 * Even though the seed (Qr-1) is valid, the private key used to authorise the block Sig(pk2) contains the public key for a different address to pk1.
 * Therefore, the mismatch between the two public keys / addresses should result in this seed not being considered valid.
 */
func TestValidButIncorrectSeedForBlockIsNotVerified(t *testing.T) {
	// Set up
	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	previouslySeenValidSeed := "0x5724245b70dbf42f029c57e57aab2d05f87b00652734465174b61aeb8bac86bd4d18463a6afa6e110c2781805af712c65b02cffa04a79b38cafda6bde759a4f100"
	currentBlockHeader.ExtendedHeader.Seed = *types.HexToSignature(previouslySeenValidSeed)

	// Authorisation for the address 0x2ef57dae52a8637bAee02A3178880c65a208Af89
	authorisationForDifferentAddress := "0xb08a3d7fd4701e7c2831aff7e372bbcacf0975270b12e8bf46d7a66b443a692c2e7dbe38386813ddb46439316bb00de79dd5a47ee54b6ecd514d77e599cece8600"
    currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(authorisationForDifferentAddress)

	// Test
	valid, err := isSeedValid(parentHeader, currentBlockHeader)

	// Verify
	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Error("Expected that the seed would not be considered valid")
	}
}

/*
 * Scenario:
 * The seed for a previously seen block, that is not the most recent parent, (not Br-1) is reused. Even if the same account re-uses a seed e.g. Qr-2(pk1) (the seed from two blocks ago that was generated by pk1)
 * it should not be considered valid.
 */
func TestSeedForAuthorisedButFromPreviousBlockIsNotVerified(t *testing.T) {
	// Set up
	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Seed for the parent seed "0FF...FFF" created by the author 0x18238f0d6F9A6765B0972AbB9F54A51d1d583503
	previouslySeenValidSeed := "0x61e7302061a8f15b677284e49f3228f862e71a27a4d87b8e86a867ef04c19c7a50330a142249c3626eeadce258d9d5e623882703d2ae0e1284ecb1b32071516100"
	currentBlockHeader.ExtendedHeader.Seed = *types.HexToSignature(previouslySeenValidSeed)

	// Authorisation for the address 0x18238f0d6F9A6765B0972AbB9F54A51d1d583503
	authorisationForDifferentAddress := "0x69716421b4fae3c4d1b3412c86be4d22ec3d062cb69f32b8a8602d20a21fefd32a41d000383e1d5d2263067fafbc861a7e50b2425784971fe2f229e4bdff53a201"
	currentBlockHeader.ExtendedHeader.Authorisation = *types.HexToSignature(authorisationForDifferentAddress)

	// Test
	valid, err := isSeedValid(parentHeader, currentBlockHeader)

	// Verify
	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Error("Expected that the seed would not be considered valid")
	}

	// Double check that test is not false-negative
	// Seed used to create previouslySeenValidSeed above
	parentHeader.ExtendedHeader.Seed = *types.HexToSignature("0x0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")

	// Test
	valid, err = isSeedValid(parentHeader, currentBlockHeader)

	// Verify
	if err != nil {
		t.Fatal(err)
	}

	if ! valid {
		t.Error("Expected that given the modified parent seed that the new seed would be considered valid")
	}
}

/*
* isSeedValid Tests END
*/