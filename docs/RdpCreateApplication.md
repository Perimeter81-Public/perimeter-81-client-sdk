# RdpCreateApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Application name | 
**Type** | **string** | Application type | 
**Network** | **string** | Application network ID. To get the ID use endpoint &#39;/networks&#39; with [network:read] permission. | 
**Host** | [**NullableOneOfFixedHostIdpHost**](oneOf&lt;FixedHost,IdpHost&gt;.md) |  | 
**Port** | [**NullableOneOfFixedPortIdpPort**](oneOf&lt;FixedPort,IdpPort&gt;.md) |  | 
**Users** | **[]string** |  | 
**Groups** | **[]string** |  | 
**Attributes** | [**RdpAttributes**](RdpAttributes.md) |  | 
**Auth** | [**ApplicationAuth**](ApplicationAuth.md) |  | 

## Methods

### NewRdpCreateApplication

`func NewRdpCreateApplication(name string, type_ string, network string, host NullableOneOfFixedHostIdpHost, port NullableOneOfFixedPortIdpPort, users []string, groups []string, attributes RdpAttributes, auth ApplicationAuth, ) *RdpCreateApplication`

NewRdpCreateApplication instantiates a new RdpCreateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdpCreateApplicationWithDefaults

`func NewRdpCreateApplicationWithDefaults() *RdpCreateApplication`

NewRdpCreateApplicationWithDefaults instantiates a new RdpCreateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RdpCreateApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RdpCreateApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RdpCreateApplication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RdpCreateApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RdpCreateApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RdpCreateApplication) SetType(v string)`

SetType sets Type field to given value.


### GetNetwork

`func (o *RdpCreateApplication) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RdpCreateApplication) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RdpCreateApplication) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetHost

`func (o *RdpCreateApplication) GetHost() OneOfFixedHostIdpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RdpCreateApplication) GetHostOk() (*OneOfFixedHostIdpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RdpCreateApplication) SetHost(v OneOfFixedHostIdpHost)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *RdpCreateApplication) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RdpCreateApplication) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *RdpCreateApplication) GetPort() OneOfFixedPortIdpPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RdpCreateApplication) GetPortOk() (*OneOfFixedPortIdpPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RdpCreateApplication) SetPort(v OneOfFixedPortIdpPort)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *RdpCreateApplication) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *RdpCreateApplication) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsers

`func (o *RdpCreateApplication) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RdpCreateApplication) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RdpCreateApplication) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetGroups

`func (o *RdpCreateApplication) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RdpCreateApplication) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RdpCreateApplication) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetAttributes

`func (o *RdpCreateApplication) GetAttributes() RdpAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RdpCreateApplication) GetAttributesOk() (*RdpAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RdpCreateApplication) SetAttributes(v RdpAttributes)`

SetAttributes sets Attributes field to given value.


### GetAuth

`func (o *RdpCreateApplication) GetAuth() ApplicationAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RdpCreateApplication) GetAuthOk() (*ApplicationAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RdpCreateApplication) SetAuth(v ApplicationAuth)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


