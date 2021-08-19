package gaia

import "go.aporeto.io/elemental"

var (
	identityNamesMap = map[string]elemental.Identity{
		"accessiblenamespace":       AccessibleNamespaceIdentity,
		"accessreport":              AccessReportIdentity,
		"account":                   AccountIdentity,
		"accountcheck":              AccountCheckIdentity,
		"activate":                  ActivateIdentity,
		"activity":                  ActivityIdentity,
		"alarm":                     AlarmIdentity,
		"apiauthorizationpolicy":    APIAuthorizationPolicyIdentity,
		"apicheck":                  APICheckIdentity,
		"app":                       AppIdentity,
		"appcredential":             AppCredentialIdentity,
		"auditprofile":              AuditProfileIdentity,
		"auditprofilemappingpolicy": AuditProfileMappingPolicyIdentity,
		"auditreport":               AuditReportIdentity,
		"authn":                     AuthnIdentity,
		"authority":                 AuthorityIdentity,
		"authz":                     AuthzIdentity,
		"automation":                AutomationIdentity,
		"automationtemplate":        AutomationTemplateIdentity,
		"cachedflowreport":          CachedFlowReportIdentity,
		"category":                  CategoryIdentity,

		"claims":              ClaimsIdentity,
		"clausesmatch":        ClauseMatchIdentity,
		"cloudaccountcleaner": CloudAccountCleanerIdentity,

		"cloudalert":     CloudAlertIdentity,
		"cloudalertrule": CloudAlertRuleIdentity,
		"cloudendpoint":  CloudEndpointIdentity,

		"cloudgraph": CloudGraphIdentity,

		"cloudmanagednetwork":   CloudManagedNetworkIdentity,
		"cloudnetworkinterface": CloudNetworkInterfaceIdentity,
		"cloudnetworkquery":     CloudNetworkQueryIdentity,

		"cloudnetworkruleset": CloudNetworkRuleSetIdentity,

		"cloudnode":   CloudNodeIdentity,
		"cloudpolicy": CloudPolicyIdentity,

		"cloudroutetable":            CloudRouteTableIdentity,
		"cloudschedulednetworkquery": CloudScheduledNetworkQueryIdentity,
		"cloudsnapshotaccount":       CloudSnapshotAccountIdentity,
		"cloudsubnet":                CloudSubnetIdentity,

		"cloudvpc": CloudVPCIdentity,

		"cnsconfig":     CNSConfigIdentity,
		"cnsrqlquery":   CNSRQLQueryIdentity,
		"cnssearch":     CNSSearchIdentity,
		"cnssuggestion": CNSSuggestionIdentity,

		"connectionexceptionreport": ConnectionExceptionReportIdentity,
		"counterreport":             CounterReportIdentity,

		"datapathcertificate":    DataPathCertificateIdentity,
		"debugbundle":            DebugBundleIdentity,
		"defaultenforcerversion": DefaultEnforcerVersionIdentity,
		"dependencymap":          DependencyMapIdentity,
		"discoverymode":          DiscoveryModeIdentity,
		"dnslookupreport":        DNSLookupReportIdentity,
		"email":                  EmailIdentity,

		"enforcer":                     EnforcerIdentity,
		"enforcerlog":                  EnforcerLogIdentity,
		"enforcerprofile":              EnforcerProfileIdentity,
		"enforcerprofilemappingpolicy": EnforcerProfileMappingPolicyIdentity,
		"enforcerrefresh":              EnforcerRefreshIdentity,
		"enforcerreport":               EnforcerReportIdentity,
		"enforcertracereport":          EnforcerTraceReportIdentity,
		"eventlog":                     EventLogIdentity,
		"export":                       ExportIdentity,
		"externalnetwork":              ExternalNetworkIdentity,
		"fileaccesspolicy":             FileAccessPolicyIdentity,
		"fileaccessreport":             FileAccessReportIdentity,
		"filepath":                     FilePathIdentity,
		"flowreport":                   FlowReportIdentity,
		"graphedge":                    GraphEdgeIdentity,

		"graphnode":                GraphNodeIdentity,
		"healthcheck":              HealthCheckIdentity,
		"hit":                      HitIdentity,
		"hookpolicy":               HookPolicyIdentity,
		"hostservice":              HostServiceIdentity,
		"hostservicemappingpolicy": HostServiceMappingPolicyIdentity,
		"httpresourcespec":         HTTPResourceSpecIdentity,
		"import":                   ImportIdentity,
		"importreference":          ImportReferenceIdentity,
		"importrequest":            ImportRequestIdentity,
		"infrastructurepolicy":     InfrastructurePolicyIdentity,
		"installedapp":             InstalledAppIdentity,
		"ipinfo":                   IPInfoIdentity,
		"isolationprofile":         IsolationProfileIdentity,
		"issue":                    IssueIdentity,
		"issueservicetoken":        IssueServiceTokenIdentity,

		"ldapprovider":           LDAPProviderIdentity,
		"localca":                LocalCAIdentity,
		"log":                    LogIdentity,
		"logout":                 LogoutIdentity,
		"message":                MessageIdentity,
		"metricsquery":           MetricsQueryIdentity,
		"metricsqueryrange":      MetricsQueryRangeIdentity,
		"namespace":              NamespaceIdentity,
		"namespaceinfo":          NamespaceInfoIdentity,
		"namespacemappingpolicy": NamespaceMappingPolicyIdentity,
		"namespacepolicyinfo":    NamespacePolicyInfoIdentity,
		"namespacerenderer":      NamespaceRendererIdentity,
		"namespacetype":          NamespaceTypeIdentity,
		"networkaccesspolicy":    NetworkAccessPolicyIdentity,

		"networkrulesetpolicy":   NetworkRuleSetPolicyIdentity,
		"oauthinfo":              OAUTHInfoIdentity,
		"oauthkey":               OAUTHKeyIdentity,
		"oidcprovider":           OIDCProviderIdentity,
		"organizationalmetadata": OrganizationalMetadataIdentity,
		"packetreport":           PacketReportIdentity,
		"passwordreset":          PasswordResetIdentity,
		"pccprovider":            PCCProviderIdentity,
		"pcsearchresult":         PCSearchResultIdentity,
		"pctimerange":            PCTimeRangeIdentity,

		"pingprobe":   PingProbeIdentity,
		"pingrequest": PingRequestIdentity,
		"pingresult":  PingResultIdentity,

		"plan":                  PlanIdentity,
		"poke":                  PokeIdentity,
		"policy":                PolicyIdentity,
		"policygraph":           PolicyGraphIdentity,
		"policyrefresh":         PolicyRefreshIdentity,
		"policyrenderer":        PolicyRendererIdentity,
		"policyrule":            PolicyRuleIdentity,
		"policyttl":             PolicyTTLIdentity,
		"pollaccount":           PollAccountIdentity,
		"processingunit":        ProcessingUnitIdentity,
		"processingunitpolicy":  ProcessingUnitPolicyIdentity,
		"processingunitrefresh": ProcessingUnitRefreshIdentity,

		"putrafficaction": PUTrafficActionIdentity,
		"quotacheck":      QuotaCheckIdentity,
		"quotapolicy":     QuotaPolicyIdentity,
		"recipe":          RecipeIdentity,

		"remoteprocessor":         RemoteProcessorIdentity,
		"renderedpolicy":          RenderedPolicyIdentity,
		"rendertemplate":          RenderTemplateIdentity,
		"report":                  ReportIdentity,
		"reportsquery":            ReportsQueryIdentity,
		"revocation":              RevocationIdentity,
		"role":                    RoleIdentity,
		"root":                    RootIdentity,
		"samlprovider":            SAMLProviderIdentity,
		"sandbox":                 SandboxIdentity,
		"search":                  SearchIdentity,
		"service":                 ServiceIdentity,
		"servicecertificate":      ServiceCertificateIdentity,
		"servicedependencypolicy": ServiceDependencyPolicyIdentity,
		"servicepublication":      ServicePublicationIdentity,
		"servicetoken":            ServiceTokenIdentity,
		"squalltag":               SquallTagIdentity,
		"sshauthority":            SSHAuthorityIdentity,
		"sshauthorizationpolicy":  SSHAuthorizationPolicyIdentity,
		"sshcertificate":          SSHCertificateIdentity,
		"sshidentity":             SSHIdentityIdentity,
		"statsinfo":               StatsInfoIdentity,
		"statsquery":              StatsQueryIdentity,
		"suggestedpolicy":         SuggestedPolicyIdentity,
		"tag":                     TagIdentity,
		"taginject":               TagInjectIdentity,
		"tagprefix":               TagPrefixIdentity,
		"tagvalue":                TagValueIdentity,
		"tenant":                  TenantIdentity,
		"textindex":               TextIndexIdentity,

		"token":            TokenIdentity,
		"tokenscopepolicy": TokenScopePolicyIdentity,

		"trigger":          TriggerIdentity,
		"trustedca":        TrustedCAIdentity,
		"trustednamespace": TrustedNamespaceIdentity,

		"useraccesspolicy":     UserAccessPolicyIdentity,
		"validaterql":          ValidateRQLReqIdentity,
		"validateuiparameter":  ValidateUIParameterIdentity,
		"vulnerability":        VulnerabilityIdentity,
		"x509certificate":      X509CertificateIdentity,
		"x509certificatecheck": X509CertificateCheckIdentity,
	}

	identitycategoriesMap = map[string]elemental.Identity{
		"accessiblenamespaces":        AccessibleNamespaceIdentity,
		"accessreports":               AccessReportIdentity,
		"accounts":                    AccountIdentity,
		"accountchecks":               AccountCheckIdentity,
		"activate":                    ActivateIdentity,
		"activities":                  ActivityIdentity,
		"alarms":                      AlarmIdentity,
		"apiauthorizationpolicies":    APIAuthorizationPolicyIdentity,
		"apichecks":                   APICheckIdentity,
		"apps":                        AppIdentity,
		"appcredentials":              AppCredentialIdentity,
		"auditprofiles":               AuditProfileIdentity,
		"auditprofilemappingpolicies": AuditProfileMappingPolicyIdentity,
		"auditreports":                AuditReportIdentity,
		"authn":                       AuthnIdentity,
		"authorities":                 AuthorityIdentity,
		"authz":                       AuthzIdentity,
		"automations":                 AutomationIdentity,
		"automationtemplates":         AutomationTemplateIdentity,
		"cachedflowreports":           CachedFlowReportIdentity,
		"categories":                  CategoryIdentity,

		"claims":              ClaimsIdentity,
		"clausesmatches":      ClauseMatchIdentity,
		"cloudaccountcleaner": CloudAccountCleanerIdentity,

		"cloudalerts":     CloudAlertIdentity,
		"cloudalertsrule": CloudAlertRuleIdentity,
		"cloudendpoints":  CloudEndpointIdentity,

		"cloudgraphs": CloudGraphIdentity,

		"cloudmanagednetworks":   CloudManagedNetworkIdentity,
		"cloudnetworkinterfaces": CloudNetworkInterfaceIdentity,
		"cloudnetworkqueries":    CloudNetworkQueryIdentity,

		"cloudnetworkrulesets": CloudNetworkRuleSetIdentity,

		"cloudnodes":    CloudNodeIdentity,
		"cloudpolicies": CloudPolicyIdentity,

		"cloudroutetables":             CloudRouteTableIdentity,
		"cloudschedulednetworkqueries": CloudScheduledNetworkQueryIdentity,
		"cloudsnapshotaccounts":        CloudSnapshotAccountIdentity,
		"cloudsubnets":                 CloudSubnetIdentity,

		"cloudvpcs": CloudVPCIdentity,

		"cnsconfigs":     CNSConfigIdentity,
		"cnsrqlquery":    CNSRQLQueryIdentity,
		"cnssearches":    CNSSearchIdentity,
		"cnssuggestions": CNSSuggestionIdentity,

		"connectionexceptionreports": ConnectionExceptionReportIdentity,
		"counterreports":             CounterReportIdentity,

		"datapathcertificates":   DataPathCertificateIdentity,
		"debugbundles":           DebugBundleIdentity,
		"defaultenforcerversion": DefaultEnforcerVersionIdentity,
		"dependencymaps":         DependencyMapIdentity,
		"discoverymode":          DiscoveryModeIdentity,
		"dnslookupreports":       DNSLookupReportIdentity,
		"emails":                 EmailIdentity,

		"enforcers":                      EnforcerIdentity,
		"enforcerlog":                    EnforcerLogIdentity,
		"enforcerprofiles":               EnforcerProfileIdentity,
		"enforcerprofilemappingpolicies": EnforcerProfileMappingPolicyIdentity,
		"enforcerrefreshes":              EnforcerRefreshIdentity,
		"enforcerreports":                EnforcerReportIdentity,
		"enforcertracereports":           EnforcerTraceReportIdentity,
		"eventlogs":                      EventLogIdentity,
		"export":                         ExportIdentity,
		"externalnetworks":               ExternalNetworkIdentity,
		"fileaccesspolicies":             FileAccessPolicyIdentity,
		"fileaccessreports":              FileAccessReportIdentity,
		"filepaths":                      FilePathIdentity,
		"flowreports":                    FlowReportIdentity,
		"graphedges":                     GraphEdgeIdentity,

		"graphnodes":                 GraphNodeIdentity,
		"healthchecks":               HealthCheckIdentity,
		"hits":                       HitIdentity,
		"hookpolicies":               HookPolicyIdentity,
		"hostservices":               HostServiceIdentity,
		"hostservicemappingpolicies": HostServiceMappingPolicyIdentity,
		"httpresourcespecs":          HTTPResourceSpecIdentity,
		"import":                     ImportIdentity,
		"importreferences":           ImportReferenceIdentity,
		"importrequests":             ImportRequestIdentity,
		"infrastructurepolicies":     InfrastructurePolicyIdentity,
		"installedapps":              InstalledAppIdentity,
		"ipinfos":                    IPInfoIdentity,
		"isolationprofiles":          IsolationProfileIdentity,
		"issue":                      IssueIdentity,
		"issueservicetokens":         IssueServiceTokenIdentity,

		"ldapproviders":            LDAPProviderIdentity,
		"localcas":                 LocalCAIdentity,
		"logs":                     LogIdentity,
		"logout":                   LogoutIdentity,
		"messages":                 MessageIdentity,
		"metricsquery":             MetricsQueryIdentity,
		"metricsqueryrange":        MetricsQueryRangeIdentity,
		"namespaces":               NamespaceIdentity,
		"namespaceinfo":            NamespaceInfoIdentity,
		"namespacemappingpolicies": NamespaceMappingPolicyIdentity,
		"namespacepolicyinfo":      NamespacePolicyInfoIdentity,
		"namespacerenderers":       NamespaceRendererIdentity,
		"namespacetypes":           NamespaceTypeIdentity,
		"networkaccesspolicies":    NetworkAccessPolicyIdentity,

		"networkrulesetpolicies": NetworkRuleSetPolicyIdentity,
		"oauthinfo":              OAUTHInfoIdentity,
		"oauthkeys":              OAUTHKeyIdentity,
		"oidcproviders":          OIDCProviderIdentity,
		"organizationalmetadata": OrganizationalMetadataIdentity,
		"packetreports":          PacketReportIdentity,
		"passwordreset":          PasswordResetIdentity,
		"pccproviders":           PCCProviderIdentity,
		"pcsearchresults":        PCSearchResultIdentity,
		"pctimeranges":           PCTimeRangeIdentity,

		"pingprobes":   PingProbeIdentity,
		"pingrequests": PingRequestIdentity,
		"pingresults":  PingResultIdentity,

		"plans":                   PlanIdentity,
		"poke":                    PokeIdentity,
		"policies":                PolicyIdentity,
		"policygraphs":            PolicyGraphIdentity,
		"policyrefreshs":          PolicyRefreshIdentity,
		"policyrenderers":         PolicyRendererIdentity,
		"policyrules":             PolicyRuleIdentity,
		"policyttls":              PolicyTTLIdentity,
		"pollaccounts":            PollAccountIdentity,
		"processingunits":         ProcessingUnitIdentity,
		"processingunitpolicies":  ProcessingUnitPolicyIdentity,
		"processingunitrefreshes": ProcessingUnitRefreshIdentity,

		"putrafficactions": PUTrafficActionIdentity,
		"quotacheck":       QuotaCheckIdentity,
		"quotapolicies":    QuotaPolicyIdentity,
		"recipes":          RecipeIdentity,

		"remoteprocessors":          RemoteProcessorIdentity,
		"renderedpolicies":          RenderedPolicyIdentity,
		"rendertemplates":           RenderTemplateIdentity,
		"reports":                   ReportIdentity,
		"reportsqueries":            ReportsQueryIdentity,
		"revocations":               RevocationIdentity,
		"roles":                     RoleIdentity,
		"root":                      RootIdentity,
		"samlproviders":             SAMLProviderIdentity,
		"sandboxes":                 SandboxIdentity,
		"search":                    SearchIdentity,
		"services":                  ServiceIdentity,
		"servicecertificates":       ServiceCertificateIdentity,
		"servicedependencypolicies": ServiceDependencyPolicyIdentity,
		"servicepublications":       ServicePublicationIdentity,
		"servicetoken":              ServiceTokenIdentity,
		"squalltags":                SquallTagIdentity,
		"sshauthorities":            SSHAuthorityIdentity,
		"sshauthorizationpolicies":  SSHAuthorizationPolicyIdentity,
		"sshcertificates":           SSHCertificateIdentity,
		"sshidentities":             SSHIdentityIdentity,
		"statsinfo":                 StatsInfoIdentity,
		"statsqueries":              StatsQueryIdentity,
		"suggestedpolicies":         SuggestedPolicyIdentity,
		"tags":                      TagIdentity,
		"taginjects":                TagInjectIdentity,
		"tagprefixes":               TagPrefixIdentity,
		"tagvalues":                 TagValueIdentity,
		"tenants":                   TenantIdentity,
		"textindexes":               TextIndexIdentity,

		"tokens":             TokenIdentity,
		"tokenscopepolicies": TokenScopePolicyIdentity,

		"triggers":          TriggerIdentity,
		"trustedcas":        TrustedCAIdentity,
		"trustednamespaces": TrustedNamespaceIdentity,

		"useraccesspolicies":    UserAccessPolicyIdentity,
		"validaterql":           ValidateRQLReqIdentity,
		"validateuiparameters":  ValidateUIParameterIdentity,
		"vulnerabilities":       VulnerabilityIdentity,
		"x509certificates":      X509CertificateIdentity,
		"x509certificatechecks": X509CertificateCheckIdentity,
	}

	aliasesMap = map[string]elemental.Identity{
		"accns":           AccessibleNamespaceIdentity,
		"apiauth":         APIAuthorizationPolicyIdentity,
		"apiauths":        APIAuthorizationPolicyIdentity,
		"appcred":         AppCredentialIdentity,
		"appcreds":        AppCredentialIdentity,
		"ap":              AuditProfileIdentity,
		"audpol":          AuditProfileMappingPolicyIdentity,
		"audpols":         AuditProfileMappingPolicyIdentity,
		"ca":              AuthorityIdentity,
		"autos":           AutomationIdentity,
		"auto":            AutomationIdentity,
		"autotmpl":        AutomationTemplateIdentity,
		"crules":          CloudNetworkRuleSetIdentity,
		"vpc":             CloudVPCIdentity,
		"vpcs":            CloudVPCIdentity,
		"pcc":             CNSConfigIdentity,
		"depmaps":         DependencyMapIdentity,
		"depmap":          DependencyMapIdentity,
		"defender":        EnforcerIdentity,
		"profile":         EnforcerProfileIdentity,
		"profiles":        EnforcerProfileIdentity,
		"enfpols":         EnforcerProfileMappingPolicyIdentity,
		"enfpol":          EnforcerProfileMappingPolicyIdentity,
		"epm":             EnforcerProfileMappingPolicyIdentity,
		"extnet":          ExternalNetworkIdentity,
		"extnets":         ExternalNetworkIdentity,
		"fp":              FilePathIdentity,
		"fps":             FilePathIdentity,
		"hook":            HookPolicyIdentity,
		"hooks":           HookPolicyIdentity,
		"hookpol":         HookPolicyIdentity,
		"hookpols":        HookPolicyIdentity,
		"hostsrv":         HostServiceIdentity,
		"hostsrvs":        HostServiceIdentity,
		"hostsrvmappol":   HostServiceMappingPolicyIdentity,
		"hostsrvmappols":  HostServiceMappingPolicyIdentity,
		"httpresource":    HTTPResourceSpecIdentity,
		"resource":        HTTPResourceSpecIdentity,
		"httpspec":        HTTPResourceSpecIdentity,
		"importref":       ImportReferenceIdentity,
		"impref":          ImportReferenceIdentity,
		"req":             ImportRequestIdentity,
		"reqs":            ImportRequestIdentity,
		"ireq":            ImportRequestIdentity,
		"ireqs":           ImportRequestIdentity,
		"infrapol":        InfrastructurePolicyIdentity,
		"infrapols":       InfrastructurePolicyIdentity,
		"iapps":           InstalledAppIdentity,
		"iapp":            InstalledAppIdentity,
		"ip":              IsolationProfileIdentity,
		"mess":            MessageIdentity,
		"mq":              MetricsQueryIdentity,
		"mqr":             MetricsQueryRangeIdentity,
		"ns":              NamespaceIdentity,
		"nspolicy":        NamespaceMappingPolicyIdentity,
		"nspolicies":      NamespaceMappingPolicyIdentity,
		"nsmap":           NamespaceMappingPolicyIdentity,
		"nsmaps":          NamespaceMappingPolicyIdentity,
		"nsrenderer":      NamespaceRendererIdentity,
		"netpol":          NetworkAccessPolicyIdentity,
		"netpols":         NetworkAccessPolicyIdentity,
		"netruleset":      NetworkRuleSetPolicyIdentity,
		"netrulesets":     NetworkRuleSetPolicyIdentity,
		"netset":          NetworkRuleSetPolicyIdentity,
		"netsets":         NetworkRuleSetPolicyIdentity,
		"networkruleset":  NetworkRuleSetPolicyIdentity,
		"networkrulesets": NetworkRuleSetPolicyIdentity,
		"om":              OrganizationalMetadataIdentity,
		"polgraph":        PolicyGraphIdentity,
		"pu":              ProcessingUnitIdentity,
		"pus":             ProcessingUnitIdentity,
		"pup":             ProcessingUnitPolicyIdentity,
		"pups":            ProcessingUnitPolicyIdentity,
		"quota":           QuotaPolicyIdentity,
		"quotas":          QuotaPolicyIdentity,
		"quotapol":        QuotaPolicyIdentity,
		"quotapols":       QuotaPolicyIdentity,
		"rcp":             RecipeIdentity,
		"hks":             RemoteProcessorIdentity,
		"hk":              RemoteProcessorIdentity,
		"rpol":            RenderedPolicyIdentity,
		"rpols":           RenderedPolicyIdentity,
		"cook":            RenderTemplateIdentity,
		"rtpl":            RenderTemplateIdentity,
		"rq":              ReportsQueryIdentity,
		"srv":             ServiceIdentity,
		"srvdep":          ServiceDependencyPolicyIdentity,
		"srvdeps":         ServiceDependencyPolicyIdentity,
		"sshpol":          SSHAuthorizationPolicyIdentity,
		"sshpols":         SSHAuthorizationPolicyIdentity,
		"si":              StatsInfoIdentity,
		"sq":              StatsQueryIdentity,
		"sugpol":          SuggestedPolicyIdentity,
		"sugpols":         SuggestedPolicyIdentity,
		"sugg":            SuggestedPolicyIdentity,
		"suggs":           SuggestedPolicyIdentity,
		"tsp":             TokenScopePolicyIdentity,
		"trustedns":       TrustedNamespaceIdentity,
		"usrpol":          UserAccessPolicyIdentity,
		"usrpols":         UserAccessPolicyIdentity,
		"validparam":      ValidateUIParameterIdentity,
		"vulns":           VulnerabilityIdentity,
		"vul":             VulnerabilityIdentity,
		"vuln":            VulnerabilityIdentity,
		"vuls":            VulnerabilityIdentity,
	}

	indexesMap = map[string][][]string{
		"accessiblenamespace": nil,
		"accessreport": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"account": {
			{"resetPasswordToken"},
			{"name"},
			{"email"},
			{"activationToken"},
			{":shard", ":unique", "zone", "zHash"},
		},
		"accountcheck": nil,
		"activate":     nil,
		"activity": {
			{"namespace", "date"},
			{"namespace", "operation"},
			{"namespace", "error.code"},
			{"namespace", "targetIdentity"},
			{"namespace", "data.ID"},
			{"namespace", "originalData.ID"},
			{":shard", ":unique", "zone", "zHash"},
			{"namespace"},
			{"namespace", "normalizedTags"},
		},
		"alarm": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"status"},
			{"namespace", "normalizedTags"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "kind"},
			{"namespace", "kind", "status"},
			{"name"},
			{"kind"},
			{"createIdempotencyKey"},
		},
		"apiauthorizationpolicy": {
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "allSubjectTags", "disabled"},
			{"namespace", "allObjectTags", "propagate"},
			{"namespace", "allSubjectTags", "propagate"},
			{"namespace", "normalizedTags"},
			{"namespace"},
			{"namespace", "allObjectTags", "disabled"},
			{"namespace", "fallback"},
			{"namespace", "disabled"},
			{":shard", ":unique", "zone", "zHash"},
			{"namespace", "name"},
			{"name"},
			{"fallback"},
			{"disabled"},
			{"createIdempotencyKey"},
		},
		"apicheck": nil,
		"app":      nil,
		"appcredential": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "disabled"},
			{"name"},
			{"disabled"},
			{"createIdempotencyKey"},
		},
		"auditprofile": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"auditprofilemappingpolicy": nil,
		"auditreport": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"authn": nil,
		"authority": {
			{"commonName"},
			{":shard", ":unique", "zone", "zHash"},
		},
		"authz": nil,
		"automation": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "disabled"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"disabled"},
			{"createIdempotencyKey"},
		},
		"automationtemplate": nil,
		"cachedflowreport": {
			{"sourceID"},
			{"namespace", "timestamp"},
			{"destinationID"},
			{":shard", "zone", "zHash", "_id"},
		},
		"category": nil,
		"claims": {
			{"namespace", "hash"},
			{":shard", ":unique", "zone", "zHash"},
			{"namespace"},
			{"namespace", "normalizedTags"},
		},
		"clausesmatch":        nil,
		"cloudaccountcleaner": nil,
		"cloudalert": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"cloudalertrule": {
			{":unique", "alertRuleID"},
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "alertRuleID"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"cloudendpoint": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "nativeID"},
			{"namespace", "accountid"},
			{"namespace", "vpcid"},
			{"createIdempotencyKey"},
		},
		"cloudgraph": nil,
		"cloudmanagednetwork": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"type"},
			{"namespace"},
			{"namespace", "nativeID"},
			{"namespace", "type"},
			{"namespace", "normalizedTags"},
			{"namespace", "accountid"},
			{"namespace", "vpcid"},
			{"key"},
			{"createIdempotencyKey"},
		},
		"cloudnetworkinterface": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "nativeID"},
			{"namespace", "accountid"},
			{"namespace", "vpcid"},
			{"createIdempotencyKey"},
		},
		"cloudnetworkquery": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"cloudnetworkruleset": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "normalizedTags"},
			{"namespace", "vpcid"},
			{"namespace"},
			{"namespace", "vpcid", "parameters"},
			{"namespace", "accountid"},
			{"namespace", "nativeID"},
			{"key"},
			{"createIdempotencyKey"},
		},
		"cloudnode": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "relatedObjectID"},
			{"namespace", "securitytags", "type", "vpcid"},
			{"namespace", "vpcid", "type"},
			{"namespace", "vpcid", "parameters"},
			{"namespace", "type", "parameters"},
			{"namespace", "nativeID"},
			{"namespace", "vpcid"},
			{"namespace"},
			{"namespace", "accountid"},
			{"namespace", "type", "subtype"},
			{"namespace", "normalizedTags"},
			{"key"},
			{"createTime"},
			{"createIdempotencyKey"},
		},
		"cloudpolicy": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "severity"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "prismaCloudPolicyID"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"cloudroutetable": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "nativeID"},
			{"namespace", "accountid"},
			{"namespace", "vpcid"},
			{"createIdempotencyKey"},
		},
		"cloudschedulednetworkquery": {
			{":shard", ":unique", "zone", "zHash"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"lastexecutiontimestamp"},
			{"alertruleid"},
			{"alertruleid", "policyid"},
		},
		"cloudsnapshotaccount": {
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"createIdempotencyKey"},
		},
		"cloudsubnet": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "nativeID"},
			{"namespace", "accountid"},
			{"namespace", "vpcid"},
			{"createIdempotencyKey"},
		},
		"cloudvpc": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "nativeID"},
			{"namespace", "accountid"},
			{"namespace", "vpcid"},
			{"createIdempotencyKey"},
		},
		"cnsconfig": {
			{"updateIdempotencyKey"},
			{"prismaID"},
			{"namespace", "prismaID"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"createIdempotencyKey"},
			{":shard", ":unique", "zone", "zHash"},
		},
		"cnsrqlquery":   nil,
		"cnssearch":     nil,
		"cnssuggestion": nil,
		"connectionexceptionreport": {
			{"processingunitnamespace", "timestamp"},
			{"namespace", "timestamp"},
			{"enforcernamespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"counterreport": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"datapathcertificate":    nil,
		"debugbundle":            nil,
		"defaultenforcerversion": nil,
		"dependencymap":          nil,
		"discoverymode": {
			{"propagate"},
			{"namespace"},
			{"namespace", "normalizedTags"},
		},
		"dnslookupreport": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
			{"namespace"},
			{"namespace", "normalizedTags"},
		},
		"email": nil,
		"enforcer": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "machineID"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "operationalStatus"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"enforcerlog": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "collectionID"},
			{"namespace", "enforcerID"},
			{"enforcerID"},
			{"createIdempotencyKey"},
			{"collectionID"},
		},
		"enforcerprofile": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"enforcerprofilemappingpolicy": nil,
		"enforcerrefresh": {
			{"propagate"},
		},
		"enforcerreport": {
			{"namespace", "timestamp"},
			{"namespace", "enforcerID"},
			{"enforcerID"},
			{":shard", "zone", "zHash", "_id"},
		},
		"enforcertracereport": {
			{"namespace", "timestamp"},
			{"namespace", "enforcerID"},
			{"enforcerID"},
		},
		"eventlog": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"export": nil,
		"externalnetwork": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"type"},
			{"propagate"},
			{"namespace", "normalizedTags"},
			{"namespace", "archived"},
			{"namespace", "type"},
			{"namespace"},
			{"namespace", "name"},
			{"name"},
			{"createIdempotencyKey"},
			{"archived"},
		},
		"fileaccesspolicy": nil,
		"fileaccessreport": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"filepath": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "normalizedTags"},
			{"namespace", "archived"},
			{"namespace"},
			{"namespace", "name"},
			{"name"},
			{"createIdempotencyKey"},
			{"archived"},
		},
		"flowreport": {
			{"sourceID"},
			{"remotenamespace", "timestamp"},
			{"namespace", "timestamp"},
			{"destinationID"},
			{":shard", "zone", "zHash", "_id"},
		},
		"graphedge": {
			{":shard", ":unique", "zone", "zHash"},
			{"namespace"},
			{"namespace", "lastSeen", "firstSeen"},
			{"lastSeen", "firstSeen"},
			{"lastSeen"},
			{"flowID", "bucketMonth"},
			{"flowID", "lastSeen"},
			{"flowID", "bucketMinute"},
			{"flowID", "bucketHour"},
			{"flowID"},
			{"flowID", "bucketDay"},
			{"firstSeen"},
		},
		"graphnode":   nil,
		"healthcheck": nil,
		"hit":         nil,
		"hookpolicy":  nil,
		"hostservice": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "normalizedTags"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "archived"},
			{"name"},
			{"createIdempotencyKey"},
			{"archived"},
		},
		"hostservicemappingpolicy": nil,
		"httpresourcespec": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "normalizedTags"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "archived"},
			{"name"},
			{"createIdempotencyKey"},
			{"archived"},
		},
		"import": nil,
		"importreference": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "label"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"importrequest": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"createIdempotencyKey"},
		},
		"infrastructurepolicy": nil,
		"installedapp": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"ipinfo": nil,
		"isolationprofile": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"issue":             nil,
		"issueservicetoken": nil,
		"ldapprovider": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"localca": nil,
		"log":     nil,
		"logout":  nil,
		"message": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"metricsquery":      nil,
		"metricsqueryrange": nil,
		"namespace": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"type"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"namespace", "name"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"namespaceinfo":          nil,
		"namespacemappingpolicy": nil,
		"namespacepolicyinfo":    nil,
		"namespacerenderer":      nil,
		"namespacetype":          nil,
		"networkaccesspolicy":    nil,
		"networkrulesetpolicy":   nil,
		"oauthinfo":              nil,
		"oauthkey":               nil,
		"oidcprovider": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"organizationalmetadata": {
			{"namespace"},
			{"namespace", "normalizedTags"},
		},
		"packetreport": {
			{"namespace", "timestamp"},
			{":shard", "zone", "zHash", "_id"},
		},
		"passwordreset": nil,
		"pccprovider": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"pcsearchresult": nil,
		"pctimerange":    nil,
		"pingprobe": {
			{"pingID"},
			{"namespace", "pingID"},
			{"namespace", "pingID", "iterationIndex"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{":shard", ":unique", "zone", "zHash"},
		},
		"pingrequest": nil,
		"pingresult": {
			{"pingID"},
			{"namespace", "pingID"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{":shard", ":unique", "zone", "zHash"},
		},
		"plan": nil,
		"poke": nil,
		"policy": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "type", "allSubjectTags", "propagate"},
			{"namespace", "type", "allObjectTags"},
			{"namespace", "type", "allSubjectTags"},
			{"namespace", "type", "allObjectTags", "disabled"},
			{"namespace", "normalizedTags"},
			{"namespace", "type"},
			{"namespace", "disabled"},
			{"namespace", "name"},
			{"namespace", "type", "allObjectTags", "propagate"},
			{"namespace"},
			{"namespace", "type", "allSubjectTags", "disabled"},
			{"namespace", "fallback"},
			{"name"},
			{"fallback"},
			{"disabled"},
			{"createIdempotencyKey"},
		},
		"policygraph":    nil,
		"policyrefresh":  nil,
		"policyrenderer": nil,
		"policyrule":     nil,
		"policyttl":      nil,
		"pollaccount":    nil,
		"processingunit": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "nativeContextID"},
			{"namespace", "normalizedTags"},
			{"namespace"},
			{"namespace", "operationalStatus", "archived"},
			{"namespace", "normalizedTags", "archived"},
			{"namespace", "name"},
			{"namespace", "archived"},
			{"name"},
			{"enforcerID"},
			{"createtime"},
			{"createIdempotencyKey"},
			{"archived", "updatetime"},
			{"archived"},
		},
		"processingunitpolicy":  nil,
		"processingunitrefresh": nil,
		"putrafficaction":       nil,
		"quotacheck":            nil,
		"quotapolicy":           nil,
		"recipe": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"remoteprocessor": nil,
		"renderedpolicy":  nil,
		"rendertemplate":  nil,
		"report":          nil,
		"reportsquery":    nil,
		"revocation": {
			{":shard", ":unique", "zone", "zHash"},
		},
		"role": nil,
		"root": nil,
		"samlprovider": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"sandbox": nil,
		"search":  nil,
		"service": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "allAPITags"},
			{"namespace"},
			{"namespace", "archived"},
			{"namespace", "allProcessingUnitsTags"},
			{"namespace", "normalizedTags"},
			{"namespace", "disabled"},
			{"namespace", "name"},
			{"name"},
			{"disabled"},
			{"createIdempotencyKey"},
			{"archived"},
			{"allProcessingUnitsTags"},
			{"allAPITags"},
		},
		"servicecertificate": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"propagate"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "fingerprint"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"fingerprint"},
			{"createIdempotencyKey"},
		},
		"servicedependencypolicy": nil,
		"servicepublication":      nil,
		"servicetoken":            nil,
		"squalltag":               nil,
		"sshauthority": {
			{":shard", ":unique", "zone", "zHash"},
			{"namespace", "name"},
			{"name"},
		},
		"sshauthorizationpolicy": nil,
		"sshcertificate":         nil,
		"sshidentity":            nil,
		"statsinfo":              nil,
		"statsquery":             nil,
		"suggestedpolicy":        nil,
		"tag": {
			{"namespace"},
			{"namespace", "normalizedTags"},
		},
		"taginject": nil,
		"tagprefix": nil,
		"tagvalue":  nil,
		"tenant":    nil,
		"textindex": {
			{"objectNamespace"},
			{"objectNamespace", "objectIdentity", "objectID"},
			{"objectIdentity"},
			{"objectID"},
			{"date"},
			{":shard", ":unique", "zone", "zHash"},
		},
		"token":            nil,
		"tokenscopepolicy": nil,
		"trigger":          nil,
		"trustedca":        nil,
		"trustednamespace": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"namespace", "name"},
			{"namespace"},
			{"namespace", "normalizedTags"},
			{"name"},
			{"createIdempotencyKey"},
		},
		"useraccesspolicy":    nil,
		"validaterql":         nil,
		"validateuiparameter": nil,
		"vulnerability": {
			{":shard", ":unique", "zone", "zHash"},
			{"updateIdempotencyKey"},
			{"severity"},
			{"propagate"},
			{"namespace"},
			{"namespace", "severity"},
			{"namespace", "CVSS2Score"},
			{"namespace", "normalizedTags"},
			{"namespace", "name"},
			{"name"},
			{"createIdempotencyKey"},
			{"CVSS2Score"},
		},
		"x509certificate":      nil,
		"x509certificatecheck": nil,
	}
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type modelManager struct{}

