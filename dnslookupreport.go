package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DNSLookupReportActionValue represents the possible values for attribute "action".
type DNSLookupReportActionValue string

const (
	// DNSLookupReportActionAccept represents the value Accept.
	DNSLookupReportActionAccept DNSLookupReportActionValue = "Accept"

	// DNSLookupReportActionFail represents the value Fail.
	DNSLookupReportActionFail DNSLookupReportActionValue = "Fail"

	// DNSLookupReportActionReject represents the value Reject.
	DNSLookupReportActionReject DNSLookupReportActionValue = "Reject"

	// DNSLookupReportActionResolve represents the value Resolve.
	DNSLookupReportActionResolve DNSLookupReportActionValue = "Resolve"
)

// DNSLookupReportIdentity represents the Identity of the object.
var DNSLookupReportIdentity = elemental.Identity{
	Name:     "dnslookupreport",
	Category: "dnslookupreports",
	Package:  "zack",
	Private:  false,
}

// DNSLookupReportsList represents a list of DNSLookupReports
type DNSLookupReportsList []*DNSLookupReport

// Identity returns the identity of the objects in the list.
func (o DNSLookupReportsList) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Copy returns a pointer to a copy the DNSLookupReportsList.
func (o DNSLookupReportsList) Copy() elemental.Identifiables {

	copy := append(DNSLookupReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DNSLookupReportsList.
func (o DNSLookupReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DNSLookupReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DNSLookupReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DNSLookupReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DNSLookupReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToSparse returns the DNSLookupReportsList converted to SparseDNSLookupReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DNSLookupReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDNSLookupReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDNSLookupReport)
	}

	return out
}

// Version returns the version of the content.
func (o DNSLookupReportsList) Version() int {

	return 1
}

// DNSLookupReport represents the model of a dnslookupreport
type DNSLookupReport struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Action of the DNS request.
	Action DNSLookupReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"a,omitempty" mapstructure:"action,omitempty"`

	// The local ID set by enforcer TODO.
	ClientLocalID string `json:"clientLocalID,omitempty" msgpack:"clientLocalID,omitempty" bson:"m,omitempty" mapstructure:"clientLocalID,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"b,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"c,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the PU.
	ProcessingUnitID string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"d,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU. This is deprecated. Use `namespace` instead.
	ProcessingUnitNamespace string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"e,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set when the lookup fails. It specifies the reason for the
	// failure.
	Reason string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"f,omitempty" mapstructure:"reason,omitempty"`

	// CNAME aliases.
	ResolvedCNAMEs []string `json:"resolvedCNAMEs,omitempty" msgpack:"resolvedCNAMEs,omitempty" bson:"k,omitempty" mapstructure:"resolvedCNAMEs,omitempty"`

	// resolved IP addresses.
	ResolvedIPs []string `json:"resolvedIPs,omitempty" msgpack:"resolvedIPs,omitempty" bson:"l,omitempty" mapstructure:"resolvedIPs,omitempty"`

	// name used for DNS resolution.
	ResolvedName string `json:"resolvedName,omitempty" msgpack:"resolvedName,omitempty" bson:"g,omitempty" mapstructure:"resolvedName,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"h,omitempty" mapstructure:"sourceIP,omitempty"`

	// Time and date of the log.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"i,omitempty" mapstructure:"timestamp,omitempty"`

	// Number of times the client saw this activity.
	Value int `json:"value,omitempty" msgpack:"value,omitempty" bson:"j,omitempty" mapstructure:"value,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDNSLookupReport returns a new *DNSLookupReport
func NewDNSLookupReport() *DNSLookupReport {

	return &DNSLookupReport{
		ModelVersion:   1,
		ResolvedIPs:    []string{},
		ResolvedCNAMEs: []string{},
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *DNSLookupReport) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DNSLookupReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DNSLookupReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DNSLookupReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDNSLookupReport{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Action = o.Action
	s.ClientLocalID = o.ClientLocalID
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.ProcessingUnitID = o.ProcessingUnitID
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.Reason = o.Reason
	s.ResolvedCNAMEs = o.ResolvedCNAMEs
	s.ResolvedIPs = o.ResolvedIPs
	s.ResolvedName = o.ResolvedName
	s.SourceIP = o.SourceIP
	s.Timestamp = o.Timestamp
	s.Value = o.Value
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DNSLookupReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDNSLookupReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Action = s.Action
	o.ClientLocalID = s.ClientLocalID
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.ProcessingUnitID = s.ProcessingUnitID
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.Reason = s.Reason
	o.ResolvedCNAMEs = s.ResolvedCNAMEs
	o.ResolvedIPs = s.ResolvedIPs
	o.ResolvedName = s.ResolvedName
	o.SourceIP = s.SourceIP
	o.Timestamp = s.Timestamp
	o.Value = s.Value
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DNSLookupReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DNSLookupReport) BleveType() string {

	return "dnslookupreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DNSLookupReport) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// Doc returns the documentation for the object
func (o *DNSLookupReport) Doc() string {

	return `A DNS lookup report is used to report a DNS lookup that is happening on
