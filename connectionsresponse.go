package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ConnectionsResponseIdentity represents the Identity of the object.
var ConnectionsResponseIdentity = elemental.Identity{
	Name:     "connectionsresponse",
	Category: "connectionsresponses",
	Package:  "guy",
	Private:  false,
}

// ConnectionsResponsesList represents a list of ConnectionsResponses
type ConnectionsResponsesList []*ConnectionsResponse

// Identity returns the identity of the objects in the list.
func (o ConnectionsResponsesList) Identity() elemental.Identity {

	return ConnectionsResponseIdentity
}

// Copy returns a pointer to a copy the ConnectionsResponsesList.
func (o ConnectionsResponsesList) Copy() elemental.Identifiables {

	copy := append(ConnectionsResponsesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ConnectionsResponsesList.
func (o ConnectionsResponsesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ConnectionsResponsesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ConnectionsResponse))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ConnectionsResponsesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ConnectionsResponsesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ConnectionsResponsesList converted to SparseConnectionsResponsesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ConnectionsResponsesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseConnectionsResponsesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseConnectionsResponse)
	}

	return out
}

// Version returns the version of the content.
func (o ConnectionsResponsesList) Version() int {

	return 1
}

// ConnectionsResponse represents the model of a connectionsresponse
type ConnectionsResponse struct {
	// Contains a batch of connections.
	Connections CurrentConnectionsList `json:"connections" msgpack:"connections" bson:"connections" mapstructure:"connections,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID string `json:"refreshID" msgpack:"refreshID" bson:"-" mapstructure:"refreshID,omitempty"`

	// Unique ID generated for each request.
	RequestID string `json:"requestID" msgpack:"requestID" bson:"-" mapstructure:"requestID,omitempty"`

	// Contains the total number of connections for the connection request.
	TotalConnections int `json:"totalConnections,omitempty" msgpack:"totalConnections,omitempty" bson:"-" mapstructure:"totalConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewConnectionsResponse returns a new *ConnectionsResponse
func NewConnectionsResponse() *ConnectionsResponse {

	return &ConnectionsResponse{
		ModelVersion: 1,
		Connections:  CurrentConnectionsList{},
	}
}

// Identity returns the Identity of the object.
func (o *ConnectionsResponse) Identity() elemental.Identity {

	return ConnectionsResponseIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ConnectionsResponse) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ConnectionsResponse) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsResponse) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesConnectionsResponse{}

	s.Connections = o.Connections

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsResponse) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesConnectionsResponse{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Connections = s.Connections

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ConnectionsResponse) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ConnectionsResponse) BleveType() string {

	return "connectionsresponse"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ConnectionsResponse) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ConnectionsResponse) Doc() string {

	return `A response from a connectionsrequest.`
}

func (o *ConnectionsResponse) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ConnectionsResponse) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseConnectionsResponse{
			Connections:      &o.Connections,
			RefreshID:        &o.RefreshID,
			RequestID:        &o.RequestID,
			TotalConnections: &o.TotalConnections,
		}
	}

	sp := &SparseConnectionsResponse{}
	for _, f := range fields {
		switch f {
		case "connections":
			sp.Connections = &(o.Connections)
		case "refreshID":
			sp.RefreshID = &(o.RefreshID)
		case "requestID":
			sp.RequestID = &(o.RequestID)
		case "totalConnections":
			sp.TotalConnections = &(o.TotalConnections)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseConnectionsResponse to the object.
func (o *ConnectionsResponse) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseConnectionsResponse)
	if so.Connections != nil {
		o.Connections = *so.Connections
	}
	if so.RefreshID != nil {
		o.RefreshID = *so.RefreshID
	}
	if so.RequestID != nil {
		o.RequestID = *so.RequestID
	}
	if so.TotalConnections != nil {
		o.TotalConnections = *so.TotalConnections
	}
}

// DeepCopy returns a deep copy if the ConnectionsResponse.
func (o *ConnectionsResponse) DeepCopy() *ConnectionsResponse {

	if o == nil {
		return nil
	}

	out := &ConnectionsResponse{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ConnectionsResponse.
func (o *ConnectionsResponse) DeepCopyInto(out *ConnectionsResponse) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ConnectionsResponse: %s", err))
	}

	*out = *target.(*ConnectionsResponse)
}

// Validate valides the current information stored into the structure.
func (o *ConnectionsResponse) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Connections {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredString("refreshID", o.RefreshID); err != nil {
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
func (*ConnectionsResponse) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ConnectionsResponseAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ConnectionsResponseLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ConnectionsResponse) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ConnectionsResponseAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ConnectionsResponse) ValueForAttribute(name string) interface{} {

	switch name {
	case "connections":
		return o.Connections
	case "refreshID":
		return o.RefreshID
	case "requestID":
		return o.RequestID
	case "totalConnections":
		return o.TotalConnections
	}

	return nil
}

