package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudTargetGroup represents the model of a cloudtargetgroup
type CloudTargetGroup struct {
	// The name of the load balancer.
	LoadBalancerName string `json:"loadBalancerName" msgpack:"loadBalancerName" bson:"loadbalancername" mapstructure:"loadBalancerName,omitempty"`

	// The port for the target group.
	Port string `json:"port" msgpack:"port" bson:"-" mapstructure:"port,omitempty"`

	// The ID for the target group.
	TargetGroupID string `json:"targetGroupID" msgpack:"targetGroupID" bson:"targetgroupid" mapstructure:"targetGroupID,omitempty"`

	// The type of the next hop object.
	TargetType string `json:"targetType" msgpack:"targetType" bson:"targettype" mapstructure:"targetType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudTargetGroup returns a new *CloudTargetGroup
func NewCloudTargetGroup() *CloudTargetGroup {

	return &CloudTargetGroup{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudTargetGroup) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudTargetGroup{}

	s.LoadBalancerName = o.LoadBalancerName
	s.TargetGroupID = o.TargetGroupID
	s.TargetType = o.TargetType

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudTargetGroup) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudTargetGroup{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.LoadBalancerName = s.LoadBalancerName
	o.TargetGroupID = s.TargetGroupID
	o.TargetType = s.TargetType

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudTargetGroup) BleveType() string {

	return "cloudtargetgroup"
}

// DeepCopy returns a deep copy if the CloudTargetGroup.
func (o *CloudTargetGroup) DeepCopy() *CloudTargetGroup {

	if o == nil {
		return nil
	}

	out := &CloudTargetGroup{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudTargetGroup.
func (o *CloudTargetGroup) DeepCopyInto(out *CloudTargetGroup) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudTargetGroup: %s", err))
	}

	*out = *target.(*CloudTargetGroup)
}

// Validate valides the current information stored into the structure.
func (o *CloudTargetGroup) Validate() error {

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

type mongoAttributesCloudTargetGroup struct {
	LoadBalancerName string `bson:"loadbalancername"`
	TargetGroupID    string `bson:"targetgroupid"`
	TargetType       string `bson:"targettype"`
}
