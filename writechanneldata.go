package main

import (
	"./package1"
	"bytes"
	"gopkg.in/avro.v0"
	"log"
)


//type RequestSession struct {
//	applicationName string
//	applicationVersion string
//	requestedProtocols []SupportedProtocol
//	supportedObjects []string
//}
//func (v *RequestSession) Marshall(encoder avro.Encoder) {
//	encoder.WriteArrayStart(requestedProtocols.Length())
//	requestedProtocols.Marshall(encoder)
//	encoder.WriteArrayNext(0)
//	encoder.WriteArrayStart(supportedObjects.Length())
//	supportedObjects.Marshall(encoder)
//	encoder.WriteArrayNext(0)
//}
//func (v *RequestSession) Unmarshall(decoder avro.Decoder) {
//}

func main() {
	ParseAllSchemas()
	writerSchema := WriterSchemaCache[""]

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(writerSchema)
	buffer := new(bytes.Buffer)
	encoder := avro.NewBinaryEncoder(buffer)


	channelData :=  package1.NewChannelData()

	dataItem := package1.NewDataItem()

	dataItem.ChannelId = 1
	dataItem.Value = package1.NewDataValue()
	dataItem.Value.Item = 1

	channelData.Data = append(channelData.Data, dataItem)

	err := writer.Write(channelData, encoder)

	if err != nil {
		log.Println("write:", err)
	}
}
