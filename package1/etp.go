package package1

import avro "gopkg.in/avro.v0"

type ArrayOfDouble struct {
	Values []float64
}

func NewArrayOfDouble() *ArrayOfDouble {
	return &ArrayOfDouble{
		Values: make([]float64, 0),
	}
}

func (o *ArrayOfDouble) Schema() avro.Schema {
	if _ArrayOfDouble_schema_err != nil {
		panic(_ArrayOfDouble_schema_err)
	}
	return _ArrayOfDouble_schema
}

type DataValue struct {
	Item interface{}
}

func NewDataValue() *DataValue {
	return &DataValue{}
}

func (o *DataValue) Schema() avro.Schema {
	if _DataValue_schema_err != nil {
		panic(_DataValue_schema_err)
	}
	return _DataValue_schema
}

type DataFrame struct {
	Index []int64
	Data  []*DataValue
}

func NewDataFrame() *DataFrame {
	return &DataFrame{
		Index: make([]int64, 0),
		Data:  make([]*DataValue, 0),
	}
}

func (o *DataFrame) Schema() avro.Schema {
	if _DataFrame_schema_err != nil {
		panic(_DataFrame_schema_err)
	}
	return _DataFrame_schema
}

type ChannelDataFrameSet struct {
	Channels []int64
	Data     []*DataFrame
}

func NewChannelDataFrameSet() *ChannelDataFrameSet {
	return &ChannelDataFrameSet{
		Channels: make([]int64, 0),
		Data:     make([]*DataFrame, 0),
	}
}

func (o *ChannelDataFrameSet) Schema() avro.Schema {
	if _ChannelDataFrameSet_schema_err != nil {
		panic(_ChannelDataFrameSet_schema_err)
	}
	return _ChannelDataFrameSet_schema
}

type Resource struct {
	Uri                 string
	ContentType         string
	Name                string
	ChannelSubscribable bool
	CustomData          map[string]string
	ResourceType        string
	HasChildren         int32
	Uuid                interface{}
	LastChanged         int64
	ObjectNotifiable    bool
}

func NewResource() *Resource {
	return &Resource{
		CustomData: make(map[string]string),
	}
}

func (o *Resource) Schema() avro.Schema {
	if _Resource_schema_err != nil {
		panic(_Resource_schema_err)
	}
	return _Resource_schema
}

type DataObject struct {
	Resource        *Resource
	ContentEncoding string
	Data            []byte
}

func NewDataObject() *DataObject {
	return &DataObject{
		Data: []byte{},
	}
}

func (o *DataObject) Schema() avro.Schema {
	if _DataObject_schema_err != nil {
		panic(_DataObject_schema_err)
	}
	return _DataObject_schema
}

// Enum values for ChannelIndexTypes
const (
	ChannelIndexTypes_Time  int32 = 0
	ChannelIndexTypes_Depth int32 = 1
)

// Enum values for IndexDirections
const (
	IndexDirections_Increasing int32 = 0
	IndexDirections_Decreasing int32 = 1
)

type IndexMetadataRecord struct {
	IndexType   *avro.GenericEnum
	Uom         string
	DepthDatum  interface{}
	Direction   *avro.GenericEnum
	Mnemonic    interface{}
	Description interface{}
	Uri         interface{}
	CustomData  map[string]*DataValue
	Scale       int32
	TimeDatum   interface{}
}

func NewIndexMetadataRecord() *IndexMetadataRecord {
	return &IndexMetadataRecord{
		IndexType:  avro.NewGenericEnum([]string{"Time", "Depth"}),
		Direction:  avro.NewGenericEnum([]string{"Increasing", "Decreasing"}),
		CustomData: make(map[string]*DataValue),
	}
}

func (o *IndexMetadataRecord) Schema() avro.Schema {
	if _IndexMetadataRecord_schema_err != nil {
		panic(_IndexMetadataRecord_schema_err)
	}
	return _IndexMetadataRecord_schema
}

// Enum values for ChannelStatuses
const (
	ChannelStatuses_Active   int32 = 0
	ChannelStatuses_Inactive int32 = 1
	ChannelStatuses_Closed   int32 = 2
)

type ChannelMetadataRecord struct {
	ChannelUri   string
	ChannelId    int64
	Indexes      []*IndexMetadataRecord
	ChannelName  string
	DataType     string
	Uom          string
	StartIndex   interface{}
	EndIndex     interface{}
	Description  string
	Status       *avro.GenericEnum
	ContentType  interface{}
	Source       string
	MeasureClass string
	Uuid         interface{}
	CustomData   map[string]*DataValue
	DomainObject *DataObject
}

func NewChannelMetadataRecord() *ChannelMetadataRecord {
	return &ChannelMetadataRecord{
		Indexes:    make([]*IndexMetadataRecord, 0),
		Status:     avro.NewGenericEnum([]string{"Active", "Inactive", "Closed"}),
		CustomData: make(map[string]*DataValue),
	}
}

func (o *ChannelMetadataRecord) Schema() avro.Schema {
	if _ChannelMetadataRecord_schema_err != nil {
		panic(_ChannelMetadataRecord_schema_err)
	}
	return _ChannelMetadataRecord_schema
}

type ChannelMetadata struct {
	Channels []*ChannelMetadataRecord
}

func NewChannelMetadata() *ChannelMetadata {
	return &ChannelMetadata{
		Channels: make([]*ChannelMetadataRecord, 0),
	}
}

func (o *ChannelMetadata) Schema() avro.Schema {
	if _ChannelMetadata_schema_err != nil {
		panic(_ChannelMetadata_schema_err)
	}
	return _ChannelMetadata_schema
}

type RequestChannelData struct {
	Uri       string
	FromIndex interface{}
	ToIndex   interface{}
}

func NewRequestChannelData() *RequestChannelData {
	return &RequestChannelData{}
}

func (o *RequestChannelData) Schema() avro.Schema {
	if _RequestChannelData_schema_err != nil {
		panic(_RequestChannelData_schema_err)
	}
	return _RequestChannelData_schema
}

type DataAttribute struct {
	AttributeId    int32
	AttributeValue *DataValue
}

func NewDataAttribute() *DataAttribute {
	return &DataAttribute{}
}

func (o *DataAttribute) Schema() avro.Schema {
	if _DataAttribute_schema_err != nil {
		panic(_DataAttribute_schema_err)
	}
	return _DataAttribute_schema
}

type DataItem struct {
	Indexes         []int64
	ChannelId       int64
	Value           *DataValue
	ValueAttributes []*DataAttribute
}

func NewDataItem() *DataItem {
	return &DataItem{
		Indexes:         make([]int64, 0),
		ValueAttributes: make([]*DataAttribute, 0),
	}
}

func (o *DataItem) Schema() avro.Schema {
	if _DataItem_schema_err != nil {
		panic(_DataItem_schema_err)
	}
	return _DataItem_schema
}

