# Model
model:
  rest_name: cloudalertrecord
  resource_name: cloudalertrecords
  entity_name: CloudAlertRecord
  package: vargid
  group: pcn/infrastructure
  description: |-
    Represents an Alert that is raised against a resource based on a RQL (policy)
    associated to an alert rule.
  get:
    description: Retrieves the Alert Record with a given id.
    global_parameters:
    - $filtering
  update:
    description: Updates the Alert Record with a given id.
  delete:
    description: Deletes the Alert Record with a given id.
  extends:
  - '@base'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@zoned'

# Indexes
indexes:
- - lastexecutiontimestamp
  - namespace
- - published
  - namespace

# Attributes
attributes:
  v1:
  - name: lastExecutionTimestamp
    description: Result of the last execution timestamp.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: metadataHash
    description: |-
      Sum of FNV-32a hashes of all the instances/interfaces grouped under the
      resource.
    type: integer
    exposed: true
    subtype: integer
    stored: true

  - name: prismaCloudAlertRuleID
    description: Prisma Cloud Alert Rule which generated the Alert Record.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: prismaCloudPolicyID
    description: Policy ID which generated the Alert Record.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: published
    description: Indicates if the alert is published to PC.
    type: boolean
    exposed: true
    stored: true
    default_value: false

  - name: resourceCount
    description: Number of interfaces/instances for which the alert is raised.
    type: integer
    exposed: true
    subtype: integer
    stored: true

  - name: resourceID
    description: Resource ID of the resource for which the Alert Record is raised.
    type: string
    exposed: true
    subtype: string
    stored: true
