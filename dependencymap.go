package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DependencyMapIdentity represents the Identity of the object.
var DependencyMapIdentity = elemental.Identity{
	Name:     "dependencymap",
	Category: "dependencymaps",
	Package:  "jenova",
	Private:  false,
}

// DependencyMapsList represents a list of DependencyMaps
type DependencyMapsList []*DependencyMap

// Identity returns the identity of the objects in the list.
func (o DependencyMapsList) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Copy returns a pointer to a copy the DependencyMapsList.
func (o DependencyMapsList) Copy() elemental.Identifiables {

	copy := append(DependencyMapsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DependencyMapsList.
func (o DependencyMapsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DependencyMapsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DependencyMap))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DependencyMapsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DependencyMapsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DependencyMapsList converted to SparseDependencyMapsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DependencyMapsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDependencyMapsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDependencyMap)
	}

	return out
}

// Version returns the version of the content.
func (o DependencyMapsList) Version() int {

	return 1
}

// DependencyMap represents the model of a dependencymap
type DependencyMap struct {
	// The edges of the map.
	Edges map[string]*GraphEdge `json:"edges" msgpack:"edges" bson:"-" mapstructure:"edges,omitempty"`

	// Provides information about the group values.
	Groups map[string]*GraphGroup `json:"groups" msgpack:"groups" bson:"-" mapstructure:"groups,omitempty"`

	// Refers to the nodes of the map.
	Nodes map[string]*GraphNode `json:"nodes" msgpack:"nodes" bson:"-" mapstructure:"nodes,omitempty"`

	// Provides suggested views based on relevant tags.
	ViewSuggestions []string `json:"viewSuggestions" msgpack:"viewSuggestions" bson:"-" mapstructure:"viewSuggestions,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDependencyMap returns a new *DependencyMap
func NewDependencyMap() *DependencyMap {

	return &DependencyMap{
		ModelVersion:    1,
		Edges:           map[string]*GraphEdge{},
		Groups:          map[string]*GraphGroup{},
		Nodes:           map[string]*GraphNode{},
		ViewSuggestions: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMap) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMap) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMap) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DependencyMap) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDependencyMap{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DependencyMap) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDependencyMap{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DependencyMap) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DependencyMap) BleveType() string {

	return "dependencymap"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DependencyMap) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DependencyMap) Doc() string {

	return `Returns a data structure representing the graph of all processing units
and their connections in a particular namespace, in a given time window. To pass
the time window you can use the query parameters ` + "`" + `startAbsolute` + "`" + `, ` + "`" + `endAbsolute` + "`" + `,
` + "`" + `startRelative` + "`" + `, ` + "`" + `endRelative` + "`" + `.

For example:

` + "`" + `/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000` + "`" + `.`
}

func (o *DependencyMap) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DependencyMap) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDependencyMap{
			Edges:           &o.Edges,
			Groups:          &o.Groups,
			Nodes:           &o.Nodes,
			ViewSuggestions: &o.ViewSuggestions,
		}
	}

	sp := &SparseDependencyMap{}
	for _, f := range fields {
		switch f {
		case "edges":
			sp.Edges = &(o.Edges)
		case "groups":
			sp.Groups = &(o.Groups)
		case "nodes":
			sp.Nodes = &(o.Nodes)
		case "viewSuggestions":
			sp.ViewSuggestions = &(o.ViewSuggestions)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDependencyMap to the object.
func (o *DependencyMap) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDependencyMap)
	if so.Edges != nil {
		o.Edges = *so.Edges
	}
	if so.Groups != nil {
		o.Groups = *so.Groups
	}
	if so.Nodes != nil {
		o.Nodes = *so.Nodes
	}
	if so.ViewSuggestions != nil {
		o.ViewSuggestions = *so.ViewSuggestions
	}
}

// DeepCopy returns a deep copy if the DependencyMap.
func (o *DependencyMap) DeepCopy() *DependencyMap {

	if o == nil {
		return nil
	}

	out := &DependencyMap{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DependencyMap.
func (o *DependencyMap) DeepCopyInto(out *DependencyMap) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DependencyMap: %s", err))
	}

	*out = *target.(*DependencyMap)
}

// Validate valides the current information stored into the structure.
func (o *DependencyMap) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Edges {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.Groups {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.Nodes {
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
func (*DependencyMap) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DependencyMapAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DependencyMapLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DependencyMap) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DependencyMap) ValueForAttribute(name string) interface{} {

	switch name {
	case "edges":
		return o.Edges
	case "groups":
		return o.Groups
	case "nodes":
		return o.Nodes
	case "viewSuggestions":
		return o.ViewSuggestions
	}

	return nil
}

// DependencyMapAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapAttributesMap = map[string]elemental.AttributeSpecification{
	"Edges": {
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `The edges of the map.`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		SubType:        "graphedge",
		Type:           "refMap",
	},
	"Groups": {
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description:    `Provides information about the group values.`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		SubType:        "graphgroup",
		Type:           "refMap",
	},
	"Nodes": {
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `Refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		SubType:        "graphnode",
		Type:           "refMap",
	},
	"ViewSuggestions": {
		AllowedChoices: []string{},
		ConvertedName:  "ViewSuggestions",
		Description:    `Provides suggested views based on relevant tags.`,
		Exposed:        true,
		Name:           "viewSuggestions",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// DependencyMapLowerCaseAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"edges": {
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `The edges of the map.`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		SubType:        "graphedge",
		Type:           "refMap",
	},
	"groups": {
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description:    `Provides information about the group values.`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		SubType:        "graphgroup",
		Type:           "refMap",
	},
	"nodes": {
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `Refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		SubType:        "graphnode",
		Type:           "refMap",
	},
	"viewsuggestions": {
		AllowedChoices: []string{},
		ConvertedName:  "ViewSuggestions",
		Description:    `Provides suggested views based on relevant tags.`,
		Exposed:        true,
		Name:           "viewSuggestions",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
}

// SparseDependencyMapsList represents a list of SparseDependencyMaps
type SparseDependencyMapsList []*SparseDependencyMap

// Identity returns the identity of the objects in the list.
func (o SparseDependencyMapsList) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Copy returns a pointer to a copy the SparseDependencyMapsList.
func (o SparseDependencyMapsList) Copy() elemental.Identifiables {

	copy := append(SparseDependencyMapsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDependencyMapsList.
func (o SparseDependencyMapsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDependencyMapsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDependencyMap))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDependencyMapsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDependencyMapsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDependencyMapsList converted to DependencyMapsList.
func (o SparseDependencyMapsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDependencyMapsList) Version() int {

	return 1
}

// SparseDependencyMap represents the sparse version of a dependencymap.
type SparseDependencyMap struct {
	// The edges of the map.
	Edges *map[string]*GraphEdge `json:"edges,omitempty" msgpack:"edges,omitempty" bson:"-" mapstructure:"edges,omitempty"`

	// Provides information about the group values.
	Groups *map[string]*GraphGroup `json:"groups,omitempty" msgpack:"groups,omitempty" bson:"-" mapstructure:"groups,omitempty"`

	// Refers to the nodes of the map.
	Nodes *map[string]*GraphNode `json:"nodes,omitempty" msgpack:"nodes,omitempty" bson:"-" mapstructure:"nodes,omitempty"`

	// Provides suggested views based on relevant tags.
	ViewSuggestions *[]string `json:"viewSuggestions,omitempty" msgpack:"viewSuggestions,omitempty" bson:"-" mapstructure:"viewSuggestions,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDependencyMap returns a new  SparseDependencyMap.
func NewSparseDependencyMap() *SparseDependencyMap {
	return &SparseDependencyMap{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDependencyMap) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDependencyMap) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDependencyMap) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDependencyMap) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDependencyMap{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDependencyMap) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDependencyMap{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDependencyMap) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDependencyMap) ToPlain() elemental.PlainIdentifiable {

	out := NewDependencyMap()
	if o.Edges != nil {
		out.Edges = *o.Edges
	}
	if o.Groups != nil {
		out.Groups = *o.Groups
	}
	if o.Nodes != nil {
		out.Nodes = *o.Nodes
	}
	if o.ViewSuggestions != nil {
		out.ViewSuggestions = *o.ViewSuggestions
	}

	return out
}

// DeepCopy returns a deep copy if the SparseDependencyMap.
func (o *SparseDependencyMap) DeepCopy() *SparseDependencyMap {

	if o == nil {
		return nil
	}

	out := &SparseDependencyMap{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDependencyMap.
func (o *SparseDependencyMap) DeepCopyInto(out *SparseDependencyMap) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDependencyMap: %s", err))
	}

	*out = *target.(*SparseDependencyMap)
}

type mongoAttributesDependencyMap struct {
}
type mongoAttributesSparseDependencyMap struct {
}
