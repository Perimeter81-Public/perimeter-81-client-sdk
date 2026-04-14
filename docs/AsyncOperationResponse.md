# AsyncOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusUrl** | Pointer to **string** |  | [optional] 
**SamplingTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewAsyncOperationResponse

`func NewAsyncOperationResponse() *AsyncOperationResponse`

NewAsyncOperationResponse instantiates a new AsyncOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncOperationResponseWithDefaults

`func NewAsyncOperationResponseWithDefaults() *AsyncOperationResponse`

NewAsyncOperationResponseWithDefaults instantiates a new AsyncOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusUrl

`func (o *AsyncOperationResponse) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *AsyncOperationResponse) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *AsyncOperationResponse) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *AsyncOperationResponse) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSamplingTime

`func (o *AsyncOperationResponse) GetSamplingTime() int32`

GetSamplingTime returns the SamplingTime field if non-nil, zero value otherwise.

### GetSamplingTimeOk

`func (o *AsyncOperationResponse) GetSamplingTimeOk() (*int32, bool)`

GetSamplingTimeOk returns a tuple with the SamplingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingTime

`func (o *AsyncOperationResponse) SetSamplingTime(v int32)`

SetSamplingTime sets SamplingTime field to given value.

### HasSamplingTime

`func (o *AsyncOperationResponse) HasSamplingTime() bool`

HasSamplingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


