# Model
model:
  rest_name: networklabelrule
  resource_name: networklabelrules
  entity_name: NetworkLabelRule
  package: squall
  group: core/policy
  description: Represents an ingress or egress network labeling rule.
  detached: true

# Attributes
attributes:
  v1:
  - name: secmarkLabel
    description: Defines the SECMARK label to apply to packets.
    type: string
    exposed: true
    required: true
    default_value: ""

  - name: name
    description: A user defined name to keep track of the rule in the reporting.
    type: string
    exposed: true
    max_length: 16
    omit_empty: true

  - name: networks
    description: A list of IP CIDRS or FQDNS that identify remote endpoints.
    type: refList
    exposed: true
    subtype: networkrulenet
    read_only: true
    omit_empty: true
    extensions:
      refMode: pointer

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
    - $atLeastOneSubExpression
    - $subExpressionsNotEmpty
    - $noDuplicateSubExpressions
    - $noDuplicateTagsInEachSubExpression
    - $tagsExpression

  - name: protocolPorts
    description: |-
      Represents the ports and protocols this policy applies to. Protocol/ports are
      defined as tcp/80, udp/22. For protocols that do not have ports, the port
      designation
      is not allowed.
    type: list
    exposed: true
    subtype: string
    validations:
    - $serviceports
