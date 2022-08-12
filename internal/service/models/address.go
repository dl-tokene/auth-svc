package models

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func NewAddressResponseModel(address data.Address) resources.AddressResponse {
	result := resources.AddressResponse{
		Data: NewAddressModel(address),
	}
	return result
}

func NewAddressListModel(addresses []data.Address) []resources.Address {
	result := make([]resources.Address, len(addresses))
	for i, addr := range addresses {
		result[i] = NewAddressModel(addr)
	}
	return result
}

func NewAddressModel(address data.Address) resources.Address {
	result := resources.Address{
		Key: resources.Key{Type: resources.ADDRESS},
		Attributes: resources.AddressAttributes{
			Address: address.Address,
		},
	}

	return result
}

func NewCPKAddressModel(address data.Address, cpkAddress common.Address) resources.CpkAddress {
	result := resources.CpkAddress{
		Key: resources.Key{
			ID:   cpkAddress.String(),
			Type: resources.CPK_ADDRESS,
		},
		Attributes: resources.CpkAddressAttributes{
			Address:    address.Address,
			CpkAddress: cpkAddress.String(),
			UserId:     address.UserID,
		},
	}
	return result
}
