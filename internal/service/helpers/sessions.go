package helpers

import (
	"net/http"
	"strings"
	"time"

	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const AuthTypeSession = "session"
const AuthTypeYesFlow = "yes-flow"
const AuthTypeQuickbuy = "quickbuy"

type standardClaims struct {
	UserID      int64  `json:"user_id"`
	UserAddress string `json:"user_address"`
	Purpose     string `json:"purpose"`
	jwt.StandardClaims
}

type refreshTokenClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

type quickbuyClaims struct {
	RememberMeID string `json:"remember_me_id"`
	Purpose      string `json:"purpose"`
	jwt.StandardClaims
}

func GenerateJWT(address *data.Address, purpose string, cfg *config.ServiceConfig) (string, error) {
	expirationTime := time.Now().Add(cfg.TokenExpireTime)
	claims := &standardClaims{
		UserID:      address.UserID,
		UserAddress: address.Address,
		Purpose:     purpose,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.TokenKey))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT")
	}
	return tokenString, nil
}

func GenerateRefreshToken(address *data.Address, cfg *config.ServiceConfig) (string, error) {
	expirationTime := time.Now().Add(cfg.RefreshTokenExpireTime)
	claims := refreshTokenClaims{
		UserID: address.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.TokenKey))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate refresh token")
	}
	return refreshTokenString, nil
}

func parseStandardJWT(tokenString string, r *http.Request) (*standardClaims, error) {
	claims := &standardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ServiceConfig(r).TokenKey), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		return claims, errors.New("invalid token")
	}
	return claims, nil
}

// TODO(rufus): split into default Authenticate and AuthenticateYesFlow and get rid of authType parameter
//   can be easily done by converting this function into an unexported one and creating two authentication functions
//   both of which call this code with correct authType passed.

func RetrieveRefreshToken(tokenString string, r *http.Request) (int64, error) {
	tokenClaims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ServiceConfig(r).TokenKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("token is invalid")
	}
	expiresAt, ok := tokenClaims["exp"].(float64) // It was parsed to tokenClaims as float64
	if !ok {
		return 0, errors.New("can't parse expiresAt")
	}
	if int64(expiresAt) < time.Now().Unix() { //Token is expired
		return 0, errors.New("token is expired")
	}

	refreshTokenUserId, ok := tokenClaims["user_id"].(float64)
	if !ok {
		return 0, errors.New("can't parse userId")
	}
	return int64(refreshTokenUserId), nil
}
func getSessionToken(authType string, r *http.Request) (string, error) {
	if authType == AuthTypeSession || authType == AuthTypeQuickbuy {
		authHeader := r.Header.Get("Authorization")
		authHeaderSplit := strings.Split(authHeader, "Bearer ")
		if len(authHeaderSplit) != 2 {
			return "", errors.New("invalid Authorization header")
		}
		return authHeaderSplit[1], nil
	} else if authType == AuthTypeYesFlow {
		sessCookie, err := r.Cookie("flow-token")
		if err != nil {
			return "", errors.Wrap(err, "invalid flow-token cookie")
		}
		return sessCookie.Value, nil
	} else {
		panic(errors.WithStack(errors.New("unexpected authentication type received")))
	}
}
func Authenticate(authType string, r *http.Request) (int64, *jsonapi.ErrorObject, error) {
	sessToken, err := getSessionToken(authType, r)
	if err != nil {
		apierr := apierrors.Unauthorized(apierrors.CodeSessionTokenNotFound)
		err = errors.Wrap(err, "session token not found")
		return 0, apierr, err
	}
	claims, err := parseStandardJWT(sessToken, r)
	if err != nil {
		apierr := apierrors.Unauthorized(apierrors.CodeSessionTokenInvalid)
		err = errors.Wrap(err, "failed to parse jwt token")
		return 0, apierr, err
	}
	return claims.UserID, nil, nil
}
