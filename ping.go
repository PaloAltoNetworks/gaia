package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingTypeValue represents the possible values for attribute "type".
type PingTypeValue string

const (
	// PingTypeRequest represents the value Request.
	PingTypeRequest PingTypeValue = "Request"

	// PingTypeResponse represents the value Response.
	PingTypeResponse PingTypeValue = "Response"
)

// PingIdentity represents the Identity of the object.
var PingIdentity = elemental.Identity{
	Name:     "ping",
	Category: "ping",
	Package:  "guy",
	Private:  false,
}

// PingsList represents a list of Pings
type PingsList []*Ping

// Identity returns the identity of the objects in the list.
func (o PingsList) Identity() elemental.Identity {

	return PingIdentity
}

// Copy returns a pointer to a copy the PingsList.
func (o PingsList) Copy() elemental.Identifiables {

	copy := append(PingsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PingsList.
func (o PingsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PingsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Ping))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PingsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PingsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PingsList converted to SparsePingsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PingsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePingsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePing)
	}

	return out
}

// Version returns the version of the content.
func (o PingsList) Version() int {

	return 1
}

// Ping represents the model of a ping
type Ping struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Time taken for a single request-response to complete.
	RTT string `json:"RTT" msgpack:"RTT" bson:"rtt" mapstructure:"RTT,omitempty"`

	// Type of the transmitter.
	TXType string `json:"TXType" msgpack:"TXType" bson:"txtype" mapstructure:"TXType,omitempty"`

	// If true, application responded to the request.
	ApplicationListening bool `json:"applicationListening" msgpack:"applicationListening" bson:"applicationlistening" mapstructure:"applicationListening,omitempty"`

	// Claims carried on the wire.
	Claims []string `json:"claims" msgpack:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// ID of the destination processing unit.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion string `json:"enforcerVersion" msgpack:"enforcerVersion" bson:"enforcerversion" mapstructure:"enforcerVersion,omitempty"`

	// A non-empty error indicates a failure.
	Error string `json:"error" msgpack:"error" bson:"error" mapstructure:"error,omitempty"`

	// If true, destination IP is in excludedNetworks.
	ExcludedNetworks bool `json:"excludedNetworks" msgpack:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple string `json:"fourTuple" msgpack:"fourTuple" bson:"fourtuple" mapstructure:"fourTuple,omitempty"`

	// IterationID unique to a single ping request-response.
	IterationID string `json:"iterationID" msgpack:"iterationID" bson:"iterationid" mapstructure:"iterationID,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize int `json:"payloadSize" msgpack:"payloadSize" bson:"payloadsize" mapstructure:"payloadSize,omitempty"`

	// PingID unique to a single ping control.
	PingID string `json:"pingID" msgpack:"pingID" bson:"pingid" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction string `json:"policyAction" msgpack:"policyAction" bson:"policyaction" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// Namespace of the reporting processing unit.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"processingunitnamespace" mapstructure:"processingUnitNamespace,omitempty"`

	// Protocol used for the communication.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Controller of the remote endpoint.
	RemoteController string `json:"remoteController" msgpack:"remoteController" bson:"remotecontroller" mapstructure:"remoteController,omitempty"`

	// Namespace hash of the remote endpoint.
	RemoteNamespaceHash string `json:"remoteNamespaceHash" msgpack:"remoteNamespaceHash" bson:"remotenamespacehash" mapstructure:"remoteNamespaceHash,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum int `json:"seqNum" msgpack:"seqNum" bson:"seqnum" mapstructure:"seqNum,omitempty"`

	// Type of the service.
	ServiceType string `json:"serviceType" msgpack:"serviceType" bson:"servicetype" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// If true, destination IP is in targetTCPNetworks.
	TargetTCPNetworks bool `json:"targetTCPNetworks" msgpack:"targetTCPNetworks" bson:"targettcpnetworks" mapstructure:"targetTCPNetworks,omitempty"`

	// Type of the report.
	Type PingTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPing returns a new *Ping
