# Model
model:
  rest_name: cloudloadbalancerroute
  resource_name: cloudloadbalancerroutes
  entity_name: CloudLoadBalancerRoute
  package: yeul
  group: pcn/infrastructure
  description: |-
    A CloudLoadBalancerRoute represents a Load Balancer Route as defined in an cloud
    provider
    (AWS/Azure/GCP etc).
    The Load Balancer Route is essentially a route from load balancer to the
     target group instances.
  aliases:
  - loadbalancerroute
  - loadbalancerroutes
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: parameters
    description: Load Balancer route related parameters.
    type: ref
    exposed: true
    subtype: cloudloadbalancerroutedata
    stored: true
    extensions:
      refMode: pointer
