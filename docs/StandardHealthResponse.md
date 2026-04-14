# StandardHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]StandardHealthCheck**](StandardHealthCheck.md) | List of health checks for standard networks (includes gateways and tunnels) | 

## Methods

### NewStandardHealthResponse

`func NewStandardHealthResponse(data []StandardHealthCheck, ) *StandardHealthResponse`

NewStandardHealthResponse instantiates a new StandardHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardHealthResponseWithDefaults

`func NewStandardHealthResponseWithDefaults() *StandardHealthResponse`

NewStandardHealthResponseWithDefaults instantiates a new StandardHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *StandardHealthResponse) GetData() []StandardHealthCheck`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StandardHealthResponse) GetDataOk() (*[]StandardHealthCheck, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StandardHealthResponse) SetData(v []StandardHealthCheck)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


