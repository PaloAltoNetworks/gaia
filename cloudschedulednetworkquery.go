package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudScheduledNetworkQueryIdentity represents the Identity of the object.
var CloudScheduledNetworkQueryIdentity = elemental.Identity{
	Name:     "cloudschedulednetworkquery",
	Category: "cloudschedulednetworkqueries",
	Package:  "vargid",
	Private:  true,
}

// CloudScheduledNetworkQueriesList represents a list of CloudScheduledNetworkQueries
type CloudScheduledNetworkQueriesList []*CloudScheduledNetworkQuery

// Identity returns the identity of the objects in the list.
func (o CloudScheduledNetworkQueriesList) Identity() elemental.Identity {

	return CloudScheduledNetworkQueryIdentity
}

// Copy returns a pointer to a copy the CloudScheduledNetworkQueriesList.
func (o CloudScheduledNetworkQueriesList) Copy() elemental.Identifiables {

	copy := append(CloudScheduledNetworkQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudScheduledNetworkQueriesList.
func (o CloudScheduledNetworkQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudScheduledNetworkQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudScheduledNetworkQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudScheduledNetworkQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudScheduledNetworkQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CloudScheduledNetworkQueriesList converted to SparseCloudScheduledNetworkQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudScheduledNetworkQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudScheduledNetworkQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudScheduledNetworkQuery)
	}

	return out
}

// Version returns the version of the content.
func (o CloudScheduledNetworkQueriesList) Version() int {

	return 1
}

// CloudScheduledNetworkQuery represents the model of a cloudschedulednetworkquery
type CloudScheduledNetworkQuery struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Prisma Cloud Alert Rule ID.
	AlertRuleID string `json:"alertRuleID" msgpack:"alertRuleID" bson:"alertruleid" mapstructure:"alertRuleID,omitempty"`

	// The result of the cloud network query.
	CloudGraphResult *CloudGraph `json:"cloudGraphResult" msgpack:"cloudGraphResult" bson:"-" mapstructure:"cloudGraphResult,omitempty"`

	// The cloud network query that should be used.
	CloudNetworkQuery *CloudNetworkQuery `json:"cloudNetworkQuery" msgpack:"cloudNetworkQuery" bson:"cloudnetworkquery" mapstructure:"cloudNetworkQuery,omitempty"`

	// Represents whether the associated policy was disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Result of the last successfully run query.
	LastExecutionTimestamp time.Time `json:"lastExecutionTimestamp" msgpack:"lastExecutionTimestamp" bson:"lastexecutiontimestamp" mapstructure:"lastExecutionTimestamp,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Prisma Cloud Policy ID.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudScheduledNetworkQuery returns a new *CloudScheduledNetworkQuery
func NewCloudScheduledNetworkQuery() *CloudScheduledNetworkQuery {

	return &CloudScheduledNetworkQuery{
		ModelVersion:      1,
		CloudGraphResult:  NewCloudGraph(),
		CloudNetworkQuery: NewCloudNetworkQuery(),
		MigrationsLog:     map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudScheduledNetworkQuery) Identity() elemental.Identity {

	return CloudScheduledNetworkQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudScheduledNetworkQuery) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudScheduledNetworkQuery) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudScheduledNetworkQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudScheduledNetworkQuery{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.AlertRuleID = o.AlertRuleID
	s.CloudNetworkQuery = o.CloudNetworkQuery
	s.Disabled = o.Disabled
	s.LastExecutionTimestamp = o.LastExecutionTimestamp
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.PolicyID = o.PolicyID
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudScheduledNetworkQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudScheduledNetworkQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.AlertRuleID = s.AlertRuleID
	o.CloudNetworkQuery = s.CloudNetworkQuery
	o.Disabled = s.Disabled
	o.LastExecutionTimestamp = s.LastExecutionTimestamp
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.PolicyID = s.PolicyID
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudScheduledNetworkQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudScheduledNetworkQuery) BleveType() string {

	return "cloudschedulednetworkquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudScheduledNetworkQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CloudScheduledNetworkQuery) Doc() string {

	return `CloudSchedulednNetworkQuery represents a CloudNetworkQuery that will be
