# ObjectsServicesProtocolRequestObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**ValueType** | **string** |  | 
**Value** | **[]int32** |  | 
**ProtocolOptions** | [**ObjectServiceProtocolOptionsICMPrequest**](ObjectServiceProtocolOptionsICMPrequest.md) |  | 

## Methods

### NewObjectsServicesProtocolRequestObj

`func NewObjectsServicesProtocolRequestObj(protocol string, valueType string, value []int32, protocolOptions ObjectServiceProtocolOptionsICMPrequest, ) *ObjectsServicesProtocolRequestObj`

NewObjectsServicesProtocolRequestObj instantiates a new ObjectsServicesProtocolRequestObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsServicesProtocolRequestObjWithDefaults

`func NewObjectsServicesProtocolRequestObjWithDefaults() *ObjectsServicesProtocolRequestObj`

NewObjectsServicesProtocolRequestObjWithDefaults instantiates a new ObjectsServicesProtocolRequestObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ObjectsServicesProtocolRequestObj) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ObjectsServicesProtocolRequestObj) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ObjectsServicesProtocolRequestObj) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetValueType

`func (o *ObjectsServicesProtocolRequestObj) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ObjectsServicesProtocolRequestObj) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ObjectsServicesProtocolRequestObj) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *ObjectsServicesProtocolRequestObj) GetValue() []int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectsServicesProtocolRequestObj) GetValueOk() (*[]int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectsServicesProtocolRequestObj) SetValue(v []int32)`

SetValue sets Value field to given value.


### GetProtocolOptions

`func (o *ObjectsServicesProtocolRequestObj) GetProtocolOptions() ObjectServiceProtocolOptionsICMPrequest`

GetProtocolOptions returns the ProtocolOptions field if non-nil, zero value otherwise.

### GetProtocolOptionsOk

`func (o *ObjectsServicesProtocolRequestObj) GetProtocolOptionsOk() (*ObjectServiceProtocolOptionsICMPrequest, bool)`

GetProtocolOptionsOk returns a tuple with the ProtocolOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolOptions

`func (o *ObjectsServicesProtocolRequestObj) SetProtocolOptions(v ObjectServiceProtocolOptionsICMPrequest)`

SetProtocolOptions sets ProtocolOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


