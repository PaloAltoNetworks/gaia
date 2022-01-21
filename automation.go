package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AutomationTriggerValue represents the possible values for attribute "trigger".
type AutomationTriggerValue string

const (
	// AutomationTriggerEvent represents the value Event.
	AutomationTriggerEvent AutomationTriggerValue = "Event"

	// AutomationTriggerRemoteCall represents the value RemoteCall.
	AutomationTriggerRemoteCall AutomationTriggerValue = "RemoteCall"

	// AutomationTriggerTime represents the value Time.
	AutomationTriggerTime AutomationTriggerValue = "Time"

	// AutomationTriggerWebhook represents the value Webhook.
	AutomationTriggerWebhook AutomationTriggerValue = "Webhook"
)

// AutomationIdentity represents the Identity of the object.
var AutomationIdentity = elemental.Identity{
	Name:     "automation",
	Category: "automations",
	Package:  "sephiroth",
	Private:  false,
}

// AutomationsList represents a list of Automations
type AutomationsList []*Automation

// Identity returns the identity of the objects in the list.
func (o AutomationsList) Identity() elemental.Identity {

	return AutomationIdentity
}

// Copy returns a pointer to a copy the AutomationsList.
func (o AutomationsList) Copy() elemental.Identifiables {

	copy := append(AutomationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationsList.
func (o AutomationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AutomationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Automation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AutomationsList converted to SparseAutomationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AutomationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAutomationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAutomation)
	}

	return out
}

// Version returns the version of the content.
func (o AutomationsList) Version() int {

	return 1
}

// Automation represents the model of a automation
type Automation struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Contains the code that will be executed if the condition is met.
	Actions []string `json:"actions" msgpack:"actions" bson:"actions" mapstructure:"actions,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Contains the Microsegmentation token used by the automation's HTTP client. This
	// token is derived from the automation's app credential attribute.
	AporetoToken string `json:"-" msgpack:"-" bson:"aporetotoken" mapstructure:"-,omitempty"`

	// Contains the app credential associated with the automation which has its roles
	// deduced from the automation's entitlements.
	AppCredential string `json:"-" msgpack:"-" bson:"appcredential" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Condition contains the code that will be executed to decide if any action(s)
	// should be executed. Providing a condition for an automation with a
	// "Webhook" trigger type will have no impact as the condition will not be
	// evaluated. If no condition is defined, then the automation action(s) will be
	// executed; this behaves akin to a condition that always succeeds.
	Condition string `json:"condition" msgpack:"condition" bson:"condition" mapstructure:"condition,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Declares which operations are allowed on which identities.
	Entitlements map[string][]elemental.Operation `json:"entitlements" msgpack:"entitlements" bson:"entitlements" mapstructure:"entitlements,omitempty"`

	// Contains the error of the last run.
	Errors []string `json:"errors" msgpack:"errors" bson:"errors" mapstructure:"errors,omitempty"`

	// Contains the identity and operation an event must have to trigger the
	// automation.
	Events map[string][]elemental.EventType `json:"events" msgpack:"events" bson:"events" mapstructure:"events,omitempty"`

	// If set and the trigger is of type Time, the automation will be run at create or
	// update before being scheduled.
	ImmediateExecution bool `json:"immediateExecution" msgpack:"immediateExecution" bson:"immediateexecution" mapstructure:"immediateExecution,omitempty"`

	// The last successful execution tine.
	LastExecTime time.Time `json:"lastExecTime" msgpack:"lastExecTime" bson:"lastexectime" mapstructure:"lastExecTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Specifies when to run the automation. Must be in valid CRON format. This
	// only applies if the trigger is set to `Time`.
	Schedule string `json:"schedule" msgpack:"schedule" bson:"schedule" mapstructure:"schedule,omitempty"`

	// Signature to validate the authenticity of the object.
	Signature string `json:"signature" msgpack:"signature" bson:"signature" mapstructure:"signature,omitempty"`

	// Contains the standard output of the last run.
	Stdout string `json:"stdout" msgpack:"stdout" bson:"stdout" mapstructure:"stdout,omitempty"`

	// Holds the unique access token used as a password to trigger the
	// authentication. It will be visible only after creation.
	Token string `json:"token" msgpack:"token" bson:"token" mapstructure:"token,omitempty"`

	// If set to `true` a new token will be issued and the previous one invalidated.
	TokenRenew bool `json:"tokenRenew" msgpack:"tokenRenew" bson:"-" mapstructure:"tokenRenew,omitempty"`

	// Controls when the automation should be triggered.
	Trigger AutomationTriggerValue `json:"trigger" msgpack:"trigger" bson:"trigger" mapstructure:"trigger,omitempty"`

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

