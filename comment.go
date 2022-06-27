package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Comment represents the model of a comment
type Comment struct {
	// The claims of the author.
	Claims []string `json:"claims" msgpack:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// The content of the comment.
	Content string `json:"content" msgpack:"content" bson:"content" mapstructure:"content,omitempty"`

	// The date of the comment.
	Date time.Time `json:"date" msgpack:"date" bson:"date" mapstructure:"date,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewComment returns a new *Comment
func NewComment() *Comment {

	return &Comment{
		ModelVersion: 1,
		Claims:       []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Comment) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesComment{}

	s.Claims = o.Claims
	s.Content = o.Content
	s.Date = o.Date

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Comment) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesComment{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Claims = s.Claims
	o.Content = s.Content
	o.Date = s.Date

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *Comment) BleveType() string {

	return "comment"
}

// DeepCopy returns a deep copy if the Comment.
func (o *Comment) DeepCopy() *Comment {

	if o == nil {
		return nil
	}

	out := &Comment{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Comment.
func (o *Comment) DeepCopyInto(out *Comment) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Comment: %s", err))
	}

	*out = *target.(*Comment)
}

// Validate valides the current information stored into the structure.
func (o *Comment) Validate() error {

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
func (*Comment) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CommentAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CommentLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Comment) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CommentAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Comment) ValueForAttribute(name string) interface{} {

	switch name {
	case "claims":
		return o.Claims
	case "content":
		return o.Content
	case "date":
		return o.Date
	}

	return nil
}

// CommentAttributesMap represents the map of attribute for Comment.
var CommentAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": {
		AllowedChoices: []string{},
		BSONFieldName:  "claims",
		ConvertedName:  "Claims",
		Description:    `The claims of the author.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Content": {
		AllowedChoices: []string{},
		BSONFieldName:  "content",
		ConvertedName:  "Content",
		Description:    `The content of the comment.`,
		Exposed:        true,
		Name:           "content",
		Stored:         true,
		Type:           "string",
	},
	"Date": {
		AllowedChoices: []string{},
		BSONFieldName:  "date",
		ConvertedName:  "Date",
		Description:    `The date of the comment.`,
		Exposed:        true,
		Name:           "date",
		Stored:         true,
		Type:           "time",
	},
}

// CommentLowerCaseAttributesMap represents the map of attribute for Comment.
var CommentLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": {
		AllowedChoices: []string{},
		BSONFieldName:  "claims",
		ConvertedName:  "Claims",
		Description:    `The claims of the author.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"content": {
		AllowedChoices: []string{},
		BSONFieldName:  "content",
		ConvertedName:  "Content",
		Description:    `The content of the comment.`,
		Exposed:        true,
		Name:           "content",
		Stored:         true,
		Type:           "string",
	},
	"date": {
		AllowedChoices: []string{},
		BSONFieldName:  "date",
		ConvertedName:  "Date",
		Description:    `The date of the comment.`,
		Exposed:        true,
		Name:           "date",
		Stored:         true,
		Type:           "time",
	},
}

type mongoAttributesComment struct {
	Claims  []string  `bson:"claims"`
	Content string    `bson:"content"`
	Date    time.Time `bson:"date"`
}
