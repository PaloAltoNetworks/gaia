package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingRequestIdentity represents the Identity of the object.
var PingRequestIdentity = elemental.Identity{
	Name:     "pingrequest",
	Category: "pingrequests",
	Package:  "guy",
	Private:  false,
}

// PingRequestsList represents a list of PingRequests
type PingRequestsList []*PingRequest

// Identity returns the identity of the objects in the list.
func (o PingRequestsList) Identity() elemental.Identity {

	return PingRequestIdentity
}

// Copy returns a pointer to a copy the PingRequestsList.
func (o PingRequestsList) Copy() elemental.Identifiables {

	copy := append(PingRequestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PingRequestsList.
func (o PingRequestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PingRequestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PingRequest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PingRequestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PingRequestsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PingRequestsList converted to SparsePingRequestsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PingRequestsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePingRequestsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePingRequest)
	}

	return out
}

// Version returns the version of the content.
func (o PingRequestsList) Version() int {

	return 1
}

// PingRequest represents the model of a pingrequest
type PingRequest struct {
	// Number of probes that will be triggered.
	Iterations int `json:"iterations" msgpack:"iterations" bson:"-" mapstructure:"iterations,omitempty"`

	// Unique ID generated for each ping request.
	PingID string `json:"pingID" msgpack:"pingID" bson:"-" mapstructure:"pingID,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID string `json:"refreshID" msgpack:"refreshID" bson:"-" mapstructure:"refreshID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingRequest returns a new *PingRequest
func NewPingRequest() *PingRequest {

	return &PingRequest{
		ModelVersion: 1,
		Iterations:   1,
	}
}

// Identity returns the Identity of the object.
func (o *PingRequest) Identity() elemental.Identity {

	return PingRequestIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PingRequest) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PingRequest) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingRequest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingRequest{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingRequest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingRequest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PingRequest) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingRequest) BleveType() string {

	return "pingrequest"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PingRequest) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PingRequest) Doc() string {

	return `Initiates a ping request for enforcer debugging.`
}

func (o *PingRequest) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PingRequest) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePingRequest{
			Iterations: &o.Iterations,
			PingID:     &o.PingID,
			RefreshID:  &o.RefreshID,
		}
	}

	sp := &SparsePingRequest{}
	for _, f := range fields {
		switch f {
		case "iterations":
			sp.Iterations = &(o.Iterations)
		case "pingID":
			sp.PingID = &(o.PingID)
		case "refreshID":
			sp.RefreshID = &(o.RefreshID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePingRequest to the object.
func (o *PingRequest) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePingRequest)
	if so.Iterations != nil {
		o.Iterations = *so.Iterations
	}
	if so.PingID != nil {
		o.PingID = *so.PingID
	}
	if so.RefreshID != nil {
		o.RefreshID = *so.RefreshID
	}
}

// DeepCopy returns a deep copy if the PingRequest.
func (o *PingRequest) DeepCopy() *PingRequest {

	if o == nil {
		return nil
	}

	out := &PingRequest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingRequest.
func (o *PingRequest) DeepCopyInto(out *PingRequest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingRequest: %s", err))
	}

	*out = *target.(*PingRequest)
}

// Validate valides the current information stored into the structure.
func (o *PingRequest) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumInt("iterations", o.Iterations, int(20), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMinimumInt("iterations", o.Iterations, int(1), false); err != nil {
		errors = errors.Append(err)
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
func (*PingRequest) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingRequestAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingRequestLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PingRequest) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingRequestAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PingRequest) ValueForAttribute(name string) interface{} {

	switch name {
	case "iterations":
		return o.Iterations
	case "pingID":
		return o.PingID
	case "refreshID":
		return o.RefreshID
	}

	return nil
}

// PingRequestAttributesMap represents the map of attribute for PingRequest.
var PingRequestAttributesMap = map[string]elemental.AttributeSpecification{
	"Iterations": {
		AllowedChoices: []string{},
		ConvertedName:  "Iterations",
		DefaultValue:   1,
		Description:    `Number of probes that will be triggered.`,
		Exposed:        true,
		MaxValue:       20,
		MinValue:       1,
		Name:           "iterations",
		Type:           "integer",
	},
	"PingID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PingID",
		Description:    `Unique ID generated for each ping request.`,
		Exposed:        true,
		Name:           "pingID",
		ReadOnly:       true,
		Type:           "string",
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
}

// PingRequestLowerCaseAttributesMap represents the map of attribute for PingRequest.
var PingRequestLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"iterations": {
		AllowedChoices: []string{},
		ConvertedName:  "Iterations",
		DefaultValue:   1,
		Description:    `Number of probes that will be triggered.`,
		Exposed:        true,
		MaxValue:       20,
		MinValue:       1,
		Name:           "iterations",
		Type:           "integer",
	},
	"pingid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PingID",
		Description:    `Unique ID generated for each ping request.`,
		Exposed:        true,
		Name:           "pingID",
		ReadOnly:       true,
		Type:           "string",
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
}

// SparsePingRequestsList represents a list of SparsePingRequests
type SparsePingRequestsList []*SparsePingRequest

// Identity returns the identity of the objects in the list.
func (o SparsePingRequestsList) Identity() elemental.Identity {

	return PingRequestIdentity
}

// Copy returns a pointer to a copy the SparsePingRequestsList.
func (o SparsePingRequestsList) Copy() elemental.Identifiables {

	copy := append(SparsePingRequestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePingRequestsList.
func (o SparsePingRequestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePingRequestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePingRequest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePingRequestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePingRequestsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePingRequestsList converted to PingRequestsList.
func (o SparsePingRequestsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePingRequestsList) Version() int {

	return 1
}

// SparsePingRequest represents the sparse version of a pingrequest.
type SparsePingRequest struct {
	// Number of probes that will be triggered.
	Iterations *int `json:"iterations,omitempty" msgpack:"iterations,omitempty" bson:"-" mapstructure:"iterations,omitempty"`

	// Unique ID generated for each ping request.
	PingID *string `json:"pingID,omitempty" msgpack:"pingID,omitempty" bson:"-" mapstructure:"pingID,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID *string `json:"refreshID,omitempty" msgpack:"refreshID,omitempty" bson:"-" mapstructure:"refreshID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePingRequest returns a new  SparsePingRequest.
func NewSparsePingRequest() *SparsePingRequest {
	return &SparsePingRequest{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePingRequest) Identity() elemental.Identity {

	return PingRequestIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePingRequest) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePingRequest) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingRequest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePingRequest{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingRequest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePingRequest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePingRequest) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePingRequest) ToPlain() elemental.PlainIdentifiable {

	out := NewPingRequest()
	if o.Iterations != nil {
		out.Iterations = *o.Iterations
	}
	if o.PingID != nil {
		out.PingID = *o.PingID
	}
	if o.RefreshID != nil {
		out.RefreshID = *o.RefreshID
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePingRequest.
func (o *SparsePingRequest) DeepCopy() *SparsePingRequest {

	if o == nil {
		return nil
	}

	out := &SparsePingRequest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePingRequest.
func (o *SparsePingRequest) DeepCopyInto(out *SparsePingRequest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePingRequest: %s", err))
	}

	*out = *target.(*SparsePingRequest)
}

type mongoAttributesPingRequest struct {
}
type mongoAttributesSparsePingRequest struct {
}
