# UpdatedNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**Dns** | **string** | DNS of the network. | 
**Subnet** | **string** | Subnet of the network. | 
**AccessType** | **string** |  | 
**Applications** | **[]string** | List of IDs of assigned applications. | 
**Tags** | **[]string** | List of network tags. | 
**Name** | **string** | Network name. | 
**IsDefault** | **bool** | Indicates that the network is default. | 
**Id** | **string** | Unique ID. | 
**TenantId** | **string** | ID of the tenant. | 
**Asn** | Pointer to [**ASN**](ASN.md) |  | [optional] 
**Regions** | **[]string** | Network regions id list. | 

## Methods

### NewUpdatedNetwork

`func NewUpdatedNetwork(createdAt time.Time, dns string, subnet string, accessType string, applications []string, tags []string, name string, isDefault bool, id string, tenantId string, regions []string, ) *UpdatedNetwork`

NewUpdatedNetwork instantiates a new UpdatedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedNetworkWithDefaults

`func NewUpdatedNetworkWithDefaults() *UpdatedNetwork`

NewUpdatedNetworkWithDefaults instantiates a new UpdatedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UpdatedNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdatedNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdatedNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdatedNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatedNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatedNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdatedNetwork) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDns

`func (o *UpdatedNetwork) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *UpdatedNetwork) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *UpdatedNetwork) SetDns(v string)`

SetDns sets Dns field to given value.


### GetSubnet

`func (o *UpdatedNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdatedNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdatedNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetAccessType

`func (o *UpdatedNetwork) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UpdatedNetwork) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UpdatedNetwork) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetApplications

`func (o *UpdatedNetwork) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *UpdatedNetwork) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *UpdatedNetwork) SetApplications(v []string)`

SetApplications sets Applications field to given value.


### GetTags

`func (o *UpdatedNetwork) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatedNetwork) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatedNetwork) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetName

`func (o *UpdatedNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatedNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatedNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *UpdatedNetwork) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdatedNetwork) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdatedNetwork) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetId

`func (o *UpdatedNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *UpdatedNetwork) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdatedNetwork) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdatedNetwork) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsn

`func (o *UpdatedNetwork) GetAsn() ASN`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *UpdatedNetwork) GetAsnOk() (*ASN, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *UpdatedNetwork) SetAsn(v ASN)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *UpdatedNetwork) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetRegions

`func (o *UpdatedNetwork) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *UpdatedNetwork) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *UpdatedNetwork) SetRegions(v []string)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


