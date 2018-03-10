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

type OpenSession struct {
	ApplicationName    string
	ApplicationVersion string
	SessionId          string
	SupportedProtocols []*SupportedProtocol
	SupportedObjects   []string
}

func DeserializeOpenSession(r io.Reader) (*OpenSession, error) {
	return readOpenSession(r)
}

func NewOpenSession() *OpenSession {
	v := &OpenSession{
		SupportedProtocols: make([]*SupportedProtocol, 0),
		SupportedObjects:   make([]string, 0),
	}

	return v
}

func (r *OpenSession) Schema() string {
	return "{\"fields\":[{\"name\":\"applicationName\",\"type\":\"string\"},{\"name\":\"applicationVersion\",\"type\":\"string\"},{\"name\":\"sessionId\",\"type\":\"string\"},{\"name\":\"supportedProtocols\",\"type\":{\"items\":{\"fields\":[{\"name\":\"protocol\",\"type\":\"int\"},{\"name\":\"protocolVersion\",\"type\":{\"fields\":[{\"name\":\"major\",\"type\":\"int\"},{\"name\":\"minor\",\"type\":\"int\"},{\"name\":\"revision\",\"type\":\"int\"},{\"name\":\"patch\",\"type\":\"int\"}],\"name\":\"Version\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"}},{\"name\":\"role\",\"type\":\"string\"},{\"name\":\"protocolCapabilities\",\"type\":{\"type\":\"map\",\"values\":{\"fields\":[{\"name\":\"item\",\"type\":[\"null\",\"double\",\"float\",\"int\",\"long\",\"string\",{\"fields\":[{\"name\":\"values\",\"type\":{\"items\":\"double\",\"type\":\"array\"}}],\"name\":\"ArrayOfDouble\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"},\"boolean\",\"bytes\"]}],\"name\":\"DataValue\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"}}}],\"name\":\"SupportedProtocol\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"supportedObjects\",\"type\":{\"items\":\"string\",\"type\":\"array\"}}],\"messageType\":\"2\",\"name\":\"OpenSession\",\"namespace\":\"Energistics.Protocol.Core\",\"protocol\":\"0\",\"protocolRoles\":\"client,server\",\"senderRole\":\"server\",\"type\":\"record\"}"
}

func (r *OpenSession) Serialize(w io.Writer) error {
	return writeOpenSession(r, w)
}
