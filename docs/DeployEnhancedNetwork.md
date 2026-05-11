# DeployEnhancedNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**DeployEnhancedNetworkNetwork**](DeployEnhancedNetworkNetwork.md) |  | 
**Regions** | [**[]EnhancedRegionCreate**](EnhancedRegionCreate.md) |  | 

## Methods

### NewDeployEnhancedNetwork

`func NewDeployEnhancedNetwork(network DeployEnhancedNetworkNetwork, regions []EnhancedRegionCreate, ) *DeployEnhancedNetwork`

NewDeployEnhancedNetwork instantiates a new DeployEnhancedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployEnhancedNetworkWithDefaults

`func NewDeployEnhancedNetworkWithDefaults() *DeployEnhancedNetwork`

NewDeployEnhancedNetworkWithDefaults instantiates a new DeployEnhancedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *DeployEnhancedNetwork) GetNetwork() DeployEnhancedNetworkNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DeployEnhancedNetwork) GetNetworkOk() (*DeployEnhancedNetworkNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DeployEnhancedNetwork) SetNetwork(v DeployEnhancedNetworkNetwork)`

SetNetwork sets Network field to given value.


### GetRegions

`func (o *DeployEnhancedNetwork) GetRegions() []EnhancedRegionCreate`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *DeployEnhancedNetwork) GetRegionsOk() (*[]EnhancedRegionCreate, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *DeployEnhancedNetwork) SetRegions(v []EnhancedRegionCreate)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


