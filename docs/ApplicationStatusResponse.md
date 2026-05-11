# ApplicationStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Shows whether the operation is completed | [optional] 
**Result** | Pointer to [**ApplicationStatusResponseResult**](ApplicationStatusResponseResult.md) |  | [optional] 

## Methods

### NewApplicationStatusResponse

`func NewApplicationStatusResponse() *ApplicationStatusResponse`

NewApplicationStatusResponse instantiates a new ApplicationStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStatusResponseWithDefaults

`func NewApplicationStatusResponseWithDefaults() *ApplicationStatusResponse`

NewApplicationStatusResponseWithDefaults instantiates a new ApplicationStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *ApplicationStatusResponse) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *ApplicationStatusResponse) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *ApplicationStatusResponse) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *ApplicationStatusResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetResult

`func (o *ApplicationStatusResponse) GetResult() ApplicationStatusResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApplicationStatusResponse) GetResultOk() (*ApplicationStatusResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApplicationStatusResponse) SetResult(v ApplicationStatusResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApplicationStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


