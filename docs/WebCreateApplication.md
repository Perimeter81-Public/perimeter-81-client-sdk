# WebCreateApplication

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

## Methods

### NewWebCreateApplication

`func NewWebCreateApplication(name string, type_ string, network string, host NullableOneOfFixedHostIdpHost, port NullableOneOfFixedPortIdpPort, users []string, groups []string, headers map[string]interface{}, ) *WebCreateApplication`

NewWebCreateApplication instantiates a new WebCreateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebCreateApplicationWithDefaults

`func NewWebCreateApplicationWithDefaults() *WebCreateApplication`

NewWebCreateApplicationWithDefaults instantiates a new WebCreateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebCreateApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebCreateApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebCreateApplication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebCreateApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebCreateApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebCreateApplication) SetType(v string)`

SetType sets Type field to given value.


### GetNetwork

`func (o *WebCreateApplication) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WebCreateApplication) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WebCreateApplication) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetHost

`func (o *WebCreateApplication) GetHost() OneOfFixedHostIdpHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *WebCreateApplication) GetHostOk() (*OneOfFixedHostIdpHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *WebCreateApplication) SetHost(v OneOfFixedHostIdpHost)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *WebCreateApplication) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *WebCreateApplication) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *WebCreateApplication) GetPort() OneOfFixedPortIdpPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WebCreateApplication) GetPortOk() (*OneOfFixedPortIdpPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WebCreateApplication) SetPort(v OneOfFixedPortIdpPort)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *WebCreateApplication) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *WebCreateApplication) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsers

`func (o *WebCreateApplication) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *WebCreateApplication) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *WebCreateApplication) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetGroups

`func (o *WebCreateApplication) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *WebCreateApplication) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *WebCreateApplication) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetHeaders

`func (o *WebCreateApplication) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebCreateApplication) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebCreateApplication) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