// NewAutomation returns a new *Automation
func NewAutomation() *Automation {

	return &Automation{
		ModelVersion:   1,
		Actions:        []string{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Entitlements:   map[string][]elemental.Operation{},
		Errors:         []string{},
		Events:         map[string][]elemental.EventType{},
		MigrationsLog:  map[string]string{},
		Parameters:     map[string]interface{}{},
		NormalizedTags: []string{},
		Trigger:        AutomationTriggerTime,
	}
}

// Identity returns the Identity of the object.
func (o *Automation) Identity() elemental.Identity {

	return AutomationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Automation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Automation) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Automation) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAutomation{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Actions = o.Actions
	s.Annotations = o.Annotations
	s.AporetoToken = o.AporetoToken
	s.AppCredential = o.AppCredential
	s.AssociatedTags = o.AssociatedTags
	s.Condition = o.Condition
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.Entitlements = o.Entitlements
	s.Errors = o.Errors
	s.Events = o.Events
	s.ImmediateExecution = o.ImmediateExecution
	s.LastExecTime = o.LastExecTime
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Parameters = o.Parameters
	s.Protected = o.Protected
	s.Schedule = o.Schedule
	s.Signature = o.Signature
	s.Stdout = o.Stdout
	s.Token = o.Token
	s.Trigger = o.Trigger
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Automation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAutomation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Actions = s.Actions
	o.Annotations = s.Annotations
	o.AporetoToken = s.AporetoToken
	o.AppCredential = s.AppCredential
	o.AssociatedTags = s.AssociatedTags
	o.Condition = s.Condition
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.Entitlements = s.Entitlements
	o.Errors = s.Errors
	o.Events = s.Events
	o.ImmediateExecution = s.ImmediateExecution
	o.LastExecTime = s.LastExecTime
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Parameters = s.Parameters
	o.Protected = s.Protected
	o.Schedule = s.Schedule
	o.Signature = s.Signature
	o.Stdout = s.Stdout
	o.Token = s.Token
	o.Trigger = s.Trigger
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Automation) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Automation) BleveType() string {

	return "automation"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Automation) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Automation) Doc() string {

	return `Allows you to define some JavaScript code and specify the conditions under which
it should
be executed.`
}

func (o *Automation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Automation) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Automation) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Automation) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Automation) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Automation) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Automation) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Automation) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Automation) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Automation) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Automation) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *Automation) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *Automation) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Automation) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Automation) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *Automation) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Automation) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Automation) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Automation) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Automation) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Automation) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Automation) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Automation) SetProtected(protected bool) {

	o.Protected = protected
}

// GetSignature returns the Signature of the receiver.
func (o *Automation) GetSignature() string {

	return o.Signature
}

