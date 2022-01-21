package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TextIndexIdentity represents the Identity of the object.
var TextIndexIdentity = elemental.Identity{
	Name:     "textindex",
	Category: "textindexes",
	Package:  "gogole",
	Private:  true,
}

// TextIndexsList represents a list of TextIndexs
type TextIndexsList []*TextIndex

// Identity returns the identity of the objects in the list.
func (o TextIndexsList) Identity() elemental.Identity {

	return TextIndexIdentity
}

// Copy returns a pointer to a copy the TextIndexsList.
func (o TextIndexsList) Copy() elemental.Identifiables {

	copy := append(TextIndexsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TextIndexsList.
func (o TextIndexsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TextIndexsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TextIndex))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TextIndexsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TextIndexsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TextIndexsList converted to SparseTextIndexsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TextIndexsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTextIndexsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTextIndex)
	}

	return out
}

// Version returns the version of the content.
func (o TextIndexsList) Version() int {

	return 1
}

// TextIndex represents the model of a textindex
type TextIndex struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// raw text data that are indexed.
	Data string `json:"-" msgpack:"-" bson:"data" mapstructure:"-,omitempty"`

	// Date of the last indexing.
	Date time.Time `json:"-" msgpack:"-" bson:"date" mapstructure:"-,omitempty"`

	// Contains the ID of the match.
	ObjectID string `json:"-" msgpack:"-" bson:"objectid" mapstructure:"-,omitempty"`

	// Contains the identity of the match.
	ObjectIdentity string `json:"-" msgpack:"-" bson:"objectidentity" mapstructure:"-,omitempty"`

	// Contains the namespace of the match.
	ObjectNamespace string `json:"-" msgpack:"-" bson:"objectnamespace" mapstructure:"-,omitempty"`

	// Contains a match score.
	Score float64 `json:"-" msgpack:"-" bson:"score" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTextIndex returns a new *TextIndex
func NewTextIndex() *TextIndex {

	return &TextIndex{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *TextIndex) Identity() elemental.Identity {

	return TextIndexIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TextIndex) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TextIndex) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TextIndex) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTextIndex{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Data = o.Data
	s.Date = o.Date
	s.ObjectID = o.ObjectID
	s.ObjectIdentity = o.ObjectIdentity
	s.ObjectNamespace = o.ObjectNamespace
	s.Score = o.Score
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TextIndex) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTextIndex{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Data = s.Data
	o.Date = s.Date
	o.ObjectID = s.ObjectID
	o.ObjectIdentity = s.ObjectIdentity
	o.ObjectNamespace = s.ObjectNamespace
	o.Score = s.Score
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TextIndex) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TextIndex) BleveType() string {

	return "textindex"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TextIndex) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TextIndex) Doc() string {

	return `Internal storage for full text search.`
}

func (o *TextIndex) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetZHash returns the ZHash of the receiver.
func (o *TextIndex) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *TextIndex) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *TextIndex) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *TextIndex) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TextIndex) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTextIndex{
			ID:              &o.ID,
			Data:            &o.Data,
			Date:            &o.Date,
			ObjectID:        &o.ObjectID,
			ObjectIdentity:  &o.ObjectIdentity,
			ObjectNamespace: &o.ObjectNamespace,
			Score:           &o.Score,
			ZHash:           &o.ZHash,
			Zone:            &o.Zone,
		}
	}

	sp := &SparseTextIndex{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "data":
			sp.Data = &(o.Data)
		case "date":
			sp.Date = &(o.Date)
		case "objectID":
			sp.ObjectID = &(o.ObjectID)
		case "objectIdentity":
			sp.ObjectIdentity = &(o.ObjectIdentity)
		case "objectNamespace":
			sp.ObjectNamespace = &(o.ObjectNamespace)
		case "score":
			sp.Score = &(o.Score)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTextIndex to the object.
func (o *TextIndex) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTextIndex)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.ObjectID != nil {
		o.ObjectID = *so.ObjectID
	}
	if so.ObjectIdentity != nil {
		o.ObjectIdentity = *so.ObjectIdentity
	}
	if so.ObjectNamespace != nil {
		o.ObjectNamespace = *so.ObjectNamespace
	}
	if so.Score != nil {
		o.Score = *so.Score
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the TextIndex.
func (o *TextIndex) DeepCopy() *TextIndex {

	if o == nil {
		return nil
	}

	out := &TextIndex{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TextIndex.
func (o *TextIndex) DeepCopyInto(out *TextIndex) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TextIndex: %s", err))
	}

	*out = *target.(*TextIndex)
}

// Validate valides the current information stored into the structure.
func (o *TextIndex) Validate() error {

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
func (*TextIndex) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TextIndexAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TextIndexLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TextIndex) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TextIndexAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TextIndex) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "data":
		return o.Data
	case "date":
		return o.Date
	case "objectID":
		return o.ObjectID
	case "objectIdentity":
		return o.ObjectIdentity
	case "objectNamespace":
		return o.ObjectNamespace
	case "score":
		return o.Score
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// TextIndexAttributesMap represents the map of attribute for TextIndex.
var TextIndexAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `raw text data that are indexed.`,
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
	"Date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `Date of the last indexing.`,
		Name:           "date",
		Stored:         true,
		Type:           "time",
	},
	"ObjectID": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `Contains the ID of the match.`,
		Name:           "objectID",
		Stored:         true,
		Type:           "string",
	},
	"ObjectIdentity": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectIdentity",
		Description:    `Contains the identity of the match.`,
		Name:           "objectIdentity",
		Stored:         true,
		Type:           "string",
	},
	"ObjectNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectNamespace",
		Description:    `Contains the namespace of the match.`,
		Name:           "objectNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Score": {
		AllowedChoices: []string{},
		ConvertedName:  "Score",
		Description:    `Contains a match score.`,
		Name:           "score",
		Stored:         true,
		Type:           "float",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// TextIndexLowerCaseAttributesMap represents the map of attribute for TextIndex.
var TextIndexLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `raw text data that are indexed.`,
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
	"date": {
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `Date of the last indexing.`,
		Name:           "date",
		Stored:         true,
		Type:           "time",
	},
	"objectid": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `Contains the ID of the match.`,
		Name:           "objectID",
		Stored:         true,
		Type:           "string",
	},
	"objectidentity": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectIdentity",
		Description:    `Contains the identity of the match.`,
		Name:           "objectIdentity",
		Stored:         true,
		Type:           "string",
	},
	"objectnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectNamespace",
		Description:    `Contains the namespace of the match.`,
		Name:           "objectNamespace",
		Stored:         true,
		Type:           "string",
	},
	"score": {
		AllowedChoices: []string{},
		ConvertedName:  "Score",
		Description:    `Contains a match score.`,
		Name:           "score",
		Stored:         true,
		Type:           "float",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseTextIndexsList represents a list of SparseTextIndexs
type SparseTextIndexsList []*SparseTextIndex

// Identity returns the identity of the objects in the list.
func (o SparseTextIndexsList) Identity() elemental.Identity {

	return TextIndexIdentity
}

// Copy returns a pointer to a copy the SparseTextIndexsList.
func (o SparseTextIndexsList) Copy() elemental.Identifiables {

	copy := append(SparseTextIndexsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTextIndexsList.
func (o SparseTextIndexsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTextIndexsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTextIndex))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTextIndexsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTextIndexsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTextIndexsList converted to TextIndexsList.
func (o SparseTextIndexsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTextIndexsList) Version() int {

	return 1
}

// SparseTextIndex represents the sparse version of a textindex.
type SparseTextIndex struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// raw text data that are indexed.
	Data *string `json:"-" msgpack:"-" bson:"data,omitempty" mapstructure:"-,omitempty"`

	// Date of the last indexing.
	Date *time.Time `json:"-" msgpack:"-" bson:"date,omitempty" mapstructure:"-,omitempty"`

	// Contains the ID of the match.
	ObjectID *string `json:"-" msgpack:"-" bson:"objectid,omitempty" mapstructure:"-,omitempty"`

	// Contains the identity of the match.
	ObjectIdentity *string `json:"-" msgpack:"-" bson:"objectidentity,omitempty" mapstructure:"-,omitempty"`

	// Contains the namespace of the match.
	ObjectNamespace *string `json:"-" msgpack:"-" bson:"objectnamespace,omitempty" mapstructure:"-,omitempty"`

	// Contains a match score.
	Score *float64 `json:"-" msgpack:"-" bson:"score,omitempty" mapstructure:"-,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTextIndex returns a new  SparseTextIndex.