type ChannelData struct {
	Data []*DataItem
}

func NewChannelData() *ChannelData {
	return &ChannelData{
		Data: make([]*DataItem, 0),
	}
}

func (o *ChannelData) Schema() avro.Schema {
	if _ChannelData_schema_err != nil {
		panic(_ChannelData_schema_err)
	}
	return _ChannelData_schema
}

type ChannelDataChange struct {
	ChannelId  int64
	StartIndex int64
	EndIndex   int64
	Data       []*DataItem
}

func NewChannelDataChange() *ChannelDataChange {
	return &ChannelDataChange{
		Data: make([]*DataItem, 0),
	}
}

func (o *ChannelDataChange) Schema() avro.Schema {
	if _ChannelDataChange_schema_err != nil {
		panic(_ChannelDataChange_schema_err)
	}
	return _ChannelDataChange_schema
}

type ChannelDescribe struct {
	Uris []string
}

func NewChannelDescribe() *ChannelDescribe {
	return &ChannelDescribe{
		Uris: make([]string, 0),
	}
}

func (o *ChannelDescribe) Schema() avro.Schema {
	if _ChannelDescribe_schema_err != nil {
		panic(_ChannelDescribe_schema_err)
	}
	return _ChannelDescribe_schema
}

type ChannelRangeInfo struct {
	ChannelId  []int64
	StartIndex int64
	EndIndex   int64
}

func NewChannelRangeInfo() *ChannelRangeInfo {
	return &ChannelRangeInfo{
		ChannelId: make([]int64, 0),
	}
}

func (o *ChannelRangeInfo) Schema() avro.Schema {
	if _ChannelRangeInfo_schema_err != nil {
		panic(_ChannelRangeInfo_schema_err)
	}
	return _ChannelRangeInfo_schema
}

type ChannelRangeRequest struct {
	ChannelRanges []*ChannelRangeInfo
}

func NewChannelRangeRequest() *ChannelRangeRequest {
	return &ChannelRangeRequest{
		ChannelRanges: make([]*ChannelRangeInfo, 0),
	}
}

func (o *ChannelRangeRequest) Schema() avro.Schema {
	if _ChannelRangeRequest_schema_err != nil {
		panic(_ChannelRangeRequest_schema_err)
	}
	return _ChannelRangeRequest_schema
}

type ChannelRemove struct {
	ChannelId    int64
	RemoveReason interface{}
}

func NewChannelRemove() *ChannelRemove {
	return &ChannelRemove{}
}

func (o *ChannelRemove) Schema() avro.Schema {
	if _ChannelRemove_schema_err != nil {
		panic(_ChannelRemove_schema_err)
	}
	return _ChannelRemove_schema
}

type ChannelStatusChange struct {
	ChannelId int64
	Status    *avro.GenericEnum
}

func NewChannelStatusChange() *ChannelStatusChange {
	return &ChannelStatusChange{
		Status: avro.NewGenericEnum([]string{"Active", "Inactive", "Closed"}),
	}
}

func (o *ChannelStatusChange) Schema() avro.Schema {
	if _ChannelStatusChange_schema_err != nil {
		panic(_ChannelStatusChange_schema_err)
	}
	return _ChannelStatusChange_schema
}

type StreamingStartIndex struct {
	Item interface{}
}

func NewStreamingStartIndex() *StreamingStartIndex {
	return &StreamingStartIndex{}
}

func (o *StreamingStartIndex) Schema() avro.Schema {
	if _StreamingStartIndex_schema_err != nil {
		panic(_StreamingStartIndex_schema_err)
	}
	return _StreamingStartIndex_schema
}

type ChannelStreamingInfo struct {
	ChannelId                 int64
	StartIndex                *StreamingStartIndex
	ReceiveChangeNotification bool
}

func NewChannelStreamingInfo() *ChannelStreamingInfo {
	return &ChannelStreamingInfo{}
}

func (o *ChannelStreamingInfo) Schema() avro.Schema {
	if _ChannelStreamingInfo_schema_err != nil {
		panic(_ChannelStreamingInfo_schema_err)
	}
	return _ChannelStreamingInfo_schema
}

type ChannelStreamingStart struct {
	Channels []*ChannelStreamingInfo
}

func NewChannelStreamingStart() *ChannelStreamingStart {
	return &ChannelStreamingStart{
		Channels: make([]*ChannelStreamingInfo, 0),
	}
}

func (o *ChannelStreamingStart) Schema() avro.Schema {
	if _ChannelStreamingStart_schema_err != nil {
		panic(_ChannelStreamingStart_schema_err)
	}
	return _ChannelStreamingStart_schema
}

type ChannelStreamingStop struct {
	Channels []int64
}

func NewChannelStreamingStop() *ChannelStreamingStop {
	return &ChannelStreamingStop{
		Channels: make([]int64, 0),
	}
}

func (o *ChannelStreamingStop) Schema() avro.Schema {
	if _ChannelStreamingStop_schema_err != nil {
		panic(_ChannelStreamingStop_schema_err)
	}
	return _ChannelStreamingStop_schema
}

type Start struct {
	MaxMessageRate int32
	MaxDataItems   int32
}

func NewStart() *Start {
	return &Start{}
}

func (o *Start) Schema() avro.Schema {
	if _Start_schema_err != nil {
		panic(_Start_schema_err)
	}
	return _Start_schema
}

type Acknowledge struct {
}

func NewAcknowledge() *Acknowledge {
	return &Acknowledge{}
}

func (o *Acknowledge) Schema() avro.Schema {
	if _Acknowledge_schema_err != nil {
		panic(_Acknowledge_schema_err)
	}
	return _Acknowledge_schema
}

type CloseSession struct {
	Reason interface{}
}

func NewCloseSession() *CloseSession {
	return &CloseSession{}
}

func (o *CloseSession) Schema() avro.Schema {
	if _CloseSession_schema_err != nil {
		panic(_CloseSession_schema_err)
	}
	return _CloseSession_schema
}

type Version struct {
	Major    int32
	Minor    int32
	Revision int32
	Patch    int32
}

func NewVersion() *Version {
	return &Version{}
}

func (o *Version) Schema() avro.Schema {
	if _Version_schema_err != nil {
		panic(_Version_schema_err)
	}
	return _Version_schema
}

type SupportedProtocol struct {
	Protocol             int32
	ProtocolVersion      *Version
	Role                 string
	ProtocolCapabilities map[string]*DataValue
}

func NewSupportedProtocol() *SupportedProtocol {
	return &SupportedProtocol{
		ProtocolCapabilities: make(map[string]*DataValue),
	}
}

func (o *SupportedProtocol) Schema() avro.Schema {
	if _SupportedProtocol_schema_err != nil {
		panic(_SupportedProtocol_schema_err)
	}
	return _SupportedProtocol_schema
}

type OpenSession struct {
	ApplicationName    string
	ApplicationVersion string
	SessionId          string
	SupportedProtocols []*SupportedProtocol
	SupportedObjects   []string
}

