package midgardmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// RootIdentity represents the Identity of the object
var RootIdentity = elemental.Identity{
	Name:     "root",
	Category: "root",
}

// Root represents the model of a root
type Root struct {
	// The identifier
	ID string `json:"ID" bson:"_id"`

	// CreatedAt represents the creation time of the object.
	CreatedAt time.Time `json:"createdAt" bson:"createdat"`

	// ParentID is the ID of the parent if any.
	ParentID string `json:"parentID" bson:"parentid"`

	// ParentType is the type of the parent if any.
	ParentType string `json:"parentType" bson:"parenttype"`

	// UpdatedAt represents the last update time of the object.
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedat"`

	Token        string  `json:"APIKey,omitempty"`
	Organization string  `json:"enterprise,omitempty"`
	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRoot returns a new *Root
func NewRoot() *Root {

	return &Root{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Root) Identity() elemental.Identity {

	return RootIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Root) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Root) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Root) Version() float64 {

	return 1.0
}

// Doc returns the documentation for the object
func (o *Root) Doc() string {
	return `Root object of the API`
}

func (o *Root) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreatedAt returns the createdAt of the receiver
func (o *Root) GetCreatedAt() time.Time {
	return o.CreatedAt
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Root) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetParentID returns the parentID of the receiver
func (o *Root) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Root) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Root) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Root) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetUpdatedAt returns the updatedAt of the receiver
func (o *Root) GetUpdatedAt() time.Time {
	return o.UpdatedAt
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Root) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Root) Validate() error {

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

// APIKey returns a the API Key
func (o *Root) APIKey() string {

	return o.Token
}

// SetAPIKey sets a the API Key
func (o *Root) SetAPIKey(key string) {

	o.Token = key
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Root) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return RootAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Root) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RootAttributesMap
}

// RootAttributesMap represents the map of attribute for Root.
var RootAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `The identifier`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt represents the creation time of the object.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent if any.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent if any.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt represents the last update time of the object.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