func NewSparseTextIndex() *SparseTextIndex {
	return &SparseTextIndex{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTextIndex) Identity() elemental.Identity {

	return TextIndexIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTextIndex) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTextIndex) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTextIndex) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTextIndex{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Data != nil {
		s.Data = o.Data
	}
	if o.Date != nil {
		s.Date = o.Date
	}
	if o.ObjectID != nil {
		s.ObjectID = o.ObjectID
	}
	if o.ObjectIdentity != nil {
		s.ObjectIdentity = o.ObjectIdentity
	}
	if o.ObjectNamespace != nil {
		s.ObjectNamespace = o.ObjectNamespace
	}
	if o.Score != nil {
		s.Score = o.Score
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTextIndex) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTextIndex{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Data != nil {
		o.Data = s.Data
	}
	if s.Date != nil {
		o.Date = s.Date
	}
	if s.ObjectID != nil {
		o.ObjectID = s.ObjectID
	}
	if s.ObjectIdentity != nil {
		o.ObjectIdentity = s.ObjectIdentity
	}
	if s.ObjectNamespace != nil {
		o.ObjectNamespace = s.ObjectNamespace
	}
	if s.Score != nil {
		o.Score = s.Score
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTextIndex) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTextIndex) ToPlain() elemental.PlainIdentifiable {

	out := NewTextIndex()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.ObjectID != nil {
		out.ObjectID = *o.ObjectID
	}
	if o.ObjectIdentity != nil {
		out.ObjectIdentity = *o.ObjectIdentity
	}
	if o.ObjectNamespace != nil {
		out.ObjectNamespace = *o.ObjectNamespace
	}
	if o.Score != nil {
		out.Score = *o.Score
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseTextIndex) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseTextIndex) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseTextIndex) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseTextIndex) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseTextIndex.
func (o *SparseTextIndex) DeepCopy() *SparseTextIndex {

	if o == nil {
		return nil
	}

	out := &SparseTextIndex{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTextIndex.
func (o *SparseTextIndex) DeepCopyInto(out *SparseTextIndex) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTextIndex: %s", err))
	}

	*out = *target.(*SparseTextIndex)
}

type mongoAttributesTextIndex struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	Data            string        `bson:"data"`
	Date            time.Time     `bson:"date"`
	ObjectID        string        `bson:"objectid"`
	ObjectIdentity  string        `bson:"objectidentity"`
	ObjectNamespace string        `bson:"objectnamespace"`
	Score           float64       `bson:"score"`
	ZHash           int           `bson:"zhash"`
	Zone            int           `bson:"zone"`
}
type mongoAttributesSparseTextIndex struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	Data            *string       `bson:"data,omitempty"`
	Date            *time.Time    `bson:"date,omitempty"`
	ObjectID        *string       `bson:"objectid,omitempty"`
	ObjectIdentity  *string       `bson:"objectidentity,omitempty"`
	ObjectNamespace *string       `bson:"objectnamespace,omitempty"`
	Score           *float64      `bson:"score,omitempty"`
	ZHash           *int          `bson:"zhash,omitempty"`
	Zone            *int          `bson:"zone,omitempty"`
}
