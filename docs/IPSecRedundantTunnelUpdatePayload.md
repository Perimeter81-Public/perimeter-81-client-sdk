# IPSecRedundantTunnelUpdatePayload

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
**TunnelId** | **string** |  | 

## Methods

### NewIPSecRedundantTunnelUpdatePayload

`func NewIPSecRedundantTunnelUpdatePayload(passphrase string, p81GWInternalIP string, remoteGWInternalIP string, remotePublicIP string, remoteASN RemoteASN, remoteID RemoteID, gatewayID string, tunnelId string, ) *IPSecRedundantTunnelUpdatePayload`

NewIPSecRedundantTunnelUpdatePayload instantiates a new IPSecRedundantTunnelUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecRedundantTunnelUpdatePayloadWithDefaults

`func NewIPSecRedundantTunnelUpdatePayloadWithDefaults() *IPSecRedundantTunnelUpdatePayload`

NewIPSecRedundantTunnelUpdatePayloadWithDefaults instantiates a new IPSecRedundantTunnelUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassphrase

`func (o *IPSecRedundantTunnelUpdatePayload) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *IPSecRedundantTunnelUpdatePayload) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetP81GWInternalIP

`func (o *IPSecRedundantTunnelUpdatePayload) GetP81GWInternalIP() string`

GetP81GWInternalIP returns the P81GWInternalIP field if non-nil, zero value otherwise.

### GetP81GWInternalIPOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetP81GWInternalIPOk() (*string, bool)`

GetP81GWInternalIPOk returns a tuple with the P81GWInternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP81GWInternalIP

`func (o *IPSecRedundantTunnelUpdatePayload) SetP81GWInternalIP(v string)`

SetP81GWInternalIP sets P81GWInternalIP field to given value.


### GetRemoteGWInternalIP

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemoteGWInternalIP() string`

GetRemoteGWInternalIP returns the RemoteGWInternalIP field if non-nil, zero value otherwise.

### GetRemoteGWInternalIPOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemoteGWInternalIPOk() (*string, bool)`

GetRemoteGWInternalIPOk returns a tuple with the RemoteGWInternalIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGWInternalIP

`func (o *IPSecRedundantTunnelUpdatePayload) SetRemoteGWInternalIP(v string)`

SetRemoteGWInternalIP sets RemoteGWInternalIP field to given value.


### GetRemotePublicIP

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemotePublicIP() string`

GetRemotePublicIP returns the RemotePublicIP field if non-nil, zero value otherwise.

### GetRemotePublicIPOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemotePublicIPOk() (*string, bool)`

GetRemotePublicIPOk returns a tuple with the RemotePublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePublicIP

`func (o *IPSecRedundantTunnelUpdatePayload) SetRemotePublicIP(v string)`

SetRemotePublicIP sets RemotePublicIP field to given value.


### GetRemoteASN

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemoteASN() RemoteASN`

GetRemoteASN returns the RemoteASN field if non-nil, zero value otherwise.

### GetRemoteASNOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemoteASNOk() (*RemoteASN, bool)`

GetRemoteASNOk returns a tuple with the RemoteASN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteASN

`func (o *IPSecRedundantTunnelUpdatePayload) SetRemoteASN(v RemoteASN)`

SetRemoteASN sets RemoteASN field to given value.


### GetRemoteID

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemoteID() RemoteID`

GetRemoteID returns the RemoteID field if non-nil, zero value otherwise.

### GetRemoteIDOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetRemoteIDOk() (*RemoteID, bool)`

GetRemoteIDOk returns a tuple with the RemoteID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteID

`func (o *IPSecRedundantTunnelUpdatePayload) SetRemoteID(v RemoteID)`

SetRemoteID sets RemoteID field to given value.


### GetGatewayID

`func (o *IPSecRedundantTunnelUpdatePayload) GetGatewayID() string`

GetGatewayID returns the GatewayID field if non-nil, zero value otherwise.

### GetGatewayIDOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetGatewayIDOk() (*string, bool)`

GetGatewayIDOk returns a tuple with the GatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayID

`func (o *IPSecRedundantTunnelUpdatePayload) SetGatewayID(v string)`

SetGatewayID sets GatewayID field to given value.


### GetTunnelId

`func (o *IPSecRedundantTunnelUpdatePayload) GetTunnelId() string`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *IPSecRedundantTunnelUpdatePayload) GetTunnelIdOk() (*string, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *IPSecRedundantTunnelUpdatePayload) SetTunnelId(v string)`

SetTunnelId sets TunnelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


