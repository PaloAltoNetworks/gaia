package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphEdgeDestinationTypeValue represents the possible values for attribute "destinationType".
type GraphEdgeDestinationTypeValue string

const (
	// GraphEdgeDestinationTypeExternalNetwork represents the value ExternalNetwork.
	GraphEdgeDestinationTypeExternalNetwork GraphEdgeDestinationTypeValue = "ExternalNetwork"

	// GraphEdgeDestinationTypeNamespace represents the value Namespace.
	GraphEdgeDestinationTypeNamespace GraphEdgeDestinationTypeValue = "Namespace"

	// GraphEdgeDestinationTypeNode represents the value Node.
	GraphEdgeDestinationTypeNode GraphEdgeDestinationTypeValue = "Node"

	// GraphEdgeDestinationTypeProcessingUnit represents the value ProcessingUnit.
	GraphEdgeDestinationTypeProcessingUnit GraphEdgeDestinationTypeValue = "ProcessingUnit"

	// GraphEdgeDestinationTypeRemoteController represents the value RemoteController.
	GraphEdgeDestinationTypeRemoteController GraphEdgeDestinationTypeValue = "RemoteController"
)

// GraphEdgeSourceTypeValue represents the possible values for attribute "sourceType".
type GraphEdgeSourceTypeValue string

const (
	// GraphEdgeSourceTypeExternalNetwork represents the value ExternalNetwork.
	GraphEdgeSourceTypeExternalNetwork GraphEdgeSourceTypeValue = "ExternalNetwork"

	// GraphEdgeSourceTypeNamespace represents the value Namespace.
	GraphEdgeSourceTypeNamespace GraphEdgeSourceTypeValue = "Namespace"

	// GraphEdgeSourceTypeNode represents the value Node.
	GraphEdgeSourceTypeNode GraphEdgeSourceTypeValue = "Node"

	// GraphEdgeSourceTypeProcessingUnit represents the value ProcessingUnit.
	GraphEdgeSourceTypeProcessingUnit GraphEdgeSourceTypeValue = "ProcessingUnit"

	// GraphEdgeSourceTypeRemoteController represents the value RemoteController.
	GraphEdgeSourceTypeRemoteController GraphEdgeSourceTypeValue = "RemoteController"
)

// GraphEdgeIdentity represents the Identity of the object.
var GraphEdgeIdentity = elemental.Identity{
	Name:     "graphedge",
	Category: "graphedges",
	Package:  "meteor",
	Private:  true,
}

// GraphEdgesList represents a list of GraphEdges
type GraphEdgesList []*GraphEdge

// Identity returns the identity of the objects in the list.
func (o GraphEdgesList) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Copy returns a pointer to a copy the GraphEdgesList.
func (o GraphEdgesList) Copy() elemental.Identifiables {

	copy := append(GraphEdgesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the GraphEdgesList.
func (o GraphEdgesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(GraphEdgesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*GraphEdge))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o GraphEdgesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o GraphEdgesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the GraphEdgesList converted to SparseGraphEdgesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o GraphEdgesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseGraphEdgesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseGraphEdge)
	}

	return out
}

// Version returns the version of the content.
func (o GraphEdgesList) Version() int {

	return 1
}

