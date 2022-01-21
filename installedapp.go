package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// InstalledAppStatusValue represents the possible values for attribute "status".
type InstalledAppStatusValue string

const (
	// InstalledAppStatusDeploying represents the value Deploying.
	InstalledAppStatusDeploying InstalledAppStatusValue = "Deploying"

	// InstalledAppStatusError represents the value Error.
	InstalledAppStatusError InstalledAppStatusValue = "Error"

	// InstalledAppStatusInitializing represents the value Initializing.
	InstalledAppStatusInitializing InstalledAppStatusValue = "Initializing"

	// InstalledAppStatusRunning represents the value Running.
	InstalledAppStatusRunning InstalledAppStatusValue = "Running"

	// InstalledAppStatusUndeploying represents the value Undeploying.
	InstalledAppStatusUndeploying InstalledAppStatusValue = "Undeploying"

	// InstalledAppStatusUnknown represents the value Unknown.
	InstalledAppStatusUnknown InstalledAppStatusValue = "Unknown"
)

// InstalledAppIdentity represents the Identity of the object.
var InstalledAppIdentity = elemental.Identity{
	Name:     "installedapp",
	Category: "installedapps",
	Package:  "highwind",
	Private:  false,
}

// InstalledAppsList represents a list of InstalledApps
type InstalledAppsList []*InstalledApp

// Identity returns the identity of the objects in the list.
func (o InstalledAppsList) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Copy returns a pointer to a copy the InstalledAppsList.
func (o InstalledAppsList) Copy() elemental.Identifiables {

	copy := append(InstalledAppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InstalledAppsList.
func (o InstalledAppsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InstalledAppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*InstalledApp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InstalledAppsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InstalledAppsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the InstalledAppsList converted to SparseInstalledAppsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o InstalledAppsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseInstalledAppsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseInstalledApp)
	}

	return out
}

// Version returns the version of the content.
func (o InstalledAppsList) Version() int {

	return 1
}

