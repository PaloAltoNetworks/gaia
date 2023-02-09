// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RenderTemplateIdentity represents the Identity of the object.
var RenderTemplateIdentity = elemental.Identity{
	Name:     "rendertemplate",
	Category: "rendertemplates",
	Package:  "ignis",
	Private:  false,
}

// RenderTemplatesList represents a list of RenderTemplates
type RenderTemplatesList []*RenderTemplate

// Identity returns the identity of the objects in the list.
func (o RenderTemplatesList) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Copy returns a pointer to a copy the RenderTemplatesList.
func (o RenderTemplatesList) Copy() elemental.Identifiables {

	copy := append(RenderTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RenderTemplatesList.
func (o RenderTemplatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RenderTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RenderTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RenderTemplatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RenderTemplatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RenderTemplatesList converted to SparseRenderTemplatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RenderTemplatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRenderTemplatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRenderTemplate)
	}

	return out
}

// Version returns the version of the content.
func (o RenderTemplatesList) Version() int {

	return 1
}

// RenderTemplate represents the model of a rendertemplate
type RenderTemplate struct {
	// Holds the rendered template.
	Output string `json:"output" msgpack:"output" bson:"-" mapstructure:"output,omitempty"`

	// Contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Template of the recipe.
	Template string `json:"template" msgpack:"template" bson:"-" mapstructure:"template,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRenderTemplate returns a new *RenderTemplate
func NewRenderTemplate() *RenderTemplate {

	return &RenderTemplate{
		ModelVersion: 1,
		Parameters:   map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *RenderTemplate) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RenderTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RenderTemplate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderTemplate) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRenderTemplate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderTemplate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRenderTemplate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *RenderTemplate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *RenderTemplate) BleveType() string {

	return "rendertemplate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *RenderTemplate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RenderTemplate) Doc() string {

	return `
This API allows to render a provided generic go template and get back the
results, after applying the provided values plus a bunch of additional
information related to the product (like namespace, organizational metadata
etc). The templating engine allows to call various helpers functions,
and aims to be one to one compatible with Kubernetes templating, so you can
use it to generate your Kubernetes template when they embeds additional
security policies needing information like namespace, API URL etc, so they
can be sent to an operator. (function ` + "`" + `env` + "`" + `, ` + "`" + `expandenv` + "`" + ` and ` + "`" + `getHostByName` + "`" + `
are disabled).

##### String functions

- ` + "`" + `trim` + "`" + `: removes space from either side of a string.
- ` + "`" + `trimAll` + "`" + `: Remove given characters from the front or back of a string.
- ` + "`" + `trimSuffix` + "`" + `: Trim just the suffix from a string.
- ` + "`" + `trimPrefix` + "`" + `: Trim just the prefix from a string.
- ` + "`" + `upper` + "`" + `: Convert the entire string to uppercase.
- ` + "`" + `lower` + "`" + `: Convert the entire string to lowercase
- ` + "`" + `title` + "`" + `: Convert to title case.
- ` + "`" + `untitle` + "`" + `: Remove title casing. ` + "`" + `untitle` + "`" + ` "Hello World" produces hello world.
- ` + "`" + `repeat` + "`" + `: Repeat a string multiple times.
- ` + "`" + `substr` + "`" + `: Get a sub string from a string. It takes three parameters (start,
end, string).
- ` + "`" + `nospace` + "`" + `: Remove all white space from a string.
- ` + "`" + `trunc` + "`" + `: Truncate a string (and add no suffix).
- ` + "`" + `abbrev` + "`" + `: Truncate a string with ellipses.
- ` + "`" + `abbrevboth` + "`" + `: Abbreviate both sides.
- ` + "`" + `initials` + "`" + `: Given multiple words, take the first letter of each word and
combine.
- ` + "`" + `randAlphaNum` + "`" + `: Generate cryptographically secure random strings ` + "`" + `(0-9azA-Z)` + "`" + `.
- ` + "`" + `randAlpha` + "`" + `: Generate cryptographically secure random strings ` + "`" + `(a-zA-Z)` + "`" + `.
- ` + "`" + `randNumeric` + "`" + `: Generate cryptographically secure random strings ` + "`" + `(0-9)` + "`" + `.
- ` + "`" + `randAscii` + "`" + `: Generate cryptographically secure random strings (printable ASCII
chars).
- ` + "`" + `wrap` + "`" + `: Lets you specify the string to wrap with.
- ` + "`" + `contains` + "`" + `: Test to see if one string is contained inside of another.
- ` + "`" + `hasSuffix` + "`" + `: Test whether a string has a given suffix.
- ` + "`" + `hasPrefix` + "`" + `: Test whether a string has a given prefix.
- ` + "`" + `quote` + "`" + `: Wrap a string in double quotes.
- ` + "`" + `squote` + "`" + `: Wrap a string in single quotes.
- ` + "`" + `cat` + "`" + `: The cat function concatenates multiple strings together into one,
separating them with spaces.
- ` + "`" + `indent` + "`" + `: Indents every line in a given string to the specified indent width.
- ` + "`" + `nindent` + "`" + `: Same as the ` + "`" + `indent` + "`" + ` function, but adds a new line to the
beginning of the string.
- ` + "`" + `replace` + "`" + `: Perform simple string replacement.
- ` + "`" + `plural` + "`" + `: Pluralize a string.
- ` + "`" + `snakecase` + "`" + `: Convert string from ` + "`" + `camelCase` + "`" + ` to ` + "`" + `snake_case` + "`" + `.
- ` + "`" + `camelcase` + "`" + `: Convert string from ` + "`" + `snake_case` + "`" + ` to ` + "`" + `CamelCase` + "`" + `.
- ` + "`" + `kebabcase` + "`" + `: Convert string from ` + "`" + `camelCase` + "`" + ` to ` + "`" + `kebab-case` + "`" + `.
- ` + "`" + `swapcase` + "`" + `: Swap the case of a string using a word based algorithm.
- ` + "`" + `shuffle` + "`" + `: Shuffle a string.
- ` + "`" + `regexMatch` + "`" + `: Returns true if the input string contains any match of the
regular expression.
- ` + "`" + `regexFindAll` + "`" + `: Returns a slice of all matches of the regular expression in
the input string.
- ` + "`" + `regexFind` + "`" + `: Return the first (left most) match of the regular expression in
the input string.
- ` + "`" + `regexReplaceAll` + "`" + `: Returns a copy of the input string, replacing matches of
the Regexp with the replacement string replacement.
- ` + "`" + `regexReplaceAllLiteral` + "`" + `: Returns a copy of the input string, replacing
matches of the Regexp with the replacement string replacement.
- ` + "`" + `regexSplit` + "`" + `: Slices the input string into sub strings separated by the
expression and returns a slice of the sub strings between those expression
matches.

##### String list functions
- ` + "`" + `join` + "`" + `: Join a list of strings into a single string, with the given separator.
- ` + "`" + `splitList` + "`" + `: Split a string into a list of strings (returns a list).
- ` + "`" + `split` + "`" + `: Split a string into a list of strings (returns a map).
- ` + "`" + `splitn` + "`" + `: Splits a string into a dictionary with a maximum length of N.
- ` + "`" + `sortAlpha` + "`" + `: Sort a list of strings into alphabetical (lexicographical) order.

##### Integer math functions
- ` + "`" + `add` + "`" + `: Sum numbers. Accepts two or more inputs.
- ` + "`" + `add1` + "`" + `: Increment by 1.
- ` + "`" + `sub` + "`" + `: Subtract two numbers.
- ` + "`" + `div` + "`" + `: Perform integer division.
- ` + "`" + `mod` + "`" + `: modulo.
- ` + "`" + `mul` + "`" + `: Multiply numbers. Accepts two or more inputs.
- ` + "`" + `max` + "`" + `: Return the largest of a series of integers.
- ` + "`" + `min` + "`" + `: Return the smallest of a series of integers
- ` + "`" + `floor` + "`" + `: Returns the greatest float value less than or equal to input value.
- ` + "`" + `round` + "`" + `: Returns a float value with the remainder rounded to the given number
to digits after the decimal point.

##### Integer slice functions
- ` + "`" + `until` + "`" + `: builds a range of integers.
- ` + "`" + `untilStep` + "`" + `: generates a list of counting integers. Allows you to define a
start, stop, and step.

##### Float math functions
- ` + "`" + `floor` + "`" + `: Returns the greatest float value less than or equal to input value.
- ` + "`" + `round` + "`" + `: Returns a float value with the remainder rounded to the given number
to digits after the decimal point.

##### Date functions
- ` + "`" + `now` + "`" + `: The current date/time.
- ` + "`" + `ago` + "`" + `: Returns duration from ` + "`" + `time.Now` + "`" + ` in seconds resolution.
- ` + "`" + `date` + "`" + `: formats a date.
- ` + "`" + `dateInZone` + "`" + `: Same as date, but with a timezone.
- ` + "`" + `unixEpoch` + "`" + `: Returns the seconds since the Unix epoch.
- ` + "`" + `dateModify` + "`" + `: Takes a modification and a date and returns the timestamp.
- ` + "`" + `htmlDate` + "`" + `: Formats a date for inserting into an HTML date picker input field.
- ` + "`" + `htmlDateInZone` + "`" + `: Same as ` + "`" + `htmlDate` + "`" + `, but with a timezone.
- ` + "`" + `toDate` + "`" + `: converts a string to a date.

##### Default functions
- ` + "`" + `default` + "`" + `: Set a default value if second value is zero.
- ` + "`" + `empty` + "`" + `: returns true if the given value is considered empty.
- ` + "`" + `coalesce` + "`" + `: takes a list of values and returns the first non-empty one.
- ` + "`" + `all` + "`" + `: takes a list of values and returns true if all values are non-empty.
- ` + "`" + `any` + "`" + `: takes a list of values and returns true if any value is non-empty.
- ` + "`" + `toJson` + "`" + `: encodes an item into a JSON string.
- ` + "`" + `toPrettyJson` + "`" + `: encodes an item into a pretty (indented) JSON string.
- ` + "`" + `ternary` + "`" + `: This is similar to the C ternary operator.

##### Encoding and decoding functions
- ` + "`" + `b64enc` + "`" + `: Encode to Base 64.
- ` + "`" + `b64dec` + "`" + `: decode Base 64.
- ` + "`" + `b32enc` + "`" + `: Encode to Base 32.
- ` + "`" + `b32dec` + "`" + `: decode Base 32.

##### Lists and list functions
- ` + "`" + `first` + "`" + `: Get the head item.
- ` + "`" + `rest` + "`" + `: get the tail of the list (everything but the first item).
- ` + "`" + `last` + "`" + `: Get the last item on a list.
- ` + "`" + `initial` + "`" + `: Returns all but the last element.
- ` + "`" + `append` + "`" + `: Append a new item to a list, creating a new list.
- ` + "`" + `prepend` + "`" + `: Push an element onto the front of a list
- ` + "`" + `reverse` + "`" + `: Produce a new list with the reversed elements of the given list.
- ` + "`" + `uniq` + "`" + `: Generate a list with all of the duplicates removed.
- ` + "`" + `without` + "`" + `: filters items out of a list.
- ` + "`" + `has` + "`" + `: Test to see if a list has a particular element,
- ` + "`" + `compact` + "`" + `: Accepts a list and removes entries with empty values.
- ` + "`" + `slice` + "`" + `: get partial elements of a List

##### Dictionaries and dictionary functions
- ` + "`" + `dict` + "`" + `: Creates dictionaries from a list of pairs.
- ` + "`" + `set` + "`" + `: add a new key/value pair to a dictionary.
- ` + "`" + `unset` + "`" + `: Given a map and a key, delete the key from the map.
- ` + "`" + `hasKey` + "`" + `: returns true if the given dictionary contains the given key.
- ` + "`" + `pluck` + "`" + `: Give one key and multiple maps, and get a list of all of the matches.
- ` + "`" + `merge` + "`" + `: Merge two or more dictionaries into one, giving precedence to the
destination dictionary.
- ` + "`" + `mergeOverwrite` + "`" + `: Merge two or more dictionaries into one, giving precedence
from right to left, effectively overwriting values in the destination
dictionary.
- ` + "`" + `keys` + "`" + `: Return a list of all of the keys in one or more dictionary types.
- ` + "`" + `pick` + "`" + `: The pick function selects just the given keys out of a dictionary,
creating a new dictionary.
- ` + "`" + `omit` + "`" + `: similar to ` + "`" + `pick` + "`" + `, except it returns a new dictionary with all the
keys that
do not match the given keys.
- ` + "`" + `values` + "`" + `: The values function is similar to keys, except it returns a new list
with all the values of the source dictionary.
- ` + "`" + `deepCopy` + "`" + `: takes a value and makes a deep copy of the value.

##### Type conversion functions
- ` + "`" + `atoi` + "`" + `: Convert a string to an integer.
- ` + "`" + `float64` + "`" + `: Convert to a ` + "`" + `float64` + "`" + `.
- ` + "`" + `int` + "`" + `: Convert to an int at the system’s width.
- ` + "`" + `int64` + "`" + `: Convert to an ` + "`" + `int64` + "`" + `.
- ` + "`" + `toDecimal` + "`" + `: Convert a Unix octal to a ` + "`" + `int64` + "`" + `.
- ` + "`" + `toString` + "`" + `: Convert to a string.
- ` + "`" + `toStrings` + "`" + `: Convert a list, slice, or array to a list of strings.

##### Paths and file path functions
- ` + "`" + `base` + "`" + `: Return the last element of a path.
- ` + "`" + `dir` + "`" + `: Return the directory, stripping the last part of the path.
- ` + "`" + `clean` + "`" + `: Clean up a path.
- ` + "`" + `ext` + "`" + `: Return the file extension.
- ` + "`" + `isAbs` + "`" + `: check whether a path is absolute.

##### Flow control function
- ` + "`" + `fail` + "`" + `: Unconditionally returns an empty string and an error with the
specified text.

##### UUID function
- ` + "`" + `uuidv4` + "`" + `: generates a UUID version 4.

##### OS function:
- ` + "`" + `env` + "`" + `: Not available. Will returns "disabled".
- ` + "`" + `expandenv` + "`" + `: not available. Will return "disabled".

##### Semantic version functions
- ` + "`" + `semver` + "`" + `: Parses a string into a Semantic Version.
- ` + "`" + `semverCompare` + "`" + `: returns true if the constraint matches, or false if it does
not match.

##### Reflection function
- ` + "`" + `kindOf` + "`" + `: Return the kind of an object.
- ` + "`" + `kindIs` + "`" + `: Returns true if the given object is of the given kind.
- ` + "`" + `typeOf` + "`" + `: returns the underlying type of a value.
- ` + "`" + `typeIs` + "`" + `: Like ` + "`" + `kindIs` + "`" + `, but for types.
- ` + "`" + `typeIsLike` + "`" + `: Works as ` + "`" + `typeIs` + "`" + `, except that it also de-references pointers.
- ` + "`" + `deepEqual` + "`" + `: returns true if two values are deeply equal.

##### Cryptographic functions
- ` + "`" + `sha1sum` + "`" + `: Computes the SHA 1 digest of the given string.
- ` + "`" + `sha256sum` + "`" + `: Computes the SHA 256 digest of the given string.
- ` + "`" + `alder32sum` + "`" + `: Computes the Alder 32 digest of the given string.
- ` + "`" + `genPrivateKey` + "`" + `: generates a new private key encoded into a PEM block.
- ` + "`" + `buildCustomCert` + "`" + `: Generate a TLS certificate from given base 64 encoded
certificate and key.
- ` + "`" + `genCA` + "`" + `: generates a new, self-signed x509 certificate authority using a
2048 bit RSA private key.
- ` + "`" + `genCAWithKey` + "`" + `: generates a new, self-signed x509 certificate authority using
a given private key.
- ` + "`" + `genSelfSignedCert` + "`" + `: generates a new, self-signed x509 certificate using a
2048 bit RSA private key.
- ` + "`" + `genSelfSignedCertWithKey` + "`" + `: generates a new, self-signed x509 certificate
using a given private key.
- ` + "`" + `genSignedCert` + "`" + `: function generates a new, x509 certificate signed by the
specified CA, using a 2048 bit RSA private key.
- ` + "`" + `genSignedCertWithKey` + "`" + `: generates a new, x509 certificate signed by the
specified CA, using a given private key.
- ` + "`" + `encryptAES` + "`" + `: encrypts text with AES 256 CBC and returns a base 64 encoded
string.
- ` + "`" + `decryptAES` + "`" + `: receives a base 64 string encoded by the AES 256 CBC algorithm
and returns the decoded text.

##### Network function
- ` + "`" + `getHostByName` + "`" + `: Not available. Will return "disabled".`
}

func (o *RenderTemplate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *RenderTemplate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRenderTemplate{
			Output:     &o.Output,
			Parameters: &o.Parameters,
			Template:   &o.Template,
		}
	}

	sp := &SparseRenderTemplate{}
	for _, f := range fields {
		switch f {
		case "output":
			sp.Output = &(o.Output)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "template":
			sp.Template = &(o.Template)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRenderTemplate to the object.
func (o *RenderTemplate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRenderTemplate)
	if so.Output != nil {
		o.Output = *so.Output
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Template != nil {
		o.Template = *so.Template
	}
}

// DeepCopy returns a deep copy if the RenderTemplate.
func (o *RenderTemplate) DeepCopy() *RenderTemplate {

	if o == nil {
		return nil
	}

	out := &RenderTemplate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RenderTemplate.
func (o *RenderTemplate) DeepCopyInto(out *RenderTemplate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RenderTemplate: %s", err))
	}

	*out = *target.(*RenderTemplate)
}

// Validate valides the current information stored into the structure.
func (o *RenderTemplate) Validate() error {

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
func (*RenderTemplate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RenderTemplateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RenderTemplateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RenderTemplate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RenderTemplateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RenderTemplate) ValueForAttribute(name string) any {

	switch name {
	case "output":
		return o.Output
	case "parameters":
		return o.Parameters
	case "template":
		return o.Template
	}

	return nil
}

// RenderTemplateAttributesMap represents the map of attribute for RenderTemplate.
var RenderTemplateAttributesMap = map[string]elemental.AttributeSpecification{
	"Output": {
		AllowedChoices: []string{},
		ConvertedName:  "Output",
		Description:    `Holds the rendered template.`,
		Exposed:        true,
		Name:           "output",
		Type:           "string",
	},
	"Parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Template": {
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe.`,
		Exposed:        true,
		Name:           "template",
		Type:           "string",
	},
}

