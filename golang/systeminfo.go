package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// SystemInfoStatusValue represents the possible values for attribute "status".
type SystemInfoStatusValue string

const (
	// SystemInfoStatusDegraded represents the value Degraded.
	SystemInfoStatusDegraded SystemInfoStatusValue = "Degraded"

	// SystemInfoStatusFailure represents the value Failure.
	SystemInfoStatusFailure SystemInfoStatusValue = "Failure"

	// SystemInfoStatusOk represents the value Ok.
	SystemInfoStatusOk SystemInfoStatusValue = "Ok"
)

// SystemInfoIdentity represents the Identity of the object
var SystemInfoIdentity = elemental.Identity{
	Name:     "systeminfo",
	Category: "systeminfos",
}

// SystemInfosList represents a list of SystemInfos
type SystemInfosList []*SystemInfo

// SystemInfo represents the model of a systeminfo
type SystemInfo struct {
	// APIVersion is the API version served by the server.
	APIVersion string `json:"APIVersion" cql:"-" bson:"-"`

	// BahamutVersion is the version of Bahamut used by the server.
	BahamutVersion string `json:"bahamutVersion" cql:"-" bson:"-"`

	// CertificateAuthority contains the main certificate authority,
	CertificateAuthority string `json:"certificateAuthority" cql:"-" bson:"-"`

	// ElementalVersion is the version of Elemental used by the server.
	ElementalVersion string `json:"elementalVersion" cql:"-" bson:"-"`

	// GaiaVersion is the version of Gaia used by the server.
	GaiaVersion string `json:"gaiaVersion" cql:"-" bson:"-"`

	// GoogleClientID is the Google oauth client ID to use to get a valid token from Google for Midgard.
	GoogleClientID string `json:"googleClientID" cql:"-" bson:"-"`

	// KairosDBURL contains the URL of the kairos db server.
	KairosDBURL string `json:"kairosDBURL" cql:"-" bson:"-"`

	// ManipulateVersion is the version of Manipulate used by the server.
	ManipulateVersion string `json:"manipulateVersion" cql:"-" bson:"-"`

	// MidgardURL contains the url to use to obtain a token.
	MidgardURL string `json:"midgardURL" cql:"-" bson:"-"`

	// PubsubService provides the end-point for the pubsub server.
	PubSubService string `json:"pubSubService" cql:"-" bson:"-"`

	// SquallVersion is the version of server.
	SquallVersion string `json:"squallVersion" cql:"-" bson:"-"`

	// Status is the overall health status of the server.
	Status SystemInfoStatusValue `json:"status" cql:"-" bson:"-"`

	// zackURL contains the URL of the Zack server.
	ZackURL string `json:"zackURL" cql:"-" bson:"-"`
}

// NewSystemInfo returns a new *SystemInfo
func NewSystemInfo() *SystemInfo {

	return &SystemInfo{}
}

// Identity returns the Identity of the object.
func (o *SystemInfo) Identity() elemental.Identity {

	return SystemInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SystemInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SystemInfo) SetIdentifier(ID string) {

}

func (o *SystemInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *SystemInfo) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Degraded", "Failure", "Ok"}, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("zackURL", o.ZackURL); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (SystemInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return SystemInfoAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (SystemInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return SystemInfoAttributesMap
}

// SystemInfoAttributesMap represents the map of attribute for SystemInfo.
var SystemInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"APIVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "string",
	},
	"BahamutVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "bahamutVersion",
		ReadOnly:       true,
		Type:           "string",
	},
	"CertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateAuthority",
		ReadOnly:       true,
		Type:           "string",
	},
	"ElementalVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "elementalVersion",
		ReadOnly:       true,
		Type:           "string",
	},
	"GaiaVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "gaiaVersion",
		ReadOnly:       true,
		Type:           "string",
	},
	"GoogleClientID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "googleClientID",
		ReadOnly:       true,
		Type:           "string",
	},
	"KairosDBURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "kairosDBURL",
		ReadOnly:       true,
		Type:           "string",
	},
	"ManipulateVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "manipulateVersion",
		ReadOnly:       true,
		Type:           "string",
	},
	"MidgardURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "midgardURL",
		ReadOnly:       true,
		Type:           "string",
	},
	"PubSubService": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "pubSubService",
		ReadOnly:       true,
		Type:           "string",
	},
	"SquallVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "squallVersion",
		ReadOnly:       true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Degraded", "Failure", "Ok"},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Type:           "enum",
	},
	"ZackURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "zackURL",
		Required:       true,
		Type:           "string",
	},
}
