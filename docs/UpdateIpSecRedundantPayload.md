# UpdateIPSecRedundantPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tunnel1** | Pointer to [**IPSecRedundantTunnelUpdatePayload**](IPSecRedundantTunnelUpdatePayload.md) |  | [optional] 
**Tunnel2** | Pointer to [**IPSecRedundantTunnelUpdatePayload**](IPSecRedundantTunnelUpdatePayload.md) |  | [optional] 
**SharedSettings** | [**IPSecSharedSettings**](IPSecSharedSettings.md) |  | 
**AdvancedSettings** | [**IPSecAdvancedSettings**](IPSecAdvancedSettings.md) |  | 

## Methods

### NewUpdateIPSecRedundantPayload

`func NewUpdateIPSecRedundantPayload(sharedSettings IPSecSharedSettings, advancedSettings IPSecAdvancedSettings, ) *UpdateIPSecRedundantPayload`

NewUpdateIPSecRedundantPayload instantiates a new UpdateIPSecRedundantPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIPSecRedundantPayloadWithDefaults

`func NewUpdateIPSecRedundantPayloadWithDefaults() *UpdateIPSecRedundantPayload`

NewUpdateIPSecRedundantPayloadWithDefaults instantiates a new UpdateIPSecRedundantPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnel1

`func (o *UpdateIPSecRedundantPayload) GetTunnel1() IPSecRedundantTunnelUpdatePayload`

GetTunnel1 returns the Tunnel1 field if non-nil, zero value otherwise.

### GetTunnel1Ok

`func (o *UpdateIPSecRedundantPayload) GetTunnel1Ok() (*IPSecRedundantTunnelUpdatePayload, bool)`

GetTunnel1Ok returns a tuple with the Tunnel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel1

`func (o *UpdateIPSecRedundantPayload) SetTunnel1(v IPSecRedundantTunnelUpdatePayload)`

SetTunnel1 sets Tunnel1 field to given value.

### HasTunnel1

`func (o *UpdateIPSecRedundantPayload) HasTunnel1() bool`

HasTunnel1 returns a boolean if a field has been set.

### GetTunnel2

`func (o *UpdateIPSecRedundantPayload) GetTunnel2() IPSecRedundantTunnelUpdatePayload`

GetTunnel2 returns the Tunnel2 field if non-nil, zero value otherwise.

### GetTunnel2Ok

`func (o *UpdateIPSecRedundantPayload) GetTunnel2Ok() (*IPSecRedundantTunnelUpdatePayload, bool)`

GetTunnel2Ok returns a tuple with the Tunnel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel2

`func (o *UpdateIPSecRedundantPayload) SetTunnel2(v IPSecRedundantTunnelUpdatePayload)`

SetTunnel2 sets Tunnel2 field to given value.

### HasTunnel2

`func (o *UpdateIPSecRedundantPayload) HasTunnel2() bool`

HasTunnel2 returns a boolean if a field has been set.

### GetSharedSettings

`func (o *UpdateIPSecRedundantPayload) GetSharedSettings() IPSecSharedSettings`

GetSharedSettings returns the SharedSettings field if non-nil, zero value otherwise.

### GetSharedSettingsOk

`func (o *UpdateIPSecRedundantPayload) GetSharedSettingsOk() (*IPSecSharedSettings, bool)`

GetSharedSettingsOk returns a tuple with the SharedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSettings

`func (o *UpdateIPSecRedundantPayload) SetSharedSettings(v IPSecSharedSettings)`

SetSharedSettings sets SharedSettings field to given value.


### GetAdvancedSettings

`func (o *UpdateIPSecRedundantPayload) GetAdvancedSettings() IPSecAdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *UpdateIPSecRedundantPayload) GetAdvancedSettingsOk() (*IPSecAdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *UpdateIPSecRedundantPayload) SetAdvancedSettings(v IPSecAdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


