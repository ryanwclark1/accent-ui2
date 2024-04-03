# TenantEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**TenantAddress**](TenantAddress.md) |  | [optional]
**Contact** | Pointer to **string** | The contact user&#39;s UUID | [optional]
**DomainNames** | Pointer to **[]string** | A list containing human readeable unique domain names, associated with a specific tenant | [optional]
**Name** | Pointer to **string** | The tenant&#39;s name | [optional]
**Phone** | Pointer to **string** | The tenant&#39;s contact phone number | [optional]

## Methods

### NewTenantEdit

`func NewTenantEdit() *TenantEdit`

NewTenantEdit instantiates a new TenantEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantEditWithDefaults

`func NewTenantEditWithDefaults() *TenantEdit`

NewTenantEditWithDefaults instantiates a new TenantEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TenantEdit) GetAddress() TenantAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TenantEdit) GetAddressOk() (*TenantAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TenantEdit) SetAddress(v TenantAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TenantEdit) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *TenantEdit) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TenantEdit) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TenantEdit) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *TenantEdit) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDomainNames

`func (o *TenantEdit) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *TenantEdit) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *TenantEdit) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *TenantEdit) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetName

`func (o *TenantEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *TenantEdit) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TenantEdit) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TenantEdit) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TenantEdit) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
