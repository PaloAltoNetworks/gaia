package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// MapEdgeIdentity represents the Identity of the object
var MapEdgeIdentity = elemental.Identity{
	Name:     "mapedge",
	Category: "mapedges",
}

// MapEdgesList represents a list of MapEdges
type MapEdgesList []*MapEdge

// MapEdge represents the model of a mapedge
type MapEdge struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-" bson:"-"`

	// AcceptedFlows tells how many accepted flows are represented by the edge
	AcceptedFlows int `json:"acceptedFlows" cql:"acceptedflows,omitempty" bson:"acceptedflows"`

	// ID of the destination resource
	DestinationID string `json:"destinationID" cql:"-" bson:"-"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// RejectedFlows tells how many flows has been rejected.
	RejectedFlows int `json:"rejectedFlows" cql:"rejectedflows,omitempty" bson:"rejectedflows"`

	// ID of the source resource
	SourceID string `json:"sourceID" cql:"-" bson:"-"`
}

// NewMapEdge returns a new *MapEdge
func NewMapEdge() *MapEdge {

	return &MapEdge{}
}

// Identity returns the Identity of the object.
func (o *MapEdge) Identity() elemental.Identity {

	return MapEdgeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MapEdge) Identifier() string {

	return o.ID
}

func (o *MapEdge) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MapEdge) SetIdentifier(ID string) {

	o.ID = ID
}

// GetName returns the name of the receiver
func (o *MapEdge) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *MapEdge) SetName(name string) {
	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *MapEdge) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (MapEdge) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return MapEdgeAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (MapEdge) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MapEdgeAttributesMap
}

// MapEdgeAttributesMap represents the map of attribute for MapEdge.
var MapEdgeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"AcceptedFlows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "acceptedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"DestinationID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "destinationID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"RejectedFlows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "rejectedFlows",
		Stored:         true,
		Type:           "integer",
	},
	"SourceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "sourceID",
		ReadOnly:       true,
		Type:           "string",
	},
}
