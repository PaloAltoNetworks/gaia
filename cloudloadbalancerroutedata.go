package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudLoadBalancerRouteData represents the model of a cloudloadbalancerroutedata
type CloudLoadBalancerRouteData struct {
	// The status of the target.
	HealthStatus string `json:"healthStatus" msgpack:"healthStatus" bson:"healthstatus" mapstructure:"healthStatus,omitempty"`

	// The port for the target group.
	Port string `json:"port" msgpack:"port" bson:"-" mapstructure:"port,omitempty"`

	// The ID for the target group.
	TargetGroupID string `json:"targetGroupID" msgpack:"targetGroupID" bson:"targetgroupid" mapstructure:"targetGroupID,omitempty"`

	// The ID of the target object.
	TargetID string `json:"targetID" msgpack:"targetID" bson:"targetid" mapstructure:"targetID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudLoadBalancerRouteData returns a new *CloudLoadBalancerRouteData
func NewCloudLoadBalancerRouteData() *CloudLoadBalancerRouteData {

	return &CloudLoadBalancerRouteData{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerRouteData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudLoadBalancerRouteData{}

	s.HealthStatus = o.HealthStatus
	s.TargetGroupID = o.TargetGroupID
	s.TargetID = o.TargetID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudLoadBalancerRouteData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudLoadBalancerRouteData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.HealthStatus = s.HealthStatus
	o.TargetGroupID = s.TargetGroupID
	o.TargetID = s.TargetID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudLoadBalancerRouteData) BleveType() string {

	return "cloudloadbalancerroutedata"
}

// DeepCopy returns a deep copy if the CloudLoadBalancerRouteData.
func (o *CloudLoadBalancerRouteData) DeepCopy() *CloudLoadBalancerRouteData {

	if o == nil {
		return nil
	}

	out := &CloudLoadBalancerRouteData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudLoadBalancerRouteData.
func (o *CloudLoadBalancerRouteData) DeepCopyInto(out *CloudLoadBalancerRouteData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudLoadBalancerRouteData: %s", err))
	}

	*out = *target.(*CloudLoadBalancerRouteData)
}

// Validate valides the current information stored into the structure.
func (o *CloudLoadBalancerRouteData) Validate() error {

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

type mongoAttributesCloudLoadBalancerRouteData struct {
	HealthStatus  string `bson:"healthstatus"`
	TargetGroupID string `bson:"targetgroupid"`
	TargetID      string `bson:"targetid"`
}
