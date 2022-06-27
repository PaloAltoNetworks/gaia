package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PollAccountCloudTypeValue represents the possible values for attribute "cloudType".
type PollAccountCloudTypeValue string

const (
	// PollAccountCloudTypeAWS represents the value AWS.
	PollAccountCloudTypeAWS PollAccountCloudTypeValue = "AWS"

	// PollAccountCloudTypeGCP represents the value GCP.
	PollAccountCloudTypeGCP PollAccountCloudTypeValue = "GCP"
)

// PollAccountIdentity represents the Identity of the object.
var PollAccountIdentity = elemental.Identity{
	Name:     "pollaccount",
	Category: "pollaccounts",
	Package:  "pandemona",
	Private:  false,
}

// PollAccountsList represents a list of PollAccounts
type PollAccountsList []*PollAccount

// Identity returns the identity of the objects in the list.
func (o PollAccountsList) Identity() elemental.Identity {

	return PollAccountIdentity
}

// Copy returns a pointer to a copy the PollAccountsList.
func (o PollAccountsList) Copy() elemental.Identifiables {

	copy := append(PollAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PollAccountsList.
func (o PollAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PollAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PollAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PollAccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PollAccountsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PollAccountsList converted to SparsePollAccountsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PollAccountsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePollAccountsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePollAccount)
	}

	return out
}

// Version returns the version of the content.
func (o PollAccountsList) Version() int {

	return 1
}

