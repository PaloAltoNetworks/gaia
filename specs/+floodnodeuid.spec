# Model
model:
  rest_name: floodnodeuid
  resource_name: floodnodeuids
  entity_name: FloodNodeUID
  package: yeul
  group: pcn/infrastructure
  description: Represents a neocna NodeUID.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: nID
    description: The node identifier.
    type: string
    exposed: true

  - name: networkAddress
    description: The network address of the node described by the NodeUID.
    type: string
    exposed: true
    omit_empty: true
