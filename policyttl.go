package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PolicyTTLIdentity represents the Identity of the object.
var PolicyTTLIdentity = elemental.Identity{
	Name:     "policyttl",
	Category: "policyttls",
	Package:  "squall",
	Private:  true,
}

// PolicyTTLsList represents a list of PolicyTTLs
type PolicyTTLsList []*PolicyTTL

// Identity returns the identity of the objects in the list.
func (o PolicyTTLsList) Identity() elemental.Identity {

	return PolicyTTLIdentity
}

// Copy returns a pointer to a copy the PolicyTTLsList.
func (o PolicyTTLsList) Copy() elemental.Identifiables {

	copy := append(PolicyTTLsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyTTLsList.
func (o PolicyTTLsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PolicyTTLsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyTTL))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyTTLsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyTTLsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PolicyTTLsList converted to SparsePolicyTTLsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PolicyTTLsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePolicyTTLsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePolicyTTL)
	}

	return out
}

// Version returns the version of the content.
func (o PolicyTTLsList) Version() int {

	return 1
}

// PolicyTTL represents the model of a policyttl
type PolicyTTL struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Time when the policy must be deleted.
	ExpirationTime time.Time `json:"-" msgpack:"-" bson:"expirationtime" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPolicyTTL returns a new *PolicyTTL
func NewPolicyTTL() *PolicyTTL {

	return &PolicyTTL{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PolicyTTL) Identity() elemental.Identity {

	return PolicyTTLIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyTTL) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyTTL) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyTTL) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPolicyTTL{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.ExpirationTime = o.ExpirationTime

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PolicyTTL) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPolicyTTL{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.ExpirationTime = s.ExpirationTime

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PolicyTTL) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PolicyTTL) BleveType() string {

	return "policyttl"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyTTL) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PolicyTTL) Doc() string {

	return `This is an unexposed api that defines a helper document allowing to handle
pushes on objects that are deleted by TTL.`
}

func (o *PolicyTTL) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PolicyTTL) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePolicyTTL{
			ID:             &o.ID,
			ExpirationTime: &o.ExpirationTime,
		}
	}

	sp := &SparsePolicyTTL{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "expirationTime":
			sp.ExpirationTime = &(o.ExpirationTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePolicyTTL to the object.
func (o *PolicyTTL) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePolicyTTL)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ExpirationTime != nil {
		o.ExpirationTime = *so.ExpirationTime
	}
}

// DeepCopy returns a deep copy if the PolicyTTL.
func (o *PolicyTTL) DeepCopy() *PolicyTTL {

	if o == nil {
		return nil
	}

	out := &PolicyTTL{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PolicyTTL.
func (o *PolicyTTL) DeepCopyInto(out *PolicyTTL) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PolicyTTL: %s", err))
	}

	*out = *target.(*PolicyTTL)
}

// Validate valides the current information stored into the structure.
func (o *PolicyTTL) Validate() error {

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
func (*PolicyTTL) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyTTLAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyTTLLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyTTL) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyTTLAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PolicyTTL) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "expirationTime":
		return o.ExpirationTime
	}

	return nil
}

// PolicyTTLAttributesMap represents the map of attribute for PolicyTTL.
var PolicyTTLAttributesMap = map[string]elemental.AttributeSpecification{
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
		Stored:         true,
		Type:           "string",
	},
	"ExpirationTime": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `Time when the policy must be deleted.`,
		Name:           "expirationTime",
		Stored:         true,
		Type:           "time",
	},
}

// PolicyTTLLowerCaseAttributesMap represents the map of attribute for PolicyTTL.
var PolicyTTLLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Stored:         true,
		Type:           "string",
	},
	"expirationtime": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `Time when the policy must be deleted.`,
		Name:           "expirationTime",
		Stored:         true,
		Type:           "time",
	},
}

// SparsePolicyTTLsList represents a list of SparsePolicyTTLs
type SparsePolicyTTLsList []*SparsePolicyTTL

// Identity returns the identity of the objects in the list.
func (o SparsePolicyTTLsList) Identity() elemental.Identity {

	return PolicyTTLIdentity
}

// Copy returns a pointer to a copy the SparsePolicyTTLsList.
func (o SparsePolicyTTLsList) Copy() elemental.Identifiables {

	copy := append(SparsePolicyTTLsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePolicyTTLsList.
func (o SparsePolicyTTLsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePolicyTTLsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePolicyTTL))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePolicyTTLsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePolicyTTLsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePolicyTTLsList converted to PolicyTTLsList.
func (o SparsePolicyTTLsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePolicyTTLsList) Version() int {

	return 1
}

// SparsePolicyTTL represents the sparse version of a policyttl.
type SparsePolicyTTL struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Time when the policy must be deleted.
	ExpirationTime *time.Time `json:"-" msgpack:"-" bson:"expirationtime,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePolicyTTL returns a new  SparsePolicyTTL.
func NewSparsePolicyTTL() *SparsePolicyTTL {
	return &SparsePolicyTTL{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePolicyTTL) Identity() elemental.Identity {

	return PolicyTTLIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePolicyTTL) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePolicyTTL) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyTTL) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePolicyTTL{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.ExpirationTime != nil {
		s.ExpirationTime = o.ExpirationTime
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePolicyTTL) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePolicyTTL{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.ExpirationTime != nil {
		o.ExpirationTime = s.ExpirationTime
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePolicyTTL) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePolicyTTL) ToPlain() elemental.PlainIdentifiable {

	out := NewPolicyTTL()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ExpirationTime != nil {
		out.ExpirationTime = *o.ExpirationTime
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePolicyTTL.
func (o *SparsePolicyTTL) DeepCopy() *SparsePolicyTTL {

	if o == nil {
		return nil
	}

	out := &SparsePolicyTTL{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePolicyTTL.
func (o *SparsePolicyTTL) DeepCopyInto(out *SparsePolicyTTL) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePolicyTTL: %s", err))
	}

	*out = *target.(*SparsePolicyTTL)
}

type mongoAttributesPolicyTTL struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	ExpirationTime time.Time     `bson:"expirationtime"`
}
type mongoAttributesSparsePolicyTTL struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	ExpirationTime *time.Time    `bson:"expirationtime,omitempty"`
}
