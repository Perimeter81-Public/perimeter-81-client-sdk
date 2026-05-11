# EnhancedNetwork

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
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 

## Methods

### NewEnhancedNetwork

`func NewEnhancedNetwork(createdAt time.Time, dns string, subnet string, accessType string, applications []string, tags []string, name string, isDefault bool, id string, tenantId string, ) *EnhancedNetwork`

NewEnhancedNetwork instantiates a new EnhancedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedNetworkWithDefaults

`func NewEnhancedNetworkWithDefaults() *EnhancedNetwork`

NewEnhancedNetworkWithDefaults instantiates a new EnhancedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EnhancedNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnhancedNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnhancedNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnhancedNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnhancedNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnhancedNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnhancedNetwork) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDns

`func (o *EnhancedNetwork) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *EnhancedNetwork) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *EnhancedNetwork) SetDns(v string)`

SetDns sets Dns field to given value.


### GetSubnet

`func (o *EnhancedNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *EnhancedNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *EnhancedNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetAccessType

`func (o *EnhancedNetwork) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *EnhancedNetwork) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *EnhancedNetwork) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetApplications

`func (o *EnhancedNetwork) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EnhancedNetwork) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EnhancedNetwork) SetApplications(v []string)`

SetApplications sets Applications field to given value.


### GetTags

`func (o *EnhancedNetwork) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EnhancedNetwork) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EnhancedNetwork) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetName

`func (o *EnhancedNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnhancedNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnhancedNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *EnhancedNetwork) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *EnhancedNetwork) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *EnhancedNetwork) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetId

`func (o *EnhancedNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *EnhancedNetwork) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EnhancedNetwork) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EnhancedNetwork) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAsn

`func (o *EnhancedNetwork) GetAsn() ASN`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *EnhancedNetwork) GetAsnOk() (*ASN, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *EnhancedNetwork) SetAsn(v ASN)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *EnhancedNetwork) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAttributes

`func (o *EnhancedNetwork) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EnhancedNetwork) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EnhancedNetwork) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EnhancedNetwork) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


