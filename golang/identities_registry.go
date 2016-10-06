package gaia

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(MyNamespaceIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(LayerIdentity)
	elemental.RegisterIdentity(DependencyMapSubviewIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(SystemInfoIdentity)
	elemental.RegisterIdentity(ImageIdentity)
	elemental.RegisterIdentity(ComputedPolicyIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(MapEdgeIdentity)
	elemental.RegisterIdentity(VulnerabilityFixedInIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(NotificationIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(IntegrationIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(VulnerabilityWithLayersIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(ClairVulnerabilityIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(AuthenticatorIdentity)
	elemental.RegisterIdentity(ComputedDependencyMapViewIdentity)
	elemental.RegisterIdentity(UserIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(ClairNotificationIdentity)
	elemental.RegisterIdentity(DependencyMapViewIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(ServerIdentity)
	elemental.RegisterIdentity(MapNodeIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
}
