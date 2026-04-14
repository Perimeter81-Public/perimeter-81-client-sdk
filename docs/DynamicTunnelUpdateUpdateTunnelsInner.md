# DynamicTunnelUpdateUpdateTunnelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** | Authentication type for tunnel (psk for pre-shared key, cert for certificate) | [optional] 
**Passphrase** | Pointer to **string** | Pre-shared key for tunnel authentication (8-64 characters). Required when authType is psk. | [optional] 
**CustomerRootCA** | Pointer to **string** | Customer root certificate authority. Required when authType is cert. | [optional] 
**RemotePublicIP** | Pointer to **string** | Remote gateway public IP address | [optional] 
**RemoteASN** | Pointer to [**ASN**](ASN.md) |  | [optional] 
**RemoteID** | Pointer to **string** | Remote gateway ID | [optional] 
**RoutingType** | Pointer to [**RoutingTypeUpdate**](RoutingTypeUpdate.md) |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewDynamicTunnelUpdateUpdateTunnelsInner

`func NewDynamicTunnelUpdateUpdateTunnelsInner(id string, ) *DynamicTunnelUpdateUpdateTunnelsInner`

NewDynamicTunnelUpdateUpdateTunnelsInner instantiates a new DynamicTunnelUpdateUpdateTunnelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicTunnelUpdateUpdateTunnelsInnerWithDefaults

`func NewDynamicTunnelUpdateUpdateTunnelsInnerWithDefaults() *DynamicTunnelUpdateUpdateTunnelsInner`

NewDynamicTunnelUpdateUpdateTunnelsInnerWithDefaults instantiates a new DynamicTunnelUpdateUpdateTunnelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetPassphrase

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCustomerRootCA

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetCustomerRootCA() string`

GetCustomerRootCA returns the CustomerRootCA field if non-nil, zero value otherwise.

### GetCustomerRootCAOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetCustomerRootCAOk() (*string, bool)`

GetCustomerRootCAOk returns a tuple with the CustomerRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRootCA

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetCustomerRootCA(v string)`

SetCustomerRootCA sets CustomerRootCA field to given value.

### HasCustomerRootCA

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasCustomerRootCA() bool`

HasCustomerRootCA returns a boolean if a field has been set.

### GetRemotePublicIP

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.

### HasRemotePublicIP

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasRemotePublicIP() bool`

HasRemotePublicIP returns a boolean if a field has been set.

### GetRemoteASN

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRemoteASN() ASN`

GetRemoteASN returns the RemoteASN field if non-nil, zero value otherwise.

### GetRemoteASNOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRemoteASNOk() (*ASN, bool)`

GetRemoteASNOk returns a tuple with the RemoteASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteASN

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetRemoteASN(v ASN)`

SetRemoteASN sets RemoteASN field to given value.

### HasRemoteASN

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasRemoteASN() bool`

HasRemoteASN returns a boolean if a field has been set.

### GetRemoteID

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRemoteID() string`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRemoteIDOk() (*string, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetRemoteID(v string)`

SetRemoteID sets RemoteID field to given value.

### HasRemoteID

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasRemoteID() bool`

HasRemoteID returns a boolean if a field has been set.

### GetRoutingType

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRoutingType() RoutingTypeUpdate`

GetRoutingType returns the RoutingType field if non-nil, zero value otherwise.

### GetRoutingTypeOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetRoutingTypeOk() (*RoutingTypeUpdate, bool)`

GetRoutingTypeOk returns a tuple with the RoutingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingType

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetRoutingType(v RoutingTypeUpdate)`

SetRoutingType sets RoutingType field to given value.

### HasRoutingType

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) HasRoutingType() bool`

HasRoutingType returns a boolean if a field has been set.

### GetId

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicTunnelUpdateUpdateTunnelsInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


