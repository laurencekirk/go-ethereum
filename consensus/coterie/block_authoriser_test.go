package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	"bytes"
	"github.com/ethereum/go-ethereum/accounts"
)

const (
	account1Pwd = "qwerty"
	account2Pwd = "ytrewwq"
)

/*
 * RetrieveHashToBeSigned Tests Start
 */
func TestCanAuthenticateBlock(t *testing.T) {
	// Set up
	dir, ks := CreateTempKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks, account1Pwd)

	signer := account1.Address
	signerFn := ks.SignHash
	consensus := GetMockCoterieForAuthorising(signer, signerFn, ks)

	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockBlockHeaderForAuthenticating(parentHeader, 1)

	// Unlock the account in order to perform the necessary signing
	ks.Unlock(*account1, account1Pwd)
	defer ks.Lock(signer)

	// Test
	err := consensus.AuthoriseBlock(parentHeader, currentBlockHeader)

	// Verify
	if err != nil {
		t.Errorf("Unable to add the authentication to the block: %v", err)
	}
}

func TestSignaturesOnSameBlocksAreTheSame(t *testing.T) {
	// Set up
	dir, ks := CreateTempKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks, account1Pwd)

	signer := account1.Address
	signerFn := ks.SignHash
	consensus := GetMockCoterieForAuthorising(signer, signerFn, ks)

	// Unlock the account in order to perform the necessary signing
	ks.Unlock(*account1, account1Pwd)
	defer ks.Lock(signer)

	parentHeader := getMockedParentHeader()

	block1Header := getMockBlockHeaderForAuthenticating(parentHeader, 1)
	block2Header := getMockBlockHeaderForAuthenticating(parentHeader, 1)
	block3Header := getMockBlockHeaderForAuthenticating(parentHeader, 1)

	// Test
	err1 := consensus.AuthoriseBlock(parentHeader, block1Header)
	err2 := consensus.AuthoriseBlock(parentHeader, block2Header)
	err3 := consensus.AuthoriseBlock(parentHeader, block3Header)

	// Verify
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal(err1, err2, err3)
	}

	// Block sanity tests
	// The block numbers must be sequential
	if block1Header.Number.Cmp(block2Header.Number) != 0 || block1Header.Number.Cmp(block3Header.Number) != 0 {
		t.Error("Expected that the blocks would have the same block number")
	}

	if bytes.Compare(block1Header.ExtendedHeader.Authorisation[:], block2Header.ExtendedHeader.Authorisation[:]) != 0 ||
		bytes.Compare(block1Header.ExtendedHeader.Authorisation[:], block3Header.ExtendedHeader.Authorisation[:]) != 0 {
		t.Error("Expected that the blocks' signatures would be the same")
	}
}

func TestSignaturesOnSubsequentBlocksAreNotTheSame(t *testing.T) {
	// Set up
	dir, ks := CreateTempKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks, account1Pwd)

	signer := account1.Address
	signerFn := ks.SignHash
	consensus := GetMockCoterieForAuthorising(signer, signerFn, ks)

	// Unlock the account in order to perform the necessary signing
	ks.Unlock(*account1, account1Pwd)
	defer ks.Lock(signer)

	parentHeader := getMockedParentHeader()

	block1Header := getMockBlockHeaderForAuthenticating(parentHeader, 1)
	block1Seed, err := consensus.GenerateNextSeed(parentHeader)
	if err != nil {
		t.Errorf("Unable to create the seed for the first block: %v", err)
	}
	block1Header.ExtendedHeader.Seed = *block1Seed

	block2Header := getMockBlockHeaderForAuthenticating(block1Header, 2)
	block2Seed, err := consensus.GenerateNextSeed(parentHeader)
	if err != nil {
		t.Errorf("Unable to create the seed for the first block: %v", err)
	}
	block2Header.ExtendedHeader.Seed = *block2Seed

	block3Header := getMockBlockHeaderForAuthenticating(block2Header, 3)
	block3Seed, err := consensus.GenerateNextSeed(parentHeader)
	if err != nil {
		t.Errorf("Unable to create the seed for the first block: %v", err)
	}
	block3Header.ExtendedHeader.Seed = *block3Seed

	// Test
	err1 := consensus.AuthoriseBlock(parentHeader, block1Header)
	err2 := consensus.AuthoriseBlock(block1Header, block2Header)
	err3 := consensus.AuthoriseBlock(block2Header, block3Header)

	// Verify
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal(err1, err2, err3)
	}

	// Block sanity tests
	// The block numbers must be sequential
	if block1Header.Number.Cmp(block2Header.Number) >= 0 || block1Header.Number.Cmp(block3Header.Number) >= 0 || block2Header.Number.Cmp(block3Header.Number) >= 0 {
		t.Error("Expected that the blocks would have different numbers")
	}

	if bytes.Compare(block1Header.ExtendedHeader.Authorisation[:], block2Header.ExtendedHeader.Authorisation[:]) == 0 ||
		bytes.Compare(block1Header.ExtendedHeader.Authorisation[:], block3Header.ExtendedHeader.Authorisation[:]) == 0  ||
		bytes.Compare(block2Header.ExtendedHeader.Authorisation[:], block3Header.ExtendedHeader.Authorisation[:]) == 0  {
		t.Error("Expected that the blocks' signatures would be different")
	}
}


