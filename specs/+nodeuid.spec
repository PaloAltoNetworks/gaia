# Model
model:
  rest_name: nodeuid
  resource_name: nodeuids
  entity_name: NodeUID
  package: yeul
  group: pcn/infrastructure
  description: Represents a neocna NodeUID.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: The node identifier.
    type: string
    exposed: true
    omit_empty: true

  - name: networkAddress
    description: The network address of the node described by the NodeUID.
    type: string
    exposed: true
    omit_empty: true
