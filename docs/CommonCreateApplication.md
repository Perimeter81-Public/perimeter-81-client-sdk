# CommonCreateApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Application name | 
**Type** | **string** | Application type | 
**Network** | **string** | Application network ID. To get the ID use endpoint &#39;/networks&#39; with [network:read] permission. | 
**Host** | [**NullableOneOfFixedHostIdpHost**](oneOf&lt;FixedHost,IdpHost&gt;.md) |  | 
**Port** | [**NullableOneOfFixedPortIdpPort**](oneOf&lt;FixedPort,IdpPort&gt;.md) |  | 
**Users** | Pointer to **[]string** | List of IDs of assigned users. To get the IDs use endpoint &#39;/users&#39; with [user:read] permission. | [optional] 
**Groups** | Pointer to **[]string** | List of IDs of assigned groups. To get the IDs use endpoint &#39;/groups&#39; with [group:read] permission. | [optional] 

## Methods

### NewCommonCreateApplication

`func NewCommonCreateApplication(name string, type_ string, network string, host NullableOneOfFixedHostIdpHost, port NullableOneOfFixedPortIdpPort, ) *CommonCreateApplication`

NewCommonCreateApplication instantiates a new CommonCreateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCreateApplicationWithDefaults

`func NewCommonCreateApplicationWithDefaults() *CommonCreateApplication`

NewCommonCreateApplicationWithDefaults instantiates a new CommonCreateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonCreateApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonCreateApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonCreateApplication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CommonCreateApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonCreateApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonCreateApplication) SetType(v string)`

SetType sets Type field to given value.


### GetNetwork

`func (o *CommonCreateApplication) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CommonCreateApplication) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CommonCreateApplication) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetHost

`func (o *CommonCreateApplication) GetHost() OneOfFixedHostIdpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CommonCreateApplication) GetHostOk() (*OneOfFixedHostIdpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CommonCreateApplication) SetHost(v OneOfFixedHostIdpHost)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *CommonCreateApplication) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *CommonCreateApplication) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *CommonCreateApplication) GetPort() OneOfFixedPortIdpPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CommonCreateApplication) GetPortOk() (*OneOfFixedPortIdpPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CommonCreateApplication) SetPort(v OneOfFixedPortIdpPort)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *CommonCreateApplication) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CommonCreateApplication) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsers

`func (o *CommonCreateApplication) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CommonCreateApplication) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CommonCreateApplication) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CommonCreateApplication) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *CommonCreateApplication) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CommonCreateApplication) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CommonCreateApplication) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CommonCreateApplication) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


