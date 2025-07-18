/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// checks if the CollectionOfPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfPermissions{}

// CollectionOfPermissions struct for CollectionOfPermissions
type CollectionOfPermissions struct {
	Value []Permission `json:"value,omitempty"`
}

// NewCollectionOfPermissions instantiates a new CollectionOfPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfPermissions() *CollectionOfPermissions {
	this := CollectionOfPermissions{}
	return &this
}

// NewCollectionOfPermissionsWithDefaults instantiates a new CollectionOfPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfPermissionsWithDefaults() *CollectionOfPermissions {
	this := CollectionOfPermissions{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfPermissions) GetValue() []Permission {
	if o == nil || IsNil(o.Value) {
		var ret []Permission
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfPermissions) GetValueOk() ([]Permission, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfPermissions) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Permission and assigns it to the Value field.
func (o *CollectionOfPermissions) SetValue(v []Permission) {
	o.Value = v
}

func (o CollectionOfPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCollectionOfPermissions struct {
	value *CollectionOfPermissions
	isSet bool
}

func (v NullableCollectionOfPermissions) Get() *CollectionOfPermissions {
	return v.value
}

func (v *NullableCollectionOfPermissions) Set(val *CollectionOfPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfPermissions(val *CollectionOfPermissions) *NullableCollectionOfPermissions {
	return &NullableCollectionOfPermissions{value: val, isSet: true}
}

func (v NullableCollectionOfPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


