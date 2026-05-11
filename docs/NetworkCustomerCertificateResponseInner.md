# NetworkCustomerCertificateResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Certificate ID | 
**DisplayName** | **string** | Certificate display name | 
**ExpiresAt** | **time.Time** | Certificate expiration date | 

## Methods

### NewNetworkCustomerCertificateResponseInner

`func NewNetworkCustomerCertificateResponseInner(id string, displayName string, expiresAt time.Time, ) *NetworkCustomerCertificateResponseInner`

NewNetworkCustomerCertificateResponseInner instantiates a new NetworkCustomerCertificateResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkCustomerCertificateResponseInnerWithDefaults

`func NewNetworkCustomerCertificateResponseInnerWithDefaults() *NetworkCustomerCertificateResponseInner`

NewNetworkCustomerCertificateResponseInnerWithDefaults instantiates a new NetworkCustomerCertificateResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkCustomerCertificateResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkCustomerCertificateResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkCustomerCertificateResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *NetworkCustomerCertificateResponseInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkCustomerCertificateResponseInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkCustomerCertificateResponseInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpiresAt

`func (o *NetworkCustomerCertificateResponseInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *NetworkCustomerCertificateResponseInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *NetworkCustomerCertificateResponseInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


