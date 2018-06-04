package gaia

import "github.com/aporeto-inc/elemental"

func init() {
	elemental.RegisterIdentity(AccountIdentity)
	elemental.RegisterIdentity(AccountCheckIdentity)
	elemental.RegisterIdentity(ActivateIdentity)
	elemental.RegisterIdentity(ActivityIdentity)
	elemental.RegisterIdentity(AlarmIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(APICheckIdentity)
	elemental.RegisterIdentity(AppIdentity)
	elemental.RegisterIdentity(AuditProfileIdentity)
	elemental.RegisterIdentity(AuthIdentity)
	elemental.RegisterIdentity(AuthorityIdentity)
	elemental.RegisterIdentity(AutomationIdentity)
	elemental.RegisterIdentity(AutomationTemplateIdentity)
	elemental.RegisterIdentity(AWSAccountIdentity)
	elemental.RegisterIdentity(CategoryIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(EmailIdentity)
	elemental.RegisterIdentity(EnforcerIdentity)
	elemental.RegisterIdentity(EnforcerProfileIdentity)
	elemental.RegisterIdentity(EnforcerProfileMappingPolicyIdentity)
	elemental.RegisterIdentity(ExportIdentity)
	elemental.RegisterIdentity(ExternalAccessIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(HookPolicyIdentity)
	elemental.RegisterIdentity(ImportIdentity)
	elemental.RegisterIdentity(InstallationIdentity)
	elemental.RegisterIdentity(InstalledAppIdentity)
	elemental.RegisterIdentity(IsolationProfileIdentity)
	elemental.RegisterIdentity(IssueIdentity)
	elemental.RegisterIdentity(JaegerbatchIdentity)
	elemental.RegisterIdentity(K8SClusterIdentity)
	elemental.RegisterIdentity(KubernetesClusterIdentity)
	elemental.RegisterIdentity(LogIdentity)
	elemental.RegisterIdentity(MessageIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
	elemental.RegisterIdentity(PasswordResetIdentity)
	elemental.RegisterIdentity(PlanIdentity)
	elemental.RegisterIdentity(PokeIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(PolicyRefreshIdentity)
	elemental.RegisterIdentity(PolicyRendererIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(PrivateKeyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(ProcessingUnitPolicyIdentity)
	elemental.RegisterIdentity(QuotaCheckIdentity)
	elemental.RegisterIdentity(QuotaPolicyIdentity)
	elemental.RegisterIdentity(RemoteProcessorIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(ReportIdentity)
	elemental.RegisterIdentity(RESTAPISpecIdentity)
	elemental.RegisterIdentity(RevocationIdentity)
	elemental.RegisterIdentity(RoleIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(ServiceIdentity)
	elemental.RegisterIdentity(StatsQueryIdentity)
	elemental.RegisterIdentity(SuggestedPolicyIdentity)
	elemental.RegisterIdentity(TabulationIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(TagInjectIdentity)
	elemental.RegisterIdentity(TokenIdentity)
	elemental.RegisterIdentity(TokenScopePolicyIdentity)
	elemental.RegisterIdentity(TriggerIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(X509CertificateIdentity)
	elemental.RegisterIdentity(X509CertificateCheckIdentity)
}

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {

	case AccountIdentity.Name:
		return NewAccount()
	case AccountCheckIdentity.Name:
		return NewAccountCheck()
	case ActivateIdentity.Name:
		return NewActivate()
	case ActivityIdentity.Name:
		return NewActivity()
	case AlarmIdentity.Name:
		return NewAlarm()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case APICheckIdentity.Name:
		return NewAPICheck()
	case AppIdentity.Name:
		return NewApp()
	case AuditProfileIdentity.Name:
		return NewAuditProfile()
	case AuthIdentity.Name:
		return NewAuth()
	case AuthorityIdentity.Name:
		return NewAuthority()
	case AutomationIdentity.Name:
		return NewAutomation()
	case AutomationTemplateIdentity.Name:
		return NewAutomationTemplate()
	case AWSAccountIdentity.Name:
		return NewAWSAccount()
	case CategoryIdentity.Name:
		return NewCategory()
	case CertificateIdentity.Name:
		return NewCertificate()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case EmailIdentity.Name:
		return NewEmail()
	case EnforcerIdentity.Name:
		return NewEnforcer()
	case EnforcerProfileIdentity.Name:
		return NewEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity.Name:
		return NewEnforcerProfileMappingPolicy()
	case ExportIdentity.Name:
		return NewExport()
	case ExternalAccessIdentity.Name:
		return NewExternalAccess()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case FileAccessIdentity.Name:
		return NewFileAccess()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case FilePathIdentity.Name:
		return NewFilePath()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case HookPolicyIdentity.Name:
		return NewHookPolicy()
	case ImportIdentity.Name:
		return NewImport()
	case InstallationIdentity.Name:
		return NewInstallation()
	case InstalledAppIdentity.Name:
		return NewInstalledApp()
	case IsolationProfileIdentity.Name:
		return NewIsolationProfile()
	case IssueIdentity.Name:
		return NewIssue()
	case JaegerbatchIdentity.Name:
		return NewJaegerbatch()
	case K8SClusterIdentity.Name:
		return NewK8SCluster()
	case KubernetesClusterIdentity.Name:
		return NewKubernetesCluster()
	case LogIdentity.Name:
		return NewLog()
	case MessageIdentity.Name:
		return NewMessage()
	case NamespaceIdentity.Name:
		return NewNamespace()
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case NetworkAccessPolicyIdentity.Name:
		return NewNetworkAccessPolicy()
	case PasswordResetIdentity.Name:
		return NewPasswordReset()
	case PlanIdentity.Name:
		return NewPlan()
	case PokeIdentity.Name:
		return NewPoke()
	case PolicyIdentity.Name:
		return NewPolicy()
	case PolicyRefreshIdentity.Name:
		return NewPolicyRefresh()
	case PolicyRendererIdentity.Name:
		return NewPolicyRenderer()
	case PolicyRuleIdentity.Name:
		return NewPolicyRule()
	case PrivateKeyIdentity.Name:
		return NewPrivateKey()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case ProcessingUnitPolicyIdentity.Name:
		return NewProcessingUnitPolicy()
	case QuotaCheckIdentity.Name:
		return NewQuotaCheck()
	case QuotaPolicyIdentity.Name:
		return NewQuotaPolicy()
	case RemoteProcessorIdentity.Name:
		return NewRemoteProcessor()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case ReportIdentity.Name:
		return NewReport()
	case RESTAPISpecIdentity.Name:
		return NewRESTAPISpec()
	case RevocationIdentity.Name:
		return NewRevocation()
	case RoleIdentity.Name:
		return NewRole()
	case RootIdentity.Name:
		return NewRoot()
	case ServiceIdentity.Name:
		return NewService()
	case StatsQueryIdentity.Name:
		return NewStatsQuery()
	case SuggestedPolicyIdentity.Name:
		return NewSuggestedPolicy()
	case TabulationIdentity.Name:
		return NewTabulation()
	case TagIdentity.Name:
		return NewTag()
	case TagInjectIdentity.Name:
		return NewTagInject()
	case TokenIdentity.Name:
		return NewToken()
	case TokenScopePolicyIdentity.Name:
		return NewTokenScopePolicy()
	case TriggerIdentity.Name:
		return NewTrigger()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	case X509CertificateIdentity.Name:
		return NewX509Certificate()
	case X509CertificateCheckIdentity.Name:
		return NewX509CertificateCheck()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {

	case AccountIdentity.Category:
		return NewAccount()
	case AccountCheckIdentity.Category:
		return NewAccountCheck()
	case ActivateIdentity.Category:
		return NewActivate()
	case ActivityIdentity.Category:
		return NewActivity()
	case AlarmIdentity.Category:
		return NewAlarm()
	case APIAuthorizationPolicyIdentity.Category:
		return NewAPIAuthorizationPolicy()
	case APICheckIdentity.Category:
		return NewAPICheck()
	case AppIdentity.Category:
		return NewApp()
	case AuditProfileIdentity.Category:
		return NewAuditProfile()
	case AuthIdentity.Category:
		return NewAuth()
	case AuthorityIdentity.Category:
		return NewAuthority()
	case AutomationIdentity.Category:
		return NewAutomation()
	case AutomationTemplateIdentity.Category:
		return NewAutomationTemplate()
	case AWSAccountIdentity.Category:
		return NewAWSAccount()
	case CategoryIdentity.Category:
		return NewCategory()
	case CertificateIdentity.Category:
		return NewCertificate()
	case DependencyMapIdentity.Category:
		return NewDependencyMap()
	case EmailIdentity.Category:
		return NewEmail()
	case EnforcerIdentity.Category:
		return NewEnforcer()
	case EnforcerProfileIdentity.Category:
		return NewEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity.Category:
		return NewEnforcerProfileMappingPolicy()
	case ExportIdentity.Category:
		return NewExport()
	case ExternalAccessIdentity.Category:
		return NewExternalAccess()
	case ExternalServiceIdentity.Category:
		return NewExternalService()
	case FileAccessIdentity.Category:
		return NewFileAccess()
	case FileAccessPolicyIdentity.Category:
		return NewFileAccessPolicy()
	case FilePathIdentity.Category:
		return NewFilePath()
	case FlowStatisticIdentity.Category:
		return NewFlowStatistic()
	case HookPolicyIdentity.Category:
		return NewHookPolicy()
	case ImportIdentity.Category:
		return NewImport()
	case InstallationIdentity.Category:
		return NewInstallation()
	case InstalledAppIdentity.Category:
		return NewInstalledApp()
	case IsolationProfileIdentity.Category:
		return NewIsolationProfile()
	case IssueIdentity.Category:
		return NewIssue()
	case JaegerbatchIdentity.Category:
		return NewJaegerbatch()
	case K8SClusterIdentity.Category:
		return NewK8SCluster()
	case KubernetesClusterIdentity.Category:
		return NewKubernetesCluster()
	case LogIdentity.Category:
		return NewLog()
	case MessageIdentity.Category:
		return NewMessage()
	case NamespaceIdentity.Category:
		return NewNamespace()
	case NamespaceMappingPolicyIdentity.Category:
		return NewNamespaceMappingPolicy()
	case NetworkAccessPolicyIdentity.Category:
		return NewNetworkAccessPolicy()
	case PasswordResetIdentity.Category:
		return NewPasswordReset()
	case PlanIdentity.Category:
		return NewPlan()
	case PokeIdentity.Category:
		return NewPoke()
	case PolicyIdentity.Category:
		return NewPolicy()
	case PolicyRefreshIdentity.Category:
		return NewPolicyRefresh()
	case PolicyRendererIdentity.Category:
		return NewPolicyRenderer()
	case PolicyRuleIdentity.Category:
		return NewPolicyRule()
	case PrivateKeyIdentity.Category:
		return NewPrivateKey()
	case ProcessingUnitIdentity.Category:
		return NewProcessingUnit()
	case ProcessingUnitPolicyIdentity.Category:
		return NewProcessingUnitPolicy()
	case QuotaCheckIdentity.Category:
		return NewQuotaCheck()
	case QuotaPolicyIdentity.Category:
		return NewQuotaPolicy()
	case RemoteProcessorIdentity.Category:
		return NewRemoteProcessor()
	case RenderedPolicyIdentity.Category:
		return NewRenderedPolicy()
	case ReportIdentity.Category:
		return NewReport()
	case RESTAPISpecIdentity.Category:
		return NewRESTAPISpec()
	case RevocationIdentity.Category:
		return NewRevocation()
	case RoleIdentity.Category:
		return NewRole()
	case RootIdentity.Category:
		return NewRoot()
	case ServiceIdentity.Category:
		return NewService()
	case StatsQueryIdentity.Category:
		return NewStatsQuery()
	case SuggestedPolicyIdentity.Category:
		return NewSuggestedPolicy()
	case TabulationIdentity.Category:
		return NewTabulation()
	case TagIdentity.Category:
		return NewTag()
	case TagInjectIdentity.Category:
		return NewTagInject()
	case TokenIdentity.Category:
		return NewToken()
	case TokenScopePolicyIdentity.Category:
		return NewTokenScopePolicy()
	case TriggerIdentity.Category:
		return NewTrigger()
	case VulnerabilityIdentity.Category:
		return NewVulnerability()
	case X509CertificateIdentity.Category:
		return NewX509Certificate()
	case X509CertificateCheckIdentity.Category:
		return NewX509CertificateCheck()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {

	case AccountIdentity.Name:
		return &AccountsList{}
	case AccountCheckIdentity.Name:
		return &AccountChecksList{}
	case ActivateIdentity.Name:
		return &ActivatesList{}
	case ActivityIdentity.Name:
		return &ActivitiesList{}
	case AlarmIdentity.Name:
		return &AlarmsList{}
	case APIAuthorizationPolicyIdentity.Name:
		return &APIAuthorizationPoliciesList{}
	case APICheckIdentity.Name:
		return &APIChecksList{}
	case AppIdentity.Name:
		return &AppsList{}
	case AuditProfileIdentity.Name:
		return &AuditProfilesList{}
	case AuthIdentity.Name:
		return &AuthsList{}
	case AuthorityIdentity.Name:
		return &AuthoritiesList{}
	case AutomationIdentity.Name:
		return &AutomationsList{}
	case AutomationTemplateIdentity.Name:
		return &AutomationTemplatesList{}
	case AWSAccountIdentity.Name:
		return &AWSAccountsList{}
	case CategoryIdentity.Name:
		return &CategoriesList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	case DependencyMapIdentity.Name:
		return &DependencyMapsList{}
	case EmailIdentity.Name:
		return &EmailsList{}
	case EnforcerIdentity.Name:
		return &EnforcersList{}
	case EnforcerProfileIdentity.Name:
		return &EnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity.Name:
		return &EnforcerProfileMappingPoliciesList{}
	case ExportIdentity.Name:
		return &ExportsList{}
	case ExternalAccessIdentity.Name:
		return &ExternalAccessList{}
	case ExternalServiceIdentity.Name:
		return &ExternalServicesList{}
	case FileAccessIdentity.Name:
		return &FileAccessList{}
	case FileAccessPolicyIdentity.Name:
		return &FileAccessPoliciesList{}
	case FilePathIdentity.Name:
		return &FilePathsList{}
	case FlowStatisticIdentity.Name:
		return &FlowStatisticsList{}
	case HookPolicyIdentity.Name:
		return &HookPoliciesList{}
	case ImportIdentity.Name:
		return &ImportsList{}
	case InstallationIdentity.Name:
		return &InstallationsList{}
	case InstalledAppIdentity.Name:
		return &InstalledAppsList{}
	case IsolationProfileIdentity.Name:
		return &IsolationProfilesList{}
	case IssueIdentity.Name:
		return &IssuesList{}
	case JaegerbatchIdentity.Name:
		return &JaegerbatchsList{}
	case K8SClusterIdentity.Name:
		return &K8SClustersList{}
	case KubernetesClusterIdentity.Name:
		return &KubernetesClustersList{}
	case LogIdentity.Name:
		return &LogsList{}
	case MessageIdentity.Name:
		return &MessagesList{}
	case NamespaceIdentity.Name:
		return &NamespacesList{}
	case NamespaceMappingPolicyIdentity.Name:
		return &NamespaceMappingPoliciesList{}
	case NetworkAccessPolicyIdentity.Name:
		return &NetworkAccessPoliciesList{}
	case PasswordResetIdentity.Name:
		return &PasswordResetsList{}
	case PlanIdentity.Name:
		return &PlansList{}
	case PokeIdentity.Name:
		return &PokesList{}
	case PolicyIdentity.Name:
		return &PoliciesList{}
	case PolicyRefreshIdentity.Name:
		return &PolicyRefreshsList{}
	case PolicyRendererIdentity.Name:
		return &PolicyRenderersList{}
	case PolicyRuleIdentity.Name:
		return &PolicyRulesList{}
	case PrivateKeyIdentity.Name:
		return &PrivateKeysList{}
	case ProcessingUnitIdentity.Name:
		return &ProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity.Name:
		return &ProcessingUnitPoliciesList{}
	case QuotaCheckIdentity.Name:
		return &QuotaChecksList{}
	case QuotaPolicyIdentity.Name:
		return &QuotaPoliciesList{}
	case RemoteProcessorIdentity.Name:
		return &RemoteProcessorsList{}
	case RenderedPolicyIdentity.Name:
		return &RenderedPoliciesList{}
	case ReportIdentity.Name:
		return &ReportsList{}
	case RESTAPISpecIdentity.Name:
		return &RESTAPISpecsList{}
	case RevocationIdentity.Name:
		return &RevocationsList{}
	case RoleIdentity.Name:
		return &RolesList{}
	case ServiceIdentity.Name:
		return &ServicesList{}
	case StatsQueryIdentity.Name:
		return &StatsQueriesList{}
	case SuggestedPolicyIdentity.Name:
		return &SuggestedPoliciesList{}
	case TabulationIdentity.Name:
		return &TabulationsList{}
	case TagIdentity.Name:
		return &TagsList{}
	case TagInjectIdentity.Name:
		return &TagInjectsList{}
	case TokenIdentity.Name:
		return &TokensList{}
	case TokenScopePolicyIdentity.Name:
		return &TokenScopePoliciesList{}
	case TriggerIdentity.Name:
		return &TriggersList{}
	case VulnerabilityIdentity.Name:
		return &VulnerabilitiesList{}
	case X509CertificateIdentity.Name:
		return &X509CertificatesList{}
	case X509CertificateCheckIdentity.Name:
		return &X509CertificateChecksList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {

	case AccountIdentity.Category:
		return &AccountsList{}
	case AccountCheckIdentity.Category:
		return &AccountChecksList{}
	case ActivateIdentity.Category:
		return &ActivatesList{}
	case ActivityIdentity.Category:
		return &ActivitiesList{}
	case AlarmIdentity.Category:
		return &AlarmsList{}
	case APIAuthorizationPolicyIdentity.Category:
		return &APIAuthorizationPoliciesList{}
	case APICheckIdentity.Category:
		return &APIChecksList{}
	case AppIdentity.Category:
		return &AppsList{}
	case AuditProfileIdentity.Category:
		return &AuditProfilesList{}
	case AuthIdentity.Category:
		return &AuthsList{}
	case AuthorityIdentity.Category:
		return &AuthoritiesList{}
	case AutomationIdentity.Category:
		return &AutomationsList{}
	case AutomationTemplateIdentity.Category:
		return &AutomationTemplatesList{}
	case AWSAccountIdentity.Category:
		return &AWSAccountsList{}
	case CategoryIdentity.Category:
		return &CategoriesList{}
	case CertificateIdentity.Category:
		return &CertificatesList{}
	case DependencyMapIdentity.Category:
		return &DependencyMapsList{}
	case EmailIdentity.Category:
		return &EmailsList{}
	case EnforcerIdentity.Category:
		return &EnforcersList{}
	case EnforcerProfileIdentity.Category:
		return &EnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity.Category:
		return &EnforcerProfileMappingPoliciesList{}
	case ExportIdentity.Category:
		return &ExportsList{}
	case ExternalAccessIdentity.Category:
		return &ExternalAccessList{}
	case ExternalServiceIdentity.Category:
		return &ExternalServicesList{}
	case FileAccessIdentity.Category:
		return &FileAccessList{}
	case FileAccessPolicyIdentity.Category:
		return &FileAccessPoliciesList{}
	case FilePathIdentity.Category:
		return &FilePathsList{}
	case FlowStatisticIdentity.Category:
		return &FlowStatisticsList{}
	case HookPolicyIdentity.Category:
		return &HookPoliciesList{}
	case ImportIdentity.Category:
		return &ImportsList{}
	case InstallationIdentity.Category:
		return &InstallationsList{}
	case InstalledAppIdentity.Category:
		return &InstalledAppsList{}
	case IsolationProfileIdentity.Category:
		return &IsolationProfilesList{}
	case IssueIdentity.Category:
		return &IssuesList{}
	case JaegerbatchIdentity.Category:
		return &JaegerbatchsList{}
	case K8SClusterIdentity.Category:
		return &K8SClustersList{}
	case KubernetesClusterIdentity.Category:
		return &KubernetesClustersList{}
	case LogIdentity.Category:
		return &LogsList{}
	case MessageIdentity.Category:
		return &MessagesList{}
	case NamespaceIdentity.Category:
		return &NamespacesList{}
	case NamespaceMappingPolicyIdentity.Category:
		return &NamespaceMappingPoliciesList{}
	case NetworkAccessPolicyIdentity.Category:
		return &NetworkAccessPoliciesList{}
	case PasswordResetIdentity.Category:
		return &PasswordResetsList{}
	case PlanIdentity.Category:
		return &PlansList{}
	case PokeIdentity.Category:
		return &PokesList{}
	case PolicyIdentity.Category:
		return &PoliciesList{}
	case PolicyRefreshIdentity.Category:
		return &PolicyRefreshsList{}
	case PolicyRendererIdentity.Category:
		return &PolicyRenderersList{}
	case PolicyRuleIdentity.Category:
		return &PolicyRulesList{}
	case PrivateKeyIdentity.Category:
		return &PrivateKeysList{}
	case ProcessingUnitIdentity.Category:
		return &ProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity.Category:
		return &ProcessingUnitPoliciesList{}
	case QuotaCheckIdentity.Category:
		return &QuotaChecksList{}
	case QuotaPolicyIdentity.Category:
		return &QuotaPoliciesList{}
	case RemoteProcessorIdentity.Category:
		return &RemoteProcessorsList{}
	case RenderedPolicyIdentity.Category:
		return &RenderedPoliciesList{}
	case ReportIdentity.Category:
		return &ReportsList{}
	case RESTAPISpecIdentity.Category:
		return &RESTAPISpecsList{}
	case RevocationIdentity.Category:
		return &RevocationsList{}
	case RoleIdentity.Category:
		return &RolesList{}
	case ServiceIdentity.Category:
		return &ServicesList{}
	case StatsQueryIdentity.Category:
		return &StatsQueriesList{}
	case SuggestedPolicyIdentity.Category:
		return &SuggestedPoliciesList{}
	case TabulationIdentity.Category:
		return &TabulationsList{}
	case TagIdentity.Category:
		return &TagsList{}
	case TagInjectIdentity.Category:
		return &TagInjectsList{}
	case TokenIdentity.Category:
		return &TokensList{}
	case TokenScopePolicyIdentity.Category:
		return &TokenScopePoliciesList{}
	case TriggerIdentity.Category:
		return &TriggersList{}
	case VulnerabilityIdentity.Category:
		return &VulnerabilitiesList{}
	case X509CertificateIdentity.Category:
		return &X509CertificatesList{}
	case X509CertificateCheckIdentity.Category:
		return &X509CertificateChecksList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AccountIdentity,
		AccountCheckIdentity,
		ActivateIdentity,
		ActivityIdentity,
		AlarmIdentity,
		APIAuthorizationPolicyIdentity,
		APICheckIdentity,
		AppIdentity,
		AuditProfileIdentity,
		AuthIdentity,
		AuthorityIdentity,
		AutomationIdentity,
		AutomationTemplateIdentity,
		AWSAccountIdentity,
		CategoryIdentity,
		CertificateIdentity,
		DependencyMapIdentity,
		EmailIdentity,
		EnforcerIdentity,
		EnforcerProfileIdentity,
		EnforcerProfileMappingPolicyIdentity,
		ExportIdentity,
		ExternalAccessIdentity,
		ExternalServiceIdentity,
		FileAccessIdentity,
		FileAccessPolicyIdentity,
		FilePathIdentity,
		FlowStatisticIdentity,
		HookPolicyIdentity,
		ImportIdentity,
		InstallationIdentity,
		InstalledAppIdentity,
		IsolationProfileIdentity,
		IssueIdentity,
		JaegerbatchIdentity,
		K8SClusterIdentity,
		KubernetesClusterIdentity,
		LogIdentity,
		MessageIdentity,
		NamespaceIdentity,
		NamespaceMappingPolicyIdentity,
		NetworkAccessPolicyIdentity,
		PasswordResetIdentity,
		PlanIdentity,
		PokeIdentity,
		PolicyIdentity,
		PolicyRefreshIdentity,
		PolicyRendererIdentity,
		PolicyRuleIdentity,
		PrivateKeyIdentity,
		ProcessingUnitIdentity,
		ProcessingUnitPolicyIdentity,
		QuotaCheckIdentity,
		QuotaPolicyIdentity,
		RemoteProcessorIdentity,
		RenderedPolicyIdentity,
		ReportIdentity,
		RESTAPISpecIdentity,
		RevocationIdentity,
		RoleIdentity,
		RootIdentity,
		ServiceIdentity,
		StatsQueryIdentity,
		SuggestedPolicyIdentity,
		TabulationIdentity,
		TagIdentity,
		TagInjectIdentity,
		TokenIdentity,
		TokenScopePolicyIdentity,
		TriggerIdentity,
		VulnerabilityIdentity,
		X509CertificateIdentity,
		X509CertificateCheckIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"apiauth":    APIAuthorizationPolicyIdentity,
	"apiauths":   APIAuthorizationPolicyIdentity,
	"ap":         AuditProfileIdentity,
	"ca":         AuthorityIdentity,
	"autos":      AutomationIdentity,
	"auto":       AutomationIdentity,
	"autotmpl":   AutomationTemplateIdentity,
	"aws":        AWSAccountIdentity,
	"awsaccs":    AWSAccountIdentity,
	"awsacc":     AWSAccountIdentity,
	"depmaps":    DependencyMapIdentity,
	"depmap":     DependencyMapIdentity,
	"profile":    EnforcerProfileIdentity,
	"profiles":   EnforcerProfileIdentity,
	"enfpols":    EnforcerProfileMappingPolicyIdentity,
	"enfpol":     EnforcerProfileMappingPolicyIdentity,
	"extacs":     ExternalAccessIdentity,
	"extac":      ExternalAccessIdentity,
	"extsrv":     ExternalServiceIdentity,
	"extsrvs":    ExternalServiceIdentity,
	"fp":         FilePathIdentity,
	"fps":        FilePathIdentity,
	"flowstat":   FlowStatisticIdentity,
	"flowstats":  FlowStatisticIdentity,
	"hook":       HookPolicyIdentity,
	"hooks":      HookPolicyIdentity,
	"hookpol":    HookPolicyIdentity,
	"hookpols":   HookPolicyIdentity,
	"iapps":      InstalledAppIdentity,
	"iapp":       InstalledAppIdentity,
	"ip":         IsolationProfileIdentity,
	"sp":         JaegerbatchIdentity,
	"mess":       MessageIdentity,
	"ns":         NamespaceIdentity,
	"nspolicy":   NamespaceMappingPolicyIdentity,
	"nspolicies": NamespaceMappingPolicyIdentity,
	"nsmap":      NamespaceMappingPolicyIdentity,
	"nsmaps":     NamespaceMappingPolicyIdentity,
	"netpol":     NetworkAccessPolicyIdentity,
	"netpols":    NetworkAccessPolicyIdentity,
	"pu":         ProcessingUnitIdentity,
	"pus":        ProcessingUnitIdentity,
	"pup":        ProcessingUnitPolicyIdentity,
	"quota":      QuotaPolicyIdentity,
	"quotas":     QuotaPolicyIdentity,
	"quotapol":   QuotaPolicyIdentity,
	"quotapols":  QuotaPolicyIdentity,
	"hks":        RemoteProcessorIdentity,
	"hk":         RemoteProcessorIdentity,
	"rpol":       RenderedPolicyIdentity,
	"rpols":      RenderedPolicyIdentity,
	"srv":        ServiceIdentity,
	"sq":         StatsQueryIdentity,
	"sugpol":     SuggestedPolicyIdentity,
	"sugpols":    SuggestedPolicyIdentity,
	"sugg":       SuggestedPolicyIdentity,
	"suggs":      SuggestedPolicyIdentity,
	"table":      TabulationIdentity,
	"tables":     TabulationIdentity,
	"tabs":       TabulationIdentity,
	"tab":        TabulationIdentity,
	"tsp":        TokenScopePolicyIdentity,
	"vulns":      VulnerabilityIdentity,
	"vul":        VulnerabilityIdentity,
	"vuln":       VulnerabilityIdentity,
	"vuls":       VulnerabilityIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case AccountIdentity:
		return []string{}
	case AccountCheckIdentity:
		return []string{}
	case ActivateIdentity:
		return []string{}
	case ActivityIdentity:
		return []string{}
	case AlarmIdentity:
		return []string{}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauth",
			"apiauths",
		}
	case APICheckIdentity:
		return []string{}
	case AppIdentity:
		return []string{}
	case AuditProfileIdentity:
		return []string{
			"ap",
		}
	case AuthIdentity:
		return []string{}
	case AuthorityIdentity:
		return []string{
			"ca",
		}
	case AutomationIdentity:
		return []string{
			"autos",
			"auto",
		}
	case AutomationTemplateIdentity:
		return []string{
			"autotmpl",
		}
	case AWSAccountIdentity:
		return []string{
			"aws",
			"awsaccs",
			"awsacc",
		}
	case CategoryIdentity:
		return []string{}
	case CertificateIdentity:
		return []string{}
	case DependencyMapIdentity:
		return []string{
			"depmaps",
			"depmap",
		}
	case EmailIdentity:
		return []string{}
	case EnforcerIdentity:
		return []string{}
	case EnforcerProfileIdentity:
		return []string{
			"profile",
			"profiles",
		}
	case EnforcerProfileMappingPolicyIdentity:
		return []string{
			"enfpols",
			"enfpol",
		}
	case ExportIdentity:
		return []string{}
	case ExternalAccessIdentity:
		return []string{
			"extacs",
			"extac",
		}
	case ExternalServiceIdentity:
		return []string{
			"extsrv",
			"extsrvs",
		}
	case FileAccessIdentity:
		return []string{}
	case FileAccessPolicyIdentity:
		return []string{}
	case FilePathIdentity:
		return []string{
			"fp",
			"fps",
		}
	case FlowStatisticIdentity:
		return []string{
			"flowstat",
			"flowstats",
		}
	case HookPolicyIdentity:
		return []string{
			"hook",
			"hooks",
			"hookpol",
			"hookpols",
		}
	case ImportIdentity:
		return []string{}
	case InstallationIdentity:
		return []string{}
	case InstalledAppIdentity:
		return []string{
			"iapps",
			"iapp",
		}
	case IsolationProfileIdentity:
		return []string{
			"ip",
		}
	case IssueIdentity:
		return []string{}
	case JaegerbatchIdentity:
		return []string{
			"sp",
		}
	case K8SClusterIdentity:
		return []string{}
	case KubernetesClusterIdentity:
		return []string{}
	case LogIdentity:
		return []string{}
	case MessageIdentity:
		return []string{
			"mess",
		}
	case NamespaceIdentity:
		return []string{
			"ns",
		}
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nspolicy",
			"nspolicies",
			"nsmap",
			"nsmaps",
		}
	case NetworkAccessPolicyIdentity:
		return []string{
			"netpol",
			"netpols",
		}
	case PasswordResetIdentity:
		return []string{}
	case PlanIdentity:
		return []string{}
	case PokeIdentity:
		return []string{}
	case PolicyIdentity:
		return []string{}
	case PolicyRefreshIdentity:
		return []string{}
	case PolicyRendererIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case PrivateKeyIdentity:
		return []string{}
	case ProcessingUnitIdentity:
		return []string{
			"pu",
			"pus",
		}
	case ProcessingUnitPolicyIdentity:
		return []string{
			"pup",
		}
	case QuotaCheckIdentity:
		return []string{}
	case QuotaPolicyIdentity:
		return []string{
			"quota",
			"quotas",
			"quotapol",
			"quotapols",
		}
	case RemoteProcessorIdentity:
		return []string{
			"hks",
			"hk",
		}
	case RenderedPolicyIdentity:
		return []string{
			"rpol",
			"rpols",
		}
	case ReportIdentity:
		return []string{}
	case RESTAPISpecIdentity:
		return []string{}
	case RevocationIdentity:
		return []string{}
	case RoleIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case ServiceIdentity:
		return []string{
			"srv",
		}
	case StatsQueryIdentity:
		return []string{
			"sq",
		}
	case SuggestedPolicyIdentity:
		return []string{
			"sugpol",
			"sugpols",
			"sugg",
			"suggs",
		}
	case TabulationIdentity:
		return []string{
			"table",
			"tables",
			"tabs",
			"tab",
		}
	case TagIdentity:
		return []string{}
	case TagInjectIdentity:
		return []string{}
	case TokenIdentity:
		return []string{}
	case TokenScopePolicyIdentity:
		return []string{
			"tsp",
		}
	case TriggerIdentity:
		return []string{}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
			"vuln",
			"vuls",
		}
	case X509CertificateIdentity:
		return []string{}
	case X509CertificateCheckIdentity:
		return []string{}
	}

	return nil
}
