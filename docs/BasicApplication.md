# BasicApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique application ID | 
**Name** | **string** | Application name | 
**Type** | **string** | Application type | 
**Enabled** | Pointer to **bool** | Application enabling status | [optional] 
**Policy** | Pointer to [**CommonApplicationPolicy**](CommonApplicationPolicy.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Application create date | [optional] 
**Network** | [**BasicApplicationAllOfNetwork**](BasicApplicationAllOfNetwork.md) |  | 
**Icon** | Pointer to [**BasicApplicationAllOfIcon**](BasicApplicationAllOfIcon.md) |  | [optional] 
**Users** | Pointer to [**[]BasicApplicationAllOfUsers**](BasicApplicationAllOfUsers.md) | List of IDs of assigned users | [optional] 
**Groups** | Pointer to [**[]BasicApplicationAllOfGroups**](BasicApplicationAllOfGroups.md) | List of IDs of assigned groups | [optional] 
**Auth** | [**ApplicationAuth**](ApplicationAuth.md) |  | 
**Alias** | [**ApplicationAlias**](ApplicationAlias.md) |  | 
**DisplayIconAtLogin** | **bool** | Determines if the application icon is displayed during user login | 
**Fqdn** | **string** | Fully qualified domain name (FQDN) associated with the application | 
**Host** | [**ApplicationHost**](ApplicationHost.md) |  | 
**Port** | [**ApplicationPort**](ApplicationPort.md) |  | 
**UpdatedAt** | Pointer to **string** | Application update date | [optional] 

## Methods

### NewBasicApplication

`func NewBasicApplication(id string, name string, type_ string, network BasicApplicationAllOfNetwork, auth ApplicationAuth, alias ApplicationAlias, displayIconAtLogin bool, fqdn string, host ApplicationHost, port ApplicationPort, ) *BasicApplication`

NewBasicApplication instantiates a new BasicApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicApplicationWithDefaults

`func NewBasicApplicationWithDefaults() *BasicApplication`

NewBasicApplicationWithDefaults instantiates a new BasicApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BasicApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BasicApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicApplication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BasicApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicApplication) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *BasicApplication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BasicApplication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BasicApplication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BasicApplication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicy

`func (o *BasicApplication) GetPolicy() CommonApplicationPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *BasicApplication) GetPolicyOk() (*CommonApplicationPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *BasicApplication) SetPolicy(v CommonApplicationPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *BasicApplication) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BasicApplication) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BasicApplication) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BasicApplication) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BasicApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *BasicApplication) GetNetwork() BasicApplicationAllOfNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BasicApplication) GetNetworkOk() (*BasicApplicationAllOfNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BasicApplication) SetNetwork(v BasicApplicationAllOfNetwork)`

SetNetwork sets Network field to given value.


### GetIcon

`func (o *BasicApplication) GetIcon() BasicApplicationAllOfIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BasicApplication) GetIconOk() (*BasicApplicationAllOfIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BasicApplication) SetIcon(v BasicApplicationAllOfIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *BasicApplication) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetUsers

`func (o *BasicApplication) GetUsers() []BasicApplicationAllOfUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BasicApplication) GetUsersOk() (*[]BasicApplicationAllOfUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BasicApplication) SetUsers(v []BasicApplicationAllOfUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BasicApplication) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *BasicApplication) GetGroups() []BasicApplicationAllOfGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BasicApplication) GetGroupsOk() (*[]BasicApplicationAllOfGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BasicApplication) SetGroups(v []BasicApplicationAllOfGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BasicApplication) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAuth

`func (o *BasicApplication) GetAuth() ApplicationAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *BasicApplication) GetAuthOk() (*ApplicationAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *BasicApplication) SetAuth(v ApplicationAuth)`

SetAuth sets Auth field to given value.


### GetAlias

`func (o *BasicApplication) GetAlias() ApplicationAlias`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *BasicApplication) GetAliasOk() (*ApplicationAlias, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *BasicApplication) SetAlias(v ApplicationAlias)`

SetAlias sets Alias field to given value.


### GetDisplayIconAtLogin

`func (o *BasicApplication) GetDisplayIconAtLogin() bool`

GetDisplayIconAtLogin returns the DisplayIconAtLogin field if non-nil, zero value otherwise.

### GetDisplayIconAtLoginOk

`func (o *BasicApplication) GetDisplayIconAtLoginOk() (*bool, bool)`

GetDisplayIconAtLoginOk returns a tuple with the DisplayIconAtLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIconAtLogin

`func (o *BasicApplication) SetDisplayIconAtLogin(v bool)`

SetDisplayIconAtLogin sets DisplayIconAtLogin field to given value.


### GetFqdn

`func (o *BasicApplication) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *BasicApplication) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *BasicApplication) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetHost

`func (o *BasicApplication) GetHost() ApplicationHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BasicApplication) GetHostOk() (*ApplicationHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BasicApplication) SetHost(v ApplicationHost)`

SetHost sets Host field to given value.


### GetPort

`func (o *BasicApplication) GetPort() ApplicationPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BasicApplication) GetPortOk() (*ApplicationPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BasicApplication) SetPort(v ApplicationPort)`

SetPort sets Port field to given value.


### GetUpdatedAt

`func (o *BasicApplication) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BasicApplication) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BasicApplication) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BasicApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


