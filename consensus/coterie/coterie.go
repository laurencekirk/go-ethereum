// Package clique implements the coterie consensus engine - implementing the consensus interface.
package coterie

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"sync"
)

var (
	errInvalidBlock = errors.New("invalid block requested for sealing")
	ErrInvalidAuth	= errors.New("invalid miner authentication signature in the header")
	ErrInvalidSeed	= errors.New("invalid seed value in the header")
)

// DirectoryLocatorFn is a callback function that is used to retrieve the location of the data dir of the Ethereum node
// This is used when locating the password file to unlock the account; as part of sealing a block.
type DirectoryLocatorFn func() string

// SignerFn is a signer callback function to request a hash to be signed by a
// backing account. Copied from the clique implementation.
//type SignerFn func(account accounts.Account, passphrase string, hash []byte) ([]byte, error)
type SignerFn func(account accounts.Account, hash []byte) ([]byte, error)


type Coterie struct {
	signer 						common.Address      // Ethereum address of the signing key
	signFn 						SignerFn            // Signer function to authorize hashes with
	ks 							*keystore.KeyStore			// Keystore which stores the Ethereum accounts
	dirLocFun					DirectoryLocatorFn  		// Data directory location function
	minersWhitelist				AuthorisedMinersWhitelist  	// Whitelist of miners governed by a smart contract
	consensusParameters			ConsensusParameters			// Smart contract which controls some of the consensus parameters

	secondLayerConsensusEngine	consensus.Engine 	// Another consensus engine e.g. Ethash PoW engine that this consensus engine 'extends' and can call into.

	lock sync.Mutex // Ensures thread safety for the in-memory caches and mining fields
}

func New(deferTo consensus.Engine, dirLocFn DirectoryLocatorFn) *Coterie {
	return &Coterie{
		dirLocFun: dirLocFn,
		secondLayerConsensusEngine:	deferTo,
	}
}

// Consensus - Engine - interface implementation

// Author retrieves the Ethereum address of the account that minted the given
// block, which may be different from the header's coinbase if a consensus
// engine is based on signatures.
// @Deprecated: there isn't enough information in the header alone to determine the author in this consensus mechanism -
// use the function `AuthorisedBy` that uses chainReader and is consistent with the other consensus mechanism functions.
func (c *Coterie) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (c *Coterie) AuthorisedBy(chain consensus.ChainReader, header *types.Header) (common.Address, error) {
	parentHeader := GetParentBlockHeader(chain, header)
	return RetrieveBlockAuthor(parentHeader, header)
}

// VerifyHeader checks whether a header conforms to the consensus rules of a
// given engine. Verifying the seal may be done optionally here, or explicitly
// via the VerifySeal method.
func (c *Coterie) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	return c.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications (the order is that of
// the input slice).
func (c *Coterie) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {

	return c.secondLayerConsensusEngine.VerifyHeaders(chain, headers, seals)
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *Coterie) verifyHeader(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// TODO replace with proper validation (determine how much should be copied from Ethash)
	// TODO check the seed value is correct
	log.Debug("GOV: verifying the block's header")


	return nil
}

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of a given engine.
func (c *Coterie) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	// Defer to the PoW verification
	return c.secondLayerConsensusEngine.VerifyUncles(chain, block)
}

// VerifySeal checks whether the crypto seal on a header is valid according to
// the consensus rules of the given engine.
func (c *Coterie) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	parentHeader := GetParentBlockHeader(chain, header)

	if valid, err := c.VerifyBlockAuthenticity(parentHeader, header); err != nil {
		return err
	} else if ! valid {
		return ErrInvalidAuth
	}

	if valid, err := isSeedValid(parentHeader, header); err != nil {
		return err
	} else if ! valid {
		return ErrInvalidSeed
	}

	return c.secondLayerConsensusEngine.VerifySeal(chain, header)
}

