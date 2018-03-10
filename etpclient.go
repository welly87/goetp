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
	//u := url.URL{Scheme: "wss", Host: "witsmlstudio.pds.software", Path: "/staging/api/etp"}
	u := url.URL{Scheme: "ws", Host: "localhost:9999", Path: "/etp"}
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
	//supportedProtocol.Role = "producer"
	supportedProtocol.Role = "consumer"
	supportedProtocol.Protocol = 1

	body.RequestedProtocols = append(body.RequestedProtocols, supportedProtocol)


	body.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())

	var quit = make(chan struct{})

	go func() {

		for {
			_, message, err := c.ReadMessage()

			buffer = bytes.NewBuffer(message)

			header, _ = energistics.DeserializeMessageHeader(buffer)


			b, err := json.Marshal(header)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s\n", b)

			switch header.Protocol {
			case 0:
				handleCoreClientProtocol(header, buffer)

				// discoverObject(c)

				// describeWell(c)
			case 1:
				handleChannelStreaming(header, buffer, c)
				//handleChannelStart(header, buffer)
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

func handleChannelStreaming(header *energistics.MessageHeader, buffer *bytes.Buffer, c *websocket.Conn ) {
	switch header.MessageType {
	case 0:
		body, _ := energistics.DeserializeStart(buffer)

		fmt.Println("MaxMessageRate => ", body.MaxMessageRate)
		fmt.Println("MaxDataItems => ", body.MaxDataItems)

		go simpleStreamer(c)
	case 1:
		describe, _ := energistics.DeserializeChannelDescribe(buffer)
		fmt.Println(describe.Uris)
	case 5:
		stop, _ := energistics.DeserializeChannelStreamingStop(buffer)
		fmt.Println(stop.Channels)
	case 1000:
		exception, _ := energistics.DeserializeProtocolException(buffer)
		fmt.Println("error", exception.ErrorMessage)

	default:
		fmt.Println("messageType : ", header.MessageType)
	}
	//body, _ := energistics.DeserializeChannelMetadata(buffer)
	//
	//for _, channel := range body.Channels {
	//	fmt.Println(channel.ChannelUri)
	//}
}
func simpleStreamer(c *websocket.Conn) {
	sendChannelMetadata(c)

	for {
		sendChannelData(c)
	}
}
func sendChannelData(c*websocket.Conn) {
	buffer := new(bytes.Buffer)

	header := energistics.NewMessageHeader()

	header.Protocol = 1

	header.MessageType = 3

	header.Serialize(buffer)

	channelData := energistics.NewChannelData()

	dataItem := energistics.NewDataItem()

	dataItem.ChannelId = 1

	dataItem.Indexes = []int64 {1}

	dataItem.Value = energistics.NewDataValue()

	dataItem.Value.Item.Double = 1.0

	dataItem.Value.Item.UnionType = energistics.UnionNullDoubleFloatIntLongStringArrayOfDoubleBoolBytesTypeEnumDouble

	channelData.Data = append(channelData.Data, dataItem)

	channelData.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
}

func sendChannelMetadata(c *websocket.Conn) {
	fmt.Println("simple streamer send channel metadata")

	channelMetadata := energistics.NewChannelMetadata()

	channelMetadataRecord := energistics.NewChannelMetadataRecord()
	indexMetadata := energistics.NewIndexMetadataRecord()

	indexMetadata.Uri.String = ""
	//indexMetadata.

	channelMetadataRecord.Indexes = append(channelMetadataRecord.Indexes, indexMetadata)


	channelMetadataRecord.ChannelId = 1
	channelMetadataRecord.ChannelUri = "eml://witsml20/Channel(fbfb17f7-3e58-4e38-91e6-f8fdf7f0613d)"

	channelMetadata.Channels = append(channelMetadata.Channels, channelMetadataRecord)

	buffer := new(bytes.Buffer)

	header := energistics.NewMessageHeader()

	header.Protocol = 1

	header.MessageType = 2

	header.Serialize(buffer)

	channelMetadata.Serialize(buffer)

	c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())
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
