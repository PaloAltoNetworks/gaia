# Model
model:
  rest_name: namespacepolicyinfo
  resource_name: namespacepolicyinfo
  entity_name: NamespacePolicyInfo
  package: squall
  group: core/namespace
  description: Returns the policy info of the specified namespace.

# Attributes
attributes:
  v1:
  - name: PUIncomingTrafficAction
    description: The processing unit action for incoming traffic for the namespace.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit

  - name: PUOutgoingTrafficAction
    description: The processing unit action for outgoing traffic for the namespace.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit

  - name: prefixes
    description: List of tag prefixes that will be used to suggest policies.
    type: list
    exposed: true
    subtype: string
    read_only: true
