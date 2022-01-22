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

type mongoAttributesTimeSeriesRow struct {
}
