# Model
model:
  rest_name: graphedgeflowdetails
  resource_name: graphedgeflowdetails
  entity_name: GraphEdgeFlowDetails
  package: meteor
  group: visualization/depmaps
  description: Contains more specific information about a particular flow between two nodes.
  detached: true

# Attributes
attributes:
  v1:
  - name: accepted
    description: Indicates whether the flow was accepted.
    type: boolean
    exposed: true

  - name: protoPort
    description: The protocol for this edge. If it is tcp or udp, the port is also included.
    type: string
    exposed: true
    example_value: tcp/443
