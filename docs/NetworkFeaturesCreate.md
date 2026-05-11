# NetworkFeaturesCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudSecurity** | Pointer to [**NetworkFeaturesCloudSecurity**](NetworkFeaturesCloudSecurity.md) |  | [optional] 
**SymmetricInnerMesh** | Pointer to [**NetworkFeaturesCloudSecurity**](NetworkFeaturesCloudSecurity.md) |  | [optional] 
**DNSServices** | Pointer to [**NetworkFeaturesDNSServices**](NetworkFeaturesDNSServices.md) |  | [optional] 

## Methods

### NewNetworkFeaturesCreate

`func NewNetworkFeaturesCreate() *NetworkFeaturesCreate`

NewNetworkFeaturesCreate instantiates a new NetworkFeaturesCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeaturesCreateWithDefaults

`func NewNetworkFeaturesCreateWithDefaults() *NetworkFeaturesCreate`

NewNetworkFeaturesCreateWithDefaults instantiates a new NetworkFeaturesCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudSecurity

`func (o *NetworkFeaturesCreate) GetCloudSecurity() NetworkFeaturesCloudSecurity`

GetCloudSecurity returns the CloudSecurity field if non-nil, zero value otherwise.

### GetCloudSecurityOk

`func (o *NetworkFeaturesCreate) GetCloudSecurityOk() (*NetworkFeaturesCloudSecurity, bool)`

GetCloudSecurityOk returns a tuple with the CloudSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSecurity

`func (o *NetworkFeaturesCreate) SetCloudSecurity(v NetworkFeaturesCloudSecurity)`

SetCloudSecurity sets CloudSecurity field to given value.

### HasCloudSecurity

`func (o *NetworkFeaturesCreate) HasCloudSecurity() bool`

HasCloudSecurity returns a boolean if a field has been set.

### GetSymmetricInnerMesh

`func (o *NetworkFeaturesCreate) GetSymmetricInnerMesh() NetworkFeaturesCloudSecurity`

GetSymmetricInnerMesh returns the SymmetricInnerMesh field if non-nil, zero value otherwise.

### GetSymmetricInnerMeshOk

`func (o *NetworkFeaturesCreate) GetSymmetricInnerMeshOk() (*NetworkFeaturesCloudSecurity, bool)`

GetSymmetricInnerMeshOk returns a tuple with the SymmetricInnerMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymmetricInnerMesh

`func (o *NetworkFeaturesCreate) SetSymmetricInnerMesh(v NetworkFeaturesCloudSecurity)`

SetSymmetricInnerMesh sets SymmetricInnerMesh field to given value.

### HasSymmetricInnerMesh

`func (o *NetworkFeaturesCreate) HasSymmetricInnerMesh() bool`

HasSymmetricInnerMesh returns a boolean if a field has been set.

### GetDNSServices

`func (o *NetworkFeaturesCreate) GetDNSServices() NetworkFeaturesDNSServices`

GetDNSServices returns the DNSServices field if non-nil, zero value otherwise.

### GetDNSServicesOk

`func (o *NetworkFeaturesCreate) GetDNSServicesOk() (*NetworkFeaturesDNSServices, bool)`

GetDNSServicesOk returns a tuple with the DNSServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNSServices

`func (o *NetworkFeaturesCreate) SetDNSServices(v NetworkFeaturesDNSServices)`

SetDNSServices sets DNSServices field to given value.

### HasDNSServices

`func (o *NetworkFeaturesCreate) HasDNSServices() bool`

HasDNSServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


