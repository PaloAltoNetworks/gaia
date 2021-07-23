# Model
model:
  rest_name: root
  resource_name: root
  entity_name: Root
  package: root
  group: core
  description: root object.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@identifiable-not-stored'
  root: true

# Relations
relations:
- rest_name: accessiblenamespace
  get:
    description: Retrieves the list of accessible namespaces.
    global_parameters:
    - $filtering

- rest_name: accessreport
  create:
    description: Create an access report.

- rest_name: account
  get:
    description: |-
      Retrieves all accounts. This is a private API that can only be done by the
      system.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: name
        description: internal parameters.
        type: string
        example_value: account

      - name: status
        description: internal parameters.
        type: string
        example_value: status
  create:
    description: Creates a new account.

- rest_name: accountcheck
  create:
    description: Verifies account credentials.

- rest_name: activate
  get:
    description: Activates a pending account.
    parameters:
      required:
      - - - token
      entries:
      - name: noRedirect
        description: If set, do not redirect the request to the web interface.
        type: boolean

      - name: token
        description: Activation token.
        type: string
        example_value: xxx-xxx-xxx-xxx

- rest_name: activity
  get:
    description: Retrieves the list of activity logs.
    global_parameters:
    - $filtering

- rest_name: alarm
  get:
    description: Retrieves all the alarms.
    global_parameters:
    - $filtering
  create:
    description: Creates a new alarm.

- rest_name: apiauthorizationpolicy
  get:
    description: Retrieves the list of API authorizations.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new API authorization.

- rest_name: apicheck
  create:
    description: Verifies the authorizations on various identities for a given token.

- rest_name: app
  get:
    description: Retrieves the list of apps.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: name
        description: internal parameter.
        type: string
        example_value: claire

- rest_name: appcredential
  get:
    description: Retrieves the list of app credentials.
    global_parameters:
    - $filtering
  create:
    description: Creates a new app credential.

- rest_name: auditprofile
  get:
    description: Retrieves the list of audit profiles.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new audit profile.

- rest_name: auditprofilemappingpolicy
  get:
    description: Retrieves the list of audit profile mapping policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new audit profile mapping policy.

- rest_name: auditreport
  create:
    description: Create a audit statistics report.

- rest_name: authn
  get:
    description: Verify the validity of a token. This is deprecated. You should use Create.
    parameters:
      entries:
      - name: token
        description: token to validate.
        type: string
        example_value: abc.def.ghi
  create:
    description: Verify the validity of a token.

- rest_name: authority
  create:
    description: Creates a new certificate authority.
    global_parameters:
    - $filtering

- rest_name: authz
  create:
    description: Verifies if a request should be accepted.
    parameters:
      entries:
      - name: forwardpolicyrules
        description: If set, return the raw resolved PolicyRulesList.
        type: boolean

- rest_name: automation
  get:
    description: Retrieves the list of automations.
    global_parameters:
    - $filtering
  create:
    description: Creates a new Automation.

- rest_name: automationtemplate
  get:
    description: Retrieves the list of automation templates.

- rest_name: cachedflowreport
  create:
    description: Create a cached flow statistics report.

- rest_name: claims
  get:
    description: Retrieves the list of claims.
    global_parameters:
    - $filtering
  create:
    description: Creates a new claims record.

- rest_name: clausesmatch
  create:
    description: Performs a clause matching.

- rest_name: cloudaccountcleaner
  create:
    description: Initiates a cloud account clean up process for all stale objects.

- rest_name: cloudalertrule
  get:
    description: Retrieves the list of cloud alert rules.
    global_parameters:
    - $filtering

- rest_name: cloudendpoint
  get:
    description: List of endpoints associated with the deployment.
    global_parameters:
    - $filtering
  create:
    description: Creates a cloud endpoint.

- rest_name: cloudgraph
  create:
    description: |-
      Creates a cloud dependency graph based on ingested data and the required
      parameters.

