package models

import (
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func NewUserResponseModel(user data.User) resources.UserResponse {
	result := resources.UserResponse{
		Data: NewUserModel(user),
	}
	return result
}

func NewUserListModel(addresses []data.User) []resources.User {
	result := make([]resources.User, len(addresses))
	for i, addr := range addresses {
		result[i] = NewUserModel(addr)
	}
	return result
}

func NewUserModel(user data.User) resources.User {
	result := resources.User{
		Key: resources.NewKeyInt64(user.ID, resources.USER),
		Attributes: resources.UserAttributes{
			Address: user.Address,
		},
	}

	return result
}
