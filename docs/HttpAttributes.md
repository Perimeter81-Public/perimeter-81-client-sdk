# HttpAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStatus** | Pointer to **int32** |  | [optional] [default to 200]
**CheckPath** | Pointer to **string** |  | [optional] [default to ""]
**StartRelativePath** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewHttpAttributes

`func NewHttpAttributes() *HttpAttributes`

NewHttpAttributes instantiates a new HttpAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpAttributesWithDefaults

`func NewHttpAttributesWithDefaults() *HttpAttributes`

NewHttpAttributesWithDefaults instantiates a new HttpAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStatus

`func (o *HttpAttributes) GetCheckStatus() int32`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *HttpAttributes) GetCheckStatusOk() (*int32, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *HttpAttributes) SetCheckStatus(v int32)`

SetCheckStatus sets CheckStatus field to given value.

### HasCheckStatus

`func (o *HttpAttributes) HasCheckStatus() bool`

HasCheckStatus returns a boolean if a field has been set.

### GetCheckPath

`func (o *HttpAttributes) GetCheckPath() string`

GetCheckPath returns the CheckPath field if non-nil, zero value otherwise.

### GetCheckPathOk

`func (o *HttpAttributes) GetCheckPathOk() (*string, bool)`

GetCheckPathOk returns a tuple with the CheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPath

`func (o *HttpAttributes) SetCheckPath(v string)`

SetCheckPath sets CheckPath field to given value.

### HasCheckPath

`func (o *HttpAttributes) HasCheckPath() bool`

HasCheckPath returns a boolean if a field has been set.

### GetStartRelativePath

`func (o *HttpAttributes) GetStartRelativePath() string`

GetStartRelativePath returns the StartRelativePath field if non-nil, zero value otherwise.

### GetStartRelativePathOk

`func (o *HttpAttributes) GetStartRelativePathOk() (*string, bool)`

GetStartRelativePathOk returns a tuple with the StartRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRelativePath

`func (o *HttpAttributes) SetStartRelativePath(v string)`

SetStartRelativePath sets StartRelativePath field to given value.

### HasStartRelativePath

`func (o *HttpAttributes) HasStartRelativePath() bool`

HasStartRelativePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


