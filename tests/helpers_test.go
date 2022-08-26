package tests

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/config"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
)

func TestVerifySignature(t *testing.T) {

	keypair, _ := ecdsa.
		GenerateKey(crypto.S256(), rand.Reader)
	msg := "Test message"
	encoded_msg := crypto.Keccak256Hash([]byte(msg)).Bytes()

	//First case
	signature, err := crypto.Sign(encoded_msg, keypair)
	if err != nil {
		(*t).Error(err)
	}
	got := helpers.VerifySignature(encoded_msg, hexutil.Encode(signature), crypto.PubkeyToAddress(keypair.PublicKey).String())

	if got != nil {
		(*t).Errorf("got %q, wanted nil", got)
	}

	//Second case
	fake_keypair, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	fake_signature, err := crypto.Sign(encoded_msg, fake_keypair)
	if err != nil {
		(*t).Error(err)
	}
	got = helpers.VerifySignature(encoded_msg, hexutil.Encode(fake_signature), crypto.PubkeyToAddress(keypair.PublicKey).String())
	want := errors.New("recovered address didn't match any of the given ones")
	if got == nil {
		(*t).Errorf("got %q, wanted %q", got, want)
	}
}

func TestJWTAuthorization(t *testing.T) {

	//gen JWT
	cfg := config.New(kv.MustFromEnv())
	jwt_token, err := helpers.GenerateRefreshToken(&data.User{ID: 1}, cfg.ServiceConfig())

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
	req = req.WithContext(helpers.CtxServiceConfig(cfg.ServiceConfig())(context.Background()))

	//Test fucntion
	id, err1, err2 := helpers.Authenticate(helpers.AuthTypeSession, req)
	if err1 != nil || err2 != nil || id != 1 {
		(*t).Errorf("got %q %q, wanted nil", err1, err2)
	}
}

func TestRetrieveRefreshToken(t *testing.T) {

	cfg := config.New(kv.MustFromEnv())
	jwt_token, err := helpers.GenerateRefreshToken(&data.User{ID: 1}, cfg.ServiceConfig())

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
	req = req.WithContext(helpers.CtxServiceConfig(cfg.ServiceConfig())(context.Background()))

	//test func
	id, err := helpers.RetrieveRefreshToken(jwt_token, req)
	if id != 1 || err != nil {
		(*t).Errorf("functions works wrong")
	}
}
