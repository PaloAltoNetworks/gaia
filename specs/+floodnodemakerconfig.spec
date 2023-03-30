# Model
model:
  rest_name: floodnodemakerconfig
  resource_name: floodnodemakerconfigs
  entity_name: FloodNodeMakerConfig
  package: yeul
  group: pcn/infrastructure
  description: |-
    Configurations to adjust the behaviour of nodemakers registered within a cached
    mux.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: cloudTypes
    description: |-
      A list of cloud types involved in flooding.  WARNING: this will eventually go
      away as we should transmit the tree. We keep it this way for backwards
      compatibility with existing code for the sake of speed.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value: gcp

  - name: optionIgnoreIPRulesForGivenAddresses
    description: A list of addresses which nodemakers will ignore when evaluating
      IP rules.
    type: list
    exposed: true
    subtype: string

  - name: optionSelectIPRulesWithFullCoverage
    description: |-
      If set, IP rule evaluation will only be considered if the target address is
      fully covered by the IP rule.
    type: boolean
    exposed: true
