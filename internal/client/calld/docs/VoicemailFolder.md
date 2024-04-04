# VoicemailFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The folder&#39;s ID | [optional]
**Name** | Pointer to **string** | The folder&#39;s name | [optional]
**Type** | Pointer to **string** | The folder&#39;s type. When a message if left on a voicemail, it is stored in the folder of type \&quot;new\&quot;, unless if it is an urgent message, in which case it is left in the folder of type \&quot;urgent\&quot;. When that messages is read, it is moved into the folder of type \&quot;old\&quot;. All other folders used the type \&quot;other\&quot;. | [optional]
**Messages** | Pointer to [**[]VoicemailMessageBase**](VoicemailMessageBase.md) | The folder&#39;s messages | [optional]

## Methods

### NewVoicemailFolder

`func NewVoicemailFolder() *VoicemailFolder`

NewVoicemailFolder instantiates a new VoicemailFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailFolderWithDefaults

`func NewVoicemailFolderWithDefaults() *VoicemailFolder`

NewVoicemailFolderWithDefaults instantiates a new VoicemailFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoicemailFolder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoicemailFolder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoicemailFolder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VoicemailFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VoicemailFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoicemailFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoicemailFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoicemailFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VoicemailFolder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoicemailFolder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoicemailFolder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VoicemailFolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessages

`func (o *VoicemailFolder) GetMessages() []VoicemailMessageBase`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VoicemailFolder) GetMessagesOk() (*[]VoicemailMessageBase, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VoicemailFolder) SetMessages(v []VoicemailMessageBase)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *VoicemailFolder) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
