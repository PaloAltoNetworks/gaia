package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// UnknownIPAddressIdentity represents the Identity of the object.
var UnknownIPAddressIdentity = elemental.Identity{
	Name:     "unknownipaddress",
	Category: "unknownipaddresses",
	Package:  "karl",
	Private:  false,
}

// UnknownIPAddressList represents a list of UnknownIPAddress
type UnknownIPAddressList []*UnknownIPAddress

// Identity returns the identity of the objects in the list.
func (o UnknownIPAddressList) Identity() elemental.Identity {

	return UnknownIPAddressIdentity
}

// Copy returns a pointer to a copy the UnknownIPAddressList.
func (o UnknownIPAddressList) Copy() elemental.Identifiables {

	copy := append(UnknownIPAddressList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the UnknownIPAddressList.
func (o UnknownIPAddressList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(UnknownIPAddressList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*UnknownIPAddress))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o UnknownIPAddressList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UnknownIPAddressList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the UnknownIPAddressList converted to SparseUnknownIPAddressList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o UnknownIPAddressList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseUnknownIPAddressList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseUnknownIPAddress)
	}

	return out
}

// Version returns the version of the content.
func (o UnknownIPAddressList) Version() int {

	return 1
}

// UnknownIPAddress represents the model of a unknownipaddress
type UnknownIPAddress struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// IP/FQDN encountered.
	Address string `json:"address" msgpack:"address" bson:"address" mapstructure:"address,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Number of times the address was seen for the source.
	Occurrences int `json:"occurrences" msgpack:"occurrences" bson:"occurrences" mapstructure:"occurrences,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Identifier of the source of the address.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// Identity of the source of the address.
	SourceIdentity string `json:"sourceIdentity" msgpack:"sourceIdentity" bson:"sourceidentity" mapstructure:"sourceIdentity,omitempty"`

	// Namespace of the source of the address.
	SourceNamespace string `json:"sourceNamespace" msgpack:"sourceNamespace" bson:"sourcenamespace" mapstructure:"sourceNamespace,omitempty"`

	// The timestamp of the source of the address.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUnknownIPAddress returns a new *UnknownIPAddress
func NewUnknownIPAddress() *UnknownIPAddress {

	return &UnknownIPAddress{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *UnknownIPAddress) Identity() elemental.Identity {

	return UnknownIPAddressIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UnknownIPAddress) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UnknownIPAddress) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UnknownIPAddress) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesUnknownIPAddress{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Address = o.Address
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Occurrences = o.Occurrences
	s.Protected = o.Protected
	s.SourceID = o.SourceID
	s.SourceIdentity = o.SourceIdentity
	s.SourceNamespace = o.SourceNamespace
	s.Timestamp = o.Timestamp
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *UnknownIPAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesUnknownIPAddress{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Address = s.Address
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Occurrences = s.Occurrences
	o.Protected = s.Protected
	o.SourceID = s.SourceID
	o.SourceIdentity = s.SourceIdentity
	o.SourceNamespace = s.SourceNamespace
	o.Timestamp = s.Timestamp
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *UnknownIPAddress) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *UnknownIPAddress) BleveType() string {

	return "unknownipaddress"
}

// DefaultOrder returns the list of default ordering fields.
func (o *UnknownIPAddress) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *UnknownIPAddress) Doc() string {

	return `Holds the IP/FQDN of flows going to somewhere (default) external network.`
}

