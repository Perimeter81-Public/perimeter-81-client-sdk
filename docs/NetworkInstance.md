# NetworkInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**Network** | **string** | ID of the network. | 
**Region** | **string** | ID of the network region. | 
**InstanceType** | **string** |  | 
**ImageType** | **string** |  | 
**ImageVersion** | **string** |  | 
**Dns** | **string** |  | 
**Ip** | **string** |  | 
**Tunnels** | [**[]NetworkTunnel**](NetworkTunnel.md) | List of network tunnels. | 
**Id** | **string** | Unique ID. | 
**TenantId** | **string** | ID of the tenant. | 

## Methods

### NewNetworkInstance

`func NewNetworkInstance(createdAt time.Time, network string, region string, instanceType string, imageType string, imageVersion string, dns string, ip string, tunnels []NetworkTunnel, id string, tenantId string, ) *NetworkInstance`

NewNetworkInstance instantiates a new NetworkInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInstanceWithDefaults

`func NewNetworkInstanceWithDefaults() *NetworkInstance`

NewNetworkInstanceWithDefaults instantiates a new NetworkInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NetworkInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkInstance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkInstance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkInstance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkInstance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkInstance) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkInstance) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkInstance) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRegion

`func (o *NetworkInstance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkInstance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkInstance) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetInstanceType

`func (o *NetworkInstance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *NetworkInstance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *NetworkInstance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetImageType

`func (o *NetworkInstance) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *NetworkInstance) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *NetworkInstance) SetImageType(v string)`

SetImageType sets ImageType field to given value.


### GetImageVersion

`func (o *NetworkInstance) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *NetworkInstance) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *NetworkInstance) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.


### GetDns

`func (o *NetworkInstance) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NetworkInstance) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NetworkInstance) SetDns(v string)`

SetDns sets Dns field to given value.


### GetIp

`func (o *NetworkInstance) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkInstance) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkInstance) SetIp(v string)`

SetIp sets Ip field to given value.


### GetTunnels

`func (o *NetworkInstance) GetTunnels() []NetworkTunnel`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *NetworkInstance) GetTunnelsOk() (*[]NetworkTunnel, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *NetworkInstance) SetTunnels(v []NetworkTunnel)`

SetTunnels sets Tunnels field to given value.


### GetId

`func (o *NetworkInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInstance) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *NetworkInstance) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetworkInstance) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetworkInstance) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