func NewOpenSession() *OpenSession {
	return &OpenSession{
		SupportedProtocols: make([]*SupportedProtocol, 0),
		SupportedObjects:   make([]string, 0),
	}
}

func (o *OpenSession) Schema() avro.Schema {
	if _OpenSession_schema_err != nil {
		panic(_OpenSession_schema_err)
	}
	return _OpenSession_schema
}

type ProtocolException struct {
	ErrorCode    int32
	ErrorMessage string
}

func NewProtocolException() *ProtocolException {
	return &ProtocolException{}
}

func (o *ProtocolException) Schema() avro.Schema {
	if _ProtocolException_schema_err != nil {
		panic(_ProtocolException_schema_err)
	}
	return _ProtocolException_schema
}

type RenewSecurityToken struct {
	Token string
}

func NewRenewSecurityToken() *RenewSecurityToken {
	return &RenewSecurityToken{}
}

func (o *RenewSecurityToken) Schema() avro.Schema {
	if _RenewSecurityToken_schema_err != nil {
		panic(_RenewSecurityToken_schema_err)
	}
	return _RenewSecurityToken_schema
}

type RequestSession struct {
	ApplicationName    string
	ApplicationVersion string
	RequestedProtocols []*SupportedProtocol
	SupportedObjects   []string
}

func NewRequestSession() *RequestSession {
	return &RequestSession{
		RequestedProtocols: make([]*SupportedProtocol, 0),
		SupportedObjects:   make([]string, 0),
	}
}

func (o *RequestSession) Schema() avro.Schema {
	if _RequestSession_schema_err != nil {
		panic(_RequestSession_schema_err)
	}
	return _RequestSession_schema
}

type ArrayOfBoolean struct {
	Values []bool
}

func NewArrayOfBoolean() *ArrayOfBoolean {
	return &ArrayOfBoolean{
		Values: make([]bool, 0),
	}
}

func (o *ArrayOfBoolean) Schema() avro.Schema {
	if _ArrayOfBoolean_schema_err != nil {
		panic(_ArrayOfBoolean_schema_err)
	}
	return _ArrayOfBoolean_schema
}

type ArrayOfInt struct {
	Values []int32
}

func NewArrayOfInt() *ArrayOfInt {
	return &ArrayOfInt{
		Values: make([]int32, 0),
	}
}

func (o *ArrayOfInt) Schema() avro.Schema {
	if _ArrayOfInt_schema_err != nil {
		panic(_ArrayOfInt_schema_err)
	}
	return _ArrayOfInt_schema
}

type ArrayOfLong struct {
	Values []int64
}

func NewArrayOfLong() *ArrayOfLong {
	return &ArrayOfLong{
		Values: make([]int64, 0),
	}
}

func (o *ArrayOfLong) Schema() avro.Schema {
	if _ArrayOfLong_schema_err != nil {
		panic(_ArrayOfLong_schema_err)
	}
	return _ArrayOfLong_schema
}

type ArrayOfFloat struct {
	Values []float32
}

func NewArrayOfFloat() *ArrayOfFloat {
	return &ArrayOfFloat{
		Values: make([]float32, 0),
	}
}

func (o *ArrayOfFloat) Schema() avro.Schema {
	if _ArrayOfFloat_schema_err != nil {
		panic(_ArrayOfFloat_schema_err)
	}
	return _ArrayOfFloat_schema
}

type AnyArray struct {
	Item *ArrayOfBoolean
}

func NewAnyArray() *AnyArray {
	return &AnyArray{}
}

func (o *AnyArray) Schema() avro.Schema {
	if _AnyArray_schema_err != nil {
		panic(_AnyArray_schema_err)
	}
	return _AnyArray_schema
}

type DataArray struct {
	Dimensions []int64
	Data       *AnyArray
}

func NewDataArray() *DataArray {
	return &DataArray{
		Dimensions: make([]int64, 0),
	}
}

func (o *DataArray) Schema() avro.Schema {
	if _DataArray_schema_err != nil {
		panic(_DataArray_schema_err)
	}
	return _DataArray_schema
}

type GetDataArray struct {
	Uri string
}

func NewGetDataArray() *GetDataArray {
	return &GetDataArray{}
}

func (o *GetDataArray) Schema() avro.Schema {
	if _GetDataArray_schema_err != nil {
		panic(_GetDataArray_schema_err)
	}
	return _GetDataArray_schema
}

type GetDataArraySlice struct {
	Uri   string
	Start []int64
	Count []int64
}

func NewGetDataArraySlice() *GetDataArraySlice {
	return &GetDataArraySlice{
		Start: make([]int64, 0),
		Count: make([]int64, 0),
	}
}

func (o *GetDataArraySlice) Schema() avro.Schema {
	if _GetDataArraySlice_schema_err != nil {
		panic(_GetDataArraySlice_schema_err)
	}
	return _GetDataArraySlice_schema
}

type PutDataArray struct {
	Uri        string
	Data       *AnyArray
	Dimensions []int64
}

func NewPutDataArray() *PutDataArray {
	return &PutDataArray{
		Dimensions: make([]int64, 0),
	}
}

func (o *PutDataArray) Schema() avro.Schema {
	if _PutDataArray_schema_err != nil {
		panic(_PutDataArray_schema_err)
	}
	return _PutDataArray_schema
}

type PutDataArraySlice struct {
	Uri        string
	Data       *AnyArray
	Dimensions []int64
	Start      []int64
	Count      []int64
}

func NewPutDataArraySlice() *PutDataArraySlice {
	return &PutDataArraySlice{
		Dimensions: make([]int64, 0),
		Start:      make([]int64, 0),
		Count:      make([]int64, 0),
	}
}

func (o *PutDataArraySlice) Schema() avro.Schema {
	if _PutDataArraySlice_schema_err != nil {
		panic(_PutDataArraySlice_schema_err)
	}
	return _PutDataArraySlice_schema
}

type GetResources struct {
	Uri string
}

func NewGetResources() *GetResources {
	return &GetResources{}
}

func (o *GetResources) Schema() avro.Schema {
	if _GetResources_schema_err != nil {
		panic(_GetResources_schema_err)
	}
	return _GetResources_schema
}

type GetResourcesResponse struct {
	Resource *Resource
}

func NewGetResourcesResponse() *GetResourcesResponse {
	return &GetResourcesResponse{}
}

func (o *GetResourcesResponse) Schema() avro.Schema {
	if _GetResourcesResponse_schema_err != nil {
		panic(_GetResourcesResponse_schema_err)
	}
	return _GetResourcesResponse_schema
}

type GrowingObjectDelete struct {
	Uri string
	Uid string
}

func NewGrowingObjectDelete() *GrowingObjectDelete {
	return &GrowingObjectDelete{}
}

