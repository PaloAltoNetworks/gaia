package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ImportReferenceConstraintValue represents the possible values for attribute "constraint".
type ImportReferenceConstraintValue string

const (
	// ImportReferenceConstraintNamespaceUnique represents the value NamespaceUnique.
	ImportReferenceConstraintNamespaceUnique ImportReferenceConstraintValue = "NamespaceUnique"

	// ImportReferenceConstraintUnique represents the value Unique.
	ImportReferenceConstraintUnique ImportReferenceConstraintValue = "Unique"

	// ImportReferenceConstraintUnrestricted represents the value Unrestricted.
	ImportReferenceConstraintUnrestricted ImportReferenceConstraintValue = "Unrestricted"
)

// ImportReferenceIdentity represents the Identity of the object.
var ImportReferenceIdentity = elemental.Identity{
	Name:     "importreference",
	Category: "importreferences",
	Package:  "yuna",
	Private:  false,
}

// ImportReferencesList represents a list of ImportReferences
type ImportReferencesList []*ImportReference

// Identity returns the identity of the objects in the list.
func (o ImportReferencesList) Identity() elemental.Identity {

	return ImportReferenceIdentity
}

// Copy returns a pointer to a copy the ImportReferencesList.
func (o ImportReferencesList) Copy() elemental.Identifiables {

	copy := append(ImportReferencesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ImportReferencesList.
func (o ImportReferencesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ImportReferencesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ImportReference))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ImportReferencesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ImportReferencesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the ImportReferencesList converted to SparseImportReferencesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ImportReferencesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseImportReferencesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseImportReference)
	}

	return out
}

// Version returns the version of the content.
func (o ImportReferencesList) Version() int {

	return 1
}

// ImportReference represents the model of a importreference
type ImportReference struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Contains the claims of the client that performed the import.
	Claims []string `json:"claims" msgpack:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// Define the import constraint. If Unrestricted, import
	// can be deployed multiple times. If Unique, only one import is allowed
	// in the current namespace and its child namespaces. If NamespaceUnique, only
	// one import is allowed in the current namespace.
	Constraint ImportReferenceConstraintValue `json:"constraint" msgpack:"constraint" bson:"constraint" mapstructure:"constraint,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Data to import.
	Data *Export `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Label used for the imported data.
	Label string `json:"label" msgpack:"label" bson:"label" mapstructure:"label,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewImportReference returns a new *ImportReference
func NewImportReference() *ImportReference {

	return &ImportReference{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Claims:         []string{},
		Constraint:     ImportReferenceConstraintUnrestricted,
		Data:           NewExport(),
		Metadata:       []string{},
		MigrationsLog:  map[string]string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *ImportReference) Identity() elemental.Identity {

	return ImportReferenceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ImportReference) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ImportReference) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ImportReference) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesImportReference{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.Claims = o.Claims
	s.Constraint = o.Constraint
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Data = o.Data
	s.Description = o.Description
	s.Label = o.Label
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ImportReference) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesImportReference{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.Claims = s.Claims
	o.Constraint = s.Constraint
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Data = s.Data
	o.Description = s.Description
	o.Label = s.Label
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ImportReference) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ImportReference) BleveType() string {

	return "importreference"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ImportReference) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ImportReference) Doc() string {

	return `Allows you to import and keep a reference.`
}

