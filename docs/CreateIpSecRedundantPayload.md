# CreateIPSecRedundantPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelName** | **string** | The name of the tunnel | 
**RegionID** | **string** |  | 
**Tunnel1** | [**IPSecRedundantTunnelPayload**](IPSecRedundantTunnelPayload.md) |  | 
**Tunnel2** | [**IPSecRedundantTunnelPayload**](IPSecRedundantTunnelPayload.md) |  | 
**SharedSettings** | [**IPSecSharedSettingsCreate**](IPSecSharedSettingsCreate.md) |  | 
**AdvancedSettings** | [**IPSecAdvancedSettings**](IPSecAdvancedSettings.md) |  | 

## Methods

### NewCreateIPSecRedundantPayload

`func NewCreateIPSecRedundantPayload(tunnelName string, regionID string, tunnel1 IPSecRedundantTunnelPayload, tunnel2 IPSecRedundantTunnelPayload, sharedSettings IPSecSharedSettingsCreate, advancedSettings IPSecAdvancedSettings, ) *CreateIPSecRedundantPayload`

NewCreateIPSecRedundantPayload instantiates a new CreateIPSecRedundantPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIPSecRedundantPayloadWithDefaults

`func NewCreateIPSecRedundantPayloadWithDefaults() *CreateIPSecRedundantPayload`

NewCreateIPSecRedundantPayloadWithDefaults instantiates a new CreateIPSecRedundantPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnelName

`func (o *CreateIPSecRedundantPayload) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *CreateIPSecRedundantPayload) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *CreateIPSecRedundantPayload) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetRegionID

`func (o *CreateIPSecRedundantPayload) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *CreateIPSecRedundantPayload) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *CreateIPSecRedundantPayload) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetTunnel1

`func (o *CreateIPSecRedundantPayload) GetTunnel1() IPSecRedundantTunnelPayload`

GetTunnel1 returns the Tunnel1 field if non-nil, zero value otherwise.

### GetTunnel1Ok

`func (o *CreateIPSecRedundantPayload) GetTunnel1Ok() (*IPSecRedundantTunnelPayload, bool)`

GetTunnel1Ok returns a tuple with the Tunnel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel1

`func (o *CreateIPSecRedundantPayload) SetTunnel1(v IPSecRedundantTunnelPayload)`

SetTunnel1 sets Tunnel1 field to given value.


### GetTunnel2

`func (o *CreateIPSecRedundantPayload) GetTunnel2() IPSecRedundantTunnelPayload`

GetTunnel2 returns the Tunnel2 field if non-nil, zero value otherwise.

### GetTunnel2Ok

`func (o *CreateIPSecRedundantPayload) GetTunnel2Ok() (*IPSecRedundantTunnelPayload, bool)`

GetTunnel2Ok returns a tuple with the Tunnel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel2

`func (o *CreateIPSecRedundantPayload) SetTunnel2(v IPSecRedundantTunnelPayload)`

SetTunnel2 sets Tunnel2 field to given value.


### GetSharedSettings

`func (o *CreateIPSecRedundantPayload) GetSharedSettings() IPSecSharedSettingsCreate`

GetSharedSettings returns the SharedSettings field if non-nil, zero value otherwise.

### GetSharedSettingsOk

`func (o *CreateIPSecRedundantPayload) GetSharedSettingsOk() (*IPSecSharedSettingsCreate, bool)`

GetSharedSettingsOk returns a tuple with the SharedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSettings

`func (o *CreateIPSecRedundantPayload) SetSharedSettings(v IPSecSharedSettingsCreate)`

SetSharedSettings sets SharedSettings field to given value.


### GetAdvancedSettings

`func (o *CreateIPSecRedundantPayload) GetAdvancedSettings() IPSecAdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *CreateIPSecRedundantPayload) GetAdvancedSettingsOk() (*IPSecAdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *CreateIPSecRedundantPayload) SetAdvancedSettings(v IPSecAdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


