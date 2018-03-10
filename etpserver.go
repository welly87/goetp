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
			handleCoreProtocol(header, buffer, c)
		case 1:
			handleConsumerChannelStreaming(header, buffer, c)

		}
	}
}
func handleConsumerChannelStreaming(header *energistics.MessageHeader, buffer *bytes.Buffer, conn *websocket.Conn) {
	switch header.MessageType {
	case 3:
		channelData, _ := energistics.DeserializeChannelData(buffer)
		fmt.Println(channelData.Data[0].Value.Item.Double)
	}
}

func handleCoreProtocol(header *energistics.MessageHeader, buffer *bytes.Buffer, c *websocket.Conn) {
	switch header.MessageType {
	case 1:
		body, _ := energistics.DeserializeRequestSession(buffer)

		handleRequestSession(header, body)

		sendStart(c)

	case 2:


	}
}
func sendStart(c *websocket.Conn) {
	buffer := new(bytes.Buffer)

	header := energistics.NewMessageHeader()

	header.Protocol = 1

	header.MessageType = 0

	header.Serialize(buffer)

	start := energistics.NewStart()

	start.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
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
