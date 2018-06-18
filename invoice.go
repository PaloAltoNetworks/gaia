package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
	"time"
)

// InvoiceBilledToProviderValue represents the possible values for attribute "billedToProvider".
type InvoiceBilledToProviderValue string

const (
	// InvoiceBilledToProviderAWS represents the value AWS.
	InvoiceBilledToProviderAWS InvoiceBilledToProviderValue = "AWS"

	// InvoiceBilledToProviderAporeto represents the value Aporeto.
	InvoiceBilledToProviderAporeto InvoiceBilledToProviderValue = "Aporeto"
)

// InvoiceIdentity represents the Identity of the object.
var InvoiceIdentity = elemental.Identity{
	Name:     "invoice",
	Category: "invoices",
	Private:  false,
}

// InvoicesList represents a list of Invoices
type InvoicesList []*Invoice

// Identity returns the identity of the objects in the list.
func (o InvoicesList) Identity() elemental.Identity {

	return InvoiceIdentity
}

// Copy returns a pointer to a copy the InvoicesList.
func (o InvoicesList) Copy() elemental.Identifiables {

	copy := append(InvoicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InvoicesList.
func (o InvoicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InvoicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Invoice))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InvoicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InvoicesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o InvoicesList) Version() int {

	return 1
}

// Invoice represents the model of a invoice
type Invoice struct {
	// ID is the id of the invoice.
	ID string `json:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// AccountID references the id of the customer that this invoice belongs to.
	AccountID string `json:"accountID" bson:"accountid" mapstructure:"accountID,omitempty"`

	// BilledToProvider holds the name of the provider that this invoice was billed to.
	BilledToProvider InvoiceBilledToProviderValue `json:"billedToProvider" bson:"billedtoprovider" mapstructure:"billedToProvider,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// EndDate holds the end date for this invoice.
	EndDate time.Time `json:"endDate" bson:"enddate" mapstructure:"endDate,omitempty"`

	// StartDate holds the start date for this invoice.
	StartDate time.Time `json:"startDate" bson:"startdate" mapstructure:"startDate,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewInvoice returns a new *Invoice
func NewInvoice() *Invoice {

	return &Invoice{
		ModelVersion:     1,
		BilledToProvider: "Aporeto",
	}
}

// Identity returns the Identity of the object.
func (o *Invoice) Identity() elemental.Identity {

	return InvoiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Invoice) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Invoice) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Invoice) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Invoice) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Invoice) Doc() string {
	return `This api allows to view invoices for Aporeto customers.`
}

func (o *Invoice) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Invoice) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("billedToProvider", string(o.BilledToProvider), []string{"Aporeto", "AWS"}, false); err != nil {
		errors = append(errors, err)
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
func (*Invoice) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InvoiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InvoiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Invoice) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InvoiceAttributesMap
}

// InvoiceAttributesMap represents the map of attribute for Invoice.
var InvoiceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the id of the invoice.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"AccountID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `AccountID references the id of the customer that this invoice belongs to.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "accountID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"BilledToProvider": elemental.AttributeSpecification{
		AllowedChoices: []string{"Aporeto", "AWS"},
		ConvertedName:  "BilledToProvider",
		DefaultValue:   InvoiceBilledToProviderAporeto,
		Description:    `BilledToProvider holds the name of the provider that this invoice was billed to.`,
		Exposed:        true,
		Name:           "billedToProvider",
		Stored:         true,
		Type:           "enum",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"EndDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EndDate",
		Description:    `EndDate holds the end date for this invoice.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "endDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"StartDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "StartDate",
		Description:    `StartDate holds the start date for this invoice.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "startDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// InvoiceLowerCaseAttributesMap represents the map of attribute for Invoice.
var InvoiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the id of the invoice.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"accountid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `AccountID references the id of the customer that this invoice belongs to.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "accountID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"billedtoprovider": elemental.AttributeSpecification{
		AllowedChoices: []string{"Aporeto", "AWS"},
		ConvertedName:  "BilledToProvider",
		DefaultValue:   InvoiceBilledToProviderAporeto,
		Description:    `BilledToProvider holds the name of the provider that this invoice was billed to.`,
		Exposed:        true,
		Name:           "billedToProvider",
		Stored:         true,
		Type:           "enum",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"enddate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EndDate",
		Description:    `EndDate holds the end date for this invoice.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "endDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"startdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "StartDate",
		Description:    `StartDate holds the start date for this invoice.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "startDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
