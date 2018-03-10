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

type Contact struct {
	OrganizationName UnionNullString
	ContactName      UnionNullString
	ContactPhone     UnionNullString
	ContactEmail     UnionNullString
}

func DeserializeContact(r io.Reader) (*Contact, error) {
	return readContact(r)
}

func NewContact() *Contact {
	v := &Contact{}

	return v
}

func (r *Contact) Schema() string {
	return "{\"fields\":[{\"name\":\"organizationName\",\"type\":[\"null\",\"string\"]},{\"name\":\"contactName\",\"type\":[\"null\",\"string\"]},{\"name\":\"contactPhone\",\"type\":[\"null\",\"string\"]},{\"name\":\"contactEmail\",\"type\":[\"null\",\"string\"]}],\"name\":\"Contact\",\"namespace\":\"Energistics.Datatypes\",\"type\":\"record\"}"
}

func (r *Contact) Serialize(w io.Writer) error {
	return writeContact(r, w)
}
