# Model
model:
  rest_name: namespaceinfo
  resource_name: namespaceinfo
  entity_name: NamespaceInfo
  package: squall
  group: core/namespace
  description: Returns the information of the specified namespace.

# Ordering
default_order:
- name

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

  - name: description
    description: Description of the namespace.
    type: string
    exposed: true
    read_only: true
    orderable: true

  - name: name
    description: Name of the namespace.
    type: string
    exposed: true
    read_only: true
    filterable: true
    orderable: true

  - name: prefixes
    description: List of tag prefixes that will be used to suggest policies.
    type: list
    exposed: true
    subtype: string
    read_only: true

  - name: protected
    description: Defines if the namespace is protected.
    type: boolean
    exposed: true
    read_only: true
    orderable: true
