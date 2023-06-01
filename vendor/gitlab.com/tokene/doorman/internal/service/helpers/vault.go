package helpers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokene/doorman/contracts/master_contracts_registry"
	"net/http"
)

func GetAddressAccManagement(r *http.Request) (common.Address, error) {
	address := RegistryConfig(r).Address

	accessM, err := master_contracts_registry.NewMasterContractsRegistry(address, EthRPCConfig(r).EthClient())
	if err != nil {
		return common.Address{}, err
	}
	return accessM.GetMasterAccessManagement(&bind.CallOpts{})
}
