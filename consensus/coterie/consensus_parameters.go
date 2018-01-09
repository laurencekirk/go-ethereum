package coterie

import (
    "math/big"
	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/consensus/coterie/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// generate the below by running <code>govendor generate +l</code> in the root of the project.
//go:generate abigen --sol contract/ppokw_parameters.sol --pkg contract --out contract/ppokw_parameters.go
//go:generate mockgen -source=consensus_parameters.go -destination=mocks/mock_consensus_parameters.go -package=mocks

var (
	errorMissingParametersContract = errors.New("The expected the consensus parameters Smart Contract is not present")
	errorDifficultyNotAdjusted = errors.New("The difficulty has not been adjusted (since the genesis block)")
	// The address that the whitelist smart contract is deployed to in the genesis block
	parametersContractAddress = common.HexToAddress("0x0000000000000000000000000000000000000043")
)

type ConsensusParameters interface {
	GetTargetCommitteeSize() (uint, error)
	HasDifficultyBeenExternallyAdjusted(currentHeader *types.Header) (bool, error)
	GetAdjustedDifficulty() (*big.Int, error)
}

type CoterieParameters struct {
	parametersContractInstance *contract.PpokwParameters
}

func NewConsensusParameters(contractBackend bind.ContractBackend) (*CoterieParameters, error) {
	parameters, err := contract.NewPpokwParameters(parametersContractAddress, contractBackend)
	if err != nil {
		return nil, err
	}
	return &CoterieParameters{
		parameters,
	}, nil
}

func (self *CoterieParameters) GetTargetCommitteeSize() (uint, error) {
	if self.parametersContractInstance == nil {
		return 0, errorMissingWhitelistContract
	}

	if size, err := self.parametersContractInstance.CommitteeSize(&callOpts); err != nil {
		return 0, err
	} else {
		return uint(size), nil
	}
}

func (self *CoterieParameters) HasDifficultyBeenExternallyAdjusted(currentHeader *types.Header) (bool, error) {
    if self.parametersContractInstance == nil {
        return false, errorMissingWhitelistContract
    }

    if mostRecentAdjustment, err := self.parametersContractInstance.Difficulty(&callOpts); err != nil {
        return false, err
    } else {
        // Situation when no transaction has been sent to adjust the difficulty since the genesis block
        if mostRecentAdjustment.BlockNumber == nil || mostRecentAdjustment.BlockNumber.Cmp(big.NewInt(0)) == 0 {
            return false, nil
        }

        currentBlockNumber := currentHeader.Number
        // Check to see if the transaction was adjusted in this block
        if currentBlockNumber.Cmp(mostRecentAdjustment.BlockNumber) == 0 {
            return true, nil
        } else {
            return false, nil
        }
    }
}

func (self *CoterieParameters) GetAdjustedDifficulty() (*big.Int, error) {
    if self.parametersContractInstance == nil {
        return big.NewInt(-1), errorMissingWhitelistContract
    }

    if mostRecentAdjustment, err := self.parametersContractInstance.Difficulty(&callOpts); err != nil {
        return big.NewInt(-1), err
    } else if mostRecentAdjustment.BlockNumber == nil {
        return big.NewInt(-1), errorDifficultyNotAdjusted
    } else {
        return mostRecentAdjustment.Difficulty, nil
    }
}