package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingResultIdentity represents the Identity of the object.
var PingResultIdentity = elemental.Identity{
	Name:     "pingresult",
	Category: "pingresults",
	Package:  "guy",
	Private:  false,
}

// PingResultsList represents a list of PingResults
type PingResultsList []*PingResult

// Identity returns the identity of the objects in the list.
func (o PingResultsList) Identity() elemental.Identity {

	return PingResultIdentity
}

// Copy returns a pointer to a copy the PingResultsList.
func (o PingResultsList) Copy() elemental.Identifiables {

	copy := append(PingResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PingResultsList.
func (o PingResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PingResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PingResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PingResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PingResultsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PingResultsList converted to SparsePingResultsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PingResultsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePingResultsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePingResult)
	}

	return out
}

// Version returns the version of the content.
func (o PingResultsList) Version() int {

	return 1
}

// PingResult represents the model of a pingresult
type PingResult struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// May contain a list of errors that have happened during the collection.
	Errors []string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"errors" mapstructure:"errors,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the Ping ID.
	PingID string `json:"pingID" msgpack:"pingID" bson:"pingid" mapstructure:"pingID,omitempty"`

	// Contains the result of aggregated ping pairs.
	PingPairs []*PingPair `json:"pingPairs" msgpack:"pingPairs" bson:"pingpairs" mapstructure:"pingPairs,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID string `json:"refreshID" msgpack:"refreshID" bson:"refreshid" mapstructure:"refreshID,omitempty"`

	// Contains information about missing probes in the result. This field will be
	// populated in the ping probe is managed by a remote controller (federation) or is
	// stored in a namespace you don't have any permissions on.
	RemoteProbes []*RemotePingProbe `json:"remoteProbes,omitempty" msgpack:"remoteProbes,omitempty" bson:"remoteprobes" mapstructure:"remoteProbes,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingResult returns a new *PingResult
func NewPingResult() *PingResult {

	return &PingResult{
		ModelVersion:  1,
		Errors:        []string{},
		MigrationsLog: map[string]string{},
		PingPairs:     []*PingPair{},
	}
}

// Identity returns the Identity of the object.
func (o *PingResult) Identity() elemental.Identity {

	return PingResultIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PingResult) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PingResult) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingResult{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreateTime = o.CreateTime
	s.Errors = o.Errors
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.PingID = o.PingID
	s.PingPairs = o.PingPairs
	s.RefreshID = o.RefreshID
	s.RemoteProbes = o.RemoteProbes
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.CreateTime = s.CreateTime
	o.Errors = s.Errors
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.PingID = s.PingID
	o.PingPairs = s.PingPairs
	o.RefreshID = s.RefreshID
	o.RemoteProbes = s.RemoteProbes
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PingResult) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingResult) BleveType() string {

	return "pingresult"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PingResult) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PingResult) Doc() string {

	return `Represents the results of a ping request.`
}

