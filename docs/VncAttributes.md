# VncAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConnections** | **int32** |  | 
**DisableClipboard** | Pointer to **bool** |  | [optional] 

## Methods

### NewVncAttributes

`func NewVncAttributes(maxConnections int32, ) *VncAttributes`

NewVncAttributes instantiates a new VncAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVncAttributesWithDefaults

`func NewVncAttributesWithDefaults() *VncAttributes`

NewVncAttributesWithDefaults instantiates a new VncAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConnections

`func (o *VncAttributes) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *VncAttributes) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *VncAttributes) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.


### GetDisableClipboard

`func (o *VncAttributes) GetDisableClipboard() bool`

GetDisableClipboard returns the DisableClipboard field if non-nil, zero value otherwise.

### GetDisableClipboardOk

`func (o *VncAttributes) GetDisableClipboardOk() (*bool, bool)`

GetDisableClipboardOk returns a tuple with the DisableClipboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableClipboard

`func (o *VncAttributes) SetDisableClipboard(v bool)`

SetDisableClipboard sets DisableClipboard field to given value.

### HasDisableClipboard

`func (o *VncAttributes) HasDisableClipboard() bool`

HasDisableClipboard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


