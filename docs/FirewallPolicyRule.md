# FirewallPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Policy rule ID. | [optional] 
**Name** | **string** | Policy rule ID. | 
**Enabled** | **bool** | wether the rule is enabled. | 
**Allowed** | **bool** | wether the rule is enabled. | 
**Sources** | [**SourcesAndDestinations**](SourcesAndDestinations.md) |  | 
**Destinations** | [**SourcesAndDestinations**](SourcesAndDestinations.md) |  | 
**Services** | Pointer to **[]string** | List of services objects IDs. | [optional] 

## Methods

### NewFirewallPolicyRule

`func NewFirewallPolicyRule(name string, enabled bool, allowed bool, sources SourcesAndDestinations, destinations SourcesAndDestinations, ) *FirewallPolicyRule`

NewFirewallPolicyRule instantiates a new FirewallPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallPolicyRuleWithDefaults

`func NewFirewallPolicyRuleWithDefaults() *FirewallPolicyRule`

NewFirewallPolicyRuleWithDefaults instantiates a new FirewallPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallPolicyRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallPolicyRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallPolicyRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallPolicyRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FirewallPolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallPolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallPolicyRule) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *FirewallPolicyRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FirewallPolicyRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FirewallPolicyRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowed

`func (o *FirewallPolicyRule) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *FirewallPolicyRule) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *FirewallPolicyRule) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetSources

`func (o *FirewallPolicyRule) GetSources() SourcesAndDestinations`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *FirewallPolicyRule) GetSourcesOk() (*SourcesAndDestinations, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *FirewallPolicyRule) SetSources(v SourcesAndDestinations)`

SetSources sets Sources field to given value.


### GetDestinations

`func (o *FirewallPolicyRule) GetDestinations() SourcesAndDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *FirewallPolicyRule) GetDestinationsOk() (*SourcesAndDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *FirewallPolicyRule) SetDestinations(v SourcesAndDestinations)`

SetDestinations sets Destinations field to given value.


### GetServices

`func (o *FirewallPolicyRule) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *FirewallPolicyRule) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *FirewallPolicyRule) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *FirewallPolicyRule) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


