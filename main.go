package main

import (
	"strings"
	"encoding/json"
	"errors"
	"github.com/ZenifiedFromI2P/ofacrypto/tools"

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
		kvp[c[0]] = c[1]
	}
	return kvp
}

// Serialize the Curve25519 public key for the CrOrder function below.
func (self *User) CvPub() string  {
	return tools.GetCvPub()
}
// Given block, (load it from AJAX), parse it to form a list of orders
func (self *User) ParseOrders(block string) js.S {
	return js.S{tools.GetOrder(block)}
}

// Create a order with target Cv25519 public key target, with the origin's pseudonym name,
// Shipping address sa and contact information contact..
// Proof of payment should be proof
func (self *User) CrOrder(target, object, name, sa, contact, proof string) {
	tools.MakeOrder(target, object, name, sa, contact, proof)
}
