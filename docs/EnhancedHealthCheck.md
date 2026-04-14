# EnhancedHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnhancedHealthCheckType**](EnhancedHealthCheckType.md) |  | 
**Meta** | [**EnhancedHealthCheckMeta**](EnhancedHealthCheckMeta.md) |  | 
**Status** | [**HealthStatus**](HealthStatus.md) |  | 

## Methods

### NewEnhancedHealthCheck

`func NewEnhancedHealthCheck(type_ EnhancedHealthCheckType, meta EnhancedHealthCheckMeta, status HealthStatus, ) *EnhancedHealthCheck`

NewEnhancedHealthCheck instantiates a new EnhancedHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedHealthCheckWithDefaults

`func NewEnhancedHealthCheckWithDefaults() *EnhancedHealthCheck`

NewEnhancedHealthCheckWithDefaults instantiates a new EnhancedHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnhancedHealthCheck) GetType() EnhancedHealthCheckType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnhancedHealthCheck) GetTypeOk() (*EnhancedHealthCheckType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnhancedHealthCheck) SetType(v EnhancedHealthCheckType)`

SetType sets Type field to given value.


### GetMeta

`func (o *EnhancedHealthCheck) GetMeta() EnhancedHealthCheckMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnhancedHealthCheck) GetMetaOk() (*EnhancedHealthCheckMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnhancedHealthCheck) SetMeta(v EnhancedHealthCheckMeta)`

SetMeta sets Meta field to given value.


### GetStatus

`func (o *EnhancedHealthCheck) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnhancedHealthCheck) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnhancedHealthCheck) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


