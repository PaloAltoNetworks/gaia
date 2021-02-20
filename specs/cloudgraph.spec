# Model
model:
  rest_name: cloudgraph
  resource_name: cloudgraphs
  entity_name: CloudGraph
  package: pcn/infrastructure
  group: pcn/infrastructure
  description: "Returns a data structure representing the graph of all cloud nodes
    \nand their connections in a particular namespace."

# Attributes
attributes:
  v1:
  - name: internalEdges
    description: The edges of the map.
    type: refMap
    exposed: true
    subtype: cloudgraphedge
    read_only: true
    extensions:
      refMode: pointer

  - name: nodes
    description: Refers to the nodes of the map.
    type: refMap
    exposed: true
    subtype: cloudgraphnode
    read_only: true
    extensions:
      refMode: pointer

  - name: publicEdges
    description: The edges of the map.
    type: refMap
    exposed: true
    subtype: cloudgraphedge
    read_only: true
    extensions:
      refMode: pointer

  - name: query
    description: |-
      The cloud network query that should be used. This requires a POST operation on
      the object.
    type: ref
    exposed: true
    subtype: cloudnetworkquery
    extensions:
      refMode: pointer