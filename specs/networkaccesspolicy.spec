# Model
model:
  rest_name: networkaccesspolicy
  resource_name: networkaccesspolicies
  entity_name: NetworkAccessPolicy
  package: squall
  description: |-
    Allows to define networking policies to allow or prevent processing units
    identitied by their tags to talk to other processing units or external services
    (also identified by their tags).
  aliases:
  - netpol
  - netpols
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@negatable-object'
  - '@negatable-subject'

# Attributes
attributes:
  v1:
  - name: action
    description: Action defines the action to apply to a flow.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    - Continue
    default_value: Allow
    filterable: true
    orderable: true

  - name: destinationPorts
    description: DestinationPorts contains the list of allowed ports and ranges.
    type: external
    exposed: true
    subtype: ports_list
    filterable: true
    orderable: true

  - name: encryptionEnabled
    description: EncryptionEnabled defines if the flow has to be encrypted.
    type: boolean
    exposed: true
    filterable: true
    orderable: true

  - name: logsEnabled
    description: LogsEnabled defines if the flow has to be logged.
    type: boolean
    exposed: true
    filterable: true
    orderable: true

  - name: object
    description: Object of the policy.
    type: external
    exposed: true
    subtype: policies_list
    orderable: true

  - name: observationEnabled
    description: If set to true, the flow will be in observation mode.
    type: boolean
    exposed: true
    filterable: true
    orderable: true

  - name: observedTrafficAction
    description: |-
      If observationEnabled is set to true, this will be the final action taken on the
      packets.
    type: enum
    exposed: true
    allowed_choices:
    - Apply
    - Continue
    default_value: Continue
    filterable: true
    orderable: true

  - name: subject
    description: Subject of the policy.
    type: external
    exposed: true
    subtype: policies_list
    orderable: true

# Relations
relations:
- rest_name: externalservice
  get:
    description: Returns the list of external services affected by a network access
      policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subjects
        - object
        default_value: objects

- rest_name: processingunit
  get:
    description: Returns the list of Processing Units affected by a network access
      policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subjects
        - object
        default_value: objects
