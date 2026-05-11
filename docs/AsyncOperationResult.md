# AsyncOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **[]string** | Reasons list. | [optional] 

## Methods

### NewAsyncOperationResult

`func NewAsyncOperationResult() *AsyncOperationResult`

NewAsyncOperationResult instantiates a new AsyncOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncOperationResultWithDefaults

`func NewAsyncOperationResultWithDefaults() *AsyncOperationResult`

NewAsyncOperationResultWithDefaults instantiates a new AsyncOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *AsyncOperationResult) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AsyncOperationResult) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AsyncOperationResult) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AsyncOperationResult) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStatusCode

`func (o *AsyncOperationResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AsyncOperationResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AsyncOperationResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AsyncOperationResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetReason

`func (o *AsyncOperationResult) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AsyncOperationResult) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AsyncOperationResult) SetReason(v []string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AsyncOperationResult) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


