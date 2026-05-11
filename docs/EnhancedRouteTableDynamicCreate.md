# EnhancedRouteTableDynamicCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelIds** | **[]string** | All IDs of dynamic tunnels | 
**Subnets** | **[]string** | List of subnet CIDR blocks | 

## Methods

### NewEnhancedRouteTableDynamicCreate

`func NewEnhancedRouteTableDynamicCreate(tunnelIds []string, subnets []string, ) *EnhancedRouteTableDynamicCreate`

NewEnhancedRouteTableDynamicCreate instantiates a new EnhancedRouteTableDynamicCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRouteTableDynamicCreateWithDefaults

`func NewEnhancedRouteTableDynamicCreateWithDefaults() *EnhancedRouteTableDynamicCreate`

NewEnhancedRouteTableDynamicCreateWithDefaults instantiates a new EnhancedRouteTableDynamicCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnelIds

`func (o *EnhancedRouteTableDynamicCreate) GetTunnelIds() []string`

GetTunnelIds returns the TunnelIds field if non-nil, zero value otherwise.

### GetTunnelIdsOk

`func (o *EnhancedRouteTableDynamicCreate) GetTunnelIdsOk() (*[]string, bool)`

GetTunnelIdsOk returns a tuple with the TunnelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIds

`func (o *EnhancedRouteTableDynamicCreate) SetTunnelIds(v []string)`

SetTunnelIds sets TunnelIds field to given value.


### GetSubnets

`func (o *EnhancedRouteTableDynamicCreate) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *EnhancedRouteTableDynamicCreate) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *EnhancedRouteTableDynamicCreate) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


