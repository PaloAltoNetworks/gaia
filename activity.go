package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ActivityIdentity represents the Identity of the object.
var ActivityIdentity = elemental.Identity{
	Name:     "activity",
	Category: "activities",
	Package:  "hojo",
	Private:  false,
}

// ActivitiesList represents a list of Activities
type ActivitiesList []*Activity

// Identity returns the identity of the objects in the list.
func (o ActivitiesList) Identity() elemental.Identity {

	return ActivityIdentity
}

// Copy returns a pointer to a copy the ActivitiesList.
func (o ActivitiesList) Copy() elemental.Identifiables {

	copy := append(ActivitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ActivitiesList.
func (o ActivitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ActivitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Activity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ActivitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ActivitiesList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToSparse returns the ActivitiesList converted to SparseActivitiesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ActivitiesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseActivitiesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseActivity)
	}

	return out
}

// Version returns the version of the content.
func (o ActivitiesList) Version() int {

	return 1
}

// Activity represents the model of a activity
type Activity struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Claims of the user who performed the operation.
	Claims interface{} `json:"claims" msgpack:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// Data of the notification.
	Data interface{} `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// Time-date stamp of the notification.
	Date time.Time `json:"date" msgpack:"date" bson:"date" mapstructure:"date,omitempty"`

	// Contains the error.
	Error interface{} `json:"error" msgpack:"error" bson:"error" mapstructure:"error,omitempty"`

	// Message of the notification.
	Message string `json:"message" msgpack:"message" bson:"message" mapstructure:"message,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Describes what kind of operation the notification represents.
	Operation string `json:"operation" msgpack:"operation" bson:"operation" mapstructure:"operation,omitempty"`

	// Contains the original data of the object that was modified.
	OriginalData interface{} `json:"originalData" msgpack:"originalData" bson:"originaldata" mapstructure:"originalData,omitempty"`

	// Contains meta information about the source.
	Source string `json:"source" msgpack:"source" bson:"source" mapstructure:"source,omitempty"`

	// The identity of the related object.
	TargetIdentity string `json:"targetIdentity" msgpack:"targetIdentity" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewActivity returns a new *Activity
func NewActivity() *Activity {

	return &Activity{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Activity) Identity() elemental.Identity {

	return ActivityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activity) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Activity) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Activity) BleveType() string {

	return "activity"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Activity) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// Doc returns the documentation for the object
func (o *Activity) Doc() string {

	return `Contains logs of all the activity that happened in a namespace. All successful or
failed actions will be available, errors, as well as the claims of
the user who triggered the actiions. This log is capped and only keeps the last
50,000 entries by default.`
}

func (o *Activity) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *Activity) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Activity) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *Activity) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Activity) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Activity) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Activity) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Activity) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseActivity{
			ID:             &o.ID,
			Claims:         &o.Claims,
			Data:           &o.Data,
			Date:           &o.Date,
			Error:          &o.Error,
			Message:        &o.Message,
			Namespace:      &o.Namespace,
			Operation:      &o.Operation,
			OriginalData:   &o.OriginalData,
			Source:         &o.Source,
			TargetIdentity: &o.TargetIdentity,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseActivity{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "claims":
			sp.Claims = &(o.Claims)
		case "data":
			sp.Data = &(o.Data)
		case "date":
			sp.Date = &(o.Date)
		case "error":
			sp.Error = &(o.Error)
		case "message":
			sp.Message = &(o.Message)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "operation":
			sp.Operation = &(o.Operation)
		case "originalData":
			sp.OriginalData = &(o.OriginalData)
		case "source":
			sp.Source = &(o.Source)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseActivity to the object.
func (o *Activity) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseActivity)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.Message != nil {
		o.Message = *so.Message
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Operation != nil {
		o.Operation = *so.Operation
	}
	if so.OriginalData != nil {
		o.OriginalData = *so.OriginalData
	}
	if so.Source != nil {
		o.Source = *so.Source
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Activity.
func (o *Activity) DeepCopy() *Activity {

	if o == nil {
		return nil
	}

	out := &Activity{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Activity.
func (o *Activity) DeepCopyInto(out *Activity) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Activity: %s", err))
	}

	*out = *target.(*Activity)
}

// Validate valides the current information stored into the structure.
func (o *Activity) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Activity) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ActivityAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ActivityLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Activity) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivityAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Activity) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "claims":
		return o.Claims
	case "data":
		return o.Data
	case "date":
		return o.Date
	case "error":
		return o.Error
	case "message":
		return o.Message
	case "namespace":
		return o.Namespace
	case "operation":
		return o.Operation
	case "originalData":
		return o.OriginalData
	case "source":
		return o.Source
	case "targetIdentity":
		return o.TargetIdentity
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ActivityAttributesMap represents the map of attribute for Activity.
var ActivityAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Data of the notification.`,
		Exposed:        true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		Description:    `Time-date stamp of the notification.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Contains the error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"Message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Message",
		Description:    `Message of the notification.`,
		Exposed:        true,
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
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
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Operation",
		Description:    `Describes what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"OriginalData": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OriginalData",
		Description:    `Contains the original data of the object that was modified.`,
		Exposed:        true,
		Name:           "originalData",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"Source": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Source",
		Description:    `Contains meta information about the source.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "source",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TargetIdentity",
		Description:    `The identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ZHash": elemental.AttributeSpecification{
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
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Exposed:   true,
		Getter:    true,
		Name:      "zone",
		ReadOnly:  true,
		Setter:    true,
		Stored:    true,
		Transient: true,
		Type:      "integer",
	},
}

// ActivityLowerCaseAttributesMap represents the map of attribute for Activity.
var ActivityLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `Data of the notification.`,
		Exposed:        true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		Description:    `Time-date stamp of the notification.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Error",
		Description:    `Contains the error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Message",
		Description:    `Message of the notification.`,
		Exposed:        true,
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
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
	"operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Operation",
		Description:    `Describes what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"originaldata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OriginalData",
		Description:    `Contains the original data of the object that was modified.`,
		Exposed:        true,
		Name:           "originalData",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "object",
	},
	"source": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Source",
		Description:    `Contains meta information about the source.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "source",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TargetIdentity",
		Description:    `The identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"zhash": elemental.AttributeSpecification{
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
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Exposed:   true,
		Getter:    true,
		Name:      "zone",
		ReadOnly:  true,
		Setter:    true,
		Stored:    true,
		Transient: true,
		Type:      "integer",
	},
}

