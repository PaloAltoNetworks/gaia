package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ConnectionsRequestIdentity represents the Identity of the object.
var ConnectionsRequestIdentity = elemental.Identity{
	Name:     "connectionsrequest",
	Category: "connectionsrequests",
	Package:  "guy",
	Private:  false,
}

// ConnectionsRequestsList represents a list of ConnectionsRequests
type ConnectionsRequestsList []*ConnectionsRequest

// Identity returns the identity of the objects in the list.
func (o ConnectionsRequestsList) Identity() elemental.Identity {

	return ConnectionsRequestIdentity
}

// Copy returns a pointer to a copy the ConnectionsRequestsList.
func (o ConnectionsRequestsList) Copy() elemental.Identifiables {

	copy := append(ConnectionsRequestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ConnectionsRequestsList.
func (o ConnectionsRequestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ConnectionsRequestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ConnectionsRequest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ConnectionsRequestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ConnectionsRequestsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ConnectionsRequestsList converted to SparseConnectionsRequestsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ConnectionsRequestsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseConnectionsRequestsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseConnectionsRequest)
	}

	return out
}

// Version returns the version of the content.
func (o ConnectionsRequestsList) Version() int {

	return 1
}

// ConnectionsRequest represents the model of a connectionsrequest
type ConnectionsRequest struct {
	// Contains the refresh ID set by processing unit refresh event.
	RefreshID string `json:"refreshID" msgpack:"refreshID" bson:"-" mapstructure:"refreshID,omitempty"`

	// Unique ID generated for each request.
	RequestID string `json:"requestID" msgpack:"requestID" bson:"-" mapstructure:"requestID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewConnectionsRequest returns a new *ConnectionsRequest
func NewConnectionsRequest() *ConnectionsRequest {

	return &ConnectionsRequest{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *ConnectionsRequest) Identity() elemental.Identity {

	return ConnectionsRequestIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ConnectionsRequest) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ConnectionsRequest) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsRequest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesConnectionsRequest{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsRequest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesConnectionsRequest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ConnectionsRequest) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ConnectionsRequest) BleveType() string {

	return "connectionsrequest"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ConnectionsRequest) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ConnectionsRequest) Doc() string {

	return `Initiates a request for a PUs current connections.`
}

func (o *ConnectionsRequest) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ConnectionsRequest) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseConnectionsRequest{
			RefreshID: &o.RefreshID,
			RequestID: &o.RequestID,
		}
	}

	sp := &SparseConnectionsRequest{}
	for _, f := range fields {
		switch f {
		case "refreshID":
			sp.RefreshID = &(o.RefreshID)
		case "requestID":
			sp.RequestID = &(o.RequestID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseConnectionsRequest to the object.
func (o *ConnectionsRequest) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseConnectionsRequest)
	if so.RefreshID != nil {
		o.RefreshID = *so.RefreshID
	}
	if so.RequestID != nil {
		o.RequestID = *so.RequestID
	}
}

// DeepCopy returns a deep copy if the ConnectionsRequest.
func (o *ConnectionsRequest) DeepCopy() *ConnectionsRequest {

	if o == nil {
		return nil
	}

	out := &ConnectionsRequest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ConnectionsRequest.
func (o *ConnectionsRequest) DeepCopyInto(out *ConnectionsRequest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ConnectionsRequest: %s", err))
	}

	*out = *target.(*ConnectionsRequest)
}

// Validate valides the current information stored into the structure.
func (o *ConnectionsRequest) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

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
func (*ConnectionsRequest) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ConnectionsRequestAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ConnectionsRequestLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ConnectionsRequest) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ConnectionsRequestAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ConnectionsRequest) ValueForAttribute(name string) interface{} {

	switch name {
	case "refreshID":
		return o.RefreshID
	case "requestID":
		return o.RequestID
	}

	return nil
}

// ConnectionsRequestAttributesMap represents the map of attribute for ConnectionsRequest.
var ConnectionsRequestAttributesMap = map[string]elemental.AttributeSpecification{
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
}

// ConnectionsRequestLowerCaseAttributesMap represents the map of attribute for ConnectionsRequest.
var ConnectionsRequestLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
}

// SparseConnectionsRequestsList represents a list of SparseConnectionsRequests
type SparseConnectionsRequestsList []*SparseConnectionsRequest

// Identity returns the identity of the objects in the list.
func (o SparseConnectionsRequestsList) Identity() elemental.Identity {

	return ConnectionsRequestIdentity
}

// Copy returns a pointer to a copy the SparseConnectionsRequestsList.
func (o SparseConnectionsRequestsList) Copy() elemental.Identifiables {

	copy := append(SparseConnectionsRequestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseConnectionsRequestsList.
func (o SparseConnectionsRequestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseConnectionsRequestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseConnectionsRequest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseConnectionsRequestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseConnectionsRequestsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseConnectionsRequestsList converted to ConnectionsRequestsList.
func (o SparseConnectionsRequestsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseConnectionsRequestsList) Version() int {

	return 1
}

// SparseConnectionsRequest represents the sparse version of a connectionsrequest.
type SparseConnectionsRequest struct {
	// Contains the refresh ID set by processing unit refresh event.
	RefreshID *string `json:"refreshID,omitempty" msgpack:"refreshID,omitempty" bson:"-" mapstructure:"refreshID,omitempty"`

	// Unique ID generated for each request.
	RequestID *string `json:"requestID,omitempty" msgpack:"requestID,omitempty" bson:"-" mapstructure:"requestID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseConnectionsRequest returns a new  SparseConnectionsRequest.
func NewSparseConnectionsRequest() *SparseConnectionsRequest {
	return &SparseConnectionsRequest{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseConnectionsRequest) Identity() elemental.Identity {

	return ConnectionsRequestIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseConnectionsRequest) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseConnectionsRequest) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsRequest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseConnectionsRequest{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsRequest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseConnectionsRequest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseConnectionsRequest) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseConnectionsRequest) ToPlain() elemental.PlainIdentifiable {

	out := NewConnectionsRequest()
	if o.RefreshID != nil {
		out.RefreshID = *o.RefreshID
	}
	if o.RequestID != nil {
		out.RequestID = *o.RequestID
	}

	return out
}

// DeepCopy returns a deep copy if the SparseConnectionsRequest.
func (o *SparseConnectionsRequest) DeepCopy() *SparseConnectionsRequest {

	if o == nil {
		return nil
	}

	out := &SparseConnectionsRequest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseConnectionsRequest.
func (o *SparseConnectionsRequest) DeepCopyInto(out *SparseConnectionsRequest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseConnectionsRequest: %s", err))
	}

	*out = *target.(*SparseConnectionsRequest)
}

type mongoAttributesConnectionsRequest struct {
}
type mongoAttributesSparseConnectionsRequest struct {
}
