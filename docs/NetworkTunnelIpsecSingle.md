# NetworkTunnelIpsecSingle

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
**KeyExchange** | **string** |  | 
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Phase2** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Right** | **string** |  | 
**RightID** | [**NetworkIpsecBaseAllOfRightID**](NetworkIpsecBaseAllOfRightID.md) |  | 
**Passphrase** | **string** |  | 
**DpdAction** | **string** | dpdAction. | 
**LeftSubnets** | **[]string** |  | 
**RightSubnets** | **[]string** |  | 

## Methods

### NewNetworkTunnelIpsecSingle

`func NewNetworkTunnelIpsecSingle(id string, network string, region string, instance string, interfaceName string, type_ string, isHA bool, tenantId string, createdAt time.Time, keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfig, phase2 IPSecPhaseConfig, right string, rightID NetworkIpsecBaseAllOfRightID, passphrase string, dpdAction string, leftSubnets []string, rightSubnets []string, ) *NetworkTunnelIpsecSingle`

NewNetworkTunnelIpsecSingle instantiates a new NetworkTunnelIpsecSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTunnelIpsecSingleWithDefaults

`func NewNetworkTunnelIpsecSingleWithDefaults() *NetworkTunnelIpsecSingle`

NewNetworkTunnelIpsecSingleWithDefaults instantiates a new NetworkTunnelIpsecSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTunnelIpsecSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTunnelIpsecSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTunnelIpsecSingle) SetId(v string)`

SetId sets Id field to given value.


### GetNetwork

`func (o *NetworkTunnelIpsecSingle) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkTunnelIpsecSingle) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkTunnelIpsecSingle) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRegion

`func (o *NetworkTunnelIpsecSingle) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkTunnelIpsecSingle) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkTunnelIpsecSingle) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstance

`func (o *NetworkTunnelIpsecSingle) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *NetworkTunnelIpsecSingle) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *NetworkTunnelIpsecSingle) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetInterfaceName

`func (o *NetworkTunnelIpsecSingle) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NetworkTunnelIpsecSingle) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NetworkTunnelIpsecSingle) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.


### GetType

`func (o *NetworkTunnelIpsecSingle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTunnelIpsecSingle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTunnelIpsecSingle) SetType(v string)`

SetType sets Type field to given value.


### GetIsHA

`func (o *NetworkTunnelIpsecSingle) GetIsHA() bool`

GetIsHA returns the IsHA field if non-nil, zero value otherwise.

### GetIsHAOk

`func (o *NetworkTunnelIpsecSingle) GetIsHAOk() (*bool, bool)`

GetIsHAOk returns a tuple with the IsHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHA

`func (o *NetworkTunnelIpsecSingle) SetIsHA(v bool)`

SetIsHA sets IsHA field to given value.


### GetTenantId

`func (o *NetworkTunnelIpsecSingle) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetworkTunnelIpsecSingle) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetworkTunnelIpsecSingle) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCreatedAt

`func (o *NetworkTunnelIpsecSingle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkTunnelIpsecSingle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkTunnelIpsecSingle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkTunnelIpsecSingle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkTunnelIpsecSingle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkTunnelIpsecSingle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkTunnelIpsecSingle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetKeyExchange

`func (o *NetworkTunnelIpsecSingle) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *NetworkTunnelIpsecSingle) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *NetworkTunnelIpsecSingle) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *NetworkTunnelIpsecSingle) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *NetworkTunnelIpsecSingle) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *NetworkTunnelIpsecSingle) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *NetworkTunnelIpsecSingle) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *NetworkTunnelIpsecSingle) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *NetworkTunnelIpsecSingle) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *NetworkTunnelIpsecSingle) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *NetworkTunnelIpsecSingle) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *NetworkTunnelIpsecSingle) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *NetworkTunnelIpsecSingle) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *NetworkTunnelIpsecSingle) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *NetworkTunnelIpsecSingle) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *NetworkTunnelIpsecSingle) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *NetworkTunnelIpsecSingle) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *NetworkTunnelIpsecSingle) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *NetworkTunnelIpsecSingle) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *NetworkTunnelIpsecSingle) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *NetworkTunnelIpsecSingle) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.


### GetRight

`func (o *NetworkTunnelIpsecSingle) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *NetworkTunnelIpsecSingle) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *NetworkTunnelIpsecSingle) SetRight(v string)`

SetRight sets Right field to given value.


### GetRightID

`func (o *NetworkTunnelIpsecSingle) GetRightID() NetworkIpsecBaseAllOfRightID`

GetRightID returns the RightID field if non-nil, zero value otherwise.

### GetRightIDOk

`func (o *NetworkTunnelIpsecSingle) GetRightIDOk() (*NetworkIpsecBaseAllOfRightID, bool)`

GetRightIDOk returns a tuple with the RightID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightID

`func (o *NetworkTunnelIpsecSingle) SetRightID(v NetworkIpsecBaseAllOfRightID)`

SetRightID sets RightID field to given value.


### GetPassphrase

`func (o *NetworkTunnelIpsecSingle) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *NetworkTunnelIpsecSingle) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *NetworkTunnelIpsecSingle) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetDpdAction

`func (o *NetworkTunnelIpsecSingle) GetDpdAction() string`

GetDpdAction returns the DpdAction field if non-nil, zero value otherwise.

### GetDpdActionOk

`func (o *NetworkTunnelIpsecSingle) GetDpdActionOk() (*string, bool)`

GetDpdActionOk returns a tuple with the DpdAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdAction

`func (o *NetworkTunnelIpsecSingle) SetDpdAction(v string)`

SetDpdAction sets DpdAction field to given value.


### GetLeftSubnets

`func (o *NetworkTunnelIpsecSingle) GetLeftSubnets() []string`

GetLeftSubnets returns the LeftSubnets field if non-nil, zero value otherwise.

### GetLeftSubnetsOk

`func (o *NetworkTunnelIpsecSingle) GetLeftSubnetsOk() (*[]string, bool)`

GetLeftSubnetsOk returns a tuple with the LeftSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftSubnets

`func (o *NetworkTunnelIpsecSingle) SetLeftSubnets(v []string)`

SetLeftSubnets sets LeftSubnets field to given value.


### GetRightSubnets

`func (o *NetworkTunnelIpsecSingle) GetRightSubnets() []string`

GetRightSubnets returns the RightSubnets field if non-nil, zero value otherwise.

### GetRightSubnetsOk

`func (o *NetworkTunnelIpsecSingle) GetRightSubnetsOk() (*[]string, bool)`

GetRightSubnetsOk returns a tuple with the RightSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightSubnets

`func (o *NetworkTunnelIpsecSingle) SetRightSubnets(v []string)`

SetRightSubnets sets RightSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