// SetSignature sets the property Signature of the receiver using the given value.
func (o *Automation) SetSignature(signature string) {

	o.Signature = signature
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Automation) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Automation) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Automation) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Automation) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Automation) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Automation) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Automation) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Automation) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Automation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAutomation{
			ID:                   &o.ID,
			Actions:              &o.Actions,
			Annotations:          &o.Annotations,
			AporetoToken:         &o.AporetoToken,
			AppCredential:        &o.AppCredential,
			AssociatedTags:       &o.AssociatedTags,
			Condition:            &o.Condition,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			Entitlements:         &o.Entitlements,
			Errors:               &o.Errors,
			Events:               &o.Events,
			ImmediateExecution:   &o.ImmediateExecution,
			LastExecTime:         &o.LastExecTime,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Parameters:           &o.Parameters,
			Protected:            &o.Protected,
			Schedule:             &o.Schedule,
			Signature:            &o.Signature,
			Stdout:               &o.Stdout,
			Token:                &o.Token,
			TokenRenew:           &o.TokenRenew,
			Trigger:              &o.Trigger,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseAutomation{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "actions":
			sp.Actions = &(o.Actions)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "aporetoToken":
			sp.AporetoToken = &(o.AporetoToken)
		case "appCredential":
			sp.AppCredential = &(o.AppCredential)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "condition":
			sp.Condition = &(o.Condition)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "entitlements":
			sp.Entitlements = &(o.Entitlements)
		case "errors":
			sp.Errors = &(o.Errors)
		case "events":
			sp.Events = &(o.Events)
		case "immediateExecution":
			sp.ImmediateExecution = &(o.ImmediateExecution)
		case "lastExecTime":
			sp.LastExecTime = &(o.LastExecTime)
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
		case "schedule":
			sp.Schedule = &(o.Schedule)
		case "signature":
			sp.Signature = &(o.Signature)
		case "stdout":
			sp.Stdout = &(o.Stdout)
		case "token":
			sp.Token = &(o.Token)
		case "tokenRenew":
			sp.TokenRenew = &(o.TokenRenew)
		case "trigger":
			sp.Trigger = &(o.Trigger)
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

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *Automation) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.AporetoToken, err = encrypter.EncryptString(o.AporetoToken); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'AporetoToken' for 'Automation' (%s): %s", o.Identifier(), err)
	}
	if o.AppCredential, err = encrypter.EncryptString(o.AppCredential); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'AppCredential' for 'Automation' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *Automation) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.AporetoToken, err = encrypter.DecryptString(o.AporetoToken); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'AporetoToken' for 'Automation' (%s): %s", o.Identifier(), err)
	}
	if o.AppCredential, err = encrypter.DecryptString(o.AppCredential); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'AppCredential' for 'Automation' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// Patch apply the non nil value of a *SparseAutomation to the object.