// SparseActivitiesList represents a list of SparseActivities
type SparseActivitiesList []*SparseActivity

// Identity returns the identity of the objects in the list.
func (o SparseActivitiesList) Identity() elemental.Identity {

	return ActivityIdentity
}

// Copy returns a pointer to a copy the SparseActivitiesList.
func (o SparseActivitiesList) Copy() elemental.Identifiables {

	copy := append(SparseActivitiesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseActivitiesList.
func (o SparseActivitiesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseActivitiesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseActivity))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseActivitiesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseActivitiesList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToPlain returns the SparseActivitiesList converted to ActivitiesList.
func (o SparseActivitiesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseActivitiesList) Version() int {

	return 1
}

// SparseActivity represents the sparse version of a activity.
type SparseActivity struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Claims of the user who performed the operation.
	Claims *interface{} `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"claims,omitempty" mapstructure:"claims,omitempty"`

	// Data of the notification.
	Data *interface{} `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// Time-date stamp of the notification.
	Date *time.Time `json:"date,omitempty" msgpack:"date,omitempty" bson:"date,omitempty" mapstructure:"date,omitempty"`

	// Contains the error.
	Error *interface{} `json:"error,omitempty" msgpack:"error,omitempty" bson:"error,omitempty" mapstructure:"error,omitempty"`

	// Message of the notification.
	Message *string `json:"message,omitempty" msgpack:"message,omitempty" bson:"message,omitempty" mapstructure:"message,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Describes what kind of operation the notification represents.
	Operation *string `json:"operation,omitempty" msgpack:"operation,omitempty" bson:"operation,omitempty" mapstructure:"operation,omitempty"`

	// Contains the original data of the object that was modified.
	OriginalData *interface{} `json:"originalData,omitempty" msgpack:"originalData,omitempty" bson:"originaldata,omitempty" mapstructure:"originalData,omitempty"`

	// Contains meta information about the source.
	Source *string `json:"source,omitempty" msgpack:"source,omitempty" bson:"source,omitempty" mapstructure:"source,omitempty"`

	// The identity of the related object.
	TargetIdentity *string `json:"targetIdentity,omitempty" msgpack:"targetIdentity,omitempty" bson:"targetidentity,omitempty" mapstructure:"targetIdentity,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseActivity returns a new  SparseActivity.
func NewSparseActivity() *SparseActivity {
	return &SparseActivity{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseActivity) Identity() elemental.Identity {

	return ActivityIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseActivity) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseActivity) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseActivity) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseActivity) ToPlain() elemental.PlainIdentifiable {

	out := NewActivity()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.Message != nil {
		out.Message = *o.Message
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Operation != nil {
		out.Operation = *o.Operation
	}
	if o.OriginalData != nil {
		out.OriginalData = *o.OriginalData
	}
	if o.Source != nil {
		out.Source = *o.Source
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseActivity) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseActivity) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseActivity) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseActivity) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseActivity) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseActivity) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseActivity.
func (o *SparseActivity) DeepCopy() *SparseActivity {

	if o == nil {
		return nil
	}

	out := &SparseActivity{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseActivity.
func (o *SparseActivity) DeepCopyInto(out *SparseActivity) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseActivity: %s", err))
	}

	*out = *target.(*SparseActivity)
}
