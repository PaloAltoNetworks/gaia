package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// ComputedPolicyIdentity represents the Identity of the object
var ComputedPolicyIdentity = elemental.Identity{
	Name:     "computedpolicy",
	Category: "computedpolicies",
}

// ComputedPoliciesList represents a list of ComputedPolicies
type ComputedPoliciesList []*ComputedPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o ComputedPoliciesList) ContentIdentity() elemental.Identity {
	return ComputedPolicyIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o ComputedPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// ComputedPolicy represents the model of a computedpolicy
type ComputedPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// Array of netowrk access policies computed
	NetworkAccessPolicies []*NetworkAccessPolicy `json:"networkAccessPolicies" bson:"networkaccesspolicies"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewComputedPolicy returns a new *ComputedPolicy
func NewComputedPolicy() *ComputedPolicy {

	return &ComputedPolicy{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *ComputedPolicy) Identity() elemental.Identity {

	return ComputedPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ComputedPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ComputedPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *ComputedPolicy) Version() float64 {

	return 1.0
}

// Doc returns the documentation for the object
func (o *ComputedPolicy) Doc() string {

	return nodocString

}

func (o *ComputedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *ComputedPolicy) Validate() error {

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
func (*ComputedPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ComputedPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ComputedPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ComputedPolicyAttributesMap
}

// ComputedPolicyAttributesMap represents the map of attribute for ComputedPolicy.
var ComputedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
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
	"NetworkAccessPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Array of netowrk access policies computed`,
		Exposed:        true,
		Name:           "networkAccessPolicies",
		Orderable:      true,
		Stored:         true,
		SubType:        "network_access_policies_list",
		Type:           "external",
	},
}
