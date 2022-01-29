package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerProfileKubernetesMetadataExtractorValue represents the possible values for attribute "kubernetesMetadataExtractor".
type EnforcerProfileKubernetesMetadataExtractorValue string

const (
	// EnforcerProfileKubernetesMetadataExtractorKubeSquall represents the value KubeSquall.
	EnforcerProfileKubernetesMetadataExtractorKubeSquall EnforcerProfileKubernetesMetadataExtractorValue = "KubeSquall"

	// EnforcerProfileKubernetesMetadataExtractorPodAtomic represents the value PodAtomic.
	EnforcerProfileKubernetesMetadataExtractorPodAtomic EnforcerProfileKubernetesMetadataExtractorValue = "PodAtomic"

	// EnforcerProfileKubernetesMetadataExtractorPodContainers represents the value PodContainers.
	EnforcerProfileKubernetesMetadataExtractorPodContainers EnforcerProfileKubernetesMetadataExtractorValue = "PodContainers"
)

// EnforcerProfileMetadataExtractorValue represents the possible values for attribute "metadataExtractor".
type EnforcerProfileMetadataExtractorValue string

const (
	// EnforcerProfileMetadataExtractorDocker represents the value Docker.
	EnforcerProfileMetadataExtractorDocker EnforcerProfileMetadataExtractorValue = "Docker"

	// EnforcerProfileMetadataExtractorECS represents the value ECS.
	EnforcerProfileMetadataExtractorECS EnforcerProfileMetadataExtractorValue = "ECS"

	// EnforcerProfileMetadataExtractorKubernetes represents the value Kubernetes.
	EnforcerProfileMetadataExtractorKubernetes EnforcerProfileMetadataExtractorValue = "Kubernetes"
)

// EnforcerProfileSyslogFormatValue represents the possible values for attribute "syslogFormat".
type EnforcerProfileSyslogFormatValue string

const (
	// EnforcerProfileSyslogFormatAuto represents the value Auto.
	EnforcerProfileSyslogFormatAuto EnforcerProfileSyslogFormatValue = "Auto"

	// EnforcerProfileSyslogFormatBSD represents the value BSD.
	EnforcerProfileSyslogFormatBSD EnforcerProfileSyslogFormatValue = "BSD"

	// EnforcerProfileSyslogFormatIETF represents the value IETF.
	EnforcerProfileSyslogFormatIETF EnforcerProfileSyslogFormatValue = "IETF"
)

// EnforcerProfileIdentity represents the Identity of the object.
var EnforcerProfileIdentity = elemental.Identity{
	Name:     "enforcerprofile",
	Category: "enforcerprofiles",
	Package:  "squall",
	Private:  false,
}

// EnforcerProfilesList represents a list of EnforcerProfiles
type EnforcerProfilesList []*EnforcerProfile

