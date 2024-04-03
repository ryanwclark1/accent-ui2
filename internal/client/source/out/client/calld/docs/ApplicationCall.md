# ApplicationCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerIdName** | Pointer to **string** |  | [optional]
**CallerIdNumber** | Pointer to **string** |  | [optional]
**CreationTime** | Pointer to **string** |  | [optional]
**DialedExtension** | Pointer to **string** |  | [optional]
**Id** | Pointer to **string** |  | [optional]
**IsCaller** | Pointer to **bool** |  | [optional]
**NodeUuid** | Pointer to **string** |  | [optional]
**OnHold** | Pointer to **bool** |  | [optional]
**Snoops** | Pointer to **map[string]interface{}** |  | [optional]
**Status** | Pointer to **string** | *Down: the call is not connected to anything yet* Ring: the call just came in the application *Progress: the call is playing a progress ringing tone* Up: the call is answered and media can be sent/received * Ringing: the call is ringing the phone  | [optional]
**Variables** | Pointer to **map[string]interface{}** |  | [optional]

## Methods

### NewApplicationCall

`func NewApplicationCall() *ApplicationCall`

NewApplicationCall instantiates a new ApplicationCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCallWithDefaults

`func NewApplicationCallWithDefaults() *ApplicationCall`

NewApplicationCallWithDefaults instantiates a new ApplicationCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerIdName

`func (o *ApplicationCall) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *ApplicationCall) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *ApplicationCall) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *ApplicationCall) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNumber

`func (o *ApplicationCall) GetCallerIdNumber() string`

GetCallerIdNumber returns the CallerIdNumber field if non-nil, zero value otherwise.

### GetCallerIdNumberOk

`func (o *ApplicationCall) GetCallerIdNumberOk() (*string, bool)`

GetCallerIdNumberOk returns a tuple with the CallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNumber

`func (o *ApplicationCall) SetCallerIdNumber(v string)`

SetCallerIdNumber sets CallerIdNumber field to given value.

### HasCallerIdNumber

`func (o *ApplicationCall) HasCallerIdNumber() bool`

HasCallerIdNumber returns a boolean if a field has been set.

### GetCreationTime

`func (o *ApplicationCall) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ApplicationCall) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ApplicationCall) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ApplicationCall) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDialedExtension

`func (o *ApplicationCall) GetDialedExtension() string`

GetDialedExtension returns the DialedExtension field if non-nil, zero value otherwise.

### GetDialedExtensionOk

`func (o *ApplicationCall) GetDialedExtensionOk() (*string, bool)`

GetDialedExtensionOk returns a tuple with the DialedExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialedExtension

`func (o *ApplicationCall) SetDialedExtension(v string)`

SetDialedExtension sets DialedExtension field to given value.

### HasDialedExtension

`func (o *ApplicationCall) HasDialedExtension() bool`

HasDialedExtension returns a boolean if a field has been set.

### GetId

`func (o *ApplicationCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCaller

`func (o *ApplicationCall) GetIsCaller() bool`

GetIsCaller returns the IsCaller field if non-nil, zero value otherwise.

### GetIsCallerOk

`func (o *ApplicationCall) GetIsCallerOk() (*bool, bool)`

GetIsCallerOk returns a tuple with the IsCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaller

`func (o *ApplicationCall) SetIsCaller(v bool)`

SetIsCaller sets IsCaller field to given value.

### HasIsCaller

`func (o *ApplicationCall) HasIsCaller() bool`

HasIsCaller returns a boolean if a field has been set.

### GetNodeUuid

`func (o *ApplicationCall) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *ApplicationCall) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *ApplicationCall) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *ApplicationCall) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetOnHold

`func (o *ApplicationCall) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *ApplicationCall) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *ApplicationCall) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *ApplicationCall) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetSnoops

`func (o *ApplicationCall) GetSnoops() map[string]interface{}`

GetSnoops returns the Snoops field if non-nil, zero value otherwise.

### GetSnoopsOk

`func (o *ApplicationCall) GetSnoopsOk() (*map[string]interface{}, bool)`

GetSnoopsOk returns a tuple with the Snoops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoops

`func (o *ApplicationCall) SetSnoops(v map[string]interface{})`

SetSnoops sets Snoops field to given value.

### HasSnoops

`func (o *ApplicationCall) HasSnoops() bool`

HasSnoops returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationCall) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationCall) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationCall) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationCall) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariables

`func (o *ApplicationCall) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ApplicationCall) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ApplicationCall) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ApplicationCall) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
