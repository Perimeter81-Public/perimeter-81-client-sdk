# RdpApplication

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
**Attributes** | Pointer to [**RdpAttributes**](RdpAttributes.md) |  | [optional] 

## Methods

### NewRdpApplication

`func NewRdpApplication(id string, name string, type_ string, network BasicApplicationAllOfNetwork, auth ApplicationAuth, alias ApplicationAlias, displayIconAtLogin bool, fqdn string, host ApplicationHost, port ApplicationPort, ) *RdpApplication`

NewRdpApplication instantiates a new RdpApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdpApplicationWithDefaults

`func NewRdpApplicationWithDefaults() *RdpApplication`

NewRdpApplicationWithDefaults instantiates a new RdpApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RdpApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RdpApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RdpApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RdpApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RdpApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RdpApplication) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RdpApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RdpApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RdpApplication) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *RdpApplication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RdpApplication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RdpApplication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RdpApplication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPolicy

`func (o *RdpApplication) GetPolicy() CommonApplicationPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RdpApplication) GetPolicyOk() (*CommonApplicationPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RdpApplication) SetPolicy(v CommonApplicationPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *RdpApplication) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RdpApplication) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RdpApplication) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RdpApplication) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RdpApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *RdpApplication) GetNetwork() BasicApplicationAllOfNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RdpApplication) GetNetworkOk() (*BasicApplicationAllOfNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RdpApplication) SetNetwork(v BasicApplicationAllOfNetwork)`

SetNetwork sets Network field to given value.


### GetIcon

`func (o *RdpApplication) GetIcon() BasicApplicationAllOfIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RdpApplication) GetIconOk() (*BasicApplicationAllOfIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RdpApplication) SetIcon(v BasicApplicationAllOfIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RdpApplication) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetUsers

`func (o *RdpApplication) GetUsers() []BasicApplicationAllOfUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RdpApplication) GetUsersOk() (*[]BasicApplicationAllOfUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RdpApplication) SetUsers(v []BasicApplicationAllOfUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RdpApplication) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetGroups

`func (o *RdpApplication) GetGroups() []BasicApplicationAllOfGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RdpApplication) GetGroupsOk() (*[]BasicApplicationAllOfGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RdpApplication) SetGroups(v []BasicApplicationAllOfGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RdpApplication) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAuth

`func (o *RdpApplication) GetAuth() ApplicationAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RdpApplication) GetAuthOk() (*ApplicationAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RdpApplication) SetAuth(v ApplicationAuth)`

SetAuth sets Auth field to given value.


### GetAlias

`func (o *RdpApplication) GetAlias() ApplicationAlias`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *RdpApplication) GetAliasOk() (*ApplicationAlias, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *RdpApplication) SetAlias(v ApplicationAlias)`

SetAlias sets Alias field to given value.


### GetDisplayIconAtLogin

`func (o *RdpApplication) GetDisplayIconAtLogin() bool`

GetDisplayIconAtLogin returns the DisplayIconAtLogin field if non-nil, zero value otherwise.

### GetDisplayIconAtLoginOk

`func (o *RdpApplication) GetDisplayIconAtLoginOk() (*bool, bool)`

GetDisplayIconAtLoginOk returns a tuple with the DisplayIconAtLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIconAtLogin

`func (o *RdpApplication) SetDisplayIconAtLogin(v bool)`

SetDisplayIconAtLogin sets DisplayIconAtLogin field to given value.


### GetFqdn

`func (o *RdpApplication) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *RdpApplication) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *RdpApplication) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetHost

`func (o *RdpApplication) GetHost() ApplicationHost`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RdpApplication) GetHostOk() (*ApplicationHost, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RdpApplication) SetHost(v ApplicationHost)`

SetHost sets Host field to given value.


### GetPort

`func (o *RdpApplication) GetPort() ApplicationPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RdpApplication) GetPortOk() (*ApplicationPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RdpApplication) SetPort(v ApplicationPort)`

SetPort sets Port field to given value.


### GetUpdatedAt

`func (o *RdpApplication) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RdpApplication) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RdpApplication) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RdpApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttributes

`func (o *RdpApplication) GetAttributes() RdpAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RdpApplication) GetAttributesOk() (*RdpAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RdpApplication) SetAttributes(v RdpAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RdpApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


