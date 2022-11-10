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
  - '@timeable'
  - '@azureresourceattrs'

# Indexes
indexes:
- - namespace
  - resourceID
- - namespace
  - subscriptionID
  - kind
