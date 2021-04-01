# Model
model:
  rest_name: cloudroute
  resource_name: cloudroutes
  entity_name: CloudRoute
  package: yeul
  group: pcn/infrastructure
  description: Describes a route in a route table.
  detached: true

# Attributes
attributes:
  v1:
  - name: destinationIPv4CIDR
    description: The Destination CIDR for the route.
    type: string
    exposed: true
    stored: true
    example_value: 10.1.1.32/24
    validations:
    - $optionalcidrorip

  - name: destinationIPv6CIDR
    description: The destination IPV6 CIDR for the route.
    type: string
    exposed: true
    stored: true
    example_value: 2001:db8::/32
    validations:
    - $optionalcidrorip

  - name: nextHopID
    description: The ID of the next hop object.
    type: string
    exposed: true
    stored: true
    example_value: gw_123444444

  - name: nextHopType
    description: The type of the next hop.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - EgressOnlyGateway
    - Gateway
    - Instance
    - LocalGateway
    - NATGateway
    - NetworkInterface
    - TransitGateway
    - VPCPeeringConnection
    - TransitGatewayAttachment
    example_value: LocalGateway