func (o *Automation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAutomation)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Actions != nil {
		o.Actions = *so.Actions
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AporetoToken != nil {
		o.AporetoToken = *so.AporetoToken
	}
	if so.AppCredential != nil {
		o.AppCredential = *so.AppCredential
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Condition != nil {
		o.Condition = *so.Condition
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Entitlements != nil {
		o.Entitlements = *so.Entitlements
	}
	if so.Errors != nil {
		o.Errors = *so.Errors
	}
	if so.Events != nil {
		o.Events = *so.Events
	}
	if so.ImmediateExecution != nil {
		o.ImmediateExecution = *so.ImmediateExecution
	}
	if so.LastExecTime != nil {
		o.LastExecTime = *so.LastExecTime
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
	if so.Schedule != nil {
		o.Schedule = *so.Schedule
	}
	if so.Signature != nil {
		o.Signature = *so.Signature
	}
	if so.Stdout != nil {
		o.Stdout = *so.Stdout
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.TokenRenew != nil {
		o.TokenRenew = *so.TokenRenew
	}
	if so.Trigger != nil {
		o.Trigger = *so.Trigger
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

// DeepCopy returns a deep copy if the Automation.
func (o *Automation) DeepCopy() *Automation {

	if o == nil {
		return nil
	}

	out := &Automation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Automation.
func (o *Automation) DeepCopyInto(out *Automation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Automation: %s", err))
	}

	*out = *target.(*Automation)
}

// Validate valides the current information stored into the structure.
func (o *Automation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("trigger", string(o.Trigger), []string{"Event", "RemoteCall", "Webhook", "Time"}, false); err != nil {
		errors = errors.Append(err)
	}

	// Custom object validation.
	if err := ValidateAutomation(o); err != nil {
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
func (*Automation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AutomationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AutomationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Automation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AutomationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Automation) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "actions":
		return o.Actions
	case "annotations":
		return o.Annotations
	case "aporetoToken":
		return o.AporetoToken
	case "appCredential":
		return o.AppCredential
	case "associatedTags":
		return o.AssociatedTags
	case "condition":
		return o.Condition
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "entitlements":
		return o.Entitlements
	case "errors":
		return o.Errors
	case "events":
		return o.Events
	case "immediateExecution":
		return o.ImmediateExecution
	case "lastExecTime":
		return o.LastExecTime
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
	case "schedule":
		return o.Schedule
	case "signature":
		return o.Signature
	case "stdout":
		return o.Stdout
	case "token":
		return o.Token
	case "tokenRenew":
		return o.TokenRenew
	case "trigger":
		return o.Trigger
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

// AutomationAttributesMap represents the map of attribute for Automation.
var AutomationAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Actions": {
		AllowedChoices: []string{},
		ConvertedName:  "Actions",
		Description:    `Contains the code that will be executed if the condition is met.`,
		Exposed:        true,
		Name:           "actions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"AporetoToken": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AporetoToken",
		Description: `Contains the Microsegmentation token used by the automation's HTTP client. This
token is derived from the automation's app credential attribute.`,
		Encrypted: true,
		Name:      "aporetoToken",
		Stored:    true,
		Type:      "string",
	},
	"AppCredential": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AppCredential",
		Description: `Contains the app credential associated with the automation which has its roles
deduced from the automation's entitlements.`,
		Encrypted: true,
		Name:      "appCredential",
		Stored:    true,
		Type:      "string",
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
	"Condition": {
		AllowedChoices: []string{},
		ConvertedName:  "Condition",
		Description: `Condition contains the code that will be executed to decide if any action(s)
should be executed. Providing a condition for an automation with a
"Webhook" trigger type will have no impact as the condition will not be
evaluated. If no condition is defined, then the automation action(s) will be
executed; this behaves akin to a condition that always succeeds.`,
		Exposed: true,
		Name:    "condition",
		Stored:  true,
		Type:    "string",
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
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": {
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Entitlements": {
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Declares which operations are allowed on which identities.`,
		Exposed:        true,
		Name:           "entitlements",
		Stored:         true,
		SubType:        "_automation_entitlements",
		Type:           "external",
	},
	"Errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `Contains the error of the last run.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Events": {
		AllowedChoices: []string{},
		ConvertedName:  "Events",
		Description: `Contains the identity and operation an event must have to trigger the
automation.`,
		Exposed: true,
		Name:    "events",
		Stored:  true,
		SubType: "_automation_events",
		Type:    "external",
	},
	"ImmediateExecution": {
		AllowedChoices: []string{},
		ConvertedName:  "ImmediateExecution",
		Description: `If set and the trigger is of type Time, the automation will be run at create or
update before being scheduled.`,
		Exposed: true,
		Name:    "immediateExecution",
		Stored:  true,
		Type:    "boolean",
	},
	"LastExecTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastExecTime",
		Description:    `The last successful execution tine.`,
		Exposed:        true,
		Name:           "lastExecTime",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
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
		Description:    `Contains the computed parameters.`,
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
	"Schedule": {
		AllowedChoices: []string{},
		ConvertedName:  "Schedule",
		Description: `Specifies when to run the automation. Must be in valid CRON format. This
only applies if the trigger is set to ` + "`" + `Time` + "`" + `.`,
		Exposed: true,
		Name:    "schedule",
		Stored:  true,
		Type:    "string",
	},
	"Signature": {
		AllowedChoices: []string{},
		ConvertedName:  "Signature",
		Description:    `Signature to validate the authenticity of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "signature",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Stdout": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Stdout",
		Description:    `Contains the standard output of the last run.`,
		Exposed:        true,
		Name:           "stdout",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description: `Holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.`,
		Exposed:   true,
		Name:      "token",
		Stored:    true,
		Transient: true,
		Type:      "string",
	},
	"TokenRenew": {
		AllowedChoices: []string{},
		ConvertedName:  "TokenRenew",
		Description:    `If set to ` + "`" + `true` + "`" + ` a new token will be issued and the previous one invalidated.`,
		Exposed:        true,
		Name:           "tokenRenew",
		Type:           "boolean",
	},
	"Trigger": {
		AllowedChoices: []string{"Event", "RemoteCall", "Webhook", "Time"},
		ConvertedName:  "Trigger",
		DefaultValue:   AutomationTriggerTime,
		Description:    `Controls when the automation should be triggered.`,
		Exposed:        true,
		Name:           "trigger",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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

// AutomationLowerCaseAttributesMap represents the map of attribute for Automation.
var AutomationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"actions": {
		AllowedChoices: []string{},
		ConvertedName:  "Actions",
		Description:    `Contains the code that will be executed if the condition is met.`,
		Exposed:        true,
		Name:           "actions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"aporetotoken": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AporetoToken",
		Description: `Contains the Microsegmentation token used by the automation's HTTP client. This
token is derived from the automation's app credential attribute.`,
		Encrypted: true,
		Name:      "aporetoToken",
		Stored:    true,
		Type:      "string",
	},
	"appcredential": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AppCredential",
		Description: `Contains the app credential associated with the automation which has its roles
deduced from the automation's entitlements.`,
		Encrypted: true,
		Name:      "appCredential",
		Stored:    true,
		Type:      "string",
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
	"condition": {
		AllowedChoices: []string{},
		ConvertedName:  "Condition",
		Description: `Condition contains the code that will be executed to decide if any action(s)
should be executed. Providing a condition for an automation with a
"Webhook" trigger type will have no impact as the condition will not be
evaluated. If no condition is defined, then the automation action(s) will be
executed; this behaves akin to a condition that always succeeds.`,
		Exposed: true,
		Name:    "condition",
		Stored:  true,
		Type:    "string",
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
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": {
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"entitlements": {
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Declares which operations are allowed on which identities.`,
		Exposed:        true,
		Name:           "entitlements",
		Stored:         true,
		SubType:        "_automation_entitlements",
		Type:           "external",
	},
	"errors": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `Contains the error of the last run.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"events": {
		AllowedChoices: []string{},
		ConvertedName:  "Events",
		Description: `Contains the identity and operation an event must have to trigger the
automation.`,
		Exposed: true,
		Name:    "events",
		Stored:  true,
		SubType: "_automation_events",
		Type:    "external",
	},
	"immediateexecution": {
		AllowedChoices: []string{},
		ConvertedName:  "ImmediateExecution",
		Description: `If set and the trigger is of type Time, the automation will be run at create or
update before being scheduled.`,
		Exposed: true,
		Name:    "immediateExecution",
		Stored:  true,
		Type:    "boolean",
	},
	"lastexectime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastExecTime",
		Description:    `The last successful execution tine.`,
		Exposed:        true,
		Name:           "lastExecTime",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
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
		Description:    `Contains the computed parameters.`,
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
	"schedule": {
		AllowedChoices: []string{},
		ConvertedName:  "Schedule",
		Description: `Specifies when to run the automation. Must be in valid CRON format. This
only applies if the trigger is set to ` + "`" + `Time` + "`" + `.`,
		Exposed: true,
		Name:    "schedule",
		Stored:  true,
		Type:    "string",
	},
	"signature": {
		AllowedChoices: []string{},
		ConvertedName:  "Signature",
		Description:    `Signature to validate the authenticity of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "signature",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"stdout": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Stdout",
		Description:    `Contains the standard output of the last run.`,
		Exposed:        true,
		Name:           "stdout",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"token": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description: `Holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.`,
		Exposed:   true,
		Name:      "token",
		Stored:    true,
		Transient: true,
		Type:      "string",
	},
	"tokenrenew": {
		AllowedChoices: []string{},
		ConvertedName:  "TokenRenew",
		Description:    `If set to ` + "`" + `true` + "`" + ` a new token will be issued and the previous one invalidated.`,
		Exposed:        true,
		Name:           "tokenRenew",
		Type:           "boolean",
	},
	"trigger": {
		AllowedChoices: []string{"Event", "RemoteCall", "Webhook", "Time"},
		ConvertedName:  "Trigger",
		DefaultValue:   AutomationTriggerTime,
		Description:    `Controls when the automation should be triggered.`,
		Exposed:        true,
		Name:           "trigger",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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

// SparseAutomationsList represents a list of SparseAutomations
type SparseAutomationsList []*SparseAutomation

// Identity returns the identity of the objects in the list.
func (o SparseAutomationsList) Identity() elemental.Identity {

	return AutomationIdentity
}

// Copy returns a pointer to a copy the SparseAutomationsList.
func (o SparseAutomationsList) Copy() elemental.Identifiables {

	copy := append(SparseAutomationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAutomationsList.
func (o SparseAutomationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAutomationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAutomation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAutomationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAutomationsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAutomationsList converted to AutomationsList.
func (o SparseAutomationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAutomationsList) Version() int {

	return 1
}

// SparseAutomation represents the sparse version of a automation.
type SparseAutomation struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Contains the code that will be executed if the condition is met.
	Actions *[]string `json:"actions,omitempty" msgpack:"actions,omitempty" bson:"actions,omitempty" mapstructure:"actions,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Contains the Microsegmentation token used by the automation's HTTP client. This
	// token is derived from the automation's app credential attribute.
	AporetoToken *string `json:"-" msgpack:"-" bson:"aporetotoken,omitempty" mapstructure:"-,omitempty"`

	// Contains the app credential associated with the automation which has its roles
	// deduced from the automation's entitlements.
	AppCredential *string `json:"-" msgpack:"-" bson:"appcredential,omitempty" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Condition contains the code that will be executed to decide if any action(s)
	// should be executed. Providing a condition for an automation with a
	// "Webhook" trigger type will have no impact as the condition will not be
	// evaluated. If no condition is defined, then the automation action(s) will be
	// executed; this behaves akin to a condition that always succeeds.
	Condition *string `json:"condition,omitempty" msgpack:"condition,omitempty" bson:"condition,omitempty" mapstructure:"condition,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Declares which operations are allowed on which identities.
	Entitlements *map[string][]elemental.Operation `json:"entitlements,omitempty" msgpack:"entitlements,omitempty" bson:"entitlements,omitempty" mapstructure:"entitlements,omitempty"`

	// Contains the error of the last run.
	Errors *[]string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"errors,omitempty" mapstructure:"errors,omitempty"`

	// Contains the identity and operation an event must have to trigger the
	// automation.
	Events *map[string][]elemental.EventType `json:"events,omitempty" msgpack:"events,omitempty" bson:"events,omitempty" mapstructure:"events,omitempty"`

	// If set and the trigger is of type Time, the automation will be run at create or
	// update before being scheduled.
	ImmediateExecution *bool `json:"immediateExecution,omitempty" msgpack:"immediateExecution,omitempty" bson:"immediateexecution,omitempty" mapstructure:"immediateExecution,omitempty"`

	// The last successful execution tine.
	LastExecTime *time.Time `json:"lastExecTime,omitempty" msgpack:"lastExecTime,omitempty" bson:"lastexectime,omitempty" mapstructure:"lastExecTime,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Specifies when to run the automation. Must be in valid CRON format. This
	// only applies if the trigger is set to `Time`.
	Schedule *string `json:"schedule,omitempty" msgpack:"schedule,omitempty" bson:"schedule,omitempty" mapstructure:"schedule,omitempty"`

	// Signature to validate the authenticity of the object.
	Signature *string `json:"signature,omitempty" msgpack:"signature,omitempty" bson:"signature,omitempty" mapstructure:"signature,omitempty"`

	// Contains the standard output of the last run.
	Stdout *string `json:"stdout,omitempty" msgpack:"stdout,omitempty" bson:"stdout,omitempty" mapstructure:"stdout,omitempty"`

	// Holds the unique access token used as a password to trigger the
	// authentication. It will be visible only after creation.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"token,omitempty" mapstructure:"token,omitempty"`

	// If set to `true` a new token will be issued and the previous one invalidated.
	TokenRenew *bool `json:"tokenRenew,omitempty" msgpack:"tokenRenew,omitempty" bson:"-" mapstructure:"tokenRenew,omitempty"`

	// Controls when the automation should be triggered.
	Trigger *AutomationTriggerValue `json:"trigger,omitempty" msgpack:"trigger,omitempty" bson:"trigger,omitempty" mapstructure:"trigger,omitempty"`

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

// NewSparseAutomation returns a new  SparseAutomation.
func NewSparseAutomation() *SparseAutomation {
	return &SparseAutomation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAutomation) Identity() elemental.Identity {

	return AutomationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAutomation) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAutomation) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAutomation) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAutomation{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Actions != nil {
		s.Actions = o.Actions
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AporetoToken != nil {
		s.AporetoToken = o.AporetoToken
	}
	if o.AppCredential != nil {
		s.AppCredential = o.AppCredential
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.Condition != nil {
		s.Condition = o.Condition
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.Entitlements != nil {
		s.Entitlements = o.Entitlements
	}
	if o.Errors != nil {
		s.Errors = o.Errors
	}
	if o.Events != nil {
		s.Events = o.Events
	}
	if o.ImmediateExecution != nil {
		s.ImmediateExecution = o.ImmediateExecution
	}
	if o.LastExecTime != nil {
		s.LastExecTime = o.LastExecTime
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
	if o.Schedule != nil {
		s.Schedule = o.Schedule
	}
	if o.Signature != nil {
		s.Signature = o.Signature
	}
	if o.Stdout != nil {
		s.Stdout = o.Stdout
	}
	if o.Token != nil {
		s.Token = o.Token
	}
	if o.Trigger != nil {
		s.Trigger = o.Trigger
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
func (o *SparseAutomation) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAutomation{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Actions != nil {
		o.Actions = s.Actions
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AporetoToken != nil {
		o.AporetoToken = s.AporetoToken
	}
	if s.AppCredential != nil {
		o.AppCredential = s.AppCredential
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.Condition != nil {
		o.Condition = s.Condition
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.Entitlements != nil {
		o.Entitlements = s.Entitlements
	}
	if s.Errors != nil {
		o.Errors = s.Errors
	}
	if s.Events != nil {
		o.Events = s.Events
	}
	if s.ImmediateExecution != nil {
		o.ImmediateExecution = s.ImmediateExecution
	}
	if s.LastExecTime != nil {
		o.LastExecTime = s.LastExecTime
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
	if s.Schedule != nil {
		o.Schedule = s.Schedule
	}
	if s.Signature != nil {
		o.Signature = s.Signature
	}
	if s.Stdout != nil {
		o.Stdout = s.Stdout
	}
	if s.Token != nil {
		o.Token = s.Token
	}
	if s.Trigger != nil {
		o.Trigger = s.Trigger
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
func (o *SparseAutomation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAutomation) ToPlain() elemental.PlainIdentifiable {

	out := NewAutomation()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Actions != nil {
		out.Actions = *o.Actions
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AporetoToken != nil {
		out.AporetoToken = *o.AporetoToken
	}
	if o.AppCredential != nil {
		out.AppCredential = *o.AppCredential
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Condition != nil {
		out.Condition = *o.Condition
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Entitlements != nil {
		out.Entitlements = *o.Entitlements
	}
	if o.Errors != nil {
		out.Errors = *o.Errors
	}
	if o.Events != nil {
		out.Events = *o.Events
	}
	if o.ImmediateExecution != nil {
		out.ImmediateExecution = *o.ImmediateExecution
	}
	if o.LastExecTime != nil {
		out.LastExecTime = *o.LastExecTime
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
	if o.Schedule != nil {
		out.Schedule = *o.Schedule
	}
	if o.Signature != nil {
		out.Signature = *o.Signature
	}
	if o.Stdout != nil {
		out.Stdout = *o.Stdout
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.TokenRenew != nil {
		out.TokenRenew = *o.TokenRenew
	}
	if o.Trigger != nil {
		out.Trigger = *o.Trigger
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

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseAutomation) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.AporetoToken, err = encrypter.EncryptString(*o.AporetoToken); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'AporetoToken' for 'SparseAutomation' (%s): %s", o.Identifier(), err)
	}
	if *o.AppCredential, err = encrypter.EncryptString(*o.AppCredential); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'AppCredential' for 'SparseAutomation' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseAutomation) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.AporetoToken, err = encrypter.DecryptString(*o.AporetoToken); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'AporetoToken' for 'SparseAutomation' (%s): %s", o.Identifier(), err)
	}
	if *o.AppCredential, err = encrypter.DecryptString(*o.AppCredential); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'AppCredential' for 'SparseAutomation' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAutomation) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAutomation) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAutomation) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAutomation) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseAutomation) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAutomation) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAutomation) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAutomation) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAutomation) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAutomation) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseAutomation) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseAutomation) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAutomation) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAutomation) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseAutomation) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAutomation) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAutomation) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAutomation) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAutomation) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAutomation) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAutomation) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseAutomation) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetSignature returns the Signature of the receiver.
func (o *SparseAutomation) GetSignature() (out string) {

	if o.Signature == nil {
		return
	}

	return *o.Signature
}

// SetSignature sets the property Signature of the receiver using the address of the given value.
func (o *SparseAutomation) SetSignature(signature string) {

	o.Signature = &signature
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseAutomation) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAutomation) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAutomation) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAutomation) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAutomation) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAutomation) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAutomation) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAutomation) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAutomation.
func (o *SparseAutomation) DeepCopy() *SparseAutomation {

	if o == nil {
		return nil
	}

	out := &SparseAutomation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAutomation.
func (o *SparseAutomation) DeepCopyInto(out *SparseAutomation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAutomation: %s", err))
	}

	*out = *target.(*SparseAutomation)
}

