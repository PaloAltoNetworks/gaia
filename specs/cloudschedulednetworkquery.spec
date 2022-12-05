# Model
model:
  rest_name: cloudschedulednetworkquery
  resource_name: cloudschedulednetworkqueries
  entity_name: CloudScheduledNetworkQuery
  package: vargid
  group: pcn/infrastructure
  description: |-
    CloudScheduledNetworkQuery represents a CloudNetworkQuery that will be
    scheduled periodically.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@named'

# Indexes
indexes:
- - namespace
  - prismacloudpolicyid
- - disabled
  - lastexecutiontimestamp

# Attributes
attributes:
  v1:
  - name: cloudGraphResultID
    description: The cloud graph result ID which is stored in MongoDB GridFS.
    type: string
    exposed: true
    subtype: string

  - name: cloudNetworkQuery
    description: The cloud network query that should be used.
    type: ref
    exposed: true
    subtype: cloudnetworkquery
    stored: true
    extensions:
      refMode: pointer

  - name: disabled
    description: Represents whether the associated policy was disabled.
    type: boolean
    exposed: true
    stored: true

  - name: lastExecutionTimestamp
    description: Timestamp of the last time the query was scheduled.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: neoCNA
    description: If set to true, neocna will be used regardless of tenant onboarding.
    type: boolean
    exposed: true

  - name: prismaCloudPolicyID
    description: Prisma Cloud Policy ID.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: successfulExecutionTimestamp
    description: |-
      Timestamp of the last time the query was successfully executed and results were
      obtained.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: successfulExecutionTimestampMap
    description: Mapping of last successful execution timestamp for every account.
    type: external
    exposed: true
    subtype: map[string]time
    stored: true

  - name: tenantPrismaID
    description: Prisma ID of the tenant in which the Alert Rule is created.
    type: string
    exposed: true
    subtype: string
    stored: true
