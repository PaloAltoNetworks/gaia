package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ConnectionExceptionReportServiceTypeValue represents the possible values for attribute "serviceType".
type ConnectionExceptionReportServiceTypeValue string

const (
	// ConnectionExceptionReportServiceTypeHTTP represents the value HTTP.
	ConnectionExceptionReportServiceTypeHTTP ConnectionExceptionReportServiceTypeValue = "HTTP"

	// ConnectionExceptionReportServiceTypeL3 represents the value L3.
	ConnectionExceptionReportServiceTypeL3 ConnectionExceptionReportServiceTypeValue = "L3"

	// ConnectionExceptionReportServiceTypeTCP represents the value TCP.
	ConnectionExceptionReportServiceTypeTCP ConnectionExceptionReportServiceTypeValue = "TCP"
)

// ConnectionExceptionReportStateValue represents the possible values for attribute "state".
type ConnectionExceptionReportStateValue string

const (
	// ConnectionExceptionReportStateAckTransmitted represents the value AckTransmitted.
	ConnectionExceptionReportStateAckTransmitted ConnectionExceptionReportStateValue = "AckTransmitted"

	// ConnectionExceptionReportStateSynAckTransmitted represents the value SynAckTransmitted.
	ConnectionExceptionReportStateSynAckTransmitted ConnectionExceptionReportStateValue = "SynAckTransmitted"

	// ConnectionExceptionReportStateSynTransmitted represents the value SynTransmitted.
	ConnectionExceptionReportStateSynTransmitted ConnectionExceptionReportStateValue = "SynTransmitted"

	// ConnectionExceptionReportStateUnknown represents the value Unknown.
	ConnectionExceptionReportStateUnknown ConnectionExceptionReportStateValue = "Unknown"
)

// ConnectionExceptionReportIdentity represents the Identity of the object.
var ConnectionExceptionReportIdentity = elemental.Identity{
	Name:     "connectionexceptionreport",
	Category: "connectionexceptionreports",
	Package:  "zack",
	Private:  false,
}

// ConnectionExceptionReportsList represents a list of ConnectionExceptionReports
type ConnectionExceptionReportsList []*ConnectionExceptionReport

// Identity returns the identity of the objects in the list.
func (o ConnectionExceptionReportsList) Identity() elemental.Identity {

	return ConnectionExceptionReportIdentity
}

// Copy returns a pointer to a copy the ConnectionExceptionReportsList.
func (o ConnectionExceptionReportsList) Copy() elemental.Identifiables {

	copy := append(ConnectionExceptionReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ConnectionExceptionReportsList.
func (o ConnectionExceptionReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ConnectionExceptionReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ConnectionExceptionReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ConnectionExceptionReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ConnectionExceptionReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ConnectionExceptionReportsList converted to SparseConnectionExceptionReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ConnectionExceptionReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseConnectionExceptionReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseConnectionExceptionReport)
	}

	return out
}

// Version returns the version of the content.
func (o ConnectionExceptionReportsList) Version() int {

	return 1
}

