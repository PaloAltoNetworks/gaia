# Model
model:
  rest_name: auditprofilemappingpolicy
  resource_name: auditprofilemappingpolicies
  entity_name: AuditProfileMappingPolicy
  package: squall
  group: policy/audit
  description: |-
    Use an audit profile mapping to define the set of enforcers that must implement
    a specific audit profile.
  aliases:
  - audpol
  - audpols
  get:
    description: Retrieves the mapping with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the mapping with the given ID.
  delete:
    description: Deletes the mapping with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: object
    description: The tag or tag expression that identifies the audit profile to be
      mapped.
    type: external
    exposed: true
    subtype: '[][]string'
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      The tag or tag expression that identifies the enforcer(s) to implement the audit
      profile.
    type: external
    exposed: true
    subtype: '[][]string'
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: auditprofile
  get:
    description: Returns the list of audit profiles that are referred to by this mapping.

- rest_name: enforcer
  get:
    description: Returns the list of enforcers that are affected by this mapping.
