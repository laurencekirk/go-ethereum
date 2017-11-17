package coterie

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"math/rand"
	"github.com/ethereum/go-ethereum/log"
)

func (c *Coterie) selectedForCurrentCommittee(parentHeader *types.Header, signer common.Address) (bool, error) {

	// First check to see if the node is in the super-set of all possible members of this block's committee
	if eligible, err := c.minersWhitelist.IsMinerInWhitelist(signer); err !=  nil {
		return false, err
	} else if ! eligible {
		return false, nil
	}

	return c.hasBeenSelectedToCommittee(signer, parentHeader.ExtendedHeader.Seed)
}

func (c *Coterie) hasBeenSelectedToCommittee(signer common.Address, seed *big.Int) (bool, error) {
	// TODO implement the proper logic
	totalNumberOfNodes := 3
	log.Info("GOV: the number of nodes in the lottery is", "number", totalNumberOfNodes)
	lotteryTicketNumber := rand.Intn(totalNumberOfNodes)
	log.Info("GOV: the ticket number is", "number", lotteryTicketNumber)
	winningTicket := seed.Int64()  % int64(totalNumberOfNodes)
	log.Info("GOV: the WINNING ticket number is", "number", winningTicket)
	return winningTicket == int64(lotteryTicketNumber), nil
}