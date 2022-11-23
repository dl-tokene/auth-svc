package gosdk

import (
	"github.com/ethereum/go-ethereum/common"
)

type NodeAdminsI interface {
	GetAdminsList() map[common.Address]Admin
	CheckAdmin(address common.Address) bool
}
type NodeAdminsMock struct {
	Admins map[common.Address]Admin
}
type Admin struct{}

func NewNodeAdminsMock(addresses ...common.Address) NodeAdminsI {
	adminsMap := make(map[common.Address]Admin)
	for _, admin := range addresses {
		adminsMap[admin] = Admin{}
	}
	return &NodeAdminsMock{
		Admins: adminsMap,
	}

}
func (node *NodeAdminsMock) CheckAdmin(address common.Address) bool {
	_, check := node.Admins[address]
	return check
}

func (node *NodeAdminsMock) GetAdminsList() map[common.Address]Admin {
	return node.Admins
}