// ConnectionExceptionReport represents the model of a connectionexceptionreport
type ConnectionExceptionReport struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Identifier of the destination controller. This should be set in
	// SynAckTransmitted state.
	DestinationController string `json:"destinationController,omitempty" msgpack:"destinationController,omitempty" bson:"destinationcontroller" mapstructure:"destinationController,omitempty"`

	// Destination IP address.
	DestinationIP string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"destinationip" mapstructure:"destinationIP,omitempty"`

	// Port of the destination.
	DestinationPort int `json:"destinationPort,omitempty" msgpack:"destinationPort,omitempty" bson:"destinationport" mapstructure:"destinationPort,omitempty"`

	// ID of the destination processing unit. This should be set in SynAckTransmitted
	// state.
	DestinationProcessingUnitID string `json:"destinationProcessingUnitID,omitempty" msgpack:"destinationProcessingUnitID,omitempty" bson:"destinationprocessingunitid" mapstructure:"destinationProcessingUnitID,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace of the processing unit that encountered this exception.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the processing unit encountered this exception.
	ProcessingUnitID string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"processingunitid" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit encountered this exception.
	ProcessingUnitNamespace string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"processingunitnamespace" mapstructure:"processingUnitNamespace,omitempty"`

	// Protocol number.
	Protocol int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol" mapstructure:"protocol,omitempty"`

	// It specifies the reason for the exception.
	Reason string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"reason" mapstructure:"reason,omitempty"`

	// Type of the service.
	ServiceType ConnectionExceptionReportServiceTypeValue `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"servicetype" mapstructure:"serviceType,omitempty"`

	// Source IP address.
	SourceIP string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"sourceip" mapstructure:"sourceIP,omitempty"`

	// Represents the current state this report was generated.
	State ConnectionExceptionReportStateValue `json:"state" msgpack:"state" bson:"state" mapstructure:"state,omitempty"`

	// Time and date of the report.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// Number of packets hit.
	Value int `json:"value,omitempty" msgpack:"value,omitempty" bson:"value" mapstructure:"value,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewConnectionExceptionReport returns a new *ConnectionExceptionReport
func NewConnectionExceptionReport() *ConnectionExceptionReport {

	return &ConnectionExceptionReport{
		ModelVersion:  1,
		ServiceType:   ConnectionExceptionReportServiceTypeL3,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *ConnectionExceptionReport) Identity() elemental.Identity {

	return ConnectionExceptionReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ConnectionExceptionReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ConnectionExceptionReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionExceptionReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesConnectionExceptionReport{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.DestinationController = o.DestinationController
	s.DestinationIP = o.DestinationIP
	s.DestinationPort = o.DestinationPort
	s.DestinationProcessingUnitID = o.DestinationProcessingUnitID
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.ProcessingUnitID = o.ProcessingUnitID
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.Protocol = o.Protocol
	s.Reason = o.Reason
	s.ServiceType = o.ServiceType
	s.SourceIP = o.SourceIP
	s.State = o.State
	s.Timestamp = o.Timestamp
	s.Value = o.Value
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionExceptionReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesConnectionExceptionReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.DestinationController = s.DestinationController
	o.DestinationIP = s.DestinationIP
	o.DestinationPort = s.DestinationPort
	o.DestinationProcessingUnitID = s.DestinationProcessingUnitID
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.ProcessingUnitID = s.ProcessingUnitID
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.Protocol = s.Protocol
	o.Reason = s.Reason
	o.ServiceType = s.ServiceType
	o.SourceIP = s.SourceIP
	o.State = s.State
	o.Timestamp = s.Timestamp
	o.Value = s.Value
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ConnectionExceptionReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ConnectionExceptionReport) BleveType() string {

	return "connectionexceptionreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ConnectionExceptionReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ConnectionExceptionReport) Doc() string {

	return `Post a new flow log.`
}

func (o *ConnectionExceptionReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *ConnectionExceptionReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *ConnectionExceptionReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *ConnectionExceptionReport) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ConnectionExceptionReport) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *ConnectionExceptionReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ConnectionExceptionReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ConnectionExceptionReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ConnectionExceptionReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ConnectionExceptionReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseConnectionExceptionReport{
			ID:                          &o.ID,
			DestinationController:       &o.DestinationController,
			DestinationIP:               &o.DestinationIP,
			DestinationPort:             &o.DestinationPort,
			DestinationProcessingUnitID: &o.DestinationProcessingUnitID,
			EnforcerID:                  &o.EnforcerID,
			EnforcerNamespace:           &o.EnforcerNamespace,
			MigrationsLog:               &o.MigrationsLog,
			Namespace:                   &o.Namespace,
			ProcessingUnitID:            &o.ProcessingUnitID,
			ProcessingUnitNamespace:     &o.ProcessingUnitNamespace,
			Protocol:                    &o.Protocol,
			Reason:                      &o.Reason,
			ServiceType:                 &o.ServiceType,
			SourceIP:                    &o.SourceIP,
			State:                       &o.State,
			Timestamp:                   &o.Timestamp,
			Value:                       &o.Value,
			ZHash:                       &o.ZHash,
			Zone:                        &o.Zone,
		}
	}

	sp := &SparseConnectionExceptionReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "destinationController":
			sp.DestinationController = &(o.DestinationController)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "destinationPort":
			sp.DestinationPort = &(o.DestinationPort)
		case "destinationProcessingUnitID":
			sp.DestinationProcessingUnitID = &(o.DestinationProcessingUnitID)
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
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "reason":
			sp.Reason = &(o.Reason)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "state":
			sp.State = &(o.State)
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

// Patch apply the non nil value of a *SparseConnectionExceptionReport to the object.
func (o *ConnectionExceptionReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseConnectionExceptionReport)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.DestinationController != nil {
		o.DestinationController = *so.DestinationController
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
	}
	if so.DestinationPort != nil {
		o.DestinationPort = *so.DestinationPort
	}
	if so.DestinationProcessingUnitID != nil {
		o.DestinationProcessingUnitID = *so.DestinationProcessingUnitID
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
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.Reason != nil {
		o.Reason = *so.Reason
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.State != nil {
		o.State = *so.State
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

// DeepCopy returns a deep copy if the ConnectionExceptionReport.
func (o *ConnectionExceptionReport) DeepCopy() *ConnectionExceptionReport {

	if o == nil {
		return nil
	}

	out := &ConnectionExceptionReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ConnectionExceptionReport.
func (o *ConnectionExceptionReport) DeepCopyInto(out *ConnectionExceptionReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ConnectionExceptionReport: %s", err))
	}

	*out = *target.(*ConnectionExceptionReport)
}

// Validate valides the current information stored into the structure.
func (o *ConnectionExceptionReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"L3", "HTTP", "TCP"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("state", string(o.State)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("state", string(o.State), []string{"SynTransmitted", "SynAckTransmitted", "AckTransmitted", "Unknown"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("value", o.Value); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	// Custom object validation.
	if err := ValidateConnectionExceptionReport(o); err != nil {
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
func (*ConnectionExceptionReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ConnectionExceptionReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ConnectionExceptionReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ConnectionExceptionReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ConnectionExceptionReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ConnectionExceptionReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "destinationController":
		return o.DestinationController
	case "destinationIP":
		return o.DestinationIP
	case "destinationPort":
		return o.DestinationPort
	case "destinationProcessingUnitID":
		return o.DestinationProcessingUnitID
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
	case "protocol":
		return o.Protocol
	case "reason":
		return o.Reason
	case "serviceType":
		return o.ServiceType
	case "sourceIP":
		return o.SourceIP
	case "state":
		return o.State
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

// ConnectionExceptionReportAttributesMap represents the map of attribute for ConnectionExceptionReport.
var ConnectionExceptionReportAttributesMap = map[string]elemental.AttributeSpecification{
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
	"DestinationController": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationController",
		Deprecated:     true,
		Description: `Identifier of the destination controller. This should be set in
