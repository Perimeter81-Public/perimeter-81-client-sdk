# TunnelAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** | Authentication type for tunnel (psk for pre-shared key, cert for certificate) | [optional] 
**Passphrase** | Pointer to **string** | Pre-shared key for tunnel authentication (8-64 characters). Required when authType is psk. | [optional] 
**CustomerRootCA** | Pointer to **string** | Customer root certificate authority. Required when authType is cert. | [optional] 

## Methods

### NewTunnelAuthConfig

`func NewTunnelAuthConfig() *TunnelAuthConfig`

NewTunnelAuthConfig instantiates a new TunnelAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelAuthConfigWithDefaults

`func NewTunnelAuthConfigWithDefaults() *TunnelAuthConfig`

NewTunnelAuthConfigWithDefaults instantiates a new TunnelAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *TunnelAuthConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *TunnelAuthConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *TunnelAuthConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *TunnelAuthConfig) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetPassphrase

`func (o *TunnelAuthConfig) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *TunnelAuthConfig) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *TunnelAuthConfig) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *TunnelAuthConfig) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCustomerRootCA

`func (o *TunnelAuthConfig) GetCustomerRootCA() string`

GetCustomerRootCA returns the CustomerRootCA field if non-nil, zero value otherwise.

### GetCustomerRootCAOk

`func (o *TunnelAuthConfig) GetCustomerRootCAOk() (*string, bool)`

GetCustomerRootCAOk returns a tuple with the CustomerRootCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRootCA

`func (o *TunnelAuthConfig) SetCustomerRootCA(v string)`

SetCustomerRootCA sets CustomerRootCA field to given value.

### HasCustomerRootCA

`func (o *TunnelAuthConfig) HasCustomerRootCA() bool`

HasCustomerRootCA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