func (o *GrowingObjectDelete) Schema() avro.Schema {
	if _GrowingObjectDelete_schema_err != nil {
		panic(_GrowingObjectDelete_schema_err)
	}
	return _GrowingObjectDelete_schema
}

type GrowingObjectIndex struct {
	Item interface{}
}

func NewGrowingObjectIndex() *GrowingObjectIndex {
	return &GrowingObjectIndex{}
}

func (o *GrowingObjectIndex) Schema() avro.Schema {
	if _GrowingObjectIndex_schema_err != nil {
		panic(_GrowingObjectIndex_schema_err)
	}
	return _GrowingObjectIndex_schema
}

type GrowingObjectDeleteRange struct {
	Uri        string
	StartIndex *GrowingObjectIndex
	EndIndex   *GrowingObjectIndex
}

func NewGrowingObjectDeleteRange() *GrowingObjectDeleteRange {
	return &GrowingObjectDeleteRange{}
}

func (o *GrowingObjectDeleteRange) Schema() avro.Schema {
	if _GrowingObjectDeleteRange_schema_err != nil {
		panic(_GrowingObjectDeleteRange_schema_err)
	}
	return _GrowingObjectDeleteRange_schema
}

type GrowingObjectGet struct {
	Uri string
	Uid string
}

func NewGrowingObjectGet() *GrowingObjectGet {
	return &GrowingObjectGet{}
}

func (o *GrowingObjectGet) Schema() avro.Schema {
	if _GrowingObjectGet_schema_err != nil {
		panic(_GrowingObjectGet_schema_err)
	}
	return _GrowingObjectGet_schema
}

type GrowingObjectGetRange struct {
	Uri        string
	StartIndex *GrowingObjectIndex
	EndIndex   *GrowingObjectIndex
	Uom        string
	DepthDatum string
}

func NewGrowingObjectGetRange() *GrowingObjectGetRange {
	return &GrowingObjectGetRange{}
}

func (o *GrowingObjectGetRange) Schema() avro.Schema {
	if _GrowingObjectGetRange_schema_err != nil {
		panic(_GrowingObjectGetRange_schema_err)
	}
	return _GrowingObjectGetRange_schema
}

type GrowingObjectPut struct {
	Uri             string
	ContentType     string
	ContentEncoding string
	Data            []byte
}

func NewGrowingObjectPut() *GrowingObjectPut {
	return &GrowingObjectPut{
		Data: []byte{},
	}
}

func (o *GrowingObjectPut) Schema() avro.Schema {
	if _GrowingObjectPut_schema_err != nil {
		panic(_GrowingObjectPut_schema_err)
	}
	return _GrowingObjectPut_schema
}

type ObjectFragment struct {
	Uri             string
	ContentType     string
	ContentEncoding string
	Data            []byte
}

func NewObjectFragment() *ObjectFragment {
	return &ObjectFragment{
		Data: []byte{},
	}
}

func (o *ObjectFragment) Schema() avro.Schema {
	if _ObjectFragment_schema_err != nil {
		panic(_ObjectFragment_schema_err)
	}
	return _ObjectFragment_schema
}

type DeleteObject struct {
	Uri string
}

func NewDeleteObject() *DeleteObject {
	return &DeleteObject{}
}

func (o *DeleteObject) Schema() avro.Schema {
	if _DeleteObject_schema_err != nil {
		panic(_DeleteObject_schema_err)
	}
	return _DeleteObject_schema
}

type GetObject struct {
	Uri string
}

func NewGetObject() *GetObject {
	return &GetObject{}
}

func (o *GetObject) Schema() avro.Schema {
	if _GetObject_schema_err != nil {
		panic(_GetObject_schema_err)
	}
	return _GetObject_schema
}

type Object struct {
	DataObject *DataObject
}

func NewObject() *Object {
	return &Object{}
}

func (o *Object) Schema() avro.Schema {
	if _Object_schema_err != nil {
		panic(_Object_schema_err)
	}
	return _Object_schema
}

type PutObject struct {
	DataObject *DataObject
}

func NewPutObject() *PutObject {
	return &PutObject{}
}

func (o *PutObject) Schema() avro.Schema {
	if _PutObject_schema_err != nil {
		panic(_PutObject_schema_err)
	}
	return _PutObject_schema
}

type CancelNotification struct {
	RequestUuid string
}

func NewCancelNotification() *CancelNotification {
	return &CancelNotification{}
}

func (o *CancelNotification) Schema() avro.Schema {
	if _CancelNotification_schema_err != nil {
		panic(_CancelNotification_schema_err)
	}
	return _CancelNotification_schema
}

// Enum values for ObjectChangeTypes
const (
	ObjectChangeTypes_Upsert int32 = 0
	ObjectChangeTypes_Delete int32 = 1
)

type ObjectChange struct {
	ChangeType *avro.GenericEnum
	ChangeTime int64
	DataObject *DataObject
}

func NewObjectChange() *ObjectChange {
	return &ObjectChange{
		ChangeType: avro.NewGenericEnum([]string{"Upsert", "Delete"}),
	}
}

func (o *ObjectChange) Schema() avro.Schema {
	if _ObjectChange_schema_err != nil {
		panic(_ObjectChange_schema_err)
	}
	return _ObjectChange_schema
}

type ChangeNotification struct {
	Change *ObjectChange
}

func NewChangeNotification() *ChangeNotification {
	return &ChangeNotification{}
}

func (o *ChangeNotification) Schema() avro.Schema {
	if _ChangeNotification_schema_err != nil {
		panic(_ChangeNotification_schema_err)
	}
	return _ChangeNotification_schema
}

type DeleteNotification struct {
	Delete *ObjectChange
}

func NewDeleteNotification() *DeleteNotification {
	return &DeleteNotification{}
}

func (o *DeleteNotification) Schema() avro.Schema {
	if _DeleteNotification_schema_err != nil {
		panic(_DeleteNotification_schema_err)
	}
	return _DeleteNotification_schema
}

type NotificationRequestRecord struct {
	Uri               string
	Uuid              string
	IncludeObjectData bool
	StartTime         int64
	ObjectTypes       []string
}

func NewNotificationRequestRecord() *NotificationRequestRecord {
	return &NotificationRequestRecord{
		ObjectTypes: make([]string, 0),
	}
}

func (o *NotificationRequestRecord) Schema() avro.Schema {
	if _NotificationRequestRecord_schema_err != nil {
		panic(_NotificationRequestRecord_schema_err)
	}
	return _NotificationRequestRecord_schema
}

type NotificationRequest struct {
	Request *NotificationRequestRecord
}

func NewNotificationRequest() *NotificationRequest {
	return &NotificationRequest{}
}

func (o *NotificationRequest) Schema() avro.Schema {
	if _NotificationRequest_schema_err != nil {
		panic(_NotificationRequest_schema_err)
	}
	return _NotificationRequest_schema
}

type WMLS_AddToStore struct {
	WMLtypeIn      string
	XMLin          string
	OptionsIn      string
	CapabilitiesIn string
}

