# PackageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pkgs** | Pointer to [**map[string]Package**](Package.md) |  | [optional]

## Methods

### NewPackageList

`func NewPackageList() *PackageList`

NewPackageList instantiates a new PackageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageListWithDefaults

`func NewPackageListWithDefaults() *PackageList`

NewPackageListWithDefaults instantiates a new PackageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkgs

`func (o *PackageList) GetPkgs() map[string]Package`

GetPkgs returns the Pkgs field if non-nil, zero value otherwise.

### GetPkgsOk

`func (o *PackageList) GetPkgsOk() (*map[string]Package, bool)`

GetPkgsOk returns a tuple with the Pkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgs

`func (o *PackageList) SetPkgs(v map[string]Package)`

SetPkgs sets Pkgs field to given value.

### HasPkgs

`func (o *PackageList) HasPkgs() bool`

HasPkgs returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
