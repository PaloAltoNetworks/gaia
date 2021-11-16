package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudLoadBalancerData represents the model of a cloudloadbalancerdata
type CloudLoadBalancerData struct {
	// The name of the load balancer.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The scheme tells whether the load balancer is internet facing or internal.
	Scheme string `json:"scheme" msgpack:"scheme" bson:"scheme" mapstructure:"scheme,omitempty"`

	// Target groups associated with this load balancer.
	TargetGrouplist []*CloudTargetGroup `json:"targetGrouplist" msgpack:"targetGrouplist" bson:"targetgrouplist" mapstructure:"targetGrouplist,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudLoadBalancerData returns a new *CloudLoadBalancerData
func NewCloudLoadBalancerData() *CloudLoadBalancerData {

	return &CloudLoadBalancerData{
		ModelVersion:    1,
		TargetGrouplist: []*CloudTargetGroup{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudLoadBalancerData{}

	s.Name = o.Name
	s.Scheme = o.Scheme
	s.TargetGrouplist = o.TargetGrouplist

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

	o.Name = s.Name
	o.Scheme = s.Scheme
	o.TargetGrouplist = s.TargetGrouplist

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

	for _, sub := range o.TargetGrouplist {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesCloudLoadBalancerData struct {
	Name            string              `bson:"name"`
	Scheme          string              `bson:"scheme"`
	TargetGrouplist []*CloudTargetGroup `bson:"targetgrouplist"`
}
