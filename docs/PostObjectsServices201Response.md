# PostObjectsServices201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Objects | 
**Description** | Pointer to **string** |  | [optional] 
**Protocols** | [**[]ObjectsServicesProtocolResponseObj**](ObjectsServicesProtocolResponseObj.md) |  | 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewPostObjectsServices201Response

`func NewPostObjectsServices201Response(name string, protocols []ObjectsServicesProtocolResponseObj, ) *PostObjectsServices201Response`

NewPostObjectsServices201Response instantiates a new PostObjectsServices201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObjectsServices201ResponseWithDefaults

`func NewPostObjectsServices201ResponseWithDefaults() *PostObjectsServices201Response`

NewPostObjectsServices201ResponseWithDefaults instantiates a new PostObjectsServices201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostObjectsServices201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostObjectsServices201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostObjectsServices201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PostObjectsServices201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostObjectsServices201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostObjectsServices201Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostObjectsServices201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocols

`func (o *PostObjectsServices201Response) GetProtocols() []ObjectsServicesProtocolResponseObj`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *PostObjectsServices201Response) GetProtocolsOk() (*[]ObjectsServicesProtocolResponseObj, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *PostObjectsServices201Response) SetProtocols(v []ObjectsServicesProtocolResponseObj)`

SetProtocols sets Protocols field to given value.


### GetId

`func (o *PostObjectsServices201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostObjectsServices201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostObjectsServices201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostObjectsServices201Response) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


