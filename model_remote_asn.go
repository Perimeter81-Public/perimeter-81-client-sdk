/*
Harmony SASE Public API

The YAML for Harmony SASE Public API.

API version: 2.3.0
*/

// Hand-edited. Openapi-generated code declared `type RemoteASN struct {}` —
// an empty struct that could never carry the BGP autonomous-system-number
// value it was supposed to represent. On the wire the API expects (and
// returns) a plain integer, so RemoteASN is redefined as a named int32.
// Callers should write e.g. `RemoteASN(65010)`; the Nullable wrapper is
// kept for backwards compatibility.
//
// Added to .swagger-codegen-ignore so a future regen does not clobber this.

package perimeter81sdk

import (
	"encoding/json"
)

// RemoteASN is a BGP Autonomous System Number that the SDK encodes as a
// JSON integer at the wire.
type RemoteASN int32

type NullableRemoteASN struct {
	value *RemoteASN
	isSet bool
}

func (v NullableRemoteASN) Get() *RemoteASN {
	return v.value
}

func (v *NullableRemoteASN) Set(val *RemoteASN) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteASN) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteASN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteASN(val *RemoteASN) *NullableRemoteASN {
	return &NullableRemoteASN{value: val, isSet: true}
}

func (v NullableRemoteASN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteASN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
