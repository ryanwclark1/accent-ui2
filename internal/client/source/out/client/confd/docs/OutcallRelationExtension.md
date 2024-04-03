# OutcallRelationExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional]
**Exten** | Pointer to **string** |  | [optional]
**Id** | Pointer to **int32** | Extension ID | [optional] [readonly]
**CallerId** | Pointer to **string** |  | [optional]
**ExternalPrefix** | Pointer to **string** | The prefix added to the outgoing call when sent on the trunk | [optional]
**Prefix** | Pointer to **string** | The prefix that the user need to dial before the extension | [optional]
**StripDigits** | Pointer to **int32** | The number of leading digits to strip | [optional]

## Methods

### NewOutcallRelationExtension

`func NewOutcallRelationExtension() *OutcallRelationExtension`

NewOutcallRelationExtension instantiates a new OutcallRelationExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcallRelationExtensionWithDefaults

`func NewOutcallRelationExtensionWithDefaults() *OutcallRelationExtension`

NewOutcallRelationExtensionWithDefaults instantiates a new OutcallRelationExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OutcallRelationExtension) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OutcallRelationExtension) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OutcallRelationExtension) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *OutcallRelationExtension) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExten

`func (o *OutcallRelationExtension) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *OutcallRelationExtension) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *OutcallRelationExtension) SetExten(v string)`

SetExten sets Exten field to given value.

### HasExten

`func (o *OutcallRelationExtension) HasExten() bool`

HasExten returns a boolean if a field has been set.

### GetId

`func (o *OutcallRelationExtension) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutcallRelationExtension) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutcallRelationExtension) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OutcallRelationExtension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCallerId

`func (o *OutcallRelationExtension) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *OutcallRelationExtension) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *OutcallRelationExtension) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *OutcallRelationExtension) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetExternalPrefix

`func (o *OutcallRelationExtension) GetExternalPrefix() string`

GetExternalPrefix returns the ExternalPrefix field if non-nil, zero value otherwise.

### GetExternalPrefixOk

`func (o *OutcallRelationExtension) GetExternalPrefixOk() (*string, bool)`

GetExternalPrefixOk returns a tuple with the ExternalPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrefix

`func (o *OutcallRelationExtension) SetExternalPrefix(v string)`

SetExternalPrefix sets ExternalPrefix field to given value.

### HasExternalPrefix

`func (o *OutcallRelationExtension) HasExternalPrefix() bool`

HasExternalPrefix returns a boolean if a field has been set.

### GetPrefix

`func (o *OutcallRelationExtension) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OutcallRelationExtension) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OutcallRelationExtension) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OutcallRelationExtension) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetStripDigits

`func (o *OutcallRelationExtension) GetStripDigits() int32`

GetStripDigits returns the StripDigits field if non-nil, zero value otherwise.

### GetStripDigitsOk

`func (o *OutcallRelationExtension) GetStripDigitsOk() (*int32, bool)`

GetStripDigitsOk returns a tuple with the StripDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripDigits

`func (o *OutcallRelationExtension) SetStripDigits(v int32)`

SetStripDigits sets StripDigits field to given value.

### HasStripDigits

`func (o *OutcallRelationExtension) HasStripDigits() bool`

HasStripDigits returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
