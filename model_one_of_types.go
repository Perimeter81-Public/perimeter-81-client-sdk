/*
Harmony SASE Public API

The YAML for Harmony SASE Public API.

API version: 2.3.0
*/

// Hand-maintained oneOf type stubs for fields the generator could not resolve.

package perimeter81sdk

import "encoding/json"

// OneOfFixedHostIdpHost represents a field that is either a FixedHost or an IdpHost.
// It is represented as interface{} since Go does not have native oneOf support.
type OneOfFixedHostIdpHost = interface{}

// NullableOneOfFixedHostIdpHost is a nullable wrapper for OneOfFixedHostIdpHost.
type NullableOneOfFixedHostIdpHost struct {
	value *interface{}
	isSet bool
}

func (v NullableOneOfFixedHostIdpHost) Get() *interface{} {
	return v.value
}

func (v *NullableOneOfFixedHostIdpHost) Set(val *interface{}) {
	v.value = val
	v.isSet = true
}

func (v NullableOneOfFixedHostIdpHost) IsSet() bool {
	return v.isSet
}

func (v *NullableOneOfFixedHostIdpHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneOfFixedHostIdpHost(val *interface{}) *NullableOneOfFixedHostIdpHost {
	return &NullableOneOfFixedHostIdpHost{value: val, isSet: true}
}

func (v NullableOneOfFixedHostIdpHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneOfFixedHostIdpHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// OneOfFixedPortIdpPort represents a field that is either a FixedPort or an IdpPort.
// It is represented as interface{} since Go does not have native oneOf support.
type OneOfFixedPortIdpPort = interface{}

// NullableOneOfFixedPortIdpPort is a nullable wrapper for OneOfFixedPortIdpPort.
type NullableOneOfFixedPortIdpPort struct {
	value *interface{}
	isSet bool
}

func (v NullableOneOfFixedPortIdpPort) Get() *interface{} {
	return v.value
}

func (v *NullableOneOfFixedPortIdpPort) Set(val *interface{}) {
	v.value = val
	v.isSet = true
}

func (v NullableOneOfFixedPortIdpPort) IsSet() bool {
	return v.isSet
}

func (v *NullableOneOfFixedPortIdpPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneOfFixedPortIdpPort(val *interface{}) *NullableOneOfFixedPortIdpPort {
	return &NullableOneOfFixedPortIdpPort{value: val, isSet: true}
}

func (v NullableOneOfFixedPortIdpPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneOfFixedPortIdpPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
