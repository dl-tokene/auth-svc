package helpers

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gitlab.com/tokene/doorman/internal/config"
	"gitlab.com/tokene/doorman/resources"
)

type standardClaims struct {
	Address string `json:"address"`
	Purpose string `json:"purpose"`
	jwt.StandardClaims
}
type refreshTokenClaims struct {
	Address string `json:"address"`
	jwt.StandardClaims
}

func GenerateJWT(user_claims resources.JwtClaimsAttributes, cfg *config.ServiceConfig) (string, int64, error) {
	expirationTime := time.Now().Add(cfg.TokenExpireTime)

	claims := &standardClaims{
		Address: strings.ToLower(user_claims.EthAddress),
		Purpose: user_claims.Purpose.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.TokenKey))
	if err != nil {
		return "", expirationTime.Unix(), errors.Wrap(err, "failed to generate JWT")
	}

	return tokenString, expirationTime.Unix(), nil
}

func GenerateRefreshToken(user_claims resources.JwtClaimsAttributes, cfg *config.ServiceConfig) (string, int64, error) {
	expirationTime := time.Now().Add(cfg.RefreshTokenExpireTime)

	claims := refreshTokenClaims{
		Address: strings.ToLower(user_claims.EthAddress),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.TokenKey))

	if err != nil {
		return "", expirationTime.Unix(), errors.Wrap(err, "failed to generate refresh token")
	}
	return refreshTokenString, expirationTime.Unix(), nil
}

func RetrieveTokenClaims(tokenString string, r *http.Request) (jwt.MapClaims, error) {
	tokenClaims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ServiceConfig(r).TokenKey), nil
	})
	if err != nil {
		return tokenClaims, err
	}
	if !token.Valid {
		return tokenClaims, errors.New("token is invalid")
	}

	expiresAt, ok := tokenClaims["exp"].(float64) // It was parsed to tokenClaims as float64
	if !ok {
		return tokenClaims, errors.New("can't parse expiresAt")
	}
	if int64(expiresAt) < time.Now().Unix() { //Token is expired
		return tokenClaims, errors.New("token is expired")
	}

	return tokenClaims, nil
}

// Returns address
func RetrieveToken(tokenString string, r *http.Request) (string, error) {
	tokenClaims, err := RetrieveTokenClaims(tokenString, r)
	if err != nil {
		return "", err
	}

	address, ok := tokenClaims["address"].(string)
	if !ok {
		return "", errors.New("can't parse address")
	}

	return address, nil
}

/*Returns access token's purpose, address and error*/
func RetrieveAccessToken(tokenString string, r *http.Request) (string, string, error) {
	tokenClaims, err := RetrieveTokenClaims(tokenString, r)
	if err != nil {
		return "", "", err
	}

	address, err := RetrieveToken(tokenString, r)
	if err != nil {
		return "", "", errors.New("can't parse address")
	}

	purpose, ok := tokenClaims["purpose"].(string)
	if !ok {
		return "", "", errors.New("cannot parse purpose")
	}

	return purpose, address, nil
}
func Authenticate(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderSplit := strings.Split(authHeader, "Bearer ")

	if len(authHeaderSplit) != 2 {
		return "", errors.New("invalid Authorization header")
	}

	return authHeaderSplit[1], nil
}

func GetAccessTokenInfo(r *http.Request) (string, string, error) {
	token, err := Authenticate(r)
	if err != nil {
		return "", "", err
	}

	purpose, address, _ := RetrieveAccessToken(token, r)

	return purpose, address, err

}

func GetTokenInfo(r *http.Request) (string, error) {
	token, err := Authenticate(r)
	if err != nil {
		return "", err
	}

	address, err := RetrieveToken(token, r)
	return address, err
}