// InstalledApp represents the model of a installedapp
type InstalledApp struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Additional configuration of the app is needed by the app itself.
	AdditionalConfiguration bool `json:"additionalConfiguration" msgpack:"additionalConfiguration" bson:"additionalconfiguration" mapstructure:"additionalConfiguration,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AppIdentifier retains the identifier for the app.
	AppIdentifier string `json:"-" msgpack:"-" bson:"appidentifier" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The category ID of the application.
	CategoryID string `json:"categoryID" msgpack:"categoryID" bson:"categoryid" mapstructure:"categoryID,omitempty"`

	// If true, will look for the public endpoints and store them as annotations in the
	// installed app.
	CheckPublicEndpoint bool `json:"checkPublicEndpoint" msgpack:"checkPublicEndpoint" bson:"checkpublicendpoint" mapstructure:"checkPublicEndpoint,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Version of the installed application.
	CurrentVersion string `json:"currentVersion" msgpack:"currentVersion" bson:"currentversion" mapstructure:"currentVersion,omitempty"`

	// DeploymentCount represents the number of expected deployment for this app.
	DeploymentCount int `json:"-" msgpack:"-" bson:"deploymentcount" mapstructure:"-,omitempty"`

	// Adds a button in the UI.
	ExternalWindowButton map[string]string `json:"externalWindowButton" msgpack:"externalWindowButton" bson:"externalwindowbutton" mapstructure:"externalWindowButton,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the computed parameters to start the application.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Status of the application.
	Status InstalledAppStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// Reason for the status of the application.
	StatusMessage string `json:"statusMessage" msgpack:"statusMessage" bson:"statusmessage" mapstructure:"statusMessage,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewInstalledApp returns a new *InstalledApp
func NewInstalledApp() *InstalledApp {

	return &InstalledApp{
		ModelVersion:         1,
		Annotations:          map[string][]string{},
		AssociatedTags:       []string{},
		ExternalWindowButton: map[string]string{},
		MigrationsLog:        map[string]string{},
		NormalizedTags:       []string{},
		Parameters:           map[string]interface{}{},
		Status:               InstalledAppStatusUnknown,
	}
}

// Identity returns the Identity of the object.
func (o *InstalledApp) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InstalledApp) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InstalledApp) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *InstalledApp) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesInstalledApp{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.AdditionalConfiguration = o.AdditionalConfiguration
	s.Annotations = o.Annotations
	s.AppIdentifier = o.AppIdentifier
	s.AssociatedTags = o.AssociatedTags
	s.CategoryID = o.CategoryID
	s.CheckPublicEndpoint = o.CheckPublicEndpoint
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.CurrentVersion = o.CurrentVersion
	s.DeploymentCount = o.DeploymentCount
	s.ExternalWindowButton = o.ExternalWindowButton
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Parameters = o.Parameters
	s.Protected = o.Protected
	s.Status = o.Status
	s.StatusMessage = o.StatusMessage
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *InstalledApp) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesInstalledApp{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.AdditionalConfiguration = s.AdditionalConfiguration
	o.Annotations = s.Annotations
	o.AppIdentifier = s.AppIdentifier
	o.AssociatedTags = s.AssociatedTags
	o.CategoryID = s.CategoryID
	o.CheckPublicEndpoint = s.CheckPublicEndpoint
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.CurrentVersion = s.CurrentVersion
	o.DeploymentCount = s.DeploymentCount
	o.ExternalWindowButton = s.ExternalWindowButton
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Parameters = s.Parameters
	o.Protected = s.Protected
	o.Status = s.Status
	o.StatusMessage = s.StatusMessage
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *InstalledApp) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *InstalledApp) BleveType() string {

	return "installedapp"
}

// DefaultOrder returns the list of default ordering fields.
func (o *InstalledApp) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *InstalledApp) Doc() string {

	return `Represents an installed application.`
}

func (o *InstalledApp) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *InstalledApp) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *InstalledApp) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *InstalledApp) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *InstalledApp) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *InstalledApp) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *InstalledApp) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *InstalledApp) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *InstalledApp) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *InstalledApp) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *InstalledApp) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *InstalledApp) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *InstalledApp) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *InstalledApp) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *InstalledApp) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *InstalledApp) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *InstalledApp) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *InstalledApp) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *InstalledApp) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *InstalledApp) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *InstalledApp) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *InstalledApp) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *InstalledApp) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *InstalledApp) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *InstalledApp) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *InstalledApp) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *InstalledApp) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *InstalledApp) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInstalledApp{
			ID:                      &o.ID,
			AdditionalConfiguration: &o.AdditionalConfiguration,
			Annotations:             &o.Annotations,
			AppIdentifier:           &o.AppIdentifier,
			AssociatedTags:          &o.AssociatedTags,
			CategoryID:              &o.CategoryID,
			CheckPublicEndpoint:     &o.CheckPublicEndpoint,
			CreateIdempotencyKey:    &o.CreateIdempotencyKey,
			CreateTime:              &o.CreateTime,
			CurrentVersion:          &o.CurrentVersion,
			DeploymentCount:         &o.DeploymentCount,
			ExternalWindowButton:    &o.ExternalWindowButton,
			MigrationsLog:           &o.MigrationsLog,
			Name:                    &o.Name,
			Namespace:               &o.Namespace,
			NormalizedTags:          &o.NormalizedTags,
			Parameters:              &o.Parameters,
			Protected:               &o.Protected,
			Status:                  &o.Status,
			StatusMessage:           &o.StatusMessage,
			UpdateIdempotencyKey:    &o.UpdateIdempotencyKey,
			UpdateTime:              &o.UpdateTime,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparseInstalledApp{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "additionalConfiguration":
			sp.AdditionalConfiguration = &(o.AdditionalConfiguration)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "appIdentifier":
			sp.AppIdentifier = &(o.AppIdentifier)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "categoryID":
			sp.CategoryID = &(o.CategoryID)
		case "checkPublicEndpoint":
			sp.CheckPublicEndpoint = &(o.CheckPublicEndpoint)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "currentVersion":
			sp.CurrentVersion = &(o.CurrentVersion)
		case "deploymentCount":
			sp.DeploymentCount = &(o.DeploymentCount)
		case "externalWindowButton":
			sp.ExternalWindowButton = &(o.ExternalWindowButton)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "protected":
			sp.Protected = &(o.Protected)
		case "status":
			sp.Status = &(o.Status)
		case "statusMessage":
			sp.StatusMessage = &(o.StatusMessage)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
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

// Patch apply the non nil value of a *SparseInstalledApp to the object.
func (o *InstalledApp) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseInstalledApp)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AdditionalConfiguration != nil {
		o.AdditionalConfiguration = *so.AdditionalConfiguration
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AppIdentifier != nil {
		o.AppIdentifier = *so.AppIdentifier
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CategoryID != nil {
		o.CategoryID = *so.CategoryID
	}
	if so.CheckPublicEndpoint != nil {
		o.CheckPublicEndpoint = *so.CheckPublicEndpoint
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.CurrentVersion != nil {
		o.CurrentVersion = *so.CurrentVersion
	}
	if so.DeploymentCount != nil {
		o.DeploymentCount = *so.DeploymentCount
	}
	if so.ExternalWindowButton != nil {
		o.ExternalWindowButton = *so.ExternalWindowButton
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.StatusMessage != nil {
		o.StatusMessage = *so.StatusMessage
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
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

// DeepCopy returns a deep copy if the InstalledApp.
func (o *InstalledApp) DeepCopy() *InstalledApp {

	if o == nil {
		return nil
	}

	out := &InstalledApp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *InstalledApp.
func (o *InstalledApp) DeepCopyInto(out *InstalledApp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy InstalledApp: %s", err))
	}

	*out = *target.(*InstalledApp)
}

// Validate valides the current information stored into the structure.
func (o *InstalledApp) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Unknown", "Deploying", "Initializing", "Running", "Undeploying", "Error"}, false); err != nil {
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
func (*InstalledApp) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InstalledAppAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InstalledAppLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*InstalledApp) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InstalledAppAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *InstalledApp) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "additionalConfiguration":
		return o.AdditionalConfiguration
	case "annotations":
		return o.Annotations
	case "appIdentifier":
		return o.AppIdentifier
	case "associatedTags":
		return o.AssociatedTags
	case "categoryID":
		return o.CategoryID
	case "checkPublicEndpoint":
		return o.CheckPublicEndpoint
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "currentVersion":
		return o.CurrentVersion
	case "deploymentCount":
		return o.DeploymentCount
	case "externalWindowButton":
		return o.ExternalWindowButton
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "parameters":
		return o.Parameters
	case "protected":
		return o.Protected
	case "status":
		return o.Status
	case "statusMessage":
		return o.StatusMessage
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// InstalledAppAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppAttributesMap = map[string]elemental.AttributeSpecification{
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
	"AdditionalConfiguration": {
		AllowedChoices: []string{},
		ConvertedName:  "AdditionalConfiguration",
		Description:    `Additional configuration of the app is needed by the app itself.`,
		Exposed:        true,
		Name:           "additionalConfiguration",
		Stored:         true,
		Type:           "boolean",
	},
	"Annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AppIdentifier": {
		AllowedChoices: []string{},
		ConvertedName:  "AppIdentifier",
		Description:    `AppIdentifier retains the identifier for the app.`,
		Name:           "appIdentifier",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CategoryID": {
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `The category ID of the application.`,
		Exposed:        true,
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CheckPublicEndpoint": {
		AllowedChoices: []string{},
		ConvertedName:  "CheckPublicEndpoint",
		Description: `If true, will look for the public endpoints and store them as annotations in the
