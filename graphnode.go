package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphNodeTypeValue represents the possible values for attribute "type".
type GraphNodeTypeValue string

const (
	// GraphNodeTypeClaim represents the value Claim.
	GraphNodeTypeClaim GraphNodeTypeValue = "Claim"

	// GraphNodeTypeDocker represents the value Docker.
	GraphNodeTypeDocker GraphNodeTypeValue = "Docker"

	// GraphNodeTypeExternalNetwork represents the value ExternalNetwork.
	GraphNodeTypeExternalNetwork GraphNodeTypeValue = "ExternalNetwork"

	// GraphNodeTypeNamespace represents the value Namespace.
	GraphNodeTypeNamespace GraphNodeTypeValue = "Namespace"

	// GraphNodeTypeNode represents the value Node.
	GraphNodeTypeNode GraphNodeTypeValue = "Node"

	// GraphNodeTypeRemoteController represents the value RemoteController.
	GraphNodeTypeRemoteController GraphNodeTypeValue = "RemoteController"

	// GraphNodeTypeVolume represents the value Volume.
	GraphNodeTypeVolume GraphNodeTypeValue = "Volume"
)

// GraphNodeIdentity represents the Identity of the object.
var GraphNodeIdentity = elemental.Identity{
	Name:     "graphnode",
	Category: "graphnodes",
	Package:  "meteor",
	Private:  true,
}

// GraphNodesList represents a list of GraphNodes
type GraphNodesList []*GraphNode

// Identity returns the identity of the objects in the list.
func (o GraphNodesList) Identity() elemental.Identity {

	return GraphNodeIdentity
}

