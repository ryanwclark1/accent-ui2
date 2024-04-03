# IAXGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **map[string]interface{}** | General IAX options. These options must be unique and unordered. Otherwise, use &#x60;ordered_options&#x60;. Option must have the following form:  &#x60;&#x60;&#x60; {   \&quot;options\&quot;: {     \&quot;name1\&quot;: \&quot;value1\&quot;,     ...   } } &#x60;&#x60;&#x60;   | [optional]
**OrderedOptions** | Pointer to **[][]string** | Any options may be repeated as many times and ordered as needed. Ordered options must have the following form:  &#x60;&#x60;&#x60; {   \&quot;ordered_options\&quot;: [     [\&quot;name1\&quot;, \&quot;value1\&quot;],     [\&quot;name2\&quot;, \&quot;value2\&quot;]   ] } &#x60;&#x60;&#x60;  The resulting configuration in iax.conf will have the following form:  &#x60;&#x60;&#x60; [general] name1&#x3D;value1 name2&#x3D;value2 &#x60;&#x60;&#x60;  | [optional]

## Methods

### NewIAXGeneral

`func NewIAXGeneral() *IAXGeneral`

NewIAXGeneral instantiates a new IAXGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAXGeneralWithDefaults

`func NewIAXGeneralWithDefaults() *IAXGeneral`

NewIAXGeneralWithDefaults instantiates a new IAXGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *IAXGeneral) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IAXGeneral) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IAXGeneral) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IAXGeneral) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOrderedOptions

`func (o *IAXGeneral) GetOrderedOptions() [][]string`

GetOrderedOptions returns the OrderedOptions field if non-nil, zero value otherwise.

### GetOrderedOptionsOk

`func (o *IAXGeneral) GetOrderedOptionsOk() (*[][]string, bool)`

GetOrderedOptionsOk returns a tuple with the OrderedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedOptions

`func (o *IAXGeneral) SetOrderedOptions(v [][]string)`

SetOrderedOptions sets OrderedOptions field to given value.

### HasOrderedOptions

`func (o *IAXGeneral) HasOrderedOptions() bool`

HasOrderedOptions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
