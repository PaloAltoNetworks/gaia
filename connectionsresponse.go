package gaia

import (
	"fmt"
	"time"

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
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Contains a batch of connections.
	Connections CurrentConnectionsList `json:"connections" msgpack:"connections" bson:"connections" mapstructure:"connections,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID string `json:"refreshID" msgpack:"refreshID" bson:"-" mapstructure:"refreshID,omitempty"`

	// Unique ID generated for each request.
	RequestID string `json:"requestID" msgpack:"requestID" bson:"-" mapstructure:"requestID,omitempty"`

	// Contains the total number of connections for the connection request.
	TotalConnections int `json:"totalConnections,omitempty" msgpack:"totalConnections,omitempty" bson:"-" mapstructure:"totalConnections,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewConnectionsResponse returns a new *ConnectionsResponse
func NewConnectionsResponse() *ConnectionsResponse {

	return &ConnectionsResponse{
		ModelVersion:  1,
		Connections:   CurrentConnectionsList{},
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *ConnectionsResponse) Identity() elemental.Identity {

	return ConnectionsResponseIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ConnectionsResponse) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ConnectionsResponse) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ConnectionsResponse) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesConnectionsResponse{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Connections = o.Connections
	s.CreateTime = o.CreateTime
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

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

	o.ID = s.ID.Hex()
	o.Connections = s.Connections
	o.CreateTime = s.CreateTime
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

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

// GetCreateTime returns the CreateTime of the receiver.
func (o *ConnectionsResponse) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ConnectionsResponse) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *ConnectionsResponse) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *ConnectionsResponse) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *ConnectionsResponse) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ConnectionsResponse) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ConnectionsResponse) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ConnectionsResponse) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *ConnectionsResponse) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ConnectionsResponse) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ConnectionsResponse) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ConnectionsResponse) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ConnectionsResponse) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseConnectionsResponse{
			ID:               &o.ID,
			Connections:      &o.Connections,
			CreateTime:       &o.CreateTime,
			MigrationsLog:    &o.MigrationsLog,
			Namespace:        &o.Namespace,
			RefreshID:        &o.RefreshID,
			RequestID:        &o.RequestID,
			TotalConnections: &o.TotalConnections,
			UpdateTime:       &o.UpdateTime,
			ZHash:            &o.ZHash,
			Zone:             &o.Zone,
		}
	}

	sp := &SparseConnectionsResponse{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "connections":
			sp.Connections = &(o.Connections)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "refreshID":
			sp.RefreshID = &(o.RefreshID)
		case "requestID":
			sp.RequestID = &(o.RequestID)
		case "totalConnections":
			sp.TotalConnections = &(o.TotalConnections)
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

// Patch apply the non nil value of a *SparseConnectionsResponse to the object.
func (o *ConnectionsResponse) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseConnectionsResponse)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Connections != nil {
		o.Connections = *so.Connections
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
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
	case "ID":
		return o.ID
	case "connections":
		return o.Connections
	case "createTime":
		return o.CreateTime
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "refreshID":
		return o.RefreshID
	case "requestID":
		return o.RequestID
	case "totalConnections":
		return o.TotalConnections
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ConnectionsResponseAttributesMap represents the map of attribute for ConnectionsResponse.
var ConnectionsResponseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
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
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
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

// ConnectionsResponseLowerCaseAttributesMap represents the map of attribute for ConnectionsResponse.
var ConnectionsResponseLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
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
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
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
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Contains a batch of connections.
	Connections *CurrentConnectionsList `json:"connections,omitempty" msgpack:"connections,omitempty" bson:"connections,omitempty" mapstructure:"connections,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the refresh ID set by processing unit refresh event.
	RefreshID *string `json:"refreshID,omitempty" msgpack:"refreshID,omitempty" bson:"-" mapstructure:"refreshID,omitempty"`

	// Unique ID generated for each request.
	RequestID *string `json:"requestID,omitempty" msgpack:"requestID,omitempty" bson:"-" mapstructure:"requestID,omitempty"`

	// Contains the total number of connections for the connection request.
	TotalConnections *int `json:"totalConnections,omitempty" msgpack:"totalConnections,omitempty" bson:"-" mapstructure:"totalConnections,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

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

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseConnectionsResponse) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseConnectionsResponse) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseConnectionsResponse{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Connections != nil {
		s.Connections = o.Connections
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
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
func (o *SparseConnectionsResponse) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseConnectionsResponse{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Connections != nil {
		o.Connections = s.Connections
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
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
func (o *SparseConnectionsResponse) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseConnectionsResponse) ToPlain() elemental.PlainIdentifiable {

	out := NewConnectionsResponse()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Connections != nil {
		out.Connections = *o.Connections
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
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
func (o *SparseConnectionsResponse) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseConnectionsResponse) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseConnectionsResponse) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseConnectionsResponse) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseConnectionsResponse) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseConnectionsResponse) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseConnectionsResponse) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseConnectionsResponse) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseConnectionsResponse) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseConnectionsResponse) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseConnectionsResponse) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseConnectionsResponse) SetZone(zone int) {

	o.Zone = &zone
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
	ID            bson.ObjectId          `bson:"_id,omitempty"`
	Connections   CurrentConnectionsList `bson:"connections"`
	CreateTime    time.Time              `bson:"createtime"`
	MigrationsLog map[string]string      `bson:"migrationslog,omitempty"`
	Namespace     string                 `bson:"namespace"`
	UpdateTime    time.Time              `bson:"updatetime"`
	ZHash         int                    `bson:"zhash"`
	Zone          int                    `bson:"zone"`
}
type mongoAttributesSparseConnectionsResponse struct {
	ID            bson.ObjectId           `bson:"_id,omitempty"`
	Connections   *CurrentConnectionsList `bson:"connections,omitempty"`
	CreateTime    *time.Time              `bson:"createtime,omitempty"`
	MigrationsLog *map[string]string      `bson:"migrationslog,omitempty"`
	Namespace     *string                 `bson:"namespace,omitempty"`
	UpdateTime    *time.Time              `bson:"updatetime,omitempty"`
	ZHash         *int                    `bson:"zhash,omitempty"`
	Zone          *int                    `bson:"zone,omitempty"`
}
