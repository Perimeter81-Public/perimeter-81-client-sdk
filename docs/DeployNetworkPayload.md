# DeployNetworkPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**CreateNetworkPayload**](CreateNetworkPayload.md) |  | 
**Regions** | [**[]CreateRegionInNetworkPayload**](CreateRegionInNetworkPayload.md) | Region list. | 

## Methods

### NewDeployNetworkPayload

`func NewDeployNetworkPayload(network CreateNetworkPayload, regions []CreateRegionInNetworkPayload, ) *DeployNetworkPayload`

NewDeployNetworkPayload instantiates a new DeployNetworkPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployNetworkPayloadWithDefaults

`func NewDeployNetworkPayloadWithDefaults() *DeployNetworkPayload`

NewDeployNetworkPayloadWithDefaults instantiates a new DeployNetworkPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *DeployNetworkPayload) GetNetwork() CreateNetworkPayload`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DeployNetworkPayload) GetNetworkOk() (*CreateNetworkPayload, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DeployNetworkPayload) SetNetwork(v CreateNetworkPayload)`

SetNetwork sets Network field to given value.


### GetRegions

`func (o *DeployNetworkPayload) GetRegions() []CreateRegionInNetworkPayload`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *DeployNetworkPayload) GetRegionsOk() (*[]CreateRegionInNetworkPayload, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *DeployNetworkPayload) SetRegions(v []CreateRegionInNetworkPayload)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


