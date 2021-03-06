// Code generated by github.com/alanctgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     AnyArray.avsc
 *     ArrayOfBoolean.avsc
 *     ArrayOfDouble.avsc
 *     ArrayOfFloat.avsc
 *     ArrayOfInt.avsc
 *     ArrayOfLong.avsc
 *     ChannelIndexTypes.avsc
 *     ChannelMetadataRecord.avsc
 *     ChannelRangeInfo.avsc
 *     ChannelStatuses.avsc
 *     ChannelStreamingInfo.avsc
 *     DataFrame.avsc
 *     DataItem.avsc
 *     IndexDirections.avsc
 *     IndexMetadataRecord.avsc
 *     Roles.avsc
 *     StreamingStartIndex.avsc
 *     Contact.avsc
 *     DataAttribute.avsc
 *     DataValue.avsc
 *     MessageHeader.avsc
 *     DataObject.avsc
 *     GrowingObjectIndex.avsc
 *     NotificationRequestRecord.avsc
 *     ObjectChange.avsc
 *     ObjectChangeTypes.avsc
 *     Resource.avsc
 *     Protocols.avsc
 *     ServerCapabilities.avsc
 *     SupportedProtocol.avsc
 *     Version.avsc
 *     ChannelDataFrameSet.avsc
 *     ChannelMetadata.avsc
 *     RequestChannelData.avsc
 *     ChannelData.avsc
 *     ChannelDataChange.avsc
 *     ChannelDescribe.avsc
 *     ChannelMetadata.avsc
 *     ChannelRangeRequest.avsc
 *     ChannelRemove.avsc
 *     ChannelStatusChange.avsc
 *     ChannelStreamingStart.avsc
 *     ChannelStreamingStop.avsc
 *     Start.avsc
 *     Acknowledge.avsc
 *     CloseSession.avsc
 *     OpenSession.avsc
 *     ProtocolException.avsc
 *     RenewSecurityToken.avsc
 *     RequestSession.avsc
 *     DataArray.avsc
 *     GetDataArray.avsc
 *     GetDataArraySlice.avsc
 *     PutDataArray.avsc
 *     PutDataArraySlice.avsc
 *     GetResources.avsc
 *     GetResourcesResponse.avsc
 *     GrowingObjectDelete.avsc
 *     GrowingObjectDeleteRange.avsc
 *     GrowingObjectGet.avsc
 *     GrowingObjectGetRange.avsc
 *     GrowingObjectPut.avsc
 *     ObjectFragment.avsc
 *     DeleteObject.avsc
 *     GetObject.avsc
 *     Object.avsc
 *     PutObject.avsc
 *     CancelNotification.avsc
 *     ChangeNotification.avsc
 *     DeleteNotification.avsc
 *     NotificationRequest.avsc
 *     WMLS_AddToStore.avsc
 *     WMLS_DeleteFromStore.avsc
 *     WMLS_GetBaseMsg.avsc
 *     WMLS_GetCap.avsc
 *     WMLS_GetFromStore.avsc
 *     WMLS_GetVersion.avsc
 *     WMLS_UpdateInStore.avsc
 *     WMSL_AddToStoreResponse.avsc
 *     WMSL_DeleteFromStoreResponse.avsc
 *     WMSL_GetBaseMsgResponse.avsc
 *     WMSL_GetCapResponse.avsc
 *     WMSL_GetFromStoreResponse.avsc
 *     WMSL_GetVersionResponse.avsc
 *     WMSL_UpdateInStoreResponse.avsc
 */

package energistics

import (
	"io"
)

type ChannelMetadataRecord struct {
	ChannelUri   string
	ChannelId    int64
	Indexes      []*IndexMetadataRecord
	ChannelName  string
	DataType     string
	Uom          string
	StartIndex   UnionNullLong
	EndIndex     UnionNullLong
	Description  string
	Status       ChannelStatuses
	ContentType  UnionNullString
	Source       string
	MeasureClass string
	Uuid         UnionNullString
	CustomData   map[string]*DataValue
	DomainObject UnionNullDataObject
}

