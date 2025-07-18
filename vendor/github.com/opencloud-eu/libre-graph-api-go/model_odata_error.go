/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OdataError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OdataError{}

// OdataError struct for OdataError
type OdataError struct {
	Error OdataErrorMain `json:"error"`
}

type _OdataError OdataError

// NewOdataError instantiates a new OdataError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOdataError(error_ OdataErrorMain) *OdataError {
	this := OdataError{}
	this.Error = error_
	return &this
}

// NewOdataErrorWithDefaults instantiates a new OdataError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOdataErrorWithDefaults() *OdataError {
	this := OdataError{}
	return &this
}

// GetError returns the Error field value
func (o *OdataError) GetError() OdataErrorMain {
	if o == nil {
		var ret OdataErrorMain
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *OdataError) GetErrorOk() (*OdataErrorMain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *OdataError) SetError(v OdataErrorMain) {
	o.Error = v
}

func (o OdataError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OdataError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *OdataError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOdataError := _OdataError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOdataError)

	if err != nil {
		return err
	}

	*o = OdataError(varOdataError)

	return err
}

type NullableOdataError struct {
	value *OdataError
	isSet bool
}

func (v NullableOdataError) Get() *OdataError {
	return v.value
}

func (v *NullableOdataError) Set(val *OdataError) {
	v.value = val
	v.isSet = true
}

func (v NullableOdataError) IsSet() bool {
	return v.isSet
}

func (v *NullableOdataError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdataError(val *OdataError) *NullableOdataError {
	return &NullableOdataError{value: val, isSet: true}
}

func (v NullableOdataError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdataError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


