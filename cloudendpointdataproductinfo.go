package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEndpointDataProductInfo represents the model of a cloudendpointdataproductinfo
type CloudEndpointDataProductInfo struct {
	// The ID of the corresponding product.
	ProductID string `json:"productID,omitempty" msgpack:"productID,omitempty" bson:"productid" mapstructure:"productID,omitempty"`

	// The type of the product.
	Type string `json:"type,omitempty" msgpack:"type,omitempty" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEndpointDataProductInfo returns a new *CloudEndpointDataProductInfo
func NewCloudEndpointDataProductInfo() *CloudEndpointDataProductInfo {

	return &CloudEndpointDataProductInfo{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointDataProductInfo) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEndpointDataProductInfo{}

	s.ProductID = o.ProductID
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointDataProductInfo) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEndpointDataProductInfo{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ProductID = s.ProductID
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEndpointDataProductInfo) BleveType() string {

	return "cloudendpointdataproductinfo"
}

// DeepCopy returns a deep copy if the CloudEndpointDataProductInfo.
func (o *CloudEndpointDataProductInfo) DeepCopy() *CloudEndpointDataProductInfo {

	if o == nil {
		return nil
	}

	out := &CloudEndpointDataProductInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEndpointDataProductInfo.
func (o *CloudEndpointDataProductInfo) DeepCopyInto(out *CloudEndpointDataProductInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEndpointDataProductInfo: %s", err))
	}

	*out = *target.(*CloudEndpointDataProductInfo)
}

// Validate valides the current information stored into the structure.
func (o *CloudEndpointDataProductInfo) Validate() error {

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

type mongoAttributesCloudEndpointDataProductInfo struct {
	ProductID string `bson:"productid,omitempty"`
	Type      string `bson:"type,omitempty"`
}
