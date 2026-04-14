# SourcesAndDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to **[]string** | List of IDs of users. | [optional] 
**Groups** | Pointer to **[]string** | List of IDs of groups. | [optional] 
**Addresses** | **[]string** | List of IDs of address objects. | 

## Methods

### NewSourcesAndDestinations

`func NewSourcesAndDestinations(addresses []string, ) *SourcesAndDestinations`

NewSourcesAndDestinations instantiates a new SourcesAndDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesAndDestinationsWithDefaults

`func NewSourcesAndDestinationsWithDefaults() *SourcesAndDestinations`

NewSourcesAndDestinationsWithDefaults instantiates a new SourcesAndDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *SourcesAndDestinations) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SourcesAndDestinations) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SourcesAndDestinations) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SourcesAndDestinations) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *SourcesAndDestinations) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SourcesAndDestinations) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SourcesAndDestinations) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SourcesAndDestinations) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAddresses

`func (o *SourcesAndDestinations) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *SourcesAndDestinations) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *SourcesAndDestinations) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


