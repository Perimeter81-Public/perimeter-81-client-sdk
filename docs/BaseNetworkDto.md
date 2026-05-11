# BaseNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Network name | [optional] 
**Tags** | Pointer to **[]string** | List of network tags | [optional] 

## Methods

### NewBaseNetworkDto

`func NewBaseNetworkDto() *BaseNetworkDto`

NewBaseNetworkDto instantiates a new BaseNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNetworkDtoWithDefaults

`func NewBaseNetworkDtoWithDefaults() *BaseNetworkDto`

NewBaseNetworkDtoWithDefaults instantiates a new BaseNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseNetworkDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseNetworkDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseNetworkDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseNetworkDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *BaseNetworkDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BaseNetworkDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BaseNetworkDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BaseNetworkDto) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


