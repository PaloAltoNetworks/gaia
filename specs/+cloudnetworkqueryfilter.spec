# Model
model:
  rest_name: cloudnetworkqueryfilter
  resource_name: cloudnetworkqueryfilters
  entity_name: CloudNetworkQueryFilter
  package: yeul
  group: pcn/infrastructure
  description: |-
    Captures the parameters allowed in a query filter for a net effective
    permissions request.
  detached: true

# Attributes
attributes:
  v1:
  - name: K8sClusterNames
    description: The cluster name of the target K8s resources. Applies only to resourceType
      K8s.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: K8sContainerImages
    description: |-
      A list of K8s images that resources can be filtered with. Applies only to
      resourceType K8s.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: K8sLabels
    description: |-
      A list of labels that apply to the queried resource. Applies only to
      resourceType K8s.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: K8sNamespaces
    description: The namespace of the target K8s resources. Applies only to resourceType
      K8s.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: K8sServiceNames
    description: The service name of the target K8s resources. Applies only to resourceType
      K8s.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: K8sServiceTypes
    description: Identifies a list of K8s Service types. Applies only to resourceType
      K8s.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: VPCIDs
    description: The VPC ID of the target resources.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: accountIDs
    description: |-
      The accounts that the search must apply to. These are the actually IDs of the
      account as provided by the cloud provider. One or more IDs can be included.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - account1
    omit_empty: true

  - name: cloudTypes
    description: The cloud types that the search must apply to.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - AWS
    omit_empty: true

  - name: imageIDs
    description: |-
      A list of imageIDs that endpoints can be filtered with. Applies only to
      resourceType Endpoint.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: objectIDs
    description: |-
      The exact object that the search applies. If ObjectIDs are defined, the rest of
      the fields are ignored. An object ID can refer to an instance, VPC endpoint, or
      network interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: paasTypes
    description: Identifies a list of Platform as a Service types.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: productInfoType
    description: Restricts the query on only endpoints with the given productInfoType.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: productInfoValue
    description: |-
      Restricts the query to only endpoints with the provided productInfoValue. Does
      not apply to other resource types.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: regions
    description: The region that the search must apply to.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - us-west-1
    omit_empty: true

  - name: resourceStatus
    description: The status of the resource.
    type: string
    exposed: true
    stored: true
    example_value: Active
    omit_empty: true

  - name: resourceType
    description: |-
      The type of endpoint resource. The resource type is a mandatory field and a
      query cannot span multiple resource types.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Instance
    - Interface
    - Service
    - ProcessingUnit
    - PaaS
    - K8s
    default_value: Instance

  - name: securityTags
    description: |-
      The list of security tags associated with the targets of the query. Security
      tags refer to security groups in AWS or network tags in GCP. So they can have
      different meaning depending on the target cloud.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: serviceNames
    description: |-
      Identifies a list of service names that should be taken into account. This is
      only valid with a resource type equal to Service.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: serviceOwners
    description: |-
      Identifies the owner of the service that the resource is attached to. Field is
      not valid if the resource type is not an interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: serviceTypes
    description: |-
      Identifies the type of service that the interface is attached to. Field is not
      valid if the resource type is not an
      interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: subnets
    description: |-
      The subnets where the resources must reside. A subnet parameter can only be
      provided for a network interface resource type.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true

  - name: tags
    description: |-
      A list of tags that select the same of endpoints for the query. These tags refer
      to the tags attached to the resources in the cloud provider definitions.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true
