# Model
model:
  rest_name: cloudgraphnode
  resource_name: cloudgraphnodes
  entity_name: CloudGraphNode
  package: yeul
  group: pcn/infrastructure
  description: |-
    Returns a data structure representing the graph of all cloud nodes and their
    connections in a particular namespace.
  detached: true

# Attributes
attributes:
  v1:
  - name: nodeData
    description: Details about the node if the query type requests full details.
    type: ref
    exposed: true
    subtype: cloudnode
    omit_empty: true
    extensions:
      refMode: pointer

  - name: type
    description: The type of the node as a string.
    type: string
    exposed: true
    omit_empty: true
