# Model
model:
  rest_name: cloudscalegroupdata
  resource_name: cloudscalegroupdata
  entity_name: CloudScaleGroupData
  package: yeul
  group: pcn/infrastructure
  description: Manages the properties associated to a Auto Scaling Group.
  detached: true

# Attributes
attributes:
  v1:
  - name: availabilityZones
    description: Availability Zones for the scale group.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - us-east-2b
    - us-east-2c

  - name: instances
    description: ID of associated instances with this scale group.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - i-abcd1234
    - i-abcd1235

  - name: vpcZoneIdentifiers
    description: One or more subnet IDs, if applicable, separated by commas.
    type: list
    exposed: true
    subtype: string
    stored: true
