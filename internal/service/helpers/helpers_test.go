package helpers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestVerifySignature(t *testing.T) {

	keypair, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	msg := "Test message"
	encoded_msg := crypto.Keccak256Hash([]byte(msg)).Bytes()
	signature, err := crypto.Sign(encoded_msg, keypair)
	if err != nil {
		(*t).Error(err)
	}
	fmt.Println(signature)
	got := VerifySignature(encoded_msg, hexutil.Encode(signature), crypto.PubkeyToAddress(keypair.PublicKey).String())

	if got != nil {
		(*t).Errorf("got %q, wanted nil", got)
	}
}
