# NetworkTunnelWireguard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID. | 
**Network** | **string** | ID of the network. | 
**Region** | **string** | ID of the network region. | 
**Instance** | **string** | ID of the network instance. | 
**InterfaceName** | **string** |  | 
**Type** | **string** |  | 
**IsHA** | **bool** | Indicates if it&#39;s a redundant tunnel. | 
**TenantId** | **string** | ID of the tenant. | 
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**LeftAllowedIP** | **[]string** |  | 
**LeftEndpoint** | **string** |  | 
**Vault** | **string** | vault. | 
**RequestConfigToken** | **string** | requestConfigToken. | 

## Methods

### NewNetworkTunnelWireguard

`func NewNetworkTunnelWireguard(id string, network string, region string, instance string, interfaceName string, type_ string, isHA bool, tenantId string, createdAt time.Time, leftAllowedIP []string, leftEndpoint string, vault string, requestConfigToken string, ) *NetworkTunnelWireguard`

NewNetworkTunnelWireguard instantiates a new NetworkTunnelWireguard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelWireguardWithDefaults

`func NewNetworkTunnelWireguardWithDefaults() *NetworkTunnelWireguard`

NewNetworkTunnelWireguardWithDefaults instantiates a new NetworkTunnelWireguard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnelWireguard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnelWireguard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnelWireguard) SetId(v string)`

SetId sets Id field to given value.


### GetNetwork

`func (o *NetworkTunnelWireguard) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkTunnelWireguard) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkTunnelWireguard) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRegion

`func (o *NetworkTunnelWireguard) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkTunnelWireguard) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkTunnelWireguard) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstance

`func (o *NetworkTunnelWireguard) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NetworkTunnelWireguard) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NetworkTunnelWireguard) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetInterfaceName

`func (o *NetworkTunnelWireguard) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkTunnelWireguard) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkTunnelWireguard) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetType

`func (o *NetworkTunnelWireguard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTunnelWireguard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTunnelWireguard) SetType(v string)`

SetType sets Type field to given value.


### GetIsHA

`func (o *NetworkTunnelWireguard) GetIsHA() bool`

GetIsHA returns the IsHA field if non-nil, zero value otherwise.

### GetIsHAOk

`func (o *NetworkTunnelWireguard) GetIsHAOk() (*bool, bool)`

GetIsHAOk returns a tuple with the IsHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHA

`func (o *NetworkTunnelWireguard) SetIsHA(v bool)`

SetIsHA sets IsHA field to given value.


### GetTenantId

`func (o *NetworkTunnelWireguard) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetworkTunnelWireguard) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetworkTunnelWireguard) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCreatedAt

`func (o *NetworkTunnelWireguard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTunnelWireguard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTunnelWireguard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkTunnelWireguard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkTunnelWireguard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkTunnelWireguard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkTunnelWireguard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLeftAllowedIP

`func (o *NetworkTunnelWireguard) GetLeftAllowedIP() []string`

GetLeftAllowedIP returns the LeftAllowedIP field if non-nil, zero value otherwise.

### GetLeftAllowedIPOk

`func (o *NetworkTunnelWireguard) GetLeftAllowedIPOk() (*[]string, bool)`

GetLeftAllowedIPOk returns a tuple with the LeftAllowedIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftAllowedIP

`func (o *NetworkTunnelWireguard) SetLeftAllowedIP(v []string)`

SetLeftAllowedIP sets LeftAllowedIP field to given value.


### GetLeftEndpoint

`func (o *NetworkTunnelWireguard) GetLeftEndpoint() string`

GetLeftEndpoint returns the LeftEndpoint field if non-nil, zero value otherwise.

### GetLeftEndpointOk

`func (o *NetworkTunnelWireguard) GetLeftEndpointOk() (*string, bool)`

GetLeftEndpointOk returns a tuple with the LeftEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftEndpoint

`func (o *NetworkTunnelWireguard) SetLeftEndpoint(v string)`

SetLeftEndpoint sets LeftEndpoint field to given value.


### GetVault

`func (o *NetworkTunnelWireguard) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *NetworkTunnelWireguard) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *NetworkTunnelWireguard) SetVault(v string)`

SetVault sets Vault field to given value.


### GetRequestConfigToken

`func (o *NetworkTunnelWireguard) GetRequestConfigToken() string`

GetRequestConfigToken returns the RequestConfigToken field if non-nil, zero value otherwise.

### GetRequestConfigTokenOk

`func (o *NetworkTunnelWireguard) GetRequestConfigTokenOk() (*string, bool)`

GetRequestConfigTokenOk returns a tuple with the RequestConfigToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigToken

`func (o *NetworkTunnelWireguard) SetRequestConfigToken(v string)`

SetRequestConfigToken sets RequestConfigToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