func DeserializeChannelMetadataRecord(r io.Reader) (*ChannelMetadataRecord, error) {
	return readChannelMetadataRecord(r)
}

func NewChannelMetadataRecord() *ChannelMetadataRecord {
	v := &ChannelMetadataRecord{
		Indexes:    make([]*IndexMetadataRecord, 0),
		CustomData: make(map[string]*DataValue),
	}

	return v
}

func (r *ChannelMetadataRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"channelUri\",\"type\":\"string\"},{\"name\":\"channelId\",\"type\":\"long\"},{\"name\":\"indexes\",\"type\":{\"items\":{\"fields\":[{\"name\":\"indexType\",\"type\":{\"name\":\"ChannelIndexTypes\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"symbols\":[\"Time\",\"Depth\"],\"type\":\"enum\"}},{\"name\":\"uom\",\"type\":\"string\"},{\"name\":\"depthDatum\",\"type\":[\"null\",\"string\"]},{\"name\":\"direction\",\"type\":{\"name\":\"IndexDirections\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"symbols\":[\"Increasing\",\"Decreasing\"],\"type\":\"enum\"}},{\"name\":\"mnemonic\",\"type\":[\"null\",\"string\"]},{\"name\":\"description\",\"type\":[\"null\",\"string\"]},{\"name\":\"uri\",\"type\":[\"null\",\"string\"]},{\"name\":\"customData\",\"type\":{\"type\":\"map\",\"values\":{\"fields\":[{\"name\":\"item\",\"type\":[\"null\",\"double\",\"float\",\"int\",\"long\",\"string\",{\"fields\":[{\"name\":\"values\",\"type\":{\"items\":\"double\",\"type\":\"array\"}}],\"name\":\"ArrayOfDouble\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"},\"boolean\",\"bytes\"]}],\"name\":\"DataValue\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"}}},{\"name\":\"scale\",\"type\":\"int\"},{\"name\":\"timeDatum\",\"type\":[\"null\",\"string\"]}],\"name\":\"IndexMetadataRecord\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"channelName\",\"type\":\"string\"},{\"name\":\"dataType\",\"type\":\"string\"},{\"name\":\"uom\",\"type\":\"string\"},{\"name\":\"startIndex\",\"type\":[\"null\",\"long\"]},{\"name\":\"endIndex\",\"type\":[\"null\",\"long\"]},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"status\",\"type\":{\"name\":\"ChannelStatuses\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"symbols\":[\"Active\",\"Inactive\",\"Closed\"],\"type\":\"enum\"}},{\"name\":\"contentType\",\"type\":[\"null\",\"string\"]},{\"name\":\"source\",\"type\":\"string\"},{\"name\":\"measureClass\",\"type\":\"string\"},{\"name\":\"uuid\",\"type\":[\"null\",\"string\"]},{\"name\":\"customData\",\"type\":{\"type\":\"map\",\"values\":\"Energistics.Datatypes.DataValue\"}},{\"name\":\"domainObject\",\"type\":[\"null\",{\"fields\":[{\"name\":\"resource\",\"type\":{\"fields\":[{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"contentType\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"channelSubscribable\",\"type\":\"boolean\"},{\"name\":\"customData\",\"type\":{\"type\":\"map\",\"values\":\"string\"}},{\"name\":\"resourceType\",\"type\":\"string\"},{\"name\":\"hasChildren\",\"type\":\"int\"},{\"name\":\"uuid\",\"type\":[\"null\",\"string\"]},{\"name\":\"lastChanged\",\"type\":\"long\"},{\"name\":\"objectNotifiable\",\"type\":\"boolean\"}],\"name\":\"Resource\",\"namespace\":\"Energistics.Datatypes.Object\",\"type\":\"record\"}},{\"name\":\"contentEncoding\",\"type\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DataObject\",\"namespace\":\"Energistics.Datatypes.Object\",\"type\":\"record\"}]}],\"name\":\"ChannelMetadataRecord\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"type\":\"record\"}"
}

func (r *ChannelMetadataRecord) Serialize(w io.Writer) error {
	return writeChannelMetadataRecord(r, w)
}
