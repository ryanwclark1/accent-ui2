# EndpointSccp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | SCCP Endpoint ID | [optional] [readonly]
**Trunk** | Pointer to [**TrunkRelationBase**](TrunkRelationBase.md) |  | [optional]
**Line** | Pointer to [**[]LineRelationBase**](LineRelationBase.md) |  | [optional] [readonly]
**Options** | Pointer to **[][]string** | Advanced configuration options. Options are appended at the end of a  SCCP account in the file &#39;sccp.conf&#39; used by asterisk. Please consult the asterisk documentation for further details on available parameters. Because of database limitations, only the following options are allowed:   *cid_name* cid_num  *allow* disallow   Options must have the following the form:  &#x60;&#x60;&#x60; {   \&quot;options\&quot;: [     [\&quot;name1\&quot;, \&quot;value1\&quot;],     [\&quot;name2\&quot;, \&quot;value2\&quot;]   ] } &#x60;&#x60;&#x60;  The resulting configuration in sip.conf will have the following form:  &#x60;&#x60;&#x60; [1000] name1&#x3D;value1 name2&#x3D;value2 &#x60;&#x60;&#x60;  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewEndpointSccp

`func NewEndpointSccp() *EndpointSccp`

NewEndpointSccp instantiates a new EndpointSccp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointSccpWithDefaults

`func NewEndpointSccpWithDefaults() *EndpointSccp`

NewEndpointSccpWithDefaults instantiates a new EndpointSccp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointSccp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointSccp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointSccp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointSccp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrunk

`func (o *EndpointSccp) GetTrunk() TrunkRelationBase`

GetTrunk returns the Trunk field if non-nil, zero value otherwise.

### GetTrunkOk

`func (o *EndpointSccp) GetTrunkOk() (*TrunkRelationBase, bool)`

GetTrunkOk returns a tuple with the Trunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunk

`func (o *EndpointSccp) SetTrunk(v TrunkRelationBase)`

SetTrunk sets Trunk field to given value.

### HasTrunk

`func (o *EndpointSccp) HasTrunk() bool`

HasTrunk returns a boolean if a field has been set.

### GetLine

`func (o *EndpointSccp) GetLine() []LineRelationBase`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *EndpointSccp) GetLineOk() (*[]LineRelationBase, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *EndpointSccp) SetLine(v []LineRelationBase)`

SetLine sets Line field to given value.

### HasLine

`func (o *EndpointSccp) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetOptions

`func (o *EndpointSccp) GetOptions() [][]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EndpointSccp) GetOptionsOk() (*[][]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EndpointSccp) SetOptions(v [][]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EndpointSccp) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTenantUuid

`func (o *EndpointSccp) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *EndpointSccp) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *EndpointSccp) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *EndpointSccp) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
