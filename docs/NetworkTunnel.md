# NetworkTunnel

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
**LeftAllowedIP** | **[]string** |  | 
**LeftEndpoint** | **string** |  | 
**Vault** | **string** | vault. | 
**RequestConfigToken** | **string** | requestConfigToken. | 
**KeyExchange** | **string** |  | 
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Phase2** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Right** | **string** |  | 
**RightID** | [**NetworkIpsecBaseAllOfRightID**](NetworkIpsecBaseAllOfRightID.md) |  | 
**DpdAction** | **string** | dpdAction. | 
**LeftSubnets** | **[]string** |  | 
**RightSubnets** | **[]string** |  | 
**HaTunnelID** | [**NetworkHaTunnelID**](NetworkHaTunnelID.md) |  | 
**RightASN** | [**ASN**](ASN.md) |  | 
**RightPrivateIP** | **string** |  | 
**LeftPrivateIP** | **string** |  | 

## Methods

### NewNetworkTunnel

`func NewNetworkTunnel(id string, network string, region string, instance string, interfaceName string, type_ string, isHA bool, tenantId string, createdAt time.Time, passphrase string, username string, leftAllowedIP []string, leftEndpoint string, vault string, requestConfigToken string, keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfig, phase2 IPSecPhaseConfig, right string, rightID NetworkIpsecBaseAllOfRightID, dpdAction string, leftSubnets []string, rightSubnets []string, haTunnelID NetworkHaTunnelID, rightASN ASN, rightPrivateIP string, leftPrivateIP string, ) *NetworkTunnel`

NewNetworkTunnel instantiates a new NetworkTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelWithDefaults

`func NewNetworkTunnelWithDefaults() *NetworkTunnel`

NewNetworkTunnelWithDefaults instantiates a new NetworkTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnel) SetId(v string)`

SetId sets Id field to given value.


### GetNetwork

`func (o *NetworkTunnel) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkTunnel) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkTunnel) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRegion

`func (o *NetworkTunnel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkTunnel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkTunnel) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstance

`func (o *NetworkTunnel) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NetworkTunnel) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NetworkTunnel) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetInterfaceName

`func (o *NetworkTunnel) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkTunnel) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkTunnel) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetType

`func (o *NetworkTunnel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTunnel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTunnel) SetType(v string)`

SetType sets Type field to given value.


### GetIsHA

`func (o *NetworkTunnel) GetIsHA() bool`

GetIsHA returns the IsHA field if non-nil, zero value otherwise.

### GetIsHAOk

`func (o *NetworkTunnel) GetIsHAOk() (*bool, bool)`

GetIsHAOk returns a tuple with the IsHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHA

`func (o *NetworkTunnel) SetIsHA(v bool)`

SetIsHA sets IsHA field to given value.


### GetTenantId

`func (o *NetworkTunnel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetworkTunnel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetworkTunnel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCreatedAt

`func (o *NetworkTunnel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTunnel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTunnel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkTunnel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkTunnel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkTunnel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkTunnel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPassphrase

`func (o *NetworkTunnel) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *NetworkTunnel) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *NetworkTunnel) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetUsername

`func (o *NetworkTunnel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NetworkTunnel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NetworkTunnel) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetLeftAllowedIP

`func (o *NetworkTunnel) GetLeftAllowedIP() []string`

GetLeftAllowedIP returns the LeftAllowedIP field if non-nil, zero value otherwise.

### GetLeftAllowedIPOk

`func (o *NetworkTunnel) GetLeftAllowedIPOk() (*[]string, bool)`

GetLeftAllowedIPOk returns a tuple with the LeftAllowedIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftAllowedIP

`func (o *NetworkTunnel) SetLeftAllowedIP(v []string)`

SetLeftAllowedIP sets LeftAllowedIP field to given value.


### GetLeftEndpoint

`func (o *NetworkTunnel) GetLeftEndpoint() string`

GetLeftEndpoint returns the LeftEndpoint field if non-nil, zero value otherwise.

### GetLeftEndpointOk

`func (o *NetworkTunnel) GetLeftEndpointOk() (*string, bool)`

GetLeftEndpointOk returns a tuple with the LeftEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftEndpoint

`func (o *NetworkTunnel) SetLeftEndpoint(v string)`

SetLeftEndpoint sets LeftEndpoint field to given value.


### GetVault

`func (o *NetworkTunnel) GetVault() string`

GetVault returns the Vault field if non-nil, zero value otherwise.

### GetVaultOk

`func (o *NetworkTunnel) GetVaultOk() (*string, bool)`

GetVaultOk returns a tuple with the Vault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVault

`func (o *NetworkTunnel) SetVault(v string)`

SetVault sets Vault field to given value.


### GetRequestConfigToken

`func (o *NetworkTunnel) GetRequestConfigToken() string`

GetRequestConfigToken returns the RequestConfigToken field if non-nil, zero value otherwise.

### GetRequestConfigTokenOk

`func (o *NetworkTunnel) GetRequestConfigTokenOk() (*string, bool)`

GetRequestConfigTokenOk returns a tuple with the RequestConfigToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfigToken

`func (o *NetworkTunnel) SetRequestConfigToken(v string)`

SetRequestConfigToken sets RequestConfigToken field to given value.


### GetKeyExchange

`func (o *NetworkTunnel) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *NetworkTunnel) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *NetworkTunnel) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *NetworkTunnel) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *NetworkTunnel) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *NetworkTunnel) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *NetworkTunnel) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *NetworkTunnel) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *NetworkTunnel) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *NetworkTunnel) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *NetworkTunnel) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *NetworkTunnel) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *NetworkTunnel) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *NetworkTunnel) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *NetworkTunnel) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *NetworkTunnel) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *NetworkTunnel) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *NetworkTunnel) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *NetworkTunnel) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *NetworkTunnel) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *NetworkTunnel) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.


