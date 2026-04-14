# WireguardTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteEndpoint** | **string** |  | 
**RemoteSubnets** | **[]string** |  | 
**RegionID** | **string** | Region ID | 
**GatewayID** | **string** | Gateway ID | 
**TunnelName** | **string** | The name of the tunnel | 
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**TunnelID** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "wireguard"]

## Methods

### NewWireguardTunnel

`func NewWireguardTunnel(remoteEndpoint string, remoteSubnets []string, regionID string, gatewayID string, tunnelName string, createdAt time.Time, ) *WireguardTunnel`

NewWireguardTunnel instantiates a new WireguardTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardTunnelWithDefaults

`func NewWireguardTunnelWithDefaults() *WireguardTunnel`

NewWireguardTunnelWithDefaults instantiates a new WireguardTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteEndpoint

`func (o *WireguardTunnel) GetRemoteEndpoint() string`

GetRemoteEndpoint returns the RemoteEndpoint field if non-nil, zero value otherwise.

### GetRemoteEndpointOk

`func (o *WireguardTunnel) GetRemoteEndpointOk() (*string, bool)`

GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndpoint

`func (o *WireguardTunnel) SetRemoteEndpoint(v string)`

SetRemoteEndpoint sets RemoteEndpoint field to given value.


### GetRemoteSubnets

`func (o *WireguardTunnel) GetRemoteSubnets() []string`

GetRemoteSubnets returns the RemoteSubnets field if non-nil, zero value otherwise.

### GetRemoteSubnetsOk

`func (o *WireguardTunnel) GetRemoteSubnetsOk() (*[]string, bool)`

GetRemoteSubnetsOk returns a tuple with the RemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnets

`func (o *WireguardTunnel) SetRemoteSubnets(v []string)`

SetRemoteSubnets sets RemoteSubnets field to given value.


### GetRegionID

`func (o *WireguardTunnel) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *WireguardTunnel) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *WireguardTunnel) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetGatewayID

`func (o *WireguardTunnel) GetGatewayID() string`

GetGatewayID returns the GatewayID field if non-nil, zero value otherwise.

### GetGatewayIDOk

`func (o *WireguardTunnel) GetGatewayIDOk() (*string, bool)`

GetGatewayIDOk returns a tuple with the GatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayID

`func (o *WireguardTunnel) SetGatewayID(v string)`

SetGatewayID sets GatewayID field to given value.


### GetTunnelName

`func (o *WireguardTunnel) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *WireguardTunnel) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *WireguardTunnel) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetCreatedAt

`func (o *WireguardTunnel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WireguardTunnel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WireguardTunnel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WireguardTunnel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WireguardTunnel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WireguardTunnel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WireguardTunnel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTunnelID

`func (o *WireguardTunnel) GetTunnelID() string`

GetTunnelID returns the TunnelID field if non-nil, zero value otherwise.

### GetTunnelIDOk

`func (o *WireguardTunnel) GetTunnelIDOk() (*string, bool)`

GetTunnelIDOk returns a tuple with the TunnelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelID

`func (o *WireguardTunnel) SetTunnelID(v string)`

SetTunnelID sets TunnelID field to given value.

### HasTunnelID

`func (o *WireguardTunnel) HasTunnelID() bool`

HasTunnelID returns a boolean if a field has been set.

### GetType

`func (o *WireguardTunnel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireguardTunnel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireguardTunnel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WireguardTunnel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


