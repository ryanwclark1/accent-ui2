# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **map[string]map[string]string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Dsize** | Pointer to **int32** |  | [optional]
**Sha1sum** | Pointer to **string** |  | [optional]
**Version** | Pointer to **string** |  | [optional]

## Methods

### NewPackage

`func NewPackage() *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *Package) GetCapabilities() map[string]map[string]string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Package) GetCapabilitiesOk() (*map[string]map[string]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Package) SetCapabilities(v map[string]map[string]string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Package) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDescription

`func (o *Package) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Package) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Package) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Package) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDsize

`func (o *Package) GetDsize() int32`

GetDsize returns the Dsize field if non-nil, zero value otherwise.

### GetDsizeOk

`func (o *Package) GetDsizeOk() (*int32, bool)`

GetDsizeOk returns a tuple with the Dsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsize

`func (o *Package) SetDsize(v int32)`

SetDsize sets Dsize field to given value.

### HasDsize

`func (o *Package) HasDsize() bool`

HasDsize returns a boolean if a field has been set.

### GetSha1sum

`func (o *Package) GetSha1sum() string`

GetSha1sum returns the Sha1sum field if non-nil, zero value otherwise.

### GetSha1sumOk

`func (o *Package) GetSha1sumOk() (*string, bool)`

GetSha1sumOk returns a tuple with the Sha1sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1sum

`func (o *Package) SetSha1sum(v string)`

SetSha1sum sets Sha1sum field to given value.

### HasSha1sum

`func (o *Package) HasSha1sum() bool`

HasSha1sum returns a boolean if a field has been set.

### GetVersion

`func (o *Package) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Package) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Package) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Package) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