// Identity returns the identity of the objects in the list.
func (o EnforcerProfilesList) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the EnforcerProfilesList.
func (o EnforcerProfilesList) Copy() elemental.Identifiables {

	copy := append(EnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerProfilesList.
func (o EnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the EnforcerProfilesList converted to SparseEnforcerProfilesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerProfilesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerProfilesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerProfile)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerProfilesList) Version() int {

	return 1
}

// EnforcerProfile represents the model of a enforcerprofile
// +k8s:openapi-gen=true
type EnforcerProfile struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Ignore traffic with a source or destination matching the specified
	// interfaces.
	ExcludedInterfaces []string `json:"excludedInterfaces" msgpack:"excludedInterfaces" bson:"excludedinterfaces" mapstructure:"excludedInterfaces,omitempty"`

	// Ignore any networks specified here and do not even report any flows.
	// This can be useful for excluding localhost loopback traffic, ignoring
	// traffic to the Kubernetes API, and using Microsegmentation for SSH only.
	ExcludedNetworks []string `json:"excludedNetworks" msgpack:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// Frequency that flow report counts are upated.
	FlowReportInterval string `json:"flowReportInterval" msgpack:"flowReportInterval" bson:"flowreportinterval" mapstructure:"flowReportInterval,omitempty"`

	// A tag expression that identifies processing units to ignore. This can be
	// useful to exclude `kube-system` pods, AWS EC2 agent pods, and third-party
	// agents.
	IgnoreExpression [][]string `json:"ignoreExpression" msgpack:"ignoreExpression" bson:"ignoreexpression" mapstructure:"ignoreExpression,omitempty"`

	// This field is kept for backward compatibility for enforcers <= 3.5.
	KubernetesMetadataExtractor EnforcerProfileKubernetesMetadataExtractorValue `json:"kubernetesMetadataExtractor" msgpack:"kubernetesMetadataExtractor" bson:"kubernetesmetadataextractor" mapstructure:"kubernetesMetadataExtractor,omitempty"`

	// This field is kept for backward compatibility for enforcers <= 3.5.
	KubernetesSupportEnabled bool `json:"kubernetesSupportEnabled" msgpack:"kubernetesSupportEnabled" bson:"kubernetessupportenabled" mapstructure:"kubernetesSupportEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// This field is kept for backward compatibility for enforcers <= 3.5.
	MetadataExtractor EnforcerProfileMetadataExtractorValue `json:"metadataExtractor" msgpack:"metadataExtractor" bson:"metadataextractor" mapstructure:"metadataExtractor,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Enables syslog functionality of enforcers using this enforcerprofile.
	SyslogEnabled bool `json:"syslogEnabled" msgpack:"syslogEnabled" bson:"syslogenabled" mapstructure:"syslogEnabled,omitempty"`

	// Contains the remote endpoint to dispatch the syslog messages.
	SyslogEndpoint string `json:"syslogEndpoint" msgpack:"syslogEndpoint" bson:"syslogendpoint" mapstructure:"syslogEndpoint,omitempty"`

	// PEM-encoded certificate to expose to the clients for TLS.
	SyslogEndpointTLSClientCertificate string `json:"syslogEndpointTLSClientCertificate" msgpack:"syslogEndpointTLSClientCertificate" bson:"syslogendpointtlsclientcertificate" mapstructure:"syslogEndpointTLSClientCertificate,omitempty"`

	// PEM-encoded server CA certificate.
	SyslogEndpointTLSServerCA string `json:"syslogEndpointTLSServerCA" msgpack:"syslogEndpointTLSServerCA" bson:"syslogendpointtlsserverca" mapstructure:"syslogEndpointTLSServerCA,omitempty"`

	// Contains the list of supported syslog facilities.
	SyslogFacility int `json:"syslogFacility" msgpack:"syslogFacility" bson:"syslogfacility" mapstructure:"syslogFacility,omitempty"`

	// Contains the list of supported syslog message format.
	SyslogFormat EnforcerProfileSyslogFormatValue `json:"syslogFormat" msgpack:"syslogFormat" bson:"syslogformat" mapstructure:"syslogFormat,omitempty"`

	// If empty, the enforcer auto-discovers the TCP networks. Auto-discovery
	// works best in Kubernetes and OpenShift deployments. You may need to manually
	// specify the TCP networks if middle boxes exist that do not comply with
	// [TCP Fast Open RFC 7413](https://tools.ietf.org/html/rfc7413).
	TargetNetworks []string `json:"targetNetworks" msgpack:"targetNetworks" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// If empty, the enforcer enforces all UDP networks. This works best when all UDP
	// networks have enforcers. If some UDP networks do not have enforcers, you
	// may need to manually specify the UDP networks that should be enforced.
	TargetUDPNetworks []string `json:"targetUDPNetworks" msgpack:"targetUDPNetworks" bson:"targetudpnetworks" mapstructure:"targetUDPNetworks,omitempty"`

	// List of trusted certificate authorities. If empty, the main chain of trust
	// will be used.
	TrustedCAs []string `json:"trustedCAs" msgpack:"trustedCAs" bson:"trustedcas" mapstructure:"trustedCAs,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerProfile returns a new *EnforcerProfile
func NewEnforcerProfile() *EnforcerProfile {

	return &EnforcerProfile{
		ModelVersion:                1,
		Annotations:                 map[string][]string{},
		ExcludedInterfaces:          []string{},
		AssociatedTags:              []string{},
		FlowReportInterval:          "30m",
		ExcludedNetworks:            []string{},
		IgnoreExpression:            [][]string{},
		MetadataExtractor:           EnforcerProfileMetadataExtractorDocker,
		SyslogFacility:              1,
		SyslogFormat:                EnforcerProfileSyslogFormatAuto,
		NormalizedTags:              []string{},
		TargetNetworks:              []string{},
		TrustedCAs:                  []string{},
		TargetUDPNetworks:           []string{},
		MigrationsLog:               map[string]string{},
		KubernetesMetadataExtractor: EnforcerProfileKubernetesMetadataExtractorPodAtomic,
		Metadata:                    []string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerProfile) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerProfile) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcerProfile{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.ExcludedInterfaces = o.ExcludedInterfaces
	s.ExcludedNetworks = o.ExcludedNetworks
	s.FlowReportInterval = o.FlowReportInterval
	s.IgnoreExpression = o.IgnoreExpression
	s.KubernetesMetadataExtractor = o.KubernetesMetadataExtractor
	s.KubernetesSupportEnabled = o.KubernetesSupportEnabled
	s.Metadata = o.Metadata
	s.MetadataExtractor = o.MetadataExtractor
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.SyslogEnabled = o.SyslogEnabled
	s.SyslogEndpoint = o.SyslogEndpoint
	s.SyslogEndpointTLSClientCertificate = o.SyslogEndpointTLSClientCertificate
	s.SyslogEndpointTLSServerCA = o.SyslogEndpointTLSServerCA
	s.SyslogFacility = o.SyslogFacility
	s.SyslogFormat = o.SyslogFormat
	s.TargetNetworks = o.TargetNetworks
	s.TargetUDPNetworks = o.TargetUDPNetworks
	s.TrustedCAs = o.TrustedCAs
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerProfile) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcerProfile{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.ExcludedInterfaces = s.ExcludedInterfaces
	o.ExcludedNetworks = s.ExcludedNetworks
	o.FlowReportInterval = s.FlowReportInterval
	o.IgnoreExpression = s.IgnoreExpression
	o.KubernetesMetadataExtractor = s.KubernetesMetadataExtractor
	o.KubernetesSupportEnabled = s.KubernetesSupportEnabled
	o.Metadata = s.Metadata
	o.MetadataExtractor = s.MetadataExtractor
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.SyslogEnabled = s.SyslogEnabled
	o.SyslogEndpoint = s.SyslogEndpoint
	o.SyslogEndpointTLSClientCertificate = s.SyslogEndpointTLSClientCertificate
	o.SyslogEndpointTLSServerCA = s.SyslogEndpointTLSServerCA
	o.SyslogFacility = s.SyslogFacility
	o.SyslogFormat = s.SyslogFormat
	o.TargetNetworks = s.TargetNetworks
	o.TargetUDPNetworks = s.TargetUDPNetworks
	o.TrustedCAs = s.TrustedCAs
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EnforcerProfile) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerProfile) BleveType() string {

	return "enforcerprofile"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *EnforcerProfile) Doc() string {

	return `Allows you to create reusable configuration profiles for your enforcers.
Enforcer profiles contain various startup information that can (for some) be
updated live. Enforcer profiles are assigned to enforcers using a enforcer
profile mapping.`
}

func (o *EnforcerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *EnforcerProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *EnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *EnforcerProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *EnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *EnforcerProfile) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *EnforcerProfile) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *EnforcerProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *EnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *EnforcerProfile) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *EnforcerProfile) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *EnforcerProfile) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *EnforcerProfile) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *EnforcerProfile) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *EnforcerProfile) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *EnforcerProfile) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *EnforcerProfile) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *EnforcerProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *EnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *EnforcerProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *EnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *EnforcerProfile) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *EnforcerProfile) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *EnforcerProfile) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *EnforcerProfile) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *EnforcerProfile) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *EnforcerProfile) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *EnforcerProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *EnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *EnforcerProfile) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *EnforcerProfile) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *EnforcerProfile) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *EnforcerProfile) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerProfile) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerProfile{
			ID:                                 &o.ID,
			Annotations:                        &o.Annotations,
			AssociatedTags:                     &o.AssociatedTags,
			CreateIdempotencyKey:               &o.CreateIdempotencyKey,
			CreateTime:                         &o.CreateTime,
			Description:                        &o.Description,
			ExcludedInterfaces:                 &o.ExcludedInterfaces,
			ExcludedNetworks:                   &o.ExcludedNetworks,
			FlowReportInterval:                 &o.FlowReportInterval,
			IgnoreExpression:                   &o.IgnoreExpression,
			KubernetesMetadataExtractor:        &o.KubernetesMetadataExtractor,
			KubernetesSupportEnabled:           &o.KubernetesSupportEnabled,
			Metadata:                           &o.Metadata,
			MetadataExtractor:                  &o.MetadataExtractor,
			MigrationsLog:                      &o.MigrationsLog,
			Name:                               &o.Name,
			Namespace:                          &o.Namespace,
			NormalizedTags:                     &o.NormalizedTags,
			Propagate:                          &o.Propagate,
			Protected:                          &o.Protected,
			SyslogEnabled:                      &o.SyslogEnabled,
			SyslogEndpoint:                     &o.SyslogEndpoint,
			SyslogEndpointTLSClientCertificate: &o.SyslogEndpointTLSClientCertificate,
			SyslogEndpointTLSServerCA:          &o.SyslogEndpointTLSServerCA,
			SyslogFacility:                     &o.SyslogFacility,
			SyslogFormat:                       &o.SyslogFormat,
			TargetNetworks:                     &o.TargetNetworks,
			TargetUDPNetworks:                  &o.TargetUDPNetworks,
			TrustedCAs:                         &o.TrustedCAs,
			UpdateIdempotencyKey:               &o.UpdateIdempotencyKey,
			UpdateTime:                         &o.UpdateTime,
			ZHash:                              &o.ZHash,
			Zone:                               &o.Zone,
		}
	}

	sp := &SparseEnforcerProfile{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "excludedInterfaces":
			sp.ExcludedInterfaces = &(o.ExcludedInterfaces)
		case "excludedNetworks":
			sp.ExcludedNetworks = &(o.ExcludedNetworks)
		case "flowReportInterval":
			sp.FlowReportInterval = &(o.FlowReportInterval)
		case "ignoreExpression":
			sp.IgnoreExpression = &(o.IgnoreExpression)
		case "kubernetesMetadataExtractor":
			sp.KubernetesMetadataExtractor = &(o.KubernetesMetadataExtractor)
		case "kubernetesSupportEnabled":
			sp.KubernetesSupportEnabled = &(o.KubernetesSupportEnabled)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "metadataExtractor":
			sp.MetadataExtractor = &(o.MetadataExtractor)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "syslogEnabled":
			sp.SyslogEnabled = &(o.SyslogEnabled)
		case "syslogEndpoint":
			sp.SyslogEndpoint = &(o.SyslogEndpoint)
		case "syslogEndpointTLSClientCertificate":
			sp.SyslogEndpointTLSClientCertificate = &(o.SyslogEndpointTLSClientCertificate)
		case "syslogEndpointTLSServerCA":
			sp.SyslogEndpointTLSServerCA = &(o.SyslogEndpointTLSServerCA)
		case "syslogFacility":
			sp.SyslogFacility = &(o.SyslogFacility)
		case "syslogFormat":
			sp.SyslogFormat = &(o.SyslogFormat)
		case "targetNetworks":
			sp.TargetNetworks = &(o.TargetNetworks)
		case "targetUDPNetworks":
			sp.TargetUDPNetworks = &(o.TargetUDPNetworks)
		case "trustedCAs":
			sp.TrustedCAs = &(o.TrustedCAs)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerProfile to the object.
func (o *EnforcerProfile) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerProfile)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.ExcludedInterfaces != nil {
		o.ExcludedInterfaces = *so.ExcludedInterfaces
	}
	if so.ExcludedNetworks != nil {
		o.ExcludedNetworks = *so.ExcludedNetworks
	}
	if so.FlowReportInterval != nil {
		o.FlowReportInterval = *so.FlowReportInterval
	}
	if so.IgnoreExpression != nil {
		o.IgnoreExpression = *so.IgnoreExpression
	}
	if so.KubernetesMetadataExtractor != nil {
		o.KubernetesMetadataExtractor = *so.KubernetesMetadataExtractor
	}
	if so.KubernetesSupportEnabled != nil {
		o.KubernetesSupportEnabled = *so.KubernetesSupportEnabled
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MetadataExtractor != nil {
		o.MetadataExtractor = *so.MetadataExtractor
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.SyslogEnabled != nil {
		o.SyslogEnabled = *so.SyslogEnabled
	}
	if so.SyslogEndpoint != nil {
		o.SyslogEndpoint = *so.SyslogEndpoint
	}
	if so.SyslogEndpointTLSClientCertificate != nil {
		o.SyslogEndpointTLSClientCertificate = *so.SyslogEndpointTLSClientCertificate
	}
	if so.SyslogEndpointTLSServerCA != nil {
		o.SyslogEndpointTLSServerCA = *so.SyslogEndpointTLSServerCA
	}
	if so.SyslogFacility != nil {
		o.SyslogFacility = *so.SyslogFacility
	}
	if so.SyslogFormat != nil {
		o.SyslogFormat = *so.SyslogFormat
	}
	if so.TargetNetworks != nil {
		o.TargetNetworks = *so.TargetNetworks
	}
	if so.TargetUDPNetworks != nil {
		o.TargetUDPNetworks = *so.TargetUDPNetworks
	}
	if so.TrustedCAs != nil {
		o.TrustedCAs = *so.TrustedCAs
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the EnforcerProfile.
func (o *EnforcerProfile) DeepCopy() *EnforcerProfile {

	if o == nil {
		return nil
	}

	out := &EnforcerProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerProfile.
func (o *EnforcerProfile) DeepCopyInto(out *EnforcerProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerProfile: %s", err))
	}

	*out = *target.(*EnforcerProfile)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTimeDuration("flowReportInterval", o.FlowReportInterval); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("ignoreExpression", o.IgnoreExpression); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("kubernetesMetadataExtractor", string(o.KubernetesMetadataExtractor), []string{"KubeSquall", "PodAtomic", "PodContainers"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("metadataExtractor", string(o.MetadataExtractor), []string{"Docker", "ECS", "Kubernetes"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateSyslogEndpoint("syslogEndpoint", o.SyslogEndpoint); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePEM("syslogEndpointTLSClientCertificate", o.SyslogEndpointTLSClientCertificate); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateCA("syslogEndpointTLSServerCA", o.SyslogEndpointTLSServerCA); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("syslogFacility", o.SyslogFacility, int(23), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMinimumInt("syslogFacility", o.SyslogFacility, int(1), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("syslogFormat", string(o.SyslogFormat), []string{"Auto", "BSD", "IETF"}, false); err != nil {
		errors = errors.Append(err)
	}

	// Custom object validation.
	if err := ValidateEnforcerProfile(o); err != nil {
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
func (*EnforcerProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerProfileAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerProfile) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "excludedInterfaces":
		return o.ExcludedInterfaces
	case "excludedNetworks":
		return o.ExcludedNetworks
	case "flowReportInterval":
		return o.FlowReportInterval
	case "ignoreExpression":
		return o.IgnoreExpression
	case "kubernetesMetadataExtractor":
		return o.KubernetesMetadataExtractor
	case "kubernetesSupportEnabled":
		return o.KubernetesSupportEnabled
	case "metadata":
		return o.Metadata
	case "metadataExtractor":
		return o.MetadataExtractor
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "syslogEnabled":
		return o.SyslogEnabled
	case "syslogEndpoint":
		return o.SyslogEndpoint
	case "syslogEndpointTLSClientCertificate":
		return o.SyslogEndpointTLSClientCertificate
	case "syslogEndpointTLSServerCA":
		return o.SyslogEndpointTLSServerCA
	case "syslogFacility":
		return o.SyslogFacility
	case "syslogFormat":
		return o.SyslogFormat
	case "targetNetworks":
		return o.TargetNetworks
	case "targetUDPNetworks":
		return o.TargetUDPNetworks
	case "trustedCAs":
		return o.TrustedCAs
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// EnforcerProfileAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileAttributesMap = map[string]elemental.AttributeSpecification{
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
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ExcludedInterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludedinterfaces",
		ConvertedName:  "ExcludedInterfaces",
		Description: `Ignore traffic with a source or destination matching the specified
interfaces.`,
		Exposed:   true,
		Name:      "excludedInterfaces",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"ExcludedNetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludednetworks",
		ConvertedName:  "ExcludedNetworks",
		Description: `Ignore any networks specified here and do not even report any flows.
This can be useful for excluding localhost loopback traffic, ignoring
traffic to the Kubernetes API, and using Microsegmentation for SSH only.`,
		Exposed:   true,
		Name:      "excludedNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"FlowReportInterval": {
		AllowedChoices: []string{},
		BSONFieldName:  "flowreportinterval",
		ConvertedName:  "FlowReportInterval",
		DefaultValue:   "30m",
		Description:    `Frequency that flow report counts are upated.`,
		Exposed:        true,
		Name:           "flowReportInterval",
		Stored:         true,
		Type:           "string",
	},
	"IgnoreExpression": {
		AllowedChoices: []string{},
		BSONFieldName:  "ignoreexpression",
		ConvertedName:  "IgnoreExpression",
		Description: `A tag expression that identifies processing units to ignore. This can be
useful to exclude ` + "`" + `kube-system` + "`" + ` pods, AWS EC2 agent pods, and third-party
agents.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"KubernetesMetadataExtractor": {
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		BSONFieldName:  "kubernetesmetadataextractor",
		ConvertedName:  "KubernetesMetadataExtractor",
		DefaultValue:   EnforcerProfileKubernetesMetadataExtractorPodAtomic,
		Deprecated:     true,
		Description:    `This field is kept for backward compatibility for enforcers <= 3.5.`,
		Exposed:        true,
		Name:           "kubernetesMetadataExtractor",
		Stored:         true,
		Type:           "enum",
	},
	"KubernetesSupportEnabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "kubernetessupportenabled",
		ConvertedName:  "KubernetesSupportEnabled",
		Deprecated:     true,
		Description:    `This field is kept for backward compatibility for enforcers <= 3.5.`,
		Exposed:        true,
		Name:           "kubernetesSupportEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"Metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"MetadataExtractor": {
		AllowedChoices: []string{"Docker", "ECS", "Kubernetes"},
		BSONFieldName:  "metadataextractor",
		ConvertedName:  "MetadataExtractor",
		DefaultValue:   EnforcerProfileMetadataExtractorDocker,
		Deprecated:     true,
		Description:    `This field is kept for backward compatibility for enforcers <= 3.5.`,
		Exposed:        true,
		Name:           "metadataExtractor",
		Stored:         true,
		Type:           "enum",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"SyslogEnabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogenabled",
		ConvertedName:  "SyslogEnabled",
		Description:    `Enables syslog functionality of enforcers using this enforcerprofile.`,
		Exposed:        true,
		Name:           "syslogEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"SyslogEndpoint": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogendpoint",
		ConvertedName:  "SyslogEndpoint",
		Description:    `Contains the remote endpoint to dispatch the syslog messages.`,
		Exposed:        true,
		Name:           "syslogEndpoint",
		Stored:         true,
		Type:           "string",
	},
	"SyslogEndpointTLSClientCertificate": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogendpointtlsclientcertificate",
		ConvertedName:  "SyslogEndpointTLSClientCertificate",
		Description:    `PEM-encoded certificate to expose to the clients for TLS.`,
		Exposed:        true,
		Name:           "syslogEndpointTLSClientCertificate",
		Stored:         true,
		Type:           "string",
	},
	"SyslogEndpointTLSServerCA": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogendpointtlsserverca",
		ConvertedName:  "SyslogEndpointTLSServerCA",
		Description:    `PEM-encoded server CA certificate.`,
		Exposed:        true,
		Name:           "syslogEndpointTLSServerCA",
		Stored:         true,
		Type:           "string",
	},
	"SyslogFacility": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogfacility",
		ConvertedName:  "SyslogFacility",
		DefaultValue:   1,
		Description:    `Contains the list of supported syslog facilities.`,
		Exposed:        true,
		MaxValue:       23,
		MinValue:       1,
		Name:           "syslogFacility",
		Stored:         true,
		Type:           "integer",
	},
	"SyslogFormat": {
		AllowedChoices: []string{"Auto", "BSD", "IETF"},
		BSONFieldName:  "syslogformat",
		ConvertedName:  "SyslogFormat",
		DefaultValue:   EnforcerProfileSyslogFormatAuto,
		Description:    `Contains the list of supported syslog message format.`,
		Exposed:        true,
		Name:           "syslogFormat",
		Stored:         true,
		Type:           "enum",
	},
	"TargetNetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetnetworks",
		ConvertedName:  "TargetNetworks",
		Description: `If empty, the enforcer auto-discovers the TCP networks. Auto-discovery
works best in Kubernetes and OpenShift deployments. You may need to manually
specify the TCP networks if middle boxes exist that do not comply with
[TCP Fast Open RFC 7413](https://tools.ietf.org/html/rfc7413).`,
		Exposed:   true,
		Name:      "targetNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"TargetUDPNetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetudpnetworks",
		ConvertedName:  "TargetUDPNetworks",
		Description: `If empty, the enforcer enforces all UDP networks. This works best when all UDP
networks have enforcers. If some UDP networks do not have enforcers, you
may need to manually specify the UDP networks that should be enforced.`,
		Exposed:   true,
		Name:      "targetUDPNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"TrustedCAs": {
		AllowedChoices: []string{},
		BSONFieldName:  "trustedcas",
		ConvertedName:  "TrustedCAs",
		Description: `List of trusted certificate authorities. If empty, the main chain of trust
will be used.`,
		Exposed: true,
		Name:    "trustedCAs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// EnforcerProfileLowerCaseAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"excludedinterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludedinterfaces",
		ConvertedName:  "ExcludedInterfaces",
		Description: `Ignore traffic with a source or destination matching the specified
interfaces.`,
		Exposed:   true,
		Name:      "excludedInterfaces",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"excludednetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "excludednetworks",
		ConvertedName:  "ExcludedNetworks",
		Description: `Ignore any networks specified here and do not even report any flows.
This can be useful for excluding localhost loopback traffic, ignoring
traffic to the Kubernetes API, and using Microsegmentation for SSH only.`,
		Exposed:   true,
		Name:      "excludedNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"flowreportinterval": {
		AllowedChoices: []string{},
		BSONFieldName:  "flowreportinterval",
		ConvertedName:  "FlowReportInterval",
		DefaultValue:   "30m",
		Description:    `Frequency that flow report counts are upated.`,
		Exposed:        true,
		Name:           "flowReportInterval",
		Stored:         true,
		Type:           "string",
	},
	"ignoreexpression": {
		AllowedChoices: []string{},
		BSONFieldName:  "ignoreexpression",
		ConvertedName:  "IgnoreExpression",
		Description: `A tag expression that identifies processing units to ignore. This can be
useful to exclude ` + "`" + `kube-system` + "`" + ` pods, AWS EC2 agent pods, and third-party
agents.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"kubernetesmetadataextractor": {
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		BSONFieldName:  "kubernetesmetadataextractor",
		ConvertedName:  "KubernetesMetadataExtractor",
		DefaultValue:   EnforcerProfileKubernetesMetadataExtractorPodAtomic,
		Deprecated:     true,
		Description:    `This field is kept for backward compatibility for enforcers <= 3.5.`,
		Exposed:        true,
		Name:           "kubernetesMetadataExtractor",
		Stored:         true,
		Type:           "enum",
	},
	"kubernetessupportenabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "kubernetessupportenabled",
		ConvertedName:  "KubernetesSupportEnabled",
		Deprecated:     true,
		Description:    `This field is kept for backward compatibility for enforcers <= 3.5.`,
		Exposed:        true,
		Name:           "kubernetesSupportEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"metadata": {
		AllowedChoices: []string{},
		BSONFieldName:  "metadata",
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"metadataextractor": {
		AllowedChoices: []string{"Docker", "ECS", "Kubernetes"},
		BSONFieldName:  "metadataextractor",
		ConvertedName:  "MetadataExtractor",
		DefaultValue:   EnforcerProfileMetadataExtractorDocker,
		Deprecated:     true,
		Description:    `This field is kept for backward compatibility for enforcers <= 3.5.`,
		Exposed:        true,
		Name:           "metadataExtractor",
		Stored:         true,
		Type:           "enum",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"syslogenabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogenabled",
		ConvertedName:  "SyslogEnabled",
		Description:    `Enables syslog functionality of enforcers using this enforcerprofile.`,
		Exposed:        true,
		Name:           "syslogEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"syslogendpoint": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogendpoint",
		ConvertedName:  "SyslogEndpoint",
		Description:    `Contains the remote endpoint to dispatch the syslog messages.`,
		Exposed:        true,
		Name:           "syslogEndpoint",
		Stored:         true,
		Type:           "string",
	},
	"syslogendpointtlsclientcertificate": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogendpointtlsclientcertificate",
		ConvertedName:  "SyslogEndpointTLSClientCertificate",
		Description:    `PEM-encoded certificate to expose to the clients for TLS.`,
		Exposed:        true,
		Name:           "syslogEndpointTLSClientCertificate",
		Stored:         true,
		Type:           "string",
	},
	"syslogendpointtlsserverca": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogendpointtlsserverca",
		ConvertedName:  "SyslogEndpointTLSServerCA",
		Description:    `PEM-encoded server CA certificate.`,
		Exposed:        true,
		Name:           "syslogEndpointTLSServerCA",
		Stored:         true,
		Type:           "string",
	},
	"syslogfacility": {
		AllowedChoices: []string{},
		BSONFieldName:  "syslogfacility",
		ConvertedName:  "SyslogFacility",
		DefaultValue:   1,
		Description:    `Contains the list of supported syslog facilities.`,
		Exposed:        true,
		MaxValue:       23,
		MinValue:       1,
		Name:           "syslogFacility",
		Stored:         true,
		Type:           "integer",
	},
	"syslogformat": {
		AllowedChoices: []string{"Auto", "BSD", "IETF"},
		BSONFieldName:  "syslogformat",
		ConvertedName:  "SyslogFormat",
		DefaultValue:   EnforcerProfileSyslogFormatAuto,
		Description:    `Contains the list of supported syslog message format.`,
		Exposed:        true,
		Name:           "syslogFormat",
		Stored:         true,
		Type:           "enum",
	},
	"targetnetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetnetworks",
		ConvertedName:  "TargetNetworks",
		Description: `If empty, the enforcer auto-discovers the TCP networks. Auto-discovery
works best in Kubernetes and OpenShift deployments. You may need to manually
specify the TCP networks if middle boxes exist that do not comply with
[TCP Fast Open RFC 7413](https://tools.ietf.org/html/rfc7413).`,
		Exposed:   true,
		Name:      "targetNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"targetudpnetworks": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetudpnetworks",
		ConvertedName:  "TargetUDPNetworks",
		Description: `If empty, the enforcer enforces all UDP networks. This works best when all UDP
networks have enforcers. If some UDP networks do not have enforcers, you
may need to manually specify the UDP networks that should be enforced.`,
		Exposed:   true,
		Name:      "targetUDPNetworks",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"trustedcas": {
		AllowedChoices: []string{},
		BSONFieldName:  "trustedcas",
		ConvertedName:  "TrustedCAs",
		Description: `List of trusted certificate authorities. If empty, the main chain of trust
will be used.`,
		Exposed: true,
		Name:    "trustedCAs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseEnforcerProfilesList represents a list of SparseEnforcerProfiles
type SparseEnforcerProfilesList []*SparseEnforcerProfile

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerProfilesList) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerProfilesList.
func (o SparseEnforcerProfilesList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerProfilesList.
func (o SparseEnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerProfilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseEnforcerProfilesList converted to EnforcerProfilesList.
func (o SparseEnforcerProfilesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerProfilesList) Version() int {

	return 1
}

// SparseEnforcerProfile represents the sparse version of a enforcerprofile.
// +k8s:openapi-gen=true
type SparseEnforcerProfile struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Ignore traffic with a source or destination matching the specified
	// interfaces.
	ExcludedInterfaces *[]string `json:"excludedInterfaces,omitempty" msgpack:"excludedInterfaces,omitempty" bson:"excludedinterfaces,omitempty" mapstructure:"excludedInterfaces,omitempty"`

	// Ignore any networks specified here and do not even report any flows.
	// This can be useful for excluding localhost loopback traffic, ignoring
	// traffic to the Kubernetes API, and using Microsegmentation for SSH only.
	ExcludedNetworks *[]string `json:"excludedNetworks,omitempty" msgpack:"excludedNetworks,omitempty" bson:"excludednetworks,omitempty" mapstructure:"excludedNetworks,omitempty"`

	// Frequency that flow report counts are upated.
	FlowReportInterval *string `json:"flowReportInterval,omitempty" msgpack:"flowReportInterval,omitempty" bson:"flowreportinterval,omitempty" mapstructure:"flowReportInterval,omitempty"`

	// A tag expression that identifies processing units to ignore. This can be
	// useful to exclude `kube-system` pods, AWS EC2 agent pods, and third-party
	// agents.
	IgnoreExpression *[][]string `json:"ignoreExpression,omitempty" msgpack:"ignoreExpression,omitempty" bson:"ignoreexpression,omitempty" mapstructure:"ignoreExpression,omitempty"`

	// This field is kept for backward compatibility for enforcers <= 3.5.
	KubernetesMetadataExtractor *EnforcerProfileKubernetesMetadataExtractorValue `json:"kubernetesMetadataExtractor,omitempty" msgpack:"kubernetesMetadataExtractor,omitempty" bson:"kubernetesmetadataextractor,omitempty" mapstructure:"kubernetesMetadataExtractor,omitempty"`

	// This field is kept for backward compatibility for enforcers <= 3.5.
	KubernetesSupportEnabled *bool `json:"kubernetesSupportEnabled,omitempty" msgpack:"kubernetesSupportEnabled,omitempty" bson:"kubernetessupportenabled,omitempty" mapstructure:"kubernetesSupportEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// This field is kept for backward compatibility for enforcers <= 3.5.
	MetadataExtractor *EnforcerProfileMetadataExtractorValue `json:"metadataExtractor,omitempty" msgpack:"metadataExtractor,omitempty" bson:"metadataextractor,omitempty" mapstructure:"metadataExtractor,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Enables syslog functionality of enforcers using this enforcerprofile.
	SyslogEnabled *bool `json:"syslogEnabled,omitempty" msgpack:"syslogEnabled,omitempty" bson:"syslogenabled,omitempty" mapstructure:"syslogEnabled,omitempty"`

	// Contains the remote endpoint to dispatch the syslog messages.
	SyslogEndpoint *string `json:"syslogEndpoint,omitempty" msgpack:"syslogEndpoint,omitempty" bson:"syslogendpoint,omitempty" mapstructure:"syslogEndpoint,omitempty"`

	// PEM-encoded certificate to expose to the clients for TLS.
	SyslogEndpointTLSClientCertificate *string `json:"syslogEndpointTLSClientCertificate,omitempty" msgpack:"syslogEndpointTLSClientCertificate,omitempty" bson:"syslogendpointtlsclientcertificate,omitempty" mapstructure:"syslogEndpointTLSClientCertificate,omitempty"`

	// PEM-encoded server CA certificate.
	SyslogEndpointTLSServerCA *string `json:"syslogEndpointTLSServerCA,omitempty" msgpack:"syslogEndpointTLSServerCA,omitempty" bson:"syslogendpointtlsserverca,omitempty" mapstructure:"syslogEndpointTLSServerCA,omitempty"`

	// Contains the list of supported syslog facilities.
	SyslogFacility *int `json:"syslogFacility,omitempty" msgpack:"syslogFacility,omitempty" bson:"syslogfacility,omitempty" mapstructure:"syslogFacility,omitempty"`

	// Contains the list of supported syslog message format.
	SyslogFormat *EnforcerProfileSyslogFormatValue `json:"syslogFormat,omitempty" msgpack:"syslogFormat,omitempty" bson:"syslogformat,omitempty" mapstructure:"syslogFormat,omitempty"`

	// If empty, the enforcer auto-discovers the TCP networks. Auto-discovery
	// works best in Kubernetes and OpenShift deployments. You may need to manually
	// specify the TCP networks if middle boxes exist that do not comply with
	// [TCP Fast Open RFC 7413](https://tools.ietf.org/html/rfc7413).
	TargetNetworks *[]string `json:"targetNetworks,omitempty" msgpack:"targetNetworks,omitempty" bson:"targetnetworks,omitempty" mapstructure:"targetNetworks,omitempty"`

	// If empty, the enforcer enforces all UDP networks. This works best when all UDP
	// networks have enforcers. If some UDP networks do not have enforcers, you
	// may need to manually specify the UDP networks that should be enforced.
	TargetUDPNetworks *[]string `json:"targetUDPNetworks,omitempty" msgpack:"targetUDPNetworks,omitempty" bson:"targetudpnetworks,omitempty" mapstructure:"targetUDPNetworks,omitempty"`

	// List of trusted certificate authorities. If empty, the main chain of trust
	// will be used.
	TrustedCAs *[]string `json:"trustedCAs,omitempty" msgpack:"trustedCAs,omitempty" bson:"trustedcas,omitempty" mapstructure:"trustedCAs,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerProfile returns a new  SparseEnforcerProfile.
func NewSparseEnforcerProfile() *SparseEnforcerProfile {
	return &SparseEnforcerProfile{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerProfile) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerProfile) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerProfile) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcerProfile{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.ExcludedInterfaces != nil {
		s.ExcludedInterfaces = o.ExcludedInterfaces
	}
	if o.ExcludedNetworks != nil {
		s.ExcludedNetworks = o.ExcludedNetworks
	}
	if o.FlowReportInterval != nil {
		s.FlowReportInterval = o.FlowReportInterval
	}
	if o.IgnoreExpression != nil {
		s.IgnoreExpression = o.IgnoreExpression
	}
	if o.KubernetesMetadataExtractor != nil {
		s.KubernetesMetadataExtractor = o.KubernetesMetadataExtractor
	}
	if o.KubernetesSupportEnabled != nil {
		s.KubernetesSupportEnabled = o.KubernetesSupportEnabled
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.MetadataExtractor != nil {
		s.MetadataExtractor = o.MetadataExtractor
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.SyslogEnabled != nil {
		s.SyslogEnabled = o.SyslogEnabled
	}
	if o.SyslogEndpoint != nil {
		s.SyslogEndpoint = o.SyslogEndpoint
	}
	if o.SyslogEndpointTLSClientCertificate != nil {
		s.SyslogEndpointTLSClientCertificate = o.SyslogEndpointTLSClientCertificate
	}
	if o.SyslogEndpointTLSServerCA != nil {
		s.SyslogEndpointTLSServerCA = o.SyslogEndpointTLSServerCA
	}
	if o.SyslogFacility != nil {
		s.SyslogFacility = o.SyslogFacility
	}
	if o.SyslogFormat != nil {
		s.SyslogFormat = o.SyslogFormat
	}
	if o.TargetNetworks != nil {
		s.TargetNetworks = o.TargetNetworks
	}
	if o.TargetUDPNetworks != nil {
		s.TargetUDPNetworks = o.TargetUDPNetworks
	}
	if o.TrustedCAs != nil {
		s.TrustedCAs = o.TrustedCAs
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerProfile) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcerProfile{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.ExcludedInterfaces != nil {
		o.ExcludedInterfaces = s.ExcludedInterfaces
	}
	if s.ExcludedNetworks != nil {
		o.ExcludedNetworks = s.ExcludedNetworks
	}
	if s.FlowReportInterval != nil {
		o.FlowReportInterval = s.FlowReportInterval
	}
	if s.IgnoreExpression != nil {
		o.IgnoreExpression = s.IgnoreExpression
	}
	if s.KubernetesMetadataExtractor != nil {
		o.KubernetesMetadataExtractor = s.KubernetesMetadataExtractor
	}
	if s.KubernetesSupportEnabled != nil {
		o.KubernetesSupportEnabled = s.KubernetesSupportEnabled
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.MetadataExtractor != nil {
		o.MetadataExtractor = s.MetadataExtractor
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.SyslogEnabled != nil {
		o.SyslogEnabled = s.SyslogEnabled
	}
	if s.SyslogEndpoint != nil {
		o.SyslogEndpoint = s.SyslogEndpoint
	}
	if s.SyslogEndpointTLSClientCertificate != nil {
		o.SyslogEndpointTLSClientCertificate = s.SyslogEndpointTLSClientCertificate
	}
	if s.SyslogEndpointTLSServerCA != nil {
		o.SyslogEndpointTLSServerCA = s.SyslogEndpointTLSServerCA
	}
	if s.SyslogFacility != nil {
		o.SyslogFacility = s.SyslogFacility
	}
	if s.SyslogFormat != nil {
		o.SyslogFormat = s.SyslogFormat
	}
	if s.TargetNetworks != nil {
		o.TargetNetworks = s.TargetNetworks
	}
	if s.TargetUDPNetworks != nil {
		o.TargetUDPNetworks = s.TargetUDPNetworks
	}
	if s.TrustedCAs != nil {
		o.TrustedCAs = s.TrustedCAs
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerProfile) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerProfile) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerProfile()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.ExcludedInterfaces != nil {
		out.ExcludedInterfaces = *o.ExcludedInterfaces
	}
	if o.ExcludedNetworks != nil {
		out.ExcludedNetworks = *o.ExcludedNetworks
	}
	if o.FlowReportInterval != nil {
		out.FlowReportInterval = *o.FlowReportInterval
	}
	if o.IgnoreExpression != nil {
		out.IgnoreExpression = *o.IgnoreExpression
	}
	if o.KubernetesMetadataExtractor != nil {
		out.KubernetesMetadataExtractor = *o.KubernetesMetadataExtractor
	}
	if o.KubernetesSupportEnabled != nil {
		out.KubernetesSupportEnabled = *o.KubernetesSupportEnabled
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MetadataExtractor != nil {
		out.MetadataExtractor = *o.MetadataExtractor
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.SyslogEnabled != nil {
		out.SyslogEnabled = *o.SyslogEnabled
	}
	if o.SyslogEndpoint != nil {
		out.SyslogEndpoint = *o.SyslogEndpoint
	}
	if o.SyslogEndpointTLSClientCertificate != nil {
		out.SyslogEndpointTLSClientCertificate = *o.SyslogEndpointTLSClientCertificate
	}
	if o.SyslogEndpointTLSServerCA != nil {
		out.SyslogEndpointTLSServerCA = *o.SyslogEndpointTLSServerCA
	}
	if o.SyslogFacility != nil {
		out.SyslogFacility = *o.SyslogFacility
	}
	if o.SyslogFormat != nil {
		out.SyslogFormat = *o.SyslogFormat
	}
	if o.TargetNetworks != nil {
		out.TargetNetworks = *o.TargetNetworks
	}
	if o.TargetUDPNetworks != nil {
		out.TargetUDPNetworks = *o.TargetUDPNetworks
	}
	if o.TrustedCAs != nil {
		out.TrustedCAs = *o.TrustedCAs
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseEnforcerProfile) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseEnforcerProfile) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseEnforcerProfile) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseEnforcerProfile) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseEnforcerProfile) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseEnforcerProfile) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseEnforcerProfile) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseEnforcerProfile) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseEnforcerProfile) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseEnforcerProfile) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseEnforcerProfile) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseEnforcerProfile) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseEnforcerProfile) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseEnforcerProfile) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEnforcerProfile) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEnforcerProfile) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEnforcerProfile) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEnforcerProfile.
func (o *SparseEnforcerProfile) DeepCopy() *SparseEnforcerProfile {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerProfile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerProfile.
func (o *SparseEnforcerProfile) DeepCopyInto(out *SparseEnforcerProfile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerProfile: %s", err))
	}

	*out = *target.(*SparseEnforcerProfile)
}

type mongoAttributesEnforcerProfile struct {
	ID                                 bson.ObjectId                                   `bson:"_id,omitempty"`
	Annotations                        map[string][]string                             `bson:"annotations"`
	AssociatedTags                     []string                                        `bson:"associatedtags"`
	CreateIdempotencyKey               string                                          `bson:"createidempotencykey"`
	CreateTime                         time.Time                                       `bson:"createtime"`
	Description                        string                                          `bson:"description"`
	ExcludedInterfaces                 []string                                        `bson:"excludedinterfaces"`
	ExcludedNetworks                   []string                                        `bson:"excludednetworks"`
	FlowReportInterval                 string                                          `bson:"flowreportinterval"`
	IgnoreExpression                   [][]string                                      `bson:"ignoreexpression"`
	KubernetesMetadataExtractor        EnforcerProfileKubernetesMetadataExtractorValue `bson:"kubernetesmetadataextractor"`
	KubernetesSupportEnabled           bool                                            `bson:"kubernetessupportenabled"`
	Metadata                           []string                                        `bson:"metadata"`
	MetadataExtractor                  EnforcerProfileMetadataExtractorValue           `bson:"metadataextractor"`
	MigrationsLog                      map[string]string                               `bson:"migrationslog,omitempty"`
	Name                               string                                          `bson:"name"`
	Namespace                          string                                          `bson:"namespace"`
	NormalizedTags                     []string                                        `bson:"normalizedtags"`
	Propagate                          bool                                            `bson:"propagate"`
	Protected                          bool                                            `bson:"protected"`
	SyslogEnabled                      bool                                            `bson:"syslogenabled"`
	SyslogEndpoint                     string                                          `bson:"syslogendpoint"`
	SyslogEndpointTLSClientCertificate string                                          `bson:"syslogendpointtlsclientcertificate"`
	SyslogEndpointTLSServerCA          string                                          `bson:"syslogendpointtlsserverca"`
	SyslogFacility                     int                                             `bson:"syslogfacility"`
	SyslogFormat                       EnforcerProfileSyslogFormatValue                `bson:"syslogformat"`
	TargetNetworks                     []string                                        `bson:"targetnetworks"`
	TargetUDPNetworks                  []string                                        `bson:"targetudpnetworks"`
	TrustedCAs                         []string                                        `bson:"trustedcas"`
	UpdateIdempotencyKey               string                                          `bson:"updateidempotencykey"`
	UpdateTime                         time.Time                                       `bson:"updatetime"`
	ZHash                              int                                             `bson:"zhash"`
	Zone                               int                                             `bson:"zone"`
}
type mongoAttributesSparseEnforcerProfile struct {
	ID                                 bson.ObjectId                                    `bson:"_id,omitempty"`
	Annotations                        *map[string][]string                             `bson:"annotations,omitempty"`
	AssociatedTags                     *[]string                                        `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey               *string                                          `bson:"createidempotencykey,omitempty"`
	CreateTime                         *time.Time                                       `bson:"createtime,omitempty"`
	Description                        *string                                          `bson:"description,omitempty"`
	ExcludedInterfaces                 *[]string                                        `bson:"excludedinterfaces,omitempty"`
	ExcludedNetworks                   *[]string                                        `bson:"excludednetworks,omitempty"`
	FlowReportInterval                 *string                                          `bson:"flowreportinterval,omitempty"`
	IgnoreExpression                   *[][]string                                      `bson:"ignoreexpression,omitempty"`
	KubernetesMetadataExtractor        *EnforcerProfileKubernetesMetadataExtractorValue `bson:"kubernetesmetadataextractor,omitempty"`
	KubernetesSupportEnabled           *bool                                            `bson:"kubernetessupportenabled,omitempty"`
	Metadata                           *[]string                                        `bson:"metadata,omitempty"`
	MetadataExtractor                  *EnforcerProfileMetadataExtractorValue           `bson:"metadataextractor,omitempty"`
	MigrationsLog                      *map[string]string                               `bson:"migrationslog,omitempty"`
	Name                               *string                                          `bson:"name,omitempty"`
	Namespace                          *string                                          `bson:"namespace,omitempty"`
	NormalizedTags                     *[]string                                        `bson:"normalizedtags,omitempty"`
	Propagate                          *bool                                            `bson:"propagate,omitempty"`
	Protected                          *bool                                            `bson:"protected,omitempty"`
	SyslogEnabled                      *bool                                            `bson:"syslogenabled,omitempty"`
	SyslogEndpoint                     *string                                          `bson:"syslogendpoint,omitempty"`
	SyslogEndpointTLSClientCertificate *string                                          `bson:"syslogendpointtlsclientcertificate,omitempty"`
	SyslogEndpointTLSServerCA          *string                                          `bson:"syslogendpointtlsserverca,omitempty"`
	SyslogFacility                     *int                                             `bson:"syslogfacility,omitempty"`
	SyslogFormat                       *EnforcerProfileSyslogFormatValue                `bson:"syslogformat,omitempty"`
	TargetNetworks                     *[]string                                        `bson:"targetnetworks,omitempty"`
	TargetUDPNetworks                  *[]string                                        `bson:"targetudpnetworks,omitempty"`
	TrustedCAs                         *[]string                                        `bson:"trustedcas,omitempty"`
	UpdateIdempotencyKey               *string                                          `bson:"updateidempotencykey,omitempty"`
	UpdateTime                         *time.Time                                       `bson:"updatetime,omitempty"`
	ZHash                              *int                                             `bson:"zhash,omitempty"`
	Zone                               *int                                             `bson:"zone,omitempty"`
}
