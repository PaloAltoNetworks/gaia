package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerReportLicenseTypeValue represents the possible values for attribute "licenseType".
type EnforcerReportLicenseTypeValue string

const (
	// EnforcerReportLicenseTypeContainer represents the value Container.
	EnforcerReportLicenseTypeContainer EnforcerReportLicenseTypeValue = "Container"

	// EnforcerReportLicenseTypeHost represents the value Host.
	EnforcerReportLicenseTypeHost EnforcerReportLicenseTypeValue = "Host"
)

// EnforcerReportIdentity represents the Identity of the object.
var EnforcerReportIdentity = elemental.Identity{
	Name:     "enforcerreport",
	Category: "enforcerreports",
	Package:  "zack",
	Private:  false,
}

// EnforcerReportsList represents a list of EnforcerReports
type EnforcerReportsList []*EnforcerReport

// Identity returns the identity of the objects in the list.
func (o EnforcerReportsList) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Copy returns a pointer to a copy the EnforcerReportsList.
func (o EnforcerReportsList) Copy() elemental.Identifiables {

	copy := append(EnforcerReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerReportsList.
func (o EnforcerReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToSparse returns the EnforcerReportsList converted to SparseEnforcerReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerReport)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerReportsList) Version() int {

	return 1
}

// EnforcerReport represents the model of a enforcerreport
type EnforcerReport struct {
	// Total CPU utilization of the enforcer as a percentage of vCPUs.
	CPULoad float64 `json:"CPULoad,omitempty" msgpack:"CPULoad,omitempty" bson:"cpuload" mapstructure:"CPULoad,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Type of license for this enforcer.
	LicenseType EnforcerReportLicenseTypeValue `json:"licenseType,omitempty" msgpack:"licenseType,omitempty" bson:"licensetype" mapstructure:"licenseType,omitempty"`

	// Total resident memory used by the enforcer in bytes.
	Memory int `json:"memory,omitempty" msgpack:"memory,omitempty" bson:"memory" mapstructure:"memory,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the enforcer.
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace of the enforcer.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Number of active processes of the enforcer.
	Processes int `json:"processes,omitempty" msgpack:"processes,omitempty" bson:"processes" mapstructure:"processes,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerReport returns a new *EnforcerReport
func NewEnforcerReport() *EnforcerReport {

	return &EnforcerReport{
		ModelVersion:  1,
		LicenseType:   EnforcerReportLicenseTypeHost,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerReport) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcerReport{}

	s.CPULoad = o.CPULoad
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.EnforcerID = o.EnforcerID
	s.LicenseType = o.LicenseType
	s.Memory = o.Memory
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.Processes = o.Processes
	s.Timestamp = o.Timestamp
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcerReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.CPULoad = s.CPULoad
	o.ID = s.ID.Hex()
	o.EnforcerID = s.EnforcerID
	o.LicenseType = s.LicenseType
	o.Memory = s.Memory
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.Processes = s.Processes
	o.Timestamp = s.Timestamp
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EnforcerReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerReport) BleveType() string {

	return "enforcerreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerReport) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// Doc returns the documentation for the object
func (o *EnforcerReport) Doc() string {

	return `Post a new enforcer statistics report.`
}

func (o *EnforcerReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *EnforcerReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *EnforcerReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *EnforcerReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *EnforcerReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *EnforcerReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *EnforcerReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerReport{
			CPULoad:       &o.CPULoad,
			ID:            &o.ID,
			EnforcerID:    &o.EnforcerID,
			LicenseType:   &o.LicenseType,
			Memory:        &o.Memory,
			MigrationsLog: &o.MigrationsLog,
			Name:          &o.Name,
			Namespace:     &o.Namespace,
			Processes:     &o.Processes,
			Timestamp:     &o.Timestamp,
			ZHash:         &o.ZHash,
			Zone:          &o.Zone,
		}
	}

	sp := &SparseEnforcerReport{}
	for _, f := range fields {
		switch f {
		case "CPULoad":
			sp.CPULoad = &(o.CPULoad)
		case "ID":
			sp.ID = &(o.ID)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "licenseType":
			sp.LicenseType = &(o.LicenseType)
		case "memory":
			sp.Memory = &(o.Memory)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "processes":
			sp.Processes = &(o.Processes)
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

// Patch apply the non nil value of a *SparseEnforcerReport to the object.
func (o *EnforcerReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerReport)
	if so.CPULoad != nil {
		o.CPULoad = *so.CPULoad
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.LicenseType != nil {
		o.LicenseType = *so.LicenseType
	}
	if so.Memory != nil {
		o.Memory = *so.Memory
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Processes != nil {
		o.Processes = *so.Processes
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

// DeepCopy returns a deep copy if the EnforcerReport.
func (o *EnforcerReport) DeepCopy() *EnforcerReport {

	if o == nil {
		return nil
	}

	out := &EnforcerReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerReport.
func (o *EnforcerReport) DeepCopyInto(out *EnforcerReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerReport: %s", err))
	}

	*out = *target.(*EnforcerReport)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("licenseType", string(o.LicenseType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("licenseType", string(o.LicenseType), []string{"Host", "Container"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	// Custom object validation.
	if err := ValidateEnforcerReport(o); err != nil {
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
func (*EnforcerReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "CPULoad":
		return o.CPULoad
	case "ID":
		return o.ID
	case "enforcerID":
		return o.EnforcerID
	case "licenseType":
		return o.LicenseType
	case "memory":
		return o.Memory
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "processes":
		return o.Processes
	case "timestamp":
		return o.Timestamp
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// EnforcerReportAttributesMap represents the map of attribute for EnforcerReport.
var EnforcerReportAttributesMap = map[string]elemental.AttributeSpecification{
	"CPULoad": {
		AllowedChoices: []string{},
		ConvertedName:  "CPULoad",
		Description:    `Total CPU utilization of the enforcer as a percentage of vCPUs.`,
		Exposed:        true,
		Name:           "CPULoad",
		Stored:         true,
		Type:           "float",
	},
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
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"LicenseType": {
		AllowedChoices: []string{"Host", "Container"},
		ConvertedName:  "LicenseType",
		DefaultValue:   EnforcerReportLicenseTypeHost,
		Description:    `Type of license for this enforcer.`,
		Exposed:        true,
		Name:           "licenseType",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Memory": {
		AllowedChoices: []string{},
		ConvertedName:  "Memory",
		Description:    `Total resident memory used by the enforcer in bytes.`,
		Exposed:        true,
		Name:           "memory",
		Stored:         true,
		Type:           "integer",
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
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the enforcer.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Processes": {
		AllowedChoices: []string{},
		ConvertedName:  "Processes",
		Description:    `Number of active processes of the enforcer.`,
		Exposed:        true,
		Name:           "processes",
		Stored:         true,
		Type:           "integer",
	},
	"Timestamp": {
		AllowedChoices: []string{},
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

// EnforcerReportLowerCaseAttributesMap represents the map of attribute for EnforcerReport.
var EnforcerReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"cpuload": {
		AllowedChoices: []string{},
		ConvertedName:  "CPULoad",
		Description:    `Total CPU utilization of the enforcer as a percentage of vCPUs.`,
		Exposed:        true,
		Name:           "CPULoad",
		Stored:         true,
		Type:           "float",
	},
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
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"licensetype": {
		AllowedChoices: []string{"Host", "Container"},
		ConvertedName:  "LicenseType",
		DefaultValue:   EnforcerReportLicenseTypeHost,
		Description:    `Type of license for this enforcer.`,
		Exposed:        true,
		Name:           "licenseType",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"memory": {
		AllowedChoices: []string{},
		ConvertedName:  "Memory",
		Description:    `Total resident memory used by the enforcer in bytes.`,
		Exposed:        true,
		Name:           "memory",
		Stored:         true,
		Type:           "integer",
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
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the enforcer.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"processes": {
		AllowedChoices: []string{},
		ConvertedName:  "Processes",
		Description:    `Number of active processes of the enforcer.`,
		Exposed:        true,
		Name:           "processes",
		Stored:         true,
		Type:           "integer",
	},
	"timestamp": {
		AllowedChoices: []string{},
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

// SparseEnforcerReportsList represents a list of SparseEnforcerReports
type SparseEnforcerReportsList []*SparseEnforcerReport

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerReportsList) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerReportsList.
func (o SparseEnforcerReportsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerReportsList.
func (o SparseEnforcerReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToPlain returns the SparseEnforcerReportsList converted to EnforcerReportsList.
func (o SparseEnforcerReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerReportsList) Version() int {

	return 1
}

// SparseEnforcerReport represents the sparse version of a enforcerreport.
type SparseEnforcerReport struct {
	// Total CPU utilization of the enforcer as a percentage of vCPUs.
	CPULoad *float64 `json:"CPULoad,omitempty" msgpack:"CPULoad,omitempty" bson:"cpuload,omitempty" mapstructure:"CPULoad,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Type of license for this enforcer.
	LicenseType *EnforcerReportLicenseTypeValue `json:"licenseType,omitempty" msgpack:"licenseType,omitempty" bson:"licensetype,omitempty" mapstructure:"licenseType,omitempty"`

	// Total resident memory used by the enforcer in bytes.
	Memory *int `json:"memory,omitempty" msgpack:"memory,omitempty" bson:"memory,omitempty" mapstructure:"memory,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the enforcer.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace of the enforcer.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Number of active processes of the enforcer.
	Processes *int `json:"processes,omitempty" msgpack:"processes,omitempty" bson:"processes,omitempty" mapstructure:"processes,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerReport returns a new  SparseEnforcerReport.
func NewSparseEnforcerReport() *SparseEnforcerReport {
	return &SparseEnforcerReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerReport) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcerReport{}

	if o.CPULoad != nil {
		s.CPULoad = o.CPULoad
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.LicenseType != nil {
		s.LicenseType = o.LicenseType
	}
	if o.Memory != nil {
		s.Memory = o.Memory
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Processes != nil {
		s.Processes = o.Processes
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
func (o *SparseEnforcerReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcerReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.CPULoad != nil {
		o.CPULoad = s.CPULoad
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.LicenseType != nil {
		o.LicenseType = s.LicenseType
	}
	if s.Memory != nil {
		o.Memory = s.Memory
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Processes != nil {
		o.Processes = s.Processes
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
func (o *SparseEnforcerReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerReport) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerReport()
	if o.CPULoad != nil {
		out.CPULoad = *o.CPULoad
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.LicenseType != nil {
		out.LicenseType = *o.LicenseType
	}
	if o.Memory != nil {
		out.Memory = *o.Memory
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Processes != nil {
		out.Processes = *o.Processes
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
func (o *SparseEnforcerReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseEnforcerReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEnforcerReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEnforcerReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEnforcerReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEnforcerReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEnforcerReport.
func (o *SparseEnforcerReport) DeepCopy() *SparseEnforcerReport {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerReport.
func (o *SparseEnforcerReport) DeepCopyInto(out *SparseEnforcerReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerReport: %s", err))
	}

	*out = *target.(*SparseEnforcerReport)
}

type mongoAttributesEnforcerReport struct {
	CPULoad       float64                        `bson:"cpuload,omitempty"`
	ID            bson.ObjectId                  `bson:"_id,omitempty"`
	EnforcerID    string                         `bson:"enforcerid,omitempty"`
	LicenseType   EnforcerReportLicenseTypeValue `bson:"licensetype,omitempty"`
	Memory        int                            `bson:"memory,omitempty"`
	MigrationsLog map[string]string              `bson:"migrationslog,omitempty"`
	Name          string                         `bson:"name,omitempty"`
	Namespace     string                         `bson:"namespace,omitempty"`
	Processes     int                            `bson:"processes,omitempty"`
	Timestamp     time.Time                      `bson:"timestamp,omitempty"`
	ZHash         int                            `bson:"zhash"`
	Zone          int                            `bson:"zone"`
}
type mongoAttributesSparseEnforcerReport struct {
	CPULoad       *float64                        `bson:"cpuload,omitempty"`
	ID            bson.ObjectId                   `bson:"_id,omitempty"`
	EnforcerID    *string                         `bson:"enforcerid,omitempty"`
	LicenseType   *EnforcerReportLicenseTypeValue `bson:"licensetype,omitempty"`
	Memory        *int                            `bson:"memory,omitempty"`
	MigrationsLog *map[string]string              `bson:"migrationslog,omitempty"`
	Name          *string                         `bson:"name,omitempty"`
	Namespace     *string                         `bson:"namespace,omitempty"`
	Processes     *int                            `bson:"processes,omitempty"`
	Timestamp     *time.Time                      `bson:"timestamp,omitempty"`
	ZHash         *int                            `bson:"zhash,omitempty"`
	Zone          *int                            `bson:"zone,omitempty"`
}
