# ObjectsServicesRequestObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Objects | 
**Description** | Pointer to **string** |  | [optional] 
**Protocols** | [**[]ObjectsServicesProtocolRequestObj**](ObjectsServicesProtocolRequestObj.md) |  | 

## Methods

### NewObjectsServicesRequestObj

`func NewObjectsServicesRequestObj(name string, protocols []ObjectsServicesProtocolRequestObj, ) *ObjectsServicesRequestObj`

NewObjectsServicesRequestObj instantiates a new ObjectsServicesRequestObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsServicesRequestObjWithDefaults

`func NewObjectsServicesRequestObjWithDefaults() *ObjectsServicesRequestObj`

NewObjectsServicesRequestObjWithDefaults instantiates a new ObjectsServicesRequestObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectsServicesRequestObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectsServicesRequestObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectsServicesRequestObj) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ObjectsServicesRequestObj) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectsServicesRequestObj) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectsServicesRequestObj) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectsServicesRequestObj) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocols

`func (o *ObjectsServicesRequestObj) GetProtocols() []ObjectsServicesProtocolRequestObj`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *ObjectsServicesRequestObj) GetProtocolsOk() (*[]ObjectsServicesProtocolRequestObj, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *ObjectsServicesRequestObj) SetProtocols(v []ObjectsServicesProtocolRequestObj)`

SetProtocols sets Protocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


