package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudRouteData represents the model of a cloudroutedata
type CloudRouteData struct {
	// The gateway that this route table is associated with.
	GatewayID string `json:"gatewayID" msgpack:"gatewayID" bson:"gatewayid" mapstructure:"gatewayID,omitempty"`

	// Indicates that this is the default route table for the VPC.
	MainTable bool `json:"mainTable" msgpack:"mainTable" bson:"maintable" mapstructure:"mainTable,omitempty"`

	// Routes associated with this route table.
	Routelist []*CloudRoute `json:"routelist" msgpack:"routelist" bson:"routelist" mapstructure:"routelist,omitempty"`

	// The list of subnets that this route table is associated with.
	SubnetAssociations []string `json:"subnetAssociations" msgpack:"subnetAssociations" bson:"subnetassociations" mapstructure:"subnetAssociations,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudRouteData returns a new *CloudRouteData
func NewCloudRouteData() *CloudRouteData {

	return &CloudRouteData{
		ModelVersion:       1,
		Routelist:          []*CloudRoute{},
		SubnetAssociations: []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRouteData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudRouteData{}

	s.GatewayID = o.GatewayID
	s.MainTable = o.MainTable
	s.Routelist = o.Routelist
	s.SubnetAssociations = o.SubnetAssociations

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudRouteData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudRouteData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.GatewayID = s.GatewayID
	o.MainTable = s.MainTable
	o.Routelist = s.Routelist
	o.SubnetAssociations = s.SubnetAssociations

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudRouteData) BleveType() string {

	return "cloudroutedata"
}

// DeepCopy returns a deep copy if the CloudRouteData.
func (o *CloudRouteData) DeepCopy() *CloudRouteData {

	if o == nil {
		return nil
	}

	out := &CloudRouteData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudRouteData.
func (o *CloudRouteData) DeepCopyInto(out *CloudRouteData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudRouteData: %s", err))
	}

	*out = *target.(*CloudRouteData)
}

// Validate valides the current information stored into the structure.
func (o *CloudRouteData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Routelist {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*CloudRouteData) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudRouteDataAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudRouteDataLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudRouteData) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudRouteDataAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudRouteData) ValueForAttribute(name string) interface{} {

	switch name {
	case "gatewayID":
		return o.GatewayID
	case "mainTable":
		return o.MainTable
	case "routelist":
		return o.Routelist
	case "subnetAssociations":
		return o.SubnetAssociations
	}

	return nil
}

// CloudRouteDataAttributesMap represents the map of attribute for CloudRouteData.
var CloudRouteDataAttributesMap = map[string]elemental.AttributeSpecification{
	"GatewayID": {
		AllowedChoices: []string{},
		BSONFieldName:  "gatewayid",
		ConvertedName:  "GatewayID",
		Description:    `The gateway that this route table is associated with.`,
		Exposed:        true,
		Name:           "gatewayID",
		Stored:         true,
		Type:           "string",
	},
	"MainTable": {
		AllowedChoices: []string{},
		BSONFieldName:  "maintable",
		ConvertedName:  "MainTable",
		Description:    `Indicates that this is the default route table for the VPC.`,
		Exposed:        true,
		Name:           "mainTable",
		Stored:         true,
		Type:           "boolean",
	},
	"Routelist": {
		AllowedChoices: []string{},
		BSONFieldName:  "routelist",
		ConvertedName:  "Routelist",
		Description:    `Routes associated with this route table.`,
		Exposed:        true,
		Name:           "routelist",
		Stored:         true,
		SubType:        "cloudroute",
		Type:           "refList",
	},
	"SubnetAssociations": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetassociations",
		ConvertedName:  "SubnetAssociations",
		Description:    `The list of subnets that this route table is associated with.`,
		Exposed:        true,
		Name:           "subnetAssociations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// CloudRouteDataLowerCaseAttributesMap represents the map of attribute for CloudRouteData.
var CloudRouteDataLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"gatewayid": {
		AllowedChoices: []string{},
		BSONFieldName:  "gatewayid",
		ConvertedName:  "GatewayID",
		Description:    `The gateway that this route table is associated with.`,
		Exposed:        true,
		Name:           "gatewayID",
		Stored:         true,
		Type:           "string",
	},
	"maintable": {
		AllowedChoices: []string{},
		BSONFieldName:  "maintable",
		ConvertedName:  "MainTable",
		Description:    `Indicates that this is the default route table for the VPC.`,
		Exposed:        true,
		Name:           "mainTable",
		Stored:         true,
		Type:           "boolean",
	},
	"routelist": {
		AllowedChoices: []string{},
		BSONFieldName:  "routelist",
		ConvertedName:  "Routelist",
		Description:    `Routes associated with this route table.`,
		Exposed:        true,
		Name:           "routelist",
		Stored:         true,
		SubType:        "cloudroute",
		Type:           "refList",
	},
	"subnetassociations": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetassociations",
		ConvertedName:  "SubnetAssociations",
		Description:    `The list of subnets that this route table is associated with.`,
		Exposed:        true,
		Name:           "subnetAssociations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesCloudRouteData struct {
	GatewayID          string        `bson:"gatewayid"`
	MainTable          bool          `bson:"maintable"`
	Routelist          []*CloudRoute `bson:"routelist"`
	SubnetAssociations []string      `bson:"subnetassociations"`
}
