# AgentRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** |  | [optional]
**Id** | Pointer to **int32** | The id of the agent | [optional] [readonly]
**Lastname** | Pointer to **string** |  | [optional]
**Number** | Pointer to **string** | Agent number. Cannot be modified after creation | [optional]

## Methods

### NewAgentRelationBase

`func NewAgentRelationBase() *AgentRelationBase`

NewAgentRelationBase instantiates a new AgentRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRelationBaseWithDefaults

`func NewAgentRelationBaseWithDefaults() *AgentRelationBase`

NewAgentRelationBaseWithDefaults instantiates a new AgentRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *AgentRelationBase) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *AgentRelationBase) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *AgentRelationBase) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *AgentRelationBase) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetId

`func (o *AgentRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgentRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastname

`func (o *AgentRelationBase) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *AgentRelationBase) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *AgentRelationBase) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *AgentRelationBase) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetNumber

`func (o *AgentRelationBase) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AgentRelationBase) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AgentRelationBase) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AgentRelationBase) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
