package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudLoadBalancerData represents the model of a cloudloadbalancerdata
type CloudLoadBalancerData struct {
	// ID of associated objects with this load balancer.
	AttachedEntities []string `json:"attachedEntities" msgpack:"attachedEntities" bson:"attachedentities" mapstructure:"attachedEntities,omitempty"`

	// Mapping of a listener to its associated target group ID list.
	Listenertargetmapping map[string][]string `json:"listenertargetmapping" msgpack:"listenertargetmapping" bson:"listenertargetmapping" mapstructure:"listenertargetmapping,omitempty"`

	// The name of the load balancer.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The scheme tells whether the load balancer is internet facing or internal.
	Scheme string `json:"scheme" msgpack:"scheme" bson:"scheme" mapstructure:"scheme,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudLoadBalancerData returns a new *CloudLoadBalancerData
func NewCloudLoadBalancerData() *CloudLoadBalancerData {

	return &CloudLoadBalancerData{
		ModelVersion:          1,
		AttachedEntities:      []string{},
		Listenertargetmapping: map[string][]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudLoadBalancerData{}

	s.AttachedEntities = o.AttachedEntities
	s.Listenertargetmapping = o.Listenertargetmapping
	s.Name = o.Name
	s.Scheme = o.Scheme

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudLoadBalancerData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AttachedEntities = s.AttachedEntities
	o.Listenertargetmapping = s.Listenertargetmapping
	o.Name = s.Name
	o.Scheme = s.Scheme

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudLoadBalancerData) BleveType() string {

	return "cloudloadbalancerdata"
}

// DeepCopy returns a deep copy if the CloudLoadBalancerData.
func (o *CloudLoadBalancerData) DeepCopy() *CloudLoadBalancerData {

	if o == nil {
		return nil
	}

	out := &CloudLoadBalancerData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudLoadBalancerData.
func (o *CloudLoadBalancerData) DeepCopyInto(out *CloudLoadBalancerData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudLoadBalancerData: %s", err))
	}

	*out = *target.(*CloudLoadBalancerData)
}

// Validate valides the current information stored into the structure.
func (o *CloudLoadBalancerData) Validate() error {

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
func (*CloudLoadBalancerData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudLoadBalancerDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudLoadBalancerDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudLoadBalancerData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudLoadBalancerDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudLoadBalancerData) ValueForAttribute(name string) interface{} {

	switch name {
	case "attachedEntities":
		return o.AttachedEntities
	case "listenertargetmapping":
		return o.Listenertargetmapping
	case "name":
		return o.Name
	case "scheme":
		return o.Scheme
	}

	return nil
}

// CloudLoadBalancerDataAttributesMap represents the map of attribute for CloudLoadBalancerData.
var CloudLoadBalancerDataAttributesMap = map[string]elemental.AttributeSpecification{
	"AttachedEntities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `ID of associated objects with this load balancer.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Listenertargetmapping": {
		AllowedChoices: []string{},
		BSONFieldName:  "listenertargetmapping",
		ConvertedName:  "Listenertargetmapping",
		Description:    `Mapping of a listener to its associated target group ID list.`,
		Exposed:        true,
		Name:           "listenertargetmapping",
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `The name of the load balancer.`,
		Exposed:        true,
		Name:           "name",
		Stored:         true,
		Type:           "string",
	},
	"Scheme": {
		AllowedChoices: []string{},
		BSONFieldName:  "scheme",
		ConvertedName:  "Scheme",
		Description:    `The scheme tells whether the load balancer is internet facing or internal.`,
		Exposed:        true,
		Name:           "scheme",
		Stored:         true,
		Type:           "string",
	},
}

// CloudLoadBalancerDataLowerCaseAttributesMap represents the map of attribute for CloudLoadBalancerData.
var CloudLoadBalancerDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"attachedentities": {
		AllowedChoices: []string{},
		BSONFieldName:  "attachedentities",
		ConvertedName:  "AttachedEntities",
		Description:    `ID of associated objects with this load balancer.`,
		Exposed:        true,
		Name:           "attachedEntities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"listenertargetmapping": {
		AllowedChoices: []string{},
		BSONFieldName:  "listenertargetmapping",
		ConvertedName:  "Listenertargetmapping",
		Description:    `Mapping of a listener to its associated target group ID list.`,
		Exposed:        true,
		Name:           "listenertargetmapping",
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `The name of the load balancer.`,
		Exposed:        true,
		Name:           "name",
		Stored:         true,
		Type:           "string",
	},
	"scheme": {
		AllowedChoices: []string{},
		BSONFieldName:  "scheme",
		ConvertedName:  "Scheme",
		Description:    `The scheme tells whether the load balancer is internet facing or internal.`,
		Exposed:        true,
		Name:           "scheme",
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesCloudLoadBalancerData struct {
	AttachedEntities      []string            `bson:"attachedentities"`
	Listenertargetmapping map[string][]string `bson:"listenertargetmapping"`
	Name                  string              `bson:"name"`
	Scheme                string              `bson:"scheme"`
}