func NewPing() *Ping {

	return &Ping{
		ModelVersion:  1,
		Claims:        []string{},
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Ping) Identity() elemental.Identity {

	return PingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Ping) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Ping) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Ping) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPing{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.RTT = o.RTT
	s.TXType = o.TXType
	s.ApplicationListening = o.ApplicationListening
	s.Claims = o.Claims
	s.CreateTime = o.CreateTime
	s.DestinationID = o.DestinationID
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.EnforcerVersion = o.EnforcerVersion
	s.Error = o.Error
	s.ExcludedNetworks = o.ExcludedNetworks
	s.FourTuple = o.FourTuple
	s.IterationID = o.IterationID
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.PayloadSize = o.PayloadSize
	s.PingID = o.PingID
	s.PolicyAction = o.PolicyAction
	s.PolicyID = o.PolicyID
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.Protocol = o.Protocol
	s.RemoteController = o.RemoteController
	s.RemoteNamespaceHash = o.RemoteNamespaceHash
	s.SeqNum = o.SeqNum
	s.ServiceType = o.ServiceType
	s.SourceID = o.SourceID
	s.TargetTCPNetworks = o.TargetTCPNetworks
	s.Type = o.Type
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Ping) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPing{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.RTT = s.RTT
	o.TXType = s.TXType
	o.ApplicationListening = s.ApplicationListening
	o.Claims = s.Claims
	o.CreateTime = s.CreateTime
	o.DestinationID = s.DestinationID
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.EnforcerVersion = s.EnforcerVersion
	o.Error = s.Error
	o.ExcludedNetworks = s.ExcludedNetworks
	o.FourTuple = s.FourTuple
	o.IterationID = s.IterationID
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.PayloadSize = s.PayloadSize
	o.PingID = s.PingID
	o.PolicyAction = s.PolicyAction
	o.PolicyID = s.PolicyID
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.Protocol = s.Protocol
	o.RemoteController = s.RemoteController
	o.RemoteNamespaceHash = s.RemoteNamespaceHash
	o.SeqNum = s.SeqNum
	o.ServiceType = s.ServiceType
	o.SourceID = s.SourceID
	o.TargetTCPNetworks = s.TargetTCPNetworks
	o.Type = s.Type
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Ping) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Ping) BleveType() string {

	return "ping"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Ping) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Ping) Doc() string {

	return `Post a new pu ping.`
}

