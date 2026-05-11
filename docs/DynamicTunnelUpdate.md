# DynamicTunnelUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunnelName** | **string** | Name of the tunnel | 
**Description** | Pointer to **string** | Optional description for the tunnel | [optional] 
**AddTunnels** | Pointer to [**[]DynamicTunnelDetails**](DynamicTunnelDetails.md) |  | [optional] 
**UpdateTunnels** | Pointer to [**[]DynamicTunnelUpdateUpdateTunnelsInner**](DynamicTunnelUpdateUpdateTunnelsInner.md) |  | [optional] 
**RemoveTunnels** | Pointer to [**[]DynamicTunnelUpdateRemoveTunnelsInner**](DynamicTunnelUpdateRemoveTunnelsInner.md) |  | [optional] 
**SharedSettings** | Pointer to [**EnhancedIPSecSharedSettingsUpdate**](EnhancedIPSecSharedSettingsUpdate.md) |  | [optional] 
**AdvancedSettings** | Pointer to [**IPSecAdvancedSettingsUpdateV23**](IPSecAdvancedSettingsUpdateV23.md) |  | [optional] 
**RoutingType** | Pointer to [**RoutingType**](RoutingType.md) |  | [optional] [default to ROUTE]

## Methods

### NewDynamicTunnelUpdate

`func NewDynamicTunnelUpdate(tunnelName string, ) *DynamicTunnelUpdate`

NewDynamicTunnelUpdate instantiates a new DynamicTunnelUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicTunnelUpdateWithDefaults

`func NewDynamicTunnelUpdateWithDefaults() *DynamicTunnelUpdate`

NewDynamicTunnelUpdateWithDefaults instantiates a new DynamicTunnelUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnelName

`func (o *DynamicTunnelUpdate) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *DynamicTunnelUpdate) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *DynamicTunnelUpdate) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetDescription

`func (o *DynamicTunnelUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicTunnelUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicTunnelUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicTunnelUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddTunnels

`func (o *DynamicTunnelUpdate) GetAddTunnels() []DynamicTunnelDetails`

GetAddTunnels returns the AddTunnels field if non-nil, zero value otherwise.

### GetAddTunnelsOk

`func (o *DynamicTunnelUpdate) GetAddTunnelsOk() (*[]DynamicTunnelDetails, bool)`

GetAddTunnelsOk returns a tuple with the AddTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTunnels

`func (o *DynamicTunnelUpdate) SetAddTunnels(v []DynamicTunnelDetails)`

SetAddTunnels sets AddTunnels field to given value.

### HasAddTunnels

`func (o *DynamicTunnelUpdate) HasAddTunnels() bool`

HasAddTunnels returns a boolean if a field has been set.

### GetUpdateTunnels

`func (o *DynamicTunnelUpdate) GetUpdateTunnels() []DynamicTunnelUpdateUpdateTunnelsInner`

GetUpdateTunnels returns the UpdateTunnels field if non-nil, zero value otherwise.

### GetUpdateTunnelsOk

`func (o *DynamicTunnelUpdate) GetUpdateTunnelsOk() (*[]DynamicTunnelUpdateUpdateTunnelsInner, bool)`

GetUpdateTunnelsOk returns a tuple with the UpdateTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTunnels

`func (o *DynamicTunnelUpdate) SetUpdateTunnels(v []DynamicTunnelUpdateUpdateTunnelsInner)`

SetUpdateTunnels sets UpdateTunnels field to given value.

### HasUpdateTunnels

`func (o *DynamicTunnelUpdate) HasUpdateTunnels() bool`

HasUpdateTunnels returns a boolean if a field has been set.

### GetRemoveTunnels

`func (o *DynamicTunnelUpdate) GetRemoveTunnels() []DynamicTunnelUpdateRemoveTunnelsInner`

GetRemoveTunnels returns the RemoveTunnels field if non-nil, zero value otherwise.

### GetRemoveTunnelsOk

`func (o *DynamicTunnelUpdate) GetRemoveTunnelsOk() (*[]DynamicTunnelUpdateRemoveTunnelsInner, bool)`

GetRemoveTunnelsOk returns a tuple with the RemoveTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTunnels

`func (o *DynamicTunnelUpdate) SetRemoveTunnels(v []DynamicTunnelUpdateRemoveTunnelsInner)`

SetRemoveTunnels sets RemoveTunnels field to given value.

### HasRemoveTunnels

`func (o *DynamicTunnelUpdate) HasRemoveTunnels() bool`

HasRemoveTunnels returns a boolean if a field has been set.

### GetSharedSettings

`func (o *DynamicTunnelUpdate) GetSharedSettings() EnhancedIPSecSharedSettingsUpdate`

GetSharedSettings returns the SharedSettings field if non-nil, zero value otherwise.

### GetSharedSettingsOk

`func (o *DynamicTunnelUpdate) GetSharedSettingsOk() (*EnhancedIPSecSharedSettingsUpdate, bool)`

GetSharedSettingsOk returns a tuple with the SharedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSettings

`func (o *DynamicTunnelUpdate) SetSharedSettings(v EnhancedIPSecSharedSettingsUpdate)`

SetSharedSettings sets SharedSettings field to given value.

### HasSharedSettings

`func (o *DynamicTunnelUpdate) HasSharedSettings() bool`

HasSharedSettings returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *DynamicTunnelUpdate) GetAdvancedSettings() IPSecAdvancedSettingsUpdateV23`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *DynamicTunnelUpdate) GetAdvancedSettingsOk() (*IPSecAdvancedSettingsUpdateV23, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *DynamicTunnelUpdate) SetAdvancedSettings(v IPSecAdvancedSettingsUpdateV23)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *DynamicTunnelUpdate) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetRoutingType

`func (o *DynamicTunnelUpdate) GetRoutingType() RoutingType`

GetRoutingType returns the RoutingType field if non-nil, zero value otherwise.

### GetRoutingTypeOk

`func (o *DynamicTunnelUpdate) GetRoutingTypeOk() (*RoutingType, bool)`

GetRoutingTypeOk returns a tuple with the RoutingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingType

`func (o *DynamicTunnelUpdate) SetRoutingType(v RoutingType)`

SetRoutingType sets RoutingType field to given value.

### HasRoutingType

`func (o *DynamicTunnelUpdate) HasRoutingType() bool`

HasRoutingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


