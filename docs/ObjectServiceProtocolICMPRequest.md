# ObjectServiceProtocolICMPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**ProtocolOptions** | [**ObjectServiceProtocolOptionsICMPrequest**](ObjectServiceProtocolOptionsICMPrequest.md) |  | 

## Methods

### NewObjectServiceProtocolICMPRequest

`func NewObjectServiceProtocolICMPRequest(protocol string, protocolOptions ObjectServiceProtocolOptionsICMPrequest, ) *ObjectServiceProtocolICMPRequest`

NewObjectServiceProtocolICMPRequest instantiates a new ObjectServiceProtocolICMPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectServiceProtocolICMPRequestWithDefaults

`func NewObjectServiceProtocolICMPRequestWithDefaults() *ObjectServiceProtocolICMPRequest`

NewObjectServiceProtocolICMPRequestWithDefaults instantiates a new ObjectServiceProtocolICMPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ObjectServiceProtocolICMPRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ObjectServiceProtocolICMPRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ObjectServiceProtocolICMPRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetProtocolOptions

`func (o *ObjectServiceProtocolICMPRequest) GetProtocolOptions() ObjectServiceProtocolOptionsICMPrequest`

GetProtocolOptions returns the ProtocolOptions field if non-nil, zero value otherwise.

### GetProtocolOptionsOk

`func (o *ObjectServiceProtocolICMPRequest) GetProtocolOptionsOk() (*ObjectServiceProtocolOptionsICMPrequest, bool)`

GetProtocolOptionsOk returns a tuple with the ProtocolOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolOptions

`func (o *ObjectServiceProtocolICMPRequest) SetProtocolOptions(v ObjectServiceProtocolOptionsICMPrequest)`

SetProtocolOptions sets ProtocolOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