- rest_name: cloudmanagednetwork
  get:
    description: Retrieves the list of cloud attachments.
    global_parameters:
    - $filtering
  create:
    description: Creates a cloud managed network.

- rest_name: cloudnetworkinterface
  get:
    description: Retrieve the list of network interfaces associated with the deployment.
    global_parameters:
    - $filtering
  create:
    description: Creates a cloud network interface.

- rest_name: cloudnetworkquery
  get:
    description: Retrieves the list of cloud network queries.
    global_parameters:
    - $filtering
  create:
    description: Creates a cloud network query.

- rest_name: cloudnetworkruleset
  get:
    description: Retrieves the list of cloud network rule set policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a cloud network ruleset.

- rest_name: cloudnode
  get:
    description: Retrieves the list of cloud nodes.
    global_parameters:
    - $filtering

- rest_name: cloudpolicy
  get:
    description: Retrieves the list of cloud policies.
    global_parameters:
    - $filtering

- rest_name: cloudroutetable
  get:
    description: Retrieves the list of routing tables.
    global_parameters:
    - $filtering
  create:
    description: Creates a new routing table.

- rest_name: cloudschedulednetworkquery
  create:
    description: |-
      Updates the last execution time of scheduled network query which are run to
      evaluate alert rules.

- rest_name: cloudsnapshotaccount
  create:
    description: Initiates a cloud account for snapshot data ingestion.

- rest_name: cloudsubnet
  get:
    description: Retrieves the list of subnets.
    global_parameters:
    - $filtering
  create:
    description: Creates a cloud subnet.

- rest_name: cloudvpc
  get:
    description: Retrieves the list of VPCs.
    global_parameters:
    - $filtering
  create:
    description: Creates a cloud VPC.

- rest_name: cnsconfig
  get:
    description: Retrieves the list of the CNS configuration objects.
    global_parameters:
    - $filtering
  create:
    description: Creates a CNS configuration object for the specified namespace.

- rest_name: cnssearch
  create:
    description: Retrieves RQL search results.

- rest_name: cnssuggestion
  create:
    description: Retrieves RQL suggestions from Microsegmentation.

- rest_name: connectionexceptionreport
  create:
    description: Create a connection exception report.

- rest_name: counterreport
  create:
    description: Create a counter report.

- rest_name: datapathcertificate
  create:
    description: Creates a new certificate for datapath.

- rest_name: defaultenforcerversion
  get:
    description: Returns the default enforcer version of the specified namespace.
  create:
    description: Modify the default enforcer version of the specified namespace.

- rest_name: dependencymap
  get:
    description: Retrieves the dependency map of a namespace.
    global_parameters:
    - $timewindow
    - $flowoffset
    - $filtering
    parameters:
      entries:
      - name: tag
        description: Only show objects with the given tags in the dependency map.
        type: string
        multiple: true
        example_value: env=prod

      - name: view
        description: Set the view query for grouping the dependency map.
        type: string
        example_value: $namespace then app

      - name: viewSuggestions
        description: Also return the view suggestions.
        type: boolean

- rest_name: discoverymode
  get:
    description: (Deprecated) Returns the list of discovery modes.
    deprecated: true
  create:
    description: (Deprecated) Deploy the discovery mode assets onto the specified namespace.
    deprecated: true

- rest_name: dnslookupreport
  create:
    description: Create a DNS Lookup report.

- rest_name: email
  create:
    description: Sends an email.

- rest_name: enforcer
  get:
    description: Retrieves the list of enforcers.
    global_parameters:
    - $filtering
  create:
    description: Creates a new enforcer.

- rest_name: enforcerlog
  get:
    description: Retrieves the list of enforcerlogs.
    global_parameters:
    - $filtering
  create:
    description: Creates a new enforcerlog.

- rest_name: enforcerprofile
  get:
    description: Retrieves the list of enforcer profiles.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new enforcer profile.

- rest_name: enforcerprofilemappingpolicy
  get:
    description: Retrieves the list of enforcer profile mappings.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new enforcer profile mappings.

