# EnhancedRegionAttributesEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | Array of IP addresses for the edge | 
**GlobalIPs** | Pointer to **[]string** | Array of Global IP addresses for the edge | [optional] 

## Methods

### NewEnhancedRegionAttributesEdge

`func NewEnhancedRegionAttributesEdge(addresses []string, ) *EnhancedRegionAttributesEdge`

NewEnhancedRegionAttributesEdge instantiates a new EnhancedRegionAttributesEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRegionAttributesEdgeWithDefaults

`func NewEnhancedRegionAttributesEdgeWithDefaults() *EnhancedRegionAttributesEdge`

NewEnhancedRegionAttributesEdgeWithDefaults instantiates a new EnhancedRegionAttributesEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *EnhancedRegionAttributesEdge) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EnhancedRegionAttributesEdge) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EnhancedRegionAttributesEdge) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetGlobalIPs

`func (o *EnhancedRegionAttributesEdge) GetGlobalIPs() []string`

GetGlobalIPs returns the GlobalIPs field if non-nil, zero value otherwise.

### GetGlobalIPsOk

`func (o *EnhancedRegionAttributesEdge) GetGlobalIPsOk() (*[]string, bool)`

GetGlobalIPsOk returns a tuple with the GlobalIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIPs

`func (o *EnhancedRegionAttributesEdge) SetGlobalIPs(v []string)`

SetGlobalIPs sets GlobalIPs field to given value.

### HasGlobalIPs

`func (o *EnhancedRegionAttributesEdge) HasGlobalIPs() bool`

HasGlobalIPs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


