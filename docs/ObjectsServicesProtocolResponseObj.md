# ObjectsServicesProtocolResponseObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**ValueType** | **string** |  | 
**Value** | **[]int32** |  | 
**ProtocolOptions** | [**ObjectServiceProtocolOptionsICMPresponse**](ObjectServiceProtocolOptionsICMPresponse.md) |  | 

## Methods

### NewObjectsServicesProtocolResponseObj

`func NewObjectsServicesProtocolResponseObj(protocol string, valueType string, value []int32, protocolOptions ObjectServiceProtocolOptionsICMPresponse, ) *ObjectsServicesProtocolResponseObj`

NewObjectsServicesProtocolResponseObj instantiates a new ObjectsServicesProtocolResponseObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsServicesProtocolResponseObjWithDefaults

`func NewObjectsServicesProtocolResponseObjWithDefaults() *ObjectsServicesProtocolResponseObj`

NewObjectsServicesProtocolResponseObjWithDefaults instantiates a new ObjectsServicesProtocolResponseObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ObjectsServicesProtocolResponseObj) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ObjectsServicesProtocolResponseObj) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ObjectsServicesProtocolResponseObj) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetValueType

`func (o *ObjectsServicesProtocolResponseObj) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ObjectsServicesProtocolResponseObj) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ObjectsServicesProtocolResponseObj) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *ObjectsServicesProtocolResponseObj) GetValue() []int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectsServicesProtocolResponseObj) GetValueOk() (*[]int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectsServicesProtocolResponseObj) SetValue(v []int32)`

SetValue sets Value field to given value.


### GetProtocolOptions

`func (o *ObjectsServicesProtocolResponseObj) GetProtocolOptions() ObjectServiceProtocolOptionsICMPresponse`

GetProtocolOptions returns the ProtocolOptions field if non-nil, zero value otherwise.

### GetProtocolOptionsOk

`func (o *ObjectsServicesProtocolResponseObj) GetProtocolOptionsOk() (*ObjectServiceProtocolOptionsICMPresponse, bool)`

GetProtocolOptionsOk returns a tuple with the ProtocolOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolOptions

`func (o *ObjectsServicesProtocolResponseObj) SetProtocolOptions(v ObjectServiceProtocolOptionsICMPresponse)`

SetProtocolOptions sets ProtocolOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