SynAckTransmitted state.`,
		Exposed: true,
		Name:    "destinationController",
		Stored:  true,
		Type:    "string",
	},
	"DestinationIP": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Destination IP address.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"DestinationPort": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Stored:         true,
		Type:           "integer",
	},
	"DestinationProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationProcessingUnitID",
		Description: `ID of the destination processing unit. This should be set in SynAckTransmitted
state.`,
		Exposed: true,
		Name:    "destinationProcessingUnitID",
		Stored:  true,
		Type:    "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Deprecated:     true,
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Stored:         true,
		Type:           "string",
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
		ConvertedName:  "Namespace",
		Description:    `Namespace of the processing unit that encountered this exception.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit encountered this exception.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Deprecated:     true,
		Description:    `Namespace of the processing unit encountered this exception.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Reason": {
		AllowedChoices: []string{},
		ConvertedName:  "Reason",
		Description:    `It specifies the reason for the exception.`,
		Exposed:        true,
		Name:           "reason",
		Stored:         true,
		Type:           "string",
	},
	"ServiceType": {
		AllowedChoices: []string{"L3", "HTTP", "TCP"},
		ConvertedName:  "ServiceType",
		DefaultValue:   ConnectionExceptionReportServiceTypeL3,
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "enum",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Source IP address.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"State": {
		AllowedChoices: []string{"SynTransmitted", "SynAckTransmitted", "AckTransmitted", "Unknown"},
		ConvertedName:  "State",
		Description:    `Represents the current state this report was generated.`,
		Exposed:        true,
		Name:           "state",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "enum",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of packets hit.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "integer",
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

// ConnectionExceptionReportLowerCaseAttributesMap represents the map of attribute for ConnectionExceptionReport.
var ConnectionExceptionReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"destinationcontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationController",
		Deprecated:     true,
		Description: `Identifier of the destination controller. This should be set in