- rest_name: enforcerrefresh
  create:
    description: Create an enforcer refresh report.

- rest_name: enforcerreport
  create:
    description: Create an enforcer statistics report.

- rest_name: enforcertracereport
  create:
    description: Create an enforcer trace report.

- rest_name: eventlog
  create:
    description: Creates a new event log for a particular entity.

- rest_name: export
  create:
    description: Exports all policies and related objects of a namespace.
    global_parameters:
    - $filtering

- rest_name: externalnetwork
  get:
    description: Retrieves the list of external networks.
    global_parameters:
    - $filtering
    - $archivable
    - $propagatable
  create:
    description: Creates a new external network.

- rest_name: fileaccesspolicy
  get:
    description: Retrieves the list of file access policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new file access policies.

- rest_name: fileaccessreport
  create:
    description: Create a file access statistics report.

- rest_name: filepath
  get:
    description: Retrieves the list of file paths.
    global_parameters:
    - $filtering
    - $archivable
    - $propagatable
  create:
    description: Create a new file path.

- rest_name: flowreport
  create:
    description: Create a flow statistics report.

- rest_name: graphedge
  get:
    description: Retrieves the graph edges.
    global_parameters:
    - $timewindow
    - $filtering

- rest_name: graphnode
  get:
    description: Retrieves the pu nodes.
    global_parameters:
    - $timewindow
    - $archivable
    - $filtering

- rest_name: healthcheck
  get:
    description: Retrieve the health of the platform.
    parameters:
      entries:
      - name: quiet
        description: If set to true, the health check endpoint will not return data but will return 200 OK if everything is fine or 218 if the controller is not operational. This is useful when you want to use the health check endpoint as a load balancer health check.
        type: boolean

- rest_name: hit
  get:
    description: Retrieve a matching hit.
    parameters:
      required:
      - - - name
          - targetID
          - targetIdentity
        - - targetID
          - targetIdentity
      entries:
      - name: name
        description: The name of the counter.
        type: string
        default_value: counter

      - name: targetID
        description: The ID of the object associated to the counter.
        type: string
        example_value: xyz

      - name: targetIdentity
        description: The identity of the object associated to the counter.
        type: string
        example_value: processingunit
  create:
    description: Manage hits.
    parameters:
      entries:
      - name: reset
        description: If set the hit will reset to 0.
        type: boolean

- rest_name: hookpolicy
  get:
    description: Retrieves the list of hooks.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new hook.

- rest_name: hostservice
  get:
    description: Retrieves the list of host services.
    global_parameters:
    - $filtering
    - $propagatable
    - $propagatable
  create:
    description: Creates a new host service.

- rest_name: hostservicemappingpolicy
  get:
    description: Retrieves the list of host service mappings.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new host service mapping.

- rest_name: httpresourcespec
  get:
    description: Retrieves the list of HTTP resource specifications.
    global_parameters:
    - $filtering
    - $propagatable
    - $archivable
  create:
    description: Creates a new HTTP resource specification.

- rest_name: import
  create:
    description: Imports data from a previous export.

- rest_name: importreference
  get:
    description: Retrieves the list of import references.
    global_parameters:
    - $filtering
  create:
    description: Imports data from a previous export and keep a reference.

- rest_name: importrequest
  get:
    description: Retrieves the list of import requests.
    global_parameters:
    - $filtering
  create:
    description: Creates a new import request.

- rest_name: infrastructurepolicy
  get:
    description: Retrieves the list of infrastructure policies.
    global_parameters:
    - $filtering
  create:
    description: Creates a new infrastructure policy.

- rest_name: installedapp
  get:
    description: Retrieves the list of installed apps.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: tag
        description: List of tags to filter on. This parameter is deprecated.
        type: string
        multiple: true
        example_value: a=a
  create:
    description: Installs a new app.

