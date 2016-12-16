package tools

/*
import (
  "encoding/json"
  "crypto/rand"

  "golang.org/x/crypto/nacl/box"
)

type Order struct {
	Name    string // The real name of the originator
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

func MakeOrder(bcvpub, name, sa, contact, proof string) {
  o := Order{name, sa, contact, proof}
  ojson, err := json.Marshal(o)
  cvpub, err := fromb64(bcvpub)
  if err != nil {
    panic(err)
  }
  mpk, mpriv, err := box.GenerateKey(rand.Reader)
  if err != nil {
    panic(err)
  }
  var peerscvpub *[32]byte
  copy(peerscvpub[:], cvpub[0:32])
  nonce := GenerateNonce()
  sealed := box.Seal([]byte{}, ojson, nonce, peerscvpub, mpriv)
  n := (*nonce)[:]
  eo := EOrder{cvpub, (*mpk)[:], sealed, n}
  eojson, err := json.Marshal(eo)
  if err != nil {
    panic(err)
  }
  SendToServer("/api/new/order", eojson)
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
  json.Unmarshal(sijson, &inp)
  var hits ROrders
  for _, eo := range inp {
    var o Order
    var nonce *[24]byte
    copy(nonce[:], eo.Nonce[0:24])
    var ppk *[32]byte
    copy(ppk[:], eo.OCvPub[0:32])
    var kp Keypair
    Store.Find("ekey", &kp)
    decrypted, valid := box.Open([]byte{}, eo.Block, nonce, ppk, kp.CvPriv)
    if !valid {
      println("Decryption error, box")
      continue
    }
    err := json.Unmarshal(decrypted, &o)
    if err != nil {
      println(err.Error())
      continue
    }
    hits = append(hits, o)
  }
  return hits
}
*/
