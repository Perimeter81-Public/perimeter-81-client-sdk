# ObjectsAddressesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** |  | 
**TotalPage** | **float32** |  | 
**ItemsTotal** | **float32** |  | 
**Data** | [**[]ObjectsAddressObj**](ObjectsAddressObj.md) |  | 

## Methods

### NewObjectsAddressesResponse

`func NewObjectsAddressesResponse(page float32, totalPage float32, itemsTotal float32, data []ObjectsAddressObj, ) *ObjectsAddressesResponse`

NewObjectsAddressesResponse instantiates a new ObjectsAddressesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsAddressesResponseWithDefaults

`func NewObjectsAddressesResponseWithDefaults() *ObjectsAddressesResponse`

NewObjectsAddressesResponseWithDefaults instantiates a new ObjectsAddressesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ObjectsAddressesResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ObjectsAddressesResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ObjectsAddressesResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetTotalPage

`func (o *ObjectsAddressesResponse) GetTotalPage() float32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *ObjectsAddressesResponse) GetTotalPageOk() (*float32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *ObjectsAddressesResponse) SetTotalPage(v float32)`

SetTotalPage sets TotalPage field to given value.


### GetItemsTotal

`func (o *ObjectsAddressesResponse) GetItemsTotal() float32`

GetItemsTotal returns the ItemsTotal field if non-nil, zero value otherwise.

### GetItemsTotalOk

`func (o *ObjectsAddressesResponse) GetItemsTotalOk() (*float32, bool)`

GetItemsTotalOk returns a tuple with the ItemsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotal

`func (o *ObjectsAddressesResponse) SetItemsTotal(v float32)`

SetItemsTotal sets ItemsTotal field to given value.


### GetData

`func (o *ObjectsAddressesResponse) GetData() []ObjectsAddressObj`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectsAddressesResponse) GetDataOk() (*[]ObjectsAddressObj, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectsAddressesResponse) SetData(v []ObjectsAddressObj)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


