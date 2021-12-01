# Model
model:
  rest_name: cloudnode
  resource_name: cloudnodes
  entity_name: CloudNode
  package: yeul
  group: pcn/infrastructure
  description: Manages the list of cloud nodes available in a cloud deployment.
  get:
    description: Retrieves the cloud node with the given ID.
  delete:
    description: Deletes the cloud node with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Indexes
indexes:
- - namespace
  - type
  - subtype
- - namespace
  - relatedObjectID
- - namespace
  - securitytags
  - type
  - vpcid
- - namespace
  - vpcid
  - type
- - namespace
  - vpcid
  - parameters
- - namespace
  - type
  - parameters
- - key
- - createTime

# Attributes
attributes:
  v1:
  - name: attachments
    description: The list of attachments for this node.
    type: list
    exposed: true
    subtype: string
    stored: true
    getter: true
    setter: true
    omit_empty: true

  - name: key
    description: Internal unique key for a resource to guarantee no overlaps at write.
    type: string
    stored: true

  - name: parameters
    description: The cloud attributes of the object.
    type: external
    exposed: true
    subtype: map[string]interface{}
    stored: true
    getter: true
    setter: true
    omit_empty: true

  - name: relatedObjectID
    description: A reference to a related object.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: securityTags
    description: List of security tags associated with the node.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: subType
    description: The sub-type of the object as found in the parameters. Used for indexing.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: type
    description: Type of the endpoint.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Endpoint
    - Subnet
    - VPC
    - Interface
    - RouteTable
    - NetworkRuleSet
    - LoadBalancer
    - LoadBalancerRoute
    example_value: Endpoint
