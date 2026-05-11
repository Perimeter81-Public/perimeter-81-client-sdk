# HttpCreateApplication

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
**Headers** | **map[string]interface{}** | Application specific headers. Keys must not contain . and $ | 
**Attributes** | [**HttpAttributes**](HttpAttributes.md) |  | 

## Methods

### NewHttpCreateApplication

`func NewHttpCreateApplication(name string, type_ string, network string, host NullableOneOfFixedHostIdpHost, port NullableOneOfFixedPortIdpPort, users []string, groups []string, headers map[string]interface{}, attributes HttpAttributes, ) *HttpCreateApplication`

NewHttpCreateApplication instantiates a new HttpCreateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpCreateApplicationWithDefaults

`func NewHttpCreateApplicationWithDefaults() *HttpCreateApplication`

NewHttpCreateApplicationWithDefaults instantiates a new HttpCreateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HttpCreateApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpCreateApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpCreateApplication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *HttpCreateApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpCreateApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpCreateApplication) SetType(v string)`

SetType sets Type field to given value.


### GetNetwork

`func (o *HttpCreateApplication) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HttpCreateApplication) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HttpCreateApplication) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetHost

`func (o *HttpCreateApplication) GetHost() OneOfFixedHostIdpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HttpCreateApplication) GetHostOk() (*OneOfFixedHostIdpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HttpCreateApplication) SetHost(v OneOfFixedHostIdpHost)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *HttpCreateApplication) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HttpCreateApplication) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HttpCreateApplication) GetPort() OneOfFixedPortIdpPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpCreateApplication) GetPortOk() (*OneOfFixedPortIdpPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpCreateApplication) SetPort(v OneOfFixedPortIdpPort)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *HttpCreateApplication) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HttpCreateApplication) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsers

`func (o *HttpCreateApplication) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *HttpCreateApplication) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *HttpCreateApplication) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetGroups

`func (o *HttpCreateApplication) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *HttpCreateApplication) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *HttpCreateApplication) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetHeaders

`func (o *HttpCreateApplication) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpCreateApplication) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpCreateApplication) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.


### GetAttributes

`func (o *HttpCreateApplication) GetAttributes() HttpAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HttpCreateApplication) GetAttributesOk() (*HttpAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HttpCreateApplication) SetAttributes(v HttpAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


