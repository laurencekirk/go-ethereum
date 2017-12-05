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
)

const PASSWORD_FILE_NAME string = "coinbasepwd"

var (
	ErrIncorrectDataDir	= errors.New("invalid Ethereum data directory")
	ErrMissingHash		= errors.New("unable to authenticate a block with a missing parent hash")
)

func (c *Coterie) AuthoriseBlock(parentHeader *types.Header, header *types.Header) (error) {
	log.Debug("GOV: the header", "header", header)

	//c.lock.RLock()
	c.lock.Lock()
	signer, signFn := c.signer, c.signFn
	//c.lock.RUnlock()
	c.lock.Unlock()

	hashToBeSigned := retrieveHashToBeSigned(parentHeader, header, BlockProducer)
	if hashToBeSigned == nil || len(hashToBeSigned) == 0 {
		return ErrMissingHash
	}

	signingAccount := accounts.Account{Address: signer}
	/*password, err := c.retrieveSignerUnlockingCredentials()
	if err != nil {
		return err
	}
	defer zeroPassword(&password)

	if err:= c.ks.Unlock(signingAccount , password); err != nil {
		return err
	}
	defer c.ks.Lock(signer)*/

	sig, err := signFn(signingAccount, hashToBeSigned)
	if err != nil {
		return err
	}

	log.Debug("GOV: the signature", "signature", sig)

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
func retrieveHashToBeSigned(parentHeader *types.Header, header *types.Header, task ConsensusTask) []byte {
	switch task {
		case BlockProducer:
			seed := parentHeader.ExtendedHeader.Seed
			round := header.Number
			taskBytes := big.NewInt(int64(task)).Bytes()
			data := append(round.Bytes()[:], taskBytes[:]...)
			data = append(data[:], seed.Bytes()[:]...)
			hash := crypto.Keccak256Hash(data)
			return hash[:]
		default:
			log.Warn("GOV: retrieving of the hash for the task has not been implemented yet", "task", task)
			return nil
	}
}

func zeroPassword(p *string) {
	*p = ""
}