func (o *UnknownIPAddress) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *UnknownIPAddress) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *UnknownIPAddress) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *UnknownIPAddress) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *UnknownIPAddress) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *UnknownIPAddress) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *UnknownIPAddress) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *UnknownIPAddress) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *UnknownIPAddress) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *UnknownIPAddress) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *UnknownIPAddress) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *UnknownIPAddress) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *UnknownIPAddress) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *UnknownIPAddress) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *UnknownIPAddress) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *UnknownIPAddress) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *UnknownIPAddress) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *UnknownIPAddress) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *UnknownIPAddress) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *UnknownIPAddress) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *UnknownIPAddress) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *UnknownIPAddress) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *UnknownIPAddress) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *UnknownIPAddress) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseUnknownIPAddress{
			ID:                   &o.ID,
			Address:              &o.Address,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Occurrences:          &o.Occurrences,
			Protected:            &o.Protected,
			SourceID:             &o.SourceID,
			SourceIdentity:       &o.SourceIdentity,
			SourceNamespace:      &o.SourceNamespace,
			Timestamp:            &o.Timestamp,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseUnknownIPAddress{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "address":
			sp.Address = &(o.Address)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "occurrences":
			sp.Occurrences = &(o.Occurrences)
		case "protected":
			sp.Protected = &(o.Protected)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceIdentity":
			sp.SourceIdentity = &(o.SourceIdentity)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseUnknownIPAddress to the object.
func (o *UnknownIPAddress) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseUnknownIPAddress)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Address != nil {
		o.Address = *so.Address
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Occurrences != nil {
		o.Occurrences = *so.Occurrences
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceIdentity != nil {
		o.SourceIdentity = *so.SourceIdentity
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the UnknownIPAddress.
func (o *UnknownIPAddress) DeepCopy() *UnknownIPAddress {

	if o == nil {
		return nil
	}

	out := &UnknownIPAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *UnknownIPAddress.
func (o *UnknownIPAddress) DeepCopyInto(out *UnknownIPAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy UnknownIPAddress: %s", err))
	}

	*out = *target.(*UnknownIPAddress)
}

// Validate valides the current information stored into the structure.
func (o *UnknownIPAddress) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("address", o.Address); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceIdentity", o.SourceIdentity); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceNamespace", o.SourceNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*UnknownIPAddress) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := UnknownIPAddressAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return UnknownIPAddressLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*UnknownIPAddress) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return UnknownIPAddressAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *UnknownIPAddress) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "address":
		return o.Address
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "occurrences":
		return o.Occurrences
	case "protected":
		return o.Protected
	case "sourceID":
		return o.SourceID
	case "sourceIdentity":
		return o.SourceIdentity
	case "sourceNamespace":
		return o.SourceNamespace
	case "timestamp":
		return o.Timestamp
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// UnknownIPAddressAttributesMap represents the map of attribute for UnknownIPAddress.
var UnknownIPAddressAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"Address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `IP/FQDN encountered.`,
		Exposed:        true,
		Name:           "address",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Occurrences": {
		AllowedChoices: []string{},
		BSONFieldName:  "occurrences",
		ConvertedName:  "Occurrences",
		Description:    `Number of times the address was seen for the source.`,
		Exposed:        true,
		Name:           "occurrences",
		Stored:         true,
		Type:           "integer",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"SourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceid",
		ConvertedName:  "SourceID",
		Description:    `Identifier of the source of the address.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceIdentity": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceidentity",
		ConvertedName:  "SourceIdentity",
		Description:    `Identity of the source of the address.`,
		Exposed:        true,
		Name:           "sourceIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcenamespace",
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the source of the address.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "timestamp",
		ConvertedName:  "Timestamp",
		Description:    `The timestamp of the source of the address.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
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
		BSONFieldName:  "zone",
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

// UnknownIPAddressLowerCaseAttributesMap represents the map of attribute for UnknownIPAddress.
var UnknownIPAddressLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"address": {
		AllowedChoices: []string{},
		BSONFieldName:  "address",
		ConvertedName:  "Address",
		Description:    `IP/FQDN encountered.`,
		Exposed:        true,
		Name:           "address",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"occurrences": {
		AllowedChoices: []string{},
		BSONFieldName:  "occurrences",
		ConvertedName:  "Occurrences",
		Description:    `Number of times the address was seen for the source.`,
		Exposed:        true,
		Name:           "occurrences",
		Stored:         true,
		Type:           "integer",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"sourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceid",
		ConvertedName:  "SourceID",
		Description:    `Identifier of the source of the address.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourceidentity": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceidentity",
		ConvertedName:  "SourceIdentity",
		Description:    `Identity of the source of the address.`,
		Exposed:        true,
		Name:           "sourceIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourcenamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcenamespace",
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the source of the address.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "timestamp",
		ConvertedName:  "Timestamp",
		Description:    `The timestamp of the source of the address.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
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
		BSONFieldName:  "zone",
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

// SparseUnknownIPAddressList represents a list of SparseUnknownIPAddress
type SparseUnknownIPAddressList []*SparseUnknownIPAddress

// Identity returns the identity of the objects in the list.
func (o SparseUnknownIPAddressList) Identity() elemental.Identity {

	return UnknownIPAddressIdentity
}

// Copy returns a pointer to a copy the SparseUnknownIPAddressList.
func (o SparseUnknownIPAddressList) Copy() elemental.Identifiables {

	copy := append(SparseUnknownIPAddressList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseUnknownIPAddressList.
func (o SparseUnknownIPAddressList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseUnknownIPAddressList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseUnknownIPAddress))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseUnknownIPAddressList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseUnknownIPAddressList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseUnknownIPAddressList converted to UnknownIPAddressList.
func (o SparseUnknownIPAddressList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseUnknownIPAddressList) Version() int {

	return 1
}

// SparseUnknownIPAddress represents the sparse version of a unknownipaddress.
type SparseUnknownIPAddress struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// IP/FQDN encountered.
	Address *string `json:"address,omitempty" msgpack:"address,omitempty" bson:"address,omitempty" mapstructure:"address,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Number of times the address was seen for the source.
	Occurrences *int `json:"occurrences,omitempty" msgpack:"occurrences,omitempty" bson:"occurrences,omitempty" mapstructure:"occurrences,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Identifier of the source of the address.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// Identity of the source of the address.
	SourceIdentity *string `json:"sourceIdentity,omitempty" msgpack:"sourceIdentity,omitempty" bson:"sourceidentity,omitempty" mapstructure:"sourceIdentity,omitempty"`

	// Namespace of the source of the address.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"sourcenamespace,omitempty" mapstructure:"sourceNamespace,omitempty"`

	// The timestamp of the source of the address.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseUnknownIPAddress returns a new  SparseUnknownIPAddress.
func NewSparseUnknownIPAddress() *SparseUnknownIPAddress {
	return &SparseUnknownIPAddress{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseUnknownIPAddress) Identity() elemental.Identity {

	return UnknownIPAddressIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseUnknownIPAddress) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseUnknownIPAddress) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseUnknownIPAddress) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseUnknownIPAddress{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Address != nil {
		s.Address = o.Address
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Occurrences != nil {
		s.Occurrences = o.Occurrences
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.SourceIdentity != nil {
		s.SourceIdentity = o.SourceIdentity
	}
	if o.SourceNamespace != nil {
		s.SourceNamespace = o.SourceNamespace
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
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
func (o *SparseUnknownIPAddress) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseUnknownIPAddress{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Address != nil {
		o.Address = s.Address
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Occurrences != nil {
		o.Occurrences = s.Occurrences
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.SourceIdentity != nil {
		o.SourceIdentity = s.SourceIdentity
	}
	if s.SourceNamespace != nil {
		o.SourceNamespace = s.SourceNamespace
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
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
func (o *SparseUnknownIPAddress) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseUnknownIPAddress) ToPlain() elemental.PlainIdentifiable {

	out := NewUnknownIPAddress()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Address != nil {
		out.Address = *o.Address
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Occurrences != nil {
		out.Occurrences = *o.Occurrences
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceIdentity != nil {
		out.SourceIdentity = *o.SourceIdentity
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseUnknownIPAddress) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseUnknownIPAddress) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseUnknownIPAddress) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseUnknownIPAddress) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseUnknownIPAddress) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseUnknownIPAddress) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseUnknownIPAddress) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseUnknownIPAddress) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseUnknownIPAddress) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseUnknownIPAddress) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseUnknownIPAddress) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseUnknownIPAddress) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseUnknownIPAddress.
func (o *SparseUnknownIPAddress) DeepCopy() *SparseUnknownIPAddress {

	if o == nil {
		return nil
	}

	out := &SparseUnknownIPAddress{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseUnknownIPAddress.
func (o *SparseUnknownIPAddress) DeepCopyInto(out *SparseUnknownIPAddress) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseUnknownIPAddress: %s", err))
	}

	*out = *target.(*SparseUnknownIPAddress)
}

type mongoAttributesUnknownIPAddress struct {
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	Address              string              `bson:"address"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Namespace            string              `bson:"namespace"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Occurrences          int                 `bson:"occurrences"`
	Protected            bool                `bson:"protected"`
	SourceID             string              `bson:"sourceid"`
	SourceIdentity       string              `bson:"sourceidentity"`
	SourceNamespace      string              `bson:"sourcenamespace"`
	Timestamp            time.Time           `bson:"timestamp"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseUnknownIPAddress struct {
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	Address              *string              `bson:"address,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Occurrences          *int                 `bson:"occurrences,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	SourceID             *string              `bson:"sourceid,omitempty"`
	SourceIdentity       *string              `bson:"sourceidentity,omitempty"`
	SourceNamespace      *string              `bson:"sourcenamespace,omitempty"`
	Timestamp            *time.Time           `bson:"timestamp,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}