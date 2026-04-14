# StandardHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StandardHealthCheckType**](StandardHealthCheckType.md) |  | 
**Meta** | [**StandardHealthCheckMeta**](StandardHealthCheckMeta.md) |  | 
**Status** | [**HealthStatus**](HealthStatus.md) |  | 

## Methods

### NewStandardHealthCheck

`func NewStandardHealthCheck(type_ StandardHealthCheckType, meta StandardHealthCheckMeta, status HealthStatus, ) *StandardHealthCheck`

NewStandardHealthCheck instantiates a new StandardHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardHealthCheckWithDefaults

`func NewStandardHealthCheckWithDefaults() *StandardHealthCheck`

NewStandardHealthCheckWithDefaults instantiates a new StandardHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StandardHealthCheck) GetType() StandardHealthCheckType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StandardHealthCheck) GetTypeOk() (*StandardHealthCheckType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StandardHealthCheck) SetType(v StandardHealthCheckType)`

SetType sets Type field to given value.


### GetMeta

`func (o *StandardHealthCheck) GetMeta() StandardHealthCheckMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StandardHealthCheck) GetMetaOk() (*StandardHealthCheckMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StandardHealthCheck) SetMeta(v StandardHealthCheckMeta)`

SetMeta sets Meta field to given value.


### GetStatus

`func (o *StandardHealthCheck) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StandardHealthCheck) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StandardHealthCheck) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


