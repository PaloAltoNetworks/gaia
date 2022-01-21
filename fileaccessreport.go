package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FileAccessReportActionValue represents the possible values for attribute "action".
type FileAccessReportActionValue string

const (
	// FileAccessReportActionAccept represents the value Accept.
	FileAccessReportActionAccept FileAccessReportActionValue = "Accept"

	// FileAccessReportActionLimit represents the value Limit.
	FileAccessReportActionLimit FileAccessReportActionValue = "Limit"

	// FileAccessReportActionReject represents the value Reject.
	FileAccessReportActionReject FileAccessReportActionValue = "Reject"
)

// FileAccessReportIdentity represents the Identity of the object.
var FileAccessReportIdentity = elemental.Identity{
	Name:     "fileaccessreport",
	Category: "fileaccessreports",
	Package:  "zack",
	Private:  false,
}

// FileAccessReportsList represents a list of FileAccessReports
type FileAccessReportsList []*FileAccessReport

// Identity returns the identity of the objects in the list.
func (o FileAccessReportsList) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Copy returns a pointer to a copy the FileAccessReportsList.
func (o FileAccessReportsList) Copy() elemental.Identifiables {

	copy := append(FileAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FileAccessReportsList.
func (o FileAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FileAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FileAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FileAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FileAccessReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToSparse returns the FileAccessReportsList converted to SparseFileAccessReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FileAccessReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFileAccessReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFileAccessReport)
	}

	return out
}

// Version returns the version of the content.
func (o FileAccessReportsList) Version() int {

	return 1
}

// FileAccessReport represents the model of a fileaccessreport
type FileAccessReport struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Action taken.
	Action FileAccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"a,omitempty" mapstructure:"action,omitempty"`

	// Host storing the file.
	Host string `json:"host,omitempty" msgpack:"host,omitempty" bson:"b,omitempty" mapstructure:"host,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Mode of file access.
	Mode string `json:"mode,omitempty" msgpack:"mode,omitempty" bson:"c,omitempty" mapstructure:"mode,omitempty"`

	// Path of the file.
	Path string `json:"path,omitempty" msgpack:"path,omitempty" bson:"d,omitempty" mapstructure:"path,omitempty"`

	// ID of the processing unit.
	ProcessingUnitID string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"e,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit.
	ProcessingUnitNamespace string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"f,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"g,omitempty" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFileAccessReport returns a new *FileAccessReport
func NewFileAccessReport() *FileAccessReport {

	return &FileAccessReport{
		ModelVersion:  1,
		Host:          "localhost",
		MigrationsLog: map[string]string{},
		Mode:          "rxw",
		Path:          "/etc/passwd",
	}
}