func (o *PingResult) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *PingResult) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *PingResult) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *PingResult) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *PingResult) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *PingResult) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *PingResult) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *PingResult) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *PingResult) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *PingResult) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *PingResult) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *PingResult) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *PingResult) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PingResult) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePingResult{
			ID:            &o.ID,
			CreateTime:    &o.CreateTime,
			Errors:        &o.Errors,
			MigrationsLog: &o.MigrationsLog,
			Namespace:     &o.Namespace,
			PingID:        &o.PingID,
			PingPairs:     &o.PingPairs,
			RefreshID:     &o.RefreshID,
			RemoteProbes:  &o.RemoteProbes,
			UpdateTime:    &o.UpdateTime,
			ZHash:         &o.ZHash,
			Zone:          &o.Zone,
		}
	}

	sp := &SparsePingResult{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "errors":
			sp.Errors = &(o.Errors)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "pingID":
			sp.PingID = &(o.PingID)
		case "pingPairs":
			sp.PingPairs = &(o.PingPairs)
		case "refreshID":
			sp.RefreshID = &(o.RefreshID)
		case "remoteProbes":
			sp.RemoteProbes = &(o.RemoteProbes)
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

// Patch apply the non nil value of a *SparsePingResult to the object.
func (o *PingResult) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePingResult)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Errors != nil {
		o.Errors = *so.Errors
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PingID != nil {
		o.PingID = *so.PingID
	}
	if so.PingPairs != nil {
		o.PingPairs = *so.PingPairs
	}
	if so.RefreshID != nil {
		o.RefreshID = *so.RefreshID
	}
	if so.RemoteProbes != nil {
		o.RemoteProbes = *so.RemoteProbes
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

// DeepCopy returns a deep copy if the PingResult.
func (o *PingResult) DeepCopy() *PingResult {

	if o == nil {
		return nil
	}

	out := &PingResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingResult.
func (o *PingResult) DeepCopyInto(out *PingResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingResult: %s", err))
	}

	*out = *target.(*PingResult)
}

// Validate valides the current information stored into the structure.
func (o *PingResult) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.PingPairs {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.RemoteProbes {
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
func (*PingResult) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingResultAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingResultLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PingResult) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingResultAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PingResult) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "createTime":
		return o.CreateTime
	case "errors":
		return o.Errors
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "pingID":
		return o.PingID
	case "pingPairs":
		return o.PingPairs
	case "refreshID":
		return o.RefreshID
	case "remoteProbes":
		return o.RemoteProbes
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// PingResultAttributesMap represents the map of attribute for PingResult.
var PingResultAttributesMap = map[string]elemental.AttributeSpecification{
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
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `May contain a list of errors that have happened during the collection.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
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
	"PingID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PingID",
		Description:    `Contains the Ping ID.`,
		Exposed:        true,
		Name:           "pingID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"PingPairs": {
		AllowedChoices: []string{},
		ConvertedName:  "PingPairs",
		Description:    `Contains the result of aggregated ping pairs.`,
		Exposed:        true,
		Name:           "pingPairs",
		Stored:         true,
		SubType:        "pingpair",
		Type:           "refList",
	},
	"RefreshID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RefreshID",
		Description:    `Contains the refresh ID set by processing unit refresh event.`,
		Exposed:        true,
		Name:           "refreshID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"RemoteProbes": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteProbes",
		Description: `Contains information about missing probes in the result. This field will be
populated in the ping probe is managed by a remote controller (federation) or is
stored in a namespace you don't have any permissions on.`,
		Exposed: true,
		Name:    "remoteProbes",
		Stored:  true,
		SubType: "remotepingprobe",
		Type:    "refList",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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

// PingResultLowerCaseAttributesMap represents the map of attribute for PingResult.
var PingResultLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `May contain a list of errors that have happened during the collection.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"migrationslog": {
		AllowedChoices: []string{},
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
	"pingid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PingID",
		Description:    `Contains the Ping ID.`,
		Exposed:        true,
		Name:           "pingID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"pingpairs": {
		AllowedChoices: []string{},
		ConvertedName:  "PingPairs",
		Description:    `Contains the result of aggregated ping pairs.`,
		Exposed:        true,
		Name:           "pingPairs",
		Stored:         true,
		SubType:        "pingpair",
		Type:           "refList",
	},
	"refreshid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RefreshID",
		Description:    `Contains the refresh ID set by processing unit refresh event.`,
		Exposed:        true,
		Name:           "refreshID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"remoteprobes": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteProbes",
		Description: `Contains information about missing probes in the result. This field will be
populated in the ping probe is managed by a remote controller (federation) or is
stored in a namespace you don't have any permissions on.`,
		Exposed: true,
		Name:    "remoteProbes",
		Stored:  true,
		SubType: "remotepingprobe",
		Type:    "refList",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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

// SparsePingResultsList represents a list of SparsePingResults
type SparsePingResultsList []*SparsePingResult

// Identity returns the identity of the objects in the list.
func (o SparsePingResultsList) Identity() elemental.Identity {

	return PingResultIdentity
}

// Copy returns a pointer to a copy the SparsePingResultsList.
func (o SparsePingResultsList) Copy() elemental.Identifiables {

	copy := append(SparsePingResultsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePingResultsList.
func (o SparsePingResultsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePingResultsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePingResult))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePingResultsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePingResultsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePingResultsList converted to PingResultsList.
func (o SparsePingResultsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePingResultsList) Version() int {

	return 1
}

// SparsePingResult represents the sparse version of a pingresult.
type SparsePingResult struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// May contain a list of errors that have happened during the collection.
	Errors *[]string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"errors,omitempty" mapstructure:"errors,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the Ping ID.
	PingID *string `json:"pingID,omitempty" msgpack:"pingID,omitempty" bson:"pingid,omitempty" mapstructure:"pingID,omitempty"`

	// Contains the result of aggregated ping pairs.
	PingPairs *[]*PingPair `json:"pingPairs,omitempty" msgpack:"pingPairs,omitempty" bson:"pingpairs,omitempty" mapstructure:"pingPairs,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID *string `json:"refreshID,omitempty" msgpack:"refreshID,omitempty" bson:"refreshid,omitempty" mapstructure:"refreshID,omitempty"`

	// Contains information about missing probes in the result. This field will be
	// populated in the ping probe is managed by a remote controller (federation) or is
	// stored in a namespace you don't have any permissions on.
	RemoteProbes *[]*RemotePingProbe `json:"remoteProbes,omitempty" msgpack:"remoteProbes,omitempty" bson:"remoteprobes,omitempty" mapstructure:"remoteProbes,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePingResult returns a new  SparsePingResult.
func NewSparsePingResult() *SparsePingResult {
	return &SparsePingResult{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePingResult) Identity() elemental.Identity {

	return PingResultIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePingResult) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePingResult) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingResult) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePingResult{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Errors != nil {
		s.Errors = o.Errors
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PingID != nil {
		s.PingID = o.PingID
	}
	if o.PingPairs != nil {
		s.PingPairs = o.PingPairs
	}
	if o.RefreshID != nil {
		s.RefreshID = o.RefreshID
	}
	if o.RemoteProbes != nil {
		s.RemoteProbes = o.RemoteProbes
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
func (o *SparsePingResult) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePingResult{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Errors != nil {
		o.Errors = s.Errors
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PingID != nil {
		o.PingID = s.PingID
	}
	if s.PingPairs != nil {
		o.PingPairs = s.PingPairs
	}
	if s.RefreshID != nil {
		o.RefreshID = s.RefreshID
	}
	if s.RemoteProbes != nil {
		o.RemoteProbes = s.RemoteProbes
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
func (o *SparsePingResult) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePingResult) ToPlain() elemental.PlainIdentifiable {

	out := NewPingResult()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Errors != nil {
		out.Errors = *o.Errors
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PingID != nil {
		out.PingID = *o.PingID
	}
	if o.PingPairs != nil {
		out.PingPairs = *o.PingPairs
	}
	if o.RefreshID != nil {
		out.RefreshID = *o.RefreshID
	}
	if o.RemoteProbes != nil {
		out.RemoteProbes = *o.RemoteProbes
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

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparsePingResult) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparsePingResult) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparsePingResult) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparsePingResult) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePingResult) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePingResult) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparsePingResult) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparsePingResult) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePingResult) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePingResult) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePingResult) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePingResult) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparsePingResult.
func (o *SparsePingResult) DeepCopy() *SparsePingResult {

	if o == nil {
		return nil
	}

	out := &SparsePingResult{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePingResult.
func (o *SparsePingResult) DeepCopyInto(out *SparsePingResult) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePingResult: %s", err))
	}

	*out = *target.(*SparsePingResult)
}

