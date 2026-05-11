# RemoveRegionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | Pointer to **string** | Region ID. | [optional] 
**Instances** | Pointer to [**[]RemoveInstancePayload**](RemoveInstancePayload.md) |  | [optional] 

## Methods

### NewRemoveRegionPayload

`func NewRemoveRegionPayload() *RemoveRegionPayload`

NewRemoveRegionPayload instantiates a new RemoveRegionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveRegionPayloadWithDefaults

`func NewRemoveRegionPayloadWithDefaults() *RemoveRegionPayload`

NewRemoveRegionPayloadWithDefaults instantiates a new RemoveRegionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *RemoveRegionPayload) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *RemoveRegionPayload) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *RemoveRegionPayload) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *RemoveRegionPayload) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetInstances

`func (o *RemoveRegionPayload) GetInstances() []RemoveInstancePayload`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *RemoveRegionPayload) GetInstancesOk() (*[]RemoveInstancePayload, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *RemoveRegionPayload) SetInstances(v []RemoveInstancePayload)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *RemoveRegionPayload) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


