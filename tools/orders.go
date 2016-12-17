package tools

import (
  "errors"
  "encoding/json"
  "encoding/hex"
  "crypto/rand"

  "golang.org/x/crypto/nacl/box"
)

type Order struct {
	Name    string // The real name of the originator
  Item    string
	SA      string // Shipping address
	Contact string // The contact info
	Proof   string // Proof of payment
}

type EOrder struct { //Encrypted order
  Target []byte `json:"target"`
	OCvPub []byte `json:"ocvpub"`
	Block  []byte `json:"block"`
	Nonce  []byte `json:"nonce"`
}

func MakeOrder(bcvpub, object, name, sa, contact, proof string) {
  o := Order{name, object, sa, contact, proof}
  ojson, err := json.Marshal(o)
  cvpub, err := hex.DecodeString(bcvpub)
  println(tob64(cvpub))
  if err != nil {
    panic(err)
  }
  mpk, mpriv, err := box.GenerateKey(rand.Reader)
  if err != nil {
    panic(err)
  }
  println("Stage 1 complete")
  var peerscvpub [32]byte
  copy(peerscvpub[:], cvpub[0:32])
  nonce := GenerateNonce()
  sealed := box.Seal([]byte{}, ojson, nonce, &peerscvpub, mpriv)
  println("Encryption done")
  n := (*nonce)[:]
  eo := EOrder{cvpub, (*mpk)[:], sealed, n}
  println("Stage 2 complete")
  eojson, err := json.Marshal(eo)
  if err != nil {
    panic(err)
  }
  SendToServer("/api/new/order", eojson)
  println("Stage 3 complete!")
  return
}

func GetCvPub() string {
  var kp Keypair
  Store.Find("ekey", &kp)
  slicedkey := (*kp.CvPub)[:]
  return tob64(slicedkey)
}

type Orders []EOrder
type ROrders []Order

func GetOrder(si string) ROrders {
  sijson := []byte(si)
  var inp Orders
  err := json.Unmarshal(sijson, &inp)
  if err != nil {
    panic(errors.New("Invalid JSON from server?!"))
  }
  var hits ROrders
  for _, eo := range inp {
    var o Order
    var nonce [24]byte
    copy(nonce[:], eo.Nonce[0:24])
    var ppk [32]byte
    copy(ppk[:], eo.OCvPub[0:32])
    var kp Keypair
    Store.Find("ekey", &kp)
    decrypted, valid := box.Open([]byte{}, eo.Block, &nonce, &ppk, kp.CvPriv)
    if !valid {
      println("Decryption error, box")
      continue
    }
    err = json.Unmarshal(decrypted, &o)
    if err != nil {
      println(err.Error())
      continue
    }
    hits = append(hits, o)
  }
  return hits
}
