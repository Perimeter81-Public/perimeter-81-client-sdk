# IPSecSharedSettingsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**P81GatewaySubnets** | **[]string** |  | 
**RemoteGatewaySubnets** | **[]string** |  | 
**PeakBandwidth** | Pointer to **int32** | Expected peak throughput of the tunnel communication in Mbps. Typical connection will be of 1000Mbps. | [optional] [default to 1000]
**P81ASN** | [**RemoteASN**](RemoteASN.md) |  | 
**Features** | Pointer to [**IPSecSharedSettingsFeatures**](IPSecSharedSettingsFeatures.md) |  | [optional] 

## Methods

### NewIPSecSharedSettingsCreate

`func NewIPSecSharedSettingsCreate(p81GatewaySubnets []string, remoteGatewaySubnets []string, p81ASN RemoteASN, ) *IPSecSharedSettingsCreate`

NewIPSecSharedSettingsCreate instantiates a new IPSecSharedSettingsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecSharedSettingsCreateWithDefaults

`func NewIPSecSharedSettingsCreateWithDefaults() *IPSecSharedSettingsCreate`

NewIPSecSharedSettingsCreateWithDefaults instantiates a new IPSecSharedSettingsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetP81GatewaySubnets

`func (o *IPSecSharedSettingsCreate) GetP81GatewaySubnets() []string`

GetP81GatewaySubnets returns the P81GatewaySubnets field if non-nil, zero value otherwise.

### GetP81GatewaySubnetsOk

`func (o *IPSecSharedSettingsCreate) GetP81GatewaySubnetsOk() (*[]string, bool)`

GetP81GatewaySubnetsOk returns a tuple with the P81GatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GatewaySubnets

`func (o *IPSecSharedSettingsCreate) SetP81GatewaySubnets(v []string)`

SetP81GatewaySubnets sets P81GatewaySubnets field to given value.


### GetRemoteGatewaySubnets

`func (o *IPSecSharedSettingsCreate) GetRemoteGatewaySubnets() []string`

GetRemoteGatewaySubnets returns the RemoteGatewaySubnets field if non-nil, zero value otherwise.

### GetRemoteGatewaySubnetsOk

`func (o *IPSecSharedSettingsCreate) GetRemoteGatewaySubnetsOk() (*[]string, bool)`

GetRemoteGatewaySubnetsOk returns a tuple with the RemoteGatewaySubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGatewaySubnets

`func (o *IPSecSharedSettingsCreate) SetRemoteGatewaySubnets(v []string)`

SetRemoteGatewaySubnets sets RemoteGatewaySubnets field to given value.


### GetPeakBandwidth

`func (o *IPSecSharedSettingsCreate) GetPeakBandwidth() int32`

GetPeakBandwidth returns the PeakBandwidth field if non-nil, zero value otherwise.

### GetPeakBandwidthOk

`func (o *IPSecSharedSettingsCreate) GetPeakBandwidthOk() (*int32, bool)`

GetPeakBandwidthOk returns a tuple with the PeakBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakBandwidth

`func (o *IPSecSharedSettingsCreate) SetPeakBandwidth(v int32)`

SetPeakBandwidth sets PeakBandwidth field to given value.

### HasPeakBandwidth

`func (o *IPSecSharedSettingsCreate) HasPeakBandwidth() bool`

HasPeakBandwidth returns a boolean if a field has been set.

### GetP81ASN

`func (o *IPSecSharedSettingsCreate) GetP81ASN() RemoteASN`

GetP81ASN returns the P81ASN field if non-nil, zero value otherwise.

### GetP81ASNOk

`func (o *IPSecSharedSettingsCreate) GetP81ASNOk() (*RemoteASN, bool)`

GetP81ASNOk returns a tuple with the P81ASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81ASN

`func (o *IPSecSharedSettingsCreate) SetP81ASN(v RemoteASN)`

SetP81ASN sets P81ASN field to given value.


### GetFeatures

`func (o *IPSecSharedSettingsCreate) GetFeatures() IPSecSharedSettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *IPSecSharedSettingsCreate) GetFeaturesOk() (*IPSecSharedSettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *IPSecSharedSettingsCreate) SetFeatures(v IPSecSharedSettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *IPSecSharedSettingsCreate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


