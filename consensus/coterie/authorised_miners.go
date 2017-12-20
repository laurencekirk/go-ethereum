package coterie

import (
	"crypto/ecdsa"
	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/consensus/coterie/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// generate the below by running <code>govendor generate +l</code> in the root of the project.
//go:generate abigen --sol contract/authorised_miners_whitelist.sol --pkg contract --out contract/authorised_miners_whitelist.go
//go:generate mockgen -source=authorised_miners.go -destination=mocks/mock_authorised_miners.go -package=mocks

var (
	errorMissingWhitelistContract = errors.New("The expected AuthorisedMinersWhitelist Smart Contract is not present")
	callOpts = bind.CallOpts{
		Pending: false,
	}
	// The address that the whitelist smart contract is deployed to in the genesis block
	whitelistContractAddress = common.HexToAddress("0x0000000000000000000000000000000000000042")
)

type AuthorisedMinersWhitelist interface {
	IsMinerInWhitelist(minerAddress common.Address) (bool, error)

	GetWhitelistSize() (uint, error)

	AddMinerToWhitelist(minerAddress common.Address, msgSender *ecdsa.PrivateKey) (*types.Transaction, error)
}

type CoterieWhitelist struct {
	whitelistContractInstance *contract.AuthorisedMinersWhitelist
}

func NewAuthorisedMinersWhitelist(contractBackend bind.ContractBackend) (*CoterieWhitelist, error) {
	whitelist, err := contract.NewAuthorisedMinersWhitelist(whitelistContractAddress, contractBackend)
	if err != nil {
		return nil, err
	}
	return &CoterieWhitelist{
		whitelist,
	}, nil
}

// 0x9e5e939fb0a23529934c061d6ecf4c93e7893d4e
// 0xea30250dd7263a4783c66463c236a2153d6b88b4
// 0x46dfb921f8f7edbbd8100458b7c1beefeabf6e15
// 0x6c80e492308f051eba48d03bcc04625682ae3e07
// 0x30ff130a7d11ef9d1efbdf19d5309556acd129cf

func (self *CoterieWhitelist) IsMinerInWhitelist(minerAddress common.Address) (bool, error) {
	if self.whitelistContractInstance == nil {
		return false, errorMissingWhitelistContract
	}

	if len(minerAddress) == 0 {
		return false, nil
	}

	if auth, err := self.whitelistContractInstance.IsAuthorisedMiner(&callOpts, minerAddress); err != nil {
		return false, err
	} else {
		return auth, nil
	}
}

func (self *CoterieWhitelist) GetWhitelistSize() (uint, error) {
	if self.whitelistContractInstance == nil {
		return 0, errorMissingWhitelistContract
	}

	if size, err := self.whitelistContractInstance.Size(&callOpts); err != nil {
		return 0, err
	} else {
		return uint(size), nil
	}
}

func (self *CoterieWhitelist) AddMinerToWhitelist(minerAddress common.Address, msgSender *ecdsa.PrivateKey) (*types.Transaction, error) {
	if self.whitelistContractInstance == nil {
		return nil, errorMissingWhitelistContract
	}

	if msgSender == nil {
		return nil, errors.New("A private key to authorise adding the miner must be provided")
	}

	if len(minerAddress) == 0 {
		return nil, errors.New("Invalid address to add to the whitelist")
	}

	return self.whitelistContractInstance.AuthoriseMiner(bind.NewKeyedTransactor(msgSender), minerAddress)
}
