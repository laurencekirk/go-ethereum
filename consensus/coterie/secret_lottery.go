package coterie

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math"
)

type ConsensusTask int

const (
	BlockProducer ConsensusTask	= iota + 1
	AnomalyDetection
	ParameterChanges
)

const (
	maxNumBytes = 8
	exponent = maxNumBytes * 2 * 4
)

func (c *Coterie) HasBeenSelectedToCommittee(signer common.Address, authorisation *types.Signature) (bool, error) {
	log.Debug("GOV: checking if the account has been selected as a block producer", "address", signer, "signature", authorisation)

	// Check if the miner is in the current whitelist
	if eligible, err := c.minersWhitelist.IsMinerInWhitelist(signer); err !=  nil {
		return false, err
	} else if ! eligible {
		return false, nil
	}

	log.Debug("GOV: the miner is in the whitelist...")

	return c.hasWonSecretLottery(authorisation)
}

func (c *Coterie) hasWonSecretLottery(authorisation *types.Signature) (bool, error) {

	if threshold, err := calculateWinningThreshold(c.consensusParameters, c.minersWhitelist); err != nil {
		return false, err
	} else {
		return calculateSignaturesRealValue(authorisation) < threshold, nil
	}
}

func calculateSignaturesRealValue(authorisation *types.Signature) float64 {
	uint64SizedSlice := authorisation[: maxNumBytes]
	sigString := common.ToHex(uint64SizedSlice)
	value, err := hexutil.DecodeUint64(sigString)
	if err != nil {
		log.Error("The error", "err", err)
	}
	log.Debug("The value", "value", value)
	asFloat := float64(value)
	devisor := math.Pow(2, exponent)
	return asFloat / devisor
}

func calculateWinningThreshold(contractParameters *ConsensusParameters, whitelist *AuthorisedMinersWhitelist) (float64, error) {
	targetCommitteeSize, err := contractParameters.GetTargetCommitteeSize()
	if err != nil {
		return float64(-1), err
	}

	log.Info("GOV: the target committee is", "number", targetCommitteeSize)

	whitelistSize, err := whitelist.GetWhitelistSize()
	if err != nil {
		return float64(-1), err
	}

	log.Info("GOV: the number of nodes in the lottery is", "number", whitelistSize)

	return float64(targetCommitteeSize / whitelistSize), nil
}