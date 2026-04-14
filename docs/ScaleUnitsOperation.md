# ScaleUnitsOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitType** | **string** | Type of scale unit. Standard units provide basic capacity, while cloudSecurity units offer enhanced security features. | 
**ScaleUnitsCount** | **int32** | Number of scale units for the region. Scale units determine the capacity and performance level of the enhanced region. Higher values provide greater throughput and connection capacity.  | 
**Idle** | Pointer to **bool** | Region is disabled for users if true | [optional] [default to true]

## Methods

### NewScaleUnitsOperation

`func NewScaleUnitsOperation(unitType string, scaleUnitsCount int32, ) *ScaleUnitsOperation`

NewScaleUnitsOperation instantiates a new ScaleUnitsOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleUnitsOperationWithDefaults

`func NewScaleUnitsOperationWithDefaults() *ScaleUnitsOperation`

NewScaleUnitsOperationWithDefaults instantiates a new ScaleUnitsOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitType

`func (o *ScaleUnitsOperation) GetUnitType() string`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *ScaleUnitsOperation) GetUnitTypeOk() (*string, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *ScaleUnitsOperation) SetUnitType(v string)`

SetUnitType sets UnitType field to given value.


### GetScaleUnitsCount

`func (o *ScaleUnitsOperation) GetScaleUnitsCount() int32`

GetScaleUnitsCount returns the ScaleUnitsCount field if non-nil, zero value otherwise.

### GetScaleUnitsCountOk

`func (o *ScaleUnitsOperation) GetScaleUnitsCountOk() (*int32, bool)`

GetScaleUnitsCountOk returns a tuple with the ScaleUnitsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUnitsCount

`func (o *ScaleUnitsOperation) SetScaleUnitsCount(v int32)`

SetScaleUnitsCount sets ScaleUnitsCount field to given value.


### GetIdle

`func (o *ScaleUnitsOperation) GetIdle() bool`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *ScaleUnitsOperation) GetIdleOk() (*bool, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *ScaleUnitsOperation) SetIdle(v bool)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *ScaleUnitsOperation) HasIdle() bool`

HasIdle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


