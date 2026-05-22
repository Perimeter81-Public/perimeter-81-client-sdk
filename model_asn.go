/*
Harmony SASE Public API

The YAML for Harmony SASE Public API.

API version: 2.3.0
*/

// Hand-edited for P81-123406 (BUG-24). Twin of model_remote_asn.go — the
// codegen also produced `type ASN struct {}` (empty struct) for the per-
// tunnel `remoteASN` field on dynamic tunnels. Redefined as a named int32
// for the same reason: the wire shape is a plain integer.
//
// Added to .swagger-codegen-ignore so a future regen does not clobber this.

package perimeter81sdk

import (
	"encoding/json"
)

// ASN — Autonomous System Number for BGP routing on dynamic tunnels. Encoded
// as a JSON integer at the wire.
type ASN int32

type NullableASN struct {
	value *ASN
	isSet bool
}

func (v NullableASN) Get() *ASN {
	return v.value
}

func (v *NullableASN) Set(val *ASN) {
	v.value = val
	v.isSet = true
}

func (v NullableASN) IsSet() bool {
	return v.isSet
}

func (v *NullableASN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASN(val *ASN) *NullableASN {
	return &NullableASN{value: val, isSet: true}
}

func (v NullableASN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableASN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
