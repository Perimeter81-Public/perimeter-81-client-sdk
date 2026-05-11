# EnhancedRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID for the enhanced region. | 
**TenantId** | **string** | ID of the tenant that owns the region. | 
**Name** | **string** |  | 
**Dns** | **string** |  | 
**Network** | **string** |  | 
**ScaleUnits** | **int32** | Number of scale units for the region. Scale units determine the capacity and performance level of the enhanced region. Higher values provide greater throughput and connection capacity.  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Attributes** | [**EnhancedRegionAttributes**](EnhancedRegionAttributes.md) |  | 

## Methods

### NewEnhancedRegion

`func NewEnhancedRegion(id string, tenantId string, name string, dns string, network string, scaleUnits int32, createdAt time.Time, attributes EnhancedRegionAttributes, ) *EnhancedRegion`

NewEnhancedRegion instantiates a new EnhancedRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRegionWithDefaults

`func NewEnhancedRegionWithDefaults() *EnhancedRegion`

NewEnhancedRegionWithDefaults instantiates a new EnhancedRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnhancedRegion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedRegion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedRegion) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *EnhancedRegion) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *EnhancedRegion) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *EnhancedRegion) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *EnhancedRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnhancedRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnhancedRegion) SetName(v string)`

SetName sets Name field to given value.


### GetDns

`func (o *EnhancedRegion) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *EnhancedRegion) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *EnhancedRegion) SetDns(v string)`

SetDns sets Dns field to given value.


### GetNetwork

`func (o *EnhancedRegion) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *EnhancedRegion) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *EnhancedRegion) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetScaleUnits

`func (o *EnhancedRegion) GetScaleUnits() int32`

GetScaleUnits returns the ScaleUnits field if non-nil, zero value otherwise.

### GetScaleUnitsOk

`func (o *EnhancedRegion) GetScaleUnitsOk() (*int32, bool)`

GetScaleUnitsOk returns a tuple with the ScaleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUnits

`func (o *EnhancedRegion) SetScaleUnits(v int32)`

SetScaleUnits sets ScaleUnits field to given value.


### GetCreatedAt

`func (o *EnhancedRegion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnhancedRegion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnhancedRegion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnhancedRegion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnhancedRegion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnhancedRegion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnhancedRegion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttributes

`func (o *EnhancedRegion) GetAttributes() EnhancedRegionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EnhancedRegion) GetAttributesOk() (*EnhancedRegionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EnhancedRegion) SetAttributes(v EnhancedRegionAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