scheduled periodically.`
}

func (o *CloudScheduledNetworkQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudScheduledNetworkQuery) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudScheduledNetworkQuery) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudScheduledNetworkQuery) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudScheduledNetworkQuery) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudScheduledNetworkQuery) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudScheduledNetworkQuery) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudScheduledNetworkQuery) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudScheduledNetworkQuery) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudScheduledNetworkQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudScheduledNetworkQuery{
			ID:                     &o.ID,
			AlertRuleID:            &o.AlertRuleID,
			CloudGraphResult:       o.CloudGraphResult,
			CloudNetworkQuery:      o.CloudNetworkQuery,
			Disabled:               &o.Disabled,
			LastExecutionTimestamp: &o.LastExecutionTimestamp,
			MigrationsLog:          &o.MigrationsLog,
			Namespace:              &o.Namespace,
			PolicyID:               &o.PolicyID,
			ZHash:                  &o.ZHash,
			Zone:                   &o.Zone,
		}
	}

	sp := &SparseCloudScheduledNetworkQuery{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "alertRuleID":
			sp.AlertRuleID = &(o.AlertRuleID)
		case "cloudGraphResult":
			sp.CloudGraphResult = o.CloudGraphResult
		case "cloudNetworkQuery":
			sp.CloudNetworkQuery = o.CloudNetworkQuery
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "lastExecutionTimestamp":
			sp.LastExecutionTimestamp = &(o.LastExecutionTimestamp)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudScheduledNetworkQuery to the object.
func (o *CloudScheduledNetworkQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudScheduledNetworkQuery)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AlertRuleID != nil {
		o.AlertRuleID = *so.AlertRuleID
	}
	if so.CloudGraphResult != nil {
		o.CloudGraphResult = so.CloudGraphResult
	}
	if so.CloudNetworkQuery != nil {
		o.CloudNetworkQuery = so.CloudNetworkQuery
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.LastExecutionTimestamp != nil {
		o.LastExecutionTimestamp = *so.LastExecutionTimestamp
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the CloudScheduledNetworkQuery.
func (o *CloudScheduledNetworkQuery) DeepCopy() *CloudScheduledNetworkQuery {

	if o == nil {
		return nil
	}

	out := &CloudScheduledNetworkQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudScheduledNetworkQuery.
func (o *CloudScheduledNetworkQuery) DeepCopyInto(out *CloudScheduledNetworkQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudScheduledNetworkQuery: %s", err))
	}

	*out = *target.(*CloudScheduledNetworkQuery)
}

// Validate valides the current information stored into the structure.
func (o *CloudScheduledNetworkQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.CloudGraphResult != nil {
		elemental.ResetDefaultForZeroValues(o.CloudGraphResult)
		if err := o.CloudGraphResult.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if o.CloudNetworkQuery != nil {
		elemental.ResetDefaultForZeroValues(o.CloudNetworkQuery)
		if err := o.CloudNetworkQuery.Validate(); err != nil {
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
func (*CloudScheduledNetworkQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudScheduledNetworkQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudScheduledNetworkQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudScheduledNetworkQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudScheduledNetworkQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudScheduledNetworkQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "alertRuleID":
		return o.AlertRuleID
	case "cloudGraphResult":
		return o.CloudGraphResult
	case "cloudNetworkQuery":
		return o.CloudNetworkQuery
	case "disabled":
		return o.Disabled
	case "lastExecutionTimestamp":
		return o.LastExecutionTimestamp
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "policyID":
		return o.PolicyID
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudScheduledNetworkQueryAttributesMap represents the map of attribute for CloudScheduledNetworkQuery.
var CloudScheduledNetworkQueryAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AlertRuleID": {
		AllowedChoices: []string{},
		BSONFieldName:  "alertruleid",
		ConvertedName:  "AlertRuleID",
		Description:    `Prisma Cloud Alert Rule ID.`,
		Exposed:        true,
		Name:           "alertRuleID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"CloudGraphResult": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudGraphResult",
		Description:    `The result of the cloud network query.`,
		Exposed:        true,
		Name:           "cloudGraphResult",
		SubType:        "cloudgraph",
		Type:           "ref",
	},
	"CloudNetworkQuery": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudnetworkquery",
		ConvertedName:  "CloudNetworkQuery",
		Description:    `The cloud network query that should be used.`,
		Exposed:        true,
		Name:           "cloudNetworkQuery",
		Stored:         true,
		SubType:        "cloudnetworkquery",
		Type:           "ref",
	},
	"Disabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "disabled",
		ConvertedName:  "Disabled",
		Description:    `Represents whether the associated policy was disabled.`,
		Exposed:        true,
		Name:           "disabled",
		Stored:         true,
		Type:           "boolean",
	},
	"LastExecutionTimestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastexecutiontimestamp",
		ConvertedName:  "LastExecutionTimestamp",
		Description:    `Result of the last successfully run query.`,
		Exposed:        true,
		Name:           "lastExecutionTimestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"PolicyID": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyid",
		ConvertedName:  "PolicyID",
		Description:    `Prisma Cloud Policy ID.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
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

