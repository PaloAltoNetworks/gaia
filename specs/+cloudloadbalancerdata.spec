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
  - name: listenertargetmapping
    description: Mapping of a listener to its associated target group ID list.
    type: external
    exposed: true
    subtype: map[string][]string
    stored: true

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
