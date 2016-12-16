package tools

import (
	"bytes"
	"net/http"
)

func SendToServer(url string, content []byte) {
	read := bytes.NewReader(content)
	go func() {
		g, e := http.Post(url, "application/json", read)
		g.Body.Close()
		if e != nil {
			panic(e)
		}
	}()
}
