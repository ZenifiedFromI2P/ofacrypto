package tools

import (
	"encoding/base64"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/pbkdf2"

	"github.com/go-humble/locstor"
)

type Keypair struct {
	CvPub  *[32]byte `json:"cvpub"`
	CvPriv *[32]byte `json:"cvpriv"`
	EdPub  []byte    `json:"edpub"`
	EdPriv []byte    `json:"edpriv"`
}

const iterations = 25e2

var Sum512 = blake2b.Sum512
var tob64 = base64.StdEncoding.EncodeToString
var fromb64 = base64.StdEncoding.DecodeString
var Store = locstor.NewDataStore(locstor.JSONEncoding)

/*
Memory-safety, etc. is just moot when it comes JS, it's so slow that I feel like
a idiot while doing crypto stuff in it, PBKDF2 with 125000 iterations is the best which we can probably get..
OWASP recommendation should be 1.04 million iterations in 2016, it hardly takes 5 minutes to do it on a modern machine, (PBKDF2 with BLAKE2b)
But, who cares, JS sucks.
*/

func wraphash() hash.Hash {
	o, err := blake2b.New512(nil)
	println("Hashed once")
	if err != nil {
		panic(err)
	}
	return o
}

func GKH(keys ...[]byte) string {
	var slice []byte
	for _, k := range keys {
		slice = append(slice, k...)
	}
	h := Sum512(slice)
	hh := h[:]
	return tob64(hh)
}

func ImportKey(password, salt []byte) {
	gsalt := blake2b.Sum512(salt) // Hash the salt to get a stronger salt
	hsalt := gsalt[:]
	Store.Save("psd", tob64(hsalt))
	// XXX: 1 is just for testing FKDFs, 125e3 is the recommended parameter for production
	s := pbkdf2.Key(password, hsalt, iterations, 32+64, wraphash)
	// TODO: Improve Ed25519 KDF
	edpub, edpriv := DeriveEd25519(s[32:])
	pub, priv := DeriveCurve25519(s[0:32])
	k := Keypair{pub, priv, edpub, edpriv}
	// We got derived key, bootstrap Curve25519
	Store.Save("ekey", k)
}
