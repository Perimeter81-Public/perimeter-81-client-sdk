# StaticTunnelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** | Authentication type for tunnel (psk for pre-shared key, cert for certificate) | [optional] 
**Passphrase** | Pointer to **string** | Pre-shared key for tunnel authentication (8-64 characters). Required when authType is psk. | [optional] 
**CustomerRootCA** | Pointer to **string** | Customer root certificate authority. Required when authType is cert. | [optional] 
**RemotePublicIP** | Pointer to **string** | Remote gateway public IP address | [optional] 
**RemoteID** | Pointer to **string** | Remote gateway ID | [optional] 
**Description** | Pointer to **string** | Optional tunnel description | [optional] 
**Features** | Pointer to [**NetworkFeaturesCreate**](NetworkFeaturesCreate.md) |  | [optional] 
**RoutingType** | Pointer to [**RoutingType**](RoutingType.md) |  | [optional] [default to ROUTE]
**RegionID** | **string** | Target region ID | 
**TunnelName** | **string** | Name of the static tunnel | 
**P81GatewaySubnets** | **[]string** | Harmony Sase gateway subnets | 
**RemoteGatewaySubnets** | **[]string** | Remote gateway subnets | 
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**KeyExchange** | **string** | IKE version for key exchange | [default to "ikev2"]
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | 
**Phase2** | [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | 

## Methods

### NewStaticTunnelCreate

`func NewStaticTunnelCreate(regionID string, tunnelName string, p81GatewaySubnets []string, remoteGatewaySubnets []string, keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfigV23, phase2 IPSecPhaseConfigV23, ) *StaticTunnelCreate`

NewStaticTunnelCreate instantiates a new StaticTunnelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticTunnelCreateWithDefaults

`func NewStaticTunnelCreateWithDefaults() *StaticTunnelCreate`

NewStaticTunnelCreateWithDefaults instantiates a new StaticTunnelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *StaticTunnelCreate) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *StaticTunnelCreate) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *StaticTunnelCreate) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *StaticTunnelCreate) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetPassphrase

`func (o *StaticTunnelCreate) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *StaticTunnelCreate) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *StaticTunnelCreate) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *StaticTunnelCreate) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCustomerRootCA

`func (o *StaticTunnelCreate) GetCustomerRootCA() string`

GetCustomerRootCA returns the CustomerRootCA field if non-nil, zero value otherwise.

### GetCustomerRootCAOk

`func (o *StaticTunnelCreate) GetCustomerRootCAOk() (*string, bool)`

GetCustomerRootCAOk returns a tuple with the CustomerRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRootCA

`func (o *StaticTunnelCreate) SetCustomerRootCA(v string)`

SetCustomerRootCA sets CustomerRootCA field to given value.

### HasCustomerRootCA

`func (o *StaticTunnelCreate) HasCustomerRootCA() bool`

HasCustomerRootCA returns a boolean if a field has been set.

### GetRemotePublicIP

`func (o *StaticTunnelCreate) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *StaticTunnelCreate) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *StaticTunnelCreate) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.

### HasRemotePublicIP

`func (o *StaticTunnelCreate) HasRemotePublicIP() bool`

HasRemotePublicIP returns a boolean if a field has been set.

### GetRemoteID

`func (o *StaticTunnelCreate) GetRemoteID() string`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *StaticTunnelCreate) GetRemoteIDOk() (*string, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *StaticTunnelCreate) SetRemoteID(v string)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *StaticTunnelCreate) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.

### GetDescription

`func (o *StaticTunnelCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticTunnelCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticTunnelCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticTunnelCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *StaticTunnelCreate) GetFeatures() NetworkFeaturesCreate`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *StaticTunnelCreate) GetFeaturesOk() (*NetworkFeaturesCreate, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *StaticTunnelCreate) SetFeatures(v NetworkFeaturesCreate)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *StaticTunnelCreate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetRoutingType

`func (o *StaticTunnelCreate) GetRoutingType() RoutingType`

GetRoutingType returns the RoutingType field if non-nil, zero value otherwise.

### GetRoutingTypeOk

`func (o *StaticTunnelCreate) GetRoutingTypeOk() (*RoutingType, bool)`

GetRoutingTypeOk returns a tuple with the RoutingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingType

`func (o *StaticTunnelCreate) SetRoutingType(v RoutingType)`

SetRoutingType sets RoutingType field to given value.

### HasRoutingType

`func (o *StaticTunnelCreate) HasRoutingType() bool`

HasRoutingType returns a boolean if a field has been set.

### GetRegionID

`func (o *StaticTunnelCreate) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *StaticTunnelCreate) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *StaticTunnelCreate) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetTunnelName

`func (o *StaticTunnelCreate) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *StaticTunnelCreate) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *StaticTunnelCreate) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetP81GatewaySubnets

`func (o *StaticTunnelCreate) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *StaticTunnelCreate) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *StaticTunnelCreate) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *StaticTunnelCreate) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *StaticTunnelCreate) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *StaticTunnelCreate) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetPeakBandwidth

`func (o *StaticTunnelCreate) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *StaticTunnelCreate) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *StaticTunnelCreate) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *StaticTunnelCreate) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetKeyExchange

`func (o *StaticTunnelCreate) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *StaticTunnelCreate) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *StaticTunnelCreate) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *StaticTunnelCreate) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *StaticTunnelCreate) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *StaticTunnelCreate) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *StaticTunnelCreate) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *StaticTunnelCreate) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *StaticTunnelCreate) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *StaticTunnelCreate) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *StaticTunnelCreate) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *StaticTunnelCreate) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *StaticTunnelCreate) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *StaticTunnelCreate) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *StaticTunnelCreate) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *StaticTunnelCreate) GetPhase1() IPSecPhaseConfigV23`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *StaticTunnelCreate) GetPhase1Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *StaticTunnelCreate) SetPhase1(v IPSecPhaseConfigV23)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *StaticTunnelCreate) GetPhase2() IPSecPhaseConfigV23`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *StaticTunnelCreate) GetPhase2Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *StaticTunnelCreate) SetPhase2(v IPSecPhaseConfigV23)`

SetPhase2 sets Phase2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


