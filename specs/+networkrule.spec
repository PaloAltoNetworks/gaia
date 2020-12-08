# Model
model:
  rest_name: networkrule
  resource_name: networkrules
  entity_name: NetworkRule
  package: squall
  group: core/policy
  description: Represents an ingress or egress network rule.
  detached: true

# Attributes
attributes:
  v1:
  - name: action
    description: |-
      Defines the action to apply to a flow.
      - `Allow`: allows the defined traffic.
      - `Reject`: rejects the defined traffic; useful in conjunction with an allow all
      policy.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Allow
    - Reject
    default_value: Allow
    orderable: true

  - name: networks
    description: A list of IP CIDRS or FQDNS that identify remote endpoints.
    type: list
    exposed: true
    subtype: string
    orderable: true
    omit_empty: true

  - name: object
    description: |-
      Identifies the set of remote workloads that the rule relates to. The selector
      will identify both processing units as well as external networks that match the
      selector.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: observationEnabled
    description: If set to `true`, the flow will be in observation mode.
    type: boolean
    exposed: true
    default_value: false
    orderable: true

  - name: protocolPorts
    description: |-
      Represents the ports and protocols this policy applies to. Protocol/ports are
      defined as tcp/80, udp/22. For protocols that do not have ports, the port
      designation
      is not allowed.
    type: list
    exposed: true
    subtype: string
    orderable: true
    validations:
    - $serviceports
