# EnhancedTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | Authentication type for tunnel (psk for pre-shared key, cert for certificate) | 
**Passphrase** | Pointer to **string** | Pre-shared key for tunnel authentication (8-64 characters). Required when authType is psk. | [optional] 
**CustomerRootCA** | Pointer to **string** | Customer root certificate authority. Required when authType is cert. | [optional] 
**RegionID** | **string** | Target region ID | 
**TunnelName** | **string** | Name of the static tunnel | 
**P81GatewaySubnets** | **[]string** | Harmony Sase gateway subnets | 
**RemoteGatewaySubnets** | **[]string** | Remote gateway subnets | 
**KeyExchange** | **string** | IKE version for key exchange | [default to "ikev2"]
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | 
**Phase2** | [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | 
**RemotePublicIP** | **string** | Remote gateway public IP address | 
**RemoteID** | **string** | Remote gateway ID | 
**Description** | Pointer to **string** | Optional tunnel description | [optional] 
**Features** | [**NetworkFeatures**](NetworkFeatures.md) |  | 
**RoutingType** | Pointer to [**RoutingType**](RoutingType.md) |  | [optional] [default to ROUTE]
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**Id** | **string** | Enhanced tunnel ID | 
**HaTunnelID** | **string** | Enhanced dynamic tunnel group ID (or tunnel ID for static tunnel) | 
**DpdAction** | **string** | Enhanced tunnel DPDTime actions. | 

## Methods

### NewEnhancedTunnel

`func NewEnhancedTunnel(authType string, regionID string, tunnelName string, p81GatewaySubnets []string, remoteGatewaySubnets []string, keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfigV23, phase2 IPSecPhaseConfigV23, remotePublicIP string, remoteID string, features NetworkFeatures, id string, haTunnelID string, dpdAction string, ) *EnhancedTunnel`

NewEnhancedTunnel instantiates a new EnhancedTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedTunnelWithDefaults

`func NewEnhancedTunnelWithDefaults() *EnhancedTunnel`

NewEnhancedTunnelWithDefaults instantiates a new EnhancedTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *EnhancedTunnel) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *EnhancedTunnel) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *EnhancedTunnel) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetPassphrase

`func (o *EnhancedTunnel) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *EnhancedTunnel) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *EnhancedTunnel) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *EnhancedTunnel) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCustomerRootCA

`func (o *EnhancedTunnel) GetCustomerRootCA() string`

GetCustomerRootCA returns the CustomerRootCA field if non-nil, zero value otherwise.

### GetCustomerRootCAOk

`func (o *EnhancedTunnel) GetCustomerRootCAOk() (*string, bool)`

GetCustomerRootCAOk returns a tuple with the CustomerRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRootCA

`func (o *EnhancedTunnel) SetCustomerRootCA(v string)`

SetCustomerRootCA sets CustomerRootCA field to given value.

### HasCustomerRootCA

`func (o *EnhancedTunnel) HasCustomerRootCA() bool`

HasCustomerRootCA returns a boolean if a field has been set.

### GetRegionID

`func (o *EnhancedTunnel) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *EnhancedTunnel) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *EnhancedTunnel) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetTunnelName

`func (o *EnhancedTunnel) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *EnhancedTunnel) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *EnhancedTunnel) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetP81GatewaySubnets

`func (o *EnhancedTunnel) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *EnhancedTunnel) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *EnhancedTunnel) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *EnhancedTunnel) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *EnhancedTunnel) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *EnhancedTunnel) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetKeyExchange

`func (o *EnhancedTunnel) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *EnhancedTunnel) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *EnhancedTunnel) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *EnhancedTunnel) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *EnhancedTunnel) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *EnhancedTunnel) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *EnhancedTunnel) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *EnhancedTunnel) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *EnhancedTunnel) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *EnhancedTunnel) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *EnhancedTunnel) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *EnhancedTunnel) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *EnhancedTunnel) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *EnhancedTunnel) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *EnhancedTunnel) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *EnhancedTunnel) GetPhase1() IPSecPhaseConfigV23`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *EnhancedTunnel) GetPhase1Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *EnhancedTunnel) SetPhase1(v IPSecPhaseConfigV23)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *EnhancedTunnel) GetPhase2() IPSecPhaseConfigV23`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *EnhancedTunnel) GetPhase2Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *EnhancedTunnel) SetPhase2(v IPSecPhaseConfigV23)`

SetPhase2 sets Phase2 field to given value.


### GetRemotePublicIP

`func (o *EnhancedTunnel) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *EnhancedTunnel) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *EnhancedTunnel) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.


### GetRemoteID

`func (o *EnhancedTunnel) GetRemoteID() string`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *EnhancedTunnel) GetRemoteIDOk() (*string, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *EnhancedTunnel) SetRemoteID(v string)`

SetRemoteID sets RemoteID field to given value.


### GetDescription

`func (o *EnhancedTunnel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnhancedTunnel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnhancedTunnel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnhancedTunnel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *EnhancedTunnel) GetFeatures() NetworkFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *EnhancedTunnel) GetFeaturesOk() (*NetworkFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *EnhancedTunnel) SetFeatures(v NetworkFeatures)`

SetFeatures sets Features field to given value.


### GetRoutingType

`func (o *EnhancedTunnel) GetRoutingType() RoutingType`

GetRoutingType returns the RoutingType field if non-nil, zero value otherwise.

### GetRoutingTypeOk

`func (o *EnhancedTunnel) GetRoutingTypeOk() (*RoutingType, bool)`

GetRoutingTypeOk returns a tuple with the RoutingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingType

`func (o *EnhancedTunnel) SetRoutingType(v RoutingType)`

SetRoutingType sets RoutingType field to given value.

### HasRoutingType

`func (o *EnhancedTunnel) HasRoutingType() bool`

HasRoutingType returns a boolean if a field has been set.

### GetPeakBandwidth

`func (o *EnhancedTunnel) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *EnhancedTunnel) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *EnhancedTunnel) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *EnhancedTunnel) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetId

`func (o *EnhancedTunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedTunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedTunnel) SetId(v string)`

SetId sets Id field to given value.


### GetHaTunnelID

`func (o *EnhancedTunnel) GetHaTunnelID() string`

GetHaTunnelID returns the HaTunnelID field if non-nil, zero value otherwise.

### GetHaTunnelIDOk

`func (o *EnhancedTunnel) GetHaTunnelIDOk() (*string, bool)`

GetHaTunnelIDOk returns a tuple with the HaTunnelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaTunnelID

`func (o *EnhancedTunnel) SetHaTunnelID(v string)`

SetHaTunnelID sets HaTunnelID field to given value.


### GetDpdAction

`func (o *EnhancedTunnel) GetDpdAction() string`

GetDpdAction returns the DpdAction field if non-nil, zero value otherwise.

### GetDpdActionOk

`func (o *EnhancedTunnel) GetDpdActionOk() (*string, bool)`

GetDpdActionOk returns a tuple with the DpdAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdAction

`func (o *EnhancedTunnel) SetDpdAction(v string)`

SetDpdAction sets DpdAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


