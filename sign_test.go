package flowpanic_test

import (
	"testing"

	"github.com/onflow/flow-go-sdk/crypto"
)

func TestSign(t *testing.T) {
	privateKey, err := crypto.GeneratePrivateKey(crypto.ECDSA_P256, make([]byte, 32))
	if err != nil {
		t.Fatal(err)
	}
	// privateKey.PublicKey() // necessary on go 1.24 to avoid panic

	signer, err := crypto.NewInMemorySigner(privateKey, crypto.SHA3_256)
	if err != nil {
		t.Fatal(err)
	}
	_, err = signer.Sign([]byte("some message"))
	if err != nil {
		t.Fatal(err)
	}
}
