# NetworkTunnelOpenvpn

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
**Passphrase** | **string** |  | 
**Username** | **string** | Openvpn username. | 

## Methods

### NewNetworkTunnelOpenvpn

`func NewNetworkTunnelOpenvpn(id string, network string, region string, instance string, interfaceName string, type_ string, isHA bool, tenantId string, createdAt time.Time, passphrase string, username string, ) *NetworkTunnelOpenvpn`

NewNetworkTunnelOpenvpn instantiates a new NetworkTunnelOpenvpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelOpenvpnWithDefaults

`func NewNetworkTunnelOpenvpnWithDefaults() *NetworkTunnelOpenvpn`

NewNetworkTunnelOpenvpnWithDefaults instantiates a new NetworkTunnelOpenvpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnelOpenvpn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnelOpenvpn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnelOpenvpn) SetId(v string)`

SetId sets Id field to given value.


### GetNetwork

`func (o *NetworkTunnelOpenvpn) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkTunnelOpenvpn) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkTunnelOpenvpn) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRegion

`func (o *NetworkTunnelOpenvpn) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkTunnelOpenvpn) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkTunnelOpenvpn) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstance

`func (o *NetworkTunnelOpenvpn) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NetworkTunnelOpenvpn) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NetworkTunnelOpenvpn) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetInterfaceName

`func (o *NetworkTunnelOpenvpn) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkTunnelOpenvpn) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkTunnelOpenvpn) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetType

`func (o *NetworkTunnelOpenvpn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTunnelOpenvpn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTunnelOpenvpn) SetType(v string)`

SetType sets Type field to given value.


### GetIsHA

`func (o *NetworkTunnelOpenvpn) GetIsHA() bool`

GetIsHA returns the IsHA field if non-nil, zero value otherwise.

### GetIsHAOk

`func (o *NetworkTunnelOpenvpn) GetIsHAOk() (*bool, bool)`

GetIsHAOk returns a tuple with the IsHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHA

`func (o *NetworkTunnelOpenvpn) SetIsHA(v bool)`

SetIsHA sets IsHA field to given value.


### GetTenantId

`func (o *NetworkTunnelOpenvpn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetworkTunnelOpenvpn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetworkTunnelOpenvpn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCreatedAt

`func (o *NetworkTunnelOpenvpn) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTunnelOpenvpn) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTunnelOpenvpn) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkTunnelOpenvpn) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkTunnelOpenvpn) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkTunnelOpenvpn) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkTunnelOpenvpn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPassphrase

`func (o *NetworkTunnelOpenvpn) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *NetworkTunnelOpenvpn) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *NetworkTunnelOpenvpn) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetUsername

`func (o *NetworkTunnelOpenvpn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NetworkTunnelOpenvpn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NetworkTunnelOpenvpn) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


