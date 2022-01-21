package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CounterReportIdentity represents the Identity of the object.
var CounterReportIdentity = elemental.Identity{
	Name:     "counterreport",
	Category: "counterreports",
	Package:  "zack",
	Private:  false,
}

// CounterReportsList represents a list of CounterReports
type CounterReportsList []*CounterReport

// Identity returns the identity of the objects in the list.
func (o CounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the CounterReportsList.
func (o CounterReportsList) Copy() elemental.Identifiables {

	copy := append(CounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CounterReportsList.
func (o CounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CounterReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToSparse returns the CounterReportsList converted to SparseCounterReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CounterReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCounterReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCounterReport)
	}

	return out
}

// Version returns the version of the content.
func (o CounterReportsList) Version() int {

	return 1
}

// CounterReport represents the model of a counterreport
type CounterReport struct {
	// Counter for sending FIN ACK received in unknown connection state.
	AckInUnknownState int `json:"AckInUnknownState,omitempty" msgpack:"AckInUnknownState,omitempty" bson:"ackinunknownstate" mapstructure:"AckInUnknownState,omitempty"`

	// Counter for ACK packet dropped because of invalid format.
	AckInvalidFormat int `json:"AckInvalidFormat,omitempty" msgpack:"AckInvalidFormat,omitempty" bson:"ackinvalidformat" mapstructure:"AckInvalidFormat,omitempty"`

	// Counter for ACK packets rejected as per policy.
	AckRejected int `json:"AckRejected,omitempty" msgpack:"AckRejected,omitempty" bson:"ackrejected" mapstructure:"AckRejected,omitempty"`

	// Counter for ACK packet dropped because signature validation failed.
	AckSigValidationFailed int `json:"AckSigValidationFailed,omitempty" msgpack:"AckSigValidationFailed,omitempty" bson:"acksigvalidationfailed" mapstructure:"AckSigValidationFailed,omitempty"`

	// Counter for TCP authentication option not found.
	AckTCPNoTCPAuthOption int `json:"AckTCPNoTCPAuthOption,omitempty" msgpack:"AckTCPNoTCPAuthOption,omitempty" bson:"acktcpnotcpauthoption" mapstructure:"AckTCPNoTCPAuthOption,omitempty"`

	// Counter for connections processed.
	ConnectionsProcessed int `json:"ConnectionsProcessed,omitempty" msgpack:"ConnectionsProcessed,omitempty" bson:"connectionsprocessed" mapstructure:"ConnectionsProcessed,omitempty"`

	// Counter for unable to find ContextID.
	ContextIDNotFound int `json:"ContextIDNotFound,omitempty" msgpack:"ContextIDNotFound,omitempty" bson:"contextidnotfound" mapstructure:"ContextIDNotFound,omitempty"`

	// Counter for no ACLs found for external services. Dropping application SYN
	// packet.
	DroppedExternalService int `json:"DroppedExternalService,omitempty" msgpack:"DroppedExternalService,omitempty" bson:"droppedexternalservice" mapstructure:"DroppedExternalService,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Counter for invalid connection state.
	InvalidConnState int `json:"InvalidConnState,omitempty" msgpack:"InvalidConnState,omitempty" bson:"invalidconnstate" mapstructure:"InvalidConnState,omitempty"`

	// Counter for invalid net state.
	InvalidNetState int `json:"InvalidNetState,omitempty" msgpack:"InvalidNetState,omitempty" bson:"invalidnetstate" mapstructure:"InvalidNetState,omitempty"`

	// Counter for invalid protocol.
	InvalidProtocol int `json:"InvalidProtocol,omitempty" msgpack:"InvalidProtocol,omitempty" bson:"invalidprotocol" mapstructure:"InvalidProtocol,omitempty"`

	// Counter for processing unit is already dead - drop SYN ACK packet.
	InvalidSynAck int `json:"InvalidSynAck,omitempty" msgpack:"InvalidSynAck,omitempty" bson:"invalidsynack" mapstructure:"InvalidSynAck,omitempty"`

	// Counter for processing unit mark not found.
	MarkNotFound int `json:"MarkNotFound,omitempty" msgpack:"MarkNotFound,omitempty" bson:"marknotfound" mapstructure:"MarkNotFound,omitempty"`

	// Counter for network SYN packet was not seen.
	NetSynNotSeen int `json:"NetSynNotSeen,omitempty" msgpack:"NetSynNotSeen,omitempty" bson:"netsynnotseen" mapstructure:"NetSynNotSeen,omitempty"`

	// Counter for no context or connection found.
	NoConnFound int `json:"NoConnFound,omitempty" msgpack:"NoConnFound,omitempty" bson:"noconnfound" mapstructure:"NoConnFound,omitempty"`

	// Counter for traffic that belongs to a non-processing unit process.
	NonPUTraffic int `json:"NonPUTraffic,omitempty" msgpack:"NonPUTraffic,omitempty" bson:"nonputraffic" mapstructure:"NonPUTraffic,omitempty"`

	// Counter for SYN ACK for flow with processed FIN ACK.
	OutOfOrderSynAck int `json:"OutOfOrderSynAck,omitempty" msgpack:"OutOfOrderSynAck,omitempty" bson:"outofordersynack" mapstructure:"OutOfOrderSynAck,omitempty"`

	// Counter for port not found.
	PortNotFound int `json:"PortNotFound,omitempty" msgpack:"PortNotFound,omitempty" bson:"portnotfound" mapstructure:"PortNotFound,omitempty"`

	// Counter for reject the packet as per policy.
	RejectPacket int `json:"RejectPacket,omitempty" msgpack:"RejectPacket,omitempty" bson:"rejectpacket" mapstructure:"RejectPacket,omitempty"`

	// Counter for post service processing failed for network packet.
	ServicePostprocessorFailed int `json:"ServicePostprocessorFailed,omitempty" msgpack:"ServicePostprocessorFailed,omitempty" bson:"servicepostprocessorfailed" mapstructure:"ServicePostprocessorFailed,omitempty"`

	// Counter for network packets that failed preprocessing.
	ServicePreprocessorFailed int `json:"ServicePreprocessorFailed,omitempty" msgpack:"ServicePreprocessorFailed,omitempty" bson:"servicepreprocessorfailed" mapstructure:"ServicePreprocessorFailed,omitempty"`

	// Counter for SYN ACK packet dropped because of bad claims.
	SynAckBadClaims int `json:"SynAckBadClaims,omitempty" msgpack:"SynAckBadClaims,omitempty" bson:"synackbadclaims" mapstructure:"SynAckBadClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of encryption mismatch.
	SynAckClaimsMisMatch int `json:"SynAckClaimsMisMatch,omitempty" msgpack:"SynAckClaimsMisMatch,omitempty" bson:"synackclaimsmismatch" mapstructure:"SynAckClaimsMisMatch,omitempty"`

	// Counter for SYN ACK from external service dropped.
	SynAckDroppedExternalService int `json:"SynAckDroppedExternalService,omitempty" msgpack:"SynAckDroppedExternalService,omitempty" bson:"synackdroppedexternalservice" mapstructure:"SynAckDroppedExternalService,omitempty"`

	// Counter for SYN ACK packet dropped because of invalid format.
	SynAckInvalidFormat int `json:"SynAckInvalidFormat,omitempty" msgpack:"SynAckInvalidFormat,omitempty" bson:"synackinvalidformat" mapstructure:"SynAckInvalidFormat,omitempty"`

	// Counter for SYN ACK packet dropped because of no claims.
	SynAckMissingClaims int `json:"SynAckMissingClaims,omitempty" msgpack:"SynAckMissingClaims,omitempty" bson:"synackmissingclaims" mapstructure:"SynAckMissingClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of missing token.
	SynAckMissingToken int `json:"SynAckMissingToken,omitempty" msgpack:"SynAckMissingToken,omitempty" bson:"synackmissingtoken" mapstructure:"SynAckMissingToken,omitempty"`

	// Counter for TCP authentication option not found.
	SynAckNoTCPAuthOption int `json:"SynAckNoTCPAuthOption,omitempty" msgpack:"SynAckNoTCPAuthOption,omitempty" bson:"synacknotcpauthoption" mapstructure:"SynAckNoTCPAuthOption,omitempty"`

	// Counter for dropping because of reject rule on transmitter.
	SynAckRejected int `json:"SynAckRejected,omitempty" msgpack:"SynAckRejected,omitempty" bson:"synackrejected" mapstructure:"SynAckRejected,omitempty"`

	// Counter for SYN packet dropped because of invalid format.
	SynDroppedInvalidFormat int `json:"SynDroppedInvalidFormat,omitempty" msgpack:"SynDroppedInvalidFormat,omitempty" bson:"syndroppedinvalidformat" mapstructure:"SynDroppedInvalidFormat,omitempty"`

	// Counter for SYN packet dropped because of invalid token.
	SynDroppedInvalidToken int `json:"SynDroppedInvalidToken,omitempty" msgpack:"SynDroppedInvalidToken,omitempty" bson:"syndroppedinvalidtoken" mapstructure:"SynDroppedInvalidToken,omitempty"`

	// Counter for SYN packet dropped because of no claims.
	SynDroppedNoClaims int `json:"SynDroppedNoClaims,omitempty" msgpack:"SynDroppedNoClaims,omitempty" bson:"syndroppednoclaims" mapstructure:"SynDroppedNoClaims,omitempty"`

	// Counter for TCP authentication option not found.
	SynDroppedTCPOption int `json:"SynDroppedTCPOption,omitempty" msgpack:"SynDroppedTCPOption,omitempty" bson:"syndroppedtcpoption" mapstructure:"SynDroppedTCPOption,omitempty"`

	// Counter for SYN packet dropped due to policy.
	SynRejectPacket int `json:"SynRejectPacket,omitempty" msgpack:"SynRejectPacket,omitempty" bson:"synrejectpacket" mapstructure:"SynRejectPacket,omitempty"`

	// Counter for received SYN packet from unknown processing unit.
	SynUnexpectedPacket int `json:"SynUnexpectedPacket,omitempty" msgpack:"SynUnexpectedPacket,omitempty" bson:"synunexpectedpacket" mapstructure:"SynUnexpectedPacket,omitempty"`

	// Counter for TCP authentication option not found.
	TCPAuthNotFound int `json:"TCPAuthNotFound,omitempty" msgpack:"TCPAuthNotFound,omitempty" bson:"tcpauthnotfound" mapstructure:"TCPAuthNotFound,omitempty"`

	// Counter for UDP ACK packet dropped due to an invalid signature.
	UDPAckInvalidSignature int `json:"UDPAckInvalidSignature,omitempty" msgpack:"UDPAckInvalidSignature,omitempty" bson:"udpackinvalidsignature" mapstructure:"UDPAckInvalidSignature,omitempty"`

	// Counter for number of processed UDP connections.
	UDPConnectionsProcessed int `json:"UDPConnectionsProcessed,omitempty" msgpack:"UDPConnectionsProcessed,omitempty" bson:"udpconnectionsprocessed" mapstructure:"UDPConnectionsProcessed,omitempty"`

	// Counter for dropped UDP data packets with no context.
	UDPDropContextNotFound int `json:"UDPDropContextNotFound,omitempty" msgpack:"UDPDropContextNotFound,omitempty" bson:"udpdropcontextnotfound" mapstructure:"UDPDropContextNotFound,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPDropFin int `json:"UDPDropFin,omitempty" msgpack:"UDPDropFin,omitempty" bson:"udpdropfin" mapstructure:"UDPDropFin,omitempty"`

	// Counter for dropped UDP in NfQueue.
	UDPDropInNfQueue int `json:"UDPDropInNfQueue,omitempty" msgpack:"UDPDropInNfQueue,omitempty" bson:"udpdropinnfqueue" mapstructure:"UDPDropInNfQueue,omitempty"`

	// Counter for dropped UDP data packets with no connection.
	UDPDropNoConnection int `json:"UDPDropNoConnection,omitempty" msgpack:"UDPDropNoConnection,omitempty" bson:"udpdropnoconnection" mapstructure:"UDPDropNoConnection,omitempty"`

	// Counter for dropped UDP data packets.
	UDPDropPacket int `json:"UDPDropPacket,omitempty" msgpack:"UDPDropPacket,omitempty" bson:"udpdroppacket" mapstructure:"UDPDropPacket,omitempty"`

	// Counter for dropped UDP queue full.
	UDPDropQueueFull int `json:"UDPDropQueueFull,omitempty" msgpack:"UDPDropQueueFull,omitempty" bson:"udpdropqueuefull" mapstructure:"UDPDropQueueFull,omitempty"`

	// Counter for dropped UDP SYN ACK handshake packets.
	UDPDropSynAck int `json:"UDPDropSynAck,omitempty" msgpack:"UDPDropSynAck,omitempty" bson:"udpdropsynack" mapstructure:"UDPDropSynAck,omitempty"`

	// Counter for UDP packets received in invalid network state.
	UDPInvalidNetState int `json:"UDPInvalidNetState,omitempty" msgpack:"UDPInvalidNetState,omitempty" bson:"udpinvalidnetstate" mapstructure:"UDPInvalidNetState,omitempty"`

	// Counter for UDP packets failing postprocessing.
	UDPPostProcessingFailed int `json:"UDPPostProcessingFailed,omitempty" msgpack:"UDPPostProcessingFailed,omitempty" bson:"udppostprocessingfailed" mapstructure:"UDPPostProcessingFailed,omitempty"`

	// Counter for UDP packets failing preprocessing.
	UDPPreProcessingFailed int `json:"UDPPreProcessingFailed,omitempty" msgpack:"UDPPreProcessingFailed,omitempty" bson:"udppreprocessingfailed" mapstructure:"UDPPreProcessingFailed,omitempty"`

	// Counter for UDP packets dropped due to policy.
	UDPRejected int `json:"UDPRejected,omitempty" msgpack:"UDPRejected,omitempty" bson:"udprejected" mapstructure:"UDPRejected,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckDropBadClaims int `json:"UDPSynAckDropBadClaims,omitempty" msgpack:"UDPSynAckDropBadClaims,omitempty" bson:"udpsynackdropbadclaims" mapstructure:"UDPSynAckDropBadClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to missing claims.
	UDPSynAckMissingClaims int `json:"UDPSynAckMissingClaims,omitempty" msgpack:"UDPSynAckMissingClaims,omitempty" bson:"udpsynackmissingclaims" mapstructure:"UDPSynAckMissingClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckPolicy int `json:"UDPSynAckPolicy,omitempty" msgpack:"UDPSynAckPolicy,omitempty" bson:"udpsynackpolicy" mapstructure:"UDPSynAckPolicy,omitempty"`

	// Counter for dropped UDP SYN transmits.
	UDPSynDrop int `json:"UDPSynDrop,omitempty" msgpack:"UDPSynDrop,omitempty" bson:"udpsyndrop" mapstructure:"UDPSynDrop,omitempty"`

	// Counter for dropped UDP SYN policy.
	UDPSynDropPolicy int `json:"UDPSynDropPolicy,omitempty" msgpack:"UDPSynDropPolicy,omitempty" bson:"udpsyndroppolicy" mapstructure:"UDPSynDropPolicy,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPSynInvalidToken int `json:"UDPSynInvalidToken,omitempty" msgpack:"UDPSynInvalidToken,omitempty" bson:"udpsyninvalidtoken" mapstructure:"UDPSynInvalidToken,omitempty"`

	// Counter for UDP SYN packet dropped due to missing claims.
	UDPSynMissingClaims int `json:"UDPSynMissingClaims,omitempty" msgpack:"UDPSynMissingClaims,omitempty" bson:"udpsynmissingclaims" mapstructure:"UDPSynMissingClaims,omitempty"`

	// Counter for unknown error.
	UnknownError int `json:"UnknownError,omitempty" msgpack:"UnknownError,omitempty" bson:"unknownerror" mapstructure:"UnknownError,omitempty"`

	// Non-zero counter indicates analyzed connections for unencrypted, encrypted,
	// and packets from endpoint applications with the TCP Fast Open option set. These
	// are not dropped counter.
	ConnectionsAnalyzed int `json:"connectionsAnalyzed,omitempty" msgpack:"connectionsAnalyzed,omitempty" bson:"connectionsanalyzed" mapstructure:"connectionsAnalyzed,omitempty"`

	// Non-zero counter indicates dropped connections because of invalid state,
	// non-processing unit traffic, or out of order packets.
	ConnectionsDropped int `json:"connectionsDropped,omitempty" msgpack:"connectionsDropped,omitempty" bson:"connectionsdropped" mapstructure:"connectionsDropped,omitempty"`

	// Non-zero counter indicates expired connections because of response not being
	// received within a certain amount of time after the request is made.
	ConnectionsExpired int `json:"connectionsExpired,omitempty" msgpack:"connectionsExpired,omitempty" bson:"connectionsexpired" mapstructure:"connectionsExpired,omitempty"`

	// Non-zero counter indicates dropped packets that did not hit any of our iptables
	// rules and queue drops.
	DroppedPackets int `json:"droppedPackets,omitempty" msgpack:"droppedPackets,omitempty" bson:"droppedpackets" mapstructure:"droppedPackets,omitempty"`

	// Non-zero counter indicates encryption processing failures of data packets.
	EncryptionFailures int `json:"encryptionFailures,omitempty" msgpack:"encryptionFailures,omitempty" bson:"encryptionfailures" mapstructure:"encryptionFailures,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report. This field is deprecated. Use the
	// 'namespace' field instead.
	// field instead.
	EnforcerNamespace string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Non-zero counter indicates connections going to and from external networks.
	// These may be drops or allowed counters.
	ExternalNetworkConnections int `json:"externalNetworkConnections,omitempty" msgpack:"externalNetworkConnections,omitempty" bson:"externalnetworkconnections" mapstructure:"externalNetworkConnections,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace of the enforcer sending the report.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Non-zero counter indicates packets dropped due to a reject policy.
	PolicyDrops int `json:"policyDrops,omitempty" msgpack:"policyDrops,omitempty" bson:"policydrops" mapstructure:"policyDrops,omitempty"`

	// PUID is the ID of the processing unit reporting the counter.
	ProcessingUnitID string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"processingunitid" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit reporting the counter.
	ProcessingUnitNamespace string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"processingunitnamespace" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// Non-zero counter indicates packets rejected due to anything related to token
	// creation/parsing failures.
	TokenDrops int `json:"tokenDrops,omitempty" msgpack:"tokenDrops,omitempty" bson:"tokendrops" mapstructure:"tokenDrops,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCounterReport returns a new *CounterReport
func NewCounterReport() *CounterReport {

	return &CounterReport{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *CounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CounterReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CounterReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CounterReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCounterReport{}

	s.AckInUnknownState = o.AckInUnknownState
	s.AckInvalidFormat = o.AckInvalidFormat
	s.AckRejected = o.AckRejected
	s.AckSigValidationFailed = o.AckSigValidationFailed
	s.AckTCPNoTCPAuthOption = o.AckTCPNoTCPAuthOption
	s.ConnectionsProcessed = o.ConnectionsProcessed
	s.ContextIDNotFound = o.ContextIDNotFound
	s.DroppedExternalService = o.DroppedExternalService
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.InvalidConnState = o.InvalidConnState
	s.InvalidNetState = o.InvalidNetState
	s.InvalidProtocol = o.InvalidProtocol
	s.InvalidSynAck = o.InvalidSynAck
	s.MarkNotFound = o.MarkNotFound
	s.NetSynNotSeen = o.NetSynNotSeen
	s.NoConnFound = o.NoConnFound
	s.NonPUTraffic = o.NonPUTraffic
	s.OutOfOrderSynAck = o.OutOfOrderSynAck
	s.PortNotFound = o.PortNotFound
	s.RejectPacket = o.RejectPacket
	s.ServicePostprocessorFailed = o.ServicePostprocessorFailed
	s.ServicePreprocessorFailed = o.ServicePreprocessorFailed
	s.SynAckBadClaims = o.SynAckBadClaims
	s.SynAckClaimsMisMatch = o.SynAckClaimsMisMatch
	s.SynAckDroppedExternalService = o.SynAckDroppedExternalService
	s.SynAckInvalidFormat = o.SynAckInvalidFormat
	s.SynAckMissingClaims = o.SynAckMissingClaims
	s.SynAckMissingToken = o.SynAckMissingToken
	s.SynAckNoTCPAuthOption = o.SynAckNoTCPAuthOption
	s.SynAckRejected = o.SynAckRejected
	s.SynDroppedInvalidFormat = o.SynDroppedInvalidFormat
	s.SynDroppedInvalidToken = o.SynDroppedInvalidToken
	s.SynDroppedNoClaims = o.SynDroppedNoClaims
	s.SynDroppedTCPOption = o.SynDroppedTCPOption
	s.SynRejectPacket = o.SynRejectPacket
	s.SynUnexpectedPacket = o.SynUnexpectedPacket
	s.TCPAuthNotFound = o.TCPAuthNotFound
	s.UDPAckInvalidSignature = o.UDPAckInvalidSignature
	s.UDPConnectionsProcessed = o.UDPConnectionsProcessed
	s.UDPDropContextNotFound = o.UDPDropContextNotFound
	s.UDPDropFin = o.UDPDropFin
	s.UDPDropInNfQueue = o.UDPDropInNfQueue
	s.UDPDropNoConnection = o.UDPDropNoConnection
	s.UDPDropPacket = o.UDPDropPacket
	s.UDPDropQueueFull = o.UDPDropQueueFull
	s.UDPDropSynAck = o.UDPDropSynAck
	s.UDPInvalidNetState = o.UDPInvalidNetState
	s.UDPPostProcessingFailed = o.UDPPostProcessingFailed
	s.UDPPreProcessingFailed = o.UDPPreProcessingFailed
	s.UDPRejected = o.UDPRejected
	s.UDPSynAckDropBadClaims = o.UDPSynAckDropBadClaims
	s.UDPSynAckMissingClaims = o.UDPSynAckMissingClaims
	s.UDPSynAckPolicy = o.UDPSynAckPolicy
	s.UDPSynDrop = o.UDPSynDrop
	s.UDPSynDropPolicy = o.UDPSynDropPolicy
	s.UDPSynInvalidToken = o.UDPSynInvalidToken
	s.UDPSynMissingClaims = o.UDPSynMissingClaims
	s.UnknownError = o.UnknownError
	s.ConnectionsAnalyzed = o.ConnectionsAnalyzed
	s.ConnectionsDropped = o.ConnectionsDropped
	s.ConnectionsExpired = o.ConnectionsExpired
	s.DroppedPackets = o.DroppedPackets
	s.EncryptionFailures = o.EncryptionFailures
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.ExternalNetworkConnections = o.ExternalNetworkConnections
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.PolicyDrops = o.PolicyDrops
	s.ProcessingUnitID = o.ProcessingUnitID
	s.ProcessingUnitNamespace = o.ProcessingUnitNamespace
	s.Timestamp = o.Timestamp
	s.TokenDrops = o.TokenDrops
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CounterReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCounterReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AckInUnknownState = s.AckInUnknownState
	o.AckInvalidFormat = s.AckInvalidFormat
	o.AckRejected = s.AckRejected
	o.AckSigValidationFailed = s.AckSigValidationFailed
	o.AckTCPNoTCPAuthOption = s.AckTCPNoTCPAuthOption
	o.ConnectionsProcessed = s.ConnectionsProcessed
	o.ContextIDNotFound = s.ContextIDNotFound
	o.DroppedExternalService = s.DroppedExternalService
	o.ID = s.ID.Hex()
	o.InvalidConnState = s.InvalidConnState
	o.InvalidNetState = s.InvalidNetState
	o.InvalidProtocol = s.InvalidProtocol
	o.InvalidSynAck = s.InvalidSynAck
	o.MarkNotFound = s.MarkNotFound
	o.NetSynNotSeen = s.NetSynNotSeen
	o.NoConnFound = s.NoConnFound
	o.NonPUTraffic = s.NonPUTraffic
	o.OutOfOrderSynAck = s.OutOfOrderSynAck
	o.PortNotFound = s.PortNotFound
	o.RejectPacket = s.RejectPacket
	o.ServicePostprocessorFailed = s.ServicePostprocessorFailed
	o.ServicePreprocessorFailed = s.ServicePreprocessorFailed
	o.SynAckBadClaims = s.SynAckBadClaims
	o.SynAckClaimsMisMatch = s.SynAckClaimsMisMatch
	o.SynAckDroppedExternalService = s.SynAckDroppedExternalService
	o.SynAckInvalidFormat = s.SynAckInvalidFormat
	o.SynAckMissingClaims = s.SynAckMissingClaims
	o.SynAckMissingToken = s.SynAckMissingToken
	o.SynAckNoTCPAuthOption = s.SynAckNoTCPAuthOption
	o.SynAckRejected = s.SynAckRejected
	o.SynDroppedInvalidFormat = s.SynDroppedInvalidFormat
	o.SynDroppedInvalidToken = s.SynDroppedInvalidToken
	o.SynDroppedNoClaims = s.SynDroppedNoClaims
	o.SynDroppedTCPOption = s.SynDroppedTCPOption
	o.SynRejectPacket = s.SynRejectPacket
	o.SynUnexpectedPacket = s.SynUnexpectedPacket
	o.TCPAuthNotFound = s.TCPAuthNotFound
	o.UDPAckInvalidSignature = s.UDPAckInvalidSignature
	o.UDPConnectionsProcessed = s.UDPConnectionsProcessed
	o.UDPDropContextNotFound = s.UDPDropContextNotFound
	o.UDPDropFin = s.UDPDropFin
	o.UDPDropInNfQueue = s.UDPDropInNfQueue
	o.UDPDropNoConnection = s.UDPDropNoConnection
	o.UDPDropPacket = s.UDPDropPacket
	o.UDPDropQueueFull = s.UDPDropQueueFull
	o.UDPDropSynAck = s.UDPDropSynAck
	o.UDPInvalidNetState = s.UDPInvalidNetState
	o.UDPPostProcessingFailed = s.UDPPostProcessingFailed
	o.UDPPreProcessingFailed = s.UDPPreProcessingFailed
	o.UDPRejected = s.UDPRejected
	o.UDPSynAckDropBadClaims = s.UDPSynAckDropBadClaims
	o.UDPSynAckMissingClaims = s.UDPSynAckMissingClaims
	o.UDPSynAckPolicy = s.UDPSynAckPolicy
	o.UDPSynDrop = s.UDPSynDrop
	o.UDPSynDropPolicy = s.UDPSynDropPolicy
	o.UDPSynInvalidToken = s.UDPSynInvalidToken
	o.UDPSynMissingClaims = s.UDPSynMissingClaims
	o.UnknownError = s.UnknownError
	o.ConnectionsAnalyzed = s.ConnectionsAnalyzed
	o.ConnectionsDropped = s.ConnectionsDropped
	o.ConnectionsExpired = s.ConnectionsExpired
	o.DroppedPackets = s.DroppedPackets
	o.EncryptionFailures = s.EncryptionFailures
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.ExternalNetworkConnections = s.ExternalNetworkConnections
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.PolicyDrops = s.PolicyDrops
	o.ProcessingUnitID = s.ProcessingUnitID
	o.ProcessingUnitNamespace = s.ProcessingUnitNamespace
	o.Timestamp = s.Timestamp
	o.TokenDrops = s.TokenDrops
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CounterReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CounterReport) BleveType() string {

	return "counterreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CounterReport) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// Doc returns the documentation for the object
func (o *CounterReport) Doc() string {

	return `Post a new counter tracing report.`
}

func (o *CounterReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CounterReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CounterReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *CounterReport) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CounterReport) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *CounterReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CounterReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CounterReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CounterReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CounterReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCounterReport{
			AckInUnknownState:            &o.AckInUnknownState,
			AckInvalidFormat:             &o.AckInvalidFormat,
			AckRejected:                  &o.AckRejected,
			AckSigValidationFailed:       &o.AckSigValidationFailed,
			AckTCPNoTCPAuthOption:        &o.AckTCPNoTCPAuthOption,
			ConnectionsProcessed:         &o.ConnectionsProcessed,
			ContextIDNotFound:            &o.ContextIDNotFound,
			DroppedExternalService:       &o.DroppedExternalService,
			ID:                           &o.ID,
			InvalidConnState:             &o.InvalidConnState,
			InvalidNetState:              &o.InvalidNetState,
			InvalidProtocol:              &o.InvalidProtocol,
			InvalidSynAck:                &o.InvalidSynAck,
			MarkNotFound:                 &o.MarkNotFound,
			NetSynNotSeen:                &o.NetSynNotSeen,
			NoConnFound:                  &o.NoConnFound,
			NonPUTraffic:                 &o.NonPUTraffic,
			OutOfOrderSynAck:             &o.OutOfOrderSynAck,
			PortNotFound:                 &o.PortNotFound,
			RejectPacket:                 &o.RejectPacket,
			ServicePostprocessorFailed:   &o.ServicePostprocessorFailed,
			ServicePreprocessorFailed:    &o.ServicePreprocessorFailed,
			SynAckBadClaims:              &o.SynAckBadClaims,
			SynAckClaimsMisMatch:         &o.SynAckClaimsMisMatch,
			SynAckDroppedExternalService: &o.SynAckDroppedExternalService,
			SynAckInvalidFormat:          &o.SynAckInvalidFormat,
			SynAckMissingClaims:          &o.SynAckMissingClaims,
			SynAckMissingToken:           &o.SynAckMissingToken,
			SynAckNoTCPAuthOption:        &o.SynAckNoTCPAuthOption,
			SynAckRejected:               &o.SynAckRejected,
			SynDroppedInvalidFormat:      &o.SynDroppedInvalidFormat,
			SynDroppedInvalidToken:       &o.SynDroppedInvalidToken,
			SynDroppedNoClaims:           &o.SynDroppedNoClaims,
			SynDroppedTCPOption:          &o.SynDroppedTCPOption,
			SynRejectPacket:              &o.SynRejectPacket,
			SynUnexpectedPacket:          &o.SynUnexpectedPacket,
			TCPAuthNotFound:              &o.TCPAuthNotFound,
			UDPAckInvalidSignature:       &o.UDPAckInvalidSignature,
			UDPConnectionsProcessed:      &o.UDPConnectionsProcessed,
			UDPDropContextNotFound:       &o.UDPDropContextNotFound,
			UDPDropFin:                   &o.UDPDropFin,
			UDPDropInNfQueue:             &o.UDPDropInNfQueue,
			UDPDropNoConnection:          &o.UDPDropNoConnection,
			UDPDropPacket:                &o.UDPDropPacket,
			UDPDropQueueFull:             &o.UDPDropQueueFull,
			UDPDropSynAck:                &o.UDPDropSynAck,
			UDPInvalidNetState:           &o.UDPInvalidNetState,
			UDPPostProcessingFailed:      &o.UDPPostProcessingFailed,
			UDPPreProcessingFailed:       &o.UDPPreProcessingFailed,
			UDPRejected:                  &o.UDPRejected,
			UDPSynAckDropBadClaims:       &o.UDPSynAckDropBadClaims,
			UDPSynAckMissingClaims:       &o.UDPSynAckMissingClaims,
			UDPSynAckPolicy:              &o.UDPSynAckPolicy,
			UDPSynDrop:                   &o.UDPSynDrop,
			UDPSynDropPolicy:             &o.UDPSynDropPolicy,
			UDPSynInvalidToken:           &o.UDPSynInvalidToken,
			UDPSynMissingClaims:          &o.UDPSynMissingClaims,
			UnknownError:                 &o.UnknownError,
			ConnectionsAnalyzed:          &o.ConnectionsAnalyzed,
			ConnectionsDropped:           &o.ConnectionsDropped,
			ConnectionsExpired:           &o.ConnectionsExpired,
			DroppedPackets:               &o.DroppedPackets,
			EncryptionFailures:           &o.EncryptionFailures,
			EnforcerID:                   &o.EnforcerID,
			EnforcerNamespace:            &o.EnforcerNamespace,
			ExternalNetworkConnections:   &o.ExternalNetworkConnections,
			MigrationsLog:                &o.MigrationsLog,
			Namespace:                    &o.Namespace,
			PolicyDrops:                  &o.PolicyDrops,
			ProcessingUnitID:             &o.ProcessingUnitID,
			ProcessingUnitNamespace:      &o.ProcessingUnitNamespace,
			Timestamp:                    &o.Timestamp,
			TokenDrops:                   &o.TokenDrops,
			ZHash:                        &o.ZHash,
			Zone:                         &o.Zone,
		}
	}

	sp := &SparseCounterReport{}
	for _, f := range fields {
		switch f {
		case "AckInUnknownState":
			sp.AckInUnknownState = &(o.AckInUnknownState)
		case "AckInvalidFormat":
			sp.AckInvalidFormat = &(o.AckInvalidFormat)
		case "AckRejected":
			sp.AckRejected = &(o.AckRejected)
		case "AckSigValidationFailed":
			sp.AckSigValidationFailed = &(o.AckSigValidationFailed)
		case "AckTCPNoTCPAuthOption":
			sp.AckTCPNoTCPAuthOption = &(o.AckTCPNoTCPAuthOption)
		case "ConnectionsProcessed":
			sp.ConnectionsProcessed = &(o.ConnectionsProcessed)
		case "ContextIDNotFound":
			sp.ContextIDNotFound = &(o.ContextIDNotFound)
		case "DroppedExternalService":
			sp.DroppedExternalService = &(o.DroppedExternalService)
		case "ID":
			sp.ID = &(o.ID)
		case "InvalidConnState":
			sp.InvalidConnState = &(o.InvalidConnState)
		case "InvalidNetState":
			sp.InvalidNetState = &(o.InvalidNetState)
		case "InvalidProtocol":
			sp.InvalidProtocol = &(o.InvalidProtocol)
		case "InvalidSynAck":
			sp.InvalidSynAck = &(o.InvalidSynAck)
		case "MarkNotFound":
			sp.MarkNotFound = &(o.MarkNotFound)
		case "NetSynNotSeen":
			sp.NetSynNotSeen = &(o.NetSynNotSeen)
		case "NoConnFound":
			sp.NoConnFound = &(o.NoConnFound)
		case "NonPUTraffic":
			sp.NonPUTraffic = &(o.NonPUTraffic)
		case "OutOfOrderSynAck":
			sp.OutOfOrderSynAck = &(o.OutOfOrderSynAck)
		case "PortNotFound":
			sp.PortNotFound = &(o.PortNotFound)
		case "RejectPacket":
			sp.RejectPacket = &(o.RejectPacket)
		case "ServicePostprocessorFailed":
			sp.ServicePostprocessorFailed = &(o.ServicePostprocessorFailed)
		case "ServicePreprocessorFailed":
			sp.ServicePreprocessorFailed = &(o.ServicePreprocessorFailed)
		case "SynAckBadClaims":
			sp.SynAckBadClaims = &(o.SynAckBadClaims)
		case "SynAckClaimsMisMatch":
			sp.SynAckClaimsMisMatch = &(o.SynAckClaimsMisMatch)
		case "SynAckDroppedExternalService":
			sp.SynAckDroppedExternalService = &(o.SynAckDroppedExternalService)
		case "SynAckInvalidFormat":
			sp.SynAckInvalidFormat = &(o.SynAckInvalidFormat)
		case "SynAckMissingClaims":
			sp.SynAckMissingClaims = &(o.SynAckMissingClaims)
		case "SynAckMissingToken":
			sp.SynAckMissingToken = &(o.SynAckMissingToken)
		case "SynAckNoTCPAuthOption":
			sp.SynAckNoTCPAuthOption = &(o.SynAckNoTCPAuthOption)
		case "SynAckRejected":
			sp.SynAckRejected = &(o.SynAckRejected)
		case "SynDroppedInvalidFormat":
			sp.SynDroppedInvalidFormat = &(o.SynDroppedInvalidFormat)
		case "SynDroppedInvalidToken":
			sp.SynDroppedInvalidToken = &(o.SynDroppedInvalidToken)
		case "SynDroppedNoClaims":
			sp.SynDroppedNoClaims = &(o.SynDroppedNoClaims)
		case "SynDroppedTCPOption":
			sp.SynDroppedTCPOption = &(o.SynDroppedTCPOption)
		case "SynRejectPacket":
			sp.SynRejectPacket = &(o.SynRejectPacket)
		case "SynUnexpectedPacket":
			sp.SynUnexpectedPacket = &(o.SynUnexpectedPacket)
		case "TCPAuthNotFound":
			sp.TCPAuthNotFound = &(o.TCPAuthNotFound)
		case "UDPAckInvalidSignature":
			sp.UDPAckInvalidSignature = &(o.UDPAckInvalidSignature)
		case "UDPConnectionsProcessed":
			sp.UDPConnectionsProcessed = &(o.UDPConnectionsProcessed)
		case "UDPDropContextNotFound":
			sp.UDPDropContextNotFound = &(o.UDPDropContextNotFound)
		case "UDPDropFin":
			sp.UDPDropFin = &(o.UDPDropFin)
		case "UDPDropInNfQueue":
			sp.UDPDropInNfQueue = &(o.UDPDropInNfQueue)
		case "UDPDropNoConnection":
			sp.UDPDropNoConnection = &(o.UDPDropNoConnection)
		case "UDPDropPacket":
			sp.UDPDropPacket = &(o.UDPDropPacket)
		case "UDPDropQueueFull":
			sp.UDPDropQueueFull = &(o.UDPDropQueueFull)
		case "UDPDropSynAck":
			sp.UDPDropSynAck = &(o.UDPDropSynAck)
		case "UDPInvalidNetState":
			sp.UDPInvalidNetState = &(o.UDPInvalidNetState)
		case "UDPPostProcessingFailed":
			sp.UDPPostProcessingFailed = &(o.UDPPostProcessingFailed)
		case "UDPPreProcessingFailed":
			sp.UDPPreProcessingFailed = &(o.UDPPreProcessingFailed)
		case "UDPRejected":
			sp.UDPRejected = &(o.UDPRejected)
		case "UDPSynAckDropBadClaims":
			sp.UDPSynAckDropBadClaims = &(o.UDPSynAckDropBadClaims)
		case "UDPSynAckMissingClaims":
			sp.UDPSynAckMissingClaims = &(o.UDPSynAckMissingClaims)
		case "UDPSynAckPolicy":
			sp.UDPSynAckPolicy = &(o.UDPSynAckPolicy)
		case "UDPSynDrop":
			sp.UDPSynDrop = &(o.UDPSynDrop)
		case "UDPSynDropPolicy":
			sp.UDPSynDropPolicy = &(o.UDPSynDropPolicy)
		case "UDPSynInvalidToken":
			sp.UDPSynInvalidToken = &(o.UDPSynInvalidToken)
		case "UDPSynMissingClaims":
			sp.UDPSynMissingClaims = &(o.UDPSynMissingClaims)
		case "UnknownError":
			sp.UnknownError = &(o.UnknownError)
		case "connectionsAnalyzed":
			sp.ConnectionsAnalyzed = &(o.ConnectionsAnalyzed)
		case "connectionsDropped":
			sp.ConnectionsDropped = &(o.ConnectionsDropped)
		case "connectionsExpired":
			sp.ConnectionsExpired = &(o.ConnectionsExpired)
		case "droppedPackets":
			sp.DroppedPackets = &(o.DroppedPackets)
		case "encryptionFailures":
			sp.EncryptionFailures = &(o.EncryptionFailures)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "externalNetworkConnections":
			sp.ExternalNetworkConnections = &(o.ExternalNetworkConnections)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "policyDrops":
			sp.PolicyDrops = &(o.PolicyDrops)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "tokenDrops":
			sp.TokenDrops = &(o.TokenDrops)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCounterReport to the object.
func (o *CounterReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCounterReport)
	if so.AckInUnknownState != nil {
		o.AckInUnknownState = *so.AckInUnknownState
	}
	if so.AckInvalidFormat != nil {
		o.AckInvalidFormat = *so.AckInvalidFormat
	}
	if so.AckRejected != nil {
		o.AckRejected = *so.AckRejected
	}
	if so.AckSigValidationFailed != nil {
		o.AckSigValidationFailed = *so.AckSigValidationFailed
	}
	if so.AckTCPNoTCPAuthOption != nil {
		o.AckTCPNoTCPAuthOption = *so.AckTCPNoTCPAuthOption
	}
	if so.ConnectionsProcessed != nil {
		o.ConnectionsProcessed = *so.ConnectionsProcessed
	}
	if so.ContextIDNotFound != nil {
		o.ContextIDNotFound = *so.ContextIDNotFound
	}
	if so.DroppedExternalService != nil {
		o.DroppedExternalService = *so.DroppedExternalService
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.InvalidConnState != nil {
		o.InvalidConnState = *so.InvalidConnState
	}
	if so.InvalidNetState != nil {
		o.InvalidNetState = *so.InvalidNetState
	}
	if so.InvalidProtocol != nil {
		o.InvalidProtocol = *so.InvalidProtocol
	}
	if so.InvalidSynAck != nil {
		o.InvalidSynAck = *so.InvalidSynAck
	}
	if so.MarkNotFound != nil {
		o.MarkNotFound = *so.MarkNotFound
	}
	if so.NetSynNotSeen != nil {
		o.NetSynNotSeen = *so.NetSynNotSeen
	}
	if so.NoConnFound != nil {
		o.NoConnFound = *so.NoConnFound
	}
	if so.NonPUTraffic != nil {
		o.NonPUTraffic = *so.NonPUTraffic
	}
	if so.OutOfOrderSynAck != nil {
		o.OutOfOrderSynAck = *so.OutOfOrderSynAck
	}
	if so.PortNotFound != nil {
		o.PortNotFound = *so.PortNotFound
	}
	if so.RejectPacket != nil {
		o.RejectPacket = *so.RejectPacket
	}
	if so.ServicePostprocessorFailed != nil {
		o.ServicePostprocessorFailed = *so.ServicePostprocessorFailed
	}
	if so.ServicePreprocessorFailed != nil {
		o.ServicePreprocessorFailed = *so.ServicePreprocessorFailed
	}
	if so.SynAckBadClaims != nil {
		o.SynAckBadClaims = *so.SynAckBadClaims
	}
	if so.SynAckClaimsMisMatch != nil {
		o.SynAckClaimsMisMatch = *so.SynAckClaimsMisMatch
	}
	if so.SynAckDroppedExternalService != nil {
		o.SynAckDroppedExternalService = *so.SynAckDroppedExternalService
	}
	if so.SynAckInvalidFormat != nil {
		o.SynAckInvalidFormat = *so.SynAckInvalidFormat
	}
	if so.SynAckMissingClaims != nil {
		o.SynAckMissingClaims = *so.SynAckMissingClaims
	}
	if so.SynAckMissingToken != nil {
		o.SynAckMissingToken = *so.SynAckMissingToken
	}
	if so.SynAckNoTCPAuthOption != nil {
		o.SynAckNoTCPAuthOption = *so.SynAckNoTCPAuthOption
	}
	if so.SynAckRejected != nil {
		o.SynAckRejected = *so.SynAckRejected
	}
	if so.SynDroppedInvalidFormat != nil {
		o.SynDroppedInvalidFormat = *so.SynDroppedInvalidFormat
	}
	if so.SynDroppedInvalidToken != nil {
		o.SynDroppedInvalidToken = *so.SynDroppedInvalidToken
	}
	if so.SynDroppedNoClaims != nil {
		o.SynDroppedNoClaims = *so.SynDroppedNoClaims
	}
	if so.SynDroppedTCPOption != nil {
		o.SynDroppedTCPOption = *so.SynDroppedTCPOption
	}
	if so.SynRejectPacket != nil {
		o.SynRejectPacket = *so.SynRejectPacket
	}
	if so.SynUnexpectedPacket != nil {
		o.SynUnexpectedPacket = *so.SynUnexpectedPacket
	}
	if so.TCPAuthNotFound != nil {
		o.TCPAuthNotFound = *so.TCPAuthNotFound
	}
	if so.UDPAckInvalidSignature != nil {
		o.UDPAckInvalidSignature = *so.UDPAckInvalidSignature
	}
	if so.UDPConnectionsProcessed != nil {
		o.UDPConnectionsProcessed = *so.UDPConnectionsProcessed
	}
	if so.UDPDropContextNotFound != nil {
		o.UDPDropContextNotFound = *so.UDPDropContextNotFound
	}
	if so.UDPDropFin != nil {
		o.UDPDropFin = *so.UDPDropFin
	}
	if so.UDPDropInNfQueue != nil {
		o.UDPDropInNfQueue = *so.UDPDropInNfQueue
	}
	if so.UDPDropNoConnection != nil {
		o.UDPDropNoConnection = *so.UDPDropNoConnection
	}
	if so.UDPDropPacket != nil {
		o.UDPDropPacket = *so.UDPDropPacket
	}
	if so.UDPDropQueueFull != nil {
		o.UDPDropQueueFull = *so.UDPDropQueueFull
	}
	if so.UDPDropSynAck != nil {
		o.UDPDropSynAck = *so.UDPDropSynAck
	}
	if so.UDPInvalidNetState != nil {
		o.UDPInvalidNetState = *so.UDPInvalidNetState
	}
	if so.UDPPostProcessingFailed != nil {
		o.UDPPostProcessingFailed = *so.UDPPostProcessingFailed
	}
	if so.UDPPreProcessingFailed != nil {
		o.UDPPreProcessingFailed = *so.UDPPreProcessingFailed
	}
	if so.UDPRejected != nil {
		o.UDPRejected = *so.UDPRejected
	}
	if so.UDPSynAckDropBadClaims != nil {
		o.UDPSynAckDropBadClaims = *so.UDPSynAckDropBadClaims
	}
	if so.UDPSynAckMissingClaims != nil {
		o.UDPSynAckMissingClaims = *so.UDPSynAckMissingClaims
	}
	if so.UDPSynAckPolicy != nil {
		o.UDPSynAckPolicy = *so.UDPSynAckPolicy
	}
	if so.UDPSynDrop != nil {
		o.UDPSynDrop = *so.UDPSynDrop
	}
	if so.UDPSynDropPolicy != nil {
		o.UDPSynDropPolicy = *so.UDPSynDropPolicy
	}
	if so.UDPSynInvalidToken != nil {
		o.UDPSynInvalidToken = *so.UDPSynInvalidToken
	}
	if so.UDPSynMissingClaims != nil {
		o.UDPSynMissingClaims = *so.UDPSynMissingClaims
	}
	if so.UnknownError != nil {
		o.UnknownError = *so.UnknownError
	}
	if so.ConnectionsAnalyzed != nil {
		o.ConnectionsAnalyzed = *so.ConnectionsAnalyzed
	}
	if so.ConnectionsDropped != nil {
		o.ConnectionsDropped = *so.ConnectionsDropped
	}
	if so.ConnectionsExpired != nil {
		o.ConnectionsExpired = *so.ConnectionsExpired
	}
	if so.DroppedPackets != nil {
		o.DroppedPackets = *so.DroppedPackets
	}
	if so.EncryptionFailures != nil {
		o.EncryptionFailures = *so.EncryptionFailures
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.ExternalNetworkConnections != nil {
		o.ExternalNetworkConnections = *so.ExternalNetworkConnections
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PolicyDrops != nil {
		o.PolicyDrops = *so.PolicyDrops
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
	if so.TokenDrops != nil {
		o.TokenDrops = *so.TokenDrops
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the CounterReport.
func (o *CounterReport) DeepCopy() *CounterReport {

	if o == nil {
		return nil
	}

	out := &CounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CounterReport.
func (o *CounterReport) DeepCopyInto(out *CounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CounterReport: %s", err))
	}

	*out = *target.(*CounterReport)
}

// Validate valides the current information stored into the structure.
func (o *CounterReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	// Custom object validation.
	if err := ValidateCounterReport(o); err != nil {
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
func (*CounterReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CounterReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CounterReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CounterReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CounterReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CounterReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "AckInUnknownState":
		return o.AckInUnknownState
	case "AckInvalidFormat":
		return o.AckInvalidFormat
	case "AckRejected":
		return o.AckRejected
	case "AckSigValidationFailed":
		return o.AckSigValidationFailed
	case "AckTCPNoTCPAuthOption":
		return o.AckTCPNoTCPAuthOption
	case "ConnectionsProcessed":
		return o.ConnectionsProcessed
	case "ContextIDNotFound":
		return o.ContextIDNotFound
	case "DroppedExternalService":
		return o.DroppedExternalService
	case "ID":
		return o.ID
	case "InvalidConnState":
		return o.InvalidConnState
	case "InvalidNetState":
		return o.InvalidNetState
	case "InvalidProtocol":
		return o.InvalidProtocol
	case "InvalidSynAck":
		return o.InvalidSynAck
	case "MarkNotFound":
		return o.MarkNotFound
	case "NetSynNotSeen":
		return o.NetSynNotSeen
	case "NoConnFound":
		return o.NoConnFound
	case "NonPUTraffic":
		return o.NonPUTraffic
	case "OutOfOrderSynAck":
		return o.OutOfOrderSynAck
	case "PortNotFound":
		return o.PortNotFound
	case "RejectPacket":
		return o.RejectPacket
	case "ServicePostprocessorFailed":
		return o.ServicePostprocessorFailed
	case "ServicePreprocessorFailed":
		return o.ServicePreprocessorFailed
	case "SynAckBadClaims":
		return o.SynAckBadClaims
	case "SynAckClaimsMisMatch":
		return o.SynAckClaimsMisMatch
	case "SynAckDroppedExternalService":
		return o.SynAckDroppedExternalService
	case "SynAckInvalidFormat":
		return o.SynAckInvalidFormat
	case "SynAckMissingClaims":
		return o.SynAckMissingClaims
	case "SynAckMissingToken":
		return o.SynAckMissingToken
	case "SynAckNoTCPAuthOption":
		return o.SynAckNoTCPAuthOption
	case "SynAckRejected":
		return o.SynAckRejected
	case "SynDroppedInvalidFormat":
		return o.SynDroppedInvalidFormat
	case "SynDroppedInvalidToken":
		return o.SynDroppedInvalidToken
	case "SynDroppedNoClaims":
		return o.SynDroppedNoClaims
	case "SynDroppedTCPOption":
		return o.SynDroppedTCPOption
	case "SynRejectPacket":
		return o.SynRejectPacket
	case "SynUnexpectedPacket":
		return o.SynUnexpectedPacket
	case "TCPAuthNotFound":
		return o.TCPAuthNotFound
	case "UDPAckInvalidSignature":
		return o.UDPAckInvalidSignature
	case "UDPConnectionsProcessed":
		return o.UDPConnectionsProcessed
	case "UDPDropContextNotFound":
		return o.UDPDropContextNotFound
	case "UDPDropFin":
		return o.UDPDropFin
	case "UDPDropInNfQueue":
		return o.UDPDropInNfQueue
	case "UDPDropNoConnection":
		return o.UDPDropNoConnection
	case "UDPDropPacket":
		return o.UDPDropPacket
	case "UDPDropQueueFull":
		return o.UDPDropQueueFull
	case "UDPDropSynAck":
		return o.UDPDropSynAck
	case "UDPInvalidNetState":
		return o.UDPInvalidNetState
	case "UDPPostProcessingFailed":
		return o.UDPPostProcessingFailed
	case "UDPPreProcessingFailed":
		return o.UDPPreProcessingFailed
	case "UDPRejected":
		return o.UDPRejected
	case "UDPSynAckDropBadClaims":
		return o.UDPSynAckDropBadClaims
	case "UDPSynAckMissingClaims":
		return o.UDPSynAckMissingClaims
	case "UDPSynAckPolicy":
		return o.UDPSynAckPolicy
	case "UDPSynDrop":
		return o.UDPSynDrop
	case "UDPSynDropPolicy":
		return o.UDPSynDropPolicy
	case "UDPSynInvalidToken":
		return o.UDPSynInvalidToken
	case "UDPSynMissingClaims":
		return o.UDPSynMissingClaims
	case "UnknownError":
		return o.UnknownError
	case "connectionsAnalyzed":
		return o.ConnectionsAnalyzed
	case "connectionsDropped":
		return o.ConnectionsDropped
	case "connectionsExpired":
		return o.ConnectionsExpired
	case "droppedPackets":
		return o.DroppedPackets
	case "encryptionFailures":
		return o.EncryptionFailures
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "externalNetworkConnections":
		return o.ExternalNetworkConnections
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "policyDrops":
		return o.PolicyDrops
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "timestamp":
		return o.Timestamp
	case "tokenDrops":
		return o.TokenDrops
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CounterReportAttributesMap represents the map of attribute for CounterReport.
var CounterReportAttributesMap = map[string]elemental.AttributeSpecification{
	"AckInUnknownState": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInUnknownState",
		Description:    `Counter for sending FIN ACK received in unknown connection state.`,
		Exposed:        true,
		Name:           "AckInUnknownState",
		Stored:         true,
		Type:           "integer",
	},
	"AckInvalidFormat": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInvalidFormat",
		Description:    `Counter for ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "AckInvalidFormat",
		Stored:         true,
		Type:           "integer",
	},
	"AckRejected": {
		AllowedChoices: []string{},
		ConvertedName:  "AckRejected",
		Description:    `Counter for ACK packets rejected as per policy.`,
		Exposed:        true,
		Name:           "AckRejected",
		Stored:         true,
		Type:           "integer",
	},
	"AckSigValidationFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "AckSigValidationFailed",
		Description:    `Counter for ACK packet dropped because signature validation failed.`,
		Exposed:        true,
		Name:           "AckSigValidationFailed",
		Stored:         true,
		Type:           "integer",
	},
	"AckTCPNoTCPAuthOption": {
		AllowedChoices: []string{},
		ConvertedName:  "AckTCPNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "AckTCPNoTCPAuthOption",
		Stored:         true,
		Type:           "integer",
	},
	"ConnectionsProcessed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsProcessed",
		Description:    `Counter for connections processed.`,
		Exposed:        true,
		Name:           "ConnectionsProcessed",
		Stored:         true,
		Type:           "integer",
	},
	"ContextIDNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "ContextIDNotFound",
		Description:    `Counter for unable to find ContextID.`,
		Exposed:        true,
		Name:           "ContextIDNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"DroppedExternalService": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedExternalService",
		Description: `Counter for no ACLs found for external services. Dropping application SYN
packet.`,
		Exposed: true,
		Name:    "DroppedExternalService",
		Stored:  true,
		Type:    "integer",
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
	"InvalidConnState": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidConnState",
		Description:    `Counter for invalid connection state.`,
		Exposed:        true,
		Name:           "InvalidConnState",
		Stored:         true,
		Type:           "integer",
	},
	"InvalidNetState": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidNetState",
		Description:    `Counter for invalid net state.`,
		Exposed:        true,
		Name:           "InvalidNetState",
		Stored:         true,
		Type:           "integer",
	},
	"InvalidProtocol": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidProtocol",
		Description:    `Counter for invalid protocol.`,
		Exposed:        true,
		Name:           "InvalidProtocol",
		Stored:         true,
		Type:           "integer",
	},
	"InvalidSynAck": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidSynAck",
		Description:    `Counter for processing unit is already dead - drop SYN ACK packet.`,
		Exposed:        true,
		Name:           "InvalidSynAck",
		Stored:         true,
		Type:           "integer",
	},
	"MarkNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "MarkNotFound",
		Description:    `Counter for processing unit mark not found.`,
		Exposed:        true,
		Name:           "MarkNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"NetSynNotSeen": {
		AllowedChoices: []string{},
		ConvertedName:  "NetSynNotSeen",
		Description:    `Counter for network SYN packet was not seen.`,
		Exposed:        true,
		Name:           "NetSynNotSeen",
		Stored:         true,
		Type:           "integer",
	},
	"NoConnFound": {
		AllowedChoices: []string{},
		ConvertedName:  "NoConnFound",
		Description:    `Counter for no context or connection found.`,
		Exposed:        true,
		Name:           "NoConnFound",
		Stored:         true,
		Type:           "integer",
	},
	"NonPUTraffic": {
		AllowedChoices: []string{},
		ConvertedName:  "NonPUTraffic",
		Description:    `Counter for traffic that belongs to a non-processing unit process.`,
		Exposed:        true,
		Name:           "NonPUTraffic",
		Stored:         true,
		Type:           "integer",
	},
	"OutOfOrderSynAck": {
		AllowedChoices: []string{},
		ConvertedName:  "OutOfOrderSynAck",
		Description:    `Counter for SYN ACK for flow with processed FIN ACK.`,
		Exposed:        true,
		Name:           "OutOfOrderSynAck",
		Stored:         true,
		Type:           "integer",
	},
	"PortNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "PortNotFound",
		Description:    `Counter for port not found.`,
		Exposed:        true,
		Name:           "PortNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"RejectPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "RejectPacket",
		Description:    `Counter for reject the packet as per policy.`,
		Exposed:        true,
		Name:           "RejectPacket",
		Stored:         true,
		Type:           "integer",
	},
	"ServicePostprocessorFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePostprocessorFailed",
		Description:    `Counter for post service processing failed for network packet.`,
		Exposed:        true,
		Name:           "ServicePostprocessorFailed",
		Stored:         true,
		Type:           "integer",
	},
	"ServicePreprocessorFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePreprocessorFailed",
		Description:    `Counter for network packets that failed preprocessing.`,
		Exposed:        true,
		Name:           "ServicePreprocessorFailed",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckBadClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckBadClaims",
		Description:    `Counter for SYN ACK packet dropped because of bad claims.`,
		Exposed:        true,
		Name:           "SynAckBadClaims",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckClaimsMisMatch": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckClaimsMisMatch",
		Description:    `Counter for SYN ACK packet dropped because of encryption mismatch.`,
		Exposed:        true,
		Name:           "SynAckClaimsMisMatch",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckDroppedExternalService": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckDroppedExternalService",
		Description:    `Counter for SYN ACK from external service dropped.`,
		Exposed:        true,
		Name:           "SynAckDroppedExternalService",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckInvalidFormat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckInvalidFormat",
		Description:    `Counter for SYN ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynAckInvalidFormat",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckMissingClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingClaims",
		Description:    `Counter for SYN ACK packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynAckMissingClaims",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckMissingToken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingToken",
		Description:    `Counter for SYN ACK packet dropped because of missing token.`,
		Exposed:        true,
		Name:           "SynAckMissingToken",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckNoTCPAuthOption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynAckNoTCPAuthOption",
		Stored:         true,
		Type:           "integer",
	},
	"SynAckRejected": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckRejected",
		Description:    `Counter for dropping because of reject rule on transmitter.`,
		Exposed:        true,
		Name:           "SynAckRejected",
		Stored:         true,
		Type:           "integer",
	},
	"SynDroppedInvalidFormat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidFormat",
		Description:    `Counter for SYN packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidFormat",
		Stored:         true,
		Type:           "integer",
	},
	"SynDroppedInvalidToken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidToken",
		Description:    `Counter for SYN packet dropped because of invalid token.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidToken",
		Stored:         true,
		Type:           "integer",
	},
	"SynDroppedNoClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedNoClaims",
		Description:    `Counter for SYN packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynDroppedNoClaims",
		Stored:         true,
		Type:           "integer",
	},
	"SynDroppedTCPOption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedTCPOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynDroppedTCPOption",
		Stored:         true,
		Type:           "integer",
	},
	"SynRejectPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynRejectPacket",
		Description:    `Counter for SYN packet dropped due to policy.`,
		Exposed:        true,
		Name:           "SynRejectPacket",
		Stored:         true,
		Type:           "integer",
	},
	"SynUnexpectedPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynUnexpectedPacket",
		Description:    `Counter for received SYN packet from unknown processing unit.`,
		Exposed:        true,
		Name:           "SynUnexpectedPacket",
		Stored:         true,
		Type:           "integer",
	},
	"TCPAuthNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "TCPAuthNotFound",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "TCPAuthNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"UDPAckInvalidSignature": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPAckInvalidSignature",
		Description:    `Counter for UDP ACK packet dropped due to an invalid signature.`,
		Exposed:        true,
		Name:           "UDPAckInvalidSignature",
		Stored:         true,
		Type:           "integer",
	},
	"UDPConnectionsProcessed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPConnectionsProcessed",
		Description:    `Counter for number of processed UDP connections.`,
		Exposed:        true,
		Name:           "UDPConnectionsProcessed",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropContextNotFound": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropContextNotFound",
		Description:    `Counter for dropped UDP data packets with no context.`,
		Exposed:        true,
		Name:           "UDPDropContextNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropFin": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropFin",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropFin",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropInNfQueue": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropInNfQueue",
		Description:    `Counter for dropped UDP in NfQueue.`,
		Exposed:        true,
		Name:           "UDPDropInNfQueue",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropNoConnection": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropNoConnection",
		Description:    `Counter for dropped UDP data packets with no connection.`,
		Exposed:        true,
		Name:           "UDPDropNoConnection",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropPacket",
		Description:    `Counter for dropped UDP data packets.`,
		Exposed:        true,
		Name:           "UDPDropPacket",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropQueueFull": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropQueueFull",
		Description:    `Counter for dropped UDP queue full.`,
		Exposed:        true,
		Name:           "UDPDropQueueFull",
		Stored:         true,
		Type:           "integer",
	},
	"UDPDropSynAck": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropSynAck",
		Description:    `Counter for dropped UDP SYN ACK handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropSynAck",
		Stored:         true,
		Type:           "integer",
	},
	"UDPInvalidNetState": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPInvalidNetState",
		Description:    `Counter for UDP packets received in invalid network state.`,
		Exposed:        true,
		Name:           "UDPInvalidNetState",
		Stored:         true,
		Type:           "integer",
	},
	"UDPPostProcessingFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPostProcessingFailed",
		Description:    `Counter for UDP packets failing postprocessing.`,
		Exposed:        true,
		Name:           "UDPPostProcessingFailed",
		Stored:         true,
		Type:           "integer",
	},
	"UDPPreProcessingFailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPreProcessingFailed",
		Description:    `Counter for UDP packets failing preprocessing.`,
		Exposed:        true,
		Name:           "UDPPreProcessingFailed",
		Stored:         true,
		Type:           "integer",
	},
	"UDPRejected": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPRejected",
		Description:    `Counter for UDP packets dropped due to policy.`,
		Exposed:        true,
		Name:           "UDPRejected",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynAckDropBadClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckDropBadClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckDropBadClaims",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynAckMissingClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckMissingClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynAckMissingClaims",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynAckPolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckPolicy",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckPolicy",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynDrop": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDrop",
		Description:    `Counter for dropped UDP SYN transmits.`,
		Exposed:        true,
		Name:           "UDPSynDrop",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynDropPolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDropPolicy",
		Description:    `Counter for dropped UDP SYN policy.`,
		Exposed:        true,
		Name:           "UDPSynDropPolicy",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynInvalidToken": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynInvalidToken",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPSynInvalidToken",
		Stored:         true,
		Type:           "integer",
	},
	"UDPSynMissingClaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynMissingClaims",
		Description:    `Counter for UDP SYN packet dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynMissingClaims",
		Stored:         true,
		Type:           "integer",
	},
	"UnknownError": {
		AllowedChoices: []string{},
		ConvertedName:  "UnknownError",
		Description:    `Counter for unknown error.`,
		Exposed:        true,
		Name:           "UnknownError",
		Stored:         true,
		Type:           "integer",
	},
	"ConnectionsAnalyzed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsAnalyzed",
		Description: `Non-zero counter indicates analyzed connections for unencrypted, encrypted,
