# Model
model:
  rest_name: cloudsubnetdata
  resource_name: cloudsubnetdata
  entity_name: CloudSubnetData
  package: yeul
  group: prisma/infrastructure
  description: Parameters associated with a subnet.
  detached: true

# Attributes
attributes:
  v1:
  - name: address
    description: Address CIDR of the Subnet.
    type: string
    exposed: true
    stored: true
    example_value: 10.0.0.0/8
    validations:
    - $cidr

  - name: zoneID
    description: The availability zone ID of the subnet.
    type: string
    exposed: true
    stored: true
    example_value: aws-east

  - name: zoneName
    description: The availability zone of the subnet.
    type: string
    exposed: true
    stored: true
    example_value: aws-east
