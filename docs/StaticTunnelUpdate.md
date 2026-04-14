# StaticTunnelUpdate

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
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**TunnelName** | Pointer to **string** | Name of the static tunnel | [optional] 
**P81GatewaySubnets** | Pointer to **[]string** | Harmony Sase gateway subnets | [optional] 
**RemoteGatewaySubnets** | Pointer to **[]string** | Remote gateway subnets | [optional] 
**KeyExchange** | Pointer to **string** | IKE version for key exchange | [optional] 
**IkeLifeTime** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **string** |  | [optional] 
**DpdDelay** | Pointer to **string** |  | [optional] 
**DpdTimeout** | Pointer to **string** |  | [optional] 
**Phase1** | Pointer to [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | [optional] 
**Phase2** | Pointer to [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | [optional] 

## Methods

### NewStaticTunnelUpdate

`func NewStaticTunnelUpdate() *StaticTunnelUpdate`

NewStaticTunnelUpdate instantiates a new StaticTunnelUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticTunnelUpdateWithDefaults

`func NewStaticTunnelUpdateWithDefaults() *StaticTunnelUpdate`

NewStaticTunnelUpdateWithDefaults instantiates a new StaticTunnelUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *StaticTunnelUpdate) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *StaticTunnelUpdate) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *StaticTunnelUpdate) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *StaticTunnelUpdate) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetPassphrase

`func (o *StaticTunnelUpdate) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *StaticTunnelUpdate) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *StaticTunnelUpdate) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *StaticTunnelUpdate) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCustomerRootCA

`func (o *StaticTunnelUpdate) GetCustomerRootCA() string`

GetCustomerRootCA returns the CustomerRootCA field if non-nil, zero value otherwise.

### GetCustomerRootCAOk

`func (o *StaticTunnelUpdate) GetCustomerRootCAOk() (*string, bool)`

GetCustomerRootCAOk returns a tuple with the CustomerRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRootCA

`func (o *StaticTunnelUpdate) SetCustomerRootCA(v string)`

SetCustomerRootCA sets CustomerRootCA field to given value.

### HasCustomerRootCA

`func (o *StaticTunnelUpdate) HasCustomerRootCA() bool`

HasCustomerRootCA returns a boolean if a field has been set.

### GetRemotePublicIP

`func (o *StaticTunnelUpdate) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *StaticTunnelUpdate) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *StaticTunnelUpdate) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.

### HasRemotePublicIP

`func (o *StaticTunnelUpdate) HasRemotePublicIP() bool`

HasRemotePublicIP returns a boolean if a field has been set.

### GetRemoteID

`func (o *StaticTunnelUpdate) GetRemoteID() string`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *StaticTunnelUpdate) GetRemoteIDOk() (*string, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *StaticTunnelUpdate) SetRemoteID(v string)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *StaticTunnelUpdate) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.

### GetDescription

`func (o *StaticTunnelUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticTunnelUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticTunnelUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticTunnelUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *StaticTunnelUpdate) GetFeatures() NetworkFeaturesCreate`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *StaticTunnelUpdate) GetFeaturesOk() (*NetworkFeaturesCreate, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *StaticTunnelUpdate) SetFeatures(v NetworkFeaturesCreate)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *StaticTunnelUpdate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetRoutingType

`func (o *StaticTunnelUpdate) GetRoutingType() RoutingType`

GetRoutingType returns the RoutingType field if non-nil, zero value otherwise.

### GetRoutingTypeOk

`func (o *StaticTunnelUpdate) GetRoutingTypeOk() (*RoutingType, bool)`

GetRoutingTypeOk returns a tuple with the RoutingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingType

`func (o *StaticTunnelUpdate) SetRoutingType(v RoutingType)`

SetRoutingType sets RoutingType field to given value.

### HasRoutingType

`func (o *StaticTunnelUpdate) HasRoutingType() bool`

HasRoutingType returns a boolean if a field has been set.

### GetPeakBandwidth

`func (o *StaticTunnelUpdate) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *StaticTunnelUpdate) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *StaticTunnelUpdate) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *StaticTunnelUpdate) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetTunnelName

`func (o *StaticTunnelUpdate) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *StaticTunnelUpdate) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *StaticTunnelUpdate) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *StaticTunnelUpdate) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.

### GetP81GatewaySubnets

`func (o *StaticTunnelUpdate) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *StaticTunnelUpdate) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *StaticTunnelUpdate) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.

### HasP81GatewaySubnets

`func (o *StaticTunnelUpdate) HasP81GatewaySubnets() bool`

HasP81GatewaySubnets returns a boolean if a field has been set.

### GetRemoteGatewaySubnets

`func (o *StaticTunnelUpdate) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *StaticTunnelUpdate) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *StaticTunnelUpdate) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.

### HasRemoteGatewaySubnets

`func (o *StaticTunnelUpdate) HasRemoteGatewaySubnets() bool`

HasRemoteGatewaySubnets returns a boolean if a field has been set.

### GetKeyExchange

`func (o *StaticTunnelUpdate) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *StaticTunnelUpdate) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *StaticTunnelUpdate) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.

### HasKeyExchange

`func (o *StaticTunnelUpdate) HasKeyExchange() bool`

HasKeyExchange returns a boolean if a field has been set.

### GetIkeLifeTime

`func (o *StaticTunnelUpdate) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *StaticTunnelUpdate) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *StaticTunnelUpdate) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.

### HasIkeLifeTime

`func (o *StaticTunnelUpdate) HasIkeLifeTime() bool`

HasIkeLifeTime returns a boolean if a field has been set.

### GetLifetime

`func (o *StaticTunnelUpdate) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *StaticTunnelUpdate) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *StaticTunnelUpdate) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *StaticTunnelUpdate) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetDpdDelay

`func (o *StaticTunnelUpdate) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *StaticTunnelUpdate) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *StaticTunnelUpdate) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.

### HasDpdDelay

`func (o *StaticTunnelUpdate) HasDpdDelay() bool`

HasDpdDelay returns a boolean if a field has been set.

### GetDpdTimeout

`func (o *StaticTunnelUpdate) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *StaticTunnelUpdate) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *StaticTunnelUpdate) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.

### HasDpdTimeout

`func (o *StaticTunnelUpdate) HasDpdTimeout() bool`

HasDpdTimeout returns a boolean if a field has been set.

### GetPhase1

`func (o *StaticTunnelUpdate) GetPhase1() IPSecPhaseConfigV23`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *StaticTunnelUpdate) GetPhase1Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *StaticTunnelUpdate) SetPhase1(v IPSecPhaseConfigV23)`

SetPhase1 sets Phase1 field to given value.

### HasPhase1

`func (o *StaticTunnelUpdate) HasPhase1() bool`

HasPhase1 returns a boolean if a field has been set.

### GetPhase2

`func (o *StaticTunnelUpdate) GetPhase2() IPSecPhaseConfigV23`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *StaticTunnelUpdate) GetPhase2Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *StaticTunnelUpdate) SetPhase2(v IPSecPhaseConfigV23)`

SetPhase2 sets Phase2 field to given value.

### HasPhase2

`func (o *StaticTunnelUpdate) HasPhase2() bool`

HasPhase2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


