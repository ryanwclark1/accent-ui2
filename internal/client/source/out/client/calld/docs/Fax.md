# Fax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | Pointer to **string** | The ID of the call that produced this fax | [optional]
**CallerId** | Pointer to **string** | The Caller ID that was presented to the fax recipient | [optional]
**Context** | Pointer to **string** | The context where the fax was sent | [optional]
**Extension** | Pointer to **string** | The extension where the fax was sent | [optional]
**Id** | Pointer to **string** | The fax ID | [optional]
**IvrExtension** | Pointer to **string** | Extension to compose before sending fax. Useful for fax in IVR | [optional]
**TenantUuid** | Pointer to **string** | The tenant UUID where the fax was sent from | [optional]
**UserUuid** | Pointer to **string** | The UUID of the user that sent the fax. May be null if the fax was sent by another service. | [optional]
**WaitTime** | Pointer to **string** | Time waiting before sending fax when destination has answered (in seconds) | [optional]

## Methods

### NewFax

`func NewFax() *Fax`

NewFax instantiates a new Fax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxWithDefaults

`func NewFaxWithDefaults() *Fax`

NewFaxWithDefaults instantiates a new Fax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *Fax) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Fax) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Fax) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *Fax) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallerId

`func (o *Fax) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *Fax) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *Fax) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *Fax) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetContext

`func (o *Fax) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Fax) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Fax) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Fax) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExtension

`func (o *Fax) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Fax) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Fax) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Fax) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetId

`func (o *Fax) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fax) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fax) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Fax) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIvrExtension

`func (o *Fax) GetIvrExtension() string`

GetIvrExtension returns the IvrExtension field if non-nil, zero value otherwise.

### GetIvrExtensionOk

`func (o *Fax) GetIvrExtensionOk() (*string, bool)`

GetIvrExtensionOk returns a tuple with the IvrExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvrExtension

`func (o *Fax) SetIvrExtension(v string)`

SetIvrExtension sets IvrExtension field to given value.

### HasIvrExtension

`func (o *Fax) HasIvrExtension() bool`

HasIvrExtension returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Fax) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Fax) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Fax) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Fax) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUserUuid

`func (o *Fax) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *Fax) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *Fax) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *Fax) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetWaitTime

`func (o *Fax) GetWaitTime() string`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *Fax) GetWaitTimeOk() (*string, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *Fax) SetWaitTime(v string)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *Fax) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