// RenderTemplateLowerCaseAttributesMap represents the map of attribute for RenderTemplate.
var RenderTemplateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"output": {
		AllowedChoices: []string{},
		ConvertedName:  "Output",
		Description:    `Holds the rendered template.`,
		Exposed:        true,
		Name:           "output",
		Type:           "string",
	},
	"parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"template": {
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe.`,
		Exposed:        true,
		Name:           "template",
		Type:           "string",
	},
}

// SparseRenderTemplatesList represents a list of SparseRenderTemplates
type SparseRenderTemplatesList []*SparseRenderTemplate

// Identity returns the identity of the objects in the list.
func (o SparseRenderTemplatesList) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Copy returns a pointer to a copy the SparseRenderTemplatesList.
func (o SparseRenderTemplatesList) Copy() elemental.Identifiables {

	copy := append(SparseRenderTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRenderTemplatesList.
func (o SparseRenderTemplatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRenderTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRenderTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRenderTemplatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRenderTemplatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRenderTemplatesList converted to RenderTemplatesList.
func (o SparseRenderTemplatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRenderTemplatesList) Version() int {

	return 1
}

// SparseRenderTemplate represents the sparse version of a rendertemplate.
type SparseRenderTemplate struct {
	// Holds the rendered template.
	Output *string `json:"output,omitempty" msgpack:"output,omitempty" bson:"-" mapstructure:"output,omitempty"`

	// Contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"-" mapstructure:"parameters,omitempty"`

	// Template of the recipe.
	Template *string `json:"template,omitempty" msgpack:"template,omitempty" bson:"-" mapstructure:"template,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseRenderTemplate returns a new  SparseRenderTemplate.
func NewSparseRenderTemplate() *SparseRenderTemplate {
	return &SparseRenderTemplate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRenderTemplate) Identity() elemental.Identity {

	return RenderTemplateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRenderTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRenderTemplate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderTemplate) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRenderTemplate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderTemplate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRenderTemplate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseRenderTemplate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRenderTemplate) ToPlain() elemental.PlainIdentifiable {

	out := NewRenderTemplate()
	if o.Output != nil {
		out.Output = *o.Output
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Template != nil {
		out.Template = *o.Template
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRenderTemplate.
func (o *SparseRenderTemplate) DeepCopy() *SparseRenderTemplate {

	if o == nil {
		return nil
	}

	out := &SparseRenderTemplate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRenderTemplate.
func (o *SparseRenderTemplate) DeepCopyInto(out *SparseRenderTemplate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRenderTemplate: %s", err))
	}

	*out = *target.(*SparseRenderTemplate)
}

type mongoAttributesRenderTemplate struct {
}
type mongoAttributesSparseRenderTemplate struct {
}
