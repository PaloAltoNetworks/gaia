# Model
model:
  rest_name: azureasset
  resource_name: azureassets
  entity_name: AzureAsset
  package: pandemona
  group: pcn/infrastructure
  description: Represents a read-only Azure cloud resource such as virtualMachines
    and subnets.
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@azureresourceattrs'
