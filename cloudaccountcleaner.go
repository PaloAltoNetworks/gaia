// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAccountCleanerIdentity represents the Identity of the object.
var CloudAccountCleanerIdentity = elemental.Identity{
	Name:     "cloudaccountcleaner",
	Category: "cloudaccountcleaner",
	Package:  "yeul",
	Private:  false,
}

// CloudAccountCleanersList represents a list of CloudAccountCleaners
type CloudAccountCleanersList []*CloudAccountCleaner

// Identity returns the identity of the objects in the list.
func (o CloudAccountCleanersList) Identity() elemental.Identity {

	return CloudAccountCleanerIdentity
}

// Copy returns a pointer to a copy the CloudAccountCleanersList.
func (o CloudAccountCleanersList) Copy() elemental.Identifiables {

	out := append(CloudAccountCleanersList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the CloudAccountCleanersList.
func (o CloudAccountCleanersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAccountCleanersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAccountCleaner))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAccountCleanersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAccountCleanersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudAccountCleanersList converted to SparseCloudAccountCleanersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAccountCleanersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAccountCleanersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAccountCleaner)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAccountCleanersList) Version() int {

	return 1
}

// CloudAccountCleaner represents the model of a cloudaccountcleaner
type CloudAccountCleaner struct {
	// The date after which objects must be cleaned.
	Date time.Time `json:"date" msgpack:"date" bson:"-" mapstructure:"date,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAccountCleaner returns a new *CloudAccountCleaner
func NewCloudAccountCleaner() *CloudAccountCleaner {

	return &CloudAccountCleaner{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *CloudAccountCleaner) Identity() elemental.Identity {

	return CloudAccountCleanerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAccountCleaner) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAccountCleaner) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAccountCleaner) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAccountCleaner{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAccountCleaner) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAccountCleaner{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudAccountCleaner) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAccountCleaner) BleveType() string {

	return "cloudaccountcleaner"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAccountCleaner) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudAccountCleaner) Doc() string {

	return `Used for garbage collection of all objects in an account that have not been
updated since the provided time.`
}

func (o *CloudAccountCleaner) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAccountCleaner) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAccountCleaner{
			Date: &o.Date,
		}
	}

	sp := &SparseCloudAccountCleaner{}
	for _, f := range fields {
		switch f {
		case "date":
			sp.Date = &(o.Date)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudAccountCleaner to the object.
func (o *CloudAccountCleaner) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAccountCleaner)
	if so.Date != nil {
		o.Date = *so.Date
	}
}

// DeepCopy returns a deep copy if the CloudAccountCleaner.
func (o *CloudAccountCleaner) DeepCopy() *CloudAccountCleaner {

	if o == nil {
		return nil
	}

	out := &CloudAccountCleaner{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAccountCleaner.
func (o *CloudAccountCleaner) DeepCopyInto(out *CloudAccountCleaner) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAccountCleaner: %s", err))
	}

	*out = *target.(*CloudAccountCleaner)
}

// Validate valides the current information stored into the structure.
func (o *CloudAccountCleaner) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredTime("date", o.Date); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*CloudAccountCleaner) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAccountCleanerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAccountCleanerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAccountCleaner) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAccountCleanerAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAccountCleaner) ValueForAttribute(name string) any {

	switch name {
	case "date":
		return o.Date
	}

	return nil
}

// CloudAccountCleanerAttributesMap represents the map of attribute for CloudAccountCleaner.
var CloudAccountCleanerAttributesMap = map[string]elemental.AttributeSpecification{
	"Date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date after which objects must be cleaned.`,
		Exposed:        true,
		Name:           "date",
		Required:       true,
		Type:           "time",
	},
}

// CloudAccountCleanerLowerCaseAttributesMap represents the map of attribute for CloudAccountCleaner.
var CloudAccountCleanerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date after which objects must be cleaned.`,
		Exposed:        true,
		Name:           "date",
		Required:       true,
		Type:           "time",
	},
}

// SparseCloudAccountCleanersList represents a list of SparseCloudAccountCleaners
type SparseCloudAccountCleanersList []*SparseCloudAccountCleaner

// Identity returns the identity of the objects in the list.
func (o SparseCloudAccountCleanersList) Identity() elemental.Identity {

	return CloudAccountCleanerIdentity
}

// Copy returns a pointer to a copy the SparseCloudAccountCleanersList.
func (o SparseCloudAccountCleanersList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAccountCleanersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAccountCleanersList.
func (o SparseCloudAccountCleanersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAccountCleanersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAccountCleaner))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAccountCleanersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAccountCleanersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudAccountCleanersList converted to CloudAccountCleanersList.
func (o SparseCloudAccountCleanersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAccountCleanersList) Version() int {

	return 1
}

// SparseCloudAccountCleaner represents the sparse version of a cloudaccountcleaner.
type SparseCloudAccountCleaner struct {
	// The date after which objects must be cleaned.
	Date *time.Time `json:"date,omitempty" msgpack:"date,omitempty" bson:"-" mapstructure:"date,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudAccountCleaner returns a new  SparseCloudAccountCleaner.
func NewSparseCloudAccountCleaner() *SparseCloudAccountCleaner {
	return &SparseCloudAccountCleaner{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAccountCleaner) Identity() elemental.Identity {

	return CloudAccountCleanerIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAccountCleaner) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAccountCleaner) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAccountCleaner) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAccountCleaner{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAccountCleaner) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAccountCleaner{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudAccountCleaner) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAccountCleaner) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAccountCleaner()
	if o.Date != nil {
		out.Date = *o.Date
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCloudAccountCleaner.
func (o *SparseCloudAccountCleaner) DeepCopy() *SparseCloudAccountCleaner {

	if o == nil {
		return nil
	}

	out := &SparseCloudAccountCleaner{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAccountCleaner.
func (o *SparseCloudAccountCleaner) DeepCopyInto(out *SparseCloudAccountCleaner) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAccountCleaner: %s", err))
	}

	*out = *target.(*SparseCloudAccountCleaner)
}

type mongoAttributesCloudAccountCleaner struct {
}
type mongoAttributesSparseCloudAccountCleaner struct {
}
