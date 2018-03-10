package main

import (
	"net/url"
	"log"
	"github.com/gorilla/websocket"
	"encoding/base64"
	"net/http"
	"./etpsrc"
	"bytes"

	//"gopkg.in/avro.v0"
	"encoding/json"
	"fmt"
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

	var quit = make(chan struct{})

	go func() {
		//defer close(done)
		for {
			_, message, err := c.ReadMessage()

			buffer = bytes.NewBuffer(message)

			header, _ = energistics.DeserializeMessageHeader(buffer)

			switch header.Protocol {
			case 0:
				handleCoreClientProtocol(header, buffer)
			}

			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("recv: %s", message)
		}
	}()

	<-quit
}
func handleCoreClientProtocol(header *energistics.MessageHeader, buffer *bytes.Buffer) {
	openSession, _ := energistics.DeserializeOpenSession(buffer)
	b, err := json.Marshal(openSession)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}
