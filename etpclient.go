package main

import (
	"net/url"
	"log"
	"github.com/gorilla/websocket"
	"encoding/base64"
	"net/http"
	"./etpsrc"
	"bytes"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}


func main() {
	//wss://witsmlstudio.pds.software/staging/api/etp
	u := url.URL{Scheme: "wss", Host: "witsmlstudio.pds.software", Path: "/staging/api/etp"}
	log.Printf("connecting to %s", u.String())

	var dialer = websocket.DefaultDialer

	dialer.Subprotocols = append(dialer.Subprotocols, "energistics-tp")

	c, _, err := dialer.Dial(u.String(), http.Header{"Authorization" : {"Basic " + basicAuth("ilab.user", "n@6C5rN!")}})
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	buffer := new(bytes.Buffer)

	header := energistics.NewMessageHeader()

	header.Protocol = 0
	header.MessageType = 1

	header.Serialize(buffer)

	requestSession := energistics.NewRequestSession()

	requestSession.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())

	go func() {
		//defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()


}
