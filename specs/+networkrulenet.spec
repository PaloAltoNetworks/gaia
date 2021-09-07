# Model
model:
  rest_name: networkrulenet
  resource_name: networkrulenets
  entity_name: NetworkRuleNet
  package: squall
  group: core/policy
  description: Represents an network contained in a NetworkRule.
  detached: true
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Attributes
attributes:
  v1:
  - name: ID
    description: The ID of the external network.
    type: string
    exposed: true
    read_only: true
    omit_empty: true

  - name: entries
    description: List of CIDRs or domain name.
    type: list
    exposed: true
    subtype: string
    read_only: true
    omit_empty: true

  - name: namespace
    description: The namespace of the external network.
    type: string
    exposed: true
    read_only: true
    omit_empty: true