- rest_name: ipinfo
  get:
    description: Returns information about an IP address given as parameters.
    parameters:
      required:
      - - - ip
      entries:
      - name: ip
        description: List of IPs to resolve.
        type: string
        multiple: true
        example_value: 1.2.3.4

- rest_name: isolationprofile
  get:
    description: Retrieves the list of isolation profiles.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new isolation profile.

- rest_name: issue
  create:
    description: Issues a new token.
    parameters:
      entries:
      - name: asCookie
        description: If set to true, the token will be delivered in a secure cookie, and not in the response body.
        type: boolean

      - name: token
        description: Token to verify.
        type: string
        example_value: xxx.yyyyyyyy.zzzz

- rest_name: issueservicetoken
  create:
    description: Internal API to issue service tokens.

- rest_name: ldapprovider
  get:
    description: Retrieves the list of the namespace LDAP providers.
    global_parameters:
    - $filtering
  create:
    description: Creates a new LDAP provider.

- rest_name: localca
  get:
    description: Returns the local and SSH certificate authorities of the namespace.
  create:
    description: Renews the local and/or SSH certificate authorities of the namespace.

- rest_name: logout
  get:
    description: Performs a logout operation.

- rest_name: message
  get:
    description: Retrieves the list of messages.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new message.

- rest_name: metricsquery
  get:
    description: |-
      Prometheus compatible endpoint to evaluate a Prometheus query expression at a
      single instant or over a range of time.
    parameters:
      required:
      - - - query
      entries:
      - name: query
        description: Prometheus expression query string.
        type: string
        example_value: flows{namespace=~"/mycompany.*"}

      - name: time
        description: Evaluation timestamp <rfc3339 | unix_timestamp>.
        type: string
        example_value: "2015-07-01T20:10:30.781Z"
  create:
    description: |-
      Prometheus compatible endpoint to evaluate a Prometheus query expression at a
      single instant or over a range of time.
      This has the same behavior as the GET request, however it is useful when
      specifying a large query that may breach server-side URL character limits. In
      such a case, you can URL-encode the parameters that would be used for a GET
      request directly in the request body by using the POST method and Content-Type:
      application/x-www-form-urlencoded header.

- rest_name: metricsqueryrange
  get:
    description: |-
      Evaluates an expression query over a range of time returning a "matrix" result
      type.
    parameters:
      required:
      - - - query
      entries:
      - name: end
        description: End timestamp <rfc3339 | unix_timestamp>.
        type: string
        example_value: "2015-07-01T20:11:00.781Z"

      - name: query
        description: Prometheus expression query string.
        type: string
        example_value: flows{namespace=~"/mycompany.*"}

      - name: start
        description: Start timestamp <rfc3339 | unix_timestamp>.
        type: string
        example_value: "2015-07-01T20:10:30.781Z"

      - name: step
        description: Query resolution step width in duration format or float number of seconds.
        type: string
        example_value: 15s
  create:
    description: |-
      Evaluates an expression query over a range of time returning a "matrix" result.
      This has the same behavior as the GET request, however it is useful when
      specifying a large query that may breach server-side URL character limits. In
      such a case, you can URL-encode the parameters that would be used for a GET
      request directly in the request body by using the POST method and Content-Type:
      application/x-www-form-urlencoded header.

- rest_name: namespace
  get:
    description: Retrieves the list of namespaces.
    global_parameters:
    - $filtering
  create:
    description: Creates a new namespace.

- rest_name: namespacedetail
  get:
    description: Returns the information of the specified namespace.

- rest_name: namespacemappingpolicy
  get:
    description: Retrieves the list namespace mappings.
    global_parameters:
    - $filtering
  create:
    description: Creates a new namespace mapping.

- rest_name: namespacepolicyinfo
  get:
    description: (Deprecated) Returns the policy info of the specified namespace.
    deprecated: true

- rest_name: namespacerenderer
  create:
    description: Renders the namespace where an object should reside.

- rest_name: namespacetype
  get:
    description: Returns the type of the specified namespace.

