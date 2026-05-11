# IPSecSingleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyExchange** | **string** |  | 
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Phase2** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**P81GatewaySubnets** | **[]string** |  | 
**RemoteGatewaySubnets** | **[]string** |  | 
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**P81ASN** | Pointer to [**RemoteASN**](RemoteASN.md) |  | [optional] 
**Features** | Pointer to [**IPSecSharedSettingsFeatures**](IPSecSharedSettingsFeatures.md) |  | [optional] 
**Passphrase** | Pointer to **string** |  | [optional] 
**RemotePublicIP** | Pointer to **string** |  | [optional] 
**RemoteID** | Pointer to [**RemoteID**](RemoteID.md) |  | [optional] 

## Methods

### NewIPSecSingleDetails

`func NewIPSecSingleDetails(keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfig, phase2 IPSecPhaseConfig, p81GatewaySubnets []string, remoteGatewaySubnets []string, ) *IPSecSingleDetails`

NewIPSecSingleDetails instantiates a new IPSecSingleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecSingleDetailsWithDefaults

`func NewIPSecSingleDetailsWithDefaults() *IPSecSingleDetails`

NewIPSecSingleDetailsWithDefaults instantiates a new IPSecSingleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyExchange

`func (o *IPSecSingleDetails) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *IPSecSingleDetails) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *IPSecSingleDetails) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *IPSecSingleDetails) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *IPSecSingleDetails) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *IPSecSingleDetails) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *IPSecSingleDetails) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *IPSecSingleDetails) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *IPSecSingleDetails) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *IPSecSingleDetails) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *IPSecSingleDetails) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *IPSecSingleDetails) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *IPSecSingleDetails) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *IPSecSingleDetails) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *IPSecSingleDetails) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *IPSecSingleDetails) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *IPSecSingleDetails) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *IPSecSingleDetails) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *IPSecSingleDetails) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *IPSecSingleDetails) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *IPSecSingleDetails) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.


### GetP81GatewaySubnets

`func (o *IPSecSingleDetails) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *IPSecSingleDetails) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *IPSecSingleDetails) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *IPSecSingleDetails) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *IPSecSingleDetails) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *IPSecSingleDetails) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetPeakBandwidth

`func (o *IPSecSingleDetails) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *IPSecSingleDetails) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *IPSecSingleDetails) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *IPSecSingleDetails) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetP81ASN

`func (o *IPSecSingleDetails) GetP81ASN() RemoteASN`

GetP81ASN returns the P81ASN field if non-nil, zero value otherwise.

### GetP81ASNOk

`func (o *IPSecSingleDetails) GetP81ASNOk() (*RemoteASN, bool)`

GetP81ASNOk returns a tuple with the P81ASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81ASN

`func (o *IPSecSingleDetails) SetP81ASN(v RemoteASN)`

SetP81ASN sets P81ASN field to given value.

### HasP81ASN

`func (o *IPSecSingleDetails) HasP81ASN() bool`

HasP81ASN returns a boolean if a field has been set.

### GetFeatures

`func (o *IPSecSingleDetails) GetFeatures() IPSecSharedSettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *IPSecSingleDetails) GetFeaturesOk() (*IPSecSharedSettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *IPSecSingleDetails) SetFeatures(v IPSecSharedSettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *IPSecSingleDetails) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPassphrase

`func (o *IPSecSingleDetails) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *IPSecSingleDetails) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *IPSecSingleDetails) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *IPSecSingleDetails) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetRemotePublicIP

`func (o *IPSecSingleDetails) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *IPSecSingleDetails) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *IPSecSingleDetails) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.

### HasRemotePublicIP

`func (o *IPSecSingleDetails) HasRemotePublicIP() bool`

HasRemotePublicIP returns a boolean if a field has been set.

### GetRemoteID

`func (o *IPSecSingleDetails) GetRemoteID() RemoteID`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *IPSecSingleDetails) GetRemoteIDOk() (*RemoteID, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *IPSecSingleDetails) SetRemoteID(v RemoteID)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *IPSecSingleDetails) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


