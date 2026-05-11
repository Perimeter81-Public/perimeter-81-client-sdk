# IPSecRedundantTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passphrase** | **string** |  | 
**P81GWInternalIP** | **string** |  | 
**RemoteGWInternalIP** | **string** |  | 
**RemotePublicIP** | **string** |  | 
**RemoteASN** | [**RemoteASN**](RemoteASN.md) |  | 
**RemoteID** | [**RemoteID**](RemoteID.md) |  | 
**GatewayID** | **string** |  | 
**TunnelID** | Pointer to **string** |  | [optional] 

## Methods

### NewIPSecRedundantTunnel

`func NewIPSecRedundantTunnel(passphrase string, p81GWInternalIP string, remoteGWInternalIP string, remotePublicIP string, remoteASN RemoteASN, remoteID RemoteID, gatewayID string, ) *IPSecRedundantTunnel`

NewIPSecRedundantTunnel instantiates a new IPSecRedundantTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecRedundantTunnelWithDefaults

`func NewIPSecRedundantTunnelWithDefaults() *IPSecRedundantTunnel`

NewIPSecRedundantTunnelWithDefaults instantiates a new IPSecRedundantTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassphrase

`func (o *IPSecRedundantTunnel) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *IPSecRedundantTunnel) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *IPSecRedundantTunnel) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetP81GWInternalIP

`func (o *IPSecRedundantTunnel) GetP81GWInternalIP() string`

GetP81GWInternalIP returns the P81GWInternalIP field if non-nil, zero value otherwise.

### GetP81GWInternalIPOk

`func (o *IPSecRedundantTunnel) GetP81GWInternalIPOk() (*string, bool)`

GetP81GWInternalIPOk returns a tuple with the P81GWInternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GWInternalIP

`func (o *IPSecRedundantTunnel) SetP81GWInternalIP(v string)`

SetP81GWInternalIP sets P81GWInternalIP field to given value.


### GetRemoteGWInternalIP

`func (o *IPSecRedundantTunnel) GetRemoteGWInternalIP() string`

GetRemoteGWInternalIP returns the RemoteGWInternalIP field if non-nil, zero value otherwise.

### GetRemoteGWInternalIPOk

`func (o *IPSecRedundantTunnel) GetRemoteGWInternalIPOk() (*string, bool)`

GetRemoteGWInternalIPOk returns a tuple with the RemoteGWInternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGWInternalIP

`func (o *IPSecRedundantTunnel) SetRemoteGWInternalIP(v string)`

SetRemoteGWInternalIP sets RemoteGWInternalIP field to given value.


### GetRemotePublicIP

`func (o *IPSecRedundantTunnel) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *IPSecRedundantTunnel) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *IPSecRedundantTunnel) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.


### GetRemoteASN

`func (o *IPSecRedundantTunnel) GetRemoteASN() RemoteASN`

GetRemoteASN returns the RemoteASN field if non-nil, zero value otherwise.

### GetRemoteASNOk

`func (o *IPSecRedundantTunnel) GetRemoteASNOk() (*RemoteASN, bool)`

GetRemoteASNOk returns a tuple with the RemoteASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteASN

`func (o *IPSecRedundantTunnel) SetRemoteASN(v RemoteASN)`

SetRemoteASN sets RemoteASN field to given value.


### GetRemoteID

`func (o *IPSecRedundantTunnel) GetRemoteID() RemoteID`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *IPSecRedundantTunnel) GetRemoteIDOk() (*RemoteID, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *IPSecRedundantTunnel) SetRemoteID(v RemoteID)`

SetRemoteID sets RemoteID field to given value.


### GetGatewayID

`func (o *IPSecRedundantTunnel) GetGatewayID() string`

GetGatewayID returns the GatewayID field if non-nil, zero value otherwise.

### GetGatewayIDOk

`func (o *IPSecRedundantTunnel) GetGatewayIDOk() (*string, bool)`

GetGatewayIDOk returns a tuple with the GatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayID

`func (o *IPSecRedundantTunnel) SetGatewayID(v string)`

SetGatewayID sets GatewayID field to given value.


### GetTunnelID

`func (o *IPSecRedundantTunnel) GetTunnelID() string`

GetTunnelID returns the TunnelID field if non-nil, zero value otherwise.

### GetTunnelIDOk

`func (o *IPSecRedundantTunnel) GetTunnelIDOk() (*string, bool)`

GetTunnelIDOk returns a tuple with the TunnelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelID

`func (o *IPSecRedundantTunnel) SetTunnelID(v string)`

SetTunnelID sets TunnelID field to given value.

### HasTunnelID

`func (o *IPSecRedundantTunnel) HasTunnelID() bool`

HasTunnelID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