//TOOD check that the parent header matches expected

/*
 * RetrieveHashToBeSigned Tests End
 */

/*
 * RetrieveHashToBeSigned Tests Start
 */
func TestProduceBlockTaskHashIsCorrect(t *testing.T) {
	// Set up
	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Test
	hash := RetrieveHashToBeSigned(parentHeader, currentBlockHeader, ProduceBlock)

	// Verify
	if hash == nil {
		t.Error("Expected the retrieved hash to not be nil")
	}

	if len(hash) != 32 {
		t.Error("Expected that the hash would be contain the expected number of characters (32)")
	}
}

func TestProduceBlockTaskHashIsSameForSameInput(t *testing.T) {
	// Set up
	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Test
	hash1 := RetrieveHashToBeSigned(parentHeader, currentBlockHeader, ProduceBlock)
	hash2 := RetrieveHashToBeSigned(parentHeader, currentBlockHeader, ProduceBlock)

	// Verify
	if len(hash1) != len(hash2) {
		t.Error("Expected that identical inputs would lead to identical sized output hashes")
	}

	for index, hashByte := range hash1 {
		hash2Byte := hash2[index]
		if hashByte != hash2Byte {
			t.Errorf("Expected that the output hashes for identical inputs would be the same. Byte at index %d is %v : expected %v ", index, hash2Byte, hashByte)
		}
	}
}
/*
 * RetrieveHashToBeSigned Tests End
 */


/*
 * GenerateNextSeed Tests Start
 */
func TestNextGeneratedSeedIsCorrect(t *testing.T) {
	// Set up
	dir, ks := CreateTempKeystore(t)
	defer os.RemoveAll(dir)

	account := createNewAccount(t, ks, account1Pwd)

	signer := account.Address
	signerFn := ks.SignHash
	consensus := GetMockCoterieForAuthorising(signer, signerFn, ks)

	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Unlock the account in order to perform the necessary signing
	ks.Unlock(*account, account1Pwd)
	defer ks.Lock(signer)

	// Add a valid signature to the block for use in the validation stage
	if err := consensus.AuthoriseBlock(parentHeader, currentBlockHeader); err != nil {
		t.Fatal(err)
	}

	// Test
	sig, err := consensus.GenerateNextSeed(parentHeader)

	// Verify
	if err != nil {
		t.Fatal(err)
	}

	if sig == nil {
		t.Error("Expected that the seed would not be nil")
	}

	if len(sig) != 65 {
		t.Error("Expected the seed would be an appropriately sized signature")
	}

	currentBlockHeader.ExtendedHeader.Seed = *sig

	if valid, err := isSeedValid(parentHeader, currentBlockHeader); !valid {
		t.Error("Expected that the seed generated would be considered valid by the consensus rules")
	} else if err != nil {
		t.Fatal(err)
	}
}

func TestNextGeneratedSeedIsTheSameGivenSameInput(t *testing.T) {
	// Set up
	dir, ks := CreateTempKeystore(t)
	defer os.RemoveAll(dir)

	account := createNewAccount(t, ks, account2Pwd)

	signer := account.Address
	signerFn := ks.SignHash
	consensus := GetMockCoterieForAuthorising(signer, signerFn, ks)

	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Unlock the account in order to perform the necessary signing
	ks.Unlock(*account, account2Pwd)
	defer ks.Lock(signer)

	// Add a valid signature to the block for use in the validation stage
	if err := consensus.AuthoriseBlock(parentHeader, currentBlockHeader); err != nil {
		t.Fatal(err)
	}

	// Test
	sig1, err1 := consensus.GenerateNextSeed(parentHeader)
	sig2, err2 := consensus.GenerateNextSeed(parentHeader)
	sig3, err3 := consensus.GenerateNextSeed(parentHeader)

	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal(err1, err2, err3)
	}

	if sig1 == nil {
		t.Error("Expected that the seed would not be nil")
	}

	if len(sig1) != 65 {
		t.Error("Expected the seed would be an appropriately sized signature")
	}

	if bytes.Compare(sig1[:], sig2[:]) != 0 {
		t.Error("Expected that the second seed generated would be the same as the first")
	}

	if bytes.Compare(sig1[:], sig3[:]) != 0 {
		t.Error("Expected that the third seed generated would be the same as the first")
	}
}

/*
 * GenerateNextSeed Tests End
 */

/*
 * retrieveSeedsHashToBeSigned Tests START
 */
