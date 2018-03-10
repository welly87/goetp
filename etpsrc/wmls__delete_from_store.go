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

type WMLS_DeleteFromStore struct {
	WMLtypeIn      string
	XMLin          string
	OptionsIn      string
	CapabilitiesIn string
}

func DeserializeWMLS_DeleteFromStore(r io.Reader) (*WMLS_DeleteFromStore, error) {
	return readWMLS_DeleteFromStore(r)
}

func NewWMLS_DeleteFromStore() *WMLS_DeleteFromStore {
	v := &WMLS_DeleteFromStore{}

	return v
}

func (r *WMLS_DeleteFromStore) Schema() string {
	return "{\"fields\":[{\"name\":\"WMLtypeIn\",\"type\":\"string\"},{\"name\":\"XMLin\",\"type\":\"string\"},{\"name\":\"OptionsIn\",\"type\":\"string\"},{\"name\":\"CapabilitiesIn\",\"type\":\"string\"}],\"messageType\":\"3\",\"name\":\"WMLS_DeleteFromStore\",\"namespace\":\"Energistics.Protocol.WitsmlSoap\",\"protocol\":\"8\",\"protocolRoles\":\"store,customer\",\"senderRole\":\"customer\",\"type\":\"record\"}"
}

func (r *WMLS_DeleteFromStore) Serialize(w io.Writer) error {
	return writeWMLS_DeleteFromStore(r, w)
}
