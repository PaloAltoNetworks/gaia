# Model
model:
  rest_name: cloudschedulednetworkquery
  resource_name: cloudschedulednetworkqueries
  entity_name: CloudScheduledNetworkQuery
  package: vargid
  group: pcn/infrastructure
  description: |-
    CloudSchedulednNetworkQuery represents a CloudNetworkQuery that will be
    scheduled periodically.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'

# Indexes
indexes:
- - alertruleid
- - lastexecutiontimestamp
- - alertruleid
  - policyid

# Attributes
attributes:
  v1:
  - name: cloudGraphResult
    description: The result of the cloud network query.
    type: ref
    exposed: true
    subtype: cloudgraph
    extensions:
      refMode: pointer

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
    description: Result of the last successfully run query.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: prismaCloudAlertRuleID
    description: Prisma Cloud Alert Rule ID.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: prismaCloudPolicyID
    description: Prisma Cloud Policy ID.
    type: string
    exposed: true
    subtype: string
    stored: true

  - name: tenantPrismaID
    description: Prisma ID of the tenant in which the Alert Rule is created.
    type: string
    exposed: true
    subtype: string
    stored: true
