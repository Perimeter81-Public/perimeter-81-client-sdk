# RouteTableUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | **[]string** | List of subnet CIDR blocks | 

## Methods

### NewRouteTableUpdate

`func NewRouteTableUpdate(subnets []string, ) *RouteTableUpdate`

NewRouteTableUpdate instantiates a new RouteTableUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableUpdateWithDefaults

`func NewRouteTableUpdateWithDefaults() *RouteTableUpdate`

NewRouteTableUpdateWithDefaults instantiates a new RouteTableUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *RouteTableUpdate) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *RouteTableUpdate) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *RouteTableUpdate) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


