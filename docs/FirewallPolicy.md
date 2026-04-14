# FirewallPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | whether the policy is enabled. | 
**Allowed** | **bool** | whether the policy is accept or drop. | 
**Id** | **string** | network policy id | 
**PolicyRules** | [**[]FirewallPolicyRule**](FirewallPolicyRule.md) | an array of policy rules. | 
**Trace** | Pointer to **bool** | whether the policy is traced. | [optional] 

## Methods

### NewFirewallPolicy

`func NewFirewallPolicy(enabled bool, allowed bool, id string, policyRules []FirewallPolicyRule, ) *FirewallPolicy`

NewFirewallPolicy instantiates a new FirewallPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallPolicyWithDefaults

`func NewFirewallPolicyWithDefaults() *FirewallPolicy`

NewFirewallPolicyWithDefaults instantiates a new FirewallPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FirewallPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FirewallPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FirewallPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowed

`func (o *FirewallPolicy) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *FirewallPolicy) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *FirewallPolicy) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetId

`func (o *FirewallPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetPolicyRules

`func (o *FirewallPolicy) GetPolicyRules() []FirewallPolicyRule`

GetPolicyRules returns the PolicyRules field if non-nil, zero value otherwise.

### GetPolicyRulesOk

`func (o *FirewallPolicy) GetPolicyRulesOk() (*[]FirewallPolicyRule, bool)`

GetPolicyRulesOk returns a tuple with the PolicyRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRules

`func (o *FirewallPolicy) SetPolicyRules(v []FirewallPolicyRule)`

SetPolicyRules sets PolicyRules field to given value.


### GetTrace

`func (o *FirewallPolicy) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *FirewallPolicy) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *FirewallPolicy) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *FirewallPolicy) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