// CloudScheduledNetworkQueryLowerCaseAttributesMap represents the map of attribute for CloudScheduledNetworkQuery.
var CloudScheduledNetworkQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"alertruleid": {
		AllowedChoices: []string{},
		BSONFieldName:  "alertruleid",
		ConvertedName:  "AlertRuleID",
		Description:    `Prisma Cloud Alert Rule ID.`,
		Exposed:        true,
		Name:           "alertRuleID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"cloudgraphresult": {
		AllowedChoices: []string{},
		ConvertedName:  "CloudGraphResult",
		Description:    `The result of the cloud network query.`,
		Exposed:        true,
		Name:           "cloudGraphResult",
		SubType:        "cloudgraph",
		Type:           "ref",
	},
	"cloudnetworkquery": {
		AllowedChoices: []string{},
		BSONFieldName:  "cloudnetworkquery",
		ConvertedName:  "CloudNetworkQuery",
		Description:    `The cloud network query that should be used.`,
		Exposed:        true,
		Name:           "cloudNetworkQuery",
		Stored:         true,
		SubType:        "cloudnetworkquery",
		Type:           "ref",
	},
	"disabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "disabled",
		ConvertedName:  "Disabled",
		Description:    `Represents whether the associated policy was disabled.`,
		Exposed:        true,
		Name:           "disabled",
		Stored:         true,
		Type:           "boolean",
	},
	"lastexecutiontimestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastexecutiontimestamp",
		ConvertedName:  "LastExecutionTimestamp",
		Description:    `Result of the last successfully run query.`,
		Exposed:        true,
		Name:           "lastExecutionTimestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"policyid": {
		AllowedChoices: []string{},
		BSONFieldName:  "policyid",
		ConvertedName:  "PolicyID",
		Description:    `Prisma Cloud Policy ID.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
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

// SparseCloudScheduledNetworkQueriesList represents a list of SparseCloudScheduledNetworkQueries
type SparseCloudScheduledNetworkQueriesList []*SparseCloudScheduledNetworkQuery

// Identity returns the identity of the objects in the list.
func (o SparseCloudScheduledNetworkQueriesList) Identity() elemental.Identity {

	return CloudScheduledNetworkQueryIdentity
}

// Copy returns a pointer to a copy the SparseCloudScheduledNetworkQueriesList.
func (o SparseCloudScheduledNetworkQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudScheduledNetworkQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudScheduledNetworkQueriesList.
func (o SparseCloudScheduledNetworkQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudScheduledNetworkQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudScheduledNetworkQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudScheduledNetworkQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudScheduledNetworkQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCloudScheduledNetworkQueriesList converted to CloudScheduledNetworkQueriesList.
func (o SparseCloudScheduledNetworkQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudScheduledNetworkQueriesList) Version() int {

	return 1
}

// SparseCloudScheduledNetworkQuery represents the sparse version of a cloudschedulednetworkquery.
type SparseCloudScheduledNetworkQuery struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Prisma Cloud Alert Rule ID.
	AlertRuleID *string `json:"alertRuleID,omitempty" msgpack:"alertRuleID,omitempty" bson:"alertruleid,omitempty" mapstructure:"alertRuleID,omitempty"`

	// The result of the cloud network query.
	CloudGraphResult *CloudGraph `json:"cloudGraphResult,omitempty" msgpack:"cloudGraphResult,omitempty" bson:"-" mapstructure:"cloudGraphResult,omitempty"`

	// The cloud network query that should be used.
	CloudNetworkQuery *CloudNetworkQuery `json:"cloudNetworkQuery,omitempty" msgpack:"cloudNetworkQuery,omitempty" bson:"cloudnetworkquery,omitempty" mapstructure:"cloudNetworkQuery,omitempty"`

	// Represents whether the associated policy was disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Result of the last successfully run query.
	LastExecutionTimestamp *time.Time `json:"lastExecutionTimestamp,omitempty" msgpack:"lastExecutionTimestamp,omitempty" bson:"lastexecutiontimestamp,omitempty" mapstructure:"lastExecutionTimestamp,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Prisma Cloud Policy ID.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudScheduledNetworkQuery returns a new  SparseCloudScheduledNetworkQuery.