func (o *ImportReference) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ImportReference) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *ImportReference) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ImportReference) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *ImportReference) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *ImportReference) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *ImportReference) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ImportReference) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ImportReference) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *ImportReference) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *ImportReference) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *ImportReference) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *ImportReference) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *ImportReference) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *ImportReference) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *ImportReference) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *ImportReference) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ImportReference) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ImportReference) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ImportReference) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *ImportReference) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ImportReference) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *ImportReference) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *ImportReference) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *ImportReference) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ImportReference) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ImportReference) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *ImportReference) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ImportReference) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ImportReference) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ImportReference) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ImportReference) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseImportReference{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Claims:               &o.Claims,
			Constraint:           &o.Constraint,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Data:                 o.Data,
			Description:          &o.Description,
			Label:                &o.Label,
			Metadata:             &o.Metadata,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseImportReference{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "claims":
			sp.Claims = &(o.Claims)
		case "constraint":
			sp.Constraint = &(o.Constraint)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "data":
			sp.Data = o.Data
		case "description":
			sp.Description = &(o.Description)
		case "label":
			sp.Label = &(o.Label)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseImportReference to the object.
func (o *ImportReference) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseImportReference)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.Constraint != nil {
		o.Constraint = *so.Constraint
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Data != nil {
		o.Data = so.Data
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Label != nil {
		o.Label = *so.Label
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the ImportReference.
func (o *ImportReference) DeepCopy() *ImportReference {

	if o == nil {
		return nil
	}

	out := &ImportReference{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ImportReference.
func (o *ImportReference) DeepCopyInto(out *ImportReference) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ImportReference: %s", err))
	}

	*out = *target.(*ImportReference)
}

// Validate valides the current information stored into the structure.
func (o *ImportReference) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("constraint", string(o.Constraint), []string{"Unrestricted", "Unique", "NamespaceUnique"}, false); err != nil {
		errors = errors.Append(err)
	}

	if o.Data != nil {
		elemental.ResetDefaultForZeroValues(o.Data)
		if err := o.Data.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*ImportReference) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ImportReferenceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ImportReferenceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ImportReference) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ImportReferenceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ImportReference) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "claims":
		return o.Claims
	case "constraint":
		return o.Constraint
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "data":
		return o.Data
	case "description":
		return o.Description
	case "label":
		return o.Label
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ImportReferenceAttributesMap represents the map of attribute for ImportReference.
var ImportReferenceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Claims": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Contains the claims of the client that performed the import.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Constraint": {
		AllowedChoices: []string{"Unrestricted", "Unique", "NamespaceUnique"},
		ConvertedName:  "Constraint",
		DefaultValue:   ImportReferenceConstraintUnrestricted,
		Description: `Define the import constraint. If Unrestricted, import
can be deployed multiple times. If Unique, only one import is allowed
in the current namespace and its child namespaces. If NamespaceUnique, only
one import is allowed in the current namespace.`,
		Exposed: true,
		Name:    "constraint",
		Stored:  true,
		Type:    "enum",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "export",
		Type:           "ref",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Label": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Label",
		Description:    `Label used for the imported data.`,
		Exposed:        true,
		Name:           "label",
		Stored:         true,
		Type:           "string",
	},
	"Metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
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
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// ImportReferenceLowerCaseAttributesMap represents the map of attribute for ImportReference.
var ImportReferenceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"claims": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Contains the claims of the client that performed the import.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"constraint": {
		AllowedChoices: []string{"Unrestricted", "Unique", "NamespaceUnique"},
		ConvertedName:  "Constraint",
		DefaultValue:   ImportReferenceConstraintUnrestricted,
		Description: `Define the import constraint. If Unrestricted, import
can be deployed multiple times. If Unique, only one import is allowed
in the current namespace and its child namespaces. If NamespaceUnique, only
one import is allowed in the current namespace.`,
		Exposed: true,
		Name:    "constraint",
		Stored:  true,
		Type:    "enum",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "export",
		Type:           "ref",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"label": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Label",
		Description:    `Label used for the imported data.`,
		Exposed:        true,
		Name:           "label",
		Stored:         true,
		Type:           "string",
	},
	"metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
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
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseImportReferencesList represents a list of SparseImportReferences
type SparseImportReferencesList []*SparseImportReference

// Identity returns the identity of the objects in the list.
func (o SparseImportReferencesList) Identity() elemental.Identity {

	return ImportReferenceIdentity
}

// Copy returns a pointer to a copy the SparseImportReferencesList.
func (o SparseImportReferencesList) Copy() elemental.Identifiables {

	copy := append(SparseImportReferencesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseImportReferencesList.
func (o SparseImportReferencesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseImportReferencesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseImportReference))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseImportReferencesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseImportReferencesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseImportReferencesList converted to ImportReferencesList.
func (o SparseImportReferencesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseImportReferencesList) Version() int {

	return 1
}

// SparseImportReference represents the sparse version of a importreference.
type SparseImportReference struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Contains the claims of the client that performed the import.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"claims,omitempty" mapstructure:"claims,omitempty"`

	// Define the import constraint. If Unrestricted, import
	// can be deployed multiple times. If Unique, only one import is allowed
	// in the current namespace and its child namespaces. If NamespaceUnique, only
	// one import is allowed in the current namespace.
	Constraint *ImportReferenceConstraintValue `json:"constraint,omitempty" msgpack:"constraint,omitempty" bson:"constraint,omitempty" mapstructure:"constraint,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Data to import.
	Data *Export `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Label used for the imported data.
	Label *string `json:"label,omitempty" msgpack:"label,omitempty" bson:"label,omitempty" mapstructure:"label,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseImportReference returns a new  SparseImportReference.
func NewSparseImportReference() *SparseImportReference {
	return &SparseImportReference{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseImportReference) Identity() elemental.Identity {

	return ImportReferenceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseImportReference) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseImportReference) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseImportReference) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseImportReference{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.Claims != nil {
		s.Claims = o.Claims
	}
	if o.Constraint != nil {
		s.Constraint = o.Constraint
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Data != nil {
		s.Data = o.Data
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Label != nil {
		s.Label = o.Label
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseImportReference) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseImportReference{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.Claims != nil {
		o.Claims = s.Claims
	}
	if s.Constraint != nil {
		o.Constraint = s.Constraint
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Data != nil {
		o.Data = s.Data
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Label != nil {
		o.Label = s.Label
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseImportReference) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseImportReference) ToPlain() elemental.PlainIdentifiable {

	out := NewImportReference()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.Constraint != nil {
		out.Constraint = *o.Constraint
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Data != nil {
		out.Data = o.Data
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Label != nil {
		out.Label = *o.Label
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseImportReference) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseImportReference) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseImportReference) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseImportReference) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseImportReference) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseImportReference) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseImportReference) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseImportReference) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseImportReference) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseImportReference) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseImportReference) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseImportReference) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseImportReference) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseImportReference) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseImportReference) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseImportReference) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseImportReference) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseImportReference) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseImportReference) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseImportReference) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseImportReference) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseImportReference) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseImportReference) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseImportReference) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseImportReference) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseImportReference) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseImportReference) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseImportReference) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseImportReference) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseImportReference) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseImportReference.
func (o *SparseImportReference) DeepCopy() *SparseImportReference {

	if o == nil {
		return nil
	}

	out := &SparseImportReference{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseImportReference.
func (o *SparseImportReference) DeepCopyInto(out *SparseImportReference) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseImportReference: %s", err))
	}

	*out = *target.(*SparseImportReference)
}

