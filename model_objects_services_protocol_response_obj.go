/*
Harmony SASE Public API

The YAML for Harmony SASE Public API.

API version: 2.3.0
*/

// Hand-edited for P81-123406 (BUG-17). The original openapi-generated code
// declared this type as a discriminator-less anyOf wrapper around
// ObjectServiceProtocolICMPResponse + ObjectServiceProtocolTCPUDP (which
// itself was a oneOf over Single/Range/List). The public-api actually emits
// a single flat object per protocol entry — { protocol, valueType, value }
// for tcp/udp and { protocol, protocolOptions } for icmp — so the generated
// nested types could never round-trip the wire payload. We replace the
// wrapper with a flat struct that matches the real shape; standard
// json.Marshal / json.Unmarshal handle it correctly in both directions.
//
// The legacy nested types (ObjectServiceProtocolTCPUDP / List / Range /
// Single / ObjectServiceProtocolICMPResponse) are left in the SDK so any
// out-of-tree callers keep compiling; they are no longer referenced by the
// flat type below.
//
// Added to .swagger-codegen-ignore so a future regen does not clobber this.

package perimeter81sdk

import (
	"encoding/json"
)

// ObjectsServicesProtocolResponseObj is one protocol entry in the response
// returned by GET /v2.3/objects/services. For tcp/udp services, `Value` and
// `ValueType` are populated; for icmp services, `ProtocolOptions` is set
// and the other two are empty.
type ObjectsServicesProtocolResponseObj struct {
	Protocol        string                                    `json:"protocol"`
	ValueType       string                                    `json:"valueType,omitempty"`
	Value           []int32                                   `json:"value,omitempty"`
	ProtocolOptions *ObjectServiceProtocolOptionsICMPresponse `json:"protocolOptions,omitempty"`
}

type NullableObjectsServicesProtocolResponseObj struct {
	value *ObjectsServicesProtocolResponseObj
	isSet bool
}

func (v NullableObjectsServicesProtocolResponseObj) Get() *ObjectsServicesProtocolResponseObj {
	return v.value
}

func (v *NullableObjectsServicesProtocolResponseObj) Set(val *ObjectsServicesProtocolResponseObj) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectsServicesProtocolResponseObj) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectsServicesProtocolResponseObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectsServicesProtocolResponseObj(val *ObjectsServicesProtocolResponseObj) *NullableObjectsServicesProtocolResponseObj {
	return &NullableObjectsServicesProtocolResponseObj{value: val, isSet: true}
}

func (v NullableObjectsServicesProtocolResponseObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectsServicesProtocolResponseObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
