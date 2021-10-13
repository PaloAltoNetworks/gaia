# Model
model:
  rest_name: infrastructurepolicy
  resource_name: infrastructurepolicies
  entity_name: InfrastructurePolicy
  package: squall
  group: policy/networking
  description: |-
    Infrastructure policies represent the network access rules of the
    underlying infrastructure. They can assist you in analyzing how AWS security
    groups,firewalls, and other access control list (ACL) mechanisms may affect
    Microsegmentation network policies. Microsegmentation's AWS integration app
    automatically populates AWS security groups.
  aliases:
  - infrapol
  - infrapols
  get:
    description: Retrieves the infrastructure policy with the given ID.
  update:
    description: Updates the infrastructure policy with the given ID.
  delete:
    description: Deletes the infrastructure policy with the given ID.
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
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: action
    description: Defines the action to apply to a flow.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    default_value: Allow
    orderable: true

  - name: applyPolicyMode
    description: |-
      Determines if the policy applies to the outgoing traffic of the `subject` or the
      incoming traffic of the `subject`. `OutgoingTraffic` (default) or
      `IncomingTraffic`.
    type: enum
    exposed: true
    allowed_choices:
    - OutgoingTraffic
    - IncomingTraffic
    default_value: OutgoingTraffic
    orderable: true

  - name: expirationTime
    description: If set the policy will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: object
    description: Object of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: Subject of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: externalnetwork
  get:
    description: Returns the list of external networks affected by an infrastructure policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object

- rest_name: processingunit
  get:
    description: Returns the list of processing units affected by an infrastructure policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object

- rest_name: service
  get:
    description: Returns the list of services affected by an infrastructure policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object