type mongoAttributesAutomation struct {
	ID                   bson.ObjectId                    `bson:"_id,omitempty"`
	Actions              []string                         `bson:"actions"`
	Annotations          map[string][]string              `bson:"annotations"`
	AporetoToken         string                           `bson:"aporetotoken"`
	AppCredential        string                           `bson:"appcredential"`
	AssociatedTags       []string                         `bson:"associatedtags"`
	Condition            string                           `bson:"condition"`
	CreateIdempotencyKey string                           `bson:"createidempotencykey"`
	CreateTime           time.Time                        `bson:"createtime"`
	Description          string                           `bson:"description"`
	Disabled             bool                             `bson:"disabled"`
	Entitlements         map[string][]elemental.Operation `bson:"entitlements"`
	Errors               []string                         `bson:"errors"`
	Events               map[string][]elemental.EventType `bson:"events"`
	ImmediateExecution   bool                             `bson:"immediateexecution"`
	LastExecTime         time.Time                        `bson:"lastexectime"`
	MigrationsLog        map[string]string                `bson:"migrationslog,omitempty"`
	Name                 string                           `bson:"name"`
	Namespace            string                           `bson:"namespace"`
	NormalizedTags       []string                         `bson:"normalizedtags"`
	Parameters           map[string]interface{}           `bson:"parameters"`
	Protected            bool                             `bson:"protected"`
	Schedule             string                           `bson:"schedule"`
	Signature            string                           `bson:"signature"`
	Stdout               string                           `bson:"stdout"`
	Token                string                           `bson:"token"`
	Trigger              AutomationTriggerValue           `bson:"trigger"`
	UpdateIdempotencyKey string                           `bson:"updateidempotencykey"`
	UpdateTime           time.Time                        `bson:"updatetime"`
	ZHash                int                              `bson:"zhash"`
	Zone                 int                              `bson:"zone"`
}
type mongoAttributesSparseAutomation struct {
	ID                   bson.ObjectId                     `bson:"_id,omitempty"`
	Actions              *[]string                         `bson:"actions,omitempty"`
	Annotations          *map[string][]string              `bson:"annotations,omitempty"`
	AporetoToken         *string                           `bson:"aporetotoken,omitempty"`
	AppCredential        *string                           `bson:"appcredential,omitempty"`
	AssociatedTags       *[]string                         `bson:"associatedtags,omitempty"`
	Condition            *string                           `bson:"condition,omitempty"`
	CreateIdempotencyKey *string                           `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                        `bson:"createtime,omitempty"`
	Description          *string                           `bson:"description,omitempty"`
	Disabled             *bool                             `bson:"disabled,omitempty"`
	Entitlements         *map[string][]elemental.Operation `bson:"entitlements,omitempty"`
	Errors               *[]string                         `bson:"errors,omitempty"`
	Events               *map[string][]elemental.EventType `bson:"events,omitempty"`
	ImmediateExecution   *bool                             `bson:"immediateexecution,omitempty"`
	LastExecTime         *time.Time                        `bson:"lastexectime,omitempty"`
	MigrationsLog        *map[string]string                `bson:"migrationslog,omitempty"`
	Name                 *string                           `bson:"name,omitempty"`
	Namespace            *string                           `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                         `bson:"normalizedtags,omitempty"`
	Parameters           *map[string]interface{}           `bson:"parameters,omitempty"`
	Protected            *bool                             `bson:"protected,omitempty"`
	Schedule             *string                           `bson:"schedule,omitempty"`
	Signature            *string                           `bson:"signature,omitempty"`
	Stdout               *string                           `bson:"stdout,omitempty"`
	Token                *string                           `bson:"token,omitempty"`
	Trigger              *AutomationTriggerValue           `bson:"trigger,omitempty"`
	UpdateIdempotencyKey *string                           `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                        `bson:"updatetime,omitempty"`
	ZHash                *int                              `bson:"zhash,omitempty"`
	Zone                 *int                              `bson:"zone,omitempty"`
}
