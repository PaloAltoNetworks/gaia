package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AuditReportIdentity represents the Identity of the object.
var AuditReportIdentity = elemental.Identity{
	Name:     "auditreport",
	Category: "auditreports",
	Package:  "zack",
	Private:  false,
}

// AuditReportsList represents a list of AuditReports
type AuditReportsList []*AuditReport

// Identity returns the identity of the objects in the list.
func (o AuditReportsList) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Copy returns a pointer to a copy the AuditReportsList.
func (o AuditReportsList) Copy() elemental.Identifiables {

	copy := append(AuditReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuditReportsList.
func (o AuditReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuditReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AuditReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuditReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuditReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToSparse returns the AuditReportsList converted to SparseAuditReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuditReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAuditReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAuditReport)
	}

	return out
}

// Version returns the version of the content.
func (o AuditReportsList) Version() int {

	return 1
}

// AuditReport represents the model of a auditreport
type AuditReport struct {
	// The login ID of the user who started the audited process.
	AUID string `json:"AUID,omitempty" msgpack:"AUID,omitempty" bson:"a,omitempty" mapstructure:"AUID,omitempty"`

	// Command working directory.
	CWD string `json:"CWD,omitempty" msgpack:"CWD,omitempty" bson:"b,omitempty" mapstructure:"CWD,omitempty"`

	// Effective group ID of the user who started the audited process.
	EGID int `json:"EGID,omitempty" msgpack:"EGID,omitempty" bson:"c,omitempty" mapstructure:"EGID,omitempty"`

	// Effective user ID of the user who started the audited process.
	EUID int `json:"EUID,omitempty" msgpack:"EUID,omitempty" bson:"d,omitempty" mapstructure:"EUID,omitempty"`

	// Path to the executable.
	EXE string `json:"EXE,omitempty" msgpack:"EXE,omitempty" bson:"e,omitempty" mapstructure:"EXE,omitempty"`

	// File system group ID of the user who started the audited process.
	FSGID int `json:"FSGID,omitempty" msgpack:"FSGID,omitempty" bson:"f,omitempty" mapstructure:"FSGID,omitempty"`

	// File system user ID of the user who started the audited process.
	FSUID int `json:"FSUID,omitempty" msgpack:"FSUID,omitempty" bson:"g,omitempty" mapstructure:"FSUID,omitempty"`

	// Full path of the file that was passed to the system call.
	FilePath string `json:"FilePath,omitempty" msgpack:"FilePath,omitempty" bson:"h,omitempty" mapstructure:"FilePath,omitempty"`

	// Group ID of the user who started the analyzed process.
	GID int `json:"GID,omitempty" msgpack:"GID,omitempty" bson:"i,omitempty" mapstructure:"GID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// File or directory permissions.
	PER int `json:"PER,omitempty" msgpack:"PER,omitempty" bson:"j,omitempty" mapstructure:"PER,omitempty"`

	// Process ID of the executable.
	PID int `json:"PID,omitempty" msgpack:"PID,omitempty" bson:"k,omitempty" mapstructure:"PID,omitempty"`

	// Process ID of the parent executable.
	PPID int `json:"PPID,omitempty" msgpack:"PPID,omitempty" bson:"l,omitempty" mapstructure:"PPID,omitempty"`

	// Set group ID of the user who started the audited process.
	SGID int `json:"SGID,omitempty" msgpack:"SGID,omitempty" bson:"m,omitempty" mapstructure:"SGID,omitempty"`

	// Set user ID of the user who started the audited process.
	SUID int `json:"SUID,omitempty" msgpack:"SUID,omitempty" bson:"n,omitempty" mapstructure:"SUID,omitempty"`

	// User ID.
	UID int `json:"UID,omitempty" msgpack:"UID,omitempty" bson:"o,omitempty" mapstructure:"UID,omitempty"`

	// First argument of the executed system call.
	A0 string `json:"a0,omitempty" msgpack:"a0,omitempty" bson:"p,omitempty" mapstructure:"a0,omitempty"`

	// Second argument of the executed system call.
	A1 string `json:"a1,omitempty" msgpack:"a1,omitempty" bson:"q,omitempty" mapstructure:"a1,omitempty"`

	// Third argument of the executed system call.
	A2 string `json:"a2,omitempty" msgpack:"a2,omitempty" bson:"r,omitempty" mapstructure:"a2,omitempty"`

	// Fourth argument of the executed system call.
	A3 string `json:"a3,omitempty" msgpack:"a3,omitempty" bson:"s,omitempty" mapstructure:"a3,omitempty"`

	// Architecture of the system of the monitored process.
	Arch string `json:"arch,omitempty" msgpack:"arch,omitempty" bson:"t,omitempty" mapstructure:"arch,omitempty"`

	// Arguments passed to the command.
	Arguments []string `json:"arguments,omitempty" msgpack:"arguments,omitempty" bson:"u,omitempty" mapstructure:"arguments,omitempty"`

	// ID of the audit profile that triggered the report.
	AuditProfileID string `json:"auditProfileID,omitempty" msgpack:"auditProfileID,omitempty" bson:"v,omitempty" mapstructure:"auditProfileID,omitempty"`

	// Namespace of the audit profile that triggered the report.
	AuditProfileNamespace string `json:"auditProfileNamespace,omitempty" msgpack:"auditProfileNamespace,omitempty" bson:"w,omitempty" mapstructure:"auditProfileNamespace,omitempty"`

	// Command issued.
	Command string `json:"command,omitempty" msgpack:"command,omitempty" bson:"x,omitempty" mapstructure:"command,omitempty"`

	// ID of the enforcer reporting.
	EnforcerID string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"y,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer reporting.
	EnforcerNamespace string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"z,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Exit code of the executed system call.
	Exit int `json:"exit,omitempty" msgpack:"exit,omitempty" bson:"aa,omitempty" mapstructure:"exit,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// ID of the processing unit originating the report.
	ProcessingUnitID string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"ab,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit originating the report.
	ProcessingUnitNamespace string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"ac,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of audit record.
	RecordType string `json:"recordType,omitempty" msgpack:"recordType,omitempty" bson:"ad,omitempty" mapstructure:"recordType,omitempty"`

	// Needs documentation.
	Sequence int `json:"sequence,omitempty" msgpack:"sequence,omitempty" bson:"ae,omitempty" mapstructure:"sequence,omitempty"`

	// Tells if the operation has been a success or a failure.
	Success bool `json:"success,omitempty" msgpack:"success,omitempty" bson:"af,omitempty" mapstructure:"success,omitempty"`

	// System call executed.
	Syscall string `json:"syscall,omitempty" msgpack:"syscall,omitempty" bson:"ag,omitempty" mapstructure:"syscall,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"ah,omitempty" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAuditReport returns a new *AuditReport
func NewAuditReport() *AuditReport {

	return &AuditReport{
		ModelVersion:  1,
		Arguments:     []string{},
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *AuditReport) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AuditReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AuditReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AuditReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAuditReport{}

	s.AUID = o.AUID
	s.CWD = o.CWD
	s.EGID = o.EGID
	s.EUID = o.EUID
	s.EXE = o.EXE
	s.FSGID = o.FSGID
	s.FSUID = o.FSUID
	s.FilePath = o.FilePath
	s.GID = o.GID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.PER = o.PER
	s.PID = o.PID
	s.PPID = o.PPID
	s.SGID = o.SGID
	s.SUID = o.SUID
	s.UID = o.UID
	s.A0 = o.A0
	s.A1 = o.A1
	s.A2 = o.A2
	s.A3 = o.A3
	s.Arch = o.Arch
	s.Arguments = o.Arguments
	s.AuditProfileID = o.AuditProfileID
	s.AuditProfileNamespace = o.AuditProfileNamespace
	s.Command = o.Command
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.Exit = o.Exit
	s.MigrationsLog = o.MigrationsLog
	s.ProcessingUnitID = o.ProcessingUnitID
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.RecordType = o.RecordType
	s.Sequence = o.Sequence
	s.Success = o.Success
	s.Syscall = o.Syscall
	s.Timestamp = o.Timestamp
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AuditReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAuditReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AUID = s.AUID
	o.CWD = s.CWD
	o.EGID = s.EGID
	o.EUID = s.EUID
	o.EXE = s.EXE
	o.FSGID = s.FSGID
	o.FSUID = s.FSUID
	o.FilePath = s.FilePath
	o.GID = s.GID
	o.ID = s.ID.Hex()
	o.PER = s.PER
	o.PID = s.PID
	o.PPID = s.PPID
	o.SGID = s.SGID
	o.SUID = s.SUID
	o.UID = s.UID
	o.A0 = s.A0
	o.A1 = s.A1
	o.A2 = s.A2
	o.A3 = s.A3
	o.Arch = s.Arch
	o.Arguments = s.Arguments
	o.AuditProfileID = s.AuditProfileID
	o.AuditProfileNamespace = s.AuditProfileNamespace
	o.Command = s.Command
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.Exit = s.Exit
	o.MigrationsLog = s.MigrationsLog
	o.ProcessingUnitID = s.ProcessingUnitID
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.RecordType = s.RecordType
	o.Sequence = s.Sequence
	o.Success = s.Success
	o.Syscall = s.Syscall
	o.Timestamp = s.Timestamp
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AuditReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AuditReport) BleveType() string {

	return "auditreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AuditReport) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// Doc returns the documentation for the object
func (o *AuditReport) Doc() string {

	return `Post a new audit report.`
}

func (o *AuditReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *AuditReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *AuditReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *AuditReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AuditReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AuditReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AuditReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AuditReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuditReport{
			AUID:                    &o.AUID,
			CWD:                     &o.CWD,
			EGID:                    &o.EGID,
			EUID:                    &o.EUID,
			EXE:                     &o.EXE,
			FSGID:                   &o.FSGID,
			FSUID:                   &o.FSUID,
			FilePath:                &o.FilePath,
			GID:                     &o.GID,
			ID:                      &o.ID,
			PER:                     &o.PER,
			PID:                     &o.PID,
			PPID:                    &o.PPID,
			SGID:                    &o.SGID,
			SUID:                    &o.SUID,
			UID:                     &o.UID,
			A0:                      &o.A0,
			A1:                      &o.A1,
			A2:                      &o.A2,
			A3:                      &o.A3,
			Arch:                    &o.Arch,
			Arguments:               &o.Arguments,
			AuditProfileID:          &o.AuditProfileID,
			AuditProfileNamespace:   &o.AuditProfileNamespace,
			Command:                 &o.Command,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			Exit:                    &o.Exit,
			MigrationsLog:           &o.MigrationsLog,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			RecordType:              &o.RecordType,
			Sequence:                &o.Sequence,
			Success:                 &o.Success,
			Syscall:                 &o.Syscall,
			Timestamp:               &o.Timestamp,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparseAuditReport{}
	for _, f := range fields {
		switch f {
		case "AUID":
			sp.AUID = &(o.AUID)
		case "CWD":
			sp.CWD = &(o.CWD)
		case "EGID":
			sp.EGID = &(o.EGID)
		case "EUID":
			sp.EUID = &(o.EUID)
		case "EXE":
			sp.EXE = &(o.EXE)
		case "FSGID":
			sp.FSGID = &(o.FSGID)
		case "FSUID":
			sp.FSUID = &(o.FSUID)
		case "FilePath":
			sp.FilePath = &(o.FilePath)
		case "GID":
			sp.GID = &(o.GID)
		case "ID":
			sp.ID = &(o.ID)
		case "PER":
			sp.PER = &(o.PER)
		case "PID":
			sp.PID = &(o.PID)
		case "PPID":
			sp.PPID = &(o.PPID)
		case "SGID":
			sp.SGID = &(o.SGID)
		case "SUID":
			sp.SUID = &(o.SUID)
		case "UID":
			sp.UID = &(o.UID)
		case "a0":
			sp.A0 = &(o.A0)
		case "a1":
			sp.A1 = &(o.A1)
		case "a2":
			sp.A2 = &(o.A2)
		case "a3":
			sp.A3 = &(o.A3)
		case "arch":
			sp.Arch = &(o.Arch)
		case "arguments":
			sp.Arguments = &(o.Arguments)
		case "auditProfileID":
			sp.AuditProfileID = &(o.AuditProfileID)
		case "auditProfileNamespace":
			sp.AuditProfileNamespace = &(o.AuditProfileNamespace)
		case "command":
			sp.Command = &(o.Command)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "exit":
			sp.Exit = &(o.Exit)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "recordType":
			sp.RecordType = &(o.RecordType)
		case "sequence":
			sp.Sequence = &(o.Sequence)
		case "success":
			sp.Success = &(o.Success)
		case "syscall":
			sp.Syscall = &(o.Syscall)
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

// Patch apply the non nil value of a *SparseAuditReport to the object.
func (o *AuditReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuditReport)
	if so.AUID != nil {
		o.AUID = *so.AUID
	}
	if so.CWD != nil {
		o.CWD = *so.CWD
	}
	if so.EGID != nil {
		o.EGID = *so.EGID
	}
	if so.EUID != nil {
		o.EUID = *so.EUID
	}
	if so.EXE != nil {
		o.EXE = *so.EXE
	}
	if so.FSGID != nil {
		o.FSGID = *so.FSGID
	}
	if so.FSUID != nil {
		o.FSUID = *so.FSUID
	}
	if so.FilePath != nil {
		o.FilePath = *so.FilePath
	}
	if so.GID != nil {
		o.GID = *so.GID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.PER != nil {
		o.PER = *so.PER
	}
	if so.PID != nil {
		o.PID = *so.PID
	}
	if so.PPID != nil {
		o.PPID = *so.PPID
	}
	if so.SGID != nil {
		o.SGID = *so.SGID
	}
	if so.SUID != nil {
		o.SUID = *so.SUID
	}
	if so.UID != nil {
		o.UID = *so.UID
	}
	if so.A0 != nil {
		o.A0 = *so.A0
	}
	if so.A1 != nil {
		o.A1 = *so.A1
	}
	if so.A2 != nil {
		o.A2 = *so.A2
	}
	if so.A3 != nil {
		o.A3 = *so.A3
	}
	if so.Arch != nil {
		o.Arch = *so.Arch
	}
	if so.Arguments != nil {
		o.Arguments = *so.Arguments
	}
	if so.AuditProfileID != nil {
		o.AuditProfileID = *so.AuditProfileID
	}
	if so.AuditProfileNamespace != nil {
		o.AuditProfileNamespace = *so.AuditProfileNamespace
	}
	if so.Command != nil {
		o.Command = *so.Command
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.Exit != nil {
		o.Exit = *so.Exit
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.RecordType != nil {
		o.RecordType = *so.RecordType
	}
	if so.Sequence != nil {
		o.Sequence = *so.Sequence
	}
	if so.Success != nil {
		o.Success = *so.Success
	}
	if so.Syscall != nil {
		o.Syscall = *so.Syscall
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

// DeepCopy returns a deep copy if the AuditReport.
func (o *AuditReport) DeepCopy() *AuditReport {

	if o == nil {
		return nil
	}

	out := &AuditReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AuditReport.
func (o *AuditReport) DeepCopyInto(out *AuditReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AuditReport: %s", err))
	}

	*out = *target.(*AuditReport)
}

// Validate valides the current information stored into the structure.
func (o *AuditReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("auditProfileID", o.AuditProfileID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("auditProfileNamespace", o.AuditProfileNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("recordType", o.RecordType); err != nil {
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
func (*AuditReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuditReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuditReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AuditReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuditReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AuditReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "AUID":
		return o.AUID
	case "CWD":
		return o.CWD
	case "EGID":
		return o.EGID
	case "EUID":
		return o.EUID
	case "EXE":
		return o.EXE
	case "FSGID":
		return o.FSGID
	case "FSUID":
		return o.FSUID
	case "FilePath":
		return o.FilePath
	case "GID":
		return o.GID
	case "ID":
		return o.ID
	case "PER":
		return o.PER
	case "PID":
		return o.PID
	case "PPID":
		return o.PPID
	case "SGID":
		return o.SGID
	case "SUID":
		return o.SUID
	case "UID":
		return o.UID
	case "a0":
		return o.A0
	case "a1":
		return o.A1
	case "a2":
		return o.A2
	case "a3":
		return o.A3
	case "arch":
		return o.Arch
	case "arguments":
		return o.Arguments
	case "auditProfileID":
		return o.AuditProfileID
	case "auditProfileNamespace":
		return o.AuditProfileNamespace
	case "command":
		return o.Command
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "exit":
		return o.Exit
	case "migrationsLog":
		return o.MigrationsLog
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "recordType":
		return o.RecordType
	case "sequence":
		return o.Sequence
	case "success":
		return o.Success
	case "syscall":
		return o.Syscall
	case "timestamp":
		return o.Timestamp
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AuditReportAttributesMap represents the map of attribute for AuditReport.
var AuditReportAttributesMap = map[string]elemental.AttributeSpecification{
	"AUID": {
		AllowedChoices: []string{},
		BSONFieldName:  "a",
		ConvertedName:  "AUID",
		Description:    `The login ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "AUID",
		Stored:         true,
		Type:           "string",
	},
	"CWD": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "CWD",
		Description:    `Command working directory.`,
		Exposed:        true,
		Name:           "CWD",
		Stored:         true,
		Type:           "string",
	},
	"EGID": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
		ConvertedName:  "EGID",
		Description:    `Effective group ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "EGID",
		Stored:         true,
		Type:           "integer",
	},
	"EUID": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "EUID",
		Description:    `Effective user ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "EUID",
		Stored:         true,
		Type:           "integer",
	},
	"EXE": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "EXE",
		Description:    `Path to the executable.`,
		Exposed:        true,
		Name:           "EXE",
		Stored:         true,
		Type:           "string",
	},
	"FSGID": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "FSGID",
		Description:    `File system group ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "FSGID",
		Stored:         true,
		Type:           "integer",
	},
	"FSUID": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "FSUID",
		Description:    `File system user ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "FSUID",
		Stored:         true,
		Type:           "integer",
	},
	"FilePath": {
		AllowedChoices: []string{},
		BSONFieldName:  "h",
		ConvertedName:  "FilePath",
		Description:    `Full path of the file that was passed to the system call.`,
		Exposed:        true,
		Name:           "FilePath",
		Stored:         true,
		Type:           "string",
	},
	"GID": {
		AllowedChoices: []string{},
		BSONFieldName:  "i",
		ConvertedName:  "GID",
		Description:    `Group ID of the user who started the analyzed process.`,
		Exposed:        true,
		Name:           "GID",
		Stored:         true,
		Type:           "integer",
	},
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
	"PER": {
		AllowedChoices: []string{},
		BSONFieldName:  "j",
		ConvertedName:  "PER",
		Description:    `File or directory permissions.`,
		Exposed:        true,
		Name:           "PER",
		Stored:         true,
		Type:           "integer",
	},
	"PID": {
		AllowedChoices: []string{},
		BSONFieldName:  "k",
		ConvertedName:  "PID",
		Description:    `Process ID of the executable.`,
		Exposed:        true,
		Name:           "PID",
		Stored:         true,
		Type:           "integer",
	},
	"PPID": {
		AllowedChoices: []string{},
		BSONFieldName:  "l",
		ConvertedName:  "PPID",
		Description:    `Process ID of the parent executable.`,
		Exposed:        true,
		Name:           "PPID",
		Stored:         true,
		Type:           "integer",
	},
	"SGID": {
		AllowedChoices: []string{},
		BSONFieldName:  "m",
		ConvertedName:  "SGID",
		Description:    `Set group ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "SGID",
		Stored:         true,
		Type:           "integer",
	},
	"SUID": {
		AllowedChoices: []string{},
		BSONFieldName:  "n",
		ConvertedName:  "SUID",
		Description:    `Set user ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "SUID",
		Stored:         true,
		Type:           "integer",
	},
	"UID": {
		AllowedChoices: []string{},
		BSONFieldName:  "o",
		ConvertedName:  "UID",
		Description:    `User ID.`,
		Exposed:        true,
		Name:           "UID",
		Stored:         true,
		Type:           "integer",
	},
	"A0": {
		AllowedChoices: []string{},
		BSONFieldName:  "p",
		ConvertedName:  "A0",
		Description:    `First argument of the executed system call.`,
		Exposed:        true,
		Name:           "a0",
		Stored:         true,
		Type:           "string",
	},
	"A1": {
		AllowedChoices: []string{},
		BSONFieldName:  "q",
		ConvertedName:  "A1",
		Description:    `Second argument of the executed system call.`,
		Exposed:        true,
		Name:           "a1",
		Stored:         true,
		Type:           "string",
	},
	"A2": {
		AllowedChoices: []string{},
		BSONFieldName:  "r",
		ConvertedName:  "A2",
		Description:    `Third argument of the executed system call.`,
		Exposed:        true,
		Name:           "a2",
		Stored:         true,
		Type:           "string",
	},
	"A3": {
		AllowedChoices: []string{},
		BSONFieldName:  "s",
		ConvertedName:  "A3",
		Description:    `Fourth argument of the executed system call.`,
		Exposed:        true,
		Name:           "a3",
		Stored:         true,
		Type:           "string",
	},
	"Arch": {
		AllowedChoices: []string{},
		BSONFieldName:  "t",
		ConvertedName:  "Arch",
		Description:    `Architecture of the system of the monitored process.`,
		Exposed:        true,
		Name:           "arch",
		Stored:         true,
		Type:           "string",
	},
	"Arguments": {
		AllowedChoices: []string{},
		BSONFieldName:  "u",
		ConvertedName:  "Arguments",
		Description:    `Arguments passed to the command.`,
		Exposed:        true,
		Name:           "arguments",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AuditProfileID": {
		AllowedChoices: []string{},
		BSONFieldName:  "v",
		ConvertedName:  "AuditProfileID",
		Description:    `ID of the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"AuditProfileNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "w",
		ConvertedName:  "AuditProfileNamespace",
		Description:    `Namespace of the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Command": {
		AllowedChoices: []string{},
		BSONFieldName:  "x",
		ConvertedName:  "Command",
		Description:    `Command issued.`,
		Exposed:        true,
		Name:           "command",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "y",
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "z",
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Exit": {
		AllowedChoices: []string{},
		BSONFieldName:  "aa",
		ConvertedName:  "Exit",
		Description:    `Exit code of the executed system call.`,
		Exposed:        true,
		Name:           "exit",
		Stored:         true,
		Type:           "integer",
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
		BSONFieldName:  "ab",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "ac",
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"RecordType": {
		AllowedChoices: []string{},
		BSONFieldName:  "ad",
		ConvertedName:  "RecordType",
		Description:    `Type of audit record.`,
		Exposed:        true,
		Name:           "recordType",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Sequence": {
		AllowedChoices: []string{},
		BSONFieldName:  "ae",
		ConvertedName:  "Sequence",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "sequence",
		Stored:         true,
		Type:           "integer",
	},
	"Success": {
		AllowedChoices: []string{},
		BSONFieldName:  "af",
		ConvertedName:  "Success",
		Description:    `Tells if the operation has been a success or a failure.`,
		Exposed:        true,
		Name:           "success",
		Stored:         true,
		Type:           "boolean",
	},
	"Syscall": {
		AllowedChoices: []string{},
		BSONFieldName:  "ag",
		ConvertedName:  "Syscall",
		Description:    `System call executed.`,
		Exposed:        true,
		Name:           "syscall",
		Stored:         true,
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "ah",
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

// AuditReportLowerCaseAttributesMap represents the map of attribute for AuditReport.
var AuditReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"auid": {
		AllowedChoices: []string{},
		BSONFieldName:  "a",
		ConvertedName:  "AUID",
		Description:    `The login ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "AUID",
		Stored:         true,
		Type:           "string",
	},
	"cwd": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "CWD",
		Description:    `Command working directory.`,
		Exposed:        true,
		Name:           "CWD",
		Stored:         true,
		Type:           "string",
	},
	"egid": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
		ConvertedName:  "EGID",
		Description:    `Effective group ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "EGID",
		Stored:         true,
		Type:           "integer",
	},
	"euid": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "EUID",
		Description:    `Effective user ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "EUID",
		Stored:         true,
		Type:           "integer",
	},
	"exe": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "EXE",
		Description:    `Path to the executable.`,
		Exposed:        true,
		Name:           "EXE",
		Stored:         true,
		Type:           "string",
	},
	"fsgid": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "FSGID",
		Description:    `File system group ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "FSGID",
		Stored:         true,
		Type:           "integer",
	},
	"fsuid": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "FSUID",
		Description:    `File system user ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "FSUID",
		Stored:         true,
		Type:           "integer",
	},
	"filepath": {
		AllowedChoices: []string{},
		BSONFieldName:  "h",
		ConvertedName:  "FilePath",
		Description:    `Full path of the file that was passed to the system call.`,
		Exposed:        true,
		Name:           "FilePath",
		Stored:         true,
		Type:           "string",
	},
	"gid": {
		AllowedChoices: []string{},
		BSONFieldName:  "i",
		ConvertedName:  "GID",
		Description:    `Group ID of the user who started the analyzed process.`,
		Exposed:        true,
		Name:           "GID",
		Stored:         true,
		Type:           "integer",
	},
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
	"per": {
		AllowedChoices: []string{},
		BSONFieldName:  "j",
		ConvertedName:  "PER",
		Description:    `File or directory permissions.`,
		Exposed:        true,
		Name:           "PER",
		Stored:         true,
		Type:           "integer",
	},
	"pid": {
		AllowedChoices: []string{},
		BSONFieldName:  "k",
		ConvertedName:  "PID",
		Description:    `Process ID of the executable.`,
		Exposed:        true,
		Name:           "PID",
		Stored:         true,
		Type:           "integer",
	},
	"ppid": {
		AllowedChoices: []string{},
		BSONFieldName:  "l",
		ConvertedName:  "PPID",
		Description:    `Process ID of the parent executable.`,
		Exposed:        true,
		Name:           "PPID",
		Stored:         true,
		Type:           "integer",
	},
	"sgid": {
		AllowedChoices: []string{},
		BSONFieldName:  "m",
		ConvertedName:  "SGID",
		Description:    `Set group ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "SGID",
		Stored:         true,
		Type:           "integer",
	},
	"suid": {
		AllowedChoices: []string{},
		BSONFieldName:  "n",
		ConvertedName:  "SUID",
		Description:    `Set user ID of the user who started the audited process.`,
		Exposed:        true,
		Name:           "SUID",
		Stored:         true,
		Type:           "integer",
	},
	"uid": {
		AllowedChoices: []string{},
		BSONFieldName:  "o",
		ConvertedName:  "UID",
		Description:    `User ID.`,
		Exposed:        true,
		Name:           "UID",
		Stored:         true,
		Type:           "integer",
	},
	"a0": {
		AllowedChoices: []string{},
		BSONFieldName:  "p",
		ConvertedName:  "A0",
		Description:    `First argument of the executed system call.`,
		Exposed:        true,
		Name:           "a0",
		Stored:         true,
		Type:           "string",
	},
	"a1": {
		AllowedChoices: []string{},
		BSONFieldName:  "q",
		ConvertedName:  "A1",
		Description:    `Second argument of the executed system call.`,
		Exposed:        true,
		Name:           "a1",
		Stored:         true,
		Type:           "string",
	},
	"a2": {
		AllowedChoices: []string{},
		BSONFieldName:  "r",
		ConvertedName:  "A2",
		Description:    `Third argument of the executed system call.`,
		Exposed:        true,
		Name:           "a2",
		Stored:         true,
		Type:           "string",
	},
	"a3": {
		AllowedChoices: []string{},
		BSONFieldName:  "s",
		ConvertedName:  "A3",
		Description:    `Fourth argument of the executed system call.`,
		Exposed:        true,
		Name:           "a3",
		Stored:         true,
		Type:           "string",
	},
	"arch": {
		AllowedChoices: []string{},
		BSONFieldName:  "t",
		ConvertedName:  "Arch",
		Description:    `Architecture of the system of the monitored process.`,
		Exposed:        true,
		Name:           "arch",
		Stored:         true,
		Type:           "string",
	},
	"arguments": {
		AllowedChoices: []string{},
		BSONFieldName:  "u",
		ConvertedName:  "Arguments",
		Description:    `Arguments passed to the command.`,
		Exposed:        true,
		Name:           "arguments",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"auditprofileid": {
		AllowedChoices: []string{},
		BSONFieldName:  "v",
		ConvertedName:  "AuditProfileID",
		Description:    `ID of the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"auditprofilenamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "w",
		ConvertedName:  "AuditProfileNamespace",
		Description:    `Namespace of the audit profile that triggered the report.`,
		Exposed:        true,
		Name:           "auditProfileNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"command": {
		AllowedChoices: []string{},
		BSONFieldName:  "x",
		ConvertedName:  "Command",
		Description:    `Command issued.`,
		Exposed:        true,
		Name:           "command",
		Stored:         true,
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "y",
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "z",
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer reporting.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"exit": {
		AllowedChoices: []string{},
		BSONFieldName:  "aa",
		ConvertedName:  "Exit",
		Description:    `Exit code of the executed system call.`,
		Exposed:        true,
		Name:           "exit",
		Stored:         true,
		Type:           "integer",
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
		BSONFieldName:  "ab",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "ac",
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit originating the report.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"recordtype": {
		AllowedChoices: []string{},
		BSONFieldName:  "ad",
		ConvertedName:  "RecordType",
		Description:    `Type of audit record.`,
		Exposed:        true,
		Name:           "recordType",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sequence": {
		AllowedChoices: []string{},
		BSONFieldName:  "ae",
		ConvertedName:  "Sequence",
		Description:    `Needs documentation.`,
		Exposed:        true,
		Name:           "sequence",
		Stored:         true,
		Type:           "integer",
	},
	"success": {
		AllowedChoices: []string{},
		BSONFieldName:  "af",
		ConvertedName:  "Success",
		Description:    `Tells if the operation has been a success or a failure.`,
		Exposed:        true,
		Name:           "success",
		Stored:         true,
		Type:           "boolean",
	},
	"syscall": {
		AllowedChoices: []string{},
		BSONFieldName:  "ag",
		ConvertedName:  "Syscall",
		Description:    `System call executed.`,
		Exposed:        true,
		Name:           "syscall",
		Stored:         true,
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "ah",
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

// SparseAuditReportsList represents a list of SparseAuditReports
type SparseAuditReportsList []*SparseAuditReport

// Identity returns the identity of the objects in the list.
func (o SparseAuditReportsList) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Copy returns a pointer to a copy the SparseAuditReportsList.
func (o SparseAuditReportsList) Copy() elemental.Identifiables {

	copy := append(SparseAuditReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuditReportsList.
func (o SparseAuditReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuditReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuditReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuditReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuditReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToPlain returns the SparseAuditReportsList converted to AuditReportsList.
func (o SparseAuditReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuditReportsList) Version() int {

	return 1
}

// SparseAuditReport represents the sparse version of a auditreport.
type SparseAuditReport struct {
	// The login ID of the user who started the audited process.
	AUID *string `json:"AUID,omitempty" msgpack:"AUID,omitempty" bson:"a,omitempty" mapstructure:"AUID,omitempty"`

	// Command working directory.
	CWD *string `json:"CWD,omitempty" msgpack:"CWD,omitempty" bson:"b,omitempty" mapstructure:"CWD,omitempty"`

	// Effective group ID of the user who started the audited process.
	EGID *int `json:"EGID,omitempty" msgpack:"EGID,omitempty" bson:"c,omitempty" mapstructure:"EGID,omitempty"`

	// Effective user ID of the user who started the audited process.
	EUID *int `json:"EUID,omitempty" msgpack:"EUID,omitempty" bson:"d,omitempty" mapstructure:"EUID,omitempty"`

	// Path to the executable.
	EXE *string `json:"EXE,omitempty" msgpack:"EXE,omitempty" bson:"e,omitempty" mapstructure:"EXE,omitempty"`

	// File system group ID of the user who started the audited process.
	FSGID *int `json:"FSGID,omitempty" msgpack:"FSGID,omitempty" bson:"f,omitempty" mapstructure:"FSGID,omitempty"`

	// File system user ID of the user who started the audited process.
	FSUID *int `json:"FSUID,omitempty" msgpack:"FSUID,omitempty" bson:"g,omitempty" mapstructure:"FSUID,omitempty"`

	// Full path of the file that was passed to the system call.
	FilePath *string `json:"FilePath,omitempty" msgpack:"FilePath,omitempty" bson:"h,omitempty" mapstructure:"FilePath,omitempty"`

	// Group ID of the user who started the analyzed process.
	GID *int `json:"GID,omitempty" msgpack:"GID,omitempty" bson:"i,omitempty" mapstructure:"GID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// File or directory permissions.
	PER *int `json:"PER,omitempty" msgpack:"PER,omitempty" bson:"j,omitempty" mapstructure:"PER,omitempty"`

	// Process ID of the executable.
	PID *int `json:"PID,omitempty" msgpack:"PID,omitempty" bson:"k,omitempty" mapstructure:"PID,omitempty"`

	// Process ID of the parent executable.
	PPID *int `json:"PPID,omitempty" msgpack:"PPID,omitempty" bson:"l,omitempty" mapstructure:"PPID,omitempty"`

	// Set group ID of the user who started the audited process.
	SGID *int `json:"SGID,omitempty" msgpack:"SGID,omitempty" bson:"m,omitempty" mapstructure:"SGID,omitempty"`

	// Set user ID of the user who started the audited process.
	SUID *int `json:"SUID,omitempty" msgpack:"SUID,omitempty" bson:"n,omitempty" mapstructure:"SUID,omitempty"`

	// User ID.
	UID *int `json:"UID,omitempty" msgpack:"UID,omitempty" bson:"o,omitempty" mapstructure:"UID,omitempty"`

	// First argument of the executed system call.
	A0 *string `json:"a0,omitempty" msgpack:"a0,omitempty" bson:"p,omitempty" mapstructure:"a0,omitempty"`

	// Second argument of the executed system call.
	A1 *string `json:"a1,omitempty" msgpack:"a1,omitempty" bson:"q,omitempty" mapstructure:"a1,omitempty"`

	// Third argument of the executed system call.
	A2 *string `json:"a2,omitempty" msgpack:"a2,omitempty" bson:"r,omitempty" mapstructure:"a2,omitempty"`

	// Fourth argument of the executed system call.
	A3 *string `json:"a3,omitempty" msgpack:"a3,omitempty" bson:"s,omitempty" mapstructure:"a3,omitempty"`

	// Architecture of the system of the monitored process.
	Arch *string `json:"arch,omitempty" msgpack:"arch,omitempty" bson:"t,omitempty" mapstructure:"arch,omitempty"`

	// Arguments passed to the command.
	Arguments *[]string `json:"arguments,omitempty" msgpack:"arguments,omitempty" bson:"u,omitempty" mapstructure:"arguments,omitempty"`

	// ID of the audit profile that triggered the report.
	AuditProfileID *string `json:"auditProfileID,omitempty" msgpack:"auditProfileID,omitempty" bson:"v,omitempty" mapstructure:"auditProfileID,omitempty"`

	// Namespace of the audit profile that triggered the report.
	AuditProfileNamespace *string `json:"auditProfileNamespace,omitempty" msgpack:"auditProfileNamespace,omitempty" bson:"w,omitempty" mapstructure:"auditProfileNamespace,omitempty"`

	// Command issued.
	Command *string `json:"command,omitempty" msgpack:"command,omitempty" bson:"x,omitempty" mapstructure:"command,omitempty"`

	// ID of the enforcer reporting.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"y,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer reporting.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"z,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Exit code of the executed system call.
	Exit *int `json:"exit,omitempty" msgpack:"exit,omitempty" bson:"aa,omitempty" mapstructure:"exit,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// ID of the processing unit originating the report.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"ab,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit originating the report.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"ac,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of audit record.
	RecordType *string `json:"recordType,omitempty" msgpack:"recordType,omitempty" bson:"ad,omitempty" mapstructure:"recordType,omitempty"`

	// Needs documentation.
	Sequence *int `json:"sequence,omitempty" msgpack:"sequence,omitempty" bson:"ae,omitempty" mapstructure:"sequence,omitempty"`

	// Tells if the operation has been a success or a failure.
	Success *bool `json:"success,omitempty" msgpack:"success,omitempty" bson:"af,omitempty" mapstructure:"success,omitempty"`

	// System call executed.
	Syscall *string `json:"syscall,omitempty" msgpack:"syscall,omitempty" bson:"ag,omitempty" mapstructure:"syscall,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"ah,omitempty" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAuditReport returns a new  SparseAuditReport.
func NewSparseAuditReport() *SparseAuditReport {
	return &SparseAuditReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuditReport) Identity() elemental.Identity {

	return AuditReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuditReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuditReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAuditReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAuditReport{}

	if o.AUID != nil {
		s.AUID = o.AUID
	}
	if o.CWD != nil {
		s.CWD = o.CWD
	}
	if o.EGID != nil {
		s.EGID = o.EGID
	}
	if o.EUID != nil {
		s.EUID = o.EUID
	}
	if o.EXE != nil {
		s.EXE = o.EXE
	}
	if o.FSGID != nil {
		s.FSGID = o.FSGID
	}
	if o.FSUID != nil {
		s.FSUID = o.FSUID
	}
	if o.FilePath != nil {
		s.FilePath = o.FilePath
	}
	if o.GID != nil {
		s.GID = o.GID
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.PER != nil {
		s.PER = o.PER
	}
	if o.PID != nil {
		s.PID = o.PID
	}
	if o.PPID != nil {
		s.PPID = o.PPID
	}
	if o.SGID != nil {
		s.SGID = o.SGID
	}
	if o.SUID != nil {
		s.SUID = o.SUID
	}
	if o.UID != nil {
		s.UID = o.UID
	}
	if o.A0 != nil {
		s.A0 = o.A0
	}
	if o.A1 != nil {
		s.A1 = o.A1
	}
	if o.A2 != nil {
		s.A2 = o.A2
	}
	if o.A3 != nil {
		s.A3 = o.A3
	}
	if o.Arch != nil {
		s.Arch = o.Arch
	}
	if o.Arguments != nil {
		s.Arguments = o.Arguments
	}
	if o.AuditProfileID != nil {
		s.AuditProfileID = o.AuditProfileID
	}
	if o.AuditProfileNamespace != nil {
		s.AuditProfileNamespace = o.AuditProfileNamespace
	}
	if o.Command != nil {
		s.Command = o.Command
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.Exit != nil {
		s.Exit = o.Exit
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.ProcessingUnitID != nil {
		s.ProcessingUnitID = o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	}
	if o.RecordType != nil {
		s.RecordType = o.RecordType
	}
	if o.Sequence != nil {
		s.Sequence = o.Sequence
	}
	if o.Success != nil {
		s.Success = o.Success
	}
	if o.Syscall != nil {
		s.Syscall = o.Syscall
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
func (o *SparseAuditReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAuditReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.AUID != nil {
		o.AUID = s.AUID
	}
	if s.CWD != nil {
		o.CWD = s.CWD
	}
	if s.EGID != nil {
		o.EGID = s.EGID
	}
	if s.EUID != nil {
		o.EUID = s.EUID
	}
	if s.EXE != nil {
		o.EXE = s.EXE
	}
	if s.FSGID != nil {
		o.FSGID = s.FSGID
	}
	if s.FSUID != nil {
		o.FSUID = s.FSUID
	}
	if s.FilePath != nil {
		o.FilePath = s.FilePath
	}
	if s.GID != nil {
		o.GID = s.GID
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.PER != nil {
		o.PER = s.PER
	}
	if s.PID != nil {
		o.PID = s.PID
	}
	if s.PPID != nil {
		o.PPID = s.PPID
	}
	if s.SGID != nil {
		o.SGID = s.SGID
	}
	if s.SUID != nil {
		o.SUID = s.SUID
	}
	if s.UID != nil {
		o.UID = s.UID
	}
	if s.A0 != nil {
		o.A0 = s.A0
	}
	if s.A1 != nil {
		o.A1 = s.A1
	}
	if s.A2 != nil {
		o.A2 = s.A2
	}
	if s.A3 != nil {
		o.A3 = s.A3
	}
	if s.Arch != nil {
		o.Arch = s.Arch
	}
	if s.Arguments != nil {
		o.Arguments = s.Arguments
	}
	if s.AuditProfileID != nil {
		o.AuditProfileID = s.AuditProfileID
	}
	if s.AuditProfileNamespace != nil {
		o.AuditProfileNamespace = s.AuditProfileNamespace
	}
	if s.Command != nil {
		o.Command = s.Command
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.Exit != nil {
		o.Exit = s.Exit
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.ProcessingUnitID != nil {
		o.ProcessingUnitID = s.ProcessingUnitID
	}
	if s.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	}
	if s.RecordType != nil {
		o.RecordType = s.RecordType
	}
	if s.Sequence != nil {
		o.Sequence = s.Sequence
	}
	if s.Success != nil {
		o.Success = s.Success
	}
	if s.Syscall != nil {
		o.Syscall = s.Syscall
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
func (o *SparseAuditReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuditReport) ToPlain() elemental.PlainIdentifiable {

	out := NewAuditReport()
	if o.AUID != nil {
		out.AUID = *o.AUID
	}
	if o.CWD != nil {
		out.CWD = *o.CWD
	}
	if o.EGID != nil {
		out.EGID = *o.EGID
	}
	if o.EUID != nil {
		out.EUID = *o.EUID
	}
	if o.EXE != nil {
		out.EXE = *o.EXE
	}
	if o.FSGID != nil {
		out.FSGID = *o.FSGID
	}
	if o.FSUID != nil {
		out.FSUID = *o.FSUID
	}
	if o.FilePath != nil {
		out.FilePath = *o.FilePath
	}
	if o.GID != nil {
		out.GID = *o.GID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.PER != nil {
		out.PER = *o.PER
	}
	if o.PID != nil {
		out.PID = *o.PID
	}
	if o.PPID != nil {
		out.PPID = *o.PPID
	}
	if o.SGID != nil {
		out.SGID = *o.SGID
	}
	if o.SUID != nil {
		out.SUID = *o.SUID
	}
	if o.UID != nil {
		out.UID = *o.UID
	}
	if o.A0 != nil {
		out.A0 = *o.A0
	}
	if o.A1 != nil {
		out.A1 = *o.A1
	}
	if o.A2 != nil {
		out.A2 = *o.A2
	}
	if o.A3 != nil {
		out.A3 = *o.A3
	}
	if o.Arch != nil {
		out.Arch = *o.Arch
	}
	if o.Arguments != nil {
		out.Arguments = *o.Arguments
	}
	if o.AuditProfileID != nil {
		out.AuditProfileID = *o.AuditProfileID
	}
	if o.AuditProfileNamespace != nil {
		out.AuditProfileNamespace = *o.AuditProfileNamespace
	}
	if o.Command != nil {
		out.Command = *o.Command
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.Exit != nil {
		out.Exit = *o.Exit
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.RecordType != nil {
		out.RecordType = *o.RecordType
	}
	if o.Sequence != nil {
		out.Sequence = *o.Sequence
	}
	if o.Success != nil {
		out.Success = *o.Success
	}
	if o.Syscall != nil {
		out.Syscall = *o.Syscall
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
func (o *SparseAuditReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAuditReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAuditReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAuditReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAuditReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAuditReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAuditReport.
func (o *SparseAuditReport) DeepCopy() *SparseAuditReport {

	if o == nil {
		return nil
	}

	out := &SparseAuditReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuditReport.
func (o *SparseAuditReport) DeepCopyInto(out *SparseAuditReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuditReport: %s", err))
	}

	*out = *target.(*SparseAuditReport)
}

type mongoAttributesAuditReport struct {
	AUID                    string            `bson:"a,omitempty"`
	CWD                     string            `bson:"b,omitempty"`
	EGID                    int               `bson:"c,omitempty"`
	EUID                    int               `bson:"d,omitempty"`
	EXE                     string            `bson:"e,omitempty"`
	FSGID                   int               `bson:"f,omitempty"`
	FSUID                   int               `bson:"g,omitempty"`
	FilePath                string            `bson:"h,omitempty"`
	GID                     int               `bson:"i,omitempty"`
	ID                      bson.ObjectId     `bson:"_id,omitempty"`
	PER                     int               `bson:"j,omitempty"`
	PID                     int               `bson:"k,omitempty"`
	PPID                    int               `bson:"l,omitempty"`
	SGID                    int               `bson:"m,omitempty"`
	SUID                    int               `bson:"n,omitempty"`
	UID                     int               `bson:"o,omitempty"`
	A0                      string            `bson:"p,omitempty"`
	A1                      string            `bson:"q,omitempty"`
	A2                      string            `bson:"r,omitempty"`
	A3                      string            `bson:"s,omitempty"`
	Arch                    string            `bson:"t,omitempty"`
	Arguments               []string          `bson:"u,omitempty"`
	AuditProfileID          string            `bson:"v,omitempty"`
	AuditProfileNamespace   string            `bson:"w,omitempty"`
	Command                 string            `bson:"x,omitempty"`
	EnforcerID              string            `bson:"y,omitempty"`
	EnforcerNamespace       string            `bson:"z,omitempty"`
	Exit                    int               `bson:"aa,omitempty"`
	MigrationsLog           map[string]string `bson:"migrationslog,omitempty"`
	ProcessingUnitID        string            `bson:"ab,omitempty"`
	ProcessingUnitNamespace string            `bson:"ac,omitempty"`
	RecordType              string            `bson:"ad,omitempty"`
	Sequence                int               `bson:"ae,omitempty"`
	Success                 bool              `bson:"af,omitempty"`
	Syscall                 string            `bson:"ag,omitempty"`
	Timestamp               time.Time         `bson:"ah,omitempty"`
	ZHash                   int               `bson:"zhash"`
	Zone                    int               `bson:"zone"`
}
type mongoAttributesSparseAuditReport struct {
	AUID                    *string            `bson:"a,omitempty"`
	CWD                     *string            `bson:"b,omitempty"`
	EGID                    *int               `bson:"c,omitempty"`
	EUID                    *int               `bson:"d,omitempty"`
	EXE                     *string            `bson:"e,omitempty"`
	FSGID                   *int               `bson:"f,omitempty"`
	FSUID                   *int               `bson:"g,omitempty"`
	FilePath                *string            `bson:"h,omitempty"`
	GID                     *int               `bson:"i,omitempty"`
	ID                      bson.ObjectId      `bson:"_id,omitempty"`
	PER                     *int               `bson:"j,omitempty"`
	PID                     *int               `bson:"k,omitempty"`
	PPID                    *int               `bson:"l,omitempty"`
	SGID                    *int               `bson:"m,omitempty"`
	SUID                    *int               `bson:"n,omitempty"`
	UID                     *int               `bson:"o,omitempty"`
	A0                      *string            `bson:"p,omitempty"`
	A1                      *string            `bson:"q,omitempty"`
	A2                      *string            `bson:"r,omitempty"`
	A3                      *string            `bson:"s,omitempty"`
	Arch                    *string            `bson:"t,omitempty"`
	Arguments               *[]string          `bson:"u,omitempty"`
	AuditProfileID          *string            `bson:"v,omitempty"`
	AuditProfileNamespace   *string            `bson:"w,omitempty"`
	Command                 *string            `bson:"x,omitempty"`
	EnforcerID              *string            `bson:"y,omitempty"`
	EnforcerNamespace       *string            `bson:"z,omitempty"`
	Exit                    *int               `bson:"aa,omitempty"`
	MigrationsLog           *map[string]string `bson:"migrationslog,omitempty"`
	ProcessingUnitID        *string            `bson:"ab,omitempty"`
	ProcessingUnitNamespace *string            `bson:"ac,omitempty"`
	RecordType              *string            `bson:"ad,omitempty"`
	Sequence                *int               `bson:"ae,omitempty"`
	Success                 *bool              `bson:"af,omitempty"`
	Syscall                 *string            `bson:"ag,omitempty"`
	Timestamp               *time.Time         `bson:"ah,omitempty"`
	ZHash                   *int               `bson:"zhash,omitempty"`
	Zone                    *int               `bson:"zone,omitempty"`
}
