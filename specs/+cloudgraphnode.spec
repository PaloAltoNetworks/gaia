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
  - name: childrenIDs
    description: The list of children for this node.
    type: external
    exposed: true
    subtype: map[string]map[string][]string
    omit_empty: true

  - name: nodeData
    description: Details about the node if the query type requests full details.
    type: ref
    exposed: true
    subtype: cloudnode
    omit_empty: true
    extensions:
      refMode: pointer

  - name: policies
    description: The policies that were applied to this node for each destination.
    type: refMap
    exposed: true
    subtype: cloudgraphnodeaction
    omit_empty: true
    extensions:
      refMode: pointer

  - name: publicChildrenIDs
    description: The list of public children for this node.
    type: external
    exposed: true
    subtype: map[string]map[string][]string
    omit_empty: true

  - name: publicPolicies
    description: |-
      The policies that were applied to this node for each destination based on public
      IPs.
    type: refMap
    exposed: true
    subtype: cloudgraphnodeaction
    omit_empty: true
    extensions:
      refMode: pointer

  - name: routeTableIDs
    description: |-
      The list of route tables IDs that forwarding was based on for the internal path,
      if routing was
      performed.
    type: external
    exposed: true
    subtype: map[string]string
    omit_empty: true

  - name: type
    description: The type of the node as a string.
    type: string
    exposed: true
    omit_empty: true
