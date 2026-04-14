# IPSecPhaseConfigV23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **[]string** |  | 
**Encryption** | **[]string** |  | 
**KeyExchangeMethod** | **[]string** | Key exchange method encryption | 

## Methods

### NewIPSecPhaseConfigV23

`func NewIPSecPhaseConfigV23(auth []string, encryption []string, keyExchangeMethod []string, ) *IPSecPhaseConfigV23`

NewIPSecPhaseConfigV23 instantiates a new IPSecPhaseConfigV23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecPhaseConfigV23WithDefaults

`func NewIPSecPhaseConfigV23WithDefaults() *IPSecPhaseConfigV23`

NewIPSecPhaseConfigV23WithDefaults instantiates a new IPSecPhaseConfigV23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *IPSecPhaseConfigV23) GetAuth() []string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IPSecPhaseConfigV23) GetAuthOk() (*[]string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IPSecPhaseConfigV23) SetAuth(v []string)`

SetAuth sets Auth field to given value.


### GetEncryption

`func (o *IPSecPhaseConfigV23) GetEncryption() []string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *IPSecPhaseConfigV23) GetEncryptionOk() (*[]string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *IPSecPhaseConfigV23) SetEncryption(v []string)`

SetEncryption sets Encryption field to given value.


### GetKeyExchangeMethod

`func (o *IPSecPhaseConfigV23) GetKeyExchangeMethod() []string`

GetKeyExchangeMethod returns the KeyExchangeMethod field if non-nil, zero value otherwise.

### GetKeyExchangeMethodOk

`func (o *IPSecPhaseConfigV23) GetKeyExchangeMethodOk() (*[]string, bool)`

GetKeyExchangeMethodOk returns a tuple with the KeyExchangeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchangeMethod

`func (o *IPSecPhaseConfigV23) SetKeyExchangeMethod(v []string)`

SetKeyExchangeMethod sets KeyExchangeMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


