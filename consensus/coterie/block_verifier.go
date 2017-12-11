package coterie

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/common"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"bytes"
)

// Useful SE answer on recovering a Public Key from a signature: https://ethereum.stackexchange.com/questions/13778/get-public-key-of-any-ethereum-account/13892

// Verify that the Block must have originated from the holder of the expected private key.
// Given the public key that is paired with the expected, unknown, private key check that Block must have been signed by
// the expected private key.
func (c *Coterie) VerifyBlockAuthenticity(parentsHeader *types.Header, currentBlockHeader *types.Header) (bool, error) {
	log.Debug("Verifying the authenticity of the block's header: ", "header", currentBlockHeader.String())

	blockAuthor, err := RetrieveBlockAuthor(parentsHeader, currentBlockHeader)
	if err != nil {
		return false, err
	}

	log.Debug("Retrieved the block signer, checking that they were selected based on the signature", "signer", blockAuthor, "signature", currentBlockHeader.ExtendedHeader.Signature)

	return c.HasBeenSelectedToCommittee(blockAuthor, &currentBlockHeader.ExtendedHeader.Signature)
}

func RetrieveBlockAuthor(parentsHeader *types.Header, currentBlockHeader *types.Header) (common.Address, error) {
	if headerErr := validateHeader(currentBlockHeader); headerErr != nil {
		return common.Address{}, headerErr
	}

	plaintext := RetrieveHashToBeSigned(parentsHeader, currentBlockHeader, ProduceBlock)
	if plaintext == nil || len(plaintext) == 0 {
		return common.Address{}, errors.New("Unable to verify a block with a missing parent hash.")
	}
	// Extract from the signature the public key that is paired with the private key; that was used to sign the block
	publicKey, err := crypto.SigToPub(plaintext, currentBlockHeader.ExtendedHeader.Signature[:])
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*publicKey), nil
}

func isSeedValid(parentsHeader *types.Header, currentBlockHeader *types.Header) (bool, error) {
	currentSeed := currentBlockHeader.ExtendedHeader.Seed
	authoriser, err := RetrieveBlockAuthor(parentsHeader, currentBlockHeader)
	if err != nil {
		return false, err
	}
	hash := retrieveSeedsHashToBeSigned(parentsHeader)

	// Extract from the signature the public key that is paired with the private key; that was used to sign the block
	publicKey, err := crypto.SigToPub(hash[:], currentSeed[:])
	if err != nil {
		return false, err
	}

	retrievedAddress := crypto.PubkeyToAddress(*publicKey)

	return bytes.Compare(authoriser[:], retrievedAddress[:]) == 0, nil
}

func validateHeader(header *types.Header) error {
	if header == nil || header.ExtendedHeader == nil || len(header.ExtendedHeader.Signature) == 0 {
		return errors.New("The Block is not correctly formatted: The Block, it's header and the extended header should not be nil")
	} else {
		return nil
	}
}