package coterie

import (
	"os"
	"io/ioutil"
	"strings"
	"github.com/ethereum/go-ethereum/core/types"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

const PASSWORD_FILE_NAME string = "coinbasepwd"

var (
	ErrIncorrectDataDir	= errors.New("invalid Ethereum data directory")
	ErrMissingHash		= errors.New("unable to authenticate a block with a missing parent hash")
)

func (c *Coterie) UnlockAccount(signer common.Address) error {
	password, err := c.retrieveSignerUnlockingCredentials()
	if err != nil {
		return err
	}
	defer zeroPassword(&password)

	signingAccount := accounts.Account{Address: signer}
	if err:= c.ks.Unlock(signingAccount , password); err != nil {
		return err
	}

	return nil
}

/*
 *
 * Note expects the account to have been locked before this method is called e.g. using UnlockAccount - this is an optimisation for devices with smaller amounts of RAM
 */
func (c *Coterie) AuthoriseBlock(parentHeader *types.Header, header *types.Header) (error) {
	c.lock.Lock()
	signer, signFn := c.signer, c.signFn
	c.lock.Unlock()

	hashToBeSigned := RetrieveHashToBeSigned(parentHeader, header, ProduceBlock)
	if hashToBeSigned == nil || len(hashToBeSigned) == 0 {
		return ErrMissingHash
	}

	signingAccount := accounts.Account{Address: signer}

	sig, err := signFn(signingAccount, hashToBeSigned)
	if err != nil {
		return err
	}

	header.SetExtendedHeader(sig)
	return nil
}

func (c *Coterie) retrieveSignerUnlockingCredentials() (string, error) {
	// c.lock.RLock()
	c.lock.Lock()
	dirLocFn := c.dirLocFun
	//c.lock.RUnlock()
	c.lock.Unlock()

	dirLoc := dirLocFn()

	if dirLoc == "" {
		return "", ErrIncorrectDataDir
	}
	pwdFilePath := dirLoc + string(os.PathSeparator) + PASSWORD_FILE_NAME

	log.Debug("GOV: Path to the password file determined as ", "path", pwdFilePath)

	if coinbasePwd, err := readPasswordFromFile(pwdFilePath); err != nil {
		return "", err
	} else {
		return coinbasePwd, nil
	}
}

func readPasswordFromFile(filePath string) (string, error) {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Only expect that there will be one line / password in the file
	lines := strings.Split(string(text), "\n")
	return strings.TrimRight(lines[0], "\r"), nil
}

/*
 * We require a 32 bit length string for the ECDSA signing function - this function assembles the known parts of a block, for a given 'task',
 * that will be used to create the fixed length string / hash.
 */
func RetrieveHashToBeSigned(parentHeader *types.Header, header *types.Header, task ConsensusTask) []byte {
	switch task {
		case ProduceBlock:
			seed := parentHeader.ExtendedHeader.Seed
			round := header.Number
			taskBytes := big.NewInt(int64(task)).Bytes()
			taskBytesLength := len(taskBytes)
			roundBytes := round.Bytes()
			roundBytesLen := len(roundBytes)
			totalDataLength := roundBytesLen + taskBytesLength + len(seed)

			// Create a byte array based off of the necessary data
			data := make([]byte, totalDataLength)
			copy(data[0:roundBytesLen], roundBytes)
			copy(data[roundBytesLen: roundBytesLen + taskBytesLength], taskBytes)
			copy(data[roundBytesLen + taskBytesLength:], seed[:])

			hash := crypto.Keccak256Hash(data)
			return hash[:]
		default:
			log.Warn("GOV: retrieving of the hash for the task has not been implemented yet", "task", task)
			return nil
	}
}

/*
 *
 * Note expects the account to have been locked before this method is called e.g. using UnlockAccount - this is an optimisation for devices with smaller amounts of RAM
 */
func (c *Coterie) GenerateNextSeed(parentHeader *types.Header) (*types.Signature, error) {
	// The signing function requires a 32 byte 'hash' to be signed and the signature is 65 bytes, so take a Keccak256Hash of it
	hashToBeSigned := retrieveSeedsHashToBeSigned(parentHeader)

	c.lock.Lock()
	signer, signFn := c.signer, c.signFn
	c.lock.Unlock()

	signingAccount := accounts.Account{Address: signer}

	sig, err := signFn(signingAccount, hashToBeSigned[:])
	if err != nil {
		return nil, err
	}

	return types.BytesToSignature(sig), nil
}

func retrieveSeedsHashToBeSigned(parentHeader *types.Header) common.Hash {
	// The new seed Q r is computed as the signature of the previous seed Q r−1
	previousSeed := parentHeader.ExtendedHeader.Seed

	return crypto.Keccak256Hash(previousSeed[:])
}

func zeroPassword(p *string) {
	*p = ""
}