package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DiscoveryModeIdentity represents the Identity of the object.
var DiscoveryModeIdentity = elemental.Identity{
	Name:     "discoverymode",
	Category: "discoverymode",
	Package:  "yuna",
	Private:  false,
}

// DiscoveryModesList represents a list of DiscoveryModes
type DiscoveryModesList []*DiscoveryMode

// Identity returns the identity of the objects in the list.
func (o DiscoveryModesList) Identity() elemental.Identity {

	return DiscoveryModeIdentity
}

// Copy returns a pointer to a copy the DiscoveryModesList.
func (o DiscoveryModesList) Copy() elemental.Identifiables {

	copy := append(DiscoveryModesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DiscoveryModesList.
func (o DiscoveryModesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DiscoveryModesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DiscoveryMode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DiscoveryModesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DiscoveryModesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DiscoveryModesList converted to SparseDiscoveryModesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DiscoveryModesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDiscoveryModesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDiscoveryMode)
	}

	return out
}

// Version returns the version of the content.
func (o DiscoveryModesList) Version() int {

	return 1
}

// DiscoveryMode represents the model of a discoverymode
type DiscoveryMode struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDiscoveryMode returns a new *DiscoveryMode
func NewDiscoveryMode() *DiscoveryMode {

	return &DiscoveryMode{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DiscoveryMode) Identity() elemental.Identity {

	return DiscoveryModeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DiscoveryMode) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DiscoveryMode) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DiscoveryMode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDiscoveryMode{}

	s.Namespace = o.Namespace
	s.Propagate = o.Propagate

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DiscoveryMode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDiscoveryMode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Namespace = s.Namespace
	o.Propagate = s.Propagate

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DiscoveryMode) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DiscoveryMode) BleveType() string {

	return "discoverymode"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DiscoveryMode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DiscoveryMode) Doc() string {

	return `(Deprecated) When discovery mode is enabled, all flows are accepted. Flows which
do not match an existing network policy will be represented by a dotted line in
your Platform view.`
}

func (o *DiscoveryMode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *DiscoveryMode) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *DiscoveryMode) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetPropagate returns the Propagate of the receiver.
func (o *DiscoveryMode) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *DiscoveryMode) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DiscoveryMode) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDiscoveryMode{
			ID:        &o.ID,
			Namespace: &o.Namespace,
			Propagate: &o.Propagate,
		}
	}

	sp := &SparseDiscoveryMode{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDiscoveryMode to the object.
func (o *DiscoveryMode) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDiscoveryMode)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
}

// DeepCopy returns a deep copy if the DiscoveryMode.
func (o *DiscoveryMode) DeepCopy() *DiscoveryMode {

	if o == nil {
		return nil
	}

	out := &DiscoveryMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DiscoveryMode.
func (o *DiscoveryMode) DeepCopyInto(out *DiscoveryMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DiscoveryMode: %s", err))
	}

	*out = *target.(*DiscoveryMode)
}

// Validate valides the current information stored into the structure.
func (o *DiscoveryMode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*DiscoveryMode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DiscoveryModeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DiscoveryModeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DiscoveryMode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DiscoveryModeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DiscoveryMode) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "namespace":
		return o.Namespace
	case "propagate":
		return o.Propagate
	}

	return nil
}

// DiscoveryModeAttributesMap represents the map of attribute for DiscoveryMode.
var DiscoveryModeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
}

// DiscoveryModeLowerCaseAttributesMap represents the map of attribute for DiscoveryMode.
var DiscoveryModeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
}

// SparseDiscoveryModesList represents a list of SparseDiscoveryModes
type SparseDiscoveryModesList []*SparseDiscoveryMode

// Identity returns the identity of the objects in the list.
func (o SparseDiscoveryModesList) Identity() elemental.Identity {

	return DiscoveryModeIdentity
}

// Copy returns a pointer to a copy the SparseDiscoveryModesList.
func (o SparseDiscoveryModesList) Copy() elemental.Identifiables {

	copy := append(SparseDiscoveryModesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDiscoveryModesList.
func (o SparseDiscoveryModesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDiscoveryModesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDiscoveryMode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDiscoveryModesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDiscoveryModesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDiscoveryModesList converted to DiscoveryModesList.
func (o SparseDiscoveryModesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDiscoveryModesList) Version() int {

	return 1
}

// SparseDiscoveryMode represents the sparse version of a discoverymode.
type SparseDiscoveryMode struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDiscoveryMode returns a new  SparseDiscoveryMode.
func NewSparseDiscoveryMode() *SparseDiscoveryMode {
	return &SparseDiscoveryMode{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDiscoveryMode) Identity() elemental.Identity {

	return DiscoveryModeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDiscoveryMode) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDiscoveryMode) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDiscoveryMode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDiscoveryMode{}

	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDiscoveryMode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDiscoveryMode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDiscoveryMode) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDiscoveryMode) ToPlain() elemental.PlainIdentifiable {

	out := NewDiscoveryMode()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseDiscoveryMode) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseDiscoveryMode) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseDiscoveryMode) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseDiscoveryMode) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// DeepCopy returns a deep copy if the SparseDiscoveryMode.
func (o *SparseDiscoveryMode) DeepCopy() *SparseDiscoveryMode {

	if o == nil {
		return nil
	}

	out := &SparseDiscoveryMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDiscoveryMode.
func (o *SparseDiscoveryMode) DeepCopyInto(out *SparseDiscoveryMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDiscoveryMode: %s", err))
	}

	*out = *target.(*SparseDiscoveryMode)
}

type mongoAttributesDiscoveryMode struct {
	Namespace string `bson:"namespace"`
	Propagate bool   `bson:"propagate"`
}
type mongoAttributesSparseDiscoveryMode struct {
	Namespace *string `bson:"namespace,omitempty"`
	Propagate *bool   `bson:"propagate,omitempty"`
}