type mongoAttributesPingResult struct {
	ID            bson.ObjectId      `bson:"_id,omitempty"`
	CreateTime    time.Time          `bson:"createtime"`
	Errors        []string           `bson:"errors,omitempty"`
	MigrationsLog map[string]string  `bson:"migrationslog,omitempty"`
	Namespace     string             `bson:"namespace"`
	PingID        string             `bson:"pingid"`
	PingPairs     []*PingPair        `bson:"pingpairs"`
	RefreshID     string             `bson:"refreshid"`
	RemoteProbes  []*RemotePingProbe `bson:"remoteprobes,omitempty"`
	UpdateTime    time.Time          `bson:"updatetime"`
	ZHash         int                `bson:"zhash"`
	Zone          int                `bson:"zone"`
}
type mongoAttributesSparsePingResult struct {
	ID            bson.ObjectId       `bson:"_id,omitempty"`
	CreateTime    *time.Time          `bson:"createtime,omitempty"`
	Errors        *[]string           `bson:"errors,omitempty"`
	MigrationsLog *map[string]string  `bson:"migrationslog,omitempty"`
	Namespace     *string             `bson:"namespace,omitempty"`
	PingID        *string             `bson:"pingid,omitempty"`
	PingPairs     *[]*PingPair        `bson:"pingpairs,omitempty"`
	RefreshID     *string             `bson:"refreshid,omitempty"`
	RemoteProbes  *[]*RemotePingProbe `bson:"remoteprobes,omitempty"`
	UpdateTime    *time.Time          `bson:"updatetime,omitempty"`
	ZHash         *int                `bson:"zhash,omitempty"`
	Zone          *int                `bson:"zone,omitempty"`
}