func NewWMLS_AddToStore() *WMLS_AddToStore {
	return &WMLS_AddToStore{}
}

func (o *WMLS_AddToStore) Schema() avro.Schema {
	if _WMLS_AddToStore_schema_err != nil {
		panic(_WMLS_AddToStore_schema_err)
	}
	return _WMLS_AddToStore_schema
}

type WMLS_DeleteFromStore struct {
	WMLtypeIn      string
	XMLin          string
	OptionsIn      string
	CapabilitiesIn string
}

func NewWMLS_DeleteFromStore() *WMLS_DeleteFromStore {
	return &WMLS_DeleteFromStore{}
}

func (o *WMLS_DeleteFromStore) Schema() avro.Schema {
	if _WMLS_DeleteFromStore_schema_err != nil {
		panic(_WMLS_DeleteFromStore_schema_err)
	}
	return _WMLS_DeleteFromStore_schema
}

type WMLS_GetBaseMsg struct {
	ReturnValueIn int32
}

func NewWMLS_GetBaseMsg() *WMLS_GetBaseMsg {
	return &WMLS_GetBaseMsg{}
}

func (o *WMLS_GetBaseMsg) Schema() avro.Schema {
	if _WMLS_GetBaseMsg_schema_err != nil {
		panic(_WMLS_GetBaseMsg_schema_err)
	}
	return _WMLS_GetBaseMsg_schema
}

type WMLS_GetCap struct {
	OptionsIn string
}

func NewWMLS_GetCap() *WMLS_GetCap {
	return &WMLS_GetCap{}
}

func (o *WMLS_GetCap) Schema() avro.Schema {
	if _WMLS_GetCap_schema_err != nil {
		panic(_WMLS_GetCap_schema_err)
	}
	return _WMLS_GetCap_schema
}

type WMLS_GetFromStore struct {
	WMLtypeIn      string
	XMLin          string
	OptionsIn      string
	CapabilitiesIn string
}

func NewWMLS_GetFromStore() *WMLS_GetFromStore {
	return &WMLS_GetFromStore{}
}

func (o *WMLS_GetFromStore) Schema() avro.Schema {
	if _WMLS_GetFromStore_schema_err != nil {
		panic(_WMLS_GetFromStore_schema_err)
	}
	return _WMLS_GetFromStore_schema
}

type WMLS_GetVersion struct {
}

func NewWMLS_GetVersion() *WMLS_GetVersion {
	return &WMLS_GetVersion{}
}

func (o *WMLS_GetVersion) Schema() avro.Schema {
	if _WMLS_GetVersion_schema_err != nil {
		panic(_WMLS_GetVersion_schema_err)
	}
	return _WMLS_GetVersion_schema
}

type WMLS_UpdateInStore struct {
	WMLtypeIn      string
	XMLin          string
	OptionsIn      string
	CapabilitiesIn string
}

func NewWMLS_UpdateInStore() *WMLS_UpdateInStore {
	return &WMLS_UpdateInStore{}
}

func (o *WMLS_UpdateInStore) Schema() avro.Schema {
	if _WMLS_UpdateInStore_schema_err != nil {
		panic(_WMLS_UpdateInStore_schema_err)
	}
	return _WMLS_UpdateInStore_schema
}

type WMSL_AddToStoreResponse struct {
	Result     int32
	SuppMsgOut string
}

func NewWMSL_AddToStoreResponse() *WMSL_AddToStoreResponse {
	return &WMSL_AddToStoreResponse{}
}

func (o *WMSL_AddToStoreResponse) Schema() avro.Schema {
	if _WMSL_AddToStoreResponse_schema_err != nil {
		panic(_WMSL_AddToStoreResponse_schema_err)
	}
	return _WMSL_AddToStoreResponse_schema
}

type WMSL_DeleteFromStoreResponse struct {
	Result     int32
	SuppMsgOut string
}

func NewWMSL_DeleteFromStoreResponse() *WMSL_DeleteFromStoreResponse {
	return &WMSL_DeleteFromStoreResponse{}
}

func (o *WMSL_DeleteFromStoreResponse) Schema() avro.Schema {
	if _WMSL_DeleteFromStoreResponse_schema_err != nil {
		panic(_WMSL_DeleteFromStoreResponse_schema_err)
	}
	return _WMSL_DeleteFromStoreResponse_schema
}

type WMSL_GetBaseMsgResponse struct {
	Result string
}

func NewWMSL_GetBaseMsgResponse() *WMSL_GetBaseMsgResponse {
	return &WMSL_GetBaseMsgResponse{}
}

func (o *WMSL_GetBaseMsgResponse) Schema() avro.Schema {
	if _WMSL_GetBaseMsgResponse_schema_err != nil {
		panic(_WMSL_GetBaseMsgResponse_schema_err)
	}
	return _WMSL_GetBaseMsgResponse_schema
}

type WMSL_GetCapResponse struct {
	Result          int32
	CapabilitiesOut string
	SuppMsgOut      string
}

func NewWMSL_GetCapResponse() *WMSL_GetCapResponse {
	return &WMSL_GetCapResponse{}
}

func (o *WMSL_GetCapResponse) Schema() avro.Schema {
	if _WMSL_GetCapResponse_schema_err != nil {
		panic(_WMSL_GetCapResponse_schema_err)
	}
	return _WMSL_GetCapResponse_schema
}

type WMSL_GetFromStoreResponse struct {
	Result     int32
	XMLout     string
	SuppMsgOut string
}

func NewWMSL_GetFromStoreResponse() *WMSL_GetFromStoreResponse {
	return &WMSL_GetFromStoreResponse{}
}

func (o *WMSL_GetFromStoreResponse) Schema() avro.Schema {
	if _WMSL_GetFromStoreResponse_schema_err != nil {
		panic(_WMSL_GetFromStoreResponse_schema_err)
	}
	return _WMSL_GetFromStoreResponse_schema
}

type WMSL_GetVersionResponse struct {
	Result string
}

func NewWMSL_GetVersionResponse() *WMSL_GetVersionResponse {
	return &WMSL_GetVersionResponse{}
}

func (o *WMSL_GetVersionResponse) Schema() avro.Schema {
	if _WMSL_GetVersionResponse_schema_err != nil {
		panic(_WMSL_GetVersionResponse_schema_err)
	}
	return _WMSL_GetVersionResponse_schema
}

type WMSL_UpdateInStoreResponse struct {
	Result     int32
	SuppMsgOut string
}

func NewWMSL_UpdateInStoreResponse() *WMSL_UpdateInStoreResponse {
	return &WMSL_UpdateInStoreResponse{}
}

func (o *WMSL_UpdateInStoreResponse) Schema() avro.Schema {
	if _WMSL_UpdateInStoreResponse_schema_err != nil {
		panic(_WMSL_UpdateInStoreResponse_schema_err)
	}
	return _WMSL_UpdateInStoreResponse_schema
}