// Identity returns the Identity of the object.
func (o *FileAccessReport) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FileAccessReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FileAccessReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FileAccessReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFileAccessReport{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Action = o.Action
	s.Host = o.Host
	s.MigrationsLog = o.MigrationsLog
	s.Mode = o.Mode
	s.Path = o.Path
	s.ProcessingUnitID = o.ProcessingUnitID
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.Timestamp = o.Timestamp
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FileAccessReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFileAccessReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Action = s.Action
	o.Host = s.Host
	o.MigrationsLog = s.MigrationsLog
	o.Mode = s.Mode
	o.Path = s.Path
	o.ProcessingUnitID = s.ProcessingUnitID
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.Timestamp = s.Timestamp
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FileAccessReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FileAccessReport) BleveType() string {

	return "fileaccessreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FileAccessReport) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// Doc returns the documentation for the object
func (o *FileAccessReport) Doc() string {

	return `Post a new file access report.`
}

func (o *FileAccessReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *FileAccessReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *FileAccessReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *FileAccessReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *FileAccessReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *FileAccessReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *FileAccessReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FileAccessReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFileAccessReport{
			ID:                      &o.ID,
			Action:                  &o.Action,
			Host:                    &o.Host,
			MigrationsLog:           &o.MigrationsLog,
			Mode:                    &o.Mode,
			Path:                    &o.Path,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Timestamp:               &o.Timestamp,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparseFileAccessReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "action":
			sp.Action = &(o.Action)
		case "host":
			sp.Host = &(o.Host)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "mode":
			sp.Mode = &(o.Mode)
		case "path":
			sp.Path = &(o.Path)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFileAccessReport to the object.
func (o *FileAccessReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFileAccessReport)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.Host != nil {
		o.Host = *so.Host
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
	if so.Path != nil {
		o.Path = *so.Path
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the FileAccessReport.
func (o *FileAccessReport) DeepCopy() *FileAccessReport {

	if o == nil {
		return nil
	}

	out := &FileAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FileAccessReport.
func (o *FileAccessReport) DeepCopyInto(out *FileAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FileAccessReport: %s", err))
	}

	*out = *target.(*FileAccessReport)
}

// Validate valides the current information stored into the structure.
func (o *FileAccessReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject", "Limit"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("host", o.Host); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("mode", o.Mode); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("path", o.Path); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
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
func (*FileAccessReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FileAccessReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FileAccessReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FileAccessReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FileAccessReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FileAccessReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "host":
		return o.Host
	case "migrationsLog":
		return o.MigrationsLog
	case "mode":
		return o.Mode
	case "path":
		return o.Path
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "timestamp":
		return o.Timestamp
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// FileAccessReportAttributesMap represents the map of attribute for FileAccessReport.
var FileAccessReportAttributesMap = map[string]elemental.AttributeSpecification{
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
		AllowedChoices: []string{"Accept", "Reject", "Limit"},
		BSONFieldName:  "a",
		ConvertedName:  "Action",
		Description:    `Action taken.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Host": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "Host",
		DefaultValue:   "localhost",
		Description:    `Host storing the file.`,
		Exposed:        true,
		Name:           "host",
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
	"Mode": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
		ConvertedName:  "Mode",
		DefaultValue:   "rxw",
		Description:    `Mode of file access.`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Path": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "Path",
		DefaultValue:   "/etc/passwd",
		Description:    `Path of the file.`,
		Exposed:        true,
		Name:           "path",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Orderable:      true,
		Required:       true,
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

// FileAccessReportLowerCaseAttributesMap represents the map of attribute for FileAccessReport.
var FileAccessReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		AllowedChoices: []string{"Accept", "Reject", "Limit"},
		BSONFieldName:  "a",
		ConvertedName:  "Action",
		Description:    `Action taken.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"host": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "Host",
		DefaultValue:   "localhost",
		Description:    `Host storing the file.`,
		Exposed:        true,
		Name:           "host",
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
	"mode": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
		ConvertedName:  "Mode",
		DefaultValue:   "rxw",
		Description:    `Mode of file access.`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"path": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "Path",
		DefaultValue:   "/etc/passwd",
		Description:    `Path of the file.`,
		Exposed:        true,
		Name:           "path",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"processingunitid": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Orderable:      true,
		Required:       true,
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

// SparseFileAccessReportsList represents a list of SparseFileAccessReports
type SparseFileAccessReportsList []*SparseFileAccessReport

// Identity returns the identity of the objects in the list.
func (o SparseFileAccessReportsList) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Copy returns a pointer to a copy the SparseFileAccessReportsList.
func (o SparseFileAccessReportsList) Copy() elemental.Identifiables {

	copy := append(SparseFileAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFileAccessReportsList.
func (o SparseFileAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFileAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFileAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFileAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFileAccessReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToPlain returns the SparseFileAccessReportsList converted to FileAccessReportsList.
func (o SparseFileAccessReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFileAccessReportsList) Version() int {

	return 1
}

// SparseFileAccessReport represents the sparse version of a fileaccessreport.
type SparseFileAccessReport struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Action taken.
	Action *FileAccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"a,omitempty" mapstructure:"action,omitempty"`

	// Host storing the file.
	Host *string `json:"host,omitempty" msgpack:"host,omitempty" bson:"b,omitempty" mapstructure:"host,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Mode of file access.
	Mode *string `json:"mode,omitempty" msgpack:"mode,omitempty" bson:"c,omitempty" mapstructure:"mode,omitempty"`

	// Path of the file.
	Path *string `json:"path,omitempty" msgpack:"path,omitempty" bson:"d,omitempty" mapstructure:"path,omitempty"`

	// ID of the processing unit.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"e,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"f,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"g,omitempty" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFileAccessReport returns a new  SparseFileAccessReport.
func NewSparseFileAccessReport() *SparseFileAccessReport {
	return &SparseFileAccessReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFileAccessReport) Identity() elemental.Identity {

	return FileAccessReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFileAccessReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFileAccessReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFileAccessReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFileAccessReport{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Action != nil {
		s.Action = o.Action
	}
	if o.Host != nil {
		s.Host = o.Host
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Mode != nil {
		s.Mode = o.Mode
	}
	if o.Path != nil {
		s.Path = o.Path
	}
	if o.ProcessingUnitID != nil {
		s.ProcessingUnitID = o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
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
func (o *SparseFileAccessReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFileAccessReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Action != nil {
		o.Action = s.Action
	}
	if s.Host != nil {
		o.Host = s.Host
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Mode != nil {
		o.Mode = s.Mode
	}
	if s.Path != nil {
		o.Path = s.Path
	}
	if s.ProcessingUnitID != nil {
		o.ProcessingUnitID = s.ProcessingUnitID
	}
	if s.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
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
func (o *SparseFileAccessReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFileAccessReport) ToPlain() elemental.PlainIdentifiable {

	out := NewFileAccessReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.Host != nil {
		out.Host = *o.Host
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}
	if o.Path != nil {
		out.Path = *o.Path
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
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
func (o *SparseFileAccessReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseFileAccessReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseFileAccessReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseFileAccessReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseFileAccessReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseFileAccessReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseFileAccessReport.
func (o *SparseFileAccessReport) DeepCopy() *SparseFileAccessReport {

	if o == nil {
		return nil
	}

	out := &SparseFileAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFileAccessReport.
func (o *SparseFileAccessReport) DeepCopyInto(out *SparseFileAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFileAccessReport: %s", err))
	}

	*out = *target.(*SparseFileAccessReport)
}

type mongoAttributesFileAccessReport struct {
	ID                      bson.ObjectId               `bson:"_id,omitempty"`
	Action                  FileAccessReportActionValue `bson:"a,omitempty"`
	Host                    string                      `bson:"b,omitempty"`
	MigrationsLog           map[string]string           `bson:"migrationslog,omitempty"`
	Mode                    string                      `bson:"c,omitempty"`
	Path                    string                      `bson:"d,omitempty"`
	ProcessingUnitID        string                      `bson:"e,omitempty"`
	ProcessingUnitNamespace string                      `bson:"f,omitempty"`
	Timestamp               time.Time                   `bson:"g,omitempty"`
	ZHash                   int                         `bson:"zhash"`
	Zone                    int                         `bson:"zone"`
}
type mongoAttributesSparseFileAccessReport struct {
	ID                      bson.ObjectId                `bson:"_id,omitempty"`
	Action                  *FileAccessReportActionValue `bson:"a,omitempty"`
	Host                    *string                      `bson:"b,omitempty"`
	MigrationsLog           *map[string]string           `bson:"migrationslog,omitempty"`
	Mode                    *string                      `bson:"c,omitempty"`
	Path                    *string                      `bson:"d,omitempty"`
	ProcessingUnitID        *string                      `bson:"e,omitempty"`
	ProcessingUnitNamespace *string                      `bson:"f,omitempty"`
	Timestamp               *time.Time                   `bson:"g,omitempty"`
	ZHash                   *int                         `bson:"zhash,omitempty"`
	Zone                    *int                         `bson:"zone,omitempty"`
}
