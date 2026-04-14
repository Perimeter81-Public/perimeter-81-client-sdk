# DeployEnhancedNetworkNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | Subnet of the network. Cannot be changed later. Allowed private subnet ranges are 10.0.0.0/12-22, 172.16.0.0/12-22, 192.168.0.0/16-22, or 198.18.0.0/15-22. | [optional] [default to "10.255.0.0/16"]
**Name** | **string** | Network name. | 
**Tags** | Pointer to **[]string** | List of network tags. | [optional] 

## Methods

### NewDeployEnhancedNetworkNetwork

`func NewDeployEnhancedNetworkNetwork(name string, ) *DeployEnhancedNetworkNetwork`

NewDeployEnhancedNetworkNetwork instantiates a new DeployEnhancedNetworkNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployEnhancedNetworkNetworkWithDefaults

`func NewDeployEnhancedNetworkNetworkWithDefaults() *DeployEnhancedNetworkNetwork`

NewDeployEnhancedNetworkNetworkWithDefaults instantiates a new DeployEnhancedNetworkNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *DeployEnhancedNetworkNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DeployEnhancedNetworkNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DeployEnhancedNetworkNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *DeployEnhancedNetworkNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetName

`func (o *DeployEnhancedNetworkNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployEnhancedNetworkNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployEnhancedNetworkNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *DeployEnhancedNetworkNetwork) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeployEnhancedNetworkNetwork) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeployEnhancedNetworkNetwork) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeployEnhancedNetworkNetwork) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