installed app.`,
		Exposed: true,
		Name:    "checkPublicEndpoint",
		Stored:  true,
		Type:    "boolean",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"CurrentVersion": {
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description:    `Version of the installed application.`,
		Exposed:        true,
		Name:           "currentVersion",
		Stored:         true,
		Type:           "string",
	},
	"DeploymentCount": {
		AllowedChoices: []string{},
		ConvertedName:  "DeploymentCount",
		Description:    `DeploymentCount represents the number of expected deployment for this app.`,
		Name:           "deploymentCount",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"ExternalWindowButton": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalWindowButton",
		Description:    `Adds a button in the UI.`,
		Exposed:        true,
		Name:           "externalWindowButton",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters to start the application.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Status": {
		AllowedChoices: []string{"Unknown", "Deploying", "Initializing", "Running", "Undeploying", "Error"},
		ConvertedName:  "Status",
		DefaultValue:   InstalledAppStatusUnknown,
		Description:    `Status of the application.`,
		Exposed:        true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"StatusMessage": {
		AllowedChoices: []string{},
		ConvertedName:  "StatusMessage",
		Description:    `Reason for the status of the application.`,
		Exposed:        true,
		Name:           "statusMessage",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// InstalledAppLowerCaseAttributesMap represents the map of attribute for InstalledApp.
var InstalledAppLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"additionalconfiguration": {
		AllowedChoices: []string{},
		ConvertedName:  "AdditionalConfiguration",
		Description:    `Additional configuration of the app is needed by the app itself.`,
		Exposed:        true,
		Name:           "additionalConfiguration",
		Stored:         true,
		Type:           "boolean",
	},
	"annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"appidentifier": {
		AllowedChoices: []string{},
		ConvertedName:  "AppIdentifier",
		Description:    `AppIdentifier retains the identifier for the app.`,
		Name:           "appIdentifier",
		Stored:         true,
		Type:           "string",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"categoryid": {
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `The category ID of the application.`,
		Exposed:        true,
		Name:           "categoryID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"checkpublicendpoint": {
		AllowedChoices: []string{},
		ConvertedName:  "CheckPublicEndpoint",
		Description: `If true, will look for the public endpoints and store them as annotations in the
