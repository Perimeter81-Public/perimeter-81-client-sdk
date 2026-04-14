# CreateInstancesInNetworkPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | Region ID. | 
**Idle** | **bool** | Create the gateway as disabled if true. | [default to true]

## Methods

### NewCreateInstancesInNetworkPayload

`func NewCreateInstancesInNetworkPayload(regionId string, idle bool, ) *CreateInstancesInNetworkPayload`

NewCreateInstancesInNetworkPayload instantiates a new CreateInstancesInNetworkPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstancesInNetworkPayloadWithDefaults

`func NewCreateInstancesInNetworkPayloadWithDefaults() *CreateInstancesInNetworkPayload`

NewCreateInstancesInNetworkPayloadWithDefaults instantiates a new CreateInstancesInNetworkPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *CreateInstancesInNetworkPayload) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CreateInstancesInNetworkPayload) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CreateInstancesInNetworkPayload) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetIdle

`func (o *CreateInstancesInNetworkPayload) GetIdle() bool`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *CreateInstancesInNetworkPayload) GetIdleOk() (*bool, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *CreateInstancesInNetworkPayload) SetIdle(v bool)`

SetIdle sets Idle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


