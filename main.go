package main

import (
	"strings"
	"encoding/json"
	"errors"
	"ofacrypto/tools"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("back", map[string]interface{}{
		"New": New,
	})
}

type User struct {
	Pseudonym string
	CI        string
}

// Create a new User object
func New(password, pseudonym, ci, pi string, dealer bool) *js.Object {
	return js.MakeWrapper(NewUser(password, pseudonym, ci, pi, dealer))
}

// Internal, don't use
func NewUser(password, pseudonym, ci, pi string, dealer bool) *User {
	tools.ImportKey([]byte(password), []byte(pseudonym))
	tools.AssureDealer(dealer, pi)
	u := User{pseudonym, ci}
	return &u
}

// Serialize the useless User into a single point JSON
func (self *User) Serialize() string {
	s, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}
	sjson := string(s)
	return sjson
}

// Create a new listing through arguments, payto is the payment address
func (self *User) NewListing(name, comment, payto, deltoken string) *js.Object {
	if name == "" || comment == "" || payto == "" || deltoken == "" {
		panic(errors.New("Any argument MAY NOT be empty"))
	}
	listing := tools.NewListing(name, comment, payto, deltoken, self.CI)
	return js.MakeWrapper(&listing)
}

// Extract all given listings through obj (string)
// It's caller's resonsiblity to parse and perform fuzzy searching on elements
func (self *User) GetListings(obj string) js.S {
	listings := tools.FListing(obj)
	println("I was called")
	return js.S{listings}
}

// Split hash argument to form key-value pairs through GET
func (self *User) ParseParam(hash string) map[string]string {
	// Level 1, split the hash to get a clear path
	println("Zenified's string argument parser, Version 420")
	g1params := strings.Split(hash, "/")
	gparams := g1params[len(g1params)-1]
	// Level 2, parse the GET parameters to form a & seperated string
	// ?key=value&test=test -> key=value&test=test
	aparams := strings.Split(gparams, "?")[1]
	nparams := strings.Split(aparams, "&")
	// Level 3, for each parameter in aparams, make it a key-value pair
	kvp := make(map[string]string)
	for _, p := range nparams {
		c := strings.Split(p, "=")
		k := c[0]
		v := c[1]
		kvp[k] = v
	}
	return kvp
}

// Export CvPub (Curve25519) public key for the fetch & create functions below
func (self *User) CvPub() string  {
	return tools.GetCvPub()
}
func (self *User) ParseOrders(block string) {

}