// Prepare initializes the consensus fields of a block header according to the
// rules of a particular engine. The changes are executed inline.
func (c *Coterie) Prepare(chain consensus.ChainReader, header *types.Header) error {

	// Covers the edge case where the node is starting up
	if c.consensusParameters != nil {
		if hasBeenAdjustedThisBlock, err := c.consensusParameters.HasDifficultyBeenExternallyAdjusted(header); err != nil {
			return err
		} else if hasBeenAdjustedThisBlock {
			if difficulty, err := c.consensusParameters.GetAdjustedDifficulty(); err != nil {
				return err
			} else {
				header.Difficulty = difficulty;
				return nil
			}
		}
	}

	// If it has not been adjusted then defer to the regular difficulty algorithm - of basing it off of the parent header
	return c.secondLayerConsensusEngine.Prepare(chain, header)
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
// Note: The block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (c *Coterie) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if header.ExtendedHeader == nil {
		header.ExtendedHeader = &types.ExtendedHeader{}
	}
	// TODO implement proper logic
	return c.secondLayerConsensusEngine.Finalize(chain, header, state, txs, uncles, receipts)
}

// Seal generates a new block for the given input block with the local miner's
// seal place on top.
func (c *Coterie) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	// c.lock.RLock()

	log.Debug("GOV: seal start")
	currentBlockHeader := block.Header()

	// Sealing the genesis block is not supported
	number := currentBlockHeader.Number.Uint64()
	if number == 0 {
		return nil, errInvalidBlock
	}

	// Hold the signer fields for the entire sealing procedure
	c.lock.Lock()
	signer := c.signer
	c.lock.Unlock()

	// We'll need information from the parent header (such as the seed)
	parentBlockHeader := GetParentBlockHeader(chain, currentBlockHeader)

	if err := c.UnlockAccount(signer); err != nil {
		return nil, err
	}
	defer c.ks.Lock(signer)

	// Create the signature which will be used as part of the secret lottery and to authorise the block
	if err := c.AuthoriseBlock(parentBlockHeader, currentBlockHeader); err != nil {
		log.Debug("GOV: An error occurred whilst creating the signature to authorise the block")
		return nil, err
	}

	// First check to see if the node is part of the current coterie / block-creator set
	inCommittee, err := c.HasBeenSelectedToCommittee(signer, &currentBlockHeader.ExtendedHeader.Authorisation)
	if err != nil {
		return nil, err
	} else if ! inCommittee {
		// TODO clique returns an error here - look into the consequences of returning a custom error instead of nil
		return nil, nil
	}

	sealedBlock, err := c.seal(block)
	if err != nil {
		return nil, err
	}

	blockWithSeed, err := c.addNextSeed(parentBlockHeader, sealedBlock)
	if err != nil {
		return nil, err
	}

	return c.secondLayerConsensusEngine.Seal(chain, blockWithSeed, stop)
}

func GetParentBlockHeader(chain consensus.ChainReader, currentBlockHeader *types.Header) (*types.Header) {
	childBlockNumber := currentBlockHeader.Number.Uint64()
	return chain.GetHeader(currentBlockHeader.ParentHash, childBlockNumber-1)
}

// APIs returns the RPC APIs this consensus engine provides.
func (c *Coterie) APIs(chain consensus.ChainReader) []rpc.API {
	// TODO implement proper logic
	return nil
}

// Coterie-specific functions / methods

//TODO figure out if this needs to be changed once the hash for the committee is known
// mine is the actual proof-of-work miner that searches for a nonce starting from
// seed that results in correct final block difficulty.
func (c *Coterie) seal(block *types.Block, ) (*types.Block, error) {
	header := block.Header()

	// Seal and return a block (if still needed)
	return block.WithSeal(header), nil
}

func (c *Coterie) addNextSeed(parentBlockHeader *types.Header, block *types.Block) (*types.Block, error) {
	header := block.Header()

	nextSeed, err := c.GenerateNextSeed(parentBlockHeader)
	if err != nil {
		return block, err
	}

	copy(header.ExtendedHeader.Seed[:], nextSeed[:])

	return block.WithSeal(header), nil
}

// Instantiates the miners whitelist and makes the consensus engine aware of it
func (c *Coterie) SetAuthorisedMinersWhitelist(contractBackend bind.ContractBackend) (error) {
	if c.minersWhitelist == nil {
		if whitelist, err := NewAuthorisedMinersWhitelist(contractBackend); err != nil {
			return err
		} else {
			c.minersWhitelist = whitelist
		}
	}

	if c.consensusParameters == nil {
		if parameters, err := NewConsensusParameters(contractBackend); err != nil {
			return err
		} else {
			c.consensusParameters = parameters
		}
	}

	return nil
}

// Authorize injects a private key into the consensus engine to mint new blocks with.
func (c *Coterie) Authorize(signer common.Address, signFn SignerFn, ks *keystore.KeyStore) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
	c.ks = ks
}