# IPSecAlgorithms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **[]string** |  | [optional] 
**Encryption** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIPSecAlgorithms

`func NewIPSecAlgorithms() *IPSecAlgorithms`

NewIPSecAlgorithms instantiates a new IPSecAlgorithms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecAlgorithmsWithDefaults

`func NewIPSecAlgorithmsWithDefaults() *IPSecAlgorithms`

NewIPSecAlgorithmsWithDefaults instantiates a new IPSecAlgorithms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *IPSecAlgorithms) GetAuth() []string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IPSecAlgorithms) GetAuthOk() (*[]string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IPSecAlgorithms) SetAuth(v []string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *IPSecAlgorithms) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetEncryption

`func (o *IPSecAlgorithms) GetEncryption() []string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *IPSecAlgorithms) GetEncryptionOk() (*[]string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *IPSecAlgorithms) SetEncryption(v []string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *IPSecAlgorithms) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


