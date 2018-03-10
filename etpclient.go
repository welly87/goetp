package main

import (
	"net/url"
	"log"
	"github.com/gorilla/websocket"
	"encoding/base64"
	"net/http"
	"./etpsrc"
	"bytes"

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

	body := energistics.NewRequestSession()


	supportedProtocol := energistics.NewSupportedProtocol()
	supportedProtocol.Role = "store"
	supportedProtocol.Protocol = 3

	body.RequestedProtocols = append(body.RequestedProtocols, supportedProtocol)

	supportedProtocol = energistics.NewSupportedProtocol()
	supportedProtocol.Role = "producer"
	supportedProtocol.Protocol = 1

	body.RequestedProtocols = append(body.RequestedProtocols, supportedProtocol)


	body.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())

	var quit = make(chan struct{})

	go func() {
		//defer close(done)
		for {
			_, message, err := c.ReadMessage()

			buffer = bytes.NewBuffer(message)

			header, _ = energistics.DeserializeMessageHeader(buffer)

			b, err := json.Marshal(body)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s\n", b)


			switch header.Protocol {
			case 0:
				handleCoreClientProtocol(header, buffer)

				// discoverObject(c)

				describeWell(c)
			case 1:
				handleChannelStreaming(header, buffer)
			case 3:
				handleResourceObject(header, buffer)
				// streamChannelData(c)
			default:
				if err != nil {
					log.Println("read:", err)
					return
				}

				log.Printf("recv: %s", message)
			}
		}
	}()

	<-quit
}
func handleChannelStreaming(header *energistics.MessageHeader, buffer *bytes.Buffer) {
	body, _ := energistics.DeserializeChannelMetadata(buffer)

	for _, channel := range body.Channels {
		fmt.Println(channel.ChannelUri)
	}
}

func describeWell(c *websocket.Conn) {
	buffer := new(bytes.Buffer)

	header := energistics.NewMessageHeader()

	header.Protocol = 1

	header.MessageType = 1

	header.Serialize(buffer)

	body := energistics.NewChannelDescribe()

	body.Uris = append(body.Uris, "eml://witsml20")

	body.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
}

func handleResourceObject(header *energistics.MessageHeader, buffer *bytes.Buffer) {
	resources, _ := energistics.DeserializeGetResourcesResponse(buffer)
	b, err := json.Marshal(resources)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}

func discoverObject(c *websocket.Conn) {

	buffer := new(bytes.Buffer)

	header := energistics.NewMessageHeader()

	header.Protocol = 3

	header.MessageType = 1

	header.Serialize(buffer)

	body := energistics.NewGetResources()

	body.Uri = "eml://witsml20" // "/'

	body.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
}

func streamChannelData(conn *websocket.Conn) {

}

func handleCoreClientProtocol(header *energistics.MessageHeader, buffer *bytes.Buffer) {
	openSession, _ := energistics.DeserializeOpenSession(buffer)
	b, err := json.Marshal(openSession)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)

}
