# Model
model:
  rest_name: policy
  resource_name: policies
  entity_name: Policy
  package: squall
  group: core/policy
  description: Represents the policy primitive used by all Microsegmentation policies.
  get:
    description: Retrieves the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@hidden'
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - namespace
  - type
- - namespace
  - type
  - allObjectTags
- - namespace
  - type
  - allSubjectTags
- - namespace
  - type
  - allObjectTags
  - disabled
- - namespace
  - type
  - allSubjectTags
  - disabled
- - namespace
  - type
  - allObjectTags
  - propagate
- - namespace
  - type
  - allSubjectTags
  - propagate

# Attributes
attributes:
  v1:
  - name: action
    description: Defines a set of actions that must be enforced when a dependency is met.
    type: external
    exposed: true
    subtype: map[string]map[string]interface{}
    stored: true

  - name: allObjectTags
    description: This is a set of all object tags for matching in the DB.
    type: list
    subtype: string
    stored: true

  - name: allSubjectTags
    description: This is a set of all subject tags for matching in the DB.
    type: list
    subtype: string
    stored: true

  - name: expirationTime
    description: If set the policy will be automatically deleted at the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: object
    description: |-
      Represents set of entities that another entity depends on. As subjects,
      objects are identified as logical operations on tags when a policy is defined.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    getter: true
    setter: true
    validations:
    - $tagsExpression

  - name: relation
    description: |-
      Describes the required operation to be performed between subjects and
      objects.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: subject
    description: |-
      Represents sets of entities that will have a dependency other entities.
      Subjects are defined as logical operations on tags. Logical operations can
      include `AND` and `OR`.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    getter: true
    setter: true
    validations:
    - $tagsExpression

  - name: type
    description: Type of the policy.
    type: enum
    exposed: true
    stored: true
    creation_only: true
    allowed_choices:
    - APIAuthorization
    - AuditProfileMapping
    - EnforcerProfile
    - File
    - Hook
    - HostServiceMapping
    - Infrastructure
    - NamespaceMapping
    - Network
    - NetworkRuleSet
    - ProcessingUnit
    - Quota
    - Service
    - ServiceDependency
    - Syscall
    - TokenScope
    - SSHAuthorization
    - UserAccess
    filterable: true
