# Model
model:
  rest_name: cloudpathelement
  resource_name: cloudpathelements
  entity_name: CloudPathElement
  package: yeul
  group: pcn/infrastructure
  description: |-
    Represents a node in a network path from source to destination. Captures the
    node and related policy information.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: nativeID
    description: The native ID of the node.
    type: string
    exposed: true
    read_only: true

  - name: policy
    description: The policy ID used at this node of the path.
    type: ref
    exposed: true
    subtype: cloudgraphnodeaction
    read_only: true
    omit_empty: true
    extensions:
      refMode: pointer

  - name: routeTables
    description: The route table ID used for the route calculation.
    type: external
    exposed: true
    subtype: map[string]string
    read_only: true
    omit_empty: true
