package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TimeSeriesRow represents the model of a timeseriesrow
type TimeSeriesRow struct {
	// Columns of the row.
	Columns []string `json:"columns" msgpack:"columns" bson:"-" mapstructure:"columns,omitempty"`

	// Name of the row.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// List of tags.
	Tags map[string]string `json:"tags" msgpack:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// List of tags.
	Values [][]interface{} `json:"values" msgpack:"values" bson:"-" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTimeSeriesRow returns a new *TimeSeriesRow
func NewTimeSeriesRow() *TimeSeriesRow {

	return &TimeSeriesRow{
		ModelVersion: 1,
		Columns:      []string{},
		Tags:         map[string]string{},
		Values:       [][]interface{}{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TimeSeriesRow) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTimeSeriesRow{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TimeSeriesRow) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTimeSeriesRow{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *TimeSeriesRow) BleveType() string {

	return "timeseriesrow"
}

// DeepCopy returns a deep copy if the TimeSeriesRow.
func (o *TimeSeriesRow) DeepCopy() *TimeSeriesRow {

	if o == nil {
		return nil
	}

	out := &TimeSeriesRow{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TimeSeriesRow.
func (o *TimeSeriesRow) DeepCopyInto(out *TimeSeriesRow) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TimeSeriesRow: %s", err))
	}

	*out = *target.(*TimeSeriesRow)
}

// Validate valides the current information stored into the structure.
func (o *TimeSeriesRow) Validate() error {

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
func (*TimeSeriesRow) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TimeSeriesRowAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TimeSeriesRowLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TimeSeriesRow) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TimeSeriesRowAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TimeSeriesRow) ValueForAttribute(name string) interface{} {

	switch name {
	case "columns":
		return o.Columns
	case "name":
		return o.Name
	case "tags":
		return o.Tags
	case "values":
		return o.Values
	}

	return nil
}

// TimeSeriesRowAttributesMap represents the map of attribute for TimeSeriesRow.
var TimeSeriesRowAttributesMap = map[string]elemental.AttributeSpecification{
	"Columns": {
		AllowedChoices: []string{},
		ConvertedName:  "Columns",
		Description:    `Columns of the row.`,
		Exposed:        true,
		Name:           "columns",
		SubType:        "string",
		Type:           "list",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the row.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Values": {
		AllowedChoices: []string{},
		ConvertedName:  "Values",
		Description:    `List of tags.`,
		Exposed:        true,
		Name:           "values",
		SubType:        "[][]interface{}",
		Type:           "external",
	},
}

// TimeSeriesRowLowerCaseAttributesMap represents the map of attribute for TimeSeriesRow.
var TimeSeriesRowLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"columns": {
		AllowedChoices: []string{},
		ConvertedName:  "Columns",
		Description:    `Columns of the row.`,
		Exposed:        true,
		Name:           "columns",
		SubType:        "string",
		Type:           "list",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the row.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"values": {
		AllowedChoices: []string{},
		ConvertedName:  "Values",
		Description:    `List of tags.`,
		Exposed:        true,
		Name:           "values",
		SubType:        "[][]interface{}",
		Type:           "external",
	},
}

type mongoAttributesTimeSeriesRow struct {
}
