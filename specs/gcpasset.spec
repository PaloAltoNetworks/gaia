# Model
model:
  rest_name: gcpasset
  resource_name: gcpassets
  entity_name: GCPAsset
  package: pandemona
  group: pcn/infrastructure
  description: Represents a read-only jkGCP cloud resource such as a virtual machine.
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@gcpresourceattrs'
