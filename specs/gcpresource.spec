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
  - nativeID
- - namespace
  - accountid

# Attributes
attributes:
  v1:
  - name: RRN
    description: Prisma Cloud RRN.
    type: string
    exposed: true
    omit_empty: true

  - name: accountID
    exposed_name: accountId
    description: Cloud account ID associated with the resource.
    type: string
    exposed: true
    stored: true
    example_value: cna-dev-123456
    omit_empty: true

  - name: data
    description: The json encoded data that represents the resource.
    type: external
    exposed: true
    subtype: '[]byte'
    stored: true
    required: true
    example_value: |-
      {
        "id": "12345",
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
    - Instance
    - Subnet
    - VirtualNetwork
    example_value: Instance

  - name: nativeID
    description: The ID of the object.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 1234567

  - name: resourceID
    description: The identifier of the resource as presented by Azure, which is a
      path.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /subscriptions/a0a00a0a-0aaa-000a-a0a0-00a00aa00000/resourceGroups/my-deployment/providers/Microsoft.Compute/virtualMachines/vm-name

  - name: selflink
    description: The link to the object.
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
