// Package clique implements the coterie consensus engine - implementing the consensus interface.
package coterie

import (
	"errors"
	//"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	/*"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethdb"*/
)

var (
	errInvalidBlock = errors.New("invalid block requested for sealing")
	ErrInvalidAuth	= errors.New("invalid miner authentication signature in the header")
)

// DirectoryLocatorFn is a callback function that is used to retrieve the location of the data dir of the Ethereum node
// This is used when locating the password file to unlock the account; as part of sealing a block.
type DirectoryLocatorFn func() string

// SignerFn is a signer callback function to request a hash to be signed by a
// backing account. Copied from the clique implementation.
type SignerFn func(account accounts.Account, passphrase string, hash []byte) ([]byte, error)

/*
type Coterie struct {
	db     			ethdb.Database           // Database to store and retrieve x
	minersWhitelist	*AuthorisedMinersWhitelist // Whitelist of miners governed by a smart contract
	signer 			common.Address           // Ethereum address of the signing key
	signFn 			SignerFn                 // Signer function to authorize hashes with
	dirLocFun		DirectoryLocatorFn        // Data directory location function
	lock   			sync.RWMutex               // Protects the coterie fields
}
*/

/*
func New(dirLocFn DirectoryLocatorFn, db ethdb.Database) *Coterie {
	return &Coterie{
		dirLocFun: dirLocFn,
		db:	db,
	}
}

// Consensus - Engine - interface implementation
*/
// Author retrieves the Ethereum address of the account that minted the given
// block, which may be different from the header's coinbase if a consensus
// engine is based on signatures.
func (c *Coterie) Author(header *types.Header) (common.Address, error) {
	return RetrieveBlockAuthor(header)
}
/*
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
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			err := c.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
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
*/
// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of a given engine.
func (c *Coterie) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	// Same as the Clique consensus - we don't expect there to be any uncles
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}
/*
// VerifySeal checks whether the crypto seal on a header is valid according to
// the consensus rules of the given engine.
func (c *Coterie) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	// TODO replace with proper validation

	// TODO confirm that this is the correct location for this logic (it is added when we create a seal, so verification here currently seems to make sense)
	// Retrieve the public key of the miner that created the block and verify that they are eligible to create the block (that there are in the whitelist)
	if valid, err := VerifyBlockAuthenticity(c.minersWhitelist, header); err != nil {
		return err
	} else if !valid {
		return ErrInvalidAuth
	}

	return nil
}

// Prepare initializes the consensus fields of a block header according to the
// rules of a particular engine. The changes are executed inline.
func (c *Coterie) Prepare(chain consensus.ChainReader, header *types.Header) error {
	chainsParent := chain.CurrentHeader().ParentHash
	log.Debug("GOV: chain's parent hash", "hash", chainsParent)
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	log.Debug("GOV: the parent block header", "blockHeader", parent)
	if chainsParent != parent.ParentHash {
		log.Debug("GOV: mismatching headers")
		return errors.New("parent hash mismatch on headers")
	}
	// TODO implement proper logic
	return nil
}

// Finalize runs any post-transaction state modifications (e.g. block rewards)
// and assembles the final block.
// Note: The block header and state database might be updated to reflect any
// consensus rules that happen at finalization (e.g. block rewards).
func (c *Coterie) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// TODO implement proper logic
	return types.NewBlock(header, txs, nil, receipts), nil
}*/

// Seal generates a new block for the given input block with the local miner's
// seal place on top.
func (c *Coterie) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {
	// c.lock.RLock()

	log.Debug("GOV: seal start")
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return nil, errInvalidBlock
	}

	// TODO implement proper logic

	// Hold the signer fields for the entire sealing procedure
	c.lock.Lock()
	signer := c.signer
	c.lock.Unlock()

	// First check to see if the node is part of the current coterie / block-creator set
	partOfCoterie, err := partOfCurrentCoterie(header, signer)
	if err != nil {
		return nil, err
	} else if ! partOfCoterie {
		// TODO clique returns an error here - look into the consequences of returning a custom error instead of nil
		return nil, nil
	}

	// Create a runner and the multiple search threads it directs
	abort := make(chan struct{})
	found := make(chan *types.Block)

	go c.seal(block, abort, found)
	// Wait until sealing is terminated or a nonce is found
	var result *types.Block
	select {
		case <-stop:
			// Outside abort, stop all miner threads
			close(abort)
		case result = <-found:
			// One of the threads found a block, abort all others
			close(abort)
	}

	return result, nil
}

// mine is the actual proof-of-work miner that searches for a nonce starting from
// seed that results in correct final block difficulty.
func (c *Coterie) seal(block *types.Block, abort chan struct{}, found chan *types.Block) {
	// Extract some data from the header
	var (
		header = block.Header()
	)
	// TODO implement the logic for the rounds here
	log.Debug("GOV: before sealing to the header", "block", block.String())

	// Add the authorisation signature to the block
	if err := c.AuthoriseBlock(header); err != nil {
		return
	}

	// Seal and return a block (if still needed)
	select {
		case found <- block.WithSeal(header):
			log.Debug("Block sealed", "sealed block", found)
		case <-abort:
			log.Debug("Block sealed, but discarded")
	}
	return
}

// APIs returns the RPC APIs this consensus engine provides.
func (c *Coterie) APIs(chain consensus.ChainReader) []rpc.API {
	// TODO implement proper logic
	return nil
}

// Coterie-specific functions / methods

// Instantiates the miners whitelist and makes the consensus engine aware of it
func (c *Coterie) SetAuthorisedMinersWhitelist(contractBackend bind.ContractBackend) (error) {
	if c.minersWhitelist == nil {
		if whitelist, err := NewAuthorisedMinersWhitelist(contractBackend); err != nil {
			return err
		} else {
			c.minersWhitelist = whitelist
		}
	}
	return nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (c *Coterie) Authorize(signer common.Address, signFn SignerFn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.signer = signer
	c.signFn = signFn
}