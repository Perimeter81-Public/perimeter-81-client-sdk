# HarmonySaseRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID for the Harmony SASE region. | 
**ObjectId** | Pointer to **string** | Alias of id | [optional] 
**DisplayName** | **string** |  | 
**Name** | **string** | Name of the Harmony SASE region. | 
**CountryCode** | **string** | ISO 3166-1 alpha-2 country code for the region. | 
**ContinentCode** | **string** | ISO 3166-1 alpha-2 continent code for the region. | 

## Methods

### NewHarmonySaseRegion

`func NewHarmonySaseRegion(id string, displayName string, name string, countryCode string, continentCode string, ) *HarmonySaseRegion`

NewHarmonySaseRegion instantiates a new HarmonySaseRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarmonySaseRegionWithDefaults

`func NewHarmonySaseRegionWithDefaults() *HarmonySaseRegion`

NewHarmonySaseRegionWithDefaults instantiates a new HarmonySaseRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HarmonySaseRegion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HarmonySaseRegion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HarmonySaseRegion) SetId(v string)`

SetId sets Id field to given value.


### GetObjectId

`func (o *HarmonySaseRegion) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *HarmonySaseRegion) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *HarmonySaseRegion) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *HarmonySaseRegion) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetDisplayName

`func (o *HarmonySaseRegion) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HarmonySaseRegion) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HarmonySaseRegion) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *HarmonySaseRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HarmonySaseRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HarmonySaseRegion) SetName(v string)`

SetName sets Name field to given value.


### GetCountryCode

`func (o *HarmonySaseRegion) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *HarmonySaseRegion) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *HarmonySaseRegion) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetContinentCode

`func (o *HarmonySaseRegion) GetContinentCode() string`

GetContinentCode returns the ContinentCode field if non-nil, zero value otherwise.

### GetContinentCodeOk

`func (o *HarmonySaseRegion) GetContinentCodeOk() (*string, bool)`

GetContinentCodeOk returns a tuple with the ContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentCode

`func (o *HarmonySaseRegion) SetContinentCode(v string)`

SetContinentCode sets ContinentCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


