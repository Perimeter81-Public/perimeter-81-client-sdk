# GetRouteTable200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Route ID | [optional] 
**Subnets** | Pointer to **[]string** |  | [optional] 
**InterfaceName** | Pointer to **string** | Route table name | [optional] 
**Propagated** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetRouteTable200ResponseInner

`func NewGetRouteTable200ResponseInner() *GetRouteTable200ResponseInner`

NewGetRouteTable200ResponseInner instantiates a new GetRouteTable200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRouteTable200ResponseInnerWithDefaults

`func NewGetRouteTable200ResponseInnerWithDefaults() *GetRouteTable200ResponseInner`

NewGetRouteTable200ResponseInnerWithDefaults instantiates a new GetRouteTable200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRouteTable200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRouteTable200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRouteTable200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetRouteTable200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubnets

`func (o *GetRouteTable200ResponseInner) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetRouteTable200ResponseInner) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetRouteTable200ResponseInner) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetRouteTable200ResponseInner) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetInterfaceName

`func (o *GetRouteTable200ResponseInner) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *GetRouteTable200ResponseInner) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *GetRouteTable200ResponseInner) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *GetRouteTable200ResponseInner) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetPropagated

`func (o *GetRouteTable200ResponseInner) GetPropagated() bool`

GetPropagated returns the Propagated field if non-nil, zero value otherwise.

### GetPropagatedOk

`func (o *GetRouteTable200ResponseInner) GetPropagatedOk() (*bool, bool)`

GetPropagatedOk returns a tuple with the Propagated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagated

`func (o *GetRouteTable200ResponseInner) SetPropagated(v bool)`

SetPropagated sets Propagated field to given value.

### HasPropagated

`func (o *GetRouteTable200ResponseInner) HasPropagated() bool`

HasPropagated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


