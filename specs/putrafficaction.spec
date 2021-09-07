# Model
model:
  rest_name: putrafficaction
  resource_name: putrafficactions
  entity_name: PUTrafficAction
  package: squall
  group: core/namespace
  description: Returns the processing unit traffic actions for the specified namespace.
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Attributes
attributes:
  v1:
  - name: Incoming
    description: The processing unit action for incoming traffic for the namespace.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit

  - name: Outgoing
    description: The processing unit action for outgoing traffic for the namespace.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit
