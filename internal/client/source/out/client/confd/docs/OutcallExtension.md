# OutcallExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerId** | Pointer to **string** |  | [optional]
**ExternalPrefix** | Pointer to **string** | The prefix added to the outgoing call when sent on the trunk | [optional]
**Prefix** | Pointer to **string** | The prefix that the user need to dial before the pattern | [optional]
**StripDigits** | Pointer to **int32** | The number of leading digits to strip | [optional]

## Methods

### NewOutcallExtension

`func NewOutcallExtension() *OutcallExtension`

NewOutcallExtension instantiates a new OutcallExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcallExtensionWithDefaults

`func NewOutcallExtensionWithDefaults() *OutcallExtension`

NewOutcallExtensionWithDefaults instantiates a new OutcallExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerId

`func (o *OutcallExtension) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *OutcallExtension) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *OutcallExtension) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *OutcallExtension) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetExternalPrefix

`func (o *OutcallExtension) GetExternalPrefix() string`

GetExternalPrefix returns the ExternalPrefix field if non-nil, zero value otherwise.

### GetExternalPrefixOk

`func (o *OutcallExtension) GetExternalPrefixOk() (*string, bool)`

GetExternalPrefixOk returns a tuple with the ExternalPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrefix

`func (o *OutcallExtension) SetExternalPrefix(v string)`

SetExternalPrefix sets ExternalPrefix field to given value.

### HasExternalPrefix

`func (o *OutcallExtension) HasExternalPrefix() bool`

HasExternalPrefix returns a boolean if a field has been set.

### GetPrefix

`func (o *OutcallExtension) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OutcallExtension) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OutcallExtension) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OutcallExtension) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetStripDigits

`func (o *OutcallExtension) GetStripDigits() int32`

GetStripDigits returns the StripDigits field if non-nil, zero value otherwise.

### GetStripDigitsOk

`func (o *OutcallExtension) GetStripDigitsOk() (*int32, bool)`

GetStripDigitsOk returns a tuple with the StripDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripDigits

`func (o *OutcallExtension) SetStripDigits(v int32)`

SetStripDigits sets StripDigits field to given value.

### HasStripDigits

`func (o *OutcallExtension) HasStripDigits() bool`

HasStripDigits returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
