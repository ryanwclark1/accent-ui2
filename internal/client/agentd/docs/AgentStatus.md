# AgentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** | The context on which the agent is logged, or null if the agent is not logged | [optional]
**Extension** | Pointer to **string** | The extension on which the agent is logged, or null if the agent is not logged | [optional]
**Id** | Pointer to **int32** | Agent&#39;s ID | [optional]
**Logged** | Pointer to **bool** | True if the agent is logged, else false | [optional]
**Number** | Pointer to **string** | Agent&#39;s number | [optional]
**OriginUuid** | Pointer to **string** | Accent&#39;s UUID | [optional]
**Paused** | Pointer to **bool** | True if the agent is paused, else false | [optional]
**PausedReason** | Pointer to **string** | The reason of the agent pause | [optional]
**StateInterface** | Pointer to **string** | The interface (e.g. SIP/alice) to determine if the agent is in use or not | [optional]
**TenantUuid** | Pointer to **string** | Tenant&#39;s UUID | [optional]

## Methods

### NewAgentStatus

`func NewAgentStatus() *AgentStatus`

NewAgentStatus instantiates a new AgentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentStatusWithDefaults

`func NewAgentStatusWithDefaults() *AgentStatus`

NewAgentStatusWithDefaults instantiates a new AgentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AgentStatus) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AgentStatus) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AgentStatus) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AgentStatus) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExtension

`func (o *AgentStatus) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *AgentStatus) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *AgentStatus) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *AgentStatus) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetId

`func (o *AgentStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgentStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogged

`func (o *AgentStatus) GetLogged() bool`

GetLogged returns the Logged field if non-nil, zero value otherwise.

### GetLoggedOk

`func (o *AgentStatus) GetLoggedOk() (*bool, bool)`

GetLoggedOk returns a tuple with the Logged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogged

`func (o *AgentStatus) SetLogged(v bool)`

SetLogged sets Logged field to given value.

### HasLogged

`func (o *AgentStatus) HasLogged() bool`

HasLogged returns a boolean if a field has been set.

### GetNumber

`func (o *AgentStatus) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AgentStatus) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AgentStatus) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AgentStatus) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetOriginUuid

`func (o *AgentStatus) GetOriginUuid() string`

GetOriginUuid returns the OriginUuid field if non-nil, zero value otherwise.

### GetOriginUuidOk

`func (o *AgentStatus) GetOriginUuidOk() (*string, bool)`

GetOriginUuidOk returns a tuple with the OriginUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginUuid

`func (o *AgentStatus) SetOriginUuid(v string)`

SetOriginUuid sets OriginUuid field to given value.

### HasOriginUuid

`func (o *AgentStatus) HasOriginUuid() bool`

HasOriginUuid returns a boolean if a field has been set.

### GetPaused

`func (o *AgentStatus) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *AgentStatus) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *AgentStatus) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *AgentStatus) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPausedReason

`func (o *AgentStatus) GetPausedReason() string`

GetPausedReason returns the PausedReason field if non-nil, zero value otherwise.

### GetPausedReasonOk

`func (o *AgentStatus) GetPausedReasonOk() (*string, bool)`

GetPausedReasonOk returns a tuple with the PausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedReason

`func (o *AgentStatus) SetPausedReason(v string)`

SetPausedReason sets PausedReason field to given value.

### HasPausedReason

`func (o *AgentStatus) HasPausedReason() bool`

HasPausedReason returns a boolean if a field has been set.

### GetStateInterface

`func (o *AgentStatus) GetStateInterface() string`

GetStateInterface returns the StateInterface field if non-nil, zero value otherwise.

### GetStateInterfaceOk

`func (o *AgentStatus) GetStateInterfaceOk() (*string, bool)`

GetStateInterfaceOk returns a tuple with the StateInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInterface

`func (o *AgentStatus) SetStateInterface(v string)`

SetStateInterface sets StateInterface field to given value.

### HasStateInterface

`func (o *AgentStatus) HasStateInterface() bool`

HasStateInterface returns a boolean if a field has been set.

### GetTenantUuid

`func (o *AgentStatus) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *AgentStatus) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *AgentStatus) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *AgentStatus) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
