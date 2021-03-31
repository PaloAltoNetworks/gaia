# Model
model:
  rest_name: cloudnetworkrule
  resource_name: cloudnetworkrules
  entity_name: CloudNetworkRule
  package: yeul
  group: pcn/infrastructure
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
    stored: true
    required: true
    allowed_choices:
    - Allow
    - Reject
    default_value: Allow

  - name: networks
    description: A list of IP CIDRS that identify remote endpoints.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
    omit_empty: true
    validations:
    - $optionalcidroriplist

  - name: object
    description: |-
      Identifies the set of remote workloads that the rule relates to. The selector
      will identify both processing units as well as external networks that match the
      selector.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    orderable: true

  - name: priority
    description: Priority of the rule. Available only for cloud ACLs.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: protocolPorts
    description: |-
      Represents the ports and protocols this policy applies to. Protocol/ports are
      defined as tcp/80, udp/22. For protocols that do not have ports, the port
      designation
      is not allowed.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $serviceports
