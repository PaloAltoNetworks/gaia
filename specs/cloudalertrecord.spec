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
  - '@named'

# Indexes
indexes:
- - namespace
  - lastexecutiontimestamp
- - namespace
  - published
- - namespace
  - resourceid
  - prismacloudpolicyid
  - prismacloudalertruleid

# Attributes
attributes:
  v1:
  - name: accountID
    description: |-
      Account ID of the resource for which the Alert Record is
      raised.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: cloudType
    description: Cloud type of the entity.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - AWS
    - AZURE
    default_value: AWS

  - name: lastExecutionTimestamp
    description: Result of the last execution timestamp.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: metadataHash
    description: |-
      Sum of FNV-32a hashes of all the instances or interfaces grouped under the
      resource.
    type: integer
    exposed: true
    subtype: integer
    stored: true

  - name: prismaCloudAlertRuleIDs
    description: Prisma Cloud Alert Rules which generated the Alert Record.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: prismaCloudAlertRuleMatched
    description: Indicates if the alert matches an alert rule.
    type: boolean
    exposed: true
    stored: true
    default_value: false

  - name: prismaCloudAlertRuleScopeLastChangedOn
    description: Result of the last changed time to the prisma cloud alert rule.
    type: integer
    exposed: true
    stored: true
    default_value: 0

  - name: prismaCloudPolicyID
    description: Policy ID which generated the Alert Record.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: prismaCloudPolicyLastModifiedOn
    description: Result of the last modified time of the prisma cloud policy.
    type: integer
    exposed: true
    stored: true
    default_value: 0

  - name: published
    description: Indicates if the alert is published to PC.
    type: boolean
    exposed: true
    stored: true
    default_value: false

  - name: region
    description: |-
      Region of the resource for which the Alert Record is
      raised.
    type: string
    exposed: true
    subtype: string
    stored: true

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

  - name: resourceType
    description: Returns the type of the resource on which alert was raised.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Interface
    - Instance
    - VPC
    - Subnet
    autogenerated: true

  - name: rrn
    description: RRN identifier for the resource.
    type: string
    exposed: true
    stored: true

  - name: verdict
    description: Policy verdict of the resource.
    type: enum
    exposed: true
    subtype: string
    stored: true
    allowed_choices:
    - PASS
    - FAIL
    default_value: FAIL
    example_value:
    - FAIL
