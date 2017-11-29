package coterie

import (
	"os"
	"io/ioutil"
	"strings"
	"github.com/ethereum/go-ethereum/core/types"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
)

const PASSWORD_FILE_NAME string = "coinbasepwd"

var (
	ErrIncorrectDataDir	= errors.New("invalid Ethereum data directory")
	ErrMissingHash		= errors.New("unable to authenticate a block with a missing parent hash")
)

func (c *Coterie) AuthoriseBlock(header *types.Header) (error) {
	//c.lock.RLock()
	c.lock.Lock()
	signer, signFn := c.signer, c.signFn
	//c.lock.RUnlock()
	c.lock.Unlock()

	/*password, err := c.retrieveSignerUnlockingCredentials()
	if err != nil {
		return err
	}
*/
	hashToBeSigned := retrieveHashToBeSigned(header)
	if hashToBeSigned == nil || len(hashToBeSigned) == 0 {
		return ErrMissingHash
	}

	// TODO remove hacky code and replace it with the previously working code
	sig, err := signFn(accounts.Account{Address: signer}, "password123", hashToBeSigned)

	if err != nil {
		sig, err = signFn(accounts.Account{Address: signer}, "passw0rd!", hashToBeSigned)
		if err != nil {
			sig, err = signFn(accounts.Account{Address: signer}, "1234567890", hashToBeSigned)
			if err != nil {
				sig, err = signFn(accounts.Account{Address: signer}, "correcthorsebatterystaple", hashToBeSigned)
				if err != nil {
					sig, err = signFn(accounts.Account{Address: signer}, "Tr0ub4dor&3", hashToBeSigned)
					if err != nil {
						return err
					}
				}
			}
		}
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

func retrieveHashToBeSigned(header *types.Header) []byte {
	return header.ParentHash[:]
}