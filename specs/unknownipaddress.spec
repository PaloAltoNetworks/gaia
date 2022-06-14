# Model
model:
  rest_name: unknownipaddress
  resource_name: unknownipaddresses
  entity_name: UnknownIPAddress
  package: karl
  group: pc/alerts
  description: Holds the IP/FQDN of flows going to somewhere (default) external network.
  extends:
  - '@base'
  - '@identifiable-stored'
  - '@namespaced'
  - '@timeable'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - address
- - namespace
  - createTime

# Attributes
attributes:
  v1:
  - name: address
    description: IP/FQDN encountered.
    type: string
    exposed: true
    stored: true

  - name: occurrences
    description: Number of times the address was seen for the source.
    type: integer
    exposed: true
    stored: true

  - name: sourceID
    description: Identifier of the source of the address.
    type: string
    exposed: true
    stored: true

  - name: sourceIdentity
    description: Identity of the source of the address.
    type: string
    exposed: true
    stored: true

  - name: sourceNamespace
    description: Namespace of the source of the address.
    type: string
    exposed: true
    stored: true

  - name: timestamp
    description: The timestamp of the source of the address.
    type: time
    exposed: true
    stored: true
