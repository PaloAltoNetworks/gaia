# Model
model:
  rest_name: awsasset
  resource_name: awsassets
  entity_name: AWSAsset
  package: pandemona
  group: pcn/infrastructure
  description: Represents a read-only AWS cloud resource such as a virtual machine.
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@awsresourceattrs'
