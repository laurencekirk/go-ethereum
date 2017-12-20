package coterie

/*import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
	"github.com/ethereum/go-ethereum/core/types"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"bytes"
	"github.com/ethereum/go-ethereum/crypto"
)




// Check to make sure that if a miner that is in the whitelist mines a block, that the verification process succeeds - the block is considered valid
func TestCanAuthoriseAndVerifyABlock(t *testing.T) {
	dir, ks := createKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks)
	whitelistMiner(t, account1)

	am := createAccountManager(ks)
	block := getNewBlock(NONCE_VALUE_2)

	if block.ExtendedHeader() != nil {
		t.Error("Expected that the block would not have been signed yet.")
	}

	err := AuthenticateBlock(block , am, &account1.Address, ACCOUNT_PASSWORD)
	if err != nil {
		t.Errorf("Unable to add the authentication to the block: %v", err)
	}

	valid := validateAuthentication(t, block)

	if !valid {
		t.Error("Expected that the block would have a *valid* signature.")
	}
}

// Check to make sure that if a miner that is not in the whitelist mines a block, that the verification process fails - the block is considered invalid
func TestCanNotVerifyABlockAuthorisedByAMinerNotInTheWhitelist(t *testing.T) {
	dir, ks := createKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks)

	am := createAccountManager(ks)
	block := getNewBlock(NONCE_VALUE_3)

	if block.ExtendedHeader() != nil {
		t.Error("Expected that the block would not have been signed yet.")
	}

	err := AuthenticateBlock(block , am, &account1.Address, ACCOUNT_PASSWORD)
	if err != nil {
		t.Errorf("Unable to add the authentication to the block: %v", err)
	}

	valid := validateAuthentication(t, block)

	if valid {
		t.Error("Expected that the block would *not* have been considered authentic because the miner is not in the whitelist.")
	}
}

func TestCanNotInjectValidSignatureIntoANewBlock(t *testing.T) {
	dir, ks := createKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks)
	whitelistMiner(t, account1)

	am := createAccountManager(ks)
	block1 := getNewBlock(NONCE_VALUE_1)

	if block1.ExtendedHeader() != nil {
		t.Error("Expected that the block would not have been signed yet.")
	}

	err := AuthenticateBlock(block1 , am, &account1.Address, ACCOUNT_PASSWORD)
	if err != nil {
		t.Errorf("Unable to add the authentication to the block: %v", err)
	}

	valid := validateAuthentication(t, block1)

	if !valid {
		t.Error("Expected that the block would have a *valid* signature.")
	}

	block2 := getNewBlock(NONCE_VALUE_2)

	// Inject valid signature from a previous block
	block2.SetExtendedHeader(*block1.ExtendedHeader())
	if bytes.Compare(*block1.ExtendedHeader(), *block2.ExtendedHeader()) != 0 {
		t.Error("Expected that the signatures on the two block would be the same")
	}

	valid = validateAuthentication(t, block2)

	if valid {
		t.Error("Expected that the injected block signature would not be considered valid.")
	}
}

func TestCanNotInjectValidSignatureIntoANewBlockWithTheSameNonce(t *testing.T) {
	dir, ks := createKeystore(t)
	defer os.RemoveAll(dir)

	account1 := createNewAccount(t, ks)
	whitelistMiner(t, account1)

	am := createAccountManager(ks)
	block1 := getNewBlock(NONCE_VALUE_1)

	if block1.ExtendedHeader() != nil {
		t.Error("Expected that the block would not have been signed yet.")
	}

	err := AuthenticateBlock(block1 , am, &account1.Address, ACCOUNT_PASSWORD)
	if err != nil {
		t.Errorf("Unable to add the authentication to the block: %v", err)
	}

	valid := validateAuthentication(t, block1)

	if !valid {
		t.Error("Expected that the block would have a *valid* signature.")
	}

	block2 := getNewBlock(NONCE_VALUE_1)

	// Inject valid signature from a previous block
	block2.SetExtendedHeader(*block1.ExtendedHeader())
	if bytes.Compare(*block1.ExtendedHeader(), *block2.ExtendedHeader()) != 0 {
		t.Error("Expected that the signatures on the two block would be the same")
	}

	valid1 := validateAuthentication(t, block2)

	if valid1 {
		t.Error("Expected that the injected block signature would not be considered valid.")
	}
}

// Functions used for setting up the test

func createKeystore(t *testing.T) (dir string, ks *keystore.KeyStore) {
	// Create a file in the current directory
	d, err := ioutil.TempDir("", "geth-keystore-test")
	if err != nil {
		t.Fatal(err)
	}
	return d, keystore.NewPlaintextKeyStore(d)
}

func createAccountManager(ks *keystore.KeyStore) (*accounts.Manager) {
	backends := []accounts.Backend{
		ks,
	}
	return accounts.NewManager(backends...)
}

func getNewBlock(nonceValue uint64) *types.Block {
	currentBlockNumber++
	if nonceValue == 0 {
		return getVanillaBlock(NONCE_VALUE_1)
	} else {
		// Create a test block to move around the database and make sure it's really new
		return getVanillaBlock(nonceValue)
	}


}

func getVanillaBlock(nonceValue uint64) *types.Block {
	return types.NewBlockWithHeader(&types.Header{
		Extra:       []byte(EXTRA_VALUE_1),
		UncleHash:   types.EmptyUncleHash,
		TxHash:      types.EmptyRootHash,
		ReceiptHash: types.EmptyRootHash,
		Number:      big.NewInt(currentBlockNumber),
		Nonce:		 types.EncodeNonce(nonceValue),
		ParentHash:  toCommonHash(currentBlockNumber-1),
	})
}

func whitelistMiner(t *testing.T, account *accounts.Account) {
	key, err := getKey(account.Address, account.URL.Path)
	if err != nil {
		t.Errorf("Unable to retrieve the key from the keystore: %v", err)
	}

	pubKey := key.PrivateKey.PublicKey

	// Add the miner / signer to the whitelist
	addMinerToWhitelist(&pubKey)
}

func validateAuthentication(t *testing.T, block *types.Block) (bool) {
	if block.ExtendedHeader() == nil {
		t.Error("Expected that the block would have a signature.")
	}

	valid, err := VerifyBlockAuthenticity(block)
	if err != nil {
		t.Errorf("Unable to validate the signature on the block: %v", err)
	}

	return valid
}

// Duplicate of the code in keystore_plain and keystore_passphrase because there seemed to be no other way to
// retrieve the key given a keystore
func getKey(addr common.Address, filename string) (*keystore.Key, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	key := new(keystore.Key)
	if err := json.NewDecoder(fd).Decode(key); err != nil {
		return nil, err
	}
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have address %x, want %x", key.Address, addr)
	}
	return key, nil
}

// Duplicate function to turn a uint64 into a common.Hash
func toCommonHash(n int64) common.Hash {
	return common.BytesToHash(crypto.Keccak256([]byte(big.NewInt(int64(n)).String())))
}*/
