# CreateIPSecSinglePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionID** | **string** | Region ID | 
**GatewayID** | **string** | Gateway ID | 
**TunnelName** | **string** | The name of the tunnel | 
**P81GatewaySubnets** | **[]string** |  | 
**RemoteGatewaySubnets** | **[]string** |  | 
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**P81ASN** | Pointer to [**RemoteASN**](RemoteASN.md) |  | [optional] 
**Features** | Pointer to [**IPSecSharedSettingsFeatures**](IPSecSharedSettingsFeatures.md) |  | [optional] 
**KeyExchange** | **string** |  | 
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Phase2** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Passphrase** | **string** |  | 
**RemotePublicIP** | **string** |  | 
**RemoteID** | Pointer to [**RemoteID**](RemoteID.md) |  | [optional] 

## Methods

### NewCreateIPSecSinglePayload

`func NewCreateIPSecSinglePayload(regionID string, gatewayID string, tunnelName string, p81GatewaySubnets []string, remoteGatewaySubnets []string, keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfig, phase2 IPSecPhaseConfig, passphrase string, remotePublicIP string, ) *CreateIPSecSinglePayload`

NewCreateIPSecSinglePayload instantiates a new CreateIPSecSinglePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIPSecSinglePayloadWithDefaults

`func NewCreateIPSecSinglePayloadWithDefaults() *CreateIPSecSinglePayload`

NewCreateIPSecSinglePayloadWithDefaults instantiates a new CreateIPSecSinglePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionID

`func (o *CreateIPSecSinglePayload) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *CreateIPSecSinglePayload) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *CreateIPSecSinglePayload) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetGatewayID

`func (o *CreateIPSecSinglePayload) GetGatewayID() string`

GetGatewayID returns the GatewayID field if non-nil, zero value otherwise.

### GetGatewayIDOk

`func (o *CreateIPSecSinglePayload) GetGatewayIDOk() (*string, bool)`

GetGatewayIDOk returns a tuple with the GatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayID

`func (o *CreateIPSecSinglePayload) SetGatewayID(v string)`

SetGatewayID sets GatewayID field to given value.


### GetTunnelName

`func (o *CreateIPSecSinglePayload) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *CreateIPSecSinglePayload) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *CreateIPSecSinglePayload) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetP81GatewaySubnets

`func (o *CreateIPSecSinglePayload) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *CreateIPSecSinglePayload) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *CreateIPSecSinglePayload) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *CreateIPSecSinglePayload) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *CreateIPSecSinglePayload) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *CreateIPSecSinglePayload) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetPeakBandwidth

`func (o *CreateIPSecSinglePayload) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *CreateIPSecSinglePayload) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *CreateIPSecSinglePayload) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *CreateIPSecSinglePayload) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetP81ASN

`func (o *CreateIPSecSinglePayload) GetP81ASN() RemoteASN`

GetP81ASN returns the P81ASN field if non-nil, zero value otherwise.

### GetP81ASNOk

`func (o *CreateIPSecSinglePayload) GetP81ASNOk() (*RemoteASN, bool)`

GetP81ASNOk returns a tuple with the P81ASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81ASN

`func (o *CreateIPSecSinglePayload) SetP81ASN(v RemoteASN)`

SetP81ASN sets P81ASN field to given value.

### HasP81ASN

`func (o *CreateIPSecSinglePayload) HasP81ASN() bool`

HasP81ASN returns a boolean if a field has been set.

### GetFeatures

`func (o *CreateIPSecSinglePayload) GetFeatures() IPSecSharedSettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateIPSecSinglePayload) GetFeaturesOk() (*IPSecSharedSettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateIPSecSinglePayload) SetFeatures(v IPSecSharedSettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateIPSecSinglePayload) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetKeyExchange

`func (o *CreateIPSecSinglePayload) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *CreateIPSecSinglePayload) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *CreateIPSecSinglePayload) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *CreateIPSecSinglePayload) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *CreateIPSecSinglePayload) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *CreateIPSecSinglePayload) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *CreateIPSecSinglePayload) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *CreateIPSecSinglePayload) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *CreateIPSecSinglePayload) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *CreateIPSecSinglePayload) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *CreateIPSecSinglePayload) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *CreateIPSecSinglePayload) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *CreateIPSecSinglePayload) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *CreateIPSecSinglePayload) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *CreateIPSecSinglePayload) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *CreateIPSecSinglePayload) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *CreateIPSecSinglePayload) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *CreateIPSecSinglePayload) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *CreateIPSecSinglePayload) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *CreateIPSecSinglePayload) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *CreateIPSecSinglePayload) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.


### GetPassphrase

`func (o *CreateIPSecSinglePayload) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CreateIPSecSinglePayload) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CreateIPSecSinglePayload) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetRemotePublicIP

`func (o *CreateIPSecSinglePayload) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *CreateIPSecSinglePayload) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *CreateIPSecSinglePayload) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.


### GetRemoteID

`func (o *CreateIPSecSinglePayload) GetRemoteID() RemoteID`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *CreateIPSecSinglePayload) GetRemoteIDOk() (*RemoteID, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *CreateIPSecSinglePayload) SetRemoteID(v RemoteID)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *CreateIPSecSinglePayload) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


