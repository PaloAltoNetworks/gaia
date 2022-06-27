package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// HealthCheckStatusValue represents the possible values for attribute "status".
type HealthCheckStatusValue string

const (
	// HealthCheckStatusDegraded represents the value Degraded.
	HealthCheckStatusDegraded HealthCheckStatusValue = "Degraded"

	// HealthCheckStatusOffline represents the value Offline.
	HealthCheckStatusOffline HealthCheckStatusValue = "Offline"

	// HealthCheckStatusOperational represents the value Operational.
	HealthCheckStatusOperational HealthCheckStatusValue = "Operational"
)

// HealthCheckTypeValue represents the possible values for attribute "type".
type HealthCheckTypeValue string

const (
	// HealthCheckTypeCache represents the value Cache.
	HealthCheckTypeCache HealthCheckTypeValue = "Cache"

	// HealthCheckTypeDatabase represents the value Database.
	HealthCheckTypeDatabase HealthCheckTypeValue = "Database"

	// HealthCheckTypeGeneral represents the value General.
	HealthCheckTypeGeneral HealthCheckTypeValue = "General"

	// HealthCheckTypeMessagingSystem represents the value MessagingSystem.
	HealthCheckTypeMessagingSystem HealthCheckTypeValue = "MessagingSystem"

	// HealthCheckTypeService represents the value Service.
	HealthCheckTypeService HealthCheckTypeValue = "Service"

	// HealthCheckTypeTSDB represents the value TSDB.
	HealthCheckTypeTSDB HealthCheckTypeValue = "TSDB"
)

// HealthCheckIdentity represents the Identity of the object.
var HealthCheckIdentity = elemental.Identity{
	Name:     "healthcheck",
	Category: "healthchecks",
	Package:  "ultros",
	Private:  false,
}

// HealthChecksList represents a list of HealthChecks
type HealthChecksList []*HealthCheck

// Identity returns the identity of the objects in the list.
func (o HealthChecksList) Identity() elemental.Identity {

	return HealthCheckIdentity
}

