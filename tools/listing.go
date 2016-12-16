package tools

import (
	"encoding/json"
	"golang.org/x/crypto/ed25519"
)

type Listing struct {
	Deltoken string `json:"deltoken"`
	Dealer   Dealer `json:"dealer"`
	Block    string `json:"block"`
	Sig      []byte `json:"sig"`
}

type RealListing struct {
	Name    string
	Comment string
	Payto   string
	Contact string
	Images  []Image
}

type Image struct {
	MIME  string
	Block []byte
}

func NewListing(name, comment, payto, deltoken, contact string) bool {
	l := RealListing{
		Comment: comment,
		Name:    name,
		Payto:   payto,
		Contact: contact,
		Images:  nil,
	}
	rljson, err := json.Marshal(l)
	if err != nil {
		println(err.Error())
		panic(err)
	}
	var k Keypair
	Store.Find("ekey", &k)
	sig := ed25519.Sign(k.EdPriv, rljson)
	list := Listing{Deltoken: deltoken, Dealer: selfDealer, Block: string(rljson), Sig: sig}
	ljson, err := json.Marshal(list)
	if err != nil {
		println(err.Error())
		panic(err)
	}
	println(string(ljson))
	SendToServer("/api/new/listing", ljson)
	return true
}

type ExtendedListing struct {
	Block  string `json:"block"`
	Sig    []byte `json:"sig"`
	Dealer Dealer `json:"dealer"`
}
type Listings []ExtendedListing
type SListing struct {
	R   RealListing
	D   Dealer
	KH  string // Human-readable KeyHash string
	Sig string
}
type SListings []SListing

func FListing(lsj string) SListings {
	var hits SListings
	lsjson := []byte(lsj)
	var ls Listings
	json.Unmarshal(lsjson, &ls)
	for _, l := range ls {
		blkbytes := []byte(l.Block)
		correct := ed25519.Verify(l.Dealer.EdPub, blkbytes, l.Sig)
		if !correct {
			println("SIGNATURE VERIFICATION FAILED, CONTINUING")
			continue
		}
		// TODO: W_HASHILY_MAP_BLAKE2
		var r RealListing
		err := json.Unmarshal(blkbytes, &r)
		if err != nil {
			println(err.Error())
			continue
		}
		kh := GKH(l.Dealer.EdPub, l.Dealer.CvPub)
		sig := tob64(l.Sig)
		sl := SListing{r, l.Dealer, kh, sig}
		hits = append(hits, sl)
	}
	return hits
}
