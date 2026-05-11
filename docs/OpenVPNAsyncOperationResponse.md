# OpenVPNAsyncOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusUrl** | Pointer to **string** |  | [optional] 
**SamplingTime** | Pointer to **int32** |  | [optional] 
**SecretAccessKey** | Pointer to **string** | Secret access key to download OpenVPN configuration file. This key will be shown only once in the response for security reasons. Make sure to store it securely. | [optional] 

## Methods

### NewOpenVPNAsyncOperationResponse

`func NewOpenVPNAsyncOperationResponse() *OpenVPNAsyncOperationResponse`

NewOpenVPNAsyncOperationResponse instantiates a new OpenVPNAsyncOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenVPNAsyncOperationResponseWithDefaults

`func NewOpenVPNAsyncOperationResponseWithDefaults() *OpenVPNAsyncOperationResponse`

NewOpenVPNAsyncOperationResponseWithDefaults instantiates a new OpenVPNAsyncOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusUrl

`func (o *OpenVPNAsyncOperationResponse) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *OpenVPNAsyncOperationResponse) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *OpenVPNAsyncOperationResponse) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.

### HasStatusUrl

`func (o *OpenVPNAsyncOperationResponse) HasStatusUrl() bool`

HasStatusUrl returns a boolean if a field has been set.

### GetSamplingTime

`func (o *OpenVPNAsyncOperationResponse) GetSamplingTime() int32`

GetSamplingTime returns the SamplingTime field if non-nil, zero value otherwise.

### GetSamplingTimeOk

`func (o *OpenVPNAsyncOperationResponse) GetSamplingTimeOk() (*int32, bool)`

GetSamplingTimeOk returns a tuple with the SamplingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingTime

`func (o *OpenVPNAsyncOperationResponse) SetSamplingTime(v int32)`

SetSamplingTime sets SamplingTime field to given value.

### HasSamplingTime

`func (o *OpenVPNAsyncOperationResponse) HasSamplingTime() bool`

HasSamplingTime returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *OpenVPNAsyncOperationResponse) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *OpenVPNAsyncOperationResponse) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *OpenVPNAsyncOperationResponse) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *OpenVPNAsyncOperationResponse) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


