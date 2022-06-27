package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RecipeOptionsAppCrendentialFormatValue represents the possible values for attribute "appCrendentialFormat".
type RecipeOptionsAppCrendentialFormatValue string

const (
	// RecipeOptionsAppCrendentialFormatJSON represents the value JSON.
	RecipeOptionsAppCrendentialFormatJSON RecipeOptionsAppCrendentialFormatValue = "JSON"

	// RecipeOptionsAppCrendentialFormatYAML represents the value YAML.
	RecipeOptionsAppCrendentialFormatYAML RecipeOptionsAppCrendentialFormatValue = "YAML"
)

// RecipeOptions represents the model of a recipeoptions
type RecipeOptions struct {
	// Indicates the format of the app credential.
	AppCrendentialFormat RecipeOptionsAppCrendentialFormatValue `json:"appCrendentialFormat" msgpack:"appCrendentialFormat" bson:"appcrendentialformat" mapstructure:"appCrendentialFormat,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRecipeOptions returns a new *RecipeOptions
func NewRecipeOptions() *RecipeOptions {

	return &RecipeOptions{
		ModelVersion:         1,
		AppCrendentialFormat: RecipeOptionsAppCrendentialFormatJSON,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RecipeOptions) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRecipeOptions{}

	s.AppCrendentialFormat = o.AppCrendentialFormat

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RecipeOptions) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRecipeOptions{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AppCrendentialFormat = s.AppCrendentialFormat

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *RecipeOptions) BleveType() string {

	return "recipeoptions"
}

// DeepCopy returns a deep copy if the RecipeOptions.
func (o *RecipeOptions) DeepCopy() *RecipeOptions {

	if o == nil {
		return nil
	}

	out := &RecipeOptions{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RecipeOptions.
func (o *RecipeOptions) DeepCopyInto(out *RecipeOptions) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RecipeOptions: %s", err))
	}

	*out = *target.(*RecipeOptions)
}

// Validate valides the current information stored into the structure.
func (o *RecipeOptions) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("appCrendentialFormat", string(o.AppCrendentialFormat), []string{"JSON", "YAML"}, false); err != nil {
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
func (*RecipeOptions) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RecipeOptionsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RecipeOptionsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RecipeOptions) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RecipeOptionsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RecipeOptions) ValueForAttribute(name string) interface{} {

	switch name {
	case "appCrendentialFormat":
		return o.AppCrendentialFormat
	}

	return nil
}

// RecipeOptionsAttributesMap represents the map of attribute for RecipeOptions.
var RecipeOptionsAttributesMap = map[string]elemental.AttributeSpecification{
	"AppCrendentialFormat": {
		AllowedChoices: []string{"JSON", "YAML"},
		BSONFieldName:  "appcrendentialformat",
		ConvertedName:  "AppCrendentialFormat",
		DefaultValue:   RecipeOptionsAppCrendentialFormatJSON,
		Description:    `Indicates the format of the app credential.`,
		Exposed:        true,
		Name:           "appCrendentialFormat",
		Stored:         true,
		Type:           "enum",
	},
}

// RecipeOptionsLowerCaseAttributesMap represents the map of attribute for RecipeOptions.
var RecipeOptionsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"appcrendentialformat": {
		AllowedChoices: []string{"JSON", "YAML"},
		BSONFieldName:  "appcrendentialformat",
		ConvertedName:  "AppCrendentialFormat",
		DefaultValue:   RecipeOptionsAppCrendentialFormatJSON,
		Description:    `Indicates the format of the app credential.`,
		Exposed:        true,
		Name:           "appCrendentialFormat",
		Stored:         true,
		Type:           "enum",
	},
}

type mongoAttributesRecipeOptions struct {
	AppCrendentialFormat RecipeOptionsAppCrendentialFormatValue `bson:"appcrendentialformat"`
}