SynAckTransmitted state.`,
		Exposed: true,
		Name:    "destinationController",
		Stored:  true,
		Type:    "string",
	},
	"destinationip": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Destination IP address.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"destinationport": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Stored:         true,
		Type:           "integer",
	},
	"destinationprocessingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationProcessingUnitID",
		Description: `ID of the destination processing unit. This should be set in SynAckTransmitted
state.`,
		Exposed: true,
		Name:    "destinationProcessingUnitID",
		Stored:  true,
		Type:    "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Deprecated:     true,
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Stored:         true,
		Type:           "string",
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
		ConvertedName:  "Namespace",
		Description:    `Namespace of the processing unit that encountered this exception.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"processingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the processing unit encountered this exception.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Deprecated:     true,
		Description:    `Namespace of the processing unit encountered this exception.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "integer",
	},
	"reason": {
		AllowedChoices: []string{},
		ConvertedName:  "Reason",
		Description:    `It specifies the reason for the exception.`,
		Exposed:        true,
		Name:           "reason",
		Stored:         true,
		Type:           "string",
	},
	"servicetype": {
		AllowedChoices: []string{"L3", "HTTP", "TCP"},
		ConvertedName:  "ServiceType",
		DefaultValue:   ConnectionExceptionReportServiceTypeL3,
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "enum",
	},
	"sourceip": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Source IP address.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"state": {
		AllowedChoices: []string{"SynTransmitted", "SynAckTransmitted", "AckTransmitted", "Unknown"},
		ConvertedName:  "State",
		Description:    `Represents the current state this report was generated.`,
		Exposed:        true,
		Name:           "state",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "enum",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of packets hit.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "integer",
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

// SparseConnectionExceptionReportsList represents a list of SparseConnectionExceptionReports
type SparseConnectionExceptionReportsList []*SparseConnectionExceptionReport

// Identity returns the identity of the objects in the list.
func (o SparseConnectionExceptionReportsList) Identity() elemental.Identity {

	return ConnectionExceptionReportIdentity
}

// Copy returns a pointer to a copy the SparseConnectionExceptionReportsList.
func (o SparseConnectionExceptionReportsList) Copy() elemental.Identifiables {

	copy := append(SparseConnectionExceptionReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseConnectionExceptionReportsList.
func (o SparseConnectionExceptionReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseConnectionExceptionReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseConnectionExceptionReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseConnectionExceptionReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseConnectionExceptionReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseConnectionExceptionReportsList converted to ConnectionExceptionReportsList.
func (o SparseConnectionExceptionReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseConnectionExceptionReportsList) Version() int {

	return 1
}

// SparseConnectionExceptionReport represents the sparse version of a connectionexceptionreport.
type SparseConnectionExceptionReport struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Identifier of the destination controller. This should be set in
	// SynAckTransmitted state.
	DestinationController *string `json:"destinationController,omitempty" msgpack:"destinationController,omitempty" bson:"destinationcontroller,omitempty" mapstructure:"destinationController,omitempty"`

	// Destination IP address.
	DestinationIP *string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"destinationip,omitempty" mapstructure:"destinationIP,omitempty"`

	// Port of the destination.
	DestinationPort *int `json:"destinationPort,omitempty" msgpack:"destinationPort,omitempty" bson:"destinationport,omitempty" mapstructure:"destinationPort,omitempty"`

	// ID of the destination processing unit. This should be set in SynAckTransmitted
	// state.
	DestinationProcessingUnitID *string `json:"destinationProcessingUnitID,omitempty" msgpack:"destinationProcessingUnitID,omitempty" bson:"destinationprocessingunitid,omitempty" mapstructure:"destinationProcessingUnitID,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace of the processing unit that encountered this exception.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// ID of the processing unit encountered this exception.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"processingunitid,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit encountered this exception.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"processingunitnamespace,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Protocol number.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol,omitempty" mapstructure:"protocol,omitempty"`

	// It specifies the reason for the exception.
	Reason *string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"reason,omitempty" mapstructure:"reason,omitempty"`

	// Type of the service.
	ServiceType *ConnectionExceptionReportServiceTypeValue `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"servicetype,omitempty" mapstructure:"serviceType,omitempty"`

	// Source IP address.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"sourceip,omitempty" mapstructure:"sourceIP,omitempty"`

	// Represents the current state this report was generated.
	State *ConnectionExceptionReportStateValue `json:"state,omitempty" msgpack:"state,omitempty" bson:"state,omitempty" mapstructure:"state,omitempty"`

	// Time and date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// Number of packets hit.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"value,omitempty" mapstructure:"value,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseConnectionExceptionReport returns a new  SparseConnectionExceptionReport.
