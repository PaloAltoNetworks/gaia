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
func (o *RemotePingProbe) GetBSON() (interface{}, error) {

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

type mongoAttributesRemotePingProbe struct {
	ControllerID  string                            `bson:"controllerid,omitempty"`
	Namespace     string                            `bson:"namespace,omitempty"`
	NamespaceType RemotePingProbeNamespaceTypeValue `bson:"namespacetype"`
	ProbeID       string                            `bson:"probeid,omitempty"`
}
