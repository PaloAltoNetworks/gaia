# Model
model:
  rest_name: cloudloadbalancerdata
  resource_name: cloudloadbalancerdata
  entity_name: CloudLoadBalancerData
  package: yeul
  group: pcn/infrastructure
  description: Parameters associated with a cloud load balancer.
  detached: true

# Attributes
attributes:
  v1:
  - name: name
    description: The name of the load balancer.
    type: string
    exposed: true
    stored: true
    example_value: lb-009251c49cf46d940

  - name: scheme
    description: The scheme tells whether the load balancer is internet facing or
      internal.
    type: string
    exposed: true
    stored: true
    example_value: internet-facing

  - name: targetGrouplist
    description: Target groups associated with this load balancer.
    type: refList
    exposed: true
    subtype: cloudtargetgroup
    stored: true
    extensions:
      refMode: pointer