- rest_name: networkaccesspolicy
  get:
    description: Retrieves the list of network policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: |-
      Creates a new network policy. This is deprecated. in favor of
      NetworkRuleSetPolicy.

- rest_name: networkrulesetpolicy
  get:
    description: Retrieves the list of network rule set policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new network rule set policy policy.

- rest_name: oidcprovider
  get:
    description: Retrieves the list of OIDC providers.
    global_parameters:
    - $filtering
  create:
    description: Creates a new OIDC provider.

- rest_name: organizationalmetadata
  get:
    description: |-
      Retrieves the list of organizational metadata for the namespace and its
      namespace hierarchy.

- rest_name: packetreport
  create:
    description: Create a packet trace report.

- rest_name: passwordreset
  get:
    description: Sends a link to the account email to reset the password.
    parameters:
      required:
      - - - email
      entries:
      - name: email
        description: Email associated to the account.
        type: string
        example_value: user@domain.com
  create:
    description: Resets the password for an account using the provided link.

- rest_name: pccprovider
  get:
    description: Retrieves the list of the PCC providers.
    global_parameters:
    - $filtering
  create:
    description: Creates a new PCC provider.

- rest_name: pingrequest
  create:
    description: Initiate a new the ping request.

- rest_name: pingresult
  get:
    description: Retrieves a ping result.
    global_parameters:
    - $filtering

- rest_name: plan
  get:
    description: Retrieves the list of plans.

- rest_name: policy
  get:
    description: Retrieves the list of policy primitives.
    global_parameters:
    - $filtering
    - $propagatable

- rest_name: policygraph
  create:
    description: Retrieve a policy graph.
    parameters:
      entries:
      - name: view
        description: Set the view query for grouping the dependency map.
        type: string
        example_value: $namespace then app

- rest_name: policyrenderer
  create:
    description: Render a policy of a given type for a given set of tags.

- rest_name: pollaccount
  create:
    description: Initiates a poll to a new account.

- rest_name: processingunit
  get:
    description: Retrieves the list of processing units.
    global_parameters:
    - $filtering
    - $archivable
  create:
    description: Creates a new processing unit.

- rest_name: processingunitpolicy
  get:
    description: Retrieves the list of processing unit policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new processing unit policy.

- rest_name: putrafficaction
  get:
    description: Returns the processing unit traffic actions of the specified namespace.
  create:
    description: Modifies the processing unit traffic actions of the specified namespace.

- rest_name: quotacheck
  create:
    description: Verifies if the quota is exceeded for a particular object.
    parameters:
      entries:
      - name: remaining
        description: Makes the system count how many object are left available in the quota.
        type: boolean

- rest_name: quotapolicy
  get:
    description: Retrieves the list of quotas.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new quota.

- rest_name: recipe
  get:
    description: Retrieves the list of recipes.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new recipe.

- rest_name: remoteprocessor
  create:
    description: This should be be here.

- rest_name: renderedpolicy
  create:
    description: Render a policy for a processing unit.
    parameters:
      entries:
      - name: renderer
        description: Select the network policy renderer to use.
        type: enum
        allowed_choices:
        - v1
        - v2

- rest_name: rendertemplate
  create:
    description: Renders a new template.

- rest_name: report
  create:
    description: Create a statistics report.

- rest_name: reportsquery
  create:
    description: Sends a query on report data.
    global_parameters:
    - $filtering
    - $timewindow

- rest_name: revocation
  get:
    description: Verify the revocation of a certificate according to parameters.
    global_parameters:
    - $filtering

- rest_name: role
  get:
    description: Retrieves the list of existing roles.

- rest_name: samlprovider
  get:
    description: Retrieves the list of the namespace SAML providers.
    global_parameters:
    - $filtering
  create:
    description: Creates a new LDAP provider.

- rest_name: sandbox
  create:
    description: Creates a temporary api sandbox.

