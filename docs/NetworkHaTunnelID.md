# NetworkHaTunnelID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**Network** | **string** | ID of the network. | 
**Tunnels** | **[]string** |  | 
**Type** | **string** |  | 
**TenantId** | **string** | tenantId. | 
**Id** | **string** | id. | 

## Methods

### NewNetworkHaTunnelID

`func NewNetworkHaTunnelID(createdAt time.Time, network string, tunnels []string, type_ string, tenantId string, id string, ) *NetworkHaTunnelID`

NewNetworkHaTunnelID instantiates a new NetworkHaTunnelID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHaTunnelIDWithDefaults

`func NewNetworkHaTunnelIDWithDefaults() *NetworkHaTunnelID`

NewNetworkHaTunnelIDWithDefaults instantiates a new NetworkHaTunnelID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NetworkHaTunnelID) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkHaTunnelID) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkHaTunnelID) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NetworkHaTunnelID) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkHaTunnelID) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkHaTunnelID) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkHaTunnelID) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkHaTunnelID) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkHaTunnelID) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkHaTunnelID) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTunnels

`func (o *NetworkHaTunnelID) GetTunnels() []string`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *NetworkHaTunnelID) GetTunnelsOk() (*[]string, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *NetworkHaTunnelID) SetTunnels(v []string)`

SetTunnels sets Tunnels field to given value.


### GetType

`func (o *NetworkHaTunnelID) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkHaTunnelID) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkHaTunnelID) SetType(v string)`

SetType sets Type field to given value.


### GetTenantId

`func (o *NetworkHaTunnelID) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NetworkHaTunnelID) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NetworkHaTunnelID) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetId

`func (o *NetworkHaTunnelID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkHaTunnelID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkHaTunnelID) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


