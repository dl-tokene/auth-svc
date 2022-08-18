package models

import (
	"time"

	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func NewRegistrationModel(token, refreshToken string, expireTime, expireRefreshTime time.Duration, user data.User) resources.RegistrationResponse {
	response := resources.RegistrationResponse{
		Data: resources.Registration{
			Key: resources.NewKeyInt64(user.ID, resources.REGISTRATION),
			Attributes: resources.RegistrationAttributes{
				AccessToken:      token,
				RefreshToken:     refreshToken,
				ExpiresIn:        int64(expireTime.Seconds()),
				RefreshExpiresIn: int64(expireRefreshTime.Seconds()),
				TokenType:        "JWT",
				User:             NewUserModel(user),
			},
		},
	}
	return response
}
