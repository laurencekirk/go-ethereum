package coterie

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math"
	"strings"
)

type ConsensusTask int

const (
	ProduceBlock     ConsensusTask = iota + 1
	DetectAnomalies
	ChangeParameters
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

	threshold, err := calculateWinningThreshold(c.consensusParameters, c.minersWhitelist)
	if err != nil {
		return false, err
	}
	if sigValue, err := calculateSignaturesRealValue(authorisation); err != nil {
		return false, err
	} else {
		log.Debug("GOV: The real value based on the signature is ", "value", sigValue)
		log.Debug("GOV: The threshold to be below is ", "threshold", threshold)
		return sigValue < threshold, nil
	}
}

func calculateSignaturesRealValue(authorisation *types.Signature) (float64, error) {
	uint64SizedSlice := authorisation[: maxNumBytes]
	sigString := common.ToHex(uint64SizedSlice)
	sigString, err := removeLeadingZeroDigits(sigString)
	if err != nil {
		return -1, err
	}
	value, err := hexutil.DecodeUint64(sigString)
	if err != nil {
		return -1, err
	}
	asFloat := float64(value)
	divisor := math.Pow(2, exponent)
	return asFloat / divisor, nil
}

func calculateWinningThreshold(contractParameters *ConsensusParameters, whitelist *AuthorisedMinersWhitelist) (float64, error) {
	targetCommitteeSize, err := contractParameters.GetTargetCommitteeSize()
	if err != nil {
		return float64(-1), err
	}

	log.Debug("GOV: the target committee is", "number", targetCommitteeSize)

	whitelistSize, err := whitelist.GetWhitelistSize()
	if err != nil {
		return float64(-1), err
	}

	log.Debug("GOV: the number of nodes in the lottery is", "number", whitelistSize)
	targetCommitteeSizeAsFloat := float64(targetCommitteeSize)
	whitelistSizeAsFloat := float64(whitelistSize)

	return targetCommitteeSizeAsFloat / whitelistSizeAsFloat, nil
}

func removeLeadingZeroDigits(hexString string) (string, error) {
	if _, err := hexutil.DecodeUint64(hexString); err != nil {
		if err == hexutil.ErrLeadingZero {
			without0xPrefix := hexString[2:]
			numLeaderZeros := 0
			for _, char := range without0xPrefix {
				if char == '0' {
					numLeaderZeros++
				} else {
					break
				}
			}

			// Catch the, very unlikely, edge case when the signature results in >= 16 leading 0s
			if numLeaderZeros == len(without0xPrefix) {
				return "0x0", nil
			} else {
				return strings.Join([]string{"0x", without0xPrefix[numLeaderZeros:]}, ""), nil
			}
		} else {
			return hexString, err
		}
	} else {
		return hexString, nil
	}

}