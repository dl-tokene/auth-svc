package models

import (
	// "errors"
	// "fmt"

	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

// func NewUserListModel(users []data.User, db data.MasterQ) ([]resources.User, error) {
// 	result := make([]resources.User, len(users))
// 	for i, user := range users {
// 		// TODO(rufus): optimize requests, don't use `Get` for each user
// 		email := db.Userdata().FilterByUserID(user.ID).GetEmail()
// 		claims := db.SMTClaims().FilterByUserID(user.ID).Get()
// 		if claims == nil {
// 			return nil, errors.New(fmt.Sprintf("smt claims don't exist for user ID %d", user.ID))
// 		}
// 		result[i] = newUserModel(user, email, claims.Amount)
// 	}
// 	return result, nil
// }

func NewUserResponseModel(user data.User, email *string, claims float64) resources.UserResponse {
	result := resources.UserResponse{
		Data: newUserModel(user, email, claims),
	}
	return result
}

func newUserModel(user data.User, email *string, claims float64) resources.User {
	result := resources.User{
		Key: resources.NewKeyInt64(user.ID, resources.USER),
		Attributes: resources.UserAttributes{
			Id:        user.ID,
			Email:     email,
			SmtClaims: claims,
		},
	}
	return result
}