// PollAccount represents the model of a pollaccount
type PollAccount struct {
	// The ID of the account.
	AccountID string `json:"accountID" msgpack:"accountID" bson:"-" mapstructure:"accountID,omitempty"`

	// The region to use for authorization.
	AuthorizationRegion string `json:"authorizationRegion" msgpack:"authorizationRegion" bson:"-" mapstructure:"authorizationRegion,omitempty"`

	// The cloud type for the account.
	CloudType PollAccountCloudTypeValue `json:"cloudType" msgpack:"cloudType" bson:"-" mapstructure:"cloudType,omitempty"`

	// The name of the account.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// The role that it should use to poll the account.
	Role string `json:"role" msgpack:"role" bson:"-" mapstructure:"role,omitempty"`

	// Limit polling to these regions only. If empty, all regions will be polled.
	TargetRegions []string `json:"targetRegions" msgpack:"targetRegions" bson:"-" mapstructure:"targetRegions,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPollAccount returns a new *PollAccount
func NewPollAccount() *PollAccount {

	return &PollAccount{
		ModelVersion:  1,
		CloudType:     PollAccountCloudTypeAWS,
		TargetRegions: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *PollAccount) Identity() elemental.Identity {

	return PollAccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PollAccount) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PollAccount) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PollAccount) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPollAccount{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PollAccount) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPollAccount{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PollAccount) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PollAccount) BleveType() string {

	return "pollaccount"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PollAccount) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PollAccount) Doc() string {

	return `Initiates a poll for a particular account. Data are stored in the current
namespace.`
}

func (o *PollAccount) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PollAccount) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePollAccount{
			AccountID:           &o.AccountID,
			AuthorizationRegion: &o.AuthorizationRegion,
			CloudType:           &o.CloudType,
			Name:                &o.Name,
			Role:                &o.Role,
			TargetRegions:       &o.TargetRegions,
		}
	}

	sp := &SparsePollAccount{}
	for _, f := range fields {
		switch f {
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "authorizationRegion":
			sp.AuthorizationRegion = &(o.AuthorizationRegion)
		case "cloudType":
			sp.CloudType = &(o.CloudType)
		case "name":
			sp.Name = &(o.Name)
		case "role":
			sp.Role = &(o.Role)
		case "targetRegions":
			sp.TargetRegions = &(o.TargetRegions)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePollAccount to the object.
func (o *PollAccount) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePollAccount)
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.AuthorizationRegion != nil {
		o.AuthorizationRegion = *so.AuthorizationRegion
	}
	if so.CloudType != nil {
		o.CloudType = *so.CloudType
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Role != nil {
		o.Role = *so.Role
	}
	if so.TargetRegions != nil {
		o.TargetRegions = *so.TargetRegions
	}
}

// DeepCopy returns a deep copy if the PollAccount.
func (o *PollAccount) DeepCopy() *PollAccount {

	if o == nil {
		return nil
	}

	out := &PollAccount{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PollAccount.
func (o *PollAccount) DeepCopyInto(out *PollAccount) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PollAccount: %s", err))
	}

	*out = *target.(*PollAccount)
}

// Validate valides the current information stored into the structure.
func (o *PollAccount) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("accountID", o.AccountID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("authorizationRegion", o.AuthorizationRegion); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("cloudType", string(o.CloudType), []string{"AWS", "GCP"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("role", o.Role); err != nil {
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
func (*PollAccount) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PollAccountAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PollAccountLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PollAccount) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PollAccountAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PollAccount) ValueForAttribute(name string) interface{} {

	switch name {
	case "accountID":
		return o.AccountID
	case "authorizationRegion":
		return o.AuthorizationRegion
	case "cloudType":
		return o.CloudType
	case "name":
		return o.Name
	case "role":
		return o.Role
	case "targetRegions":
		return o.TargetRegions
	}

	return nil
}

// PollAccountAttributesMap represents the map of attribute for PollAccount.
var PollAccountAttributesMap = map[string]elemental.AttributeSpecification{
	"AccountID": {
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `The ID of the account.`,
		Exposed:        true,
		Name:           "accountID",
		Required:       true,
		Type:           "string",
	},
	"AuthorizationRegion": {
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationRegion",
		Description:    `The region to use for authorization.`,
		Exposed:        true,
		Name:           "authorizationRegion",
		Required:       true,
		Type:           "string",
	},
	"CloudType": {
		AllowedChoices: []string{"AWS", "GCP"},
		ConvertedName:  "CloudType",
		DefaultValue:   PollAccountCloudTypeAWS,
		Description:    `The cloud type for the account.`,
		Exposed:        true,
		Name:           "cloudType",
		Type:           "enum",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of the account.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"Role": {
		AllowedChoices: []string{},
		ConvertedName:  "Role",
		Description:    `The role that it should use to poll the account.`,
		Exposed:        true,
		Name:           "role",
		Required:       true,
		Type:           "string",
	},
	"TargetRegions": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetRegions",
		Description:    `Limit polling to these regions only. If empty, all regions will be polled.`,
		Exposed:        true,
		Name:           "targetRegions",
		SubType:        "string",
		Type:           "list",
	},
}

// PollAccountLowerCaseAttributesMap represents the map of attribute for PollAccount.
var PollAccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"accountid": {
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `The ID of the account.`,
		Exposed:        true,
		Name:           "accountID",
		Required:       true,
		Type:           "string",
	},
	"authorizationregion": {
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationRegion",
		Description:    `The region to use for authorization.`,
		Exposed:        true,
		Name:           "authorizationRegion",
		Required:       true,
		Type:           "string",
	},
	"cloudtype": {
		AllowedChoices: []string{"AWS", "GCP"},
		ConvertedName:  "CloudType",
		DefaultValue:   PollAccountCloudTypeAWS,
		Description:    `The cloud type for the account.`,
		Exposed:        true,
		Name:           "cloudType",
		Type:           "enum",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of the account.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"role": {
		AllowedChoices: []string{},
		ConvertedName:  "Role",
		Description:    `The role that it should use to poll the account.`,
		Exposed:        true,
		Name:           "role",
		Required:       true,
		Type:           "string",
	},
	"targetregions": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetRegions",
		Description:    `Limit polling to these regions only. If empty, all regions will be polled.`,
		Exposed:        true,
		Name:           "targetRegions",
		SubType:        "string",
		Type:           "list",
	},
}

// SparsePollAccountsList represents a list of SparsePollAccounts
type SparsePollAccountsList []*SparsePollAccount

// Identity returns the identity of the objects in the list.
func (o SparsePollAccountsList) Identity() elemental.Identity {

	return PollAccountIdentity
}

// Copy returns a pointer to a copy the SparsePollAccountsList.
func (o SparsePollAccountsList) Copy() elemental.Identifiables {

	copy := append(SparsePollAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePollAccountsList.
func (o SparsePollAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePollAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePollAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePollAccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePollAccountsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePollAccountsList converted to PollAccountsList.
func (o SparsePollAccountsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePollAccountsList) Version() int {

	return 1
}

// SparsePollAccount represents the sparse version of a pollaccount.
type SparsePollAccount struct {
	// The ID of the account.
	AccountID *string `json:"accountID,omitempty" msgpack:"accountID,omitempty" bson:"-" mapstructure:"accountID,omitempty"`

	// The region to use for authorization.
	AuthorizationRegion *string `json:"authorizationRegion,omitempty" msgpack:"authorizationRegion,omitempty" bson:"-" mapstructure:"authorizationRegion,omitempty"`

	// The cloud type for the account.
	CloudType *PollAccountCloudTypeValue `json:"cloudType,omitempty" msgpack:"cloudType,omitempty" bson:"-" mapstructure:"cloudType,omitempty"`

	// The name of the account.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// The role that it should use to poll the account.
	Role *string `json:"role,omitempty" msgpack:"role,omitempty" bson:"-" mapstructure:"role,omitempty"`

	// Limit polling to these regions only. If empty, all regions will be polled.
	TargetRegions *[]string `json:"targetRegions,omitempty" msgpack:"targetRegions,omitempty" bson:"-" mapstructure:"targetRegions,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePollAccount returns a new  SparsePollAccount.
func NewSparsePollAccount() *SparsePollAccount {
	return &SparsePollAccount{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePollAccount) Identity() elemental.Identity {

	return PollAccountIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePollAccount) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePollAccount) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePollAccount) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePollAccount{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePollAccount) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePollAccount{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePollAccount) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePollAccount) ToPlain() elemental.PlainIdentifiable {

	out := NewPollAccount()
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.AuthorizationRegion != nil {
		out.AuthorizationRegion = *o.AuthorizationRegion
	}
	if o.CloudType != nil {
		out.CloudType = *o.CloudType
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Role != nil {
		out.Role = *o.Role
	}
	if o.TargetRegions != nil {
		out.TargetRegions = *o.TargetRegions
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePollAccount.
func (o *SparsePollAccount) DeepCopy() *SparsePollAccount {

	if o == nil {
		return nil
	}

	out := &SparsePollAccount{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePollAccount.
func (o *SparsePollAccount) DeepCopyInto(out *SparsePollAccount) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePollAccount: %s", err))
	}

	*out = *target.(*SparsePollAccount)
}

type mongoAttributesPollAccount struct {
}
type mongoAttributesSparsePollAccount struct {
}
