// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RemotePingProbeNamespaceTypeValue represents the possible values for attribute "namespaceType".
type RemotePingProbeNamespaceTypeValue string

const (
	// RemotePingProbeNamespaceTypeHash represents the value Hash.
	RemotePingProbeNamespaceTypeHash RemotePingProbeNamespaceTypeValue = "Hash"

	// RemotePingProbeNamespaceTypePlain represents the value Plain.
	RemotePingProbeNamespaceTypePlain RemotePingProbeNamespaceTypeValue = "Plain"
)

// RemotePingProbe represents the model of a remotepingprobe
type RemotePingProbe struct {
	// The controller ID that manages the ping report.
	ControllerID string `json:"controllerID,omitempty" msgpack:"controllerID,omitempty" bson:"controllerid,omitempty" mapstructure:"controllerID,omitempty"`

	// The namespace where the ping report is stored. Only applicable when the remote
	// controller is empty.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Type of the namespace reported. It can be hash or plain, depending on various
	// factors.
	NamespaceType RemotePingProbeNamespaceTypeValue `json:"namespaceType" msgpack:"namespaceType" bson:"namespacetype" mapstructure:"namespaceType,omitempty"`

	// The ID of the probe. Only applicable when the remote controller is empty.
	ProbeID string `json:"probeID,omitempty" msgpack:"probeID,omitempty" bson:"probeid,omitempty" mapstructure:"probeID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRemotePingProbe returns a new *RemotePingProbe
func NewRemotePingProbe() *RemotePingProbe {

	return &RemotePingProbe{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RemotePingProbe) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRemotePingProbe{}

	s.ControllerID = o.ControllerID
	s.Namespace = o.Namespace
	s.NamespaceType = o.NamespaceType
	s.ProbeID = o.ProbeID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RemotePingProbe) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRemotePingProbe{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ControllerID = s.ControllerID
	o.Namespace = s.Namespace
	o.NamespaceType = s.NamespaceType
	o.ProbeID = s.ProbeID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *RemotePingProbe) BleveType() string {

	return "remotepingprobe"
}

// DeepCopy returns a deep copy if the RemotePingProbe.
func (o *RemotePingProbe) DeepCopy() *RemotePingProbe {

	if o == nil {
		return nil
	}

	out := &RemotePingProbe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RemotePingProbe.
func (o *RemotePingProbe) DeepCopyInto(out *RemotePingProbe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RemotePingProbe: %s", err))
	}

	*out = *target.(*RemotePingProbe)
}

// Validate valides the current information stored into the structure.
func (o *RemotePingProbe) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("namespaceType", string(o.NamespaceType), []string{"Plain", "Hash"}, true); err != nil {
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
func (*RemotePingProbe) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RemotePingProbeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RemotePingProbeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RemotePingProbe) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RemotePingProbeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RemotePingProbe) ValueForAttribute(name string) any {

	switch name {
	case "controllerID":
		return o.ControllerID
	case "namespace":
		return o.Namespace
	case "namespaceType":
		return o.NamespaceType
	case "probeID":
		return o.ProbeID
	}

	return nil
}

// RemotePingProbeAttributesMap represents the map of attribute for RemotePingProbe.
var RemotePingProbeAttributesMap = map[string]elemental.AttributeSpecification{
	"ControllerID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "controllerid",
		ConvertedName:  "ControllerID",
		Description:    `The controller ID that manages the ping report.`,
		Exposed:        true,
		Name:           "controllerID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description: `The namespace where the ping report is stored. Only applicable when the remote
controller is empty.`,
		Exposed:  true,
		Name:     "namespace",
		ReadOnly: true,
		Stored:   true,
		Type:     "string",
	},
	"NamespaceType": {
		AllowedChoices: []string{"Plain", "Hash"},
		Autogenerated:  true,
		BSONFieldName:  "namespacetype",
		ConvertedName:  "NamespaceType",
		Description: `Type of the namespace reported. It can be hash or plain, depending on various
factors.`,
		Exposed:  true,
		Name:     "namespaceType",
		ReadOnly: true,
		Stored:   true,
		Type:     "enum",
	},
	"ProbeID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "probeid",
		ConvertedName:  "ProbeID",
		Description:    `The ID of the probe. Only applicable when the remote controller is empty.`,
		Exposed:        true,
		Name:           "probeID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// RemotePingProbeLowerCaseAttributesMap represents the map of attribute for RemotePingProbe.
var RemotePingProbeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"controllerid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "controllerid",
		ConvertedName:  "ControllerID",
		Description:    `The controller ID that manages the ping report.`,
		Exposed:        true,
		Name:           "controllerID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description: `The namespace where the ping report is stored. Only applicable when the remote
controller is empty.`,
		Exposed:  true,
		Name:     "namespace",
		ReadOnly: true,
		Stored:   true,
		Type:     "string",
	},
	"namespacetype": {
		AllowedChoices: []string{"Plain", "Hash"},
		Autogenerated:  true,
		BSONFieldName:  "namespacetype",
		ConvertedName:  "NamespaceType",
		Description: `Type of the namespace reported. It can be hash or plain, depending on various
factors.`,
		Exposed:  true,
		Name:     "namespaceType",
		ReadOnly: true,
		Stored:   true,
		Type:     "enum",
	},
	"probeid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "probeid",
		ConvertedName:  "ProbeID",
		Description:    `The ID of the probe. Only applicable when the remote controller is empty.`,
		Exposed:        true,
		Name:           "probeID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesRemotePingProbe struct {
	ControllerID  string                            `bson:"controllerid,omitempty"`
	Namespace     string                            `bson:"namespace,omitempty"`
	NamespaceType RemotePingProbeNamespaceTypeValue `bson:"namespacetype"`
	ProbeID       string                            `bson:"probeid,omitempty"`
}