and packets from endpoint applications with the TCP Fast Open option set. These
are not dropped counter.`,
		Exposed: true,
		Name:    "connectionsAnalyzed",
		Stored:  true,
		Type:    "integer",
	},
	"ConnectionsDropped": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsDropped",
		Description: `Non-zero counter indicates dropped connections because of invalid state,
non-processing unit traffic, or out of order packets.`,
		Exposed: true,
		Name:    "connectionsDropped",
		Stored:  true,
		Type:    "integer",
	},
	"ConnectionsExpired": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsExpired",
		Description: `Non-zero counter indicates expired connections because of response not being
received within a certain amount of time after the request is made.`,
		Exposed: true,
		Name:    "connectionsExpired",
		Stored:  true,
		Type:    "integer",
	},
	"DroppedPackets": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedPackets",
		Description: `Non-zero counter indicates dropped packets that did not hit any of our iptables
rules and queue drops.`,
		Exposed: true,
		Name:    "droppedPackets",
		Stored:  true,
		Type:    "integer",
	},
	"EncryptionFailures": {
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionFailures",
		Description:    `Non-zero counter indicates encryption processing failures of data packets.`,
		Exposed:        true,
		Name:           "encryptionFailures",
		Stored:         true,
		Type:           "integer",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer sending the report.`,
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
		Description: `Namespace of the enforcer sending the report. This field is deprecated. Use the
'namespace' field instead.
field instead.`,
		Exposed: true,
		Name:    "enforcerNamespace",
		Stored:  true,
		Type:    "string",
	},
	"ExternalNetworkConnections": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworkConnections",
		Description: `Non-zero counter indicates connections going to and from external networks.
These may be drops or allowed counters.`,
		Exposed: true,
		Name:    "externalNetworkConnections",
		Stored:  true,
		Type:    "integer",
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
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"PolicyDrops": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyDrops",
		Description:    `Non-zero counter indicates packets dropped due to a reject policy.`,
		Exposed:        true,
		Name:           "policyDrops",
		Stored:         true,
		Type:           "integer",
	},
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"TokenDrops": {
		AllowedChoices: []string{},
		ConvertedName:  "TokenDrops",
		Description: `Non-zero counter indicates packets rejected due to anything related to token
creation/parsing failures.`,
		Exposed: true,
		Name:    "tokenDrops",
		Stored:  true,
		Type:    "integer",
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

// CounterReportLowerCaseAttributesMap represents the map of attribute for CounterReport.
var CounterReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"ackinunknownstate": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInUnknownState",
		Description:    `Counter for sending FIN ACK received in unknown connection state.`,
		Exposed:        true,
		Name:           "AckInUnknownState",
		Stored:         true,
		Type:           "integer",
	},
	"ackinvalidformat": {
		AllowedChoices: []string{},
		ConvertedName:  "AckInvalidFormat",
		Description:    `Counter for ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "AckInvalidFormat",
		Stored:         true,
		Type:           "integer",
	},
	"ackrejected": {
		AllowedChoices: []string{},
		ConvertedName:  "AckRejected",
		Description:    `Counter for ACK packets rejected as per policy.`,
		Exposed:        true,
		Name:           "AckRejected",
		Stored:         true,
		Type:           "integer",
	},
	"acksigvalidationfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "AckSigValidationFailed",
		Description:    `Counter for ACK packet dropped because signature validation failed.`,
		Exposed:        true,
		Name:           "AckSigValidationFailed",
		Stored:         true,
		Type:           "integer",
	},
	"acktcpnotcpauthoption": {
		AllowedChoices: []string{},
		ConvertedName:  "AckTCPNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "AckTCPNoTCPAuthOption",
		Stored:         true,
		Type:           "integer",
	},
	"connectionsprocessed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsProcessed",
		Description:    `Counter for connections processed.`,
		Exposed:        true,
		Name:           "ConnectionsProcessed",
		Stored:         true,
		Type:           "integer",
	},
	"contextidnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "ContextIDNotFound",
		Description:    `Counter for unable to find ContextID.`,
		Exposed:        true,
		Name:           "ContextIDNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"droppedexternalservice": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedExternalService",
		Description: `Counter for no ACLs found for external services. Dropping application SYN
packet.`,
		Exposed: true,
		Name:    "DroppedExternalService",
		Stored:  true,
		Type:    "integer",
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
	"invalidconnstate": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidConnState",
		Description:    `Counter for invalid connection state.`,
		Exposed:        true,
		Name:           "InvalidConnState",
		Stored:         true,
		Type:           "integer",
	},
	"invalidnetstate": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidNetState",
		Description:    `Counter for invalid net state.`,
		Exposed:        true,
		Name:           "InvalidNetState",
		Stored:         true,
		Type:           "integer",
	},
	"invalidprotocol": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidProtocol",
		Description:    `Counter for invalid protocol.`,
		Exposed:        true,
		Name:           "InvalidProtocol",
		Stored:         true,
		Type:           "integer",
	},
	"invalidsynack": {
		AllowedChoices: []string{},
		ConvertedName:  "InvalidSynAck",
		Description:    `Counter for processing unit is already dead - drop SYN ACK packet.`,
		Exposed:        true,
		Name:           "InvalidSynAck",
		Stored:         true,
		Type:           "integer",
	},
	"marknotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "MarkNotFound",
		Description:    `Counter for processing unit mark not found.`,
		Exposed:        true,
		Name:           "MarkNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"netsynnotseen": {
		AllowedChoices: []string{},
		ConvertedName:  "NetSynNotSeen",
		Description:    `Counter for network SYN packet was not seen.`,
		Exposed:        true,
		Name:           "NetSynNotSeen",
		Stored:         true,
		Type:           "integer",
	},
	"noconnfound": {
		AllowedChoices: []string{},
		ConvertedName:  "NoConnFound",
		Description:    `Counter for no context or connection found.`,
		Exposed:        true,
		Name:           "NoConnFound",
		Stored:         true,
		Type:           "integer",
	},
	"nonputraffic": {
		AllowedChoices: []string{},
		ConvertedName:  "NonPUTraffic",
		Description:    `Counter for traffic that belongs to a non-processing unit process.`,
		Exposed:        true,
		Name:           "NonPUTraffic",
		Stored:         true,
		Type:           "integer",
	},
	"outofordersynack": {
		AllowedChoices: []string{},
		ConvertedName:  "OutOfOrderSynAck",
		Description:    `Counter for SYN ACK for flow with processed FIN ACK.`,
		Exposed:        true,
		Name:           "OutOfOrderSynAck",
		Stored:         true,
		Type:           "integer",
	},
	"portnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "PortNotFound",
		Description:    `Counter for port not found.`,
		Exposed:        true,
		Name:           "PortNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"rejectpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "RejectPacket",
		Description:    `Counter for reject the packet as per policy.`,
		Exposed:        true,
		Name:           "RejectPacket",
		Stored:         true,
		Type:           "integer",
	},
	"servicepostprocessorfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePostprocessorFailed",
		Description:    `Counter for post service processing failed for network packet.`,
		Exposed:        true,
		Name:           "ServicePostprocessorFailed",
		Stored:         true,
		Type:           "integer",
	},
	"servicepreprocessorfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "ServicePreprocessorFailed",
		Description:    `Counter for network packets that failed preprocessing.`,
		Exposed:        true,
		Name:           "ServicePreprocessorFailed",
		Stored:         true,
		Type:           "integer",
	},
	"synackbadclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckBadClaims",
		Description:    `Counter for SYN ACK packet dropped because of bad claims.`,
		Exposed:        true,
		Name:           "SynAckBadClaims",
		Stored:         true,
		Type:           "integer",
	},
	"synackclaimsmismatch": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckClaimsMisMatch",
		Description:    `Counter for SYN ACK packet dropped because of encryption mismatch.`,
		Exposed:        true,
		Name:           "SynAckClaimsMisMatch",
		Stored:         true,
		Type:           "integer",
	},
	"synackdroppedexternalservice": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckDroppedExternalService",
		Description:    `Counter for SYN ACK from external service dropped.`,
		Exposed:        true,
		Name:           "SynAckDroppedExternalService",
		Stored:         true,
		Type:           "integer",
	},
	"synackinvalidformat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckInvalidFormat",
		Description:    `Counter for SYN ACK packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynAckInvalidFormat",
		Stored:         true,
		Type:           "integer",
	},
	"synackmissingclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingClaims",
		Description:    `Counter for SYN ACK packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynAckMissingClaims",
		Stored:         true,
		Type:           "integer",
	},
	"synackmissingtoken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckMissingToken",
		Description:    `Counter for SYN ACK packet dropped because of missing token.`,
		Exposed:        true,
		Name:           "SynAckMissingToken",
		Stored:         true,
		Type:           "integer",
	},
	"synacknotcpauthoption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckNoTCPAuthOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynAckNoTCPAuthOption",
		Stored:         true,
		Type:           "integer",
	},
	"synackrejected": {
		AllowedChoices: []string{},
		ConvertedName:  "SynAckRejected",
		Description:    `Counter for dropping because of reject rule on transmitter.`,
		Exposed:        true,
		Name:           "SynAckRejected",
		Stored:         true,
		Type:           "integer",
	},
	"syndroppedinvalidformat": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidFormat",
		Description:    `Counter for SYN packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidFormat",
		Stored:         true,
		Type:           "integer",
	},
	"syndroppedinvalidtoken": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedInvalidToken",
		Description:    `Counter for SYN packet dropped because of invalid token.`,
		Exposed:        true,
		Name:           "SynDroppedInvalidToken",
		Stored:         true,
		Type:           "integer",
	},
	"syndroppednoclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedNoClaims",
		Description:    `Counter for SYN packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "SynDroppedNoClaims",
		Stored:         true,
		Type:           "integer",
	},
	"syndroppedtcpoption": {
		AllowedChoices: []string{},
		ConvertedName:  "SynDroppedTCPOption",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "SynDroppedTCPOption",
		Stored:         true,
		Type:           "integer",
	},
	"synrejectpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynRejectPacket",
		Description:    `Counter for SYN packet dropped due to policy.`,
		Exposed:        true,
		Name:           "SynRejectPacket",
		Stored:         true,
		Type:           "integer",
	},
	"synunexpectedpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "SynUnexpectedPacket",
		Description:    `Counter for received SYN packet from unknown processing unit.`,
		Exposed:        true,
		Name:           "SynUnexpectedPacket",
		Stored:         true,
		Type:           "integer",
	},
	"tcpauthnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "TCPAuthNotFound",
		Description:    `Counter for TCP authentication option not found.`,
		Exposed:        true,
		Name:           "TCPAuthNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"udpackinvalidsignature": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPAckInvalidSignature",
		Description:    `Counter for UDP ACK packet dropped due to an invalid signature.`,
		Exposed:        true,
		Name:           "UDPAckInvalidSignature",
		Stored:         true,
		Type:           "integer",
	},
	"udpconnectionsprocessed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPConnectionsProcessed",
		Description:    `Counter for number of processed UDP connections.`,
		Exposed:        true,
		Name:           "UDPConnectionsProcessed",
		Stored:         true,
		Type:           "integer",
	},
	"udpdropcontextnotfound": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropContextNotFound",
		Description:    `Counter for dropped UDP data packets with no context.`,
		Exposed:        true,
		Name:           "UDPDropContextNotFound",
		Stored:         true,
		Type:           "integer",
	},
	"udpdropfin": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropFin",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropFin",
		Stored:         true,
		Type:           "integer",
	},
	"udpdropinnfqueue": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropInNfQueue",
		Description:    `Counter for dropped UDP in NfQueue.`,
		Exposed:        true,
		Name:           "UDPDropInNfQueue",
		Stored:         true,
		Type:           "integer",
	},
	"udpdropnoconnection": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropNoConnection",
		Description:    `Counter for dropped UDP data packets with no connection.`,
		Exposed:        true,
		Name:           "UDPDropNoConnection",
		Stored:         true,
		Type:           "integer",
	},
	"udpdroppacket": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropPacket",
		Description:    `Counter for dropped UDP data packets.`,
		Exposed:        true,
		Name:           "UDPDropPacket",
		Stored:         true,
		Type:           "integer",
	},
	"udpdropqueuefull": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropQueueFull",
		Description:    `Counter for dropped UDP queue full.`,
		Exposed:        true,
		Name:           "UDPDropQueueFull",
		Stored:         true,
		Type:           "integer",
	},
	"udpdropsynack": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPDropSynAck",
		Description:    `Counter for dropped UDP SYN ACK handshake packets.`,
		Exposed:        true,
		Name:           "UDPDropSynAck",
		Stored:         true,
		Type:           "integer",
	},
	"udpinvalidnetstate": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPInvalidNetState",
		Description:    `Counter for UDP packets received in invalid network state.`,
		Exposed:        true,
		Name:           "UDPInvalidNetState",
		Stored:         true,
		Type:           "integer",
	},
	"udppostprocessingfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPostProcessingFailed",
		Description:    `Counter for UDP packets failing postprocessing.`,
		Exposed:        true,
		Name:           "UDPPostProcessingFailed",
		Stored:         true,
		Type:           "integer",
	},
	"udppreprocessingfailed": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPPreProcessingFailed",
		Description:    `Counter for UDP packets failing preprocessing.`,
		Exposed:        true,
		Name:           "UDPPreProcessingFailed",
		Stored:         true,
		Type:           "integer",
	},
	"udprejected": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPRejected",
		Description:    `Counter for UDP packets dropped due to policy.`,
		Exposed:        true,
		Name:           "UDPRejected",
		Stored:         true,
		Type:           "integer",
	},
	"udpsynackdropbadclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckDropBadClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckDropBadClaims",
		Stored:         true,
		Type:           "integer",
	},
	"udpsynackmissingclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckMissingClaims",
		Description:    `Counter for UDP SYN ACK packets dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynAckMissingClaims",
		Stored:         true,
		Type:           "integer",
	},
	"udpsynackpolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynAckPolicy",
		Description:    `Counter for UDP SYN ACK packets dropped due to bad claims.`,
		Exposed:        true,
		Name:           "UDPSynAckPolicy",
		Stored:         true,
		Type:           "integer",
	},
	"udpsyndrop": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDrop",
		Description:    `Counter for dropped UDP SYN transmits.`,
		Exposed:        true,
		Name:           "UDPSynDrop",
		Stored:         true,
		Type:           "integer",
	},
	"udpsyndroppolicy": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynDropPolicy",
		Description:    `Counter for dropped UDP SYN policy.`,
		Exposed:        true,
		Name:           "UDPSynDropPolicy",
		Stored:         true,
		Type:           "integer",
	},
	"udpsyninvalidtoken": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynInvalidToken",
		Description:    `Counter for dropped UDP FIN handshake packets.`,
		Exposed:        true,
		Name:           "UDPSynInvalidToken",
		Stored:         true,
		Type:           "integer",
	},
	"udpsynmissingclaims": {
		AllowedChoices: []string{},
		ConvertedName:  "UDPSynMissingClaims",
		Description:    `Counter for UDP SYN packet dropped due to missing claims.`,
		Exposed:        true,
		Name:           "UDPSynMissingClaims",
		Stored:         true,
		Type:           "integer",
	},
	"unknownerror": {
		AllowedChoices: []string{},
		ConvertedName:  "UnknownError",
		Description:    `Counter for unknown error.`,
		Exposed:        true,
		Name:           "UnknownError",
		Stored:         true,
		Type:           "integer",
	},
	"connectionsanalyzed": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsAnalyzed",
		Description: `Non-zero counter indicates analyzed connections for unencrypted, encrypted,
and packets from endpoint applications with the TCP Fast Open option set. These
are not dropped counter.`,
		Exposed: true,
		Name:    "connectionsAnalyzed",
		Stored:  true,
		Type:    "integer",
	},
	"connectionsdropped": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsDropped",
		Description: `Non-zero counter indicates dropped connections because of invalid state,
non-processing unit traffic, or out of order packets.`,
		Exposed: true,
		Name:    "connectionsDropped",
		Stored:  true,
		Type:    "integer",
	},
	"connectionsexpired": {
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsExpired",
		Description: `Non-zero counter indicates expired connections because of response not being
received within a certain amount of time after the request is made.`,
		Exposed: true,
		Name:    "connectionsExpired",
		Stored:  true,
		Type:    "integer",
	},
	"droppedpackets": {
		AllowedChoices: []string{},
		ConvertedName:  "DroppedPackets",
		Description: `Non-zero counter indicates dropped packets that did not hit any of our iptables
rules and queue drops.`,
		Exposed: true,
		Name:    "droppedPackets",
		Stored:  true,
		Type:    "integer",
	},
	"encryptionfailures": {
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionFailures",
		Description:    `Non-zero counter indicates encryption processing failures of data packets.`,
		Exposed:        true,
		Name:           "encryptionFailures",
		Stored:         true,
		Type:           "integer",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer sending the report.`,
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
		Description: `Namespace of the enforcer sending the report. This field is deprecated. Use the
'namespace' field instead.
field instead.`,
		Exposed: true,
		Name:    "enforcerNamespace",
		Stored:  true,
		Type:    "string",
	},
	"externalnetworkconnections": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworkConnections",
		Description: `Non-zero counter indicates connections going to and from external networks.
These may be drops or allowed counters.`,
		Exposed: true,
		Name:    "externalNetworkConnections",
		Stored:  true,
		Type:    "integer",
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
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"policydrops": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyDrops",
		Description:    `Non-zero counter indicates packets dropped due to a reject policy.`,
		Exposed:        true,
		Name:           "policyDrops",
		Stored:         true,
		Type:           "integer",
	},
	"processingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"processingunitnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the processing unit reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Stored:         true,
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"tokendrops": {
		AllowedChoices: []string{},
		ConvertedName:  "TokenDrops",
		Description: `Non-zero counter indicates packets rejected due to anything related to token
creation/parsing failures.`,
		Exposed: true,
		Name:    "tokenDrops",
		Stored:  true,
		Type:    "integer",
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

// SparseCounterReportsList represents a list of SparseCounterReports
type SparseCounterReportsList []*SparseCounterReport

// Identity returns the identity of the objects in the list.
func (o SparseCounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the SparseCounterReportsList.
func (o SparseCounterReportsList) Copy() elemental.Identifiables {

	copy := append(SparseCounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCounterReportsList.
func (o SparseCounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCounterReportsList) DefaultOrder() []string {

	return []string{
		"timestamp",
	}
}

// ToPlain returns the SparseCounterReportsList converted to CounterReportsList.
func (o SparseCounterReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCounterReportsList) Version() int {

	return 1
}

// SparseCounterReport represents the sparse version of a counterreport.
type SparseCounterReport struct {
	// Counter for sending FIN ACK received in unknown connection state.
	AckInUnknownState *int `json:"AckInUnknownState,omitempty" msgpack:"AckInUnknownState,omitempty" bson:"ackinunknownstate,omitempty" mapstructure:"AckInUnknownState,omitempty"`

	// Counter for ACK packet dropped because of invalid format.
	AckInvalidFormat *int `json:"AckInvalidFormat,omitempty" msgpack:"AckInvalidFormat,omitempty" bson:"ackinvalidformat,omitempty" mapstructure:"AckInvalidFormat,omitempty"`

	// Counter for ACK packets rejected as per policy.
	AckRejected *int `json:"AckRejected,omitempty" msgpack:"AckRejected,omitempty" bson:"ackrejected,omitempty" mapstructure:"AckRejected,omitempty"`

	// Counter for ACK packet dropped because signature validation failed.
	AckSigValidationFailed *int `json:"AckSigValidationFailed,omitempty" msgpack:"AckSigValidationFailed,omitempty" bson:"acksigvalidationfailed,omitempty" mapstructure:"AckSigValidationFailed,omitempty"`

	// Counter for TCP authentication option not found.
	AckTCPNoTCPAuthOption *int `json:"AckTCPNoTCPAuthOption,omitempty" msgpack:"AckTCPNoTCPAuthOption,omitempty" bson:"acktcpnotcpauthoption,omitempty" mapstructure:"AckTCPNoTCPAuthOption,omitempty"`

	// Counter for connections processed.
	ConnectionsProcessed *int `json:"ConnectionsProcessed,omitempty" msgpack:"ConnectionsProcessed,omitempty" bson:"connectionsprocessed,omitempty" mapstructure:"ConnectionsProcessed,omitempty"`

	// Counter for unable to find ContextID.
	ContextIDNotFound *int `json:"ContextIDNotFound,omitempty" msgpack:"ContextIDNotFound,omitempty" bson:"contextidnotfound,omitempty" mapstructure:"ContextIDNotFound,omitempty"`

	// Counter for no ACLs found for external services. Dropping application SYN
	// packet.
	DroppedExternalService *int `json:"DroppedExternalService,omitempty" msgpack:"DroppedExternalService,omitempty" bson:"droppedexternalservice,omitempty" mapstructure:"DroppedExternalService,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Counter for invalid connection state.
	InvalidConnState *int `json:"InvalidConnState,omitempty" msgpack:"InvalidConnState,omitempty" bson:"invalidconnstate,omitempty" mapstructure:"InvalidConnState,omitempty"`

	// Counter for invalid net state.
	InvalidNetState *int `json:"InvalidNetState,omitempty" msgpack:"InvalidNetState,omitempty" bson:"invalidnetstate,omitempty" mapstructure:"InvalidNetState,omitempty"`

	// Counter for invalid protocol.
	InvalidProtocol *int `json:"InvalidProtocol,omitempty" msgpack:"InvalidProtocol,omitempty" bson:"invalidprotocol,omitempty" mapstructure:"InvalidProtocol,omitempty"`

	// Counter for processing unit is already dead - drop SYN ACK packet.
	InvalidSynAck *int `json:"InvalidSynAck,omitempty" msgpack:"InvalidSynAck,omitempty" bson:"invalidsynack,omitempty" mapstructure:"InvalidSynAck,omitempty"`

	// Counter for processing unit mark not found.
	MarkNotFound *int `json:"MarkNotFound,omitempty" msgpack:"MarkNotFound,omitempty" bson:"marknotfound,omitempty" mapstructure:"MarkNotFound,omitempty"`

	// Counter for network SYN packet was not seen.
	NetSynNotSeen *int `json:"NetSynNotSeen,omitempty" msgpack:"NetSynNotSeen,omitempty" bson:"netsynnotseen,omitempty" mapstructure:"NetSynNotSeen,omitempty"`

	// Counter for no context or connection found.
	NoConnFound *int `json:"NoConnFound,omitempty" msgpack:"NoConnFound,omitempty" bson:"noconnfound,omitempty" mapstructure:"NoConnFound,omitempty"`

	// Counter for traffic that belongs to a non-processing unit process.
	NonPUTraffic *int `json:"NonPUTraffic,omitempty" msgpack:"NonPUTraffic,omitempty" bson:"nonputraffic,omitempty" mapstructure:"NonPUTraffic,omitempty"`

	// Counter for SYN ACK for flow with processed FIN ACK.
	OutOfOrderSynAck *int `json:"OutOfOrderSynAck,omitempty" msgpack:"OutOfOrderSynAck,omitempty" bson:"outofordersynack,omitempty" mapstructure:"OutOfOrderSynAck,omitempty"`

	// Counter for port not found.
	PortNotFound *int `json:"PortNotFound,omitempty" msgpack:"PortNotFound,omitempty" bson:"portnotfound,omitempty" mapstructure:"PortNotFound,omitempty"`

	// Counter for reject the packet as per policy.
	RejectPacket *int `json:"RejectPacket,omitempty" msgpack:"RejectPacket,omitempty" bson:"rejectpacket,omitempty" mapstructure:"RejectPacket,omitempty"`

	// Counter for post service processing failed for network packet.
	ServicePostprocessorFailed *int `json:"ServicePostprocessorFailed,omitempty" msgpack:"ServicePostprocessorFailed,omitempty" bson:"servicepostprocessorfailed,omitempty" mapstructure:"ServicePostprocessorFailed,omitempty"`

	// Counter for network packets that failed preprocessing.
	ServicePreprocessorFailed *int `json:"ServicePreprocessorFailed,omitempty" msgpack:"ServicePreprocessorFailed,omitempty" bson:"servicepreprocessorfailed,omitempty" mapstructure:"ServicePreprocessorFailed,omitempty"`

	// Counter for SYN ACK packet dropped because of bad claims.
	SynAckBadClaims *int `json:"SynAckBadClaims,omitempty" msgpack:"SynAckBadClaims,omitempty" bson:"synackbadclaims,omitempty" mapstructure:"SynAckBadClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of encryption mismatch.
	SynAckClaimsMisMatch *int `json:"SynAckClaimsMisMatch,omitempty" msgpack:"SynAckClaimsMisMatch,omitempty" bson:"synackclaimsmismatch,omitempty" mapstructure:"SynAckClaimsMisMatch,omitempty"`

	// Counter for SYN ACK from external service dropped.
	SynAckDroppedExternalService *int `json:"SynAckDroppedExternalService,omitempty" msgpack:"SynAckDroppedExternalService,omitempty" bson:"synackdroppedexternalservice,omitempty" mapstructure:"SynAckDroppedExternalService,omitempty"`

	// Counter for SYN ACK packet dropped because of invalid format.
	SynAckInvalidFormat *int `json:"SynAckInvalidFormat,omitempty" msgpack:"SynAckInvalidFormat,omitempty" bson:"synackinvalidformat,omitempty" mapstructure:"SynAckInvalidFormat,omitempty"`

	// Counter for SYN ACK packet dropped because of no claims.
	SynAckMissingClaims *int `json:"SynAckMissingClaims,omitempty" msgpack:"SynAckMissingClaims,omitempty" bson:"synackmissingclaims,omitempty" mapstructure:"SynAckMissingClaims,omitempty"`

	// Counter for SYN ACK packet dropped because of missing token.
	SynAckMissingToken *int `json:"SynAckMissingToken,omitempty" msgpack:"SynAckMissingToken,omitempty" bson:"synackmissingtoken,omitempty" mapstructure:"SynAckMissingToken,omitempty"`

	// Counter for TCP authentication option not found.
	SynAckNoTCPAuthOption *int `json:"SynAckNoTCPAuthOption,omitempty" msgpack:"SynAckNoTCPAuthOption,omitempty" bson:"synacknotcpauthoption,omitempty" mapstructure:"SynAckNoTCPAuthOption,omitempty"`

	// Counter for dropping because of reject rule on transmitter.
	SynAckRejected *int `json:"SynAckRejected,omitempty" msgpack:"SynAckRejected,omitempty" bson:"synackrejected,omitempty" mapstructure:"SynAckRejected,omitempty"`

	// Counter for SYN packet dropped because of invalid format.
	SynDroppedInvalidFormat *int `json:"SynDroppedInvalidFormat,omitempty" msgpack:"SynDroppedInvalidFormat,omitempty" bson:"syndroppedinvalidformat,omitempty" mapstructure:"SynDroppedInvalidFormat,omitempty"`

	// Counter for SYN packet dropped because of invalid token.
	SynDroppedInvalidToken *int `json:"SynDroppedInvalidToken,omitempty" msgpack:"SynDroppedInvalidToken,omitempty" bson:"syndroppedinvalidtoken,omitempty" mapstructure:"SynDroppedInvalidToken,omitempty"`

	// Counter for SYN packet dropped because of no claims.
	SynDroppedNoClaims *int `json:"SynDroppedNoClaims,omitempty" msgpack:"SynDroppedNoClaims,omitempty" bson:"syndroppednoclaims,omitempty" mapstructure:"SynDroppedNoClaims,omitempty"`

	// Counter for TCP authentication option not found.
	SynDroppedTCPOption *int `json:"SynDroppedTCPOption,omitempty" msgpack:"SynDroppedTCPOption,omitempty" bson:"syndroppedtcpoption,omitempty" mapstructure:"SynDroppedTCPOption,omitempty"`

	// Counter for SYN packet dropped due to policy.
	SynRejectPacket *int `json:"SynRejectPacket,omitempty" msgpack:"SynRejectPacket,omitempty" bson:"synrejectpacket,omitempty" mapstructure:"SynRejectPacket,omitempty"`

	// Counter for received SYN packet from unknown processing unit.
	SynUnexpectedPacket *int `json:"SynUnexpectedPacket,omitempty" msgpack:"SynUnexpectedPacket,omitempty" bson:"synunexpectedpacket,omitempty" mapstructure:"SynUnexpectedPacket,omitempty"`

	// Counter for TCP authentication option not found.
	TCPAuthNotFound *int `json:"TCPAuthNotFound,omitempty" msgpack:"TCPAuthNotFound,omitempty" bson:"tcpauthnotfound,omitempty" mapstructure:"TCPAuthNotFound,omitempty"`

	// Counter for UDP ACK packet dropped due to an invalid signature.
	UDPAckInvalidSignature *int `json:"UDPAckInvalidSignature,omitempty" msgpack:"UDPAckInvalidSignature,omitempty" bson:"udpackinvalidsignature,omitempty" mapstructure:"UDPAckInvalidSignature,omitempty"`

	// Counter for number of processed UDP connections.
	UDPConnectionsProcessed *int `json:"UDPConnectionsProcessed,omitempty" msgpack:"UDPConnectionsProcessed,omitempty" bson:"udpconnectionsprocessed,omitempty" mapstructure:"UDPConnectionsProcessed,omitempty"`

	// Counter for dropped UDP data packets with no context.
	UDPDropContextNotFound *int `json:"UDPDropContextNotFound,omitempty" msgpack:"UDPDropContextNotFound,omitempty" bson:"udpdropcontextnotfound,omitempty" mapstructure:"UDPDropContextNotFound,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPDropFin *int `json:"UDPDropFin,omitempty" msgpack:"UDPDropFin,omitempty" bson:"udpdropfin,omitempty" mapstructure:"UDPDropFin,omitempty"`

	// Counter for dropped UDP in NfQueue.
	UDPDropInNfQueue *int `json:"UDPDropInNfQueue,omitempty" msgpack:"UDPDropInNfQueue,omitempty" bson:"udpdropinnfqueue,omitempty" mapstructure:"UDPDropInNfQueue,omitempty"`

	// Counter for dropped UDP data packets with no connection.
	UDPDropNoConnection *int `json:"UDPDropNoConnection,omitempty" msgpack:"UDPDropNoConnection,omitempty" bson:"udpdropnoconnection,omitempty" mapstructure:"UDPDropNoConnection,omitempty"`

	// Counter for dropped UDP data packets.
	UDPDropPacket *int `json:"UDPDropPacket,omitempty" msgpack:"UDPDropPacket,omitempty" bson:"udpdroppacket,omitempty" mapstructure:"UDPDropPacket,omitempty"`

	// Counter for dropped UDP queue full.
	UDPDropQueueFull *int `json:"UDPDropQueueFull,omitempty" msgpack:"UDPDropQueueFull,omitempty" bson:"udpdropqueuefull,omitempty" mapstructure:"UDPDropQueueFull,omitempty"`

	// Counter for dropped UDP SYN ACK handshake packets.
	UDPDropSynAck *int `json:"UDPDropSynAck,omitempty" msgpack:"UDPDropSynAck,omitempty" bson:"udpdropsynack,omitempty" mapstructure:"UDPDropSynAck,omitempty"`

	// Counter for UDP packets received in invalid network state.
	UDPInvalidNetState *int `json:"UDPInvalidNetState,omitempty" msgpack:"UDPInvalidNetState,omitempty" bson:"udpinvalidnetstate,omitempty" mapstructure:"UDPInvalidNetState,omitempty"`

	// Counter for UDP packets failing postprocessing.
	UDPPostProcessingFailed *int `json:"UDPPostProcessingFailed,omitempty" msgpack:"UDPPostProcessingFailed,omitempty" bson:"udppostprocessingfailed,omitempty" mapstructure:"UDPPostProcessingFailed,omitempty"`

	// Counter for UDP packets failing preprocessing.
	UDPPreProcessingFailed *int `json:"UDPPreProcessingFailed,omitempty" msgpack:"UDPPreProcessingFailed,omitempty" bson:"udppreprocessingfailed,omitempty" mapstructure:"UDPPreProcessingFailed,omitempty"`

	// Counter for UDP packets dropped due to policy.
	UDPRejected *int `json:"UDPRejected,omitempty" msgpack:"UDPRejected,omitempty" bson:"udprejected,omitempty" mapstructure:"UDPRejected,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckDropBadClaims *int `json:"UDPSynAckDropBadClaims,omitempty" msgpack:"UDPSynAckDropBadClaims,omitempty" bson:"udpsynackdropbadclaims,omitempty" mapstructure:"UDPSynAckDropBadClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to missing claims.
	UDPSynAckMissingClaims *int `json:"UDPSynAckMissingClaims,omitempty" msgpack:"UDPSynAckMissingClaims,omitempty" bson:"udpsynackmissingclaims,omitempty" mapstructure:"UDPSynAckMissingClaims,omitempty"`

	// Counter for UDP SYN ACK packets dropped due to bad claims.
	UDPSynAckPolicy *int `json:"UDPSynAckPolicy,omitempty" msgpack:"UDPSynAckPolicy,omitempty" bson:"udpsynackpolicy,omitempty" mapstructure:"UDPSynAckPolicy,omitempty"`

	// Counter for dropped UDP SYN transmits.
	UDPSynDrop *int `json:"UDPSynDrop,omitempty" msgpack:"UDPSynDrop,omitempty" bson:"udpsyndrop,omitempty" mapstructure:"UDPSynDrop,omitempty"`

	// Counter for dropped UDP SYN policy.
	UDPSynDropPolicy *int `json:"UDPSynDropPolicy,omitempty" msgpack:"UDPSynDropPolicy,omitempty" bson:"udpsyndroppolicy,omitempty" mapstructure:"UDPSynDropPolicy,omitempty"`

	// Counter for dropped UDP FIN handshake packets.
	UDPSynInvalidToken *int `json:"UDPSynInvalidToken,omitempty" msgpack:"UDPSynInvalidToken,omitempty" bson:"udpsyninvalidtoken,omitempty" mapstructure:"UDPSynInvalidToken,omitempty"`

	// Counter for UDP SYN packet dropped due to missing claims.
	UDPSynMissingClaims *int `json:"UDPSynMissingClaims,omitempty" msgpack:"UDPSynMissingClaims,omitempty" bson:"udpsynmissingclaims,omitempty" mapstructure:"UDPSynMissingClaims,omitempty"`

	// Counter for unknown error.
	UnknownError *int `json:"UnknownError,omitempty" msgpack:"UnknownError,omitempty" bson:"unknownerror,omitempty" mapstructure:"UnknownError,omitempty"`

	// Non-zero counter indicates analyzed connections for unencrypted, encrypted,
	// and packets from endpoint applications with the TCP Fast Open option set. These
	// are not dropped counter.
	ConnectionsAnalyzed *int `json:"connectionsAnalyzed,omitempty" msgpack:"connectionsAnalyzed,omitempty" bson:"connectionsanalyzed,omitempty" mapstructure:"connectionsAnalyzed,omitempty"`

	// Non-zero counter indicates dropped connections because of invalid state,
	// non-processing unit traffic, or out of order packets.
	ConnectionsDropped *int `json:"connectionsDropped,omitempty" msgpack:"connectionsDropped,omitempty" bson:"connectionsdropped,omitempty" mapstructure:"connectionsDropped,omitempty"`

	// Non-zero counter indicates expired connections because of response not being
	// received within a certain amount of time after the request is made.
	ConnectionsExpired *int `json:"connectionsExpired,omitempty" msgpack:"connectionsExpired,omitempty" bson:"connectionsexpired,omitempty" mapstructure:"connectionsExpired,omitempty"`

	// Non-zero counter indicates dropped packets that did not hit any of our iptables
	// rules and queue drops.
	DroppedPackets *int `json:"droppedPackets,omitempty" msgpack:"droppedPackets,omitempty" bson:"droppedpackets,omitempty" mapstructure:"droppedPackets,omitempty"`

	// Non-zero counter indicates encryption processing failures of data packets.
	EncryptionFailures *int `json:"encryptionFailures,omitempty" msgpack:"encryptionFailures,omitempty" bson:"encryptionfailures,omitempty" mapstructure:"encryptionFailures,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report. This field is deprecated. Use the
	// 'namespace' field instead.
	// field instead.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Non-zero counter indicates connections going to and from external networks.
	// These may be drops or allowed counters.
	ExternalNetworkConnections *int `json:"externalNetworkConnections,omitempty" msgpack:"externalNetworkConnections,omitempty" bson:"externalnetworkconnections,omitempty" mapstructure:"externalNetworkConnections,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace of the enforcer sending the report.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Non-zero counter indicates packets dropped due to a reject policy.
	PolicyDrops *int `json:"policyDrops,omitempty" msgpack:"policyDrops,omitempty" bson:"policydrops,omitempty" mapstructure:"policyDrops,omitempty"`

	// PUID is the ID of the processing unit reporting the counter.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"processingunitid,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the processing unit reporting the counter.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"processingunitnamespace,omitempty" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// Non-zero counter indicates packets rejected due to anything related to token
	// creation/parsing failures.
	TokenDrops *int `json:"tokenDrops,omitempty" msgpack:"tokenDrops,omitempty" bson:"tokendrops,omitempty" mapstructure:"tokenDrops,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCounterReport returns a new  SparseCounterReport.
func NewSparseCounterReport() *SparseCounterReport {
	return &SparseCounterReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCounterReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCounterReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCounterReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCounterReport{}

	if o.AckInUnknownState != nil {
		s.AckInUnknownState = o.AckInUnknownState
	}
	if o.AckInvalidFormat != nil {
		s.AckInvalidFormat = o.AckInvalidFormat
	}
	if o.AckRejected != nil {
		s.AckRejected = o.AckRejected
	}
	if o.AckSigValidationFailed != nil {
		s.AckSigValidationFailed = o.AckSigValidationFailed
	}
	if o.AckTCPNoTCPAuthOption != nil {
		s.AckTCPNoTCPAuthOption = o.AckTCPNoTCPAuthOption
	}
	if o.ConnectionsProcessed != nil {
		s.ConnectionsProcessed = o.ConnectionsProcessed
	}
	if o.ContextIDNotFound != nil {
		s.ContextIDNotFound = o.ContextIDNotFound
	}
	if o.DroppedExternalService != nil {
		s.DroppedExternalService = o.DroppedExternalService
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.InvalidConnState != nil {
		s.InvalidConnState = o.InvalidConnState
	}
	if o.InvalidNetState != nil {
		s.InvalidNetState = o.InvalidNetState
	}
	if o.InvalidProtocol != nil {
		s.InvalidProtocol = o.InvalidProtocol
	}
	if o.InvalidSynAck != nil {
		s.InvalidSynAck = o.InvalidSynAck
	}
	if o.MarkNotFound != nil {
		s.MarkNotFound = o.MarkNotFound
	}
	if o.NetSynNotSeen != nil {
		s.NetSynNotSeen = o.NetSynNotSeen
	}
	if o.NoConnFound != nil {
		s.NoConnFound = o.NoConnFound
	}
	if o.NonPUTraffic != nil {
		s.NonPUTraffic = o.NonPUTraffic
	}
	if o.OutOfOrderSynAck != nil {
		s.OutOfOrderSynAck = o.OutOfOrderSynAck
	}
	if o.PortNotFound != nil {
		s.PortNotFound = o.PortNotFound
	}
	if o.RejectPacket != nil {
		s.RejectPacket = o.RejectPacket
	}
	if o.ServicePostprocessorFailed != nil {
		s.ServicePostprocessorFailed = o.ServicePostprocessorFailed
	}
	if o.ServicePreprocessorFailed != nil {
		s.ServicePreprocessorFailed = o.ServicePreprocessorFailed
	}
	if o.SynAckBadClaims != nil {
		s.SynAckBadClaims = o.SynAckBadClaims
	}
	if o.SynAckClaimsMisMatch != nil {
		s.SynAckClaimsMisMatch = o.SynAckClaimsMisMatch
	}
	if o.SynAckDroppedExternalService != nil {
		s.SynAckDroppedExternalService = o.SynAckDroppedExternalService
	}
	if o.SynAckInvalidFormat != nil {
		s.SynAckInvalidFormat = o.SynAckInvalidFormat
	}
	if o.SynAckMissingClaims != nil {
		s.SynAckMissingClaims = o.SynAckMissingClaims
	}
	if o.SynAckMissingToken != nil {
		s.SynAckMissingToken = o.SynAckMissingToken
	}
	if o.SynAckNoTCPAuthOption != nil {
		s.SynAckNoTCPAuthOption = o.SynAckNoTCPAuthOption
	}
	if o.SynAckRejected != nil {
		s.SynAckRejected = o.SynAckRejected
	}
	if o.SynDroppedInvalidFormat != nil {
		s.SynDroppedInvalidFormat = o.SynDroppedInvalidFormat
	}
	if o.SynDroppedInvalidToken != nil {
		s.SynDroppedInvalidToken = o.SynDroppedInvalidToken
	}
	if o.SynDroppedNoClaims != nil {
		s.SynDroppedNoClaims = o.SynDroppedNoClaims
	}
	if o.SynDroppedTCPOption != nil {
		s.SynDroppedTCPOption = o.SynDroppedTCPOption
	}
	if o.SynRejectPacket != nil {
		s.SynRejectPacket = o.SynRejectPacket
	}
	if o.SynUnexpectedPacket != nil {
		s.SynUnexpectedPacket = o.SynUnexpectedPacket
	}
	if o.TCPAuthNotFound != nil {
		s.TCPAuthNotFound = o.TCPAuthNotFound
	}
	if o.UDPAckInvalidSignature != nil {
		s.UDPAckInvalidSignature = o.UDPAckInvalidSignature
	}
	if o.UDPConnectionsProcessed != nil {
		s.UDPConnectionsProcessed = o.UDPConnectionsProcessed
	}
	if o.UDPDropContextNotFound != nil {
		s.UDPDropContextNotFound = o.UDPDropContextNotFound
	}
	if o.UDPDropFin != nil {
		s.UDPDropFin = o.UDPDropFin
	}
	if o.UDPDropInNfQueue != nil {
		s.UDPDropInNfQueue = o.UDPDropInNfQueue
	}
	if o.UDPDropNoConnection != nil {
		s.UDPDropNoConnection = o.UDPDropNoConnection
	}
	if o.UDPDropPacket != nil {
		s.UDPDropPacket = o.UDPDropPacket
	}
	if o.UDPDropQueueFull != nil {
		s.UDPDropQueueFull = o.UDPDropQueueFull
	}
	if o.UDPDropSynAck != nil {
		s.UDPDropSynAck = o.UDPDropSynAck
	}
	if o.UDPInvalidNetState != nil {
		s.UDPInvalidNetState = o.UDPInvalidNetState
	}
	if o.UDPPostProcessingFailed != nil {
		s.UDPPostProcessingFailed = o.UDPPostProcessingFailed
	}
	if o.UDPPreProcessingFailed != nil {
		s.UDPPreProcessingFailed = o.UDPPreProcessingFailed
	}
	if o.UDPRejected != nil {
		s.UDPRejected = o.UDPRejected
	}
	if o.UDPSynAckDropBadClaims != nil {
		s.UDPSynAckDropBadClaims = o.UDPSynAckDropBadClaims
	}
	if o.UDPSynAckMissingClaims != nil {
		s.UDPSynAckMissingClaims = o.UDPSynAckMissingClaims
	}
	if o.UDPSynAckPolicy != nil {
		s.UDPSynAckPolicy = o.UDPSynAckPolicy
	}
	if o.UDPSynDrop != nil {
		s.UDPSynDrop = o.UDPSynDrop
	}
	if o.UDPSynDropPolicy != nil {
		s.UDPSynDropPolicy = o.UDPSynDropPolicy
	}
	if o.UDPSynInvalidToken != nil {
		s.UDPSynInvalidToken = o.UDPSynInvalidToken
	}
	if o.UDPSynMissingClaims != nil {
		s.UDPSynMissingClaims = o.UDPSynMissingClaims
	}
	if o.UnknownError != nil {
		s.UnknownError = o.UnknownError
	}
	if o.ConnectionsAnalyzed != nil {
		s.ConnectionsAnalyzed = o.ConnectionsAnalyzed
	}
	if o.ConnectionsDropped != nil {
		s.ConnectionsDropped = o.ConnectionsDropped
	}
	if o.ConnectionsExpired != nil {
		s.ConnectionsExpired = o.ConnectionsExpired
	}
	if o.DroppedPackets != nil {
		s.DroppedPackets = o.DroppedPackets
	}
	if o.EncryptionFailures != nil {
		s.EncryptionFailures = o.EncryptionFailures
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.ExternalNetworkConnections != nil {
		s.ExternalNetworkConnections = o.ExternalNetworkConnections
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PolicyDrops != nil {
		s.PolicyDrops = o.PolicyDrops
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
	if o.TokenDrops != nil {
		s.TokenDrops = o.TokenDrops
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
func (o *SparseCounterReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCounterReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.AckInUnknownState != nil {
		o.AckInUnknownState = s.AckInUnknownState
	}
	if s.AckInvalidFormat != nil {
		o.AckInvalidFormat = s.AckInvalidFormat
	}
	if s.AckRejected != nil {
		o.AckRejected = s.AckRejected
	}
	if s.AckSigValidationFailed != nil {
		o.AckSigValidationFailed = s.AckSigValidationFailed
	}
	if s.AckTCPNoTCPAuthOption != nil {
		o.AckTCPNoTCPAuthOption = s.AckTCPNoTCPAuthOption
	}
	if s.ConnectionsProcessed != nil {
		o.ConnectionsProcessed = s.ConnectionsProcessed
	}
	if s.ContextIDNotFound != nil {
		o.ContextIDNotFound = s.ContextIDNotFound
	}
	if s.DroppedExternalService != nil {
		o.DroppedExternalService = s.DroppedExternalService
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.InvalidConnState != nil {
		o.InvalidConnState = s.InvalidConnState
	}
	if s.InvalidNetState != nil {
		o.InvalidNetState = s.InvalidNetState
	}
	if s.InvalidProtocol != nil {
		o.InvalidProtocol = s.InvalidProtocol
	}
	if s.InvalidSynAck != nil {
		o.InvalidSynAck = s.InvalidSynAck
	}
	if s.MarkNotFound != nil {
		o.MarkNotFound = s.MarkNotFound
	}
	if s.NetSynNotSeen != nil {
		o.NetSynNotSeen = s.NetSynNotSeen
	}
	if s.NoConnFound != nil {
		o.NoConnFound = s.NoConnFound
	}
	if s.NonPUTraffic != nil {
		o.NonPUTraffic = s.NonPUTraffic
	}
	if s.OutOfOrderSynAck != nil {
		o.OutOfOrderSynAck = s.OutOfOrderSynAck
	}
	if s.PortNotFound != nil {
		o.PortNotFound = s.PortNotFound
	}
	if s.RejectPacket != nil {
		o.RejectPacket = s.RejectPacket
	}
	if s.ServicePostprocessorFailed != nil {
		o.ServicePostprocessorFailed = s.ServicePostprocessorFailed
	}
	if s.ServicePreprocessorFailed != nil {
		o.ServicePreprocessorFailed = s.ServicePreprocessorFailed
	}
	if s.SynAckBadClaims != nil {
		o.SynAckBadClaims = s.SynAckBadClaims
	}
	if s.SynAckClaimsMisMatch != nil {
		o.SynAckClaimsMisMatch = s.SynAckClaimsMisMatch
	}
	if s.SynAckDroppedExternalService != nil {
		o.SynAckDroppedExternalService = s.SynAckDroppedExternalService
	}
	if s.SynAckInvalidFormat != nil {
		o.SynAckInvalidFormat = s.SynAckInvalidFormat
	}
	if s.SynAckMissingClaims != nil {
		o.SynAckMissingClaims = s.SynAckMissingClaims
	}
	if s.SynAckMissingToken != nil {
		o.SynAckMissingToken = s.SynAckMissingToken
	}
	if s.SynAckNoTCPAuthOption != nil {
		o.SynAckNoTCPAuthOption = s.SynAckNoTCPAuthOption
	}
	if s.SynAckRejected != nil {
		o.SynAckRejected = s.SynAckRejected
	}
	if s.SynDroppedInvalidFormat != nil {
		o.SynDroppedInvalidFormat = s.SynDroppedInvalidFormat
	}
	if s.SynDroppedInvalidToken != nil {
		o.SynDroppedInvalidToken = s.SynDroppedInvalidToken
	}
	if s.SynDroppedNoClaims != nil {
		o.SynDroppedNoClaims = s.SynDroppedNoClaims
	}
	if s.SynDroppedTCPOption != nil {
		o.SynDroppedTCPOption = s.SynDroppedTCPOption
	}
	if s.SynRejectPacket != nil {
		o.SynRejectPacket = s.SynRejectPacket
	}
	if s.SynUnexpectedPacket != nil {
		o.SynUnexpectedPacket = s.SynUnexpectedPacket
	}
	if s.TCPAuthNotFound != nil {
		o.TCPAuthNotFound = s.TCPAuthNotFound
	}
	if s.UDPAckInvalidSignature != nil {
		o.UDPAckInvalidSignature = s.UDPAckInvalidSignature
	}
	if s.UDPConnectionsProcessed != nil {
		o.UDPConnectionsProcessed = s.UDPConnectionsProcessed
	}
	if s.UDPDropContextNotFound != nil {
		o.UDPDropContextNotFound = s.UDPDropContextNotFound
	}
	if s.UDPDropFin != nil {
		o.UDPDropFin = s.UDPDropFin
	}
	if s.UDPDropInNfQueue != nil {
		o.UDPDropInNfQueue = s.UDPDropInNfQueue
	}
	if s.UDPDropNoConnection != nil {
		o.UDPDropNoConnection = s.UDPDropNoConnection
	}
	if s.UDPDropPacket != nil {
		o.UDPDropPacket = s.UDPDropPacket
	}
	if s.UDPDropQueueFull != nil {
		o.UDPDropQueueFull = s.UDPDropQueueFull
	}
	if s.UDPDropSynAck != nil {
		o.UDPDropSynAck = s.UDPDropSynAck
	}
	if s.UDPInvalidNetState != nil {
		o.UDPInvalidNetState = s.UDPInvalidNetState
	}
	if s.UDPPostProcessingFailed != nil {
		o.UDPPostProcessingFailed = s.UDPPostProcessingFailed
	}
	if s.UDPPreProcessingFailed != nil {
		o.UDPPreProcessingFailed = s.UDPPreProcessingFailed
	}
	if s.UDPRejected != nil {
		o.UDPRejected = s.UDPRejected
	}
	if s.UDPSynAckDropBadClaims != nil {
		o.UDPSynAckDropBadClaims = s.UDPSynAckDropBadClaims
	}
	if s.UDPSynAckMissingClaims != nil {
		o.UDPSynAckMissingClaims = s.UDPSynAckMissingClaims
	}
	if s.UDPSynAckPolicy != nil {
		o.UDPSynAckPolicy = s.UDPSynAckPolicy
	}
	if s.UDPSynDrop != nil {
		o.UDPSynDrop = s.UDPSynDrop
	}
	if s.UDPSynDropPolicy != nil {
		o.UDPSynDropPolicy = s.UDPSynDropPolicy
	}
	if s.UDPSynInvalidToken != nil {
		o.UDPSynInvalidToken = s.UDPSynInvalidToken
	}
	if s.UDPSynMissingClaims != nil {
		o.UDPSynMissingClaims = s.UDPSynMissingClaims
	}
	if s.UnknownError != nil {
		o.UnknownError = s.UnknownError
	}
	if s.ConnectionsAnalyzed != nil {
		o.ConnectionsAnalyzed = s.ConnectionsAnalyzed
	}
	if s.ConnectionsDropped != nil {
		o.ConnectionsDropped = s.ConnectionsDropped
	}
	if s.ConnectionsExpired != nil {
		o.ConnectionsExpired = s.ConnectionsExpired
	}
	if s.DroppedPackets != nil {
		o.DroppedPackets = s.DroppedPackets
	}
	if s.EncryptionFailures != nil {
		o.EncryptionFailures = s.EncryptionFailures
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.ExternalNetworkConnections != nil {
		o.ExternalNetworkConnections = s.ExternalNetworkConnections
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PolicyDrops != nil {
		o.PolicyDrops = s.PolicyDrops
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
	if s.TokenDrops != nil {
		o.TokenDrops = s.TokenDrops
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
func (o *SparseCounterReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCounterReport) ToPlain() elemental.PlainIdentifiable {

	out := NewCounterReport()
	if o.AckInUnknownState != nil {
		out.AckInUnknownState = *o.AckInUnknownState
	}
	if o.AckInvalidFormat != nil {
		out.AckInvalidFormat = *o.AckInvalidFormat
	}
	if o.AckRejected != nil {
		out.AckRejected = *o.AckRejected
	}
	if o.AckSigValidationFailed != nil {
		out.AckSigValidationFailed = *o.AckSigValidationFailed
	}
	if o.AckTCPNoTCPAuthOption != nil {
		out.AckTCPNoTCPAuthOption = *o.AckTCPNoTCPAuthOption
	}
	if o.ConnectionsProcessed != nil {
		out.ConnectionsProcessed = *o.ConnectionsProcessed
	}
	if o.ContextIDNotFound != nil {
		out.ContextIDNotFound = *o.ContextIDNotFound
	}
	if o.DroppedExternalService != nil {
		out.DroppedExternalService = *o.DroppedExternalService
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.InvalidConnState != nil {
		out.InvalidConnState = *o.InvalidConnState
	}
	if o.InvalidNetState != nil {
		out.InvalidNetState = *o.InvalidNetState
	}
	if o.InvalidProtocol != nil {
		out.InvalidProtocol = *o.InvalidProtocol
	}
	if o.InvalidSynAck != nil {
		out.InvalidSynAck = *o.InvalidSynAck
	}
	if o.MarkNotFound != nil {
		out.MarkNotFound = *o.MarkNotFound
	}
	if o.NetSynNotSeen != nil {
		out.NetSynNotSeen = *o.NetSynNotSeen
	}
	if o.NoConnFound != nil {
		out.NoConnFound = *o.NoConnFound
	}
	if o.NonPUTraffic != nil {
		out.NonPUTraffic = *o.NonPUTraffic
	}
	if o.OutOfOrderSynAck != nil {
		out.OutOfOrderSynAck = *o.OutOfOrderSynAck
	}
	if o.PortNotFound != nil {
		out.PortNotFound = *o.PortNotFound
	}
	if o.RejectPacket != nil {
		out.RejectPacket = *o.RejectPacket
	}
	if o.ServicePostprocessorFailed != nil {
		out.ServicePostprocessorFailed = *o.ServicePostprocessorFailed
	}
	if o.ServicePreprocessorFailed != nil {
		out.ServicePreprocessorFailed = *o.ServicePreprocessorFailed
	}
	if o.SynAckBadClaims != nil {
		out.SynAckBadClaims = *o.SynAckBadClaims
	}
	if o.SynAckClaimsMisMatch != nil {
		out.SynAckClaimsMisMatch = *o.SynAckClaimsMisMatch
	}
	if o.SynAckDroppedExternalService != nil {
		out.SynAckDroppedExternalService = *o.SynAckDroppedExternalService
	}
	if o.SynAckInvalidFormat != nil {
		out.SynAckInvalidFormat = *o.SynAckInvalidFormat
	}
	if o.SynAckMissingClaims != nil {
		out.SynAckMissingClaims = *o.SynAckMissingClaims
	}
	if o.SynAckMissingToken != nil {
		out.SynAckMissingToken = *o.SynAckMissingToken
	}
	if o.SynAckNoTCPAuthOption != nil {
		out.SynAckNoTCPAuthOption = *o.SynAckNoTCPAuthOption
	}
	if o.SynAckRejected != nil {
		out.SynAckRejected = *o.SynAckRejected
	}
	if o.SynDroppedInvalidFormat != nil {
		out.SynDroppedInvalidFormat = *o.SynDroppedInvalidFormat
	}
	if o.SynDroppedInvalidToken != nil {
		out.SynDroppedInvalidToken = *o.SynDroppedInvalidToken
	}
	if o.SynDroppedNoClaims != nil {
		out.SynDroppedNoClaims = *o.SynDroppedNoClaims
	}
	if o.SynDroppedTCPOption != nil {
		out.SynDroppedTCPOption = *o.SynDroppedTCPOption
	}
	if o.SynRejectPacket != nil {
		out.SynRejectPacket = *o.SynRejectPacket
	}
	if o.SynUnexpectedPacket != nil {
		out.SynUnexpectedPacket = *o.SynUnexpectedPacket
	}
	if o.TCPAuthNotFound != nil {
		out.TCPAuthNotFound = *o.TCPAuthNotFound
	}
	if o.UDPAckInvalidSignature != nil {
		out.UDPAckInvalidSignature = *o.UDPAckInvalidSignature
	}
	if o.UDPConnectionsProcessed != nil {
		out.UDPConnectionsProcessed = *o.UDPConnectionsProcessed
	}
	if o.UDPDropContextNotFound != nil {
		out.UDPDropContextNotFound = *o.UDPDropContextNotFound
	}
	if o.UDPDropFin != nil {
		out.UDPDropFin = *o.UDPDropFin
	}
	if o.UDPDropInNfQueue != nil {
		out.UDPDropInNfQueue = *o.UDPDropInNfQueue
	}
	if o.UDPDropNoConnection != nil {
		out.UDPDropNoConnection = *o.UDPDropNoConnection
	}
	if o.UDPDropPacket != nil {
		out.UDPDropPacket = *o.UDPDropPacket
	}
	if o.UDPDropQueueFull != nil {
		out.UDPDropQueueFull = *o.UDPDropQueueFull
	}
	if o.UDPDropSynAck != nil {
		out.UDPDropSynAck = *o.UDPDropSynAck
	}
	if o.UDPInvalidNetState != nil {
		out.UDPInvalidNetState = *o.UDPInvalidNetState
	}
	if o.UDPPostProcessingFailed != nil {
		out.UDPPostProcessingFailed = *o.UDPPostProcessingFailed
	}
	if o.UDPPreProcessingFailed != nil {
		out.UDPPreProcessingFailed = *o.UDPPreProcessingFailed
	}
	if o.UDPRejected != nil {
		out.UDPRejected = *o.UDPRejected
	}
	if o.UDPSynAckDropBadClaims != nil {
		out.UDPSynAckDropBadClaims = *o.UDPSynAckDropBadClaims
	}
	if o.UDPSynAckMissingClaims != nil {
		out.UDPSynAckMissingClaims = *o.UDPSynAckMissingClaims
	}
	if o.UDPSynAckPolicy != nil {
		out.UDPSynAckPolicy = *o.UDPSynAckPolicy
	}
	if o.UDPSynDrop != nil {
		out.UDPSynDrop = *o.UDPSynDrop
	}
	if o.UDPSynDropPolicy != nil {
		out.UDPSynDropPolicy = *o.UDPSynDropPolicy
	}
	if o.UDPSynInvalidToken != nil {
		out.UDPSynInvalidToken = *o.UDPSynInvalidToken
	}
	if o.UDPSynMissingClaims != nil {
		out.UDPSynMissingClaims = *o.UDPSynMissingClaims
	}
	if o.UnknownError != nil {
		out.UnknownError = *o.UnknownError
	}
	if o.ConnectionsAnalyzed != nil {
		out.ConnectionsAnalyzed = *o.ConnectionsAnalyzed
	}
	if o.ConnectionsDropped != nil {
		out.ConnectionsDropped = *o.ConnectionsDropped
	}
	if o.ConnectionsExpired != nil {
		out.ConnectionsExpired = *o.ConnectionsExpired
	}
	if o.DroppedPackets != nil {
		out.DroppedPackets = *o.DroppedPackets
	}
	if o.EncryptionFailures != nil {
		out.EncryptionFailures = *o.EncryptionFailures
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.ExternalNetworkConnections != nil {
		out.ExternalNetworkConnections = *o.ExternalNetworkConnections
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PolicyDrops != nil {
		out.PolicyDrops = *o.PolicyDrops
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
	if o.TokenDrops != nil {
		out.TokenDrops = *o.TokenDrops
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
func (o *SparseCounterReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCounterReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCounterReport) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCounterReport) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCounterReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCounterReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCounterReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCounterReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCounterReport.
func (o *SparseCounterReport) DeepCopy() *SparseCounterReport {

	if o == nil {
		return nil
	}

	out := &SparseCounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCounterReport.
func (o *SparseCounterReport) DeepCopyInto(out *SparseCounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCounterReport: %s", err))
	}

	*out = *target.(*SparseCounterReport)
}

type mongoAttributesCounterReport struct {
	AckInUnknownState            int               `bson:"ackinunknownstate,omitempty"`
	AckInvalidFormat             int               `bson:"ackinvalidformat,omitempty"`
	AckRejected                  int               `bson:"ackrejected,omitempty"`
	AckSigValidationFailed       int               `bson:"acksigvalidationfailed,omitempty"`
	AckTCPNoTCPAuthOption        int               `bson:"acktcpnotcpauthoption,omitempty"`
	ConnectionsProcessed         int               `bson:"connectionsprocessed,omitempty"`
	ContextIDNotFound            int               `bson:"contextidnotfound,omitempty"`
	DroppedExternalService       int               `bson:"droppedexternalservice,omitempty"`
	ID                           bson.ObjectId     `bson:"_id,omitempty"`
	InvalidConnState             int               `bson:"invalidconnstate,omitempty"`
	InvalidNetState              int               `bson:"invalidnetstate,omitempty"`
	InvalidProtocol              int               `bson:"invalidprotocol,omitempty"`
	InvalidSynAck                int               `bson:"invalidsynack,omitempty"`
	MarkNotFound                 int               `bson:"marknotfound,omitempty"`
	NetSynNotSeen                int               `bson:"netsynnotseen,omitempty"`
	NoConnFound                  int               `bson:"noconnfound,omitempty"`
	NonPUTraffic                 int               `bson:"nonputraffic,omitempty"`
	OutOfOrderSynAck             int               `bson:"outofordersynack,omitempty"`
	PortNotFound                 int               `bson:"portnotfound,omitempty"`
	RejectPacket                 int               `bson:"rejectpacket,omitempty"`
	ServicePostprocessorFailed   int               `bson:"servicepostprocessorfailed,omitempty"`
	ServicePreprocessorFailed    int               `bson:"servicepreprocessorfailed,omitempty"`
	SynAckBadClaims              int               `bson:"synackbadclaims,omitempty"`
	SynAckClaimsMisMatch         int               `bson:"synackclaimsmismatch,omitempty"`
	SynAckDroppedExternalService int               `bson:"synackdroppedexternalservice,omitempty"`
	SynAckInvalidFormat          int               `bson:"synackinvalidformat,omitempty"`
	SynAckMissingClaims          int               `bson:"synackmissingclaims,omitempty"`
	SynAckMissingToken           int               `bson:"synackmissingtoken,omitempty"`
	SynAckNoTCPAuthOption        int               `bson:"synacknotcpauthoption,omitempty"`
	SynAckRejected               int               `bson:"synackrejected,omitempty"`
	SynDroppedInvalidFormat      int               `bson:"syndroppedinvalidformat,omitempty"`
	SynDroppedInvalidToken       int               `bson:"syndroppedinvalidtoken,omitempty"`
	SynDroppedNoClaims           int               `bson:"syndroppednoclaims,omitempty"`
	SynDroppedTCPOption          int               `bson:"syndroppedtcpoption,omitempty"`
	SynRejectPacket              int               `bson:"synrejectpacket,omitempty"`
	SynUnexpectedPacket          int               `bson:"synunexpectedpacket,omitempty"`
	TCPAuthNotFound              int               `bson:"tcpauthnotfound,omitempty"`
	UDPAckInvalidSignature       int               `bson:"udpackinvalidsignature,omitempty"`
	UDPConnectionsProcessed      int               `bson:"udpconnectionsprocessed,omitempty"`
	UDPDropContextNotFound       int               `bson:"udpdropcontextnotfound,omitempty"`
	UDPDropFin                   int               `bson:"udpdropfin,omitempty"`
	UDPDropInNfQueue             int               `bson:"udpdropinnfqueue,omitempty"`
	UDPDropNoConnection          int               `bson:"udpdropnoconnection,omitempty"`
	UDPDropPacket                int               `bson:"udpdroppacket,omitempty"`
	UDPDropQueueFull             int               `bson:"udpdropqueuefull,omitempty"`
	UDPDropSynAck                int               `bson:"udpdropsynack,omitempty"`
	UDPInvalidNetState           int               `bson:"udpinvalidnetstate,omitempty"`
	UDPPostProcessingFailed      int               `bson:"udppostprocessingfailed,omitempty"`
	UDPPreProcessingFailed       int               `bson:"udppreprocessingfailed,omitempty"`
	UDPRejected                  int               `bson:"udprejected,omitempty"`
	UDPSynAckDropBadClaims       int               `bson:"udpsynackdropbadclaims,omitempty"`
	UDPSynAckMissingClaims       int               `bson:"udpsynackmissingclaims,omitempty"`
	UDPSynAckPolicy              int               `bson:"udpsynackpolicy,omitempty"`
	UDPSynDrop                   int               `bson:"udpsyndrop,omitempty"`
	UDPSynDropPolicy             int               `bson:"udpsyndroppolicy,omitempty"`
	UDPSynInvalidToken           int               `bson:"udpsyninvalidtoken,omitempty"`
	UDPSynMissingClaims          int               `bson:"udpsynmissingclaims,omitempty"`
	UnknownError                 int               `bson:"unknownerror,omitempty"`
	ConnectionsAnalyzed          int               `bson:"connectionsanalyzed,omitempty"`
	ConnectionsDropped           int               `bson:"connectionsdropped,omitempty"`
	ConnectionsExpired           int               `bson:"connectionsexpired,omitempty"`
	DroppedPackets               int               `bson:"droppedpackets,omitempty"`
	EncryptionFailures           int               `bson:"encryptionfailures,omitempty"`
	EnforcerID                   string            `bson:"enforcerid,omitempty"`
	EnforcerNamespace            string            `bson:"enforcernamespace,omitempty"`
	ExternalNetworkConnections   int               `bson:"externalnetworkconnections,omitempty"`
	MigrationsLog                map[string]string `bson:"migrationslog,omitempty"`
	Namespace                    string            `bson:"namespace,omitempty"`
	PolicyDrops                  int               `bson:"policydrops,omitempty"`
	ProcessingUnitID             string            `bson:"processingunitid,omitempty"`
	ProcessingUnitNamespace      string            `bson:"processingunitnamespace,omitempty"`
	Timestamp                    time.Time         `bson:"timestamp,omitempty"`
	TokenDrops                   int               `bson:"tokendrops,omitempty"`
	ZHash                        int               `bson:"zhash"`
	Zone                         int               `bson:"zone"`
}
type mongoAttributesSparseCounterReport struct {
	AckInUnknownState            *int               `bson:"ackinunknownstate,omitempty"`
	AckInvalidFormat             *int               `bson:"ackinvalidformat,omitempty"`
	AckRejected                  *int               `bson:"ackrejected,omitempty"`
	AckSigValidationFailed       *int               `bson:"acksigvalidationfailed,omitempty"`
	AckTCPNoTCPAuthOption        *int               `bson:"acktcpnotcpauthoption,omitempty"`
	ConnectionsProcessed         *int               `bson:"connectionsprocessed,omitempty"`
	ContextIDNotFound            *int               `bson:"contextidnotfound,omitempty"`
	DroppedExternalService       *int               `bson:"droppedexternalservice,omitempty"`
	ID                           bson.ObjectId      `bson:"_id,omitempty"`
	InvalidConnState             *int               `bson:"invalidconnstate,omitempty"`
	InvalidNetState              *int               `bson:"invalidnetstate,omitempty"`
	InvalidProtocol              *int               `bson:"invalidprotocol,omitempty"`
	InvalidSynAck                *int               `bson:"invalidsynack,omitempty"`
	MarkNotFound                 *int               `bson:"marknotfound,omitempty"`
	NetSynNotSeen                *int               `bson:"netsynnotseen,omitempty"`
	NoConnFound                  *int               `bson:"noconnfound,omitempty"`
	NonPUTraffic                 *int               `bson:"nonputraffic,omitempty"`
	OutOfOrderSynAck             *int               `bson:"outofordersynack,omitempty"`
	PortNotFound                 *int               `bson:"portnotfound,omitempty"`
	RejectPacket                 *int               `bson:"rejectpacket,omitempty"`
	ServicePostprocessorFailed   *int               `bson:"servicepostprocessorfailed,omitempty"`
	ServicePreprocessorFailed    *int               `bson:"servicepreprocessorfailed,omitempty"`
	SynAckBadClaims              *int               `bson:"synackbadclaims,omitempty"`
	SynAckClaimsMisMatch         *int               `bson:"synackclaimsmismatch,omitempty"`
	SynAckDroppedExternalService *int               `bson:"synackdroppedexternalservice,omitempty"`
	SynAckInvalidFormat          *int               `bson:"synackinvalidformat,omitempty"`
	SynAckMissingClaims          *int               `bson:"synackmissingclaims,omitempty"`
	SynAckMissingToken           *int               `bson:"synackmissingtoken,omitempty"`
	SynAckNoTCPAuthOption        *int               `bson:"synacknotcpauthoption,omitempty"`
	SynAckRejected               *int               `bson:"synackrejected,omitempty"`
	SynDroppedInvalidFormat      *int               `bson:"syndroppedinvalidformat,omitempty"`
	SynDroppedInvalidToken       *int               `bson:"syndroppedinvalidtoken,omitempty"`
	SynDroppedNoClaims           *int               `bson:"syndroppednoclaims,omitempty"`
	SynDroppedTCPOption          *int               `bson:"syndroppedtcpoption,omitempty"`
	SynRejectPacket              *int               `bson:"synrejectpacket,omitempty"`
	SynUnexpectedPacket          *int               `bson:"synunexpectedpacket,omitempty"`
	TCPAuthNotFound              *int               `bson:"tcpauthnotfound,omitempty"`
	UDPAckInvalidSignature       *int               `bson:"udpackinvalidsignature,omitempty"`
	UDPConnectionsProcessed      *int               `bson:"udpconnectionsprocessed,omitempty"`
	UDPDropContextNotFound       *int               `bson:"udpdropcontextnotfound,omitempty"`
	UDPDropFin                   *int               `bson:"udpdropfin,omitempty"`
	UDPDropInNfQueue             *int               `bson:"udpdropinnfqueue,omitempty"`
	UDPDropNoConnection          *int               `bson:"udpdropnoconnection,omitempty"`
	UDPDropPacket                *int               `bson:"udpdroppacket,omitempty"`
	UDPDropQueueFull             *int               `bson:"udpdropqueuefull,omitempty"`
	UDPDropSynAck                *int               `bson:"udpdropsynack,omitempty"`
	UDPInvalidNetState           *int               `bson:"udpinvalidnetstate,omitempty"`
	UDPPostProcessingFailed      *int               `bson:"udppostprocessingfailed,omitempty"`
	UDPPreProcessingFailed       *int               `bson:"udppreprocessingfailed,omitempty"`
	UDPRejected                  *int               `bson:"udprejected,omitempty"`
	UDPSynAckDropBadClaims       *int               `bson:"udpsynackdropbadclaims,omitempty"`
	UDPSynAckMissingClaims       *int               `bson:"udpsynackmissingclaims,omitempty"`
	UDPSynAckPolicy              *int               `bson:"udpsynackpolicy,omitempty"`
	UDPSynDrop                   *int               `bson:"udpsyndrop,omitempty"`
	UDPSynDropPolicy             *int               `bson:"udpsyndroppolicy,omitempty"`
	UDPSynInvalidToken           *int               `bson:"udpsyninvalidtoken,omitempty"`
	UDPSynMissingClaims          *int               `bson:"udpsynmissingclaims,omitempty"`
	UnknownError                 *int               `bson:"unknownerror,omitempty"`
	ConnectionsAnalyzed          *int               `bson:"connectionsanalyzed,omitempty"`
	ConnectionsDropped           *int               `bson:"connectionsdropped,omitempty"`
	ConnectionsExpired           *int               `bson:"connectionsexpired,omitempty"`
	DroppedPackets               *int               `bson:"droppedpackets,omitempty"`
	EncryptionFailures           *int               `bson:"encryptionfailures,omitempty"`
	EnforcerID                   *string            `bson:"enforcerid,omitempty"`
	EnforcerNamespace            *string            `bson:"enforcernamespace,omitempty"`
	ExternalNetworkConnections   *int               `bson:"externalnetworkconnections,omitempty"`
	MigrationsLog                *map[string]string `bson:"migrationslog,omitempty"`
	Namespace                    *string            `bson:"namespace,omitempty"`
	PolicyDrops                  *int               `bson:"policydrops,omitempty"`
	ProcessingUnitID             *string            `bson:"processingunitid,omitempty"`
	ProcessingUnitNamespace      *string            `bson:"processingunitnamespace,omitempty"`
	Timestamp                    *time.Time         `bson:"timestamp,omitempty"`
	TokenDrops                   *int               `bson:"tokendrops,omitempty"`
	ZHash                        *int               `bson:"zhash,omitempty"`
	Zone                         *int               `bson:"zone,omitempty"`
}