// Copy returns a pointer to a copy the GraphNodesList.
func (o GraphNodesList) Copy() elemental.Identifiables {

	copy := append(GraphNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the GraphNodesList.
func (o GraphNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(GraphNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*GraphNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o GraphNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o GraphNodesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the GraphNodesList converted to SparseGraphNodesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o GraphNodesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseGraphNodesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseGraphNode)
	}

	return out
}

// Version returns the version of the content.
func (o GraphNodesList) Version() int {

	return 1
}

// GraphNode represents the model of a graphnode
type GraphNode struct {
	// Identifier of object represented by the node.
	ID string `json:"ID" msgpack:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// Enforcement status of processing unit represented by the node.
	EnforcementStatus string `json:"enforcementStatus" msgpack:"enforcementStatus" bson:"enforcementstatus" mapstructure:"enforcementStatus,omitempty"`

	// Contains the date when the edge was first seen.
	FirstSeen time.Time `json:"firstSeen" msgpack:"firstSeen" bson:"firstseen" mapstructure:"firstSeen,omitempty"`

	// ID of the group the node is eventually part of.
	GroupID string `json:"groupID" msgpack:"groupID" bson:"groupid" mapstructure:"groupID,omitempty"`

	// List of images.
	Images []string `json:"images" msgpack:"images" bson:"images" mapstructure:"images,omitempty"`

	// Contains the date when the edge was last seen.
	LastSeen time.Time `json:"lastSeen" msgpack:"lastSeen" bson:"lastseen" mapstructure:"lastSeen,omitempty"`

	// Name of object represented by the node.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace of object represented by the node.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Status of object represented by the node.
	Status string `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// Tags of object represented by the node.
	Tags []string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags" mapstructure:"tags,omitempty"`

	// Type of object represented by the node.
	Type GraphNodeTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// If `true` the node is marked as unreachable.
	Unreachable bool `json:"unreachable" msgpack:"unreachable" bson:"unreachable" mapstructure:"unreachable,omitempty"`

	// Tags of object represented by the node.
	VulnerabilityLevel string `json:"vulnerabilityLevel" msgpack:"vulnerabilityLevel" bson:"vulnerabilitylevel" mapstructure:"vulnerabilityLevel,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGraphNode returns a new *GraphNode
func NewGraphNode() *GraphNode {

	return &GraphNode{
		ModelVersion: 1,
		Images:       []string{},
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *GraphNode) Identity() elemental.Identity {

	return GraphNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GraphNode) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GraphNode) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GraphNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesGraphNode{}

	s.ID = o.ID
	s.EnforcementStatus = o.EnforcementStatus
	s.FirstSeen = o.FirstSeen
	s.GroupID = o.GroupID
	s.Images = o.Images
	s.LastSeen = o.LastSeen
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.Status = o.Status
	s.Tags = o.Tags
	s.Type = o.Type
	s.Unreachable = o.Unreachable
	s.VulnerabilityLevel = o.VulnerabilityLevel

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *GraphNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesGraphNode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID
	o.EnforcementStatus = s.EnforcementStatus
	o.FirstSeen = s.FirstSeen
	o.GroupID = s.GroupID
	o.Images = s.Images
	o.LastSeen = s.LastSeen
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.Status = s.Status
	o.Tags = s.Tags
	o.Type = s.Type
	o.Unreachable = s.Unreachable
	o.VulnerabilityLevel = s.VulnerabilityLevel

	return nil
}

// Version returns the hardcoded version of the model.
func (o *GraphNode) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *GraphNode) BleveType() string {

	return "graphnode"
}

// DefaultOrder returns the list of default ordering fields.
func (o *GraphNode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *GraphNode) Doc() string {

	return `Represents an node from the dependency map.`
}

func (o *GraphNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *GraphNode) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseGraphNode{
			ID:                 &o.ID,
			EnforcementStatus:  &o.EnforcementStatus,
			FirstSeen:          &o.FirstSeen,
			GroupID:            &o.GroupID,
			Images:             &o.Images,
			LastSeen:           &o.LastSeen,
			Name:               &o.Name,
			Namespace:          &o.Namespace,
			Status:             &o.Status,
			Tags:               &o.Tags,
			Type:               &o.Type,
			Unreachable:        &o.Unreachable,
			VulnerabilityLevel: &o.VulnerabilityLevel,
		}
	}

	sp := &SparseGraphNode{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "enforcementStatus":
			sp.EnforcementStatus = &(o.EnforcementStatus)
		case "firstSeen":
			sp.FirstSeen = &(o.FirstSeen)
		case "groupID":
			sp.GroupID = &(o.GroupID)
		case "images":
			sp.Images = &(o.Images)
		case "lastSeen":
			sp.LastSeen = &(o.LastSeen)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "status":
			sp.Status = &(o.Status)
		case "tags":
			sp.Tags = &(o.Tags)
		case "type":
			sp.Type = &(o.Type)
		case "unreachable":
			sp.Unreachable = &(o.Unreachable)
		case "vulnerabilityLevel":
			sp.VulnerabilityLevel = &(o.VulnerabilityLevel)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseGraphNode to the object.
func (o *GraphNode) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseGraphNode)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.EnforcementStatus != nil {
		o.EnforcementStatus = *so.EnforcementStatus
	}
	if so.FirstSeen != nil {
		o.FirstSeen = *so.FirstSeen
	}
	if so.GroupID != nil {
		o.GroupID = *so.GroupID
	}
	if so.Images != nil {
		o.Images = *so.Images
	}
	if so.LastSeen != nil {
		o.LastSeen = *so.LastSeen
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.Unreachable != nil {
		o.Unreachable = *so.Unreachable
	}
	if so.VulnerabilityLevel != nil {
		o.VulnerabilityLevel = *so.VulnerabilityLevel
	}
}

// DeepCopy returns a deep copy if the GraphNode.
func (o *GraphNode) DeepCopy() *GraphNode {

	if o == nil {
		return nil
	}

	out := &GraphNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphNode.
func (o *GraphNode) DeepCopyInto(out *GraphNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphNode: %s", err))
	}

	*out = *target.(*GraphNode)
}

// Validate valides the current information stored into the structure.
func (o *GraphNode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Docker", "ExternalNetwork", "Volume", "Claim", "Node", "Namespace", "RemoteController"}, false); err != nil {
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
func (*GraphNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := GraphNodeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return GraphNodeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*GraphNode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return GraphNodeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *GraphNode) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "enforcementStatus":
		return o.EnforcementStatus
	case "firstSeen":
		return o.FirstSeen
	case "groupID":
		return o.GroupID
	case "images":
		return o.Images
	case "lastSeen":
		return o.LastSeen
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "status":
		return o.Status
	case "tags":
		return o.Tags
	case "type":
		return o.Type
	case "unreachable":
		return o.Unreachable
	case "vulnerabilityLevel":
		return o.VulnerabilityLevel
	}

	return nil
}

// GraphNodeAttributesMap represents the map of attribute for GraphNode.
var GraphNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Identifier of object represented by the node.`,
		Exposed:        true,
		Name:           "ID",
		Stored:         true,
		Type:           "string",
	},
	"EnforcementStatus": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcementStatus",
		Description:    `Enforcement status of processing unit represented by the node.`,
		Exposed:        true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "string",
	},
	"FirstSeen": {
		AllowedChoices: []string{},
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date when the edge was first seen.`,
		Exposed:        true,
		Name:           "firstSeen",
		Stored:         true,
		Type:           "time",
	},
	"GroupID": {
		AllowedChoices: []string{},
		ConvertedName:  "GroupID",
		Description:    `ID of the group the node is eventually part of.`,
		Exposed:        true,
		Name:           "groupID",
		Stored:         true,
		Type:           "string",
	},
	"Images": {
		AllowedChoices: []string{},
		ConvertedName:  "Images",
		Description:    `List of images.`,
		Exposed:        true,
		Name:           "images",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"LastSeen": {
		AllowedChoices: []string{},
		ConvertedName:  "LastSeen",
		Description:    `Contains the date when the edge was last seen.`,
		Exposed:        true,
		Name:           "lastSeen",
		Stored:         true,
		Type:           "time",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of object represented by the node.`,
		Exposed:        true,
		Name:           "name",
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of object represented by the node.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `Status of object represented by the node.`,
		Exposed:        true,
		Name:           "status",
		Stored:         true,
		Type:           "string",
	},
	"Tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags of object represented by the node.`,
		Exposed:        true,
		Name:           "tags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Type": {
		AllowedChoices: []string{"Docker", "ExternalNetwork", "Volume", "Claim", "Node", "Namespace", "RemoteController"},
		ConvertedName:  "Type",
		Description:    `Type of object represented by the node.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
	},
	"Unreachable": {
		AllowedChoices: []string{},
		ConvertedName:  "Unreachable",
		Description:    `If ` + "`" + `true` + "`" + ` the node is marked as unreachable.`,
		Exposed:        true,
		Name:           "unreachable",
		Stored:         true,
		Type:           "boolean",
	},
	"VulnerabilityLevel": {
		AllowedChoices: []string{},
		ConvertedName:  "VulnerabilityLevel",
		Description:    `Tags of object represented by the node.`,
		Exposed:        true,
		Name:           "vulnerabilityLevel",
		Stored:         true,
		Type:           "string",
	},
}

// GraphNodeLowerCaseAttributesMap represents the map of attribute for GraphNode.
var GraphNodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Identifier of object represented by the node.`,
		Exposed:        true,
		Name:           "ID",
		Stored:         true,
		Type:           "string",
	},
	"enforcementstatus": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcementStatus",
		Description:    `Enforcement status of processing unit represented by the node.`,
		Exposed:        true,
		Name:           "enforcementStatus",
		Stored:         true,
		Type:           "string",
	},
	"firstseen": {
		AllowedChoices: []string{},
		ConvertedName:  "FirstSeen",
		Description:    `Contains the date when the edge was first seen.`,
		Exposed:        true,
		Name:           "firstSeen",
		Stored:         true,
		Type:           "time",
	},
	"groupid": {
		AllowedChoices: []string{},
		ConvertedName:  "GroupID",
		Description:    `ID of the group the node is eventually part of.`,
		Exposed:        true,
		Name:           "groupID",
		Stored:         true,
		Type:           "string",
	},
	"images": {
		AllowedChoices: []string{},
		ConvertedName:  "Images",
		Description:    `List of images.`,
		Exposed:        true,
		Name:           "images",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"lastseen": {
		AllowedChoices: []string{},
		ConvertedName:  "LastSeen",
		Description:    `Contains the date when the edge was last seen.`,
		Exposed:        true,
		Name:           "lastSeen",
		Stored:         true,
		Type:           "time",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of object represented by the node.`,
		Exposed:        true,
		Name:           "name",
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of object represented by the node.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `Status of object represented by the node.`,
		Exposed:        true,
		Name:           "status",
		Stored:         true,
		Type:           "string",
	},
	"tags": {
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags of object represented by the node.`,
		Exposed:        true,
		Name:           "tags",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"type": {
		AllowedChoices: []string{"Docker", "ExternalNetwork", "Volume", "Claim", "Node", "Namespace", "RemoteController"},
		ConvertedName:  "Type",
		Description:    `Type of object represented by the node.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
	},
	"unreachable": {
		AllowedChoices: []string{},
		ConvertedName:  "Unreachable",
		Description:    `If ` + "`" + `true` + "`" + ` the node is marked as unreachable.`,
		Exposed:        true,
		Name:           "unreachable",
		Stored:         true,
		Type:           "boolean",
	},
	"vulnerabilitylevel": {
		AllowedChoices: []string{},
		ConvertedName:  "VulnerabilityLevel",
		Description:    `Tags of object represented by the node.`,
		Exposed:        true,
		Name:           "vulnerabilityLevel",
		Stored:         true,
		Type:           "string",
	},
}

// SparseGraphNodesList represents a list of SparseGraphNodes
type SparseGraphNodesList []*SparseGraphNode

// Identity returns the identity of the objects in the list.
func (o SparseGraphNodesList) Identity() elemental.Identity {

	return GraphNodeIdentity
}

// Copy returns a pointer to a copy the SparseGraphNodesList.
func (o SparseGraphNodesList) Copy() elemental.Identifiables {

	copy := append(SparseGraphNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseGraphNodesList.
func (o SparseGraphNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseGraphNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseGraphNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseGraphNodesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseGraphNodesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseGraphNodesList converted to GraphNodesList.
func (o SparseGraphNodesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseGraphNodesList) Version() int {

	return 1
}

// SparseGraphNode represents the sparse version of a graphnode.
type SparseGraphNode struct {
	// Identifier of object represented by the node.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"id,omitempty" mapstructure:"ID,omitempty"`

	// Enforcement status of processing unit represented by the node.
	EnforcementStatus *string `json:"enforcementStatus,omitempty" msgpack:"enforcementStatus,omitempty" bson:"enforcementstatus,omitempty" mapstructure:"enforcementStatus,omitempty"`

	// Contains the date when the edge was first seen.
	FirstSeen *time.Time `json:"firstSeen,omitempty" msgpack:"firstSeen,omitempty" bson:"firstseen,omitempty" mapstructure:"firstSeen,omitempty"`

	// ID of the group the node is eventually part of.
	GroupID *string `json:"groupID,omitempty" msgpack:"groupID,omitempty" bson:"groupid,omitempty" mapstructure:"groupID,omitempty"`

	// List of images.
	Images *[]string `json:"images,omitempty" msgpack:"images,omitempty" bson:"images,omitempty" mapstructure:"images,omitempty"`

	// Contains the date when the edge was last seen.
	LastSeen *time.Time `json:"lastSeen,omitempty" msgpack:"lastSeen,omitempty" bson:"lastseen,omitempty" mapstructure:"lastSeen,omitempty"`

	// Name of object represented by the node.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace of object represented by the node.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Status of object represented by the node.
	Status *string `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// Tags of object represented by the node.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags,omitempty" mapstructure:"tags,omitempty"`

	// Type of object represented by the node.
	Type *GraphNodeTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// If `true` the node is marked as unreachable.
	Unreachable *bool `json:"unreachable,omitempty" msgpack:"unreachable,omitempty" bson:"unreachable,omitempty" mapstructure:"unreachable,omitempty"`

	// Tags of object represented by the node.
	VulnerabilityLevel *string `json:"vulnerabilityLevel,omitempty" msgpack:"vulnerabilityLevel,omitempty" bson:"vulnerabilitylevel,omitempty" mapstructure:"vulnerabilityLevel,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseGraphNode returns a new  SparseGraphNode.
func NewSparseGraphNode() *SparseGraphNode {
	return &SparseGraphNode{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseGraphNode) Identity() elemental.Identity {

	return GraphNodeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseGraphNode) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseGraphNode) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseGraphNode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseGraphNode{}

	if o.ID != nil {
		s.ID = o.ID
	}
	if o.EnforcementStatus != nil {
		s.EnforcementStatus = o.EnforcementStatus
	}
	if o.FirstSeen != nil {
		s.FirstSeen = o.FirstSeen
	}
	if o.GroupID != nil {
		s.GroupID = o.GroupID
	}
	if o.Images != nil {
		s.Images = o.Images
	}
	if o.LastSeen != nil {
		s.LastSeen = o.LastSeen
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Status != nil {
		s.Status = o.Status
	}
	if o.Tags != nil {
		s.Tags = o.Tags
	}
	if o.Type != nil {
		s.Type = o.Type
	}
	if o.Unreachable != nil {
		s.Unreachable = o.Unreachable
	}
	if o.VulnerabilityLevel != nil {
		s.VulnerabilityLevel = o.VulnerabilityLevel
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseGraphNode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseGraphNode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.ID != nil {
		o.ID = s.ID
	}
	if s.EnforcementStatus != nil {
		o.EnforcementStatus = s.EnforcementStatus
	}
	if s.FirstSeen != nil {
		o.FirstSeen = s.FirstSeen
	}
	if s.GroupID != nil {
		o.GroupID = s.GroupID
	}
	if s.Images != nil {
		o.Images = s.Images
	}
	if s.LastSeen != nil {
		o.LastSeen = s.LastSeen
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Status != nil {
		o.Status = s.Status
	}
	if s.Tags != nil {
		o.Tags = s.Tags
	}
	if s.Type != nil {
		o.Type = s.Type
	}
	if s.Unreachable != nil {
		o.Unreachable = s.Unreachable
	}
	if s.VulnerabilityLevel != nil {
		o.VulnerabilityLevel = s.VulnerabilityLevel
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseGraphNode) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseGraphNode) ToPlain() elemental.PlainIdentifiable {

	out := NewGraphNode()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.EnforcementStatus != nil {
		out.EnforcementStatus = *o.EnforcementStatus
	}
	if o.FirstSeen != nil {
		out.FirstSeen = *o.FirstSeen
	}
	if o.GroupID != nil {
		out.GroupID = *o.GroupID
	}
	if o.Images != nil {
		out.Images = *o.Images
	}
	if o.LastSeen != nil {
		out.LastSeen = *o.LastSeen
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.Unreachable != nil {
		out.Unreachable = *o.Unreachable
	}
	if o.VulnerabilityLevel != nil {
		out.VulnerabilityLevel = *o.VulnerabilityLevel
	}

	return out
}

// DeepCopy returns a deep copy if the SparseGraphNode.
func (o *SparseGraphNode) DeepCopy() *SparseGraphNode {

	if o == nil {
		return nil
	}

	out := &SparseGraphNode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseGraphNode.
func (o *SparseGraphNode) DeepCopyInto(out *SparseGraphNode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseGraphNode: %s", err))
	}

	*out = *target.(*SparseGraphNode)
}

type mongoAttributesGraphNode struct {
	ID                 string             `bson:"id"`
	EnforcementStatus  string             `bson:"enforcementstatus"`
	FirstSeen          time.Time          `bson:"firstseen"`
	GroupID            string             `bson:"groupid"`
	Images             []string           `bson:"images"`
	LastSeen           time.Time          `bson:"lastseen"`
	Name               string             `bson:"name"`
	Namespace          string             `bson:"namespace"`
	Status             string             `bson:"status"`
	Tags               []string           `bson:"tags,omitempty"`
	Type               GraphNodeTypeValue `bson:"type"`
	Unreachable        bool               `bson:"unreachable"`
	VulnerabilityLevel string             `bson:"vulnerabilitylevel"`
}
type mongoAttributesSparseGraphNode struct {
	ID                 *string             `bson:"id,omitempty"`
	EnforcementStatus  *string             `bson:"enforcementstatus,omitempty"`
	FirstSeen          *time.Time          `bson:"firstseen,omitempty"`
	GroupID            *string             `bson:"groupid,omitempty"`
	Images             *[]string           `bson:"images,omitempty"`
	LastSeen           *time.Time          `bson:"lastseen,omitempty"`
	Name               *string             `bson:"name,omitempty"`
	Namespace          *string             `bson:"namespace,omitempty"`
	Status             *string             `bson:"status,omitempty"`
	Tags               *[]string           `bson:"tags,omitempty"`
	Type               *GraphNodeTypeValue `bson:"type,omitempty"`
	Unreachable        *bool               `bson:"unreachable,omitempty"`
	VulnerabilityLevel *string             `bson:"vulnerabilitylevel,omitempty"`
}