func (o *Ping) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Ping) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Ping) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Ping) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Ping) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *Ping) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Ping) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Ping) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Ping) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Ping) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Ping) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Ping) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Ping) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Ping) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePing{
			ID:                      &o.ID,
			RTT:                     &o.RTT,
			TXType:                  &o.TXType,
			ApplicationListening:    &o.ApplicationListening,
			Claims:                  &o.Claims,
			CreateTime:              &o.CreateTime,
			DestinationID:           &o.DestinationID,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			EnforcerVersion:         &o.EnforcerVersion,
			Error:                   &o.Error,
			ExcludedNetworks:        &o.ExcludedNetworks,
			FourTuple:               &o.FourTuple,
			IterationID:             &o.IterationID,
			MigrationsLog:           &o.MigrationsLog,
			Namespace:               &o.Namespace,
			PayloadSize:             &o.PayloadSize,
			PingID:                  &o.PingID,
			PolicyAction:            &o.PolicyAction,
			PolicyID:                &o.PolicyID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Protocol:                &o.Protocol,
			RemoteController:        &o.RemoteController,
			RemoteNamespaceHash:     &o.RemoteNamespaceHash,
			SeqNum:                  &o.SeqNum,
			ServiceType:             &o.ServiceType,
			SourceID:                &o.SourceID,
			TargetTCPNetworks:       &o.TargetTCPNetworks,
			Type:                    &o.Type,
			UpdateTime:              &o.UpdateTime,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparsePing{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "RTT":
			sp.RTT = &(o.RTT)
		case "TXType":
			sp.TXType = &(o.TXType)
		case "applicationListening":
			sp.ApplicationListening = &(o.ApplicationListening)
		case "claims":
			sp.Claims = &(o.Claims)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "enforcerVersion":
			sp.EnforcerVersion = &(o.EnforcerVersion)
		case "error":
			sp.Error = &(o.Error)
		case "excludedNetworks":
			sp.ExcludedNetworks = &(o.ExcludedNetworks)
		case "fourTuple":
			sp.FourTuple = &(o.FourTuple)
		case "iterationID":
			sp.IterationID = &(o.IterationID)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "payloadSize":
			sp.PayloadSize = &(o.PayloadSize)
		case "pingID":
			sp.PingID = &(o.PingID)
		case "policyAction":
			sp.PolicyAction = &(o.PolicyAction)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "remoteController":
			sp.RemoteController = &(o.RemoteController)
		case "remoteNamespaceHash":
			sp.RemoteNamespaceHash = &(o.RemoteNamespaceHash)
		case "seqNum":
			sp.SeqNum = &(o.SeqNum)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "targetTCPNetworks":
			sp.TargetTCPNetworks = &(o.TargetTCPNetworks)
		case "type":
			sp.Type = &(o.Type)
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

// Patch apply the non nil value of a *SparsePing to the object.
func (o *Ping) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePing)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.RTT != nil {
		o.RTT = *so.RTT
	}
	if so.TXType != nil {
		o.TXType = *so.TXType
	}
	if so.ApplicationListening != nil {
		o.ApplicationListening = *so.ApplicationListening
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.EnforcerVersion != nil {
		o.EnforcerVersion = *so.EnforcerVersion
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.ExcludedNetworks != nil {
		o.ExcludedNetworks = *so.ExcludedNetworks
	}
	if so.FourTuple != nil {
		o.FourTuple = *so.FourTuple
	}
	if so.IterationID != nil {
		o.IterationID = *so.IterationID
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PayloadSize != nil {
		o.PayloadSize = *so.PayloadSize
	}
	if so.PingID != nil {
		o.PingID = *so.PingID
	}
	if so.PolicyAction != nil {
		o.PolicyAction = *so.PolicyAction
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.RemoteController != nil {
		o.RemoteController = *so.RemoteController
	}
	if so.RemoteNamespaceHash != nil {
		o.RemoteNamespaceHash = *so.RemoteNamespaceHash
	}
	if so.SeqNum != nil {
		o.SeqNum = *so.SeqNum
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.TargetTCPNetworks != nil {
		o.TargetTCPNetworks = *so.TargetTCPNetworks
	}
	if so.Type != nil {
		o.Type = *so.Type
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

// DeepCopy returns a deep copy if the Ping.
func (o *Ping) DeepCopy() *Ping {

	if o == nil {
		return nil
	}

	out := &Ping{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Ping.
func (o *Ping) DeepCopyInto(out *Ping) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Ping: %s", err))
	}

	*out = *target.(*Ping)
}

// Validate valides the current information stored into the structure.
func (o *Ping) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("pingID", o.PingID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Request", "Response"}, false); err != nil {
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
func (*Ping) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Ping) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Ping) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "RTT":
		return o.RTT
	case "TXType":
		return o.TXType
	case "applicationListening":
		return o.ApplicationListening
	case "claims":
		return o.Claims
	case "createTime":
		return o.CreateTime
	case "destinationID":
		return o.DestinationID
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "enforcerVersion":
		return o.EnforcerVersion
	case "error":
		return o.Error
	case "excludedNetworks":
		return o.ExcludedNetworks
	case "fourTuple":
		return o.FourTuple
	case "iterationID":
		return o.IterationID
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "payloadSize":
		return o.PayloadSize
	case "pingID":
		return o.PingID
	case "policyAction":
		return o.PolicyAction
	case "policyID":
		return o.PolicyID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "protocol":
		return o.Protocol
	case "remoteController":
		return o.RemoteController
	case "remoteNamespaceHash":
		return o.RemoteNamespaceHash
	case "seqNum":
		return o.SeqNum
	case "serviceType":
		return o.ServiceType
	case "sourceID":
		return o.SourceID
	case "targetTCPNetworks":
		return o.TargetTCPNetworks
	case "type":
		return o.Type
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// PingAttributesMap represents the map of attribute for Ping.
var PingAttributesMap = map[string]elemental.AttributeSpecification{
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
	"RTT": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Stored:         true,
		Type:           "string",
	},
	"TXType": {
		AllowedChoices: []string{},
		ConvertedName:  "TXType",
		Description:    `Type of the transmitter.`,
		Exposed:        true,
		Name:           "TXType",
		Stored:         true,
		Type:           "string",
	},
	"ApplicationListening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Stored:         true,
		Type:           "boolean",
	},
	"Claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims carried on the wire.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"DestinationID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination processing unit.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
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
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerVersion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Stored:         true,
		Type:           "string",
	},
	"Error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `A non-empty error indicates a failure.`,
		Exposed:        true,
		Name:           "error",
		Stored:         true,
		Type:           "string",
	},
	"ExcludedNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description:    `If true, destination IP is in excludedNetworks.`,
		Exposed:        true,
		Name:           "excludedNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"FourTuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Stored:         true,
		Type:           "string",
	},
	"IterationID": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationID",
		Description:    `IterationID unique to a single ping request-response.`,
		Exposed:        true,
		Name:           "iterationID",
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
	"PayloadSize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Stored:         true,
		Type:           "integer",
	},
	"PingID": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"PolicyAction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Stored:         true,
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the reporting processing unit.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"RemoteController": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteController",
		Description:    `Controller of the remote endpoint.`,
		Exposed:        true,
		Name:           "remoteController",
		Stored:         true,
		Type:           "string",
	},
	"RemoteNamespaceHash": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespaceHash",
		Description:    `Namespace hash of the remote endpoint.`,
		Exposed:        true,
		Name:           "remoteNamespaceHash",
		Stored:         true,
		Type:           "string",
	},
	"SeqNum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
		Stored:         true,
		Type:           "integer",
	},
	"ServiceType": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "string",
	},
	"SourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source PU.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"TargetTCPNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetTCPNetworks",
		Description:    `If true, destination IP is in targetTCPNetworks.`,
		Exposed:        true,
		Name:           "targetTCPNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"Type": {
		AllowedChoices: []string{"Request", "Response"},
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// PingLowerCaseAttributesMap represents the map of attribute for Ping.
var PingLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"rtt": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Stored:         true,
		Type:           "string",
	},
	"txtype": {
		AllowedChoices: []string{},
		ConvertedName:  "TXType",
		Description:    `Type of the transmitter.`,
		Exposed:        true,
		Name:           "TXType",
		Stored:         true,
		Type:           "string",
	},
	"applicationlistening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Stored:         true,
		Type:           "boolean",
	},
	"claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims carried on the wire.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"destinationid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination processing unit.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
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
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcerversion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Stored:         true,
		Type:           "string",
	},
	"error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `A non-empty error indicates a failure.`,
		Exposed:        true,
		Name:           "error",
		Stored:         true,
		Type:           "string",
	},
	"excludednetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description:    `If true, destination IP is in excludedNetworks.`,
		Exposed:        true,
		Name:           "excludedNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"fourtuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Stored:         true,
		Type:           "string",
	},
	"iterationid": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationID",
		Description:    `IterationID unique to a single ping request-response.`,
		Exposed:        true,
		Name:           "iterationID",
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
	"payloadsize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Stored:         true,
		Type:           "integer",
	},
	"pingid": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"policyaction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Stored:         true,
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the reporting processing unit.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"remotecontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteController",
		Description:    `Controller of the remote endpoint.`,
		Exposed:        true,
		Name:           "remoteController",
		Stored:         true,
		Type:           "string",
	},
	"remotenamespacehash": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespaceHash",
		Description:    `Namespace hash of the remote endpoint.`,
		Exposed:        true,
		Name:           "remoteNamespaceHash",
		Stored:         true,
		Type:           "string",
	},
	"seqnum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
		Stored:         true,
		Type:           "integer",
	},
	"servicetype": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "string",
	},
	"sourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source PU.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"targettcpnetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetTCPNetworks",
		Description:    `If true, destination IP is in targetTCPNetworks.`,
		Exposed:        true,
		Name:           "targetTCPNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"type": {
		AllowedChoices: []string{"Request", "Response"},
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparsePingsList represents a list of SparsePings
type SparsePingsList []*SparsePing

