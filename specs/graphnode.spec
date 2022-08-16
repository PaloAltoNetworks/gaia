# Model
model:
  rest_name: graphnode
  resource_name: graphnodes
  entity_name: GraphNode
  package: meteor
  group: visualization/depmaps
  description: Represents an node from the dependency map.
  private: true
  extensions:
    forceDocumentation: true

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of object represented by the node.
    type: string
    exposed: true
    stored: true

  - name: enforcementStatus
    description: Enforcement status of processing unit represented by the node.
    type: string
    exposed: true
    stored: true

  - name: firstSeen
    description: Contains the date when the edge was first seen.
    type: time
    exposed: true
    stored: true

  - name: groupID
    description: ID of the group the node is eventually part of.
    type: string
    exposed: true
    stored: true

  - name: images
    description: List of images.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: lastSeen
    description: Contains the date when the edge was last seen.
    type: time
    exposed: true
    stored: true

  - name: name
    description: Name of object represented by the node.
    type: string
    exposed: true
    stored: true

  - name: namespace
    description: Namespace of object represented by the node.
    type: string
    exposed: true
    stored: true

  - name: status
    description: Status of object represented by the node.
    type: string
    exposed: true
    stored: true

  - name: tags
    description: Tags of object represented by the node.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: type
    description: Type of object represented by the node.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Docker
    - ExternalNetwork
    - Volume
    - Claim
    - Node
    - Namespace
    - RemoteController
    - APIGateway
    - Host
    - HostService
    - LinuxService
    - WindowsService
    - RKT
    - User
    - SSHSession
    - ECS

  - name: unreachable
    description: If `true` the node is marked as unreachable.
    type: boolean
    exposed: true
    stored: true

  - name: vulnerabilityLevel
    description: Tags of object represented by the node.
    type: string
    exposed: true
    stored: true