// ConnectionsResponseAttributesMap represents the map of attribute for ConnectionsResponse.
var ConnectionsResponseAttributesMap = map[string]elemental.AttributeSpecification{
	"Connections": {
		AllowedChoices: []string{},
		BSONFieldName:  "connections",
		ConvertedName:  "Connections",
		Description:    `Contains a batch of connections.`,
		Exposed:        true,
		Name:           "connections",
		Stored:         true,
		SubType:        "currentconnection",
		Type:           "refList",
	},
	"RefreshID": {
		AllowedChoices: []string{},
		ConvertedName:  "RefreshID",
		Description:    `Contains the refresh ID set by processing unit refresh event.`,
		Exposed:        true,
		Name:           "refreshID",
		Required:       true,
		Type:           "string",
	},
	"RequestID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequestID",
		Description:    `Unique ID generated for each request.`,
		Exposed:        true,
		Name:           "requestID",
		ReadOnly:       true,
		Type:           "string",
	},
	"TotalConnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TotalConnections",
		Description:    `Contains the total number of connections for the connection request.`,
		Exposed:        true,
		Name:           "totalConnections",
		Type:           "integer",
	},
}

// ConnectionsResponseLowerCaseAttributesMap represents the map of attribute for ConnectionsResponse.
var ConnectionsResponseLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"connections": {
		AllowedChoices: []string{},
		BSONFieldName:  "connections",
		ConvertedName:  "Connections",
		Description:    `Contains a batch of connections.`,
		Exposed:        true,
		Name:           "connections",
		Stored:         true,
		SubType:        "currentconnection",
		Type:           "refList",
	},
	"refreshid": {
		AllowedChoices: []string{},
		ConvertedName:  "RefreshID",
		Description:    `Contains the refresh ID set by processing unit refresh event.`,
		Exposed:        true,
		Name:           "refreshID",
		Required:       true,
		Type:           "string",
	},
	"requestid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequestID",
		Description:    `Unique ID generated for each request.`,
		Exposed:        true,
		Name:           "requestID",
		ReadOnly:       true,
		Type:           "string",
	},
	"totalconnections": {
		AllowedChoices: []string{},
		ConvertedName:  "TotalConnections",
		Description:    `Contains the total number of connections for the connection request.`,
		Exposed:        true,
		Name:           "totalConnections",
		Type:           "integer",
	},
}

// SparseConnectionsResponsesList represents a list of SparseConnectionsResponses
type SparseConnectionsResponsesList []*SparseConnectionsResponse

// Identity returns the identity of the objects in the list.
func (o SparseConnectionsResponsesList) Identity() elemental.Identity {

	return ConnectionsResponseIdentity
}

// Copy returns a pointer to a copy the SparseConnectionsResponsesList.
func (o SparseConnectionsResponsesList) Copy() elemental.Identifiables {

	copy := append(SparseConnectionsResponsesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseConnectionsResponsesList.
func (o SparseConnectionsResponsesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseConnectionsResponsesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseConnectionsResponse))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseConnectionsResponsesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseConnectionsResponsesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseConnectionsResponsesList converted to ConnectionsResponsesList.
func (o SparseConnectionsResponsesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseConnectionsResponsesList) Version() int {

	return 1
}

// SparseConnectionsResponse represents the sparse version of a connectionsresponse.
type SparseConnectionsResponse struct {
	// Contains a batch of connections.
	Connections *CurrentConnectionsList `json:"connections,omitempty" msgpack:"connections,omitempty" bson:"connections,omitempty" mapstructure:"connections,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID *string `json:"refreshID,omitempty" msgpack:"refreshID,omitempty" bson:"-" mapstructure:"refreshID,omitempty"`

	// Unique ID generated for each request.
	RequestID *string `json:"requestID,omitempty" msgpack:"requestID,omitempty" bson:"-" mapstructure:"requestID,omitempty"`

	// Contains the total number of connections for the connection request.
	TotalConnections *int `json:"totalConnections,omitempty" msgpack:"totalConnections,omitempty" bson:"-" mapstructure:"totalConnections,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseConnectionsResponse returns a new  SparseConnectionsResponse.
func NewSparseConnectionsResponse() *SparseConnectionsResponse {
	return &SparseConnectionsResponse{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseConnectionsResponse) Identity() elemental.Identity {

	return ConnectionsResponseIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseConnectionsResponse) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseConnectionsResponse) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsResponse) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseConnectionsResponse{}

	if o.Connections != nil {
		s.Connections = o.Connections
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsResponse) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseConnectionsResponse{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Connections != nil {
		o.Connections = s.Connections
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseConnectionsResponse) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseConnectionsResponse) ToPlain() elemental.PlainIdentifiable {

	out := NewConnectionsResponse()
	if o.Connections != nil {
		out.Connections = *o.Connections
	}
	if o.RefreshID != nil {
		out.RefreshID = *o.RefreshID
	}
	if o.RequestID != nil {
		out.RequestID = *o.RequestID
	}
	if o.TotalConnections != nil {
		out.TotalConnections = *o.TotalConnections
	}

	return out
}

// DeepCopy returns a deep copy if the SparseConnectionsResponse.
func (o *SparseConnectionsResponse) DeepCopy() *SparseConnectionsResponse {

	if o == nil {
		return nil
	}

	out := &SparseConnectionsResponse{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseConnectionsResponse.
func (o *SparseConnectionsResponse) DeepCopyInto(out *SparseConnectionsResponse) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseConnectionsResponse: %s", err))
	}

	*out = *target.(*SparseConnectionsResponse)
}

type mongoAttributesConnectionsResponse struct {
	Connections CurrentConnectionsList `bson:"connections"`
}
type mongoAttributesSparseConnectionsResponse struct {
	Connections *CurrentConnectionsList `bson:"connections,omitempty"`
}
