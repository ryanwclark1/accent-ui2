# TenantCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**TenantAddress**](TenantAddress.md) |  | [optional]
**Contact** | Pointer to **string** | The contact user&#39;s UUID | [optional]
**DomainNames** | Pointer to **[]string** | A list containing human readeable unique domain names, associated with a specific tenant | [optional]
**Name** | Pointer to **string** | The tenant&#39;s name | [optional]
**Phone** | Pointer to **string** | The tenant&#39;s contact phone number | [optional]
**Slug** | Pointer to **string** | A unique, human readeable identifier for this tenant. This field cannot be modified and will be auto-generated if missing. | [optional]
**Uuid** | Pointer to **string** | The tenant&#39;s UUID | [optional]

## Methods

### NewTenantCreate

`func NewTenantCreate() *TenantCreate`

NewTenantCreate instantiates a new TenantCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreateWithDefaults

`func NewTenantCreateWithDefaults() *TenantCreate`

NewTenantCreateWithDefaults instantiates a new TenantCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TenantCreate) GetAddress() TenantAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TenantCreate) GetAddressOk() (*TenantAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TenantCreate) SetAddress(v TenantAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TenantCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *TenantCreate) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TenantCreate) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TenantCreate) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *TenantCreate) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDomainNames

`func (o *TenantCreate) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *TenantCreate) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *TenantCreate) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *TenantCreate) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetName

`func (o *TenantCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *TenantCreate) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TenantCreate) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TenantCreate) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TenantCreate) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetSlug

`func (o *TenantCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TenantCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TenantCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TenantCreate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUuid

`func (o *TenantCreate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TenantCreate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TenantCreate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TenantCreate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
