# Model
model:
  rest_name: processingunitpolicy
  resource_name: processingunitpolicies
  entity_name: ProcessingUnitPolicy
  package: squall
  group: policy/processingunits
  description: |-
    Processing unit policies allow you to define special behavior for
    processing units. For example you can associate an isolation profile
    with a set of processing units or select a specific datapath.
  aliases:
  - pup
  - pups
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
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
  validations:
  - $processingunitpolicy

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: action
    description: |-
      Action determines the action to take while enforcing the isolation profile.
      NOTE: Choose `Default` if your processing unit is not supposed to make a
      decision on isolation profiles at all.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Default
    - Delete
    - Enforce
    - LogCompliance
    - Reject
    - Snapshot
    - Stop
    default_value: Default
    orderable: true

  - name: datapathType
    description: |-
      The datapath type that processing units selected by `subject` should implement:
      - `Default`: This policy is not making a decision for the datapath
      - `Aporeto`: The enforcer is managing and handling the datapath
      - `EnvoyAuthorizer`: The enforcer is serving Envoy-compatible gRPC APIs for
      every processing unit that for example can be used by an Envoy proxy to use the
      Microsegmentation PKI and implement Microsegmentation network policies. NOTE:
      The enforcer is not going to own the datapath in this example. It is merely
      providing an authorizer API.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Default
    - Aporeto
    - EnvoyAuthorizer
    default_value: Default
    filterable: true

  - name: isolationProfileSelector
    description: |-
      The isolation profiles to be mapped. Only applies to `Enforce` and
      `LogCompliance` actions.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      Contains the tag expression the tags need to match for the policy to
      apply.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: isolationprofile
  get:
    description: Returns the list of isolation profiles associated with the mapping.

- rest_name: processingunit
  get:
    description: Returns the list of processing units referenced by the mapping.
