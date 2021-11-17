package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudLoadBalancerData represents the model of a cloudloadbalancerdata
type CloudLoadBalancerData struct {
	// Mapping of target group ID list associated with a listener.
	Listenertargetmapping map[string][]string `json:"listenertargetmapping" msgpack:"listenertargetmapping" bson:"listenertargetmapping" mapstructure:"listenertargetmapping,omitempty"`

	// The name of the load balancer.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The scheme tells whether the load balancer is internet facing or internal.
	Scheme string `json:"scheme" msgpack:"scheme" bson:"scheme" mapstructure:"scheme,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudLoadBalancerData returns a new *CloudLoadBalancerData
func NewCloudLoadBalancerData() *CloudLoadBalancerData {

	return &CloudLoadBalancerData{
		ModelVersion:          1,
		Listenertargetmapping: map[string][]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudLoadBalancerData{}

	s.Listenertargetmapping = o.Listenertargetmapping
	s.Name = o.Name
	s.Scheme = o.Scheme

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudLoadBalancerData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Listenertargetmapping = s.Listenertargetmapping
	o.Name = s.Name
	o.Scheme = s.Scheme

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudLoadBalancerData) BleveType() string {

	return "cloudloadbalancerdata"
}

// DeepCopy returns a deep copy if the CloudLoadBalancerData.
func (o *CloudLoadBalancerData) DeepCopy() *CloudLoadBalancerData {

	if o == nil {
		return nil
	}

	out := &CloudLoadBalancerData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudLoadBalancerData.
func (o *CloudLoadBalancerData) DeepCopyInto(out *CloudLoadBalancerData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudLoadBalancerData: %s", err))
	}

	*out = *target.(*CloudLoadBalancerData)
}

// Validate valides the current information stored into the structure.
func (o *CloudLoadBalancerData) Validate() error {

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

type mongoAttributesCloudLoadBalancerData struct {
	Listenertargetmapping map[string][]string `bson:"listenertargetmapping"`
	Name                  string              `bson:"name"`
	Scheme                string              `bson:"scheme"`
}