func NewSparseCloudScheduledNetworkQuery() *SparseCloudScheduledNetworkQuery {
	return &SparseCloudScheduledNetworkQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudScheduledNetworkQuery) Identity() elemental.Identity {

	return CloudScheduledNetworkQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudScheduledNetworkQuery) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudScheduledNetworkQuery) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudScheduledNetworkQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudScheduledNetworkQuery{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.AlertRuleID != nil {
		s.AlertRuleID = o.AlertRuleID
	}
	if o.CloudNetworkQuery != nil {
		s.CloudNetworkQuery = o.CloudNetworkQuery
	}
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.LastExecutionTimestamp != nil {
		s.LastExecutionTimestamp = o.LastExecutionTimestamp
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
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
func (o *SparseCloudScheduledNetworkQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudScheduledNetworkQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.AlertRuleID != nil {
		o.AlertRuleID = s.AlertRuleID
	}
	if s.CloudNetworkQuery != nil {
		o.CloudNetworkQuery = s.CloudNetworkQuery
	}
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.LastExecutionTimestamp != nil {
		o.LastExecutionTimestamp = s.LastExecutionTimestamp
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
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
func (o *SparseCloudScheduledNetworkQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudScheduledNetworkQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudScheduledNetworkQuery()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AlertRuleID != nil {
		out.AlertRuleID = *o.AlertRuleID
	}
	if o.CloudGraphResult != nil {
		out.CloudGraphResult = o.CloudGraphResult
	}
	if o.CloudNetworkQuery != nil {
		out.CloudNetworkQuery = o.CloudNetworkQuery
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.LastExecutionTimestamp != nil {
		out.LastExecutionTimestamp = *o.LastExecutionTimestamp
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudScheduledNetworkQuery) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudScheduledNetworkQuery) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudScheduledNetworkQuery) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudScheduledNetworkQuery) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudScheduledNetworkQuery) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudScheduledNetworkQuery) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudScheduledNetworkQuery) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudScheduledNetworkQuery) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudScheduledNetworkQuery.
func (o *SparseCloudScheduledNetworkQuery) DeepCopy() *SparseCloudScheduledNetworkQuery {

	if o == nil {
		return nil
	}

	out := &SparseCloudScheduledNetworkQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudScheduledNetworkQuery.
func (o *SparseCloudScheduledNetworkQuery) DeepCopyInto(out *SparseCloudScheduledNetworkQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudScheduledNetworkQuery: %s", err))
	}

	*out = *target.(*SparseCloudScheduledNetworkQuery)
}

type mongoAttributesCloudScheduledNetworkQuery struct {
	ID                     bson.ObjectId      `bson:"_id,omitempty"`
	AlertRuleID            string             `bson:"alertruleid"`
	CloudNetworkQuery      *CloudNetworkQuery `bson:"cloudnetworkquery"`
	Disabled               bool               `bson:"disabled"`
	LastExecutionTimestamp time.Time          `bson:"lastexecutiontimestamp"`
	MigrationsLog          map[string]string  `bson:"migrationslog,omitempty"`
	Namespace              string             `bson:"namespace"`
	PolicyID               string             `bson:"policyid"`
	ZHash                  int                `bson:"zhash"`
	Zone                   int                `bson:"zone"`
}
type mongoAttributesSparseCloudScheduledNetworkQuery struct {
	ID                     bson.ObjectId      `bson:"_id,omitempty"`
	AlertRuleID            *string            `bson:"alertruleid,omitempty"`
	CloudNetworkQuery      *CloudNetworkQuery `bson:"cloudnetworkquery,omitempty"`
	Disabled               *bool              `bson:"disabled,omitempty"`
	LastExecutionTimestamp *time.Time         `bson:"lastexecutiontimestamp,omitempty"`
	MigrationsLog          *map[string]string `bson:"migrationslog,omitempty"`
	Namespace              *string            `bson:"namespace,omitempty"`
	PolicyID               *string            `bson:"policyid,omitempty"`
	ZHash                  *int               `bson:"zhash,omitempty"`
	Zone                   *int               `bson:"zone,omitempty"`
}