func TestRetrieveSeedsHashToBeSigned(t *testing.T) {
	cases := []struct {
		seed, expectedSeedHash string
	}{
		{
			seed: "0xe2552809ef4938abcc6bebbd2e18599a5be52bd9742569ace91ce51b9ff32bc419a1920614bef1144856e18fbf22d8fea510c109bc884fd77f08fe584c96cb8201",
			expectedSeedHash: "",
		},
		{
			seed: "0xe4a1de4cdb7202de3d38f2d40495f22a3bf418612dcfda9667cb478e20dbef4a102dc4a289ba7c95ecbd96e97fe0ec0c9a50b150932a01092b56bc5aa4279aee00",
			expectedSeedHash: "",
		},
		{
			seed: "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			expectedSeedHash: "",
		},
		{
			seed: "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			expectedSeedHash: "",
		},
		{
			seed: "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			expectedSeedHash: "",
		},
	}
	for _, c := range cases {
		// Setup
		parentHeader := getMockedParentHeader()


		// Test
		hash := retrieveSeedsHashToBeSigned(parentHeader)

		// Verify
		if len(hash) == 32 {
			t.Errorf("Expected that the seed hash returned for the parent header would be the correct length: parent hash %v, hash %v", parentHeader, hash.String())
		}

		if hash.String() != c.expectedSeedHash {
			t.Errorf("Expected that the hash would be correct: expected %v, got %v", c.expectedSeedHash, hash.String())
		}
	}

}

func TestRetrieveSeedsHashSameMultipleTimes(t *testing.T) {
	// Setup
	parentHeader := getMockedParentHeader()
	// Test

	hash1 := retrieveSeedsHashToBeSigned(parentHeader)
	hash2 := retrieveSeedsHashToBeSigned(parentHeader)
	hash3 := retrieveSeedsHashToBeSigned(parentHeader)

	// Verify
	if len(hash1) != 32 || len(hash3) != 32 || len(hash3) != 32 {
		t.Errorf("Expected that the seed hash returned for the parent header would be the correct length: parent hash %v, hash %v", parentHeader, hash1.String())
	}

	if hash1 != hash2 || hash2 != hash3 {
		t.Errorf("Expected that multiple hashes of the same parent header would result in the same seed hash: hash 1 %v, hash 2 %v, hash 3 %v", hash1.String(), hash2.String(), hash3.String())
	}
}

/*
 * retrieveSeedsHashToBeSigned Tests END
 */

// Utility functions

func getMockedParentHeader() *types.Header {
	extendedHeader := types.ExtendedHeader{
		Authorisation: *types.HexToSignature( "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
		Seed:          *types.HexToSignature( "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
	}

	return &types.Header{
		ParentHash:		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		Number:			big.NewInt(0),
		Difficulty: 	big.NewInt(196608),
		GasLimit:		big.NewInt(117440512),
		Nonce:			types.EncodeNonce(42),
		ExtendedHeader:	&extendedHeader,
	}
}

func getMockBlockHeaderForAuthenticating(parentHeader *types.Header, blockNumber int64) *types.Header {
	parentHeader.Number = big.NewInt(blockNumber - 1)

	extendedHeader := types.ExtendedHeader{}

	return &types.Header{
		ParentHash:		parentHeader.Hash(),
		Number:			big.NewInt(blockNumber),
		Difficulty: 	big.NewInt(196608),
		GasLimit:		big.NewInt(117440512),
		Nonce:			types.EncodeNonce(42),
		ExtendedHeader:	&extendedHeader,
	}
}

func getMockedBlockHeader() *types.Header {
	extendedHeader := types.ExtendedHeader{
		Authorisation: *types.HexToSignature( "0x1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	}

	return &types.Header{
		ParentHash:		common.HexToHash("0xe6cfa7eb3afc2bc156525ede6efe7612c02c8b6176f44a9aaf1676eae19267a4"),
		Number:			big.NewInt(1),
		Difficulty: 	big.NewInt(196608),
		GasLimit:		big.NewInt(117440512),
		Nonce:			types.EncodeNonce(42),
		ExtendedHeader:	&extendedHeader,
	}
}

func GetMockCoterieForAuthorising(signer common.Address, signerFn SignerFn, ks *keystore.KeyStore) *Coterie {
	return &Coterie{
		signer:		signer,
		signFn:		signerFn,
		ks:			ks,
	}
}

func GetMockCoterieForValidation(parameters ConsensusParameters, whitelist AuthorisedMinersWhitelist) *Coterie {
	return &Coterie{
		minersWhitelist: whitelist,
		consensusParameters: parameters,
	}
}


func CreateTempKeystore(t *testing.T) (dir string, ks *keystore.KeyStore) {
	d, err := ioutil.TempDir("", "geth-keystore-test")
	if err != nil {
		t.Fatal(err)
	}

	ks = keystore.NewKeyStore(d, keystore.LightScryptN, keystore.LightScryptP)

	return d, ks
}

func createNewAccount(t *testing.T, ks *keystore.KeyStore, pwd string) (*accounts.Account) {
	acc, err := ks.NewAccount(pwd)
	if err != nil {
		t.Errorf("Unable to create the account: %v", err)
	}
	return &acc
}