// GraphEdge represents the model of a graphedge
type GraphEdge struct {
	// DB Identifier of the edge.
	ID string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// Number of accepted flows in the edge.
	AcceptedFlows bool `json:"acceptedFlows" msgpack:"acceptedFlows" bson:"acceptedflows" mapstructure:"acceptedFlows,omitempty"`

	// The date for the day bucket.
	BucketDay time.Time `json:"-" msgpack:"-" bson:"bucketday" mapstructure:"-,omitempty"`

	// The date for the hour bucket.
	BucketHour time.Time `json:"-" msgpack:"-" bson:"buckethour" mapstructure:"-,omitempty"`

	// The date for the minute bucket.
	BucketMinute time.Time `json:"-" msgpack:"-" bson:"bucketminute" mapstructure:"-,omitempty"`

	// The date for the month bucket.
	BucketMonth time.Time `json:"-" msgpack:"-" bson:"bucketmonth" mapstructure:"-,omitempty"`

	// The truth value that indicates if there are accepted default flows.
	DefaultAcceptedFlows bool `json:"defaultAcceptedFlows" msgpack:"defaultAcceptedFlows" bson:"defaultacceptedflows" mapstructure:"defaultAcceptedFlows,omitempty"`

	// The truth value that indicates if there are rejected default flows.
	DefaultRejectedFlows bool `json:"defaultRejectedFlows" msgpack:"defaultRejectedFlows" bson:"defaultrejectedflows" mapstructure:"defaultRejectedFlows,omitempty"`

	// Name of the remote destination controller if different than the current one.
	DestinationController string `json:"destinationController,omitempty" msgpack:"destinationController,omitempty" bson:"destinationcontroller,omitempty" mapstructure:"destinationController,omitempty"`

	// ID of the destination `GraphNode` of the edge.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// Type of the destination `GraphNode` of the edge.
	DestinationType GraphEdgeDestinationTypeValue `json:"destinationType" msgpack:"destinationType" bson:"destinationtype" mapstructure:"destinationType,omitempty"`

	// The number of encrypted flows in the edge.
	Encrypted bool `json:"encrypted" msgpack:"encrypted" bson:"encrypted" mapstructure:"encrypted,omitempty"`

	// Contains the date when the edge was first seen.
	FirstSeen time.Time `json:"firstSeen,omitempty" msgpack:"firstSeen,omitempty" bson:"firstseen,omitempty" mapstructure:"firstSeen,omitempty"`

	// Identifier of the edge.
	FlowID string `json:"ID" msgpack:"ID" bson:"flowid" mapstructure:"ID,omitempty"`

	// Contains the date when the edge was last seen.
	LastSeen time.Time `json:"lastSeen,omitempty" msgpack:"lastSeen,omitempty" bson:"lastseen,omitempty" mapstructure:"lastSeen,omitempty"`

	// Namespace of the object that reported the flow.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Number of accepted observed flows.
	ObservedAcceptedFlows bool `json:"observedAcceptedFlows" msgpack:"observedAcceptedFlows" bson:"observedacceptedflows" mapstructure:"observedAcceptedFlows,omitempty"`

	// Number of encrypted observed flows.
	ObservedEncrypted bool `json:"observedEncrypted" msgpack:"observedEncrypted" bson:"observedencrypted" mapstructure:"observedEncrypted,omitempty"`

	// Number of rejected observed flows.
	ObservedRejectedFlows bool `json:"observedRejectedFlows" msgpack:"observedRejectedFlows" bson:"observedrejectedflows" mapstructure:"observedRejectedFlows,omitempty"`

	// the destination protocol for the flow. If the protocol is TCP or UDP, it will also include the port.
	ProtoPort string `json:"protoPort" msgpack:"protoPort" bson:"protoport" mapstructure:"protoPort,omitempty"`

	// Number of rejected flows in the edge.
	RejectedFlows bool `json:"rejectedFlows" msgpack:"rejectedFlows" bson:"rejectedflows" mapstructure:"rejectedFlows,omitempty"`

	// Namespace of the object that was targeted by the flow.
	RemoteNamespace string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"remotenamespace,omitempty" mapstructure:"remoteNamespace,omitempty"`

	// Name of the remote source controller if different than the current one.
	SourceController string `json:"sourceController,omitempty" msgpack:"sourceController,omitempty" bson:"sourcecontroller,omitempty" mapstructure:"sourceController,omitempty"`

	// ID of the source `GraphNode` of the edge.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// Type of the source `GraphNode` of the edge.
	SourceType GraphEdgeSourceTypeValue `json:"sourceType" msgpack:"sourceType" bson:"sourcetype" mapstructure:"sourceType,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGraphEdge returns a new *GraphEdge
func NewGraphEdge() *GraphEdge {

	return &GraphEdge{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *GraphEdge) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GraphEdge) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GraphEdge) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GraphEdge) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesGraphEdge{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.AcceptedFlows = o.AcceptedFlows
	s.BucketDay = o.BucketDay
	s.BucketHour = o.BucketHour
	s.BucketMinute = o.BucketMinute
	s.BucketMonth = o.BucketMonth
	s.DefaultAcceptedFlows = o.DefaultAcceptedFlows
	s.DefaultRejectedFlows = o.DefaultRejectedFlows
	s.DestinationController = o.DestinationController
	s.DestinationID = o.DestinationID
	s.DestinationType = o.DestinationType
	s.Encrypted = o.Encrypted
	s.FirstSeen = o.FirstSeen
	s.FlowID = o.FlowID
	s.LastSeen = o.LastSeen
	s.Namespace = o.Namespace
	s.ObservedAcceptedFlows = o.ObservedAcceptedFlows
	s.ObservedEncrypted = o.ObservedEncrypted
	s.ObservedRejectedFlows = o.ObservedRejectedFlows
	s.ProtoPort = o.ProtoPort
	s.RejectedFlows = o.RejectedFlows
	s.RemoteNamespace = o.RemoteNamespace
	s.SourceController = o.SourceController
	s.SourceID = o.SourceID
	s.SourceType = o.SourceType
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GraphEdge) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesGraphEdge{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.AcceptedFlows = s.AcceptedFlows
	o.BucketDay = s.BucketDay
	o.BucketHour = s.BucketHour
	o.BucketMinute = s.BucketMinute
	o.BucketMonth = s.BucketMonth
	o.DefaultAcceptedFlows = s.DefaultAcceptedFlows
	o.DefaultRejectedFlows = s.DefaultRejectedFlows
	o.DestinationController = s.DestinationController
	o.DestinationID = s.DestinationID
	o.DestinationType = s.DestinationType
	o.Encrypted = s.Encrypted
	o.FirstSeen = s.FirstSeen
	o.FlowID = s.FlowID
	o.LastSeen = s.LastSeen
	o.Namespace = s.Namespace
	o.ObservedAcceptedFlows = s.ObservedAcceptedFlows
	o.ObservedEncrypted = s.ObservedEncrypted
	o.ObservedRejectedFlows = s.ObservedRejectedFlows
	o.ProtoPort = s.ProtoPort
	o.RejectedFlows = s.RejectedFlows
	o.RemoteNamespace = s.RemoteNamespace
	o.SourceController = s.SourceController
	o.SourceID = s.SourceID
	o.SourceType = s.SourceType
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *GraphEdge) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *GraphEdge) BleveType() string {

	return "graphedge"
}

// DefaultOrder returns the list of default ordering fields.
func (o *GraphEdge) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *GraphEdge) Doc() string {

	return `Represents an edge from the dependency map.`
}

