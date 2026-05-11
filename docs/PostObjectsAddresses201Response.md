# PostObjectsAddresses201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the Objects | 
**Description** | Pointer to **string** |  | [optional] 
**ValueType** | **string** |  | 
**Value** | **[]string** |  | 

## Methods

### NewPostObjectsAddresses201Response

`func NewPostObjectsAddresses201Response(name string, valueType string, value []string, ) *PostObjectsAddresses201Response`

NewPostObjectsAddresses201Response instantiates a new PostObjectsAddresses201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObjectsAddresses201ResponseWithDefaults

`func NewPostObjectsAddresses201ResponseWithDefaults() *PostObjectsAddresses201Response`

NewPostObjectsAddresses201ResponseWithDefaults instantiates a new PostObjectsAddresses201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostObjectsAddresses201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostObjectsAddresses201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostObjectsAddresses201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostObjectsAddresses201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PostObjectsAddresses201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostObjectsAddresses201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostObjectsAddresses201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostObjectsAddresses201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostObjectsAddresses201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostObjectsAddresses201Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostObjectsAddresses201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueType

`func (o *PostObjectsAddresses201Response) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *PostObjectsAddresses201Response) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *PostObjectsAddresses201Response) SetValueType(v string)`

SetValueType sets ValueType field to given value.


### GetValue

`func (o *PostObjectsAddresses201Response) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PostObjectsAddresses201Response) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PostObjectsAddresses201Response) SetValue(v []string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