// Identity returns the identity of the objects in the list.
func (o SparsePingsList) Identity() elemental.Identity {

	return PingIdentity
}

// Copy returns a pointer to a copy the SparsePingsList.
func (o SparsePingsList) Copy() elemental.Identifiables {

	copy := append(SparsePingsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePingsList.
func (o SparsePingsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePingsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePing))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePingsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePingsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePingsList converted to PingsList.
func (o SparsePingsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePingsList) Version() int {

	return 1
}

// SparsePing represents the sparse version of a ping.
type SparsePing struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Time taken for a single request-response to complete.
	RTT *string `json:"RTT,omitempty" msgpack:"RTT,omitempty" bson:"rtt,omitempty" mapstructure:"RTT,omitempty"`

	// Type of the transmitter.
	TXType *string `json:"TXType,omitempty" msgpack:"TXType,omitempty" bson:"txtype,omitempty" mapstructure:"TXType,omitempty"`

	// If true, application responded to the request.
	ApplicationListening *bool `json:"applicationListening,omitempty" msgpack:"applicationListening,omitempty" bson:"applicationlistening,omitempty" mapstructure:"applicationListening,omitempty"`

	// Claims carried on the wire.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"claims,omitempty" mapstructure:"claims,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// ID of the destination processing unit.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"destinationid,omitempty" mapstructure:"destinationID,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion *string `json:"enforcerVersion,omitempty" msgpack:"enforcerVersion,omitempty" bson:"enforcerversion,omitempty" mapstructure:"enforcerVersion,omitempty"`

	// A non-empty error indicates a failure.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"error,omitempty" mapstructure:"error,omitempty"`

	// If true, destination IP is in excludedNetworks.
	ExcludedNetworks *bool `json:"excludedNetworks,omitempty" msgpack:"excludedNetworks,omitempty" bson:"excludednetworks,omitempty" mapstructure:"excludedNetworks,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple *string `json:"fourTuple,omitempty" msgpack:"fourTuple,omitempty" bson:"fourtuple,omitempty" mapstructure:"fourTuple,omitempty"`

	// IterationID unique to a single ping request-response.
	IterationID *string `json:"iterationID,omitempty" msgpack:"iterationID,omitempty" bson:"iterationid,omitempty" mapstructure:"iterationID,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize *int `json:"payloadSize,omitempty" msgpack:"payloadSize,omitempty" bson:"payloadsize,omitempty" mapstructure:"payloadSize,omitempty"`

	// PingID unique to a single ping control.
	PingID *string `json:"pingID,omitempty" msgpack:"pingID,omitempty" bson:"pingid,omitempty" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction *string `json:"policyAction,omitempty" msgpack:"policyAction,omitempty" bson:"policyaction,omitempty" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// Namespace of the reporting processing unit.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"processingunitnamespace,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Protocol used for the communication.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol,omitempty" mapstructure:"protocol,omitempty"`

	// Controller of the remote endpoint.
	RemoteController *string `json:"remoteController,omitempty" msgpack:"remoteController,omitempty" bson:"remotecontroller,omitempty" mapstructure:"remoteController,omitempty"`

	// Namespace hash of the remote endpoint.
	RemoteNamespaceHash *string `json:"remoteNamespaceHash,omitempty" msgpack:"remoteNamespaceHash,omitempty" bson:"remotenamespacehash,omitempty" mapstructure:"remoteNamespaceHash,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum *int `json:"seqNum,omitempty" msgpack:"seqNum,omitempty" bson:"seqnum,omitempty" mapstructure:"seqNum,omitempty"`

	// Type of the service.
	ServiceType *string `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"servicetype,omitempty" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// If true, destination IP is in targetTCPNetworks.
	TargetTCPNetworks *bool `json:"targetTCPNetworks,omitempty" msgpack:"targetTCPNetworks,omitempty" bson:"targettcpnetworks,omitempty" mapstructure:"targetTCPNetworks,omitempty"`

	// Type of the report.
	Type *PingTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePing returns a new  SparsePing.
func NewSparsePing() *SparsePing {
	return &SparsePing{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePing) Identity() elemental.Identity {

	return PingIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePing) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePing) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePing) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePing{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.RTT != nil {
		s.RTT = o.RTT
	}
	if o.TXType != nil {
		s.TXType = o.TXType
	}
	if o.ApplicationListening != nil {
		s.ApplicationListening = o.ApplicationListening
	}
	if o.Claims != nil {
		s.Claims = o.Claims
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.DestinationID != nil {
		s.DestinationID = o.DestinationID
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.EnforcerVersion != nil {
		s.EnforcerVersion = o.EnforcerVersion
	}
	if o.Error != nil {
		s.Error = o.Error
	}
	if o.ExcludedNetworks != nil {
		s.ExcludedNetworks = o.ExcludedNetworks
	}
	if o.FourTuple != nil {
		s.FourTuple = o.FourTuple
	}
	if o.IterationID != nil {
		s.IterationID = o.IterationID
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PayloadSize != nil {
		s.PayloadSize = o.PayloadSize
	}
	if o.PingID != nil {
		s.PingID = o.PingID
	}
	if o.PolicyAction != nil {
		s.PolicyAction = o.PolicyAction
	}
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
	}
	if o.ProcessingUnitNamespace != nil {
		s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	}
	if o.Protocol != nil {
		s.Protocol = o.Protocol
	}
	if o.RemoteController != nil {
		s.RemoteController = o.RemoteController
	}
	if o.RemoteNamespaceHash != nil {
		s.RemoteNamespaceHash = o.RemoteNamespaceHash
	}
	if o.SeqNum != nil {
		s.SeqNum = o.SeqNum
	}
	if o.ServiceType != nil {
		s.ServiceType = o.ServiceType
	}
	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.TargetTCPNetworks != nil {
		s.TargetTCPNetworks = o.TargetTCPNetworks
	}
	if o.Type != nil {
		s.Type = o.Type
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
func (o *SparsePing) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePing{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.RTT != nil {
		o.RTT = s.RTT
	}
	if s.TXType != nil {
		o.TXType = s.TXType
	}
	if s.ApplicationListening != nil {
		o.ApplicationListening = s.ApplicationListening
	}
	if s.Claims != nil {
		o.Claims = s.Claims
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.DestinationID != nil {
		o.DestinationID = s.DestinationID
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.EnforcerVersion != nil {
		o.EnforcerVersion = s.EnforcerVersion
	}
	if s.Error != nil {
		o.Error = s.Error
	}
	if s.ExcludedNetworks != nil {
		o.ExcludedNetworks = s.ExcludedNetworks
	}
	if s.FourTuple != nil {
		o.FourTuple = s.FourTuple
	}
	if s.IterationID != nil {
		o.IterationID = s.IterationID
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PayloadSize != nil {
		o.PayloadSize = s.PayloadSize
	}
	if s.PingID != nil {
		o.PingID = s.PingID
	}
	if s.PolicyAction != nil {
		o.PolicyAction = s.PolicyAction
	}
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
	}
	if s.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	}
	if s.Protocol != nil {
		o.Protocol = s.Protocol
	}
	if s.RemoteController != nil {
		o.RemoteController = s.RemoteController
	}
	if s.RemoteNamespaceHash != nil {
		o.RemoteNamespaceHash = s.RemoteNamespaceHash
	}
	if s.SeqNum != nil {
		o.SeqNum = s.SeqNum
	}
	if s.ServiceType != nil {
		o.ServiceType = s.ServiceType
	}
	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.TargetTCPNetworks != nil {
		o.TargetTCPNetworks = s.TargetTCPNetworks
	}
	if s.Type != nil {
		o.Type = s.Type
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
func (o *SparsePing) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePing) ToPlain() elemental.PlainIdentifiable {

	out := NewPing()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.RTT != nil {
		out.RTT = *o.RTT
	}
	if o.TXType != nil {
		out.TXType = *o.TXType
	}
	if o.ApplicationListening != nil {
		out.ApplicationListening = *o.ApplicationListening
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.EnforcerVersion != nil {
		out.EnforcerVersion = *o.EnforcerVersion
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.ExcludedNetworks != nil {
		out.ExcludedNetworks = *o.ExcludedNetworks
	}
	if o.FourTuple != nil {
		out.FourTuple = *o.FourTuple
	}
	if o.IterationID != nil {
		out.IterationID = *o.IterationID
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PayloadSize != nil {
		out.PayloadSize = *o.PayloadSize
	}
	if o.PingID != nil {
		out.PingID = *o.PingID
	}
	if o.PolicyAction != nil {
		out.PolicyAction = *o.PolicyAction
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.RemoteController != nil {
		out.RemoteController = *o.RemoteController
	}
	if o.RemoteNamespaceHash != nil {
		out.RemoteNamespaceHash = *o.RemoteNamespaceHash
	}
	if o.SeqNum != nil {
		out.SeqNum = *o.SeqNum
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.TargetTCPNetworks != nil {
		out.TargetTCPNetworks = *o.TargetTCPNetworks
	}
	if o.Type != nil {
		out.Type = *o.Type
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
func (o *SparsePing) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparsePing) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparsePing) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparsePing) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePing) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePing) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparsePing) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparsePing) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePing) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePing) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePing) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePing) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparsePing.
func (o *SparsePing) DeepCopy() *SparsePing {

	if o == nil {
		return nil
	}

	out := &SparsePing{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePing.
func (o *SparsePing) DeepCopyInto(out *SparsePing) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePing: %s", err))
	}

	*out = *target.(*SparsePing)
}

type mongoAttributesPing struct {
	ID                      bson.ObjectId     `bson:"_id,omitempty"`
	RTT                     string            `bson:"rtt"`
	TXType                  string            `bson:"txtype"`
	ApplicationListening    bool              `bson:"applicationlistening"`
	Claims                  []string          `bson:"claims"`
	CreateTime              time.Time         `bson:"createtime"`
	DestinationID           string            `bson:"destinationid"`
	EnforcerID              string            `bson:"enforcerid"`
	EnforcerNamespace       string            `bson:"enforcernamespace"`
	EnforcerVersion         string            `bson:"enforcerversion"`
	Error                   string            `bson:"error"`
	ExcludedNetworks        bool              `bson:"excludednetworks"`
	FourTuple               string            `bson:"fourtuple"`
	IterationID             string            `bson:"iterationid"`
	MigrationsLog           map[string]string `bson:"migrationslog,omitempty"`
	Namespace               string            `bson:"namespace"`
	PayloadSize             int               `bson:"payloadsize"`
	PingID                  string            `bson:"pingid"`
	PolicyAction            string            `bson:"policyaction"`
	PolicyID                string            `bson:"policyid"`
	ProcessingUnitNamespace string            `bson:"processingunitnamespace"`
	Protocol                int               `bson:"protocol"`
	RemoteController        string            `bson:"remotecontroller"`
	RemoteNamespaceHash     string            `bson:"remotenamespacehash"`
	SeqNum                  int               `bson:"seqnum"`
	ServiceType             string            `bson:"servicetype"`
	SourceID                string            `bson:"sourceid"`
	TargetTCPNetworks       bool              `bson:"targettcpnetworks"`
	Type                    PingTypeValue     `bson:"type"`
	UpdateTime              time.Time         `bson:"updatetime"`
	ZHash                   int               `bson:"zhash"`
	Zone                    int               `bson:"zone"`
}
type mongoAttributesSparsePing struct {
	ID                      bson.ObjectId      `bson:"_id,omitempty"`
	RTT                     *string            `bson:"rtt,omitempty"`
	TXType                  *string            `bson:"txtype,omitempty"`
	ApplicationListening    *bool              `bson:"applicationlistening,omitempty"`
	Claims                  *[]string          `bson:"claims,omitempty"`
	CreateTime              *time.Time         `bson:"createtime,omitempty"`
	DestinationID           *string            `bson:"destinationid,omitempty"`
	EnforcerID              *string            `bson:"enforcerid,omitempty"`
	EnforcerNamespace       *string            `bson:"enforcernamespace,omitempty"`
	EnforcerVersion         *string            `bson:"enforcerversion,omitempty"`
	Error                   *string            `bson:"error,omitempty"`
	ExcludedNetworks        *bool              `bson:"excludednetworks,omitempty"`
	FourTuple               *string            `bson:"fourtuple,omitempty"`
	IterationID             *string            `bson:"iterationid,omitempty"`
	MigrationsLog           *map[string]string `bson:"migrationslog,omitempty"`
	Namespace               *string            `bson:"namespace,omitempty"`
	PayloadSize             *int               `bson:"payloadsize,omitempty"`
	PingID                  *string            `bson:"pingid,omitempty"`
	PolicyAction            *string            `bson:"policyaction,omitempty"`
	PolicyID                *string            `bson:"policyid,omitempty"`
	ProcessingUnitNamespace *string            `bson:"processingunitnamespace,omitempty"`
	Protocol                *int               `bson:"protocol,omitempty"`
	RemoteController        *string            `bson:"remotecontroller,omitempty"`
	RemoteNamespaceHash     *string            `bson:"remotenamespacehash,omitempty"`
	SeqNum                  *int               `bson:"seqnum,omitempty"`
	ServiceType             *string            `bson:"servicetype,omitempty"`
	SourceID                *string            `bson:"sourceid,omitempty"`
	TargetTCPNetworks       *bool              `bson:"targettcpnetworks,omitempty"`
	Type                    *PingTypeValue     `bson:"type,omitempty"`
	UpdateTime              *time.Time         `bson:"updatetime,omitempty"`
	ZHash                   *int               `bson:"zhash,omitempty"`
	Zone                    *int               `bson:"zone,omitempty"`
}
