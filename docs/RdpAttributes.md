# RdpAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientType** | Pointer to **string** | RDP client type | [optional] [default to "web"]
**DisableClipboard** | Pointer to **bool** |  | [optional] [default to false]
**DisablePrinting** | Pointer to **bool** |  | [optional] [default to false]
**IgnoreTLSVerification** | Pointer to **bool** |  | [optional] [default to true]
**AdminConsole** | Pointer to **bool** |  | [optional] [default to false]
**MaxConnections** | Pointer to [**RdpAttributesMaxConnections**](RdpAttributesMaxConnections.md) |  | [optional] 
**RdpSecurity** | Pointer to **string** |  | [optional] 
**EnableMultiMonitors** | Pointer to **bool** |  | [optional] 

## Methods

### NewRdpAttributes

`func NewRdpAttributes() *RdpAttributes`

NewRdpAttributes instantiates a new RdpAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdpAttributesWithDefaults

`func NewRdpAttributesWithDefaults() *RdpAttributes`

NewRdpAttributesWithDefaults instantiates a new RdpAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientType

`func (o *RdpAttributes) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *RdpAttributes) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *RdpAttributes) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *RdpAttributes) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDisableClipboard

`func (o *RdpAttributes) GetDisableClipboard() bool`

GetDisableClipboard returns the DisableClipboard field if non-nil, zero value otherwise.

### GetDisableClipboardOk

`func (o *RdpAttributes) GetDisableClipboardOk() (*bool, bool)`

GetDisableClipboardOk returns a tuple with the DisableClipboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableClipboard

`func (o *RdpAttributes) SetDisableClipboard(v bool)`

SetDisableClipboard sets DisableClipboard field to given value.

### HasDisableClipboard

`func (o *RdpAttributes) HasDisableClipboard() bool`

HasDisableClipboard returns a boolean if a field has been set.

### GetDisablePrinting

`func (o *RdpAttributes) GetDisablePrinting() bool`

GetDisablePrinting returns the DisablePrinting field if non-nil, zero value otherwise.

### GetDisablePrintingOk

`func (o *RdpAttributes) GetDisablePrintingOk() (*bool, bool)`

GetDisablePrintingOk returns a tuple with the DisablePrinting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePrinting

`func (o *RdpAttributes) SetDisablePrinting(v bool)`

SetDisablePrinting sets DisablePrinting field to given value.

### HasDisablePrinting

`func (o *RdpAttributes) HasDisablePrinting() bool`

HasDisablePrinting returns a boolean if a field has been set.

### GetIgnoreTLSVerification

`func (o *RdpAttributes) GetIgnoreTLSVerification() bool`

GetIgnoreTLSVerification returns the IgnoreTLSVerification field if non-nil, zero value otherwise.

### GetIgnoreTLSVerificationOk

`func (o *RdpAttributes) GetIgnoreTLSVerificationOk() (*bool, bool)`

GetIgnoreTLSVerificationOk returns a tuple with the IgnoreTLSVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTLSVerification

`func (o *RdpAttributes) SetIgnoreTLSVerification(v bool)`

SetIgnoreTLSVerification sets IgnoreTLSVerification field to given value.

### HasIgnoreTLSVerification

`func (o *RdpAttributes) HasIgnoreTLSVerification() bool`

HasIgnoreTLSVerification returns a boolean if a field has been set.

### GetAdminConsole

`func (o *RdpAttributes) GetAdminConsole() bool`

GetAdminConsole returns the AdminConsole field if non-nil, zero value otherwise.

### GetAdminConsoleOk

`func (o *RdpAttributes) GetAdminConsoleOk() (*bool, bool)`

GetAdminConsoleOk returns a tuple with the AdminConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminConsole

`func (o *RdpAttributes) SetAdminConsole(v bool)`

SetAdminConsole sets AdminConsole field to given value.

### HasAdminConsole

`func (o *RdpAttributes) HasAdminConsole() bool`

HasAdminConsole returns a boolean if a field has been set.

### GetMaxConnections

`func (o *RdpAttributes) GetMaxConnections() RdpAttributesMaxConnections`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *RdpAttributes) GetMaxConnectionsOk() (*RdpAttributesMaxConnections, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *RdpAttributes) SetMaxConnections(v RdpAttributesMaxConnections)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *RdpAttributes) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetRdpSecurity

`func (o *RdpAttributes) GetRdpSecurity() string`

GetRdpSecurity returns the RdpSecurity field if non-nil, zero value otherwise.

### GetRdpSecurityOk

`func (o *RdpAttributes) GetRdpSecurityOk() (*string, bool)`

GetRdpSecurityOk returns a tuple with the RdpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpSecurity

`func (o *RdpAttributes) SetRdpSecurity(v string)`

SetRdpSecurity sets RdpSecurity field to given value.

### HasRdpSecurity

`func (o *RdpAttributes) HasRdpSecurity() bool`

HasRdpSecurity returns a boolean if a field has been set.

### GetEnableMultiMonitors

`func (o *RdpAttributes) GetEnableMultiMonitors() bool`

GetEnableMultiMonitors returns the EnableMultiMonitors field if non-nil, zero value otherwise.

### GetEnableMultiMonitorsOk

`func (o *RdpAttributes) GetEnableMultiMonitorsOk() (*bool, bool)`

GetEnableMultiMonitorsOk returns a tuple with the EnableMultiMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiMonitors

`func (o *RdpAttributes) SetEnableMultiMonitors(v bool)`

SetEnableMultiMonitors sets EnableMultiMonitors field to given value.

### HasEnableMultiMonitors

`func (o *RdpAttributes) HasEnableMultiMonitors() bool`

HasEnableMultiMonitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


