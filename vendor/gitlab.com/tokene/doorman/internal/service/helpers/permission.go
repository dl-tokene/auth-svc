package helpers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/doorman/contracts/master_access_management"
)

func CheckPermissionsByAddress(contractAddress, userAddress common.Address, client *ethclient.Client, permission, resource string) (bool, error) {
	if client == nil {
		return false, errors.New("Cant connect to node")
	}
	contract, err := master_access_management.NewMasterAccessManagement(contractAddress, client)
	if err != nil {
		return false, err
	}
	return contract.MasterAccessManagementCaller.HasPermission(&bind.CallOpts{}, userAddress, resource, permission)
}
