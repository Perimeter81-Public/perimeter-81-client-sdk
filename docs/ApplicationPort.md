# ApplicationPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**ApplicationPortValue**](ApplicationPortValue.md) |  | 
**Source** | **string** |  | 

## Methods

### NewApplicationPort

`func NewApplicationPort(value ApplicationPortValue, source string, ) *ApplicationPort`

NewApplicationPort instantiates a new ApplicationPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPortWithDefaults

`func NewApplicationPortWithDefaults() *ApplicationPort`

NewApplicationPortWithDefaults instantiates a new ApplicationPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ApplicationPort) GetValue() ApplicationPortValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationPort) GetValueOk() (*ApplicationPortValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationPort) SetValue(v ApplicationPortValue)`

SetValue sets Value field to given value.


### GetSource

`func (o *ApplicationPort) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApplicationPort) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApplicationPort) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


