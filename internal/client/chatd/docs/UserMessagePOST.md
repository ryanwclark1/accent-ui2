# UserMessagePOST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Alias/nickname of the sender |
**Content** | **string** | The content of the message |

## Methods

### NewUserMessagePOST

`func NewUserMessagePOST(alias string, content string, ) *UserMessagePOST`

NewUserMessagePOST instantiates a new UserMessagePOST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMessagePOSTWithDefaults

`func NewUserMessagePOSTWithDefaults() *UserMessagePOST`

NewUserMessagePOSTWithDefaults instantiates a new UserMessagePOST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *UserMessagePOST) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *UserMessagePOST) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *UserMessagePOST) SetAlias(v string)`

SetAlias sets Alias field to given value.

### GetContent

`func (o *UserMessagePOST) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UserMessagePOST) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UserMessagePOST) SetContent(v string)`

SetContent sets Content field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
