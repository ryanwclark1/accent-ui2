# Nat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Param** | Pointer to **string** | Set to &#x60;1&#x60; if all the devices are behind a NAT | [optional]

## Methods

### NewNat

`func NewNat() *Nat`

NewNat instantiates a new Nat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatWithDefaults

`func NewNatWithDefaults() *Nat`

NewNatWithDefaults instantiates a new Nat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParam

`func (o *Nat) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *Nat) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *Nat) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *Nat) HasParam() bool`

HasParam returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