func (o *GraphEdge) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetZHash returns the ZHash of the receiver.
func (o *GraphEdge) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *GraphEdge) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *GraphEdge) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *GraphEdge) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *GraphEdge) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseGraphEdge{
			ID:                    &o.ID,
			AcceptedFlows:         &o.AcceptedFlows,
			BucketDay:             &o.BucketDay,
			BucketHour:            &o.BucketHour,
			BucketMinute:          &o.BucketMinute,
			BucketMonth:           &o.BucketMonth,
			DefaultAcceptedFlows:  &o.DefaultAcceptedFlows,
			DefaultRejectedFlows:  &o.DefaultRejectedFlows,
			DestinationController: &o.DestinationController,
			DestinationID:         &o.DestinationID,
			DestinationType:       &o.DestinationType,
			Encrypted:             &o.Encrypted,
			FirstSeen:             &o.FirstSeen,
			FlowID:                &o.FlowID,
			LastSeen:              &o.LastSeen,
			Namespace:             &o.Namespace,
			ObservedAcceptedFlows: &o.ObservedAcceptedFlows,
			ObservedEncrypted:     &o.ObservedEncrypted,
			ObservedRejectedFlows: &o.ObservedRejectedFlows,
			ProtoPort:             &o.ProtoPort,
			RejectedFlows:         &o.RejectedFlows,
			RemoteNamespace:       &o.RemoteNamespace,
			SourceController:      &o.SourceController,
			SourceID:              &o.SourceID,
			SourceType:            &o.SourceType,
			ZHash:                 &o.ZHash,
			Zone:                  &o.Zone,
		}
	}

	sp := &SparseGraphEdge{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "acceptedFlows":
			sp.AcceptedFlows = &(o.AcceptedFlows)
		case "bucketDay":
			sp.BucketDay = &(o.BucketDay)
		case "bucketHour":
			sp.BucketHour = &(o.BucketHour)
		case "bucketMinute":
			sp.BucketMinute = &(o.BucketMinute)
		case "bucketMonth":
			sp.BucketMonth = &(o.BucketMonth)
		case "defaultAcceptedFlows":
			sp.DefaultAcceptedFlows = &(o.DefaultAcceptedFlows)
		case "defaultRejectedFlows":
			sp.DefaultRejectedFlows = &(o.DefaultRejectedFlows)
		case "destinationController":
			sp.DestinationController = &(o.DestinationController)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationType":
			sp.DestinationType = &(o.DestinationType)
		case "encrypted":
			sp.Encrypted = &(o.Encrypted)
		case "firstSeen":
			sp.FirstSeen = &(o.FirstSeen)
		case "flowID":
			sp.FlowID = &(o.FlowID)
		case "lastSeen":
			sp.LastSeen = &(o.LastSeen)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "observedAcceptedFlows":
			sp.ObservedAcceptedFlows = &(o.ObservedAcceptedFlows)
		case "observedEncrypted":
			sp.ObservedEncrypted = &(o.ObservedEncrypted)
		case "observedRejectedFlows":
			sp.ObservedRejectedFlows = &(o.ObservedRejectedFlows)
		case "protoPort":
			sp.ProtoPort = &(o.ProtoPort)
		case "rejectedFlows":
			sp.RejectedFlows = &(o.RejectedFlows)
		case "remoteNamespace":
			sp.RemoteNamespace = &(o.RemoteNamespace)
		case "sourceController":
			sp.SourceController = &(o.SourceController)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceType":
			sp.SourceType = &(o.SourceType)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseGraphEdge to the object.
func (o *GraphEdge) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseGraphEdge)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AcceptedFlows != nil {
		o.AcceptedFlows = *so.AcceptedFlows
	}
	if so.BucketDay != nil {
		o.BucketDay = *so.BucketDay
	}
	if so.BucketHour != nil {
		o.BucketHour = *so.BucketHour
	}
	if so.BucketMinute != nil {
		o.BucketMinute = *so.BucketMinute
	}
	if so.BucketMonth != nil {
		o.BucketMonth = *so.BucketMonth
	}
	if so.DefaultAcceptedFlows != nil {
		o.DefaultAcceptedFlows = *so.DefaultAcceptedFlows
	}
	if so.DefaultRejectedFlows != nil {
		o.DefaultRejectedFlows = *so.DefaultRejectedFlows
	}
	if so.DestinationController != nil {
		o.DestinationController = *so.DestinationController
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.DestinationType != nil {
		o.DestinationType = *so.DestinationType
	}
	if so.Encrypted != nil {
		o.Encrypted = *so.Encrypted
	}
	if so.FirstSeen != nil {
		o.FirstSeen = *so.FirstSeen
	}
	if so.FlowID != nil {
		o.FlowID = *so.FlowID
	}
	if so.LastSeen != nil {
		o.LastSeen = *so.LastSeen
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ObservedAcceptedFlows != nil {
		o.ObservedAcceptedFlows = *so.ObservedAcceptedFlows
	}
	if so.ObservedEncrypted != nil {
		o.ObservedEncrypted = *so.ObservedEncrypted
	}
	if so.ObservedRejectedFlows != nil {
		o.ObservedRejectedFlows = *so.ObservedRejectedFlows
	}
	if so.ProtoPort != nil {
		o.ProtoPort = *so.ProtoPort
	}
	if so.RejectedFlows != nil {
		o.RejectedFlows = *so.RejectedFlows
	}
	if so.RemoteNamespace != nil {
		o.RemoteNamespace = *so.RemoteNamespace
	}
	if so.SourceController != nil {
		o.SourceController = *so.SourceController
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceType != nil {
		o.SourceType = *so.SourceType
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the GraphEdge.
func (o *GraphEdge) DeepCopy() *GraphEdge {

	if o == nil {
		return nil
	}

	out := &GraphEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphEdge.
func (o *GraphEdge) DeepCopyInto(out *GraphEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphEdge: %s", err))
	}

	*out = *target.(*GraphEdge)
}

// Validate valides the current information stored into the structure.
func (o *GraphEdge) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("destinationType", string(o.DestinationType), []string{"ProcessingUnit", "ExternalNetwork", "Namespace", "Node", "RemoteController"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("sourceType", string(o.SourceType), []string{"ProcessingUnit", "ExternalNetwork", "Namespace", "Node", "RemoteController"}, false); err != nil {
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
func (*GraphEdge) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := GraphEdgeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return GraphEdgeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*GraphEdge) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return GraphEdgeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *GraphEdge) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "acceptedFlows":
		return o.AcceptedFlows
	case "bucketDay":
		return o.BucketDay
	case "bucketHour":
		return o.BucketHour
	case "bucketMinute":
		return o.BucketMinute
	case "bucketMonth":
		return o.BucketMonth
	case "defaultAcceptedFlows":
		return o.DefaultAcceptedFlows
	case "defaultRejectedFlows":
		return o.DefaultRejectedFlows
	case "destinationController":
		return o.DestinationController
	case "destinationID":
		return o.DestinationID
	case "destinationType":
		return o.DestinationType
	case "encrypted":
		return o.Encrypted
	case "firstSeen":
		return o.FirstSeen
	case "flowID":
		return o.FlowID
	case "lastSeen":
		return o.LastSeen
	case "namespace":
		return o.Namespace
	case "observedAcceptedFlows":
		return o.ObservedAcceptedFlows
	case "observedEncrypted":
		return o.ObservedEncrypted
	case "observedRejectedFlows":
		return o.ObservedRejectedFlows
	case "protoPort":
		return o.ProtoPort
	case "rejectedFlows":
		return o.RejectedFlows
	case "remoteNamespace":
		return o.RemoteNamespace
	case "sourceController":
		return o.SourceController
	case "sourceID":
		return o.SourceID
	case "sourceType":
		return o.SourceType
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// GraphEdgeAttributesMap represents the map of attribute for GraphEdge.
var GraphEdgeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `DB Identifier of the edge.`,
		Identifier:     true,
		Name:           "ID",
		Stored:         true,
		Type:           "string",
	},
	"AcceptedFlows": {
		AllowedChoices: []string{},
		BSONFieldName:  "acceptedflows",
		ConvertedName:  "AcceptedFlows",
		Description:    `Number of accepted flows in the edge.`,
		Exposed:        true,
		Name:           "acceptedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"BucketDay": {
		AllowedChoices: []string{},
		BSONFieldName:  "bucketday",
		ConvertedName:  "BucketDay",
		Description:    `The date for the day bucket.`,
		Name:           "bucketDay",
		Stored:         true,
		Type:           "time",
	},
	"BucketHour": {
		AllowedChoices: []string{},
		BSONFieldName:  "buckethour",
		ConvertedName:  "BucketHour",
		Description:    `The date for the hour bucket.`,
		Name:           "bucketHour",
		Stored:         true,
		Type:           "time",
	},
	"BucketMinute": {
		AllowedChoices: []string{},
		BSONFieldName:  "bucketminute",
		ConvertedName:  "BucketMinute",
		Description:    `The date for the minute bucket.`,
		Name:           "bucketMinute",
		Stored:         true,
		Type:           "time",
	},
	"BucketMonth": {
		AllowedChoices: []string{},
		BSONFieldName:  "bucketmonth",
		ConvertedName:  "BucketMonth",
		Description:    `The date for the month bucket.`,
		Name:           "bucketMonth",
		Stored:         true,
		Type:           "time",
	},
	"DefaultAcceptedFlows": {
		AllowedChoices: []string{},
		BSONFieldName:  "defaultacceptedflows",
		ConvertedName:  "DefaultAcceptedFlows",
		Description:    `The truth value that indicates if there are accepted default flows.`,
		Exposed:        true,
		Name:           "defaultAcceptedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"DefaultRejectedFlows": {
		AllowedChoices: []string{},
		BSONFieldName:  "defaultrejectedflows",
		ConvertedName:  "DefaultRejectedFlows",
		Description:    `The truth value that indicates if there are rejected default flows.`,
		Exposed:        true,
		Name:           "defaultRejectedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"DestinationController": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationcontroller",
		ConvertedName:  "DestinationController",
		Description:    `Name of the remote destination controller if different than the current one.`,
		Exposed:        true,
		Name:           "destinationController",
		Stored:         true,
		Type:           "string",
	},
	"DestinationID": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationid",
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
	},
	"DestinationType": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Namespace", "Node", "RemoteController"},
		BSONFieldName:  "destinationtype",
		ConvertedName:  "DestinationType",
		Description:    `Type of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationType",
		Stored:         true,
		Type:           "enum",
	},
	"Encrypted": {
		AllowedChoices: []string{},
		BSONFieldName:  "encrypted",
		ConvertedName:  "Encrypted",
		Description:    `The number of encrypted flows in the edge.`,
		Exposed:        true,
		Name:           "encrypted",
		Stored:         true,
		Type:           "boolean",
	},
	"FirstSeen": {
		AllowedChoices: []string{},
		BSONFieldName:  "firstseen",
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date when the edge was first seen.`,
		Exposed:        true,
		Name:           "firstSeen",
		Stored:         true,
		Type:           "time",
	},
	"FlowID": {
		AllowedChoices: []string{},
		BSONFieldName:  "flowid",
		ConvertedName:  "FlowID",
		Description:    `Identifier of the edge.`,
		Exposed:        true,
		Name:           "flowID",
		Stored:         true,
		Type:           "string",
	},
	"LastSeen": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastseen",
		ConvertedName:  "LastSeen",
		Description:    `Contains the date when the edge was last seen.`,
		Exposed:        true,
		Name:           "lastSeen",
		Stored:         true,
		Type:           "time",
	},
	"Namespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace of the object that reported the flow.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"ObservedAcceptedFlows": {
		AllowedChoices: []string{},
		BSONFieldName:  "observedacceptedflows",
		ConvertedName:  "ObservedAcceptedFlows",
		Description:    `Number of accepted observed flows.`,
		Exposed:        true,
		Name:           "observedAcceptedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"ObservedEncrypted": {
		AllowedChoices: []string{},
		BSONFieldName:  "observedencrypted",
		ConvertedName:  "ObservedEncrypted",
		Description:    `Number of encrypted observed flows.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Stored:         true,
		Type:           "boolean",
	},
	"ObservedRejectedFlows": {
		AllowedChoices: []string{},
		BSONFieldName:  "observedrejectedflows",
		ConvertedName:  "ObservedRejectedFlows",
		Description:    `Number of rejected observed flows.`,
		Exposed:        true,
		Name:           "observedRejectedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"ProtoPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "protoport",
		ConvertedName:  "ProtoPort",
		Description:    `the destination protocol for the flow. If the protocol is TCP or UDP, it will also include the port.`,
		Exposed:        true,
		Name:           "protoPort",
		Stored:         true,
		Type:           "string",
	},
	"RejectedFlows": {
		AllowedChoices: []string{},
		BSONFieldName:  "rejectedflows",
		ConvertedName:  "RejectedFlows",
		Description:    `Number of rejected flows in the edge.`,
		Exposed:        true,
		Name:           "rejectedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"RemoteNamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "remotenamespace",
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the object that was targeted by the flow.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Stored:         true,
		Type:           "string",
	},
	"SourceController": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcecontroller",
		ConvertedName:  "SourceController",
		Description:    `Name of the remote source controller if different than the current one.`,
		Exposed:        true,
		Name:           "sourceController",
		Stored:         true,
		Type:           "string",
	},
	"SourceID": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceid",
		ConvertedName:  "SourceID",
		Description:    `ID of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"SourceType": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Namespace", "Node", "RemoteController"},
		BSONFieldName:  "sourcetype",
		ConvertedName:  "SourceType",
		Description:    `Type of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceType",
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

// GraphEdgeLowerCaseAttributesMap represents the map of attribute for GraphEdge.
var GraphEdgeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `DB Identifier of the edge.`,
		Identifier:     true,
		Name:           "ID",
		Stored:         true,
		Type:           "string",
	},
	"acceptedflows": {
		AllowedChoices: []string{},
		BSONFieldName:  "acceptedflows",
		ConvertedName:  "AcceptedFlows",
		Description:    `Number of accepted flows in the edge.`,
		Exposed:        true,
		Name:           "acceptedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"bucketday": {
		AllowedChoices: []string{},
		BSONFieldName:  "bucketday",
		ConvertedName:  "BucketDay",
		Description:    `The date for the day bucket.`,
		Name:           "bucketDay",
		Stored:         true,
		Type:           "time",
	},
	"buckethour": {
		AllowedChoices: []string{},
		BSONFieldName:  "buckethour",
		ConvertedName:  "BucketHour",
		Description:    `The date for the hour bucket.`,
		Name:           "bucketHour",
		Stored:         true,
		Type:           "time",
	},
	"bucketminute": {
		AllowedChoices: []string{},
		BSONFieldName:  "bucketminute",
		ConvertedName:  "BucketMinute",
		Description:    `The date for the minute bucket.`,
		Name:           "bucketMinute",
		Stored:         true,
		Type:           "time",
	},
	"bucketmonth": {
		AllowedChoices: []string{},
		BSONFieldName:  "bucketmonth",
		ConvertedName:  "BucketMonth",
		Description:    `The date for the month bucket.`,
		Name:           "bucketMonth",
		Stored:         true,
		Type:           "time",
	},
	"defaultacceptedflows": {
		AllowedChoices: []string{},
		BSONFieldName:  "defaultacceptedflows",
		ConvertedName:  "DefaultAcceptedFlows",
		Description:    `The truth value that indicates if there are accepted default flows.`,
		Exposed:        true,
		Name:           "defaultAcceptedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"defaultrejectedflows": {
		AllowedChoices: []string{},
		BSONFieldName:  "defaultrejectedflows",
		ConvertedName:  "DefaultRejectedFlows",
		Description:    `The truth value that indicates if there are rejected default flows.`,
		Exposed:        true,
		Name:           "defaultRejectedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"destinationcontroller": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationcontroller",
		ConvertedName:  "DestinationController",
		Description:    `Name of the remote destination controller if different than the current one.`,
		Exposed:        true,
		Name:           "destinationController",
		Stored:         true,
		Type:           "string",
	},
	"destinationid": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationid",
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
	},
	"destinationtype": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Namespace", "Node", "RemoteController"},
		BSONFieldName:  "destinationtype",
		ConvertedName:  "DestinationType",
		Description:    `Type of the destination ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "destinationType",
		Stored:         true,
		Type:           "enum",
	},
	"encrypted": {
		AllowedChoices: []string{},
		BSONFieldName:  "encrypted",
		ConvertedName:  "Encrypted",
		Description:    `The number of encrypted flows in the edge.`,
		Exposed:        true,
		Name:           "encrypted",
		Stored:         true,
		Type:           "boolean",
	},
	"firstseen": {
		AllowedChoices: []string{},
		BSONFieldName:  "firstseen",
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date when the edge was first seen.`,
		Exposed:        true,
		Name:           "firstSeen",
		Stored:         true,
		Type:           "time",
	},
	"flowid": {
		AllowedChoices: []string{},
		BSONFieldName:  "flowid",
		ConvertedName:  "FlowID",
		Description:    `Identifier of the edge.`,
		Exposed:        true,
		Name:           "flowID",
		Stored:         true,
		Type:           "string",
	},
	"lastseen": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastseen",
		ConvertedName:  "LastSeen",
		Description:    `Contains the date when the edge was last seen.`,
		Exposed:        true,
		Name:           "lastSeen",
		Stored:         true,
		Type:           "time",
	},
	"namespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace of the object that reported the flow.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"observedacceptedflows": {
		AllowedChoices: []string{},
		BSONFieldName:  "observedacceptedflows",
		ConvertedName:  "ObservedAcceptedFlows",
		Description:    `Number of accepted observed flows.`,
		Exposed:        true,
		Name:           "observedAcceptedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"observedencrypted": {
		AllowedChoices: []string{},
		BSONFieldName:  "observedencrypted",
		ConvertedName:  "ObservedEncrypted",
		Description:    `Number of encrypted observed flows.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Stored:         true,
		Type:           "boolean",
	},
	"observedrejectedflows": {
		AllowedChoices: []string{},
		BSONFieldName:  "observedrejectedflows",
		ConvertedName:  "ObservedRejectedFlows",
		Description:    `Number of rejected observed flows.`,
		Exposed:        true,
		Name:           "observedRejectedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"protoport": {
		AllowedChoices: []string{},
		BSONFieldName:  "protoport",
		ConvertedName:  "ProtoPort",
		Description:    `the destination protocol for the flow. If the protocol is TCP or UDP, it will also include the port.`,
		Exposed:        true,
		Name:           "protoPort",
		Stored:         true,
		Type:           "string",
	},
	"rejectedflows": {
		AllowedChoices: []string{},
		BSONFieldName:  "rejectedflows",
		ConvertedName:  "RejectedFlows",
		Description:    `Number of rejected flows in the edge.`,
		Exposed:        true,
		Name:           "rejectedFlows",
		Stored:         true,
		Type:           "boolean",
	},
	"remotenamespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "remotenamespace",
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the object that was targeted by the flow.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Stored:         true,
		Type:           "string",
	},
	"sourcecontroller": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcecontroller",
		ConvertedName:  "SourceController",
		Description:    `Name of the remote source controller if different than the current one.`,
		Exposed:        true,
		Name:           "sourceController",
		Stored:         true,
		Type:           "string",
	},
	"sourceid": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourceid",
		ConvertedName:  "SourceID",
		Description:    `ID of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"sourcetype": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Namespace", "Node", "RemoteController"},
		BSONFieldName:  "sourcetype",
		ConvertedName:  "SourceType",
		Description:    `Type of the source ` + "`" + `GraphNode` + "`" + ` of the edge.`,
		Exposed:        true,
		Name:           "sourceType",
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

// SparseGraphEdgesList represents a list of SparseGraphEdges
type SparseGraphEdgesList []*SparseGraphEdge

// Identity returns the identity of the objects in the list.
func (o SparseGraphEdgesList) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Copy returns a pointer to a copy the SparseGraphEdgesList.
func (o SparseGraphEdgesList) Copy() elemental.Identifiables {

	copy := append(SparseGraphEdgesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseGraphEdgesList.
func (o SparseGraphEdgesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseGraphEdgesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseGraphEdge))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseGraphEdgesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseGraphEdgesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseGraphEdgesList converted to GraphEdgesList.
func (o SparseGraphEdgesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseGraphEdgesList) Version() int {

	return 1
}

// SparseGraphEdge represents the sparse version of a graphedge.
type SparseGraphEdge struct {
	// DB Identifier of the edge.
	ID *string `json:"-" msgpack:"-" bson:"-" mapstructure:"-,omitempty"`

	// Number of accepted flows in the edge.
	AcceptedFlows *bool `json:"acceptedFlows,omitempty" msgpack:"acceptedFlows,omitempty" bson:"acceptedflows,omitempty" mapstructure:"acceptedFlows,omitempty"`

	// The date for the day bucket.
	BucketDay *time.Time `json:"-" msgpack:"-" bson:"bucketday,omitempty" mapstructure:"-,omitempty"`

	// The date for the hour bucket.
	BucketHour *time.Time `json:"-" msgpack:"-" bson:"buckethour,omitempty" mapstructure:"-,omitempty"`

	// The date for the minute bucket.
	BucketMinute *time.Time `json:"-" msgpack:"-" bson:"bucketminute,omitempty" mapstructure:"-,omitempty"`

	// The date for the month bucket.
	BucketMonth *time.Time `json:"-" msgpack:"-" bson:"bucketmonth,omitempty" mapstructure:"-,omitempty"`

	// The truth value that indicates if there are accepted default flows.
	DefaultAcceptedFlows *bool `json:"defaultAcceptedFlows,omitempty" msgpack:"defaultAcceptedFlows,omitempty" bson:"defaultacceptedflows,omitempty" mapstructure:"defaultAcceptedFlows,omitempty"`

	// The truth value that indicates if there are rejected default flows.
	DefaultRejectedFlows *bool `json:"defaultRejectedFlows,omitempty" msgpack:"defaultRejectedFlows,omitempty" bson:"defaultrejectedflows,omitempty" mapstructure:"defaultRejectedFlows,omitempty"`

	// Name of the remote destination controller if different than the current one.
	DestinationController *string `json:"destinationController,omitempty" msgpack:"destinationController,omitempty" bson:"destinationcontroller,omitempty" mapstructure:"destinationController,omitempty"`

	// ID of the destination `GraphNode` of the edge.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"destinationid,omitempty" mapstructure:"destinationID,omitempty"`

	// Type of the destination `GraphNode` of the edge.
	DestinationType *GraphEdgeDestinationTypeValue `json:"destinationType,omitempty" msgpack:"destinationType,omitempty" bson:"destinationtype,omitempty" mapstructure:"destinationType,omitempty"`

	// The number of encrypted flows in the edge.
	Encrypted *bool `json:"encrypted,omitempty" msgpack:"encrypted,omitempty" bson:"encrypted,omitempty" mapstructure:"encrypted,omitempty"`

	// Contains the date when the edge was first seen.
	FirstSeen *time.Time `json:"firstSeen,omitempty" msgpack:"firstSeen,omitempty" bson:"firstseen,omitempty" mapstructure:"firstSeen,omitempty"`

	// Identifier of the edge.
	FlowID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"flowid,omitempty" mapstructure:"ID,omitempty"`

	// Contains the date when the edge was last seen.
	LastSeen *time.Time `json:"lastSeen,omitempty" msgpack:"lastSeen,omitempty" bson:"lastseen,omitempty" mapstructure:"lastSeen,omitempty"`

	// Namespace of the object that reported the flow.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Number of accepted observed flows.
	ObservedAcceptedFlows *bool `json:"observedAcceptedFlows,omitempty" msgpack:"observedAcceptedFlows,omitempty" bson:"observedacceptedflows,omitempty" mapstructure:"observedAcceptedFlows,omitempty"`

	// Number of encrypted observed flows.
	ObservedEncrypted *bool `json:"observedEncrypted,omitempty" msgpack:"observedEncrypted,omitempty" bson:"observedencrypted,omitempty" mapstructure:"observedEncrypted,omitempty"`

	// Number of rejected observed flows.
	ObservedRejectedFlows *bool `json:"observedRejectedFlows,omitempty" msgpack:"observedRejectedFlows,omitempty" bson:"observedrejectedflows,omitempty" mapstructure:"observedRejectedFlows,omitempty"`

	// the destination protocol for the flow. If the protocol is TCP or UDP, it will also include the port.
	ProtoPort *string `json:"protoPort,omitempty" msgpack:"protoPort,omitempty" bson:"protoport,omitempty" mapstructure:"protoPort,omitempty"`

	// Number of rejected flows in the edge.
	RejectedFlows *bool `json:"rejectedFlows,omitempty" msgpack:"rejectedFlows,omitempty" bson:"rejectedflows,omitempty" mapstructure:"rejectedFlows,omitempty"`

	// Namespace of the object that was targeted by the flow.
	RemoteNamespace *string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"remotenamespace,omitempty" mapstructure:"remoteNamespace,omitempty"`

	// Name of the remote source controller if different than the current one.
	SourceController *string `json:"sourceController,omitempty" msgpack:"sourceController,omitempty" bson:"sourcecontroller,omitempty" mapstructure:"sourceController,omitempty"`

	// ID of the source `GraphNode` of the edge.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// Type of the source `GraphNode` of the edge.
	SourceType *GraphEdgeSourceTypeValue `json:"sourceType,omitempty" msgpack:"sourceType,omitempty" bson:"sourcetype,omitempty" mapstructure:"sourceType,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseGraphEdge returns a new  SparseGraphEdge.
func NewSparseGraphEdge() *SparseGraphEdge {
	return &SparseGraphEdge{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseGraphEdge) Identity() elemental.Identity {

	return GraphEdgeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseGraphEdge) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseGraphEdge) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseGraphEdge) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseGraphEdge{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.AcceptedFlows != nil {
		s.AcceptedFlows = o.AcceptedFlows
	}
	if o.BucketDay != nil {
		s.BucketDay = o.BucketDay
	}
	if o.BucketHour != nil {
		s.BucketHour = o.BucketHour
	}
	if o.BucketMinute != nil {
		s.BucketMinute = o.BucketMinute
	}
	if o.BucketMonth != nil {
		s.BucketMonth = o.BucketMonth
	}
	if o.DefaultAcceptedFlows != nil {
		s.DefaultAcceptedFlows = o.DefaultAcceptedFlows
	}
	if o.DefaultRejectedFlows != nil {
		s.DefaultRejectedFlows = o.DefaultRejectedFlows
	}
	if o.DestinationController != nil {
		s.DestinationController = o.DestinationController
	}
	if o.DestinationID != nil {
		s.DestinationID = o.DestinationID
	}
	if o.DestinationType != nil {
		s.DestinationType = o.DestinationType
	}
	if o.Encrypted != nil {
		s.Encrypted = o.Encrypted
	}
	if o.FirstSeen != nil {
		s.FirstSeen = o.FirstSeen
	}
	if o.FlowID != nil {
		s.FlowID = o.FlowID
	}
	if o.LastSeen != nil {
		s.LastSeen = o.LastSeen
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.ObservedAcceptedFlows != nil {
		s.ObservedAcceptedFlows = o.ObservedAcceptedFlows
	}
	if o.ObservedEncrypted != nil {
		s.ObservedEncrypted = o.ObservedEncrypted
	}
	if o.ObservedRejectedFlows != nil {
		s.ObservedRejectedFlows = o.ObservedRejectedFlows
	}
	if o.ProtoPort != nil {
		s.ProtoPort = o.ProtoPort
	}
	if o.RejectedFlows != nil {
		s.RejectedFlows = o.RejectedFlows
	}
	if o.RemoteNamespace != nil {
		s.RemoteNamespace = o.RemoteNamespace
	}
	if o.SourceController != nil {
		s.SourceController = o.SourceController
	}
	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.SourceType != nil {
		s.SourceType = o.SourceType
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
func (o *SparseGraphEdge) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseGraphEdge{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.AcceptedFlows != nil {
		o.AcceptedFlows = s.AcceptedFlows
	}
	if s.BucketDay != nil {
		o.BucketDay = s.BucketDay
	}
	if s.BucketHour != nil {
		o.BucketHour = s.BucketHour
	}
	if s.BucketMinute != nil {
		o.BucketMinute = s.BucketMinute
	}
	if s.BucketMonth != nil {
		o.BucketMonth = s.BucketMonth
	}
	if s.DefaultAcceptedFlows != nil {
		o.DefaultAcceptedFlows = s.DefaultAcceptedFlows
	}
	if s.DefaultRejectedFlows != nil {
		o.DefaultRejectedFlows = s.DefaultRejectedFlows
	}
	if s.DestinationController != nil {
		o.DestinationController = s.DestinationController
	}
	if s.DestinationID != nil {
		o.DestinationID = s.DestinationID
	}
	if s.DestinationType != nil {
		o.DestinationType = s.DestinationType
	}
	if s.Encrypted != nil {
		o.Encrypted = s.Encrypted
	}
	if s.FirstSeen != nil {
		o.FirstSeen = s.FirstSeen
	}
	if s.FlowID != nil {
		o.FlowID = s.FlowID
	}
	if s.LastSeen != nil {
		o.LastSeen = s.LastSeen
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.ObservedAcceptedFlows != nil {
		o.ObservedAcceptedFlows = s.ObservedAcceptedFlows
	}
	if s.ObservedEncrypted != nil {
		o.ObservedEncrypted = s.ObservedEncrypted
	}
	if s.ObservedRejectedFlows != nil {
		o.ObservedRejectedFlows = s.ObservedRejectedFlows
	}
	if s.ProtoPort != nil {
		o.ProtoPort = s.ProtoPort
	}
	if s.RejectedFlows != nil {
		o.RejectedFlows = s.RejectedFlows
	}
	if s.RemoteNamespace != nil {
		o.RemoteNamespace = s.RemoteNamespace
	}
	if s.SourceController != nil {
		o.SourceController = s.SourceController
	}
	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.SourceType != nil {
		o.SourceType = s.SourceType
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
func (o *SparseGraphEdge) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseGraphEdge) ToPlain() elemental.PlainIdentifiable {

	out := NewGraphEdge()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AcceptedFlows != nil {
		out.AcceptedFlows = *o.AcceptedFlows
	}
	if o.BucketDay != nil {
		out.BucketDay = *o.BucketDay
	}
	if o.BucketHour != nil {
		out.BucketHour = *o.BucketHour
	}
	if o.BucketMinute != nil {
		out.BucketMinute = *o.BucketMinute
	}
	if o.BucketMonth != nil {
		out.BucketMonth = *o.BucketMonth
	}
	if o.DefaultAcceptedFlows != nil {
		out.DefaultAcceptedFlows = *o.DefaultAcceptedFlows
	}
	if o.DefaultRejectedFlows != nil {
		out.DefaultRejectedFlows = *o.DefaultRejectedFlows
	}
	if o.DestinationController != nil {
		out.DestinationController = *o.DestinationController
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.DestinationType != nil {
		out.DestinationType = *o.DestinationType
	}
	if o.Encrypted != nil {
		out.Encrypted = *o.Encrypted
	}
	if o.FirstSeen != nil {
		out.FirstSeen = *o.FirstSeen
	}
	if o.FlowID != nil {
		out.FlowID = *o.FlowID
	}
	if o.LastSeen != nil {
		out.LastSeen = *o.LastSeen
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ObservedAcceptedFlows != nil {
		out.ObservedAcceptedFlows = *o.ObservedAcceptedFlows
	}
	if o.ObservedEncrypted != nil {
		out.ObservedEncrypted = *o.ObservedEncrypted
	}
	if o.ObservedRejectedFlows != nil {
		out.ObservedRejectedFlows = *o.ObservedRejectedFlows
	}
	if o.ProtoPort != nil {
		out.ProtoPort = *o.ProtoPort
	}
	if o.RejectedFlows != nil {
		out.RejectedFlows = *o.RejectedFlows
	}
	if o.RemoteNamespace != nil {
		out.RemoteNamespace = *o.RemoteNamespace
	}
	if o.SourceController != nil {
		out.SourceController = *o.SourceController
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceType != nil {
		out.SourceType = *o.SourceType
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseGraphEdge) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseGraphEdge) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseGraphEdge) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseGraphEdge) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseGraphEdge.
func (o *SparseGraphEdge) DeepCopy() *SparseGraphEdge {

	if o == nil {
		return nil
	}

	out := &SparseGraphEdge{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseGraphEdge.
func (o *SparseGraphEdge) DeepCopyInto(out *SparseGraphEdge) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseGraphEdge: %s", err))
	}

	*out = *target.(*SparseGraphEdge)
}

type mongoAttributesGraphEdge struct {
	ID                    bson.ObjectId                 `bson:"_id,omitempty"`
	AcceptedFlows         bool                          `bson:"acceptedflows"`
	BucketDay             time.Time                     `bson:"bucketday"`
	BucketHour            time.Time                     `bson:"buckethour"`
	BucketMinute          time.Time                     `bson:"bucketminute"`
	BucketMonth           time.Time                     `bson:"bucketmonth"`
	DefaultAcceptedFlows  bool                          `bson:"defaultacceptedflows"`
	DefaultRejectedFlows  bool                          `bson:"defaultrejectedflows"`
	DestinationController string                        `bson:"destinationcontroller,omitempty"`
	DestinationID         string                        `bson:"destinationid"`
	DestinationType       GraphEdgeDestinationTypeValue `bson:"destinationtype"`
	Encrypted             bool                          `bson:"encrypted"`
	FirstSeen             time.Time                     `bson:"firstseen,omitempty"`
	FlowID                string                        `bson:"flowid"`
	LastSeen              time.Time                     `bson:"lastseen,omitempty"`
	Namespace             string                        `bson:"namespace"`
	ObservedAcceptedFlows bool                          `bson:"observedacceptedflows"`
	ObservedEncrypted     bool                          `bson:"observedencrypted"`
	ObservedRejectedFlows bool                          `bson:"observedrejectedflows"`
	ProtoPort             string                        `bson:"protoport"`
	RejectedFlows         bool                          `bson:"rejectedflows"`
	RemoteNamespace       string                        `bson:"remotenamespace,omitempty"`
	SourceController      string                        `bson:"sourcecontroller,omitempty"`
	SourceID              string                        `bson:"sourceid"`
	SourceType            GraphEdgeSourceTypeValue      `bson:"sourcetype"`
	ZHash                 int                           `bson:"zhash"`
	Zone                  int                           `bson:"zone"`
}
type mongoAttributesSparseGraphEdge struct {
	ID                    bson.ObjectId                  `bson:"_id,omitempty"`
	AcceptedFlows         *bool                          `bson:"acceptedflows,omitempty"`
	BucketDay             *time.Time                     `bson:"bucketday,omitempty"`
	BucketHour            *time.Time                     `bson:"buckethour,omitempty"`
	BucketMinute          *time.Time                     `bson:"bucketminute,omitempty"`
	BucketMonth           *time.Time                     `bson:"bucketmonth,omitempty"`
	DefaultAcceptedFlows  *bool                          `bson:"defaultacceptedflows,omitempty"`
	DefaultRejectedFlows  *bool                          `bson:"defaultrejectedflows,omitempty"`
	DestinationController *string                        `bson:"destinationcontroller,omitempty"`
	DestinationID         *string                        `bson:"destinationid,omitempty"`
	DestinationType       *GraphEdgeDestinationTypeValue `bson:"destinationtype,omitempty"`
	Encrypted             *bool                          `bson:"encrypted,omitempty"`
	FirstSeen             *time.Time                     `bson:"firstseen,omitempty"`
	FlowID                *string                        `bson:"flowid,omitempty"`
	LastSeen              *time.Time                     `bson:"lastseen,omitempty"`
	Namespace             *string                        `bson:"namespace,omitempty"`
	ObservedAcceptedFlows *bool                          `bson:"observedacceptedflows,omitempty"`
	ObservedEncrypted     *bool                          `bson:"observedencrypted,omitempty"`
	ObservedRejectedFlows *bool                          `bson:"observedrejectedflows,omitempty"`
	ProtoPort             *string                        `bson:"protoport,omitempty"`
	RejectedFlows         *bool                          `bson:"rejectedflows,omitempty"`
	RemoteNamespace       *string                        `bson:"remotenamespace,omitempty"`
	SourceController      *string                        `bson:"sourcecontroller,omitempty"`
	SourceID              *string                        `bson:"sourceid,omitempty"`
	SourceType            *GraphEdgeSourceTypeValue      `bson:"sourcetype,omitempty"`
	ZHash                 *int                           `bson:"zhash,omitempty"`
	Zone                  *int                           `bson:"zone,omitempty"`
}
