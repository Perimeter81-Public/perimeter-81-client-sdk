# IPSecPhaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **[]string** |  | 
**Encryption** | **[]string** |  | 
**Dh** | **[]int32** | Diffie Helman encryption | 

## Methods

### NewIPSecPhaseConfig

`func NewIPSecPhaseConfig(auth []string, encryption []string, dh []int32, ) *IPSecPhaseConfig`

NewIPSecPhaseConfig instantiates a new IPSecPhaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecPhaseConfigWithDefaults

`func NewIPSecPhaseConfigWithDefaults() *IPSecPhaseConfig`

NewIPSecPhaseConfigWithDefaults instantiates a new IPSecPhaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *IPSecPhaseConfig) GetAuth() []string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IPSecPhaseConfig) GetAuthOk() (*[]string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IPSecPhaseConfig) SetAuth(v []string)`

SetAuth sets Auth field to given value.


### GetEncryption

`func (o *IPSecPhaseConfig) GetEncryption() []string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *IPSecPhaseConfig) GetEncryptionOk() (*[]string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *IPSecPhaseConfig) SetEncryption(v []string)`

SetEncryption sets Encryption field to given value.


### GetDh

`func (o *IPSecPhaseConfig) GetDh() []int32`

GetDh returns the Dh field if non-nil, zero value otherwise.

### GetDhOk

`func (o *IPSecPhaseConfig) GetDhOk() (*[]int32, bool)`

GetDhOk returns a tuple with the Dh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDh

`func (o *IPSecPhaseConfig) SetDh(v []int32)`

SetDh sets Dh field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


