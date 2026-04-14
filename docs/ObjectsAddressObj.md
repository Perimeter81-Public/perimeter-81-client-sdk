# ObjectsAddressObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Address ID | [optional] 
**Name** | **string** | The name of the Objects | 
**Description** | Pointer to **string** |  | [optional] 
**ValueType** | **string** |  | 
**Value** | **[]string** |  | 

## Methods

### NewObjectsAddressObj

`func NewObjectsAddressObj(name string, valueType string, value []string, ) *ObjectsAddressObj`

NewObjectsAddressObj instantiates a new ObjectsAddressObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsAddressObjWithDefaults

`func NewObjectsAddressObjWithDefaults() *ObjectsAddressObj`

NewObjectsAddressObjWithDefaults instantiates a new ObjectsAddressObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectsAddressObj) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectsAddressObj) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectsAddressObj) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectsAddressObj) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ObjectsAddressObj) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectsAddressObj) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectsAddressObj) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ObjectsAddressObj) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectsAddressObj) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectsAddressObj) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectsAddressObj) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueType

`func (o *ObjectsAddressObj) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ObjectsAddressObj) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ObjectsAddressObj) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *ObjectsAddressObj) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectsAddressObj) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectsAddressObj) SetValue(v []string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


