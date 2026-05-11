# ApplicationAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasEnabled** | **bool** |  | 
**Cname** | Pointer to **string** |  | [optional] 
**SslCertificate** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationAlias

`func NewApplicationAlias(aliasEnabled bool, ) *ApplicationAlias`

NewApplicationAlias instantiates a new ApplicationAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAliasWithDefaults

`func NewApplicationAliasWithDefaults() *ApplicationAlias`

NewApplicationAliasWithDefaults instantiates a new ApplicationAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasEnabled

`func (o *ApplicationAlias) GetAliasEnabled() bool`

GetAliasEnabled returns the AliasEnabled field if non-nil, zero value otherwise.

### GetAliasEnabledOk

`func (o *ApplicationAlias) GetAliasEnabledOk() (*bool, bool)`

GetAliasEnabledOk returns a tuple with the AliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasEnabled

`func (o *ApplicationAlias) SetAliasEnabled(v bool)`

SetAliasEnabled sets AliasEnabled field to given value.


### GetCname

`func (o *ApplicationAlias) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *ApplicationAlias) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *ApplicationAlias) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *ApplicationAlias) HasCname() bool`

HasCname returns a boolean if a field has been set.

### GetSslCertificate

`func (o *ApplicationAlias) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *ApplicationAlias) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *ApplicationAlias) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *ApplicationAlias) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


