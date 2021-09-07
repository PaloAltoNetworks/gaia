package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// KubernetesClusterIdentity represents the Identity of the object.
var KubernetesClusterIdentity = elemental.Identity{
	Name:     "kubernetescluster",
	Category: "kubernetesclusters",
	Package:  "squall",
	Private:  false,
}

// KubernetesClustersList represents a list of KubernetesClusters
type KubernetesClustersList []*KubernetesCluster

// Identity returns the identity of the objects in the list.
func (o KubernetesClustersList) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Copy returns a pointer to a copy the KubernetesClustersList.
func (o KubernetesClustersList) Copy() elemental.Identifiables {

	copy := append(KubernetesClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the KubernetesClustersList.
func (o KubernetesClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(KubernetesClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*KubernetesCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o KubernetesClustersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o KubernetesClustersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the KubernetesClustersList converted to SparseKubernetesClustersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o KubernetesClustersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseKubernetesClustersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseKubernetesCluster)
	}

	return out
}

// Version returns the version of the content.
func (o KubernetesClustersList) Version() int {

	return 1
}

// KubernetesCluster represents the model of a kubernetescluster
type KubernetesCluster struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// API versions supported by the API server.
	Apiversions []string `json:"apiversions" msgpack:"apiversions" bson:"apiversions" mapstructure:"apiversions,omitempty"`

	// Cluster external IP address.
	ExternalIP string `json:"externalIP" msgpack:"externalIP" bson:"externalip" mapstructure:"externalIP,omitempty"`

	// Cluster internal IP address.
	InternalIP string `json:"internalIP" msgpack:"internalIP" bson:"internalip" mapstructure:"internalIP,omitempty"`

	// Namespace associated with the cluster.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewKubernetesCluster returns a new *KubernetesCluster
func NewKubernetesCluster() *KubernetesCluster {

	return &KubernetesCluster{
		ModelVersion: 1,
		Annotations:  map[string][]string{},
		Apiversions:  []string{},
	}
}

// Identity returns the Identity of the object.
func (o *KubernetesCluster) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KubernetesCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KubernetesCluster) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *KubernetesCluster) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesKubernetesCluster{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.Apiversions = o.Apiversions
	s.ExternalIP = o.ExternalIP
	s.InternalIP = o.InternalIP
	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *KubernetesCluster) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesKubernetesCluster{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.Apiversions = s.Apiversions
	o.ExternalIP = s.ExternalIP
	o.InternalIP = s.InternalIP
	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *KubernetesCluster) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *KubernetesCluster) BleveType() string {

	return "kubernetescluster"
}

// DefaultOrder returns the list of default ordering fields.
func (o *KubernetesCluster) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *KubernetesCluster) Doc() string {

	return `Used to represent an instance of a Kubernetes API server.`
}

func (o *KubernetesCluster) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *KubernetesCluster) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseKubernetesCluster{
			ID:          &o.ID,
			Annotations: &o.Annotations,
			Apiversions: &o.Apiversions,
			ExternalIP:  &o.ExternalIP,
			InternalIP:  &o.InternalIP,
			Namespace:   &o.Namespace,
		}
	}

	sp := &SparseKubernetesCluster{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "apiversions":
			sp.Apiversions = &(o.Apiversions)
		case "externalIP":
			sp.ExternalIP = &(o.ExternalIP)
		case "internalIP":
			sp.InternalIP = &(o.InternalIP)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseKubernetesCluster to the object.
func (o *KubernetesCluster) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseKubernetesCluster)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.Apiversions != nil {
		o.Apiversions = *so.Apiversions
	}
	if so.ExternalIP != nil {
		o.ExternalIP = *so.ExternalIP
	}
	if so.InternalIP != nil {
		o.InternalIP = *so.InternalIP
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the KubernetesCluster.
func (o *KubernetesCluster) DeepCopy() *KubernetesCluster {

	if o == nil {
		return nil
	}

	out := &KubernetesCluster{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *KubernetesCluster.
func (o *KubernetesCluster) DeepCopyInto(out *KubernetesCluster) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy KubernetesCluster: %s", err))
	}

	*out = *target.(*KubernetesCluster)
}

// Validate valides the current information stored into the structure.
func (o *KubernetesCluster) Validate() error {

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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*KubernetesCluster) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := KubernetesClusterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return KubernetesClusterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*KubernetesCluster) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return KubernetesClusterAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *KubernetesCluster) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "apiversions":
		return o.Apiversions
	case "externalIP":
		return o.ExternalIP
	case "internalIP":
		return o.InternalIP
	case "namespace":
		return o.Namespace
	}

	return nil
}