// Generated by codegen. Please do not modify.
var _ArrayOfDouble_schema, _ArrayOfDouble_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "ArrayOfDouble",
    "fields": [
        {
            "name": "values",
            "type": {
                "type": "array",
                "items": "double"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DataValue_schema, _DataValue_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "DataValue",
    "fields": [
        {
            "name": "item",
            "default": null,
            "type": [
                "null",
                "double",
                "float",
                "int",
                "long",
                "string",
                "ArrayOfDouble",
                "boolean",
                "bytes"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DataFrame_schema, _DataFrame_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "DataFrame",
    "fields": [
        {
            "name": "index",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "data",
            "type": {
                "type": "array",
                "items": "DataValue"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelDataFrameSet_schema, _ChannelDataFrameSet_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelDataFrame",
    "name": "ChannelDataFrameSet",
    "fields": [
        {
            "name": "channels",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "data",
            "type": {
                "type": "array",
                "items": "DataFrame"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Resource_schema, _Resource_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.Object",
    "name": "Resource",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "contentType",
            "type": "string"
        },
        {
            "name": "name",
            "type": "string"
        },
        {
            "name": "channelSubscribable",
            "type": "boolean"
        },
        {
            "name": "customData",
            "type": {
                "type": "map",
                "values": "string"
            }
        },
        {
            "name": "resourceType",
            "type": "string"
        },
        {
            "name": "hasChildren",
            "type": "int"
        },
        {
            "name": "uuid",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "lastChanged",
            "type": "long"
        },
        {
            "name": "objectNotifiable",
            "type": "boolean"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DataObject_schema, _DataObject_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.Object",
    "name": "DataObject",
    "fields": [
        {
            "name": "resource",
            "type": "Resource"
        },
        {
            "name": "contentEncoding",
            "type": "string"
        },
        {
            "name": "data",
            "type": "bytes"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _IndexMetadataRecord_schema, _IndexMetadataRecord_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "IndexMetadataRecord",
    "fields": [
        {
            "name": "indexType",
            "type": {
                "type": "enum",
                "namespace": "Energistics.Datatypes.ChannelData",
                "name": "ChannelIndexTypes",
                "symbols": [
                    "Time",
                    "Depth"
                ]
            }
        },
        {
            "name": "uom",
            "type": "string"
        },
        {
            "name": "depthDatum",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "direction",
            "type": {
                "type": "enum",
                "namespace": "Energistics.Datatypes.ChannelData",
                "name": "IndexDirections",
                "symbols": [
                    "Increasing",
                    "Decreasing"
                ]
            }
        },
        {
            "name": "mnemonic",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "description",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "uri",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "customData",
            "type": {
                "type": "map",
                "values": "DataValue"
            }
        },
        {
            "name": "scale",
            "type": "int"
        },
        {
            "name": "timeDatum",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelMetadataRecord_schema, _ChannelMetadataRecord_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "ChannelMetadataRecord",
    "fields": [
        {
            "name": "channelUri",
            "type": "string"
        },
        {
            "name": "channelId",
            "type": "long"
        },
        {
            "name": "indexes",
            "type": {
                "type": "array",
                "items": "IndexMetadataRecord"
            }
        },
        {
            "name": "channelName",
            "type": "string"
        },
        {
            "name": "dataType",
            "type": "string"
        },
        {
            "name": "uom",
            "type": "string"
        },
        {
            "name": "startIndex",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        },
        {
            "name": "endIndex",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        },
        {
            "name": "description",
            "type": "string"
        },
        {
            "name": "status",
            "type": {
                "type": "enum",
                "namespace": "Energistics.Datatypes.ChannelData",
                "name": "ChannelStatuses",
                "symbols": [
                    "Active",
                    "Inactive",
                    "Closed"
                ]
            }
        },
        {
            "name": "contentType",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "source",
            "type": "string"
        },
        {
            "name": "measureClass",
            "type": "string"
        },
        {
            "name": "uuid",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        },
        {
            "name": "customData",
            "type": {
                "type": "map",
                "values": "DataValue"
            }
        },
        {
            "name": "domainObject",
            "default": null,
            "type": [
                "null",
                "DataObject"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelMetadata_schema, _ChannelMetadata_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelDataFrame",
    "name": "ChannelMetadata",
    "fields": [
        {
            "name": "channels",
            "type": {
                "type": "array",
                "items": "ChannelMetadataRecord"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _RequestChannelData_schema, _RequestChannelData_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelDataFrame",
    "name": "RequestChannelData",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "fromIndex",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        },
        {
            "name": "toIndex",
            "default": null,
            "type": [
                "null",
                "long"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DataAttribute_schema, _DataAttribute_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "DataAttribute",
    "fields": [
        {
            "name": "attributeId",
            "type": "int"
        },
        {
            "name": "attributeValue",
            "type": "DataValue"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DataItem_schema, _DataItem_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "DataItem",
    "fields": [
        {
            "name": "indexes",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "channelId",
            "type": "long"
        },
        {
            "name": "value",
            "type": "DataValue"
        },
        {
            "name": "valueAttributes",
            "type": {
                "type": "array",
                "items": "DataAttribute"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelData_schema, _ChannelData_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelData",
    "fields": [
        {
            "name": "data",
            "type": {
                "type": "array",
                "items": "DataItem"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelDataChange_schema, _ChannelDataChange_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelDataChange",
    "fields": [
        {
            "name": "channelId",
            "type": "long"
        },
        {
            "name": "startIndex",
            "type": "long"
        },
        {
            "name": "endIndex",
            "type": "long"
        },
        {
            "name": "data",
            "type": {
                "type": "array",
                "items": "DataItem"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelDescribe_schema, _ChannelDescribe_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelDescribe",
    "fields": [
        {
            "name": "uris",
            "type": {
                "type": "array",
                "items": "string"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelRangeInfo_schema, _ChannelRangeInfo_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "ChannelRangeInfo",
    "fields": [
        {
            "name": "channelId",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "startIndex",
            "type": "long"
        },
        {
            "name": "endIndex",
            "type": "long"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelRangeRequest_schema, _ChannelRangeRequest_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelRangeRequest",
    "fields": [
        {
            "name": "channelRanges",
            "type": {
                "type": "array",
                "items": "ChannelRangeInfo"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelRemove_schema, _ChannelRemove_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelRemove",
    "fields": [
        {
            "name": "channelId",
            "type": "long"
        },
        {
            "name": "removeReason",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelStatusChange_schema, _ChannelStatusChange_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelStatusChange",
    "fields": [
        {
            "name": "channelId",
            "type": "long"
        },
        {
            "name": "status",
            "type": {
                "type": "enum",
                "namespace": "Energistics.Datatypes.ChannelData",
                "name": "ChannelStatuses",
                "symbols": [
                    "Active",
                    "Inactive",
                    "Closed"
                ]
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _StreamingStartIndex_schema, _StreamingStartIndex_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "StreamingStartIndex",
    "fields": [
        {
            "name": "item",
            "default": null,
            "type": [
                "null",
                "int",
                "long"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelStreamingInfo_schema, _ChannelStreamingInfo_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.ChannelData",
    "name": "ChannelStreamingInfo",
    "fields": [
        {
            "name": "channelId",
            "type": "long"
        },
        {
            "name": "startIndex",
            "type": "StreamingStartIndex"
        },
        {
            "name": "receiveChangeNotification",
            "type": "boolean"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelStreamingStart_schema, _ChannelStreamingStart_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelStreamingStart",
    "fields": [
        {
            "name": "channels",
            "type": {
                "type": "array",
                "items": "ChannelStreamingInfo"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChannelStreamingStop_schema, _ChannelStreamingStop_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "ChannelStreamingStop",
    "fields": [
        {
            "name": "channels",
            "type": {
                "type": "array",
                "items": "long"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Start_schema, _Start_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.ChannelStreaming",
    "name": "Start",
    "fields": [
        {
            "name": "maxMessageRate",
            "type": "int"
        },
        {
            "name": "maxDataItems",
            "type": "int"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Acknowledge_schema, _Acknowledge_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Core",
    "name": "Acknowledge",
    "fields": []
}`)

// Generated by codegen. Please do not modify.
var _CloseSession_schema, _CloseSession_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Core",
    "name": "CloseSession",
    "fields": [
        {
            "name": "reason",
            "default": null,
            "type": [
                "null",
                "string"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Version_schema, _Version_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "Version",
    "fields": [
        {
            "name": "major",
            "type": "int"
        },
        {
            "name": "minor",
            "type": "int"
        },
        {
            "name": "revision",
            "type": "int"
        },
        {
            "name": "patch",
            "type": "int"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _SupportedProtocol_schema, _SupportedProtocol_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "SupportedProtocol",
    "fields": [
        {
            "name": "protocol",
            "type": "int"
        },
        {
            "name": "protocolVersion",
            "type": "Version"
        },
        {
            "name": "role",
            "type": "string"
        },
        {
            "name": "protocolCapabilities",
            "type": {
                "type": "map",
                "values": "DataValue"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _OpenSession_schema, _OpenSession_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Core",
    "name": "OpenSession",
    "fields": [
        {
            "name": "applicationName",
            "type": "string"
        },
        {
            "name": "applicationVersion",
            "type": "string"
        },
        {
            "name": "sessionId",
            "type": "string"
        },
        {
            "name": "supportedProtocols",
            "type": {
                "type": "array",
                "items": "SupportedProtocol"
            }
        },
        {
            "name": "supportedObjects",
            "type": {
                "type": "array",
                "items": "string"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ProtocolException_schema, _ProtocolException_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Core",
    "name": "ProtocolException",
    "fields": [
        {
            "name": "errorCode",
            "type": "int"
        },
        {
            "name": "errorMessage",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _RenewSecurityToken_schema, _RenewSecurityToken_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Core",
    "name": "RenewSecurityToken",
    "fields": [
        {
            "name": "token",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _RequestSession_schema, _RequestSession_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Core",
    "name": "RequestSession",
    "fields": [
        {
            "name": "applicationName",
            "type": "string"
        },
        {
            "name": "applicationVersion",
            "type": "string"
        },
        {
            "name": "requestedProtocols",
            "type": {
                "type": "array",
                "items": "SupportedProtocol"
            }
        },
        {
            "name": "supportedObjects",
            "type": {
                "type": "array",
                "items": "string"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ArrayOfBoolean_schema, _ArrayOfBoolean_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "ArrayOfBoolean",
    "fields": [
        {
            "name": "values",
            "type": {
                "type": "array",
                "items": "boolean"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ArrayOfInt_schema, _ArrayOfInt_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "ArrayOfInt",
    "fields": [
        {
            "name": "values",
            "type": {
                "type": "array",
                "items": "int"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ArrayOfLong_schema, _ArrayOfLong_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "ArrayOfLong",
    "fields": [
        {
            "name": "values",
            "type": {
                "type": "array",
                "items": "long"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ArrayOfFloat_schema, _ArrayOfFloat_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "ArrayOfFloat",
    "fields": [
        {
            "name": "values",
            "type": {
                "type": "array",
                "items": "float"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _AnyArray_schema, _AnyArray_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes",
    "name": "AnyArray",
    "fields": [
        {
            "name": "item",
            "default": null,
            "type": [
                "null",
                "ArrayOfBoolean",
                "bytes",
                "ArrayOfInt",
                "ArrayOfLong",
                "ArrayOfFloat",
                "ArrayOfDouble"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DataArray_schema, _DataArray_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.DataArray",
    "name": "DataArray",
    "fields": [
        {
            "name": "dimensions",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "data",
            "type": "AnyArray"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GetDataArray_schema, _GetDataArray_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.DataArray",
    "name": "GetDataArray",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GetDataArraySlice_schema, _GetDataArraySlice_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.DataArray",
    "name": "GetDataArraySlice",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "start",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "count",
            "type": {
                "type": "array",
                "items": "long"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _PutDataArray_schema, _PutDataArray_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.DataArray",
    "name": "PutDataArray",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "data",
            "type": "AnyArray"
        },
        {
            "name": "dimensions",
            "type": {
                "type": "array",
                "items": "long"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _PutDataArraySlice_schema, _PutDataArraySlice_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.DataArray",
    "name": "PutDataArraySlice",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "data",
            "type": "AnyArray"
        },
        {
            "name": "dimensions",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "start",
            "type": {
                "type": "array",
                "items": "long"
            }
        },
        {
            "name": "count",
            "type": {
                "type": "array",
                "items": "long"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GetResources_schema, _GetResources_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Discovery",
    "name": "GetResources",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GetResourcesResponse_schema, _GetResourcesResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Discovery",
    "name": "GetResourcesResponse",
    "fields": [
        {
            "name": "resource",
            "type": "Resource"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GrowingObjectDelete_schema, _GrowingObjectDelete_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.GrowingObject",
    "name": "GrowingObjectDelete",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "uid",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GrowingObjectIndex_schema, _GrowingObjectIndex_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.Object",
    "name": "GrowingObjectIndex",
    "fields": [
        {
            "name": "item",
            "default": null,
            "type": [
                "null",
                "long",
                "double"
            ]
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GrowingObjectDeleteRange_schema, _GrowingObjectDeleteRange_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.GrowingObject",
    "name": "GrowingObjectDeleteRange",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "startIndex",
            "type": "GrowingObjectIndex"
        },
        {
            "name": "endIndex",
            "type": "GrowingObjectIndex"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GrowingObjectGet_schema, _GrowingObjectGet_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.GrowingObject",
    "name": "GrowingObjectGet",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "uid",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GrowingObjectGetRange_schema, _GrowingObjectGetRange_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.GrowingObject",
    "name": "GrowingObjectGetRange",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "startIndex",
            "type": "GrowingObjectIndex"
        },
        {
            "name": "endIndex",
            "type": "GrowingObjectIndex"
        },
        {
            "name": "uom",
            "type": "string"
        },
        {
            "name": "depthDatum",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GrowingObjectPut_schema, _GrowingObjectPut_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.GrowingObject",
    "name": "GrowingObjectPut",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "contentType",
            "type": "string"
        },
        {
            "name": "contentEncoding",
            "type": "string"
        },
        {
            "name": "data",
            "type": "bytes"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ObjectFragment_schema, _ObjectFragment_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.GrowingObject",
    "name": "ObjectFragment",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "contentType",
            "type": "string"
        },
        {
            "name": "contentEncoding",
            "type": "string"
        },
        {
            "name": "data",
            "type": "bytes"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DeleteObject_schema, _DeleteObject_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Store",
    "name": "DeleteObject",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _GetObject_schema, _GetObject_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Store",
    "name": "GetObject",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _Object_schema, _Object_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Store",
    "name": "Object",
    "fields": [
        {
            "name": "dataObject",
            "type": "DataObject"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _PutObject_schema, _PutObject_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.Store",
    "name": "PutObject",
    "fields": [
        {
            "name": "dataObject",
            "type": "DataObject"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _CancelNotification_schema, _CancelNotification_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.StoreNotification",
    "name": "CancelNotification",
    "fields": [
        {
            "name": "requestUuid",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ObjectChange_schema, _ObjectChange_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.Object",
    "name": "ObjectChange",
    "fields": [
        {
            "name": "changeType",
            "type": {
                "type": "enum",
                "namespace": "Energistics.Datatypes.Object",
                "name": "ObjectChangeTypes",
                "symbols": [
                    "Upsert",
                    "Delete"
                ]
            }
        },
        {
            "name": "changeTime",
            "type": "long"
        },
        {
            "name": "dataObject",
            "type": "DataObject"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _ChangeNotification_schema, _ChangeNotification_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.StoreNotification",
    "name": "ChangeNotification",
    "fields": [
        {
            "name": "change",
            "type": "ObjectChange"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _DeleteNotification_schema, _DeleteNotification_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.StoreNotification",
    "name": "DeleteNotification",
    "fields": [
        {
            "name": "delete",
            "type": "ObjectChange"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _NotificationRequestRecord_schema, _NotificationRequestRecord_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Datatypes.Object",
    "name": "NotificationRequestRecord",
    "fields": [
        {
            "name": "uri",
            "type": "string"
        },
        {
            "name": "uuid",
            "type": "string"
        },
        {
            "name": "includeObjectData",
            "type": "boolean"
        },
        {
            "name": "startTime",
            "type": "long"
        },
        {
            "name": "objectTypes",
            "type": {
                "type": "array",
                "items": "string"
            }
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _NotificationRequest_schema, _NotificationRequest_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.StoreNotification",
    "name": "NotificationRequest",
    "fields": [
        {
            "name": "request",
            "type": "NotificationRequestRecord"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMLS_AddToStore_schema, _WMLS_AddToStore_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_AddToStore",
    "fields": [
        {
            "name": "WMLtypeIn",
            "type": "string"
        },
        {
            "name": "XMLin",
            "type": "string"
        },
        {
            "name": "OptionsIn",
            "type": "string"
        },
        {
            "name": "CapabilitiesIn",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMLS_DeleteFromStore_schema, _WMLS_DeleteFromStore_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_DeleteFromStore",
    "fields": [
        {
            "name": "WMLtypeIn",
            "type": "string"
        },
        {
            "name": "XMLin",
            "type": "string"
        },
        {
            "name": "OptionsIn",
            "type": "string"
        },
        {
            "name": "CapabilitiesIn",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMLS_GetBaseMsg_schema, _WMLS_GetBaseMsg_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_GetBaseMsg",
    "fields": [
        {
            "name": "ReturnValueIn",
            "type": "int"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMLS_GetCap_schema, _WMLS_GetCap_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_GetCap",
    "fields": [
        {
            "name": "OptionsIn",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMLS_GetFromStore_schema, _WMLS_GetFromStore_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_GetFromStore",
    "fields": [
        {
            "name": "WMLtypeIn",
            "type": "string"
        },
        {
            "name": "XMLin",
            "type": "string"
        },
        {
            "name": "OptionsIn",
            "type": "string"
        },
        {
            "name": "CapabilitiesIn",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMLS_GetVersion_schema, _WMLS_GetVersion_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_GetVersion",
    "fields": []
}`)

// Generated by codegen. Please do not modify.
var _WMLS_UpdateInStore_schema, _WMLS_UpdateInStore_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMLS_UpdateInStore",
    "fields": [
        {
            "name": "WMLtypeIn",
            "type": "string"
        },
        {
            "name": "XMLin",
            "type": "string"
        },
        {
            "name": "OptionsIn",
            "type": "string"
        },
        {
            "name": "CapabilitiesIn",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_AddToStoreResponse_schema, _WMSL_AddToStoreResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_AddToStoreResponse",
    "fields": [
        {
            "name": "Result",
            "type": "int"
        },
        {
            "name": "SuppMsgOut",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_DeleteFromStoreResponse_schema, _WMSL_DeleteFromStoreResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_DeleteFromStoreResponse",
    "fields": [
        {
            "name": "Result",
            "type": "int"
        },
        {
            "name": "SuppMsgOut",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_GetBaseMsgResponse_schema, _WMSL_GetBaseMsgResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_GetBaseMsgResponse",
    "fields": [
        {
            "name": "Result",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_GetCapResponse_schema, _WMSL_GetCapResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_GetCapResponse",
    "fields": [
        {
            "name": "Result",
            "type": "int"
        },
        {
            "name": "CapabilitiesOut",
            "type": "string"
        },
        {
            "name": "SuppMsgOut",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_GetFromStoreResponse_schema, _WMSL_GetFromStoreResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_GetFromStoreResponse",
    "fields": [
        {
            "name": "Result",
            "type": "int"
        },
        {
            "name": "XMLout",
            "type": "string"
        },
        {
            "name": "SuppMsgOut",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_GetVersionResponse_schema, _WMSL_GetVersionResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_GetVersionResponse",
    "fields": [
        {
            "name": "Result",
            "type": "string"
        }
    ]
}`)

// Generated by codegen. Please do not modify.
var _WMSL_UpdateInStoreResponse_schema, _WMSL_UpdateInStoreResponse_schema_err = avro.ParseSchema(`{
    "type": "record",
    "namespace": "Energistics.Protocol.WitsmlSoap",
    "name": "WMSL_UpdateInStoreResponse",
    "fields": [
        {
            "name": "Result",
            "type": "int"
        },
        {
            "name": "SuppMsgOut",
            "type": "string"
        }
    ]
}`)
