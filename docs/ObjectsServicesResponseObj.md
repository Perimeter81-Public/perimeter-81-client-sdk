# ObjectsServicesResponseObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Objects | 
**Description** | Pointer to **string** |  | [optional] 
**Protocols** | [**[]ObjectsServicesProtocolResponseObj**](ObjectsServicesProtocolResponseObj.md) |  | 

## Methods

### NewObjectsServicesResponseObj

`func NewObjectsServicesResponseObj(name string, protocols []ObjectsServicesProtocolResponseObj, ) *ObjectsServicesResponseObj`

NewObjectsServicesResponseObj instantiates a new ObjectsServicesResponseObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsServicesResponseObjWithDefaults

`func NewObjectsServicesResponseObjWithDefaults() *ObjectsServicesResponseObj`

NewObjectsServicesResponseObjWithDefaults instantiates a new ObjectsServicesResponseObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectsServicesResponseObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectsServicesResponseObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectsServicesResponseObj) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ObjectsServicesResponseObj) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectsServicesResponseObj) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectsServicesResponseObj) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectsServicesResponseObj) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocols

`func (o *ObjectsServicesResponseObj) GetProtocols() []ObjectsServicesProtocolResponseObj`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ObjectsServicesResponseObj) GetProtocolsOk() (*[]ObjectsServicesProtocolResponseObj, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ObjectsServicesResponseObj) SetProtocols(v []ObjectsServicesProtocolResponseObj)`

SetProtocols sets Protocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


