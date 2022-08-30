package models

import (
	"time"

	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func NewAdminLoginModel(token, refreshToken string, expireTime, refreshExpireTime time.Duration) resources.LoginResponse {
	response := resources.LoginResponse{
		Data: resources.Login{
			Key: resources.Key{Type: resources.ADMIN_LOGIN},
			Attributes: resources.LoginAttributes{
				AccessToken:      token,
				RefreshToken:     refreshToken,
				ExpiresIn:        int64(expireTime.Seconds()),
				RefreshExpiresIn: int64(refreshExpireTime.Seconds()),
				TokenType:        "JWT",
			},
		},
	}
	return response
}