type mongoAttributesImportReference struct {
	ID                   bson.ObjectId                  `bson:"_id,omitempty"`
	Annotations          map[string][]string            `bson:"annotations"`
	AssociatedTags       []string                       `bson:"associatedtags"`
	Claims               []string                       `bson:"claims"`
	Constraint           ImportReferenceConstraintValue `bson:"constraint"`
	CreateIdempotencyKey string                         `bson:"createidempotencykey"`
	CreateTime           time.Time                      `bson:"createtime"`
	Data                 *Export                        `bson:"data"`
	Description          string                         `bson:"description"`
	Label                string                         `bson:"label"`
	Metadata             []string                       `bson:"metadata"`
	MigrationsLog        map[string]string              `bson:"migrationslog,omitempty"`
	Name                 string                         `bson:"name"`
	Namespace            string                         `bson:"namespace"`
	NormalizedTags       []string                       `bson:"normalizedtags"`
	Protected            bool                           `bson:"protected"`
	UpdateIdempotencyKey string                         `bson:"updateidempotencykey"`
	UpdateTime           time.Time                      `bson:"updatetime"`
	ZHash                int                            `bson:"zhash"`
	Zone                 int                            `bson:"zone"`
}
type mongoAttributesSparseImportReference struct {
	ID                   bson.ObjectId                   `bson:"_id,omitempty"`
	Annotations          *map[string][]string            `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                       `bson:"associatedtags,omitempty"`
	Claims               *[]string                       `bson:"claims,omitempty"`
	Constraint           *ImportReferenceConstraintValue `bson:"constraint,omitempty"`
	CreateIdempotencyKey *string                         `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                      `bson:"createtime,omitempty"`
	Data                 *Export                         `bson:"data,omitempty"`
	Description          *string                         `bson:"description,omitempty"`
	Label                *string                         `bson:"label,omitempty"`
	Metadata             *[]string                       `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string              `bson:"migrationslog,omitempty"`
	Name                 *string                         `bson:"name,omitempty"`
	Namespace            *string                         `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                       `bson:"normalizedtags,omitempty"`
	Protected            *bool                           `bson:"protected,omitempty"`
	UpdateIdempotencyKey *string                         `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                      `bson:"updatetime,omitempty"`
	ZHash                *int                            `bson:"zhash,omitempty"`
	Zone                 *int                            `bson:"zone,omitempty"`
}
