# StandardHealthCheckMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | ID of the network | 
**InstanceId** | **string** | ID of the instance (gateway or tunnel) | 
**TunnelName** | Pointer to **string** | Name of the tunnel (only for tunnel type) | [optional] 

## Methods

### NewStandardHealthCheckMeta

`func NewStandardHealthCheckMeta(networkId string, instanceId string, ) *StandardHealthCheckMeta`

NewStandardHealthCheckMeta instantiates a new StandardHealthCheckMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardHealthCheckMetaWithDefaults

`func NewStandardHealthCheckMetaWithDefaults() *StandardHealthCheckMeta`

NewStandardHealthCheckMetaWithDefaults instantiates a new StandardHealthCheckMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *StandardHealthCheckMeta) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *StandardHealthCheckMeta) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *StandardHealthCheckMeta) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetInstanceId

`func (o *StandardHealthCheckMeta) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *StandardHealthCheckMeta) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *StandardHealthCheckMeta) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetTunnelName

`func (o *StandardHealthCheckMeta) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *StandardHealthCheckMeta) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *StandardHealthCheckMeta) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *StandardHealthCheckMeta) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


