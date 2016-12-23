package frontend

import (
  "bytes"
  "io/ioutil"
  "html/template"

  "github.com/ZenifiedFromI2P/bindata"
)

const tmpl = html.New("navbar.html").Parse(string(bindata.Asset("navbar.html")))

type TmplInputs struct {
  IsDealer bool
  Logged bool
}

func Generate(I TmplInputs) string {
  out := new(bytes.Buffer)
  tmpl.Execute(out, I)
  b, _ := ioutil.ReadAll(out)
  return string(b)
}
