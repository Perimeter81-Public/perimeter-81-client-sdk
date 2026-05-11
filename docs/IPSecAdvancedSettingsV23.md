# IPSecAdvancedSettingsV23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyExchange** | **string** |  | 
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | 
**Phase2** | [**IPSecPhaseConfigV23**](IPSecPhaseConfigV23.md) |  | 

## Methods

### NewIPSecAdvancedSettingsV23

`func NewIPSecAdvancedSettingsV23(keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfigV23, phase2 IPSecPhaseConfigV23, ) *IPSecAdvancedSettingsV23`

NewIPSecAdvancedSettingsV23 instantiates a new IPSecAdvancedSettingsV23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecAdvancedSettingsV23WithDefaults

`func NewIPSecAdvancedSettingsV23WithDefaults() *IPSecAdvancedSettingsV23`

NewIPSecAdvancedSettingsV23WithDefaults instantiates a new IPSecAdvancedSettingsV23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyExchange

`func (o *IPSecAdvancedSettingsV23) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *IPSecAdvancedSettingsV23) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *IPSecAdvancedSettingsV23) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *IPSecAdvancedSettingsV23) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *IPSecAdvancedSettingsV23) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *IPSecAdvancedSettingsV23) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *IPSecAdvancedSettingsV23) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *IPSecAdvancedSettingsV23) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *IPSecAdvancedSettingsV23) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *IPSecAdvancedSettingsV23) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *IPSecAdvancedSettingsV23) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *IPSecAdvancedSettingsV23) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *IPSecAdvancedSettingsV23) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *IPSecAdvancedSettingsV23) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *IPSecAdvancedSettingsV23) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *IPSecAdvancedSettingsV23) GetPhase1() IPSecPhaseConfigV23`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *IPSecAdvancedSettingsV23) GetPhase1Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *IPSecAdvancedSettingsV23) SetPhase1(v IPSecPhaseConfigV23)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *IPSecAdvancedSettingsV23) GetPhase2() IPSecPhaseConfigV23`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *IPSecAdvancedSettingsV23) GetPhase2Ok() (*IPSecPhaseConfigV23, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *IPSecAdvancedSettingsV23) SetPhase2(v IPSecPhaseConfigV23)`

SetPhase2 sets Phase2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


