package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"./codegen"
	"./package1"
	"github.com/gorilla/websocket"
	"gopkg.in/avro.v0"
	"strconv"

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

var readerSchemaCache = make(map[string]avro.Schema)

var WriterSchemaCache = make(map[string]avro.Schema)

var typeRegistry = make(map[string]reflect.Type)

func ParseAllSchemas() {
	schemas := codegen.EtpSortedSchemaList()

	for _, val := range schemas {
		filePath := strings.Replace(val, ".", "/", -1) + ".avsc"
		schema := parseSchema(filePath)
		protocol, ok1 := schema.Prop("protocol")
		messageType, ok2 := schema.Prop("messageType")

		if ok1 && ok2 {
			readerSchemaCache[protocol.(string)+"-"+messageType.(string)] = schema
			WriterSchemaCache[schema.GetName()] = schema
		}
	}

	// TODO need to list all the possible message
	typeRegistry[reflect.TypeOf(package1.RequestSession{}).Name()] = reflect.TypeOf(package1.RequestSession{})
	typeRegistry[reflect.TypeOf(package1.CloseSession{}).Name()] = reflect.TypeOf(package1.CloseSession{})
	typeRegistry[reflect.TypeOf(package1.Start{}).Name()] = reflect.TypeOf(package1.Start{})
	//fmt.Println(reflect.TypeOf(package1.RequestSession{}).Name())
}

func handle(header *MessageHeader, body interface{}, c *websocket.Conn) {
	switch header.Protocol {
	case 0:
		handleCoreProtocol(header, body, c)
	case 1:
		handleStreamingProtocol(header, body, c)
	}
}

func handleCoreProtocol(header *MessageHeader, body interface{}, c *websocket.Conn) {
	fmt.Println("handle core protocol ", reflect.TypeOf(body))

	switch msg := body.(type) {
	case *package1.RequestSession:
		handleRequestSession(header, msg, c)
	}
}

func handleStreamingProtocol(header *MessageHeader, body interface{}, c *websocket.Conn) {
	fmt.Println("handle streaming protocol ", reflect.TypeOf(body))

	switch msg := body.(type) {
	case *package1.Start:
		fmt.Println(msg.MaxMessageRate, msg.MaxDataItems)
		go func() {
			for  {
				sendChannelData(header, c)
			}
		}()

	}
}
func sendChannelData(header *MessageHeader, c *websocket.Conn) {

	schema, err := avro.ParseSchema(messageHeaderSchema)
	if err != nil {
		// Should not happen if the schema is valid
		panic(err)
	}

	writer := avro.NewSpecificDatumWriter()
	// SetSchema must be called before calling Write
	writer.SetSchema(schema)

	// Create a new Buffer and Encoder to write to this Buffer
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	header.Protocol = 1
	header.MessageType = 3

	writer.Write(header, encoder)

	// write channel data

	encoder.WriteArrayStart(1)

	encoder.WriteArrayStart(1)
	encoder.WriteLong(1) // write index
	encoder.WriteArrayNext(0)

	encoder.WriteLong(1) // write channel id

	encoder.WriteLong(1)
	encoder.WriteDouble(1.5) // write channel value

	encoder.WriteArrayStart(0)
	//encoder.WriteLong(1) // write value attribute
	encoder.WriteArrayNext(0)

	encoder.WriteArrayNext(0)

	err = c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())

	if err != nil {
		log.Println("write:", err)
	}
}

func handleRequestSession(header *MessageHeader, body *package1.RequestSession, c *websocket.Conn) {
	fmt.Println(body.ApplicationName)
	fmt.Println(body.ApplicationVersion)

	openSession := package1.NewOpenSession()

	openSession.ApplicationName = "Go-ETP"
	openSession.ApplicationVersion = "1.0.0.1"
	openSession.SupportedProtocols = body.RequestedProtocols
	openSession.SupportedObjects = body.SupportedObjects

	writeReply(header, openSession, c)
}

func writeReply(header *MessageHeader, body interface{}, c *websocket.Conn) {
	schema, err := avro.ParseSchema(messageHeaderSchema)
	if err != nil {
		// Should not happen if the schema is valid
		panic(err)
	}

	writer := avro.NewSpecificDatumWriter()
	// SetSchema must be called before calling Write
	writer.SetSchema(schema)

	// Create a new Buffer and Encoder to write to this Buffer
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	writerSchema := WriterSchemaCache[reflect.TypeOf(body).Elem().Name()]

	m, _ := writerSchema.Prop("messageType")
	i64, err := strconv.ParseInt(m.(string), 10, 32)
	header.MessageType = int32(i64)

	p, _ := writerSchema.Prop("protocol")
	i64, err = strconv.ParseInt(p.(string), 10, 32)
	header.Protocol = int32(i64)

	// Write the header
	writer.Write(header, encoder)

	//fmt.Println(buffer.Len())

	writer.SetSchema(writerSchema)

	fmt.Println("writer schema", writerSchema)

	err = writer.Write(body, encoder)

	//fmt.Println(buffer.Len())

	if err != nil {
		log.Println("write:", err)
	}

	err = c.WriteMessage(websocket.BinaryMessage, buffer.Bytes())

	if err != nil {
		log.Println("write:", err)
	}
}



func makeInstance(name string) interface{} {
	v := reflect.New(typeRegistry[name]).Elem()
	// Maybe fill in fields here if necessary
	return v.Interface()
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
	// Parse the schema first
	schema, err := avro.ParseSchema(messageHeaderSchema)
	if err != nil {
		// Should not happen if the schema is valid
		panic(err)
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()

		fmt.Println("message type : ", mt)

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

		fmt.Println("Message Header")

		b, err := json.Marshal(header)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", b)

		// Create a new RequestSession to decode data into

		bodySchema := readerSchemaCache[fmt.Sprintf("%v-%v", header.Protocol, header.MessageType)]

		body := reflect.New(typeRegistry[bodySchema.GetName()]).Interface()

		reader.SetSchema(bodySchema)

		// Read data into a given record with a given Decoder.
		err = reader.Read(body, decoder)

		//fmt.Println(err)

		handle(header, body, c)
	}
}


func test() {
	ParseAllSchemas()
	writerSchema := WriterSchemaCache["ChannelData"]

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(writerSchema)
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)

	fmt.Println(buffer.Len())

	channelData :=  package1.NewChannelData()

	dataItem := package1.NewDataItem()

	dataItem.ChannelId = 1
	dataItem.Value = package1.NewDataValue()
	dataItem.Value.Item = 1

	channelData.Data = append(channelData.Data, dataItem)

	encoder.WriteArrayStart(1)

	encoder.WriteArrayStart(1)
	encoder.WriteLong(1) // write index
	encoder.WriteArrayNext(0)

	encoder.WriteLong(1) // write channel id

	encoder.WriteLong(1)
	encoder.WriteDouble(1.5)

	encoder.WriteArrayStart(0)
	//encoder.WriteLong(1) // write index
	encoder.WriteArrayNext(0)

	encoder.WriteArrayNext(0)

	fmt.Println(buffer.Len())

	//if err != nil {
	//	log.Println("write:", err)
	//}


}

func main() {

	//test()


	ParseAllSchemas()

	// fmt.Println(SchemaCache)
	//fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/etp", echo)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
