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

type ChannelStreamingStart struct {
	Channels []*ChannelStreamingInfo
}

func DeserializeChannelStreamingStart(r io.Reader) (*ChannelStreamingStart, error) {
	return readChannelStreamingStart(r)
}

func NewChannelStreamingStart() *ChannelStreamingStart {
	v := &ChannelStreamingStart{
		Channels: make([]*ChannelStreamingInfo, 0),
	}

	return v
}

func (r *ChannelStreamingStart) Schema() string {
	return "{\"fields\":[{\"name\":\"channels\",\"type\":{\"items\":{\"fields\":[{\"name\":\"channelId\",\"type\":\"long\"},{\"name\":\"startIndex\",\"type\":{\"fields\":[{\"name\":\"item\",\"type\":[\"null\",\"int\",\"long\"]}],\"name\":\"StreamingStartIndex\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"type\":\"record\"}},{\"name\":\"receiveChangeNotification\",\"type\":\"boolean\"}],\"name\":\"ChannelStreamingInfo\",\"namespace\":\"Energistics.Datatypes.ChannelData\",\"type\":\"record\"},\"type\":\"array\"}}],\"messageType\":\"4\",\"name\":\"ChannelStreamingStart\",\"namespace\":\"Energistics.Protocol.ChannelStreaming\",\"protocol\":\"1\",\"protocolRoles\":\"producer,consumer\",\"senderRole\":\"consumer\",\"type\":\"record\"}"
}

func (r *ChannelStreamingStart) Serialize(w io.Writer) error {
	return writeChannelStreamingStart(r, w)
}
