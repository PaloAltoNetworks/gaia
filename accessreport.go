package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AccessReportActionValue represents the possible values for attribute "action".
type AccessReportActionValue string

const (
	// AccessReportActionAccept represents the value Accept.
	AccessReportActionAccept AccessReportActionValue = "Accept"

	// AccessReportActionReject represents the value Reject.
	AccessReportActionReject AccessReportActionValue = "Reject"
)

// AccessReportTypeValue represents the possible values for attribute "type".
type AccessReportTypeValue string

const (
	// AccessReportTypeSSHLogin represents the value SSHLogin.
	AccessReportTypeSSHLogin AccessReportTypeValue = "SSHLogin"

	// AccessReportTypeSSHLogout represents the value SSHLogout.
	AccessReportTypeSSHLogout AccessReportTypeValue = "SSHLogout"

	// AccessReportTypeSudoEnter represents the value SudoEnter.
	AccessReportTypeSudoEnter AccessReportTypeValue = "SudoEnter"

	// AccessReportTypeSudoExit represents the value SudoExit.
	AccessReportTypeSudoExit AccessReportTypeValue = "SudoExit"
)

// AccessReportIdentity represents the Identity of the object.
var AccessReportIdentity = elemental.Identity{
	Name:     "accessreport",
	Category: "accessreports",
	Package:  "zack",
	Private:  false,
}

// AccessReportsList represents a list of AccessReports
type AccessReportsList []*AccessReport

// Identity returns the identity of the objects in the list.
func (o AccessReportsList) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Copy returns a pointer to a copy the AccessReportsList.
func (o AccessReportsList) Copy() elemental.Identifiables {

	copy := append(AccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccessReportsList.
func (o AccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AccessReportsList converted to SparseAccessReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccessReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAccessReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAccessReport)
	}

	return out
}

// Version returns the version of the content.
func (o AccessReportsList) Version() int {

	return 1
}

// AccessReport represents the model of a accessreport
type AccessReport struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Action applied to the access.
	Action AccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"a,omitempty" mapstructure:"action,omitempty"`

	// Hash of the claims used to communicate.
	ClaimHash string `json:"claimHash,omitempty" msgpack:"claimHash,omitempty" bson:"b,omitempty" mapstructure:"claimHash,omitempty"`

	// Identifier of the enforcer.
	EnforcerID string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"c,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"d,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// ID of the processing unit of the report.
	ProcessingUnitID string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"e,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Name of the processing unit of the report.
	ProcessingUnitName string `json:"processingUnitName,omitempty" msgpack:"processingUnitName,omitempty" bson:"f,omitempty" mapstructure:"processingUnitName,omitempty"`

	// Namespace of the processing unit of the report.
	ProcessingUnitNamespace string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"g,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	Reason string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"h,omitempty" mapstructure:"reason,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"i,omitempty" mapstructure:"timestamp,omitempty"`

	// Type of the report.
	Type AccessReportTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"j,omitempty" mapstructure:"type,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAccessReport returns a new *AccessReport
