# Model
model:
  rest_name: cloudvpcdata
  resource_name: cloudvpcdata
  entity_name: CloudVPCData
  package: yeul
  group: pcn/infrastructure
  description: Managed the list of IP addresses associated with an interface.
  detached: true

# Attributes
attributes:
  v1:
  - name: addresses
    description: Address CIDRs of the VPC.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - 10.8.0.0/16
    - 172.24.0.0/16
    validations:
    - $cidrs
