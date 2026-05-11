# EnhancedHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EnhancedHealthCheck**](EnhancedHealthCheck.md) | List of health checks for enhanced networks (tunnels only) | 

## Methods

### NewEnhancedHealthResponse

`func NewEnhancedHealthResponse(data []EnhancedHealthCheck, ) *EnhancedHealthResponse`

NewEnhancedHealthResponse instantiates a new EnhancedHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedHealthResponseWithDefaults

`func NewEnhancedHealthResponseWithDefaults() *EnhancedHealthResponse`

NewEnhancedHealthResponseWithDefaults instantiates a new EnhancedHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnhancedHealthResponse) GetData() []EnhancedHealthCheck`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnhancedHealthResponse) GetDataOk() (*[]EnhancedHealthCheck, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnhancedHealthResponse) SetData(v []EnhancedHealthCheck)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


