# Model
model:
  rest_name: gcpresource
  resource_name: gcpresources
  entity_name: GCPResource
  package: pandemona
  group: pcn/infrastructure
  description: Represents a GCP cloud resource such as a virtual machine.
  private: true
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@gcpresourceattrs'

# Indexes
indexes:
- - namespace
  - selflink
- - namespace
  - numericID
- - namespace
  - kind
