package tools

import (
	"encoding/json"
)

type Dealer struct {
	EdPub []byte `json:"edpub"`
	CvPub []byte `json:"cvpub"`
	Name  string `json:"name"`
}

var selfDealer Dealer

func AssureDealer(should bool, name string) {
	if !should { //If we should not
		return // Prematurely quit
	} // else we assure the dealer obj exists otherwise server will complain
	// Prepare payload
	var k Keypair
	Store.Find("ekey", &k)
	d := Dealer{k.EdPub, (*k.CvPub)[:], name}
	selfDealer = d
	djson, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	SendToServer("/api/new/dealer", djson)
}
