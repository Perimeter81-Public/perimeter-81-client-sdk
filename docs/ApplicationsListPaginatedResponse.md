# ApplicationsListPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ApplicationsListObject**](ApplicationsListObject.md) | List of retrieved applications | [optional] 
**Page** | Pointer to **int32** | Current page number | [optional] 
**TotalPage** | Pointer to **int32** | Number of all pages | [optional] 
**ItemsTotal** | Pointer to **int32** | Number of all items | [optional] 

## Methods

### NewApplicationsListPaginatedResponse

`func NewApplicationsListPaginatedResponse() *ApplicationsListPaginatedResponse`

NewApplicationsListPaginatedResponse instantiates a new ApplicationsListPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsListPaginatedResponseWithDefaults

`func NewApplicationsListPaginatedResponseWithDefaults() *ApplicationsListPaginatedResponse`

NewApplicationsListPaginatedResponseWithDefaults instantiates a new ApplicationsListPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApplicationsListPaginatedResponse) GetData() []ApplicationsListObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApplicationsListPaginatedResponse) GetDataOk() (*[]ApplicationsListObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApplicationsListPaginatedResponse) SetData(v []ApplicationsListObject)`

SetData sets Data field to given value.

### HasData

`func (o *ApplicationsListPaginatedResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPage

`func (o *ApplicationsListPaginatedResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ApplicationsListPaginatedResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ApplicationsListPaginatedResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ApplicationsListPaginatedResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPage

`func (o *ApplicationsListPaginatedResponse) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *ApplicationsListPaginatedResponse) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *ApplicationsListPaginatedResponse) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *ApplicationsListPaginatedResponse) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### GetItemsTotal

`func (o *ApplicationsListPaginatedResponse) GetItemsTotal() int32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ApplicationsListPaginatedResponse) GetItemsTotalOk() (*int32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ApplicationsListPaginatedResponse) SetItemsTotal(v int32)`

SetItemsTotal sets ItemsTotal field to given value.

### HasItemsTotal

`func (o *ApplicationsListPaginatedResponse) HasItemsTotal() bool`

HasItemsTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


