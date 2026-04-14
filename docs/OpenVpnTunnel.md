# OpenVPNTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionID** | **string** | Region ID | 
**GatewayID** | **string** | Gateway ID | 
**TunnelName** | **string** | The name of the tunnel | 
**CreatedAt** | **time.Time** | The date when this record was created. | 
**UpdatedAt** | Pointer to **time.Time** | The date of last update of the record. | [optional] 
**TunnelID** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "OpenVPN"]
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** | Secret access key for the OpenVPN tunnel. This is only the hashed value. The original value will not be retrievable after creation. | [optional] 

## Methods

### NewOpenVPNTunnel

`func NewOpenVPNTunnel(regionID string, gatewayID string, tunnelName string, createdAt time.Time, ) *OpenVPNTunnel`

NewOpenVPNTunnel instantiates a new OpenVPNTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenVPNTunnelWithDefaults

`func NewOpenVPNTunnelWithDefaults() *OpenVPNTunnel`

NewOpenVPNTunnelWithDefaults instantiates a new OpenVPNTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionID

`func (o *OpenVPNTunnel) GetRegionID() string`

GetRegionID returns the RegionID field if non-nil, zero value otherwise.

### GetRegionIDOk

`func (o *OpenVPNTunnel) GetRegionIDOk() (*string, bool)`

GetRegionIDOk returns a tuple with the RegionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionID

`func (o *OpenVPNTunnel) SetRegionID(v string)`

SetRegionID sets RegionID field to given value.


### GetGatewayID

`func (o *OpenVPNTunnel) GetGatewayID() string`

GetGatewayID returns the GatewayID field if non-nil, zero value otherwise.

### GetGatewayIDOk

`func (o *OpenVPNTunnel) GetGatewayIDOk() (*string, bool)`

GetGatewayIDOk returns a tuple with the GatewayID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayID

`func (o *OpenVPNTunnel) SetGatewayID(v string)`

SetGatewayID sets GatewayID field to given value.


### GetTunnelName

`func (o *OpenVPNTunnel) GetTunnelName() string`

GetTunnelName returns the TunnelName field if non-nil, zero value otherwise.

### GetTunnelNameOk

`func (o *OpenVPNTunnel) GetTunnelNameOk() (*string, bool)`

GetTunnelNameOk returns a tuple with the TunnelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelName

`func (o *OpenVPNTunnel) SetTunnelName(v string)`

SetTunnelName sets TunnelName field to given value.


### GetCreatedAt

`func (o *OpenVPNTunnel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenVPNTunnel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenVPNTunnel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OpenVPNTunnel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpenVPNTunnel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpenVPNTunnel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpenVPNTunnel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTunnelID

`func (o *OpenVPNTunnel) GetTunnelID() string`

GetTunnelID returns the TunnelID field if non-nil, zero value otherwise.

### GetTunnelIDOk

`func (o *OpenVPNTunnel) GetTunnelIDOk() (*string, bool)`

GetTunnelIDOk returns a tuple with the TunnelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelID

`func (o *OpenVPNTunnel) SetTunnelID(v string)`

SetTunnelID sets TunnelID field to given value.

### HasTunnelID

`func (o *OpenVPNTunnel) HasTunnelID() bool`

HasTunnelID returns a boolean if a field has been set.

### GetType

`func (o *OpenVPNTunnel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenVPNTunnel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenVPNTunnel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenVPNTunnel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *OpenVPNTunnel) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *OpenVPNTunnel) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *OpenVPNTunnel) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *OpenVPNTunnel) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *OpenVPNTunnel) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *OpenVPNTunnel) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *OpenVPNTunnel) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *OpenVPNTunnel) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