// KubernetesClusterAttributesMap represents the map of attribute for KubernetesCluster.
var KubernetesClusterAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"Apiversions": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiversions",
		ConvertedName:  "Apiversions",
		Description:    `API versions supported by the API server.`,
		Exposed:        true,
		Name:           "apiversions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ExternalIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "externalip",
		ConvertedName:  "ExternalIP",
		Description:    `Cluster external IP address.`,
		Exposed:        true,
		Name:           "externalIP",
		Stored:         true,
		Type:           "string",
	},
	"InternalIP": {
		AllowedChoices: []string{},
		BSONFieldName:  "internalip",
		ConvertedName:  "InternalIP",
		Description:    `Cluster internal IP address.`,
		Exposed:        true,
		Name:           "internalIP",
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace associated with the cluster.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// KubernetesClusterLowerCaseAttributesMap represents the map of attribute for KubernetesCluster.
var KubernetesClusterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"apiversions": {
		AllowedChoices: []string{},
		BSONFieldName:  "apiversions",
		ConvertedName:  "Apiversions",
		Description:    `API versions supported by the API server.`,
		Exposed:        true,
		Name:           "apiversions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"externalip": {
		AllowedChoices: []string{},
		BSONFieldName:  "externalip",
		ConvertedName:  "ExternalIP",
		Description:    `Cluster external IP address.`,
		Exposed:        true,
		Name:           "externalIP",
		Stored:         true,
		Type:           "string",
	},
	"internalip": {
		AllowedChoices: []string{},
		BSONFieldName:  "internalip",
		ConvertedName:  "InternalIP",
		Description:    `Cluster internal IP address.`,
		Exposed:        true,
		Name:           "internalIP",
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace associated with the cluster.`,
		Exposed:        true,
		Name:           "namespace",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseKubernetesClustersList represents a list of SparseKubernetesClusters
type SparseKubernetesClustersList []*SparseKubernetesCluster

// Identity returns the identity of the objects in the list.
func (o SparseKubernetesClustersList) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Copy returns a pointer to a copy the SparseKubernetesClustersList.
func (o SparseKubernetesClustersList) Copy() elemental.Identifiables {

	copy := append(SparseKubernetesClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseKubernetesClustersList.
func (o SparseKubernetesClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseKubernetesClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseKubernetesCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseKubernetesClustersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseKubernetesClustersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseKubernetesClustersList converted to KubernetesClustersList.
func (o SparseKubernetesClustersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseKubernetesClustersList) Version() int {

	return 1
}

// SparseKubernetesCluster represents the sparse version of a kubernetescluster.
type SparseKubernetesCluster struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// API versions supported by the API server.
	Apiversions *[]string `json:"apiversions,omitempty" msgpack:"apiversions,omitempty" bson:"apiversions,omitempty" mapstructure:"apiversions,omitempty"`

	// Cluster external IP address.
	ExternalIP *string `json:"externalIP,omitempty" msgpack:"externalIP,omitempty" bson:"externalip,omitempty" mapstructure:"externalIP,omitempty"`

	// Cluster internal IP address.
	InternalIP *string `json:"internalIP,omitempty" msgpack:"internalIP,omitempty" bson:"internalip,omitempty" mapstructure:"internalIP,omitempty"`

	// Namespace associated with the cluster.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseKubernetesCluster returns a new  SparseKubernetesCluster.
func NewSparseKubernetesCluster() *SparseKubernetesCluster {
	return &SparseKubernetesCluster{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseKubernetesCluster) Identity() elemental.Identity {

	return KubernetesClusterIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseKubernetesCluster) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseKubernetesCluster) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseKubernetesCluster) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseKubernetesCluster{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.Apiversions != nil {
		s.Apiversions = o.Apiversions
	}
	if o.ExternalIP != nil {
		s.ExternalIP = o.ExternalIP
	}
	if o.InternalIP != nil {
		s.InternalIP = o.InternalIP
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseKubernetesCluster) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseKubernetesCluster{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.Apiversions != nil {
		o.Apiversions = s.Apiversions
	}
	if s.ExternalIP != nil {
		o.ExternalIP = s.ExternalIP
	}
	if s.InternalIP != nil {
		o.InternalIP = s.InternalIP
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseKubernetesCluster) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseKubernetesCluster) ToPlain() elemental.PlainIdentifiable {

	out := NewKubernetesCluster()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.Apiversions != nil {
		out.Apiversions = *o.Apiversions
	}
	if o.ExternalIP != nil {
		out.ExternalIP = *o.ExternalIP
	}
	if o.InternalIP != nil {
		out.InternalIP = *o.InternalIP
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// DeepCopy returns a deep copy if the SparseKubernetesCluster.
func (o *SparseKubernetesCluster) DeepCopy() *SparseKubernetesCluster {

	if o == nil {
		return nil
	}

	out := &SparseKubernetesCluster{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseKubernetesCluster.
func (o *SparseKubernetesCluster) DeepCopyInto(out *SparseKubernetesCluster) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseKubernetesCluster: %s", err))
	}

	*out = *target.(*SparseKubernetesCluster)
}

type mongoAttributesKubernetesCluster struct {
	ID          bson.ObjectId       `bson:"_id,omitempty"`
	Annotations map[string][]string `bson:"annotations"`
	Apiversions []string            `bson:"apiversions"`
	ExternalIP  string              `bson:"externalip"`
	InternalIP  string              `bson:"internalip"`
	Namespace   string              `bson:"namespace"`
}
type mongoAttributesSparseKubernetesCluster struct {
	ID          bson.ObjectId        `bson:"_id,omitempty"`
	Annotations *map[string][]string `bson:"annotations,omitempty"`
	Apiversions *[]string            `bson:"apiversions,omitempty"`
	ExternalIP  *string              `bson:"externalip,omitempty"`
	InternalIP  *string              `bson:"internalip,omitempty"`
	Namespace   *string              `bson:"namespace,omitempty"`
}
