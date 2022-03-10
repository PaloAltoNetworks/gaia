# Model
model:
  rest_name: discoverymode
  resource_name: discoverymode
  entity_name: DiscoveryMode
  package: yuna
  group: core
  description: |-
    (Deprecated) When discovery mode is enabled, all flows are accepted. Flows which
    do not match an existing network policy will be represented by a dotted line in
    your Platform view.
  get:
    description: (Deprecated) Retrieve the discovery mode with the given import reference
      ID.
    deprecated: true
  delete:
    description: (Deprecated) Remove the discovery mode assets with the given import
      reference ID.
    deprecated: true
  extends:
  - '@identifiable-not-stored'
  - '@propagated'
  - '@namespaced'
