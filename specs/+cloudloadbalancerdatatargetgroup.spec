# Model
model:
  rest_name: cloudloadbalancerdatatargetgroup
  resource_name: cloudloadbalancerdatatargetgroup
  entity_name: CloudLoadBalancerDataTargetGroup
  package: yeul
  group: pcn/infrastructure
  description: Represents a load balancing rule target group.
  detached: true

# Attributes
attributes:
  v1:
  - name: DirectServerReturnEnabled
    description: Defines if Direct Server Return is enabled for the group.
    type: boolean
    exposed: true
    stored: true

  - name: TargetIDs
    description: The list of IDs of the target.
    type: list
    exposed: true
    subtype: string
    stored: true
