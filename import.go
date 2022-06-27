package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ImportModeValue represents the possible values for attribute "mode".
type ImportModeValue string

const (
	// ImportModeImport represents the value Import.
	ImportModeImport ImportModeValue = "Import"

	// ImportModeRemove represents the value Remove.
	ImportModeRemove ImportModeValue = "Remove"

	// ImportModeReplacePartial represents the value ReplacePartial.
	ImportModeReplacePartial ImportModeValue = "ReplacePartial"
)

// ImportIdentity represents the Identity of the object.
var ImportIdentity = elemental.Identity{
	Name:     "import",
	Category: "import",
	Package:  "squall",
	Private:  false,
}

// ImportsList represents a list of Imports
type ImportsList []*Import

// Identity returns the identity of the objects in the list.
func (o ImportsList) Identity() elemental.Identity {

	return ImportIdentity
}

// Copy returns a pointer to a copy the ImportsList.
func (o ImportsList) Copy() elemental.Identifiables {

	copy := append(ImportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ImportsList.
func (o ImportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ImportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Import))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ImportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ImportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ImportsList converted to SparseImportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ImportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseImportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseImport)
	}

	return out
}

// Version returns the version of the content.
func (o ImportsList) Version() int {

	return 1
}

// Import represents the model of a import
type Import struct {
	// Data to import.
	Data *Export `json:"data" msgpack:"data" bson:"-" mapstructure:"data,omitempty"`

	// How to import the data: `ReplacePartial`, `Import` (default), or `Remove`.
	// `ReplacePartial`
	// is deprecated. Use `Import` instead. While you can use `ReplacePartial` it will
	// be interpreted
	// as `Import`.
	Mode ImportModeValue `json:"mode" msgpack:"mode" bson:"-" mapstructure:"mode,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewImport returns a new *Import
func NewImport() *Import {

	return &Import{
		ModelVersion: 1,
		Data:         NewExport(),
		Mode:         ImportModeImport,
	}
}

// Identity returns the Identity of the object.
func (o *Import) Identity() elemental.Identity {

	return ImportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Import) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Import) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Import) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesImport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Import) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesImport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Import) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Import) BleveType() string {

	return "import"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Import) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Import) Doc() string {

	return `Imports an export of policies and related objects into the namespace.`
}

func (o *Import) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Import) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseImport{
			Data: o.Data,
			Mode: &o.Mode,
		}
	}

	sp := &SparseImport{}
	for _, f := range fields {
		switch f {
		case "data":
			sp.Data = o.Data
		case "mode":
			sp.Mode = &(o.Mode)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseImport to the object.
func (o *Import) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseImport)
	if so.Data != nil {
		o.Data = so.Data
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
}

// DeepCopy returns a deep copy if the Import.
func (o *Import) DeepCopy() *Import {

	if o == nil {
		return nil
	}

	out := &Import{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Import.
func (o *Import) DeepCopyInto(out *Import) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Import: %s", err))
	}

	*out = *target.(*Import)
}

// Validate valides the current information stored into the structure.
func (o *Import) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Data != nil {
		elemental.ResetDefaultForZeroValues(o.Data)
		if err := o.Data.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"ReplacePartial", "Import", "Remove"}, false); err != nil {
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
func (*Import) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ImportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ImportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Import) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ImportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Import) ValueForAttribute(name string) interface{} {

	switch name {
	case "data":
		return o.Data
	case "mode":
		return o.Mode
	}

	return nil
}

// ImportAttributesMap represents the map of attribute for Import.
var ImportAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		SubType:        "export",
		Type:           "ref",
	},
	"Mode": {
		AllowedChoices: []string{"ReplacePartial", "Import", "Remove"},
		ConvertedName:  "Mode",
		DefaultValue:   ImportModeImport,
		Description: `How to import the data: ` + "`" + `ReplacePartial` + "`" + `, ` + "`" + `Import` + "`" + ` (default), or ` + "`" + `Remove` + "`" + `.
` + "`" + `ReplacePartial` + "`" + `
is deprecated. Use ` + "`" + `Import` + "`" + ` instead. While you can use ` + "`" + `ReplacePartial` + "`" + ` it will
be interpreted
as ` + "`" + `Import` + "`" + `.`,
		Exposed: true,
		Name:    "mode",
		Type:    "enum",
	},
}

// ImportLowerCaseAttributesMap represents the map of attribute for Import.
var ImportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		SubType:        "export",
		Type:           "ref",
	},
	"mode": {
		AllowedChoices: []string{"ReplacePartial", "Import", "Remove"},
		ConvertedName:  "Mode",
		DefaultValue:   ImportModeImport,
		Description: `How to import the data: ` + "`" + `ReplacePartial` + "`" + `, ` + "`" + `Import` + "`" + ` (default), or ` + "`" + `Remove` + "`" + `.
` + "`" + `ReplacePartial` + "`" + `
is deprecated. Use ` + "`" + `Import` + "`" + ` instead. While you can use ` + "`" + `ReplacePartial` + "`" + ` it will
be interpreted
as ` + "`" + `Import` + "`" + `.`,
		Exposed: true,
		Name:    "mode",
		Type:    "enum",
	},
}

// SparseImportsList represents a list of SparseImports
type SparseImportsList []*SparseImport

// Identity returns the identity of the objects in the list.
func (o SparseImportsList) Identity() elemental.Identity {

	return ImportIdentity
}

// Copy returns a pointer to a copy the SparseImportsList.
func (o SparseImportsList) Copy() elemental.Identifiables {

	copy := append(SparseImportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseImportsList.
func (o SparseImportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseImportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseImport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseImportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseImportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseImportsList converted to ImportsList.
func (o SparseImportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseImportsList) Version() int {

	return 1
}

// SparseImport represents the sparse version of a import.
type SparseImport struct {
	// Data to import.
	Data *Export `json:"data,omitempty" msgpack:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	// How to import the data: `ReplacePartial`, `Import` (default), or `Remove`.
	// `ReplacePartial`
	// is deprecated. Use `Import` instead. While you can use `ReplacePartial` it will
	// be interpreted
	// as `Import`.
	Mode *ImportModeValue `json:"mode,omitempty" msgpack:"mode,omitempty" bson:"-" mapstructure:"mode,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseImport returns a new  SparseImport.
func NewSparseImport() *SparseImport {
	return &SparseImport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseImport) Identity() elemental.Identity {

	return ImportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseImport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseImport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseImport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseImport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseImport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseImport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseImport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseImport) ToPlain() elemental.PlainIdentifiable {

	out := NewImport()
	if o.Data != nil {
		out.Data = o.Data
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}

	return out
}

// DeepCopy returns a deep copy if the SparseImport.
func (o *SparseImport) DeepCopy() *SparseImport {

	if o == nil {
		return nil
	}

	out := &SparseImport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseImport.
func (o *SparseImport) DeepCopyInto(out *SparseImport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseImport: %s", err))
	}

	*out = *target.(*SparseImport)
}

type mongoAttributesImport struct {
}
type mongoAttributesSparseImport struct {
}
