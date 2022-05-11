# Model
model:
  rest_name: cloudloadbalancer
  resource_name: cloudloadbalancers
  entity_name: CloudLoadBalancer
  package: yeul
  group: pcn/infrastructure
  description: |-
    A CloudLoadBalancer represents a Load Balancer as defined in an cloud provider
    (AWS/Azure/GCP etc).
    The Load Balancer is essentially an L4,L7 or gateway load balancer with at least
    1 target group attached and it defines how a load balancing happens.
  aliases:
  - loadbalancer
  - loadbalancers
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
    description: Load Balancer related parameters.
    type: ref
    exposed: true
    subtype: cloudloadbalancerdata
    stored: true
    extensions:
      refMode: pointer
