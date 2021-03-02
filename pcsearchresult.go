package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PCSearchResultIdentity represents the Identity of the object.
var PCSearchResultIdentity = elemental.Identity{
	Name:     "pcsearchresult",
	Category: "pcsearchresults",
	Package:  "karl",
	Private:  false,
}

// PCSearchResultsList represents a list of PCSearchResults
type PCSearchResultsList []*PCSearchResult

// Identity returns the identity of the objects in the list.
func (o PCSearchResultsList) Identity() elemental.Identity {

	return PCSearchResultIdentity
}

// Copy returns a pointer to a copy the PCSearchResultsList.
func (o PCSearchResultsList) Copy() elemental.Identifiables {

	copy := append(PCSearchResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PCSearchResultsList.
func (o PCSearchResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PCSearchResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PCSearchResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PCSearchResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PCSearchResultsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PCSearchResultsList converted to SparsePCSearchResultsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PCSearchResultsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePCSearchResultsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePCSearchResult)
	}

	return out
}

// Version returns the version of the content.
func (o PCSearchResultsList) Version() int {

	return 1
}

// PCSearchResult represents the model of a pcsearchresult
type PCSearchResult struct {
	// The payload of the search result.
	Items *ReportsQuery `json:"items" msgpack:"items" bson:"-" mapstructure:"items,omitempty"`

	// The pagination token for next page.
	NextPageToken string `json:"nextPageToken" msgpack:"nextPageToken" bson:"-" mapstructure:"nextPageToken,omitempty"`

	// The total number of result items.
	TotalRows int `json:"totalRows" msgpack:"totalRows" bson:"-" mapstructure:"totalRows,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPCSearchResult returns a new *PCSearchResult
func NewPCSearchResult() *PCSearchResult {

	return &PCSearchResult{
		ModelVersion: 1,
		Items:        NewReportsQuery(),
	}
}

// Identity returns the Identity of the object.
func (o *PCSearchResult) Identity() elemental.Identity {

	return PCSearchResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PCSearchResult) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PCSearchResult) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCSearchResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPCSearchResult{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCSearchResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPCSearchResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PCSearchResult) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PCSearchResult) BleveType() string {

	return "pcsearchresult"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PCSearchResult) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PCSearchResult) Doc() string {

	return `Represents the result data of RQL search.`
}

func (o *PCSearchResult) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PCSearchResult) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePCSearchResult{
			Items:         o.Items,
			NextPageToken: &o.NextPageToken,
			TotalRows:     &o.TotalRows,
		}
	}

	sp := &SparsePCSearchResult{}
	for _, f := range fields {
		switch f {
		case "items":
			sp.Items = o.Items
		case "nextPageToken":
			sp.NextPageToken = &(o.NextPageToken)
		case "totalRows":
			sp.TotalRows = &(o.TotalRows)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePCSearchResult to the object.
func (o *PCSearchResult) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePCSearchResult)
	if so.Items != nil {
		o.Items = so.Items
	}
	if so.NextPageToken != nil {
		o.NextPageToken = *so.NextPageToken
	}
	if so.TotalRows != nil {
		o.TotalRows = *so.TotalRows
	}
}

// DeepCopy returns a deep copy if the PCSearchResult.
func (o *PCSearchResult) DeepCopy() *PCSearchResult {

	if o == nil {
		return nil
	}

	out := &PCSearchResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PCSearchResult.
func (o *PCSearchResult) DeepCopyInto(out *PCSearchResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PCSearchResult: %s", err))
	}

	*out = *target.(*PCSearchResult)
}

// Validate valides the current information stored into the structure.
func (o *PCSearchResult) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Items != nil {
		elemental.ResetDefaultForZeroValues(o.Items)
		if err := o.Items.Validate(); err != nil {
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
func (*PCSearchResult) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PCSearchResultAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PCSearchResultLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PCSearchResult) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PCSearchResultAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PCSearchResult) ValueForAttribute(name string) interface{} {

	switch name {
	case "items":
		return o.Items
	case "nextPageToken":
		return o.NextPageToken
	case "totalRows":
		return o.TotalRows
	}

	return nil
}

// PCSearchResultAttributesMap represents the map of attribute for PCSearchResult.
var PCSearchResultAttributesMap = map[string]elemental.AttributeSpecification{
	"Items": {
		AllowedChoices: []string{},
		ConvertedName:  "Items",
		Description:    `The payload of the search result.`,
		Exposed:        true,
		Name:           "items",
		ReadOnly:       true,
		SubType:        "reportsquery",
		Type:           "ref",
	},
	"NextPageToken": {
		AllowedChoices: []string{},
		ConvertedName:  "NextPageToken",
		Description:    `The pagination token for next page.`,
		Exposed:        true,
		Name:           "nextPageToken",
		ReadOnly:       true,
		Type:           "string",
	},
	"TotalRows": {
		AllowedChoices: []string{},
		ConvertedName:  "TotalRows",
		Description:    `The total number of result items.`,
		Exposed:        true,
		Name:           "totalRows",
		ReadOnly:       true,
		Type:           "integer",
	},
}

// PCSearchResultLowerCaseAttributesMap represents the map of attribute for PCSearchResult.
var PCSearchResultLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"items": {
		AllowedChoices: []string{},
		ConvertedName:  "Items",
		Description:    `The payload of the search result.`,
		Exposed:        true,
		Name:           "items",
		ReadOnly:       true,
		SubType:        "reportsquery",
		Type:           "ref",
	},
	"nextpagetoken": {
		AllowedChoices: []string{},
		ConvertedName:  "NextPageToken",
		Description:    `The pagination token for next page.`,
		Exposed:        true,
		Name:           "nextPageToken",
		ReadOnly:       true,
		Type:           "string",
	},
	"totalrows": {
		AllowedChoices: []string{},
		ConvertedName:  "TotalRows",
		Description:    `The total number of result items.`,
		Exposed:        true,
		Name:           "totalRows",
		ReadOnly:       true,
		Type:           "integer",
	},
}

// SparsePCSearchResultsList represents a list of SparsePCSearchResults
type SparsePCSearchResultsList []*SparsePCSearchResult

// Identity returns the identity of the objects in the list.
func (o SparsePCSearchResultsList) Identity() elemental.Identity {

	return PCSearchResultIdentity
}

// Copy returns a pointer to a copy the SparsePCSearchResultsList.
func (o SparsePCSearchResultsList) Copy() elemental.Identifiables {

	copy := append(SparsePCSearchResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePCSearchResultsList.
func (o SparsePCSearchResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePCSearchResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePCSearchResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePCSearchResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePCSearchResultsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePCSearchResultsList converted to PCSearchResultsList.
func (o SparsePCSearchResultsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePCSearchResultsList) Version() int {

	return 1
}

// SparsePCSearchResult represents the sparse version of a pcsearchresult.
type SparsePCSearchResult struct {
	// The payload of the search result.
	Items *ReportsQuery `json:"items,omitempty" msgpack:"items,omitempty" bson:"-" mapstructure:"items,omitempty"`

	// The pagination token for next page.
	NextPageToken *string `json:"nextPageToken,omitempty" msgpack:"nextPageToken,omitempty" bson:"-" mapstructure:"nextPageToken,omitempty"`

	// The total number of result items.
	TotalRows *int `json:"totalRows,omitempty" msgpack:"totalRows,omitempty" bson:"-" mapstructure:"totalRows,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePCSearchResult returns a new  SparsePCSearchResult.
func NewSparsePCSearchResult() *SparsePCSearchResult {
	return &SparsePCSearchResult{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePCSearchResult) Identity() elemental.Identity {

	return PCSearchResultIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePCSearchResult) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePCSearchResult) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCSearchResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePCSearchResult{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCSearchResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePCSearchResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePCSearchResult) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePCSearchResult) ToPlain() elemental.PlainIdentifiable {

	out := NewPCSearchResult()
	if o.Items != nil {
		out.Items = o.Items
	}
	if o.NextPageToken != nil {
		out.NextPageToken = *o.NextPageToken
	}
	if o.TotalRows != nil {
		out.TotalRows = *o.TotalRows
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePCSearchResult.
func (o *SparsePCSearchResult) DeepCopy() *SparsePCSearchResult {

	if o == nil {
		return nil
	}

	out := &SparsePCSearchResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePCSearchResult.
func (o *SparsePCSearchResult) DeepCopyInto(out *SparsePCSearchResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePCSearchResult: %s", err))
	}

	*out = *target.(*SparsePCSearchResult)
}

type mongoAttributesPCSearchResult struct {
}
type mongoAttributesSparsePCSearchResult struct {
}