### GetRight

`func (o *NetworkTunnel) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *NetworkTunnel) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *NetworkTunnel) SetRight(v string)`

SetRight sets Right field to given value.


### GetRightID

`func (o *NetworkTunnel) GetRightID() NetworkIpsecBaseAllOfRightID`

GetRightID returns the RightID field if non-nil, zero value otherwise.

### GetRightIDOk

`func (o *NetworkTunnel) GetRightIDOk() (*NetworkIpsecBaseAllOfRightID, bool)`

GetRightIDOk returns a tuple with the RightID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightID

`func (o *NetworkTunnel) SetRightID(v NetworkIpsecBaseAllOfRightID)`

SetRightID sets RightID field to given value.


### GetDpdAction

`func (o *NetworkTunnel) GetDpdAction() string`

GetDpdAction returns the DpdAction field if non-nil, zero value otherwise.

### GetDpdActionOk

`func (o *NetworkTunnel) GetDpdActionOk() (*string, bool)`

GetDpdActionOk returns a tuple with the DpdAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdAction

`func (o *NetworkTunnel) SetDpdAction(v string)`

SetDpdAction sets DpdAction field to given value.


### GetLeftSubnets

`func (o *NetworkTunnel) GetLeftSubnets() []string`

GetLeftSubnets returns the LeftSubnets field if non-nil, zero value otherwise.

### GetLeftSubnetsOk

`func (o *NetworkTunnel) GetLeftSubnetsOk() (*[]string, bool)`

GetLeftSubnetsOk returns a tuple with the LeftSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftSubnets

`func (o *NetworkTunnel) SetLeftSubnets(v []string)`

SetLeftSubnets sets LeftSubnets field to given value.


### GetRightSubnets

`func (o *NetworkTunnel) GetRightSubnets() []string`

GetRightSubnets returns the RightSubnets field if non-nil, zero value otherwise.

### GetRightSubnetsOk

`func (o *NetworkTunnel) GetRightSubnetsOk() (*[]string, bool)`

GetRightSubnetsOk returns a tuple with the RightSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightSubnets

`func (o *NetworkTunnel) SetRightSubnets(v []string)`

SetRightSubnets sets RightSubnets field to given value.


### GetHaTunnelID

`func (o *NetworkTunnel) GetHaTunnelID() NetworkHaTunnelID`

GetHaTunnelID returns the HaTunnelID field if non-nil, zero value otherwise.

### GetHaTunnelIDOk

`func (o *NetworkTunnel) GetHaTunnelIDOk() (*NetworkHaTunnelID, bool)`

GetHaTunnelIDOk returns a tuple with the HaTunnelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaTunnelID

`func (o *NetworkTunnel) SetHaTunnelID(v NetworkHaTunnelID)`

SetHaTunnelID sets HaTunnelID field to given value.


### GetRightASN

`func (o *NetworkTunnel) GetRightASN() ASN`

GetRightASN returns the RightASN field if non-nil, zero value otherwise.

### GetRightASNOk

`func (o *NetworkTunnel) GetRightASNOk() (*ASN, bool)`

GetRightASNOk returns a tuple with the RightASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightASN

`func (o *NetworkTunnel) SetRightASN(v ASN)`

SetRightASN sets RightASN field to given value.


### GetRightPrivateIP

`func (o *NetworkTunnel) GetRightPrivateIP() string`

GetRightPrivateIP returns the RightPrivateIP field if non-nil, zero value otherwise.

### GetRightPrivateIPOk

`func (o *NetworkTunnel) GetRightPrivateIPOk() (*string, bool)`

GetRightPrivateIPOk returns a tuple with the RightPrivateIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightPrivateIP

`func (o *NetworkTunnel) SetRightPrivateIP(v string)`

SetRightPrivateIP sets RightPrivateIP field to given value.


### GetLeftPrivateIP

`func (o *NetworkTunnel) GetLeftPrivateIP() string`

GetLeftPrivateIP returns the LeftPrivateIP field if non-nil, zero value otherwise.

### GetLeftPrivateIPOk

`func (o *NetworkTunnel) GetLeftPrivateIPOk() (*string, bool)`

GetLeftPrivateIPOk returns a tuple with the LeftPrivateIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftPrivateIP

`func (o *NetworkTunnel) SetLeftPrivateIP(v string)`

SetLeftPrivateIP sets LeftPrivateIP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


