# WireGuradDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteEndpoint** | **string** |  | 
**RemoteSubnets** | **[]string** |  | 

## Methods

### NewWireGuradDetails

`func NewWireGuradDetails(remoteEndpoint string, remoteSubnets []string, ) *WireGuradDetails`

NewWireGuradDetails instantiates a new WireGuradDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireGuradDetailsWithDefaults

`func NewWireGuradDetailsWithDefaults() *WireGuradDetails`

NewWireGuradDetailsWithDefaults instantiates a new WireGuradDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteEndpoint

`func (o *WireGuradDetails) GetRemoteEndpoint() string`

GetRemoteEndpoint returns the RemoteEndpoint field if non-nil, zero value otherwise.

### GetRemoteEndpointOk

`func (o *WireGuradDetails) GetRemoteEndpointOk() (*string, bool)`

GetRemoteEndpointOk returns a tuple with the RemoteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndpoint

`func (o *WireGuradDetails) SetRemoteEndpoint(v string)`

SetRemoteEndpoint sets RemoteEndpoint field to given value.


### GetRemoteSubnets

`func (o *WireGuradDetails) GetRemoteSubnets() []string`

GetRemoteSubnets returns the RemoteSubnets field if non-nil, zero value otherwise.

### GetRemoteSubnetsOk

`func (o *WireGuradDetails) GetRemoteSubnetsOk() (*[]string, bool)`

GetRemoteSubnetsOk returns a tuple with the RemoteSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnets

`func (o *WireGuradDetails) SetRemoteSubnets(v []string)`

SetRemoteSubnets sets RemoteSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


