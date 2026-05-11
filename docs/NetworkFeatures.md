# NetworkFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudSecurity** | Pointer to [**NetworkFeaturesCloudSecurity**](NetworkFeaturesCloudSecurity.md) |  | [optional] 
**SymmetricInnerMesh** | Pointer to [**NetworkFeaturesCloudSecurity**](NetworkFeaturesCloudSecurity.md) |  | [optional] 
**DNSServices** | Pointer to [**NetworkFeaturesDNSServices**](NetworkFeaturesDNSServices.md) |  | [optional] 

## Methods

### NewNetworkFeatures

`func NewNetworkFeatures() *NetworkFeatures`

NewNetworkFeatures instantiates a new NetworkFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeaturesWithDefaults

`func NewNetworkFeaturesWithDefaults() *NetworkFeatures`

NewNetworkFeaturesWithDefaults instantiates a new NetworkFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudSecurity

`func (o *NetworkFeatures) GetCloudSecurity() NetworkFeaturesCloudSecurity`

GetCloudSecurity returns the CloudSecurity field if non-nil, zero value otherwise.

### GetCloudSecurityOk

`func (o *NetworkFeatures) GetCloudSecurityOk() (*NetworkFeaturesCloudSecurity, bool)`

GetCloudSecurityOk returns a tuple with the CloudSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSecurity

`func (o *NetworkFeatures) SetCloudSecurity(v NetworkFeaturesCloudSecurity)`

SetCloudSecurity sets CloudSecurity field to given value.

### HasCloudSecurity

`func (o *NetworkFeatures) HasCloudSecurity() bool`

HasCloudSecurity returns a boolean if a field has been set.

### GetSymmetricInnerMesh

`func (o *NetworkFeatures) GetSymmetricInnerMesh() NetworkFeaturesCloudSecurity`

GetSymmetricInnerMesh returns the SymmetricInnerMesh field if non-nil, zero value otherwise.

### GetSymmetricInnerMeshOk

`func (o *NetworkFeatures) GetSymmetricInnerMeshOk() (*NetworkFeaturesCloudSecurity, bool)`

GetSymmetricInnerMeshOk returns a tuple with the SymmetricInnerMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymmetricInnerMesh

`func (o *NetworkFeatures) SetSymmetricInnerMesh(v NetworkFeaturesCloudSecurity)`

SetSymmetricInnerMesh sets SymmetricInnerMesh field to given value.

### HasSymmetricInnerMesh

`func (o *NetworkFeatures) HasSymmetricInnerMesh() bool`

HasSymmetricInnerMesh returns a boolean if a field has been set.

### GetDNSServices

`func (o *NetworkFeatures) GetDNSServices() NetworkFeaturesDNSServices`

GetDNSServices returns the DNSServices field if non-nil, zero value otherwise.

### GetDNSServicesOk

`func (o *NetworkFeatures) GetDNSServicesOk() (*NetworkFeaturesDNSServices, bool)`

GetDNSServicesOk returns a tuple with the DNSServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSServices

`func (o *NetworkFeatures) SetDNSServices(v NetworkFeaturesDNSServices)`

SetDNSServices sets DNSServices field to given value.

### HasDNSServices

`func (o *NetworkFeatures) HasDNSServices() bool`

HasDNSServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


