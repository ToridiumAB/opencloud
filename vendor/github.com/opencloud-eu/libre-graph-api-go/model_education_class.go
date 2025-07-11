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

// checks if the EducationClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EducationClass{}

// EducationClass And extension of group representing a class or course
type EducationClass struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// An optional description for the group. Returned by default.
	Description *string `json:"description,omitempty"`
	// The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $search and $orderBy.
	DisplayName string `json:"displayName"`
	// Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand.
	Members []User `json:"members,omitempty"`
	// A list of member references to the members to be added. Up to 20 members can be added with a single request
	MembersodataBind []string `json:"members@odata.bind,omitempty"`
	// Classification of the group, i.e. \"class\" or \"course\"
	Classification string `json:"classification"`
	// An external unique ID for the class
	ExternalId *string `json:"externalId,omitempty"`
}

type _EducationClass EducationClass

// NewEducationClass instantiates a new EducationClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEducationClass(displayName string, classification string) *EducationClass {
	this := EducationClass{}
	this.DisplayName = displayName
	this.Classification = classification
	return &this
}

// NewEducationClassWithDefaults instantiates a new EducationClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEducationClassWithDefaults() *EducationClass {
	this := EducationClass{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EducationClass) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationClass) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EducationClass) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EducationClass) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EducationClass) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationClass) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EducationClass) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EducationClass) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *EducationClass) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EducationClass) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EducationClass) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *EducationClass) GetMembers() []User {
	if o == nil || IsNil(o.Members) {
		var ret []User
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationClass) GetMembersOk() ([]User, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *EducationClass) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []User and assigns it to the Members field.
func (o *EducationClass) SetMembers(v []User) {
	o.Members = v
}

// GetMembersodataBind returns the MembersodataBind field value if set, zero value otherwise.
func (o *EducationClass) GetMembersodataBind() []string {
	if o == nil || IsNil(o.MembersodataBind) {
		var ret []string
		return ret
	}
	return o.MembersodataBind
}

// GetMembersodataBindOk returns a tuple with the MembersodataBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationClass) GetMembersodataBindOk() ([]string, bool) {
	if o == nil || IsNil(o.MembersodataBind) {
		return nil, false
	}
	return o.MembersodataBind, true
}

// HasMembersodataBind returns a boolean if a field has been set.
func (o *EducationClass) HasMembersodataBind() bool {
	if o != nil && !IsNil(o.MembersodataBind) {
		return true
	}

	return false
}

// SetMembersodataBind gets a reference to the given []string and assigns it to the MembersodataBind field.
func (o *EducationClass) SetMembersodataBind(v []string) {
	o.MembersodataBind = v
}

// GetClassification returns the Classification field value
func (o *EducationClass) GetClassification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value
// and a boolean to check if the value has been set.
func (o *EducationClass) GetClassificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classification, true
}

// SetClassification sets field value
func (o *EducationClass) SetClassification(v string) {
	o.Classification = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EducationClass) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EducationClass) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EducationClass) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EducationClass) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o EducationClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EducationClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.MembersodataBind) {
		toSerialize["members@odata.bind"] = o.MembersodataBind
	}
	toSerialize["classification"] = o.Classification
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	return toSerialize, nil
}

func (o *EducationClass) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"classification",
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

	varEducationClass := _EducationClass{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEducationClass)

	if err != nil {
		return err
	}

	*o = EducationClass(varEducationClass)

	return err
}

type NullableEducationClass struct {
	value *EducationClass
	isSet bool
}

func (v NullableEducationClass) Get() *EducationClass {
	return v.value
}

func (v *NullableEducationClass) Set(val *EducationClass) {
	v.value = val
	v.isSet = true
}

func (v NullableEducationClass) IsSet() bool {
	return v.isSet
}

func (v *NullableEducationClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEducationClass(val *EducationClass) *NullableEducationClass {
	return &NullableEducationClass{value: val, isSet: true}
}

func (v NullableEducationClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEducationClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


