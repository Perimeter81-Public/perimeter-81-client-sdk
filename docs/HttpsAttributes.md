# HttpsAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStatus** | Pointer to **int32** |  | [optional] [default to 200]
**CheckPath** | Pointer to **string** |  | [optional] [default to ""]
**StartRelativePath** | Pointer to **string** |  | [optional] [default to ""]
**SslCertificateVerification** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewHttpsAttributes

`func NewHttpsAttributes() *HttpsAttributes`

NewHttpsAttributes instantiates a new HttpsAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpsAttributesWithDefaults

`func NewHttpsAttributesWithDefaults() *HttpsAttributes`

NewHttpsAttributesWithDefaults instantiates a new HttpsAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStatus

`func (o *HttpsAttributes) GetCheckStatus() int32`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *HttpsAttributes) GetCheckStatusOk() (*int32, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *HttpsAttributes) SetCheckStatus(v int32)`

SetCheckStatus sets CheckStatus field to given value.

### HasCheckStatus

`func (o *HttpsAttributes) HasCheckStatus() bool`

HasCheckStatus returns a boolean if a field has been set.

### GetCheckPath

`func (o *HttpsAttributes) GetCheckPath() string`

GetCheckPath returns the CheckPath field if non-nil, zero value otherwise.

### GetCheckPathOk

`func (o *HttpsAttributes) GetCheckPathOk() (*string, bool)`

GetCheckPathOk returns a tuple with the CheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPath

`func (o *HttpsAttributes) SetCheckPath(v string)`

SetCheckPath sets CheckPath field to given value.

### HasCheckPath

`func (o *HttpsAttributes) HasCheckPath() bool`

HasCheckPath returns a boolean if a field has been set.

### GetStartRelativePath

`func (o *HttpsAttributes) GetStartRelativePath() string`

GetStartRelativePath returns the StartRelativePath field if non-nil, zero value otherwise.

### GetStartRelativePathOk

`func (o *HttpsAttributes) GetStartRelativePathOk() (*string, bool)`

GetStartRelativePathOk returns a tuple with the StartRelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRelativePath

`func (o *HttpsAttributes) SetStartRelativePath(v string)`

SetStartRelativePath sets StartRelativePath field to given value.

### HasStartRelativePath

`func (o *HttpsAttributes) HasStartRelativePath() bool`

HasStartRelativePath returns a boolean if a field has been set.

### GetSslCertificateVerification

`func (o *HttpsAttributes) GetSslCertificateVerification() bool`

GetSslCertificateVerification returns the SslCertificateVerification field if non-nil, zero value otherwise.

### GetSslCertificateVerificationOk

`func (o *HttpsAttributes) GetSslCertificateVerificationOk() (*bool, bool)`

GetSslCertificateVerificationOk returns a tuple with the SslCertificateVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificateVerification

`func (o *HttpsAttributes) SetSslCertificateVerification(v bool)`

SetSslCertificateVerification sets SslCertificateVerification field to given value.

### HasSslCertificateVerification

`func (o *HttpsAttributes) HasSslCertificateVerification() bool`

HasSslCertificateVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


