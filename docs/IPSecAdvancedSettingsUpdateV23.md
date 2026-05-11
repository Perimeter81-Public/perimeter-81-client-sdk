# IPSecAdvancedSettingsUpdateV23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyExchange** | Pointer to **string** |  | [optional] 
**IkeLifeTime** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **string** |  | [optional] 
**DpdDelay** | Pointer to **string** |  | [optional] 
**DpdTimeout** | Pointer to **string** |  | [optional] 
**Phase1** | Pointer to [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | [optional] 
**Phase2** | Pointer to [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | [optional] 

## Methods

### NewIPSecAdvancedSettingsUpdateV23

`func NewIPSecAdvancedSettingsUpdateV23() *IPSecAdvancedSettingsUpdateV23`

NewIPSecAdvancedSettingsUpdateV23 instantiates a new IPSecAdvancedSettingsUpdateV23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecAdvancedSettingsUpdateV23WithDefaults

`func NewIPSecAdvancedSettingsUpdateV23WithDefaults() *IPSecAdvancedSettingsUpdateV23`

NewIPSecAdvancedSettingsUpdateV23WithDefaults instantiates a new IPSecAdvancedSettingsUpdateV23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyExchange

`func (o *IPSecAdvancedSettingsUpdateV23) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *IPSecAdvancedSettingsUpdateV23) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *IPSecAdvancedSettingsUpdateV23) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.

### HasKeyExchange

`func (o *IPSecAdvancedSettingsUpdateV23) HasKeyExchange() bool`

HasKeyExchange returns a boolean if a field has been set.

### GetIkeLifeTime

`func (o *IPSecAdvancedSettingsUpdateV23) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *IPSecAdvancedSettingsUpdateV23) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *IPSecAdvancedSettingsUpdateV23) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.

### HasIkeLifeTime

`func (o *IPSecAdvancedSettingsUpdateV23) HasIkeLifeTime() bool`

HasIkeLifeTime returns a boolean if a field has been set.

### GetLifetime

`func (o *IPSecAdvancedSettingsUpdateV23) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *IPSecAdvancedSettingsUpdateV23) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *IPSecAdvancedSettingsUpdateV23) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *IPSecAdvancedSettingsUpdateV23) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetDpdDelay

`func (o *IPSecAdvancedSettingsUpdateV23) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *IPSecAdvancedSettingsUpdateV23) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *IPSecAdvancedSettingsUpdateV23) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.

### HasDpdDelay

`func (o *IPSecAdvancedSettingsUpdateV23) HasDpdDelay() bool`

HasDpdDelay returns a boolean if a field has been set.

### GetDpdTimeout

`func (o *IPSecAdvancedSettingsUpdateV23) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *IPSecAdvancedSettingsUpdateV23) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *IPSecAdvancedSettingsUpdateV23) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.

### HasDpdTimeout

`func (o *IPSecAdvancedSettingsUpdateV23) HasDpdTimeout() bool`

HasDpdTimeout returns a boolean if a field has been set.

### GetPhase1

`func (o *IPSecAdvancedSettingsUpdateV23) GetPhase1() IPSecPhaseConfigV23`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *IPSecAdvancedSettingsUpdateV23) GetPhase1Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *IPSecAdvancedSettingsUpdateV23) SetPhase1(v IPSecPhaseConfigV23)`

SetPhase1 sets Phase1 field to given value.

### HasPhase1

`func (o *IPSecAdvancedSettingsUpdateV23) HasPhase1() bool`

HasPhase1 returns a boolean if a field has been set.

### GetPhase2

`func (o *IPSecAdvancedSettingsUpdateV23) GetPhase2() IPSecPhaseConfigV23`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *IPSecAdvancedSettingsUpdateV23) GetPhase2Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *IPSecAdvancedSettingsUpdateV23) SetPhase2(v IPSecPhaseConfigV23)`

SetPhase2 sets Phase2 field to given value.

### HasPhase2

`func (o *IPSecAdvancedSettingsUpdateV23) HasPhase2() bool`

HasPhase2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