behalf of a processing unit. If the DNS server is on the standard UDP port 53
then the enforcer can proxy the DNS traffic and make a report. The report
indicate whether or not the lookup was successful.`
}

func (o *DNSLookupReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *DNSLookupReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *DNSLookupReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *DNSLookupReport) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *DNSLookupReport) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *DNSLookupReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *DNSLookupReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *DNSLookupReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *DNSLookupReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DNSLookupReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDNSLookupReport{
			ID:                      &o.ID,
			Action:                  &o.Action,
			ClientLocalID:           &o.ClientLocalID,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			MigrationsLog:           &o.MigrationsLog,
			Namespace:               &o.Namespace,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Reason:                  &o.Reason,
			ResolvedCNAMEs:          &o.ResolvedCNAMEs,
			ResolvedIPs:             &o.ResolvedIPs,
			ResolvedName:            &o.ResolvedName,
			SourceIP:                &o.SourceIP,
			Timestamp:               &o.Timestamp,
			Value:                   &o.Value,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparseDNSLookupReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "action":
			sp.Action = &(o.Action)
		case "clientLocalID":
			sp.ClientLocalID = &(o.ClientLocalID)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "reason":
			sp.Reason = &(o.Reason)
		case "resolvedCNAMEs":
			sp.ResolvedCNAMEs = &(o.ResolvedCNAMEs)
		case "resolvedIPs":
			sp.ResolvedIPs = &(o.ResolvedIPs)
		case "resolvedName":
			sp.ResolvedName = &(o.ResolvedName)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDNSLookupReport to the object.
func (o *DNSLookupReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDNSLookupReport)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.ClientLocalID != nil {
		o.ClientLocalID = *so.ClientLocalID
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
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Reason != nil {
		o.Reason = *so.Reason
	}
	if so.ResolvedCNAMEs != nil {
		o.ResolvedCNAMEs = *so.ResolvedCNAMEs
	}
	if so.ResolvedIPs != nil {
		o.ResolvedIPs = *so.ResolvedIPs
	}
	if so.ResolvedName != nil {
		o.ResolvedName = *so.ResolvedName
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the DNSLookupReport.
func (o *DNSLookupReport) DeepCopy() *DNSLookupReport {

	if o == nil {
		return nil
	}

	out := &DNSLookupReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DNSLookupReport.
func (o *DNSLookupReport) DeepCopyInto(out *DNSLookupReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DNSLookupReport: %s", err))
	}

	*out = *target.(*DNSLookupReport)
}

// Validate valides the current information stored into the structure.
func (o *DNSLookupReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Fail", "Reject", "Resolve"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateOptionalCIDRorIPList("resolvedIPs", o.ResolvedIPs); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("resolvedName", o.ResolvedName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceIP", o.SourceIP); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("value", o.Value); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	// Custom object validation.
	if err := ValidateDNSLookupReport(o); err != nil {
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
func (*DNSLookupReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DNSLookupReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DNSLookupReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DNSLookupReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DNSLookupReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DNSLookupReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "clientLocalID":
		return o.ClientLocalID
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "reason":
		return o.Reason
	case "resolvedCNAMEs":
		return o.ResolvedCNAMEs
	case "resolvedIPs":
		return o.ResolvedIPs
	case "resolvedName":
		return o.ResolvedName
	case "sourceIP":
		return o.SourceIP
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// DNSLookupReportAttributesMap represents the map of attribute for DNSLookupReport.
var DNSLookupReportAttributesMap = map[string]elemental.AttributeSpecification{
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
		AllowedChoices: []string{"Accept", "Fail", "Reject", "Resolve"},
		BSONFieldName:  "a",
		ConvertedName:  "Action",
		Description:    `Action of the DNS request.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"ClientLocalID": {
		AllowedChoices: []string{},
		BSONFieldName:  "m",
		ConvertedName:  "ClientLocalID",
		Description:    `The local ID set by enforcer TODO.`,
		Exposed:        true,
		Name:           "clientLocalID",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
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
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the PU.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "ProcessingUnitNamespace",
		Deprecated:     true,
		Description:    `Namespace of the PU. This is deprecated. Use ` + "`" + `namespace` + "`" + ` instead.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Reason": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "Reason",
		Description: `This field is only set when the lookup fails. It specifies the reason for the
failure.`,
		Exposed: true,
		Name:    "reason",
		Stored:  true,
		Type:    "string",
	},
	"ResolvedCNAMEs": {
		AllowedChoices: []string{},
		BSONFieldName:  "k",
		ConvertedName:  "ResolvedCNAMEs",
		Description:    `CNAME aliases.`,
		Exposed:        true,
		Name:           "resolvedCNAMEs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ResolvedIPs": {
		AllowedChoices: []string{},
		BSONFieldName:  "l",
		ConvertedName:  "ResolvedIPs",
		Description:    `resolved IP addresses.`,
		Exposed:        true,
		Name:           "resolvedIPs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ResolvedName": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "ResolvedName",
		Description:    `name used for DNS resolution.`,
		Exposed:        true,
		Name:           "resolvedName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "h",
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "i",
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Value": {
		AllowedChoices: []string{},
		BSONFieldName:  "j",
		ConvertedName:  "Value",
		Description:    `Number of times the client saw this activity.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "integer",
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

// DNSLookupReportLowerCaseAttributesMap represents the map of attribute for DNSLookupReport.
var DNSLookupReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		AllowedChoices: []string{"Accept", "Fail", "Reject", "Resolve"},
		BSONFieldName:  "a",
		ConvertedName:  "Action",
		Description:    `Action of the DNS request.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"clientlocalid": {
		AllowedChoices: []string{},
		BSONFieldName:  "m",
		ConvertedName:  "ClientLocalID",
		Description:    `The local ID set by enforcer TODO.`,
		Exposed:        true,
		Name:           "clientLocalID",
		Stored:         true,
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		BSONFieldName:  "b",
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "c",
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
	"processingunitid": {
		AllowedChoices: []string{},
		BSONFieldName:  "d",
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the PU.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "e",
		ConvertedName:  "ProcessingUnitNamespace",
		Deprecated:     true,
		Description:    `Namespace of the PU. This is deprecated. Use ` + "`" + `namespace` + "`" + ` instead.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"reason": {
		AllowedChoices: []string{},
		BSONFieldName:  "f",
		ConvertedName:  "Reason",
		Description: `This field is only set when the lookup fails. It specifies the reason for the
failure.`,
		Exposed: true,
		Name:    "reason",
		Stored:  true,
		Type:    "string",
	},
	"resolvedcnames": {
		AllowedChoices: []string{},
		BSONFieldName:  "k",
		ConvertedName:  "ResolvedCNAMEs",
		Description:    `CNAME aliases.`,
		Exposed:        true,
		Name:           "resolvedCNAMEs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"resolvedips": {
		AllowedChoices: []string{},
		BSONFieldName:  "l",
		ConvertedName:  "ResolvedIPs",
		Description:    `resolved IP addresses.`,
		Exposed:        true,
		Name:           "resolvedIPs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"resolvedname": {
		AllowedChoices: []string{},
		BSONFieldName:  "g",
		ConvertedName:  "ResolvedName",
		Description:    `name used for DNS resolution.`,
		Exposed:        true,
		Name:           "resolvedName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourceip": {
		AllowedChoices: []string{},
		BSONFieldName:  "h",
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		BSONFieldName:  "i",
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"value": {
		AllowedChoices: []string{},
		BSONFieldName:  "j",
		ConvertedName:  "Value",
		Description:    `Number of times the client saw this activity.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "integer",
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

// SparseDNSLookupReportsList represents a list of SparseDNSLookupReports
type SparseDNSLookupReportsList []*SparseDNSLookupReport

// Identity returns the identity of the objects in the list.
func (o SparseDNSLookupReportsList) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Copy returns a pointer to a copy the SparseDNSLookupReportsList.
func (o SparseDNSLookupReportsList) Copy() elemental.Identifiables {

	copy := append(SparseDNSLookupReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDNSLookupReportsList.
func (o SparseDNSLookupReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDNSLookupReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDNSLookupReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDNSLookupReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDNSLookupReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToPlain returns the SparseDNSLookupReportsList converted to DNSLookupReportsList.
func (o SparseDNSLookupReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDNSLookupReportsList) Version() int {

	return 1
}

// SparseDNSLookupReport represents the sparse version of a dnslookupreport.
type SparseDNSLookupReport struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Action of the DNS request.
	Action *DNSLookupReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"a,omitempty" mapstructure:"action,omitempty"`

	// The local ID set by enforcer TODO.
	ClientLocalID *string `json:"clientLocalID,omitempty" msgpack:"clientLocalID,omitempty" bson:"m,omitempty" mapstructure:"clientLocalID,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"b,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"c,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// ID of the PU.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"d,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU. This is deprecated. Use `namespace` instead.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"e,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set when the lookup fails. It specifies the reason for the
	// failure.
	Reason *string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"f,omitempty" mapstructure:"reason,omitempty"`

	// CNAME aliases.
	ResolvedCNAMEs *[]string `json:"resolvedCNAMEs,omitempty" msgpack:"resolvedCNAMEs,omitempty" bson:"k,omitempty" mapstructure:"resolvedCNAMEs,omitempty"`

	// resolved IP addresses.
	ResolvedIPs *[]string `json:"resolvedIPs,omitempty" msgpack:"resolvedIPs,omitempty" bson:"l,omitempty" mapstructure:"resolvedIPs,omitempty"`

	// name used for DNS resolution.
	ResolvedName *string `json:"resolvedName,omitempty" msgpack:"resolvedName,omitempty" bson:"g,omitempty" mapstructure:"resolvedName,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"h,omitempty" mapstructure:"sourceIP,omitempty"`

	// Time and date of the log.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"i,omitempty" mapstructure:"timestamp,omitempty"`

	// Number of times the client saw this activity.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"j,omitempty" mapstructure:"value,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDNSLookupReport returns a new  SparseDNSLookupReport.
func NewSparseDNSLookupReport() *SparseDNSLookupReport {
	return &SparseDNSLookupReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDNSLookupReport) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDNSLookupReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDNSLookupReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDNSLookupReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDNSLookupReport{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Action != nil {
		s.Action = o.Action
	}
	if o.ClientLocalID != nil {
		s.ClientLocalID = o.ClientLocalID
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
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.ProcessingUnitID != nil {
		s.ProcessingUnitID = o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	}
	if o.Reason != nil {
		s.Reason = o.Reason
	}
	if o.ResolvedCNAMEs != nil {
		s.ResolvedCNAMEs = o.ResolvedCNAMEs
	}
	if o.ResolvedIPs != nil {
		s.ResolvedIPs = o.ResolvedIPs
	}
	if o.ResolvedName != nil {
		s.ResolvedName = o.ResolvedName
	}
	if o.SourceIP != nil {
		s.SourceIP = o.SourceIP
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
	}
	if o.Value != nil {
		s.Value = o.Value
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
func (o *SparseDNSLookupReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDNSLookupReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Action != nil {
		o.Action = s.Action
	}
	if s.ClientLocalID != nil {
		o.ClientLocalID = s.ClientLocalID
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
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.ProcessingUnitID != nil {
		o.ProcessingUnitID = s.ProcessingUnitID
	}
	if s.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	}
	if s.Reason != nil {
		o.Reason = s.Reason
	}
	if s.ResolvedCNAMEs != nil {
		o.ResolvedCNAMEs = s.ResolvedCNAMEs
	}
	if s.ResolvedIPs != nil {
		o.ResolvedIPs = s.ResolvedIPs
	}
	if s.ResolvedName != nil {
		o.ResolvedName = s.ResolvedName
	}
	if s.SourceIP != nil {
		o.SourceIP = s.SourceIP
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
	}
	if s.Value != nil {
		o.Value = s.Value
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
func (o *SparseDNSLookupReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDNSLookupReport) ToPlain() elemental.PlainIdentifiable {

	out := NewDNSLookupReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.ClientLocalID != nil {
		out.ClientLocalID = *o.ClientLocalID
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
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Reason != nil {
		out.Reason = *o.Reason
	}
	if o.ResolvedCNAMEs != nil {
		out.ResolvedCNAMEs = *o.ResolvedCNAMEs
	}
	if o.ResolvedIPs != nil {
		out.ResolvedIPs = *o.ResolvedIPs
	}
	if o.ResolvedName != nil {
		out.ResolvedName = *o.ResolvedName
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Value != nil {
		out.Value = *o.Value
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
func (o *SparseDNSLookupReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseDNSLookupReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseDNSLookupReport) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseDNSLookupReport) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseDNSLookupReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseDNSLookupReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseDNSLookupReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseDNSLookupReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseDNSLookupReport.
func (o *SparseDNSLookupReport) DeepCopy() *SparseDNSLookupReport {

	if o == nil {
		return nil
	}

	out := &SparseDNSLookupReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDNSLookupReport.
func (o *SparseDNSLookupReport) DeepCopyInto(out *SparseDNSLookupReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDNSLookupReport: %s", err))
	}

	*out = *target.(*SparseDNSLookupReport)
}

type mongoAttributesDNSLookupReport struct {
	ID                      bson.ObjectId              `bson:"_id,omitempty"`
	Action                  DNSLookupReportActionValue `bson:"a,omitempty"`
	ClientLocalID           string                     `bson:"m,omitempty"`
	EnforcerID              string                     `bson:"b,omitempty"`
	EnforcerNamespace       string                     `bson:"c,omitempty"`
	MigrationsLog           map[string]string          `bson:"migrationslog,omitempty"`
	Namespace               string                     `bson:"namespace"`
	ProcessingUnitID        string                     `bson:"d,omitempty"`
	ProcessingUnitNamespace string                     `bson:"e,omitempty"`
	Reason                  string                     `bson:"f,omitempty"`
	ResolvedCNAMEs          []string                   `bson:"k,omitempty"`
	ResolvedIPs             []string                   `bson:"l,omitempty"`
	ResolvedName            string                     `bson:"g,omitempty"`
	SourceIP                string                     `bson:"h,omitempty"`
	Timestamp               time.Time                  `bson:"i,omitempty"`
	Value                   int                        `bson:"j,omitempty"`
	ZHash                   int                        `bson:"zhash"`
	Zone                    int                        `bson:"zone"`
}
type mongoAttributesSparseDNSLookupReport struct {
	ID                      bson.ObjectId               `bson:"_id,omitempty"`
	Action                  *DNSLookupReportActionValue `bson:"a,omitempty"`
	ClientLocalID           *string                     `bson:"m,omitempty"`
	EnforcerID              *string                     `bson:"b,omitempty"`
	EnforcerNamespace       *string                     `bson:"c,omitempty"`
	MigrationsLog           *map[string]string          `bson:"migrationslog,omitempty"`
	Namespace               *string                     `bson:"namespace,omitempty"`
	ProcessingUnitID        *string                     `bson:"d,omitempty"`
	ProcessingUnitNamespace *string                     `bson:"e,omitempty"`
	Reason                  *string                     `bson:"f,omitempty"`
	ResolvedCNAMEs          *[]string                   `bson:"k,omitempty"`
	ResolvedIPs             *[]string                   `bson:"l,omitempty"`
	ResolvedName            *string                     `bson:"g,omitempty"`
	SourceIP                *string                     `bson:"h,omitempty"`
	Timestamp               *time.Time                  `bson:"i,omitempty"`
	Value                   *int                        `bson:"j,omitempty"`
	ZHash                   *int                        `bson:"zhash,omitempty"`
	Zone                    *int                        `bson:"zone,omitempty"`
}
