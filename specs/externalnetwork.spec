# Model
model:
  rest_name: externalnetwork
  resource_name: externalnetworks
  entity_name: ExternalNetwork
  package: squall
  group: policy/networking
  description: |-
    An external network represents a random network or IP address that is not
    managed by Microsegmentation. External networks can be used in network policies
    to
    allow traffic from or to the declared network or IP, using the provided
    protocol and port (or range of ports). If you want to describe the internet
    (i.e., anywhere), use `0.0.0.0/0` as the address and `1-65000` for the ports.
    You must assign the external network one or more tags. These allow you to
    reference the external network from your network policies.
  aliases:
  - extnet
  - extnets
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@archivable'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@timeable'
  extensions:
    commentFlags:
    - +k8s:openapi-gen=true

# Indexes
indexes:
- - namespace
  - type
- - type

# Attributes
attributes:
  v1:
  - name: entries
    description: List of CIDRs or domain name.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $networksorhostnames

  - name: servicePorts
    description: List of protocol/ports `(tcp/80)` or `(udp/80:100)`.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - tcp/80
    - udp/80:100
    validations:
    - $serviceports

  - name: type
    description: The type of external network (default `Subnet`).
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ENI
    - RDSCluster
    - RDSInstance
    - SecurityGroup
    - Subnet
    default_value: Subnet
    filterable: true
