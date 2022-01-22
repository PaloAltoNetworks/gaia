package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TraceRecord represents the model of a tracerecord
type TraceRecord struct {
	// The time to live (TTL) value of the packet.
	TTL int `json:"TTL" msgpack:"TTL" bson:"ttl" mapstructure:"TTL,omitempty"`

	// Chain that the trace was collected from.
	Chain string `json:"chain" msgpack:"chain" bson:"chain" mapstructure:"chain,omitempty"`

	// The destination IP.
	DestinationIP string `json:"destinationIP" msgpack:"destinationIP" bson:"destinationip" mapstructure:"destinationIP,omitempty"`

	// The destination interface of the packet.
	DestinationInterface string `json:"destinationInterface" msgpack:"destinationInterface" bson:"destinationinterface" mapstructure:"destinationInterface,omitempty"`

	// The destination UPD or TCP port of the packet.
	DestinationPort int `json:"destinationPort" msgpack:"destinationPort" bson:"destinationport" mapstructure:"destinationPort,omitempty"`

	// Length of the observed packet.
	Length int `json:"length" msgpack:"length" bson:"length" mapstructure:"length,omitempty"`

	// The IP packet header ID.
	PacketID int `json:"packetID" msgpack:"packetID" bson:"packetid" mapstructure:"packetID,omitempty"`

	// The protocol of the packet.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Priority index of the iptables entry that was hit.
	RuleID int `json:"ruleID" msgpack:"ruleID" bson:"ruleid" mapstructure:"ruleID,omitempty"`

	// Source IP of the packet.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"sourceip" mapstructure:"sourceIP,omitempty"`

	// Source interface of the packet.
	SourceInterface string `json:"sourceInterface" msgpack:"sourceInterface" bson:"sourceinterface" mapstructure:"sourceInterface,omitempty"`

	// Source TCP or UDP port of the packet.
	SourcePort int `json:"sourcePort" msgpack:"sourcePort" bson:"sourceport" mapstructure:"sourcePort,omitempty"`

	// The iptables name that the trace collected.
	TableName string `json:"tableName" msgpack:"tableName" bson:"tablename" mapstructure:"tableName,omitempty"`

	// The time-date stamp of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTraceRecord returns a new *TraceRecord
func NewTraceRecord() *TraceRecord {

	return &TraceRecord{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TraceRecord) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTraceRecord{}

	s.TTL = o.TTL
	s.Chain = o.Chain
	s.DestinationIP = o.DestinationIP
	s.DestinationInterface = o.DestinationInterface
	s.DestinationPort = o.DestinationPort
	s.Length = o.Length
	s.PacketID = o.PacketID
	s.Protocol = o.Protocol
	s.RuleID = o.RuleID
	s.SourceIP = o.SourceIP
	s.SourceInterface = o.SourceInterface
	s.SourcePort = o.SourcePort
	s.TableName = o.TableName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TraceRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTraceRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.TTL = s.TTL
	o.Chain = s.Chain
	o.DestinationIP = s.DestinationIP
	o.DestinationInterface = s.DestinationInterface
	o.DestinationPort = s.DestinationPort
	o.Length = s.Length
	o.PacketID = s.PacketID
	o.Protocol = s.Protocol
	o.RuleID = s.RuleID
	o.SourceIP = s.SourceIP
	o.SourceInterface = s.SourceInterface
	o.SourcePort = s.SourcePort
	o.TableName = s.TableName

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *TraceRecord) BleveType() string {

	return "tracerecord"
}

// DeepCopy returns a deep copy if the TraceRecord.
func (o *TraceRecord) DeepCopy() *TraceRecord {

	if o == nil {
		return nil
	}

	out := &TraceRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TraceRecord.
func (o *TraceRecord) DeepCopyInto(out *TraceRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TraceRecord: %s", err))
	}

	*out = *target.(*TraceRecord)
}

// Validate valides the current information stored into the structure.
func (o *TraceRecord) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredInt("TTL", o.TTL); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("TTL", o.TTL, int(255), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("chain", o.Chain); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("destinationIP", o.DestinationIP); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("destinationPort", o.DestinationPort); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("destinationPort", o.DestinationPort, int(65536), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMinimumInt("destinationPort", o.DestinationPort, int(1), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("length", o.Length); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("length", o.Length, int(65536), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("packetID", o.PacketID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("protocol", o.Protocol, int(65536), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("ruleID", o.RuleID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceIP", o.SourceIP); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("sourcePort", o.SourcePort); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("sourcePort", o.SourcePort, int(65536), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMinimumInt("sourcePort", o.SourcePort, int(1), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("tableName", o.TableName); err != nil {
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

type mongoAttributesTraceRecord struct {
	TTL                  int    `bson:"ttl"`
	Chain                string `bson:"chain"`
	DestinationIP        string `bson:"destinationip"`
	DestinationInterface string `bson:"destinationinterface"`
	DestinationPort      int    `bson:"destinationport"`
	Length               int    `bson:"length"`
	PacketID             int    `bson:"packetid"`
	Protocol             int    `bson:"protocol"`
	RuleID               int    `bson:"ruleid"`
	SourceIP             string `bson:"sourceip"`
	SourceInterface      string `bson:"sourceinterface"`
	SourcePort           int    `bson:"sourceport"`
	TableName            string `bson:"tablename"`
}
