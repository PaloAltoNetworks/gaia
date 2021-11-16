# Model
model:
  rest_name: cloudtargetgroup
  resource_name: cloudtargetgroups
  entity_name: CloudTargetGroup
  package: yeul
  group: pcn/infrastructure
  description: Describes a target group in a load balancer.
  detached: true

# Attributes
attributes:
  v1:
  - name: Port
    description: The port for the target group.
    type: string
    exposed: true

  - name: TargetGroupID
    description: The ID for the target group.
    type: string
    exposed: true
    stored: true

  - name: TargetType
    description: The type of the next hop object.
    type: string
    exposed: true
    stored: true
    example_value: i-1234

  - name: loadBalancerName
    description: The name of the load balancer.
    type: string
    exposed: true
    stored: true
    example_value: lb-1234
