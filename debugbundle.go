package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DebugBundleIdentity represents the Identity of the object.
var DebugBundleIdentity = elemental.Identity{
	Name:     "debugbundle",
	Category: "debugbundles",
	Package:  "guy",
	Private:  false,
}

// DebugBundlesList represents a list of DebugBundles
type DebugBundlesList []*DebugBundle

// Identity returns the identity of the objects in the list.
func (o DebugBundlesList) Identity() elemental.Identity {

	return DebugBundleIdentity
}

// Copy returns a pointer to a copy the DebugBundlesList.
func (o DebugBundlesList) Copy() elemental.Identifiables {

	copy := append(DebugBundlesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DebugBundlesList.
func (o DebugBundlesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DebugBundlesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DebugBundle))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DebugBundlesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DebugBundlesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DebugBundlesList converted to SparseDebugBundlesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DebugBundlesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDebugBundlesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDebugBundle)
	}

	return out
}

// Version returns the version of the content.
func (o DebugBundlesList) Version() int {

	return 1
}

// DebugBundle represents the model of a debugbundle
type DebugBundle struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Can be used to correlate with an EnforcerRefresh.
	DebugID string `json:"debugID,omitempty" msgpack:"debugID,omitempty" bson:"-" mapstructure:"debugID,omitempty"`

	// The ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDebugBundle returns a new *DebugBundle
func NewDebugBundle() *DebugBundle {

	return &DebugBundle{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DebugBundle) Identity() elemental.Identity {

	return DebugBundleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DebugBundle) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DebugBundle) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DebugBundle) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDebugBundle{}

	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DebugBundle) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDebugBundle{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DebugBundle) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DebugBundle) BleveType() string {

	return "debugbundle"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DebugBundle) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DebugBundle) Doc() string {

	return `Represents a file that can be uploaded.`
}

func (o *DebugBundle) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *DebugBundle) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *DebugBundle) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DebugBundle) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDebugBundle{
			ID:         &o.ID,
			DebugID:    &o.DebugID,
			EnforcerID: &o.EnforcerID,
			Namespace:  &o.Namespace,
		}
	}

	sp := &SparseDebugBundle{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "debugID":
			sp.DebugID = &(o.DebugID)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDebugBundle to the object.
func (o *DebugBundle) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDebugBundle)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.DebugID != nil {
		o.DebugID = *so.DebugID
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the DebugBundle.
func (o *DebugBundle) DeepCopy() *DebugBundle {

	if o == nil {
		return nil
	}

	out := &DebugBundle{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DebugBundle.
func (o *DebugBundle) DeepCopyInto(out *DebugBundle) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DebugBundle: %s", err))
	}

	*out = *target.(*DebugBundle)
}

// Validate valides the current information stored into the structure.
func (o *DebugBundle) Validate() error {

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
func (*DebugBundle) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DebugBundleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DebugBundleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DebugBundle) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DebugBundleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DebugBundle) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "debugID":
		return o.DebugID
	case "enforcerID":
		return o.EnforcerID
	case "namespace":
		return o.Namespace
	}

	return nil
}

// DebugBundleAttributesMap represents the map of attribute for DebugBundle.
var DebugBundleAttributesMap = map[string]elemental.AttributeSpecification{
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
	"DebugID": {
		AllowedChoices: []string{},
		ConvertedName:  "DebugID",
		Description:    `Can be used to correlate with an EnforcerRefresh.`,
		Exposed:        true,
		Name:           "debugID",
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `The ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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
}

// DebugBundleLowerCaseAttributesMap represents the map of attribute for DebugBundle.
var DebugBundleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"debugid": {
		AllowedChoices: []string{},
		ConvertedName:  "DebugID",
		Description:    `Can be used to correlate with an EnforcerRefresh.`,
		Exposed:        true,
		Name:           "debugID",
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `The ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		ReadOnly:       true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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
}

// SparseDebugBundlesList represents a list of SparseDebugBundles
type SparseDebugBundlesList []*SparseDebugBundle

// Identity returns the identity of the objects in the list.
func (o SparseDebugBundlesList) Identity() elemental.Identity {

	return DebugBundleIdentity
}

// Copy returns a pointer to a copy the SparseDebugBundlesList.
func (o SparseDebugBundlesList) Copy() elemental.Identifiables {

	copy := append(SparseDebugBundlesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDebugBundlesList.
func (o SparseDebugBundlesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDebugBundlesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDebugBundle))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDebugBundlesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDebugBundlesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDebugBundlesList converted to DebugBundlesList.
func (o SparseDebugBundlesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDebugBundlesList) Version() int {

	return 1
}

// SparseDebugBundle represents the sparse version of a debugbundle.
type SparseDebugBundle struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Can be used to correlate with an EnforcerRefresh.
	DebugID *string `json:"debugID,omitempty" msgpack:"debugID,omitempty" bson:"-" mapstructure:"debugID,omitempty"`

	// The ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDebugBundle returns a new  SparseDebugBundle.
func NewSparseDebugBundle() *SparseDebugBundle {
	return &SparseDebugBundle{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDebugBundle) Identity() elemental.Identity {

	return DebugBundleIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDebugBundle) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDebugBundle) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDebugBundle) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDebugBundle{}

	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDebugBundle) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDebugBundle{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDebugBundle) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDebugBundle) ToPlain() elemental.PlainIdentifiable {

	out := NewDebugBundle()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.DebugID != nil {
		out.DebugID = *o.DebugID
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseDebugBundle) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseDebugBundle) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseDebugBundle.
func (o *SparseDebugBundle) DeepCopy() *SparseDebugBundle {

	if o == nil {
		return nil
	}

	out := &SparseDebugBundle{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDebugBundle.
func (o *SparseDebugBundle) DeepCopyInto(out *SparseDebugBundle) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDebugBundle: %s", err))
	}

	*out = *target.(*SparseDebugBundle)
}

type mongoAttributesDebugBundle struct {
	Namespace string `bson:"namespace"`
}
type mongoAttributesSparseDebugBundle struct {
	Namespace *string `bson:"namespace,omitempty"`
}