installed app.`,
		Exposed: true,
		Name:    "checkPublicEndpoint",
		Stored:  true,
		Type:    "boolean",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"currentversion": {
		AllowedChoices: []string{},
		ConvertedName:  "CurrentVersion",
		Description:    `Version of the installed application.`,
		Exposed:        true,
		Name:           "currentVersion",
		Stored:         true,
		Type:           "string",
	},
	"deploymentcount": {
		AllowedChoices: []string{},
		ConvertedName:  "DeploymentCount",
		Description:    `DeploymentCount represents the number of expected deployment for this app.`,
		Name:           "deploymentCount",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"externalwindowbutton": {
		AllowedChoices: []string{},
		ConvertedName:  "ExternalWindowButton",
		Description:    `Adds a button in the UI.`,
		Exposed:        true,
		Name:           "externalWindowButton",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters to start the application.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"status": {
		AllowedChoices: []string{"Unknown", "Deploying", "Initializing", "Running", "Undeploying", "Error"},
		ConvertedName:  "Status",
		DefaultValue:   InstalledAppStatusUnknown,
		Description:    `Status of the application.`,
		Exposed:        true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"statusmessage": {
		AllowedChoices: []string{},
		ConvertedName:  "StatusMessage",
		Description:    `Reason for the status of the application.`,
		Exposed:        true,
		Name:           "statusMessage",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SparseInstalledAppsList represents a list of SparseInstalledApps
type SparseInstalledAppsList []*SparseInstalledApp

// Identity returns the identity of the objects in the list.
func (o SparseInstalledAppsList) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Copy returns a pointer to a copy the SparseInstalledAppsList.
func (o SparseInstalledAppsList) Copy() elemental.Identifiables {

	copy := append(SparseInstalledAppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseInstalledAppsList.
func (o SparseInstalledAppsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseInstalledAppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseInstalledApp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseInstalledAppsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseInstalledAppsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseInstalledAppsList converted to InstalledAppsList.
func (o SparseInstalledAppsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseInstalledAppsList) Version() int {

	return 1
}

// SparseInstalledApp represents the sparse version of a installedapp.
type SparseInstalledApp struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Additional configuration of the app is needed by the app itself.
	AdditionalConfiguration *bool `json:"additionalConfiguration,omitempty" msgpack:"additionalConfiguration,omitempty" bson:"additionalconfiguration,omitempty" mapstructure:"additionalConfiguration,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AppIdentifier retains the identifier for the app.
	AppIdentifier *string `json:"-" msgpack:"-" bson:"appidentifier,omitempty" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The category ID of the application.
	CategoryID *string `json:"categoryID,omitempty" msgpack:"categoryID,omitempty" bson:"categoryid,omitempty" mapstructure:"categoryID,omitempty"`

	// If true, will look for the public endpoints and store them as annotations in the
	// installed app.
	CheckPublicEndpoint *bool `json:"checkPublicEndpoint,omitempty" msgpack:"checkPublicEndpoint,omitempty" bson:"checkpublicendpoint,omitempty" mapstructure:"checkPublicEndpoint,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Version of the installed application.
	CurrentVersion *string `json:"currentVersion,omitempty" msgpack:"currentVersion,omitempty" bson:"currentversion,omitempty" mapstructure:"currentVersion,omitempty"`

	// DeploymentCount represents the number of expected deployment for this app.
	DeploymentCount *int `json:"-" msgpack:"-" bson:"deploymentcount,omitempty" mapstructure:"-,omitempty"`

	// Adds a button in the UI.
	ExternalWindowButton *map[string]string `json:"externalWindowButton,omitempty" msgpack:"externalWindowButton,omitempty" bson:"externalwindowbutton,omitempty" mapstructure:"externalWindowButton,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Contains the computed parameters to start the application.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Status of the application.
	Status *InstalledAppStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// Reason for the status of the application.
	StatusMessage *string `json:"statusMessage,omitempty" msgpack:"statusMessage,omitempty" bson:"statusmessage,omitempty" mapstructure:"statusMessage,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseInstalledApp returns a new  SparseInstalledApp.