func (f modelManager) IdentityFromName(name string) elemental.Identity {

	return identityNamesMap[name]
}

func (f modelManager) IdentityFromCategory(category string) elemental.Identity {

	return identitycategoriesMap[category]
}

func (f modelManager) IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

func (f modelManager) IdentityFromAny(any string) (i elemental.Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f modelManager) Identifiable(identity elemental.Identity) elemental.Identifiable {

	switch identity {

	case AccessibleNamespaceIdentity:
		return NewAccessibleNamespace()
	case AccessReportIdentity:
		return NewAccessReport()
	case AccountIdentity:
		return NewAccount()
	case AccountCheckIdentity:
		return NewAccountCheck()
	case ActivateIdentity:
		return NewActivate()
	case ActivityIdentity:
		return NewActivity()
	case AlarmIdentity:
		return NewAlarm()
	case APIAuthorizationPolicyIdentity:
		return NewAPIAuthorizationPolicy()
	case APICheckIdentity:
		return NewAPICheck()
	case AppIdentity:
		return NewApp()
	case AppCredentialIdentity:
		return NewAppCredential()
	case AuditProfileIdentity:
		return NewAuditProfile()
	case AuditProfileMappingPolicyIdentity:
		return NewAuditProfileMappingPolicy()
	case AuditReportIdentity:
		return NewAuditReport()
	case AuthnIdentity:
		return NewAuthn()
	case AuthorityIdentity:
		return NewAuthority()
	case AuthzIdentity:
		return NewAuthz()
	case AutomationIdentity:
		return NewAutomation()
	case AutomationTemplateIdentity:
		return NewAutomationTemplate()
	case CachedFlowReportIdentity:
		return NewCachedFlowReport()
	case CategoryIdentity:
		return NewCategory()
	case ClaimsIdentity:
		return NewClaims()
	case ClauseMatchIdentity:
		return NewClauseMatch()
	case CloudAccountCleanerIdentity:
		return NewCloudAccountCleaner()
	case CloudAlertIdentity:
		return NewCloudAlert()
	case CloudAlertRuleIdentity:
		return NewCloudAlertRule()
	case CloudEndpointIdentity:
		return NewCloudEndpoint()
	case CloudGraphIdentity:
		return NewCloudGraph()
	case CloudManagedNetworkIdentity:
		return NewCloudManagedNetwork()
	case CloudNetworkInterfaceIdentity:
		return NewCloudNetworkInterface()
	case CloudNetworkQueryIdentity:
		return NewCloudNetworkQuery()
	case CloudNetworkRuleSetIdentity:
		return NewCloudNetworkRuleSet()
	case CloudNodeIdentity:
		return NewCloudNode()
	case CloudPolicyIdentity:
		return NewCloudPolicy()
	case CloudRouteTableIdentity:
		return NewCloudRouteTable()
	case CloudScheduledNetworkQueryIdentity:
		return NewCloudScheduledNetworkQuery()
	case CloudSnapshotAccountIdentity:
		return NewCloudSnapshotAccount()
	case CloudSubnetIdentity:
		return NewCloudSubnet()
	case CloudVPCIdentity:
		return NewCloudVPC()
	case CNSConfigIdentity:
		return NewCNSConfig()
	case CNSRQLQueryIdentity:
		return NewCNSRQLQuery()
	case CNSSearchIdentity:
		return NewCNSSearch()
	case CNSSuggestionIdentity:
		return NewCNSSuggestion()
	case ConnectionExceptionReportIdentity:
		return NewConnectionExceptionReport()
	case CounterReportIdentity:
		return NewCounterReport()
	case DataPathCertificateIdentity:
		return NewDataPathCertificate()
	case DebugBundleIdentity:
		return NewDebugBundle()
	case DefaultEnforcerVersionIdentity:
		return NewDefaultEnforcerVersion()
	case DependencyMapIdentity:
		return NewDependencyMap()
	case DiscoveryModeIdentity:
		return NewDiscoveryMode()
	case DNSLookupReportIdentity:
		return NewDNSLookupReport()
	case EmailIdentity:
		return NewEmail()
	case EnforcerIdentity:
		return NewEnforcer()
	case EnforcerLogIdentity:
		return NewEnforcerLog()
	case EnforcerProfileIdentity:
		return NewEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity:
		return NewEnforcerProfileMappingPolicy()
	case EnforcerRefreshIdentity:
		return NewEnforcerRefresh()
	case EnforcerReportIdentity:
		return NewEnforcerReport()
	case EnforcerTraceReportIdentity:
		return NewEnforcerTraceReport()
	case EventLogIdentity:
		return NewEventLog()
	case ExportIdentity:
		return NewExport()
	case ExternalNetworkIdentity:
		return NewExternalNetwork()
	case FileAccessPolicyIdentity:
		return NewFileAccessPolicy()
	case FileAccessReportIdentity:
		return NewFileAccessReport()
	case FilePathIdentity:
		return NewFilePath()
	case FlowReportIdentity:
		return NewFlowReport()
	case GraphEdgeIdentity:
		return NewGraphEdge()
	case GraphNodeIdentity:
		return NewGraphNode()
	case HealthCheckIdentity:
		return NewHealthCheck()
	case HitIdentity:
		return NewHit()
	case HookPolicyIdentity:
		return NewHookPolicy()
	case HostServiceIdentity:
		return NewHostService()
	case HostServiceMappingPolicyIdentity:
		return NewHostServiceMappingPolicy()
	case HTTPResourceSpecIdentity:
		return NewHTTPResourceSpec()
	case ImportIdentity:
		return NewImport()
	case ImportReferenceIdentity:
		return NewImportReference()
	case ImportRequestIdentity:
		return NewImportRequest()
	case InfrastructurePolicyIdentity:
		return NewInfrastructurePolicy()
	case InstalledAppIdentity:
		return NewInstalledApp()
	case IPInfoIdentity:
		return NewIPInfo()
	case IsolationProfileIdentity:
		return NewIsolationProfile()
	case IssueIdentity:
		return NewIssue()
	case IssueServiceTokenIdentity:
		return NewIssueServiceToken()
	case LDAPProviderIdentity:
		return NewLDAPProvider()
	case LocalCAIdentity:
		return NewLocalCA()
	case LogIdentity:
		return NewLog()
	case LogoutIdentity:
		return NewLogout()
	case MessageIdentity:
		return NewMessage()
	case MetricsQueryIdentity:
		return NewMetricsQuery()
	case MetricsQueryRangeIdentity:
		return NewMetricsQueryRange()
	case NamespaceIdentity:
		return NewNamespace()
	case NamespaceInfoIdentity:
		return NewNamespaceInfo()
	case NamespaceMappingPolicyIdentity:
		return NewNamespaceMappingPolicy()
	case NamespacePolicyInfoIdentity:
		return NewNamespacePolicyInfo()
	case NamespaceRendererIdentity:
		return NewNamespaceRenderer()
	case NamespaceTypeIdentity:
		return NewNamespaceType()
	case NetworkAccessPolicyIdentity:
		return NewNetworkAccessPolicy()
	case NetworkRuleSetPolicyIdentity:
		return NewNetworkRuleSetPolicy()
	case OAUTHInfoIdentity:
		return NewOAUTHInfo()
	case OAUTHKeyIdentity:
		return NewOAUTHKey()
	case OIDCProviderIdentity:
		return NewOIDCProvider()
	case OrganizationalMetadataIdentity:
		return NewOrganizationalMetadata()
	case PacketReportIdentity:
		return NewPacketReport()
	case PasswordResetIdentity:
		return NewPasswordReset()
	case PCCProviderIdentity:
		return NewPCCProvider()
	case PCSearchResultIdentity:
		return NewPCSearchResult()
	case PCTimeRangeIdentity:
		return NewPCTimeRange()
	case PingProbeIdentity:
		return NewPingProbe()
	case PingRequestIdentity:
		return NewPingRequest()
	case PingResultIdentity:
		return NewPingResult()
	case PlanIdentity:
		return NewPlan()
	case PokeIdentity:
		return NewPoke()
	case PolicyIdentity:
		return NewPolicy()
	case PolicyGraphIdentity:
		return NewPolicyGraph()
	case PolicyRefreshIdentity:
		return NewPolicyRefresh()
	case PolicyRendererIdentity:
		return NewPolicyRenderer()
	case PolicyRuleIdentity:
		return NewPolicyRule()
	case PolicyTTLIdentity:
		return NewPolicyTTL()
	case PollAccountIdentity:
		return NewPollAccount()
	case ProcessingUnitIdentity:
		return NewProcessingUnit()
	case ProcessingUnitPolicyIdentity:
		return NewProcessingUnitPolicy()
	case ProcessingUnitRefreshIdentity:
		return NewProcessingUnitRefresh()
	case PUTrafficActionIdentity:
		return NewPUTrafficAction()
	case QuotaCheckIdentity:
		return NewQuotaCheck()
	case QuotaPolicyIdentity:
		return NewQuotaPolicy()
	case RecipeIdentity:
		return NewRecipe()
	case RemoteProcessorIdentity:
		return NewRemoteProcessor()
	case RenderedPolicyIdentity:
		return NewRenderedPolicy()
	case RenderTemplateIdentity:
		return NewRenderTemplate()
	case ReportIdentity:
		return NewReport()
	case ReportsQueryIdentity:
		return NewReportsQuery()
	case RevocationIdentity:
		return NewRevocation()
	case RoleIdentity:
		return NewRole()
	case RootIdentity:
		return NewRoot()
	case SAMLProviderIdentity:
		return NewSAMLProvider()
	case SandboxIdentity:
		return NewSandbox()
	case SearchIdentity:
		return NewSearch()
	case ServiceIdentity:
		return NewService()
	case ServiceCertificateIdentity:
		return NewServiceCertificate()
	case ServiceDependencyPolicyIdentity:
		return NewServiceDependencyPolicy()
	case ServicePublicationIdentity:
		return NewServicePublication()
	case ServiceTokenIdentity:
		return NewServiceToken()
	case SquallTagIdentity:
		return NewSquallTag()
	case SSHAuthorityIdentity:
		return NewSSHAuthority()
	case SSHAuthorizationPolicyIdentity:
		return NewSSHAuthorizationPolicy()
	case SSHCertificateIdentity:
		return NewSSHCertificate()
	case SSHIdentityIdentity:
		return NewSSHIdentity()
	case StatsInfoIdentity:
		return NewStatsInfo()
	case StatsQueryIdentity:
		return NewStatsQuery()
	case SuggestedPolicyIdentity:
		return NewSuggestedPolicy()
	case TagIdentity:
		return NewTag()
	case TagInjectIdentity:
		return NewTagInject()
	case TagPrefixIdentity:
		return NewTagPrefix()
	case TagValueIdentity:
		return NewTagValue()
	case TenantIdentity:
		return NewTenant()
	case TextIndexIdentity:
		return NewTextIndex()
	case TokenIdentity:
		return NewToken()
	case TokenScopePolicyIdentity:
		return NewTokenScopePolicy()
	case TriggerIdentity:
		return NewTrigger()
	case TrustedCAIdentity:
		return NewTrustedCA()
	case TrustedNamespaceIdentity:
		return NewTrustedNamespace()
	case UserAccessPolicyIdentity:
		return NewUserAccessPolicy()
	case ValidateRQLReqIdentity:
		return NewValidateRQLReq()
	case ValidateUIParameterIdentity:
		return NewValidateUIParameter()
	case VulnerabilityIdentity:
		return NewVulnerability()
	case X509CertificateIdentity:
		return NewX509Certificate()
	case X509CertificateCheckIdentity:
		return NewX509CertificateCheck()
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiable(identity elemental.Identity) elemental.SparseIdentifiable {

	switch identity {

	case AccessibleNamespaceIdentity:
		return NewSparseAccessibleNamespace()
	case AccessReportIdentity:
		return NewSparseAccessReport()
	case AccountIdentity:
		return NewSparseAccount()
	case AccountCheckIdentity:
		return NewSparseAccountCheck()
	case ActivateIdentity:
		return NewSparseActivate()
	case ActivityIdentity:
		return NewSparseActivity()
	case AlarmIdentity:
		return NewSparseAlarm()
	case APIAuthorizationPolicyIdentity:
		return NewSparseAPIAuthorizationPolicy()
	case APICheckIdentity:
		return NewSparseAPICheck()
	case AppIdentity:
		return NewSparseApp()
	case AppCredentialIdentity:
		return NewSparseAppCredential()
	case AuditProfileIdentity:
		return NewSparseAuditProfile()
	case AuditProfileMappingPolicyIdentity:
		return NewSparseAuditProfileMappingPolicy()
	case AuditReportIdentity:
		return NewSparseAuditReport()
	case AuthnIdentity:
		return NewSparseAuthn()
	case AuthorityIdentity:
		return NewSparseAuthority()
	case AuthzIdentity:
		return NewSparseAuthz()
	case AutomationIdentity:
		return NewSparseAutomation()
	case AutomationTemplateIdentity:
		return NewSparseAutomationTemplate()
	case CachedFlowReportIdentity:
		return NewSparseCachedFlowReport()
	case CategoryIdentity:
		return NewSparseCategory()
	case ClaimsIdentity:
		return NewSparseClaims()
	case ClauseMatchIdentity:
		return NewSparseClauseMatch()
	case CloudAccountCleanerIdentity:
		return NewSparseCloudAccountCleaner()
	case CloudAlertIdentity:
		return NewSparseCloudAlert()
	case CloudAlertRuleIdentity:
		return NewSparseCloudAlertRule()
	case CloudEndpointIdentity:
		return NewSparseCloudEndpoint()
	case CloudGraphIdentity:
		return NewSparseCloudGraph()
	case CloudManagedNetworkIdentity:
		return NewSparseCloudManagedNetwork()
	case CloudNetworkInterfaceIdentity:
		return NewSparseCloudNetworkInterface()
	case CloudNetworkQueryIdentity:
		return NewSparseCloudNetworkQuery()
	case CloudNetworkRuleSetIdentity:
		return NewSparseCloudNetworkRuleSet()
	case CloudNodeIdentity:
		return NewSparseCloudNode()
	case CloudPolicyIdentity:
		return NewSparseCloudPolicy()
	case CloudRouteTableIdentity:
		return NewSparseCloudRouteTable()
	case CloudScheduledNetworkQueryIdentity:
		return NewSparseCloudScheduledNetworkQuery()
	case CloudSnapshotAccountIdentity:
		return NewSparseCloudSnapshotAccount()
	case CloudSubnetIdentity:
		return NewSparseCloudSubnet()
	case CloudVPCIdentity:
		return NewSparseCloudVPC()
	case CNSConfigIdentity:
		return NewSparseCNSConfig()
	case CNSRQLQueryIdentity:
		return NewSparseCNSRQLQuery()
	case CNSSearchIdentity:
		return NewSparseCNSSearch()
	case CNSSuggestionIdentity:
		return NewSparseCNSSuggestion()
	case ConnectionExceptionReportIdentity:
		return NewSparseConnectionExceptionReport()
	case CounterReportIdentity:
		return NewSparseCounterReport()
	case DataPathCertificateIdentity:
		return NewSparseDataPathCertificate()
	case DebugBundleIdentity:
		return NewSparseDebugBundle()
	case DefaultEnforcerVersionIdentity:
		return NewSparseDefaultEnforcerVersion()
	case DependencyMapIdentity:
		return NewSparseDependencyMap()
	case DiscoveryModeIdentity:
		return NewSparseDiscoveryMode()
	case DNSLookupReportIdentity:
		return NewSparseDNSLookupReport()
	case EmailIdentity:
		return NewSparseEmail()
	case EnforcerIdentity:
		return NewSparseEnforcer()
	case EnforcerLogIdentity:
		return NewSparseEnforcerLog()
	case EnforcerProfileIdentity:
		return NewSparseEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity:
		return NewSparseEnforcerProfileMappingPolicy()
	case EnforcerRefreshIdentity:
		return NewSparseEnforcerRefresh()
	case EnforcerReportIdentity:
		return NewSparseEnforcerReport()
	case EnforcerTraceReportIdentity:
		return NewSparseEnforcerTraceReport()
	case EventLogIdentity:
		return NewSparseEventLog()
	case ExportIdentity:
		return NewSparseExport()
	case ExternalNetworkIdentity:
		return NewSparseExternalNetwork()
	case FileAccessPolicyIdentity:
		return NewSparseFileAccessPolicy()
	case FileAccessReportIdentity:
		return NewSparseFileAccessReport()
	case FilePathIdentity:
		return NewSparseFilePath()
	case FlowReportIdentity:
		return NewSparseFlowReport()
	case GraphEdgeIdentity:
		return NewSparseGraphEdge()
	case GraphNodeIdentity:
		return NewSparseGraphNode()
	case HealthCheckIdentity:
		return NewSparseHealthCheck()
	case HitIdentity:
		return NewSparseHit()
	case HookPolicyIdentity:
		return NewSparseHookPolicy()
	case HostServiceIdentity:
		return NewSparseHostService()
	case HostServiceMappingPolicyIdentity:
		return NewSparseHostServiceMappingPolicy()
	case HTTPResourceSpecIdentity:
		return NewSparseHTTPResourceSpec()
	case ImportIdentity:
		return NewSparseImport()
	case ImportReferenceIdentity:
		return NewSparseImportReference()
	case ImportRequestIdentity:
		return NewSparseImportRequest()
	case InfrastructurePolicyIdentity:
		return NewSparseInfrastructurePolicy()
	case InstalledAppIdentity:
		return NewSparseInstalledApp()
	case IPInfoIdentity:
		return NewSparseIPInfo()
	case IsolationProfileIdentity:
		return NewSparseIsolationProfile()
	case IssueIdentity:
		return NewSparseIssue()
	case IssueServiceTokenIdentity:
		return NewSparseIssueServiceToken()
	case LDAPProviderIdentity:
		return NewSparseLDAPProvider()
	case LocalCAIdentity:
		return NewSparseLocalCA()
	case LogIdentity:
		return NewSparseLog()
	case LogoutIdentity:
		return NewSparseLogout()
	case MessageIdentity:
		return NewSparseMessage()
	case MetricsQueryIdentity:
		return NewSparseMetricsQuery()
	case MetricsQueryRangeIdentity:
		return NewSparseMetricsQueryRange()
	case NamespaceIdentity:
		return NewSparseNamespace()
	case NamespaceInfoIdentity:
		return NewSparseNamespaceInfo()
	case NamespaceMappingPolicyIdentity:
		return NewSparseNamespaceMappingPolicy()
	case NamespacePolicyInfoIdentity:
		return NewSparseNamespacePolicyInfo()
	case NamespaceRendererIdentity:
		return NewSparseNamespaceRenderer()
	case NamespaceTypeIdentity:
		return NewSparseNamespaceType()
	case NetworkAccessPolicyIdentity:
		return NewSparseNetworkAccessPolicy()
	case NetworkRuleSetPolicyIdentity:
		return NewSparseNetworkRuleSetPolicy()
	case OAUTHInfoIdentity:
		return NewSparseOAUTHInfo()
	case OAUTHKeyIdentity:
		return NewSparseOAUTHKey()
	case OIDCProviderIdentity:
		return NewSparseOIDCProvider()
	case OrganizationalMetadataIdentity:
		return NewSparseOrganizationalMetadata()
	case PacketReportIdentity:
		return NewSparsePacketReport()
	case PasswordResetIdentity:
		return NewSparsePasswordReset()
	case PCCProviderIdentity:
		return NewSparsePCCProvider()
	case PCSearchResultIdentity:
		return NewSparsePCSearchResult()
	case PCTimeRangeIdentity:
		return NewSparsePCTimeRange()
	case PingProbeIdentity:
		return NewSparsePingProbe()
	case PingRequestIdentity:
		return NewSparsePingRequest()
	case PingResultIdentity:
		return NewSparsePingResult()
	case PlanIdentity:
		return NewSparsePlan()
	case PokeIdentity:
		return NewSparsePoke()
	case PolicyIdentity:
		return NewSparsePolicy()
	case PolicyGraphIdentity:
		return NewSparsePolicyGraph()
	case PolicyRefreshIdentity:
		return NewSparsePolicyRefresh()
	case PolicyRendererIdentity:
		return NewSparsePolicyRenderer()
	case PolicyRuleIdentity:
		return NewSparsePolicyRule()
	case PolicyTTLIdentity:
		return NewSparsePolicyTTL()
	case PollAccountIdentity:
		return NewSparsePollAccount()
	case ProcessingUnitIdentity:
		return NewSparseProcessingUnit()
	case ProcessingUnitPolicyIdentity:
		return NewSparseProcessingUnitPolicy()
	case ProcessingUnitRefreshIdentity:
		return NewSparseProcessingUnitRefresh()
	case PUTrafficActionIdentity:
		return NewSparsePUTrafficAction()
	case QuotaCheckIdentity:
		return NewSparseQuotaCheck()
	case QuotaPolicyIdentity:
		return NewSparseQuotaPolicy()
	case RecipeIdentity:
		return NewSparseRecipe()
	case RemoteProcessorIdentity:
		return NewSparseRemoteProcessor()
	case RenderedPolicyIdentity:
		return NewSparseRenderedPolicy()
	case RenderTemplateIdentity:
		return NewSparseRenderTemplate()
	case ReportIdentity:
		return NewSparseReport()
	case ReportsQueryIdentity:
		return NewSparseReportsQuery()
	case RevocationIdentity:
		return NewSparseRevocation()
	case RoleIdentity:
		return NewSparseRole()
	case SAMLProviderIdentity:
		return NewSparseSAMLProvider()
	case SandboxIdentity:
		return NewSparseSandbox()
	case SearchIdentity:
		return NewSparseSearch()
	case ServiceIdentity:
		return NewSparseService()
	case ServiceCertificateIdentity:
		return NewSparseServiceCertificate()
	case ServiceDependencyPolicyIdentity:
		return NewSparseServiceDependencyPolicy()
	case ServicePublicationIdentity:
		return NewSparseServicePublication()
	case ServiceTokenIdentity:
		return NewSparseServiceToken()
	case SquallTagIdentity:
		return NewSparseSquallTag()
	case SSHAuthorityIdentity:
		return NewSparseSSHAuthority()
	case SSHAuthorizationPolicyIdentity:
		return NewSparseSSHAuthorizationPolicy()
	case SSHCertificateIdentity:
		return NewSparseSSHCertificate()
	case SSHIdentityIdentity:
		return NewSparseSSHIdentity()
	case StatsInfoIdentity:
		return NewSparseStatsInfo()
	case StatsQueryIdentity:
		return NewSparseStatsQuery()
	case SuggestedPolicyIdentity:
		return NewSparseSuggestedPolicy()
	case TagIdentity:
		return NewSparseTag()
	case TagInjectIdentity:
		return NewSparseTagInject()
	case TagPrefixIdentity:
		return NewSparseTagPrefix()
	case TagValueIdentity:
		return NewSparseTagValue()
	case TenantIdentity:
		return NewSparseTenant()
	case TextIndexIdentity:
		return NewSparseTextIndex()
	case TokenIdentity:
		return NewSparseToken()
	case TokenScopePolicyIdentity:
		return NewSparseTokenScopePolicy()
	case TriggerIdentity:
		return NewSparseTrigger()
	case TrustedCAIdentity:
		return NewSparseTrustedCA()
	case TrustedNamespaceIdentity:
		return NewSparseTrustedNamespace()
	case UserAccessPolicyIdentity:
		return NewSparseUserAccessPolicy()
	case ValidateRQLReqIdentity:
		return NewSparseValidateRQLReq()
	case ValidateUIParameterIdentity:
		return NewSparseValidateUIParameter()
	case VulnerabilityIdentity:
		return NewSparseVulnerability()
	case X509CertificateIdentity:
		return NewSparseX509Certificate()
	case X509CertificateCheckIdentity:
		return NewSparseX509CertificateCheck()
	default:
		return nil
	}
}

func (f modelManager) Indexes(identity elemental.Identity) [][]string {

	return indexesMap[identity.Name]
}

func (f modelManager) IdentifiableFromString(any string) elemental.Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f modelManager) Identifiables(identity elemental.Identity) elemental.Identifiables {

	switch identity {

	case AccessibleNamespaceIdentity:
		return &AccessibleNamespacesList{}
	case AccessReportIdentity:
		return &AccessReportsList{}
	case AccountIdentity:
		return &AccountsList{}
	case AccountCheckIdentity:
		return &AccountChecksList{}
	case ActivateIdentity:
		return &ActivatesList{}
	case ActivityIdentity:
		return &ActivitiesList{}
	case AlarmIdentity:
		return &AlarmsList{}
	case APIAuthorizationPolicyIdentity:
		return &APIAuthorizationPoliciesList{}
	case APICheckIdentity:
		return &APIChecksList{}
	case AppIdentity:
		return &AppsList{}
	case AppCredentialIdentity:
		return &AppCredentialsList{}
	case AuditProfileIdentity:
		return &AuditProfilesList{}
	case AuditProfileMappingPolicyIdentity:
		return &AuditProfileMappingPoliciesList{}
	case AuditReportIdentity:
		return &AuditReportsList{}
	case AuthnIdentity:
		return &AuthnsList{}
	case AuthorityIdentity:
		return &AuthoritiesList{}
	case AuthzIdentity:
		return &AuthzsList{}
	case AutomationIdentity:
		return &AutomationsList{}
	case AutomationTemplateIdentity:
		return &AutomationTemplatesList{}
	case CachedFlowReportIdentity:
		return &CachedFlowReportsList{}
	case CategoryIdentity:
		return &CategoriesList{}
	case ClaimsIdentity:
		return &ClaimsList{}
	case ClauseMatchIdentity:
		return &ClauseMatchesList{}
	case CloudAccountCleanerIdentity:
		return &CloudAccountCleanersList{}
	case CloudAlertIdentity:
		return &CloudAlertsList{}
	case CloudAlertRuleIdentity:
		return &CloudAlertRulesList{}
	case CloudEndpointIdentity:
		return &CloudEndpointsList{}
	case CloudGraphIdentity:
		return &CloudGraphsList{}
	case CloudManagedNetworkIdentity:
		return &CloudManagedNetworksList{}
	case CloudNetworkInterfaceIdentity:
		return &CloudNetworkInterfacesList{}
	case CloudNetworkQueryIdentity:
		return &CloudNetworkQueriesList{}
	case CloudNetworkRuleSetIdentity:
		return &CloudNetworkRuleSetsList{}
	case CloudNodeIdentity:
		return &CloudNodesList{}
	case CloudPolicyIdentity:
		return &CloudPoliciesList{}
	case CloudRouteTableIdentity:
		return &CloudRouteTablesList{}
	case CloudScheduledNetworkQueryIdentity:
		return &CloudScheduledNetworkQueriesList{}
	case CloudSnapshotAccountIdentity:
		return &CloudSnapshotAccountsList{}
	case CloudSubnetIdentity:
		return &CloudSubnetsList{}
	case CloudVPCIdentity:
		return &CloudVPCsList{}
	case CNSConfigIdentity:
		return &CNSConfigsList{}
	case CNSRQLQueryIdentity:
		return &CNSRQLQueriesList{}
	case CNSSearchIdentity:
		return &CNSSearchesList{}
	case CNSSuggestionIdentity:
		return &CNSSuggestionsList{}
	case ConnectionExceptionReportIdentity:
		return &ConnectionExceptionReportsList{}
	case CounterReportIdentity:
		return &CounterReportsList{}
	case DataPathCertificateIdentity:
		return &DataPathCertificatesList{}
	case DebugBundleIdentity:
		return &DebugBundlesList{}
	case DefaultEnforcerVersionIdentity:
		return &DefaultEnforcerVersionsList{}
	case DependencyMapIdentity:
		return &DependencyMapsList{}
	case DiscoveryModeIdentity:
		return &DiscoveryModesList{}
	case DNSLookupReportIdentity:
		return &DNSLookupReportsList{}
	case EmailIdentity:
		return &EmailsList{}
	case EnforcerIdentity:
		return &EnforcersList{}
	case EnforcerLogIdentity:
		return &EnforcerLogsList{}
	case EnforcerProfileIdentity:
		return &EnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity:
		return &EnforcerProfileMappingPoliciesList{}
	case EnforcerRefreshIdentity:
		return &EnforcerRefreshsList{}
	case EnforcerReportIdentity:
		return &EnforcerReportsList{}
	case EnforcerTraceReportIdentity:
		return &EnforcerTraceReportsList{}
	case EventLogIdentity:
		return &EventLogsList{}
	case ExportIdentity:
		return &ExportsList{}
	case ExternalNetworkIdentity:
		return &ExternalNetworksList{}
	case FileAccessPolicyIdentity:
		return &FileAccessPoliciesList{}
	case FileAccessReportIdentity:
		return &FileAccessReportsList{}
	case FilePathIdentity:
		return &FilePathsList{}
	case FlowReportIdentity:
		return &FlowReportsList{}
	case GraphEdgeIdentity:
		return &GraphEdgesList{}
	case GraphNodeIdentity:
		return &GraphNodesList{}
	case HealthCheckIdentity:
		return &HealthChecksList{}
	case HitIdentity:
		return &HitsList{}
	case HookPolicyIdentity:
		return &HookPoliciesList{}
	case HostServiceIdentity:
		return &HostServicesList{}
	case HostServiceMappingPolicyIdentity:
		return &HostServiceMappingPoliciesList{}
	case HTTPResourceSpecIdentity:
		return &HTTPResourceSpecsList{}
	case ImportIdentity:
		return &ImportsList{}
	case ImportReferenceIdentity:
		return &ImportReferencesList{}
	case ImportRequestIdentity:
		return &ImportRequestsList{}
	case InfrastructurePolicyIdentity:
		return &InfrastructurePoliciesList{}
	case InstalledAppIdentity:
		return &InstalledAppsList{}
	case IPInfoIdentity:
		return &IPInfosList{}
	case IsolationProfileIdentity:
		return &IsolationProfilesList{}
	case IssueIdentity:
		return &IssuesList{}
	case IssueServiceTokenIdentity:
		return &IssueServiceTokensList{}
	case LDAPProviderIdentity:
		return &LDAPProvidersList{}
	case LocalCAIdentity:
		return &LocalCAsList{}
	case LogIdentity:
		return &LogsList{}
	case LogoutIdentity:
		return &LogoutsList{}
	case MessageIdentity:
		return &MessagesList{}
	case MetricsQueryIdentity:
		return &MetricsQueriesList{}
	case MetricsQueryRangeIdentity:
		return &MetricsQueryRangesList{}
	case NamespaceIdentity:
		return &NamespacesList{}
	case NamespaceInfoIdentity:
		return &NamespaceInfosList{}
	case NamespaceMappingPolicyIdentity:
		return &NamespaceMappingPoliciesList{}
	case NamespacePolicyInfoIdentity:
		return &NamespacePolicyInfosList{}
	case NamespaceRendererIdentity:
		return &NamespaceRenderersList{}
	case NamespaceTypeIdentity:
		return &NamespaceTypesList{}
	case NetworkAccessPolicyIdentity:
		return &NetworkAccessPoliciesList{}
	case NetworkRuleSetPolicyIdentity:
		return &NetworkRuleSetPoliciesList{}
	case OAUTHInfoIdentity:
		return &OAUTHInfosList{}
	case OAUTHKeyIdentity:
		return &OAUTHKeysList{}
	case OIDCProviderIdentity:
		return &OIDCProvidersList{}
	case OrganizationalMetadataIdentity:
		return &OrganizationalMetadatasList{}
	case PacketReportIdentity:
		return &PacketReportsList{}
	case PasswordResetIdentity:
		return &PasswordResetsList{}
	case PCCProviderIdentity:
		return &PCCProvidersList{}
	case PCSearchResultIdentity:
		return &PCSearchResultsList{}
	case PCTimeRangeIdentity:
		return &PCTimeRangesList{}
	case PingProbeIdentity:
		return &PingProbesList{}
	case PingRequestIdentity:
		return &PingRequestsList{}
	case PingResultIdentity:
		return &PingResultsList{}
	case PlanIdentity:
		return &PlansList{}
	case PokeIdentity:
		return &PokesList{}
	case PolicyIdentity:
		return &PoliciesList{}
	case PolicyGraphIdentity:
		return &PolicyGraphsList{}
	case PolicyRefreshIdentity:
		return &PolicyRefreshsList{}
	case PolicyRendererIdentity:
		return &PolicyRenderersList{}
	case PolicyRuleIdentity:
		return &PolicyRulesList{}
	case PolicyTTLIdentity:
		return &PolicyTTLsList{}
	case PollAccountIdentity:
		return &PollAccountsList{}
	case ProcessingUnitIdentity:
		return &ProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity:
		return &ProcessingUnitPoliciesList{}
	case ProcessingUnitRefreshIdentity:
		return &ProcessingUnitRefreshsList{}
	case PUTrafficActionIdentity:
		return &PUTrafficActionsList{}
	case QuotaCheckIdentity:
		return &QuotaChecksList{}
	case QuotaPolicyIdentity:
		return &QuotaPoliciesList{}
	case RecipeIdentity:
		return &RecipesList{}
	case RemoteProcessorIdentity:
		return &RemoteProcessorsList{}
	case RenderedPolicyIdentity:
		return &RenderedPoliciesList{}
	case RenderTemplateIdentity:
		return &RenderTemplatesList{}
	case ReportIdentity:
		return &ReportsList{}
	case ReportsQueryIdentity:
		return &ReportsQueriesList{}
	case RevocationIdentity:
		return &RevocationsList{}
	case RoleIdentity:
		return &RolesList{}
	case SAMLProviderIdentity:
		return &SAMLProvidersList{}
	case SandboxIdentity:
		return &SandboxsList{}
	case SearchIdentity:
		return &SearchesList{}
	case ServiceIdentity:
		return &ServicesList{}
	case ServiceCertificateIdentity:
		return &ServiceCertificatesList{}
	case ServiceDependencyPolicyIdentity:
		return &ServiceDependencyPoliciesList{}
	case ServicePublicationIdentity:
		return &ServicePublicationsList{}
	case ServiceTokenIdentity:
		return &ServiceTokensList{}
	case SquallTagIdentity:
		return &SquallTagsList{}
	case SSHAuthorityIdentity:
		return &SSHAuthoritiesList{}
	case SSHAuthorizationPolicyIdentity:
		return &SSHAuthorizationPoliciesList{}
	case SSHCertificateIdentity:
		return &SSHCertificatesList{}
	case SSHIdentityIdentity:
		return &SSHIdentitiesList{}
	case StatsInfoIdentity:
		return &StatsInfosList{}
	case StatsQueryIdentity:
		return &StatsQueriesList{}
	case SuggestedPolicyIdentity:
		return &SuggestedPoliciesList{}
	case TagIdentity:
		return &TagsList{}
	case TagInjectIdentity:
		return &TagInjectsList{}
	case TagPrefixIdentity:
		return &TagPrefixsList{}
	case TagValueIdentity:
		return &TagValuesList{}
	case TenantIdentity:
		return &TenantsList{}
	case TextIndexIdentity:
		return &TextIndexsList{}
	case TokenIdentity:
		return &TokensList{}
	case TokenScopePolicyIdentity:
		return &TokenScopePoliciesList{}
	case TriggerIdentity:
		return &TriggersList{}
	case TrustedCAIdentity:
		return &TrustedCAsList{}
	case TrustedNamespaceIdentity:
		return &TrustedNamespacesList{}
	case UserAccessPolicyIdentity:
		return &UserAccessPoliciesList{}
	case ValidateRQLReqIdentity:
		return &ValidateRQLReqsList{}
	case ValidateUIParameterIdentity:
		return &ValidateUIParametersList{}
	case VulnerabilityIdentity:
		return &VulnerabilitiesList{}
	case X509CertificateIdentity:
		return &X509CertificatesList{}
	case X509CertificateCheckIdentity:
		return &X509CertificateChecksList{}
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiables(identity elemental.Identity) elemental.SparseIdentifiables {

	switch identity {

	case AccessibleNamespaceIdentity:
		return &SparseAccessibleNamespacesList{}
	case AccessReportIdentity:
		return &SparseAccessReportsList{}
	case AccountIdentity:
		return &SparseAccountsList{}
	case AccountCheckIdentity:
		return &SparseAccountChecksList{}
	case ActivateIdentity:
		return &SparseActivatesList{}
	case ActivityIdentity:
		return &SparseActivitiesList{}
	case AlarmIdentity:
		return &SparseAlarmsList{}
	case APIAuthorizationPolicyIdentity:
		return &SparseAPIAuthorizationPoliciesList{}
	case APICheckIdentity:
		return &SparseAPIChecksList{}
	case AppIdentity:
		return &SparseAppsList{}
	case AppCredentialIdentity:
		return &SparseAppCredentialsList{}
	case AuditProfileIdentity:
		return &SparseAuditProfilesList{}
	case AuditProfileMappingPolicyIdentity:
		return &SparseAuditProfileMappingPoliciesList{}
	case AuditReportIdentity:
		return &SparseAuditReportsList{}
	case AuthnIdentity:
		return &SparseAuthnsList{}
	case AuthorityIdentity:
		return &SparseAuthoritiesList{}
	case AuthzIdentity:
		return &SparseAuthzsList{}
	case AutomationIdentity:
		return &SparseAutomationsList{}
	case AutomationTemplateIdentity:
		return &SparseAutomationTemplatesList{}
	case CachedFlowReportIdentity:
		return &SparseCachedFlowReportsList{}
	case CategoryIdentity:
		return &SparseCategoriesList{}
	case ClaimsIdentity:
		return &SparseClaimsList{}
	case ClauseMatchIdentity:
		return &SparseClauseMatchesList{}
	case CloudAccountCleanerIdentity:
		return &SparseCloudAccountCleanersList{}
	case CloudAlertIdentity:
		return &SparseCloudAlertsList{}
	case CloudAlertRuleIdentity:
		return &SparseCloudAlertRulesList{}
	case CloudEndpointIdentity:
		return &SparseCloudEndpointsList{}
	case CloudGraphIdentity:
		return &SparseCloudGraphsList{}
	case CloudManagedNetworkIdentity:
		return &SparseCloudManagedNetworksList{}
	case CloudNetworkInterfaceIdentity:
		return &SparseCloudNetworkInterfacesList{}
	case CloudNetworkQueryIdentity:
		return &SparseCloudNetworkQueriesList{}
	case CloudNetworkRuleSetIdentity:
		return &SparseCloudNetworkRuleSetsList{}
	case CloudNodeIdentity:
		return &SparseCloudNodesList{}
	case CloudPolicyIdentity:
		return &SparseCloudPoliciesList{}
	case CloudRouteTableIdentity:
		return &SparseCloudRouteTablesList{}
	case CloudScheduledNetworkQueryIdentity:
		return &SparseCloudScheduledNetworkQueriesList{}
	case CloudSnapshotAccountIdentity:
		return &SparseCloudSnapshotAccountsList{}
	case CloudSubnetIdentity:
		return &SparseCloudSubnetsList{}
	case CloudVPCIdentity:
		return &SparseCloudVPCsList{}
	case CNSConfigIdentity:
		return &SparseCNSConfigsList{}
	case CNSRQLQueryIdentity:
		return &SparseCNSRQLQueriesList{}
	case CNSSearchIdentity:
		return &SparseCNSSearchesList{}
	case CNSSuggestionIdentity:
		return &SparseCNSSuggestionsList{}
	case ConnectionExceptionReportIdentity:
		return &SparseConnectionExceptionReportsList{}
	case CounterReportIdentity:
		return &SparseCounterReportsList{}
	case DataPathCertificateIdentity:
		return &SparseDataPathCertificatesList{}
	case DebugBundleIdentity:
		return &SparseDebugBundlesList{}
	case DefaultEnforcerVersionIdentity:
		return &SparseDefaultEnforcerVersionsList{}
	case DependencyMapIdentity:
		return &SparseDependencyMapsList{}
	case DiscoveryModeIdentity:
		return &SparseDiscoveryModesList{}
	case DNSLookupReportIdentity:
		return &SparseDNSLookupReportsList{}
	case EmailIdentity:
		return &SparseEmailsList{}
	case EnforcerIdentity:
		return &SparseEnforcersList{}
	case EnforcerLogIdentity:
		return &SparseEnforcerLogsList{}
	case EnforcerProfileIdentity:
		return &SparseEnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity:
		return &SparseEnforcerProfileMappingPoliciesList{}
	case EnforcerRefreshIdentity:
		return &SparseEnforcerRefreshsList{}
	case EnforcerReportIdentity:
		return &SparseEnforcerReportsList{}
	case EnforcerTraceReportIdentity:
		return &SparseEnforcerTraceReportsList{}
	case EventLogIdentity:
		return &SparseEventLogsList{}
	case ExportIdentity:
		return &SparseExportsList{}
	case ExternalNetworkIdentity:
		return &SparseExternalNetworksList{}
	case FileAccessPolicyIdentity:
		return &SparseFileAccessPoliciesList{}
	case FileAccessReportIdentity:
		return &SparseFileAccessReportsList{}
	case FilePathIdentity:
		return &SparseFilePathsList{}
	case FlowReportIdentity:
		return &SparseFlowReportsList{}
	case GraphEdgeIdentity:
		return &SparseGraphEdgesList{}
	case GraphNodeIdentity:
		return &SparseGraphNodesList{}
	case HealthCheckIdentity:
		return &SparseHealthChecksList{}
	case HitIdentity:
		return &SparseHitsList{}
	case HookPolicyIdentity:
		return &SparseHookPoliciesList{}
	case HostServiceIdentity:
		return &SparseHostServicesList{}
	case HostServiceMappingPolicyIdentity:
		return &SparseHostServiceMappingPoliciesList{}
	case HTTPResourceSpecIdentity:
		return &SparseHTTPResourceSpecsList{}
	case ImportIdentity:
		return &SparseImportsList{}
	case ImportReferenceIdentity:
		return &SparseImportReferencesList{}
	case ImportRequestIdentity:
		return &SparseImportRequestsList{}
	case InfrastructurePolicyIdentity:
		return &SparseInfrastructurePoliciesList{}
	case InstalledAppIdentity:
		return &SparseInstalledAppsList{}
	case IPInfoIdentity:
		return &SparseIPInfosList{}
	case IsolationProfileIdentity:
		return &SparseIsolationProfilesList{}
	case IssueIdentity:
		return &SparseIssuesList{}
	case IssueServiceTokenIdentity:
		return &SparseIssueServiceTokensList{}
	case LDAPProviderIdentity:
		return &SparseLDAPProvidersList{}
	case LocalCAIdentity:
		return &SparseLocalCAsList{}
	case LogIdentity:
		return &SparseLogsList{}
	case LogoutIdentity:
		return &SparseLogoutsList{}
	case MessageIdentity:
		return &SparseMessagesList{}
	case MetricsQueryIdentity:
		return &SparseMetricsQueriesList{}
	case MetricsQueryRangeIdentity:
		return &SparseMetricsQueryRangesList{}
	case NamespaceIdentity:
		return &SparseNamespacesList{}
	case NamespaceInfoIdentity:
		return &SparseNamespaceInfosList{}
	case NamespaceMappingPolicyIdentity:
		return &SparseNamespaceMappingPoliciesList{}
	case NamespacePolicyInfoIdentity:
		return &SparseNamespacePolicyInfosList{}
	case NamespaceRendererIdentity:
		return &SparseNamespaceRenderersList{}
	case NamespaceTypeIdentity:
		return &SparseNamespaceTypesList{}
	case NetworkAccessPolicyIdentity:
		return &SparseNetworkAccessPoliciesList{}
	case NetworkRuleSetPolicyIdentity:
		return &SparseNetworkRuleSetPoliciesList{}
	case OAUTHInfoIdentity:
		return &SparseOAUTHInfosList{}
	case OAUTHKeyIdentity:
		return &SparseOAUTHKeysList{}
	case OIDCProviderIdentity:
		return &SparseOIDCProvidersList{}
	case OrganizationalMetadataIdentity:
		return &SparseOrganizationalMetadatasList{}
	case PacketReportIdentity:
		return &SparsePacketReportsList{}
	case PasswordResetIdentity:
		return &SparsePasswordResetsList{}
	case PCCProviderIdentity:
		return &SparsePCCProvidersList{}
	case PCSearchResultIdentity:
		return &SparsePCSearchResultsList{}
	case PCTimeRangeIdentity:
		return &SparsePCTimeRangesList{}
	case PingProbeIdentity:
		return &SparsePingProbesList{}
	case PingRequestIdentity:
		return &SparsePingRequestsList{}
	case PingResultIdentity:
		return &SparsePingResultsList{}
	case PlanIdentity:
		return &SparsePlansList{}
	case PokeIdentity:
		return &SparsePokesList{}
	case PolicyIdentity:
		return &SparsePoliciesList{}
	case PolicyGraphIdentity:
		return &SparsePolicyGraphsList{}
	case PolicyRefreshIdentity:
		return &SparsePolicyRefreshsList{}
	case PolicyRendererIdentity:
		return &SparsePolicyRenderersList{}
	case PolicyRuleIdentity:
		return &SparsePolicyRulesList{}
	case PolicyTTLIdentity:
		return &SparsePolicyTTLsList{}
	case PollAccountIdentity:
		return &SparsePollAccountsList{}
	case ProcessingUnitIdentity:
		return &SparseProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity:
		return &SparseProcessingUnitPoliciesList{}
	case ProcessingUnitRefreshIdentity:
		return &SparseProcessingUnitRefreshsList{}
	case PUTrafficActionIdentity:
		return &SparsePUTrafficActionsList{}
	case QuotaCheckIdentity:
		return &SparseQuotaChecksList{}
	case QuotaPolicyIdentity:
		return &SparseQuotaPoliciesList{}
	case RecipeIdentity:
		return &SparseRecipesList{}
	case RemoteProcessorIdentity:
		return &SparseRemoteProcessorsList{}
	case RenderedPolicyIdentity:
		return &SparseRenderedPoliciesList{}
	case RenderTemplateIdentity:
		return &SparseRenderTemplatesList{}
	case ReportIdentity:
		return &SparseReportsList{}
	case ReportsQueryIdentity:
		return &SparseReportsQueriesList{}
	case RevocationIdentity:
		return &SparseRevocationsList{}
	case RoleIdentity:
		return &SparseRolesList{}
	case SAMLProviderIdentity:
		return &SparseSAMLProvidersList{}
	case SandboxIdentity:
		return &SparseSandboxsList{}
	case SearchIdentity:
		return &SparseSearchesList{}
	case ServiceIdentity:
		return &SparseServicesList{}
	case ServiceCertificateIdentity:
		return &SparseServiceCertificatesList{}
	case ServiceDependencyPolicyIdentity:
		return &SparseServiceDependencyPoliciesList{}
	case ServicePublicationIdentity:
		return &SparseServicePublicationsList{}
	case ServiceTokenIdentity:
		return &SparseServiceTokensList{}
	case SquallTagIdentity:
		return &SparseSquallTagsList{}
	case SSHAuthorityIdentity:
		return &SparseSSHAuthoritiesList{}
	case SSHAuthorizationPolicyIdentity:
		return &SparseSSHAuthorizationPoliciesList{}
	case SSHCertificateIdentity:
		return &SparseSSHCertificatesList{}
	case SSHIdentityIdentity:
		return &SparseSSHIdentitiesList{}
	case StatsInfoIdentity:
		return &SparseStatsInfosList{}
	case StatsQueryIdentity:
		return &SparseStatsQueriesList{}
	case SuggestedPolicyIdentity:
		return &SparseSuggestedPoliciesList{}
	case TagIdentity:
		return &SparseTagsList{}
	case TagInjectIdentity:
		return &SparseTagInjectsList{}
	case TagPrefixIdentity:
		return &SparseTagPrefixsList{}
	case TagValueIdentity:
		return &SparseTagValuesList{}
	case TenantIdentity:
		return &SparseTenantsList{}
	case TextIndexIdentity:
		return &SparseTextIndexsList{}
	case TokenIdentity:
		return &SparseTokensList{}
	case TokenScopePolicyIdentity:
		return &SparseTokenScopePoliciesList{}
	case TriggerIdentity:
		return &SparseTriggersList{}
	case TrustedCAIdentity:
		return &SparseTrustedCAsList{}
	case TrustedNamespaceIdentity:
		return &SparseTrustedNamespacesList{}
	case UserAccessPolicyIdentity:
		return &SparseUserAccessPoliciesList{}
	case ValidateRQLReqIdentity:
		return &SparseValidateRQLReqsList{}
	case ValidateUIParameterIdentity:
		return &SparseValidateUIParametersList{}
	case VulnerabilityIdentity:
		return &SparseVulnerabilitiesList{}
	case X509CertificateIdentity:
		return &SparseX509CertificatesList{}
	case X509CertificateCheckIdentity:
		return &SparseX509CertificateChecksList{}
	default:
		return nil
	}
}

func (f modelManager) IdentifiablesFromString(any string) elemental.Identifiables {

	return f.Identifiables(f.IdentityFromAny(any))
}

func (f modelManager) Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func (f modelManager) AllIdentities() []elemental.Identity {
	return AllIdentities()
}

var manager = modelManager{}

// Manager returns the model elemental.ModelManager.
func Manager() elemental.ModelManager { return manager }

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AccessibleNamespaceIdentity,
		AccessReportIdentity,
		AccountIdentity,
		AccountCheckIdentity,
		ActivateIdentity,
		ActivityIdentity,
		AlarmIdentity,
		APIAuthorizationPolicyIdentity,
		APICheckIdentity,
		AppIdentity,
		AppCredentialIdentity,
		AuditProfileIdentity,
		AuditProfileMappingPolicyIdentity,
		AuditReportIdentity,
		AuthnIdentity,
		AuthorityIdentity,
		AuthzIdentity,
		AutomationIdentity,
		AutomationTemplateIdentity,
		CachedFlowReportIdentity,
		CategoryIdentity,
		ClaimsIdentity,
		ClauseMatchIdentity,
		CloudAccountCleanerIdentity,
		CloudAlertIdentity,
		CloudAlertRuleIdentity,
		CloudEndpointIdentity,
		CloudGraphIdentity,
		CloudManagedNetworkIdentity,
		CloudNetworkInterfaceIdentity,
		CloudNetworkQueryIdentity,
		CloudNetworkRuleSetIdentity,
		CloudNodeIdentity,
		CloudPolicyIdentity,
		CloudRouteTableIdentity,
		CloudScheduledNetworkQueryIdentity,
		CloudSnapshotAccountIdentity,
		CloudSubnetIdentity,
		CloudVPCIdentity,
		CNSConfigIdentity,
		CNSRQLQueryIdentity,
		CNSSearchIdentity,
		CNSSuggestionIdentity,
		ConnectionExceptionReportIdentity,
		CounterReportIdentity,
		DataPathCertificateIdentity,
		DebugBundleIdentity,
		DefaultEnforcerVersionIdentity,
		DependencyMapIdentity,
		DiscoveryModeIdentity,
		DNSLookupReportIdentity,
		EmailIdentity,
		EnforcerIdentity,
		EnforcerLogIdentity,
		EnforcerProfileIdentity,
		EnforcerProfileMappingPolicyIdentity,
		EnforcerRefreshIdentity,
		EnforcerReportIdentity,
		EnforcerTraceReportIdentity,
		EventLogIdentity,
		ExportIdentity,
		ExternalNetworkIdentity,
		FileAccessPolicyIdentity,
		FileAccessReportIdentity,
		FilePathIdentity,
		FlowReportIdentity,
		GraphEdgeIdentity,
		GraphNodeIdentity,
		HealthCheckIdentity,
		HitIdentity,
		HookPolicyIdentity,
		HostServiceIdentity,
		HostServiceMappingPolicyIdentity,
		HTTPResourceSpecIdentity,
		ImportIdentity,
		ImportReferenceIdentity,
		ImportRequestIdentity,
		InfrastructurePolicyIdentity,
		InstalledAppIdentity,
		IPInfoIdentity,
		IsolationProfileIdentity,
		IssueIdentity,
		IssueServiceTokenIdentity,
		LDAPProviderIdentity,
		LocalCAIdentity,
		LogIdentity,
		LogoutIdentity,
		MessageIdentity,
		MetricsQueryIdentity,
		MetricsQueryRangeIdentity,
		NamespaceIdentity,
		NamespaceInfoIdentity,
		NamespaceMappingPolicyIdentity,
		NamespacePolicyInfoIdentity,
		NamespaceRendererIdentity,
		NamespaceTypeIdentity,
		NetworkAccessPolicyIdentity,
		NetworkRuleSetPolicyIdentity,
		OAUTHInfoIdentity,
		OAUTHKeyIdentity,
		OIDCProviderIdentity,
		OrganizationalMetadataIdentity,
		PacketReportIdentity,
		PasswordResetIdentity,
		PCCProviderIdentity,
		PCSearchResultIdentity,
		PCTimeRangeIdentity,
		PingProbeIdentity,
		PingRequestIdentity,
		PingResultIdentity,
		PlanIdentity,
		PokeIdentity,
		PolicyIdentity,
		PolicyGraphIdentity,
		PolicyRefreshIdentity,
		PolicyRendererIdentity,
		PolicyRuleIdentity,
		PolicyTTLIdentity,
		PollAccountIdentity,
		ProcessingUnitIdentity,
		ProcessingUnitPolicyIdentity,
		ProcessingUnitRefreshIdentity,
		PUTrafficActionIdentity,
		QuotaCheckIdentity,
		QuotaPolicyIdentity,
		RecipeIdentity,
		RemoteProcessorIdentity,
		RenderedPolicyIdentity,
		RenderTemplateIdentity,
		ReportIdentity,
		ReportsQueryIdentity,
		RevocationIdentity,
		RoleIdentity,
		RootIdentity,
		SAMLProviderIdentity,
		SandboxIdentity,
		SearchIdentity,
		ServiceIdentity,
		ServiceCertificateIdentity,
		ServiceDependencyPolicyIdentity,
		ServicePublicationIdentity,
		ServiceTokenIdentity,
		SquallTagIdentity,
		SSHAuthorityIdentity,
		SSHAuthorizationPolicyIdentity,
		SSHCertificateIdentity,
		SSHIdentityIdentity,
		StatsInfoIdentity,
		StatsQueryIdentity,
		SuggestedPolicyIdentity,
		TagIdentity,
		TagInjectIdentity,
		TagPrefixIdentity,
		TagValueIdentity,
		TenantIdentity,
		TextIndexIdentity,
		TokenIdentity,
		TokenScopePolicyIdentity,
		TriggerIdentity,
		TrustedCAIdentity,
		TrustedNamespaceIdentity,
		UserAccessPolicyIdentity,
		ValidateRQLReqIdentity,
		ValidateUIParameterIdentity,
		VulnerabilityIdentity,
		X509CertificateIdentity,
		X509CertificateCheckIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case AccessibleNamespaceIdentity:
		return []string{
			"accns",
		}
	case AccessReportIdentity:
		return []string{}
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
	case AppCredentialIdentity:
		return []string{
			"appcred",
			"appcreds",
		}
	case AuditProfileIdentity:
		return []string{
			"ap",
		}
	case AuditProfileMappingPolicyIdentity:
		return []string{
			"audpol",
			"audpols",
		}
	case AuditReportIdentity:
		return []string{}
	case AuthnIdentity:
		return []string{}
	case AuthorityIdentity:
		return []string{
			"ca",
		}
	case AuthzIdentity:
		return []string{}
	case AutomationIdentity:
		return []string{
			"autos",
			"auto",
		}
	case AutomationTemplateIdentity:
		return []string{
			"autotmpl",
		}
	case CachedFlowReportIdentity:
		return []string{}
	case CategoryIdentity:
		return []string{}
	case ClaimsIdentity:
		return []string{}
	case ClauseMatchIdentity:
		return []string{}
	case CloudAccountCleanerIdentity:
		return []string{}
	case CloudAlertIdentity:
		return []string{}
	case CloudAlertRuleIdentity:
		return []string{}
	case CloudEndpointIdentity:
		return []string{}
	case CloudGraphIdentity:
		return []string{}
	case CloudManagedNetworkIdentity:
		return []string{}
	case CloudNetworkInterfaceIdentity:
		return []string{}
	case CloudNetworkQueryIdentity:
		return []string{}
	case CloudNetworkRuleSetIdentity:
		return []string{
			"crules",
		}
	case CloudNodeIdentity:
		return []string{}
	case CloudPolicyIdentity:
		return []string{}
	case CloudRouteTableIdentity:
		return []string{}
	case CloudScheduledNetworkQueryIdentity:
		return []string{}
	case CloudSnapshotAccountIdentity:
		return []string{}
	case CloudSubnetIdentity:
		return []string{}
	case CloudVPCIdentity:
		return []string{
			"vpc",
			"vpcs",
		}
	case CNSConfigIdentity:
		return []string{
			"pcc",
		}
	case CNSRQLQueryIdentity:
		return []string{}
	case CNSSearchIdentity:
		return []string{}
	case CNSSuggestionIdentity:
		return []string{}
	case ConnectionExceptionReportIdentity:
		return []string{}
	case CounterReportIdentity:
		return []string{}
	case DataPathCertificateIdentity:
		return []string{}
	case DebugBundleIdentity:
		return []string{}
	case DefaultEnforcerVersionIdentity:
		return []string{}
	case DependencyMapIdentity:
		return []string{
			"depmaps",
			"depmap",
		}
	case DiscoveryModeIdentity:
		return []string{}
	case DNSLookupReportIdentity:
		return []string{}
	case EmailIdentity:
		return []string{}
	case EnforcerIdentity:
		return []string{
			"defender",
		}
	case EnforcerLogIdentity:
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
			"epm",
		}
	case EnforcerRefreshIdentity:
		return []string{}
	case EnforcerReportIdentity:
		return []string{}
	case EnforcerTraceReportIdentity:
		return []string{}
	case EventLogIdentity:
		return []string{}
	case ExportIdentity:
		return []string{}
	case ExternalNetworkIdentity:
		return []string{
			"extnet",
			"extnets",
		}
	case FileAccessPolicyIdentity:
		return []string{}
	case FileAccessReportIdentity:
		return []string{}
	case FilePathIdentity:
		return []string{
			"fp",
			"fps",
		}
	case FlowReportIdentity:
		return []string{}
	case GraphEdgeIdentity:
		return []string{}
	case GraphNodeIdentity:
		return []string{}
	case HealthCheckIdentity:
		return []string{}
	case HitIdentity:
		return []string{}
	case HookPolicyIdentity:
		return []string{
			"hook",
			"hooks",
			"hookpol",
			"hookpols",
		}
	case HostServiceIdentity:
		return []string{
			"hostsrv",
			"hostsrvs",
		}
	case HostServiceMappingPolicyIdentity:
		return []string{
			"hostsrvmappol",
			"hostsrvmappols",
		}
	case HTTPResourceSpecIdentity:
		return []string{
			"httpresource",
			"resource",
			"httpspec",
		}
	case ImportIdentity:
		return []string{}
	case ImportReferenceIdentity:
		return []string{
			"importref",
			"impref",
		}
	case ImportRequestIdentity:
		return []string{
			"req",
			"reqs",
			"ireq",
			"ireqs",
		}
	case InfrastructurePolicyIdentity:
		return []string{
			"infrapol",
			"infrapols",
		}
	case InstalledAppIdentity:
		return []string{
			"iapps",
			"iapp",
		}
	case IPInfoIdentity:
		return []string{}
	case IsolationProfileIdentity:
		return []string{
			"ip",
		}
	case IssueIdentity:
		return []string{}
	case IssueServiceTokenIdentity:
		return []string{}
	case LDAPProviderIdentity:
		return []string{}
	case LocalCAIdentity:
		return []string{}
	case LogIdentity:
		return []string{}
	case LogoutIdentity:
		return []string{}
	case MessageIdentity:
		return []string{
			"mess",
		}
	case MetricsQueryIdentity:
		return []string{
			"mq",
		}
	case MetricsQueryRangeIdentity:
		return []string{
			"mqr",
		}
	case NamespaceIdentity:
		return []string{
			"ns",
		}
	case NamespaceInfoIdentity:
		return []string{}
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nspolicy",
			"nspolicies",
			"nsmap",
			"nsmaps",
		}
	case NamespacePolicyInfoIdentity:
		return []string{}
	case NamespaceRendererIdentity:
		return []string{
			"nsrenderer",
		}
	case NamespaceTypeIdentity:
		return []string{}
	case NetworkAccessPolicyIdentity:
		return []string{
			"netpol",
			"netpols",
		}
	case NetworkRuleSetPolicyIdentity:
		return []string{
			"netruleset",
			"netrulesets",
			"netset",
			"netsets",
			"networkruleset",
			"networkrulesets",
		}
	case OAUTHInfoIdentity:
		return []string{}
	case OAUTHKeyIdentity:
		return []string{}
	case OIDCProviderIdentity:
		return []string{}
	case OrganizationalMetadataIdentity:
		return []string{
			"om",
		}
	case PacketReportIdentity:
		return []string{}
	case PasswordResetIdentity:
		return []string{}
	case PCCProviderIdentity:
		return []string{}
	case PCSearchResultIdentity:
		return []string{}
	case PCTimeRangeIdentity:
		return []string{}
	case PingProbeIdentity:
		return []string{}
	case PingRequestIdentity:
		return []string{}
	case PingResultIdentity:
		return []string{}
	case PlanIdentity:
		return []string{}
	case PokeIdentity:
		return []string{}
	case PolicyIdentity:
		return []string{}
	case PolicyGraphIdentity:
		return []string{
			"polgraph",
		}
	case PolicyRefreshIdentity:
		return []string{}
	case PolicyRendererIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case PolicyTTLIdentity:
		return []string{}
	case PollAccountIdentity:
		return []string{}
	case ProcessingUnitIdentity:
		return []string{
			"pu",
			"pus",
		}
	case ProcessingUnitPolicyIdentity:
		return []string{
			"pup",
			"pups",
		}
	case ProcessingUnitRefreshIdentity:
		return []string{}
	case PUTrafficActionIdentity:
		return []string{}
	case QuotaCheckIdentity:
		return []string{}
	case QuotaPolicyIdentity:
		return []string{
			"quota",
			"quotas",
			"quotapol",
			"quotapols",
		}
	case RecipeIdentity:
		return []string{
			"rcp",
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
	case RenderTemplateIdentity:
		return []string{
			"cook",
			"rtpl",
		}
	case ReportIdentity:
		return []string{}
	case ReportsQueryIdentity:
		return []string{
			"rq",
		}
	case RevocationIdentity:
		return []string{}
	case RoleIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case SAMLProviderIdentity:
		return []string{}
	case SandboxIdentity:
		return []string{}
	case SearchIdentity:
		return []string{}
	case ServiceIdentity:
		return []string{
			"srv",
		}
	case ServiceCertificateIdentity:
		return []string{}
	case ServiceDependencyPolicyIdentity:
		return []string{
			"srvdep",
			"srvdeps",
		}
	case ServicePublicationIdentity:
		return []string{}
	case ServiceTokenIdentity:
		return []string{}
	case SquallTagIdentity:
		return []string{}
	case SSHAuthorityIdentity:
		return []string{}
	case SSHAuthorizationPolicyIdentity:
		return []string{
			"sshpol",
			"sshpols",
		}
	case SSHCertificateIdentity:
		return []string{}
	case SSHIdentityIdentity:
		return []string{}
	case StatsInfoIdentity:
		return []string{
			"si",
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
	case TagIdentity:
		return []string{}
	case TagInjectIdentity:
		return []string{}
	case TagPrefixIdentity:
		return []string{}
	case TagValueIdentity:
		return []string{}
	case TenantIdentity:
		return []string{}
	case TextIndexIdentity:
		return []string{}
	case TokenIdentity:
		return []string{}
	case TokenScopePolicyIdentity:
		return []string{
			"tsp",
		}
	case TriggerIdentity:
		return []string{}
	case TrustedCAIdentity:
		return []string{}
	case TrustedNamespaceIdentity:
		return []string{
			"trustedns",
		}
	case UserAccessPolicyIdentity:
		return []string{
			"usrpol",
			"usrpols",
		}
	case ValidateRQLReqIdentity:
		return []string{}
	case ValidateUIParameterIdentity:
		return []string{
			"validparam",
		}
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
