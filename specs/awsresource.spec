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
  - accountID
- - namespace
  - accountID
- - prismaID
- - prismaID
  - kind
  - arn
  - denormedFields
- - prismaID
  - kind
  - resourceID
  - denormedFields
- - prismaID
  - kind
  - prismaRegion
  - denormedFields
- - prismaID
  - kind
  - name
  - denormedFields
- - prismaID
  - kind
  - denormedFields
- - prismaID
  - kind
  - accountID
  - denormedFields
- - prismaID
  - arn
- - prismaID
  - resourceID
- - prismaID
  - accountID
- - prismaID
  - name
- - prismaID
  - denormedFields