func NewSparseInstalledApp() *SparseInstalledApp {
	return &SparseInstalledApp{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseInstalledApp) Identity() elemental.Identity {

	return InstalledAppIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseInstalledApp) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseInstalledApp) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseInstalledApp) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseInstalledApp{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.AdditionalConfiguration != nil {
		s.AdditionalConfiguration = o.AdditionalConfiguration
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AppIdentifier != nil {
		s.AppIdentifier = o.AppIdentifier
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CategoryID != nil {
		s.CategoryID = o.CategoryID
	}
	if o.CheckPublicEndpoint != nil {
		s.CheckPublicEndpoint = o.CheckPublicEndpoint
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.CurrentVersion != nil {
		s.CurrentVersion = o.CurrentVersion
	}
	if o.DeploymentCount != nil {
		s.DeploymentCount = o.DeploymentCount
	}
	if o.ExternalWindowButton != nil {
		s.ExternalWindowButton = o.ExternalWindowButton
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Parameters != nil {
		s.Parameters = o.Parameters
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Status != nil {
		s.Status = o.Status
	}
	if o.StatusMessage != nil {
		s.StatusMessage = o.StatusMessage
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
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
func (o *SparseInstalledApp) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseInstalledApp{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.AdditionalConfiguration != nil {
		o.AdditionalConfiguration = s.AdditionalConfiguration
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AppIdentifier != nil {
		o.AppIdentifier = s.AppIdentifier
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CategoryID != nil {
		o.CategoryID = s.CategoryID
	}
	if s.CheckPublicEndpoint != nil {
		o.CheckPublicEndpoint = s.CheckPublicEndpoint
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.CurrentVersion != nil {
		o.CurrentVersion = s.CurrentVersion
	}
	if s.DeploymentCount != nil {
		o.DeploymentCount = s.DeploymentCount
	}
	if s.ExternalWindowButton != nil {
		o.ExternalWindowButton = s.ExternalWindowButton
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Parameters != nil {
		o.Parameters = s.Parameters
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Status != nil {
		o.Status = s.Status
	}
	if s.StatusMessage != nil {
		o.StatusMessage = s.StatusMessage
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
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
func (o *SparseInstalledApp) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseInstalledApp) ToPlain() elemental.PlainIdentifiable {

	out := NewInstalledApp()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AdditionalConfiguration != nil {
		out.AdditionalConfiguration = *o.AdditionalConfiguration
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AppIdentifier != nil {
		out.AppIdentifier = *o.AppIdentifier
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CategoryID != nil {
		out.CategoryID = *o.CategoryID
	}
	if o.CheckPublicEndpoint != nil {
		out.CheckPublicEndpoint = *o.CheckPublicEndpoint
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.CurrentVersion != nil {
		out.CurrentVersion = *o.CurrentVersion
	}
	if o.DeploymentCount != nil {
		out.DeploymentCount = *o.DeploymentCount
	}
	if o.ExternalWindowButton != nil {
		out.ExternalWindowButton = *o.ExternalWindowButton
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.StatusMessage != nil {
		out.StatusMessage = *o.StatusMessage
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
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

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseInstalledApp) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseInstalledApp) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseInstalledApp) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseInstalledApp) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseInstalledApp) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseInstalledApp) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseInstalledApp) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseInstalledApp) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseInstalledApp) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseInstalledApp) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseInstalledApp) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseInstalledApp) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseInstalledApp) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseInstalledApp) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseInstalledApp.
func (o *SparseInstalledApp) DeepCopy() *SparseInstalledApp {

	if o == nil {
		return nil
	}

	out := &SparseInstalledApp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseInstalledApp.
func (o *SparseInstalledApp) DeepCopyInto(out *SparseInstalledApp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseInstalledApp: %s", err))
	}

	*out = *target.(*SparseInstalledApp)
}

type mongoAttributesInstalledApp struct {
	ID                      bson.ObjectId           `bson:"_id,omitempty"`
	AdditionalConfiguration bool                    `bson:"additionalconfiguration"`
	Annotations             map[string][]string     `bson:"annotations"`
	AppIdentifier           string                  `bson:"appidentifier"`
	AssociatedTags          []string                `bson:"associatedtags"`
	CategoryID              string                  `bson:"categoryid"`
	CheckPublicEndpoint     bool                    `bson:"checkpublicendpoint"`
	CreateIdempotencyKey    string                  `bson:"createidempotencykey"`
	CreateTime              time.Time               `bson:"createtime"`
	CurrentVersion          string                  `bson:"currentversion"`
	DeploymentCount         int                     `bson:"deploymentcount"`
	ExternalWindowButton    map[string]string       `bson:"externalwindowbutton"`
	MigrationsLog           map[string]string       `bson:"migrationslog,omitempty"`
	Name                    string                  `bson:"name"`
	Namespace               string                  `bson:"namespace"`
	NormalizedTags          []string                `bson:"normalizedtags"`
	Parameters              map[string]interface{}  `bson:"parameters"`
	Protected               bool                    `bson:"protected"`
	Status                  InstalledAppStatusValue `bson:"status"`
	StatusMessage           string                  `bson:"statusmessage"`
	UpdateIdempotencyKey    string                  `bson:"updateidempotencykey"`
	UpdateTime              time.Time               `bson:"updatetime"`
	ZHash                   int                     `bson:"zhash"`
	Zone                    int                     `bson:"zone"`
}
type mongoAttributesSparseInstalledApp struct {
	ID                      bson.ObjectId            `bson:"_id,omitempty"`
	AdditionalConfiguration *bool                    `bson:"additionalconfiguration,omitempty"`
	Annotations             *map[string][]string     `bson:"annotations,omitempty"`
	AppIdentifier           *string                  `bson:"appidentifier,omitempty"`
	AssociatedTags          *[]string                `bson:"associatedtags,omitempty"`
	CategoryID              *string                  `bson:"categoryid,omitempty"`
	CheckPublicEndpoint     *bool                    `bson:"checkpublicendpoint,omitempty"`
	CreateIdempotencyKey    *string                  `bson:"createidempotencykey,omitempty"`
	CreateTime              *time.Time               `bson:"createtime,omitempty"`
	CurrentVersion          *string                  `bson:"currentversion,omitempty"`
	DeploymentCount         *int                     `bson:"deploymentcount,omitempty"`
	ExternalWindowButton    *map[string]string       `bson:"externalwindowbutton,omitempty"`
	MigrationsLog           *map[string]string       `bson:"migrationslog,omitempty"`
	Name                    *string                  `bson:"name,omitempty"`
	Namespace               *string                  `bson:"namespace,omitempty"`
	NormalizedTags          *[]string                `bson:"normalizedtags,omitempty"`
	Parameters              *map[string]interface{}  `bson:"parameters,omitempty"`
	Protected               *bool                    `bson:"protected,omitempty"`
	Status                  *InstalledAppStatusValue `bson:"status,omitempty"`
	StatusMessage           *string                  `bson:"statusmessage,omitempty"`
	UpdateIdempotencyKey    *string                  `bson:"updateidempotencykey,omitempty"`
	UpdateTime              *time.Time               `bson:"updatetime,omitempty"`
	ZHash                   *int                     `bson:"zhash,omitempty"`
	Zone                    *int                     `bson:"zone,omitempty"`
}
