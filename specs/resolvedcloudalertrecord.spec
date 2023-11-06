# Model
model:
  rest_name: resolvedcloudalertrecord
  resource_name: resolvedcloudalertrecords
  entity_name: ResolvedCloudAlertRecord
  package: vargid
  group: pcn/infrastructure
  description: Represents an Alert that was raised and later resolved.
  extends:
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - record.prismacloudpolicyid

# Attributes
attributes:
  v1:
  - name: record
    description: The record that was raised and later resolved.
    type: ref
    exposed: true
    subtype: cloudalertrecord
    stored: true
    extensions:
      refMode: pointer