// Copy returns a pointer to a copy the HealthChecksList.
func (o HealthChecksList) Copy() elemental.Identifiables {

	copy := append(HealthChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the HealthChecksList.
func (o HealthChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(HealthChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*HealthCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o HealthChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o HealthChecksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the HealthChecksList converted to SparseHealthChecksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o HealthChecksList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseHealthChecksList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseHealthCheck)
	}

	return out
}

// Version returns the version of the content.
func (o HealthChecksList) Version() int {

	return 1
}

// HealthCheck represents the model of a healthcheck
type HealthCheck struct {
	// A human readable alert list describing the current state of the sub system if
	// available.
	Alerts []string `json:"alerts,omitempty" msgpack:"alerts,omitempty" bson:"-" mapstructure:"alerts,omitempty"`

	// The name of the observed sub system if applicable.
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// The response time of the observed sub system if applicable.
	ResponseTime string `json:"responseTime,omitempty" msgpack:"responseTime,omitempty" bson:"-" mapstructure:"responseTime,omitempty"`

	// The current health of the observed sub system.
	Status HealthCheckStatusValue `json:"status" msgpack:"status" bson:"-" mapstructure:"status,omitempty"`

	// The type of the observed sub system.
	Type HealthCheckTypeValue `json:"type" msgpack:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewHealthCheck returns a new *HealthCheck
func NewHealthCheck() *HealthCheck {

	return &HealthCheck{
		ModelVersion: 1,
		Alerts:       []string{},
	}
}

// Identity returns the Identity of the object.
func (o *HealthCheck) Identity() elemental.Identity {

	return HealthCheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HealthCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HealthCheck) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *HealthCheck) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesHealthCheck{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *HealthCheck) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesHealthCheck{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *HealthCheck) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *HealthCheck) BleveType() string {

	return "healthcheck"
}

// DefaultOrder returns the list of default ordering fields.
func (o *HealthCheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *HealthCheck) Doc() string {

	return `This API allows to retrieve a generic health state of the platform.
A return code different from 200 OK means the platform is not operational.
The health check contains the list of observed sub system.`
}

func (o *HealthCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *HealthCheck) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseHealthCheck{
			Alerts:       &o.Alerts,
			Name:         &o.Name,
			ResponseTime: &o.ResponseTime,
			Status:       &o.Status,
			Type:         &o.Type,
		}
	}

	sp := &SparseHealthCheck{}
	for _, f := range fields {
		switch f {
		case "alerts":
			sp.Alerts = &(o.Alerts)
		case "name":
			sp.Name = &(o.Name)
		case "responseTime":
			sp.ResponseTime = &(o.ResponseTime)
		case "status":
			sp.Status = &(o.Status)
		case "type":
			sp.Type = &(o.Type)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseHealthCheck to the object.
func (o *HealthCheck) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseHealthCheck)
	if so.Alerts != nil {
		o.Alerts = *so.Alerts
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.ResponseTime != nil {
		o.ResponseTime = *so.ResponseTime
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
}

// DeepCopy returns a deep copy if the HealthCheck.
func (o *HealthCheck) DeepCopy() *HealthCheck {

	if o == nil {
		return nil
	}

	out := &HealthCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *HealthCheck.
func (o *HealthCheck) DeepCopyInto(out *HealthCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy HealthCheck: %s", err))
	}

	*out = *target.(*HealthCheck)
}

// Validate valides the current information stored into the structure.
func (o *HealthCheck) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Degraded", "Offline", "Operational"}, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Cache", "Database", "General", "MessagingSystem", "Service", "TSDB"}, true); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*HealthCheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := HealthCheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return HealthCheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*HealthCheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return HealthCheckAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *HealthCheck) ValueForAttribute(name string) interface{} {

	switch name {
	case "alerts":
		return o.Alerts
	case "name":
		return o.Name
	case "responseTime":
		return o.ResponseTime
	case "status":
		return o.Status
	case "type":
		return o.Type
	}

	return nil
}

// HealthCheckAttributesMap represents the map of attribute for HealthCheck.
var HealthCheckAttributesMap = map[string]elemental.AttributeSpecification{
	"Alerts": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Alerts",
		Description: `A human readable alert list describing the current state of the sub system if
available.`,
		Exposed:  true,
		Name:     "alerts",
		ReadOnly: true,
		SubType:  "string",
		Type:     "list",
	},
	"Name": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `The name of the observed sub system if applicable.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"ResponseTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResponseTime",
		Description:    `The response time of the observed sub system if applicable.`,
		Exposed:        true,
		Name:           "responseTime",
		ReadOnly:       true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{"Degraded", "Offline", "Operational"},
		Autogenerated:  true,
		ConvertedName:  "Status",
		Description:    `The current health of the observed sub system.`,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Type": {
		AllowedChoices: []string{"Cache", "Database", "General", "MessagingSystem", "Service", "TSDB"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `The type of the observed sub system.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
}

// HealthCheckLowerCaseAttributesMap represents the map of attribute for HealthCheck.
var HealthCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"alerts": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Alerts",
		Description: `A human readable alert list describing the current state of the sub system if
available.`,
		Exposed:  true,
		Name:     "alerts",
		ReadOnly: true,
		SubType:  "string",
		Type:     "list",
	},
	"name": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `The name of the observed sub system if applicable.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
	"responsetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResponseTime",
		Description:    `The response time of the observed sub system if applicable.`,
		Exposed:        true,
		Name:           "responseTime",
		ReadOnly:       true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{"Degraded", "Offline", "Operational"},
		Autogenerated:  true,
		ConvertedName:  "Status",
		Description:    `The current health of the observed sub system.`,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Type:           "enum",
	},
	"type": {
		AllowedChoices: []string{"Cache", "Database", "General", "MessagingSystem", "Service", "TSDB"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		Description:    `The type of the observed sub system.`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
}

// SparseHealthChecksList represents a list of SparseHealthChecks
type SparseHealthChecksList []*SparseHealthCheck

// Identity returns the identity of the objects in the list.
func (o SparseHealthChecksList) Identity() elemental.Identity {

	return HealthCheckIdentity
}

// Copy returns a pointer to a copy the SparseHealthChecksList.
func (o SparseHealthChecksList) Copy() elemental.Identifiables {

	copy := append(SparseHealthChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseHealthChecksList.
func (o SparseHealthChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseHealthChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseHealthCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseHealthChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseHealthChecksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseHealthChecksList converted to HealthChecksList.
func (o SparseHealthChecksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseHealthChecksList) Version() int {

	return 1
}

// SparseHealthCheck represents the sparse version of a healthcheck.
type SparseHealthCheck struct {
	// A human readable alert list describing the current state of the sub system if
	// available.
	Alerts *[]string `json:"alerts,omitempty" msgpack:"alerts,omitempty" bson:"-" mapstructure:"alerts,omitempty"`

	// The name of the observed sub system if applicable.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// The response time of the observed sub system if applicable.
	ResponseTime *string `json:"responseTime,omitempty" msgpack:"responseTime,omitempty" bson:"-" mapstructure:"responseTime,omitempty"`

	// The current health of the observed sub system.
	Status *HealthCheckStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"-" mapstructure:"status,omitempty"`

	// The type of the observed sub system.
	Type *HealthCheckTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseHealthCheck returns a new  SparseHealthCheck.
func NewSparseHealthCheck() *SparseHealthCheck {
	return &SparseHealthCheck{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseHealthCheck) Identity() elemental.Identity {

	return HealthCheckIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseHealthCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseHealthCheck) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseHealthCheck) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseHealthCheck{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseHealthCheck) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseHealthCheck{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseHealthCheck) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseHealthCheck) ToPlain() elemental.PlainIdentifiable {

	out := NewHealthCheck()
	if o.Alerts != nil {
		out.Alerts = *o.Alerts
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.ResponseTime != nil {
		out.ResponseTime = *o.ResponseTime
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.Type != nil {
		out.Type = *o.Type
	}

	return out
}

// DeepCopy returns a deep copy if the SparseHealthCheck.
func (o *SparseHealthCheck) DeepCopy() *SparseHealthCheck {

	if o == nil {
		return nil
	}

	out := &SparseHealthCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseHealthCheck.
func (o *SparseHealthCheck) DeepCopyInto(out *SparseHealthCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseHealthCheck: %s", err))
	}

	*out = *target.(*SparseHealthCheck)
}

type mongoAttributesHealthCheck struct {
}
type mongoAttributesSparseHealthCheck struct {
}
