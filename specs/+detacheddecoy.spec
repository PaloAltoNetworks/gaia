# Model
model:
  rest_name: detacheddecoy
  resource_name: detacheddecoys
  entity_name: DetachedDecoy
  package: yeul
  group: pcn/infrastructure
  description: |-
    Represents a Decoy, partially, suitable for over-the-wire transmission. WARNING:
    this will eventually go away as we should transmit the tree. We keep it this way
    for backwards compatibility with existing code for the sake of speed.
  detached: true

# Attributes
attributes:
  v1:
  - name: hasErrNotPermitted
    description: Indicates whether the decoy experienced ErrNotPermitted.
    type: boolean
    exposed: true

  - name: nodeUID
    description: The NodeUID which the decoy passed through.
    type: ref
    exposed: true
    subtype: nodeuid
    extensions:
      refMode: pointer
