# Model
model:
  rest_name: policyrule
  resource_name: policyrules
  entity_name: PolicyRule
  package: squall
  group: core/policy
  description: Allows services to retrieve a policy resolution (internal).
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@identifiable-not-stored'
  - '@named'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: action
    description: Defines set of actions that must be enforced when a dependency is met.
    type: external
    exposed: true
    subtype: map[string]map[string]interface{}

  - name: auditProfiles
    description: Provides the audit profiles that must be applied.
    type: refList
    exposed: true
    subtype: auditprofile
    deprecated: true
    omit_empty: true
    extensions:
      noInit: true

  - name: enforcerProfiles
    description: Provides information about the enforcer profile.
    type: refList
    exposed: true
    subtype: enforcerprofile
    omit_empty: true
    extensions:
      noInit: true

  - name: externalNetworks
    description: Provides the external network that the policy targets.
    type: refList
    exposed: true
    subtype: externalnetwork
    omit_empty: true
    extensions:
      noInit: true

  - name: filePaths
    description: Provides the file paths that the policy targets.
    type: refList
    exposed: true
    subtype: filepath
    deprecated: true
    omit_empty: true
    extensions:
      noInit: true

  - name: hostServices
    description: Provides the list of host services that must be instantiated.
    type: refList
    exposed: true
    subtype: hostservice
    omit_empty: true
    extensions:
      noInit: true

  - name: isolationProfiles
    description: Provides the isolation profiles of the rule.
    type: refList
    exposed: true
    subtype: isolationprofile
    deprecated: true
    omit_empty: true
    extensions:
      noInit: true

  - name: namespaces
    description: The namespace that the policy targets.
    type: refList
    exposed: true
    subtype: namespace
    omit_empty: true
    extensions:
      noInit: true

  - name: policyNamespace
    description: The namespace of the policy that created this rule.
    type: string
    exposed: true

  - name: policyUpdateTime
    description: Last time the policy was updated.
    type: time
    exposed: true

  - name: propagated
    description: Indicates if the policy is propagated.
    type: boolean
    exposed: true

  - name: relation
    description: Describes the required operation to be performed between subjects and objects.
    type: list
    exposed: true
    subtype: string

  - name: services
    description: Provides the services of this policy rule.
    type: refList
    exposed: true
    subtype: service
    omit_empty: true
    extensions:
      noInit: true

  - name: tagClauses
    description: Policy target tags.
    type: external
    exposed: true
    subtype: '[][]string'
    omit_empty: true
    validations:
    - $tagsExpression
