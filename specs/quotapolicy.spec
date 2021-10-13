# Model
model:
  rest_name: quotapolicy
  resource_name: quotapolicies
  entity_name: QuotaPolicy
  package: squall
  group: policy/quota
  description: |-
    Allows you to set quotas on the number of objects that can be
    created in a namespace.
  aliases:
  - quota
  - quotas
  - quotapol
  - quotapols
  get:
    description: Retrieves the quota with the given ID.
  update:
    description: Updates the quota with the given ID.
  delete:
    description: Deletes the quota with the given ID.
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
  - '@hidden'
  - '@fallback'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: expirationTime
    description: If set the quota will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: identities
    description: Contains the list of identity names where the quota will be applied.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - processingunit
    - enforcer

  - name: quota
    description: |-
      Specifies the maximum number of objects matching the policy subject that can be
      created.
    type: integer
    exposed: true

  - name: targetNamespace
    description: Contains the base namespace from where the count will be done.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
