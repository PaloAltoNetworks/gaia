# Model
model:
  rest_name: azureresource
  resource_name: azureresources
  entity_name: AzureResource
  package: pandemona
  group: pcn/infrastructure
  description: |-
    Represents an Azure cloud resource such as virtualMachines and subnets. Only
    required attributes need to be set when creating the resource. Optional
    attributes will be ignored as they are set by the processor.
  private: true
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - resourceID
- - namespace
  - subscriptionID
  - kind

# Attributes
attributes:
  v1:
  - name: data
    description: The json-encoded data that represents the resource.
    type: external
    exposed: true
    subtype: '[]byte'
    stored: true
    required: true
    example_value: |-
      {
        "id": "/subscriptions/.../subnets/default",
        "name": "default",
        "routeTableId": "/subscriptions/.../Microsoft.Network/routeTables/my-table",
        ...
      }

  - name: kind
    description: The specific kind of the resource.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Pending
    - VirtualMachine
    - NetworkInterface
    - Subnet
    - IPConfiguration
    - VirtualNetwork
    - NetworkSecurityGroup
    - NATGateway
    - PublicIPAddress
    - PublicIPPrefix
    - VirtualMachineScaleSet
    - VirtualMachineScaleSetVM
    - LoadBalancer
    - BackendAddressPool
    - OutboundRule
    - FrontendIPConfiguration
    - DatabaseAccount
    - FlexibleServer
    - Server
    autogenerated: true
    default_value: Pending
    example_value: VirtualMachine

  - name: name
    description: The name of this resource.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    example_value: vm-name

  - name: provider
    description: The major type of the resource.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Pending
    - MicrosoftCompute
    - MicrosoftNetwork
    - MicrosoftDocumentDB
    - MicrosoftDBforMySQL
    autogenerated: true
    default_value: Pending
    example_value: MicrosoftCompute

  - name: resourceGroup
    description: The name of the logical subcontainer of cloud resources.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    example_value: my-deployment

  - name: resourceID
    description: The identifier of the resource as presented by Azure, which is a
      path.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    example_value: /subscriptions/a0a00a0a-0aaa-000a-a0a0-00a00aa00000/resourceGroups/my-deployment/providers/Microsoft.Compute/virtualMachines/vm-name

  - name: resourceTags
    description: Field 'tags' inside the Azure resource.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true
    read_only: true
    autogenerated: true

  - name: subscriptionID
    description: The logical ID of the container which contains the cloud resources.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    example_value: a0a00a0a-0aaa-000a-a0a0-00a00aa00000

  - name: tags
    description: |-
      Contextual values that can be used to narrow searching of resources if the
      resourceID is not known. For instance, it could be used to store a resource's
      location or public IP addresses to support cross-cloud analysis.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
