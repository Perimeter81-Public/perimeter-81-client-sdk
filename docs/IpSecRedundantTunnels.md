# IPSecRedundantTunnels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**TunnelName** | Pointer to **string** | The name of the tunnel | [optional] 
**RegionID** | Pointer to **string** |  | [optional] 
**Tunnel1** | Pointer to [**IPSecRedundantTunnel**](IPSecRedundantTunnel.md) |  | [optional] 
**Tunnel2** | Pointer to [**IPSecRedundantTunnel**](IPSecRedundantTunnel.md) |  | [optional] 
**SharedSettings** | Pointer to [**IPSecSharedSettings**](IPSecSharedSettings.md) |  | [optional] 
**AdvancedSettings** | Pointer to [**IPSecAdvancedSettings**](IPSecAdvancedSettings.md) |  | [optional] 

## Methods

### NewIPSecRedundantTunnels

`func NewIPSecRedundantTunnels(createdAt time.Time, ) *IPSecRedundantTunnels`

NewIPSecRedundantTunnels instantiates a new IPSecRedundantTunnels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecRedundantTunnelsWithDefaults

`func NewIPSecRedundantTunnelsWithDefaults() *IPSecRedundantTunnels`

NewIPSecRedundantTunnelsWithDefaults instantiates a new IPSecRedundantTunnels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IPSecRedundantTunnels) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPSecRedundantTunnels) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPSecRedundantTunnels) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IPSecRedundantTunnels) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IPSecRedundantTunnels) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IPSecRedundantTunnels) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IPSecRedundantTunnels) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTunnelName

`func (o *IPSecRedundantTunnels) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *IPSecRedundantTunnels) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *IPSecRedundantTunnels) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.

### HasTunnelName

`func (o *IPSecRedundantTunnels) HasTunnelName() bool`

HasTunnelName returns a boolean if a field has been set.

### GetRegionID

`func (o *IPSecRedundantTunnels) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *IPSecRedundantTunnels) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *IPSecRedundantTunnels) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.

### HasRegionID

`func (o *IPSecRedundantTunnels) HasRegionID() bool`

HasRegionID returns a boolean if a field has been set.

### GetTunnel1

`func (o *IPSecRedundantTunnels) GetTunnel1() IPSecRedundantTunnel`

GetTunnel1 returns the Tunnel1 field if non-nil, zero value otherwise.

### GetTunnel1Ok

`func (o *IPSecRedundantTunnels) GetTunnel1Ok() (*IPSecRedundantTunnel, bool)`

GetTunnel1Ok returns a tuple with the Tunnel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel1

`func (o *IPSecRedundantTunnels) SetTunnel1(v IPSecRedundantTunnel)`

SetTunnel1 sets Tunnel1 field to given value.

### HasTunnel1

`func (o *IPSecRedundantTunnels) HasTunnel1() bool`

HasTunnel1 returns a boolean if a field has been set.

### GetTunnel2

`func (o *IPSecRedundantTunnels) GetTunnel2() IPSecRedundantTunnel`

GetTunnel2 returns the Tunnel2 field if non-nil, zero value otherwise.

### GetTunnel2Ok

`func (o *IPSecRedundantTunnels) GetTunnel2Ok() (*IPSecRedundantTunnel, bool)`

GetTunnel2Ok returns a tuple with the Tunnel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel2

`func (o *IPSecRedundantTunnels) SetTunnel2(v IPSecRedundantTunnel)`

SetTunnel2 sets Tunnel2 field to given value.

### HasTunnel2

`func (o *IPSecRedundantTunnels) HasTunnel2() bool`

HasTunnel2 returns a boolean if a field has been set.

### GetSharedSettings

`func (o *IPSecRedundantTunnels) GetSharedSettings() IPSecSharedSettings`

GetSharedSettings returns the SharedSettings field if non-nil, zero value otherwise.

### GetSharedSettingsOk

`func (o *IPSecRedundantTunnels) GetSharedSettingsOk() (*IPSecSharedSettings, bool)`

GetSharedSettingsOk returns a tuple with the SharedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSettings

`func (o *IPSecRedundantTunnels) SetSharedSettings(v IPSecSharedSettings)`

SetSharedSettings sets SharedSettings field to given value.

### HasSharedSettings

`func (o *IPSecRedundantTunnels) HasSharedSettings() bool`

HasSharedSettings returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *IPSecRedundantTunnels) GetAdvancedSettings() IPSecAdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *IPSecRedundantTunnels) GetAdvancedSettingsOk() (*IPSecAdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *IPSecRedundantTunnels) SetAdvancedSettings(v IPSecAdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *IPSecRedundantTunnels) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


