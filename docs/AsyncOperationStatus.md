# AsyncOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**AsyncOperationResult**](AsyncOperationResult.md) |  | [optional] 

## Methods

### NewAsyncOperationStatus

`func NewAsyncOperationStatus() *AsyncOperationStatus`

NewAsyncOperationStatus instantiates a new AsyncOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncOperationStatusWithDefaults

`func NewAsyncOperationStatusWithDefaults() *AsyncOperationStatus`

NewAsyncOperationStatusWithDefaults instantiates a new AsyncOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *AsyncOperationStatus) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *AsyncOperationStatus) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *AsyncOperationStatus) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *AsyncOperationStatus) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetResult

`func (o *AsyncOperationStatus) GetResult() AsyncOperationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AsyncOperationStatus) GetResultOk() (*AsyncOperationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AsyncOperationStatus) SetResult(v AsyncOperationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *AsyncOperationStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


