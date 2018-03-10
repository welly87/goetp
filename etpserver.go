package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"./etpsrc"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9999", "http service address")

var upgrader = websocket.Upgrader{
	Subprotocols: []string{"energistics-tp"},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	go handleClient(c)
}

func handleClient(c *websocket.Conn) {

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		fmt.Println("message type : ", mt)

		if err != nil {
			log.Println("read:", err)
			break
		}

		buffer := bytes.NewBuffer(message)

		header, _ := energistics.DeserializeMessageHeader(buffer)

		fmt.Println("Message Header")

		b, err := json.Marshal(header)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", b)

		switch header.Protocol {
		case 0:
			handleCoreProtocol(header, buffer)
		}

	}
}

func handleCoreProtocol(header *energistics.MessageHeader, buffer *bytes.Buffer) {
	switch header.MessageType {
	case 1:
		body, _ := energistics.DeserializeRequestSession(buffer)

		handleRequestSession(header, body)

	case 2:


	}
}
func handleRequestSession(header *energistics.MessageHeader, requestSession *energistics.RequestSession) {
	b, err := json.Marshal(requestSession)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)

}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/etp", echo)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
