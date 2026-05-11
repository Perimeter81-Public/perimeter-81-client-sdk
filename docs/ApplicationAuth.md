# ApplicationAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEnabled** | **bool** |  | 
**AuthUsername** | Pointer to **string** |  | [optional] 
**AuthPassword** | Pointer to **string** |  | [optional] 
**AuthDomain** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationAuth

`func NewApplicationAuth(authEnabled bool, ) *ApplicationAuth`

NewApplicationAuth instantiates a new ApplicationAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAuthWithDefaults

`func NewApplicationAuthWithDefaults() *ApplicationAuth`

NewApplicationAuthWithDefaults instantiates a new ApplicationAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEnabled

`func (o *ApplicationAuth) GetAuthEnabled() bool`

GetAuthEnabled returns the AuthEnabled field if non-nil, zero value otherwise.

### GetAuthEnabledOk

`func (o *ApplicationAuth) GetAuthEnabledOk() (*bool, bool)`

GetAuthEnabledOk returns a tuple with the AuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEnabled

`func (o *ApplicationAuth) SetAuthEnabled(v bool)`

SetAuthEnabled sets AuthEnabled field to given value.


### GetAuthUsername

`func (o *ApplicationAuth) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *ApplicationAuth) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *ApplicationAuth) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *ApplicationAuth) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetAuthPassword

`func (o *ApplicationAuth) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *ApplicationAuth) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *ApplicationAuth) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *ApplicationAuth) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthDomain

`func (o *ApplicationAuth) GetAuthDomain() string`

GetAuthDomain returns the AuthDomain field if non-nil, zero value otherwise.

### GetAuthDomainOk

`func (o *ApplicationAuth) GetAuthDomainOk() (*string, bool)`

GetAuthDomainOk returns a tuple with the AuthDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomain

`func (o *ApplicationAuth) SetAuthDomain(v string)`

SetAuthDomain sets AuthDomain field to given value.

### HasAuthDomain

`func (o *ApplicationAuth) HasAuthDomain() bool`

HasAuthDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


