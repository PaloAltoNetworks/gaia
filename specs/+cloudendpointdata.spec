# Model
model:
  rest_name: cloudendpointdata
  resource_name: cloudendpointdata
  entity_name: CloudEndpointData
  package: yeul
  group: pcn/infrastructure
  description: Parameters associated with a cloud endpoint.
  detached: true

# Attributes
attributes:
  v1:
  - name: SubnetAttachments
    description: The list of Subnets that this endpoint is directly attached to.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: VPCAttached
    description: |-
      Indicates that the endpoint is directly attached to the VPC. In this case the
      attachedInterfaces is empty. In general this is only valid for endpoint type
      Gateway and Peering Connection.
    type: boolean
    exposed: true
    stored: true

  - name: VPCAttachments
    description: The list of VPCs that this endpoint is directly attached to.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: associatedRouteTables
    description: |-
      List of route tables associated with this endpoint. Depending on cloud provider
      it can apply in some gateways.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: attachedEntities
    description: A list of entities that are associated to a given endpoint.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - eni-12344
    - subnet-abcd

  - name: attachedInterfaces
    description: |-
      A list of interfaces attached with the endpoint. In some cases endpoints can
      have more than one interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - eni-12344
    - eni-33333

  - name: availabilityZone
    description: |-
      The availabilityZone of the endpoint. Available for instances. This can be the
      placement in AWS or availability zone  or Azure.
    type: string
    exposed: true
    stored: true
    example_value:
    - us-east-2a
    - 2
    omit_empty: true

  - name: forwardingEnabled
    description: |-
      If the endpoint has multiple connections and forwarding can be enabled between
      them.
    type: boolean
    exposed: true
    stored: true

  - name: hasPublicIP
    description: If the endpoint has a public IP or is directly exposed.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true

  - name: imageID
    description: |-
      The imageID of running in the endpoint. Available for instances and
      potentially other 3rd parties. This can be the AMI ID in AWS or corresponding
      instance imageID in other clouds.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: instanceType
    description: |-
      The instanceType of the endpoint. Available for instances. This can be the
      instance type in AWS or virtual machine size in Azure.
    type: string
    exposed: true
    stored: true
    example_value:
    - t2.micro
    - Standard_F8s_v2
    omit_empty: true

  - name: productInfo
    description: Product related metadata associated with this endpoint.
    type: refList
    exposed: true
    subtype: cloudendpointdataproductinfo
    stored: true
    omit_empty: true
    extensions:
      refMode: pointer

  - name: publicIPAddresses
    description: if the endpoint has a public IP we store the IP address in this field.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: resourceStatus
    description: The status of the resource.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Active
    - Inactive
    default_value: Active

  - name: serviceName
    description: Identifies the name of the service for service endpoints.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: serviceType
    description: |-
      Identifies the service type that this endpoint represents (example Gateway Load
      Balancer).
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Interface
    - Gateway
    - GatewayLoadBalancer
    - ManagedService
    - NotApplicable
    default_value: NotApplicable

  - name: type
    description: Type of the endpoint.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Instance
    - LoadBalancer
    - PeeringConnection
    - Service
    - Gateway
    - TransitGateway
    - NATGateway
    - PublicIPAddress
    example_value: Instance
