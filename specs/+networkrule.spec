# Model
model:
  rest_name: networkrule
  resource_name: networkrules
  entity_name: NetworkRule
  package: squall
  group: core/policy
  description: Represents an ingress or egress network rule.
  detached: true
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

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

  - name: appidRestrictions
    description: appid restrictions.
    type: string
    exposed: true
    subtype: string

  - name: logsDisabled
    description: |-
      If `true`, the relevant flows will not be reported to the Microsegmentation
      Console.
      Under some advanced scenarios you may wish to set this to `true`, such as to
      save space or improve performance.
    type: boolean
    exposed: true

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

  - name: observationEnabled
    description: If set to `true`, the flow will be in observation mode.
    type: boolean
    exposed: true
    default_value: false

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
