# EnhancedRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Route table entry ID | 
**TunnelIds** | **[]string** | List of tunnels ID | 
**Subnets** | **[]string** | List of subnet CIDR blocks | 
**Propagated** | **bool** | Whether the route is propagated automatically | [default to false]

## Methods

### NewEnhancedRouteTable

`func NewEnhancedRouteTable(id string, tunnelIds []string, subnets []string, propagated bool, ) *EnhancedRouteTable`

NewEnhancedRouteTable instantiates a new EnhancedRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRouteTableWithDefaults

`func NewEnhancedRouteTableWithDefaults() *EnhancedRouteTable`

NewEnhancedRouteTableWithDefaults instantiates a new EnhancedRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnhancedRouteTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedRouteTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedRouteTable) SetId(v string)`

SetId sets Id field to given value.


### GetTunnelIds

`func (o *EnhancedRouteTable) GetTunnelIds() []string`

GetTunnelIds returns the TunnelIds field if non-nil, zero value otherwise.

### GetTunnelIdsOk

`func (o *EnhancedRouteTable) GetTunnelIdsOk() (*[]string, bool)`

GetTunnelIdsOk returns a tuple with the TunnelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelIds

`func (o *EnhancedRouteTable) SetTunnelIds(v []string)`

SetTunnelIds sets TunnelIds field to given value.


### GetSubnets

`func (o *EnhancedRouteTable) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *EnhancedRouteTable) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *EnhancedRouteTable) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.


### GetPropagated

`func (o *EnhancedRouteTable) GetPropagated() bool`

GetPropagated returns the Propagated field if non-nil, zero value otherwise.

### GetPropagatedOk

`func (o *EnhancedRouteTable) GetPropagatedOk() (*bool, bool)`

GetPropagatedOk returns a tuple with the Propagated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagated

`func (o *EnhancedRouteTable) SetPropagated(v bool)`

SetPropagated sets Propagated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


