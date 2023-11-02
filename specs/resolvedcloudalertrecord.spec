# Model
model:
  rest_name: resolvedcloudalertrecord
  resource_name: resolvedcloudalertrecords
  entity_name: ResolvedCloudAlertRecord
  package: vargid
  group: pcn/infrastructure
  description: Represents an Alert that was raised and later resolved.
  get:
    description: Retrieves the Resolved Alert Record with a given id.
    global_parameters:
    - $filtering
  extends:
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@zoned'

# Indexes
indexes:
- - namespace
  - Record.prismaCloudPolicyID

# Attributes
attributes:
  v1:
  - name: Record
    description: The record that was raised and later resolved.
    type: ref
    exposed: true
    subtype: cloudalertrecord
    stored: true
    extensions:
      refMode: pointer
