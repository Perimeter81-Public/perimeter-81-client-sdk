# EnhancedIPSecSharedSettingsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**P81GatewaySubnets** | **[]string** |  | 
**RemoteGatewaySubnets** | **[]string** |  | 
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**P81ASN** | Pointer to [**RemoteASN**](RemoteASN.md) |  | [optional] 
**Features** | [**NetworkFeaturesCreate**](NetworkFeaturesCreate.md) |  | 
**LeftASN** | [**RemoteASN**](RemoteASN.md) |  | 

## Methods

### NewEnhancedIPSecSharedSettingsCreate

`func NewEnhancedIPSecSharedSettingsCreate(p81GatewaySubnets []string, remoteGatewaySubnets []string, features NetworkFeaturesCreate, leftASN RemoteASN, ) *EnhancedIPSecSharedSettingsCreate`

NewEnhancedIPSecSharedSettingsCreate instantiates a new EnhancedIPSecSharedSettingsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedIPSecSharedSettingsCreateWithDefaults

`func NewEnhancedIPSecSharedSettingsCreateWithDefaults() *EnhancedIPSecSharedSettingsCreate`

NewEnhancedIPSecSharedSettingsCreateWithDefaults instantiates a new EnhancedIPSecSharedSettingsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetP81GatewaySubnets

`func (o *EnhancedIPSecSharedSettingsCreate) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *EnhancedIPSecSharedSettingsCreate) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *EnhancedIPSecSharedSettingsCreate) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *EnhancedIPSecSharedSettingsCreate) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *EnhancedIPSecSharedSettingsCreate) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *EnhancedIPSecSharedSettingsCreate) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetPeakBandwidth

`func (o *EnhancedIPSecSharedSettingsCreate) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *EnhancedIPSecSharedSettingsCreate) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *EnhancedIPSecSharedSettingsCreate) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *EnhancedIPSecSharedSettingsCreate) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetP81ASN

`func (o *EnhancedIPSecSharedSettingsCreate) GetP81ASN() RemoteASN`

GetP81ASN returns the P81ASN field if non-nil, zero value otherwise.

### GetP81ASNOk

`func (o *EnhancedIPSecSharedSettingsCreate) GetP81ASNOk() (*RemoteASN, bool)`

GetP81ASNOk returns a tuple with the P81ASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81ASN

`func (o *EnhancedIPSecSharedSettingsCreate) SetP81ASN(v RemoteASN)`

SetP81ASN sets P81ASN field to given value.

### HasP81ASN

`func (o *EnhancedIPSecSharedSettingsCreate) HasP81ASN() bool`

HasP81ASN returns a boolean if a field has been set.

### GetFeatures

`func (o *EnhancedIPSecSharedSettingsCreate) GetFeatures() NetworkFeaturesCreate`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *EnhancedIPSecSharedSettingsCreate) GetFeaturesOk() (*NetworkFeaturesCreate, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *EnhancedIPSecSharedSettingsCreate) SetFeatures(v NetworkFeaturesCreate)`

SetFeatures sets Features field to given value.


### GetLeftASN

`func (o *EnhancedIPSecSharedSettingsCreate) GetLeftASN() RemoteASN`

GetLeftASN returns the LeftASN field if non-nil, zero value otherwise.

### GetLeftASNOk

`func (o *EnhancedIPSecSharedSettingsCreate) GetLeftASNOk() (*RemoteASN, bool)`

GetLeftASNOk returns a tuple with the LeftASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftASN

`func (o *EnhancedIPSecSharedSettingsCreate) SetLeftASN(v RemoteASN)`

SetLeftASN sets LeftASN field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