func NewSparseConnectionExceptionReport() *SparseConnectionExceptionReport {
	return &SparseConnectionExceptionReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseConnectionExceptionReport) Identity() elemental.Identity {

	return ConnectionExceptionReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseConnectionExceptionReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseConnectionExceptionReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionExceptionReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseConnectionExceptionReport{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.DestinationController != nil {
		s.DestinationController = o.DestinationController
	}
	if o.DestinationIP != nil {
		s.DestinationIP = o.DestinationIP
	}
	if o.DestinationPort != nil {
		s.DestinationPort = o.DestinationPort
	}
	if o.DestinationProcessingUnitID != nil {
		s.DestinationProcessingUnitID = o.DestinationProcessingUnitID
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
	if o.Protocol != nil {
		s.Protocol = o.Protocol
	}
	if o.Reason != nil {
		s.Reason = o.Reason
	}
	if o.ServiceType != nil {
		s.ServiceType = o.ServiceType
	}
	if o.SourceIP != nil {
		s.SourceIP = o.SourceIP
	}
	if o.State != nil {
		s.State = o.State
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
func (o *SparseConnectionExceptionReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseConnectionExceptionReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.DestinationController != nil {
		o.DestinationController = s.DestinationController
	}
	if s.DestinationIP != nil {
		o.DestinationIP = s.DestinationIP
	}
	if s.DestinationPort != nil {
		o.DestinationPort = s.DestinationPort
	}
	if s.DestinationProcessingUnitID != nil {
		o.DestinationProcessingUnitID = s.DestinationProcessingUnitID
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
	if s.Protocol != nil {
		o.Protocol = s.Protocol
	}
	if s.Reason != nil {
		o.Reason = s.Reason
	}
	if s.ServiceType != nil {
		o.ServiceType = s.ServiceType
	}
	if s.SourceIP != nil {
		o.SourceIP = s.SourceIP
	}
	if s.State != nil {
		o.State = s.State
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
func (o *SparseConnectionExceptionReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseConnectionExceptionReport) ToPlain() elemental.PlainIdentifiable {

	out := NewConnectionExceptionReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.DestinationController != nil {
		out.DestinationController = *o.DestinationController
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
	}
	if o.DestinationPort != nil {
		out.DestinationPort = *o.DestinationPort
	}
	if o.DestinationProcessingUnitID != nil {
		out.DestinationProcessingUnitID = *o.DestinationProcessingUnitID
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
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.Reason != nil {
		out.Reason = *o.Reason
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.State != nil {
		out.State = *o.State
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
func (o *SparseConnectionExceptionReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseConnectionExceptionReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseConnectionExceptionReport) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseConnectionExceptionReport) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseConnectionExceptionReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseConnectionExceptionReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseConnectionExceptionReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseConnectionExceptionReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseConnectionExceptionReport.
func (o *SparseConnectionExceptionReport) DeepCopy() *SparseConnectionExceptionReport {

	if o == nil {
		return nil
	}

	out := &SparseConnectionExceptionReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseConnectionExceptionReport.
func (o *SparseConnectionExceptionReport) DeepCopyInto(out *SparseConnectionExceptionReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseConnectionExceptionReport: %s", err))
	}

	*out = *target.(*SparseConnectionExceptionReport)
}

type mongoAttributesConnectionExceptionReport struct {
	ID                          bson.ObjectId                             `bson:"_id,omitempty"`
	DestinationController       string                                    `bson:"destinationcontroller,omitempty"`
	DestinationIP               string                                    `bson:"destinationip,omitempty"`
	DestinationPort             int                                       `bson:"destinationport,omitempty"`
	DestinationProcessingUnitID string                                    `bson:"destinationprocessingunitid,omitempty"`
	EnforcerID                  string                                    `bson:"enforcerid,omitempty"`
	EnforcerNamespace           string                                    `bson:"enforcernamespace,omitempty"`
	MigrationsLog               map[string]string                         `bson:"migrationslog,omitempty"`
	Namespace                   string                                    `bson:"namespace,omitempty"`
	ProcessingUnitID            string                                    `bson:"processingunitid,omitempty"`
	ProcessingUnitNamespace     string                                    `bson:"processingunitnamespace,omitempty"`
	Protocol                    int                                       `bson:"protocol,omitempty"`
	Reason                      string                                    `bson:"reason,omitempty"`
	ServiceType                 ConnectionExceptionReportServiceTypeValue `bson:"servicetype,omitempty"`
	SourceIP                    string                                    `bson:"sourceip,omitempty"`
	State                       ConnectionExceptionReportStateValue       `bson:"state"`
	Timestamp                   time.Time                                 `bson:"timestamp,omitempty"`
	Value                       int                                       `bson:"value,omitempty"`
	ZHash                       int                                       `bson:"zhash"`
	Zone                        int                                       `bson:"zone"`
}
type mongoAttributesSparseConnectionExceptionReport struct {
	ID                          bson.ObjectId                              `bson:"_id,omitempty"`
	DestinationController       *string                                    `bson:"destinationcontroller,omitempty"`
	DestinationIP               *string                                    `bson:"destinationip,omitempty"`
	DestinationPort             *int                                       `bson:"destinationport,omitempty"`
	DestinationProcessingUnitID *string                                    `bson:"destinationprocessingunitid,omitempty"`
	EnforcerID                  *string                                    `bson:"enforcerid,omitempty"`
	EnforcerNamespace           *string                                    `bson:"enforcernamespace,omitempty"`
	MigrationsLog               *map[string]string                         `bson:"migrationslog,omitempty"`
	Namespace                   *string                                    `bson:"namespace,omitempty"`
	ProcessingUnitID            *string                                    `bson:"processingunitid,omitempty"`
	ProcessingUnitNamespace     *string                                    `bson:"processingunitnamespace,omitempty"`
	Protocol                    *int                                       `bson:"protocol,omitempty"`
	Reason                      *string                                    `bson:"reason,omitempty"`
	ServiceType                 *ConnectionExceptionReportServiceTypeValue `bson:"servicetype,omitempty"`
	SourceIP                    *string                                    `bson:"sourceip,omitempty"`
	State                       *ConnectionExceptionReportStateValue       `bson:"state,omitempty"`
	Timestamp                   *time.Time                                 `bson:"timestamp,omitempty"`
	Value                       *int                                       `bson:"value,omitempty"`
	ZHash                       *int                                       `bson:"zhash,omitempty"`
	Zone                        *int                                       `bson:"zone,omitempty"`
}
