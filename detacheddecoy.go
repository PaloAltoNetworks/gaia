// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DetachedDecoy represents the model of a detacheddecoy
type DetachedDecoy struct {
	// If not empty, it means the decoy has ErrNotPermitted.
	ErrNotPermitted string `json:"errNotPermitted" msgpack:"errNotPermitted" bson:"-" mapstructure:"errNotPermitted,omitempty"`

	// The NodeUID which the decoy passed through.
	NodeUID *NodeUID `json:"nodeUID" msgpack:"nodeUID" bson:"-" mapstructure:"nodeUID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDetachedDecoy returns a new *DetachedDecoy
func NewDetachedDecoy() *DetachedDecoy {

	return &DetachedDecoy{
		ModelVersion: 1,
		NodeUID:      NewNodeUID(),
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DetachedDecoy) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDetachedDecoy{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DetachedDecoy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDetachedDecoy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *DetachedDecoy) BleveType() string {

	return "detacheddecoy"
}

// DeepCopy returns a deep copy if the DetachedDecoy.
func (o *DetachedDecoy) DeepCopy() *DetachedDecoy {

	if o == nil {
		return nil
	}

	out := &DetachedDecoy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DetachedDecoy.
func (o *DetachedDecoy) DeepCopyInto(out *DetachedDecoy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DetachedDecoy: %s", err))
	}

	*out = *target.(*DetachedDecoy)
}

// Validate valides the current information stored into the structure.
func (o *DetachedDecoy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.NodeUID != nil {
		elemental.ResetDefaultForZeroValues(o.NodeUID)
		if err := o.NodeUID.Validate(); err != nil {
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
func (*DetachedDecoy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DetachedDecoyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DetachedDecoyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DetachedDecoy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DetachedDecoyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DetachedDecoy) ValueForAttribute(name string) any {

	switch name {
	case "errNotPermitted":
		return o.ErrNotPermitted
	case "nodeUID":
		return o.NodeUID
	}

	return nil
}

// DetachedDecoyAttributesMap represents the map of attribute for DetachedDecoy.
var DetachedDecoyAttributesMap = map[string]elemental.AttributeSpecification{
	"ErrNotPermitted": {
		AllowedChoices: []string{},
		ConvertedName:  "ErrNotPermitted",
		Description:    `If not empty, it means the decoy has ErrNotPermitted.`,
		Exposed:        true,
		Name:           "errNotPermitted",
		Type:           "string",
	},
	"NodeUID": {
		AllowedChoices: []string{},
		ConvertedName:  "NodeUID",
		Description:    `The NodeUID which the decoy passed through.`,
		Exposed:        true,
		Name:           "nodeUID",
		SubType:        "nodeuid",
		Type:           "ref",
	},
}

// DetachedDecoyLowerCaseAttributesMap represents the map of attribute for DetachedDecoy.
var DetachedDecoyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"errnotpermitted": {
		AllowedChoices: []string{},
		ConvertedName:  "ErrNotPermitted",
		Description:    `If not empty, it means the decoy has ErrNotPermitted.`,
		Exposed:        true,
		Name:           "errNotPermitted",
		Type:           "string",
	},
	"nodeuid": {
		AllowedChoices: []string{},
		ConvertedName:  "NodeUID",
		Description:    `The NodeUID which the decoy passed through.`,
		Exposed:        true,
		Name:           "nodeUID",
		SubType:        "nodeuid",
		Type:           "ref",
	},
}

type mongoAttributesDetachedDecoy struct {
}
