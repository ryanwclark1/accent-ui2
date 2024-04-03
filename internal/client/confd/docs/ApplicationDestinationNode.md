# ApplicationDestinationNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **bool** | Automatically answer the call when the call enters the destination node | [optional]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional]
**Type** | Pointer to **string** | type of the default node | [optional]

## Methods

### NewApplicationDestinationNode

`func NewApplicationDestinationNode() *ApplicationDestinationNode`

NewApplicationDestinationNode instantiates a new ApplicationDestinationNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDestinationNodeWithDefaults

`func NewApplicationDestinationNodeWithDefaults() *ApplicationDestinationNode`

NewApplicationDestinationNodeWithDefaults instantiates a new ApplicationDestinationNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *ApplicationDestinationNode) GetAnswer() bool`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *ApplicationDestinationNode) GetAnswerOk() (*bool, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *ApplicationDestinationNode) SetAnswer(v bool)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *ApplicationDestinationNode) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *ApplicationDestinationNode) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *ApplicationDestinationNode) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *ApplicationDestinationNode) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *ApplicationDestinationNode) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetType

`func (o *ApplicationDestinationNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationDestinationNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationDestinationNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationDestinationNode) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
