# Model
model:
  rest_name: graphedge
  resource_name: graphedges
  entity_name: GraphEdge
  package: meteor
  group: visualization/depmaps
  description: Represents an edge from the dependency map.
  private: true
  extends:
  - '@zoned'
  extensions:
    forceDocumentation: true

# Indexes
indexes:
- - namespace
- - namespace
  - lastSeen
  - firstSeen
- - lastSeen
  - firstSeen
- - lastSeen
- - firstSeen
- - flowID
- - flowID
  - lastSeen
- - flowID
  - bucketMinute
- - flowID
  - bucketHour
- - flowID
  - bucketDay
- - flowID
  - bucketMonth

# Attributes
attributes:
  v1:
  - name: ID
    description: DB Identifier of the edge.
    type: string
    stored: true
    identifier: true

  - name: acceptedFlows
    description: Number of accepted flows in the edge.
    type: boolean
    exposed: true
    stored: true

  - name: bucketDay
    description: The date for the day bucket.
    type: time
    stored: true

  - name: bucketHour
    description: The date for the hour bucket.
    type: time
    stored: true

  - name: bucketMinute
    description: The date for the minute bucket.
    type: time
    stored: true

  - name: bucketMonth
    description: The date for the month bucket.
    type: time
    stored: true

  - name: defaultAcceptedFlows
    description: The truth value that indicates if there are accepted default flows.
    type: boolean
    exposed: true
    stored: true

  - name: defaultRejectedFlows
    description: The truth value that indicates if there are rejected default flows.
    type: boolean
    exposed: true
    stored: true

  - name: destinationController
    description: Name of the remote destination controller if different than the current one.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: destinationID
    description: ID of the destination `GraphNode` of the edge.
    type: string
    exposed: true
    stored: true

  - name: destinationType
    description: Type of the destination `GraphNode` of the edge.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Namespace
    - Node
    - RemoteController

  - name: details
    description: Contains more flow details grouped by their destination protocol/ports.
    type: external
    exposed: true
    subtype: map[string]graphedgeflowdetails
    omit_empty: true

  - name: encrypted
    description: The number of encrypted flows in the edge.
    type: boolean
    exposed: true
    stored: true

  - name: firstSeen
    description: Contains the date when the edge was first seen.
    type: time
    exposed: true
    stored: true
    omit_empty: true

  - name: flowID
    exposed_name: ID
    description: Identifier of the edge.
    type: string
    exposed: true
    stored: true

  - name: lastSeen
    description: Contains the date when the edge was last seen.
    type: time
    exposed: true
    stored: true
    omit_empty: true

  - name: namespace
    description: Namespace of the object that reported the flow.
    type: string
    exposed: true
    stored: true

  - name: observedAcceptedFlows
    description: Number of accepted observed flows.
    type: boolean
    exposed: true
    stored: true

  - name: observedEncrypted
    description: Number of encrypted observed flows.
    type: boolean
    exposed: true
    stored: true

  - name: observedRejectedFlows
    description: Number of rejected observed flows.
    type: boolean
    exposed: true
    stored: true

  - name: rejectedFlows
    description: Number of rejected flows in the edge.
    type: boolean
    exposed: true
    stored: true

  - name: remoteNamespace
    description: Namespace of the object that was targeted by the flow.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: sourceController
    description: Name of the remote source controller if different than the current one.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: sourceID
    description: ID of the source `GraphNode` of the edge.
    type: string
    exposed: true
    stored: true

  - name: sourceType
    description: Type of the source `GraphNode` of the edge.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Namespace
    - Node
    - RemoteController
