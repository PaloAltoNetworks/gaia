# Model
model:
  rest_name: cloudloadbalancerroutedata
  resource_name: cloudloadbalancerroutedata
  entity_name: CloudLoadBalancerRouteData
  package: yeul
  group: pcn/infrastructure
  description: Parameters associated with a cloud load balancer route.
  detached: true

# Attributes
attributes:
  v1:
  - name: port
    description: The port for the target group.
    type: string
    exposed: true

  - name: targetGroupID
    description: The ID for the target group.
    type: string
    exposed: true
    stored: true

  - name: targetID
    description: The ID of the target object.
    type: string
    exposed: true
    stored: true
    example_value: i-d867b126f574d
