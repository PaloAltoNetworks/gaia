# Model
model:
  rest_name: gcpresource
  resource_name: gcpresources
  entity_name: GCPResource
  package: pandemona
  group: pcn/infrastructure
  description: Represents a GCP cloud resource such as a virtual machine.
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - selflink
- - namespace
  - numericID

# Attributes
attributes:
  v1:
  - name: data
    description: The JSON-encoded data that represents the resource.
    type: external
    exposed: true
    subtype: '[]byte'
    stored: true
    required: true
    example_value: |-
      {
        "id": "0000000000000000000",
        "kind": "compute#instance"
        ...
      }

  - name: kind
    description: The specific kind of the resource.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - ComputeInstance
    - ComputeSubnetwork
    - ComputeNetwork
    example_value: ComputeInstance

  - name: numericID
    description: A numeric resource ID that will mainly be used in RQL queries.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: "0000000000000000000"

  - name: prismaCloudRRN
    description: The identifier used by Prisma Cloud to locate the same resource.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: selflink
    description: The identifier of the resource as presented by GCP, which is a URL.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://www.googleapis.com/compute/v1/projects/foobar/global/networks/abc

  - name: tags
    description: |-
      Contextual values that can be used to narrow searching of resources if the
      resourceID is not known. For instance, it could be used to store a resource's
      location or public IP addresses to support cross-cloud analysis.
    type: list
    exposed: true
    subtype: string
    stored: true
