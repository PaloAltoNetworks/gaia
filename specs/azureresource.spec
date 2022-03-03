# Model
model:
  rest_name: azureresource
  resource_name: azureresource
  entity_name: AzureResource
  package: pandemona
  group: pcn/infrastructure
  description: Represents an Azure cloud resource
  private: true
  extends:
  - '@migratable'
  - '@namespaced'
  - '@zoned'

indexes:
- - namespace
  - resourceID
- - namespace
  - subscriptionID
  - kind

# Attributes
attributes:
  v1:
  - name: resourceID
    description: The identifier of the resource as presented by Azure, which is a path.
    exposed: true
    stored: true
    type: string
    example_value: /subscriptions/a0a00a0a-0aaa-000a-a0a0-00a00aa00000/resourceGroups/my-deployment/providers/Microsoft.Compute/virtualMachines/vm-name

  - name: name
    description: The name of this resource.
    exposed: true
    stored: true
    type: string
    example_value: vm-name


  - name: subscriptionID
    description: The logical ID of the container which contains the cloud resources.
    exposed: true
    stored: true
    type: string
    example_value: a0a00a0a-0aaa-000a-a0a0-00a00aa00000

  - name: resourceGroup
    description: The name of the logical subcontainer of cloud resources.
    exposed: true
    stored: true
    type: string
    example_value: my-deployment

  - name: provider
    description: The major type of the resource.
    exposed: true
    stored: true
    type: string
    example_value: Microsoft.Compute

  - name: kind
    description: The specific kind of the resource.
    exposed: true
    stored: true
    type: string
    example_value: virtualMachines

  - name: data
    description: The json-encoded data that represents the resource.
    exposed: true
    stored: true
    type: string
    example_value: |-
      {
        "id": "/subscriptions/.../subnets/default",
        "name": "default",
        "routeTableId": "/subscriptions/.../Microsoft.Network/routeTables/my-table",
        ...
      }

  - name: tags
    description: |-
      A contextual key-value field that can be used to narrow searching of resources if the
      resourceID is not known. For instance, it could be used to store resource location or
      public IP addresses to support cross-cloud analysis.
    exposed: true
    stored: true
    type: external
    subtype: map[string][]string
