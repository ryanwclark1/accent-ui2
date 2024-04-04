# CallFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to [**CallFilterRecipientUsers**](CallFilterRecipientUsers.md) |  | [optional]
**Surrogates** | Pointer to [**CallFilterSurrogateUsers**](CallFilterSurrogateUsers.md) |  | [optional]
**CallerIdMode** | Pointer to **string** | How the caller_id_name will be treated | [optional]
**CallerIdName** | Pointer to **string** | Name to display | [optional]
**Description** | Pointer to **string** | Additional information about the call filter | [optional]
**Enabled** | Pointer to **bool** | Disable or enable the call filter | [optional] [default to true]
**Id** | Pointer to **int32** | The id of the call filter | [optional] [readonly]
**Name** | **string** | The name of the call filter |
**Source** | **string** | Call type to apply call filter |
**Strategy** | **string** | Determine which will ring. |
**SurrogatesTimeout** | Pointer to **int32** | Number of seconds the filter&#39;s surrogates will ring before falling back | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewCallFilter

`func NewCallFilter(name string, source string, strategy string, ) *CallFilter`

NewCallFilter instantiates a new CallFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallFilterWithDefaults

`func NewCallFilterWithDefaults() *CallFilter`

NewCallFilterWithDefaults instantiates a new CallFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *CallFilter) GetRecipients() CallFilterRecipientUsers`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CallFilter) GetRecipientsOk() (*CallFilterRecipientUsers, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CallFilter) SetRecipients(v CallFilterRecipientUsers)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CallFilter) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSurrogates

`func (o *CallFilter) GetSurrogates() CallFilterSurrogateUsers`

GetSurrogates returns the Surrogates field if non-nil, zero value otherwise.

### GetSurrogatesOk

`func (o *CallFilter) GetSurrogatesOk() (*CallFilterSurrogateUsers, bool)`

GetSurrogatesOk returns a tuple with the Surrogates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurrogates

`func (o *CallFilter) SetSurrogates(v CallFilterSurrogateUsers)`

SetSurrogates sets Surrogates field to given value.

### HasSurrogates

`func (o *CallFilter) HasSurrogates() bool`

HasSurrogates returns a boolean if a field has been set.

### GetCallerIdMode

`func (o *CallFilter) GetCallerIdMode() string`

GetCallerIdMode returns the CallerIdMode field if non-nil, zero value otherwise.

### GetCallerIdModeOk

`func (o *CallFilter) GetCallerIdModeOk() (*string, bool)`

GetCallerIdModeOk returns a tuple with the CallerIdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdMode

`func (o *CallFilter) SetCallerIdMode(v string)`

SetCallerIdMode sets CallerIdMode field to given value.

### HasCallerIdMode

`func (o *CallFilter) HasCallerIdMode() bool`

HasCallerIdMode returns a boolean if a field has been set.

### GetCallerIdName

`func (o *CallFilter) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *CallFilter) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *CallFilter) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *CallFilter) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetDescription

`func (o *CallFilter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CallFilter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CallFilter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CallFilter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CallFilter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CallFilter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CallFilter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CallFilter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *CallFilter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallFilter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallFilter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CallFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CallFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallFilter) SetName(v string)`

SetName sets Name field to given value.

### GetSource

`func (o *CallFilter) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CallFilter) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CallFilter) SetSource(v string)`

SetSource sets Source field to given value.

### GetStrategy

`func (o *CallFilter) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *CallFilter) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *CallFilter) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### GetSurrogatesTimeout

`func (o *CallFilter) GetSurrogatesTimeout() int32`

GetSurrogatesTimeout returns the SurrogatesTimeout field if non-nil, zero value otherwise.

### GetSurrogatesTimeoutOk

`func (o *CallFilter) GetSurrogatesTimeoutOk() (*int32, bool)`

GetSurrogatesTimeoutOk returns a tuple with the SurrogatesTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurrogatesTimeout

`func (o *CallFilter) SetSurrogatesTimeout(v int32)`

SetSurrogatesTimeout sets SurrogatesTimeout field to given value.

### HasSurrogatesTimeout

`func (o *CallFilter) HasSurrogatesTimeout() bool`

HasSurrogatesTimeout returns a boolean if a field has been set.

### GetTenantUuid

`func (o *CallFilter) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *CallFilter) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *CallFilter) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *CallFilter) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
