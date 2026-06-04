/*
Harmony SASE Public API

The YAML for Harmony SASE Public API.

API version: 2.3.0
*/

// Hand-edited. Same flat-struct rewrite as
// model_objects_services_protocol_response_obj.go, but for the request side.
// The public-api accepts protocols in the same flat shape it emits.
//
// Added to .swagger-codegen-ignore so a future regen does not clobber this.

package perimeter81sdk

import (
	"encoding/json"
)

// ObjectsServicesProtocolRequestObj is one protocol entry in the request
// body for POST/PUT /v2.3/objects/services. Mirrors the response shape.
type ObjectsServicesProtocolRequestObj struct {
	Protocol        string                                   `json:"protocol"`
	ValueType       string                                   `json:"valueType,omitempty"`
	Value           []int32                                  `json:"value,omitempty"`
	ProtocolOptions *ObjectServiceProtocolOptionsICMPrequest `json:"protocolOptions,omitempty"`
}

type NullableObjectsServicesProtocolRequestObj struct {
	value *ObjectsServicesProtocolRequestObj
	isSet bool
}

func (v NullableObjectsServicesProtocolRequestObj) Get() *ObjectsServicesProtocolRequestObj {
	return v.value
}

func (v *NullableObjectsServicesProtocolRequestObj) Set(val *ObjectsServicesProtocolRequestObj) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectsServicesProtocolRequestObj) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectsServicesProtocolRequestObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectsServicesProtocolRequestObj(val *ObjectsServicesProtocolRequestObj) *NullableObjectsServicesProtocolRequestObj {
	return &NullableObjectsServicesProtocolRequestObj{value: val, isSet: true}
}

func (v NullableObjectsServicesProtocolRequestObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectsServicesProtocolRequestObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