- rest_name: search
  get:
    description: Perform a full text search on the database.
    parameters:
      required:
      - - - q
      entries:
      - name: q
        description: search query.
        type: string
        example_value: my enforcer

- rest_name: service
  get:
    description: Retrieves the list of services.
    global_parameters:
    - $filtering
    - $archivable
    - $propagatable
  create:
    description: Creates a new service.

- rest_name: servicecertificate
  get:
    description: Retrieves the list of service certificates.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new service certificate.

- rest_name: servicedependencypolicy
  get:
    description: Retrieves the list of service dependency policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new service dependency.

- rest_name: servicepublication
  create:
    description: Creates a service publication for a given service object.

- rest_name: servicetoken
  create:
    description: Creates an OAUTH compatible service token.

- rest_name: squalltag
  get:
    description: Retrieves a computed list of tags from squall for caching.
    parameters:
      required:
      - - - identity
      entries:
      - name: identity
        description: Search for all the tags used for the this identity.
        type: string
        example_value: processingunit

- rest_name: sshauthority
  create:
    description: Creates a new SSH certificate authority.
  delete:
    description: Deletes an existing SSH certificate authority.

- rest_name: sshauthorizationpolicy
  get:
    description: Retrieves the list of SSH authorizations.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new SSH authorizations.

- rest_name: sshcertificate
  create:
    description: Creates a new SSH certificate.

- rest_name: sshidentity
  create:
    description: Creates a new SSH certificate.

- rest_name: statsinfo
  create:
    description: Retrieves information about the content of the stats measurement.

- rest_name: statsquery
  create:
    description: Sends a query on statistical data.
    global_parameters:
    - $timewindow
    - $filtering

- rest_name: suggestedpolicy
  get:
    description: Retrieves a list of network policy suggestions.
    global_parameters:
    - $timewindow
    - $flowoffset
    parameters:
      entries:
      - name: filterAction
        description: Action to take with the filter tags.
        type: enum
        allowed_choices:
        - include
        - exclude

      - name: filterTags
        description: Tags to filter in the policy suggestions.
        type: string
        multiple: true
        example_value: a=a

- rest_name: tag
  get:
    description: Retrieves the list of existing tags in the system.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: onlyPolicyTags
        description: if set to true, only return tags that match the tag prefixes.
        type: boolean

- rest_name: taginject
  create:
    description: Internal API to inject tags.

- rest_name: tagprefix
  get:
    description: Returns the tag prefixes of the specified namespace.
  create:
    description: Modify the tag prefixes of the specified namespace.

- rest_name: tagvalue
  get:
    description: Retrieves the list of existing values for the given tag keys.
    parameters:
      required:
      - - - key
      entries:
      - name: key
        description: Keys of the tag you want to get the values of.
        type: string
        multiple: true
        example_value: $name

- rest_name: tenant
  create:
    description: Creates the tenant's namespace and API authorization policy.

- rest_name: token
  create:
    description: Creates a new token.

- rest_name: tokenscopepolicy
  get:
    description: Retrieves the list of token scope policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new token scope policy.

- rest_name: trustedca
  get:
    description: Retrieves the trusted CAs of a namespace.
    parameters:
      entries:
      - name: type
        description: The type of certificates that it should return.
        type: enum
        allowed_choices:
        - Any
        - X509
        - SSH
        - JWT
        default_value: Any

- rest_name: trustednamespace
  get:
    description: Retrieves the list of trusted namespaces.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new trusted namespace.

- rest_name: useraccesspolicy
  get:
    description: Retrieves the list of user access policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new enforcer policy.

- rest_name: validateuiparameter
  create:
    description: Validates some UI parameters.

- rest_name: vulnerability
  get:
    description: Retrieves the list of vulnerabilities.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new vulnerability.

- rest_name: x509certificate
  create:
    description: Creates a new X.509 certificate.

- rest_name: x509certificatecheck
  get:
    description: Verifies if an X.509 certificate is valid.
    global_parameters:
    - $filtering
