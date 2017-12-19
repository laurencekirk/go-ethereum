package coterie

import (
	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/consensus/coterie/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// generate the below by running <code>govendor generate +l</code> in the root of the project.
//go:generate abigen --sol contract/ppokw_parameters.sol --pkg contract --out contract/ppokw_parameters.go

var (
	errorMissingParametersContract = errors.New("The expected the consensus parameters Smart Contract is not present")
	// The address that the whitelist smart contract is deployed to in the genesis block
	parametersContractAddress = common.HexToAddress("0x0000000000000000000000000000000000000043")
)

type ConsensusParameters interface {
	GetTargetCommitteeSize() (uint, error)
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