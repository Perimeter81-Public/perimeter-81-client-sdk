# IPSecSingleTunnel

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
**Passphrase** | Pointer to **string** |  | [optional] 
**RemotePublicIP** | Pointer to **string** |  | [optional] 
**RemoteID** | Pointer to [**RemoteID**](RemoteID.md) |  | [optional] 
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "ipsec"]
**TunnelID** | Pointer to **string** |  | [optional] 

## Methods

### NewIPSecSingleTunnel

`func NewIPSecSingleTunnel(regionID string, gatewayID string, tunnelName string, p81GatewaySubnets []string, remoteGatewaySubnets []string, keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfig, phase2 IPSecPhaseConfig, createdAt time.Time, ) *IPSecSingleTunnel`

NewIPSecSingleTunnel instantiates a new IPSecSingleTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecSingleTunnelWithDefaults

`func NewIPSecSingleTunnelWithDefaults() *IPSecSingleTunnel`

NewIPSecSingleTunnelWithDefaults instantiates a new IPSecSingleTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionID

`func (o *IPSecSingleTunnel) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *IPSecSingleTunnel) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *IPSecSingleTunnel) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetGatewayID

`func (o *IPSecSingleTunnel) GetGatewayID() string`

GetGatewayID returns the GatewayID field if non-nil, zero value otherwise.

### GetGatewayIDOk

`func (o *IPSecSingleTunnel) GetGatewayIDOk() (*string, bool)`

GetGatewayIDOk returns a tuple with the GatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayID

`func (o *IPSecSingleTunnel) SetGatewayID(v string)`

SetGatewayID sets GatewayID field to given value.


### GetTunnelName

`func (o *IPSecSingleTunnel) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *IPSecSingleTunnel) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *IPSecSingleTunnel) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetP81GatewaySubnets

`func (o *IPSecSingleTunnel) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *IPSecSingleTunnel) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *IPSecSingleTunnel) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *IPSecSingleTunnel) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *IPSecSingleTunnel) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *IPSecSingleTunnel) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetPeakBandwidth

`func (o *IPSecSingleTunnel) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *IPSecSingleTunnel) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *IPSecSingleTunnel) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *IPSecSingleTunnel) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetP81ASN

`func (o *IPSecSingleTunnel) GetP81ASN() RemoteASN`

GetP81ASN returns the P81ASN field if non-nil, zero value otherwise.

### GetP81ASNOk

`func (o *IPSecSingleTunnel) GetP81ASNOk() (*RemoteASN, bool)`

GetP81ASNOk returns a tuple with the P81ASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81ASN

`func (o *IPSecSingleTunnel) SetP81ASN(v RemoteASN)`

SetP81ASN sets P81ASN field to given value.

### HasP81ASN

`func (o *IPSecSingleTunnel) HasP81ASN() bool`

HasP81ASN returns a boolean if a field has been set.

### GetFeatures

`func (o *IPSecSingleTunnel) GetFeatures() IPSecSharedSettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *IPSecSingleTunnel) GetFeaturesOk() (*IPSecSharedSettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *IPSecSingleTunnel) SetFeatures(v IPSecSharedSettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *IPSecSingleTunnel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetKeyExchange

`func (o *IPSecSingleTunnel) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *IPSecSingleTunnel) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *IPSecSingleTunnel) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *IPSecSingleTunnel) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *IPSecSingleTunnel) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *IPSecSingleTunnel) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *IPSecSingleTunnel) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *IPSecSingleTunnel) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *IPSecSingleTunnel) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *IPSecSingleTunnel) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *IPSecSingleTunnel) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *IPSecSingleTunnel) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *IPSecSingleTunnel) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *IPSecSingleTunnel) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *IPSecSingleTunnel) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *IPSecSingleTunnel) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *IPSecSingleTunnel) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *IPSecSingleTunnel) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *IPSecSingleTunnel) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *IPSecSingleTunnel) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *IPSecSingleTunnel) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.


### GetPassphrase

`func (o *IPSecSingleTunnel) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *IPSecSingleTunnel) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *IPSecSingleTunnel) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *IPSecSingleTunnel) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetRemotePublicIP

`func (o *IPSecSingleTunnel) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *IPSecSingleTunnel) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *IPSecSingleTunnel) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.

### HasRemotePublicIP

`func (o *IPSecSingleTunnel) HasRemotePublicIP() bool`

HasRemotePublicIP returns a boolean if a field has been set.

### GetRemoteID

`func (o *IPSecSingleTunnel) GetRemoteID() RemoteID`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *IPSecSingleTunnel) GetRemoteIDOk() (*RemoteID, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *IPSecSingleTunnel) SetRemoteID(v RemoteID)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *IPSecSingleTunnel) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPSecSingleTunnel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPSecSingleTunnel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPSecSingleTunnel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IPSecSingleTunnel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IPSecSingleTunnel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IPSecSingleTunnel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IPSecSingleTunnel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *IPSecSingleTunnel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPSecSingleTunnel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPSecSingleTunnel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IPSecSingleTunnel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTunnelID

`func (o *IPSecSingleTunnel) GetTunnelID() string`

GetTunnelID returns the TunnelID field if non-nil, zero value otherwise.

### GetTunnelIDOk

`func (o *IPSecSingleTunnel) GetTunnelIDOk() (*string, bool)`

GetTunnelIDOk returns a tuple with the TunnelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelID

`func (o *IPSecSingleTunnel) SetTunnelID(v string)`

SetTunnelID sets TunnelID field to given value.

### HasTunnelID

`func (o *IPSecSingleTunnel) HasTunnelID() bool`

HasTunnelID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


