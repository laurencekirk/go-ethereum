package coterie

import (
	"testing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
)

const (
	account1Pwd = "qwerty"
)

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

	account, err := ks.NewAccount(account1Pwd)
	if err != nil {
		t.Fatal(err)
	}

	signer := account.Address
	signerFn := ks.SignHash
	consensus := GetMockCoterieForAuthorising(signer, signerFn, ks)

	parentHeader := getMockedParentHeader()
	currentBlockHeader := getMockedBlockHeader()

	// Unlock the account in order to perform the necessary signing
	ks.Unlock(account, account1Pwd)
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
		t.Error("Expected that the seed would not be null")
	}

	if len(sig) != 65 {
		t.Error("Expected the seed would be an appropriately sized signature")
	}

	output := sig.String()
	if len(output) == 0 {

	}

	currentBlockHeader.ExtendedHeader.Seed = *sig

	if valid, err := isSeedValid(parentHeader, currentBlockHeader); !valid {
		t.Error("Expected that the seed generated would be considered valid by the consensus rules")
	} else if err != nil {
		t.Fatal(err)
	}
}

/*
 * GenerateNextSeed Tests End
 */

// Utility functions

func getMockedParentHeader() *types.Header {
	extendedHeader := types.ExtendedHeader{
		Signature:		*types.HexToSignature( "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
		Seed:			*types.HexToSignature( "0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
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

func getMockedBlockHeader() *types.Header {
	extendedHeader := types.ExtendedHeader{
		Signature:		*types.HexToSignature( "0x1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
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


func CreateTempKeystore(t *testing.T) (dir string, ks *keystore.KeyStore) {
	d, err := ioutil.TempDir("", "geth-keystore-test")
	if err != nil {
		t.Fatal(err)
	}

	ks = keystore.NewKeyStore(d, keystore.LightScryptN, keystore.LightScryptP)

	return d, ks
}