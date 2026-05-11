# IPSecAdvancedSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyExchange** | Pointer to **string** |  | [optional] 
**IkeLifeTime** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **string** |  | [optional] 
**DpdDelay** | Pointer to **string** |  | [optional] 
**DpdTimeout** | Pointer to **string** |  | [optional] 
**Phase1** | Pointer to [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | [optional] 
**Phase2** | Pointer to [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | [optional] 

## Methods

### NewIPSecAdvancedSettingsUpdate

`func NewIPSecAdvancedSettingsUpdate() *IPSecAdvancedSettingsUpdate`

NewIPSecAdvancedSettingsUpdate instantiates a new IPSecAdvancedSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecAdvancedSettingsUpdateWithDefaults

`func NewIPSecAdvancedSettingsUpdateWithDefaults() *IPSecAdvancedSettingsUpdate`

NewIPSecAdvancedSettingsUpdateWithDefaults instantiates a new IPSecAdvancedSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyExchange

`func (o *IPSecAdvancedSettingsUpdate) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *IPSecAdvancedSettingsUpdate) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *IPSecAdvancedSettingsUpdate) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.

### HasKeyExchange

`func (o *IPSecAdvancedSettingsUpdate) HasKeyExchange() bool`

HasKeyExchange returns a boolean if a field has been set.

### GetIkeLifeTime

`func (o *IPSecAdvancedSettingsUpdate) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *IPSecAdvancedSettingsUpdate) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *IPSecAdvancedSettingsUpdate) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.

### HasIkeLifeTime

`func (o *IPSecAdvancedSettingsUpdate) HasIkeLifeTime() bool`

HasIkeLifeTime returns a boolean if a field has been set.

### GetLifetime

`func (o *IPSecAdvancedSettingsUpdate) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *IPSecAdvancedSettingsUpdate) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *IPSecAdvancedSettingsUpdate) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *IPSecAdvancedSettingsUpdate) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetDpdDelay

`func (o *IPSecAdvancedSettingsUpdate) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *IPSecAdvancedSettingsUpdate) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *IPSecAdvancedSettingsUpdate) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.

### HasDpdDelay

`func (o *IPSecAdvancedSettingsUpdate) HasDpdDelay() bool`

HasDpdDelay returns a boolean if a field has been set.

### GetDpdTimeout

`func (o *IPSecAdvancedSettingsUpdate) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *IPSecAdvancedSettingsUpdate) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *IPSecAdvancedSettingsUpdate) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.

### HasDpdTimeout

`func (o *IPSecAdvancedSettingsUpdate) HasDpdTimeout() bool`

HasDpdTimeout returns a boolean if a field has been set.

### GetPhase1

`func (o *IPSecAdvancedSettingsUpdate) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *IPSecAdvancedSettingsUpdate) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *IPSecAdvancedSettingsUpdate) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.

### HasPhase1

`func (o *IPSecAdvancedSettingsUpdate) HasPhase1() bool`

HasPhase1 returns a boolean if a field has been set.

### GetPhase2

`func (o *IPSecAdvancedSettingsUpdate) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *IPSecAdvancedSettingsUpdate) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *IPSecAdvancedSettingsUpdate) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.

### HasPhase2

`func (o *IPSecAdvancedSettingsUpdate) HasPhase2() bool`

HasPhase2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


