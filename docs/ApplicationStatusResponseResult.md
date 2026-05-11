# ApplicationStatusResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | Response status code | [optional] 
**Resource** | Pointer to **string** | Endpoint to get the resource if exists | [optional] 
**Reason** | Pointer to **[]string** | Failure reasons | [optional] 

## Methods

### NewApplicationStatusResponseResult

`func NewApplicationStatusResponseResult() *ApplicationStatusResponseResult`

NewApplicationStatusResponseResult instantiates a new ApplicationStatusResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStatusResponseResultWithDefaults

`func NewApplicationStatusResponseResultWithDefaults() *ApplicationStatusResponseResult`

NewApplicationStatusResponseResultWithDefaults instantiates a new ApplicationStatusResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ApplicationStatusResponseResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApplicationStatusResponseResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApplicationStatusResponseResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ApplicationStatusResponseResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetResource

`func (o *ApplicationStatusResponseResult) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApplicationStatusResponseResult) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApplicationStatusResponseResult) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ApplicationStatusResponseResult) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetReason

`func (o *ApplicationStatusResponseResult) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApplicationStatusResponseResult) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApplicationStatusResponseResult) SetReason(v []string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApplicationStatusResponseResult) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


