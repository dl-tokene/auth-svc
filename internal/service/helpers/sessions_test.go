package helpers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

func TestJWTAuthorization(t *testing.T) {

	//gen JWT
	cfg := config.New(kv.MustFromEnv())
	jwt_token, err := GenerateRefreshToken(&data.User{ID: 1}, cfg.ServiceConfig())

	if err != nil || jwt_token == "" {
		(*t).Errorf("got %q", err)
	}

	//Create request
	jsonBody := []byte(fmt.Sprintf(`{"refresh_token": "%s"}`, jwt_token))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/", bodyReader)
	req.Header.Add("Authorization", "Bearer "+jwt_token)
	if err != nil {
		fmt.Println(err)
	}
	req = req.WithContext(CtxServiceConfig(cfg.ServiceConfig())(context.Background()))

	//Test fucntion
	id, _, err1, err2 := Authenticate(AuthTypeSession, req)
	if err1 != nil || err2 != nil || id != 1 {
		(*t).Errorf("got %q %q, wanted nil", err1, err2)
	}
}

func TestRetrieveRefreshToken(t *testing.T) {

	cfg := config.New(kv.MustFromEnv())
	jwt_token, err := GenerateRefreshToken(&data.User{ID: 1}, cfg.ServiceConfig())

	if err != nil || jwt_token == "" {
		(*t).Errorf("got %q", err)
	}

	//Create request
	jsonBody := []byte(fmt.Sprintf(`{"refresh_token": "%s"}`, jwt_token))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/", bodyReader)
	req.Header.Add("Authorization", "Bearer "+jwt_token)
	if err != nil {
		fmt.Println(err)
	}
	req = req.WithContext(CtxServiceConfig(cfg.ServiceConfig())(context.Background()))

	//test func
	id, err := RetrieveRefreshToken(jwt_token, req)
	if id != 1 || err != nil {
		(*t).Errorf("functions works wrong")
	}
}
