# RewritingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | Pointer to **string** | The local address matching rule. Supports regular expressions | [optional]
**Replacement** | Pointer to **string** | The replacement for the matched rule | [optional]

## Methods

### NewRewritingRule

`func NewRewritingRule() *RewritingRule`

NewRewritingRule instantiates a new RewritingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewritingRuleWithDefaults

`func NewRewritingRuleWithDefaults() *RewritingRule`

NewRewritingRuleWithDefaults instantiates a new RewritingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *RewritingRule) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *RewritingRule) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *RewritingRule) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *RewritingRule) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetReplacement

`func (o *RewritingRule) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *RewritingRule) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *RewritingRule) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *RewritingRule) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
