# ApplicationsListObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique application ID | 
**Name** | **string** | Application name | 
**Type** | **string** | Application type | 
**Enabled** | Pointer to **bool** | Application enabling status | [optional] 
**Policy** | Pointer to [**CommonApplicationPolicy**](CommonApplicationPolicy.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Application create date | [optional] 
**Network** | [**ApplicationsListObjectAllOfNetwork**](ApplicationsListObjectAllOfNetwork.md) |  | 
**Host** | **string** | Application host | 
**Port** | [**ApplicationsListObjectAllOfPort**](ApplicationsListObjectAllOfPort.md) |  | 

## Methods

### NewApplicationsListObject

`func NewApplicationsListObject(id string, name string, type_ string, network ApplicationsListObjectAllOfNetwork, host string, port ApplicationsListObjectAllOfPort, ) *ApplicationsListObject`

NewApplicationsListObject instantiates a new ApplicationsListObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsListObjectWithDefaults

`func NewApplicationsListObjectWithDefaults() *ApplicationsListObject`

NewApplicationsListObjectWithDefaults instantiates a new ApplicationsListObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationsListObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationsListObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationsListObject) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationsListObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationsListObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationsListObject) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApplicationsListObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationsListObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationsListObject) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *ApplicationsListObject) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationsListObject) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationsListObject) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationsListObject) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicy

`func (o *ApplicationsListObject) GetPolicy() CommonApplicationPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ApplicationsListObject) GetPolicyOk() (*CommonApplicationPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ApplicationsListObject) SetPolicy(v CommonApplicationPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ApplicationsListObject) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationsListObject) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationsListObject) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationsListObject) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationsListObject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *ApplicationsListObject) GetNetwork() ApplicationsListObjectAllOfNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ApplicationsListObject) GetNetworkOk() (*ApplicationsListObjectAllOfNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ApplicationsListObject) SetNetwork(v ApplicationsListObjectAllOfNetwork)`

SetNetwork sets Network field to given value.


### GetHost

`func (o *ApplicationsListObject) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ApplicationsListObject) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ApplicationsListObject) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ApplicationsListObject) GetPort() ApplicationsListObjectAllOfPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApplicationsListObject) GetPortOk() (*ApplicationsListObjectAllOfPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApplicationsListObject) SetPort(v ApplicationsListObjectAllOfPort)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


