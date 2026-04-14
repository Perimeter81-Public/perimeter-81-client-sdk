# CreateRegionInNetworkPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HarmonySaseRegionId** | **string** | Harmony SASE region ID, Take id from GET Harmony sases regions endpoint (/networks/enhanced/harmony-sase-regions). | 
**Idle** | **bool** | Create the gateway as disabled if true. | [default to true]
**ScaleUnits** | Pointer to **int32** | Number of scale units for the region. Scale units determine the capacity and performance level of the enhanced region. Higher values provide greater throughput and connection capacity.  | [optional] [default to 1]

## Methods

### NewCreateRegionInNetworkPayload

`func NewCreateRegionInNetworkPayload(harmonySaseRegionId string, idle bool, ) *CreateRegionInNetworkPayload`

NewCreateRegionInNetworkPayload instantiates a new CreateRegionInNetworkPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRegionInNetworkPayloadWithDefaults

`func NewCreateRegionInNetworkPayloadWithDefaults() *CreateRegionInNetworkPayload`

NewCreateRegionInNetworkPayloadWithDefaults instantiates a new CreateRegionInNetworkPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHarmonySaseRegionId

`func (o *CreateRegionInNetworkPayload) GetHarmonySaseRegionId() string`

GetHarmonySaseRegionId returns the HarmonySaseRegionId field if non-nil, zero value otherwise.

### GetHarmonySaseRegionIdOk

`func (o *CreateRegionInNetworkPayload) GetHarmonySaseRegionIdOk() (*string, bool)`

GetHarmonySaseRegionIdOk returns a tuple with the HarmonySaseRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonySaseRegionId

`func (o *CreateRegionInNetworkPayload) SetHarmonySaseRegionId(v string)`

SetHarmonySaseRegionId sets HarmonySaseRegionId field to given value.


### GetIdle

`func (o *CreateRegionInNetworkPayload) GetIdle() bool`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *CreateRegionInNetworkPayload) GetIdleOk() (*bool, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *CreateRegionInNetworkPayload) SetIdle(v bool)`

SetIdle sets Idle field to given value.


### GetScaleUnits

`func (o *CreateRegionInNetworkPayload) GetScaleUnits() int32`

GetScaleUnits returns the ScaleUnits field if non-nil, zero value otherwise.

### GetScaleUnitsOk

`func (o *CreateRegionInNetworkPayload) GetScaleUnitsOk() (*int32, bool)`

GetScaleUnitsOk returns a tuple with the ScaleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUnits

`func (o *CreateRegionInNetworkPayload) SetScaleUnits(v int32)`

SetScaleUnits sets ScaleUnits field to given value.

### HasScaleUnits

`func (o *CreateRegionInNetworkPayload) HasScaleUnits() bool`

HasScaleUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


