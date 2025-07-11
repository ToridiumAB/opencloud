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

// checks if the CollectionOfUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfUsers{}

// CollectionOfUsers struct for CollectionOfUsers
type CollectionOfUsers struct {
	Value []User `json:"value,omitempty"`
}

// NewCollectionOfUsers instantiates a new CollectionOfUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfUsers() *CollectionOfUsers {
	this := CollectionOfUsers{}
	return &this
}

// NewCollectionOfUsersWithDefaults instantiates a new CollectionOfUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfUsersWithDefaults() *CollectionOfUsers {
	this := CollectionOfUsers{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfUsers) GetValue() []User {
	if o == nil || IsNil(o.Value) {
		var ret []User
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUsers) GetValueOk() ([]User, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfUsers) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []User and assigns it to the Value field.
func (o *CollectionOfUsers) SetValue(v []User) {
	o.Value = v
}

func (o CollectionOfUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCollectionOfUsers struct {
	value *CollectionOfUsers
	isSet bool
}

func (v NullableCollectionOfUsers) Get() *CollectionOfUsers {
	return v.value
}

func (v *NullableCollectionOfUsers) Set(val *CollectionOfUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfUsers(val *CollectionOfUsers) *NullableCollectionOfUsers {
	return &NullableCollectionOfUsers{value: val, isSet: true}
}

func (v NullableCollectionOfUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


