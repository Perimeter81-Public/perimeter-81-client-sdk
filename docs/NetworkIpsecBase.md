# NetworkIpsecBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyExchange** | **string** |  | 
**IkeLifeTime** | **string** |  | 
**Lifetime** | **string** |  | 
**DpdDelay** | **string** |  | 
**DpdTimeout** | **string** |  | 
**Phase1** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Phase2** | [**IPSecPhaseConfig**](IPSecPhaseConfig.md) |  | 
**Right** | **string** |  | 
**RightID** | [**NetworkIpsecBaseAllOfRightID**](NetworkIpsecBaseAllOfRightID.md) |  | 
**Passphrase** | **string** |  | 
**DpdAction** | **string** | dpdAction. | 
**LeftSubnets** | **[]string** |  | 
**RightSubnets** | **[]string** |  | 

## Methods

### NewNetworkIpsecBase

`func NewNetworkIpsecBase(keyExchange string, ikeLifeTime string, lifetime string, dpdDelay string, dpdTimeout string, phase1 IPSecPhaseConfig, phase2 IPSecPhaseConfig, right string, rightID NetworkIpsecBaseAllOfRightID, passphrase string, dpdAction string, leftSubnets []string, rightSubnets []string, ) *NetworkIpsecBase`

NewNetworkIpsecBase instantiates a new NetworkIpsecBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIpsecBaseWithDefaults

`func NewNetworkIpsecBaseWithDefaults() *NetworkIpsecBase`

NewNetworkIpsecBaseWithDefaults instantiates a new NetworkIpsecBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyExchange

`func (o *NetworkIpsecBase) GetKeyExchange() string`

GetKeyExchange returns the KeyExchange field if non-nil, zero value otherwise.

### GetKeyExchangeOk

`func (o *NetworkIpsecBase) GetKeyExchangeOk() (*string, bool)`

GetKeyExchangeOk returns a tuple with the KeyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchange

`func (o *NetworkIpsecBase) SetKeyExchange(v string)`

SetKeyExchange sets KeyExchange field to given value.


### GetIkeLifeTime

`func (o *NetworkIpsecBase) GetIkeLifeTime() string`

GetIkeLifeTime returns the IkeLifeTime field if non-nil, zero value otherwise.

### GetIkeLifeTimeOk

`func (o *NetworkIpsecBase) GetIkeLifeTimeOk() (*string, bool)`

GetIkeLifeTimeOk returns a tuple with the IkeLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifeTime

`func (o *NetworkIpsecBase) SetIkeLifeTime(v string)`

SetIkeLifeTime sets IkeLifeTime field to given value.


### GetLifetime

`func (o *NetworkIpsecBase) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *NetworkIpsecBase) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *NetworkIpsecBase) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.


### GetDpdDelay

`func (o *NetworkIpsecBase) GetDpdDelay() string`

GetDpdDelay returns the DpdDelay field if non-nil, zero value otherwise.

### GetDpdDelayOk

`func (o *NetworkIpsecBase) GetDpdDelayOk() (*string, bool)`

GetDpdDelayOk returns a tuple with the DpdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdDelay

`func (o *NetworkIpsecBase) SetDpdDelay(v string)`

SetDpdDelay sets DpdDelay field to given value.


### GetDpdTimeout

`func (o *NetworkIpsecBase) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *NetworkIpsecBase) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *NetworkIpsecBase) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.


### GetPhase1

`func (o *NetworkIpsecBase) GetPhase1() IPSecPhaseConfig`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *NetworkIpsecBase) GetPhase1Ok() (*IPSecPhaseConfig, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *NetworkIpsecBase) SetPhase1(v IPSecPhaseConfig)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *NetworkIpsecBase) GetPhase2() IPSecPhaseConfig`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *NetworkIpsecBase) GetPhase2Ok() (*IPSecPhaseConfig, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *NetworkIpsecBase) SetPhase2(v IPSecPhaseConfig)`

SetPhase2 sets Phase2 field to given value.


### GetRight

`func (o *NetworkIpsecBase) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *NetworkIpsecBase) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *NetworkIpsecBase) SetRight(v string)`

SetRight sets Right field to given value.


### GetRightID

`func (o *NetworkIpsecBase) GetRightID() NetworkIpsecBaseAllOfRightID`

GetRightID returns the RightID field if non-nil, zero value otherwise.

### GetRightIDOk

`func (o *NetworkIpsecBase) GetRightIDOk() (*NetworkIpsecBaseAllOfRightID, bool)`

GetRightIDOk returns a tuple with the RightID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightID

`func (o *NetworkIpsecBase) SetRightID(v NetworkIpsecBaseAllOfRightID)`

SetRightID sets RightID field to given value.


### GetPassphrase

`func (o *NetworkIpsecBase) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *NetworkIpsecBase) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *NetworkIpsecBase) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetDpdAction

`func (o *NetworkIpsecBase) GetDpdAction() string`

GetDpdAction returns the DpdAction field if non-nil, zero value otherwise.

### GetDpdActionOk

`func (o *NetworkIpsecBase) GetDpdActionOk() (*string, bool)`

GetDpdActionOk returns a tuple with the DpdAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdAction

`func (o *NetworkIpsecBase) SetDpdAction(v string)`

SetDpdAction sets DpdAction field to given value.


### GetLeftSubnets

`func (o *NetworkIpsecBase) GetLeftSubnets() []string`

GetLeftSubnets returns the LeftSubnets field if non-nil, zero value otherwise.

### GetLeftSubnetsOk

`func (o *NetworkIpsecBase) GetLeftSubnetsOk() (*[]string, bool)`

GetLeftSubnetsOk returns a tuple with the LeftSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftSubnets

`func (o *NetworkIpsecBase) SetLeftSubnets(v []string)`

SetLeftSubnets sets LeftSubnets field to given value.


### GetRightSubnets

`func (o *NetworkIpsecBase) GetRightSubnets() []string`

GetRightSubnets returns the RightSubnets field if non-nil, zero value otherwise.

### GetRightSubnetsOk

`func (o *NetworkIpsecBase) GetRightSubnetsOk() (*[]string, bool)`

GetRightSubnetsOk returns a tuple with the RightSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightSubnets

`func (o *NetworkIpsecBase) SetRightSubnets(v []string)`

SetRightSubnets sets RightSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


