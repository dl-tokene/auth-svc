package helpers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/logan/v3/errors"
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
	got := VerifySignature(encoded_msg, hexutil.Encode(signature), crypto.PubkeyToAddress(keypair.PublicKey).String())

	if got != nil {
		(*t).Errorf("got %q, wanted nil", got)
	}

	//Second case
	fake_keypair, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	fake_signature, err := crypto.Sign(encoded_msg, fake_keypair)
	if err != nil {
		(*t).Error(err)
	}
	got = VerifySignature(encoded_msg, hexutil.Encode(fake_signature), crypto.PubkeyToAddress(keypair.PublicKey).String())
	want := errors.New("recovered address didn't match any of the given ones")
	if got == nil {
		(*t).Errorf("got %q, wanted %q", got, want)
	}
}
