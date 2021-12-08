# Model
model:
  rest_name: connectionsresponse
  resource_name: connectionsresponses
  entity_name: ConnectionsResponse
  package: guy
  group: core/enforcer
  description: A response from a connectionsrequest.
  extends:
  - '@zoned'
  - '@migratable'
  - '@timeable'
  - '@namespaced'
  - '@identifiable-stored'

# Indexes
indexes:
- - requestID
- - namespace
  - requestID


# Attributes
attributes:
  v1:
  - name: requestID
    description: Unique ID generated for each request.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: refreshID
    description: Contains the refresh ID set by processing unit refresh event.
    type: string
    exposed: true
    required: true
    example_value: xxxx-xxxx-xxxx

  - name: totalConnections
    description: Contains the total number of connections for the connection request.
    type: integer
    exposed: true
    omit_empty: true

  - name: connections
    description: Contains a batch of connections.
    type: refList
    exposed: true
    subtype: currentconnection
    stored: true   