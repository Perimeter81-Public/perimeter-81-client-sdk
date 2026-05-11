# EnhancedRegionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HarmonySaseRegionId** | **string** | Harmony SASE region ID, Take id from GET Harmony sases regions endpoint (/networks/enhanced/harmony-sase-regions). | 
**ScaleUnits** | Pointer to **int32** | Number of scale units for the region. Scale units determine the capacity and performance level of the enhanced region. Higher values provide greater throughput and connection capacity.  | [optional] [default to 1]
**Idle** | Pointer to **bool** | Region is disabled for users if true | [optional] [default to true]

## Methods

### NewEnhancedRegionCreate

`func NewEnhancedRegionCreate(harmonySaseRegionId string, ) *EnhancedRegionCreate`

NewEnhancedRegionCreate instantiates a new EnhancedRegionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRegionCreateWithDefaults

`func NewEnhancedRegionCreateWithDefaults() *EnhancedRegionCreate`

NewEnhancedRegionCreateWithDefaults instantiates a new EnhancedRegionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHarmonySaseRegionId

`func (o *EnhancedRegionCreate) GetHarmonySaseRegionId() string`

GetHarmonySaseRegionId returns the HarmonySaseRegionId field if non-nil, zero value otherwise.

### GetHarmonySaseRegionIdOk

`func (o *EnhancedRegionCreate) GetHarmonySaseRegionIdOk() (*string, bool)`

GetHarmonySaseRegionIdOk returns a tuple with the HarmonySaseRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarmonySaseRegionId

`func (o *EnhancedRegionCreate) SetHarmonySaseRegionId(v string)`

SetHarmonySaseRegionId sets HarmonySaseRegionId field to given value.


### GetScaleUnits

`func (o *EnhancedRegionCreate) GetScaleUnits() int32`

GetScaleUnits returns the ScaleUnits field if non-nil, zero value otherwise.

### GetScaleUnitsOk

`func (o *EnhancedRegionCreate) GetScaleUnitsOk() (*int32, bool)`

GetScaleUnitsOk returns a tuple with the ScaleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUnits

`func (o *EnhancedRegionCreate) SetScaleUnits(v int32)`

SetScaleUnits sets ScaleUnits field to given value.

### HasScaleUnits

`func (o *EnhancedRegionCreate) HasScaleUnits() bool`

HasScaleUnits returns a boolean if a field has been set.

### GetIdle

`func (o *EnhancedRegionCreate) GetIdle() bool`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *EnhancedRegionCreate) GetIdleOk() (*bool, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *EnhancedRegionCreate) SetIdle(v bool)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *EnhancedRegionCreate) HasIdle() bool`

HasIdle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


