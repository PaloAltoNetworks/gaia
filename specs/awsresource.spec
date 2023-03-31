# Model
model:
  rest_name: awsresource
  resource_name: awsresources
  entity_name: AWSResource
  package: pandemona
  group: pcn/infrastructure
  description: Represents a AWS cloud resource such as a virtual machine.
  private: true
  extends:
  - '@identifiable-stored'
  - '@migratable'
  - '@namespaced'
  - '@zoned'
  - '@timeable'
  - '@awsresourceattrs'

# Indexes
indexes:
- - namespace
  - arn
- - namespace
  - resourceID
- - namespace
  - kind
