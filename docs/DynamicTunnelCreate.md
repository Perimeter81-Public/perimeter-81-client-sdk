# DynamicTunnelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelName** | **string** | Name of the tunnel | 
**Description** | Pointer to **string** | Optional description for the tunnel | [optional] 
**Tunnels** | [**[]DynamicTunnelDetails**](DynamicTunnelDetails.md) | List of tunnels | 
**SharedSettings** | [**EnhancedIPSecSharedSettingsCreate**](EnhancedIPSecSharedSettingsCreate.md) |  | 
**AdvancedSettings** | [**IPSecAdvancedSettingsV23**](IPSecAdvancedSettingsV23.md) |  | 

## Methods

### NewDynamicTunnelCreate

`func NewDynamicTunnelCreate(tunnelName string, tunnels []DynamicTunnelDetails, sharedSettings EnhancedIPSecSharedSettingsCreate, advancedSettings IPSecAdvancedSettingsV23, ) *DynamicTunnelCreate`

NewDynamicTunnelCreate instantiates a new DynamicTunnelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicTunnelCreateWithDefaults

`func NewDynamicTunnelCreateWithDefaults() *DynamicTunnelCreate`

NewDynamicTunnelCreateWithDefaults instantiates a new DynamicTunnelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnelName

`func (o *DynamicTunnelCreate) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *DynamicTunnelCreate) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *DynamicTunnelCreate) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetDescription

`func (o *DynamicTunnelCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicTunnelCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicTunnelCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicTunnelCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTunnels

`func (o *DynamicTunnelCreate) GetTunnels() []DynamicTunnelDetails`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *DynamicTunnelCreate) GetTunnelsOk() (*[]DynamicTunnelDetails, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *DynamicTunnelCreate) SetTunnels(v []DynamicTunnelDetails)`

SetTunnels sets Tunnels field to given value.


### GetSharedSettings

`func (o *DynamicTunnelCreate) GetSharedSettings() EnhancedIPSecSharedSettingsCreate`

GetSharedSettings returns the SharedSettings field if non-nil, zero value otherwise.

### GetSharedSettingsOk

`func (o *DynamicTunnelCreate) GetSharedSettingsOk() (*EnhancedIPSecSharedSettingsCreate, bool)`

GetSharedSettingsOk returns a tuple with the SharedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSettings

`func (o *DynamicTunnelCreate) SetSharedSettings(v EnhancedIPSecSharedSettingsCreate)`

SetSharedSettings sets SharedSettings field to given value.


### GetAdvancedSettings

`func (o *DynamicTunnelCreate) GetAdvancedSettings() IPSecAdvancedSettingsV23`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *DynamicTunnelCreate) GetAdvancedSettingsOk() (*IPSecAdvancedSettingsV23, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *DynamicTunnelCreate) SetAdvancedSettings(v IPSecAdvancedSettingsV23)`

SetAdvancedSettings sets AdvancedSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


