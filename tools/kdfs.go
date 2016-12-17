package tools

import (
	"bytes"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/nacl/box"
)

func DeriveEd25519(keys []byte) ([]byte, []byte) {
	if len(keys) != 64 {
		println("It's recommended to use 64-bytes for deriving a Ed25519 key")
		panic(errors.New("len(keys) != 64"))
	}
	reader := bytes.NewReader(keys) // OK we got a reader
	edpub, edpriv, err := ed25519.GenerateKey(reader)
	if err != nil {
		panic(err)
	}
	return edpub, edpriv
}

func DeriveCurve25519(keys []byte) (pub, priv *[32]byte) {
	if len(keys) != 32 {
		println("It's recommended to atleast use 32-bytes for deriving a Ed25519 key")
		panic(errors.New("len(keys) != 32"))
	}
	reader := bytes.NewReader(keys)
	pub, priv, err := box.GenerateKey(reader)
	if err != nil {
		panic(err)
	}
	return
}

func GenerateNonce() *[24]byte {
	var nonce [24]byte
	n := make([]byte, 24)
	rand.Read(n)
	copy(nonce[:], n)
	return &nonce
}
