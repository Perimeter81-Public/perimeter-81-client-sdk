# DynamicTunnelDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** | Authentication type for tunnel (psk for pre-shared key, cert for certificate) | [optional] 
**Passphrase** | Pointer to **string** | Pre-shared key for tunnel authentication (8-64 characters). Required when authType is psk. | [optional] 
**CustomerRootCA** | Pointer to **string** | Customer root certificate authority. Required when authType is cert. | [optional] 
**P81GWInternalIP** | Pointer to **string** | Harmony Sase gateway internal IP address | [optional] 
**RemoteGWInternalIP** | Pointer to **string** | Remote gateway internal IP address | [optional] 
**RemotePublicIP** | Pointer to **string** | Remote gateway public IP address | [optional] 
**RemoteASN** | Pointer to [**ASN**](ASN.md) |  | [optional] 
**RemoteID** | Pointer to **string** | Remote gateway ID | [optional] 
**RoutingType** | Pointer to [**RoutingType**](RoutingType.md) |  | [optional] [default to ROUTE]
**RegionID** | **string** | Dynamic tunnel enhanced region ID | 

## Methods

### NewDynamicTunnelDetails

`func NewDynamicTunnelDetails(regionID string, ) *DynamicTunnelDetails`

NewDynamicTunnelDetails instantiates a new DynamicTunnelDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicTunnelDetailsWithDefaults

`func NewDynamicTunnelDetailsWithDefaults() *DynamicTunnelDetails`

NewDynamicTunnelDetailsWithDefaults instantiates a new DynamicTunnelDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *DynamicTunnelDetails) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DynamicTunnelDetails) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DynamicTunnelDetails) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *DynamicTunnelDetails) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetPassphrase

`func (o *DynamicTunnelDetails) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *DynamicTunnelDetails) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *DynamicTunnelDetails) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *DynamicTunnelDetails) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCustomerRootCA

`func (o *DynamicTunnelDetails) GetCustomerRootCA() string`

GetCustomerRootCA returns the CustomerRootCA field if non-nil, zero value otherwise.

### GetCustomerRootCAOk

`func (o *DynamicTunnelDetails) GetCustomerRootCAOk() (*string, bool)`

GetCustomerRootCAOk returns a tuple with the CustomerRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRootCA

`func (o *DynamicTunnelDetails) SetCustomerRootCA(v string)`

SetCustomerRootCA sets CustomerRootCA field to given value.

### HasCustomerRootCA

`func (o *DynamicTunnelDetails) HasCustomerRootCA() bool`

HasCustomerRootCA returns a boolean if a field has been set.

### GetP81GWInternalIP

`func (o *DynamicTunnelDetails) GetP81GWInternalIP() string`

GetP81GWInternalIP returns the P81GWInternalIP field if non-nil, zero value otherwise.

### GetP81GWInternalIPOk

`func (o *DynamicTunnelDetails) GetP81GWInternalIPOk() (*string, bool)`

GetP81GWInternalIPOk returns a tuple with the P81GWInternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GWInternalIP

`func (o *DynamicTunnelDetails) SetP81GWInternalIP(v string)`

SetP81GWInternalIP sets P81GWInternalIP field to given value.

### HasP81GWInternalIP

`func (o *DynamicTunnelDetails) HasP81GWInternalIP() bool`

HasP81GWInternalIP returns a boolean if a field has been set.

### GetRemoteGWInternalIP

`func (o *DynamicTunnelDetails) GetRemoteGWInternalIP() string`

GetRemoteGWInternalIP returns the RemoteGWInternalIP field if non-nil, zero value otherwise.

### GetRemoteGWInternalIPOk

`func (o *DynamicTunnelDetails) GetRemoteGWInternalIPOk() (*string, bool)`

GetRemoteGWInternalIPOk returns a tuple with the RemoteGWInternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGWInternalIP

`func (o *DynamicTunnelDetails) SetRemoteGWInternalIP(v string)`

SetRemoteGWInternalIP sets RemoteGWInternalIP field to given value.

### HasRemoteGWInternalIP

`func (o *DynamicTunnelDetails) HasRemoteGWInternalIP() bool`

HasRemoteGWInternalIP returns a boolean if a field has been set.

### GetRemotePublicIP

`func (o *DynamicTunnelDetails) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *DynamicTunnelDetails) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *DynamicTunnelDetails) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.

### HasRemotePublicIP

`func (o *DynamicTunnelDetails) HasRemotePublicIP() bool`

HasRemotePublicIP returns a boolean if a field has been set.

### GetRemoteASN

`func (o *DynamicTunnelDetails) GetRemoteASN() ASN`

GetRemoteASN returns the RemoteASN field if non-nil, zero value otherwise.

### GetRemoteASNOk

`func (o *DynamicTunnelDetails) GetRemoteASNOk() (*ASN, bool)`

GetRemoteASNOk returns a tuple with the RemoteASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteASN

`func (o *DynamicTunnelDetails) SetRemoteASN(v ASN)`

SetRemoteASN sets RemoteASN field to given value.

### HasRemoteASN

`func (o *DynamicTunnelDetails) HasRemoteASN() bool`

HasRemoteASN returns a boolean if a field has been set.

### GetRemoteID

`func (o *DynamicTunnelDetails) GetRemoteID() string`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *DynamicTunnelDetails) GetRemoteIDOk() (*string, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *DynamicTunnelDetails) SetRemoteID(v string)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *DynamicTunnelDetails) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.

### GetRoutingType

`func (o *DynamicTunnelDetails) GetRoutingType() RoutingType`

GetRoutingType returns the RoutingType field if non-nil, zero value otherwise.

### GetRoutingTypeOk

`func (o *DynamicTunnelDetails) GetRoutingTypeOk() (*RoutingType, bool)`

GetRoutingTypeOk returns a tuple with the RoutingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingType

`func (o *DynamicTunnelDetails) SetRoutingType(v RoutingType)`

SetRoutingType sets RoutingType field to given value.

### HasRoutingType

`func (o *DynamicTunnelDetails) HasRoutingType() bool`

HasRoutingType returns a boolean if a field has been set.

### GetRegionID

`func (o *DynamicTunnelDetails) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *DynamicTunnelDetails) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *DynamicTunnelDetails) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


