package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// RenderedPolicyIdentity represents the Identity of the object
var RenderedPolicyIdentity = elemental.Identity{
	Name:     "renderedpolicy",
	Category: "renderedpolicies",
}

// RenderedPoliciesList represents a list of RenderedPolicies
type RenderedPoliciesList []*RenderedPolicy

// RenderedPolicy represents the model of a renderedpolicy
type RenderedPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-"`

	// EgressPolicies lists all the egress policies attached to ProcessingUnit
	EgressPolicies map[string][]PolicyRule `json:"egressPolicies" cql:"-"`

	// IngressPolicies lists all the ingress policies attached to ProcessingUnit
	IngressPolicies map[string][]PolicyRule `json:"ingressPolicies" cql:"-"`

	// Identifier of the ProcessingUnit
	ProcessingUnitID string `json:"processingUnitID" cql:"-"`
}

// NewRenderedPolicy returns a new *RenderedPolicy
func NewRenderedPolicy() *RenderedPolicy {

	return &RenderedPolicy{}
}

// Identity returns the Identity of the object.
func (o *RenderedPolicy) Identity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RenderedPolicy) Identifier() string {

	return o.ProcessingUnitID
}

func (o *RenderedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RenderedPolicy) SetIdentifier(ID string) {

	o.ProcessingUnitID = ID
}

// Validate valides the current information stored into the structure.
func (o *RenderedPolicy) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o RenderedPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return RenderedPolicyAttributesMap[name]
}

// RenderedPolicyAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
	"EgressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "egressPolicies",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"IngressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "ingressPolicies",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "processingUnitID",
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
}