func NewAccessReport() *AccessReport {

	return &AccessReport{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *AccessReport) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AccessReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AccessReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AccessReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAccessReport{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Action = o.Action
	s.ClaimHash = o.ClaimHash
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.MigrationsLog = o.MigrationsLog
	s.ProcessingUnitID = o.ProcessingUnitID
	s.ProcessingUnitName = o.ProcessingUnitName
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.Reason = o.Reason
	s.Timestamp = o.Timestamp
	s.Type = o.Type
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AccessReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAccessReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Action = s.Action
	o.ClaimHash = s.ClaimHash
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.MigrationsLog = s.MigrationsLog
	o.ProcessingUnitID = s.ProcessingUnitID
	o.ProcessingUnitName = s.ProcessingUnitName
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.Reason = s.Reason
	o.Timestamp = s.Timestamp
	o.Type = s.Type
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AccessReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AccessReport) BleveType() string {

	return "accessreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AccessReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AccessReport) Doc() string {

	return `Represents any access made by the user.`
}

func (o *AccessReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *AccessReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *AccessReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *AccessReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AccessReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AccessReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AccessReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AccessReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccessReport{
			ID:                      &o.ID,
			Action:                  &o.Action,
			ClaimHash:               &o.ClaimHash,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			MigrationsLog:           &o.MigrationsLog,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitName:      &o.ProcessingUnitName,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Reason:                  &o.Reason,
			Timestamp:               &o.Timestamp,
			Type:                    &o.Type,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparseAccessReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "action":
			sp.Action = &(o.Action)
		case "claimHash":
			sp.ClaimHash = &(o.ClaimHash)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitName":
			sp.ProcessingUnitName = &(o.ProcessingUnitName)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "reason":
			sp.Reason = &(o.Reason)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "type":
			sp.Type = &(o.Type)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAccessReport to the object.
func (o *AccessReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAccessReport)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.ClaimHash != nil {
		o.ClaimHash = *so.ClaimHash
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitName != nil {
		o.ProcessingUnitName = *so.ProcessingUnitName
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Reason != nil {
		o.Reason = *so.Reason
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the AccessReport.
func (o *AccessReport) DeepCopy() *AccessReport {

	if o == nil {
		return nil
	}

	out := &AccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AccessReport.
func (o *AccessReport) DeepCopyInto(out *AccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AccessReport: %s", err))
	}

	*out = *target.(*AccessReport)
}

// Validate valides the current information stored into the structure.
func (o *AccessReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("type", string(o.Type)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"SSHLogin", "SSHLogout", "SudoEnter", "SudoExit"}, false); err != nil {
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
func (*AccessReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccessReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccessReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AccessReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccessReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AccessReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "claimHash":
		return o.ClaimHash
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "migrationsLog":
		return o.MigrationsLog
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitName":
		return o.ProcessingUnitName
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "reason":
		return o.Reason
	case "timestamp":
		return o.Timestamp
	case "type":
		return o.Type
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AccessReportAttributesMap represents the map of attribute for AccessReport.
var AccessReportAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Action": {
		AllowedChoices: []string{"Accept", "Reject"},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		Description:    `Action applied to the access.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"ClaimHash": {
		AllowedChoices: []string{},
		BSONFieldName:  "claimhash",
		ConvertedName:  "ClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "claimHash",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "enforcerid",
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "enforcernamespace",
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		BSONFieldName:  "processingunitid",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitName": {
		AllowedChoices: []string{},
		BSONFieldName:  "processingunitname",
		ConvertedName:  "ProcessingUnitName",
		Description:    `Name of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitName",
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "processingunitnamespace",
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Reason": {
		AllowedChoices: []string{},
		BSONFieldName:  "reason",
		ConvertedName:  "Reason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "reason",
		Stored:  true,
		Type:    "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "timestamp",
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"Type": {
		AllowedChoices: []string{"SSHLogin", "SSHLogout", "SudoEnter", "SudoExit"},
		BSONFieldName:  "type",
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
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

// AccessReportLowerCaseAttributesMap represents the map of attribute for AccessReport.
var AccessReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"action": {
		AllowedChoices: []string{"Accept", "Reject"},
		BSONFieldName:  "a",
		ConvertedName:  "Action",
		Description:    `Action applied to the access.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"claimhash": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "ClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "claimHash",
		Stored:         true,
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"processingunitid": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"processingunitname": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "ProcessingUnitName",
		Description:    `Name of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitName",
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit of the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"reason": {
		AllowedChoices: []string{},
		BSONFieldName:  "h",
		ConvertedName:  "Reason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "reason",
		Stored:  true,
		Type:    "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "i",
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"type": {
		AllowedChoices: []string{"SSHLogin", "SSHLogout", "SudoEnter", "SudoExit"},
		BSONFieldName:  "j",
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Stored:         true,
		Type:           "enum",
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

// SparseAccessReportsList represents a list of SparseAccessReports
type SparseAccessReportsList []*SparseAccessReport

// Identity returns the identity of the objects in the list.
func (o SparseAccessReportsList) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Copy returns a pointer to a copy the SparseAccessReportsList.
func (o SparseAccessReportsList) Copy() elemental.Identifiables {

	copy := append(SparseAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccessReportsList.
func (o SparseAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAccessReportsList converted to AccessReportsList.
func (o SparseAccessReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccessReportsList) Version() int {

	return 1
}

// SparseAccessReport represents the sparse version of a accessreport.
type SparseAccessReport struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Action applied to the access.
	Action *AccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"a,omitempty" mapstructure:"action,omitempty"`

	// Hash of the claims used to communicate.
	ClaimHash *string `json:"claimHash,omitempty" msgpack:"claimHash,omitempty" bson:"b,omitempty" mapstructure:"claimHash,omitempty"`

	// Identifier of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"c,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"d,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// ID of the processing unit of the report.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"e,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Name of the processing unit of the report.
	ProcessingUnitName *string `json:"processingUnitName,omitempty" msgpack:"processingUnitName,omitempty" bson:"f,omitempty" mapstructure:"processingUnitName,omitempty"`

	// Namespace of the processing unit of the report.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"g,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	Reason *string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"h,omitempty" mapstructure:"reason,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"i,omitempty" mapstructure:"timestamp,omitempty"`

	// Type of the report.
	Type *AccessReportTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"j,omitempty" mapstructure:"type,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAccessReport returns a new  SparseAccessReport.
func NewSparseAccessReport() *SparseAccessReport {
	return &SparseAccessReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccessReport) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccessReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccessReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccessReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAccessReport{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Action != nil {
		s.Action = o.Action
	}
	if o.ClaimHash != nil {
		s.ClaimHash = o.ClaimHash
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.ProcessingUnitID != nil {
		s.ProcessingUnitID = o.ProcessingUnitID
	}
	if o.ProcessingUnitName != nil {
		s.ProcessingUnitName = o.ProcessingUnitName
	}
	if o.ProcessingUnitNamespace != nil {
		s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	}
	if o.Reason != nil {
		s.Reason = o.Reason
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
	}
	if o.Type != nil {
		s.Type = o.Type
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
func (o *SparseAccessReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAccessReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Action != nil {
		o.Action = s.Action
	}
	if s.ClaimHash != nil {
		o.ClaimHash = s.ClaimHash
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.ProcessingUnitID != nil {
		o.ProcessingUnitID = s.ProcessingUnitID
	}
	if s.ProcessingUnitName != nil {
		o.ProcessingUnitName = s.ProcessingUnitName
	}
	if s.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	}
	if s.Reason != nil {
		o.Reason = s.Reason
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
	}
	if s.Type != nil {
		o.Type = s.Type
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
func (o *SparseAccessReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccessReport) ToPlain() elemental.PlainIdentifiable {

	out := NewAccessReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.ClaimHash != nil {
		out.ClaimHash = *o.ClaimHash
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitName != nil {
		out.ProcessingUnitName = *o.ProcessingUnitName
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Reason != nil {
		out.Reason = *o.Reason
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Type != nil {
		out.Type = *o.Type
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
func (o *SparseAccessReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAccessReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAccessReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAccessReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAccessReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAccessReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAccessReport.
func (o *SparseAccessReport) DeepCopy() *SparseAccessReport {

	if o == nil {
		return nil
	}

	out := &SparseAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccessReport.
func (o *SparseAccessReport) DeepCopyInto(out *SparseAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccessReport: %s", err))
	}

	*out = *target.(*SparseAccessReport)
}

type mongoAttributesAccessReport struct {
	ID                      bson.ObjectId           `bson:"_id,omitempty"`
	Action                  AccessReportActionValue `bson:"a,omitempty"`
	ClaimHash               string                  `bson:"b,omitempty"`
	EnforcerID              string                  `bson:"c,omitempty"`
	EnforcerNamespace       string                  `bson:"d,omitempty"`
	MigrationsLog           map[string]string       `bson:"migrationslog,omitempty"`
	ProcessingUnitID        string                  `bson:"e,omitempty"`
	ProcessingUnitName      string                  `bson:"f,omitempty"`
	ProcessingUnitNamespace string                  `bson:"g,omitempty"`
	Reason                  string                  `bson:"h,omitempty"`
	Timestamp               time.Time               `bson:"i,omitempty"`
	Type                    AccessReportTypeValue   `bson:"j,omitempty"`
	ZHash                   int                     `bson:"zhash"`
	Zone                    int                     `bson:"zone"`
}
type mongoAttributesSparseAccessReport struct {
	ID                      bson.ObjectId            `bson:"_id,omitempty"`
	Action                  *AccessReportActionValue `bson:"a,omitempty"`
	ClaimHash               *string                  `bson:"b,omitempty"`
	EnforcerID              *string                  `bson:"c,omitempty"`
	EnforcerNamespace       *string                  `bson:"d,omitempty"`
	MigrationsLog           *map[string]string       `bson:"migrationslog,omitempty"`
	ProcessingUnitID        *string                  `bson:"e,omitempty"`
	ProcessingUnitName      *string                  `bson:"f,omitempty"`
	ProcessingUnitNamespace *string                  `bson:"g,omitempty"`
	Reason                  *string                  `bson:"h,omitempty"`
	Timestamp               *time.Time               `bson:"i,omitempty"`
	Type                    *AccessReportTypeValue   `bson:"j,omitempty"`
	ZHash                   *int                     `bson:"zhash,omitempty"`
	Zone                    *int                     `bson:"zone,omitempty"`
}
