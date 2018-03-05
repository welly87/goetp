package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"./codegen"
	"./package1"
	"github.com/gorilla/websocket"
	"gopkg.in/avro.v0"
)

var addr = flag.String("addr", "localhost:9999", "http service address")

var upgrader = websocket.Upgrader{
	Subprotocols: []string{"energistics-tp"},
}

// Define the schema to read
var messageHeaderSchema = `{
	"type": "record",
	"namespace": "Energistics.Datatypes",	
	"name": "MessageHeader",
	"fields":
	[
		{ "name": "protocol", "type": "int" },
		{ "name": "messageType", "type": "int" },
		{ "name": "correlationId", "type": "long" },
		{ "name": "messageId", "type": "long" },
		{ "name": "messageFlags", "type": "int" }
	]
}`

type MessageHeader struct {
	Protocol      int32
	MessageType   int32
	CorrelationId int64
	MessageId     int64
	MessageFlags  int32
}

func parseSchema(filePath string) avro.Schema {
	bytes, _ := ioutil.ReadFile(filePath)

	schema, err := avro.ParseSchema(string(bytes))
	if err != nil {
		// Should not happen if the schema is valid
		panic(err)
	}

	return schema
}

var SchemaCache = make(map[string]avro.Schema)

func parseAllSchemas() {
	schemas := codegen.EtpSortedSchemaList()

	for _, val := range schemas {
		// schemas[i] = strings.Replace(val, ".", "/", -1) + ".avsc"

		filePath := strings.Replace(val, ".", "/", -1) + ".avsc"

		SchemaCache[val] = parseSchema(filePath)
	}

	fmt.Println(SchemaCache["Energistics.Protocol.Core.RequestSession"].Prop("messageType"))

}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// Parse the schema first
	schema, err := avro.ParseSchema(messageHeaderSchema)
	if err != nil {
		// Should not happen if the schema is valid
		panic(err)
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		// Create a new Decoder with a given buffer
		decoder := avro.NewBinaryDecoder(message)

		reader := avro.NewSpecificDatumReader()
		// SetSchema must be called before calling Read
		reader.SetSchema(schema)

		// Create a new MessageHeader to decode data into
		header := new(MessageHeader)

		// Read data into a given record with a given Decoder.
		err = reader.Read(header, decoder)
		if err != nil {
			panic(err)
		}

		// fmt.Println("Message Header")

		// b, err := json.Marshal(header)

		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("%s\n", b)

		if header.Protocol == 0 {
			if header.MessageType == 1 {
				// Create a new RequestSession to decode data into
				body := package1.NewRequestSession()

				bodySchema := SchemaCache["Energistics.Protocol.Core.RequestSession"]

				reader.SetSchema(bodySchema)

				// Read data into a given record with a given Decoder.
				err = reader.Read(body, decoder)

				writer := avro.NewSpecificDatumWriter()
				// SetSchema must be called before calling Write
				writer.SetSchema(schema)

				// Create a new Buffer and Encoder to write to this Buffer
				buffer := new(bytes.Buffer)
				encoder := avro.NewBinaryEncoder(buffer)

				header.MessageType = 2
				// Write the header
				writer.Write(header, encoder)

				writer.SetSchema(SchemaCache["Energistics.Protocol.Core.OpenSession"])

				openSession := package1.NewOpenSession()

				openSession.ApplicationName = "Go-ETP"
				openSession.ApplicationVersion = "1.0.0.1"
				openSession.SupportedProtocols = body.RequestedProtocols
				openSession.SupportedObjects = body.SupportedObjects

				writer.Write(openSession, encoder)

				err = c.WriteMessage(mt, buffer.Bytes())

				if err != nil {
					log.Println("write:", err)
					break
				}
			}

		} else if header.Protocol == 1 {
			if header.MessageType == 0 {
				// Create a new RequestSession to decode data into
				body := package1.NewStart()

				bodySchema := SchemaCache["Energistics.Protocol.ChannelStreaming.Start"]

				reader.SetSchema(bodySchema)

				// Read data into a given record with a given Decoder.
				err = reader.Read(body, decoder)

				fmt.Println(body.MaxMessageRate)
				fmt.Println(body.MaxDataItems)
			}
		}

	}
}

func main() {
	parseAllSchemas()

	// fmt.Println(SchemaCache)

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
