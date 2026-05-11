# EnhancedRouteTableStaticCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelId** | **string** | Static tunnel ID | 
**Subnets** | **[]string** | List of subnet CIDR blocks | 

## Methods

### NewEnhancedRouteTableStaticCreate

`func NewEnhancedRouteTableStaticCreate(tunnelId string, subnets []string, ) *EnhancedRouteTableStaticCreate`

NewEnhancedRouteTableStaticCreate instantiates a new EnhancedRouteTableStaticCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRouteTableStaticCreateWithDefaults

`func NewEnhancedRouteTableStaticCreateWithDefaults() *EnhancedRouteTableStaticCreate`

NewEnhancedRouteTableStaticCreateWithDefaults instantiates a new EnhancedRouteTableStaticCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnelId

`func (o *EnhancedRouteTableStaticCreate) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *EnhancedRouteTableStaticCreate) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *EnhancedRouteTableStaticCreate) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.


### GetSubnets

`func (o *EnhancedRouteTableStaticCreate) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *EnhancedRouteTableStaticCreate) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *EnhancedRouteTableStaticCreate) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


